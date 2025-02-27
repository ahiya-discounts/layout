package dep

import "github.com/google/wire"

var DepProviderSet = wire.NewSet(NewZapLogger)
