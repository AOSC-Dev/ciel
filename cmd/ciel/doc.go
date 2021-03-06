package main

import "fmt"

func docHelp() {
	fmt.Print(`Usage:
	ciel version
	ciel init
	ciel load-os [TAR_FILE]    // unpack OS tarball or fetch the latest BuildKit from internet directly
	ciel load-tree [GIT_URL]   // clone package tree from your link or AOSC OS ABBS at GitHub

	ciel update-os -- [params] // similar to 'apt-get update && apt-get dist-upgrade', params are appended to 'apt-get dist-upgrade'
	ciel update-tree           // similar to 'git pull'


	ciel [list]
	ciel add INSTANCE
	ciel del INSTANCE
	ciel shell -i INSTANCE         // start an interactive shell
	ciel shell -i INSTANCE "SHELL COMMAND LINE"
	ciel config (-i INSTANCE | -g) // configure system and toolchain for building (interactively)
	ciel build -i INSTANCE PACKAGE
	ciel rollback -i INSTANCE

	ciel down [-i INSTANCE]    // shutdown & unmount all or one instance
	ciel mount [-i INSTANCE]   // mount all or one instance

Rarely used:
	ciel stop -i INSTANCE      // shutdown an instance
	ciel run -i INSTANCE ABSPATH_TO_EXE ARG1 ARG2 ...
	                  // lower-level version of 'shell', without login environment,
	                  // without sourcing ~/.bash_profile
	ciel farewell  // DELETE ALL CIEL THINGS, except OUTPUT, TREE etc.
	               // equals to 'ciel down && rm -r .ciel'

	ciel doctor    // diagnose problems

Altering OS & Releasing OS:
	ciel load-os
	ciel update-os  // see above
	ciel generate
	             // (plugin) install packages and set up environment by RECIPE
	ciel factory-reset -i INSTANCE
	             // delete all out-of-dpkg files
	ciel commit -i INSTANCE
	             // commit changes onto the shared underlying OS
	ciel release VARIANT THREADS
	             // (plugin) make a .tar.xz release for the underlying OS

Global flags:
	-C CIEL_DIR    // use CIEL_DIR as workdir instead of current directory
	-i INSTANCE    // specify the INSTANCE to manipulate
	-batch         // batch mode, no input is required
	-n             // do not start 'init' (systemd)
`)
}
