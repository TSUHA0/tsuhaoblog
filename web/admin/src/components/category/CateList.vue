<template>
  <div>
    <h3>分类列表页面</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addCateclick">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Catelist"
        @change="handleTableChange"
        bordered
      >
        <template slot="btnslot" slot-scope="record">
          <a-button type="primary" @click="editCateclick(record.id)" style="margin-right:15px">编辑</a-button>
          <a-button type="danger" @click="deleteCate(record.id)">删除</a-button>
        </template>
      </a-table>
    </a-card>
    <!-- 新增分类 -->
    <a-modal :closable="false" :visible="addCateModalVisible" title="新增分类" @ok="AddCateOk" @cancel="AddCateCancel">
      <a-form-model ref="refaddCate" :rules="addrules" :model="CateInfo">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="CateInfo.name" placeholder="请输入分类名"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑分类 -->
    <a-modal
      :closable="false"
      :visible="editCateModalVisible"
      title="编辑分类"
      @ok="EditCateOk"
      @cancel="EditCateCancel"
    >
      <a-form-model ref="refeditCate" :rules="editrules" :model="CateInfo">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="CateInfo.name" placeholder="请输入分类名"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
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
      Catelist: [],
      columns,
      queryParam: {
        name: '',
        pagesize: 5,
        pagenum: 1
      },
      addCateModalVisible: false,
      editCateModalVisible: false,
      CateInfo: {
        id: 0,
        name: '',
      },
      addrules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' },
          { min: 0, max: 12, message: '分类名长度不应超过20字符', trigger: 'blur' }
        ]
      },
      editrules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' },
          { min: 0, max: 12, message: '分类名长度不应超过20字符', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getCateList()
  },
  methods: {
    //查询分类列表
    async getCateList() {
      const { data: res } = await this.$http.get('cates', {
        params: {
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
      this.Catelist = res.data
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
      this.getCateList()
    },
    async deleteCate(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该分类吗？一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`cate/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.success('已取消删除')
        }
      })
    },
    //添加分类
    addCateclick() {
      this.CateInfo.name = ''
      this.addCateModalVisible = true
    },
    AddCateOk() {
      this.$refs.refaddCate.validate(async valid => {
        if (!valid) {
          return this.$message.error('输入数据不合法')
        } else {
          const { data: res } = await this.$http.post('cate/add', {
            name: this.CateInfo.name
          })

          if (res.status != 200) {
            for (let s in res) if (s == 'error') return this.$message.error(res.error)
            return this.$message.error(res.message)
          }
          this.$message.info('添加分类成功')
          this.getCateList()
          this.addCateModalVisible = false
        }
      })
    },
    AddCateCancel() {
      this.$message.info('已取消添加分类')
      this.CateInfo.name=""
      this.addCateModalVisible = false
    },
    adminChange(checked) {
      if (checked) {
        this.CateInfo.role = 1
      } else {
        this.CateInfo.role = 2
      }
    },
    //编辑分类
    async editCateclick(id) {
      const {data:res} = await this.$http.get(`cate/${id}`)
      this.CateInfo = res.data
      this.editCateModalVisible = true
    },
    EditCateOk() {
      this.$refs.refeditCate.validate(async valid => {
        if (!valid) {
          return this.$message.error('输入数据不合法')
        } else {
          const { data: res } = await this.$http.put(`cate/${parseInt(this.CateInfo.id)}`, {
            name: this.CateInfo.name
          })

          if (res.status != 200) {
            for (let s in res) if (s == 'error') return this.$message.error(res.error)
            return this.$message.error(res.message)
          }
          this.$message.info('编辑分类成功')
          this.getCateList()
          this.editCateModalVisible = false
        }
      })
    },
    EditCateCancel() {
      this.$message.info('编辑已取消')
      this.$refs.refeditCate.resetFields()
      this.editCateModalVisible = false
    }
  }
}
</script>

<style scoped></style>
