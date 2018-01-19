package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	msgs = map[int]*Message{}
	seq  = 1
)

// Message 消息结构体
type Message struct {
	ID     int     `json:"id"`
	Token  string  `json:"token"`
	Type   string  `json:"type"` //donation
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Author string  `json:"author"`
	Intro  string  `json:"intro"`
	File   string  `json:"file"`
}

//Donate  捐款记录
type Donate struct {
	gorm.Model
	Name   string
	Amount float64
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "mmbwgwxc.db")
	if err != nil {
		panic("连接数据库失败")
	}
	// 自动迁移模式
	db.AutoMigrate(&Donate{})
	defer db.Close()
	initList()
}

func initList() {

	var donates []Donate
	db.Find(&donates)

	if len(donates) > 0 {
		for _, v := range donates {
			msg := Message{
				ID: seq,
			}
			seq++

			msg.Name = v.Name
			msg.Author = ""
			msg.Token = ""
			msg.Amount = v.Amount
			msgs[msg.ID] = &msg
		}
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", conn)
	e.GET("/api", api)
	e.GET("/list", list)
	go handleMessages()
	e.Logger.Fatal(e.Start(":3344"))
}

func api(c echo.Context) error {

	msg := Message{
		ID: seq,
	}
	
	token := c.QueryParam("token")
	msgtype := c.QueryParam("type")
	if msgtype == "donate" && token="112233" {
		seq++
		msg.ID = seq

		msg.Name = c.QueryParam("name")
		msg.Author = ""
		msg.Token = ""
		msg.Type = msgtype
		msg.Amount, _ = strconv.ParseFloat(c.QueryParam("amount"), 32)
		msgs[msg.ID] = &msg
		var donate = Donate{Name: msg.Name, Amount: msg.Amount}

		db, err := gorm.Open("sqlite3", "mmbwgwxc.db")
		if err != nil {

		}
		// 自动迁移模式
		defer db.Close()
		db.Create(&donate)
		fmt.Println(donate)
		broadcast <- msg
	}
	return c.JSON(http.StatusCreated, msg)
}

func list(c echo.Context) error {
	return c.JSON(http.StatusCreated, msgs)
}
func conn(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		seq++
		msg.ID = seq
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("handle error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
	return nil
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
