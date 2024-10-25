const { Router } = require("express");
const { registerUser } = require("../controllers/register");
const { userLogin } = require("../controllers/login");

const userRoutes = new Router();

userRoutes.post("/user/register", registerUser);
userRoutes.post("/user/login", userLogin);

module.exports = userRoutes;
