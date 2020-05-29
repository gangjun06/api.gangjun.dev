const DB = require("./../helpers/db");
const JWT = require("jsonwebtoken");
const { JWT_SECRET } = require("./../config");

const signToken = (user) => {
  return JWT.sign(
    {
      iss: "gangjun",
      sub: user,
      iat: new Date().getTime(),
      exp: new Date().setDate(new Date().getDate() + 1),
    },
    JWT_SECRET
  );
};

module.exports = {
  signUp: async (req, res, next) => {
    const { email, password } = req.value.body;
    const loEmail = email.toLowerCase();
    let user;

    try {
      user = await await DB("user").insert({ email: loEmail, password });
      user = user[0];
    } catch (error) {
      if (error.code === "ER_DUP_ENTRY") {
        return res.status(403).json({ error: "Email is already in use" });
      }
      return res.json({ error });
    }

    // res.json({ user: user[0] });
    const token = signToken(user);
    return res.status(200).json({ user, token });
  },
  signIn: async (req, res, next) => {
    console.log("SignIn called");
    res.end();
  },
  secret: async (req, res, next) => {
    console.log("secret called");
    res.end();
  },
};
