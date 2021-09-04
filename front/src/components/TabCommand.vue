<template>
    <div class="rating-list">
        <div class="rate-line" v-for="(ratingItem,index) of ratingList" :key="index">
            <div class="rating-item">
                <p>{{index + 1}}</p>
            </div>
            <div class="rating-item">
                <img class="photo" :src="img" alt="">
            </div>
            <div class="name">
                <p>{{ratingItem.name}}</p>
            </div>
            <div class="rating-item">
                <p>{{ratingItem.points}}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Command',
    data(){
        return{
            img: 'https://via.placeholder.com/50',
            ratingList:[
                {
                    name: 'Ящерицы',
                    points: 9000

                },
                {
                    name: 'Анонимусы',
                    points: 9000},
                {
                    name: 'Лягушки',
                    points: 9000
                },
                {
                    name: 'РикАнутые',
                    points: 9000
                },
                {
                    name: 'Котики',
                    points: 9000
                },
                {
                    name: 'ПростоТа',
                    points: 9000
                },
                {
                    name: 'ИСИТко',
                    points: 9000
                },
                
            ]
        }
    }
}
</script>

<style>
    .rating-list {
        max-height: 360px;
        overflow-y: scroll;
        margin-bottom: 15px;
    }

    .command {
        background: #FF8562;
        color: #DCF5FF;
        display: flex;
        justify-content: center;
        border-radius: 0 4px 4px 0;
    }
</style>