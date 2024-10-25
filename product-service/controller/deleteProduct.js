const productModel = require("../models/product");

exports.deleteProduct = async (req, res, next) => {
    try {
        const product = await productModel.findOneAndUpdate(
            {
                _id: req.query.id,
                isDeleted: false,
            },
            {
                $set: {
                    isDeleted: true,
                },
            }
        );

        if (!product) {
            return res.status(400).json({ message: "Product not found" });
        }

        return res.status(200).json({
            message: "Product delete successfully",
        });
    } catch (error) {
        next(error);
    }
};
