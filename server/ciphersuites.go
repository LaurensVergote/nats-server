// Copyright 2016-2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"crypto/tls"
)

// Where we maintain all of the available ciphers
var cipherMap map[string]uint16
var cipherMapByID map[uint16]string

func init() {
	const approxCount = 40 // be a little larger
	cipherMap = make(map[string]uint16, approxCount)
	cipherMapByID = make(map[uint16]string, approxCount)
	for _, cs := range tls.CipherSuites() {
		cipherMap[cs.Name] = cs.ID
		cipherMapByID[cs.ID] = cs.Name
	}
	for _, cs := range tls.InsecureCipherSuites() {
		cipherMap[cs.Name] = cs.ID
		cipherMapByID[cs.ID] = cs.Name
	}
}

func defaultCipherSuites() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	}
}

// Where we maintain available curve preferences
var curvePreferenceMap = map[string]tls.CurveID{
	"X25519":    tls.X25519,
	"CurveP256": tls.CurveP256,
	"CurveP384": tls.CurveP384,
	"CurveP521": tls.CurveP521,
}

// reorder to default to the highest level of security.  See:
// https://blog.bracebin.com/achieving-perfect-ssl-labs-score-with-go
func defaultCurvePreferences() []tls.CurveID {
	return []tls.CurveID{
		tls.X25519, // faster than P256, arguably more secure
		tls.CurveP256,
		tls.CurveP384,
		tls.CurveP521,
	}
}
