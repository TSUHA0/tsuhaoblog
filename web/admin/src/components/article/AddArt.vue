<template>
  <div>
    <h3>{{ Object.keys(this.$route.params).indexOf('id') >= 0 ? '编辑文章' : '新建文章' }}</h3>
    <a-card>
      <a-form-model>
          <a-form-model-item label="文章标题"></a-form-model-item>
          <a-input></a-input>
        <a-form-model-item label="文章分类"></a-form-model-item>
        <a-col><a-input></a-input></a-col>
        <a-form-model-item label="文章描述"></a-form-model-item>
        <a-input></a-input>
        <a-form-model-item label="文章缩略图"></a-form-model-item>
        <a-input></a-input>
        <a-form-model-item label="文章内容"></a-form-model-item>
        <a-input></a-input>
        <a-form-model-item class="btn">
          <a-button type="danger" style="margin-right:15px">提交</a-button>
          <a-button type="primary">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>
<script>
export default {
  artInfo: {
    id: 0,
    title: '',
    cid: 0,
    desc: '',
    content: '',
    img: ''
  }
}
</script>
<style></style>
