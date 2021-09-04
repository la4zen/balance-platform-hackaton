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
                <p class="point">{{ratingItem.points}}</p>
            </div>
            </div>
        </div>
</template>

<script>
export default {
    name: 'Personal',
    data(){
        return{
            img: 'https://via.placeholder.com/50',
            ratingList:[
                {
                    name: 'Андрей',
                    points: 10000
                },
                {
                    name: 'Дарья',
                    points: 10000},
                {
                    name: 'Мария',
                    points: 10000
                },
                {
                    name: 'Олег',
                    points: 10000
                },
                {
                    name: 'Филипп',
                    points: 10000
                },
                {
                    name: 'Кристина',
                    points: 10000
                },
                {
                    name: 'Максим',
                    points: 10000
                },
                {
                    name: 'Дмитрий',
                    points: 10000
                },
                {
                    name: 'Вячеслав',
                    points: 10000
                },
                {
                    name: 'Анна',
                    points: 10000
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

    .personal {
        background: #FF8562;
        color: #FFFCFC;
        display: flex;
        justify-content: center;
        border-radius: 4px 0 0 4px;
        border-right: 2px solid #C95A0A;
    }

    .rating-item {
        width: 60px;
        text-align: center;
    }

    .name {
        width: 100px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
    }

    .photo {
        width: 50px;
        height: 50px;
        border-radius: 50px;
        margin-right: 10px;
    }
</style>