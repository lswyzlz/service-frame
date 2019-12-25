package main

import (
	"fmt"
	"service-frame/pkg/setting"
)

func main() {
	str := setting.AppConfig.GetString("Author")
	t := setting.AppConfig.GetTime("TimeStamp")
	t1 := setting.AppConfig.GetInt("Favorite.LuckyNumber1")

	fmt.Println(str)
	fmt.Println(t)
	fmt.Println(t1)
}
