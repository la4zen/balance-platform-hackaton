<template>
  <div id="app">
    <HeaderProject/>
    <main-project></main-project>
    <FooterProject/>
  </div>
</template>

<script>
// import axios from "axios";
import FooterProject from './components/FooterProject.vue'
import MainProject from './components/MainProject.vue'
import HeaderProject from './components/HeaderProject.vue'

export default {
  name: 'App',
  components: {
    HeaderProject,
    FooterProject,
    MainProject
    // HeaderMenu
  }
}
</script>

<style>
  body {
    margin: 0;
    padding: 0;
    font-family: 'Roboto Slab', serif;
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
