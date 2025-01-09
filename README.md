# Sylix-Go-Webserver

这是一个基于 GIN 框架的 Web 服务器示例项目。该项目展示了如何使用 GIN 框架来构建高效的 Web 应用，同时利用了 Go 1.16 引入的 `embed` 包来嵌入静态文件系统（embedfs）。

## 特性

- **GIN 框架**：使用 GIN 框架构建高性能的 HTTP 服务器。
- **embedfs**：利用 Go 1.16 的 `embed` 包嵌入静态文件，简化文件管理。
- **SylixOS 支持**：替换了 `go.mod` 中有关 SylixOS 的包，以确保兼容性。

## 运行示例

该项目是一个演示性质的项目，旨在展示如何结合使用 GIN 框架和 `embed` 包来构建 Web 服务器。你可以根据需要进行修改和扩展。

```bash
# 克隆仓库
git clone https://github.com/yourusername/sylix-go-webserver.git

# 进入项目目录
cd sylix-go-webserver

# 运行项目
./build.sh
```

## 许可证

本项目采用 MIT 许可证，详细信息请参阅 LICENSE 文件。