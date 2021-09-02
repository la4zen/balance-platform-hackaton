<template>
    <div class="rating-list">
        <div class="rate-line" v-for="(ratingItem,index) of ratingList" :key="index">
            <p>{{index + 1}}</p>
            <img class="photo" :src="img" alt="">
            <p>{{ratingItem.name}}</p>
            <p>{{ratingItem.points}}</p>
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
                    name: 'fggjkf',
                    points: 9000

                },
                {
                    name: 'fghgf',
                    points: 9000},
                {
                    name: 'ytrew',
                    points: 9000
                },
                {
                    name: 'nbvcx',
                    points: 9000
                },
                {
                    name: 'jhgfd',
                    points: 9000
                },
                {
                    name: 'gfds',
                    points: 9000
                },
                {
                    name: 'iu7y6tre',
                    points: 9000
                },
                {
                    name: 'kjuytr',
                    points: 9000
                },
                {
                    name: ',mnbvc',
                    points: 9000
                },
                {
                    name: 'trfdsa',
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
        background: #FFCC9A;
        color: #DC7905;
        display: flex;
        justify-content: center;
        border-radius: 0 4px 4px 0;
    }
</style>