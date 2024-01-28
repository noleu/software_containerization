<script setup lang="ts">
import EventList from '@/components/EventList.vue'
import EventSearch from '@/components/EventSearch.vue'
import CreateEventModal from '@/components/CreateEventModal.vue'
import { useEventStore } from '@/stores/event'
import { ref, onMounted } from 'vue'
import { Modal } from 'flowbite'
import { type CreateEvent } from '@/interfaces/event'

const eventStore = useEventStore()
const createEventModal = ref<Modal>()

eventStore.getEvents()

onMounted(() => {
  const $createEventModalElement: HTMLElement | null = document.querySelector('#create-event-modal')
  createEventModal.value = new Modal($createEventModalElement)
})

function createEvent(eventData: CreateEvent) {
  eventStore.createEvent(eventData)
  createEventModal.value?.hide()
}
</script>

<template>
  <header class="mt-5 mb-2 text-center">
    <h1 class="text-3xl font-bold">Find some cool events</h1>
  </header>
  <div class="flex justify-center">
    <EventSearch class="mr-2" />
    <div class="flex items-center">
      <button
        @click="createEventModal?.show()"
        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        + Create Event
      </button>
    </div>
  </div>
  <EventList v-if="eventStore.events" class="mt-2" :events="eventStore.events" />

  <CreateEventModal @close="createEventModal?.hide()" @create-event="createEvent" />
</template>
