package utils

import (
	"errors"
	"github.com/speps/go-hashids/v2"
)

var (
	hd *hashids.HashID
)

// InitHashID 初始化HashID
func InitHashID(salt string, minLength int) error {
	hdConfig := hashids.NewData()
	hdConfig.Salt = salt           // 加盐，提高安全性
	hdConfig.MinLength = minLength // 生成的最小长度

	var err error
	hd, err = hashids.NewWithData(hdConfig)
	return err
}

// EncodeID 将数字ID编码为哈希字符串
func EncodeID(id int64) (string, error) {
	if hd == nil {
		return "", errors.New("hashid not initialized")
	}
	return hd.EncodeInt64([]int64{id})
}

// DecodeID 将哈希字符串解码为原始数字ID
func DecodeID(hash string) (int64, error) {
	if hd == nil {
		return 0, errors.New("hashid not initialized")
	}

	ids, err := hd.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}

	if len(ids) == 0 {
		return 0, errors.New("invalid hash")
	}

	return ids[0], nil
}
