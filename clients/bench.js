/*
 * Mumbo - A fast  in-memory key value store
 * Copyright(c) 2016-present @GavinDmello
 * BSD Licensed
 */

var WebSocket = require('ws');
var ws = new WebSocket('ws://localhost:2700/');
var recs = 0
var st, et
var done = false

ws.on('open', function open() {
    console.log('Starting test...')
    startSetTest()
});

ws.on('message', function(data, flags) {
    recs++
    if (recs == 100000) {
        recs = 0
        et = Date.now()
        console.log(100000 / ((et - st) / 1000), 'recs per sec')
        if (!done) {
            done = true
            setTimeout(startGetTest, 1000)
        }
    }
});


function startGetTest() {
    console.log('**** starting gets test **** ')
    for (var i = 0; i <= 100000; i++) {
        if (i == 0) {
            st = Date.now()
        }
        ws.send(JSON.stringify({
            cmd: 'get',
            key: i.toString()
        }));
    }
}

function startSetTest() {
    console.log('**** starting sets test **** ')
    for (var i = 0; i <= 100000; i++) {
        if (i == 0) {
            st = Date.now()
        }
        //console.log('da')
        ws.send(JSON.stringify({
            cmd: 'set',
            key: i.toString(),
            value: i.toString()
        }));
    }
}

