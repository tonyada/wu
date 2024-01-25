package aes

import (
	"testing"
	. "wu"
)

func printByte(b []byte) {
	print("[]byte{")
	for i := 0; i < len(b); i++ {
		if i != len(b)-1 {
			print(Sprintf("%v,", b[i]))
		} else {
			// last byte without ,
			print(Sprintf("%v", b[i]))
		}
	}
	println("}")
}

func TestAes(t *testing.T) {
	// iv 初始向量，长度必须为16个字节(128bit) UTF8中文 = 3 bytes
	en, err := Encrypt([]byte("abc"))
	if err != nil {
		t.Errorf("Encrypt err %v", err)
	}
	de, err := Decrypt(en)
	if err != nil {
		t.Errorf("Decrypt err  %v", err)
	}
	// output good bytes for copy
	printByte(en)
	myEncodePass := []byte{47, 137, 70}
	Printfln("%v encrypt bytes: %v", len(en), en)
	Println(len(de), " decrypt string: ", string(de))
	mycode, _ := Decrypt(myEncodePass)
	Println("mycode:", string(mycode))
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Encrypt([]byte("abc"))
	}
}

func BenchmarkDecrypt(b *testing.B) {
	en, _ := Encrypt([]byte("abc"))
	for i := 0; i < b.N; i++ {
		_, _ = Decrypt(en)
	}
}
