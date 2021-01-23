<template>
  <div class="flex justify-center mt-16 text-sm px-4">
    <div class="w-full sm:w-64" ref="loginWrapper">
      <div v-if="loginMode">
        <LoginForm @changeView="changeView"/>
      </div>
      <div v-else>
        <RegisterForm @changeView="changeView"/>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import LoginForm from '@/components/LoginForm'
import RegisterForm from '@/components/RegisterForm'

export default {
  components: { LoginForm, RegisterForm },
  setup () {
    const loginMode = ref(true)

    function changeView () {
      loginMode.value = !loginMode.value
    }

    return {
      loginMode,
      changeView
    }
  }
}
</script>
