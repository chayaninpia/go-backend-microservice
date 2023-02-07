package handler

import (
	"github.com/chayaninpia/go-backend-microservice/models"
	"github.com/chayaninpia/go-backend-microservice/pdst001/util"
	"xorm.io/xorm"
)

type Pdst001Handler struct {
	pdst001I *models.Pdst001I
	pdst001O *models.Pdst001O
}

func NewHandler() *Pdst001Handler {
	return &Pdst001Handler{}
}

func (h *Pdst001Handler) Pdst001(req *models.Pdst001I) (*models.Pdst001O, error) {

	h.pdst001I = req
	e, err := util.InitXorm()
	if err != nil {
		return nil, err
	}

	if err := h.QueryStock(e); err != nil {
		return nil, err
	}

	defer func() error {
		if err := e.Close(); err != nil {
			return err
		}
		return err
	}()

	return h.pdst001O, nil
}

func (h *Pdst001Handler) QueryStock(e *xorm.Engine) error {

	res := make(models.Pdst001O, 0)

	qs := e.Select(`tp.id, tp.barcode_id, tp.product_name , tps.quantity`).Table(`t_product_stock`).Alias(`tps`).
		Join(`INNER`, `t_product AS tp`, `tp.id = tps.product_id`)

	if h.pdst001I.BarcodeId != `` {
		qs.Where(`tp.barcode_id = ?`, h.pdst001I.BarcodeId)
	}

	if h.pdst001I.ProductName != `` {
		qs.Where(`tp.product_name = ?`, h.pdst001I.ProductName)
	}

	if h.pdst001I.ProductId != `` {
		qs.Where(`tps.product_id = ?`, h.pdst001I.ProductId)
	}

	qRes, err := qs.QueryInterface()
	if err != nil {
		return err
	}

	for _, v := range qRes {
		res = append(res, models.Pdst001IChild{
			ProductId:   v[`id`].(string),
			BarcodeId:   v[`barcode_id`].(string),
			ProductName: v[`product_name`].(string),
			Quantity:    v[`quantity`].(int32),
		})
	}

	h.pdst001O = &res
	return nil
}
