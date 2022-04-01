package client

import (
	"context"
	"fmt"
	"gitee.com/qhr0605/blockchain-middleware-sdk/config"
	rpc_model "gitee.com/qhr0605/blockchain-middleware-sdk/model/rpc-model"
	"testing"
)

func TestMiddleWareClient_CallHandleQueryBills(t *testing.T) {
	config.Init(false, "", "")
	client, err := NewClient()
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	responseWithPageMeta, err := client.CallHandleQueryBills(context.Background(), &rpc_model.QueryBillWithPage{
		QueryString: "{\"selector\":{\"docType\":\"bill\"}}",
		PageSize:    3,
		BookMark:    "g1AAAAA-eJzLYWBgYMpgSmHgKy5JLCrJTq2MT8lPzkzJBYqzGxmbmJqZW4CkOWDSyBJZAPEID6k",
	})
	if err != nil {
		fmt.Printf("call query Bills faield! err : %s", err.Error())
		t.FailNow()
	}
	fmt.Printf("%+v", *responseWithPageMeta)
}

func TestMiddleWareClient_CallHandleUpdateStatus(t *testing.T) {
}

func TestMiddleWareClient_CallHandleInitBill(t *testing.T) {
	config.Init(false, "", "")
	client, err := NewClient()
	if err != nil {
		fmt.Printf(err.Error())
		t.FailNow()
	}
	var remark = "测试用"
	var address = "http://hnu.edu.cn"
	bill := rpc_model.Bill{
		DocType:            "bill",
		BillStatus:         "0",
		BillType:           "增值税专用电子发票",
		MachineId:          "283737611",
		BillCode:           "123456",
		BillNumber:         "8383837",
		CheckCode:          "9918 1212 2123467 323",
		TotalAmount:        "788.0",
		IssueDate:          "2022-04-01",
		InvoiceTime:        "21:23:00",
		InvoicingPartyCode: "29129291",
		InvoicingPartyName: "232323",
		RecName:            "中国科学院",
		RecAcct:            "121289",
		RecOpBk:            "中国银行北京分行",
		RecAddress:         "三里屯",
		PayerPartyType:     "2",
		PayerPartyCode:     "92988123",
		PayerPartyName:     "湖南大学",
		PayerAcct:          "94947292",
		PayerOpBk:          "中国银行湖南大学分行",
		PayerAddress:       "麓山南路",
		ItemCode:           "23223",
		ItemName:           "机器人手臂",
		ItemQuantity:       2,
		ItemUnitPrice:      200.00,
		ItemUnit:           "件",
		ItemAmount:         454.0,
		ItemTaxRate:        0.13,
		ItemTaxAmount:      54.0,
		Remarks:            remark,
		MemoryAddress:      address,
		HandlingPerson:     "管理员",
		Checker:            "管理员",
	}
	response, err := client.CallHandleInitBill(context.Background(), &bill)
	if err != nil {
		return
	}
	fmt.Printf("%+v", response)

}
