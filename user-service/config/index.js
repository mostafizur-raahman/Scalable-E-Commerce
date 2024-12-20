require("dotenv").config();

module.exports = {
    port: process.env.PORT || 5000,
    mongo_uri: process.env.MONGO_URI,
    jwt_secret: process.env.JWT_SECRET,
    jwt_ttl: process.env.JWT_TTL,
};
