package service

import "github.com/chayaninpia/go-backend-microservice/pdst001/util"

func Pdst001() error {
	e, err := util.InitXorm()
	if err != nil {
		return err
	}

	// res, err := req.QueryStock(e)
	// if err != nil {
	// 	return err
	// }

	defer func() error {
		if err := e.Close(); err != nil {
			return err
		}
		return err
	}()

	return nil
}
