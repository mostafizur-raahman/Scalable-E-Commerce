const productModel = require("../models/product");

exports.updateProduct = async (req, res, next) => {
    try {
        const product = await productModel.findOneAndUpdate(
            {
                _id: req.query.id,
                isDeleted: false,
            },
            {
                ...req.body,
            },
            {
                new: true,
            }
        );

        if (!product) {
            return res
                .status(400)
                .json({ message: "Product not found or slug already in use" });
        }

        return res.status(200).json({
            message: "Product updated successfully",
            data: product,
        });
    } catch (error) {
        next(error);
    }
};
