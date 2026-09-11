package iroh_ffi

// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/libs/x86_64-unknown-linux-musl -liroh -lm
// #cgo linux,arm64 LDFLAGS: -L${SRCDIR}/libs/aarch64-unknown-linux-musl -liroh -lm
// #cgo windows,amd64 LDFLAGS: -L${SRCDIR}/libs/x86_64-pc-windows-gnu -liroh -lws2_32 -luserenv -lbcrypt -lntdll -liphlpapi -ladvapi32 -lcrypt32 -lsecur32
import "C"
