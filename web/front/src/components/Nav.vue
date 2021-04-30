<template>
  <v-container>
    <v-img src="../assets/nav.png">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="white">
            <img src="../assets/pic.png" class="widthSet" alt />
          </v-avatar>
          <div class="ma-4 white--text">渠昊</div>
        </v-col>
      </v-card-title>
      <v-divider></v-divider>
    </v-img>  

    <v-divider color="indigo"></v-divider>

    <v-list nav dense>
      <v-card-title>About Me:</v-card-title>
      <v-card-text >中科大软件学院学生</v-card-text>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="blue darken-2">{{ "mdi-qqchat" }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">1355738338</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="green darken-2">{{ "mdi-wechat" }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">quhaohi</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="black darken-2">{{ "mdi-github" }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">
          <a href="https://github.com/TSUHA0" target="_blank" style="color:grey;">Github</a>
        </v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="black darken-2">{{ "mdi-code-tags" }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">
          <a href="https://leetcode-cn.com/u/tsuhao/" target="_blank" style="color:grey;">Leetcode力扣</a>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-container>
</template>
<script>
export default {};
</script>
<style scoped>
</style>
