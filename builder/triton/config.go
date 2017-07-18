package triton

import (
	"github.com/idanya/packer/common"
	"github.com/idanya/packer/helper/communicator"
	"github.com/idanya/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	AccessConfig        `mapstructure:",squash"`
	SourceMachineConfig `mapstructure:",squash"`
	TargetImageConfig   `mapstructure:",squash"`

	Comm communicator.Config `mapstructure:",squash"`

	ctx interpolate.Context
}
