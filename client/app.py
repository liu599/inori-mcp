import streamlit as st
import asyncio
from dotenv import load_dotenv
import os
from mcp_client import MCPClient

# 加载环境变量
load_dotenv()

# 设置 Streamlit 服务器配置
st.set_page_config(
    page_title="MCP 客户端",
    page_icon="🤖",
    layout="wide"
)

# 标题
st.title("🤖 MCP 客户端")

# 侧边栏配置
with st.sidebar:
    st.header("配置")
    mcp_server_url = st.text_input(
        "MCP 服务器地址",
        value=os.getenv("MCP_SERVER_URL", "http://localhost:8080"),
        help="输入MCP服务器的地址"
    )
    api_key = st.text_input(
        "API Key",
        type="password",
        value=os.getenv("API_KEY", ""),
        help="输入API密钥"
    )

# 主界面
st.header("对话")

# 初始化会话状态
if "messages" not in st.session_state:
    st.session_state.messages = []

# 显示历史消息
for message in st.session_state.messages:
    with st.chat_message(message["role"]):
        st.markdown(message["content"])

# 用户输入
if prompt := st.chat_input("请输入您的问题"):
    # 添加用户消息到历史
    st.session_state.messages.append({"role": "user", "content": prompt})
    with st.chat_message("user"):
        st.markdown(prompt)
    
    # 创建 MCP 客户端实例
    client = MCPClient(mcp_server_url, api_key)
    
    # 处理查询并显示结果
    with st.chat_message("assistant"):
        try:
            # 使用 asyncio 运行异步函数
            response = asyncio.run(client.process_query(prompt))
            st.markdown(response)
            st.session_state.messages.append({"role": "assistant", "content": response})
        except Exception as e:
            st.error(f"发生错误: {str(e)}") 