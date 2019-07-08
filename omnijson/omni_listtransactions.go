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
type OmniListTransactionsResult = []struct {
	Txid             string  `json:"txid"`
	Fee              float64 `json:"fee"`
	SendingAddress   string  `json:"sendingaddress"`
	ReferenceAddress string  `json:"referenceaddress"`
	IsMine           bool    `json:"ismine"`
	Version          int     `json:"version"`
	TypeInt          int     `json:"type_int"`
	Type             string  `json:"type"`
	PropertyId       int     `json:"propertyid"`
	Divisible        bool    `json:"divisible"`
	Amount           float64 `json:"amount"`
	Valid            bool    `json:"valid"`
	BlockHash        string  `json:"blockhash"`
	BlockTime        int64   `json:"blocktime"`
	PositionInBlock  int     `json:"positioninblock"`
	Block            int64   `json:"block"`
	Confirmations    int     `json:"confirmations"`
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
