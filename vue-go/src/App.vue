/* eslint-disable */
<template>
  <div id="app">
    <!-- vue logo -->
    <img src="./assets/logo.png" id="logo">
    <!-- 導覽列 -->
    <nav>
      <router-link to="/">首頁</router-link>
      <router-link to="/down">連結二</router-link>
    </nav>
    <!-- 路由匹配到的组件將會渲染在這 -->
    <router-view/>
     <!-- 插入 footer.vue -->
    <SiteFooter></SiteFooter>
  </div>
</template>

<script>
import SiteFooter from '@/components/SiteFooter'// 載入組件
export default {
  name: 'app',
  components: {
    SiteFooter// 宣告組件
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
#logo{
  max-width: 60px;
}
</style>
