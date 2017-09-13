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

import (
	"strings"
	"github.com/RackSec/srslog"
)

// Syslog extends RackSec srs log with embedded configuration information
// about the syslog daemon. It is part of the systemwide configuration under
// Config.Syslogs.
type Syslog struct {
	*srslog.Writer
	Tag string
	Port string
	Protocol string
	Hostname string
	Priority srslog.Priority
}

// NewSyslog instantiates a new syslog client.
func NewSyslog(tag, port, proto, host string, pri srslog.Priority) (this *Syslog, err error) {

	this = &Syslog{Tag: tag, Port: port, Protocol: proto, Hostname: host, Priority: pri}
	raddr := strings.Join([]string{host, port}, ":")

	this.Writer, err = srslog.Dial(proto, raddr, pri, tag)

	if err != nil {
		return nil, ErrorDecorator(err)
	}

	return this, nil
}

// Init initializes syslogs in the systemwide configuration under Config.Syslogs.
func (this *Syslog) Init() (err error) {

	raddr := strings.Join([]string{this.Hostname, this.Port}, ":")

	this.Writer, err = srslog.Dial(this.Protocol, raddr, this.Priority, this.Tag)

	if err != nil {
		err = ErrorDecorator(err)
	}

	return err
}