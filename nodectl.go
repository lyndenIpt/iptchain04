package main

import (
	"os"
	"sort"

	_ "IPT/cli"
	"IPT/cli/asset"
	"IPT/cli/bookkeeper"
	. "IPT/cli/common"
	"IPT/cli/consensus"
	"IPT/cli/data"
	"IPT/cli/debug"
	"IPT/cli/info"
	"IPT/cli/multisig"
	"IPT/cli/privpayload"
	"IPT/cli/recover"
	"IPT/cli/smartcontract"
	"IPT/cli/test"
	"IPT/cli/wallet"

	"github.com/urfave/cli"
)

var Version string

func main() {
	app := cli.NewApp()
	app.Name = "nodectl"
	app.Version = Version
	app.HelpName = "nodectl"
	app.Usage = "command line tool for IPT blockchain"
	app.UsageText = "nodectl [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//global options
	app.Flags = []cli.Flag{
		NewIpFlag(),
		NewPortFlag(),
	}
	//commands
	app.Commands = []cli.Command{
		*consensus.NewCommand(),
		*debug.NewCommand(),
		*info.NewCommand(),
		*test.NewCommand(),
		*wallet.NewCommand(),
		*asset.NewCommand(),
		*privpayload.NewCommand(),
		*data.NewCommand(),
		*bookkeeper.NewCommand(),
		*recover.NewCommand(),
		*multisig.NewCommand(),
		*smartcontract.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
