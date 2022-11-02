const io = require('socket.io-client');
// get port from argv
const port = process.argv[3];
const socket = io.connect('ws://localhost:' + port, { transports: ['websocket'], query: {username: process.argv[2]} });

// get username from console
const username = process.argv[2];

socket.on("new-message", (data) => {
    if (data.success && data.text) {
        console.log(data.body.from + ': ' + data.text)
    }
})

socket.on("connect", () => {
    socket.emit("register", {text: username})
    socket.emit("join", {text: "room"})
})

// get message from console and send it to server
process.stdin.on('data', (data) => {
    socket.emit('new-message', { text: data.toString().trim(), body: {target: "room"}, guid: username });
})