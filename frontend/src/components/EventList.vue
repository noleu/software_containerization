<script setup lang="ts">
import { defineProps } from 'vue'
import { type ReturnEvent } from '@/interfaces/event'

defineProps({
  events: { type: Array<ReturnEvent>, required: true }
})

function formattedEventDate(eventDate: Date) {
  const months = [
    'Jan',
    'Feb',
    'Mar',
    'Apr',
    'May',
    'Jun',
    'Jul',
    'Aug',
    'Sep',
    'Oct',
    'Nov',
    'Dec'
  ]

  const monthIndex = eventDate.getMonth()
  const day = eventDate.getDate()

  const formattedDate = `${months[monthIndex]} ${day}`

  return formattedDate
}
</script>

<template>
  <div>
    <ul v-if="events.length > 0">
      <li v-for="event in events" :key="event.id" class="bg-white shadow-lg rounded-lg mb-4">
        <div class="flex">
          <div class="flex flex-col bg-blue-600 rounded-tl-lg rounded-bl-lg p-4">
            <h1 class="text-white">{{ formattedEventDate(event.date_time) }}</h1>
            <p class="text-sm text-white">{{ event.location }}</p>
          </div>
          <div class="p-4">
            <h1 class="text-xl font-semibold">{{ event.title }}</h1>
            <p class="text-sm overflow-auto break-words">{{ event.description }}</p>
          </div>
        </div>
      </li>
    </ul>
    <p v-else class="text-gray-600">No events found</p>
  </div>
</template>
