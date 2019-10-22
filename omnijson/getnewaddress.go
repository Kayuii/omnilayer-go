package omnijson

/*
Result
{
	"result":"mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4",
	"error":null,
	"id":"curltext"
}
*/

//GetNewAddressResult ...
type GetNewAddressResult = string

//GetNewAddressCommand ...
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
