/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

var WebSocket = require('ws');
var ws = new WebSocket('ws://localhost:2700');

ws.on('error', function(error) {
    console.log('error', error)
})

ws.on('open', function open() {
    ws.send(JSON.stringify({
        cmd: 'set',
        key: 'hello',
        value: 'world'
    }));

    ws.send(JSON.stringify({
        cmd: 'set',
        key: 'list',
        value: [1, 2, 3]
    }));

    ws.send(JSON.stringify({
        cmd: 'del',
        key: 'hello'
    }));

    ws.send(JSON.stringify({
        cmd: 'exist',
        key: 'hello'
    }));

    ws.send(JSON.stringify({
        cmd: 'listpush',
        key: 'list',
        item: 'newitem'
    }));

    ws.send(JSON.stringify({
        cmd: 'get',
        key: 'list',
        item: 'newitem'
    }));

    ws.send(JSON.stringify({
        cmd: 'listremove',
        key: 'list',
        item: 'newitem'
    }));

    ws.send(JSON.stringify({
        cmd: 'batchget',
        keylist: ['hello', 'list']
    }));
});

ws.on('message', function(data, flags) {
    console.log(data)
});

ws.on('close', function close() {
  console.log('disconnected');
});
