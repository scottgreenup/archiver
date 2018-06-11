#!/usr/bin/env bash

function create_file() {
    name="$1"
    user="$2"
    group="$3"
    perm="$4"

    filename="${name}"

    touch "${filename}"
    chown $user:$group ${filename}
    chmod $perm ${filename}
}

function create_permissions() {
    create_file "user_user_r--------" $USER $USER 400
    create_file "user_user_---r-----" $USER $USER 040
    create_file "user_user_------r--" $USER $USER 004
}

mkdir -p permissions
cd permissions
create_permissions
