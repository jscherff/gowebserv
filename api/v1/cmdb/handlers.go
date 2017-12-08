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

package cmdb

import (
	`net/http`
	`github.com/jscherff/gox/log`
)

// AuthSetTokenV1 authenticates client using basic authentication and
// issues a JWT for API authentication if successful.
func AuthSetTokenV1(w http.ResponseWriter, r *http.Request) {

	if user, pass, ok := r.BasicAuth(); !ok {
		errLog.Print(`missing credentials`)
		http.Error(w, `missing credentials`, http.StatusUnauthorized)
	} else if token, err := createAuthToken(user, pass, r.URL.Host); err != nil {
		errLog.Print(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	} else if cookie, err := createAuthCookie(token); err != nil {
		errLog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		sysLog.Printf(`issuing auth token to %q at %q`, user, r.RemoteAddr)
		http.SetCookie(w, cookie)
	}
}