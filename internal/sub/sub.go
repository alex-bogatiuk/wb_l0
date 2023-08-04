package sub

import (
	"github.com/alex-bogatiuk/wb_l0/internal/storage"
	"github.com/gookit/slog"
	"github.com/nats-io/stan.go"
)

type NatsSub struct {
	Nc stan.Conn
	Ss storage.OrderStorageService
}

func CreateSub(oss storage.OrderStorageService) *NatsSub {
	sc := NatsSub{
		Ss: oss,
	}

	return &sc
}

func (nSub *NatsSub) Connect(clusterID string, clientID string, URL string) error {
	nc, err := stan.Connect(clusterID, clientID, stan.NatsURL(URL))
	if err != nil {
		return err
	}
	nSub.Nc = nc
	return err
}

func (nSub *NatsSub) Close() {
	if nSub.Nc != nil {
		nSub.Nc.Close()
	}
}

func (nSub *NatsSub) Subscribe(channel string, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	sub, err := nSub.Nc.Subscribe(channel, nSub.SaveOrderHandler, opts...)
	if err != nil {
		slog.Error("nats subscribe error:", err)
	}
	return sub, err
}

func (nSub *NatsSub) SaveOrderHandler(msg *stan.Msg) {
	slog.Info(msg.Data)
	err := nSub.Ss.SaveOrder(msg.Data)
	if err != nil {
		slog.Error("save order data error:", err)

	}
}
