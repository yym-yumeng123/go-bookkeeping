> HTTP 是无状态的, 每次发起请求, 服务器不知道是谁发起的, 每次请求是独立的, 服务器不会记录任何东西

### Cookie

Cookie 是一种存储方式, 为了解决HTTP无状态导致无法跟踪用户信息而出现的

#### Cookie 实现状态维持流程:

1. 客户端向服务端发送请求时，服务端在 `HTTP 响应头中添加一个 Set-Cookie` 字段，这个字段的值就是 Cookie 的值。包含了唯一的会话标识，客户端浏览器会把它存到本地
2. 当客户端再次向服务端发送请求时，都会自动带上这个 Cookie，服务端就可以通过这个 Cookie 来识别客户端的身份

```js
HTTP/1.0 200 OK
Content-type: text/html
Set-Cookie: theme=light
Set-Cookie: sessionToken=abc123; Expires=Wed, 09 Jun 2021 10:18:14 GMT
...
```


### Session

`Session` 也是 Web 应用程序中常用的会话跟踪机制。它是一种在`服务器端存储用户状态信息的机制`，通常用于存储用户的身份认证信息、会话标识符等敏感数据

#### Session登录流程

session 是通过 cookie 实现的

1. 首先客户端登录网站，发送账号密码给服务端。服务端校验用户是否存在
2. 生成一个 `SessionId`, 把登录状态存到`服务端的 session` 中
3. 通过 `Set-Cookie` 把 `SessionId 写入到 Cookie` 中，返回给客户端 `{sessionId: 'xxxxx'}`
4. 此后浏览器再请求，都会自动带上 cookie
5. 服务端会根据 cookie 中的 SessionId 找到对应的 session，从而判断用户是否登录
6. 成功后，返回数据给客户端



> 引用 [阮一峰的JWT](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html), 相当于读写了一遍

### 服务端/客户端

- Client 端
- Server 端

用户认证的流程:

1. client 向 server 发送用户名和密码
2. server 验证通过后, 在当前会话(session) 保存相关数据, 比如用户信息, 登录时间等.
3. server 向 client 返回一个 session_id, 写入用户的 cookie
4. client 随后的每一次请求, 都会通过 Cookie, 将 `session_id` 传回 server.
5. server 收到 session_id, 找到前期保存的数据, 得知 client 端用户的身份

上面的模式在于: 扩展性不好, 单机没问题, 如果是服务器集群，或者是跨域的服务导向架构，就要求 session 数据共享，每台服务器都能够读取 session

举例: A 网站和 B 网站是同一家公司的关联服务。现在要求，用户只要在其中一个网站登录，再访问另一个网站就会自动登录，请问怎么实现？

1.  session 数据持久化，写入数据库或别的持久层。各种服务收到请求后，都向持久层请求数据
2.  服务器不保存 session 数据了, 所有数据都保存在客户端. 每次请求都发回服务器. JWT 就是这种方案的一个代表


### JWT

#### JWT 是什么

`JWT(JSON Web Token)`是一种加密的、可变长度的字符串，用于表示互联网上的用户身份验证和授权。它被广泛应用于应用程序的认证、授权和加密通信。

#### JWT 原理

原理是: 服务器认证以后, 生成一个 JSON 对象, 发回给用户, 类似下面

```json
{
  "name": "张三",
  "role": "admin",
  "time": "2023-11-11"
}
```

以后，用户与服务端通信的时候，都要发回这个 JSON 对象. 服务器完全只靠这个对象认定用户身份, 为了防止用户篡改数据，服务器在生成这个对象的时候，会加上签名

服务器就不保存任何 session 数据了，也就是说，服务器变成无状态了，从而比较容易实现扩展

#### JWT 的数据结构

是一个很长的字符串, 中间用(.)分隔成三个部分. 内部没有换行, 下面是为了方便展示

1. Header(头部)
2. Payload(负载)
3. Signature(签名)

```bash
Header.Payload.Signature
# Header
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
# payload
eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
# Signature
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

**Header**

Header 部分是一个 JSON 对象, 描述 JWT 元数据

```json
{
  "alg": "HS256", // alg 签名的算法, 默认是 HMSC SHA256(写成 HS256)
  "typ": "jwt" // typ 表示这个 token 的类型, JWT 令牌统一写 JWT
}
```

**Payload**

`Payload` 部分也是一个 JSON 对象，用来存放实际需要传递的数据。JWT 规定了 7 个官方字段，供选用

- iss(issuer) 签发人
- exp(expiration time) 过期时间
- sub(subject) 主题
- aud(audience) 受众
- nbf(Not Before) 生效时间
- iat(Issued At) 签发时间
- jti(JWT ID) 编号

```json
// 还可以定义私有字段
{
  "sub": "121212",
  "name": "yym",
  "admin": true
}
```

**Signature**

Signature 部分是对前两部分的签名，防止数据篡改。

首先, 需要指定一个秘钥(secret). 这个秘钥只有服务器知道, 不能泄露给用户, 然后，使用 Header 里面指定的签名算法（默认是 HMAC SHA256），按照下面的公式产生签名

```js
HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)
```

算出签名以后，把 Header、Payload、Signature 三个部分拼成一个字符串，每个部分之间用"点"（.）分隔，就可以返回给用户

**Base64URL**

前面提到，Header 和 Payload 串型化的算法是 Base64URL。这个算法跟 Base64 算法基本类似，但有一些小的不同。

JWT 作为一个令牌（token），有些场合可能会放到 URL（比如 api.example.com/?token=xxx）。Base64 有三个字符+、/和=，在 URL 里面有特殊含义，所以要被替换掉：=被省略、+替换成-，/替换成_ 。这就是 `Base64URL` 算法


#### JWT的使用方式

客户端收到服务器返回的 JWT，可以储存在 Cookie 里面，也可以储存在 localStorage

此后，客户端每次与服务器通信，都要带上这个 JWT。你可以把它放在 Cookie 里面自动发送，但是这样不能跨域，所以更好的做法是放在 HTTP 请求的头信息 `Authorization` 字段里面

```
Authorization: Bearer <token>
```

另一种做法是，跨域的时候，JWT 就放在 POST 请求的数据体里面

#### JWT 特点

1. JWT默认不加密
2. JWT不加密的情况下, 不能把秘密数据写入 JWT
3. JWT不仅可以用于认证, 也可以用于交换信息,有效使用 JWT，可以降低服务器查询数据库的次数
4. JWT 的最大缺点是，由于服务器不保存 session 状态，因此无法在使用过程中废止某个 token，或者更改 token 的权限。也就是说，一旦 JWT 签发了，在到期之前就会始终有效，除非服务器部署额外的逻辑。
5. JWT 本身包含了认证信息，一旦泄露，任何人都可以获得该令牌的所有权限。为了减少盗用，JWT 的有效期应该设置得比较短。
6. 为了减少盗用，JWT 不应该使用 HTTP 协议明码传输，要使用 HTTPS 协议传输

### JWT 与 Refresh Token

1. 客户端在某一刻发请求给服务端登录, 服务端会返回一个 token(access_token), 客户端在后面的请求, 每次都在 header 携带 token(认证用户身份).
2. 假如设置2个小时后, JWT过期, 用户只能重新登录获取 token, 体验比较差, 所以如何解决?
3. 我们在原来token(access_token) 的基础上, 添加一个 `refresh_token` 用来刷新 token, 服务端返回新的 `token(acces_token)` 和 `refresh_token`
4. `access_token` 用来做登录鉴权, 如果token 有效访问接口, 如果无效, 重新登录, 这时校验`refresh_token`, token 有效, 返回新的 `access_token` 
5. access_token设置30分钟过期, refresh_token 设置 7天过期
   1. 7天内, access_token 过期了, 可以用 refresh_token 刷新, 拿到新的 token
   2. 不超过7天未访问系统, 一直是登录状态, 可以无限续签
   3. 超过7天, refresh_token 过期了, 需要重新登录