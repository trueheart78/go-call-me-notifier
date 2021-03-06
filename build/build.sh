#!/bin/bash

set -e -x

echo "Creating release dir..."
mkdir -p release

# variables as defined by "go tool nm"
OSVAR=github.com/trueheart78/go-call-me-notifier/cmd.BuildOS
ARCHVAR=github.com/trueheart78/go-call-me-notifier/cmd.BuildARCH
ARMVAR=github.com/trueheart78/go-call-me-notifier/cmd.BuildARM

# handle alternate binary name for pre-releases
BINNAME=${NAME:-go-call-me-notifier}

createRelease() {
	os=$1
	arch=$2
	arm=$3

	if [ "$os" = darwin ]
	then
		osname='mac'
	else
		osname=$os
	fi
	if [ "$arch" = amd64 ]
	then
		osarch=64bit
	else
		osarch=32bit
	fi

	ldflags="-X $OSVAR=$os -X $ARCHVAR=$arch"
	if [ "$arm" ]
	then
		osarch=arm-v$arm
		ldflags="$ldflags -X $ARMVAR=$arm"
	elif [ "$arch" = arm64 ]
	then
		osarch=arm-v8
		ldflags="$ldflags -X $ARMVAR=8"
	fi

	binname=$BINNAME
	if [ "$osname" = windows ]
	then
		binname="$binname.exe"
	fi

	relname="../release/$BINNAME-$osname-$osarch"
	echo "Creating $os/$arch binary..."

	if [ "$arm" ]
	then
		GOOS=$os GOARCH=$arch GOARM=$arm go build -ldflags "$ldflags" -o "out/$binname" cmd/go-call-me-notifier/go-call-me-notifier.go
	else
		GOOS=$os GOARCH=$arch go build -ldflags "$ldflags" -o "out/$binname" cmd/go-call-me-notifier/go-call-me-notifier.go
	fi

	cd out

	if [ "$osname" = windows ]
	then
		zip "$relname.zip" "$binname"
	else
		tar cvzf "$relname.tgz" "$binname"
	fi
	cd ..
	rm -rf out/
}

createRelease darwin amd64
