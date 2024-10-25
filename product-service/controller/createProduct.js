const productModel = require("../models/product");

exports.createProduct = async (req, res, next) => {
    try {
        const newProduct = await productModel.create(req.body);

        return res.status(201).json({
            message: "Product create succesfully",
            doc: newProduct,
        });
    } catch (error) {
        next(error);
    }
};
