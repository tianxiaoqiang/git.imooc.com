package services

const (
	DefaultBlessing = "恭喜发财,洪富猪到"
)

//订单类型:发布订单，退款订单

type OrderType int

const (
	OrderTypeSending OrderType = 1
	OrderTypeRefund  OrderType = 2
)

//支付状态：未支付，支付中，已支付，支付失败

type PayStatus int

const (
	PayNothing PayStatus = 1
	Paying     PayStatus = 2
	Payed      PayStatus = 3
	PayFailure PayStatus = 4
)

//活动状态：创建，激活，过期，失效
type ActivityStatus int

const (
	ActivityCreate    ActivityStatus = 1
	ActivityActivated ActivityStatus = 2
	ActivityExpired   ActivityStatus = 3
	ActivityDisabled  ActivityStatus = 4
)

//订单状态 1:订单创建 2:订单发布 3:订单过期，4:订单失效
type OrderStatus int

const (
	OrderCreate   OrderStatus = 1
	OrderSending  OrderStatus = 2
	OrderExpired  OrderStatus = 3
	OrderDisabled OrderStatus = 4
)

type envelopeType int

//红包类型：普通红包，碰运气红包
const (
	GeneralEnvelopeType envelopeType = 1
	LuckyEnvelopeType   envelopeType = 2
)
