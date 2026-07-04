# CS2Lab

CS2 全竞技地图道具教学工具，覆盖投掷点位、准星参照、道具组合配合。

## Features

- 7 张竞技地图的烟雾弹、闪光弹、燃烧弹、高爆手雷点位图鉴
- 每处投掷标注站位、准星参照截图和投掷方式（左键/右键/跳投）
- 多道具配合说明，地图叠加显示覆盖区域
- 按场景描述智能匹配已有教学内容
- 用户投稿未收录点位，审核后发布
- 收藏常用点位并添加个人笔记

## Architecture

```text
Browser
  ↓
React Frontend (Oxelia51 unified UI, image annotation)
  ↓
Go API Layer (content CRUD, search, submission review)
  ↓
PostgreSQL / SQLite (maps, lineups, user data)
```

在线版运行于 Oxelia51 平台。道具截图和标注数据通过 Go 后端管理。桌面版使用 SQLite 替代 PostgreSQL，地图数据打包在二进制中，无需联网。

## Requirements

- 在线版：Oxelia51 平台（Go + PostgreSQL + React）
- 桌面版：独立可执行文件，无需运行时依赖

## Installation

### 桌面版

从 [GitHub Releases](https://github.com/XiaoleC05/CS2Lab/releases) 下载 `CS2Lab.exe`。

### 在线版

在线版集成于 Oxelia51 平台，参见 [Oxelia51 部署指南](https://github.com/XiaoleC05/Oxelia51)。

## Usage

### 在线

1. 访问 [oxelia51.com](https://oxelia51.com) 注册并登录
2. 进入 CS2Lab 工具页
3. 选择地图和道具类型，浏览点位

### 桌面

1. 双击 `CS2Lab.exe` 启动
2. 数据内置，无需联网

## Roadmap

- [ ] 首张地图（Dust2）完整道具覆盖
- [ ] 扩展至全部 7 张竞技地图
- [ ] 智能搜索
- [ ] 用户投稿与审核

## Contributing

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/xxx`)
3. 提交变更 (`git commit -m 'Add xxx'`)
4. 推送分支 (`git push origin feature/xxx`)
5. 提交 Pull Request

## License

This project is licensed under the MIT License.
