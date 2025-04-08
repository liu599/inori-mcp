import requests
import json
import os
from typing import Dict, Any, List, Optional
from openai import OpenAI

class MCPClient:
    def __init__(self, server_url: str, api_key: str):
        self.server_url = server_url.rstrip('/')
        self.api_key = api_key
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {api_key}"
        }
        self.client = OpenAI()
    
    async def list_tools(self) -> List[Dict[str, Any]]:
        """获取所有可用的工具列表"""
        response = requests.get(
            f"{self.server_url}/tools",
            headers=self.headers
        )
        response.raise_for_status()
        return response.json()
    
    async def call_tool(self, tool_name: str, arguments: Dict[str, Any]) -> Dict[str, Any]:
        """调用指定的工具"""
        payload = {
            "params": {
                "arguments": arguments
            }
        }
        
        response = requests.post(
            f"{self.server_url}/call/{tool_name}",
            json=payload,
            headers=self.headers
        )
        response.raise_for_status()
        return response.json()
    
    async def process_query(self, query: str) -> str:
        """处理用户查询，使用 MCP 工具生成回答"""
        # 构建系统提示
        system_prompt = (
            "You are a helpful assistant. "
            "You have access to the following tools: "
            f"{json.dumps(tools, ensure_ascii=False)}. "
            "Please use these tools to help answer the user's question. "
            "When using a tool, make sure to provide all required parameters."
        )
        
        # 构建消息
        messages = [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": query}
        ]

        # 获取所有 mcp 服务器工具列表信息
        tools = await self.list_tools()
        
        # 生成 function call 的描述信息
        available_tools = [{
            "type": "function",
            "function": {
                "name": tool["name"],
                "description": tool["description"],
                "parameters": tool["inputSchema"]
            }
        } for tool in tools]

        # 请求 OpenAI，function call 的描述信息通过 tools 参数传入
        response = self.client.chat.completions.create(
            model=os.getenv("OPENAI_MODEL", "gpt-3.5-turbo"),
            messages=messages,
            tools=available_tools
        )

        # 处理返回的内容
        content = response.choices[0]
        if content.finish_reason == "tool_calls":
            # 如果需要使用工具，就解析工具
            tool_call = content.message.tool_calls[0]
            tool_name = tool_call.function.name
            tool_args = json.loads(tool_call.function.arguments)

            # 执行工具
            result = await self.call_tool(tool_name, tool_args)
            print(f"\n\n[Calling tool {tool_name} with args {tool_args}]\n\n")
            
            # 将 OpenAI 返回的调用哪个工具数据和工具执行完成后的数据都存入 messages 中
            messages.append({
                "role": "assistant",
                "content": None,
                "tool_calls": [{
                    "id": tool_call.id,
                    "type": "function",
                    "function": {
                        "name": tool_name,
                        "arguments": tool_call.function.arguments
                    }
                }]
            })
            messages.append({
                "role": "tool",
                "content": result.get("text", ""),
                "tool_call_id": tool_call.id
            })

            # 将上面的结果再返回给 OpenAI 用于生成最终的结果
            response = self.client.chat.completions.create(
                model=os.getenv("OPENAI_MODEL", "gpt-3.5-turbo"),
                messages=messages
            )
            return response.choices[0].message.content

        return content.message.content 