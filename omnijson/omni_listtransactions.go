package omnijson

type OmniListTransactionsResult = []OmniGettransactionResult

type OmniListTransactionsCommand struct {
	Txid       string
	Count      int64
	Skip       int64
	Startblock int64
	Endblock   int64
}

func (OmniListTransactionsCommand) Method() string {
	return "omni_listtransactions"
}

func (OmniListTransactionsCommand) ID() string {
	return "1"
}

func (cmd OmniListTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.Txid, cmd.Count, cmd.Skip, cmd.Startblock, cmd.Endblock}
}
