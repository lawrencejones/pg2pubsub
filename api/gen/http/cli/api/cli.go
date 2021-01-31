// Code generated by goa v3.2.6, DO NOT EDIT.
//
// api HTTP client CLI support package
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	healthc "github.com/lawrencejones/pgsink/api/gen/http/health/client"
	importsc "github.com/lawrencejones/pgsink/api/gen/http/imports/client"
	tablesc "github.com/lawrencejones/pgsink/api/gen/http/tables/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `health check
tables list
imports list
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` health check` + "\n" +
		os.Args[0] + ` tables list --schema "public,payments"` + "\n" +
		os.Args[0] + ` imports list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		healthFlags = flag.NewFlagSet("health", flag.ContinueOnError)

		healthCheckFlags = flag.NewFlagSet("check", flag.ExitOnError)

		tablesFlags = flag.NewFlagSet("tables", flag.ContinueOnError)

		tablesListFlags      = flag.NewFlagSet("list", flag.ExitOnError)
		tablesListSchemaFlag = tablesListFlags.String("schema", "public", "")

		importsFlags = flag.NewFlagSet("imports", flag.ContinueOnError)

		importsListFlags = flag.NewFlagSet("list", flag.ExitOnError)
	)
	healthFlags.Usage = healthUsage
	healthCheckFlags.Usage = healthCheckUsage

	tablesFlags.Usage = tablesUsage
	tablesListFlags.Usage = tablesListUsage

	importsFlags.Usage = importsUsage
	importsListFlags.Usage = importsListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "health":
			svcf = healthFlags
		case "tables":
			svcf = tablesFlags
		case "imports":
			svcf = importsFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "health":
			switch epn {
			case "check":
				epf = healthCheckFlags

			}

		case "tables":
			switch epn {
			case "list":
				epf = tablesListFlags

			}

		case "imports":
			switch epn {
			case "list":
				epf = importsListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "health":
			c := healthc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "check":
				endpoint = c.Check()
				data = nil
			}
		case "tables":
			c := tablesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = tablesc.BuildListPayload(*tablesListSchemaFlag)
			}
		case "imports":
			c := importsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// healthUsage displays the usage of the health command and its subcommands.
func healthUsage() {
	fmt.Fprintf(os.Stderr, `Provide service health information
Usage:
    %s [globalflags] health COMMAND [flags]

COMMAND:
    check: Health check for probes

Additional help:
    %s health COMMAND --help
`, os.Args[0], os.Args[0])
}
func healthCheckUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] health check

Health check for probes

Example:
    `+os.Args[0]+` health check
`, os.Args[0])
}

// tablesUsage displays the usage of the tables command and its subcommands.
func tablesUsage() {
	fmt.Fprintf(os.Stderr, `Expose Postgres tables, and their import/sync status
Usage:
    %s [globalflags] tables COMMAND [flags]

COMMAND:
    list: List all tables

Additional help:
    %s tables COMMAND --help
`, os.Args[0], os.Args[0])
}
func tablesListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tables list -schema STRING

List all tables
    -schema STRING: 

Example:
    `+os.Args[0]+` tables list --schema "public,payments"
`, os.Args[0])
}

// importsUsage displays the usage of the imports command and its subcommands.
func importsUsage() {
	fmt.Fprintf(os.Stderr, `Manage table imports, scoped to the server subscription ID
Usage:
    %s [globalflags] imports COMMAND [flags]

COMMAND:
    list: List all imports

Additional help:
    %s imports COMMAND --help
`, os.Args[0], os.Args[0])
}
func importsListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] imports list

List all imports

Example:
    `+os.Args[0]+` imports list
`, os.Args[0])
}