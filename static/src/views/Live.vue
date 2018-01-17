//index.vue
<template>

<div>


<div class="weui-panel_access">
  <div class="weui-cells__title">带说明的列表项</div>
    <div class="weui-cell">
        <div class="weui-cell__bd">
            <p>标题文字</p>
        </div>
        <div class="weui-cell__ft">说明文字</div>
    </div>
  </div>
</div>


</template>
<style>
.article-title{
  color: #000;
}
</style>
<script>

import ArticleList from '@/components/ArticleList';
import InfiniteLoading from 'vue-infinite-loading'
import api from '../api';

export default {
  
  name: 'indexP',
  
  
    components: {
      InfiniteLoading,
      ArticleList
    },
    data () {
      return {
        list: [
          {name:"yizenghui",amount:100},
          {name:"asxsa",amount:100},
          {name:"asdasdf",amount:100},
          {name:"asd",amount:100},
        ],
      }
    },
   
}
</script>



