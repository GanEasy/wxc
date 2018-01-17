<template>
  <div>
  <div class="page__hd">
      <h1 class="page__title">List</h1>
      <p class="page__desc">列表</p>
  </div>
    <div class="weui-cells">
        <div class="weui-cell" v-for="item in list" :key="item.name" >
            <div class="weui-cell__bd">
                <p>{{item.name}}</p>
            </div>
            <div class="weui-cell__ft">{{item.amount}}</div>
        </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false, // True if email and username have been filled in
      list: [
        {name:"易增辉",amount:100},
        {name:"甘霞",amount:100},
        {name:"小灰灰",amount:100},
        {name:"富豪",amount:1000},
      ],
    }
  },
  
    created: function() {
        var self = this;
        self.ws = new WebSocket('ws://localhost:3344/ws');
        self.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            console.log("ret:",e.data,msg)
        });
        self.email = "121258121@qq.com"
        self.username = "易增辉"
        self.newMsg = "say hi"
        setInterval(function() {

          self.ws.send(
              JSON.stringify({
                  email: self.email,
                  username: self.username,
                  message: self.newMsg // Strip out html
              }
          ));
          // ws.send('Hello, Server!');
        }, 10000);
        // this.newMsg = ''; // Reset newMsg
      },

    
    methods: {
    }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page__hd {
    padding: 40px;
}
</style>
