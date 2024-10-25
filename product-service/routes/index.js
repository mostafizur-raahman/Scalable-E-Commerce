const { Router } = require("express");
const { createProduct } = require("../controller/createProduct");
const { read } = require("../controller/readProduct");
const { updateProduct } = require("../controller/updateProduct");
const { deleteProduct } = require("../controller/deleteProduct");

const productRoutes = new Router();

productRoutes.post("/product/create", createProduct);
productRoutes.get("/products", read);
productRoutes.patch("/product/update", updateProduct);
productRoutes.patch("/product/delete", deleteProduct);

module.exports = productRoutes;
