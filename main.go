//    Copyright 2018 Yoshi Yamaguchi
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"fmt"
	"os"
	"path"
)

const SelfCommandName = "hako" // TODO: make this replacable on build time.

func cmdAndOpts() (string, []string) {
	cmdPath := os.Args[0]
	if path.Base(cmdPath) == SelfCommandName {
		if len(os.Args) > 1 {
			return os.Args[1], os.Args[2:]
		}
		return "", []string{}
	}
	if len(os.Args) > 1 {
		return os.Args[0], os.Args[1:]
	}
	return os.Args[0], []string{}
}

func main() {
	command, opts := cmdAndOpts()
	fmt.Println(command, opts) // DEBUG
}
