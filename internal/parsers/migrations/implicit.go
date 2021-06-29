package migrations

import (
	"encoding/json"

	"github.com/baking-bad/bcdhub/internal/bcd/ast"
	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/bigmap"
	"github.com/baking-bad/bcdhub/internal/models/operation"
	"github.com/baking-bad/bcdhub/internal/models/types"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/baking-bad/bcdhub/internal/parsers"
	"github.com/baking-bad/bcdhub/internal/parsers/contract"
)

// ImplicitParser -
type ImplicitParser struct {
	ctx     *config.Context
	network types.Network
	rpc     noderpc.INode
}

// NewImplicitParser -
func NewImplicitParser(ctx *config.Context, network types.Network, rpc noderpc.INode) *ImplicitParser {
	return &ImplicitParser{ctx, network, rpc}
}

// Parse -
func (p *ImplicitParser) Parse(metadata noderpc.Metadata, head noderpc.Header) (*parsers.Result, error) {
	if len(metadata.ImplicitOperationsResults) == 0 {
		return nil, nil
	}

	protocol, err := p.ctx.CachedProtocolByHash(p.network, head.Protocol)
	if err != nil {
		return nil, err
	}

	parserResult := parsers.NewResult()
	for i := range metadata.ImplicitOperationsResults {
		switch metadata.ImplicitOperationsResults[i].Kind {
		case consts.Origination:
			if err := p.origination(metadata.ImplicitOperationsResults[i], head, protocol.ID, parserResult); err != nil {
				return nil, err
			}
		case consts.Transaction:
			p.transaction(metadata.ImplicitOperationsResults[i], head, protocol.ID, parserResult)
		}
	}
	return parserResult, nil
}

func (p *ImplicitParser) transaction(implicit noderpc.ImplicitOperationsResult, head noderpc.Header, protocolID int64, result *parsers.Result) {
	result.Operations = append(result.Operations, &operation.Operation{
		Network:             p.network,
		ProtocolID:          protocolID,
		Level:               head.Level,
		Timestamp:           head.Timestamp,
		Kind:                types.NewOperationKind(implicit.Kind),
		ConsumedGas:         implicit.ConsumedGas,
		PaidStorageSizeDiff: implicit.PaidStorageSizeDiff,
		StorageSize:         implicit.StorageSize,
		DeffatedStorage:     implicit.Storage,
	})
	logger.Info().Str("kind", consts.Transaction).Msg("Implicit operation found")
}

func (p *ImplicitParser) origination(implicit noderpc.ImplicitOperationsResult, head noderpc.Header, protocolID int64, result *parsers.Result) error {
	origination := operation.Operation{
		Network:             p.network,
		ProtocolID:          protocolID,
		Level:               head.Level,
		Timestamp:           head.Timestamp,
		Kind:                types.NewOperationKind(implicit.Kind),
		Destination:         implicit.OriginatedContracts[0],
		ConsumedGas:         implicit.ConsumedGas,
		PaidStorageSizeDiff: implicit.PaidStorageSizeDiff,
		StorageSize:         implicit.StorageSize,
		DeffatedStorage:     implicit.Storage,
	}

	logger.Info().Str("address", origination.Destination).Str("kind", consts.Origination).Msg("Implicit operation found")

	script, err := p.rpc.GetRawScript(origination.Destination, origination.Level)
	if err != nil {
		return err
	}
	origination.Script = script

	contractParser := contract.NewParser(p.ctx, contract.WithShareDir(p.ctx.SharePath))
	contractResult, err := contractParser.Parse(&origination)
	if err != nil {
		return err
	}
	if contractResult != nil {
		result.Merge(contractResult)
	}

	if origination.DeffatedStorage == nil {
		rawStorage, err := p.rpc.GetScriptStorageRaw(origination.Destination, origination.Level)
		if err != nil {
			return err
		}
		origination.DeffatedStorage = rawStorage
	}

	result.Operations = append(result.Operations, &origination)

	origination.AST, err = ast.NewScript(origination.Script)
	if err != nil {
		return err
	}

	storageType, err := origination.AST.StorageType()
	if err != nil {
		return err
	}

	if err := storageType.SettleFromBytes(origination.DeffatedStorage); err != nil {
		return err
	}

	ptrs := storageType.FindBigMapByPtr()
	for ptr, bigMap := range ptrs {
		keyType, err := json.Marshal(bigMap.KeyType)
		if err != nil {
			return err
		}
		valueType, err := json.Marshal(bigMap.ValueType)
		if err != nil {
			return err
		}
		result.BigMaps = append(result.BigMaps, &bigmap.BigMap{
			Network:      origination.Network,
			Contract:     origination.Destination,
			Ptr:          ptr,
			KeyType:      keyType,
			ValueType:    valueType,
			Name:         bigMap.FieldName,
			CreatedLevel: head.Level,
			CreatedAt:    head.Timestamp,
		})
	}

	return nil
}
