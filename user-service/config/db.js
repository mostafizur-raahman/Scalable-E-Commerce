const mongoose = require("mongoose");
const config = require("../config");

exports.connectDB = async () => {
    try {
        // await mongoose.connect(config.mongo_uri);

        console.log("User service connect to database succesfully...");
    } catch (error) {
        console.log("Failed to connect database user.service");
        throw error;
    }
};
