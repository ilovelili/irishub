package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

// RegisterCodec registers concrete types on the codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgDefineService{}, "irishub/service/MsgDefineService", nil)
	cdc.RegisterConcrete(MsgSvcBind{}, "irishub/service/MsgSvcBinding", nil)
	cdc.RegisterConcrete(MsgSvcBindingUpdate{}, "irishub/service/MsgSvcBindingUpdate", nil)
	cdc.RegisterConcrete(MsgSvcDisable{}, "irishub/service/MsgSvcDisable", nil)
	cdc.RegisterConcrete(MsgSvcEnable{}, "irishub/service/MsgSvcEnable", nil)
	cdc.RegisterConcrete(MsgSvcRefundDeposit{}, "irishub/service/MsgSvcRefundDeposit", nil)
	cdc.RegisterConcrete(MsgSvcRequest{}, "irishub/service/MsgSvcRequest", nil)
	cdc.RegisterConcrete(MsgSvcResponse{}, "irishub/service/MsgSvcResponse", nil)
	cdc.RegisterConcrete(MsgSvcRefundFees{}, "irishub/service/MsgSvcRefundFees", nil)
	cdc.RegisterConcrete(MsgSvcWithdrawFees{}, "irishub/service/MsgSvcWithdrawFees", nil)
	cdc.RegisterConcrete(MsgSvcWithdrawTax{}, "irishub/service/MsgSvcWithdrawTax", nil)
	cdc.RegisterConcrete(ServiceDefinition{}, "irishub/service/ServiceDefinition", nil)
	cdc.RegisterConcrete(SvcBinding{}, "irishub/service/SvcBinding", nil)
	cdc.RegisterConcrete(SvcRequest{}, "irishub/service/SvcRequest", nil)
	cdc.RegisterConcrete(SvcResponse{}, "irishub/service/SvcResponse", nil)
	cdc.RegisterConcrete(IncomingFee{}, "irishub/service/IncomingFee", nil)
	cdc.RegisterConcrete(ReturnedFee{}, "irishub/service/ReturnedFee", nil)
	cdc.RegisterConcrete(&Params{}, "irishub/service/Params", nil)
}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}