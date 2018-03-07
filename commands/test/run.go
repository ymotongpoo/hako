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

package test

import (
	"flag"
	"fmt"
)

const CommandName = "test"

// Test implements commands.Runner, and represents the type for 'test' command.
type Test struct {
	Args    []string
	FlagSet *flag.FlagSet
}

// New returns new Test instance.
func New(args []string) Test {
	return Test{
		Args: args,
	}
}

// Run is the operation command to launch 'test' command.
func (t Test) Run() {
	fmt.Printf("[test] %v\n", t.Args) // DEBUG
}
