<template>
    <div id="num">
        <van-form @submit="onSubmit" class="forms" >
            <router-link to="/take"><el-button icon="el-icon-back" id="back">返回</el-button></router-link>
            <p class="title">语音取餐</p>
            <p class="submits" v-show="!numok">请报出取餐码...</p>
            <p class="ok" v-show="numok">取餐成功，识别结果为2005,请到智享柜1号取餐</p>
        </van-form>
    </div>
</template>

<script>
    import { Toast } from 'vant';
    import {checkvoice, kaisuo, voice} from "../../network/outsell"
    export default {
        name: "Voice",
        data() {
            return {
                show1: false,
                value: '',
                codes:'',
                numok:false
            };
        },
        methods:{
            onSubmit1(values) {
                // console.log('submit', values);
                checkvoice().then((res) => {
                    if(res.status === 200){
                        kaisuo();
                        this.numok = true;
                        Toast.success('取餐成功,请到1号柜前取餐');
                    }else{
                        Toast.fail('取餐码错误或此餐已取');
                    }
                })
            },
        },
        mounted(){
            this.numok = false;
            this.onSubmit1({codes:11})
        }
    }
</script>

<style scoped>
    #num{
        width: 100%;
        height: 100%;
    }
    .forms{
        width: 70vh;
        padding-top: 30vh;
        margin-left: 55vh;
    }
    .submits{
        padding-left: 16vh;
        padding-top: 13vh;
        font-size: 5vh;
    }
    #back{
        margin-left: -25px;
    }
    .title{
        position: absolute;
        top: 20vh;
        font: 30px Extra large;
        left: 76vh;
    }
    .ok{
        text-align: center;
        font: 25px Extra large;
    }
</style>
