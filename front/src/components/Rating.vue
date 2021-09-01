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
                <button class="personal btnTab" type="button">
                    <p>Персональный</p>
                    <img class="icon" src="../assets/person.png" alt="">
                </button>
                <button class="command btnTab" type="button">
                    <p>Командный</p>
                    <img class="icon" src="../assets/command.png" alt="">
                </button>
            </div>
            <div class="rating-list">
               
                    <div class="rate-line" v-for="(ratingItem,index) of ratingList" :key="index">
                        <!-- <img src="ratingItem.place"> -->
                        <p>{{ratingItem.place}}</p>
                        <img class="photo" :src="img" alt="">
                        <p>{{ratingItem.name}}</p>
                        <p>{{ratingItem.points}}</p>
                    </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    name: 'Rating',
    components:{
        
    },
    props:['visibility'],
    data(){
        return{
            img: 'https://via.placeholder.com/50',
            ratingList:[
                {
                    place: 1,
                    name: 'Lorem',
                    points: 10000

                },
                {
                    place: 2,
                    name: 'Lorem',
                    points: 10000},
                {
                    place: 3,
                    name: 'Lorem',
                    points: 10000
                },
                {
                    place: 4,
                    name: 'Lorem',
                    points: 10000
                },
                {
                    place: 5,
                    name: 'Lorem',
                    points: 10000
                },
                {
                    place: 6,
                    name: 'Lorem',
                    points: 10000},
                    {
                    place: 7,
                    name: 'Lorem',
                    points: 10000},
                    {
                    place: 8,
                    name: 'Lorem',
                    points: 10000},
                    {
                    place: 9,
                    name: 'Lorem',
                    points: 10000},
                    {
                    place: 10,
                    name: 'Lorem',
                    points: 10000},
                
            ]
        }
    }
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
        /* padding: 0 25px; */
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

    .personal {
        background: #FF8562;
        color: #FFFCFC;
        display: flex;
        justify-content: center;
        border-radius: 4px 0 0 4px;
        border-right: 2px solid #C95A0A;
    }

    .command {
        background: #FFCC9A;
        color: #DC7905;
        display: flex;
        justify-content: center;
        border-radius: 0 4px 4px 0;
    }

    .rate-line {
        display: flex;
        justify-content: space-around;
        margin: 0 15px;
        margin-bottom: 10px;
    }

    .rating-list {
        max-height: 360px;
        overflow-y: scroll;
        margin-bottom: 15px;
    }

    .photo {
        width: 50px;
        height: 50px;
        border-radius: 50px;
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