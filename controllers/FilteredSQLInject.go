package controllers

import (
	"github.com/astaxie/beego/logs"
	"regexp"
)

// 正则过滤sql注入的方法
// 参数 : 要匹配的语句
func FilteredSQLInject(input string) bool {
	//过滤 ‘
	//ORACLE 注解 --  /**/
	//关键字过滤 update ,delete
	// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
	//如果字符串中存在注入语句将返回true
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
	} else {
		if re.MatchString(input) == true {
			logs.Info("出现sql注入")
			return true
		}
	}
	return false
}
