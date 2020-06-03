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

router.route("/signout").get(passportJWT, UsersController.signOut);
// router
//   .route("/oauth/facebook")
//   .post(
//     passport.authenticate("facebookToken", { session: false }),
//     UserController.facebookOAuth
//   );

// router
//   .route("/oauth/link/facebook")
//   .post(
//     passportJwt,
//     passport.authenticate("facebookToken", { session: false }),
//     UserController.linkFacebookOAuth
//   );

// router
//   .route("/oauth/unlink/facebook")
//   .post(passportJwt, UserController.unlinkFacebookOAuth);

router.route("/status").get(passportJwt, UserController.checkAuth);

module.exports = router;
