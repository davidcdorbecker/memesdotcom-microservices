import { createRouter, createWebHistory } from 'vue-router'
import { isUserLogged } from '@/functions/auth'

const routes = [
  {
    name: 'login',
    path: '/login',
    beforeEnter (to, from, next) {
      if (isUserLogged.value) next({ name: 'feed' })
      else next()
    },
    component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue')
  },
  {
    name: 'feed',
    path: '/',
    meta: { requiresAuth: true },
    component: () => import(/* webpackChunkName: "feed" */ '../views/Feed.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  if (!to.matched.some(record => record.meta.requiresAuth)) next()
  else if (isUserLogged.value) next()
  else next({ name: 'login', query: { redirect: to.fullPath } })
})

export default router
