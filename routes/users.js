const router = require("express-promise-router")();
const { validateBody, schemas } = require("./../helpers/routeHelpers");
const UserController = require("../controllers/users");
const passport = require("passport");
const passportConf = require("./../helpers/passport");

const passportSingIn = passport.authenticate("local", { session: false });
const passportJwt = passport.authenticate("jwt", { session: false });

router
  .route("/signup")
  .post(validateBody(schemas.authSchema), UserController.signUp);

router
  .route("/signin")
  .post(
    validateBody(schemas.authSchema),
    passportSingIn,
    UserController.signIn
  );

router.route("/secret").get(passportJwt, UserController.secret);

module.exports = router;
