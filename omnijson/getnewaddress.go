package omnijson

type GetNewAddressResult = string

type GetNewAddressCommand struct {
	Account string
}

func (GetNewAddressCommand) Method() string {
	return "getnewaddress"
}

func (GetNewAddressCommand) ID() string {
	return "1"
}

func (cmd GetNewAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Account}
}
