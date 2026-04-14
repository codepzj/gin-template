# apiwrap

apiwrap 是一个用于包装 API 的工具，简化开发逻辑

响应无论对错，http状态码始终返回200

正确响应：
```json
{
	"code": 200,
	"msg": "success",
	"data": {}
}
```

错误响应：
```json
{
	"code": 500,
	"msg": "error",
	"error": "error message"
}
```

200,400,404,500这些是通用的，并不是http状态码，可以适当添加自己的业务代码

后端请求一般不返回repo相应的详细错误，因为可能容易泄露一些敏感信息

一般对于repo层的错误，service层是这样处理的
- logger打印对应repo的err日志到控制台
- 真正返回msg的时候填写xxx failed就ok了
- 一般不要暴露太多信息给前端，前端一般用相应体中的code来判断是否请求成功