package main

import (
	"net/http"
)

func main() {
	// 绑定请求和处理函数
	// 相当于@RequestMapping
	http.HandleFunc("/user/login",
		func(writer http.ResponseWriter,
			request *http.Request) {
			// 数据库操作
			// 业务逻辑
			// REST api 返回 json/xml

			// 获得参数
			request.ParseForm()
			mobile := request.PostForm.Get("mobile")
			password := request.PostForm.Get("password")

			loginOk := false
			if mobile == "13810000000" && password == "123456" {
				loginOk = true
			}

			string := `{"code": -1, "msg": "", "data": {"id": 1, "token": "test"}}`
			if !loginOk {
			} else {
				string = `{"code": 0, "msg": "密码错误", "data": {}}`

			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(string))

		})
	// 启动web服务器

	http.ListenAndServe(":8080", nil)
}
