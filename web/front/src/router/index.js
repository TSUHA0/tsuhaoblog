import Vue from 'vue'
import VueRouter from 'vue-router'

const ArticleList = () =>
  import(/* webpackChunkName: "group-index" */ '../components/Artlist.vue')
const Detail = () =>
  import(/* webpackChunkName: "group-detail" */ '../components/Details.vue')
const Category = () =>
  import(/* webpackChunkName: "group-category" */ '../components/CateArt.vue')
const Search = () =>
  import(/* webpackChunkName: "group-search" */ '../components/Search.vue')

Vue.use(VueRouter)

// 路由重复点击捕获错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject)
    return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [
  { path: '/', component: ArticleList, meta: { title: '欢迎来到TsuhaoBlog' } },
  {
    path: '/article/detail/:id',
    component: Detail,
    meta: { title: '文章详情' },
    props: true
  },
  {
    path: '/category/:cid',
    component: Category,
    meta: { title: '分类信息' },
    props: true
  },
  {
    path: '/search/:title',
    component: Search,
    meta: { title: '搜索结果' },
    props: true
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

export default router
