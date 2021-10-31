package toy

import (
    "crypto/rand"
    "crypto/sha1"
    "encoding/base64"
)

func GenRandomString() string {
    buf := make([]byte, 16)
    _, err := rand.Read(buf)
    if err != nil {
        return ""
    }

    return base64.URLEncoding.EncodeToString(buf)
}

func HashBySha1(input string) string {
    buf := sha1.Sum([]byte(input))

    return base64.URLEncoding.EncodeToString(buf[:])
}
