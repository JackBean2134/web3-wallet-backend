# Web3 Wallet Backend

<div align="center">

以太坊钱包管理后端服务，基于 Go 和 Gin 框架构建

![Go](https://img.shields.io/badge/Go-1.25.6-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-1.12.0-00ADD8?style=for-the-badge&logo=go)
![Ethereum](https://img.shields.io/badge/Ethereum-go--ethereum-3C3C3D?style=for-the-badge&logo=ethereum)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

[功能特性](#-功能特性) • [快速开始](#-快速开始) • [API 文档](#-api-文档) • [核心实现](#-核心实现)

</div>

---

## ✨ 功能特性

-  **创建钱包** - 一键生成以太坊地址和私钥对
- 💰 **余额查询** - 实时查询任意地址的 ETH 余额
- 🔄 **转账功能** - 完整的 ETH 转账流程（签名、广播、确认）
- 🔍 **交易追踪** - 查询交易状态、确认数和 Gas 使用情况
- ✅ **地址验证** - 智能验证以太坊地址格式，防止错误
- 🛡️ **错误处理** - 完善的参数验证和错误提示机制

## 🛠 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.25.6 | 高性能后端语言 |
| Gin | 1.12.0 | 轻量级 Web 框架 |
| go-ethereum | 1.17.2 | 以太坊官方 Go 实现 |
| Ankr RPC | Public | 公共以太坊节点 |

##  项目结构
web3-wallet-backend/
── config/ # 配置管理（环境变量加载）
├── controller/ # HTTP 控制器层
│ ├── wallet.go # 钱包相关接口
│ └── transfer.go # 转账相关接口
├── model/ # 数据模型定义
├── router/ # 路由配置
├── service/ # 业务逻辑层
├── utils/ # 工具函数（地址验证、单位转换等）
├── main.go # 程序入口
├── go.mod # 依赖管理
├── go.sum # 依赖校验
├── .gitignore # Git 忽略文件
└── README.md # 项目文档
##  快速开始

### 前置要求

- Go 1.25.6 或更高版本
- 网络连接（用于访问以太坊 RPC 节点）

### 安装依赖

```bash
git clone https://github.com/JackBean2134/web3-wallet-backend.git
cd web3-wallet-backend
go mod download

配置环境变量（可选）
创建 .env 文件：
```bash
SERVER_PORT=8080
ETH_RPC_URL=https://rpc.ankr.com/eth

运行服务
```bash
go run main.go
服务将在 http://localhost:8080 启动

测试接口
1. 创建钱包
```bash
curl -X POST http://localhost:8080/wallet/create

响应示例：
{
  "address": "0x063FF56B25e73f6f3fD0f05ebD3c997381651f1E",
  "private_key": "2963e00590c97df15becb62514b5fc102f518d60142b96224151a092b49bf754"
}

2. 查询余额
```bash
curl "http://localhost:8080/wallet/balance?address=0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc"

响应示例：
{
  "address": "0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc",
  "balance": "0"
}

📡 API 文档
1. 创建钱包
接口: POST /wallet/create说明: 生成新的以太坊钱包地址和私钥
请求体: 无
响应:
{
  "address": "0x...",
  "private_key": "..."
}

2. 查询余额
接口: GET /wallet/balance说明: 查询指定地址的 ETH 余额查询参数:
参数	类型	必填	说明
address	string	✅	以太坊地址
示例:
```bash
GET /wallet/balance?address=0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc

响应:
{
  "address": "0xE7381B9e9AaB14E90089Bba889D9e01fCa1F34bc",
  "balance": "1.234567890123456789"
}

3. 转账 ETH
接口: POST /wallet/transfer
说明: 从指定钱包向目标地址转账 ETH
请求体:
{
  "from_address": "0x...",
  "private_key": "...",
  "to_address": "0x...",
  "amount": "0.1",
  "gas_limit": 21000,
  "gas_price": "20000000000"
}
参数说明:
参数	类型	必填	说明
from_address	string	✅	发送方地址
private_key	string	✅	发送方私钥
to_address	string	✅	接收方地址
amount	string	✅	转账金额（ETH）
gas_limit	uint64	❌	Gas 限制（默认自动估算）
gas_price	string		Gas 价格（Wei，默认自动获取）

响应:
{
  "tx_hash": "0x...",
  "from": "0x...",
  "to": "0x...",
  "amount": "0.1",
  "status": "pending"
}

4. 查询交易状态
接口: GET /wallet/transaction/status
说明: 查询交易的状态和确认数
查询参数:
参数	类型	必填	说明
tx_hash	string	✅	交易哈希
示例:
```bash
GET /wallet/transaction/status?tx_hash=0x...

响应:
{
  "tx_hash": "0x...",
  "status": "success",
  "block_number": 12345678,
  "gas_used": 21000,
  "confirmations": 12
}

🔧 核心实现
私钥签名
使用 EIP-155 签名方案，支持不同链的 ChainID：
signer := types.LatestSignerForChainID(chainID)
signedTx, err := types.SignTx(tx, signer, privateKey)

Nonce 管理
自动获取待处理交易的 nonce，避免冲突：
nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
