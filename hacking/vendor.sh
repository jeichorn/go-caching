#!/usr/bin/env bash

set -e

# clone function for downloading a vendor 
clone() {
    vcs=$1
    package=$2
    commit=$3

    url_package=https://$package
    directory_target=src/$package

    echo -n "$package @ $commit: "

    if [ -d $directory_target ]; then
        echo -n "removing old, "
        rm -fr $directory_target
    fi

    echo -n "cloning, "
    case $vcs in
        git)
            git clone --quiet --no-checkout $url_package $directory_target
            ( cd $directory_target && git reset --quiet --hard $comment )
            ;;
        hg)
            hg clone --quiet --updaterev $commit $url_package $directory_target
            ;;
    esac

    echo -n "removing vcs, "
    ( cd $directory_target && rm -rf .{git,hg} )

    echo "done"
}

# Ensure vendor directory exists and enter for downloading
directory_origin=$(cd `dirname $BASH_SOURCE`/.. && pwd -P)
directory_vendor=$directory_origin/_vendor
mkdir -p $directory_vendor
cd $directory_vendor

# Download packages
echo "Downloading vendors..."
clone git github.com/rainycape/memcache     756dcc4797 # v2: Jul 25, 2014
echo "Done"

cd $directory_origin

# Tip for setup GOPATH
echo ""
echo "You need run following code to setup GOPATH for vendors: "
echo ""
echo "export GOPATH=$directory_vendor:\$GOPATH"
echo ""
echo "Done"
