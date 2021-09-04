<template>
  <header class="header-project">
    <div class="container">
      <div class="header-items">
          <button class="logo" @click="showAccount = !showAccount">
            <img src="../assets/lk.png" alt="">
            <p>Мария</p>
          </button>
          <account :visibility="showAccount" @closeDialog="closeAccount"></account>
        <div class="right-item">
          <div class="coins">
            <img src="../assets/Stamp.png" alt="">
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
import Account from '../components/Account.vue'
import ButtonItem from '../components/ButtonItem.vue'

export default {
  name: 'HeaderProject',
  props: [
  ],
  components: {
    ButtonItem,
    Account,
    Shop
  },
  data(){
    return {
      visibility: false,
      showShop: false,
      showAccount: false,
    }
  },
  methods: {
    closeAccount(){
      this.showAccount = false
    },
    test() {
      this.showShop = false
    },
    closeMenu(e) {
      if (e.target.className === 'he') console.log(e.target.className);
    },
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
    background: #5b1b0898;
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
    width: 60px;
    height: 60px;
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
    align-items: center;
    margin-right: 36px;
  }

  .coins img {
    width: 50px;
    height: 45px;
  }

  .coins p {
    font-size: 18px;
    font-weight: 800;
    color: white;
    margin-left: 10px;
    color: rgb(255, 196, 0);
  }
</style>
