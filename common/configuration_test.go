package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var (
	tuser              = "someuser"
	tfingerprint       = "somefingerprint"
	tkeyfile           = "somelocation"
	ttenancy           = "sometenancy"
	tregion            = "someregion"
	testPrivateKeyConf = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDCFENGw33yGihy92pDjZQhl0C36rPJj+CvfSC8+q28hxA161QF
NUd13wuCTUcq0Qd2qsBe/2hFyc2DCJJg0h1L78+6Z4UMR7EOcpfdUE9Hf3m/hs+F
UR45uBJeDK1HSFHD8bHKD6kv8FPGfJTotc+2xjJwoYi+1hqp1fIekaxsyQIDAQAB
AoGBAJR8ZkCUvx5kzv+utdl7T5MnordT1TvoXXJGXK7ZZ+UuvMNUCdN2QPc4sBiA
QWvLw1cSKt5DsKZ8UETpYPy8pPYnnDEz2dDYiaew9+xEpubyeW2oH4Zx71wqBtOK
kqwrXa/pzdpiucRRjk6vE6YY7EBBs/g7uanVpGibOVAEsqH1AkEA7DkjVH28WDUg
f1nqvfn2Kj6CT7nIcE3jGJsZZ7zlZmBmHFDONMLUrXR/Zm3pR5m0tCmBqa5RK95u
412jt1dPIwJBANJT3v8pnkth48bQo/fKel6uEYyboRtA5/uHuHkZ6FQF7OUkGogc
mSJluOdc5t6hI1VsLn0QZEjQZMEOWr+wKSMCQQCC4kXJEsHAve77oP6HtG/IiEn7
kpyUXRNvFsDE0czpJJBvL/aRFUJxuRK91jhjC68sA7NsKMGg5OXb5I5Jj36xAkEA
gIT7aFOYBFwGgQAQkWNKLvySgKbAZRTeLBacpHMuQdl1DfdntvAyqpAZ0lY0RKmW
G6aFKaqQfOXKCyWoUiVknQJAXrlgySFci/2ueKlIE1QqIiLSZ8V8OlpFLRnb1pzI
7U1yQXnTAEFYM560yJlzUpOb1V4cScGd365tiSMvxLOvTA==
-----END RSA PRIVATE KEY-----`
)

func removeFileFn(filename string) func() {
	return func() {
		os.Remove(filename)
	}
}

func writeTempFile(data string) (filename string) {
	f, _ := ioutil.TempFile("", "gosdkTest")
	f.WriteString(data)
	filename = f.Name()
	return
}

func TestFileConfigurationProvider_parseConfigFileData(t *testing.T) {
	data := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=somelocation
tenancy=sometenancy
compartment = somecompartment
region=someregion
`
	c, e := parseConfigFile([]byte(data))
	assert.NoError(t, e)
	assert.Equal(t, c.UserOcid, tuser)
	assert.Equal(t, c.Fingerprint, tfingerprint)
	assert.Equal(t, c.KeyFilePath, tkeyfile)
	assert.Equal(t, c.TenancyOcid, ttenancy)
	assert.Equal(t, c.Region, tregion)
}

func TestFileConfigurationProvider_ParseEmptyFile(t *testing.T) {
	data := ``
	_, e := parseConfigFile([]byte(data))
	assert.Error(t, e)
}

func TestFileConfigurationProvider_FromFile(t *testing.T) {
	expected := []string{ttenancy, tuser, tfingerprint, tkeyfile}
	data := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=somelocation
tenancy=sometenancy
compartment = somecompartment
region=someregion
`
	filename := writeTempFile(data)
	defer removeFileFn(filename)

	c := fileConfigurationProvider{ConfigPath: filename}
	fns := []func() (string, error){c.TenancyOCID, c.UserOCID, c.KeyFingerPrint}

	for i, fn := range fns {
		val, e := fn()
		assert.NoError(t, e)
		assert.Equal(t, expected[i], val)
	}
}

func TestFileConfigurationProvider_NoFile(t *testing.T) {
	c := fileConfigurationProvider{ConfigPath: "/no/file"}
	fns := []func() (string, error){c.TenancyOCID, c.UserOCID, c.KeyFingerPrint}

	for _, fn := range fns {
		_, e := fn()
		assert.Error(t, e)
	}
}

func TestFileConfigurationProvider_KeyProvider(t *testing.T) {
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
tenancy=sometenancy
compartment = somecompartment
region=someregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	c := fileConfigurationProvider{ConfigPath: tmpConfFile}
	rskey, e := c.PrivateRSAKey()
	keyId, e1 := c.KeyID()
	assert.NoError(t, e)
	assert.NotEmpty(t, rskey)
	assert.NoError(t, e1)
	assert.NotEmpty(t, keyId)
}

func TestFileConfigurationProvider_FromFileFn(t *testing.T) {
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
tenancy=sometenancy
compartment = somecompartment
region=someregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	c, e0 := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, e0)
	rskey, e := c.PrivateRSAKey()
	keyId, e1 := c.KeyID()
	assert.NoError(t, e)
	assert.NotEmpty(t, rskey)
	assert.NoError(t, e1)
	assert.NotEmpty(t, keyId)
}

func TestFileConfigurationProvider_FromFileIncomplete(t *testing.T) {
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
compartment = somecompartment
region=someregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	c, e0 := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, e0)
	_, e1 := c.KeyID()
	assert.NoError(t, e1)
	_, e1 = c.TenancyOCID()
	assert.Error(t, e1)

}

func TestFileConfigurationProvider_FromFileIncomplete2(t *testing.T) {
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
compartment = somecompartment
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(keyFile)

	c, e0 := ConfigurationProviderFromFile(tmpConfFile, "")
	assert.NoError(t, e0)
	_, e1 := c.KeyID()
	assert.NoError(t, e1)
	_, e1 = c.TenancyOCID()
	assert.Error(t, e1)
	_, e1 = c.Region()
	assert.Error(t, e1)

}

func TestComposingConfigurationProvider_MultipleFiles(t *testing.T) {
	dataTpl0 := ``
	dataTpl := `[DEFAULT]
user=someuser
fingerprint=somefingerprint
key_file=%s
tenancy=sometenancy
compartment = somecompartment
region=someregion
`

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile0 := writeTempFile(dataTpl0)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(tmpConfFile0)
	defer removeFileFn(keyFile)

	c0, _ := ConfigurationProviderFromFile(tmpConfFile, "")
	c, _ := ConfigurationProviderFromFile(tmpConfFile, "")

	provider, ec := ComposingConfigurationProvider([]ConfigurationProvider{c0, c})
	assert.NoError(t, ec)
	fns := []func() (string, error){provider.TenancyOCID, provider.UserOCID, provider.KeyFingerPrint}

	for _, fn := range fns {
		val, e := fn()
		assert.NoError(t, e)
		assert.NotEmpty(t, val)
	}
	key, _ := provider.PrivateRSAKey()
	assert.NotNil(t, key)
}

func TestComposingConfigurationProvider_MultipleFilesNoConf(t *testing.T) {
	dataTpl0 := ``
	dataTpl := ` `

	keyFile := writeTempFile(testPrivateKeyConf)
	data := fmt.Sprintf(dataTpl, keyFile)
	tmpConfFile0 := writeTempFile(dataTpl0)
	tmpConfFile := writeTempFile(data)

	defer removeFileFn(tmpConfFile)
	defer removeFileFn(tmpConfFile0)
	defer removeFileFn(keyFile)

	c0, _ := ConfigurationProviderFromFile(tmpConfFile, "")
	c, _ := ConfigurationProviderFromFile(tmpConfFile, "")

	provider, ec := ComposingConfigurationProvider([]ConfigurationProvider{c0, c})
	assert.NoError(t, ec)
	fns := []func() (string, error){provider.TenancyOCID, provider.UserOCID, provider.KeyFingerPrint}

	for _, fn := range fns {
		_, e := fn()
		assert.Error(t, e)
	}
}