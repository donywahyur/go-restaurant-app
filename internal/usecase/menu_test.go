package usecase

import (
	"context"
	"errors"
	"go-restaurant-app/internal/mocks"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_menuUsecase_GetMenuByType(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		u       *menuUsecase
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		{
			name: "success get menu",
			u: func() *menuUsecase {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockMenuRepository(ctrl)

				mock.EXPECT().GetMenuByType(gomock.Any(), string(constant.MenuTypeFood)).
					Times(1).Return([]model.MenuItem{
					{
						Name:      "Hamburger",
						OrderCode: "HAMBURGER",
						Price:     10000,
						Type:      constant.MenuTypeFood,
					},
				}, nil)

				return NewMenuUsecase(mock)
			}(),
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want: []model.MenuItem{
				{
					Name:      "Hamburger",
					OrderCode: "HAMBURGER",
					Price:     10000,
					Type:      constant.MenuTypeFood,
				},
			},
			wantErr: false,
		},
		{
			name: "Failed to get menu",
			u: func() *menuUsecase {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockMenuRepository(ctrl)

				mock.EXPECT().GetMenuByType(gomock.Any(), "").
					Times(1).
					Return(nil, errors.New("failed to get menu"))

				return NewMenuUsecase(mock)
			}(),
			args: args{
				ctx:      context.Background(),
				menuType: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetMenuByType(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuUsecase.GetMenuByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuUsecase.GetMenuByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_menuUsecase_GetMenuByOrderCode(t *testing.T) {
	type args struct {
		ctx       context.Context
		orderCode string
	}
	tests := []struct {
		name    string
		u       *menuUsecase
		args    args
		want    model.MenuItem
		wantErr bool
	}{
		{
			name: "Success get menu by code",
			u: func() *menuUsecase {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockMenuRepository(ctrl)

				mock.EXPECT().GetMenuByOrderCode(gomock.Any(), "HAMBURGER").
					Times(1).
					Return(model.MenuItem{

						Name:      "Hamburger",
						OrderCode: "HAMBURGER",
						Price:     10000,
						Type:      constant.MenuTypeFood,
					}, nil)

				return NewMenuUsecase(mock)
			}(),
			args: args{
				ctx:       context.Background(),
				orderCode: "HAMBURGER",
			},
			want: model.MenuItem{
				Name:      "Hamburger",
				OrderCode: "HAMBURGER",
				Price:     10000,
				Type:      constant.MenuTypeFood,
			},
			wantErr: false,
		},
		{
			name: "Failed get menu by code",
			u: func() *menuUsecase {
				ctrl := gomock.NewController(t)
				mock := mocks.NewMockMenuRepository(ctrl)

				mock.EXPECT().GetMenuByOrderCode(gomock.Any(), "HAMBURGER2").
					Times(1).
					Return(model.MenuItem{}, errors.New("failed to get menu"))

				return NewMenuUsecase(mock)
			}(),
			args: args{
				ctx:       context.Background(),
				orderCode: "HAMBURGER2",
			},
			want:    model.MenuItem{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetMenuByOrderCode(tt.args.ctx, tt.args.orderCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuUsecase.GetMenuByOrderCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuUsecase.GetMenuByOrderCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
