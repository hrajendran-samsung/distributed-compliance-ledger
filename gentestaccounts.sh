#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
    # Mac OS X sed needs the file extension when -i flag is used. Keeping it empty as we don't need backupfile
    SED_EXT="''"
fi

rm -rf localnet/client/*

PASSWD=test1234
NUMUSERS="${1:-10}"

for i in $(seq 1 "$NUMUSERS"); do
    userN="$(printf "%05d\n" "$i")"
    username="tu${userN}"
    echo $PASSWD"" | dclcli keys add "$username"

    tu_address="$(dclcli keys show "$username" -a)"
    tu_pubkey="$(dclcli keys show "$username" -p)"

    dcld add-genesis-account --address=$tu_address --pubkey=$tu_pubkey --roles="Vendor"
    echo "$username generated"
done

dcld validate-genesis

cp -r ~/.dclcli/* localnet/client
cp -f ~/.dcld/config/genesis.json localnet/node0/config/
cp -f ~/.dcld/config/genesis.json localnet/node1/config/
cp -f ~/.dcld/config/genesis.json localnet/node2/config/
cp -f ~/.dcld/config/genesis.json localnet/node3/config/
