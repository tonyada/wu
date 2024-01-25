package aes

// package main
import (
	"crypto/aes"
	"crypto/cipher"
	"wu"
)

// The key argument should be the AES key

type AESCrypt struct {
	key []byte // 16 = AES-128 24 = AES-192 32 = AES-256
	iv  []byte // iv 初始向量，长度必须为16个字节(128bit) UTF8中文 = 3 bytes
}

var AesCrypt = &AESCrypt{key: []byte("٩(•̮̮̃•̃)۶->钥匙kEy6"),
	iv: []byte("这Iv干毛的?6")[:aes.BlockSize]} //! WARNING!!! change key and iv for security

// set key
func SetKey(key string)               { AesCrypt.SetKey(key) }
func (a *AESCrypt) SetKey(key string) { AesCrypt.key = []byte(key) }

// set iv
func SetIv(key string)               { AesCrypt.SetIv(key) }
func (a *AESCrypt) SetIv(key string) { AesCrypt.iv = []byte(key)[:aes.BlockSize] }

// AES 加密 - msg 明文
func Encrypt(msg []byte) ([]byte, error) {
	encrypted := make([]byte, len(msg))
	err := AesCrypt.EncryptAESCFB(encrypted, msg)
	if wu.Err(err) {
		return nil, err
	}
	return encrypted, nil
}

// AES 加密 string - msg 明文
func EncryptStr(msg []byte) string {
	en, _ := Encrypt([]byte(msg))
	return string(en)
}

// AES 解密 msg 密文
func Decrypt(msg []byte) ([]byte, error) {
	decrypted := make([]byte, len(msg))
	err := AesCrypt.DecryptAESCFB(decrypted, msg)
	if wu.Err(err) {
		return nil, err
	}
	return decrypted, nil
}

// AES 解密 string - msg 明文
func DecryptStr(msg []byte) string {
	en, _ := Decrypt([]byte(msg))
	return string(en)
}

func (a *AESCrypt) EncryptAESCFB(dst, src []byte) error {
	// 得到块，用于加密
	aesBlockEncrypter, err := aes.NewCipher([]byte(AesCrypt.key))
	if wu.Err(err) {
		return err
	}
	// 加密，使用CFB模式(密文反馈模式)，其他模式参见crypto/cipher
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, AesCrypt.iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func (a *AESCrypt) DecryptAESCFB(dst, src []byte) error {
	// 得到块，用于解密
	aesBlockDecrypter, err := aes.NewCipher([]byte(AesCrypt.key))
	if wu.Err(err) {
		return nil
	}
	// 解密
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, AesCrypt.iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}

// func main() {
// 	const key16 = "1234567890123456"
// 	const key24 = "123456789012345678901234"
// 	const key32 = "12345678901234567890123456789012"
// 	const key = "٩(•̮̮̃•̃)۶->钥匙key6"
// 	var msg = "你好啊啊"

// 	enstr, _ := Encrypt(msg, key,"abcdef1234567890")
// 	Println(enstr)
// 	destr, _ := Decrypt(enstr, key,"abcdef1234567890")
// 	Println(destr)
// 	// // Encrypt
// 	// encrypted := make([]byte, len(msg))
// 	// err = EncryptAESCFB(encrypted, []byte(msg), []byte(key), iv)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Printf("Encrypting %v %s -> %v\n", []byte(msg), msg, encrypted)

// 	// // Decrypt
// 	// decrypted := make([]byte, len(msg))
// 	// err = DecryptAESCFB(decrypted, encrypted, []byte(key), iv)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Printf("Decrypting %v -> %v %s\n", encrypted, decrypted, decrypted)
// }

// // AESEncrypt encrypts text and given key with AES.
// func AESEncrypt(key, text []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b := base64.StdEncoding.EncodeToString(text)
// 	ciphertext := make([]byte, aes.BlockSize+len(b))
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return nil, err
// 	}
// 	cfb := cipher.NewCFBEncrypter(block, iv)
// 	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
// 	return ciphertext, nil
// }

// // AESDecrypt decrypts text and given key with AES.
// func AESDecrypt(key, text []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(text) < aes.BlockSize {
// 		return nil, errors.New("ciphertext too short")
// 	}
// 	iv := text[:aes.BlockSize]
// 	text = text[aes.BlockSize:]
// 	cfb := cipher.NewCFBDecrypter(block, iv)
// 	cfb.XORKeyStream(text, text)
// 	data, err := base64.StdEncoding.DecodeString(string(text))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }
