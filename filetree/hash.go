package filetree

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/kalafut/imohash"
	"github.com/minio/sha256-simd"
)

func MD5(content []byte) string {
	digest := md5.New()
	digest.Write(content)
	sum := digest.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func Sha256(content []byte) string {
	digest := sha256.New()
	digest.Write(content)
	sum := digest.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func ImoHash(file string) string {
	hash, err := imohash.SumFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%016x", hash)
}
