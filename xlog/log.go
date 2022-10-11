package xlog

import "context"

// Logger 是接口
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(err error, msg string, keysAndValues ...interface{})
}

// Debug 打印开发者需要知道的输出, keyAndValues是键值对，并且不会自动合并重复的键，如果是奇数个，那么会忽略最后一个key
func Debug(ctx context.Context, msg string, keyAndValues ...interface{}) {

}

// Info 打印需要用户知道的输出, keyAndValues是键值对，并且不会自动合并重复的键，如果是奇数个，那么会忽略最后一个key
func Info(ctx context.Context, msg string, keyAndValues ...interface{}) {

}

// Error 打印错误，这里的错误打印，不应该将堆栈打印出来，堆栈信息应该附着在error上
func Error(ctx context.Context, err error) {

}