var WebSocket = require('ws');
var ws = new WebSocket('ws://localhost:3000/');

ws.on('open', function open() {
    ws.send(JSON.stringify({
        cmd: 'set',
        key: 'hello',
        value: 'world'
    }));

    ws.send(JSON.stringify({
        cmd: 'get',
        key: 'hello'
    }));
});

ws.on('message', function(data, flags) {
    console.log(data)
});
