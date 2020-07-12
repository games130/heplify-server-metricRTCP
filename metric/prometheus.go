package metric

import (
	"encoding/binary"
	"fmt"
	"strings"
	"sync"
	"regexp"
	"time"
	"os"
	"bufio"

	"github.com/games130/logp"
	"github.com/games130/heplify-server-metricRTCP/config"
	"github.com/games130/heplify-server-metricRTCP/decoder"
	
)

type Prometheus struct {
	TargetEmpty   bool
	TargetIP      []string
	TargetName    []string
	TargetMap     map[string]string
	TargetConf    *sync.RWMutex
	perMSGDebug   bool
	count         int64
}

func (p *Prometheus) setup() (err error) {
	p.TargetConf = new(sync.RWMutex)
	p.TargetIP = strings.Split(cutSpace(config.Setting.PromTargetIP), ",")
	p.TargetName = strings.Split(cutSpace(config.Setting.PromTargetName), ",")
	p.perMSGDebug = config.Setting.PerMSGDebug
	p.count = 1
	
	if len(p.TargetIP) == len(p.TargetName) && p.TargetIP != nil && p.TargetName != nil {
		if len(p.TargetIP[0]) == 0 || len(p.TargetName[0]) == 0 {
			logp.Info("expose metrics without or unbalanced targets")
			p.TargetIP[0] = ""
			p.TargetName[0] = ""
			p.TargetEmpty = true
		} else {
			for i := range p.TargetName {
				logp.Info("prometheus tag assignment %d: %s -> %s", i+1, p.TargetIP[i], p.TargetName[i])
			}
			p.TargetMap = make(map[string]string)
			for i := 0; i < len(p.TargetName); i++ {
				p.TargetMap[p.TargetIP[i]] = p.TargetName[i]
			}
		}
	} else {
		logp.Info("please give every PromTargetIP a unique IP and PromTargetName a unique name")
		return fmt.Errorf("faulty PromTargetIP or PromTargetName")
	}

	return err
}

func (p *Prometheus) expose(hCh chan *decoder.HEP) {
	for pkt := range hCh {
		if p.perMSGDebug {
				logp.Info("perMSGDebug-prom: ,Count,%s, SrcIP,%s, DstIP,%s, CID,%s, FirstMethod,%s, FromUser,%s, ToUser,%s", p.count, pkt.SrcIP, pkt.DstIP, pkt.CallID, pkt.FirstMethod, pkt.FromUser, pkt.ToUser)
				p.count++
		}
		
		//logp.Info("exposing some packet %s and %s", pkt.CID, pkt.FirstMethod)
		//fmt.Println("exposing some packet %s and %s", pkt.CID, pkt.FirstMethod)

		var st, dt string
		if pkt != nil && pkt.ProtoType == 5 {
			if !p.TargetEmpty {
				p.checkTargetPrefix(pkt)
			}

		}
	}
}

//new
func (p *Prometheus) checkTargetPrefix(pkt *decoder.HEP) {
	st, sOk := p.TargetMap[pkt.SrcIP]
	if sOk {
		p.dissectRTCPStats(st, pkt)
	}
	
	dt, dOk := p.TargetMap[pkt.DstIP]
	if dOk {
		p.dissectRTCPStats(dt, pkt)
	}
}

func (p *Prometheus) end(){
}
