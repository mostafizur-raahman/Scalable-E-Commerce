const { ApiResponse } = require("../middleware/ApiResponse");
const userModel = require("../models/user");

exports.userLogin = async (req, res, next) => {
    try {
        const { email, password } = req.body;

        const user = await userModel.findOne({
            email,
            isDeleted: false,
        });

        if (!user) {
            return ApiResponse(res, 404, false, "User not found");
        }

        if (!email || !password) {
            return ApiResponse(
                res,
                404,
                false,
                "Email and password are required"
            );
        }
    } catch (error) {
        next(error);
    }
};
