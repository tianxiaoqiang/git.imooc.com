package envelopes

import (
	"github.com/segmentio/ksuid"
	"git.imooc.com/wendell1000/account/services"
)

type goodsDomain struct {
	//引入持久化对象
	RedEnvelopeGoods
}


//生成一个红包编号

func (d *goodsDomain) CreateEnvelopeNo(){
	d.EnvelopeNo = ksuid.New().Next().String()
}

//创建一个红包商品对象(传入一个红包商品的DTO)
func (d *goodsDomain) Create(goods services.RedEnvelopeGoodsDTO)  {
	//存入到RedEnvelopeGoods中
	d.RedEnvelopeGoods.FromDTO(&goods)


}