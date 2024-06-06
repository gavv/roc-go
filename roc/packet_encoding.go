// DO NOT EDIT! Code generated by generate_enums script from roc-toolkit
// roc-toolkit git tag: v0.2.6, commit: 3cbb3d13

package roc

// Packet encoding.
//
//go:generate stringer -type PacketEncoding -trimprefix PacketEncoding -output packet_encoding_string.go
type PacketEncoding int

const (
	// PCM signed 16-bit.
	//
	// "L16" encoding from RTP A/V Profile (RFC 3551). Uncompressed samples coded
	// as interleaved 16-bit signed big-endian integers in two's complement
	// notation.
	PacketEncodingAvpL16 PacketEncoding = 2
)