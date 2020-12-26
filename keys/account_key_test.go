package keys

import (
	"github.com/GuoxiW/fractal/utils/log"
	"os"
	"testing"
)

func TestAccountKeySign(t *testing.T) {
	log.SetDefaultLogger(log.InitLog15Logger(log.LvlDebug, os.Stdout))
	key, err := LoadAccountKey("./key.json", "12345")
	if err != nil {
		return
	}
	key.PrivKey.Sign([]byte{})
}
