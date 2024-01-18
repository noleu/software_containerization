<script setup lang="ts">
import { Modal } from 'flowbite'
import { onMounted, ref } from 'vue'
import LoginModal from './LoginModal.vue'
import RegisterModal from './RegisterModal.vue'

const loginModal = ref<Modal>()
const registerModal = ref<Modal>()

onMounted(() => {
  const $loginModalElement: HTMLElement | null = document.querySelector('#login-modal')
  loginModal.value = new Modal($loginModalElement)
  const $registerModalElement: HTMLElement | null = document.querySelector('#register-modal')
  registerModal.value = new Modal($registerModalElement)
})
</script>

<template>
  <nav
    class="bg-white dark:bg-gray-900 sticky w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600"
  >
    <div class="max-w-screen-xl flex flex-wrap items-center justify-end mx-auto p-4">
      <div class="flex gap-3">
        <button
          id="loginButton"
          type="button"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          @click="loginModal?.show()"
        >
          Login
        </button>
        <button
          type="button"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          @click="registerModal?.show()"
        >
          Register
        </button>
      </div>
    </div>
  </nav>
  <LoginModal
    @close="loginModal?.hide()"
    @register="
      () => {
        loginModal?.hide()
        registerModal?.show()
      }
    "
  />
  <RegisterModal
    @close="registerModal?.hide()"
    @login="
      () => {
        registerModal?.hide()
        loginModal?.show()
      }
    "
  />
</template>
