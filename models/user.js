const DB = require("./../helpers/db");
const bcrypt = require("bcryptjs");

module.exports = {
  db: DB("user"),
  isValidPassword: (password, input) => {
    try {
      return bcrypt.compare(input, password);
    } catch (error) {
      throw new Error(error);
    }
  },
  createUser: async (email, password) => {
    email = email.toLowerCase();
    const salt = await bcrypt.genSalt(10);
    const passwordHash = await bcrypt.hash(password, salt);
    const result = await DB("user").insert({
      email,
      password: passwordHash,
    });
    return result[0];
  },
  createfUser: async (id, email) => {
    const result = await DB("user").insert({
      f_id: id,
      f_email: email,
    });
    return result[0];
  },
  findUserById: async (id) => {
    const result = await DB("user").where("id", id);
    return result[0];
  },
  findUserByfId: async (id) => {
    const result = await DB("user").where("f_id", id);
    return result[0];
  },
  findUserByEmail: async (email) => {
    const result = await DB("user").where("email", email);
    return result[0];
  },
};
