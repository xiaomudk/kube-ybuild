declare namespace API {
  /** 登录参数 */
  type LoginParams = {
    password: string;
    username: string;
  };

  /** 登录成功结果 */
  type LoginResult = {
    token: string;
  };
}
