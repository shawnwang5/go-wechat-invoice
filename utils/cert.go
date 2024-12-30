package utils

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

// ReadPemCert 读取证书
//
// 参数：
//   - req: 请求参数
//
// 返回值：
//   - res: 响应数据
//   - err: error
func ReadPemCert(certPath string) (cert *x509.Certificate, err error) {
	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return cert, err
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return cert, err
	}

	cert, err = x509.ParseCertificate(block.Bytes)
	return cert, err
}
