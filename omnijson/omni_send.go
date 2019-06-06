package omnijson

type OmniSendResult string

type OmniSendCommand struct {
	FromAddress string
	ToAddress   string
	PropertyId  uint
	Amount      string
	//RedeemAddress         string
	//ReferenceAmount       float64
}

func (OmniSendCommand) Method() string {
	return "omni_send"
}

func (OmniSendCommand) ID() string {
	return "1"
}

func (cmd OmniSendCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyId, cmd.Amount}
}
