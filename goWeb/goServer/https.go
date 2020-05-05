package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main(){
	creatPrivateKey()//生成个人使用的ssl证书以及服务器私钥
	server:=http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           nil,
		//TLSConfig:         nil,
		//ReadTimeout:       0,
		//ReadHeaderTimeout: 0,
		//WriteTimeout:      0,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      nil,
		//ConnState:         nil,
		//ErrorLog:          nil,
		//BaseContext:       nil,
		//ConnContext:       nil,
	}
	//cert.pem是ssl证书，key.pem是服务器私钥
	server.ListenAndServeTLS("cert.pem","key.pem")

}
//生成个人使用的ssl证书以及服务器私钥
func creatPrivateKey()  {
	max:=new(big.Int).Lsh(big.NewInt(1),128)
	serialNumber,_:=rand.Int(rand.Reader,max)//证书序列号
	subject:=pkix.Name{
		Organization:       []string{"Manning  Publications Co"},
		OrganizationalUnit: []string{"books"},
		CommonName:         "Go Web Programming",
	}
	template:=x509.Certificate{//单词certificate（证明）usage（用法）
		SerialNumber:                serialNumber,              //证书序列号
		Subject:                     subject,					//专有名称，证书的标题
		NotBefore:                   time.Now(),//time.Time{},	//下面这两行是证书的有效期，这里是1年
		NotAfter:                    time.Now().Add(365*24*time.Hour),//time.Time{},
		KeyUsage:                    x509.KeyUsageKeyEncipherment|x509.KeyUsageDigitalSignature,//证书的用途
		ExtKeyUsage:                 []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:                 []net.IP{net.ParseIP("127.0.0.1")},

	}
	pk,_:=rsa.GenerateKey(rand.Reader,2048)//生成RSA私钥，里面包含了一个公开访问的公钥
	derBytes,_:=x509.CreateCertificate(rand.Reader,&template,&template,&pk.PublicKey,pk)//创建ssl证书
	//使用encoding/pem标准库将证书编码到cert.pem
	certOut,_:=os.Create("cert.pem")
	pem.Encode(certOut,&pem.Block{Type:"CERTIFICATE",Bytes:derBytes})
	certOut.Close()
	//以pem编码的方式把之前生成的密钥编码并保存到key.pem里
	keyOut,_:=os.Create("key.pem")
	pem.Encode(keyOut,&pem.Block{Type:"RSA PRIVATE KEY",Bytes:x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}