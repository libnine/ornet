import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import History from '@/components/History'
import Search from '@/components/Search'
import Calendar from '@/components/Calendar'
import Case from '@/components/Case'
import Create from '@/components/Create'
import Add from '@/components/Add'
import User from '@/components/User'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/search',
      name: 'Search',
      component: Search
    },
    {
      path: '/calendar',
      name: 'Calendar',
      component: Calendar
    },
    {
      path: '/calendar/:case_date',
      name: 'Case',
      component: Case,
      props: true
    },
    {
      path: '/create',
      name: 'Create',
      component: Create,
    },
    {
      path: '/history',
      name: 'History',
      component: History
    },
    {
      path: '/add_user',
      name: 'Add',
      component: Add,
      props: true
    },
    {
      path: '/:user',
      name: 'User',
      component: User,
      props: true
    }
  ]
})
