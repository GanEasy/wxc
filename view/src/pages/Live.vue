<template>
  <div>
  <div class="page__hd" >
      <h1 class="page__title" style="font-size: 1.6em;">“迎春送福•领墨宝献爱心”活动捐款清单</h1>
  </div>
  <img src="../assets/cover.jpg" style="max-width: 100%;">
  
  <div class="page__hd" style="">
      <p class="page__desc">捐款人次：{{list.length}}</p>
      <p class="page__desc">捐款总金额：{{total}} 元</p>
      <p class="page__desc">排名不分先后，名单实时更新！</p>
  </div>
    <div class="weui-cells">
        <div class="weui-cell" v-for="item in list" :key="item.id" >
            <div class="weui-cell__bd">
                <p>{{item.name}}</p>
            </div>
            <div class="weui-cell__ft">{{item.amount}} 元</div>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      token:"",
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      chatContent: '', // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false, // True if email and username have been filled in
      list: [
      ],
      total:0,
    }
  },
  
   watch: {
    list: function (newQuestion) {
        var _t =0;
        for(var i=0;i<this.list.length;i++) {
            var _tt =  parseFloat(this.list[i].amount)
            if(_tt){
                _t += _tt
            }
        }
        this.total = Math.floor(_t * 100) / 100 
    }
  },
  
   
    mounted() {
      this.token = this.$route.query.token
        console.log(this.token)

    },
    created: function() {
        var self = this;
        self.axios.get("http://119.23.237.215:3344/list").then((response) => {
        // console.log(response.data)
            for(var k in response.data) {
                var item = response.data[k]
                var amount = Math.floor(item.amount * 100) / 100 
                self.list.push({id:item.id,name:item.name,amount:amount})
            }
        })


        // self.ws = new WebSocket('ws://localhost:3344/ws');
        self.ws = new WebSocket('ws://119.23.237.215:3344/ws');
        self.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
             if(msg.type == "donate"){
                var amount = Math.floor(msg.amount * 100) / 100 
                self.list.push({id:msg.id,name:msg.name,amount:amount})
            }
        });
      },

    
    methods: {
    }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page__hd {
    padding: 20px;
}
.page__desc{
    font-size: 1.0em;
    color:black;
    font-weight: bold;
}
</style>
