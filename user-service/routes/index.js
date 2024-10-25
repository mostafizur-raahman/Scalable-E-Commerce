const { Router } = require("express");
const { registerUser } = require("../controllers/register");
const { userLogin } = require("../controllers/login");

const userRoutes = new Router();

userRoutes.post("/register", registerUser);
userRoutes.post("/login", userLogin);

module.exports = userRoutes;
