module.exports = {
  signUp: async (req, res, next) => {
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
