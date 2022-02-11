/*================================================================
*
*  文件名称：errno.go
*  创 建 者: mongia
*  创建日期：2021年12月29日
*
================================================================*/

package util

type Errno struct {
	Code int32
	Msg  string
}

func (err Errno) Error() string {
	return err.Msg
}

var (
	EOK    = &Errno{0, "success"}
	EParam = &Errno{100001, "input param error"}

	EDB       = &Errno{200001, "db error"}
	EDBNoData = &Errno{200002, "no data found in db"}

	ERedis = &Errno{201001, "redis error"}
)
