const io = require('socket.io-client');
const socket = io.connect('ws://localhost:9999', { transports: ['websocket'] });

socket.emit("echo", { text: "Hello World"}, (data) => {
    console.log(data);
});

socket.on("hello", (data) => {
    console.log(data);
})