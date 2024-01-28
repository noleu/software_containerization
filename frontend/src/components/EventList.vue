<script setup lang="ts">
import { defineProps } from 'vue'
import { type ReturnEvent } from '@/interfaces/event'

defineProps({
  events: { type: Array<ReturnEvent>, required: true }
})
</script>

<template>
  <ul class="space-y-2" v-if="events.length > 0">
    <li v-for="event in events" :key="event.id" class="border rounded-md p-2">
      <span>{{ event.date_time }}</span>
      <h1>{{ event.title }}</h1>
      <p>{{ event.description }}</p>
      <span>{{ event.location }}</span>
    </li>
  </ul>
  <span v-else>No events found</span>
</template>
