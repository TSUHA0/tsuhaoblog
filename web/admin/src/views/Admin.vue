<template lang="">
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
          <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
          <!-- 设置key可以帮助刷新页面，使缓存信息及时更新，
          如果不设置，就会出现编写文章界面更新了但是仍然不跳转的现象
          比如/admin/addart/2 直接修改为/admin/addart/就不会刷新 -->
          <router-view :key="$route.path"></router-view>
      </a-layout-content>
        <Footer></Footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Footer from '../components/admin/Footer'
import Nav from '../components/admin/Nav'
import Header from '../components/admin/Header'
export default {
  components: { Nav, Footer,Header }
}
</script>
<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
