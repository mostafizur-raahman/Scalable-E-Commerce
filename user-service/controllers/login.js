const { ApiResponse } = require("../middleware/ApiResponse");
const userModel = require("../models/user");
const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");
const config = require("../config");

exports.userLogin = async (req, res, next) => {
    try {
        const { email, password } = req.body;

        const user = await userModel.findOne({
            email,
            isDeleted: false,
        });

        if (!user) {
            ApiResponse(res, 404, false, "User not found");
        }

        if (!email || !password) {
            ApiResponse(res, 404, false, "Email and password are required");
        }

        const isValidPassword = await bcrypt.compare(password, user.password);

        if (!isValidPassword) {
            ApiResponse(res, 400, false, "Invalid credentials");
        }

        const token = jwt.sign(
            { _id: user._id, email: user.email },
            config.jwt_secret,
            { expiresIn: config.jwt_ttl }
        );

        return res.status(200).json({
            token,
        });
    } catch (error) {
        next(error);
    }
};
