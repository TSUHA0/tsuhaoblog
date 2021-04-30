<template>
  <v-app app>
    <Topbar></Topbar>
    <v-main class="grey">
      <v-container
        class="mb-2"
        fluid
        style="justify-center margin: 0px; padding: 0px; width: 90%"
      >
        <v-row class="mt-3">
          <v-col cols="12" md="3">
            <Nav></Nav>
          </v-col>

          <v-col cols="12" md="8">
            <v-sheet min-height="80vh" rounded="lg">
              <router-view :key="$route.path"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer></Footer>
  </v-app>
</template>

<script>
import Footer from '../components/Footer'
import Topbar from '../components/TopBar'
import Nav from '../components/Nav'
export default {
  components: {
    Footer,
    Topbar,
    Nav
  }
}
</script>

<style scoped></style>
