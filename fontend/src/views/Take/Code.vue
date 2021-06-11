<template>
    <div class="home">
        <router-link id="back" to="/take"><el-button icon="el-icon-back" @click="closeMedia()">返回</el-button></router-link>
        <p class="p1">二维码取餐</p>
        <p class="codes">请将二维码放置在扫描处</p>

        <div>
            <div id="loadingMessage">请稍等</div>
            <canvas id="canvas" hidden></canvas>
            <div id="output" hidden>
                <div id="outputMessage"></div>
                <div hidden><b>Data:</b> <span id="outputData"></span></div>
            </div>
        </div>

    </div>
</template>

<script>
    import { Toast } from 'vant';
    import jsQR from "jsqr"
    import {kaisuo, takeqr} from "../../network/outsell"
    export default {
        name: "Code",
        data(){
            return{
                videos : {}
            }
        },
        methods: {
            getMedia(){
                var video = document.createElement("video");
                var canvasElement = document.getElementById("canvas");
                var canvas = canvasElement.getContext("2d");
                var loadingMessage = document.getElementById("loadingMessage");
                var outputContainer = document.getElementById("output");
                var outputMessage = document.getElementById("outputMessage");
                var outputData = document.getElementById("outputData");
                var aa = true;
                function drawLine(begin, end, color) {
                    canvas.beginPath();
                    canvas.moveTo(begin.x, begin.y);
                    canvas.lineTo(end.x, end.y);
                    canvas.lineWidth = 4;
                    canvas.strokeStyle = color;
                    canvas.stroke();
                }

                // Use facingMode: environment to attemt to get the front camera on phones
                navigator.mediaDevices.getUserMedia({ video: { facingMode: "environment" } }).then(function(stream) {
                    // console.log(typeof stream)
                    video.srcObject = stream;
                    video.setAttribute("playsinline", true); // required to tell iOS safari we don't want fullscreen
                    video.play();
                    // TODO 关闭摄像头
                    // if (document.getElementById("back").onclick){
                    //     video.srcObject.getVideoTracks()[0].stop()
                    // }
                    requestAnimationFrame(tick);

                });

                function tick() {
                    loadingMessage.innerText = "⌛ Loading video..."
                    if (video.readyState === video.HAVE_ENOUGH_DATA) {
                        loadingMessage.hidden = true;
                        canvasElement.hidden = false;
                        outputContainer.hidden = false;

                        canvasElement.height = "250";
                        canvasElement.width = "350";
                        canvas.drawImage(video, 0, 0, canvasElement.width, canvasElement.height);
                        var imageData = canvas.getImageData(0, 0, canvasElement.width, canvasElement.height);
                        var code = jsQR(imageData.data, imageData.width, imageData.height, {
                            inversionAttempts: "dontInvert",
                        });

                        if (code) {
                            drawLine(code.location.topLeftCorner, code.location.topRightCorner, "#FF3B58");
                            drawLine(code.location.topRightCorner, code.location.bottomRightCorner, "#FF3B58");
                            drawLine(code.location.bottomRightCorner, code.location.bottomLeftCorner, "#FF3B58");
                            drawLine(code.location.bottomLeftCorner, code.location.topLeftCorner, "#FF3B58");
                            console.log(code.data)
                            //     outputMessage.hidden = true;
                            //     outputData.parentElement.hidden = false;
                            //     outputData.innerText = code.data;
                            takeqr(code.data).then((res) => {
                                if(res.status === 200) {
                                    if (aa) {
                                        kaisuo();
                                        console.log(11234);
                                        Toast.success('取餐成功,请到1号柜前取餐');
                                        aa = false;
                                    }
                                }
                                    else{
                                        Toast.fail('二维码错误或此餐已取');
                                }
                            })
                        } else {
                            outputMessage.hidden = false;
                            outputData.parentElement.hidden = true;
                        }
                    }
                    requestAnimationFrame(tick);
                }
            },

        },
        mounted() {
            this.getMedia()
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
    .p1{
        position: absolute;
        top: 20vh;
        font: 30px Extra large;
    }
    #back{
        position: absolute;
        top: 28vh;
    }
    .codes{
        margin-bottom: 10vh;
        font: 30px Extra large;
    }
</style>
