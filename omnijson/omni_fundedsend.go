package omnijson

type OmniFundedSendResult string

type OmniFundedSendCommand struct {
	FromAddress string
	ToAddress   string
	PropertyId  uint
	Amount      string
	FeeAddress  string
}

func (OmniFundedSendCommand) Method() string {
	return "omni_funded_send"
}

func (OmniFundedSendCommand) ID() string {
	return "1"
}

func (cmd OmniFundedSendCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyId, cmd.Amount, cmd.FeeAddress}
}
