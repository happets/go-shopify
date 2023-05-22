package goshopify

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
	"time"
)

func TestFulfillmentOrderList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/orders/123/fulfillment_orders.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"fulfillment_orders": [{"id":1},{"id":2}]}`))

	fulfillmentService := &FulfillmentOrderServiceOp{client: client}

	fulfillmentOrders, err := fulfillmentService.List(123, nil)
	if err != nil {
		t.Errorf("FulfillmentOrder.List returned error: %v", err)
	}

	expected := []FulfillmentOrder{{Id: 1}, {Id: 2}}
	if !reflect.DeepEqual(fulfillmentOrders, expected) {
		t.Errorf("FulfillmentOrder.List returned %+v, expected %+v", fulfillmentOrders, expected)
	}
}

// TODO
func TestFulfillmentOrderGet(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
		options       interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Get(tt.args.fulfillmentID, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderCancel(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Cancel(tt.args.fulfillmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cancel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderClose(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
		message       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Close(tt.args.fulfillmentID, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Close() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderHold(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
		notify        bool
		reason        FulfillmentOrderHoldReason
		notes         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Hold(tt.args.fulfillmentID, tt.args.notify, tt.args.reason, tt.args.notes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hold() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderMove(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
		moveRequest   FulfillmentOrderMoveRequest
		options       interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrderMoveResource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Move(tt.args.fulfillmentID, tt.args.moveRequest, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderOpen(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Open(tt.args.fulfillmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Open() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderReleaseHold(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.ReleaseHold(tt.args.fulfillmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReleaseHold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReleaseHold() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderReschedule(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FulfillmentOrder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			got, err := s.Reschedule(tt.args.fulfillmentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reschedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reschedule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO
func TestFulfillmentOrderSetDeadline(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		fulfillmentIDs []int64
		deadline       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FulfillmentOrderServiceOp{
				client: tt.fields.client,
			}
			if err := s.SetDeadline(tt.args.fulfillmentIDs, tt.args.deadline); (err != nil) != tt.wantErr {
				t.Errorf("SetDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
