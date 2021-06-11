import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'


//take
import TakeHome from "../views/Take/TakeHome";
import Code from "../views/Take/Code";
import Num from "../views/Take/Num";
import Voice from "../views/Take/Voice";
//put
import PutHome from "../views/Put/PutHome";
import Face from "../views/Put/Face";

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/take',
    name: "Take",
    component: TakeHome
  },
  {
    path: "/take/num",
    name: "Num",
    component: Num
  },
  {
    path:"/take/code",
    name:"Code",
    component:Code
  },
  {
    path:"/take/voice",
    name:"Voice",
    component:Voice
  },
  //PUTS
  {
    path:"/put",
    name:"put",
    component:PutHome
  },
  {
    path:"/put/face",
    name:"face",
    component:Face
  }
];

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes
});

export default router
