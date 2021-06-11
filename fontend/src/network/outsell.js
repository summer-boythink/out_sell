import {request} from './index'

export function putfood(tele,userid){
    return request({
        url:`putfood?tele=${tele}&userid=${userid}`
    })
}

export function takenum(code){
    return request({
        url:`takenum?code=${code}`
    })
}

export function takeqr(id){
    return request({
        url:`takeqr?id=${id}`
    })
}

export function voice(code) {
    return request({
        url:`voice?code=${code}`
    })
}


//chuangan
export function temp() {
    return request({
        url:`devices/temp`
    })
}

export function weight() {
    return request({
        url:`devices/weight`
    })
}

export function kaisuo() {
    return request({
        url:`devices/kaisuo`
    })
}

export function checkvoice() {
    return request({
        url:`devices/checkvoice`
    })
}

