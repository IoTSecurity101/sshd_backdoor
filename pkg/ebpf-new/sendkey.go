package ebpfnew

import (
	"ebpf_common/pkg/generate"
	"fmt"
	"golang.org/x/sys/unix"
)

func formatUint8SliceToUint8Array(slice []uint8) [450]uint8 {
	var array [450]uint8
	for i, v := range slice {
		array[i] = v
	}
	return array
}

// SendKey func sends ssh keys to map_payload_buffer.
func (c *CiliumEBPFRuntime) SendKey(key string) error {
	keyBytes, err := unix.ByteSliceFromString(key)
	if err != nil {
		return err
	}
	keyLen := len(keyBytes)
	if keyLen > 450 {
		return fmt.Errorf("key too long")
	}

	return c.Objects.MapPayloadBuffer.Put(uint8(0), generate.BpfCustomPayload{
		RawBuf:     formatUint8SliceToUint8Array(keyBytes),
		PayloadLen: uint32(keyLen),
	})
}