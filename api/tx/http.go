package tx

type QueryInternalTx struct {
	TxHash  string `query:"txHash"`
	JsonRPC string `query:"jsonRpc"`
}
