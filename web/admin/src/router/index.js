import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'



// 路由重复点击捕获错误
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

const Index = () => import(/* webpackChunkName: "Index" */ '../components/admin/Index.vue')
const AddArt = () => import(/* webpackChunkName: "AddArt" */ '../components/article/AddArt.vue')
const ArtList = () => import(/* webpackChunkName: "ArtList" */ '../components/article/ArtList.vue')
const CateList = () => import(/* webpackChunkName: "CateList" */ '../components/category/CateList.vue')
const UserList = () => import(/* webpackChunkName: "UserList" */ '../components/user/UserList.vue')

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '请登录'
    },
    component: Login
  },
  {
    path: '/',
    name: 'admin',
    meta: {
      title: 'TsuhaoBlog 后台管理页面'
    },
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index,
        meta: {
          title: 'TsuhaoBlog 后台管理页面'
        }
      },
      {
        path: 'addart',
        component: AddArt,
        meta: {
          title: '新增文章'
        }
      },
      {
        path: 'addart/:id',
        component: AddArt,
        meta: {
          title: '编辑文章'
        },
        props: true
      },
      {
        path: 'artlist',
        component: ArtList,
        meta: {
          title: '文章列表'
        }
      },
      {
        path: 'catelist',
        component: CateList,
        meta: {
          title: '分类列表'
        }
      },
      {
        path: 'userlist',
        component: UserList,
        meta: {
          title: '用户列表'
        }
      },
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  // to and from are both route objects. must call `next`.
  const token = window.sessionStorage.getItem('token')
  if (to.path == '/') {
    next('/login')
  }
  if (to.path == '/login') {
    return next()
  }
  if (!token && to.path == '/admin') {
    next('/login')
  } else {
    next()
  }
})

export default router
