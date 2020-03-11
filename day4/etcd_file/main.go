package main

import (
	"sync"
)

/*
etcd常用场景:服务注册&发现/全局配置/分布式锁
特点:1.部署简单, 2.基于http提供接口  3.基于raft算法保证强一致性  4.天生就是为来了集群化设计
基于raft协议:1.选举/2.过期/3.异常处理

raft选举: 基于大多数原则.
选举过程中为3个角色follower(跟随者),candidate(候选人),leader(领导者).
选举过程:
1.leader下线后,所有follower变为candidate进行拉票(每个节点150-300ms的随机时间达到错开每个节点成为candidate的时间)成为候选人后首先投给自己然后进行拉票请求
2.当其中一个candidate首先得到超过半数以上的投票则该节点选举成功成为leader,
3.其余candidate重新变回follower
4.leader掉线后集群节点收到leader的心跳超时将重新触发选举
*/

/*
简单实现raft选举
*/
const raftCount = 3 // 节点数

// leader对象
type leader struct {
	Term     int // 任期
	LeaderId int
}

type raft struct {
	mu              sync.Mutex // 锁
	me              int        // 节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为哪个节点投票
	state           int        // 状态: 1:follower;2:candidate;3:leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 当前节点leader
	message         chan bool  // 节点间发数据的管道
	electCh         chan bool  // 选举管道
	heartBeat       chan bool  // 心跳管道
	heartBeatRe     chan bool  // 心跳返回管道
	timeOut         int        // 超时时间(150ms~300ms随机值)
}
