<template>
    <div class="info" v-show="visibility">
        <p></p>
        <div class="competence" v-for="(competence,index) of competences" :key="index">
            <h4>{{ competence.name }}<br>
            <span>({{ competence.forWho}})</span></h4>
            <ul>
                <li class="info-subitem" v-for="item of competence.items" :key="item">{{ item }}</li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
    name: 'InfoCompetence',
    props: ['visibility'],
    data(){
        return {
            competences: [
                {
                    name: 'Компетенция I',
                    forWho: 'мастер',
                    items: ['Иностранный паспорт', 'Трудовая книжка', 'Справки МВД', 'Справки ПФР']
                },
                {
                    name: 'Компетенция I',
                    forWho: 'любитель',
                    items: ['Анкета-заявка из банка', 'Cправка 2-НДФЛ', 'Cправка по форме банка']
                },
                {
                    name: 'Компетенция I',
                    forWho: 'новичок',
                    items: ['Паспорт РФ', 'ИНН', 'Снилс', 'Водительское удостоверение']
                }
            ]
        }
    }
}
</script>

<style>
    .info {
        display: flex;
        position: absolute;
        top: 55px;
        left: -30px;
        width: 600px;
        padding-top: 20px;
        height: 200px;
        border-radius: 10px;
        background: white;
        box-shadow: 4px 4px 5px rgb(0, 0, 0, 0.25);
    }

    ul {
        margin: 0;
    }

    h4 {
        color: #5F4A2F;
        font-size: 14px;
        margin: 0;
        margin-bottom: 10px;
    }

    .competence{
    padding-left: 35px;
    width: 33.333%;
    }

    .info-subitem {
        color: #87B2C5;
        font-size: 14px;
        font-weight: 400;
    }

    span {
      color: #da9336;  
    }
</style>