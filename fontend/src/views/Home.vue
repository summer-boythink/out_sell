<template>
    <div class="home">
        <p class="p1">请选择模式</p>
        <router-link to="/take"><el-button class="btn" type="success">我要取餐</el-button></router-link>
        <router-link to="/put"><el-button class="btn" type="success">我要放餐</el-button></router-link>
    </div>
</template>

<script>
    export default {
        name: "Home"
    }
</script>

<style scoped>
    .home{
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        align-items: center;
        height: 100%;
    }
    .btn{
        width: 300px;
        height: 150px;
    }
    .p1{
        position: absolute;
        top: 20vh;
        font: 30px Extra large;
    }
</style>
