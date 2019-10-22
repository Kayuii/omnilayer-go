package omnijson

type GetBlockCountResult = int64

type GetBlockCountCmd struct{}

func (GetBlockCountCmd) Method() string {
	return "getblockcount"
}

func (GetBlockCountCmd) ID() string {
	return "1"
}

func (GetBlockCountCmd) Params() []interface{} {
	return nil
}
