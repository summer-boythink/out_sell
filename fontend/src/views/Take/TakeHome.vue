<template>
  <div class="home">
    <router-link id="back" to="/"><el-button icon="el-icon-back">返回</el-button></router-link>
    <p class="p1">取餐界面</p>
    <router-link to="/take/num"><el-button class="btn" type="primary">数字取餐</el-button></router-link>
    <router-link to="/take/code"><el-button class="btn" type="primary">二维码取餐</el-button></router-link>
    <router-link to="/take/voice"><el-button class="btn" type="primary">语音取餐</el-button></router-link>
  </div>
</template>

<script>

export default {
  name: 'TakeHome',
  components: {

  }
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
  #back{
    position: absolute;
    top: 28vh;
  }
</style>
