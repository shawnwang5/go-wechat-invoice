package officialaccount

import (
	"net/http"

	"github.com/shawnwang5/go-wechat-invoice/config"
	officialAccountModel "github.com/shawnwang5/go-wechat-invoice/model/officialAccount"
	"github.com/shawnwang5/go-wechat-invoice/utils"
	requestUtils "github.com/shawnwang5/go-wechat-invoice/utils/request"
)

const (
	BASE_URL = config.BASE_URL
)

// GetPreInvoicingUrl 获取预开票 URL
//
// 参数：
//   - accessToken: access_token
//
// 返回值：
//   - res: 响应数据
//   - err: error
func GetPreInvoicingUrl(accessToken string) (res *officialAccountModel.GetPreInvoicingUrlRes, err error) {
	url := BASE_URL + "/card/invoice/seturl?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	res, err = requestUtils.HttpRequest[officialAccountModel.GetPreInvoicingUrlRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SetMerchantContact 设置商户联系方式
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func SetMerchantContact(accessToken string, req officialAccountModel.SetMerchantContactReq) (res *officialAccountModel.SetMerchantContactRes, err error) {
	url := BASE_URL + "/card/invoice/setbizattr?action=set_contact&access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	contact := make(map[string]interface{})
	contact["time_out"] = req.Contact.TimeOut
	contact["phone"] = req.Contact.Phone
	data["contact"] = contact
	res, err = requestUtils.HttpRequest[officialAccountModel.SetMerchantContactRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryMerchantContact 查询商户联系方式
//
// 参数：
//   - accessToken: access_token
//
// 返回值：
//   - res: 响应数据
//   - err: error
func QueryMerchantContact(accessToken string) (res *officialAccountModel.QueryMerchantContactRes, err error) {
	url := BASE_URL + "/card/invoice/setbizattr?action=set_contact&access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	res, err = requestUtils.HttpRequest[officialAccountModel.QueryMerchantContactRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetTicket 获取授权页 ticket
//
// 参数：
//   - accessToken: access_token
//
// 返回值：
//   - res: 响应数据
//   - err: error
func GetTicket(accessToken string) (res *officialAccountModel.GetTicketRes, err error) {
	url := BASE_URL + "/cgi-bin/ticket/getticket?type=wx_card&access_token=" + accessToken
	method := http.MethodGet
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	res, err = requestUtils.HttpRequest[officialAccountModel.GetTicketRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAuthUrl 获取授权页链接
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func GetAuthUrl(accessToken string, req officialAccountModel.GetAuthUrlReq) (res *officialAccountModel.GetAuthUrlRes, err error) {
	url := BASE_URL + "/card/invoice/setbizattr?action=set_contact&access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := utils.StructToMap(req)
	res, err = requestUtils.HttpRequest[officialAccountModel.GetAuthUrlRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserAuthStatus 查询用户授权完成状态
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func GetUserAuthStatus(accessToken string, req officialAccountModel.GetUserAuthStatusReq) (res *officialAccountModel.GetUserAuthStatusRes, err error) {
	url := BASE_URL + "/card/invoice/getauthdata?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := utils.StructToMap(req)
	res, err = requestUtils.HttpRequest[officialAccountModel.GetUserAuthStatusRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateInvoiceCardTemplate 创建发票卡券模板
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func CreateInvoiceCardTemplate(accessToken string, req officialAccountModel.CreateInvoiceCardTemplateReq) (res *officialAccountModel.CreateInvoiceCardTemplateRes, err error) {
	url := BASE_URL + "/card/invoice/platform/createcard?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := make(map[string]interface{})
	dataBaseInfo := utils.StructToMap(req.InvoiceInfo.BaseInfo)
	dataInvoiceInfo := utils.StructToMap(req.InvoiceInfo)
	dataInvoiceInfo["base_info"] = dataBaseInfo
	data["invoice_info"] = dataInvoiceInfo
	res, err = requestUtils.HttpRequest[officialAccountModel.CreateInvoiceCardTemplateRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UploadInvoicePdf 上传发票 PDF
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func UploadInvoicePdf(accessToken string, req officialAccountModel.UploadInvoicePdfReq) (res *officialAccountModel.UploadInvoicePdfRes, err error) {
	url := BASE_URL + "/card/invoice/platform/setpdf?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	headers["Content-Type"] = "multipart/form-data"
	params := make(map[string]string)
	data := make(map[string]interface{})
	data["pdf"] = req.Pdf
	res, err = requestUtils.HttpRequest[officialAccountModel.UploadInvoicePdfRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QueryInvoicePdf 查询已上传的 PDF 文件
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func QueryInvoicePdf(accessToken string, req officialAccountModel.QueryInvoicePdfReq) (res *officialAccountModel.QueryInvoicePdfRes, err error) {
	url := BASE_URL + "/card/invoice/platform/getpdf?action=get_url&access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := utils.StructToMap(req)
	data["action"] = "get_url"
	res, err = requestUtils.HttpRequest[officialAccountModel.QueryInvoicePdfRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InsertInvoiceToUserCard 将电子发票卡券插入用户卡包
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func InsertInvoiceToUserCard(accessToken string, req officialAccountModel.InsertInvoiceToUserCardReq) (res *officialAccountModel.InsertInvoiceToUserCardRes, err error) {
	url := BASE_URL + "/card/invoice/insert?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	dataUserCardInfoList := make([]map[string]interface{}, 0)
	for _, item := range req.CardExt.UserCard.Info {
		dataUserCardInfoList = append(dataUserCardInfoList, utils.StructToMap(item))
	}
	dataUserCard := utils.StructToMap(req.CardExt.UserCard)
	dataUserCard["info"] = dataUserCardInfoList
	dataCardExt := utils.StructToMap(req.CardExt)
	dataCardExt["user_card"] = dataUserCard
	data := utils.StructToMap(req)
	data["card_ext"] = dataCardExt
	res, err = requestUtils.HttpRequest[officialAccountModel.InsertInvoiceToUserCardRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateInvoiceStatus 更新发票卡券状态
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func UpdateInvoiceStatus(accessToken string, req officialAccountModel.UpdateInvoiceStatusReq) (res *officialAccountModel.UpdateInvoiceStatusRes, err error) {
	url := BASE_URL + "/card/invoice/platform/updatestatus?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := utils.StructToMap(req)
	res, err = requestUtils.HttpRequest[officialAccountModel.UpdateInvoiceStatusRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DecodeInvoiceCode 解码 code 接口
//
// 参数：
//   - accessToken: access_token
//   - req: 请求数据
//
// 返回值：
//   - res: 响应数据
//   - err: error
func DecodeInvoiceCode(accessToken string, req officialAccountModel.DecodeInvoiceCodeReq) (res *officialAccountModel.DecodeInvoiceCodeRes, err error) {
	url := BASE_URL + "/card/invoice/platform/updatestatus?access_token=" + accessToken
	method := http.MethodPost
	headers := make(map[string]string)
	params := make(map[string]string)
	data := utils.StructToMap(req)
	res, err = requestUtils.HttpRequest[officialAccountModel.DecodeInvoiceCodeRes](url,
		method,
		headers,
		params,
		data,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
