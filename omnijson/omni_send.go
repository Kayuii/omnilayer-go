package omnijson

type OmniSendResult string

type OmniSendCommand struct {

	FromAddress     string
	ToAddress       string
	PropertyID      int
	Amount          string
	RedeemAddress   string
	ReferenceAmount string
	// FromAddress string
	// ToAddress   string
	// PropertyId  uint32
	// Amount      string
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
	ret := []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyID, cmd.Amount}
	if cmd.RedeemAddress != "" {
		ret = append(ret, cmd.RedeemAddress)
		if cmd.ReferenceAmount != "" {
			ret = append(ret, cmd.ReferenceAmount)
		}
	}
	return ret
}
