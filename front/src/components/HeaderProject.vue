<template>
  <header class="header-project">
    <div class="container">
      <div class="header-items">
          <button class="logo" @click="showMenu = !showMenu">
            <img src="https://via.placeholder.com/60" alt="">
            <p>Имя</p>
          </button>
        <header-menu @close="closeMenu" :visibility="showMenu"></header-menu>
        <div class="right-item">
          <div class="coins">
            <img src="../assets/2.svg" alt="">
            <p>100</p>
          </div>
          <button-item @click="showShop = !showShop">Магазин</button-item>
          <shop :visibility="showShop" @closeDialog="test">
            <!-- <h3 v-if="!cartItems.length">Товары скоро появятся</h3> -->
          </shop>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import Shop from '../components/HeaderShop.vue'
import HeaderMenu from '../components/HeaderMenu.vue'
import ButtonItem from '../components/ButtonItem.vue'

export default {
  name: 'HeaderProject',
  props: [
  ],
  components: {
    ButtonItem,
    HeaderMenu,
    Shop
  },
  data(){
    return {
      visibility: false,
      showMenu: false,
      showShop: false,
      // cartItems: []
    }
  },
  methods: {
    test() {
      console.log('click');
      this.showShop = false
    },
    closeMenu(e) {
      if (e.target.className === 'he')
      console.log(e.target.className);
    }
  }
}
</script>

<style scoped>
  .header-project {
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    height: 80px;
    background: #FF8562;
  }

  .container{
    height: 100%;
    width: 1186px;
    margin: 0 auto;
  }

  .header-items {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .logo img{
    border-radius: 50px;
    margin-right: 15px;
  }

  .logo {
    border: none;
    background: none;
    display: flex;
    align-items: center;
    color: white;
    font-weight: 600;
  }

  .right-item {
    display: flex;
    align-items: center;
  }

  .coins {
    display: flex;
    margin-right: 36px;
  }

  .coins p {
    font-size: 18px;
    font-weight: bold;
    color: white;
    margin-left: 10px;
  }
</style>
