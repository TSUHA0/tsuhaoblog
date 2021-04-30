<template>
  <div class="container">
    <div class="loginbox">
      <p class="logintitle">
        TsuhaoBlog 后台管理系统
      </p>
      <a-form-model ref="refform" :rules="rules" class="loginform" :model="formdata">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="请输入用户名" @keyup.enter="loginfn">
            <a-icon slot="prefix" type="user" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>

        <a-form-model-item prop="password">
          <a-input v-model="formdata.password" placeholder="请输入密码" type="password" @keyup.enter="loginfn">
            <a-icon slot="prefix" type="lock" style="color:rgba(0,0,0,.25)" />
          </a-input>
        </a-form-model-item>

        <a-form-model-item class="loginbtn">
          <a-button type="primary" style="margin:10px" @click="loginfn">登陆</a-button>
          <a-button type="danger" style="margin:10px" @click="cancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名长度应在4~12字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度应在6~20字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    cancel: function() {
      this.$refs.refform.resetFields()
    },
    loginfn: function() {
      this.$refs.refform.validate(async valid => {
        if (!valid) {
          this.$message.error('输入非法登录数据，请重新输入')
          return false
        } else {
          const { data: res } = await this.$http.post('login', this.formdata)
          if (res.code != 200) {
            this.$message.error(res.message)
            return
          }
          window.sessionStorage.setItem('token', res.token)
          this.$router.push('admin/index')
        }
      })
    }
  }
}
</script>
<style scoped>
.logintitle {
  position: relative;
  margin: 5px;
  font-size: 20px;
  color: black;
}
.container {
  height: 100%;
  background-image: linear-gradient(to bottom right,rgb(101, 208, 235),rgb(130, 150, 196));
}
.loginbox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: fixed;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 10px;
  box-shadow: 0px 10px 10px 0px rgb(62, 170, 233);
}

.loginform {
  width: 100%;
  position: fixed;
  bottom: 10%;
  padding: 0px 20px;
  box-sizing: border-box;
}

.loginbtn {
  display: flex;
  justify-content: flex-end;
}
</style>
