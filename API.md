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
    "message": "ProblemInfo"
}
```

`ProblemInfo`

***

```go
package dispatch
type ProblemInfo struct {
	OJ           int
	ProblemID    int
	ProblemName  string
	Description  string
	Source       string
	TimeLimit    string
	MemoryLimit  string
	Input        string
	Output       string
	SampleInput  []string
	SampleOutput []string
	Language     []struct {
		TypeName string
		ID       int
	}
}
```

##### Sample

GET				 `SubURL/1/1004`

Response

```json
{
    "message": {
        "OJ": 1,
        "ProblemID": 1004,
        "ProblemName": "Financial Management",
        "Description": "Larry graduated this year and finally has a job. He's making a lot of money, but somehow never seems to have enough. Larry has decided that he needs to grab hold of his financial portfolio and solve his financing problems. The first step is to figure out what's been going on with his money. Larry has his bank account statements and wants to see how much money he has. Help Larry by writing a program to take his closing balance from each of the past twelve months and calculate his average account balance. ",
        "Source": "Mid-Atlantic 2001",
        "TimeLimit": "1000MS",
        "MemoryLimit": "10000K",
        "Input": "The input will be twelve lines. Each line will contain the closing balance of his bank account for a particular month. Each number will be positive and displayed to the penny. No dollar sign will be included. ",
        "Output": "The output will be a single number, the average (mean) of the closing balances for the twelve months. It will be rounded to the nearest penny, preceded immediately by a dollar sign, and followed by the end-of-line. There will be no other spaces or characters in the output. ",
        "SampleInput": [
            "100.00\n489.12\n12454.12\n1234.10\n823.05\n109.20\n5.27\n1542.25\n839.18\n83.99\n1295.01\n1.75"
        ],
        "SampleOutput": [
            "$1581.42"
        ],
        "Language": [
            {
                "TypeName": "G++",
                "ID": 0
            },
            {
                "TypeName": "GCC",
                "ID": 1
            },
            {
                "TypeName": "Java",
                "ID": 2
            },
            {
                "TypeName": "Pascal",
                "ID": 3
            },
            {
                "TypeName": "C++",
                "ID": 4
            },
            {
                "TypeName": "C",
                "ID": 5
            },
            {
                "TypeName": "Fortran",
                "ID": 6
            }
        ]
    },
    "status": 0
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

### 获取Virtual Judge的题目信息

需要提供OJ和题号查看题目信息

| type     | value                   |
| -------- | ----------------------- |
| method   | GET                     |
| `SubURL` | `RootURL/problems`      |
| `URL`    | `SubURL/:OJ/:ProblemID` |

**Request**

***

```json

```

**Response**

***

```json
{
    
}
```





