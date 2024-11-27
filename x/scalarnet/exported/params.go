package exported

import (
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
)

const (
	// ModuleName exposes scalarnet module name
	ModuleName = "Scalarnet"
)

var (
	// NativeAsset is the native asset on ScalarNet
	NativeAsset = "sclr"

	Scalarnet = exported.Chain{
		Name:                  "ScalarNet",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                ModuleName,
	}
)
