package uniqueid

import (
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenerateId() uint64 {
	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextId failed with:%s \n", err)
	}
	return id
}
