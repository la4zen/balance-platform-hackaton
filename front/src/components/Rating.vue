<template>
    <div class="rating" v-show="visibility">
        <div class="rating-head">
            <h3>Топ-10</h3>
            <img src="../assets/Star.png" alt="">
        </div>
        <div class="rating-content">
            <div class="rating-search">
                <button class="search-button">
                    <img class="loupe" src="../assets/loupe.svg" alt="">
                </button>
                <form action="#">
                    <input placeholder="Найти" type="text" name="" id="">
                </form>
            </div>
            <div class="rating-tabs">
                <button @click="test(tab)" :class="[tab.class, {active: selected === tab}]" class="btnTab" v-for="(tab,index) in tabs" :key="index">
                    <p>{{tab.text}}</p>
                    <img class="icon" :src="tab.img" alt="">
                </button>
            </div>
            <component :is="selected"></component>
        </div>
    </div>
</template>

<script>
import Personal from '../components/TabPersonal.vue'
import Command from '../components/TabCommand.vue'
export default {
    name: 'Rating',
    components:{
        Personal,
        Command
    },
    props:['visibility','img'],
    data(){
        return{
            selected: 'Personal',
            tabs:[
                {
                    text: 'Персональный',
                    class: 'personal',
                    img: require('../assets/person.png')
                },
                {
                    text: 'Командный',
                    class: 'command',
                    img: require('../assets/command.png')
                }
            ],
        }
    },

    methods: {
        test(tab) {
            if (tab.class === 'personal') this.selected = 'Personal';
            else this.selected = 'Command';
        }
    }
    // computed: {
    //       currentTabComponent: function() {
    //         return "tab-" + this.currentTab.toLowerCase();
    //       }
    //     }
}
</script>

<style>
    .rating{
        position: absolute;
        border-radius:0 0 4px 4px;
        bottom: 70px;
        width: 440px;
        background: #FFFCFC;
        box-shadow: 8px 8px 8px rgb(0, 0, 0, 0.12);
    }

    .rating-head {
        display: flex;
        justify-content: center;
        color: #FD5525;
        margin: 0;
        font-weight: 800;
        font-size: 20px;
        background: #5ED3BB;
        box-shadow: 0px 4px 4px rgb(0, 0, 0, 0.12);
    }

    .rating-content{
        padding: 0 25px;
    }

    .rating-search {
        border: 1px solid #E7E2E2;
        padding: 10px 0;
        display: flex; 
        margin: 10px 0;
    }

    form {
        box-sizing: border-box;
        text-align: center;
    }

    input {
        outline: none;
        width: 100%;
        margin: 0 auto;
        padding: 5px;
        border: none;
        box-sizing: border-box;
    }

    input::placeholder{
        color: #E1D7D7;
    }

    .btnTab{
        width: 100%;
        border: none;
        display: flex;
        align-items: center;
    }

    .rating-tabs {
        width: 100%;
        display: flex;
        justify-content: center;
        margin-bottom: 15px;
    }

    .rate-line {
        display: flex;
        justify-content: space-between;
        margin: 0 15px;
        margin-bottom: 10px;
    }

    .icon {
        display: block;
        margin-left: 10px;
    }

    .search-button{
        border: none;
        background: none;
    }

    .loupe {
        width: 20px;
        margin-left: 20px;
    }
</style>