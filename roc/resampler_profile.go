// Code generated by generate_bindings.py script from roc-streaming/bindgen
// roc-toolkit git tag: v0.3.0, commit: 57b932b8

package roc

// Resampler profile.
//
// Affects CPU usage and quality. Each resampler backend treats profile in its
// own way.
//
//go:generate stringer -type ResamplerProfile -trimprefix ResamplerProfile -output resampler_profile_string.go
type ResamplerProfile int

const (
	// Default profile.
	//
	// Current default is ResamplerProfileMedium.
	ResamplerProfileDefault ResamplerProfile = 0

	// High quality, higher CPU usage.
	ResamplerProfileHigh ResamplerProfile = 1

	// Medium quality, medium CPU usage.
	ResamplerProfileMedium ResamplerProfile = 2

	// Low quality, lower CPU usage.
	ResamplerProfileLow ResamplerProfile = 3
)
