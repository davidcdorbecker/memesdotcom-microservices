<template>
  <div class="spinner items-center" v-if="isLoading">
    <img class="h-24 w-24" :src="loader" alt="loader">
  </div>
</template>

<script>
import { isLoading } from '@/functions/spinner'
import loader from '@/assets/img/loader.svg'

export default {
  setup () {
    return {
      isLoading,
      loader
    }
  }
}
</script>

<style scoped>
.spinner {
  z-index: 9999;
  position: fixed;
  display: flex;
  justify-content: center;
  top: 0;
  left: 0;
  background-color: #1e272e;
  width: 100%;
  height: 100%;
}
</style>
