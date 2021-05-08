import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://119.45.205.191:80/api/v1/'
Vue.prototype.$http = axios
