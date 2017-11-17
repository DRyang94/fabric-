package parse

import (
	"bytes"
	"encoding/xml"
	//"fmt"
	"log"
	//"reflect"
)

type MsgHead struct {
	XMLName    xml.Name   `xml:"MsgHead"`
	MsgVer     MsgVer     `xml:"MsgVer"`
	MsgId      MsgId      `xml:"MsgId"`
	MsgType    MsgType    `xml:"MsgType"`
	ProcId     ProcId     `xml:"ProcId"`
	ProcTime   ProcTime   `xml:"ProcTime"`
	Sender     Sender     `xml:"Sender"`
	SenderId   SenderId   `xml:"SenderId"`
	Receiver   Receiver   `xml:"Receiver"`
	ReceiverId ReceiverId `xml:"ReceiverId"`
	RspCode    RspCode    `xml:"RspCode"`
	ErrMsg     ErrMsg     `xml:"ErrMsg"`
	SignData   SignData   `xml:"SignData"`
}

type StringMsgReq0007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0007 `xml:"MsgCont"`
}

type StringMsgRsp0007 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0007 `xml:"MsgCont"`
}

type StringMsgReq0011 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0011 `xml:"MsgCont"`
}

type StringMsgRsp0011 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0011 `xml:"MsgCont"`
}

type StringMsgReq0012 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0012 `xml:"MsgCont"`
}

type StringMsgRsp0012 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0012 `xml:"MsgCont"`
}

type StringMsgReq0013 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0013 `xml:"MsgCont"`
}

type StringMsgReq0014 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0014 `xml:"MsgCont"`
}
type StringMsgReq0015 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0015 `xml:"MsgCont"`
}

type StringMsgReq0001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0001 `xml:"MsgCont"`
}

type StringMsgRsp0001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0001 `xml:"MsgCont"`
}

type StringMsgReq0002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0002 `xml:"MsgCont"`
}

type StringMsgRsp0002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0002 `xml:"MsgCont"`
}

type StringMsgRsp1001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1001 `xml:"MsgCont"`
}

type StringMsgRsp1002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1002 `xml:"MsgCont"`
}

type StringMsgRsp1003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp1003 `xml:"MsgCont"`
}

// mengyang start
type StringMsgReq0004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0004 `xml:"MsgCont"`
}

//贷款审核拒绝
type StringMsgReq0006 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0006 `xml:"MsgCont"`
}

//贷款放款
type StringMsgReq0003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0003 `xml:"MsgCont"`
}

//贷款还款期数跟新
type StringMsgReq0005 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0005 `xml:"MsgCont"`
}

//贷款还款期数跟新
type StringMsgReq0008 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0008 `xml:"MsgCont"`
}

//贷款还款
type StringMsgReq0010 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0010 `xml:"MsgCont"`
}

//贷款审核通过
type StringMsgRsp0004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0004 `xml:"MsgCont"`
}

//贷款审核拒绝
type StringMsgRsp0006 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0006 `xml:"MsgCont"`
}

//贷款还款
type StringMsgRsp0003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0003 `xml:"MsgCont"`
}

//贷款还款期数跟新
type StringMsgRsp0005 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0005 `xml:"MsgCont"`
}
type StringMsgRsp0008 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0008 `xml:"MsgCont"`
}
type StringMsgRsp0010 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp0010 `xml:"MsgCont"`
}

type StringMsgReq0009 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq0009 `xml:"MsgCont"`
}

type StringMsgRsp2001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp2001 `xml:"MsgCont"`
}

type StringMsgReq2001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq2001 `xml:"MsgCont"`
}

type StringMsgRsp2002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContRsp2002 `xml:"MsgCont"`
}

type StringMsgReq2002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq2002 `xml:"MsgCont"`
}

type StringMsgReq2003 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq2003 `xml:"MsgCont"`
}

type StringMsgReq2004 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq2004 `xml:"MsgCont"`
}

type StringMsgReq2005 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq2005 `xml:"MsgCont"`
}

type StringMsgReq3001 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq3001 `xml:"MsgCont"`
}

type MsgContReq3001 struct {
	XMLName     xml.Name    `xml:"MsgCont"`
	BlockHeight BlockHeight `xml:"BlockHeight"`
}

type StringMsgReq3002 struct {
	XMLName    xml.Name       `xml:"Msg"`
	MsgMsgHead MsgHead        `xml:"MsgHead"`
	MsgMsgCont MsgContReq3002 `xml:"MsgCont"`
}

//贷款审核通过
type MsgContReq0004 struct {
	XMLName  xml.Name    `xml:"MsgCont"`
	LoanInfo LoanInfo004 `xml:"LoanInfo"`
}

type MsgContReq0013 struct {
	XMLName xml.Name `xml:"MsgCont"`
	TxId    TxId     `xml:"TxId"`
}

type MsgContReq0014 struct {
	XMLName xml.Name    `xml:"MsgCont"`
	OrgInfo OrgInfo0014 `xml:"OrgInfo"`
}
type MsgContReq2001 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
}
type MsgContRsp2001 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
	ChainCodeStatus  ChainCodeStatus  `xml:"ChainCodeStatus"`
}

type MsgContReq2002 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
}
type MsgContRsp2002 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
	ChainCodeStatus  ChainCodeStatus  `xml:"ChainCodeStatus"`
}

type MsgContReq2003 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
}
type MsgContReq2004 struct {
	XMLName xml.Name `xml:"MsgCont"`
	PriKey  PriKey   `xml:"PriKey"`
	PubKey  PubKey   `xml:"PubKey"`
	KeySize KeySize  `xml:"KeySize"`
	NodeId  NodeId   `xml:"NodeId"`
}

type MsgContReq2005 struct {
	XMLName xml.Name `xml:"MsgCont"`
	//	ChainCodeID      ChainCodeID      `xml:"ChainCodeID"`
	//	ChainCodePath    ChainCodePath    `xml:"ChainCodePath"`
	//	ChainCodeVersion ChainCodeVersion `xml:"ChainCodeVersion"`
}

type MsgContReq3002 struct {
	XMLName  xml.Name `xml:"MsgCont"`
	BlockNum BlockNum `xml:"BlockNum"`
}

type LoanInfo004 struct {
	XMLName    xml.Name   `xml:"LoanInfo"`
	LenderName LenderName `xml:"LenderName"`
	IdType     IdType     `xml:"IdType"`
	IdNo       IdNo       `xml:"IdNo"`
	LoanNo     LoanNo     `xml:"LoanNo"`
	LoanStatus LoanStatus `xml:"LoanStatus"`
	LoanAmt    LoanAmt    `xml:"LoanAmt"`
	//	RepayType   RepayType   `xml:"RepayType"`
	//	RepayPeriod RepayPeriod `xml:"RepayPeriod"`
	//	RepayAmt    RepayAmt    `xml:"RepayAmt"`
}

type MsgContReq0006 struct {
	XMLName     xml.Name     `xml:"MsgCont"`
	LoanInfoRef LoanInfo0006 `xml:"LoanInfo"`
}
type LoanInfo0006 struct {
	XMLName    xml.Name   `xml:"LoanInfo"`
	LenderName LenderName `xml:"LenderName"`
	IdType     IdType     `xml:"IdType"`
	IdNo       IdNo       `xml:"IdNo"`
	LoanNo     LoanNo     `xml:"LoanNo"`
	LoanStatus LoanStatus `xml:"LoanStatus"`
	LoanType   LoanType   `xml:"LoanType"`
	LoanDesc   LoanDesc   `xml:"LoanDesc"`
	RejMsg     RejMsg     `xml:"RejMsg"`
}

type MsgContReq0003 struct {
	XMLName             xml.Name            `xml:"MsgCont"`
	LoanBaseInfo0003req LoanBaseInfo0003req `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0003req struct {
	XMLName     xml.Name    `xml:"LoanBaseInfo"`
	LoanNo      LoanNo      `xml:"LoanNo"`
	ReleaseNo   ReleaseNo   `xml:"ReleaseNo"`
	Rate        Rate        `xml:"Rate"`
	LoanAmt     LoanAmt     `xml:"LoanAmt"`
	RepayType   RepayType   `xml:"RepayType"`
	RepayPeriod RepayPeriod `xml:"RepayPeriod"`
	RepayAmt    RepayAmt    `xml:"RepayAmt"`
}

type MsgContReq0005 struct {
	XMLName             xml.Name            `xml:"MsgCont"`
	LoanBaseInfo0005req LoanBaseInfo0005req `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0005req struct {
	XMLName     xml.Name    `xml:"LoanBaseInfo"`
	LoanNo      LoanNo      `xml:"LoanNo"`
	RepayPeriod RepayPeriod `xml:"RepayPeriod"`
	ReleaseNo   ReleaseNo   `xml:"ReleaseNo"`
}

type MsgContReq0009 struct {
	XMLName      xml.Name            `xml:"MsgCont"`
	LoanBaseInfo LoanBaseInfo0009Req `xml:"LoanBaseInfo"`
}

type LoanBaseInfo0009Req struct {
	XMLName  xml.Name `xml:"LoanBaseInfo"`
	LoanNo   LoanNo   `xml:"LoanNo"`
	IdType   IdType   `xml:"IdType"`
	IdNo     IdNo     `xml:"IdNo"`
	LoanType LoanType `xml:"LoanType"`
}

type MsgContReq0008 struct {
	XMLName      xml.Name     `xml:"MsgCont"`
	LoanInfo0008 LoanInfo0008 `xml:"LoanInfo"`
}
type LoanInfo0008 struct {
	XMLName   xml.Name  `xml:"LoanInfo"`
	LoanNo    LoanNo    `xml:"LoanNo"`
	IdType    IdType    `xml:"IdType"`
	IdNo      IdNo      `xml:"IdNo"`
	Amt       Amt       `xml:"Amt"`
	ReleaseNo ReleaseNo `xml:"ReleaseNo"`
}
type MsgContReq0010 struct {
	XMLName        xml.Name       `xml:"MsgCont"`
	LenderInfo0010 LenderInfo0010 `xml:"LenderInfo"`
	Repay0010      Repay0010      `xml:"Repay"`
}

type LenderInfo0010 struct {
	XMLName xml.Name `xml:"LenderInfo"`
	//LoanId   LoanId   `xml:"LoanId"`
	IdType IdType `xml:"IdType"`
	IdNo   IdNo   `xml:"IdNo"`
}
type Repay0010 struct {
	XMLName     xml.Name    `xml:"Repay"`
	LoanNo      LoanNo      `xml:"LoanNo"`
	RepayPlanNo RepayPlanNo `xml:"RepayPlanNo"`
	RepayAmt    RepayAmt    `xml:"RepayAmt"`
	Interest    Interest    `xml:"Interest"`
	FineAmt     FineAmt     `xml:"FineAmt"`
	Amt         Amt         `xml:"Amt"`
	RepayStatus RepayStatus `xml:"RepayStatus"`
	RepayDate   RepayDate   `xml:"RepayDate"`
	ReleaseNo   ReleaseNo   `xml:"ReleaseNo"`
}

type MsgContRsp0004 struct {
	XMLName      xml.Name         `xml:"MsgCont"`
	LoanBaseInfo LoanBaseInfo0004 `xml:"LoanBaseInfo"`
}

type LoanBaseInfo0004 struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}

type MsgContRsp0006 struct {
	XMLName         xml.Name         `xml:"MsgCont"`
	LoanBaseInfoRef LoanBaseInfo0006 `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0006 struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}
type MsgContRsp0003 struct {
	XMLName             xml.Name            `xml:"MsgCont"`
	LoanBaseInfo0003rsp LoanBaseInfo0003rsp `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0003rsp struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}
type MsgContRsp0005 struct {
	XMLName             xml.Name            `xml:"MsgCont"`
	LoanBaseInfo0005rsp LoanBaseInfo0005rsp `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0005rsp struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}
type MsgContRsp0008 struct {
	XMLName          xml.Name         `xml:"MsgCont"`
	LoanBaseInfo0008 LoanBaseInfo0008 `xml:"LoanBaseInfo"`
}
type LoanBaseInfo0008 struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}
type MsgContRsp0010 struct {
	XMLName       xml.Name      `xml:"MsgCont"`
	RepayInfo0010 RepayInfo0010 `xml:"RepayInfo"`
}

type RepayInfo0010 struct {
	XMLName xml.Name `xml:"RepayInfo"`
	RepayId RepayId  `xml:"RepayId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}

//end

type MsgContReq0007 struct {
	XMLName      xml.Name            `xml:"MsgCont"`
	LoanBaseInfo LoanBaseInfoReq0007 `xml:"LoanBaseInfo"`
}

type MsgContRsp0007 struct {
	XMLName      xml.Name            `xml:"MsgCont"`
	LoanBaseInfo LoanBaseInfoRsp0007 `xml:"LoanBaseInfo"`
}

type MsgContReq0011 struct {
	XMLName     xml.Name       `xml:"MsgCont"`
	IsEndMonth  IsEndMonth     `xml:"IsEndMonth"`
	LoanInfo    LoanInfo0011   `xml:"LoanInfo"`
	OverdueInfo OverdueInfoReq `xml:"OverdueInfo"`
}

type MsgContRsp0011 struct {
	XMLName     xml.Name       `xml:"MsgCont"`
	OverdueInfo OverdueInfoRsp `xml:"OverdueInfo"`
}

type MsgContReq0012 struct {
	XMLName xml.Name `xml:"MsgCont"`
	OrgInfo OrgInfo  `xml:"OrgInfo"`
}

type MsgContRsp0012 struct {
	XMLName xml.Name   `xml:"MsgCont"`
	OrgInfo OrgInfoRsp `xml:"OrgInfo"`
}
type MsgContReq0015 struct {
	XMLName     xml.Name    `xml:"MsgCont"`
	Flag        Flag        `xml:"Flag"`
	TimeStamp   TimeStamp   `xml:"TimeStamp"`
	TSSignature TSSignature `xml:"TSSignature"`
}

type LoanInfo0011 struct {
	XMLName xml.Name `xml:"LoanInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	IdType  IdType   `xml:"IdType"`
	IdNo    IdNo     `xml:"IdNo"`
}

type OverdueInfoReq struct {
	XMLName     xml.Name    `xml:"OverdueInfo"`
	LoanId      LoanId      `xml:"LoanId"`
	RepayPlanNo RepayPlanNo `xml:"RepayPlanNo"`
	RepayAmt    RepayAmt    `xml:"RepayAmt"`
	Interest    Interest    `xml:"Interest"`
	FineAmt     FineAmt     `xml:"FineAmt"`
	OverdueDay  OverdueDay  `xml:"OverdueDay"`
	SendDate    SendDate    `xml:"SendDate"`
	ReleaseNo   ReleaseNo   `xml:"ReleaseNo"`
}

type OverdueInfoRsp struct {
	XMLName   xml.Name  `xml:"OverdueInfo"`
	OverdueId OverdueId `xml:"OverdueId"`
	CrtTime   CrtTime   `xml:"CrtTime"`
	UpdTime   UpdTime   `xml:"UpdTime"`
}

type OrgInfo struct {
	XMLName           xml.Name          `xml:"OrgInfo"`
	OrgName           OrgName           `xml:"OrgName"`
	IdType            IdType            `xml:"IdType"`
	IdNo              IdNo              `xml:"IdNo"`
	DigitalCurrency   DigitalCurrency   `xml:"DigitalCurrency"`
	LoanTotal         LoanTotal         `xml:"LoanTotal"`
	LoanNum           LoanNum           `xml:"LoanNum"`
	OperateState      OperateState      `xml:"OperateState"`
	CancelReason      CancelReason      `xml:"CancelReason"`
	Describe          Describe          `xml:"Describe"`
	InitialReward     InitialReward     `xml:"InitialReward"`
	FollowUpReward    FollowUpReward    `xml:"FollowUpReward"`
	AmountConsume     AmountConsume     `xml:"AmountConsume"`
	InitialRewardCon  InitialRewardCon  `xml:"InitialRewardCon"`
	FollowUpRewardCon FollowUpRewardCon `xml:"FollowUpRewardCon"`
	AmountConsumeCon  AmountConsumeCon  `xml:"AmountConsumeCon"`
}
type OrgInfo0014 struct {
	XMLName xml.Name `xml:"OrgInfo"`
	IdType  IdType   `xml:"IdType"`
	IdNo    IdNo     `xml:"IdNo"`
}

type OrgInfoRsp struct {
	XMLName xml.Name `xml:"OrgInfo"`
	OrgId   OrgId    `xml:"OrgId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}

type Description struct {
	XMLName   xml.Name `xml:"Description"`
	InnerText string   `xml:",innerxml"`
}

type MsgContReq0001 struct {
	XMLName          xml.Name        `xml:"MsgCont"`
	LoanBaseInfo     LoanBaseInfo    `xml:"LoanBaseInfo"`
	LenderInfo       LenderInfo      `xml:"LenderInfo"`
	DLenderInfo      DLenderInfo     `xml:"DLenderInfo"`
	WarrantorInfoLst []WarrantorInfo `xml:"WarrantorInfo"`
	PledgeInfoLst    []PledgeInfo    `xml:"PledgeInfo"`
}

type MsgContRsp0001 struct {
	XMLName         xml.Name        `xml:"MsgCont"`
	LoanBaseInfoRes LoanBaseInfoRes `xml:"LoanBaseInfo"`
}

type MsgContReq0002 struct {
	XMLName   xml.Name  `xml:"MsgCont"`
	QueryInfo QueryInfo `xml:"QueryInfo"`
}

type MsgContRsp0002 struct {
	XMLName          xml.Name           `xml:"MsgCont"`
	LoanInfo         LoanInfo           `xml:"LoanInfo"`
	PledgeInfoLst    []PledgeInfoQry    `xml:"PledgeInfo"`
	WarrantorInfoLst []WarrantorInfoQry `xml:"WarrantorInfo"`
	ApplyInfo        ApplyInfo          `xml:"ApplyInfo"`
}

type MsgContRsp1001 struct {
	XMLName      xml.Name         `xml:"MsgCont"`
	LoanBaseInfo LoanBaseInfo1001 `xml:"LoanBaseInfo"`
	LenderInfo   LenderInfo1001   `xml:"LenderInfo"`
}

type MsgContRsp1002 struct {
	XMLName       xml.Name          `xml:"MsgCont"`
	LoanBaseInfo  LoanBaseInfo1002  `xml:"LoanBaseInfo"`
	WarrantorInfo WarrantorInfo1002 `xml:"WarrantorInfo"`
}

type MsgContRsp1003 struct {
	XMLName     xml.Name    `xml:"MsgCont"`
	IdType      IdType      `xml:"IdType"`
	IdNo        IdNo        `xml:"IdNo"`
	LoanNo      LoanNo      `xml:"LoanNo"`
	LoanStatus  LoanStatus  `xml:"LoanStatus"`
	EndType     EndType     `xml:"EndType"`
	LoanAmt     LoanAmt     `xml:"LoanAmt"`
	RepayType   RepayType   `xml:"RepayType"`
	RepayPeriod RepayPeriod `xml:"RepayPeriod"`
	RepayAmt    RepayAmt    `xml:"RepayAmt"`
	CrtTime     CrtTime     `xml:"CrtTime"`
	LoanDesc    LoanDesc    `xml:"LoanDesc"`
}

type ApplyInfo struct {
	XMLName       xml.Name      `xml:"ApplyInfo"`
	LenderName    LenderName    `xml:"LenderName"`
	IdType        IdType        `xml:"IdType"`
	IdNoS         IdNoS         `xml:"IdNoS"`
	LenderType    LenderType    `xml:"LenderType"`
	IsBlacklist   IsBlacklist   `xml:"IsBlacklist"`
	OverdueCntNum OverdueCntNum `xml:"OverdueCntNum"`
	UnpayCntNum   UnpayCntNum   `xml:"UnpayCntNum"`
	OverdueCntDay OverdueCntDay `xml:"OverdueCntDay"`
	CrtUser       CrtUser       `xml:"CrtUser"`
	CrtTime       CrtTime       `xml:"CrtTime"`
}

type LoanBaseInfo struct {
	XMLName   xml.Name  `xml:"LoanBaseInfo"`
	LoanNo    LoanNo    `xml:"LoanNo"`
	LoanType  LoanType  `xml:"LoanType"`
	ApplyAmt  ApplyAmt  `xml:"ApplyAmt"`
	LoanDesc  LoanDesc  `xml:"LoanDesc"`
	ApplyTime ApplyTime `xml:"ApplyTime"`
}

type LoanBaseInfo1001 struct {
	XMLName  xml.Name `xml:"LoanBaseInfo"`
	LoanType LoanType `xml:"LoanType"`
	ApplyAmt ApplyAmt `xml:"ApplyAmt"`
	CrtTime  CrtTime  `xml:"CrtTime"`
	LoanDesc LoanDesc `xml:"LoanDesc"`
}

type LoanBaseInfo1002 struct {
	XMLName       xml.Name      `xml:"LoanBaseInfo"`
	LoanType      LoanType      `xml:"LoanType"`
	ApplyAmt      ApplyAmt      `xml:"ApplyAmt"`
	LoanDirection LoanDirection `xml:"LoanDirection"`
	CrtTime       CrtTime       `xml:"CrtTime"`
	LoanDesc      LoanDesc      `xml:"LoanDesc"`
}

type LoanBaseInfoReq0007 struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanNo  LoanNo   `xml:"LoanNo"`
	EndType EndType  `xml:"EndType"`
	EndMsg  EndMsg   `xml:"EndMsg"`
}

type LoanBaseInfoRsp0007 struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
	UpdTime UpdTime  `xml:"UpdTime"`
}

type LoanInfo struct {
	XMLName         xml.Name        `xml:"LoanInfo"`
	OrgId           OrgId           `xml:"OrgId"`
	LoanType        LoanType        `xml:"LoanType"`
	PledgeType      PledgeType      `xml:"PledgeType"`
	PledgeNo        PledgeNo        `xml:"PledgeNo"`
	WarrantorIdType WarrantorIdType `xml:"WarrantorIdType"`
	WarrantorIdNo   WarrantorIdNo   `xml:"WarrantorIdNo"`
	LoanStatus      LoanStatus      `xml:"LoanStatus"`
	LoanTime        LoanTime        `xml:"LoanTime"`
	ApplyAmt        ApplyAmt        `xml:"ApplyAmt"`
	LoanAmt         LoanAmt         `xml:"LoanAmt"`
	UnrepayAmt      UnrepayAmt      `xml:"UnrepayAmt"`
	OverdueNum      OverdueNum      `xml:"OverdueNum"`
	OverdueDay      OverdueDay      `xml:"OverdueDay"`
}

type LenderInfo struct {
	XMLName     xml.Name    `xml:"LenderInfo"`
	LenderType  LenderType  `xml:"LenderType"`
	LenderName  LenderName  `xml:"LenderName"`
	IdType      IdType      `xml:"IdType"`
	IdNo        IdNo        `xml:"IdNo"`
	IsBlackList IsBlackList `xml:"IsBlackList"`
}

type QueryInfo struct {
	XMLName    xml.Name   `xml:"QueryInfo"`
	QueryType  QueryType  `xml:"QueryType"`
	LoanNo     LoanNo     `xml:"LoanNo"`
	IdNo       IdNo       `xml:"IdNo"`
	IdType     IdType     `xml:"IdType"`
	DIdNo      DIdNo      `xml:"DIdNo"`
	DIdType    DIdType    `xml:"DIdType"`
	WIdNo      WIdNo      `xml:"WIdNo"`
	WIdType    WIdType    `xml:"WIdType"`
	PledgeNo   PledgeNo   `xml:"PledgeNo"`
	PledgeType PledgeType `xml:"PledgeType"`
}

type LenderInfo1001 struct {
	XMLName    xml.Name   `xml:"LenderInfo"`
	LenderType LenderType `xml:"LenderType"`
	IdType     IdType     `xml:"IdType"`
	IdNo       IdNo       `xml:"IdNo"`
}

type DLenderInfo struct {
	XMLName     xml.Name    `xml:"DLenderInfo"`
	DLenderName DLenderName `xml:"DLenderName"`
	DIdType     DIdType     `xml:"DIdType"`
	DIdNo       DIdNo       `xml:"DIdNo"`
}

//type WarrantorForArray struct {
//	WarrantorName string
//	WIdType       string
//	WIdNo         string
//}

type WarrantorInfo struct {
	WarrantorName WarrantorName `xml:"WarrantorName"`
	WIdType       WIdType       `xml:"WIdType"`
	WIdNo         WIdNo         `xml:"WIdNo"`
}

type WarrantorInfoQry struct {
	XMLName         xml.Name        `xml:"Warrantor"`
	WarrantorName   WarrantorName   `xml:"WarrantorName"`
	IdType          IdType          `xml:"IdType"`
	IdNoS           IdNoS           `xml:"IdNoS"`
	WarrantorStatus WarrantorStatus `xml:"WarrantorStatus"`

	ApplyComId       ApplyComId       `xml:"ApplyComId"`
	CrtTime          CrtTime          `xml:"CrtTime"`
	WarrantorId      WarrantorId      `xml:"WarrantorId"`
	WarrantorHistory WarrantorHistory `xml:"WarrantorHistory"`
}

type WarrantorInfo1002 struct {
	XMLName xml.Name `xml:"WarrantorInfo"`
	WIdType WIdType  `xml:"WIdType"`
	WIdNo   WIdNo    `xml:"WIdNo"`
}

type WarrantorHistory struct {
	XMLName      xml.Name    `xml:"WarrantorHistory"`
	WarrantorLst []Warrantor `xml:"Warrantor"`
}

type Warrantor struct {
	ApplyTime  string `xml:"ApplyTime"`
	ApplyComId string `xml:"ApplyComId"`
	LoanStatus string `xml:"LoanStatus"`
	IdType     string `xml:"IdType"`
	IdNoS      string `xml:"IdNoS"`
	LoanTime   string `xml:"LoanTime"`
	ApplyAmt   string `xml:"ApplyAmt"`
	LoanAmt    string `xml:"LoanAmt"`
	UnrepayAmt string `xml:"UnrepayAmt"`
	OverdueNum string `xml:"OverdueNum"`
	OverdueDay string `xml:"OverdueDay"`
}

type PledgeInfo struct {
	XMLName    xml.Name   `xml:"PledgeInfo"`
	PledgeType PledgeType `xml:"PledgeType"`
	PledgeNo   PledgeNo   `xml:"PledgeNo"`
	PledgeDesc PledgeDesc `xml:"PledgeDesc"`
}

type PledgeForArray struct {
	PledgeType string
	PledgeNo   string
	PledgeDesc string
}

type PledgeInfoQry struct {
	XMLName       xml.Name      `xml:"PledgeInfo"`
	PledgeType    PledgeType    `xml:"PledgeType"`
	PledgeNo      PledgeNo      `xml:"PledgeNo"`
	PledgeDesc    PledgeDesc    `xml:"PledgeDesc"`
	PledgeStatus  PledgeStatus  `xml:"PledgeStatus"`
	ApplyComId    ApplyComId    `xml:"ApplyComId"`
	CrtTime       CrtTime       `xml:"CrtTime"`
	PledgeHistory PledgeHistory `xml:"PledgeHistory"`
}

type PledgeHistory struct {
	XMLName   xml.Name `xml:"PledgeHistory"`
	PledgeLst []Pledge `xml:"Pledge"`
}

type Pledge struct {
	ApplyTime  string `xml:"ApplyTime"`
	ApplyComId string `xml:"ApplyComId"`
	IdType     string `xml:"IdType"`
	IdNoS      string `xml:"IdNoS"`
	LoanStatus string `xml:"LoanStatus"`
	LoanTime   string `xml:"LoanTime"`
	ApplyAmt   string `xml:"ApplyAmt"`
	LoanAmt    string `xml:"LoanAmt"`
	UnrepayAmt string `xml:"UnrepayAmt"`
	OverdueNum string `xml:"OverdueNum"`
	OverdueDay string `xml:"OverdueDay"`
}

type LoanBaseInfoRes struct {
	XMLName xml.Name `xml:"LoanBaseInfo"`
	LoanId  LoanId   `xml:"LoanId"`
	CrtTime CrtTime  `xml:"CrtTime"`
}

type MsgVer struct {
	XMLName   xml.Name `xml:"MsgVer"`
	InnerText string   `xml:",innerxml"`
}

type MsgId struct {
	XMLName   xml.Name `xml:"MsgId"`
	InnerText string   `xml:",innerxml"`
}

type MsgType struct {
	XMLName   xml.Name `xml:"MsgType"`
	InnerText string   `xml:",innerxml"`
}

type ProcId struct {
	XMLName   xml.Name `xml:"ProcId"`
	InnerText string   `xml:",innerxml"`
}

type ProcTime struct {
	XMLName   xml.Name `xml:"ProcTime"`
	InnerText string   `xml:",innerxml"`
}

type Sender struct {
	XMLName   xml.Name `xml:"Sender"`
	InnerText string   `xml:",innerxml"`
}

type SenderId struct {
	XMLName   xml.Name `xml:"SenderId"`
	InnerText string   `xml:",innerxml"`
}

type Receiver struct {
	XMLName   xml.Name `xml:"Receiver"`
	InnerText string   `xml:",innerxml"`
}

type ReceiverId struct {
	XMLName   xml.Name `xml:"ReceiverId"`
	InnerText string   `xml:",innerxml"`
}

type RspCode struct {
	XMLName   xml.Name `xml:"RspCode"`
	InnerText string   `xml:",innerxml"`
}

type ErrMsg struct {
	XMLName   xml.Name `xml:"ErrMsg"`
	InnerText string   `xml:",innerxml"`
}

type Response struct {
	XMLName   xml.Name `xml:"Response"`
	InnerText string   `xml:",innerxml"`
}

type ResMsg struct {
	XMLName   xml.Name `xml:"ResMsg"`
	InnerText string   `xml:",innerxml"`
}

type OrgId struct {
	XMLName   xml.Name `xml:"OrgId"`
	InnerText string   `xml:",innerxml"`
}

type OrgName struct {
	XMLName   xml.Name `xml:"OrgName"`
	InnerText string   `xml:",innerxml"`
}

type IdType struct {
	XMLName   xml.Name `xml:"IdType"`
	InnerText string   `xml:",innerxml"`
}

type IdNoS struct {
	XMLName   xml.Name `xml:"IdNoS"`
	InnerText string   `xml:",innerxml"`
}

type Describe struct {
	XMLName   xml.Name `xml:"Describe"`
	InnerText string   `xml:",innerxml"`
}
type InitialReward struct {
	XMLName   xml.Name `xml:"InitialReward"`
	InnerText string   `xml:",innerxml"`
}
type FollowUpReward struct {
	XMLName   xml.Name `xml:"FollowUpReward"`
	InnerText string   `xml:",innerxml"`
}
type AmountConsume struct {
	XMLName   xml.Name `xml:"AmountConsume"`
	InnerText string   `xml:",innerxml"`
}

type InitialRewardCon struct {
	XMLName   xml.Name `xml:"InitialRewardCon"`
	InnerText string   `xml:",innerxml"`
}
type FollowUpRewardCon struct {
	XMLName   xml.Name `xml:"FollowUpRewardCon"`
	InnerText string   `xml:",innerxml"`
}
type AmountConsumeCon struct {
	XMLName   xml.Name `xml:"AmountConsumeCon"`
	InnerText string   `xml:",innerxml"`
}

type RejMsg struct {
	XMLName   xml.Name `xml:"RejMsg"`
	InnerText string   `xml:",innerxml"`
}

type OperateState struct {
	XMLName   xml.Name `xml:"OperateState"`
	InnerText string   `xml:",innerxml"`
}

type CancelReason struct {
	XMLName   xml.Name `xml:"CancelReason"`
	InnerText string   `xml:",innerxml"`
}
type CrtTime struct {
	XMLName   xml.Name `xml:"CrtTime"`
	InnerText string   `xml:",innerxml"`
}

type LoanId struct {
	XMLName   xml.Name `xml:"LoanId"`
	InnerText string   `xml:",innerxml"`
}
type Flag struct {
	XMLName   xml.Name `xml:"Flag"`
	InnerText string   `xml:",innerxml"`
}
type TimeStamp struct {
	XMLName   xml.Name `xml:"TimeStamp"`
	InnerText string   `xml:",innerxml"`
}
type TSSignature struct {
	XMLName   xml.Name `xml:"TSSignature"`
	InnerText string   `xml:",innerxml"`
}
type LoanNo struct {
	XMLName   xml.Name `xml:"LoanNo"`
	InnerText string   `xml:",innerxml"`
}
type ReleaseNo struct {
	XMLName   xml.Name `xml:"ReleaseNo"`
	InnerText string   `xml:",innerxml"`
}
type LoanType struct {
	XMLName   xml.Name `xml:"LoanType"`
	InnerText string   `xml:",innerxml"`
}

type ApplyAmt struct {
	XMLName   xml.Name `xml:"ApplyAmt"`
	InnerText string   `xml:",innerxml"`
}

type LoanDesc struct {
	XMLName   xml.Name `xml:"LoanDesc"`
	InnerText string   `xml:",innerxml"`
}

type LenderType struct {
	XMLName   xml.Name `xml:"LenderType"`
	InnerText string   `xml:",innerxml"`
}
type LenderName struct {
	XMLName   xml.Name `xml:"LenderName"`
	InnerText string   `xml:",innerxml"`
}

type IdNo struct {
	XMLName   xml.Name `xml:"IdNo"`
	InnerText string   `xml:",innerxml"`
}

type DLenderName struct {
	XMLName   xml.Name `xml:"DLenderName"`
	InnerText string   `xml:",innerxml"`
}

type DIdType struct {
	XMLName   xml.Name `xml:"DIdType"`
	InnerText string   `xml:",innerxml"`
}

type DIdNo struct {
	XMLName   xml.Name `xml:"DIdNo"`
	InnerText string   `xml:",innerxml"`
}

type WarrantorName struct {
	XMLName   xml.Name `xml:"WarrantorName"`
	InnerText string   `xml:",innerxml"`
}

type WIdType struct {
	XMLName   xml.Name `xml:"WIdType"`
	InnerText string   `xml:",innerxml"`
}

type WIdNo struct {
	XMLName   xml.Name `xml:"WIdNo"`
	InnerText string   `xml:",innerxml"`
}

type PledgeType struct {
	XMLName   xml.Name `xml:"PledgeType"`
	InnerText string   `xml:",innerxml"`
}

type PledgeNo struct {
	XMLName   xml.Name `xml:"PledgeNo"`
	InnerText string   `xml:",innerxml"`
}

type PledgeDesc struct {
	XMLName   xml.Name `xml:"PledgeDesc"`
	InnerText string   `xml:",innerxml"`
}

type WarrantorIdType struct {
	XMLName   xml.Name `xml:"WarrantorIdType"`
	InnerText string   `xml:",innerxml"`
}

type WarrantorIdNo struct {
	XMLName   xml.Name `xml:"WarrantorIdNo"`
	InnerText string   `xml:",innerxml"`
}

type LoanStatus struct {
	XMLName   xml.Name `xml:"LoanStatus"`
	InnerText string   `xml:",innerxml"`
}

type LoanTime struct {
	XMLName   xml.Name `xml:"LoanTime"`
	InnerText string   `xml:",innerxml"`
}

type LoanAmt struct {
	XMLName   xml.Name `xml:"LoanAmt"`
	InnerText string   `xml:",innerxml"`
}

type UnrepayAmt struct {
	XMLName   xml.Name `xml:"UnrepayAmt"`
	InnerText string   `xml:",innerxml"`
}

type OverdueNum struct {
	XMLName   xml.Name `xml:"OverdueNum"`
	InnerText string   `xml:",innerxml"`
}

type OverdueDay struct {
	XMLName   xml.Name `xml:"OverdueDay"`
	InnerText string   `xml:",innerxml"`
}

type PledgeStatus struct {
	XMLName   xml.Name `xml:"PledgeStatus"`
	InnerText string   `xml:",innerxml"`
}

type ApplyComId struct {
	XMLName   xml.Name `xml:"ApplyComId"`
	InnerText string   `xml:",innerxml"`
}

type ApplyTime struct {
	XMLName   xml.Name `xml:"ApplyTime"`
	InnerText string   `xml:",innerxml"`
}

type WarrantorId struct {
	XMLName   xml.Name `xml:"WarrantorId"`
	InnerText string   `xml:",innerxml"`
}

type CrtUser struct {
	XMLName   xml.Name `xml:"CrtUser"`
	InnerText string   `xml:",innerxml"`
}

type OverdueCntDay struct {
	XMLName   xml.Name `xml:"OverdueCntDay"`
	InnerText string   `xml:",innerxml"`
}

type UnpayCntNum struct {
	XMLName   xml.Name `xml:"UnpayCntNum"`
	InnerText string   `xml:",innerxml"`
}

type OverdueCntNum struct {
	XMLName   xml.Name `xml:"OverdueCntNum"`
	InnerText string   `xml:",innerxml"`
}

type IsBlacklist struct {
	XMLName   xml.Name `xml:"IsBlacklist"`
	InnerText string   `xml:",innerxml"`
}

type WarrantorStatus struct {
	XMLName   xml.Name `xml:"WarrantorStatus"`
	InnerText string   `xml:",innerxml"`
}

type LoanDirection struct {
	XMLName   xml.Name `xml:"LoanDirection"`
	InnerText string   `xml:",innerxml"`
}

type EndType struct {
	XMLName   xml.Name `xml:"EndType"`
	InnerText string   `xml:",innerxml"`
}

type RepayType struct {
	XMLName   xml.Name `xml:"RepayType"`
	InnerText string   `xml:",innerxml"`
}

type RepayPeriod struct {
	XMLName   xml.Name `xml:"RepayPeriod"`
	InnerText string   `xml:",innerxml"`
}

type RepayAmt struct {
	XMLName   xml.Name `xml:"RepayAmt"`
	InnerText string   `xml:",innerxml"`
}

type DigitalCurrency struct {
	XMLName   xml.Name `xml:"DigitalCurrency"`
	InnerText string   `xml:",innerxml"`
}

type LoanTotal struct {
	XMLName   xml.Name `xml:"LoanTotal"`
	InnerText string   `xml:",innerxml"`
}

type LoanNum struct {
	XMLName   xml.Name `xml:"LoanNum"`
	InnerText string   `xml:",innerxml"`
}

type UpdTime struct {
	XMLName   xml.Name `xml:"UpdTime"`
	InnerText string   `xml:",innerxml"`
}

type IsEndMonth struct {
	XMLName   xml.Name `xml:"IsEndMonth"`
	InnerText string   `xml:",innerxml"`
}

type RepayPlanNo struct {
	XMLName   xml.Name `xml:"RepayPlanNo"`
	InnerText string   `xml:",innerxml"`
}

type Interest struct {
	XMLName   xml.Name `xml:"Interest"`
	InnerText string   `xml:",innerxml"`
}

type FineAmt struct {
	XMLName   xml.Name `xml:"FineAmt"`
	InnerText string   `xml:",innerxml"`
}

type OverdueId struct {
	XMLName   xml.Name `xml:"OverdueId"`
	InnerText string   `xml:",innerxml"`
}

type EndMsg struct {
	XMLName   xml.Name `xml:"EndMsg"`
	InnerText string   `xml:",innerxml"`
}

type Amt struct {
	XMLName   xml.Name `xml:"Amt"`
	InnerText string   `xml:",innerxml"`
}

type RepayStatus struct {
	XMLName   xml.Name `xml:"RepayStatus"`
	InnerText string   `xml:",innerxml"`
}

type UnpayNum struct {
	XMLName   xml.Name `xml:"UnpayNum"`
	InnerText string   `xml:",innerxml"`
}

type Overdue struct {
	XMLName   xml.Name `xml:"Overdue"`
	InnerText string   `xml:",innerxml"`
}

type RepayId struct {
	XMLName   xml.Name `xml:"RepayId"`
	InnerText string   `xml:",innerxml"`
}

type Rate struct {
	XMLName   xml.Name `xml:"Rate"`
	InnerText string   `xml:",innerxml"`
}

type ChainCodeID struct {
	XMLName   xml.Name `xml:"ChainCodeID"`
	InnerText string   `xml:",innerxml"`
}

type ChainCodePath struct {
	XMLName   xml.Name `xml:"ChainCodePath"`
	InnerText string   `xml:",innerxml"`
}

type ChainCodeVersion struct {
	XMLName   xml.Name `xml:"ChainCodeVersion"`
	InnerText string   `xml:",innerxml"`
}

type PriKey struct {
	XMLName   xml.Name `xml:"PriKey"`
	InnerText string   `xml:",innerxml"`
}
type PubKey struct {
	XMLName   xml.Name `xml:"PubKey"`
	InnerText string   `xml:",innerxml"`
}
type KeySize struct {
	XMLName   xml.Name `xml:"KeySize"`
	InnerText string   `xml:",innerxml"`
}
type NodeId struct {
	XMLName   xml.Name `xml:"NodeId"`
	InnerText string   `xml:",innerxml"`
}

type BlockNum struct {
	XMLName   xml.Name `xml:"BlockNum"`
	InnerText string   `xml:",innerxml"`
}

type ChainCodeStatus struct {
	XMLName   xml.Name `xml:"ChainCodeStatus"`
	InnerText string   `xml:",innerxml"`
}

type IsBlackList struct {
	XMLName   xml.Name `xml:"IsBlackList"`
	InnerText string   `xml:",innerxml"`
}

type TxId struct {
	XMLName   xml.Name `xml:"TxId"`
	InnerText string   `xml:",innerxml"`
}

type RepayDate struct {
	XMLName   xml.Name `xml:"RepayDate"`
	InnerText string   `xml:",innerxml"`
}

type SendDate struct {
	XMLName   xml.Name `xml:"SendDate"`
	InnerText string   `xml:",innerxml"`
}

type SignData struct {
	XMLName   xml.Name `xml:"SignData"`
	InnerText string   `xml:",innerxml"`
}

type BlockHeight struct {
	XMLName   xml.Name `xml:"BlockHeight"`
	InnerText string   `xml:",innerxml"`
}

type QueryType struct {
	XMLName   xml.Name `xml:"QueryType"`
	InnerText string   `xml:",innerxml"`
}

func WriteXML0004(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0004 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML0004(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0004
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanInfoMap := make(map[string]interface{})
	LoanInfoMap["LenderName"] = result.MsgMsgCont.LoanInfo.LenderName.InnerText
	LoanInfoMap["IdType"] = result.MsgMsgCont.LoanInfo.IdType.InnerText
	LoanInfoMap["IdNo"] = result.MsgMsgCont.LoanInfo.IdNo.InnerText
	LoanInfoMap["LoanNo"] = result.MsgMsgCont.LoanInfo.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanInfo.LoanNo.InnerText
	LoanInfoMap["LoanStatus"] = result.MsgMsgCont.LoanInfo.LoanStatus.InnerText
	LoanInfoMap["LoanAmt"] = result.MsgMsgCont.LoanInfo.LoanAmt.InnerText
	//	LoanInfoMap["RepayType"] = result.MsgMsgCont.LoanInfo.RepayType.InnerText
	//	LoanInfoMap["RepayPeriod"] = result.MsgMsgCont.LoanInfo.RepayPeriod.InnerText
	//	LoanInfoMap["RepayAmt"] = result.MsgMsgCont.LoanInfo.RepayAmt.InnerText
	resultMap["LoanInfo"] = LoanInfoMap
	return resultMap, err
}

func WriteXML0006(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0006 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0006(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0006
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanInfoMap := make(map[string]interface{})
	LoanInfoMap["LenderName"] = result.MsgMsgCont.LoanInfoRef.LenderName.InnerText
	LoanInfoMap["IdType"] = result.MsgMsgCont.LoanInfoRef.IdType.InnerText
	LoanInfoMap["IdNo"] = result.MsgMsgCont.LoanInfoRef.IdNo.InnerText
	LoanInfoMap["LoanNo"] = result.MsgMsgCont.LoanInfoRef.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanInfoRef.LoanNo.InnerText
	LoanInfoMap["LoanStatus"] = result.MsgMsgCont.LoanInfoRef.LoanStatus.InnerText

	log.Println(GetGID(), "The current LoanType is ", result.MsgMsgCont.LoanInfoRef.LoanType.InnerText)
	if result.MsgMsgCont.LoanInfoRef.LoanType.InnerText == "3" {
		log.Println(GetGID(), "The current LoanType is 3")

		LoanInfoMap["LoanType"] = "0"
	} else {
		LoanInfoMap["LoanType"] = result.MsgMsgCont.LoanInfoRef.LoanType.InnerText
	}

	LoanInfoMap["RejMsg"] = result.MsgMsgCont.LoanInfoRef.RejMsg.InnerText
	LoanInfoMap["LoanDesc"] = result.MsgMsgCont.LoanInfoRef.LoanDesc.InnerText
	resultMap["LoanInfo"] = LoanInfoMap
	return resultMap, err
}

func WriteXML0009(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0009 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0009(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0009
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanInfoMap := make(map[string]interface{})
	LoanInfoMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText
	LoanInfoMap["IdType"] = result.MsgMsgCont.LoanBaseInfo.IdType.InnerText
	LoanInfoMap["IdNo"] = result.MsgMsgCont.LoanBaseInfo.IdNo.InnerText

	log.Println(GetGID(), "The current LoanType is ", result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText)
	if result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText == "3" {
		log.Println(GetGID(), "The current LoanType is 3")
		LoanInfoMap["LoanType"] = "0"
	} else {
		LoanInfoMap["LoanType"] = result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText
	}
	resultMap["LoanBaseInfo"] = LoanInfoMap
	return resultMap, err
}

func ReadXML0013(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0013
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["TxId"] = result.MsgMsgCont.TxId.InnerText
	return resultMap, err
}

func WriteXML0013(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0013 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML0014(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0014
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	OrgInfoMap := make(map[string]interface{})
	OrgInfoMap["IdType"] = result.MsgMsgCont.OrgInfo.IdType.InnerText
	OrgInfoMap["IdNo"] = result.MsgMsgCont.OrgInfo.IdNo.InnerText
	resultMap["OrgInfo"] = OrgInfoMap

	return resultMap, err
}
func WriteXML0014(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0013 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func WriteXML0003(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0003 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0003(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0003
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanBaseInfoMap := make(map[string]interface{})
	LoanBaseInfoMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo0003req.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo0003req.LoanNo.InnerText
	LoanBaseInfoMap["Rate"] = result.MsgMsgCont.LoanBaseInfo0003req.Rate.InnerText
	LoanBaseInfoMap["LoanAmt"] = result.MsgMsgCont.LoanBaseInfo0003req.LoanAmt.InnerText
	//新增放款编号
	LoanBaseInfoMap["ReleaseNo"] = result.MsgMsgCont.LoanBaseInfo0003req.ReleaseNo.InnerText
	//新增三字段
	LoanBaseInfoMap["RepayType"] = result.MsgMsgCont.LoanBaseInfo0003req.RepayType.InnerText
	LoanBaseInfoMap["RepayPeriod"] = result.MsgMsgCont.LoanBaseInfo0003req.RepayPeriod.InnerText
	LoanBaseInfoMap["RepayAmt"] = result.MsgMsgCont.LoanBaseInfo0003req.RepayAmt.InnerText

	resultMap["LoanBaseInfo"] = LoanBaseInfoMap
	return resultMap, err
}

func WriteXML0005(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0005 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0005(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0005
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanBaseInfoMap := make(map[string]interface{})
	LoanBaseInfoMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo0005req.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo0005req.LoanNo.InnerText
	LoanBaseInfoMap["RepayPeriod"] = result.MsgMsgCont.LoanBaseInfo0005req.RepayPeriod.InnerText
	LoanBaseInfoMap["ReleaseNo"] = result.MsgMsgCont.LoanBaseInfo0005req.ReleaseNo.InnerText
	resultMap["LoanBaseInfo"] = LoanBaseInfoMap

	return resultMap, err
}
func WriteXML0008(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0008 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0008(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0008
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanInfoMap := make(map[string]interface{})
	LoanInfoMap["LoanNo"] = result.MsgMsgCont.LoanInfo0008.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanInfo0008.LoanNo.InnerText
	LoanInfoMap["IdType"] = result.MsgMsgCont.LoanInfo0008.IdType.InnerText
	LoanInfoMap["IdNo"] = result.MsgMsgCont.LoanInfo0008.IdNo.InnerText
	LoanInfoMap["Amt"] = result.MsgMsgCont.LoanInfo0008.Amt.InnerText
	LoanInfoMap["ReleaseNo"] = result.MsgMsgCont.LoanInfo0008.ReleaseNo.InnerText
	resultMap["LoanInfo"] = LoanInfoMap
	return resultMap, err
}

func WriteXML0010(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0010 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}
func ReadXML0010(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0010
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LenderInfoMap := make(map[string]interface{})
	//LenderInfoMap["LoanId"] = result.MsgMsgCont.LenderInfo0010.LoanId.InnerText
	LenderInfoMap["IdType"] = result.MsgMsgCont.LenderInfo0010.IdType.InnerText
	LenderInfoMap["IdNo"] = result.MsgMsgCont.LenderInfo0010.IdNo.InnerText
	resultMap["LenderInfo"] = LenderInfoMap

	RepayMap := make(map[string]interface{})
	RepayMap["Amt"] = result.MsgMsgCont.Repay0010.Amt.InnerText
	RepayMap["LoanNo"] = result.MsgMsgCont.Repay0010.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.Repay0010.LoanNo.InnerText

	RepayMap["RepayPlanNo"] = result.MsgMsgCont.Repay0010.RepayPlanNo.InnerText
	RepayMap["RepayAmt"] = result.MsgMsgCont.Repay0010.RepayAmt.InnerText
	RepayMap["Interest"] = result.MsgMsgCont.Repay0010.Interest.InnerText
	RepayMap["FineAmt"] = result.MsgMsgCont.Repay0010.FineAmt.InnerText
	RepayMap["Amt"] = result.MsgMsgCont.Repay0010.Amt.InnerText
	RepayMap["RepayStatus"] = result.MsgMsgCont.Repay0010.RepayStatus.InnerText //本笔贷款是否已还清的的标志
	RepayMap["RepayDate"] = result.MsgMsgCont.Repay0010.RepayDate.InnerText
	RepayMap["ReleaseNo"] = result.MsgMsgCont.Repay0010.ReleaseNo.InnerText
	resultMap["Repay"] = RepayMap
	return resultMap, err
}

func WriteXML0007(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0007 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXML0007(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0007
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanBaseInfoMap := make(map[string]interface{})
	LoanBaseInfoMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText
	LoanBaseInfoMap["EndType"] = result.MsgMsgCont.LoanBaseInfo.EndType.InnerText
	LoanBaseInfoMap["EndMsg"] = result.MsgMsgCont.LoanBaseInfo.EndMsg.InnerText
	resultMap["LoanBaseInfo"] = LoanBaseInfoMap
	return resultMap, err
}

func WriteXML0011(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0011 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXML0011(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0011
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["IsEndMonth"] = result.MsgMsgCont.IsEndMonth.InnerText
	LoanInfoMap := make(map[string]interface{})
	LoanInfoMap["LoanNo"] = result.MsgMsgCont.LoanInfo.LoanId.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanInfo.LoanId.InnerText
	LoanInfoMap["IdType"] = result.MsgMsgCont.LoanInfo.IdType.InnerText
	LoanInfoMap["IdNo"] = result.MsgMsgCont.LoanInfo.IdNo.InnerText
	resultMap["LoanInfo"] = LoanInfoMap

	OverdueInfoMap := make(map[string]interface{})
	OverdueInfoMap["LoanNo"] = result.MsgMsgCont.OverdueInfo.LoanId.InnerText
	OverdueInfoMap["RepayPlanNo"] = result.MsgMsgCont.OverdueInfo.RepayPlanNo.InnerText
	OverdueInfoMap["RepayAmt"] = result.MsgMsgCont.OverdueInfo.RepayAmt.InnerText
	OverdueInfoMap["Interest"] = result.MsgMsgCont.OverdueInfo.Interest.InnerText
	OverdueInfoMap["FineAmt"] = result.MsgMsgCont.OverdueInfo.FineAmt.InnerText
	OverdueInfoMap["OverdueDay"] = result.MsgMsgCont.OverdueInfo.OverdueDay.InnerText
	OverdueInfoMap["SendDate"] = result.MsgMsgCont.OverdueInfo.SendDate.InnerText
	OverdueInfoMap["ReleaseNo"] = result.MsgMsgCont.OverdueInfo.ReleaseNo.InnerText

	resultMap["OverdueInfo"] = OverdueInfoMap
	return resultMap, err
}

func WriteXML0012(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0012 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXML0012(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0012
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	OrgInfoMap := make(map[string]interface{})
	OrgInfoMap["OrgName"] = result.MsgMsgCont.OrgInfo.OrgName.InnerText
	OrgInfoMap["IdType"] = result.MsgMsgCont.OrgInfo.IdType.InnerText
	OrgInfoMap["IdNo"] = result.MsgMsgCont.OrgInfo.IdNo.InnerText
	OrgInfoMap["DigitalCurrency"] = result.MsgMsgCont.OrgInfo.DigitalCurrency.InnerText
	OrgInfoMap["LoanTotal"] = result.MsgMsgCont.OrgInfo.LoanTotal.InnerText
	OrgInfoMap["LoanNum"] = result.MsgMsgCont.OrgInfo.LoanNum.InnerText
	OrgInfoMap["Describe"] = result.MsgMsgCont.OrgInfo.Describe.InnerText
	OrgInfoMap["OperateState"] = result.MsgMsgCont.OrgInfo.OperateState.InnerText
	OrgInfoMap["CancelReason"] = result.MsgMsgCont.OrgInfo.CancelReason.InnerText
	//新增数字货币字段
	OrgInfoMap["InitialReward"] = result.MsgMsgCont.OrgInfo.InitialReward.InnerText
	OrgInfoMap["FollowUpReward"] = result.MsgMsgCont.OrgInfo.FollowUpReward.InnerText
	OrgInfoMap["AmountConsume"] = result.MsgMsgCont.OrgInfo.AmountConsume.InnerText
	OrgInfoMap["InitialRewardCon"] = result.MsgMsgCont.OrgInfo.InitialRewardCon.InnerText
	OrgInfoMap["FollowUpRewardCon"] = result.MsgMsgCont.OrgInfo.FollowUpRewardCon.InnerText
	OrgInfoMap["AmountConsumeCon"] = result.MsgMsgCont.OrgInfo.AmountConsumeCon.InnerText
	resultMap["OrgInfo"] = OrgInfoMap
	log.Println(GetGID(), "0012 OrgInfo------", resultMap)
	return resultMap, err
}

func WriteXML0001(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0001 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXML0001(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0001
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	LoanBaseInfoMap := make(map[string]interface{})
	LoanBaseInfoMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText
	resultMap["LoanNo"] = result.MsgMsgCont.LoanBaseInfo.LoanNo.InnerText

	log.Println(GetGID(), "The current LoanType is ", result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText)

	if result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText == "3" {
		log.Println(GetGID(), "The current LoanType is 3")

		LoanBaseInfoMap["LoanType"] = "0"
	} else {
		LoanBaseInfoMap["LoanType"] = result.MsgMsgCont.LoanBaseInfo.LoanType.InnerText
	}
	LoanBaseInfoMap["ApplyAmt"] = result.MsgMsgCont.LoanBaseInfo.ApplyAmt.InnerText
	LoanBaseInfoMap["LoanDesc"] = result.MsgMsgCont.LoanBaseInfo.LoanDesc.InnerText
	LoanBaseInfoMap["ApplyTime"] = result.MsgMsgCont.LoanBaseInfo.ApplyTime.InnerText
	resultMap["LoanBaseInfo"] = LoanBaseInfoMap

	LenderMap := make(map[string]interface{})
	LenderMap["LenderType"] = result.MsgMsgCont.LenderInfo.LenderType.InnerText
	LenderMap["LenderName"] = result.MsgMsgCont.LenderInfo.LenderName.InnerText
	LenderMap["IdType"] = result.MsgMsgCont.LenderInfo.IdType.InnerText
	LenderMap["IdNo"] = result.MsgMsgCont.LenderInfo.IdNo.InnerText
	LenderMap["IsBlackList"] = result.MsgMsgCont.LenderInfo.IsBlackList.InnerText
	resultMap["LenderInfo"] = LenderMap

	DLenderMap := make(map[string]interface{})
	DLenderMap["DLenderName"] = result.MsgMsgCont.DLenderInfo.DLenderName.InnerText
	DLenderMap["DIdType"] = result.MsgMsgCont.DLenderInfo.DIdType.InnerText
	DLenderMap["DIdNo"] = result.MsgMsgCont.DLenderInfo.DIdNo.InnerText
	resultMap["DLenderInfo"] = DLenderMap

	if len(result.MsgMsgCont.WarrantorInfoLst) > 0 {
		WarrantorInfoLst := []map[string]string{}

		for i := 0; i < len(result.MsgMsgCont.WarrantorInfoLst); i++ {
			WarrantorInfoMap := make(map[string]string)

			WarrantorInfoMap["WarrantorName"] = result.MsgMsgCont.WarrantorInfoLst[i].WarrantorName.InnerText
			WarrantorInfoMap["WIdType"] = result.MsgMsgCont.WarrantorInfoLst[i].WIdType.InnerText
			WarrantorInfoMap["WIdNo"] = result.MsgMsgCont.WarrantorInfoLst[i].WIdNo.InnerText
			WarrantorInfoLst = append(WarrantorInfoLst, WarrantorInfoMap)

		}
		resultMap["WarrantorInfo"] = WarrantorInfoLst
	}
	if len(result.MsgMsgCont.PledgeInfoLst) > 0 {
		PledgeInfoLst := []map[string]string{}
		for j := 0; j < len(result.MsgMsgCont.PledgeInfoLst); j++ {
			PledgeInfoMap := make(map[string]string)
			PledgeInfoMap["PledgeType"] = result.MsgMsgCont.PledgeInfoLst[j].PledgeType.InnerText
			PledgeInfoMap["PledgeNo"] = result.MsgMsgCont.PledgeInfoLst[j].PledgeNo.InnerText
			PledgeInfoMap["PledgeDesc"] = result.MsgMsgCont.PledgeInfoLst[j].PledgeDesc.InnerText
			PledgeInfoLst = append(PledgeInfoLst, PledgeInfoMap)
		}
		resultMap["PledgeInfo"] = PledgeInfoLst

	}

	return resultMap, err
}

func WriteXML0015(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0015 inputMap------", inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML0015(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0015
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText
	//MsgMsgCont
	resultMap["Flag"] = result.MsgMsgCont.Flag.InnerText
	resultMap["TimeStamp"] = result.MsgMsgCont.TimeStamp.InnerText
	resultMap["TSSignature"] = result.MsgMsgCont.TSSignature.InnerText

	return resultMap, err
}

//func WriteXML0002(inputMap map[string]interface{}) (string, error) {
//	var msg, loanType, fileName string
//	var err error

//	inputMap = transferMap(inputMap)
//	LoanInfoMap := inputMap["LoanInfo"].(map[string]string)
//	loanType = LoanInfoMap["LoanType"]
//	if inputMap["LoanType"] == nil {
//		log.Println(GetGID(),"LoanType is nil")
//	}
//	if strings.EqualFold(loanType, "0") {
//		//<PledgeInfo><PledgeType>1</PledgeType><PledgeNo>1</PledgeNo><PledgeDesc>1</PledgeDesc><PledgeStatus>1</PledgeStatus><ApplyComId>1</ApplyComId><CrtTime>1</CrtTime><PledgeHistory><Pledge><ApplyTime>1</ApplyTime><ApplyComId>1</ApplyComId><IdType>1</IdType><IdNoS>1</IdNoS><LoanStatus>1</LoanStatus><LoanTime>1</LoanTime><ApplyAmt>1</ApplyAmt><LoanAmt>1</LoanAmt><UnrepayAmt>1</UnrepayAmt><OverdueNum>1</OverdueNum><OverdueDay>1</OverdueDay></Pledge><Pledge><ApplyTime>1</ApplyTime><ApplyComId>1</ApplyComId><IdType>1</IdType><IdNoS>1</IdNoS><LoanStatus>1</LoanStatus><LoanTime>1</LoanTime><ApplyAmt>1</ApplyAmt><LoanAmt>1</LoanAmt><UnrepayAmt>1</UnrepayAmt><OverdueNum>1</OverdueNum><OverdueDay>1</OverdueDay></Pledge></PledgeHistory></PledgeInfo>
//		fileName = "Rsp0002_0.xml"
//		log.Println(GetGID(),"fileName")
//		content, err := ioutil.ReadFile(GetPath() + "/" + fileName)
//		log.Println(GetGID(),content)

//		// 动态生成抵押物历史xml结构
//		s := `<?xml version='1.0' encoding='UTF-8'?><Msg>	<MsgHead>		<MsgVer>1</MsgVer>		<MsgId>0002</MsgId>		<MsgType>1</MsgType>		<ProcID></ProcID>		<ProcTime></ProcTime>		<Sender></Sender>		<SenderID></SenderID>		<Receiver></Receiver>		<ReceiverID></ReceiverID>		<RspCode></RspCode>		<ErrMsg></ErrMsg>	</MsgHead>	<MsgCont>		<LoanInfo>			<OrgId>1</OrgId>			<LoanType>1</LoanType>			<PledgeType>1</PledgeType>			<PledgeNo>1</PledgeNo>			<WarrantorIdType>1</WarrantorIdType>			<WarrantorIdNo>1</WarrantorIdNo>			<LoanStatus>1</LoanStatus>			<LoanTime>1</LoanTime>			<ApplyAmt>1</ApplyAmt>			<LoanAmt>1</LoanAmt>			<UnrepayAmt>1</UnrepayAmt>			<OverdueNum>1</OverdueNum>			<OverdueDay>1</OverdueDay>								</LoanInfo>		<ApplyInfo>					<LenderName>1</LenderName>																									<IdType>1</IdType>							<IdNoS>1</IdNoS>																									<LenderType>1</LenderType>							<IsBlacklist>1</IsBlacklist>																									<OverdueCntNum>1</OverdueCntNum>							<UnpayCntNum>1</UnpayCntNum>																									<OverdueCntDay>1</OverdueCntDay>							<CrtUser>1</CrtUser>																									<CrtTime>1</CrtTime>							</ApplyInfo>	</MsgCont></Msg>`
//		el, err := LoadByXml(s)
//		if err != nil {
//			log.Println(GetGID(),"err", err)
//			return "", err
//		}
//		g := el.Node("MsgCont")
//		g.AddNode(NewElement("WarrantorInfo", "<WarrantorName>1</WarrantorName><IdType>1</IdType><IdNoS>1</IdNoS><WarrantorStatus>1</WarrantorStatus><ApplyComId>1</ApplyComId><CrtTime>1</CrtTime><WarrantorId>1</WarrantorId>"))
//		log.Println(GetGID(),el.SyncToXml())

//		WarrantorInfo := g.Node("WarrantorInfo")
//		WarrantorInfo.AddNode(NewElement("WarrantorHistory", ""))
//		WarrantorHistory := WarrantorInfo.Node("WarrantorHistory")
//		WarrantorHistory.AddNode(NewElement("Warrantor", "<ApplyTime>1</ApplyTime>					<ApplyComId>1</ApplyComId>						<IdType>1</IdType>						<IdNoS>1</IdNoS>						<LoanStatus>1</LoanStatus>						<LoanTime>1</LoanTime>						<ApplyAmt>1</ApplyAmt>						<LoanAmt>1</LoanAmt>						<UnrepayAmt>1</UnrepayAmt>						<OverdueNum>1</OverdueNum>						<OverdueDay>1</OverdueDay>	"))
//		WarrantorHistory.AddNode(NewElement("Warrantor", "<ApplyTime>1</ApplyTime>					<ApplyComId>1</ApplyComId>						<IdType>1</IdType>						<IdNoS>1</IdNoS>						<LoanStatus>1</LoanStatus>						<LoanTime>1</LoanTime>						<ApplyAmt>1</ApplyAmt>						<LoanAmt>1</LoanAmt>						<UnrepayAmt>1</UnrepayAmt>						<OverdueNum>1</OverdueNum>						<OverdueDay>1</OverdueDay>	"))

//		log.Println(GetGID(),el.SyncToXml())

//	} else if strings.EqualFold(loanType, "1") {
//		fileName = "Rsp0002_1.xml"
//	} else if strings.EqualFold(loanType, "2") {
//		//<WarrantorInfo><WarrantorName>1</WarrantorName><IdType>1</IdType><IdNoS>1</IdNoS><WarrantorStatus>1</WarrantorStatus><ApplyComId>1</ApplyComId><CrtTime>1</CrtTime><WarrantorId>1</WarrantorId><WarrantorHistory><Warrantor><ApplyTime>1</ApplyTime><ApplyComId>1</ApplyComId><IdType>1</IdType><IdNoS>1</IdNoS><LoanStatus>1</LoanStatus><LoanTime>1</LoanTime><ApplyAmt>1</ApplyAmt><LoanAmt>1</LoanAmt><UnrepayAmt>1</UnrepayAmt><OverdueNum>1</OverdueNum><OverdueDay>1</OverdueDay></Warrantor><Warrantor><ApplyTime>1</ApplyTime><ApplyComId>1</ApplyComId><IdType>1</IdType><IdNoS>1</IdNoS><LoanStatus>1</LoanStatus><LoanTime>1</LoanTime><ApplyAmt>1</ApplyAmt><LoanAmt>1</LoanAmt><UnrepayAmt>1</UnrepayAmt><OverdueNum>1</OverdueNum><OverdueDay>1</OverdueDay></Warrantor></WarrantorHistory></WarrantorInfo>
//		fileName = "Rsp0002_2.xml"
//	} else {
//		return "", err
//	}
//	content, err := ioutil.ReadFile(GetPath() + "/" + fileName)
//	if err != nil {
//		return "", err
//	}
//	var result StringMsgRsp0002
//	err = xml.Unmarshal(content, &result)
//	//	if err != nil {
//	//		return "", err
//	//	}
//	log.Println(GetGID(),result)
//	log.Println(GetGID(),result.MsgMsgHead)
//	//MsgHead部分处理
//	result.MsgMsgHead.MsgVer.InnerText = "1"
//	result.MsgMsgHead.MsgId.InnerText = inputMap["MsgId"].(string)
//	result.MsgMsgHead.MsgType.InnerText = inputMap["MsgType"].(string)
//	result.MsgMsgHead.ProcId.InnerText = inputMap["ProcId"].(string)
//	result.MsgMsgHead.ProcTime.InnerText = time.Now().Format("20060102150405")
//	result.MsgMsgHead.Sender.InnerText = inputMap["Sender"].(string)
//	result.MsgMsgHead.SenderId.InnerText = inputMap["SenderId"].(string)

//	result.MsgMsgHead.Receiver.InnerText = inputMap["Receiver"].(string)
//	result.MsgMsgHead.ReceiverId.InnerText = inputMap["ReceiverId"].(string)
//	result.MsgMsgHead.RspCode.InnerText = inputMap["RspCode"].(string)
//	result.MsgMsgHead.ErrMsg.InnerText = inputMap["ErrMsg"].(string)

//	// 以下贷款信息的获取具体看业务适配器返回值TODO
//	loanInfoMap := make(map[string]string)
//	loanInfoMap = inputMap["LoanInfo"].(map[string]string)
//	result.MsgMsgCont.LoanInfo.OrgId.InnerText = loanInfoMap["OrgId"]
//	result.MsgMsgCont.LoanInfo.LoanType.InnerText = loanInfoMap["LoanType"]
//	result.MsgMsgCont.LoanInfo.PledgeType.InnerText = loanInfoMap["PledgeType"]
//	result.MsgMsgCont.LoanInfo.PledgeNo.InnerText = loanInfoMap["PledgeNo"]
//	result.MsgMsgCont.LoanInfo.WarrantorIdType.InnerText = loanInfoMap["WarrantorIdType"]
//	result.MsgMsgCont.LoanInfo.WarrantorIdNo.InnerText = loanInfoMap["WarrantorIdNo"]
//	result.MsgMsgCont.LoanInfo.LoanStatus.InnerText = loanInfoMap["LoanStatus"]
//	result.MsgMsgCont.LoanInfo.LoanTime.InnerText = loanInfoMap["LoanTime"]
//	result.MsgMsgCont.LoanInfo.ApplyAmt.InnerText = loanInfoMap["ApplyAmt"]
//	result.MsgMsgCont.LoanInfo.LoanAmt.InnerText = loanInfoMap["LoanAmt"]
//	result.MsgMsgCont.LoanInfo.UnrepayAmt.InnerText = loanInfoMap["UnrepayAmt"]
//	result.MsgMsgCont.LoanInfo.OverdueNum.InnerText = loanInfoMap["OverdueNum"]
//	result.MsgMsgCont.LoanInfo.OverdueDay.InnerText = loanInfoMap["OverdueDay"]
//	if inputMap["PledgeInfo"] != nil {
//		// 如何获取TODO
//		PledgeInfoLst := inputMap["PledgeInfo"].([]map[string]interface{})
//		log.Println(GetGID(),"len(PledgeInfoLst)=", len(PledgeInfoLst))

//		for i := 0; i < len(PledgeInfoLst); i++ {
//			log.Println(GetGID(),"i=", i)

//			result.MsgMsgCont.PledgeInfoLst[i].PledgeType.InnerText = PledgeInfoLst[i]["PledgeType"].(string)
//			result.MsgMsgCont.PledgeInfoLst[i].PledgeNo.InnerText = PledgeInfoLst[i]["PledgeNo"].(string)
//			result.MsgMsgCont.PledgeInfoLst[i].PledgeDesc.InnerText = PledgeInfoLst[i]["PledgeDesc"].(string)
//			result.MsgMsgCont.PledgeInfoLst[i].PledgeStatus.InnerText = PledgeInfoLst[i]["PledgeStatus"].(string)
//			result.MsgMsgCont.PledgeInfoLst[i].ApplyComId.InnerText = PledgeInfoLst[i]["ApplyComId"].(string)
//			result.MsgMsgCont.PledgeInfoLst[i].CrtTime.InnerText = PledgeInfoLst[i]["CrtTime"].(string)

//			var PledgeHistoryLst []map[string]string
//			PledgeHistoryLst = PledgeInfoLst[i]["PledgeHistory"].([]map[string]string)

//			for j := 0; j < len(PledgeHistoryLst); j++ {
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].ApplyTime = PledgeHistoryLst[j]["ApplyTime"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].ApplyComId = PledgeHistoryLst[j]["ApplyComId"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].IdType = PledgeHistoryLst[j]["IdType"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].IdNoS = PledgeHistoryLst[j]["IdNoS"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].LoanStatus = PledgeHistoryLst[j]["LoanStatus"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].LoanTime = PledgeHistoryLst[j]["LoanTime"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].ApplyAmt = PledgeHistoryLst[j]["ApplyAmt"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].LoanAmt = PledgeHistoryLst[j]["LoanAmt"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].UnrepayAmt = PledgeHistoryLst[j]["UnrepayAmt"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].OverdueNum = PledgeHistoryLst[j]["OverdueNum"]
//				result.MsgMsgCont.PledgeInfoLst[i].PledgeHistory.PledgeLst[j].OverdueDay = PledgeHistoryLst[j]["OverdueDay"]
//			}
//		}
//	}

//	if inputMap["WarrantorInfo"] != nil {
//		var WarrantorLst []map[string]interface{}
//		WarrantorLst = inputMap["WarrantorInfo"].([]map[string]interface{})
//		for i := 0; i < len(WarrantorLst); i++ {
//			result.MsgMsgCont.WarrantorInfoLst[i].WarrantorName.InnerText = WarrantorLst[i]["WarrantorName"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].IdType.InnerText = WarrantorLst[i]["IdType"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].IdNoS.InnerText = WarrantorLst[i]["IdNoS"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].WarrantorStatus.InnerText = WarrantorLst[i]["WarrantorStatus"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].ApplyComId.InnerText = WarrantorLst[i]["ApplyComId"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].CrtTime.InnerText = WarrantorLst[i]["CrtTime"].(string)
//			result.MsgMsgCont.WarrantorInfoLst[i].WarrantorId.InnerText = WarrantorLst[i]["WarrantorId"].(string)

//			var WarrantorHistoryLst []map[string]string
//			WarrantorHistoryLst = WarrantorLst[i]["WarrantorHistory"].([]map[string]string)

//			for j := 0; j < len(WarrantorHistoryLst); j++ {
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].ApplyTime = WarrantorHistoryLst[j]["ApplyTime"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].ApplyComId = WarrantorHistoryLst[j]["ApplyComId"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].IdType = WarrantorHistoryLst[j]["IdType"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].IdNoS = WarrantorHistoryLst[j]["IdNoS"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].LoanStatus = WarrantorHistoryLst[j]["LoanStatus"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].LoanTime = WarrantorHistoryLst[j]["LoanTime"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].ApplyAmt = WarrantorHistoryLst[j]["ApplyAmt"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].LoanAmt = WarrantorHistoryLst[j]["LoanAmt"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].UnrepayAmt = WarrantorHistoryLst[j]["UnrepayAmt"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].OverdueNum = WarrantorHistoryLst[j]["OverdueNum"]
//				result.MsgMsgCont.WarrantorInfoLst[i].WarrantorHistory.WarrantorLst[j].OverdueDay = WarrantorHistoryLst[j]["OverdueDay"]
//			}

//		}
//	}

//	//保存修改后的内容
//	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "")
//	if outPutErr == nil {
//		//加入XML头
//		headerBytes := []byte(strings.Replace(xml.Header, "\n", "", -1))
//		//拼接XML头和实际XML内容
//		xmlOutPutData := append(headerBytes, xmlOutPut...)
//		msg = string(xmlOutPutData)
//		log.Println(GetGID(),"the result msg:", msg)

//		//写入文件
//		//		ioutil.WriteFile("studygolang_test.xml", xmlOutPutDataNew, os.ModeAppend)
//		log.Println(GetGID(),"OK~")
//		//		return string(xmlOutPutDataNew)
//	} else {
//		log.Println(GetGID(),outPutErr)
//		//		return ""
//	}
//	return msg, err
//}

func WriteXML0002(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "0002 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

// 以下方法是把request对应的xml字符串解析成map
func ReadXML0002(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq0002
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont

	QueryInfoMap := make(map[string]interface{})

	QueryInfoMap["QueryType"] = result.MsgMsgCont.QueryInfo.QueryType.InnerText
	QueryInfoMap["LoanNo"] = result.MsgMsgCont.QueryInfo.LoanNo.InnerText
	//	resultMap["LoanNo"] = result.MsgMsgCont.QueryInfo.LoanNo.InnerText

	QueryInfoMap["IdType"] = result.MsgMsgCont.QueryInfo.IdType.InnerText
	QueryInfoMap["IdNo"] = result.MsgMsgCont.QueryInfo.IdNo.InnerText
	QueryInfoMap["DIdType"] = result.MsgMsgCont.QueryInfo.DIdType.InnerText
	QueryInfoMap["DIdNo"] = result.MsgMsgCont.QueryInfo.DIdNo.InnerText
	QueryInfoMap["WIdType"] = result.MsgMsgCont.QueryInfo.WIdType.InnerText
	QueryInfoMap["WIdNo"] = result.MsgMsgCont.QueryInfo.WIdNo.InnerText
	QueryInfoMap["PledgeNo"] = result.MsgMsgCont.QueryInfo.PledgeNo.InnerText
	QueryInfoMap["PledgeType"] = result.MsgMsgCont.QueryInfo.PledgeType.InnerText
	resultMap["QueryInfo"] = QueryInfoMap
	return resultMap, err
}

func WriteXML1001(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "1001 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func WriteXML1002(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "1002 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func WriteXML1003(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "1003 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML2001(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq2001
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["ChainCodeID"] = result.MsgMsgCont.ChainCodeID.InnerText
	resultMap["ChainCodePath"] = result.MsgMsgCont.ChainCodePath.InnerText
	resultMap["ChainCodeVersion"] = result.MsgMsgCont.ChainCodeVersion.InnerText
	return resultMap, err
}

func WriteXML2001(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "2001 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML2002(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq2002
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["ChainCodeID"] = result.MsgMsgCont.ChainCodeID.InnerText
	resultMap["ChainCodePath"] = result.MsgMsgCont.ChainCodePath.InnerText
	resultMap["ChainCodeVersion"] = result.MsgMsgCont.ChainCodeVersion.InnerText
	return resultMap, err
}

func WriteXML2002(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "2002 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func WriteXML2003(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "2003 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML2003(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq2003
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["ChainCodeID"] = result.MsgMsgCont.ChainCodeID.InnerText
	resultMap["ChainCodePath"] = result.MsgMsgCont.ChainCodePath.InnerText
	resultMap["ChainCodeVersion"] = result.MsgMsgCont.ChainCodeVersion.InnerText
	return resultMap, err
}

func ReadXML2004(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq2004
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["PriKey"] = result.MsgMsgCont.PriKey.InnerText
	resultMap["PubKey"] = result.MsgMsgCont.PubKey.InnerText
	resultMap["KeySize"] = result.MsgMsgCont.KeySize.InnerText
	resultMap["NodeId"] = result.MsgMsgCont.NodeId.InnerText
	return resultMap, err
}
func WriteXML2004(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "2004 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML2005(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq2005
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	//	resultMap["ChainCodeID"] = result.MsgMsgCont.ChainCodeID.InnerText
	//	resultMap["ChainCodePath"] = result.MsgMsgCont.ChainCodePath.InnerText
	//	resultMap["ChainCodeVersion"] = result.MsgMsgCont.ChainCodeVersion.InnerText
	return resultMap, err
}
func WriteXML2005(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "2005 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML3001(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq3001
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["BlockHeight"] = result.MsgMsgCont.BlockHeight.InnerText
	return resultMap, err
}

func WriteXML3001(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "3001 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err
}

func ReadXML3002(msg string) (map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	var result StringMsgReq3002
	var content []byte = []byte(msg)
	var err error
	err = xml.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	log.Println(GetGID(), result)
	log.Println(GetGID(), result.MsgMsgHead)
	//MsgHead部分处理
	resultMap["MsgVer"] = result.MsgMsgHead.MsgVer.InnerText
	resultMap["MsgId"] = result.MsgMsgHead.MsgId.InnerText
	resultMap["MsgType"] = result.MsgMsgHead.MsgType.InnerText
	resultMap["ProcId"] = result.MsgMsgHead.ProcId.InnerText
	resultMap["ProcTime"] = result.MsgMsgHead.ProcTime.InnerText
	resultMap["Sender"] = result.MsgMsgHead.Sender.InnerText
	resultMap["SenderId"] = result.MsgMsgHead.SenderId.InnerText
	resultMap["Receiver"] = result.MsgMsgHead.Receiver.InnerText
	resultMap["ReceiverId"] = result.MsgMsgHead.ReceiverId.InnerText
	resultMap["RspCode"] = result.MsgMsgHead.RspCode.InnerText
	resultMap["ErrMsg"] = result.MsgMsgHead.ErrMsg.InnerText
	resultMap["SignData"] = result.MsgMsgHead.SignData.InnerText

	//MsgMsgCont
	resultMap["BlockNum"] = result.MsgMsgCont.BlockNum.InnerText
	return resultMap, err
}

func WriteXML3002(inputMap map[string]interface{}) (string, error) {
	var msg string
	var err error
	log.Println(GetGID(), "3002 inputMap------", inputMap)

	//	inputMap = transferMap(inputMap)
	msg = PackResponseMsg(inputMap)
	return msg, err

	//	var MsgHeadMap = make(map[string]interface{})
	//	MsgHeadMap["MsgVer"] = "MsgVer"
	//	MsgHeadMap["MsgId"] = "MsgId"
	//	MsgHeadMap["MsgType"] = "MsgType"
	//	MsgHeadMap["ProcId"] = "ProcId"
	//	MsgHeadMap["ProcTime"] = "ProcTime"
	//	MsgHeadMap["Sender"] = "Sender"
	//	MsgHeadMap["SenderId"] = "SenderId"
	//	MsgHeadMap["Receiver"] = "Receiver"
	//	MsgHeadMap["ReceiverId"] = "ReceiverId"
	//	MsgHeadMap["RspCode"] = "RspCode"
	//	MsgHeadMap["ErrMsg"] = "ErrMsg"
	//	var buffer bytes.Buffer
	//	var buffer_MsgHead bytes.Buffer
	//	var buffer_MsgCont bytes.Buffer
	//	buffer.WriteString("<?xml version='1.0' encoding='UTF-8'?><Msg>")
	//	buffer_MsgHead.WriteString("<MsgHead>")
	//	buffer_MsgCont.WriteString("<MsgCont>")
	//	//crtUser := inputMap["CrtUser"].(string)
	//	for k, v := range inputMap {
	//		tempMap := make(map[string]interface{})
	//		log.Println(GetGID(),"MsgHeadMap[k]:", MsgHeadMap[k])
	//		log.Println(GetGID(),"v:", v)
	//		if MsgHeadMap[k] != nil {
	//			tempMap[k] = v
	//			buffer_MsgHead.WriteString(analysisMap(tempMap))
	//		}
	//	}

	//	buffer_MsgCont.WriteString("<Block>")

	//	//	var blockData string
	//	//	for _, block := range inputMap["Block"].(map[string]interface{}) {
	//	//		blockData += block
	//	//	}
	//	//	log.Println(GetGID(),"blockData:   ", blockData)
	//	buffer_MsgCont.WriteString(inputMap["Block"])
	//	buffer_MsgCont.WriteString("</Block>")
	//	buffer_MsgCont.WriteString("<BlockNum>")
	//	buffer_MsgCont.WriteString(inputMap["BlockNum"].(string))
	//	buffer_MsgCont.WriteString("</BlockNum>")

	//	buffer_MsgHead.WriteString("</MsgHead>")
	//	buffer_MsgCont.WriteString("</MsgCont>")

	//	buffer.WriteString(string(buffer_MsgHead.Bytes()))
	//	buffer.WriteString(string(buffer_MsgCont.Bytes()))
	//	buffer.WriteString("</Msg>")
	//	msg = string(buffer.Bytes())
	//	log.Println(GetGID(),"3002 msg------", msg)
	//	return msg, err
}

func GetPath() string {
	//dir := "./parse/template"
	dir := "./cfg/template"
	return dir
}

//    personDB := make(map[string][2]PersonInfo)

//    //初始化，注意对数组的初始化
//    personDB["test1"] = [2]PersonInfo{{"12345", "Tom", "aaa"}, {"12346", "Xym", "bbb"}}

func transferMap(inputMap map[string]interface{}) map[string]interface{} {
	log.Println(GetGID(), "########## transferMap start#####inputMap#####", inputMap)
	if inputMap["Receiver"] == nil {
		inputMap["Receiver"] = ""
	}
	if inputMap["ReceiverId"] == nil {
		inputMap["ReceiverId"] = ""
	}
	for key, value := range inputMap {
		//log.Println(GetGID(),"%f=%f", key, value)
		if value == nil {
			inputMap[key] = ""
		}
	}
	log.Println(GetGID(), "########## transferMap end #####inputMap#####", inputMap)
	return inputMap
}

func PackResponseMsg(data map[string]interface{}) string {
	log.Println(GetGID(), data)

	var MsgHeadMap = make(map[string]interface{})
	MsgHeadMap["MsgVer"] = "MsgVer"
	MsgHeadMap["MsgId"] = "MsgId"
	MsgHeadMap["MsgType"] = "MsgType"
	MsgHeadMap["ProcId"] = "ProcId"
	MsgHeadMap["ProcTime"] = "ProcTime"
	MsgHeadMap["Sender"] = "Sender"
	MsgHeadMap["SenderId"] = "SenderId"
	MsgHeadMap["Receiver"] = "Receiver"
	MsgHeadMap["ReceiverId"] = "ReceiverId"
	MsgHeadMap["RspCode"] = "RspCode"
	MsgHeadMap["ErrMsg"] = "ErrMsg"
	var buffer bytes.Buffer
	var buffer_MsgHead bytes.Buffer
	var buffer_MsgCont bytes.Buffer
	buffer.WriteString("<?xml version='1.0' encoding='UTF-8'?><Msg>")
	buffer_MsgHead.WriteString("<MsgHead>")
	buffer_MsgCont.WriteString("<MsgCont>")

	for k, v := range data {
		tempMap := make(map[string]interface{})
		//log.Println(GetGID(),"MsgHeadMap[k]:", MsgHeadMap[k])
		//log.Println(GetGID(),"v:", v)

		if MsgHeadMap[k] != nil {
			tempMap[k] = v
			buffer_MsgHead.WriteString(analysisMap(tempMap))
		} else {
			tempMap[k] = v
			//log.Println(GetGID(),"tempMap:", tempMap)

			buffer_MsgCont.WriteString(analysisMap(tempMap))
		}
	}
	buffer_MsgHead.WriteString("</MsgHead>")
	buffer_MsgCont.WriteString("</MsgCont>")

	buffer.WriteString(string(buffer_MsgHead.Bytes()))
	buffer.WriteString(string(buffer_MsgCont.Bytes()))
	buffer.WriteString("</Msg>")
	return string(buffer.Bytes())
}

func analysisMap(data map[string]interface{}) string {
	var buffer bytes.Buffer
	for k, v := range data {
		//s_type := fmt.Sprint(reflect.TypeOf(v))
		//log.Println(GetGID(),s_type)
		switch v.(type) {
		case string:

			buffer.WriteString("<")
			buffer.WriteString(k)
			buffer.WriteString(">")
			buffer.WriteString(v.(string))
			buffer.WriteString("</")
			buffer.WriteString(k)
			buffer.WriteString(">")
		case map[string]interface{}:
			v_sz, _ := v.(map[string]interface{})
			buffer.WriteString("<")
			buffer.WriteString(k)
			buffer.WriteString(">")
			buffer.WriteString(analysisMap(v_sz))
			buffer.WriteString("</")
			buffer.WriteString(k)
			buffer.WriteString(">")

		case []map[string]interface{}:

			v_sz, _ := v.([]map[string]interface{})

			if k == "LoanInfo" {
				// 贷款查询专用
				for i := 0; i < len(v_sz); i++ {
					buffer.WriteString("<")
					buffer.WriteString(k)
					buffer.WriteString(">")
					log.Println(GetGID(), "--1--")
					vv := v_sz[i]

					buffer.WriteString(analysisMap(vv))
					buffer.WriteString("</")
					buffer.WriteString(k)
					buffer.WriteString(">")
				}

			} else {
				buffer.WriteString("<")
				buffer.WriteString(k)
				buffer.WriteString(">")
				for i := 0; i < len(v_sz); i++ {
					log.Println(GetGID(), "--1--")
					vv := v_sz[i]
					buffer.WriteString(analysisMap(vv))

				}
				buffer.WriteString("</")
				buffer.WriteString(k)
				buffer.WriteString(">")
			}
		}
	}
	return string(buffer.Bytes())
}

//func sortMapByKeys(data map[string]interface{}) string {
//	// get the list of keys and sort them
//	keys := []string{}
//	for key := range colors {
//		keys = append(keys, key)
//	}
//	sort.Strings(keys)

//	for _, val := range keys {
//		log.Println(GetGID(),val, colors[val])
//	}
//}
