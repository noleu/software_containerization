<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue'
</script>

<template>
  <header>
    <NavBar />
  </header>
  <div class="flex justify-center items-center">
    <div class="w-2/3">
      <RouterView />
    </div>
  </div>
</template>
