package migrations

import (
	"github.com/baking-bad/bcdhub/internal/bcd/consts"
	"github.com/baking-bad/bcdhub/internal/config"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/models/migration"
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

	protocol, err := p.ctx.Cache.ProtocolByHash(p.network, head.Protocol)
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
		}
	}
	return parserResult, nil
}

func (p *ImplicitParser) origination(implicit noderpc.ImplicitOperationsResult, head noderpc.Header, protocolID int64, result *parsers.Result) error {
	result.Migrations = append(result.Migrations, &migration.Migration{
		Network:    p.network,
		ProtocolID: protocolID,
		Level:      head.Level,
		Timestamp:  head.Timestamp,
		Kind:       types.MigrationKindBootstrap,
		Address:    implicit.OriginatedContracts[0],
	})
	logger.Info().Msg("Implicit bootstrap migration found")

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

	script, err := p.rpc.GetRawScript(origination.Destination, origination.Level)
	if err != nil {
		return err
	}
	origination.Script = script

	contractParser := contract.NewParser(p.ctx)
	contractResult, err := contractParser.Parse(&origination)
	if err != nil {
		return err
	}
	if contractResult != nil {
		result.Merge(contractResult)
	}
	return nil
}
