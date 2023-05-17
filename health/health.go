// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package health

import (
	"errors"
	"fmt"
	"sync"
)

type component string

var (
	VulnDb component = "vuln-db"
)

// Every package can report its health status by calling NewHealth.
// Launcher will listen to the channel and report the status to the fleet server.
var Reporter = &reporter{
	ch:     make(chan error, 1),
	errors: map[component]error{},
	mut:    sync.Mutex{},
}

type reporter struct {
	ch     chan error
	errors map[component]error
	mut    sync.Mutex
}

func (r *reporter) NewHealth(com component, err error) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.errors[com] = err
	r.ch <- r.getHealth()
}

func (r *reporter) getHealth() error {
	list := make([]error, 0, len(r.errors))
	for c, err := range r.errors {
		if err != nil {
			list = append(list, fmt.Errorf("component %s is unhealthy: %w", c, err))
		}
	}

	return errors.Join(list...)
}

func (r *reporter) Channel() <-chan error {
	return r.ch
}

func (r *reporter) Stop() {
	close(r.ch)
}
