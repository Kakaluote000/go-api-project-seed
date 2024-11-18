# go-api-project-seed

`go-api-project-seed` 是一个开箱即用的 Golang 项目模板，旨在帮助开发者快速搭建现代化的 RESTful API 服务。项目集成了主流框架和工具，提供了强大的代码生成能力和通用的 CURD 模板，适用于生产环境，能够显著提高开发效率。

## 核心特性

### 1. 框架和库集成
- **Gin**: 高性能的 HTTP Web 框架，用于路由和中间件管理。
- **GORM**: 功能强大的 ORM 框架，支持与 MySQL 的高效交互。
- **go-redis**: 轻量级 Redis 客户端，支持分布式锁（Redsync）。
- **JWT-go**: 用于身份验证的 JSON Web Token 集成。
- **Logrus**: 灵活强大的日志工具，支持多种日志格式和输出。
- **Swagger**: 自动生成 API 文档，提升文档维护效率。
- **Prometheus**: 提供监控和指标采集能力。
- **Viper**: 配置文件管理工具，支持多种配置格式（如 YAML）。

### 2. 生产级功能
- 支持分页查询、权限控制等通用业务需求。
- 支持基于 GORM 的代码生成工具，可自动生成模型、控制器、服务和仓储层代码。
- 高度可扩展的 CURD 接口，自动处理路由注册与逻辑实现。

### 3. 代码生成工具
- 用户可指定 MySQL 数据库及表，自动生成完整的代码结构。
- 模板文件（如 `model.tmpl`, `controller.tmpl`）高度抽象，易于扩展和定制化。
- 生成代码后，自动格式化（`gofmt`）和修复引用（`goimports`）。

### 4. 开发者友好
- 提供统一的配置管理和环境隔离。
- 集成 Swagger，API 文档生成与接口调试无缝对接。
- 提供全面的注释与文档，适合团队协作与长期维护。

## 项目结构

```plaintext
go-api-project-seed/
├── cmd/                    # 程序入口目录
│   └── main.go             # 主程序文件
├── configs/                # 配置文件目录
│   └── config.yaml         # 项目主配置文件
├── docs/                   # Swagger 文档目录
├── internal/               # 内部业务逻辑目录
│   ├── controller/         # 控制器层
│   ├── service/            # 服务层
│   ├── repository/         # 数据仓储层
│   └── model/              # 数据模型层
├── templates/              # 模板文件目录（代码生成工具使用）
│   ├── model.tmpl          # 数据模型模板
│   ├── controller.tmpl     # 控制器模板
│   ├── service.tmpl        # 服务模板
│   └── repository.tmpl     # 数据仓储模板
├── tools/                  # 工具集
│   └── generate.go         # 自动代码生成工具
├── go.mod                  # Go 模块依赖
└── README.md               # 项目说明文档
```

## 使用方法

### 1. 克隆项目
```bash
git clone https://github.com/Kakaluote000/go-api-project-seed.git
cd go-api-project-seed
```

### 2. 配置文件
修改 `configs/config.yaml`，配置 MySQL 数据库、Redis 等信息。

### 3. 生成代码
使用 `generate.go` 工具生成模型、控制器、服务和仓储代码：
```bash
go run tools/generate.go --table users --output internal
```

### 4. 运行项目
```bash
go run cmd/main.go
```

### 5. 访问 Swagger 文档  
启动服务后，访问 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) 查看 API 文档。

## 适用场景
- 中小型后端服务开发。
- 快速原型开发和代码迭代。
- 有自动化代码生成需求的项目。
- RESTful 风格的 API 服务。

## 未来规划
- 增加 GraphQL 支持。  
- 集成更多监控与报警工具（如 Jaeger）。  
- 支持多种数据库（如 PostgreSQL、MongoDB）。  
- 提供更多通用模板（如多租户、软删除等）。  

## 欢迎贡献和反馈  
如果你有任何问题或建议，请提交 [Issue](https://github.com/Kakaluote000/go-api-project-seed/issues) 或发起 PR！
```

---
