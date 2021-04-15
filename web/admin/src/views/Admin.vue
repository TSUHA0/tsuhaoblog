<template lang="">
  <a-layout class="container">
    <Nav></Nav>
    <a-layout>
          <Header></Header>
      <a-layout-content>
          <router-view></router-view>
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
  position: fixed;
  height: 100%;
  width: 100%;
}

</style>
