package hash

import (
	"errors"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/speps/go-hashids"
)

const (
	DonateId = iota // DonateId 捐赠信息
	UserId          // UserId 用户信息
)

var (
	ErrTypeNotExist = errors.New("type is not exist")
)

// Encode 对给定数据计算HashID
func Encode(v []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = conf.SystemConfig.HashIDSalt

	h := hashids.NewWithData(hd)

	id, err := h.Encode(v)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Decode 对给定数据计算原始数据
func Decode(raw string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = conf.SystemConfig.HashIDSalt

	h := hashids.NewWithData(hd)

	return h.DecodeWithError(raw)
}

// Id 计算数据库内主键对应的HashID
func Id(id uint, t int) string {
	v, _ := Encode([]int{int(id), t})
	return v
}

// DecodeID 计算HashID对应的数据库ID
func DecodeID(id string, t int) (uint, error) {
	v, _ := Decode(id)
	if len(v) != 2 || v[1] != t {
		return 0, ErrTypeNotExist
	}

	return uint(v[0]), nil
}
