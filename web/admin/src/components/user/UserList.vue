<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="输入用户名查找"
            enter-button
            allowClear
            @search="searchUser"
          ></a-input-search>
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUserclick">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="ID"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        @change="handleTableChange"
        bordered
      >
        <template slot="roleslot" slot-scope="text">{{ text == 1 ? '管理员' : '订阅者' }}</template>
        <template slot="btnslot" slot-scope="record">
          <a-button type="primary" @click="editUserclick(record.ID)" style="margin-right:15px">编辑</a-button>
          <a-button type="danger" @click="deleteUser(record.ID)">删除</a-button>
        </template>
      </a-table>
    </a-card>
    <!-- 新增用户 -->
    <a-modal :closable="false" :visible="addUserModalVisible" title="新增用户" @ok="AddUserOk" @cancel="AddUserCancel">
      <a-form-model ref="refaddUser" :rules="addrules" :model="UserInfo">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="UserInfo.username" placeholder="请输入用户名"></a-input>
        </a-form-model-item>
        <a-form-model-item hasFeedback label="密码" prop="password">
          <a-input-password v-model="UserInfo.password" placeholder="请输入密码"></a-input-password>
        </a-form-model-item>
        <a-form-model-item hasFeedback label="确认密码" prop="checkpass">
          <a-input-password v-model="UserInfo.checkpass" placeholder="请重新输入密码"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户 -->
    <a-modal
      :closable="false"
      :visible="editUserModalVisible"
      title="编辑用户"
      @ok="EditUserOk"
      @cancel="EditUserCancel"
    >
      <a-form-model ref="refeditUser" :rules="editrules" :model="UserInfo">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="UserInfo.username" placeholder="请输入用户名"></a-input>
        </a-form-model-item>

        <a-form-model-item label="是否为管理员" prop="role">
          <a-switch
            :checked="IsAdmin"
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'roleslot' }
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'btnslot' }
  }
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: total => `共${total}条`
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      addUserModalVisible: false,
      editUserModalVisible: false,
      UserInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      addrules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名长度应在4~12字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码长度应在6~20字符', trigger: 'blur' }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (value === '') {
                callback(new Error('请重复输入密码'))
              } else {
                if (this.UserInfo.checkpass !== this.UserInfo.password) {
                  callback(new Error('两次密码输入不同'))
                }
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editrules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名长度应在4~12字符', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getUserList()
  },
  computed: {
    IsAdmin: function () {
      if (this.UserInfo.role === 1) {
        return true
      } else {
        return false
      }
    },
  },
  methods: {
    //查询用户列表
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.userlist = res.data
      this.pagination.total = res.total
    },
    //分页
    handleTableChange(pagination) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagenum = pagination.current
      this.queryParam.pagesize = pagination.pageSize

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    //查找用户
    async searchUser() {
      this.queryParam.pagenum = 1
      this.pagination.current = 1
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.pagination.total = res.total
    },
    async deleteUser(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该用户吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.success('已取消删除')
        }
      })
    },
    //添加用户
    addUserclick(){
      this.UserInfo.username=""
      this.UserInfo.password=""
      this.addUserModalVisible=true
    },
    AddUserOk() {
      this.$refs.refaddUser.validate(async valid => {
        if (!valid) {
          return this.$message.error('输入数据不合法')
        } else {
          const { data: res } = await this.$http.post('user/add', {
            username: this.UserInfo.username,
            password: this.UserInfo.password,
            role: parseInt(this.UserInfo.role)
          })

          if (res.status != 200) {
            for (let s in res) if (s == 'error') return this.$message.error(res.error)
            return this.$message.error(res.message)
          }
          this.$message.info('添加用户成功')
          this.getUserList()
          this.addUserModalVisible = false
        }
      })
    },
    AddUserCancel() {
      this.$refs.refaddUser.resetFields()
      this.addUserModalVisible = false
    },
    adminChange(checked) {
      if (checked) {
        this.UserInfo.role = 1
      } else {
        this.UserInfo.role = 2
      }
    },
    //编辑用户
    async editUserclick(id) {
      const { data: res } = await this.$http.get(`user/${id}`)
      this.UserInfo=res.data
      this.editUserModalVisible = true
    },
    EditUserOk() {
      this.$refs.refeditUser.validate(async valid => {
        if (!valid) {
          return this.$message.error('输入数据不合法')
        } else {
          const { data: res } = await this.$http.put(`user/${parseInt(this.UserInfo.id)}`, {
            username: this.UserInfo.username,
            role: parseInt(this.UserInfo.role)
          })

          if (res.status != 200) {
            for (let s in res) if (s == 'error') return this.$message.error(res.error)
            return this.$message.error(res.message)
          }
          this.$message.info('编辑用户成功')
          this.getUserList()
          this.editUserModalVisible = false
        }
      })
    },
    EditUserCancel() {
      this.$message.info('编辑已取消')
      this.$refs.refeditUser.resetFields()
      this.editUserModalVisible = false
    }
  }
}
</script>

<style scoped></style>
