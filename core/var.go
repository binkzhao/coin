package core

import "math"

const dbFile       = "blockChain.db"
const blocksBucket = "blocks" // 区块链在数据库里面的键
const subsidy      = 10 // 一个区块币的数量
const maxNonce     = math.MaxInt64 // nonce计算器最大值
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// 目前我们并不会实现一个动态调整目标的算法，所以将难度定义为一个全局的常量即可
var targetBits = 16 // 挖矿难度值, 值越大，挖矿越难
