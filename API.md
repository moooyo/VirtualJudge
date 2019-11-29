# API 使用指南

项目提供rest风格API，下文中出现的通用标识符含义为

| 标识符        | 含义                      |
| ------------- | ------------------------- |
| `RootURL`     | 项目部署的根地址          |
| `SubURL`      | 分模块定义的URL           |
| `ParentURL`   | 表示分模块的父模块的URL   |
| `FunctionURL` | 定义某模块URL，项目中唯一 |
| `{id}`        | ${id}属性                 |

对于整个项目拥有一个唯一的`RootURL`，对于不同子模块

## 提交到Virtual Judge

提交代码到对应的题目

| type     | value            |
| -------- | ---------------- |
| method   | POST             |
| `SubURL` | `RootURL/submit` |

**Request**

Virtual Judge只接收单次提交单份代码，提交格式为

```json
{
 "uid":1, 
 "sid":1, 
 "oj":1, 
 "problem":1001, 
 "language":1, 
 "source":"#include<iostream>...."
}
```

**Response**

Status: 200 OK

***

```json
{
    "status": "$code",
    "message": ""
}
```



## 获取Virtual Judge的评测情况

使用submit id查询评测情况

| type     | value                  |
| -------- | ---------------------- |
| method   | GET                    |
| `SubURL` | `RootURL/submit/{sid}` |

**Request**

```json
{
    "sid": "$sid"
}
```

**Response**

Status: 200 OK

***

```json
{
    "uid": 1001,
    "oj": 1,
    "problem": 1001,
    "status": "$StatusCode"
}
```

