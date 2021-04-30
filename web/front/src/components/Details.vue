<template>
  <div>
    <div class="d-flex justify-center pa-3 ma-1 text-h4 font-weight-bold">
      {{ artInfo.title }}
    </div>
    <div class="d-flex justify-center align-center">
      <div class="d-flex mx-10 justify-center">
        <v-icon class="mr-1" color="indigo" small>
          {{ "mdi-calendar-month" }}
        </v-icon>
        <span>{{ artInfo.CreatedAt | dateformat("YYYY-MM-DD") }}</span>
      </div>
      <div class="d-flex mr-10 justify-center">
        <v-icon class="mr-1" color="pink" small>{{ "mdi-comment" }}</v-icon>
        <span>{{ total }}</span>
      </div>
      <div class="d-flex mr-10 justify-center">
        <v-icon class="mr-1" color="green" small>{{ "mdi-eye" }}</v-icon>
        <span>{{ artInfo.read_count }}</span>
      </div>
    </div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert
      class="ma-4"
      elevation="1"
      color="indigo"
      dark
      border="left"
      outlined
      >{{ artInfo.desc }}</v-alert
    >

    <div v-html="artInfo.content" class="content ma-5 pa-3 text-justify"></div>
  </div>
</template>
<script>
export default {
  props: ["id"],
  data() {
    return {
      artInfo: {},
      commentList: [],
      total: 0,
      headers: {
        username: "",
        user_id: 0,
      },
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },
    };
  },
  created() {
    this.getArtInfo();
    this.headers = {};
  },
  methods: {
    // 查询文章
    async getArtInfo() {
      const { data: res } = await this.$http.get(`art/${this.id}`);
      this.artInfo = res.data;
    },
  },
};
</script>
<style scoped>
.content >>> div,
img,
span {
  width: auto;
  max-width: 100%;
}

.content >>> pre,
code {
  margin: 10px;
  padding: 14px;
  overflow: auto;
  font-size: 85%;
  line-height: 1.45;
  background-color: rgba(27, 31, 35, 0.05);
  border-left-width: 0.5rem;
  border-left-style: solid;
  border-color: #303f9f;
}
</style>
