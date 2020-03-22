package datalistener

import (
	"context"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

const minID = 2
const maxID = 8

//Data struct
type Data struct {
	Time int64     `json:"time"`
	Data []float64 `json:"raw_data"`
}

type Item struct {
	Pressure     float64   `json:"Pressure"`
	Humidity     float64   `josn:"Humidity"`
	TemperatureR float64   `josn:"TemperatureR"`
	TemperatureA float64   `json:"TemperatureA"`
	PH           float64   `json:"pH"`
	FlowRate     float64   `json:"FlowRate"`
	CO           float64   `json:"CO"`
	EventTime    time.Time `json:"EventTime"`
}

func NilData() Item {
	return Item{
		Pressure:     0,
		Humidity:     0,
		TemperatureR: 0,
		TemperatureA: 0,
		PH:           0,
		FlowRate:     0,
		CO:           0,
		EventTime:    time.Now(),
	}
}

func GetClient(opcEnd string) *opcua.Client {
	/*
		var endpointFlag = flag.Lookup("endpoint")
		if endpointFlag == nil {
			flag.String("endpoint", opcEnd, "OPC UA Endpoint URL")
		}
		flag.Parse()*/
	ctx := context.Background()

	var c *opcua.Client = opcua.NewClient(opcEnd, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		return nil
	}

	return c
}

func GetData(client *opcua.Client) (Item, []float64, error) {
	req := ua.ReadRequest{
		MaxAge:             2000,
		TimestampsToReturn: ua.TimestampsToReturnBoth,
		NodesToRead:        []*ua.ReadValueID{},
	}

	for i := uint32(minID); i <= maxID; i++ {
		var nodeid = ua.NewNumericNodeID(2, i)
		req.NodesToRead = append(req.NodesToRead, &ua.ReadValueID{NodeID: nodeid})
	}

	resp, err := client.Read(&req)
	if err != nil {
		return NilData(), nil, err
	}

	if resp.Results[0].Status != ua.StatusOK {
		return NilData(), nil, err
	}

	var out []float64

	for _, res := range resp.Results {
		out = append(out, float64(res.Value.Int())+res.Value.Float())
	}

	return Item{
		Pressure:     out[0],
		Humidity:     out[1],
		TemperatureR: out[2],
		TemperatureA: out[3],
		PH:           out[4],
		FlowRate:     out[5],
		CO:           out[6],
		EventTime:    resp.Header().Timestamp,
	}, out, nil
}
