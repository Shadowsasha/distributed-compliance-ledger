// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"
)

func TestAccountRole_Validate(t *testing.T) {
	tests := []struct {
		name    string
		role    AccountRole
		wantErr bool
	}{
		{
			name:    "invalid role",
			wantErr: true,
		},
		{
			name:    "valid  vendor role",
			role:    Vendor,
			wantErr: false,
		},
		{
			name:    "valid  certification center role",
			role:    CertificationCenter,
			wantErr: false,
		},
		{
			name:    "valid  trustee role",
			role:    Trustee,
			wantErr: false,
		},
		{
			name:    "valid  node admin role",
			role:    NodeAdmin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.role.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AccountRole.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	type args struct {
		ba         *types.BaseAccount
		roles      AccountRoles
		approvals  []*Grant
		rejects    []*Grant
		vendorID   int32
		productIDs []*commontypes.Uint16Range
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "valid account all roles",
			args: args{
				ba:         &types.BaseAccount{},
				roles:      []AccountRole{Vendor, CertificationCenter, Trustee, NodeAdmin},
				approvals:  []*Grant{},
				rejects:    []*Grant{},
				vendorID:   1,
				productIDs: []*commontypes.Uint16Range{},
			},
			want: &Account{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, CertificationCenter, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				Rejects:     []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{},
			},
		},
		{
			name: "invalid account vendor role",
			args: args{
				ba:         &types.BaseAccount{},
				roles:      []AccountRole{Vendor},
				approvals:  []*Grant{},
				rejects:    []*Grant{},
				vendorID:   2,
				productIDs: []*commontypes.Uint16Range{},
			},
			want: &Account{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				Rejects:     []*Grant{},
				VendorID:    2,
				ProductIDs:  []*commontypes.Uint16Range{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.ba, tt.args.roles, tt.args.approvals, tt.args.rejects, tt.args.vendorID, []*commontypes.Uint16Range{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid account with vendor ID",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			wantErr: false,
		},
		{
			name: "valid account with certification center role",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{CertificationCenter},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			wantErr: false,
		},
		{
			name: "invalid vendor account with missing vendor ID",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if err := acc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_GetRoles(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   []AccountRole
	}{
		{
			name: "account with Vendor and Trustee roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			want: []AccountRole{Vendor, Trustee},
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			want: []AccountRole{Vendor, Trustee, NodeAdmin},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.GetRoles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetVendorID(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "account with vendor ID 45",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals:   []*Grant{},
				VendorID:    45,
			},
			want: 45,
		},
		{
			name: "account with vendor ID 0",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Trustee},
				Approvals:   []*Grant{},
				VendorID:    0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.GetVendorID(); got != tt.want {
				t.Errorf("Account.GetVendorID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_HasRole(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	type args struct {
		targetRole AccountRole
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "account with Vendor and Trustee roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: Vendor,
			},
			want: true,
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: NodeAdmin,
			},
			want: true,
		},
		{
			name: "account with Vendor, Trustee and NodeAdmin roles",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
			},
			args: args{
				targetRole: CertificationCenter,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.HasRole(tt.args.targetRole); got != tt.want {
				t.Errorf("Account.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_HasRightsToChange(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
		ProductIDs  []*commontypes.Uint16Range
	}
	type args struct {
		pid int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Account with associated ProductIDs: [1-100], wants to modify product with pid=1",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee},
				Approvals:   []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{{Min: 1, Max: 100}},
			},
			args: args{
				pid: 1,
			},
			want: true,
		},
		{
			name: "Account with associated ProductIDs: [1-100,200-300], wants to modify product with pid=300",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{{Min: 1, Max: 100}, {Min: 200, Max: 300}},
			},
			args: args{
				pid: 300,
			},
			want: true,
		},
		{
			name: "Account with associated ProductIDs: [100-100], wants to modify product with pid=100",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{{Min: 100, Max: 100}},
			},
			args: args{
				pid: 100,
			},
			want: true,
		},
		{
			name: "Account without associated ProductIDs: [1-100], wants to modify product with pid=101",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{{Min: 1, Max: 100}},
			},
			args: args{
				pid: 101,
			},
			want: false,
		},
		{
			name: "Account without associated ProductIDs: [100-100], wants to modify product with pid=99",
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor, Trustee, NodeAdmin},
				Approvals:   []*Grant{},
				VendorID:    1,
				ProductIDs:  []*commontypes.Uint16Range{{Min: 100, Max: 100}},
			},
			args: args{
				pid: 99,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
				ProductIDs:  tt.fields.ProductIDs,
			}
			if got := acc.HasRightsToChange(tt.args.pid); got != tt.want {
				t.Errorf("Account.HasRightsToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPendingAccount_HasApprovalFrom(t *testing.T) {
	type fields struct {
		Account *Account
	}
	type args struct {
		address sdk.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := PendingAccount{
				Account: tt.fields.Account,
			}
			if got := acc.HasApprovalFrom(tt.args.address); got != tt.want {
				t.Errorf("PendingAccount.HasApprovalFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_GetApprovals(t *testing.T) {
	type fields struct {
		BaseAccount *types.BaseAccount
		Roles       []AccountRole
		Approvals   []*Grant
		VendorID    int32
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Grant
	}{
		{
			name: `account having 2 approvals`,
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals: []*Grant{
					{
						Address: "address1",
						Info:    "info1",
						Time:    123,
					},
					{
						Address: "address2",
						Info:    "",
						Time:    456,
					},
				},
				VendorID: 1,
			},
			want: []*Grant{
				{
					Address: "address1",
					Info:    "info1",
					Time:    123,
				},
				{
					Address: "address2",
					Info:    "",
					Time:    456,
				},
			},
		},
		{
			name: `account having 1 approval`,
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				Approvals: []*Grant{
					{
						Address: "address1",
						Info:    "info1",
						Time:    123,
					},
				},
				VendorID: 1,
			},
			want: []*Grant{
				{
					Address: "address1",
					Info:    "info1",
					Time:    123,
				},
			},
		},
		{
			name: `account having 0 approvals`,
			fields: fields{
				BaseAccount: &types.BaseAccount{},
				Roles:       []AccountRole{Vendor},
				VendorID:    1,
				Approvals:   []*Grant{},
			},
			want: []*Grant{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := Account{
				BaseAccount: tt.fields.BaseAccount,
				Roles:       tt.fields.Roles,
				Approvals:   tt.fields.Approvals,
				VendorID:    tt.fields.VendorID,
			}
			if got := acc.GetApprovals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetApprovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
