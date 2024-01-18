import axios from 'axios'
import { useAxios } from '@vueuse/integrations/useAxios'

const instance = axios.create({
  baseURL: 'http://localhost:8080/api'
})

export default instance
