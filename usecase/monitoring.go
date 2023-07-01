package usecase

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/luqman-v1/orbit-pro-monitor/model"
	"github.com/luqman-v1/orbit-pro-monitor/repo"
	"github.com/luqman-v1/orbit-pro-monitor/util"
	"github.com/sirupsen/logrus"
)

type Monitor struct {
	Repo   repo.IOrbit
	cookie string
	model.EngInfo
	model.DashboardInfo
	model.StatCommonData
}
type IMonitor interface {
	Monitor(ctx context.Context) error
}

func NewMonitor(uc Monitor) IMonitor {
	return &Monitor{
		Repo: uc.Repo,
	}
}

func (m *Monitor) Monitor(ctx context.Context) error {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	cookie, err := m.authOrbit(ctx)
	if err != nil {
		return err
	}
	m.cookie = cookie

	tickerCount := 1
	_ = m.draw(tickerCount, Draw{
		Sinr:       m.getGaugeSINR(),
		Rsrp:       m.getGaugeRSRP(),
		Rsrq:       m.getGaugeRSRQ(),
		Rssi:       m.getGaugeRSSI(),
		TextHeader: m.getHeader(),
	})
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return nil
			}
		case <-ticker:
			_ = m.sync(ctx)
			_ = m.draw(tickerCount, Draw{
				Sinr:       m.getGaugeSINR(),
				Rsrp:       m.getGaugeRSRP(),
				Rsrq:       m.getGaugeRSRQ(),
				Rssi:       m.getGaugeRSSI(),
				TextHeader: m.getHeader(),
			})
			tickerCount++
		}
	}
	return nil
}
func (m *Monitor) getTitleGauge(title string, value int) string {
	return fmt.Sprintf("%v(%vdB)", title, value)
}

func (m *Monitor) sync(ctx context.Context) error {
	result, err := m.getEngInfo(ctx)
	if err != nil {
		return err
	}
	info, err := m.getDashboardInfo(ctx)
	if err != nil {
		return err
	}
	statsCommonData, err := m.getStatCommonData(ctx)
	if err != nil {
		return err
	}
	m.StatCommonData = statsCommonData
	m.EngInfo = result
	m.DashboardInfo = info
	return nil
}

type Draw struct {
	Sinr       *widgets.Gauge
	Rsrp       *widgets.Gauge
	Rsrq       *widgets.Gauge
	Rssi       *widgets.Gauge
	TextHeader *widgets.Paragraph
}

func (m *Monitor) draw(count int, draw Draw) error {
	result := m.EngInfo
	sinr, _ := strconv.Atoi(result.Eng.Lte.Sinr)
	rsrp, _ := strconv.Atoi(result.Eng.Lte.Rsrp)
	rsrq, _ := strconv.Atoi(result.Eng.Lte.Rsrq)
	rssi, _ := strconv.Atoi(result.Eng.Lte.Rssi)
	draw.Sinr.Title = m.getTitleGauge("SINR", sinr)
	draw.Sinr.Percent = util.GetSirn(sinr)

	draw.Rsrp.Title = m.getTitleGauge("RSRP", rsrp)
	draw.Rsrp.Percent = util.GetRsp(rsrp)

	draw.Rsrq.Title = m.getTitleGauge("RSRQ", rsrq)
	draw.Rsrq.Percent = util.GetRsrq(rsrp)
	rssi = util.GetRSSI(rssi)
	draw.Rssi.Title = m.getTitleGauge("RSSI", rssi)
	draw.Rssi.Percent = util.GetRssiPercentage(rssi)

	ui.Render(draw.TextHeader, draw.Sinr, draw.Rsrp, draw.Rsrq, draw.Rssi)
	return nil
}

func (m *Monitor) getGaugeSINR() *widgets.Gauge {
	g := widgets.NewGauge()
	g.SetRect(0, 13, 25, 16)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan
	return g
}

func (m *Monitor) getGaugeRSSI() *widgets.Gauge {
	g := widgets.NewGauge()
	g.SetRect(25, 13, 50, 16)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan
	return g
}

func (m *Monitor) getGaugeRSRQ() *widgets.Gauge {
	g := widgets.NewGauge()
	g.SetRect(0, 10, 25, 13)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan
	return g
}

func (m *Monitor) getGaugeRSRP() *widgets.Gauge {
	g := widgets.NewGauge()
	g.SetRect(25, 10, 50, 13)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan
	return g
}

func (m *Monitor) getHeader() *widgets.Paragraph {
	info := m.DashboardInfo
	engInfo := m.EngInfo
	statCommonData := m.StatCommonData

	dl, _ := strconv.ParseFloat(statCommonData.Statistics.Dnthrpt, 64)
	ul, _ := strconv.ParseFloat(statCommonData.Statistics.Upthrpt, 64)
	p := widgets.NewParagraph()
	p.Title = "Orbit Pro Monitor by luqman"
	p.Text = fmt.Sprintf("PRESS q TO QUIT \n"+
		"ISP : %v          Band : %v \n"+
		"APN : %v         PCI  :  %v \n"+
		"CellID : %v     EARFCN : %v \n"+
		"ULBW : %vM             DLBW : %vM \n"+
		"Recieved speed : %vMb   Sent speed : %vMb \n",
		info.CelluarBasicInfo.NetworkName,
		engInfo.Eng.Lte.Band1+"/"+engInfo.Eng.Lte.Band2,
		info.Contextlist.Item.List.LteApn,
		engInfo.Eng.Lte.PhyCellID,
		engInfo.Eng.Lte.CellID,
		engInfo.Eng.Lte.DlEuarfcn,
		engInfo.Eng.Lte.UlBandwidth,
		engInfo.Eng.Lte.DlBandwidth,
		util.ByteToMB(dl),
		util.ByteToMB(ul),
	)

	p.SetRect(0, 0, 50, 10)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan
	return p
}

func (m *Monitor) updateParagraph(count int, p *widgets.Paragraph) *widgets.Paragraph {
	if count%2 == 0 {
		p.TextStyle.Fg = ui.ColorRed
	} else {
		p.TextStyle.Fg = ui.ColorWhite
	}
	return p
}

func (m *Monitor) authOrbit(ctx context.Context) (string, error) {
	cookie, err := m.Repo.Auth(ctx)
	if err != nil {
		logrus.Error("err at auth", err)
		return "", err
	}
	return cookie, nil
}

func (m *Monitor) getEngInfo(ctx context.Context) (model.EngInfo, error) {
	var result model.EngInfo
	info, err := m.Repo.SetInfo(ctx, model.RequestParam{
		Param: model.Parameter{
			Method:    "call",
			Session:   "000",
			ObjPath:   "cm",
			ObjMethod: "get_eng_info",
		},
	}, m.cookie)
	if err != nil {
		logrus.Error("err at getEngInfo", err)
		return result, err
	}
	_ = xml.Unmarshal(info, &result)
	return result, nil
}

func (m *Monitor) getStatCommonData(ctx context.Context) (model.StatCommonData, error) {
	var result model.StatCommonData
	info, err := m.Repo.SetInfo(ctx, model.RequestParam{
		Param: model.Parameter{
			Method:    "call",
			Session:   "000",
			ObjPath:   "statistics",
			ObjMethod: "stat_get_common_data",
		},
	}, m.cookie)
	if err != nil {
		logrus.Error("err at getStatCommonData", err)
		return result, err
	}
	_ = xml.Unmarshal(info, &result)
	return result, nil
}

func (m *Monitor) getDashboardInfo(ctx context.Context) (model.DashboardInfo, error) {
	var result model.DashboardInfo
	info, err := m.Repo.SetInfo(ctx, model.RequestParam{
		Param: model.Parameter{
			Method:    "call",
			Session:   "000",
			ObjPath:   "cm",
			ObjMethod: "get_link_context",
		},
	}, m.cookie)
	if err != nil {
		logrus.Error("err at getDashboardInfo", err)
		return result, err
	}
	_ = xml.Unmarshal(info, &result)
	return result, nil
}
