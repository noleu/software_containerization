import { useAxios } from '@vueuse/integrations/useAxios'
import { ref } from 'vue'
import { defineStore } from 'pinia'
import instance from '@/libs/axios'
import type { CreateEvent, ReturnEvent } from '@/interfaces/event'

export const useEventStore = defineStore('event', () => {
  const events = ref<ReturnEvent[]>()

  async function createEvent(event: CreateEvent) {
    try {
      const { data } = await useAxios('/events', { method: 'POST', data: event }, instance)
      events.value = data.value
      return data
    } catch (error) {
      console.log(error)
      return error
    }
  }

  async function getEvents() {
    try {
      const { data } = await useAxios<ReturnEvent[]>('/events', { method: 'GET' }, instance)
      events.value = data.value?.map((event) => ({
        ...event,
        date_time: new Date(event.date_time)
      }))
      return data
    } catch (error) {
      console.log(error)
      return error
    }
  }

  return { events, createEvent, getEvents }
})
