package main

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

func OpenDouDian() {
	// 打开浏览器
	explore := rod.New().MustConnect()
	defer explore.MustClose()
	// 打开页面
	page := explore.MustPage("https://fxg.jinritemai.com/")
	// 点击抖音登录
	loginButton := page.MustWaitStable().MustElement("#oauth-login > div.index_oauthLoginBody__2c_95 > div:nth-child(1)")
	loginButton.MustClick()
	// 监测 URL 变动
	isChangeSuccess := listenUrlChange(page)
	// 若 URL 变动成功，则说明登录成功
	if <-isChangeSuccess {
		fmt.Println("登录成功")
	}
}

func listenUrlChange(page *rod.Page) chan bool {
	// 检测 URL 变动
	// url 中是否包含 fxg.jinritemai.com/ffa
	isChangeSuccess := make(chan bool)
	go func() {
		for {
			if strings.Contains(page.MustInfo().URL, "fxg.jinritemai.com/ffa") {
				isChangeSuccess <- true
				break
			}
		}
	}()
	return isChangeSuccess
}

func createCommodity(page *rod.Page) {
	// 查找“商品创建”按钮 #app-main > div.style_container-responsive__3rYw_.styles_responsive-wrapper__1qtNL > div.style_left-side-bar__2OzUz > div > div > div > div > div:nth-child(4) > div.style_content__3lQ6V > div > div:nth-child(1) > div > div
	//*[@id="app-main"]/div[4]/div[1]/div/div/div/div/div[4]/div[2]/div/div[1]/div/div
	///html/body/div[1]/div/div[4]/div[4]/div[1]/div/div/div/div/div[4]/div[2]/div/div[1]/div/div
	createButton := page.MustWaitStable().MustElement("#app-main > div.style_container-responsive__3rYw_.styles_responsive-wrapper__1qtNL > div.style_left-side-bar__2OzUz > div > div > div > div > div:nth-child(4) > div.style_content__3lQ6V > div > div:nth-child(1) > div > div")
	createButton.MustClick()
}

func main() {
	OpenDouDian()
}
