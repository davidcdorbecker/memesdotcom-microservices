<template>
  <div ref="registerWrapper">
    <div class="flex justify-between mb-2">
      <div class="text-gray-100">Sign up</div>
      <button class="text-gray-100 underline" @click="$emit('changeView')">Or login</button>
    </div>
    <input type="text" v-model="name" class="w-full text-white p-3 rounded my-2" style="background-color: #364048" placeholder="Username">
    <input type="text" v-model="email" class="w-full text-white p-3 rounded my-2" style="background-color: #364048" placeholder="Email">
    <input type="password" v-model="password" class="w-full p-3 rounded my-2" style="background-color: #364048" placeholder="Password">
    <a href="#"><div class="text-xs underline text-gray-500">Forgot password</div></a>

    <button @click="register" class="w-full mt-4 p-3 rounded text-gray-800 font-bold" style="background-color: #A2F537">
      Sign up
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
import axios from 'axios'

export default {
  data () {
    return {
      name: '',
      email: '',
      password: ''
    }
  },
  methods: {
    async register () {
      if (!this.name || !this.password || !this.email) {
        this.shakeElement(this.$refs.registerWrapper)
        return
      }
      const url = this.$store.state.database.remoteAuthUrl + '/register'
      const requestBody = {
        name: this.name,
        email: this.email,
        password: this.password
      }
      this.$store.dispatch('startLoading')
      try {
        const { data } = await axios.post(url, requestBody)
        localStorage.setItem('__mdb__', data.database)
        localStorage.setItem('__mtk__', data.token)
        this.$store.dispatch('setLoggerVisibility', this.username === 'dev')
        this.$router.push({ path: '/' })
      } catch (error) {
        console.log(error)
        const errors = error?.response?.data?.errors
        if (errors) {
          for (const error of errors) {
            this.$toast.add({ severity: 'error', summary: '', detail: error.msg, life: 3000 })
          }
        }
        this.shakeElement(this.$refs.registerWrapper)
        this.$store.dispatch('endLoading')
      }
    }
  }
}
</script>
