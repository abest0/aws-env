# aws-env.go
**aws-env.go** is a simple command line app, written in Go, that provides easy access to the **AWS ACCESS KEY ID** and **AWS SECRET ACCESS KEY** associated with the various profiles in the AWS credentials file.

####Disclaimer
Remember this is just a tool to facilitate local development.  ***Friends don't let friends share secret keys***

##Overview
Credentials for the AWS Command Line Interface are typically set using `aws configure`.  The credentials file for the **aws-cli** is structured as follows:
```
[default]
aws_access_key_id = <some aws access key id>
aws_secret_access_key = <some really long key>
[staging]
aws_access_key_id = <another aws access key id>
aws_secret_access_key = <another really long key>
```
This credentials profile contains credentials for 2 profiles.  **aws-env.go** provides easy access to the values for setting environment variables or for use in scripts

##Installation
1. [Install Go][install] obvs.
1. Ensure `$GOPATH/bin` is included in your path
1. Install the app using the following command:
``` sh
go get github.com/abest0/aws-env
```

##Usage
So, you want to set your environment variables?
``` sh
eval "$(aws-env)"
```

Need to set a variable for a script?
``` sh
ak="$(aws-env access-key)"
echo $ak
```


##Help
Like most cli apps, if you have any questions on usage, type `aws-env --help`

``` shell
NAME:
   aws-env - extract AWS Access Key Id and Secret Access Key

USAGE:
   aws-env [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   access-key   retrieve and output the AWS access key id for the provided profile
   secret-key   retrieve and output the AWS secret key for the provided profile
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose                    display more output
   -f, --file "credentials"     aws credentials file
   -p, --profile "default"      profile to extract from aws credentials file
   --aws-home                   location of aws home [$AWS_HOME]
   --help, -h                   show help
   --version, -v                print the version
```

[install]: http://golang.org/doc/install
