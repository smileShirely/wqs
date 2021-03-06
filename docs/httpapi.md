## 队列接口
**http://ip:port/queue** <br>
**参数列表：**<br>

| 参数名 | 是否必填 | 说明 | 
| ---- | ---- | ----|
| action | 必填 | 创建/删除/变更/查看队列，参数为create,remove,update,lookup |
| queue | 必填 | 队列名称，查看队列时选填 |

**示例：** <br>
**创建队列：** <br>
curl -d "action=create&queue=menglong\_queue1" "http://127.0.0.1:8080/queue" <br>
{"action":"create","result":true} <br>

**删除队列：** <br>
curl -d "action=remove&queue=menglong\_queue1" "http://127.0.0.1:8080/queue"
{"action":"create","result":true} <br>

**变更队列：** <br>
暂不提供 <br>

**查看队列：** <br>
curl "http://127.0.0.1:8080/queue?action=lookup"<br>
curl "http://127.0.0.1:8080/queue?action=lookup&queue=menglong\_queue1"<br>
curl "http://127.0.0.1:8080/queue?action=lookup&queue=menglong\_queue1&group=menglong\_group1"<br>

## 业务接口
**http://ip:port/group** <br>
**参数列表：** <br>

| 参数名 | 是否必填 | 说明 | 
| ---- | ---- | ----|
| action | 必填 | 增加/删除/变更/查看业务方，参数为add,remove,update,lookup |
| group | 必填 | 业务标识，查看业务方时选填 |
| queue | 必填 | 队列名称 |
| write | 选填 | 默认false，增加和变更业务方时使用 |
| read | 选填 | 默认false，增加和变更业务方时使用 |
| url | 选填 | 业务方使用的域名，增加和变更业务方时使用 |
| ips | 选填 | 域名对应的ip，多个ip用逗号分隔，增加和变更业务方时使用 |

**示例：** <br>
**增加业务方：** <br>
curl -d "action=add&group=menglong\_group1&queue=menglong\_queue1&write=true&read=true" "http://127.0.0.1:8080/group" <br>
{"action":"add","result":true} <br>

**删除业务方：** <br>
curl -d "action=remove&group=menglong\_group1&queue=menglong\_queue1" "http://127.0.0.1:8080/group" <br>
{"action":"remove","result":true} <br>

**变更业务方：** <br>
curl -d "action=update&group=menglong\_group1&queue=menglong\_queue1&write=false" "http://127.0.0.1:8080/group"
{"action":"update","result":true}<br>

**查看业务方：** <br>
curl "http://127.0.0.1:8080/group?action=lookup" <br>
curl "http://127.0.0.1:8080/group?action=lookup&group=menglong\_group1"<br>


## 消息接口
**http://ip:port/message**
**参数列表：** <br>

| 参数名 | 是否必填 | 说明 | 
| ---- | ---- | ----|
| action | 必填 | 取消息/发消息/确认消息，参数为receive,send,ack |
| queue | 必填 | 队列名称 |
| group | 必填 | 业务名称 |
| msg | 必填 | 消息体，发送消息使用 |

**示例：** <br>
**发送消息：** <br>
curl -d "action=send&queue=remind&group=if&msg=helloworld" "http://127.0.0.1:8080/msg" <br>
{"action":"send","result":true} <br>

**接收消息：** <br>
curl "http://127.0.0.1:8080/msg?action=receive&queue=remind&group=if" <br>
{"action":"receive","msg":"helloworld2"} <br>

**确认消息：** <br>
curl -d "action=ack&queue=remind&group=if&id=xxxx" "http://127.0.0.1:8080/msg" <br>
{"action":"ack","result":true} <br>

## 监控接口
**http://ip:port/monitor** <br>
**参数列表：** <br>

| 参数名 | 是否必填 | 说明 | 
| ---- | ---- | ----|
| type | 必填 | send/receive 消息发送or接收 |
| group | 必填 | 业务标识，查看业务方时选填 |
| queue | 必填 | 队列名称 |
| start | 选填 | 开始时间，秒级时间戳 |
| end | 选填 | 结束时间，秒级时间戳 |
| interval | 选填 | 时间间隔，默认为1倍，1倍为10s，如果填2，即为20S |

**示例：** <br>
**消息发送量：** <br>
curl "http://127.0.0.1:8080/monitor?type=send&queue=menglong_queue1&group=menglong\_group1" <br>

**消息接收量：** <br>
curl "http://127.0.0.1:8080/monitor?type=receive&queue=menglong\_queue1&group=menglong\_group1" <br>

## 报警接口(定义中)
**http://ip:port/alarm** <br>
type：heap，send.second，receive.second <br>