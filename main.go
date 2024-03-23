package main

import (
	"fmt"

	"dcupdate/app"
	"dcupdate/app/environment"
)

var GITCOMMIT string = "Development"
var VERSION string = "Development"

func main() {

	if VERSION == "Development" {
		fmt.Println(`
________  _________  ____ ___            .___       __
\______ \ \_   ___ \|    |   \______   __| _/____ _/  |_  ____
 |    |  \/    \  \/|    |   /\____ \ / __ |\__  \\   __\/ __ \
 |        \     \___|    |  / |  |_> > /_/ | / __ \|  | \  ___/
/_______  /\______  /______/  |   __/\____ |(____  /__|  \___  >
        \/        \/          |__|        \/     \/          \/
 `)
	}

	environment.VERSION = VERSION
	environment.SetUpEnv()

	app.Start()

}
