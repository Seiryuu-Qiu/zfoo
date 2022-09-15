package protocol

type UdpHelloRequest struct {
	message string
}

func (protocol UdpHelloRequest) protocolId() int16 {
	return 1200
}

func (protocol UdpHelloRequest) write(buffer *ByteBuffer, packet any) {
	if buffer.WritePacketFlag(packet) {
		return
	}
	var message = packet.(*UdpHelloRequest)
	buffer.WriteString(message.message)
}

func (protocol UdpHelloRequest) read(buffer *ByteBuffer) any {
	var packet = new(UdpHelloRequest)
	if !buffer.ReadBool() {
		return packet
	}
	var result0 = buffer.ReadString()
    packet.message = result0
	return packet
}


type UdpHelloResponse struct {
	message string
}

func (protocol UdpHelloResponse) protocolId() int16 {
	return 1201
}

func (protocol UdpHelloResponse) write(buffer *ByteBuffer, packet any) {
	if buffer.WritePacketFlag(packet) {
		return
	}
	var message = packet.(*UdpHelloResponse)
	buffer.WriteString(message.message)
}

func (protocol UdpHelloResponse) read(buffer *ByteBuffer) any {
	var packet = new(UdpHelloResponse)
	if !buffer.ReadBool() {
		return packet
	}
	var result0 = buffer.ReadString()
    packet.message = result0
	return packet
}