package roc

import "time"

// Sender configuration.
// You can zero-initialize this struct to get a default config.
// See also Sender.
type SenderConfig struct {
	// The rate of the samples in the frames passed to sender.
	//
	// Number of samples per channel per second. If FrameSampleRate and
	// PacketSampleRate are different, resampler should be enabled. Should be set.
	FrameSampleRate uint32

	// The channel set in the frames passed to sender.
	//
	// Should be set.
	FrameChannels ChannelSet

	// The sample encoding in the frames passed to sender.
	//
	// Should be set.
	FrameEncoding FrameEncoding

	// The rate of the samples in the packets generated by sender.
	//
	// Number of samples per channel per second. If zero, default value is used.
	PacketSampleRate uint32

	// The channel set in the packets generated by sender.
	//
	// If zero, default value is used.
	PacketChannels ChannelSet

	// The sample encoding in the packets generated by sender.
	//
	// If zero, default value is used.
	PacketEncoding PacketEncoding

	// The length of the packets produced by sender, in nanoseconds.
	//
	// Number of nanoseconds encoded per packet. The samples written to the sender
	// are buffered until the full packet is accumulated or the sender is flushed
	// or closed. Larger number reduces packet overhead but also increases latency.
	// If zero, default value is used.
	PacketLength time.Duration

	// Enable packet interleaving.
	//
	// If true, the sender shuffles packets before sending them. This may
	// increase robustness but also increases latency.
	PacketInterleaving bool

	// Clock source to use.
	//
	// Defines whether write operation will be blocking or non-blocking. If zero,
	// default value is used.
	ClockSource ClockSource

	// Resampler backend to use.
	ResamplerBackend ResamplerBackend

	// Resampler profile to use.
	//
	// If non-zero, the sender employs resampler if the frame sample rate differs
	// from the packet sample rate.
	ResamplerProfile ResamplerProfile

	// FEC encoding to use.
	//
	// If non-zero, the sender employs a FEC encoding to generate redundant packets
	// which may be used on receiver to restore lost packets. This requires both
	// sender and receiver to use two separate source and repair endpoints.
	FecEncoding FecEncoding

	// Number of source packets per FEC block.
	//
	// Used if some FEC encoding is selected. Larger number increases robustness
	// but also increases latency. If zero, default value is used.
	FecBlockSourcePackets uint32

	// Number of repair packets per FEC block.
	//
	// Used if some FEC encoding is selected. Larger number increases robustness
	// but also increases traffic. If zero, default value is used.
	FecBlockRepairPackets uint32
}