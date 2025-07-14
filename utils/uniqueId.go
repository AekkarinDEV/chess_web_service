package utils

import (
	"crypto/rand"
	"encoding/hex"
	"sync/atomic"
	"time"
)

var counter uint32 = 0

func GenerateRecordID() string {
    b := make([]byte, 12)

    // 4 bytes: current timestamp
    timestamp := uint32(time.Now().Unix())
    b[0] = byte(timestamp >> 24)
    b[1] = byte(timestamp >> 16)
    b[2] = byte(timestamp >> 8)
    b[3] = byte(timestamp)

    // 5 bytes: random machine + process id
    rand.Read(b[4:9])

    // 3 bytes: incrementing counter (wraps at 2^24)
    c := atomic.AddUint32(&counter, 1) % 0xFFFFFF
    b[9] = byte(c >> 16)
    b[10] = byte(c >> 8)
    b[11] = byte(c)

    return hex.EncodeToString(b)
}
