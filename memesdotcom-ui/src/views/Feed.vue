<template>
  <div class="w-full sm:w-144 mx-auto mt-12">
    <MemeCard v-for="index in 10" :key="index"/>
  </div>
</template>

<script>
import MemeCard from '@/components/MemeCard'
export default {
  components: { MemeCard }
}
</script>
