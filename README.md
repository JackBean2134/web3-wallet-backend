# Web3 Wallet Backend

<div align="center">

基于 Go + Gin 的以太坊钱包管理后端服务

![Go](https://img.shields.io/badge/Go-1.25.6-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-1.12.0-00ADD8?style=for-the-badge)
![Ethereum](https://img.shields.io/badge/Ethereum-go--ethereum-3C3C3D?style=for-the-badge&logo=ethereum)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)
---
</div>

## ✨ 功能
- 🔐 创建以太坊钱包（地址 + 私钥）
- 💰 查询 ETH 余额
- 🔄 ETH 转账（签名、广播、确认）
- 🔍 交易状态查询

## 🚀 快速开始

```bash
# 克隆项目
git clone https://github.com/JackBean2134/web3-wallet-backend.git
cd web3-wallet-backend

# 下载依赖
go mod download

# 启动服务
go run main.go
```

服务运行在 `http://localhost:8080`

## 📡 API 接口

### 1. 创建钱包

bash

```
POST http://localhost:8080/wallet/create
```

**响应：**

json

```json
{  
	"address": "0x063FF56B25e73f6f3fD0f05ebD3c997381651f1E",  
	 "private_key": "2963e00590c97df15becb62514b5fc102f518d60142b96224151a092b49bf754"
}
```

### 2. 查询余额

```bash
GET http://localhost:8080/wallet/balance?address=0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc
```

**响应：**

json

```json
{  
	"address": "0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc",  
	"balance": "1.234567890123456789" 
}
```

### 3. 转账 ETH

bash

```bash
POST http://localhost:8080/wallet/transfer 
Content-Type: application/json 
{  
	"from_address": "0x...",  
	"private_key": "...",  
	"to_address": "0x...",  
	"amount": "0.1" 
}
```

**响应：**

json

```json
{  
	"tx_hash": "0x...",  
	"status": "pending" 
}
```

### 4. 查询交易状态

bash

```bash
GET http://localhost:8080/wallet/transaction/status?tx_hash=0x...
```

**响应：**

json

```json
{  
   "tx_hash": "0x...",  
   "status": "success",  
   "confirmations": 12 
}
```

## 🔧 核心技术

- **私钥签名**: EIP-155 标准，支持多链
- **Nonce 管理**: 自动获取 pending nonce，避免冲突
- **Gas 估算**: 智能估算，失败 fallback 到 21000
- **地址验证**: 完整校验以太坊地址格式

## ⚙️ 环境变量

| 变量          | 说明     | 默认值                   |
| :------------ | :------- | :----------------------- |
| `SERVER_PORT` | 服务端口 | 8080                     |
| `ETH_RPC_URL` | RPC 地址 | https://rpc.ankr.com/eth |

## 📦 项目结构

```plaintext
├── config/       # 配置管理 
├── controller/   # HTTP 控制器 
├── model/        # 数据模型 
├── router/       # 路由配置 
├── service/      # 业务逻辑 
├── utils/        # 工具函数 
└── main.go       # 程序入口
```

## 🛡️ 安全提示

> ⚠️ 本项目用于学习/测试，生产环境需要：
>
> - 使用私有 RPC 节点
> - 启用 HTTPS
> - 添加 API 认证
> - 实现密钥加密存储
>
> ## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

------

<div align="center">

⭐ **如果这个项目对您有帮助，请给个 Star！**

Made with ❤️ by [JackBean2134](https://github.com/JackBean2134)

</div>
