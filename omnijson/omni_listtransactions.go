package omnijson

/*
Result
[
{
    "txid": "c428b59a4ecd6b5f0525afdb42512777428d36bb38cb64c7bec4317c4f307db8",
    "fee": "0.00000405",
    "sendingaddress": "n1QF57ai9BKsJWFocqNdK726MiqW4tffqg",
    "referenceaddress": "mxzoNNDrGtYWzh3NFD6y6nNdeCkQbLKJNZ",
    "ismine": true,
    "version": 0,
    "type_int": 0,
    "type": "Simple Send",
    "propertyid": 1,
    "divisible": true,
    "amount": "1.00000000",
    "valid": true,
    "blockhash": "000000000000082e5a59ebc0d58830171ec3354d5fce9221828f8c52bca26c59",
    "blocktime": 1561607522,
    "positioninblock": 46,
    "block": 1565779,
    "confirmations": 2279
  },
]
*/
type OmniListTransactionsResult = [] struct {
	Txid             string  `json:"txid"`
	SendingAddress   string  `json:"sendingaddress"`
	ReferenceAddress string  `json:"referenceaddress"`
	BlockHash        string  `json:"blockhash"`
	Type             string  `json:"type"`
	Fee              float64 `json:"fee"`
	Amount           float64 `json:"amount"`
	BlockTime        int64   `json:"blocktime"`
	Version          uint32  `json:"version"`
	TypeInt          uint32  `json:"type_int"`
	PropertyID       uint32  `json:"propertyid"`
	Confirmations    uint32  `json:"confirmations"`
	Block            int32   `json:"block"`
	PositionInBlock  int     `json:"positioninblock"`
	IsMine           bool    `json:"ismine"`
	Divisible        bool    `json:"divisible"`
	Valid            bool    `json:"valid"`
}

type OmniListTransactionsCommand struct {
	Txid       string
	Count      int
	Skip       int
	StartBlock int
	EndBlock   int
}

func (OmniListTransactionsCommand) Method() string {
	return "omni_listtransactions"
}

func (OmniListTransactionsCommand) ID() string {
	return "1"
}

func (cmd OmniListTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.Txid, cmd.Count, cmd.Skip, cmd.StartBlock, cmd.EndBlock}
}

func NewOmniListTransactionsCommand(count, startBlock int) OmniListTransactionsCommand {
	return OmniListTransactionsCommand{
		Txid:       "*",
		Count:      count,
		Skip:       0,
		StartBlock: startBlock,
		EndBlock:   999999999,
	}
}
