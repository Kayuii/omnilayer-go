package omnilayer

import (
	"encoding/json"

	"github.com/kayuii/omnilayer-go/omnijson"
)

type futureOmniFoundedSend chan *response

func (f futureOmniFoundedSend) Receive() (omnijson.OmniFoundedSendResult, error) {
	var result omnijson.OmniFoundedSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreatePayloadSimpleSend chan *response

func (f futureOmniCreatePayloadSimpleSend) Receive() (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	var result omnijson.OmniCreatePayloadSimpleSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxChange chan *response

func (f futureOmniCreateRawTxChange) Receive() (omnijson.OmniCreateRawTxChangeResult, error) {
	var result omnijson.OmniCreateRawTxChangeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxOpReturn chan *response

func (f futureOmniCreateRawTxOpReturn) Receive() (omnijson.OmniCreateRawTxOpReturnResult, error) {
	var result omnijson.OmniCreateRawTxOpReturnResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxReference chan *response

func (f futureOmniCreateRawTxReference) Receive() (omnijson.OmniCreateRawTxReferenceResult, error) {
	var result omnijson.OmniCreateRawTxReferenceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetInfo chan *response

func (f futureGetInfo) Receive() (omnijson.OmniGetInfoResult, error) {
	var result omnijson.OmniGetInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetTransaction chan *response

func (f futureOmniGetTransaction) Receive() (omnijson.OmniGettransactionResult, error) {
	var result omnijson.OmniGettransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniListBlockTransactions chan *response

func (f futureOmniListBlockTransactions) Receive() (omnijson.OmniListBlockTransactionsResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(omnijson.OmniListBlockTransactionsResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetBalance chan *response

func (f futureOmniGetBalance) Receive() (omnijson.OmniGetBalanceResult, error) {
	var result omnijson.OmniGetBalanceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniSend chan *response

func (f futureOmniSend) Receive() (omnijson.OmniSendResult, error) {
	var result omnijson.OmniSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniFundedSend chan *response

func (f futureOmniFundedSend) Receive() (omnijson.OmniFundedSendResult, error) {
	var result omnijson.OmniFundedSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniListTransactions chan *response

func (f futureOmniListTransactions) Receive() (omnijson.OmniListTransactionsResult, error) {
	var result omnijson.OmniListTransactionsResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
