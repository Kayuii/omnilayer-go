package omnijson

type OmniSendResult string

type OmniSendCommand struct {
	FromAddress     string
	ToAddress       string
	PropertyId      int
	Amount          string
	RedeemAddress   string
	ReferenceAmount string
}

func (OmniSendCommand) Method() string {
	return "omni_send"
}

func (OmniSendCommand) ID() string {
	return "1"
}

func (cmd OmniSendCommand) Params() []interface{} {
	ret := []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyId, cmd.Amount, cmd.RedeemAddress, cmd.ReferenceAmount}
	if cmd.RedeemAddress != "" {
		ret = append(ret, cmd.RedeemAddress)
		if cmd.ReferenceAmount != "" {
			ret = append(ret, cmd.ReferenceAmount)
		}
	}
	return ret
}
