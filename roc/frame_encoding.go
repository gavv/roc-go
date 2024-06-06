// DO NOT EDIT! Code generated by generate_enums script from roc-toolkit
// roc-toolkit git tag: v0.2.6, commit: 3cbb3d13

package roc

// Frame encoding.
//
//go:generate stringer -type FrameEncoding -trimprefix FrameEncoding -output frame_encoding_string.go
type FrameEncoding int

const (
	// PCM floats.
	//
	// Uncompressed samples coded as floats in range [-1; 1]. Channels are
	// interleaved, e.g. two channels are encoded as "L R L R...".
	FrameEncodingPcmFloat FrameEncoding = 1
)