package constant

import "time"

const (
	// JumpOutToken 跳出获取历史记录的token数量
	JumpOutToken = 1000
	// DefaultContextModel 默认的上下文大模型ID
	DefaultContextModel = 0
	// DefaultMaxToken 不输入token时默认定义的token最大数量
	DefaultMaxToken = 1000
	// DefaultTitle 默认标题
	DefaultTitle = "init"
	//DefaultAdminUID 默认的管理员uid （官方调试所用uid）
	DefaultAdminUID = "0"
	// DefaultMaxLimitedTime 请求默认超时时间 方便调试默认关闭
	DefaultMaxLimitedTime = time.Minute / 2
	// ApiServerOpenAI OpenAI-API服务器默认网址
	//ApiServerOpenAI = "https://api.openai.com/v1/completions"
	ApiServerOpenAI = "https://api.openai.com/v1/chat/completions"
	// InstructModel 初始模型
	InstructModel = "gpt-3.5-turbo-instruct"
	// DefaultModel 默认模型
	DefaultModel = "gpt-3.5-turbo-0125"
	// ReplaceCharFromDefaultToCustomize 自定义唯一标识符 选了个挺少见的 可优化算法
	ReplaceCharFromDefaultToCustomize = '¶'
	// OfficialBotPrefix 创建新机器人的前缀
	OfficialBotPrefix = "OfficialBot"
	// UserCachePrefix 用户chat缓存前缀
	UserCachePrefix = "UserCache"
	// OfficialBotIdList redis中存储官方机器人id 维护的便于id查找的list
	OfficialBotIdList = "OfficialBotIdList"
	// ChatCache redis中存储以往chat记录的缓存前缀
	ChatCache = "ChatCache"
	// ChatCacheExpire redis中存储chat记录的限时
	ChatCacheExpire = 30 * time.Minute
	// HistoryChatPrompt 告诉chatGPT以往聊天记录的prompt模板 可改进
	HistoryChatPrompt = "Here is the chat history which I have talked with you,please according to the history give me generation:"
	// SystemRole 系统角色
	SystemRole = "system"
	// UserRole 用户角色
	UserRole = "user"
	// GPTRole GPT角色
	GPTRole = "assistant"
	// NowAsk 当前的一次询问
	NowAsk = "And Here is my question this time:  "
	// ChatHistoryWeight 发送上下文历史记录的权重设置
	ChatHistoryWeight = 3
	// APIExecuteSuccessStatus 执行API成功后返回的状态码
	APIExecuteSuccessStatus = 200
	// ReferenceRecordPrompt 告诉chatGPT要他回复回应中的某个部分
	ReferenceRecordPrompt = "Here is a record we have been talked,And I have confused about parts of your generation,please fairly and clearly explain about it and my question:"
	// DefaultShareSecretKeyDestroyTime 默认分享密钥存活时间
	DefaultShareSecretKeyDestroyTime = 24 * time.Hour * 3
	// DefaultRecycledTime 回收站存放时间 超过此时间将被回收 实际运行值为30 多出来的一天是给定时任务进行删除的
	DefaultRecycledTime = 31
	//DefaultRecycledDeleteTTL 执行定时任务的时间
	DefaultRecycledDeleteTTL = 1
	// DefaultRecycledPrefix 回收站中放入redis的前缀
	DefaultRecycledPrefix = "recycled_"
	// DefaultRecycledList 回收站的逻辑表
	DefaultRecycleListPrefix = "recycled_list_"
)
