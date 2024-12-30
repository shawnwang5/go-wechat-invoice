package officialaccount

import "os"

// 获取预开票 URL 的响应数据
type GetPreInvoicingUrlRes struct {
	Errcode    string `json:"errcode" mapstructure:"errcode"`         // 错误码
	Errmsg     string `json:"errmsg" mapstructure:"errmsg"`           // 错误信息
	InvoiceUrl string `json:"invoice_url" mapstructure:"invoice_url"` // 该开票平台专用的授权链接。开票平台须将 invoice_url 内的 s_pappid 给到服务的商户，商户在请求授权链接时会向微信传入该参数，标识所使用的开票平台是哪家
}

// 设置商户联系方式
type SetMerchantContactReqContact struct {
	TimeOut int    `json:"time_out" mapstructure:"time_out"` // 开票超时时间
	Phone   string `json:"phone" mapstructure:"phone"`       // 联系电话
}
type SetMerchantContactReq struct {
	Contact SetMerchantContactReqContact `json:"contact" mapstructure:"contact"` // 联系方式信息
}
type SetMerchantContactRes struct {
	Errcode string `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
}

// 查询商户联系方式
type QueryMerchantContactResContact struct {
	TimeOut int    `json:"time_out" mapstructure:"time_out"` // 开票超时时间
	Phone   string `json:"phone" mapstructure:"phone"`       // 联系电话
}
type QueryMerchantContactRes struct {
	Errcode string                         `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string                         `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
	Contact QueryMerchantContactResContact `json:"contact" mapstructure:"contact"` // 联系方式信息
}

// 获取授权页 ticket
type GetTicketRes struct {
	Errcode   int    `json:"errcode" mapstructure:"errcode"`       // 错误码，含义见错误码
	Errmsg    string `json:"errmsg" mapstructure:"errmsg"`         // 错误信息，含义见错误码
	Ticket    string `json:"ticket" mapstructure:"ticket"`         // 临时票据，用于在获取授权链接时作为参数传入
	ExpiresIn int    `json:"expires_in" mapstructure:"expires_in"` // ticket 的有效期，一般为 7200 秒
}

// 获取授权页链接
type GetAuthUrlReq struct {
	SPappId     string `json:"s_pappid" mapstructure:"s_pappid"`         // 开票平台在微信的标识号，商户需要找开票平台提供
	OrderId     string `json:"order_id" mapstructure:"order_id"`         // 订单id，在商户内单笔开票请求的唯一识别号，
	Money       int    `json:"money" mapstructure:"money"`               // 订单金额，以分为单位
	Timestamp   int    `json:"timestamp" mapstructure:"timestamp"`       // 时间戳
	Source      string `json:"source" mapstructure:"source"`             // 开票来源，app：app开票，web：微信h5开票，wxa：小程序开发票，wap：普通网页开票
	RedirectUrl string `json:"redirect_url" mapstructure:"redirect_url"` // 授权成功后跳转页面。本字段只有在source为H5的时候需要填写，引导用户在微信中进行下一步流程。app开票因为从外部app拉起微信授权页，授权完成后自动回到原来的app，故无需填写。
	Ticket      string `json:"ticket" mapstructure:"ticket"`             // 从上一环节中获取
	Type        int    `json:"type" mapstructure:"type"`                 // 授权类型，0：开票授权，1：填写字段开票授权，2：领票授权
}
type GetAuthUrlRes struct {
	Errcode int    `json:"errcode" mapstructure:"errcode"`   // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`     // 错误信息
	AuthUrl string `json:"auth_url" mapstructure:"auth_url"` // 授权链接
	AppId   string `json:"appid" mapstructure:"appid"`       // source为wxa时才有
}

// 查询用户授权完成状态
type GetUserAuthStatusReq struct {
	OrderId string `json:"order_id" mapstructure:"order_id"` // 发票order_id
	SPappId string `json:"s_pappid" mapstructure:"s_pappid"` // 开票平台在微信的标识，由开票平台告知商户
}
type GetUserAuthStatusResCustomField struct {
	Key   string `json:"key" mapstructure:"key"`     // key
	Value string `json:"value" mapstructure:"value"` // value
}
type GetUserAuthStatusResUserField struct {
	Title       string                            `json:"title" mapstructure:"title"`               // title
	Phone       string                            `json:"phone" mapstructure:"phone"`               // phone
	Email       string                            `json:"email" mapstructure:"email"`               // email
	CustomField []GetUserAuthStatusResCustomField `json:"custom_field" mapstructure:"custom_field"` // custom_field
}
type GetUserAuthStatusResBizField struct {
	Title       string                            `json:"title" mapstructure:"title"`               // title
	TaxNo       string                            `json:"tax_no" mapstructure:"tax_no"`             // tax_no
	Addr        string                            `json:"addr" mapstructure:"addr"`                 // addr
	Phone       string                            `json:"phone" mapstructure:"phone"`               // phone
	BankType    string                            `json:"bank_type" mapstructure:"bank_type"`       // bank_type
	BankNo      string                            `json:"bank_no" mapstructure:"bank_no"`           // bank_no
	CustomField []GetUserAuthStatusResCustomField `json:"custom_field" mapstructure:"custom_field"` // custom_field
}
type GetUserAuthStatusResUserAuthInfo struct {
	UserField GetUserAuthStatusResUserField `json:"user_field" mapstructure:"user_field"` // 个人抬头
	BizField  GetUserAuthStatusResBizField  `json:"biz_field" mapstructure:"biz_field"`   // 单位抬头
}
type GetUserAuthStatusRes struct {
	Errcode       int                              `json:"errcode" mapstructure:"errcode"`               // 错误码
	Errmsg        string                           `json:"errmsg" mapstructure:"errmsg"`                 // 错误信息
	InvoiceStatus string                           `json:"invoice_status" mapstructure:"invoice_status"` // 订单授权状态，当errcode为0时会出现
	AuthTime      int                              `json:"auth_time" mapstructure:"auth_time"`           // 授权时间，为十位时间戳（utc+8），当errcode为0时会出现
	UserAuthInfo  GetUserAuthStatusResUserAuthInfo `json:"user_auth_info" mapstructure:"user_auth_info"` // 用户授权信息结构体，仅在授权页为type=1时出现
}

// 创建发票卡券模板
type CreateInvoiceCardTemplateReqBaseInfo struct {
	LogoUrl              string `json:"logo_url" mapstructure:"logo_url"`                               // 发票商家 LOGO ，请参考 新增永久素材
	Title                string `json:"title" mapstructure:"title"`                                     // 收款方（显示在列表），上限为 9 个汉字，建议填入商户简称
	CustomUrlName        string `json:"custom_url_name" mapstructure:"custom_url_name"`                 // 开票平台自定义入口名称，与 custom_url 字段共同使用，长度限制在 5 个汉字内
	CustomUrl            string `json:"custom_url" mapstructure:"custom_url"`                           // 开票平台自定义入口跳转外链的地址链接 , 发票外跳的链接会带有发票参数，用于标识是从哪张发票跳出的链接
	CustomUrlSubTitle    string `json:"custom_url_sub_title" mapstructure:"custom_url_sub_title"`       // 显示在入口右侧的 tips ，长度限制在 6 个汉字内
	PromotionUrlName     string `json:"promotion_url_name" mapstructure:"promotion_url_name"`           // 营销场景的自定义入口
	PromotionUrl         string `json:"promotion_url" mapstructure:"promotion_url"`                     // 入口跳转外链的地址链接，发票外跳的链接会带有发票参数，用于标识是从那张发票跳出的链接
	PromotionUrlSubTitle string `json:"promotion_url_sub_title" mapstructure:"promotion_url_sub_title"` // 显示在入口右侧的 tips ，长度限制在 6 个汉字内
}
type CreateInvoiceCardTemplateReqInvoiceInfo struct {
	BaseInfo CreateInvoiceCardTemplateReqBaseInfo `json:"base_info" mapstructure:"base_info"` // 发票卡券模板基础信息
	Payee    string                               `json:"payee" mapstructure:"payee"`         // 收款方（开票方）全称，显示在发票详情内。故建议一个收款方对应一个发票卡券模板
	Type     string                               `json:"type" mapstructure:"type"`           // 发票类型
}
type CreateInvoiceCardTemplateReq struct {
	InvoiceInfo CreateInvoiceCardTemplateReqInvoiceInfo `json:"invoice_info" mapstructure:"invoice_info"` // 发票模板对象
}
type CreateInvoiceCardTemplateRes struct {
	Errcode string `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
	CardId  string `json:"card_id" mapstructure:"card_id"` // 当错误码为 0 时，返回发票卡券模板的编号，用于后续该商户发票生成后，作为必填参数在调用插卡接口时传入
}

// 上传发票 PDF
type UploadInvoicePdfReq struct {
	Pdf *os.File `json:"pdf" mapstructure:"pdf"` // form-data中媒体文件标识，有filename、filelength、content-type等信息
}
type UploadInvoicePdfRes struct {
	Errcode  int    `json:"errcode" mapstructure:"errcode"`       // 错误码
	Errmsg   string `json:"errmsg" mapstructure:"errmsg"`         // 错误信息
	SMediaId string `json:"s_media_id" mapstructure:"s_media_id"` // 64位整数，在 将发票卡券插入用户卡包 时使用用于关联pdf和发票卡券，s_media_id有效期有3天，3天内若未将s_media_id关联到发票卡券，pdf将自动销毁
}

// 查询已上传的PDF文件
type QueryInvoicePdfReq struct {
	SMediaId string `json:"s_media_id" mapstructure:"s_media_id"` // s_media_id
}
type QueryInvoicePdfRes struct {
	Errcode          int    `json:"errcode" mapstructure:"errcode"`                         // 错误码
	Errmsg           string `json:"errmsg" mapstructure:"errmsg"`                           // 错误信息
	PdfUrl           string `json:"pdf_url" mapstructure:"pdf_url"`                         // pdf 的 url ，两个小时有效期
	PdfUrlExpireTime int    `json:"pdf_url_expire_time" mapstructure:"pdf_url_expire_time"` // pdf_url 过期时间， 7200 秒
}

// 将电子发票卡券插入用户卡包
type InsertInvoiceToUserCardReqCardInfo struct {
	Name  string `json:"name" mapstructure:"name"`   // 是 项目的名称
	Num   int    `json:"num" mapstructure:"num"`     // 否 项目的数量
	Unit  string `json:"unit" mapstructure:"unit"`   // 否 项目的单位，如个
	Price int    `json:"price" mapstructure:"price"` // 是 项目的单价
}
type InsertInvoiceToUserCardReqUserCard struct {
	Fee                   int                                  `json:"fee" mapstructure:"fee"`                                           // 发票的金额，以分为单位
	Title                 string                               `json:"title" mapstructure:"title"`                                       // 发票的抬头
	BillingTime           int                                  `json:"billing_time" mapstructure:"billing_time"`                         // 发票的开票时间，为10位时间戳（utc+8）
	BillingNo             string                               `json:"billing_no" mapstructure:"billing_no"`                             // 发票的发票号码；数电发票传20位发票号码
	BillingCode           string                               `json:"billing_code" mapstructure:"billing_code"`                         // 发票的发票代码；数电发票发票代码为空
	Info                  []InsertInvoiceToUserCardReqCardInfo `json:"info" mapstructure:"info"`                                         // 商品详情结构，见下方
	FeeWithoutTax         int                                  `json:"fee_without_tax" mapstructure:"fee_without_tax"`                   // 不含税金额，以分为单位
	Tax                   int                                  `json:"tax" mapstructure:"tax"`                                           // 税额，以分为单位
	SPdfMediaId           string                               `json:"s_pdf_media_id" mapstructure:"s_pdf_media_id"`                     // 发票pdf文件上传到微信发票平台后，会生成一个发票s_media_id，该s_media_id可以直接用于关联发票PDF和发票卡券。发票上传参考“ 3 上传PDF ”一节
	STripPdfMediaId       string                               `json:"s_trip_pdf_media_id" mapstructure:"s_trip_pdf_media_id"`           // 其它消费附件的PDF，如行程单、水单等，PDF上传方式参考“ 3 上传PDF ”一节
	CheckCode             string                               `json:"check_code" mapstructure:"check_code"`                             // 校验码，发票pdf右上角，开票日期下的校验码；数电发票发票校验码为空
	BuyerNumber           string                               `json:"buyer_number" mapstructure:"buyer_number"`                         // 购买方纳税人识别号
	BuyerAddressAndPhone  string                               `json:"buyer_address_and_phone" mapstructure:"buyer_address_and_phone"`   // 购买方地址、电话
	BuyerBankAccount      string                               `json:"buyer_bank_account" mapstructure:"buyer_bank_account"`             // 购买方开户行及账号
	SellerNumber          string                               `json:"seller_number" mapstructure:"seller_number"`                       // 销售方纳税人识别号
	SellerAddressAndPhone string                               `json:"seller_address_and_phone" mapstructure:"seller_address_and_phone"` // 销售方地址、电话
	SellerBankAccount     string                               `json:"seller_bank_account" mapstructure:"seller_bank_account"`           // 销售方开户行及账号
	Remarks               string                               `json:"remarks" mapstructure:"remarks"`                                   // 备注，发票右下角初
	Cashier               string                               `json:"cashier" mapstructure:"cashier"`                                   // 收款人，发票左下角处
	Maker                 string                               `json:"maker" mapstructure:"maker"`                                       // 开票人，发票下方处
}
type InsertInvoiceToUserCardReqCardExt struct {
	NonceStr string                             `json:"nonce_str" mapstructure:"nonce_str"` // 随机字符串，防止重复
	UserCard InsertInvoiceToUserCardReqUserCard `json:"user_card" mapstructure:"user_card"` // 用户信息结构体
}
type InsertInvoiceToUserCardReq struct {
	OrderId string                            `json:"order_id" mapstructure:"order_id"` // 发票order_id，既商户给用户授权开票的订单号
	CardId  string                            `json:"card_id" mapstructure:"card_id"`   // 发票card_id
	AppId   string                            `json:"appid" mapstructure:"appid"`       // 该订单号授权时使用的appid，一般为商户appid
	CardExt InsertInvoiceToUserCardReqCardExt `json:"card_ext" mapstructure:"card_ext"` // 发票具体内容
}
type InsertInvoiceToUserCardRes struct {
	Errcode int    `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
	Code    string `json:"code" mapstructure:"code"`       // 发票code
	OpenId  string `json:"openid" mapstructure:"openid"`   // 获得发票用户的openid
	Unionid string `json:"unionid" mapstructure:"unionid"` // 只有在用户将公众号绑定到微信开放平台账号后，才会出现该字段
}

// 更新发票卡券状态
type UpdateInvoiceStatusReq struct {
	CardId          string `json:"card_id" mapstructure:"card_id"`                   // 发票 id
	Code            string `json:"code" mapstructure:"code"`                         // 发票 code
	ReimburseStatus string `json:"reimburse_status" mapstructure:"reimburse_status"` // 发票报销状态
}
type UpdateInvoiceStatusRes struct {
	Errcode int    `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
}

// 解码 code 接口
type DecodeInvoiceCodeReq struct {
	EncryptCode string `json:"encrypt_code" mapstructure:"encrypt_code"` // 在发票卡券发起访问外链的时候后缀的加密发票code，指向一张具体的发票卡券
}
type DecodeInvoiceCodeRes struct {
	Errcode int    `json:"errcode" mapstructure:"errcode"` // 错误码
	Errmsg  string `json:"errmsg" mapstructure:"errmsg"`   // 错误信息
	Code    string `json:"code" mapstructure:"code"`       // 解密后获取的真实发票卡券Code码
}
