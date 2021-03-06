package krb5

import (
	"strings"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var asRequestPacket []byte = []byte{
	0x00, 0x03, 0xff, 0xa6, 0xab, 0x0c, 0x00, 0x03,
	0xff, 0xa7, 0xab, 0x0c, 0x08, 0x00, 0x45, 0x00,
	0x01, 0x3e, 0x01, 0xf7, 0x00, 0x00, 0x80, 0x11,
	0x14, 0xb0, 0x0a, 0x01, 0x0c, 0x02, 0x0a, 0x05,
	0x03, 0x01, 0x04, 0x40, 0x00, 0x58, 0x01, 0x2a,
	0x4b, 0x2d, 0x6a, 0x82, 0x01, 0x1e, 0x30, 0x82,
	0x01, 0x1a, 0xa1, 0x03, 0x02, 0x01, 0x05, 0xa2,
	0x03, 0x02, 0x01, 0x0a, 0xa3, 0x5f, 0x30, 0x5d,
	0x30, 0x48, 0xa1, 0x03, 0x02, 0x01, 0x02, 0xa2,
	0x41, 0x04, 0x3f, 0x30, 0x3d, 0xa0, 0x03, 0x02,
	0x01, 0x17, 0xa2, 0x36, 0x04, 0x34, 0x71, 0x31,
	0x9a, 0x93, 0xd6, 0x05, 0x31, 0xfc, 0xb4, 0x43,
	0xf7, 0xe9, 0x60, 0x39, 0xf5, 0x40, 0xad, 0xdb,
	0xe6, 0x7c, 0xcf, 0x9d, 0xd3, 0xc3, 0xda, 0x9e,
	0x23, 0x36, 0x12, 0x81, 0x6c, 0x57, 0x20, 0x44,
	0x7a, 0xe2, 0x02, 0xcf, 0xe7, 0xa8, 0x4a, 0x71,
	0x9e, 0x1e, 0xf7, 0x0b, 0x93, 0xbc, 0xef, 0x49,
	0x78, 0x6f, 0x30, 0x11, 0xa1, 0x04, 0x02, 0x02,
	0x00, 0x80, 0xa2, 0x09, 0x04, 0x07, 0x30, 0x05,
	0xa0, 0x03, 0x01, 0x01, 0xff, 0xa4, 0x81, 0xac,
	0x30, 0x81, 0xa9, 0xa0, 0x07, 0x03, 0x05, 0x00,
	0x40, 0x81, 0x00, 0x10, 0xa1, 0x0f, 0x30, 0x0d,
	0xa0, 0x03, 0x02, 0x01, 0x01, 0xa1, 0x06, 0x30,
	0x04, 0x1b, 0x02, 0x75, 0x35, 0xa2, 0x08, 0x1b,
	0x06, 0x44, 0x45, 0x4e, 0x59, 0x44, 0x43, 0xa3,
	0x1b, 0x30, 0x19, 0xa0, 0x03, 0x02, 0x01, 0x02,
	0xa1, 0x12, 0x30, 0x10, 0x1b, 0x06, 0x6b, 0x72,
	0x62, 0x74, 0x67, 0x74, 0x1b, 0x06, 0x44, 0x45,
	0x4e, 0x59, 0x44, 0x43, 0xa5, 0x11, 0x18, 0x0f,
	0x32, 0x30, 0x33, 0x37, 0x30, 0x39, 0x31, 0x33,
	0x30, 0x32, 0x34, 0x38, 0x30, 0x35, 0x5a, 0xa6,
	0x11, 0x18, 0x0f, 0x32, 0x30, 0x33, 0x37, 0x30,
	0x39, 0x31, 0x33, 0x30, 0x32, 0x34, 0x38, 0x30,
	0x35, 0x5a, 0xa7, 0x06, 0x02, 0x04, 0x32, 0x0f,
	0xe8, 0xac, 0xa8, 0x19, 0x30, 0x17, 0x02, 0x01,
	0x17, 0x02, 0x02, 0xff, 0x7b, 0x02, 0x01, 0x80,
	0x02, 0x01, 0x03, 0x02, 0x01, 0x01, 0x02, 0x01,
	0x18, 0x02, 0x02, 0xff, 0x79, 0xa9, 0x1d, 0x30,
	0x1b, 0x30, 0x19, 0xa0, 0x03, 0x02, 0x01, 0x14,
	0xa1, 0x12, 0x04, 0x10, 0x58, 0x50, 0x31, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20,
}

var krbString = "$krb5$23$u5$DENYDC$nodata$71319a93d60531fcb443f7e96039f540addbe67ccf9dd3c3da9e233612816c5720447ae202cfe7a84a719e1ef70b93bcef49786f"

func TestHandlePackage(t *testing.T) {
	h := NewKrbHandler()
	h.HandlePacket(createIPv4UDPPacket(asRequestPacket))

	if len(h.KrbRequests) != 1 {
		t.Fatalf("wanted: one packet got: %v", len(h.KrbRequests))
	}
}

func TestKRBString(t *testing.T) {
	h := NewKrbHandler()
	h.HandlePacket(createIPv4UDPPacket(asRequestPacket))
	result, err := h.KrbRequests[0].String()

	if err != nil {
		t.Fatal("Error then parsing packet data")
	}

	if strings.Compare(result, krbString) != 0 {
		t.Fatalf("wanted: %v got: %v", krbString, result)
	}
}

func createIPv4UDPPacket(packet []byte) gopacket.Packet {
	return gopacket.NewPacket(packet, layers.LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
}
