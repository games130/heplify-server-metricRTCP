package metric

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/games130/logp"
	"github.com/games130/heplify-server-metricRTCP/decoder"
)

var (
	errNoComma  = fmt.Errorf("no comma in string")
	errNoTwoVal = fmt.Errorf("no two values in string")
)

/*
func (p *Prometheus) dissectRTCPXRStats(nodeID, stats string) {
	if nlr, err := strconv.ParseFloat(extractXR("NLR=", stats), 64); err == nil {
		vqrtcpxrNLR.WithLabelValues(nodeID).Set(nlr)
	}
	if jdr, err := strconv.ParseFloat(extractXR("JDR=", stats), 64); err == nil {
		vqrtcpxrJDR.WithLabelValues(nodeID).Set(jdr)
	}
	if iaj, err := strconv.ParseFloat(extractXR("IAJ=", stats), 64); err == nil {
		vqrtcpxrIAJ.WithLabelValues(nodeID).Set(iaj)
	}
	if moslq, err := strconv.ParseFloat(extractXR("MOSLQ=", stats), 64); err == nil {
		vqrtcpxrMOSLQ.WithLabelValues(nodeID).Set(moslq)
	}
	if moscq, err := strconv.ParseFloat(extractXR("MOSCQ=", stats), 64); err == nil {
		vqrtcpxrMOSCQ.WithLabelValues(nodeID).Set(moscq)
	}
}

func (p *Prometheus) dissectXRTPStats(tn, stats string) {
	var err error
	plr, pls, jir, jis, dle, r, mos := 0, 0, 0, 0, 0, 0.0, 0.0

	if cs, err := strconv.ParseFloat(extractXR("CS=", stats), 64); err == nil {
		xrtpCS.WithLabelValues(tn).Set(cs / 1000)
	}

	if plt := extractXR("PL=", stats); len(plt) > 1 {
		if plr, pls, err = splitCommaInt(plt); err == nil {
			xrtpPLR.WithLabelValues(tn).Set(float64(plr))
			xrtpPLS.WithLabelValues(tn).Set(float64(pls))
		}
	}

	if jit := extractXR("JI=", stats); len(jit) > 1 {
		if jir, jis, err = splitCommaInt(jit); err == nil {
			xrtpJIR.WithLabelValues(tn).Set(float64(jir))
			xrtpJIS.WithLabelValues(tn).Set(float64(jis))
		}
	}

	if dlt := extractXR("DL=", stats); len(dlt) > 1 {
		if dle, _, err = splitCommaInt(dlt); err == nil || dle > 0 {
			xrtpDLE.WithLabelValues(tn).Set(float64(dle))
		}
	}

	pr, _ := strconv.Atoi(extractXR("PR=", stats))
	ps, _ := strconv.Atoi(extractXR("PS=", stats))
	if pr == 0 && ps == 0 {
		pr, ps = 1, 1
	}

	loss := ((plr + pls) * 100) / (pr + ps)
	el := (jir * 2) + (dle + 10)

	if el < 160 {
		r = 93.2 - (float64(el) / 40)
	} else {
		r = 93.2 - (float64(el-120) / 10)
	}
	r = r - (float64(loss) * 2.5)

	mos = 1 + (0.035)*r + (0.000007)*r*(r-60)*(100-r)
	if mos < 1 || mos > 5 {
		mos = 1
	}
	xrtpMOS.WithLabelValues(tn).Set(mos)
}
*/

func (p *Prometheus) dissectRTCPStats(nodeID string, pkt *decoder.HEP) {
	jsonparser.EachKey([]byte(pkt.Payload), func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		switch idx {
		case 0:
			if fractionLost, err := jsonparser.ParseFloat(value); err == nil {
				rtcpFractionLost.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(normMax(fractionLost))
			}
		case 1:
			if packetsLost, err := jsonparser.ParseFloat(value); err == nil {
				rtcpPacketsLost.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(normMax(packetsLost))
			}
		case 2:
			if iaJitter, err := jsonparser.ParseFloat(value); err == nil {
				rtcpJitter.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(normMax(iaJitter))
			}
		case 3:
			if dlsr, err := jsonparser.ParseFloat(value); err == nil {
				rtcpDLSR.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(normMax(dlsr))
			}
		case 4:
			if fractionLost, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrFractionLost.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(fractionLost)
			}
		case 5:
			if fractionDiscard, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrFractionDiscard.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(fractionDiscard)
			}
		case 6:
			if burstDensity, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrBurstDensity.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(burstDensity)
			}
		case 7:
			if gapDensity, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrGapDensity.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(gapDensity)
			}
		case 8:
			if burstDuration, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrBurstDuration.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(burstDuration)
			}
		case 9:
			if gapDuration, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrGapDuration.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(gapDuration)
			}
		case 10:
			if roundTripDelay, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrRoundTripDelay.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(roundTripDelay)
			}
		case 11:
			if endSystemDelay, err := jsonparser.ParseFloat(value); err == nil {
				rtcpxrEndSystemDelay.WithLabelValues(nodeID, pkt.SrcIP, pkt.DstIP).Set(endSystemDelay)
			}
		}
	}, rtcpPaths...)
}

func normMax(val float64) float64 {
	if val > 10000000 {
		return 0
	}
	return val
}

/*
func splitCommaInt(str string) (int, int, error) {
	var err error
	var one, two int
	sp := strings.IndexRune(str, ',')
	if sp == -1 {
		return one, two, errNoComma
	}

	if one, err = strconv.Atoi(str[0:sp]); err != nil {
		return one, two, err
	}
	if len(str)-1 >= sp+1 {
		if two, err = strconv.Atoi(str[sp+1:]); err != nil {
			return one, two, err
		}
	} else {
		return one, two, errNoTwoVal
	}
	return one, two, nil
}
*/
