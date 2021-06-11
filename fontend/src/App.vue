<template>
  <div id="app">
    <p class="temp">现在的温度是{{temp}}℃</p>
<!--    <p class="temp">现在的重量是{{weight}}g</p>-->
    <router-view/>

      <img id="guangao1" src="http://imgs-1304695318.cos.ap-shanghai.myqcloud.com/1621757566516.jpg" alt="">
      <img id="guangao2" src="http://imgs-1304695318.cos.ap-shanghai.myqcloud.com/1621757566660.jpg" alt="">

  </div>
</template>
<script>
  import {temp, weight} from "./network/outsell";

  export default {
    data(){
      return{
        temp:25.2,
        weight
      }
    },
    mounted() {
      document.oncontextmenu = function () {
        return false
      };
      setInterval(() => {
        temp().then((res) => {
        console.log(res);
        if(res.status === 200){
          this.temp = res.temp
          // if(this.$store.state.t1 === true) {
          //   if(this.temp < 55){
          //     this.temp += 2.2
          //   }
          // }
        }
      })},2000);

      // setInterval(() => {weight().then((res) => {
      //   console.log(res);
      //   if(res.status === 200){
      //     this.weight = res.weight
      //   }
      // })},2000);
    }
  }
</script>
<style>
  *{
    margin: 0 auto
  }
  #app{
    width: 100%;
    height: 92%;
    position: absolute;
    user-select: none;
  }
  .temp{
    text-align: center;
    position: relative;
    /* bottom: -20px; */
    top: 12vh;
    color: red;
  }
  #guangao1{
    width: 50%;
    position: fixed;
    bottom: 0;
    height: 100px;
    /*float: left;*/
  }
  #guangao2{
    width: 50%;
    position: fixed;
    bottom: 0;
    height: 100px;
    right: 0;
  }
</style>
