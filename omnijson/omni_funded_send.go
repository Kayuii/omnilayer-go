package omnijson

//OmniFundedSendResult ...
type OmniFundedSendResult string

//OmniFoundedSendCommand ...
type OmniFundedSendCommand struct {
	From     string `json:"fromaddress"`
	To       string `json:"toaddress"`
	ProperID int64  `json:"propertyid"`
	Amount   string `json:"amount"`
	Fee      string `json:"feeaddress"`
}

//Method ...
func (OmniFundedSendCommand) Method() string {
	return "omni_funded_send"
}

//ID ...
func (OmniFundedSendCommand) ID() string {
	return "1"
}

//Params ...
func (cmd OmniFundedSendCommand) Params() []interface{} {
	return []interface{}{cmd.From, cmd.To, cmd.ProperID, cmd.Amount, cmd.Fee}
}
