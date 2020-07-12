package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// X-RTP-Stat Metrics
/*	xrtpCS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_cs",
		Help: "XRTP call setup time"},
		[]string{"target_name"})
	xrtpJIR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_jir",
		Help: "XRTP received jitter"},
		[]string{"target_name"})
	xrtpJIS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_jis",
		Help: "XRTP sent jitter"},
		[]string{"target_name"})
	xrtpPLR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_plr",
		Help: "XRTP received packets lost"},
		[]string{"target_name"})
	xrtpPLS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_pls",
		Help: "XRTP sent packets lost"},
		[]string{"target_name"})
	xrtpDLE = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_dle",
		Help: "XRTP mean rtt"},
		[]string{"target_name"})
	xrtpMOS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_mos",
		Help: "XRTP mos"},
		[]string{"target_name"})
*/

	// RTCP Metrics
	rtcpFractionLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcp_fraction_lost",
		Help: "RTCP fraction lost"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpPacketsLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcp_packets_lost",
		Help: "RTCP packets lost"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcp_jitter",
		Help: "RTCP jitter"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpDLSR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcp_dlsr",
		Help: "RTCP dlsr"},
		[]string{"target_name","source_ip", "destination_ip"})

	// RTCP-XR Metrics
	rtcpxrFractionLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_fraction_lost",
		Help: "RTCPXR fraction lost"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrFractionDiscard = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_fraction_discard",
		Help: "RTCPXR fraction discard"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrBurstDensity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_burst_density",
		Help: "RTCPXR burst density"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrBurstDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_burst_duration",
		Help: "RTCPXR burst duration"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrGapDensity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_gap_density",
		Help: "RTCPXR gap density"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrGapDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_gap_duration",
		Help: "RTCPXR gap duration"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrRoundTripDelay = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_round_trip_delay",
		Help: "RTCPXR round trip delay"},
		[]string{"target_name","source_ip", "destination_ip"})
	rtcpxrEndSystemDelay = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplifyown_rtcpxr_end_system_delay",
		Help: "RTCPXR end system delay"},
		[]string{"target_name","source_ip", "destination_ip"})

	// VQ-RTCP-XR Metrics
/*	vqrtcpxrNLR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_nlr",
		Help: "VQ-RTCPXR network packet loss rate"},
		[]string{"node_id"})
	vqrtcpxrJDR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_jdr",
		Help: "VQ-RTCPXR jitter buffer discard rate"},
		[]string{"node_id"})
	vqrtcpxrIAJ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_iaj",
		Help: "VQ-RTCPXR interarrival jitter"},
		[]string{"node_id"})
	vqrtcpxrMOSLQ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_moslq",
		Help: "VQ-RTCPXR MOS listening voice quality"},
		[]string{"node_id"})
	vqrtcpxrMOSCQ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_moscq",
		Help: "VQ-RTCPXR MOS conversation voice quality"},
		[]string{"node_id"})
*/

	// JSON Paths

	rtcpPaths = [][]string{
		[]string{"report_blocks", "[0]", "fraction_lost"},
		[]string{"report_blocks", "[0]", "packets_lost"},
		[]string{"report_blocks", "[0]", "ia_jitter"},
		[]string{"report_blocks", "[0]", "dlsr"},
		[]string{"report_blocks_xr", "fraction_lost"},
		[]string{"report_blocks_xr", "fraction_discard"},
		[]string{"report_blocks_xr", "burst_density"},
		[]string{"report_blocks_xr", "gap_density"},
		[]string{"report_blocks_xr", "burst_duration"},
		[]string{"report_blocks_xr", "gap_duration"},
		[]string{"report_blocks_xr", "round_trip_delay"},
		[]string{"report_blocks_xr", "end_system_delay"},
	}
)
