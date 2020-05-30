const User = require("./../models/user");
const JWT = require("jsonwebtoken");
const config = require("./../config");

const signToken = (id) => {
  return JWT.sign(
    {
      iss: "gangjun",
      sub: id,
      iat: new Date().getTime(),
      exp: new Date().setDate(new Date().getDate() + 1),
    },
    config.JWT_SECRET
  );
};

module.exports = {
  signUp: async (req, res, next) => {
    const { email, password } = req.value.body;
    let user;

    try {
      user = await User.createUser(email, password);
    } catch (error) {
      if (error.code === "ER_DUP_ENTRY") {
        return res.status(403).json({ error: "Email is already in use" });
      }
      return res.json({ error });
    }

    // res.json({ user: user[0] });
    const token = signToken(user);
    return res.status(200).json({ token });
  },
  signIn: async (req, res, next) => {
    const token = signToken(req.user.id);
    res.status(200).json({ token });
  },
  facebookOAuth: (req, res, next) => {
    const token = signToken(req.user.f_id);
    res.status(200).json({ token });
  },
  secret: async (req, res, next) => {
    res.status(200).json({ secret: "resource" });
  },
};
