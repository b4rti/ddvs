package main

import (
	"github.com/docker/go-plugins-helpers/volume"
)

func (d DDVSDriver) Path(req *volume.PathRequest) (*volume.PathResponse, error) {
	debug(req)

	res := &volume.PathResponse{
		Mountpoint: "",
	}

	return res, nil
}
