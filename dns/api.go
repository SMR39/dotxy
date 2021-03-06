package dns

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// ResolveDNSQuery function sends the dns query to the dns resolver
func ResolveDNSQuery(verbose bool, resolverAddr string, query []byte) ([]byte, error) {
	conn, err := tls.Dial("tcp", resolverAddr, &tls.Config{})
	if err != nil {
		return nil, err
	}

	req := make([]byte, len(query)+2)
	binary.BigEndian.PutUint16(req[0:2], uint16(len(query)))
	copy(req[2:], query)

	_, err = conn.Write(req)
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Println("sent query on to server", resolverAddr)
	}

	resp, err := readDNSResponse(conn)
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Println("received response from the DNS server", resolverAddr, "with sizeof", len(resp), "bytes")
	}

	return resp, nil
}

func readDNSResponse(r io.Reader) ([]byte, error) {
	length, err := readResponseLength(r)
	if err != nil {
		return nil, err
	}

	return readDNSMessage(r, length)
}

func readResponseLength(r io.Reader) (int, error) {
	bytes := make([]byte, 2)
	n, err := r.Read(bytes)
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, fmt.Errorf("reading length did not receive enough bytes")
	}
	length := int(binary.BigEndian.Uint16(bytes))
	return length, nil
}

func readDNSMessage(r io.Reader, length int) ([]byte, error) {
	resp := make([]byte, length)
	offset := 0
	for offset < length {
		n, err := r.Read(resp[offset:])
		if err != nil {
			return nil, err
		}
		offset += n
	}
	return resp, nil
}
