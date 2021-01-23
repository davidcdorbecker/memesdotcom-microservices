<template>
<div class="text-gray-100 text-xs pb-5">
  <div class="flex p-2 justify-between items-center">
    <div class="flex items-center">
      <div class="w-6 h-6 bg-green-500 rounded-full mr-2"/>
      <div>
        <div class="leading-none">davidhernandez</div>
        <div class="leading-none text-xxs text-gray-400">16h</div>
      </div>
    </div>
    <div>
      <Icon
        icon="ellipsis-v"
        class="cursor-pointer"
      />
    </div>
  </div>
  <img :src="meme" alt="meme">
</div>
</template>

<script>
import meme from '@/assets/img/memes/meme.jpeg'
export default {
  setup () {
    return {
      meme
    }
  }
}
</script>

<style scoped>

</style>
