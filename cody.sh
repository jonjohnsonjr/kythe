#!/bin/bash -e
#
# Copyright 2018 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
 
##############################
# Run this in the root of the Kythe repo
##############################
ROOT="$PWD"
 
lein() {
  local bin="$(which lein)"
  if [[ -z "$bin" ]]; then
    bin="$ROOT/third_party/leiningen/lein"
  fi
  "$bin" "$@"
}
 
brun() {
  local target="$1"
  shift
  bazel run -c opt "$target" -- "$@"
}
 
if [[ ! -r kythe/web/ui/resources/public/js/main.js ]]; then
  pushd kythe/web/ui
  lein cljsbuild once prod
  popd
fi
 
INCLUDE=
if [[ "$1" == "--include_imports" ]]; then
  shift
  INCLUDE=imports
elif [[ "$1" == "--include_deps" ]]; then
  shift
  INCLUDE=deps
fi
 
go get -d "$@"
packages=($(go list "$@"))
all_packages=($(go list -f $'{{.ImportPath}}\n{{join .Deps "\\n"}}' "$@" | sort -u))
go install -i "${all_packages[@]}"
 
if [[ "$INCLUDE" == "imports" ]]; then
  packages=($(go list -f $'{{.ImportPath}}\n{{ join .Imports "\\n" }}' "$@" | sort -u))
elif [[ "$INCLUDE" == "deps" ]]; then
  packages=(${all_packages[@]})
fi
 
brun //kythe/go/extractors/cmd/gotool --corpus="GOPATH" --goroot="$(go env GOROOT)" --output=/tmp/out.kzip --continue "${packages[@]}"
brun //kythe/go/indexer/cmd/go_indexer /tmp/out.kzip >/tmp/out.entries
rm -rf /tmp/out
brun //kythe/go/serving/tools/write_tables --experimental_beam_pipeline --entries /tmp/out.entries --out /tmp/out
 
brun //kythe/go/serving/tools/http_server --listen :8080 --public_resources "$ROOT"/kythe/web/ui/resources/public --serving_table /tmp/out
