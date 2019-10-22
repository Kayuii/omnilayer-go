package omnijson

type ImportAddressCommand struct {
	Address  	string
	Label    	string
	Rescan 		bool
}

func (ImportAddressCommand) Method() string {
	return "importaddress"
}

func (ImportAddressCommand) ID() string {
	return "1"
}

func (cmd ImportAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Address, cmd.Label, cmd.Rescan}
}
