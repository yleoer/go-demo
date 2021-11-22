package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yleoer/go-demo/mockDemo/equipment"
)

func TestPersion_dayLife(t *testing.T) {
	type fields struct {
		name string
		phone equipment.Phone
	}

	// 生成 mockPhone 对象
	mockCtl := gomock.NewController(t)
	mockPhone := equipment.NewMockPhone(mockCtl)

	// 设置 mockPhone 对象的接口方法返回值
	mockPhone.EXPECT().ZhiHu().Return(true)
	mockPhone.EXPECT().WeiXin().Return(true)
	mockPhone.EXPECT().WangZhe().Return(true)

	tests := []struct{
		name string
		fields fields
		want bool
	} {
		{"case1", fields{"iphone6s", equipment.NewIphone6s()}, true},
		{"case2", fields{"mocked phone", mockPhone}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &Person{
				name: tt.fields.name,
				phone: tt.fields.phone,
			}
			assert.Equal(t, tt.want, x.dayLife())
		})
	}
}
