# 词库维护脚本

运行维护前用 `RIME_DIR` 指定待处理的 Rime 配置副本；未设置时沿用
`~/Library/Rime`。`go run .` 会改写词库，建议在 Git 工作区执行后检查 diff。

```sh
RIME_DIR=/path/to/rime go run .
```

验证构建和路径配置：

```sh
RIME_DIR=/path/to/rime go test ./...
```

测试会读取该目录中的词库和检查规则，不执行主程序的词库更新流程。
这不是 Rime 引擎的部署或输入行为测试。
