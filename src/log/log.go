/*================================================================
*
*  文件名称：log.go
*  创 建 者: mongia
*  创建日期：2022年01月11日
*
================================================================*/

package util

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

type Field = zap.Field

func InitLog() {
	//	hook := lumberjack.Logger{
	//		Filename:   "./logs/pharmacy.log", // 日志文件路径
	//		MaxSize:    128,                   // 每个日志文件保存的最大尺寸 单位：M
	//		MaxBackups: 30,                    // 日志文件最多保存多少个备份
	//		MaxAge:     7,                     // 文件最多保存多少天
	//		Compress:   true,                  // 是否压缩
	//	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 跳过调用个数
	callerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("serviceName", "pharmacyerp"))
	// 构造日志
	logger := zap.New(core, caller, callerSkip, development, field)

	Log = logger
}

func Info(msg string, fields ...Field) {
	Log.Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	Log.Debug(msg, fields...)
}

func Error(msg string, fields ...Field) {
	Log.Debug(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	Log.Warn(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	Log.Fatal(msg, fields...)
}

func Bool(key string, val bool) Field {
	return zap.Bool(key, val)
}

func Boolp(key string, val *bool) Field {
	return zap.Boolp(key, val)
}

func ByteString(key string, val []byte) Field {
	return zap.ByteString(key, val)
}

func Complex128(key string, val complex128) Field {
	return zap.Complex128(key, val)
}

func Complex128p(key string, val *complex128) Field {
	return zap.Complex128p(key, val)
}

func Complex64(key string, val complex64) Field {
	return zap.Complex64(key, val)
}

func Complex64p(key string, val *complex64) Field {
	return zap.Complex64p(key, val)
}

func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

func Float64p(key string, val *float64) Field {
	return zap.Float64p(key, val)
}

func Float32(key string, val float32) Field {
	return zap.Float32(key, val)
}

func Float32p(key string, val *float32) Field {
	return zap.Float32p(key, val)
}

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func Intp(key string, val *int) Field {
	return zap.Intp(key, val)
}

func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

func Int64p(key string, val *int64) Field {
	return zap.Int64p(key, val)
}

func Int32(key string, val int32) Field {
	return zap.Int32(key, val)
}

func Int32p(key string, val *int32) Field {
	return zap.Int32p(key, val)
}

func Int16(key string, val int16) Field {
	return zap.Int16(key, val)
}

func Int16p(key string, val *int16) Field {
	return zap.Int16p(key, val)
}

func Int8(key string, val int8) Field {
	return zap.Int8(key, val)
}

func Int8p(key string, val *int8) Field {
	return zap.Int8p(key, val)
}

func String(key string, val string) Field {
	return zap.String(key, val)
}

func Stringp(key string, val *string) Field {
	return zap.Stringp(key, val)
}

func Uint(key string, val uint) Field {
	return zap.Uint(key, val)
}

func Uintp(key string, val *uint) Field {
	return zap.Uintp(key, val)
}

func Uint64(key string, val uint64) Field {
	return zap.Uint64(key, val)
}

func Uint64p(key string, val *uint64) Field {
	return zap.Uint64p(key, val)
}

func Uint32(key string, val uint32) Field {
	return zap.Uint32(key, val)
}

func Uint32p(key string, val *uint32) Field {
	return zap.Uint32p(key, val)
}

func Uint16(key string, val uint16) Field {
	return zap.Uint16(key, val)
}

func Uint16p(key string, val *uint16) Field {
	return zap.Uint16p(key, val)
}

func Uint8(key string, val uint8) Field {
	return zap.Uint8(key, val)
}

func Uint8p(key string, val *uint8) Field {
	return zap.Uint8p(key, val)
}

func Uintptr(key string, val uintptr) Field {
	return zap.Uintptr(key, val)
}

func Uintptrp(key string, val *uintptr) Field {
	return zap.Uintptrp(key, val)
}

func Reflect(key string, val interface{}) Field {
	return zap.Reflect(key, val)
}

func Namespace(key string) Field {
	return zap.Namespace(key)
}

func Stringer(key string, val fmt.Stringer) Field {
	return zap.Stringer(key, val)
}

func Time(key string, val time.Time) Field {
	return zap.Time(key, val)
}

func Timep(key string, val *time.Time) Field {
	return zap.Timep(key, val)
}

func Stack(key string) Field {
	return zap.Stack(key)
}

func StackSkip(key string, skip int) Field {
	return zap.StackSkip(key, skip)
}

func Duration(key string, val time.Duration) Field {
	return zap.Duration(key, val)
}

func Durationp(key string, val *time.Duration) Field {
	return zap.Durationp(key, val)
}

func Object(key string, val zapcore.ObjectMarshaler) Field {
	return zap.Object(key, val)
}

func Inline(val zapcore.ObjectMarshaler) Field {
	return zap.Inline(val)
}

func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}
