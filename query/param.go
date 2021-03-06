package query

import (
	"github.com/lino-network/lino-go/model"
)

// GetEvaluateOfContentValueParam returns the EvaluateOfContentValueParam.
func (query *Query) GetEvaluateOfContentValueParam() (*model.EvaluateOfContentValueParam, error) {
	resp, err := query.transport.Query(getEvaluateOfContentValueParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.EvaluateOfContentValueParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetGlobalAllocationParam returns the GlobalAllocationParam.
func (query *Query) GetGlobalAllocationParam() (*model.GlobalAllocationParam, error) {
	resp, err := query.transport.Query(getGlobalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.GlobalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetInfraInternalAllocationParam returns the InfraInternalAllocationParam.
func (query *Query) GetInfraInternalAllocationParam() (*model.InfraInternalAllocationParam, error) {
	resp, err := query.transport.Query(getInfraInternalAllocationParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.InfraInternalAllocationParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetDeveloperParam returns the DeveloperParam.
func (query *Query) GetDeveloperParam() (*model.DeveloperParam, error) {
	resp, err := query.transport.Query(getDeveloperParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.DeveloperParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetVoteParam returns the VoteParam.
func (query *Query) GetVoteParam() (*model.VoteParam, error) {
	resp, err := query.transport.Query(getVoteParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.VoteParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetProposalParam returns the ProposalParam.
func (query *Query) GetProposalParam() (*model.ProposalParam, error) {
	resp, err := query.transport.Query(getProposalParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ProposalParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetValidatorParam returns the ValidatorParam.
func (query *Query) GetValidatorParam() (*model.ValidatorParam, error) {
	resp, err := query.transport.Query(getValidatorParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.ValidatorParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetCoinDayParam returns the CoinDayParam.
func (query *Query) GetCoinDayParam() (*model.CoinDayParam, error) {
	resp, err := query.transport.Query(getCoinDayParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.CoinDayParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetBandwidthParam returns the BandwidthParam.
func (query *Query) GetBandwidthParam() (*model.BandwidthParam, error) {
	resp, err := query.transport.Query(getBandwidthParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.BandwidthParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetAccountParam returns the AccountParam.
func (query *Query) GetAccountParam() (*model.AccountParam, error) {
	resp, err := query.transport.Query(getAccountParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.AccountParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}

// GetPostParam returns the PostParam.
func (query *Query) GetPostParam() (*model.PostParam, error) {
	resp, err := query.transport.Query(getPostParamKey(), ParamKVStoreKey)
	if err != nil {
		return nil, err
	}

	param := new(model.PostParam)
	if err := query.transport.Cdc.UnmarshalJSON(resp, param); err != nil {
		return nil, err
	}
	return param, nil
}
