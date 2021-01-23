<template>
  <Spinner/>
  <Layout/>
</template>
<script>
import Layout from '@/components/Layout'
export default {
  components: { Layout }
}
</script>
<style>
@font-face {
  font-family: 'Nunito Regular';
  src: url('assets/fonts/Nunito-Regular.ttf');
}
#app {
  font-family: 'Nunito Regular', serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
