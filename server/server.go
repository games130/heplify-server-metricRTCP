package input

import (
	"runtime"
	"sync"
	"sync/atomic"
	"context"
	//"time"

	"github.com/games130/logp"
	"github.com/games130/heplify-server-metricRTCP/config"
	"github.com/games130/heplify-server-metricRTCP/decoder"
	"github.com/games130/heplify-server-metricRTCP/metric"
	//proto "github.com/games130/heplify-server-metricRTCP/proto"
	proto "github.com/games130/microProtocRTCP"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/nats"
)

type HEPInput struct {
	inCh	      chan *proto.RTCPpkt
	pmCh          chan *decoder.HEP
	wg            *sync.WaitGroup
	quit          chan bool
	perMSGDebug   bool
	stats         HEPStats
}

type HEPStats struct {
	HEPCount 		uint64
	INVITECount		uint64
	REGISTERCount		uint64
	BYECount		uint64
	PRACKCount		uint64
	R180Count		uint64
	R183Count 		uint64
	R200Count 		uint64
	R400Count 		uint64
	R404Count 		uint64
	R406Count 		uint64
	R408Count 		uint64
	R416Count 		uint64
	R420Count 		uint64
	R422Count 		uint64
	R480Count 		uint64
	R481Count 		uint64
	R484Count 		uint64
	R485Count 		uint64
	R488Count 		uint64
	R500Count 		uint64
	R502Count 		uint64
	R503Count 		uint64
	R504Count 		uint64
	R603Count 		uint64
	R604Count 		uint64
	OtherCount 		uint64
}

func (h *HEPInput) subEv(ctx context.Context, event *proto.RTCPpkt) error {
	//log.Logf("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	//fmt.Println("received %s and %s", event.GetCID(), event.GetFirstMethod())
	
	// do something with event
	atomic.AddUint64(&h.stats.HEPCount, 1)
	h.inCh <- event
	
	return nil
}

func NewHEPInput() *HEPInput {
	h := &HEPInput{
		inCh:      make(chan *proto.RTCPpkt, 40000),
		pmCh:	   make(chan *decoder.HEP, 40000),
		wg:        &sync.WaitGroup{},
		quit:      make(chan bool),
	}
	
	h.perMSGDebug = config.Setting.PerMSGDebug

	return h
}

func (h *HEPInput) Run() {
	logp.Info("creating hepWorker totaling: %s", runtime.NumCPU()*4)
	for n := 0; n < runtime.NumCPU()*4; n++ {
		h.wg.Add(1)
		go h.hepWorker()
	}
	
	b := nats.NewBroker(
		broker.Addrs(config.Setting.BrokerAddr),
	)
	
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.srv.metric"),
		micro.Broker(b),
	)
	// parse command line
	service.Init()
	
	// register subscriber
	micro.RegisterSubscriber(config.Setting.BrokerTopic, service.Server(), h.subEv, server.SubscriberQueue(config.Setting.BrokerQueue))

	m := metric.New("prometheus")
	m.Chan = h.pmCh
	
	//fmt.Println("micro server before start")
	go func (){
		if err := service.Run(); err != nil {
			logp.Err("%v", err)
		}
	}()	

	//fmt.Println("metric server before start")
	if err := m.Run(); err != nil {
		logp.Err("%v", err)
	}
	defer m.End()
	h.wg.Wait()
}

func (h *HEPInput) End() {
	logp.Info("stopping heplify-server...")

	h.quit <- true
	<-h.quit

	logp.Info("heplify-server has been stopped")
}

func (h *HEPInput) hepWorker() {
	for {
		select {
		case <-h.quit:
			h.quit <- true
			h.wg.Done()
			return
		case msg := <-h.inCh:
			//fmt.Println("want to start decoding %s and %s", msg.GetCID(), msg.GetFirstMethod())
			hepPkt, _ := decoder.DecodeHEP(msg)
			
			if h.perMSGDebug {
				logp.Info("perMSGDebug: ,HEPCount,%s, SrcIP,%s, DstIP,%s, CID,%s", h.stats.HEPCount, hepPkt.SrcIP, hepPkt.DstIP, hepPkt.CID)
			}
			
			h.pmCh <- hepPkt
		}
	}
}


