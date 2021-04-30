<template>
  <v-app-bar app>
    <!-- 
    <v-container  class="py-0 fill-height">
      <v-btn v-for="item in catelist" :key=item.id class="mr-2" text>{{item.name}}</v-btn>
    </v-container> -->

    <v-tabs show-arrows>
      <v-tabs-slider color="red"></v-tabs-slider>
      <v-tab @click="backhome">主页</v-tab>
      <v-tab
        v-for="item in catelist"
        :key="item.id"
        @click="$router.push(`/category/${item.id}`)"
      >
        {{ item.name }}
      </v-tab>
    </v-tabs>

    <v-spacer></v-spacer>

    <v-switch
      v-model="$vuetify.theme.dark"
      hint="切换黑暗模式"
      inset
      label="关灯"
      persistent-hint
    ></v-switch>
    
<!-- @click="pushadmin" -->
    <v-btn icon @click="pushadmin">
      <v-icon>mdi-account-key</v-icon>
    </v-btn>

    <v-btn icon @click="$router.push(`/search/${title}`)">
      <v-icon>mdi-magnify</v-icon>
    </v-btn>

    <v-responsive max-width="260" color="white">
      <v-text-field
        flat
        dense
        hide-details
        rounded
        black
        solo-inverted
        placeholder="请输入文章标题查找"
        v-model="title"
      ></v-text-field>
    </v-responsive>
  </v-app-bar>
</template>
<script>
export default {
  data() {
    return {
      catelist: [],
      title: ''
    }
  },
  created() {
    this.getCatelist()
  },
  methods: {
    async getCatelist() {
      const { data: res } = await this.$http.get('cates')
      this.catelist = res.data
    },
    backhome() {
      this.$router.push('/')
    },
    pushadmin(){
      let domain = window.location.host
      console.log('domain: ', domain);
      window.location.href = "http://"+domain+"/admin";
    }

  }
}
</script>
<style></style>
