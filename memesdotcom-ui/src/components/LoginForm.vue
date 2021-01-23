<template>
  <div ref="loginWrapper" @keypress.enter="() => {}">
    <div class="flex justify-between mb-2">
      <div class="text-gray-100">Log in</div>
      <button class="text-gray-100 underline" @click="$emit('changeView')">Or create an account</button>
    </div>
    <input type="text" v-model="name" class="w-full text-white p-3 rounded my-2" style="background-color: #364048" placeholder="Username | Email">
    <input type="password" v-model="password" class="w-full text-white p-3 rounded my-2" style="background-color: #364048" placeholder="Password">
    <a href="#"><div class="text-xs underline text-gray-500">Forgot password</div></a>

    <button @click="login" class="w-full mt-4 p-3 rounded text-gray-800 font-bold" style="background-color: #A2F537">
      Sign in
      <Icon class="ml-1" icon="arrow-right"/>
    </button>
    <hr class="text-white my-5" data-divider="Or sign in with"/>
    <div class="flex flex-wrap">
      <button class="text-gray-100 w-1/2 p-2 bg-red-500 hover:bg-red-400 rounded font-bold">
        <Icon icon="envelope"/>
        Email
      </button>
    </div>
  </div>
</template>

<script>
import { startLoading } from '@/functions/spinner'
import { login } from '@/functions/auth'
import { ref } from 'vue'

export default {
  setup () {
    const name = ref('')
    const password = ref('')

    return {
      name,
      password,
      startLoading,
      login
    }
  }
}
</script>
