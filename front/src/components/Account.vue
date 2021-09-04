<template>
    <div class="account" v-show="visibility">
        <button class="cross" @click="close()">
            <img src="../assets/Rectangle 24.png" alt="">
        </button>
        <div class="wrapper">
            <div class="top-blocks">
                <div class="personal-info">
                    <h1>Мария Авинникова</h1>
                    <p><b>Уровень:</b> Любитель</p>
                    <p class="age"><b>Возраст:</b> 23 года</p>
                    <ProgressBar
                        class="bordered"
                        :options="options"
                        :value="value"
                    />
                    <p class="progress-line">До следующего уровня 30 наград</p>
                </div>
                <div class="awards">
                    <div class="my-awards">
                        <div class="achievements">
                            <img src="../assets/rare.png" alt="">
                            <p class="quantity">32</p>
                        </div>
                        <div class="achievements">
                            <img src="../assets/eye-epic.png" alt="">
                            <p class="quantity">12</p>
                        </div>
                    </div>
                    <div class="levels">
                        <div class="level-icons">
                            <div class="ordinary level"></div>
                            <div class="rare level"></div>
                            <div class="epic level"></div>
                            <div class="legendary level"></div>
                        </div>
                        <div class="progress-level">
                            <img src="../assets/стрела1.png" alt="">
                        </div>
                        <p class="info-level">Выполняйте больше заданий, получайте выше награду, повышайте уровень заработной платы!<br>
                            Совершенству нет предела <br>
                            Получите 30 наград, чтоб перейти к следующему уровню</p>
                    </div>
                </div>
            </div>
            <div class="statistics">
                <h2>Статистика по периодам</h2>
                <div class="graphics">
                    <div class="schema">
                        <img class="graphic" src="../assets/stock-graph 1.png" alt="">
                        <p>Средняя скорость работы за последние 2 смены</p>
                    </div>
                    <div class="static-week">
                        <img class="week" src="../assets/Group 23.png" alt="">
                        <p>Статистика за неделю</p>
                    </div>
                    <div class="circle-bar">
                        <img class="circle" src="../assets/krugovoj-indikator-vypolneniya-progress-bar-1.png" alt="">
                        <p>Производительность труда за месяц 78%</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import ProgressBar from 'vuejs-progress-bar';
export default {
    name: 'Account',
    props: ['visibility'],
    components: {
        ProgressBar
    },
    methods:{
        close(){
            this.$emit('closeDialog');  
        }
    },
    data(){
        return {
            value: 40,
            options: {
                text: {
                    color: '#FFFFFF',
                    shadowEnable: true,
                    shadowColor: '#000000',
                    fontSize: 14,
                    fontFamily: 'Helvetica',
                    dynamicPosition: false,
                    hideText: false
                },
                progress: {
                    color: '#CFC187',
                    backgroundColor: '#CFD0F0'
                },
                layout: {
                    height: 35,
                    width: 140,
                    verticalTextAlign: 61,
                    horizontalTextAlign: 43,
                    zeroOffset: 0,
                    strokeWidth: 30,
                    progressPadding: 0,
                    type: 'line'
                }
            }
        }
    }
}
</script>

<style>
    .account {
        background: #FCFAFA;
        position: absolute;
        width: 100vw;
        height: 600px;
        top: 80px;
        left: 0;
    }

    .top-blocks{
        display: flex;
        height: 300px;
        border-bottom: 2px solid rgba(236, 168, 65, 0.802);
    }

    .personal-info {
        padding: 10px 20px;
        box-sizing: border-box;
        width: 700px;
        height: 100%;
        border-right: 2px solid rgba(236, 168, 65, 0.802);
    }

    h1 {
        color: chocolate;
    }

    h2 {
       color: rgb(181, 85, 16); 
    }

    .age {
        margin-bottom: 35px;
    }

    .awards {
        display: flex;
        width: 100%;
    }

    .my-awards {
        width: 600px;
        display: flex;
        gap: 30px;
        padding: 20px;
    }

    .achievements {
        width: 70px;

    }

    .achievements img {
        width: 100%;
    }

    .quantity {
        text-align: center;
        color: red;
        margin: 0;
    }

    .levels {
        width: 300px;
        padding: 20px 20px 0 20px;
    }

    .level-icons {
        display: flex;
        justify-content: space-around;
    }

    .info-level {
        color: rgb(138, 135, 135);
        padding-left: 20px;
        font-size: 14px;
        margin: 0;
    }

    .progress-level img{
        width: 310px;
    }

    .level {
        border-radius: 50px;
        width: 50px;
        height: 50px;
        background: #e5e5e5;
        opacity: 0.8;
    }

    .statistics {
        height: 100%;
    }
    
    .ordinary {
        background: url('../assets/ord.png') center center/ contain no-repeat;
        width: 50px;
        box-sizing: border-box;
    }

    .rare {
        background: url('../assets/rare.png') center center/ contain no-repeat;
        width: 50px;
        box-sizing: border-box;
    }

    .epic {
        background: url('../assets/epic.png') center center/ contain no-repeat;
        width: 50px;
        box-sizing: border-box;
    }

    .legendary {
        background: url('../assets/leg.png') center center/ contain no-repeat;
        width: 50px;
        box-sizing: border-box;
    }

    ul {
        padding-left: 0;
        padding-top: 10px;
    }

    .progress-line {
        color: rgba(160, 118, 55, 0.802);
    }

    ul li {
        list-style-type: none;
        font-weight: 400;
        padding-bottom: 10px;
        cursor: pointer;
    }

    .graphics {
    margin-top: 20px;
        display: flex;
        align-items: center;
        justify-content: space-around;
    }


    .graphic {
        width: 300px;
        height: 150px;
    }

    .circle-bar p {
        width: 200px;
        margin-right: 25px;
    }

    .circle {
        width: 300px;
    }

    .week {
        width: 300px;
    }

    .week p {
        text-align: center;
    }
</style>