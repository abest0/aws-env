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
	cli.StringFlag{
		Name:   "aws-home",
		Usage:  "location of aws home",
		EnvVar: "AWS_HOME",
	},
}

// Params are the options that will used to pull data out
// out of the credentials file
type Params struct {
	awsHome  string
	fileName string
	profile  string
}

func entering(n string) string {
	log.Printf("Entering [%s]", n)
	return n
}

func exiting(n string) {
	log.Printf("Exiting [%s]", n)
}

func setup(ctx *cli.Context) error {
	if !ctx.Bool("verbose") {
		log.SetOutput(ioutil.Discard)
	}
	return nil
}

func buildCredentialsPath(awsHome, fileName string) string {

	if "" == awsHome {
		fmt.Fprintln(os.Stderr, "Environment variable [$AWS_HOME] is not set. Is the aws-cli installed? If not, install it or use --aws-home option")
		os.Exit(1)
	}

	return path.Join(awsHome, fileName)
}

func process(params Params) (map[string]string, error) {
	credentialPath := buildCredentialsPath(params.awsHome, params.fileName)

	log.Printf("Inspecting file: %s\n", credentialPath)
	log.Printf("Extracting profile [%s]\n", params.profile)

	cfg, err := ini.Load(credentialPath)

	if nil != err {
		fmt.Println("An error occurred loading the file")
		return nil, err
	}

	section, err := cfg.GetSection(params.profile)

	if nil != err {
		fmt.Println("An error occurred extracting the profile")
		return nil, err
	}

	return section.KeysHash(), nil
}

// CmdProcess extracts the AWS Keys from the profile of the of the specified
// credentials file.  These values will then be set as the environment vaiables
func CmdProcess(ctx *cli.Context) {
	defer exiting(entering("CmdProcess"))
	params := Params{ctx.String("aws-home"), ctx.String("file"), ctx.String("profile")}
	m, err := process(params)

	if nil != err {
		log.Fatalln(err)
	}

	for k, v := range m {
		k := strings.ToUpper(k)
		fmt.Printf("export %s=\"%v\"\n", k, v)
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
