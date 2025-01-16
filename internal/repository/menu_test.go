package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_menuRepository_GetMenuByType(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name     string
		args     args
		want     []model.MenuItem
		wantErr  bool
		initMock func() (*sql.DB, sqlmock.Sqlmock, error)
	}{
		{
			name: "Success list menu",
			args: args{
				ctx:      context.Background(),
				menuType: constant.MenuTypeFood,
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()

				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnRows(sqlmock.NewRows([]string{
					"name",
					"order_code",
					"price",
					"type",
				}).AddRow("Hamburger", "HAMBURGER", 10000, constant.MenuTypeFood))

				return db, mock, err
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
			name: "Failed list menu",
			args: args{
				ctx:      context.Background(),
				menuType: constant.MenuTypeFood,
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()

				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnError(errors.New("mock error"))

				return db, mock, err
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, err := tt.initMock()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				DSN:                  "sqlmock_db_0",
				DriverName:           "postgres",
				Conn:                 db,
				PreferSimpleProtocol: true,
			}))

			if err != nil {
				t.Error(err)
			}

			repo := &menuRepository{
				db: gormDB,
			}

			got, err := repo.GetMenuByType(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuRepository.GetMenuByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuRepository.GetMenuByType() = %v, want %v", got, tt.want)
			}
			if err := dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("expectations were not met: %s", err.Error())
			}
		})
	}
}

func Test_menuRepository_GetMenuByOrderCode(t *testing.T) {
	type args struct {
		ctx       context.Context
		orderCode string
	}
	tests := []struct {
		name    string
		r       *menuRepository
		args    args
		want    model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetMenuByOrderCode(tt.args.ctx, tt.args.orderCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuRepository.GetMenuByOrderCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuRepository.GetMenuByOrderCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
