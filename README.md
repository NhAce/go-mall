# Go商城后端项目

这是一个基于Go语言开发的微信小程序商城后端项目，主要售卖农副产品。

## 技术栈

- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **配置管理**: Viper
- **日志**: Zap

## 项目结构

```
go-mall/
├── main.go                 # 程序入口
├── go.mod                  # Go模块文件
├── config/                 # 配置文件
│   ├── config.yaml        # 配置文件
│   └── config.go          # 配置结构体
├── model/                  # 数据模型
│   ├── user.go            # 用户模型
│   ├── product.go         # 商品模型
│   ├── cart.go            # 购物车模型
│   └── order.go           # 订单模型
├── controller/             # 控制器
│   ├── user.go            # 用户控制器
│   └── product.go         # 商品控制器
├── middleware/             # 中间件
│   └── auth.go            # JWT认证中间件
├── router/                 # 路由
│   └── router.go          # 路由注册
├── utils/                  # 工具函数
│   ├── database.go        # 数据库连接
│   ├── response.go        # 响应工具
│   └── jwt.go             # JWT工具
└── README.md              # 项目说明
```

## 安装和运行

### 1. 环境要求

- Go 1.21+
- MySQL 5.7+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

1. 创建MySQL数据库：
```sql
CREATE DATABASE go_mall CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改配置文件 `config/config.yaml`：
```yaml
database:
  host: 127.0.0.1
  port: 3306
  user: your_username
  password: your_password
  dbname: go_mall
```

### 4. 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API接口

### 用户相关

- `POST /api/user/login` - 微信小程序登录
- `GET /api/user/info` - 获取用户信息（需要认证）
- `POST /api/user/update` - 更新用户信息（需要认证）

### 商品相关

- `GET /api/product/list` - 获取商品列表
- `GET /api/product/detail/:id` - 获取商品详情
- `GET /api/category/list` - 获取分类列表

## 开发计划

- [ ] 购物车功能
- [ ] 订单管理
- [ ] 地址管理
- [ ] 微信支付集成
- [ ] 商品评价
- [ ] 后台管理

## 注意事项

1. 微信登录功能需要配置真实的微信小程序AppID和Secret
2. 生产环境请修改JWT密钥
3. 建议使用HTTPS协议
4. 注意数据库备份和日志管理 