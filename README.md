# Mumbo

A simple key value store written in Golang.

## Download
You can check out the latest releases and download the binary.
<a href="https://github.com/GavinDmello/Mumbo/releases">https://github.com/GavinDmello/Mumbo/releases</a>

## Operations list

 - Get
 - Set
 - Set with TTL
 - Exist
 - Delete
 - BatchGet
 - ListPush
 - ListRemove

## Configuration
The config file needs to be a json file. You can place this file in the `/etc/mumbo-conf.json`
folder. If the file is not present, a default configuration will be extended.
The following options are honoured as of now :-

- `gcInterval` The intervals after which garbage collections will happen of dead keys
The value needs to be be given in milliseconds.
- `persistence` If you want data to be persisted.
- `diskWriteInterval` The interval after which the data will be dumped to the disk. Should be given in milliseconds.
- `port` The port on which you want the server to run

Example config in `/etc/mumbo-conf.json`
```
{
    "gcInterval" : 100,
    "persistence" : true,
    "diskWriteInterval" : 300000,
    "port" : 2700

}
```

## Benchmarks
This was done on a box with 8 GB RAM and i5 processor
This benchmark was done on the same box without persistence. If done on different boxes the network might slow you down.

 -  Reads per sec =  25,873
 -  Writes per sec = 26,427

## License
BSD
