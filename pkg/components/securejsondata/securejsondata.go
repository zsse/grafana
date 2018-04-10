package securejsondata

import (
	"github.com/rkrikbaev/grafana/pkg/log"
	"github.com/rkrikbaev/grafana/pkg/setting"
	"github.com/rkrikbaev/grafana/pkg/util"
)

type SecureJsonData map[string][]byte

func (s SecureJsonData) Decrypt() map[string]string {
	decrypted := make(map[string]string)
	for key, data := range s {
		decryptedData, err := util.Decrypt(data, setting.SecretKey)
		if err != nil {
			log.Fatal(4, err.Error())
		}

		decrypted[key] = string(decryptedData)
	}
	return decrypted
}

func GetEncryptedJsonData(sjd map[string]string) SecureJsonData {
	encrypted := make(SecureJsonData)
	for key, data := range sjd {
		encryptedData, err := util.Encrypt([]byte(data), setting.SecretKey)
		if err != nil {
			log.Fatal(4, err.Error())
		}

		encrypted[key] = encryptedData
	}
	return encrypted
}
