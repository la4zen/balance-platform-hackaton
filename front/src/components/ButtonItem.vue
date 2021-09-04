<template>
    <button type="button" @click="$emit('click', $event)" class="btn">
        <slot></slot>
    </button>
</template>

<script>
export default {
    name: 'ButtonItem'
}
</script>

<style>
    .btn{
        font-size: 18px;
        height: 45px;
        color: #DCF5FF;
        background: #00abefce;
        border: 2px solid #1487E3;
        border-radius: 6px;
        padding: 0 15px;
        box-sizing: border-box;
        box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
    }
</style>