package filetree

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func MD5(input []byte) string {
	digest := md5.New()
	digest.Write(input)
	sum := digest.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func Sha256(content []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(content))
}
