package InoriMCP

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 配置结构体
type Config struct {
	LLM struct {
		BaseURL string `yaml:"base_url"`
		APIKey  string `yaml:"api_key"`
	} `yaml:"llm"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig(configPath string) (*Config, error) {
	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析配置
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 检查必要的配置项
	if config.LLM.APIKey == "" {
		return nil, fmt.Errorf("API密钥不能为空")
	}

	return &config, nil
}

// GetLLMConfig 获取LLM配置
func (c *Config) GetLLMConfig() (string, string) {
	return c.LLM.BaseURL, c.LLM.APIKey
} 