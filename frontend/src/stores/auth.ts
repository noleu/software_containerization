import { useAxios } from '@vueuse/integrations/useAxios'
import { ref } from 'vue'
import { defineStore } from 'pinia'
import instance from '@/libs/axios'

export const useAuthStore = defineStore('auth', () => {
  const me = ref([])

  function login({ email, password }: { email: string; password: string }) {
    try {
      const { data } = useAxios('/login', { method: 'POST', data: { email, password } }, instance)
      me.value = data.value
    } catch (error) {
      console.log(error)
    }
  }

  function register({ email, password }: { email: string; password: string }) {
    try {
      const { data } = useAxios(
        '/register',
        { method: 'POST', data: { email, password } },
        instance
      )
      me.value = data.value
    } catch (error) {
      console.log(error)
    }
  }

  return { me, login, register }
})
