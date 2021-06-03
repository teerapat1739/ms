
const WebSocket = require('ws');
const redis = require('redis');

require('dotenv').config()
// Configuration: adapt to your environment
const REDIS_SERVER = process.env.REDIS_SERVER ||"redis://localhost:6379";
console.log("ssss", process.env.REDIS_SERVER)
const WEB_SOCKET_PORT = 5555;

// Connect to Redis and subscribe to "app:notifications" channel
var redisClient = redis.createClient(REDIS_SERVER);
redisClient.subscribe('testttt');

// Create & Start the WebSocket server
const server = new WebSocket.Server({ port : WEB_SOCKET_PORT });

// Register event for client connection
server.on('connection', function connection(ws) {

  // broadcast on web socket when receving a Redis PUB/SUB Event
  redisClient.on('message', function(channel, message){
    console.log(message);
    ws.send(message);
  })

});

console.log("WebSocket server started connect to ws://locahost:"+ WEB_SOCKET_PORT);