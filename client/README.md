# MCP 客户端

这是一个使用Streamlit构建的MCP客户端应用，用于与MCP服务器交互。

## 功能特点

- 支持配置MCP服务器地址
- 支持API密钥认证
- 提供友好的用户界面
- 支持调用hello_world工具

## 安装

1. 确保已安装Python 3.8或更高版本
2. 安装依赖：
```bash
pip install -r requirements.txt
```

## 运行

1. 启动MCP服务器（确保在8080端口运行）
2. 运行Streamlit应用：
```bash
streamlit run app.py
```

## 使用说明

1. 在侧边栏配置MCP服务器地址和API密钥
2. 在主界面选择要使用的工具
3. 根据工具要求输入参数
4. 点击"调用工具"按钮执行操作

## 环境变量

可以创建`.env`文件来设置默认配置：

```
MCP_SERVER_URL=http://localhost:8080
API_KEY=your_api_key
``` 