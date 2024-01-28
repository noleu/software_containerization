<script setup lang="ts">
import EventList from '@/components/EventList.vue'
import EventSearch from '@/components/EventSearch.vue'
import CreateEventModal from '@/components/CreateEventModal.vue'
import { useEventStore } from '@/stores/event'
import { ref, onMounted, computed } from 'vue'
import { Modal } from 'flowbite'
import { type CreateEvent } from '@/interfaces/event'

const eventStore = useEventStore()
const createEventModal = ref<Modal>()

eventStore.getEvents()

const searchString = ref('')

onMounted(() => {
  const $createEventModalElement: HTMLElement | null = document.querySelector('#create-event-modal')
  createEventModal.value = new Modal($createEventModalElement)
})

const searchedEvents = computed(() => {
  if (eventStore.events) {
    return eventStore.events.filter((event) => {
      return event.title.toLowerCase().includes(searchString.value.toLowerCase())
    })
  }
  return []
})

function createEvent(eventData: CreateEvent) {
  eventStore
    .createEvent(eventData)
    .then((res) => {
      if (res) {
        eventStore.getEvents()
      }
      createEventModal.value?.hide()
    })
    .catch((err) => {
      console.error('Error creating event:', err)
    })
}
</script>

<template>
  <header class="mt-5 mb-2 text-center">
    <h1 class="text-3xl font-bold">Find some cool events</h1>
  </header>
  <div class="flex justify-center">
    <EventSearch v-model="searchString" class="mr-2" />
    <div class="flex items-center">
      <button
        @click="createEventModal?.show()"
        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        + Create Event
      </button>
    </div>
  </div>
  <EventList v-if="searchedEvents" class="mt-2" :events="searchedEvents" />

  <CreateEventModal @close="createEventModal?.hide()" @create-event="createEvent" />
</template>
