<template>
        <div class="shop" v-show="visibility">
            <h2 class="shop-head">Магазин</h2>
            <button class="cross" @click="close()">
                <img src="../assets/Rectangle 24.png" alt="">
            </button>
            <shop-items></shop-items>
        </div>
</template>

<script>
import ShopItems from './ShopItems.vue';

// import ClickOutside from 'vue-click-outside'
export default {
  components: { ShopItems },
    name: 'Shop',
    props:['visibility'],
    data(){
        return{
        }
    },
    methods:{
        close(){
            this.$emit('closeDialog');  
        }
    }
}
</script>

<style>
    .shop {
        border-radius: 6px;
        position: absolute;
        top: 95px;
        right: 170px;
        width: 1186px;
        height: 560px;
        background: #FFFCFC;
    }

    .shop-head {
        position: relative;
        text-align: center;
        margin: 0;
        border-radius: 4px 4px 0 0;
        width: 100%;
        height: 55px;
        background: #eb9969de;
        padding-top: 10px;
        color: #f0eeee;
        box-sizing: border-box;
    }

    .cross {
        position: absolute;
        width: 55px;
        height: 55px;
        background: tomato;
        border: none;
        top: 0;
        right: 0;
    }

    .cross img {
        width: 20px;
    }
</style>