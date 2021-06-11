const router = require('koa-router')()
const db = require('../tool/db')
const telepost = require('../tool/tele')
router.prefix('/outsell')

function exec(yuju){
    return new Promise((res,rej) => {
        db.query(yuju,function (err,data){
            if(err){
                console.log(err);
            }else{
                res(data)
            }
        })
    })
}

let num = 2000

// http://localhost:3000/outsell/putfood?tele=15268119204&userid=1
router.get('/putfood',async (ctx,next) => {
    // ctx.body = JSON.stringify(ctx.query)
    let {tele,userid} = ctx.query;
    let code =await telepost(tele,num)
    num++
    let result = await exec(`insert into food (name,imgurl,remark,userid,code,status) values ("红烧肉","https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg2-cloud.itouchtv.cn%2Fupload%2F20181031%2FYkznQdzYGJ1540955556.png&refer=http%3A%2F%2Fimg2-cloud.itouchtv.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1619431323&t=d51bcc1ff0706df75856e5fcd339e1ae","无",${userid},${num-1},0)`)
    // console.log(result.affectedRows);
    if(result.affectedRows>0 && code.return_code == '00000'){
        ctx.body = JSON.stringify({message:"ok",status:200})
    }else{
        ctx.body = JSON.stringify({message:"fail",status:404})
    }
    
})

// http://localhost:3000/outsell/takenum?code=1999
router.get('/takenum',async (ctx,next) => {
    let {code} = ctx.query;
    let result = await exec(`UPDATE food set status=1 WHERE code=${code}`)
    console.log(result);
    if(result.changedRows>0){
        ctx.body = JSON.stringify({message:"ok",status:200})
    }else{
        ctx.body = JSON.stringify({message:"fail",status:404})
    }
})

// http://localhost:3000/outsell/takenum?id=1
router.get('/takeqr',async (ctx,next) => {
    let {id} = ctx.query;
    let result = await exec(`UPDATE food set status=1 WHERE id=${id}`)
    // console.log(result.affectedRows);
    if(result.affectedRows>0){
        ctx.body = JSON.stringify({message:"ok",status:200})
    }else{
        ctx.body = JSON.stringify({message:"fail",status:404})
    }
})

module.exports = router