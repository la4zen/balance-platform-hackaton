<template>
    <div class="task-list">
        <div class="task" :class="task.class" v-for="(task, index) of tasks" :key="index">
            <div class="category">{{task.competence}}</div>
            <div class="task-text">
                <button class="start">Приступить к заданию</button>
            </div>
            <div class="check">
                <img src="../assets/Stamp.png" alt="">
                <p>{{task.coins}}</p>
            </div>
            <!-- <input type="checkbox" name="" id=""> -->
        </div>
    </div>
</template>

<script>
export default {
    name: 'DayTask',
    data(){
        return{
            tasks: [
                {
                    competence: 'Компетенция I',
                    class: 'important',
                    coins: 30
                },
                {
                    competence: 'Компетенция II',
                    class: 'urgent',
                    coins: 50
                },
                {
                    competence: 'Компетенция III',
                    class: 'simple',
                    coins: 10
                }
            ]
        }
    },
}
</script>

<style>
    .task-list {
        width: 100%;
    }

    .task{
        margin: 15px 30px;
        padding: 0 50px;
        font-weight: 600;
        color: #1AB192;
        box-sizing: border-box;
        display: flex;
        height: 80px;
        border-radius: 4px;
        border: 2px solid #1AB192;
        background: white;
        align-items: center;
        box-shadow: 4px 4px 5px rgb(0, 0, 0, 0.25);
    }

    .category {
        width: 500px;
    }

    .check{
        display: flex;
        align-items: center;
        margin-left: 80px;
    }

    .task-text {
        color: #000000;
        margin-left: 90px;
        max-width: 700px;
    }

    .important{
        color: red;
        border: 2px solid red;
    }

    .urgent {
        color: purple;
        border: 2px solid purple; 
    }

    .start {
        color: #FF8562;
        background: none;
        border: none;
        font-size: 18px;
    }
</style>