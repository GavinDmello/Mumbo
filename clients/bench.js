var WebSocket = require('ws');
var ws = new WebSocket('ws://localhost:3000/');
v = 0
ws.on('open', function open() {

    for (var i = 0; i<=100000; i++) {
        if (i == 0) {
            console.log(Date.now())
        }
        ws.send(JSON.stringify({
            cmd: 'set',
            key: i.toString(),
            value: i.toString()
        }));


    }
});

ws.on('message', function(data, flags) {
    v++
    if (v == 100000) {
        console.log(Date.now())
    }
});
