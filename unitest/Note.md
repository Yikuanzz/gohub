## Golang Web 分层测试思路

1、repo 数据层：直接访问数据库或者缓存这一层的单元测试。

2、service 逻辑层：针对业务逻辑层的单元测试。

3、API 接口层：http 的接口层的单元测试。

4、其他测试：表格驱动测试、benchmark 测试、测试覆盖率等等。



### 数据层

最重要的就是测试各个数据字段的编写是是否正确，以及 SQL 等查询条件是否能被正常筛选。

- 工具：https://github.com/ory/dockertest



1、用 docker 启动数据源进行测试；

2、通过 ORM 或 导入 SQL 进行数据初始化；

3、测试完单个方法保证测试前后的数据一致性，不影响其他单元测试。



### 逻辑层

这一层描述业务逻辑，主要依赖数据层，因此重要的是对数据层进行 mock。

- 工具：https://github.com/uber-go/mock

- 优化：`//go:generate mockgen -source=./user.go -destination=../mock/user_mock.go -package=mock`



### 接口层

接口层内容可以根据不同的框架进行测试，在接口层则是对逻辑层进行 mock。

1、在实际场景中往往需要进行鉴权，实际可以在前面添加中间件进行处理；

2、对于不同的异常情况，是需要进行测试。



### 其他技能

**打桩测试**：用来模拟某些依赖性的行为，可以直接对某个方法进行 mock ，直接给出想要的返回。

- 工具：https://github.com/agiledragon/gomonkey

```shell
go test -gcflags=all=-l .
```



**压力测试**：对方法进行压力测试，不是对接口的压力测试。

- 工具：Go 提供的 `b *testing.B` 。

```shell
go test -bench=. .
```





**测试覆盖率**：统计单元测试的覆盖率。

- 工具：Go 提供的 `-cover` 参数。

```shell
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```



**表格驱动测试**：将测试数据做成表格的形式，测试代码不同只需增加测试数据。

- 工具：https://github.com/cweill/gotests

