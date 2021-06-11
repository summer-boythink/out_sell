const router = require('koa-router')()

router.prefix('/devices')
var child_process = require('child_process');
 


router.get('/temp',async (ctx,next) => {

    function exec(){
        return new Promise((resolve,reject) =>{
            var workerProcess = child_process.exec('node ../devices/ds18b20.js ', function (error, stdout, stderr) {
            if (error) {
                console.log(error.stack);
                console.log('Error code: '+error.code);
                console.log('Signal received: '+error.signal);
            }
            console.log('stdout: ' + stdout);
            console.log('stderr: ' + stderr);
            //emitter.emit("temp",stdout)
            // temp1 = stdout
            resolve(stdout)
            }); 
        })
    }

    await exec().then((data) => {
        ctx.body = {
                status:200,
                temp:data
            }
    })
})

router.get('/weight',async (ctx,next) => {
    function exec(){
        return new Promise((resolve,reject) =>{
            var workerProcess = child_process.exec('python ../devices/hx711py/example.py ', function (error, stdout, stderr) {
            if (error) {
                console.log(error.stack);
                console.log('Error code: '+error.code);
                console.log('Signal received: '+error.signal);
            }
            console.log('stdout: ' + stdout);
            console.log('stderr: ' + stderr);
            //emitter.emit("temp",stdout)
            // temp1 = stdout
            resolve(stdout)
            }); 
        })
    }

    await exec().then((data) => {
        ctx.body = {
                status:200,
                weight:data
            }
    })
})

router.get('/kaisuo',async (ctx,next) => {
    function exec(){
        return new Promise((resolve,reject) =>{
            var workerProcess = child_process.exec('node ../devices/led1.js ', function (error, stdout, stderr) {
            if (error) {
                console.log(error.stack);
                console.log('Error code: '+error.code);
                console.log('Signal received: '+error.signal);
            }
            console.log('stdout: ' + stdout);
            console.log('stderr: ' + stderr);
            //emitter.emit("temp",stdout)
            // temp1 = stdout
            resolve(stdout)
            }); 
        })
    }

    await exec().then((data) => {
        ctx.body = {
                status:200
            }
    })
})


module.exports = router