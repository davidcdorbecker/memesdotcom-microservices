<template>
  <div class="h-screen bg-dark">
    <div class="fixed top-0 w-full bg-dark flex justify-between items-center p-3 h-12 shadow-md">
      <img class="h-3" :src="logo" alt="logo">
      <div class="text-white text-lg">
        <div class="flex items-center" v-if="isUserLogged">
          <Navbar class="hidden sm:flex"/>
          <Icon
            icon="bell"
            class="mx-4 cursor-pointer hover:bg-red"
          />
          <Icon
            @click="logout"
            icon="power-off"
            class="mx-4 cursor-pointer hover:text-red-500"
          />
        </div>
      </div>
    </div>
    <div>
      <RouterView/>
    </div>
    <div
        v-if="isUserLogged"
        class="fixed bottom-0 w-full h-12 shadow-lg block sm:hidden"
        style="background-color: #1b2228"
    >
      <Navbar/>
    </div>
  </div>
</template>
<script>
import logo from '@/assets/img/logo.svg'
import { isUserLogged, logout } from '@/functions/auth'
import Navbar from '@/components/Navbar'

export default {
  components: { Navbar },
  setup () {
    return {
      logo,
      isUserLogged,
      logout
    }
  }
}
</script>
