import { useAxios } from '@vueuse/integrations/useAxios'
import { ref } from 'vue'
import { defineStore } from 'pinia'
import instance from '@/libs/axios'

export const useEventStore = defineStore('event', () => {
  const events = ref([])

  function getEvents() {
    try {
      const { data } = useAxios('/events', { method: 'GET' }, instance)
      events.value = data.value
    } catch (error) {
      console.log(error)
    }
  }

  return { events, getEvents }
})
