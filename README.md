#使用说明
### 1 参数说明

#### 1.1 client参数说明

|  参数名   | 类型  | 是否必填 |说明|
|  ----  | ----  | --- | ----|
| eventSink  | string | 是 | 事件中心的url

#### 1.2 event参数说明

|  参数名   | 类型  | 是否必填 |说明|
|  ----  | ----  | --- | ----|
| subject  | string | 是 |事件主题：operate、alter、log等
| source  | string | 是 |事件来源：inspurcloud.产品代码（组件编码
| type  | string | 是 |事件类型：inspurcloud.产品代码.资源.操作
| resource  | string | 否 |资源类型：资源的类型，由各产品自己定义
| accountid  | string | 否 |账户
| action  | string | 否| 动作类型：在subject为operate时可以设为操作的类型。可选为create,list,modify,delete。
| region  | string | 否 |区域
| body  | interface{} | 是 |报文体：事件数据的报文体

### 2 示例


```
package main

import (
	"context"
	"fmt"

	icclent "cloud.inspur.com/event-center-sdk-go/client"
	icevent "cloud.inspur.com/event-center-sdk-go/event"
)

func main()  {
	c, _ := icclent.Init("http://localhost:8081")

	event := icevent.NewEvent()
	event.SetSubject("subject1")
	event.SetType("type")
	event.SetSource("source1")
	event.SetResource("rsource")
	event.SetAccountid("accoun")
	event.SetAction("list")
	event.SetRegion("rego")
	var data = map[string]string{"name":"asdf", "id":"47178080808"}
	err := event.SetBody(data)
	err = c.Send(context.Background(), event)
	if err != nil {
		fmt.Println(err)
	}
}
```