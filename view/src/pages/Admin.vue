<template>
  <div>

<div v-if="showMsg" id="toast">
    <div class="weui-mask_transparent"></div>
    <div class="weui-toast">
        <i class="weui-icon-success-no-circle weui-icon_toast"></i>
        <p class="weui-toast__content">{{msg}}</p>
    </div>
</div>

  <div class="page__hd">
      <h1 class="page__title">捐赠人名单</h1>
      <p class="page__desc">排名不分先后，本名单实时更新！</p>
      <p class="page__desc">捐款人次：{{list.length}}</p>
      <p class="page__desc">捐款总金额：{{total}}</p>
  </div>



  <div v-if="token">
  

        <div class="weui-tab">
            <div class="weui-navbar">
                <div class="weui-navbar__item" v-on:click="showpanel =1" :class="showpanel==1? 'weui-bar__item_on':''">
                    新增捐款
                </div>
                <div class="weui-navbar__item" v-on:click="showpanel =2"  :class="showpanel==2? 'weui-bar__item_on':''">
                    捐款记录
                </div>
                <div class="weui-navbar__item"  v-on:click="showpanel =3" :class="showpanel==3? 'weui-bar__item_on':''">
                    推送消息
                </div>
            </div>
            <div class="weui-tab__panel">

<!-- 发布新捐款信息s -->
                <div v-if="showpanel==1">
                    <div class="weui-cells__title">发布新捐款信息</div>
                        <div class="weui-cells">
                            <div class="weui-cell">
                                <div class="weui-cell weui-cell_select weui-cell_select-before">
                                    <div class="weui-cell__hd">
                                        <input class="weui-input" placeholder="捐款人" v-model="inputName">
                                    </div>
                                    <div class="weui-cell__bd">
                                        <input class="weui-input" type="number" pattern="[0-9]*" placeholder="金额" v-model="inputAmount">
                                    </div>
                                </div>
                            </div>
                        </div>

                    <div class="weui-btn-area">
                        <a class="weui-btn weui-btn_primary" href="javascript:"  v-on:click="pushDonate">提交捐款消息</a>
                    </div>
                </div><!-- 发布新捐款信息e -->


                <!-- 捐款人名单s -->
                <div v-if="showpanel==2">

                    <div class="weui-cells" v-if="list.length>0">
                        <div class="weui-cell" v-for="item in list" :key="item.id"  >
                            <div class="weui-cell__bd">
                                <p>{{item.name}}</p>
                            </div>
                            <div class="weui-cell__ft">{{item.amount}} 元</div>
                        </div>
                    </div>

                    <div class="weui-cells" v-if="list.length>0">
                        <div class="weui-cell" >
                            <div class="weui-cell__bd">
                                <p>还没有捐款信息，快去新增一个吧</p>
                            </div>
                        </div>
                    </div>

                </div><!-- 捐款人名单e -->

                
                <!--消息推送s-->
                <div v-if="showpanel==3">
                    <div class="weui-cells__title">消息推送</div>

                    <div class="weui-cells weui-cells_form">
                        <div class="weui-cell">
                            <div class="weui-cell__bd">
                                <textarea class="weui-textarea" placeholder="请输入文本" rows="3"></textarea>
                            </div>
                        </div>
                    </div>
                    <div class="weui-btn-area">
                        <a class="weui-btn weui-btn_primary" href="javascript:"  v-on:click="pushMsg">提交该消息</a>
                    </div>
                </div><!--消息推送e-->

            </div><!--展页内容end-->

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
      showpanel:1,
      inputToken:"",
      inputName:"",
      inputAmount:null,   
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
      total:0,
      showMsg:false,
      msg:"",
    }
  },
  
   watch: {
    list: function (newQuestion) {
        var _t =0;
        for(var i=0;i<this.list.length;i++) {
            var _tt =  parseFloat(this.list[i].amount)
            if(_tt){
                _t += Math.floor(_tt * 100) / 100 
            }
        }
        this.total = Math.floor(_t * 100) / 100 
    }
  },
   
    mounted() {
      this.token = this.$route.query.token
        // console.log(this.token)
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

        self.ws = new WebSocket('ws://119.23.237.215:3344/ws');
        self.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            // console.log("ret:",msg)
            if(msg.type == "donate"){
                var amount = Math.floor(msg.amount * 100) / 100 
                self.list.push({id:msg.id,name:msg.name,amount:amount})
            }
        });

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
            if(self.inputName==""){
                return false
            }else if(self.inputAmount == ""){
                return false
            }
            // console.log(self.inputAmount)
            var url = "http://119.23.237.215:3344/api?type=donate&name="+self.inputName+"&amount="+parseFloat(self.inputAmount)+"&token="+self.token
            self.axios.get(url).then((response) => {
                self.inputName = ""
                self.inputAmount = ""
                console.log(response.data)
                self.msg = "已增加捐款记录"
                self.showMsg = true
                setTimeout(function(){self.showMsg = false},1500);
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
