const { ApiResponse } = require("../middleware/ApiResponse");
const userModel = require("../models/user");

exports.registerUser = async (req, res, next) => {
    try {
        const { name, email, password } = req.body;

        const existingUser = await userModel.findOne({
            email: email,
            isDeleted: false,
        });

        if (existingUser) {
            return ApiResponse(
                res,
                409,
                false,
                "User already exists, sign in."
            );
        }

        const newUser = await userModel.create(req.body);

        return res.status(201).json({
            message: "User created successfully",
            user: newUser,
        });
    } catch (error) {
        next(error);
    }
};
