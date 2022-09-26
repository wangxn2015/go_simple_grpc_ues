package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/wangxn2015/go_simple_grpc_ues/api/track_msg"
	"math/rand"
	"time"
)

type UeStore interface {
	Search(ctx context.Context, found func(ue *track_msg.UEInfo) error) error
}

type InMemoryUeStore struct {
	data map[uint64]*track_msg.UEInfo
}

func NewInMemoryUeStore() *InMemoryUeStore {
	return &InMemoryUeStore{
		data: make(map[uint64]*track_msg.UEInfo),
	}
}

func (store *InMemoryUeStore) Search(ctx context.Context, found func(ue *track_msg.UEInfo) error) error {

	if len(store.data) == 0 {
		store.generateData()
	}

	for j := 0; j < 100; j++ {
		store.moveUes()
		for _, ue := range store.data {
			if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
				fmt.Println("context is cancelled")
				return nil
			}
			other, err := deepCopy(ue)
			if err != nil {
				return err
			}
			err = found(other)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Second)
	}

	return nil
}

func (store *InMemoryUeStore) generateData() {
	for i := 0; i < 3; i++ {
		store.data[uint64(2000+i)] = &track_msg.UEInfo{
			Imsi: uint64(2000 + i),
			Loca: &track_msg.Location{
				Lat: 40.075 + float64(i+1)*0.0005*rand.Float64(),
				Lng: 116.24 + float64(i+1)*0.0005*rand.Float64(),
			},
		}
	}

}

func (store *InMemoryUeStore) moveUes() {
	for i := 0; i < 3; i++ {
		Location := &track_msg.Location{
			Lat: 40.075 + float64(i+1)*0.0005*rand.Float64(),
			Lng: 116.24 + float64(i+1)*0.0005*rand.Float64(),
		}

		store.data[uint64(2000+i)].Loca = Location
	}

}

func deepCopy(ue *track_msg.UEInfo) (*track_msg.UEInfo, error) {
	other := &track_msg.UEInfo{}

	err := copier.Copy(other, ue)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
