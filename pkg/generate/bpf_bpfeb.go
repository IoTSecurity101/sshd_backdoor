// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package generate

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type BpfCustomPayload struct {
	RawBuf     [450]uint8
	_          [2]byte
	PayloadLen uint32
}

type BpfSyscallReadLogging struct {
	BufferAddr  uint64
	CallingSize int64
}

// LoadBpf returns the embedded CollectionSpec for Bpf.
func LoadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Bpf: %w", err)
	}

	return spec, err
}

// LoadBpfObjects loads Bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*BpfObjects
//	*BpfPrograms
//	*BpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// BpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfSpecs struct {
	BpfProgramSpecs
	BpfMapSpecs
}

// BpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfProgramSpecs struct {
	HandleOpenatEnter *ebpf.ProgramSpec `ebpf:"handle_openat_enter"`
	HandleOpenatExit  *ebpf.ProgramSpec `ebpf:"handle_openat_exit"`
	HandleReadEnter   *ebpf.ProgramSpec `ebpf:"handle_read_enter"`
	HandleReadExit    *ebpf.ProgramSpec `ebpf:"handle_read_exit"`
}

// BpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type BpfMapSpecs struct {
	MapBuffAddrs     *ebpf.MapSpec `ebpf:"map_buff_addrs"`
	MapFds           *ebpf.MapSpec `ebpf:"map_fds"`
	MapPayloadBuffer *ebpf.MapSpec `ebpf:"map_payload_buffer"`
}

// BpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfObjects struct {
	BpfPrograms
	BpfMaps
}

func (o *BpfObjects) Close() error {
	return _BpfClose(
		&o.BpfPrograms,
		&o.BpfMaps,
	)
}

// BpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfMaps struct {
	MapBuffAddrs     *ebpf.Map `ebpf:"map_buff_addrs"`
	MapFds           *ebpf.Map `ebpf:"map_fds"`
	MapPayloadBuffer *ebpf.Map `ebpf:"map_payload_buffer"`
}

func (m *BpfMaps) Close() error {
	return _BpfClose(
		m.MapBuffAddrs,
		m.MapFds,
		m.MapPayloadBuffer,
	)
}

// BpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type BpfPrograms struct {
	HandleOpenatEnter *ebpf.Program `ebpf:"handle_openat_enter"`
	HandleOpenatExit  *ebpf.Program `ebpf:"handle_openat_exit"`
	HandleReadEnter   *ebpf.Program `ebpf:"handle_read_enter"`
	HandleReadExit    *ebpf.Program `ebpf:"handle_read_exit"`
}

func (p *BpfPrograms) Close() error {
	return _BpfClose(
		p.HandleOpenatEnter,
		p.HandleOpenatExit,
		p.HandleReadEnter,
		p.HandleReadExit,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfeb.o
var _BpfBytes []byte
