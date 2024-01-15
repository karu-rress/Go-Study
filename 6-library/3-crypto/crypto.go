package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"io"
)

func Hash() {
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s))

	sha := sha512.New()
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("world!"))

	h2 := sha.Sum(nil)
	fmt.Printf("%x\n%x\n", h1, h2)
}

func SymmetricECB() {
	// 모두 16 Bytes로 맞춤
	key := "Hello, Key 12345"
	s := "Hello, world! 12"

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)

	fmt.Println(string(plaintext))
}

func encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크기의 배수여야 합니다.")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(b, iv)

	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

func SymmetricCBC() {
	key := "Hello Key 123456"

	s := `환영합니다, Rolling Ress의 카루입니다.
샘플 테스트로 AES CBC 암호화를 진행해보겠습니다.
이걸 Go로 하는 정신나간 사람이 있다..
그게 바로 저예요!`

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	ciphertext := encrypt(block, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext)
	fmt.Println(string(plaintext))

}

func Asymmetric() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	publicKey := &privateKey.PublicKey

	s := `환영합니다, Rolling Ress의 카루입니다.
샘플 테스트로 AES CBC 암호화를 진행해보겠습니다.
이걸 Go로 하는 정신나간 사람이 있다..
그게 바로 저예요!`

	ciphertext, _ := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)
	fmt.Printf("%x\n", ciphertext)

	plaintext, _ := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)
	fmt.Println(string(plaintext))
}

func DigitalSign() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	publicKey := &privateKey.PublicKey

	message := "안녕하세요 카루님"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, _ := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		h1, digest,
	)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(
		publicKey,
		h2, digest, signature,
	)

	if err != nil {
		fmt.Println("검증 성공")
	}
}

func main() {
	SymmetricCBC()
}
