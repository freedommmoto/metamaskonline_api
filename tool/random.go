package tool

import (
	"math/rand"
	"strings"
)

const alphabetandint = "a1s2d3f4h5s6d7k8f9j0h1w2e3u4r5g6b7r8o9q0w1h2s3d4h5d6g7t8y9p0h1j2f3g4v"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetandint)
	for i := 0; i < n; i++ {
		c := alphabetandint[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomWallet() string {
	return RandomString(62)
}

func GetMockOneBlockData() string {
	return "{\"blockNumber\":\"21046285\",\"timeStamp\":\"1657783148\",\"hash\":\"0xcb60fae3d6ee35e9d34246dbc20295105aaca29029c2b92d9bc0631d1985e3f2\",\"nonce\":\"4\",\"blockHash\":\"0x57a550e421d0fb0ef0e53a85a64210544c4169c068404141aff3298d49d5a6da\",\"transactionIndex\":\"5\",\"from\":\"0x891b68d6b21c64d56db262d066b38ea76b6468f6\",\"to\":\"0x85956f45e5439c15441868f734d09ca8d85133e5\",\"value\":\"2340000000000000\",\"gas\":\"21000\",\"gasPrice\":\"10000000000\",\"isError\":\"0\",\"txreceipt_status\":\"1\",\"input\":\"0x\",\"contractAddress\":\"\",\"cumulativeGasUsed\":\"1989149\",\"gasUsed\":\"21000\",\"confirmations\":\"22341\"}"
}

func GetMockOneLineRquest() string {
	return "{\n    \"to\": \"Ue5308cc32ee5ca607c596e87877715b6\",\n    \"messages\":[\n        {\n            \"type\":\"text\",\n            \"text\":\"Hello, world1\"\n        },\n        {\n            \"type\":\"text\",\n            \"text\":\"Hello, world2\"\n        }\n    ]\n}"
}
