package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/codegangsta/cli"
	"gopkg.in/ini.v1"
)

// Contant val
const awsHome = "AWS_HOME"

var flags = []cli.Flag{
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Display more output",
	},
	cli.StringFlag{
		Name:  "f, file",
		Value: "credentials",
		Usage: "aws credentials file",
	},
	cli.StringFlag{
		Name:  "p, profile",
		Value: "default",
		Usage: "profile to extract from aws credentials file",
	},
}

// ExtractOptions are the options that will used to pull data out
// out of the credentials file
type ExtractOptions struct {
	credentialsFile string
	profile         string
}

func setup(ctx *cli.Context) error {
	if !ctx.Bool("verbose") {
		log.SetOutput(ioutil.Discard)
	}
	return nil
}

func (exo *ExtractOptions) process() (map[string]string, error) {
	awsHome := os.Getenv(awsHome)
	p := path.Join(awsHome, exo.credentialsFile)

	log.Printf("Inspecting file: %s\n", p)
	log.Printf("Extracting profile [%s]\n", exo.profile)

	cfg, err := ini.Load(p)

	if nil != err {
		fmt.Println("An error occurred loading the file")
		return nil, err
	}

	section, err := cfg.GetSection(exo.profile)

	if nil != err {
		fmt.Println("An error occurred extracting the profile")
		return nil, err
	}

	return section.KeysHash(), nil
}

// CmdProcess extracts the AWS Keys from the profile of the of the specified
// credentials file.  These values will then be set as the environment vaiables
func CmdProcess(c *cli.Context) {

	fp := ExtractOptions{c.String("file"), c.String("profile")}
	m, err := fp.process()

	if nil != err {
		log.Panic(err)
	}

	for k, v := range m {
		k := strings.ToUpper(k)
		fmt.Printf("export %s=\"%v\"\n", k, v)
		err := os.Setenv(k, v)
		if nil != err {
			log.Panic(err)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "aws-env"
	app.Usage = "extract AWS Secret Key Id and Access Keys"
	app.Flags = flags
	app.Before = setup
	app.Action = CmdProcess

	app.Run(os.Args)
}
