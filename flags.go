// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "flag"

var (
	fWsConfigFile = flag.String("config-server", "config.json", "Web server configuration file")
	fDbConfigFile = flag.String("config-database", "mysql.json", "Database configuration file")
	fEnableConsole = flag.Bool("enable-console", false, "Enable console logging")
	fEnableLogFiles = flag.Bool("enable-logfiles", false, "Enable logfile logging")
	fEnableSyslog = flag.Bool("enable-syslog", false, "Enable syslog logging")
)