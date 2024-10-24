const express = require("express");
const http = require("http");
const { connectDB } = require("./config/db");
const config = require("./config");
const userRoutes = require("./routes");
const app = express();

const server = http.createServer(app);

// middleware
app.use(express.json());

// database connect
connectDB()
    .then(() => {
        server.listen(config.port, () => {
            console.log(`user.service running on port ${config.port}`);
        });
    })
    .catch((err) => {
        console.error("Failed to connect to the database user.service", err);
        process.exit(1);
    });

// routes
app.use("/api/v1/user", userRoutes);

// 404 Not Found Middleware
app.use((req, res, next) => {
    res.status(404).json({
        success: false,
        message: "Resource not found",
    });
});

// Global Error Handling Middleware
app.use((err, req, res, next) => {
    console.error(err.stack);

    const statusCode = err.status || 500;
    res.status(statusCode).json({
        success: false,
        message: err.message || "Internal Server Error",
    });
});
