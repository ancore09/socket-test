const io = require('socket.io-client');
// get port from argv
const port = process.argv[3];
const socket = io.connect('ws://localhost:' + port, { transports: ['websocket'] });

socket.on("new-message", (data) => {
    if (data.success && data.text) {
        console.log(data.text)
    }
})

// get username from console
const username = process.argv[2];

// get message from console and send it to server
process.stdin.on('data', (data) => {
    socket.emit('new-message', { text: username + ': ' + data.toString().trim() });
})