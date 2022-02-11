/*================================================================
*
*  文件名称：user_cache.go
*  创 建 者: mongia
*  创建日期：2022年01月05日
*
================================================================*/

package middleaware

import (
	"errors"
	"strconv"

	userpb "pharmacyerp/pb/user"
	cache "pharmacyerp/third"

	"github.com/gomodule/redigo/redis"
)

const (
	ucPrefix = "puc_"
)

type UserCache struct {
}

func (uc *UserCache) MakeCacheKey(key string) string {
	return ucPrefix + key
}

func (uc *UserCache) GetDirectKey(key string, keyType int) (string, error) {
	if int(userpb.UserItem_USER_ID) == keyType {
		return "", nil
	}

	return cache.GetString(key)
}

func (uc *UserCache) CheckInCache(key string, keyType int) (bool, error) {
	rdKey := uc.MakeCacheKey(key)
	if ucPrefix == rdKey {
		return false, errors.New("checkInCache function param error")
	}

	r, err := uc.GetDirectKey(key, keyType)
	if nil != err {
		return false, err
	}

	if "" != r {
		rdKey = r
	}

	return cache.Exists(rdKey)
}

func (uc *UserCache) GetInCache(key string, keyType int) (*UserInfo, error) {
	rdKey := uc.MakeCacheKey(key)
	if ucPrefix == rdKey {
		return nil, errors.New("GetInCache function param error")
	}

	r, err := uc.GetDirectKey(key, keyType)
	if nil != err {
		return nil, err
	}

	if "" != r {
		rdKey = r
	}

	user := &UserInfo{}
	err = cache.HGetALL(user, rdKey)
	if nil != err {
		return nil, err
	}
	if 0 == user.UserId {
		return nil, errors.New("No data in Cache")
	}
	return user, nil
}

func (uc *UserCache) RemoveInCache(users []*UserInfo) error {
	c := cache.GetRedisPool().Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return c.Err()
	}
	c.Send("MULTI")
	for _, user := range users {
		c.Send("DEL", uc.MakeCacheKey(strconv.FormatInt(user.UserId, 10)))

		if 0 != len(user.Telephone) {
			c.Send("DEL", uc.MakeCacheKey(user.Telephone))
		}
		if 0 != len(user.Username) {
			c.Send("DEL", uc.MakeCacheKey(user.Username))
		}
		if 0 != len(user.Card) {
			c.Send("DEL", uc.MakeCacheKey(user.Card))
		}
	}
	_, err := redis.Values(c.Do("EXEC"))
	if nil != err {
		return err
	}

	return nil
}

func (uc *UserCache) SetInCache(user *UserInfo) error {
	c := cache.GetRedisPool().Get()
	defer c.Close() // 不管连接获取正常与否，都将连接返回给连接池

	if nil != c.Err() {
		return c.Err()
	}

	key := uc.MakeCacheKey(strconv.FormatInt(user.UserId, 10))
	c.Send("MULTI")

	if 0 != len(user.Telephone) {
		c.Send("SET", uc.MakeCacheKey(user.Telephone), key)
	}
	if 0 != len(user.Username) {
		c.Send("SET", uc.MakeCacheKey(user.Username), key)
	}
	if 0 != len(user.Card) {
		c.Send("SET", uc.MakeCacheKey(user.Card), key)
	}
	c.Send("HMSET", redis.Args{}.Add(key).AddFlat(user)...)

	_, err := redis.Values(c.Do("EXEC"))
	if nil != err {
		return err
	}
	return nil
}
