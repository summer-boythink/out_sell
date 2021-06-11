<template>
    <div id="num">
        <van-form @submit="onSubmit" class="forms" >
            <router-link to="/"><el-button icon="el-icon-back" id="back">返回</el-button></router-link>
            <p class="title">存餐界面</p>
            <van-field
                    v-model="tele"
                    name="tele"
                    label="用户手机号"
                    placeholder="请输入用户手机号"
                    @touchstart.native.stop="show1 = true,show2= false"
                    autocomplete="off"
            />
            <van-field
                    v-model="id"
                    name="id"
                    label="智享Id"
                    placeholder="请输入智享Id"
                    @touchstart.native.stop="show2 = true,show1 = false"
                    autocomplete="off"
            />

            <van-number-keyboard
                    v-model="tele"
                    :show="show1"
                    :maxlength="15"
                    @blur="show1 = false"
            />
            <van-number-keyboard
                    v-model="id"
                    :show="show2"
                    :maxlength="15"
                    @blur="show2 = false"
            />

            <div style="margin: 16px;" class="submits">
                <van-button round block type="info" native-type="submit">存餐</van-button>
                <van-button class="btn2" round block type="primary" @click="trap">点击扫脸存餐</van-button>
            </div>
        </van-form>
    </div>
</template>

<script>
    import { Toast } from 'vant';
    import {kaisuo, putfood} from '../../network/outsell'
    export default {
        name: "PutHome",
        data() {
            return {
                show1: false,
                show2: false,
                value: '',
                tele:'',
                id:"",
                menu:""
            };
        },
        methods:{
            onSubmit(values) {
                // console.log('submit', values);
                putfood(values.tele,values.id).then((res) => {
                    console.log(res);
                    if(res.status === 200){
                         Toast.success('1号柜已打开,请放置');
                         kaisuo();
                         this.$store.state.t1 = true;

                    }else{
                        Toast.fail('输入有错误噢');
                    }
                })
            },
            trap(){
                window.location.href = "https://www.tangqihang.top/face/face.html"
            }
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
    .btn2{
        margin-top:20px ;
    }
</style>
