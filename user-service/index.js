const express = require("express");
const http = require("http");
const { connectDB } = require("./config/db");
const config = require("./config");
const app = express();

const server = http.createServer(app);

// middleware
app.use(express.json());

// database connect
connectDB();

server.listen(config.port, () => [
    console.log(`user.service running on port ${config.port}`),
]);
