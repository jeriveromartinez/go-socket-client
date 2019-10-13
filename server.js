const app = require('express')();
const http = require('http').createServer(app);
const io = require('socket.io')(http);
const port = 3000;

app.get('/', function (req, res) {
    res.end('Hello');
});

io.on('connection', function (socket) {
    console.log('a user connected');

    setInterval(function () {
        const date = new Date();
        socket.emit('message', `Hello boy!!! ${date}`);
    }, 5 * 1000);

    socket.on('disconnect', function () {
        console.log('user disconnected');
    });
});

http.listen(port, function () {
    console.log(`listening on *:${port}`);
});