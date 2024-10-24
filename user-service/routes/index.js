const { Router } = require("express");
const { registerUser } = require("../controllers/register");

const userRoutes = new Router();

userRoutes.post("/register", registerUser);

module.exports = userRoutes;
