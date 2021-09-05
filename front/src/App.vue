<template>
  <div id="app">
    <HeaderProject/>
    <!-- <header-menu></header-menu> -->
    <FooterProject/>
  </div>
</template>

<script>
import FooterProject from './components/FooterProject.vue'
// import HeaderMenu from './components/HeaderMenu.vue'
import HeaderProject from './components/HeaderProject.vue'

export default {
  name: 'App',
  components: {
    HeaderProject,
    FooterProject,
    // HeaderMenu
  }
}
</script>

<style>
  body {
    margin: 0;
    padding: 0;
    font-family: Roboto;
  }

  .container{
    height: 100%;
    width: 1186px;
    margin: 0 auto;
  }

  button{
    cursor: pointer;
  }
</style>
