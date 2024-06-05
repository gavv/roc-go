// Code generated by generate_bindings.py script from roc-streaming/bindgen
// roc-toolkit git tag: v0.3.0, commit: 57b932b8

package roc

// Media encoding.
//
// Defines format and parameters of samples encoded in frames or packets.
type MediaEncoding struct {
	// Sample frequency.
	//
	// Defines number of samples per channel per second (e.g. 44100).
	Rate uint32

	// Sample format.
	//
	// Defines sample precision and encoding.
	Format Format

	// Channel layout.
	//
	// Defines number of channels and meaning of each channel.
	Channels ChannelLayout

	// Multi-track channel count.
	//
	// If Channels is ChannelLayoutMultitrack, defines number of channels (which
	// represent independent "tracks"). For other channel layouts should be zero.
	//
	// Should be in range [1; 1024].
	Tracks uint32
}
