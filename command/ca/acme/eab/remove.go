package eab

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/smallstep/cli/flags"
	"github.com/smallstep/cli/utils/cautils"
	"github.com/urfave/cli"
	"go.step.sm/cli-utils/errs"
)

func removeCommand() cli.Command {
	return cli.Command{
		Name:   "remove",
		Action: cli.ActionFunc(removeAction),
		Usage:  "remove an ACME EAB Key from the CA",
		UsageText: `**step beta ca acme eab remove** <provisioner> <key_id>
[**--admin-cert**=<file>] [**--admin-key**=<file>]
[**--admin-provisioner**=<string>] [**--admin-subject**=<string>]
[**--password-file**=<file>] [**--ca-url**=<uri>] [**--root**=<file>]
[**--context**=<name>]`,
		Flags: []cli.Flag{
			flags.AdminCert,
			flags.AdminKey,
			flags.AdminProvisioner,
			flags.AdminSubject,
			flags.PasswordFile,
			flags.CaURL,
			flags.Root,
			flags.Context,
		},
		Description: `**step beta ca acme eab remove** removes an ACME EAB Key from the CA.

## POSITIONAL ARGUMENTS

<provisioner>
: Name of the provisioner to remove an ACME EAB key for

<key_id>
: The ACME EAB Key ID to remove

## EXAMPLES

Remove ACME EAB Key with Key ID "zFGdKC1sHmNf3Wsx3OujY808chxwEdmr" from my_acme_provisioner:
'''
$ step beta ca acme eab remove my_acme_provisioner zFGdKC1sHmNf3Wsx3OujY808chxwEdmr
'''
`,
	}
}

func removeAction(ctx *cli.Context) error {
	if err := errs.NumberOfArguments(ctx, 2); err != nil {
		return err
	}

	args := ctx.Args()
	provisioner := args.Get(0)
	keyID := args.Get(1)

	client, err := cautils.NewAdminClient(ctx)
	if err != nil {
		return errors.Wrap(err, "error creating admin client")
	}

	err = client.RemoveExternalAccountKey(provisioner, keyID)
	if err != nil {
		return errors.Wrap(err, "error removing ACME EAB key")
	}

	fmt.Println("Key was deleted successfully!")

	return nil
}
