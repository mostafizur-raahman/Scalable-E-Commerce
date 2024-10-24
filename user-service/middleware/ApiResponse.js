// api response
exports.ApiResponse = (res, statusCode, success, message) => {
    return res.status(statusCode).json({
        success: success,
        message: message,
    });
};
