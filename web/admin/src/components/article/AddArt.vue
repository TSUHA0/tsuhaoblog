<template>
  <div>
    <h3>{{ Add ? '新建文章' : '编辑文章' }}</h3>
    <a-card>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
        <a-row :gutter="16">
          <a-col :span="16">
            <a-form-model-item label="文章标题" prop="title">
              <a-input v-model="artInfo.title"></a-input>
            </a-form-model-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :lg="{ span: 6 }" :xs="{ span: 8 }">
            <a-form-model-item label="文章分类" prop="cid">
              <a-select v-model="artInfo.cid" @change="CateChange" placeholder="请选择分类">
                <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">
                  {{ item.name }}
                </a-select-option>
              </a-select>
            </a-form-model-item>
          </a-col>
        </a-row>

        <a-row :gutter="16">
          <a-col :span="16">
            <a-form-model-item label="文章描述" prop="desc">
              <a-input type="textarea" v-model="artInfo.desc"></a-input>
            </a-form-model-item>
          </a-col>
        </a-row>

        <a-form-model-item label="文章缩略图" prop="img">
          <a-upload list-type="picture" name="file" :action="upUrl" :headers="headers" @change="upChange">
            <a-button> <a-icon type="upload" /> Upload </a-button>
          </a-upload>
          <img src="artInfo.img" />
        </a-form-model-item>

        <a-form-model-item label="文章内容">
          <mavon-editor ref="md" v-model="artInfo.content" @imgAdd="$imgAdd"></mavon-editor>
        </a-form-model-item>
        <a-form-model-item class="btn">
          <a-button type="danger" style="margin-right:15px" @click="submitClk">提交</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http.js'
import axios from 'axios'
export default {
  data() {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      Add: false,
      CurrCateName: '请选择分类',
      Catelist: [],
      headers: {},
      upUrl: Url + 'upload',
      token: `Bearer ${window.sessionStorage.getItem('token')}`,
      artInfoRules: {
        title: [{ required: true, message: '请输入文章标题', trigger: 'change' }],
        cid: [{ required: true, message: '请选择文章分类', trigger: 'change' }],
        desc: [
          { required: true, message: '请输入文章描述', trigger: 'change' },
          { max: 120, message: '描述最多可写120个字符', trigger: 'change' }
        ]
      }
    }
  },
  mounted() {
    this.getCateList()
    this.Add = Object.keys(this.$route.params).indexOf('id') < 0
    this.headers = { Authorization: this.token }
    if (this.$route.params.id) {
      this.getArtInfo(this.$route.params.id)
    }
  },
  methods: {
    CateChange(value) {
      this.artInfo.cid = value
    },
    // 绑定@imgAdd event
    async $imgAdd(pos, $file) {
      const $vm = this.$refs.md
      // 第一步.将图片上传到服务器.
      var formdata = new FormData()
      formdata.append('file', $file)
      axios({
        url: Url + 'upload',
        method: 'post',
        data: formdata,
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: this.token
        }
      }).then(url => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         */
        $vm.$img2Url(pos, url.data.url)
      })
    },
    // 查询文章信息
    async getArtInfo(id) {
      const { data: res } = await this.$http.get(`art/${id}`)
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    // 获取分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('cates')
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.Catelist = res.data
    },
    submitClk() {
      this.$refs.artInfoRef.validate(async valid => {
        if (!valid) return this.$message.error('参数验证未通过，请按要求录入文章内容')
        let res
        if (this.Add) res = await this.$http.post('art/add', this.artInfo)
        else res = await this.$http.put(`art/${this.$route.params.id}`, this.artInfo)
        if (res.status !== 200) return this.$message.error(res.message)
        this.$router.push('/admin/artlist')
        this.$message.success('添加文章成功')
      })
    },
    // 上传图片
    upChange(info) {
      if (info.file.status === 'done') {
        let s = info.file.response.status
        if (s === 1004 || s === 1005 || s === 1006 || s === 1007) return this.$message.error(info.file.response.message)
        this.$message.success(`图片上传成功`)
        this.artInfo.img = info.file.response.url
      } else if (info.file.status === 'error') {
        this.$message.error(`图片上传失败`)
      }
    },
  }
}
</script>
<style></style>
