//Package dbmanager Provides a postgres database manager using sqlx.

//It uses maps to create manage an save data in postgres database

package dbmanage

import (
	"reflect"
	"testing"
)

func TestDBManager_TableExists(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.database.TableExists(tt.args.tableName); got != tt.want {
				t.Errorf("DBManager.TableExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_columnFromTableExists(t *testing.T) {
	type args struct {
		tableName  string
		columnName string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.database.columnFromTableExists(tt.args.tableName, tt.args.columnName); got != tt.want {
				t.Errorf("DBManager.columnFromTableExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_getAllTableNames(t *testing.T) {
	tests := []struct {
		name     string
		database *DBManager
		want     []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.database.getAllTableNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.getAllTableNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_Connect(t *testing.T) {
	type args struct {
		dbName     string
		dbUser     string
		dbPassword string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.Connect(tt.args.dbName, tt.args.dbUser, tt.args.dbPassword); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_ConnectURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.ConnectURL(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.ConnectURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_addToSchemaManager(t *testing.T) {
	type args struct {
		tableName string
		fields    Value
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.database.addToSchemaManager(tt.args.tableName, tt.args.fields)
		})
	}
}

func TestDBManager_CreateTable(t *testing.T) {
	type args struct {
		tableName string
		fields    []Field
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.CreateTable(tt.args.tableName, tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_GetAllRows(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     []Value
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.database.GetAllRows(tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.GetAllRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.GetAllRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_FilerRowsBy(t *testing.T) {
	type args struct {
		tableName string
		filterBy  Value
		orderBy   []string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     []Value
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.database.FilerRowsBy(tt.args.tableName, tt.args.filterBy, tt.args.orderBy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.FilerRowsBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.FilerRowsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_SearchRowsBy(t *testing.T) {
	type args struct {
		tableName string
		filterBy  Value
		orderBy   []string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     []Value
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.database.SearchRowsBy(tt.args.tableName, tt.args.filterBy, tt.args.orderBy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.SearchRowsBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.SearchRowsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_DeleteRowBy(t *testing.T) {
	type args struct {
		tableName string
		element   Value
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.DeleteRowBy(tt.args.tableName, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.DeleteRowBy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_DeleteAllRows(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.DeleteAllRows(tt.args.tableName); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.DeleteAllRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_InsertElement(t *testing.T) {
	type args struct {
		tableName string
		element   Value
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.InsertElement(tt.args.tableName, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.InsertElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_DropTable(t *testing.T) {
	type args struct {
		tableName string
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.DropTable(tt.args.tableName); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.DropTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_UpdateRowBy(t *testing.T) {
	type args struct {
		tableName string
		filter    Value
		elem      Value
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.database.UpdateRowBy(tt.args.tableName, tt.args.filter, tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.UpdateRowBy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_PagedQuery(t *testing.T) {
	type args struct {
		tableName string
		pageInfo  Value
	}
	tests := []struct {
		name     string
		database *DBManager
		args     args
		want     []Value
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.database.PagedQuery(tt.args.tableName, tt.args.pageInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.PagedQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.PagedQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
