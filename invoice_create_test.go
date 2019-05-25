package uculqi

import (
	"reflect"
	"testing"
)

func Test_newInvoiceWithMinimal(t *testing.T) {
	type args struct {
		companyInfo *CompanyInfo
		info        *MinimalInformation
	}
	tests := []struct {
		name    string
		args    args
		want    *Invoice
		wantErr bool
	}{
		{name: "null", args: args{nil, nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newInvoiceWithMinimal(tt.args.companyInfo, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInvoiceWithMinimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInvoiceWithMinimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newFullInvoice(t *testing.T) {
	type args struct {
		companyInfo *CompanyInfo
		order       *Order
	}
	tests := []struct {
		name    string
		args    args
		want    *Invoice
		wantErr bool
	}{
		{name: "null", args: args{nil, nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newFullInvoice(tt.args.companyInfo, tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("newFullInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFullInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
