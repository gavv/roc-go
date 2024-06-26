package roc

func makeContextConfig() ContextConfig {
	return ContextConfig{
		MaxPacketSize: 2000,
		MaxFrameSize:  4000,
	}
}

func makeSenderConfig() SenderConfig {
	return SenderConfig{
		FrameEncoding:  makeMediaEncoding(),
		PacketEncoding: PacketEncodingAvpL16Stereo,
	}
}

func makeReceiverConfig() ReceiverConfig {
	return ReceiverConfig{
		FrameEncoding: makeMediaEncoding(),
	}
}

func makeInterfaceConfig() InterfaceConfig {
	return InterfaceConfig{
		OutgoingAddress: "127.0.0.1",
		ReuseAddress:    true,
	}
}

func makeMediaEncoding() MediaEncoding {
	return MediaEncoding{
		Rate:     44100,
		Format:   FormatPcmFloat32,
		Channels: ChannelLayoutStereo,
	}
}
