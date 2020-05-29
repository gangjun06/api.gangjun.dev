const passport = require("passport");
const JwtStrategy = require("passport-jwt").Strategy;
const ExtractJwt = require("passport-jwt").ExtractJwt;
const LocalStrategy = require("passport-local").Strategy;
const { JWT_SECRET } = require("./../config");
const User = require("./../models/user");

passport.use(
  new JwtStrategy(
    {
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken("Authorization"),
      secretOrKey: JWT_SECRET,
    },
    async (payLoad, done) => {
      try {
        const user = await User.findUserById(payLoad.sub);

        if (user == undefined) {
          console.log("undefined");
          return done(null, false);
        }

        done(null, user);
      } catch (error) {
        done(error, false);
      }
    }
  )
);

passport.use(
  new LocalStrategy(
    {
      usernameField: "email",
    },
    async (email, password, done) => {
      try {
        const user = await User.findUserByEmail(email);

        if (user == undefined) {
          return done(null, false);
        }

        const isMatch = await User.isValidPassword(user.password, password);
        if (!isMatch) {
          return done(null, false);
        }

        done(null, user);
      } catch (error) {
        done(error, false);
      }
    }
  )
);
