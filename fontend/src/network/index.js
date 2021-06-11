import axios from 'axios'

export function request(config){

    const instance1 = axios.create({
        // baseURL:'https://service-0gep6mo4-1304695318.gz.apigw.tencentcs.com/release/outsell/'
        baseURL:'http://127.0.0.1:3000/outsell/'
    });

//axios拦截器
    instance1.interceptors.request.use(config => {
        // console.log(config);
        // if (/get/i.test(config.method)) { //判断get请求
        //     config.params  =  config.params || {};
        //     config.params.t = Date.parse(new Date())/1000; //添加时间戳
        // }
        return config;
    },err => {

    });

    instance1.interceptors.response.use(res => {
        return res.data
    },err => {

    });

    return instance1(config)
}
