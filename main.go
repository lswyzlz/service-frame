package main

import (
	"service-frame/pkg/loger"
	"service-frame/pkg/setting"
)

func main() {
	str := setting.AppConfig.GetString("Author")
	t := setting.AppConfig.GetTime("TimeStamp")
	t1 := setting.AppConfig.GetInt("Favorite.LuckyNumber1")

	loger.Apploger.WithField("str", str).Warnln("this is a warn")
	loger.Apploger.WithField("t", t).Warnln("this is a warn")
	loger.Apploger.WithField("t1", t1).Warnln("this is a warn")
}
