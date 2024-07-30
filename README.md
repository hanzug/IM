

高性能通用通讯服务，支持即时通讯，站内/系统消息，消息中台，物联网通讯，音视频信令，直播弹幕，客服系统，AI通讯，即时社区等场景。

`本项目需要在go1.20.0或以上环境编译`



分布式IM重要特性： 故障自动转移，去中心化设计，节点之间数据互备，支持集群快速自动扩容，代理节点机制


![](http://localhost:63342/markdownPreview/713228306/docs/logo.png?_ijt=jplti0t3trcq6irctv35vaufb0)

- **官网**: https://githubim.com
- **通讯协议**: [WuKongIM协议](https://githubim.com/guide/proto.html)
- **提问**: https://github.com/WuKongIM/WuKongIM/issues
- **文档**: https://githubim.com

[![](https://img.shields.io/github/license/WuKongIM/WuKongIM?color=yellow&style=flat-square)](file:///D:/github/IM/LICENSE) [![](https://img.shields.io/badge/go-%3E%3D1.20-30dff3?style=flat-square&logo=go)](https://github.com/WuKongIM/WuKongIM) [![](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/WuKongIM/WuKongIM) [![](https://img.shields.io/badge/Slack-99%2B-blueviolet?logo=slack&logoColor=white)](https://join.slack.com/t/wukongim/shared_invite/zt-22o7we8on-2iKNUmgigB9ERdF9XUivmw)

## 架构图

![](http://localhost:63342/markdownPreview/713228306/docs/architecture/cluster.png?_ijt=jplti0t3trcq6irctv35vaufb0)

## 节点故障转移演示

![](http://localhost:63342/markdownPreview/713228306/docs/architecture/cluster-failover.webp?_ijt=jplti0t3trcq6irctv35vaufb0)

## 演示

**聊天Demo**

web聊天场景演示： [http://imdemo.githubim.com](http://imdemo.githubim.com/)

后端监控演示： [http://monitor.githubim.com/web](http://monitor.githubim.com/web)


## 特点

🎦**独特性**

群成员无上限，轻松支持10万人群聊，消息可永久存储。

📚**资源消耗低**

自研二进制协议，心跳包只有1字节，省流量，省电量，传输更迅速。

🔐**安全性**

消息通道和消息内容全程加密，防中间人攻击和窜改消息内容，服务端数据实时备份，数据不丢失。

🚀 **性能**

基于pebble kv数据库，研发了针对于IM这种服务的特有分布式数据库，省了其他数据库为了通用性而带来的性能损耗， 因为存储快，所以消息快。

🔥**高可用**

通过魔改raft分布式协议，实现了自动容灾，一台机器宕机，另一台机器自动接管，对外无感知。

去中心化，无单点，无中心节点，每个节点都是独立且平等的，都可以提供服务。

扩容方便，只需增加机器，不需要停机，不需要迁移数据，自动按策略分配数据。

## 功能特性

- [x] 支持自定义消息
- [x] 支持订阅/发布者模式
- [x] 支持个人/群聊/客服/社区资讯频道
- [x] 支持频道黑明单
- [x] 支持频道白名单
- [x] 支持消息永久漫游，换设备登录，消息不丢失
- [x] 支持在线状态，支持同账号多设备同时在线
- [x] 支持多设备消息实时同步
- [x] 支持用户最近会话列表服务端维护
- [x] 支持指令消息
- [x] 支持离线指令接口
- [x] 支持Webhook，轻松对接自己的业务系统
- [x] 支持Datasource，无缝对接自己的业务系统数据源
- [x] 支持Websocket连接
- [x] 支持TLS 1.3
- [x] 支持Prometheus监控
- [x] 监控系统开发
- [x] 支持Windows系统(仅开发用)
- [x] 支持流式消息，类似chatgpt的结果输出流
- [x] 支持分布式
  - [x] 去中心化设计，任意一个节点宕机，集群自动修复
  - [x] 集群节点之间数据互备，任意一个节点损害，不影响数据完整性
  - [x] 支持集群快速自动扩容
  - [ ] 支持长连接CDN，解决跨国跨地区长连接不稳定问题

## 快速运行

### Docker部署（单机）
```shell
docker run -d -p 15001:5001 -p 15100:5100 -p 15172:5172 -p 15200:5200 -p 15210:5210 -p 15300:5300  --name wukongim -v ./wukongim:/root/wukongim  wukongim/wukongim:v2.0.0-beta-20240428
```

### Docker部署（分布式）
```shell
git clone https://github.com/WuKongIM/WuKongIM.git

cd ./WuKongIM/docker/cluster

sudo docker compose up -d
```

## 源码开发

### 单机

```shell
go run main.go

(或 go run main.go --config config/wk.yaml)

```

### 分布式

```yaml
# 启动第一个节点
go run main.go --config  ./exampleconfig/cluster1.yaml

# 启动第二个节点
go run main.go --config  ./exampleconfig/cluster2.yaml

# 启动第三个节点
go run main.go --config  ./exampleconfig/cluster3.yaml

```

### 访问

查询系统信息: [http://127.0.0.1:15001/varz](http://127.0.0.1:15001/varz)

查看监控信息： [http://127.0.0.1:15300/web](http://127.0.0.1:15300/web)

客户端演示地址：[http://127.0.0.1:15172/chatdemo](http://127.0.0.1:15172/chatdemo) (分布式地址为：[http://127.0.0.1:15172/login](http://127.0.0.1:15172/login))

端口解释:

`15001: api端口 15100: tcp长连接端口 15172: demo端口 15200: websocket长连接端口 15300: 监控系统端口`



## 图解

总体架构图


![](http://localhost:63342/markdownPreview/713228306/docs/architecture/architecture2.png?_ijt=jplti0t3trcq6irctv35vaufb0)

业务系统对接

![[Pasted image 20240731070930.png]]

Webhook对接图

![image](http://localhost:63342/markdownPreview/713228306/docs/webhook.png?_ijt=jplti0t3trcq6irctv35vaufb0)

## 适用场景

#### 即时通讯

- 群频道支持
- 个人频道支持
- 消息永久存储
- 离线消息推送支持
- 最近会话维护

#### 消息推送/站内消息

- 群频道支持
- 个人频道支持
- 离线消息推送支持

#### 物联网通讯

- mqtt协议支持（待开发）
- 支持发布与订阅

#### 音视频信令服务器

- 支持临时指令消息投递

#### 直播弹幕

- 临时消息投递

- 临时订阅者支持


#### 客服系统

- 客服频道支持

- 消息支持投递给第三方服务器

- 第三方服务器可决定分配指定的订阅者成组投递


#### 实时AI反馈

- 支持客户端发的消息推送给第三方服务器，第三方服务器反馈给AI后返回的结果再推送给客户端

#### 即时社区

- 社区频道支持
- 支持topic模式的消息投递

## 监控截图

![image](http://localhost:63342/markdownPreview/713228306/docs/screen1.png?_ijt=jplti0t3trcq6irctv35vaufb0) ![image](http://localhost:63342/markdownPreview/713228306/docs/screen2.png?_ijt=jplti0t3trcq6irctv35vaufb0) ![image](http://localhost:63342/markdownPreview/713228306/docs/screen3.png?_ijt=jplti0t3trcq6irctv35vaufb0) ![image](http://localhost:63342/markdownPreview/713228306/docs/screen4.png?_ijt=jplti0t3trcq6irctv35vaufb0) ![image](http://localhost:63342/markdownPreview/713228306/docs/screen5.png?_ijt=jplti0t3trcq6irctv35vaufb0)

## Star

我们团队一直致力于即时通讯的研发，需要您的鼓励，如果您觉得本项目对您有帮助，欢迎点个star，您的支持是我们最大的动力。

## 案例展示

**项目名**

TangSengDaoDao

**开源地址**

[https://github.com/TangSengDaoDao/TangSengDaoDaoServer](https://github.com/TangSengDaoDao/TangSengDaoDaoServer)

**截图**


|![](http://localhost:63342/markdownPreview/713228306/docs/case/tsdaodao/screenshot/conversationlist.webp?_ijt=jplti0t3trcq6irctv35vaufb0)|![](http://localhost:63342/markdownPreview/713228306/docs/case/tsdaodao/screenshot/messages.webp?_ijt=jplti0t3trcq6irctv35vaufb0)|![](http://localhost:63342/markdownPreview/713228306/docs/case/tsdaodao/screenshot/robot.webp?_ijt=jplti0t3trcq6irctv35vaufb0)|



![](http://localhost:63342/markdownPreview/713228306/docs/case/tsdaodao/screenshot/pc11.png?_ijt=jplti0t3trcq6irctv35vaufb0)
