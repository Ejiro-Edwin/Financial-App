package v1

import (
	"encoding/json"
	"github.com/ejiro-edwin/finance-app/internal/api/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ServerVersion struct {
	Version string `json:"version"`
}

var versionJSON []byte


func init(){
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,

	})
	if err != nil {
		panic(err)
	}
}

func VersionHandler(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(400)
	if _,err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version.")
	}
}