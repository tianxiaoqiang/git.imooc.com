package envelopes

import (
	"github.com/tietang/dbx"
	"github.com/prometheus/common/log"
	"github.com/shopspring/decimal"
	"git.imooc.com/wendell1000/account/services"
	"time"
)
//红包商品表的dao
type RedEnvelopeGoodsDao struct {

	runner *dbx.TxRunner
}

//插入商品

func (dao *RedEnvelopeGoodsDao) Insert (po *RedEnvelopeGoods) (int64,error) {

	rs,err := dao.runner.Insert(po)
	if err != nil{
		return 0,err
	}
	return rs.LastInsertId()

}

//更新红包余额和数量

func (dao *RedEnvelopeGoodsDao) GetOne(envelopeNo string) *RedEnvelopeGoods {
	po := &RedEnvelopeGoods{EnvelopeNo:envelopeNo}
	ok,err := dao.runner.GetOne(po)
	if !ok || err != nil{
		log.Error(err)
		return nil
	}
	return po
}

//更新红包余额和数量
//不再使用事务行锁来更新红包余额和数量
//使用乐观锁来保证更新红包余额和数量的安全，避免负库存问题
//通过在where子句中判断红包剩余金额和数量来解决2个问题：
//1. 负库存问题，避免红包剩余金额和数量不够时进行扣减更新
//2. 减少实际的数据更新，也就是过滤掉部分无效的更新，提高总体性能
func (dao *RedEnvelopeGoodsDao) UpdateBalance(envelopeNo string,amount decimal.Decimal) (int64,error) {

	sql := "update red_envelope_goods " +
		" set remain_amount=remain_amount-CAST(? AS DECIMAL(30,6)), " +
		" remain_quantity=remain_quantity-1 " +
		" where envelope_no=? " +
	//最重要的，乐观锁的关键
		" and remain_quantity>0" +
		" and remain_amount >= CAST(? AS DECIMAL(30,6)) "

	rs, err := dao.runner.Exec(sql, amount.String(), envelopeNo, amount.String())
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

//更新订单状态
func (dao *RedEnvelopeGoodsDao) UpdateStatus (envelopeNo string,orderStatus services.OrderStatus) (int64,error) {

	sql := "update red_envelope_goods set status = ? where envelope_no = ?"
	rs,err := dao.runner.Exec(sql,orderStatus,envelopeNo)

	if err !=nil{
		return 0,err
	}
	return rs.RowsAffected()
}


//更新订单支付状态
func (dao * RedEnvelopeGoodsDao) UpdatePayStatus (envelopeNo string,payStatus services.PayStatus) (int64,error){

	sql := "update red_envelop_goods " +
		" set pay_status = ?" +
		" where envelope_no = ?"

	rs,err := dao.runner.Exec(sql,payStatus,envelopeNo)

	if err != nil{

		return 0,err
	}
	return rs.RowsAffected()
}


//过期，把过期的所有红包都查询出来，分页，limit,offset size
func (dao *RedEnvelopeGoodsDao) FindExpired (limit,offset int) []RedEnvelopeGoods {
	var goods []RedEnvelopeGoods
	now := time.Now()
	sql := " select * from red_envelope_goods " +
		" where order_type=1 and remain_quantity>0  and expired_at>? and status in(1,2,3,6) " +
		" limit ?,?"
	err := dao.runner.Find(&goods,sql,now,limit,offset)
	if err != nil{
		log.Error(err)
	}
	return goods
}


