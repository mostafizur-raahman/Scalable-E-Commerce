const mongoose = require("mongoose");
const config = require("../config");

exports.connectDB = async () => {
    try {
        await mongoose.connect(config.mongo_uri);

        console.log("Product service connect to database succesfully...");
    } catch (error) {
        console.log("Failed to connect database product.service");
        throw error;
    }
};
