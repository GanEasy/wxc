import Vue from 'vue'
import Router from 'vue-router'
import LivePage from '@/pages/Live'
import AdminPage from '@/pages/Admin'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'live',
      component: LivePage
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminPage
    },
  ]
})
