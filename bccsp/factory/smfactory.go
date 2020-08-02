/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package factory

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/sm"
	"github.com/pkg/errors"
)

const (
	SMFactoryName = "SM"
)

type SMFactory struct{}

// Name returns the name of this factory
func (f *SMFactory) Name() string {
	return SMFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *SMFactory) Get(config *FactoryOpts) (bccsp.BCCSP, error) {
	// Validate arguments
	if config == nil || config.SM == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	smOpts := config.SM

	var ks bccsp.KeyStore
	switch {
	case smOpts.FileKeystore != nil:
		fks, err := sm.NewFileBasedKeyStore(nil, smOpts.FileKeystore.KeyStorePath, false)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to initialize software key store")
		}
		ks = fks
	default:
		// Default to ephemeral key store
		ks = sm.NewDummyKeyStore()
	}

	return sm.NewWithParams(smOpts.Security, smOpts.Hash, ks)
}

// SMOpts contains options for the SMFactory
type SMOpts struct {
	// Default algorithms when not specified (Deprecated?)
	Security     int               `json:"security" yaml:"Security"`
	Hash         string            `json:"hash" yaml:"Hash"`
	FileKeystore *FileKeystoreOpts `json:"filekeystore,omitempty" yaml:"FileKeyStore,omitempty"`
}

// Pluggable Keystores,
type FileSMKeystoreOpts struct {
	KeyStorePath string `yaml:"KeyStore"`
}
