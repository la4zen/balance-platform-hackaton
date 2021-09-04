<template>
    <footer class="footer-project">
        <div class="container">
            <div class="footer-items">
                <button-item @click="showRating = !showRating" ref="dropup">Рейтинг</button-item>
                <rating @vue-click-outside="onClickOutside"  :visibility="showRating" v-show="showRating"></rating>
                <button-item @click="showTasks = !showTasks">Задачи</button-item>
                <tasks :visibility="showTasks" @closeDialog="test"></tasks>
            </div>
        </div>
    </footer>
</template>

<script>
import Tasks from '../components/Tasks.vue'
import ButtonItem from '../components/ButtonItem.vue'
import Rating from '../components/Rating.vue'

export default {
    name: 'FooterProject',
    props: ['visibility'],
    components: {
        ButtonItem,
        Rating,
        Tasks
    },
    data(){
        return{
            showRating: false,
            showTasks: false
        }
    },
    methods: {
        test() {
            console.log('click');
            this.showTasks = false
        },
        onClickOutside(event) {
            console.log('Clicked outside. Event: ', event);
        },
        // closeRate(e) {
        //     if (e.target.className === 'he')
        //     console.log(e.target.className);
        // },
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
        background: #9e6528bd;
    }

    .footer-items {
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
</style>