export interface User {
  username: string;
  fullname: string;
  email: string;
}

export interface Followee {
    username: string;
    createTime: string;
}

export class UserModel {
    static loggedIn: boolean = false;

    static username: string = '';
    private static usernameError: string = '';

    static fullname: string = '';
    private static fullnameError: string = '';

    static email: string = '';
    private static emailError: string = '';

    static followees: Array<Followee> = [];

    static empty() {
      UserModel.username = '';
      UserModel.fullname = '';
      UserModel.password = '';
      UserModel.email = '';
      UserModel.loggedIn = false;
      UserModel.followees = [];
    };

    static isLoggedIn(): boolean {
      return UserModel.loggedIn;
    };

    static setLoggedIn(state: boolean) {
      UserModel.loggedIn = state;
    }

    static setUsername(user: string) {
      // Disallow usernames with whitespaces and URL control/unsafe characters.
      const re = new RegExp(/[\s:;,+$\/\\?!*\'()@=&#]/);

      UserModel.username = user;

      if (UserModel.username.length < 1) {
        UserModel.usernameError =
          'Username must be longer than 0 characters';
      } else if (UserModel.username.length >= 30) {
        UserModel.usernameError =
          'Username must not be longer than 30 characters';
      } else if (re.test(UserModel.username)) {
        UserModel.usernameError =
          'Username must not contain spaces or any of :;,+$\/\\?!*\'()@=&#';
      } else {
        UserModel.usernameError = '';
      }
    };

    static getUsername(): string {
      return UserModel.username;
    };

    static validUsername(): boolean {
      return UserModel.usernameError == '';
    };

    static errorUsername(): string {
      return UserModel.usernameError;
    };

    static setFullname(user: string) {
      UserModel.fullname = user;
      if (!(UserModel.fullname.length < 255)) {
        UserModel.fullnameError = 'Name must be shorter than 255 characters';
      } else if (UserModel.fullname.length < 1) {
        UserModel.fullnameError = 'Name can not be empty';
      } else {
        UserModel.fullnameError = '';
      }
    };

    static getFullname(): string {
      return UserModel.fullname;
    };

    static validFullname(): boolean {
      return UserModel.fullnameError == '';
    };

    static errorFullname(): string {
      return UserModel.fullnameError;
    };

    static setPassword(user: string) {
      UserModel.password = user;
      if (UserModel.password.length >= 4096) {
        UserModel.passwordError = `\
Password must be shorter than 4096 characters, \
that is more than 1 googol^98 in entropy`;
      } else if (UserModel.password.length < 8) {
        UserModel.passwordError = 'Password must be atleast 8 characters';
      } else {
        UserModel.passwordError = '';
      }
    };

    static getPassword(): string {
      return UserModel.password;
    };

    static setEmail(user: string) {
      // The database only cares that the e-mail contains an @ and is longer
      // than 2 characters. But we're using this crazy RegExp here that _should_
      // take care of _most_ cases.
      /* eslint max-len: ["error", { "ignoreRegExpLiterals": true }] */
      const re = new RegExp(/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);

      UserModel.email = user;
      if (!(UserModel.email.length < 255)) {
        UserModel.emailError = 'Email must be shorter than 255 characters';
      } else if (!(re.test(UserModel.email))) {
        UserModel.emailError = 'Must be an valid email';
      } else {
        UserModel.emailError = '';
      }
    };

    static getEmail(): string {
      return UserModel.email;
    };

    static validEmail(): boolean {
      return UserModel.emailError == '';
    };

    static errorEmail(): string {
      return UserModel.emailError;
    };

    static allErrors(): Array<string> {
      const errors = [];

      errors.push(UserModel.errorUsername());
      errors.push(UserModel.errorPassword());
      errors.push(UserModel.errorFullname());
      errors.push(UserModel.errorEmail());

      return errors.filter((e) => e != '');
    };

    static validAll(): boolean {
      return (
        UserModel.validUsername() && UserModel.validEmail() &&
        UserModel.validFullname() && UserModel.validPassword()
      ) && (
        (UserModel.getUsername() != '') && (UserModel.getEmail() != '') &&
        (UserModel.getFullname() != '') && (UserModel.getPassword() != ''));
    };

    static validLogin(): boolean {
      return (
        UserModel.validUsername() && UserModel.validPassword()
      ) && (
        (UserModel.getUsername() != '') && (UserModel.getPassword() != ''));
    };

    // TODO: Should we move this so it isn't inside the user model?
    static password: string = '';
    private static passwordError: string = '';

    static errorPassword(): string {
      return UserModel.passwordError;
    };

    static validPassword(): boolean {
      return UserModel.passwordError == '';
    };

    static emptyFollowees() {
      UserModel.followees = [];
    };
};
