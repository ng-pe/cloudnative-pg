/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2020 2ndQuadrant Italia SRL. Exclusively licensed to 2ndQuadrant Limited.
*/

// Package app contains all the code that is directly run by the
// operator executable
package app

import (
	"fmt"
	"os"

	"github.com/2ndquadrant/cloud-native-postgresql/pkg/fileutils"
)

// BootstrapInto is called by the controller manager to copy the operator executable
// inside a certain location. This is useful to insert the controller in a
// volume to be used by the actual PostgreSQL controller
func BootstrapInto(executablePath string, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: manager bootstrap <target>")
		os.Exit(1)
	}

	dest := args[0]

	err := fileutils.CopyFile(executablePath, dest)
	if err != nil {
		panic(err)
	}

	err = os.Chmod(dest, 0755) // #nosec
	if err != nil {
		panic(err)
	}
}