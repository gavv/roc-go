// Code generated by generate_bindings.py script from roc-streaming/bindgen
// roc-toolkit git tag: v0.4.0, commit: 62401be9

package roc

// Latency tuner profile.
//
// Defines whether latency tuning is enabled and which algorithm is used.
//
//go:generate stringer -type LatencyTunerProfile -trimprefix LatencyTunerProfile -output latency_tuner_profile_string.go
type LatencyTunerProfile int

const (
	// Default profile.
	//
	// On receiver, when LatencyTunerBackendNiq is used, selects
	// LatencyTunerProfileResponsive if target latency is low, and
	// LatencyTunerProfileGradual if target latency is high.
	//
	// On sender, selects LatencyTunerProfileIntact.
	LatencyTunerProfileDefault LatencyTunerProfile = 0

	// No latency tuning.
	//
	// In this mode, clock speed is not adjusted. Default on sender.
	//
	// You can set this mode on receiver, and set some other mode on sender, to do
	// latency tuning on sender side instead of recever side. It's useful when
	// receiver is CPU-constrained and sender is not, because latency tuner relies
	// on resampling, which is CPU-demanding.
	//
	// You can also set this mode on both sender and receiver if you don't need
	// latency tuning at all. However, if sender and receiver have independent
	// clocks (which is typically the case), clock drift will lead to periodic
	// playback disruptions caused by underruns and overruns.
	LatencyTunerProfileIntact LatencyTunerProfile = 1

	// Responsive latency tuning.
	//
	// Clock speed is adjusted quickly and accurately.
	//
	// Requires high precision clock adjustment, hence recommended for use with
	// ResamplerBackendBuiltin.
	//
	// Pros:
	//  - allows very low latency and synchronization error
	//
	// Cons:
	//  - does not work well with some resampler backends
	//  - does not work well with LatencyTunerBackendNiq if network jitter is high
	LatencyTunerProfileResponsive LatencyTunerProfile = 2

	// Gradual latency tuning.
	//
	// Clock speed is adjusted slowly and smoothly.
	//
	// Pros:
	//  - works well even with high network jitter
	//  - works well with any resampler backend
	//
	// Cons:
	//  - does not allow very low latency and synchronization error
	LatencyTunerProfileGradual LatencyTunerProfile = 3
)
