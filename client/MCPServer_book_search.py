# -*- coding: utf-8 -*-
# coding: utf-8
# @Time        : 2025/4/10
# @Author      : liuboyuan
# @Description : MCP 服务器实现，提供图书搜索服务

import json

import httpx
from mcp.server import FastMCP

# 初始化 FastMCP 服务器，服务名称为 'book-search'
app = FastMCP('book-search')

@app.tool()
async def book_search(query: str) -> str:
    """
    图书搜索工具实现
    
    这是一个模拟的图书搜索服务，用于演示 MCP 工具调用功能。
    在实际应用中，这里可以替换为真实的图书搜索 API 调用。

    Args:
        query: 用户输入的搜索关键词

    Returns:
        str: 搜索结果的字符串，每本书之间用换行符分隔
    """
    # 模拟图书搜索结果
    books = [
        '钢铁是怎样炼成的',
        '魔法少女小圆',
        '线性代数',
        '丹特丽安的书架'
    ]

    # 将结果用换行符连接，便于阅读
    return '\n\n\n'.join(books)

if __name__ == "__main__":
    # 启动 MCP 服务器，使用 stdio 传输方式
    print('starting MCP book search server...')
    app.run(transport='stdio')