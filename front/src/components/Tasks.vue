<template>
    <div class="tasks" v-show="visibility">
        <div class="tasks-head">
                <h2>{{title}}</h2>
                <img @click="showInfo = !showInfo" class="warn" src="../assets/warning.png" alt="">
        </div>
        <info-competence :visibility="showInfo" v-show="showInfo"></info-competence>
        <component :is="selected"></component>
        <button class="cross" @click="close()">
            <img src="../assets/Rectangle 24.png" alt="">
        </button>
    </div>
</template>

<script>
import DayTask from '../components/DayTask.vue'
import InfoCompetence from './InfoCompetence.vue';
export default {
    name: 'Tasks',
    props:['visibility'],
    components: {
        DayTask,
        InfoCompetence
    },
    methods:{
        close(){
                this.$emit('closeDialog');  
            }
    },
    data(){
        return {
            selected: 'DayTask',
            title:'Задачи',
            showInfo: false
        }
    }
}
</script>

<style>
    .tasks {
        border-radius: 6px;
        position: absolute;
        bottom: 95px;
        left: 170px;
        width: 1186px;
        height: 560px;
        background: #FFFCFC;
    }

    .tasks-head {
        position: relative;
        border-radius: 4px 4px 0 0;
        width: 100%;
        height: 56px;
        background: #eb9969de;
        color: white;
        box-sizing: border-box;
    }

    .warn {
        cursor: pointer;
        position: absolute;
        width: 35px;
        top: 8px;
        left: 35px;
    }

    h2 {
        padding-top: 10px;
        margin: 0;
        height: 100%;
        text-align: center;
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

    .task-item{
        width: 200px;
        height: 55px;
        text-align: center;
        border: none;
        background: none;
        color: #944B16;
        font-weight: 600;
        text-transform: capitalize;
    }

    .active {
        color: white;
    }

    .cross img {
        width: 20px;
    }
</style>