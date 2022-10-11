package server

import (
	"app/pkg/utils"
	"io/ioutil"
	nhttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func transmitRequestDecoder(r *nhttp.Request, v interface{}) error {
	codec, ok := http.CodecForRequest(r, "Content-Type")
	if !ok {
		return errors.BadRequest("CODEC", r.Header.Get("Content-Type"))
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if err = codec.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	ip := utils.ConvIP(r.Header)
	r.Header.Set("x-app-global-requestIP", ip)
	return nil
}
