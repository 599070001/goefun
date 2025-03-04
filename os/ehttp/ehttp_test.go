package ehttp

import (
	E "github.com/599070001/goefun/core"
	"testing"
)

func TestNewHttp(t *testing.T) {
	http := NewHttp()
	ret, flag := http.Get("https://www.taobao.com/help/getip.php")
	if flag {
		t.Log("访问失败", ret)
	} else {
		t.Log(flag, ret)
	}

	ret, flag = http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=myip")
	if flag {
		t.Log("访问失败", ret)
	} else {
		t.Log(flag, ret)
	}

	retByte, flag := http.GetByte("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1596964553726&di=e77c87e7f2c5a8b0b12bdf6c13fbefe9&imgtype=0&src=http%3A%2F%2Fa2.att.hudong.com%2F36%2F48%2F19300001357258133412489354717.jpg")
	if flag {
		t.Log("访问失败", ret)
	} else {
		t.Log(flag, retByte)
	}

	E.E写到文件("1.jpg", retByte)

}
