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
	"os"
	"path"

	"github.com/ymotongpoo/hako/commands"
	"github.com/ymotongpoo/hako/commands/test"
)

const SelfCommandName = "hako" // TODO: make this replacable on build time.

type Command struct {
	Name string
	Args []string
}

func cmdAndOpts(args []string) Command {
	cmdPath := args[0]
	if path.Base(cmdPath) == SelfCommandName {
		if len(os.Args) > 1 {
			return Command{os.Args[1], os.Args[2:]}
		}
		return Command{"", []string{}}
	}
	if len(os.Args) > 1 {
		return Command{os.Args[0], os.Args[1:]}
	}
	return Command{os.Args[0], []string{}}
}

func router(c Command) commands.Runner {
	switch c.Name {
	case "test":
		return test.New(c.Args)
	default:
		return nil
	}
}

func main() {
	command := cmdAndOpts(os.Args)
	runner := router(command)
	runner.Run() // DEBUG
}
