const passport = require("passport");
const JwtStrategy = require("passport-jwt").Strategy;
const ExtractJwt = require("passport-jwt").ExtractJwt;
const LocalStrategy = require("passport-local").Strategy;
const FaceBookTokenStrategy = require("passport-facebook-token");
const config = require("./../config");
const User = require("./../models/user");

passport.use(
  new JwtStrategy(
    {
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken("Authorization"),
      secretOrKey: config.JWT_SECRET,
    },
    async (payLoad, done) => {
      try {
        const user = await User.findUserById(payLoad.sub);

        if (user == undefined) {
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
  "facebookToken",
  new FaceBookTokenStrategy(
    {
      clientID: config.facebook.id,
      clientSecret: config.facebook.secret,
    },
    async (accessToken, refreshToken, profile, done) => {
      try {
        const existingUser = await User.findUserByfId(profile.id);
        if (existingUser == undefined) {
          const newUser = await User.createfUser(
            profile.id,
            profile.emails[0].value
          );
          return done(null, {
            id: newUser,
            email: null,
            password: null,
            f_id: profile.id,
            f_email: profile.emails[0].value,
          });
        }
        return done(null, existingUser);
      } catch (error) {
        done(error, null);
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
