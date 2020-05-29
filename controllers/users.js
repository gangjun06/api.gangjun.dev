const DB = require("./../helpers/db");

module.exports = {
  signUp: async (req, res, next) => {
    const { email, password } = req.value.body;
    const loEmail = email.toLowerCase();

    try {
      const user = await await DB("user")
        .insert({ email: loEmail, password })
      res.json({ user: user[0] });
    } catch (error) {
      if (error.code === "ER_DUP_ENTRY") {
        return res.json({ error: "Email is already in use" });
      }
      return res.json({ error });
    }
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
