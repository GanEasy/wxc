<template>
  <div>
  <div class="page__hd">
      <h1 class="page__title">捐赠人名单</h1>
      <p class="page__desc">排名不分先后，本名单实时更新！</p>
  </div>

  <div v-if="token">
    <div class="weui-cells">
        <div class="weui-cell" v-for="item in list" :key="item.id"  >
            <div class="weui-cell__bd">
                <p>{{item.name}}</p>
            </div>
            <div class="weui-cell__ft">{{item.amount}} 元</div>
        </div>
    </div>
  
    <div class="weui-cells__title">消息推送</div>

    <div class="weui-cells weui-cells_form">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <textarea class="weui-textarea" placeholder="请输入文本" rows="3"></textarea>
            </div>
        </div>
    </div>
    <div class="weui-btn-area">
        <a class="weui-btn weui-btn_primary" href="javascript:"  v-on:click="pushMsg">推送该消息</a>
    </div>


    <div class="weui-cells__title">发布新捐款信息</div>
        <div class="weui-cells">
            <div class="weui-cell weui-cell_select weui-cell_select-before">
                <div class="weui-cell__hd" style="padding-left:10px">
                    <input class="weui-input" placeholder="请输入捐款人名字" v-model="inputName">
                </div>
                <div class="weui-cell__bd">
                    <input class="weui-input" type="number" pattern="[0-9]*" placeholder="请输入捐款金额" v-model="inputAmount">
                </div>
            </div>
        </div>

        <div class="weui-btn-area">
            <a class="weui-btn weui-btn_primary" href="javascript:"  v-on:click="pushDonate">推送该捐款消息</a>
        </div>
  </div>

  <div v-if="!token">
    <div class="weui-cells__title">安全验证</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input class="weui-input" type="text" placeholder="请输入token" v-model="inputToken">
            </div>
        </div>
    </div>
    <div class="weui-btn-area">
            <a class="weui-btn weui-btn_primary" href="javascript:"  v-on:click="saveToken">确定</a>
        </div>
  </div>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      inputToken:"",
      inputName:"",
      inputAmount:"",
      msg:"",
      token:"",
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false, // True if email and username have been filled in
      list: [
      ],
    }
  },
  
  
   
    mounted() {
      this.token = this.$route.query.token
        console.log(this.token)

    },
    created: function() {
        var self = this;
        self.axios.get("http://localhost:3344/list").then((response) => {
        console.log(response.data)

        for(var k in response.data) {
            var item = response.data[k]
        //    console.log()
           self.list.push({id:item.id,name:item.name,amount:item.amount})

        }
        //  for(var t=0;t<response.data.length;t++){
        //         site.articles.push(response.data[t])
        //     }
        })





        self.ws = new WebSocket('ws://localhost:3344/ws');
        self.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            console.log("ret:",msg)
            self.list.push({id:msg.id,name:msg.name,amount:msg.amount})
        });
        // self.email = "121258121@qq.com"
        // self.username = "易增辉"
        // self.newMsg = "say hi"
        setInterval(function() {
        //   self.ws.send( 
        //       JSON.stringify({
        //           token: self.token,
        //           email: self.email,
        //           type:"text",
        //           amount:"112",
        //           name: self.username,
        //           message: self.newMsg // Strip out html
        //       }
        //   ));
          // ws.send('Hello, Server!');
        }, 20000);
        // this.newMsg = ''; // Reset newMsg
      },

    
    methods: {
        saveToken:function(){
            // todo check token
            this.token = this.inputToken
        },
        pushMsg:function(){
            // todo check token
            this.token = this.inputToken
        },
        pushDonate:function(){
             var self = this;
             var url = "http://localhost:3344/api?name="+self.inputName+"&amount="+self.inputAmount+"&token="+self.token
            self.axios.get(url).then((response) => {
                console.log(response.data)
            })
        }
    }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page__hd {
    padding: 40px;
}
</style>
