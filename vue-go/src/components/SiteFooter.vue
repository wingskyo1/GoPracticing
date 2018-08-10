<!-- Sitefooter.vue -->

<template>
  <div class="hello2">
    <h2>{{ msg }}</h2>
  </div>
</template>
<script>
export default {
  name: 'sitfooter',
  data () {
    return {
      msg: '住就是底部啦'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .hello2  {
    background-color:rgb(41, 128, 99);
    color:rgb(20, 50, 77) ;
    position: fixed;
    left: 0;
    bottom: 0;
    text-align: center;
    width: 100%;
  }
</style>
