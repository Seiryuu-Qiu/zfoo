{}
const {} = function() {
    {}
};

{}.PROTOCOL_ID = {};

{}.prototype.protocolId = function() {
    return {}.PROTOCOL_ID;
};

{}.write = function(buffer, packet) {
    if (packet === null) {
        buffer.writeInt(0);
        return;
    }
    {}
};

{}.read = function(buffer) {
    const length = buffer.readInt();
    if (length === 0) {
        return null;
    }
    const beforeReadIndex = buffer.getReadOffset();
    const packet = new {}();
    {}
    if (length > 0) {
        buffer.setReadOffset(beforeReadIndex + length);
    }
    return packet;
};

export default {};
