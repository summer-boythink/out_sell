
<template>
    <div id="num">
            <van-form @submit="onSubmit" class="forms" >
                <router-link to="/take"><el-button icon="el-icon-back" id="back">返回</el-button></router-link>
                <p class="title">数字取餐</p>
                <van-field
                        v-model="codes"
                        name="codes"
                        label="取餐码"
                        placeholder="请输入取餐码"
                        @touchstart.native.stop="show1 = true"
                        autocomplete="off"
                />

            <van-number-keyboard
                    v-model="codes"
                    :show="show1"
                    :maxlength="6"
                    @blur="show1 = false"
            />
                <div style="margin: 16px;" class="submits">
                    <van-button round block type="info" native-type="submit">取餐</van-button>
                </div>
                <p class="ok" v-show="numok">取餐成功，请到智享柜1号取餐</p>
            </van-form>
    </div>
</template>

<script>
    import { Toast } from 'vant';
    import {kaisuo, takenum} from "../../network/outsell"
    export default {
        name: "Num",
        data() {
            return {
                show1: false,
                value: '',
                codes:'',
                numok:false
            };
        },
        methods:{
            onSubmit(values) {
                // console.log('submit', values);
                takenum(values.codes).then((res) => {
                     if(res.status === 200){
                         kaisuo();
                        Toast.success('取餐成功,请到1号柜前取餐');
                    }else{
                        Toast.fail('取餐码错误或此餐已取');
                    }
                })
            },
        },
        mounted(){
            this.numok = false
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
        width: 30vh;
        padding-left: 15vh;
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
