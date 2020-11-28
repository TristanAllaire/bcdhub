package events

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/noderpc"
	"github.com/tidwall/gjson"
)

// Event -
type Event interface {
	GetCode() (gjson.Result, error)
	Parse(response gjson.Result) []TokenBalance
	Normalize(parameter string) gjson.Result
}

// Context -
type Context struct {
	Network                  string
	Parameters               string
	Source                   string
	Initiator                string
	Entrypoint               string
	ChainID                  string
	HardGasLimitPerOperation int64
	Amount                   int64
}

// Sections -
type Sections struct {
	Parameter  json.RawMessage
	ReturnType json.RawMessage
	Code       json.RawMessage
}

// TokenBalance -
type TokenBalance struct {
	Address string
	TokenID int64
	Value   int64
}

// GetCode -
func (sections Sections) GetCode() (gjson.Result, error) {
	return gjson.Parse(fmt.Sprintf(`[{
		"prim": "parameter",
		"args": [%s]
	},{
		"prim": "storage",
		"args": [%s]
	},{
		"prim": "code",
		"args": [%s]
	}]`, string(sections.Parameter), string(sections.ReturnType), string(sections.Code))), nil
}

// Execute -
func Execute(rpc noderpc.INode, event Event, ctx Context) ([]TokenBalance, error) {
	parameter := event.Normalize(ctx.Parameters)
	storage := gjson.Parse(`[]`)
	code, err := event.GetCode()
	if err != nil {
		return nil, err
	}

	response, err := rpc.RunCode(code, storage, parameter, ctx.ChainID, ctx.Source, ctx.Initiator, ctx.Entrypoint, ctx.Amount, ctx.HardGasLimitPerOperation)
	if err != nil {
		return nil, err
	}
	logger.Debug(response)
	return event.Parse(response), nil
}

// NormalizeName -
func NormalizeName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "-", "")
	return strings.ReplaceAll(name, "_", "")
}