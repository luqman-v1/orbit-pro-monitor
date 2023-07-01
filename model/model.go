package model

import (
	"encoding/json"
	"encoding/xml"

	"github.com/sirupsen/logrus"
)

type EngInfo struct {
	XMLName xml.Name `xml:"RGW"`
	Text    string   `xml:",chardata"`
	Reqid   string   `xml:"reqid"`
	Eng     struct {
		Text string `xml:",chardata"`
		Lte  struct {
			Text             string `xml:",chardata"`
			Mcc              string `xml:"mcc"`
			MncLen           string `xml:"mnc_len"`
			Mnc              string `xml:"mnc"`
			Tac              string `xml:"tac"`
			PhyCellID        string `xml:"phy_cell_id"`
			CellID           string `xml:"cell_id"`
			DlEuarfcn        string `xml:"dl_euarfcn"`
			UlEuarfcn        string `xml:"ul_euarfcn"`
			Band             string `xml:"band"`
			DlBandwidth      string `xml:"dl_bandwidth"`
			TransmissionMode string `xml:"transmission_mode"`
			Rsrp             string `xml:"rsrp"`
			Rsrq             string `xml:"rsrq"`
			Sinr             string `xml:"sinr"`
			MainRsrp         string `xml:"main_rsrp"`
			DiversityRsrp    string `xml:"diversity_rsrp"`
			MainRsrq         string `xml:"main_rsrq"`
			DiversityRsrq    string `xml:"diversity_rsrq"`
			Rssi             string `xml:"rssi"`
			Cqi              string `xml:"cqi"`
			DlBler           string `xml:"dl_bler"`
			UlBler           string `xml:"ul_bler"`
			DlThroughput     string `xml:"dl_throughput"`
			DlPeakThroughput string `xml:"dl_peak_throughput"`
			UlThroughput     string `xml:"ul_throughput"`
			UlPeakThroughput string `xml:"ul_peak_throughput"`
			RankIndicator    string `xml:"rank_indicator"`
			TxPower          string `xml:"tx_power"`
			MainDLMcs        string `xml:"mainDLMcs"`
			DiversityDLMcs   string `xml:"diversityDLMcs"`
			ULFREQ           string `xml:"ULFREQ"`
			DLFREQ           string `xml:"DLFREQ"`
			DlEarfcnStart    string `xml:"dl_earfcn_start"`
			DlEarfcnEnd      string `xml:"dl_earfcn_end"`
			UlEarfcnStart    string `xml:"ul_earfcn_start"`
			UlEarfcnEnd      string `xml:"ul_earfcn_end"`
			FrequencyRange   string `xml:"frequency_range"`
			Iccid            string `xml:"iccid"`
			OPERATIONMODE    string `xml:"OPERATION_MODE"`
			UlBandwidth      string `xml:"ul_bandwidth"`
			PLMN             string `xml:"PLMN"`
			EutranCellid     string `xml:"eutran_cellid"`
			ECGI             string `xml:"ECGI"`
			CELLID           string `xml:"CELL_ID"`
			Band1            string `xml:"band1"`
			Band2            string `xml:"band2"`
		} `xml:"lte"`
	} `xml:"eng"`
	Response struct {
		Text           string `xml:",chardata"`
		ResponseStatus string `xml:"response_status"`
	} `xml:"response"`
}

func (r EngInfo) ToJSON() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return b
}

type RequestParam struct {
	XMLName xml.Name  `xml:"RGW"`
	Text    string    `xml:",chardata"`
	Param   Parameter `xml:"param"`
}

type Parameter struct {
	Text      string `xml:",chardata"`
	Method    string `xml:"method"`
	Session   string `xml:"session"`
	ObjPath   string `xml:"obj_path"`
	ObjMethod string `xml:"obj_method"`
}

type DashboardInfo struct {
	XMLName          xml.Name `xml:"RGW"`
	Text             string   `xml:",chardata"`
	ID               string   `xml:"id"`
	CelluarBasicInfo struct {
		Text               string `xml:",chardata"`
		SysMode            string `xml:"sys_mode"`
		DataMode           string `xml:"data_mode"`
		Rssi               string `xml:"rssi"`
		RegStatus          string `xml:"RegStatus"`
		IMEI               string `xml:"IMEI"`
		IMSI               string `xml:"IMSI"`
		MSISDN             string `xml:"MSISDN"`
		NetworkName        string `xml:"network_name"`
		RoamingNetworkName string `xml:"roaming_network_name"`
		RoamingStatus      string `xml:"roaming_status"`
	} `xml:"celluar_basic_info"`
	Contextlist struct {
		Text string `xml:",chardata"`
		Item struct {
			Text  string `xml:",chardata"`
			Index string `xml:"index,attr"`
			List  struct {
				Text             string `xml:",chardata"`
				WanType          string `xml:"wan_type"`
				ConnectionNum    string `xml:"connection_num"`
				ConnectionStatus string `xml:"connection_status"`
				PdpType          string `xml:"pdp_type"`
				IpType           string `xml:"ip_type"`
				PrimaryCid       string `xml:"primary_cid"`
				Qci              string `xml:"qci"`
				Apn              string `xml:"apn"`
				LteApn           string `xml:"lte_apn"`
				Usr2g3g          string `xml:"usr_2g3g"`
				Pswd2g3g         string `xml:"pswd_2g3g"`
				Authtype2g3g     string `xml:"authtype_2g3g"`
				Usr4g            string `xml:"usr_4g"`
				Pswd4g           string `xml:"pswd_4g"`
				Authtype4g       string `xml:"authtype_4g"`
				Mtu              string `xml:"mtu"`
				AutoApn          string `xml:"auto_apn"`
				ConnectMode      string `xml:"connect_mode"`
				DataOnRoaming    string `xml:"data_on_roaming"`
				Ipv4Ip           string `xml:"ipv4_ip"`
				Ipv4Dns1         string `xml:"ipv4_dns1"`
				Ipv4Dns2         string `xml:"ipv4_dns2"`
				Ipv4Gateway      string `xml:"ipv4_gateway"`
				Ipv4Submask      string `xml:"ipv4_submask"`
			} `xml:"list"`
		} `xml:"Item"`
	} `xml:"contextlist"`
}

type StatCommonData struct {
	XMLName    xml.Name `xml:"RGW"`
	Text       string   `xml:",chardata"`
	Statistics struct {
		Text             string `xml:",chardata"`
		RxBytes          string `xml:"rx_bytes"`
		TxBytes          string `xml:"tx_bytes"`
		RxTxBytes        string `xml:"rx_tx_bytes"`
		ErrorBytes       string `xml:"error_bytes"`
		MonthRxBytes     string `xml:"month_rx_bytes"`
		MonthTxBytes     string `xml:"month_tx_bytes"`
		TotalRxBytes     string `xml:"total_rx_bytes"`
		TotalTxBytes     string `xml:"total_tx_bytes"`
		TotalRxTxBytes   string `xml:"total_rx_tx_bytes"`
		TotalErrorBytes  string `xml:"total_error_bytes"`
		TotalRxPackets   string `xml:"total_rx_packets"`
		TotalTxPackets   string `xml:"total_tx_packets"`
		TotalRxTxPackets string `xml:"total_rx_tx_packets"`
		Upthrpt          string `xml:"upthrpt"`
		Dnthrpt          string `xml:"dnthrpt"`
		AvgUpthrpt       string `xml:"avg_upthrpt"`
		AvgDnthrpt       string `xml:"avg_dnthrpt"`
	} `xml:"statistics"`
}
