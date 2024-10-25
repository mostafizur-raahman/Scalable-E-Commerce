const productModel = require("../models/product");
const paginate = require("../utils/paigination");

exports.read = async (req, res, next) => {
    try {
        const page = parseInt(req.query.page) || 1;
        const limit = parseInt(req.query.limit) || 10;

        const products = await productModel.find({
            isDeleted: false,
        });

        const result = paginate(products, page, limit);

        return res.status(200).json({
            message: "Product fetch successfully",
            ...result,
        });
    } catch (error) {
        next(error);
    }
};
