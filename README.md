# 大学生信息管理系统

基于 [Eino](https://github.com/cloudwego/eino) 框架构建的大学生信息管理系统，集成 AI 智能助手，通过 Skill 文档对外提供可复用技能（兼容 OpenClaw 等 AI Agent 平台）。

## 技术栈

- **后端**: Go + Gin + GORM + SQLite
- **前端**: React 18 + TypeScript + Ant Design + Vite
- **AI**: Eino ReAct Agent + DeepSeek API
- **对外能力**: Skill 技能文档 + REST API

## 功能特性

- 学生信息管理（增删改查、搜索筛选）
- 课程管理
- 成绩管理（录入、统计、GPA 计算）
- AI 智能助手（自然语言操作系统）
- 会话历史持久化（多轮对话自动保存）
- 学业预警（GPA 低于阈值自动检测）
- Skill 技能文档（供外部 AI Agent 通过 REST API 调用系统能力）

## 快速开始

### 1. 启动后端

```bash
cd backend

# 安装依赖
go mod tidy

# 配置环境变量（复制模板并填入你的 API Key）
cp ../.env.example ../.env
# 编辑 ../.env，填入 DEEPSEEK_API_KEY

# 启动服务
go run ./cmd/server/
```

### 2. 启动前端

```bash
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

访问 http://localhost:3000 即可使用。

### 3. Skill 对外技能

`skills/` 目录下的 Skill 文档可被外部 AI Agent（如 OpenClaw）加载：

- `skills/student-manage/` - 学生管理技能（查询/增删改学生和课程）
- `skills/grade-analysis/` - 成绩分析技能（统计、GPA、预警、报告）

外部 Agent 加载 Skill 后，通过 `curl` 调用 REST API 完成操作。需设置环境变量 `EINO_OA_BASE_URL` 指向后端地址（如 `http://localhost:8080/api`）。

## 环境变量

配置文件 `.env` 放在项目根目录，后端启动时自动加载。

| 变量名 | 说明 | 默认值 |
| --- | --- | --- |
| `SERVER_PORT` | HTTP 服务端口 | 8080 |
| `DB_PATH` | SQLite 数据库路径 | data/eino-oa.db |
| `DEEPSEEK_API_KEY` | DeepSeek API 密钥 | (必填，启用 AI) |
| `DEEPSEEK_MODEL` | 模型名称 | deepseek-chat |
| `DEEPSEEK_BASE_URL` | API 基础地址 | https://api.deepseek.com/v1 |

## 项目结构

```
eino-oa/
  backend/
    cmd/
      server/          主服务入口（HTTP API + AI Agent）
    internal/
      model/           数据模型
      repository/      数据库操作层
      service/         业务逻辑层
      handler/         HTTP Handler
      agent/           Eino Agent + Tools
    config/            配置管理
  frontend/
    src/
      pages/           业务页面
      services/        API 封装
      components/      通用组件
  skills/
    student-manage/    学生管理技能（REST API）
    grade-analysis/    成绩分析技能（REST API）
```

## AI 助手用法示例

- "查一下张三的成绩"
- "分析一下数据结构的成绩分布"
- "哪些学生的 GPA 低于 2.0"
- "帮我录入一个新学生，学号 2024010，姓名刘明"
- "生成王五的综合成绩报告"
