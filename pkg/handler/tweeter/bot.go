package tweeter

import (
	"context"

	"github.com/mattmoor/bindings/pkg/twitter"

	"github.com/mattmoor/knobots/pkg/handler"
)

type tweeter struct {
}

var _ handler.Interface = (*tweeter)(nil)

func New(context.Context) handler.Interface {
	return &tweeter{}
}

func (*tweeter) GetType() interface{} {
	return &Tweet{}
}

func (gt *tweeter) Handle(ctx context.Context, x interface{}) (handler.Response, error) {
	rrr := x.(*Tweet)

	client, err := twitter.NewUserClient(ctx)
	if err != nil {
		return nil, err
	}

	_, _, err = client.Statuses.Update(rrr.Message, nil)
	return nil, err
}

type Tweet struct {
	Message string
}

var _ handler.Response = (*Tweet)(nil)

func (*Tweet) GetSource() string {
	return "https://github.com/mattmoor/knobots/cmd/tweeter"
}

func (*Tweet) GetType() string {
	return "dev.mattmoor.knobots.tweeter"
}
