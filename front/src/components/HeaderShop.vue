<template>
        <div class="shop" v-show="visibility">
            <h2 class="shop-head">Магазин</h2>
            <button class="cross" @click="close()">
                <img src="../assets/close.svg" alt="">
            </button>
        </div>
</template>

<script>

// import ClickOutside from 'vue-click-outside'
export default {
    name: 'Shop',
    props:['visibility'],
    data(){
        return{
            cartItems: [],
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
        background: lightblue;
    }

    .shop-head {
        position: relative;
        text-align: center;
        margin: 0;
        border-radius: 4px 4px 0 0;
        width: 100%;
        height: 55px;
        background: #6980d3;
        padding-top: 10px;
        color: white;
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