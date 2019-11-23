# Homer API接口

1. 查询sip消息
- 说明  
执行查询操作

- URL描述  
  - /api/v1/search/data
  - Method:POST
  - Content type: application/json

- 输入参数:  

|  参数名  |  是否必填  |  数据类型  |     说明     |  
|----|----|----|----|
|   param    |     是    |    Json   |    查询条件                           |
| timestamp  |     是    |    Json   | 查询时间范围，以unix时间表示，单位为毫秒 |

- param格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|    --        |--         |--         |--            |
|transaction|是   |Json  |事务类型          |             
|limit      |是   |Int   |限制查询条数      |
|search     |是   |Json  |其它查询条件      |
|location   |是   |Json  |数据库节点        |
|timezone   |是   |Json  |时区             |

- transaction格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--        |--         |--         |--            |
|call         |是   |Bool  |是否查询与呼叫和通话相关的sip消息  |
|registration |是   |Bool  |是否查询与注册有关的sip消息        |
|rest         |是   |Bool  |是否查询其他sip消息               |

- search格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--        |--         |--         |--            |
|limit     |是   |String  |返回的结果条数          |
|callid    |否   |String  |callid                 |
|to_user   |否   |String  |被叫                   |
|from_user |否   |String  |主叫                   |

- location格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--        |--         |--         |--            |
|node   |是   |JsonArray   |数据库node节点         |

- node格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--       |--         |--         |--             |
|id       |是   |String   |node的id               |
|name     |是   |String   |node的名字             |

- timezone格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--        |--         |--         |--            |
|value   |是   |String |相对于GMT的偏移，单位为分钟数，如:"-480"  |
|name    |是   |String |时区名称,如"GMT+8 CCT"                  |
|offset  |是   |String |相对于GMT的偏移,如"+0800"               |

- timestamp格式定义  

|  参数名  |  是否必填  |  数据类型  |     说明     |
|--        |--         |--         |--            |
|from   |是  |Int  |开始时间，以unix时间表示，如:1535506514176|
|to     |是  |Int  |结束时间，以unix时间表示，如:1535507414176|

- 发送请求示例
```json
{
    "param": {
        "transaction": {
            "call": true,
            "registration": true,
            "rest": true
        },
        "limit": 200,
        "search": {
            "limit": "2000",
            "callid": "sdf",
            "to_user": "sdfs",
            "from_user": "sfsd"
        },
        "location": {
            "node": [
                {
                    "id": "1",
                    "name": "noded1"
                }
            ]
        },
        "timezone": {
            "value": "-480",
            "name": "GMT+8 CCT",
            "offset": "+0800"
        }
    },
    "timestamp": {
        "from": 1535506514176,
        "to": 1535507414176
    }
}
```


- 返回结果  

|  参数名  |  数据类型  |     说明     |
|--        |--         |--            |
|status  |Int       |返回的http状态码               |
|sid     |String    |唯一标识id                     |
|Auth    |String    |是否经过授权                   |
|message |String    |当前消息:"ok","wrong-session"  |
|data    |JsonArray |返回的数据                     |

- data格式定义  

|  参数名  |  数据类型  |     说明     |
|--        |--         |--            |
|id                  |String   |数据库中该条数据的编号
|date                |String   |该条数据的日期,如:"2018-08-28 15:20:32"
|milli_ts            |Int      |unix时间，毫秒为单位，如:1535440832935
|micro_ts            |Int      |unix时间，微秒为单位，如:1535440832935179
|method              |String   |sip方法，如:"INVITE","100"
|reply_reason        |String   |reply reason,如:"Ringing"
|ruri                |String   |request URI
|ruri_user           |String   |request URI中的用户名
|ruri_domain         |String   |request URI中的域
|from_user           |String   |主叫名字
|from_domain         |String   |主叫URI
|from_tag            |String   |From tag
|to_user             |String   |目的地用户名
|to_domain           |String   |目的地SIP URI
|to_tag              |String   |To tag
|pid_user            |String   |P-Asserted-Identity user
|contact_user        |String   |Contact user
|auth_user           |String   |Auth user
|callid              |String   |Callid
|callid_aleg         |String   |Callid of Aleg
|via_1               |String   |First via
|via_1_branch        |String   |First via brach
|cseq                |String   |Cseq,如:"1 INVITE"
|diversion           |String   |Diversion
|reason              |String   |
|content_type        |String   |
|auth                |String   |
|user_agent          |String   |用户代理工具名称
|source_ip           |String   |源ip地址
|source_port         |String   |源端口
|destination_ip      |String   |目的ip地址
|destination_port    |String   |目的端口
|contact_ip          |String   |contact的ip地址
|contact_port        |String   |contact的端口
|originator_ip       |String   |message的发起者的ip地址
|originator_port     |String   |message的发起者的端口
|expires             |String   |过期时间
|correlation_id      |String   |关联id
|proto               |String   |传输协议
|family              |String   |
|rtp_stat            |String   |rtp统计
|type                |String   |封装类型
|node                |String   |message的存储节点
|custom_field1       |String   |
|custom_field2       |String   |
|custom_field3       |String   |
|trans               |String   |message的消息类型
|dbnode              |String   |message的dbnode类型
|source_alias        |String   |源地址别名
|destination_alias   |String   |目的地址别名
|msg                 |String   |sip消息体

- 返回示例
```json
{
    "status": 200,
    "sid": "rdilbfpmbr7p95lsf6c0pv0634",
    "Auth": "true",
    "message": "ok",
    "data": [
        {
            "id": "1",
            "date": "2018-08-28 15:20:32",
            "milli_ts": 1535440832935,
            "micro_ts": 1535440832935179,
            "method": "100",
            "reply_reason": "Trying",
            "ruri": "",
            "ruri_user": "",
            "ruri_domain": "",
            "from_user": "3000",
            "from_domain": "10.10.5.80",
            "from_tag": "195",
            "to_user": "3001",
            "to_domain": "10.10.5.39",
            "to_tag": "",
            "pid_user": "",
            "contact_user": "",
            "auth_user": "",
            "callid": "195-14457@10.10.5.80",
            "callid_aleg": "",
            "via_1": "",
            "via_1_branch": "",
            "cseq": "1 INVITE",
            "diversion": "",
            "reason": "",
            "content_type": "",
            "auth": "",
            "user_agent": "",
            "source_ip": "10.10.5.39",
            "source_port": "32766",
            "destination_ip": "10.10.5.80",
            "destination_port": "17768",
            "contact_ip": "",
            "contact_port": "0",
            "originator_ip": "",
            "originator_port": "0",
            "expires": "-1",
            "correlation_id": "195-14457@10.10.5.80",
            "proto": "17",
            "family": "2",
            "rtp_stat": "",
            "type": "2",
            "node": "homer01:0",
            "custom_field1": "",
            "custom_field2": "",
            "custom_field3": "",
            "trans": "",
            "dbnode": "",
            "source_alias": "",
            "destination_alias": "",
            "msg": ""
        }
    ]
}
```