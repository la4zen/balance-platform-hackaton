<template>
    <footer class="footer-project">
        <div class="container">
            <div class="footer-items">
                <button-item @click="showRating = !showRating">Рейтинг</button-item>
                <rating :visibility="showRating"></rating>
                <button-item>Задачи</button-item>
            </div>
        </div>
    </footer>
</template>

<script>
import ButtonItem from '../components/ButtonItem.vue'
import Rating from '../components/Rating.vue'

export default {
    name: 'FooterProject',
    props: ['visibility'],
    components: {
        ButtonItem,
        Rating
    },
    data(){
        return{
            showRating: false
        }
    }
}
</script>

<style>
    .footer-project{
        position: fixed;
        right: 0;
        left: 0;
        bottom: 0;
        height: 80px;
        background: #1AB192;
    }

    .footer-items {
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
</style>