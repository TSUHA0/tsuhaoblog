<template>
  <a-layout-header class="headerbtn">
      <a-button type="danger" @click="logout">退出</a-button>
  </a-layout-header>
</template>
<script>
export default {
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
<style scoped>
.headerbtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
