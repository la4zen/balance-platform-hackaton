<template>
    <div class="header-menu" v-show="visibility" @click="close">
        <ul v-for="(menuli, idx) of menulist" :key="idx">
            <li>{{ menuli }}</li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'HeaderMenu',
    props:['visibility'],
    data(){
        return {
            menulist: ['Мой профиль', 'Мои награды', 'Моя статистика']
        }
    },

    methods: {
        close(e) {
            this.$emit('close', e)
        }
    }
}
</script>

<style>
    .header-menu {
        position: absolute;
        top: 75px;
        left: 180px;
        background: #FCFAFA;
        box-sizing: border-box;
        border-radius: 4px;
        box-shadow: 5px 4px 4px rgba(0, 0, 0, 0.1);
        padding: 20px;
    }

    ul {
        padding-left: 0;
    }

    ul li {
        list-style-type: none;
        font-weight: 600;
        cursor: pointer;
    }
</style>