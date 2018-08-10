<!-- HardenDown.vue -->

<template>
  <div class="hello">
    <!-- <img src="../assets/harden.jpg"> -->
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'HardenDown',
  data () {
    return {
      msg: '第二頁 不要亂點啦!!'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .hello img {
    transform: perspective(600px) rotateX(80deg);
  }
</style>
