import Vue from 'vue'
import axios from 'axios'

const config = {
    baseURL: "http://localhost:8081/"
}

axios.create(config)

Vue.prototype.$http = axios
