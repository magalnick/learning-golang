package main

import (
	"fmt"
	"github.com/magalnick/my-nice-project/helpers"
)

/**
 * To tell this program that packages will be used:
 * go mod init github-or-lab.com/account-name/path-to-project
 * For example:
 * github.com/magalnick/my-nice-project
 * ...or...
 * gitlab.com/truefootageinc/backend-development/truefootage-monorepo/src
 *
 * This will create a file called go.mod that can then have the package definitions in it.
 * The github part is optional, but standard practice since a project with packages will likely
 * have code stored in a git repo somewhere.
 *
 * Then create a sub-folder, and in there a new .go file.
 */

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "type name"
	myVar.TypeNumber = 2
	myVar.PrivateType("private or not?")

	fmt.Println(myVar)
}
