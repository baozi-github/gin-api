package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"gin-api/util"
	"log"
	"strings"
)

// FormatterPublicKey 格式化公钥
func FormatterPublicKey(publicKey string) string {
	if util.Strpos(publicKey, "-----BEGIN PUBLIC KEY-----", 0) != -1 {
		return publicKey
	}

	str := util.ChunkSplit(publicKey, 64, "\n")
	publicKey = "-----BEGIN PUBLIC KEY-----" + "\n" + str + "-----END PUBLIC KEY-----"
	return publicKey
}

// FormatPrivateKey 格式化私钥
func FormatPrivateKey(privateKey string) string {
	if strings.Index(privateKey, "-----BEGIN RSA PRIVATE KEY-----") != -1 {
		return privateKey
	}
	str := util.ChunkSplit(privateKey, 64, "\n")
	privateKey = "-----BEGIN RSA PRIVATE KEY-----" + "\n" + str + "-----END RSA PRIVATE KEY-----"
	return privateKey
}

// Encrypt 公钥分段加密
func Encrypt(src []byte, publicKeyByte string) string {
	block, _ := pem.Decode([]byte(FormatterPublicKey(publicKeyByte)))
	if block == nil {
		log.Println("encrypt block error")
		return ""
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Println("parse error ", err.Error())
		return ""
	}

	keySize, srcSize := publicKey.Size(), len(src)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src[offSet:endIndex])
		if err != nil {
			log.Println("err2", err.Error())
			//return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt := buffer.Bytes()
	return base64.StdEncoding.EncodeToString(bytesEncrypt)
}

// RsaDecrypt 私钥分段解密
func RsaDecrypt(str string, privateKeyBytes string) string {
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("decrypt decode error ", err.Error())
		return ""
	}

	block, _ := pem.Decode([]byte(FormatterPublicKey(privateKeyBytes)))
	if block == nil {
		log.Println("decrypt pem decode error ")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("decrypt x509 error ", err.Error())
		return ""
	}

	keySize := privateKey.Size()
	srcSize := len(decodeString)
	log.Println("密钥长度：", keySize, "\t密文长度：\t", srcSize)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeString[offSet:endIndex])
		if err != nil {
			log.Println("decrypt DecryptPKCS1v15 error ", err.Error())
			return ""
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt := buffer.Bytes()
	return string(bytesDecrypt)

}
