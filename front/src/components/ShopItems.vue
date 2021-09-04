<template>
    <div class="shop-list">
            <div class="shop-item" v-for="(shopItem,index) of shopItems" :key="index">
                <img class="product-image" :src="shopItem.img" alt="">
                <div class="desc">
                    <h4 class="item-name">{{shopItem.name}}</h4>
                    <button class="points-price">
                        <img src="../assets/Stamp.png" alt="">
                        <p>{{shopItem.price}}</p>
                    </button>
                </div>
            </div>
    </div>
</template>

<script>
export default {
    name: 'ShopItems',
    data(){
        return{
            shopItems: [
                {
                    name: 'Футболка белая',
                    price: 10000,
                    img: require('../assets/white-shirt.png')
                },
                {
                    name: 'Футболка черная',
                    price: 10000,
                    img: require('../assets/black-shirt.png')
                },
                {
                    name: 'Чашка',
                    price: 3000,
                    img: require('../assets/13.png')
                },
                {
                    name: 'Термокружка',
                    price: 4000,
                    img: require('../assets/5ув.png')
                },
                {
                    name: 'Блокнот',
                    price: 2000,
                    img: require('../assets/blocknote.png')
                },
                {
                    name: 'Свитшот',
                    price: 13000,
                    img: require('../assets/jacket.png')
                },

            ],
        }
    }
}
</script>

<style>

    .shop-list {
        margin: 15px 0;
        padding: 20px;
        display: flex;
        flex-wrap: wrap;
        max-height: 450px;
        gap: 30px;
        overflow-x: auto;
    }

    .shop-item {
        cursor: pointer;
        border-radius: 5px;
        box-shadow: -2px 2px 10px rgba(0, 0, 0, 0.1);
    }

    .shop-item:hover {
        transform: scale(1.05, 1.05);
        transition: transform 0.3s ease-in-out;
        }

    .product-image {
        width: 350px;
        height: 200px;
    }

    .desc {
        background: #F9D8CF;
        text-align: center;
        padding-top: 20px;
        padding-bottom: 15px;
    }

    .points-price {
        margin: 0 auto;
        border: none;
        padding: 5px 10px;
        border-radius: 8px;
        background: #f19379;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .points-price img{
        width: 30px;
        height: 30px;
    }

    .points-price p{
        color: rgb(245, 235, 225);
        font-size: 18px;
        font-weight: 600;
        margin: 0;
    }
</style>


