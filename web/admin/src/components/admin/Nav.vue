<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" key="index">
      <a-menu-item><a-icon type="dashboard"></a-icon><span>仪表盘</span></a-menu-item>

      <a-sub-menu>
        <span slot="title"><a-icon type="file"></a-icon><span>文章管理</span></span>
        <a-menu-item key="addart"><a-icon type="edit"></a-icon><span>撰写文章</span></a-menu-item>
        <a-menu-item key="artlist"><a-icon type="snippets"></a-icon> <span>文章列表</span></a-menu-item>
      </a-sub-menu>
      
      <a-menu-item key="catelist"><a-icon type="tag" /><span>分类管理</span></a-menu-item>

      <a-menu-item key="userlist"><a-icon type="user" /><span>用戶管理</span></a-menu-item>

    </a-menu>
  </a-layout-sider>
</template>
<script>
export default {
  data() {
    return { collapsed: false }
  }
}
</script>
<style scoped>
.log {
  margin: 16px;
  height: 32px;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
}
</style>
