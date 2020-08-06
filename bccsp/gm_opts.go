package bccsp

const (
	SM2 = "SM2"
	SM3 = "SM3"
	SM4 = "SM4"
)

var _ KeyGenOpts = &SM2KeyGenOpts{}
var _ KeyGenOpts = &SM2ImportKeyOpts{}
var _ KeyGenOpts = &SM4KeyGenOpts{}
var _ KeyGenOpts = &SM4ImportKeyOpts{}
var _ HashOpts = &SM3Opts{}

//sm2 opts
type SM2KeyGenOpts struct {
	Temporary bool
}

func (opts *SM2KeyGenOpts) Algorithm() string {
	return SM2
}

func (opts *SM2KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

type SM2ImportKeyOpts struct {
	Temporary bool
}

func (opts *SM2ImportKeyOpts) Algorithm() string {
	return SM2
}

func (opts *SM2ImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

//sm3 opts
type SM3Opts struct{}

func (opts *SM3Opts) Algorithm() string {
	return SM3
}

//sm4 opts
type SM4KeyGenOpts struct {
	Temporary bool
}

func (opts *SM4KeyGenOpts) Algorithm() string {
	return SM4
}

func (opts *SM4KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

type SM4ImportKeyOpts struct {
	Temporary bool
}

func (opts *SM4ImportKeyOpts) Algorithm() string {
	return SM4
}

func (opts *SM4ImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}
