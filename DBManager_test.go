package dbmanager

import (
	"reflect"
	"testing"
)

func TestDBManager_tableExists(t *testing.T) {
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
			if got := tt.database.tableExists(tt.args.tableName); got != tt.want {
				t.Errorf("DBManager.tableExists() = %v, want %v", got, tt.want)
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

func TestDBManager_GetAllTableNames(t *testing.T) {
	tests := []struct {
		name     string
		database *DBManager
		want     []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.database.GetAllTableNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.GetAllTableNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_Connect(t *testing.T) {
	type args struct {
		dbType     string
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
			if err := tt.database.Connect(tt.args.dbType, tt.args.dbName, tt.args.dbUser, tt.args.dbPassword); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.Connect() error = %v, wantErr %v", err, tt.wantErr)
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

func TestDBManager_GetAllTableElements(t *testing.T) {
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
			got, err := tt.database.GetAllTableElements(tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.GetAllTableElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.GetAllTableElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_FilerTableElementsBy(t *testing.T) {
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
			got, err := tt.database.FilerTableElementsBy(tt.args.tableName, tt.args.filterBy, tt.args.orderBy...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBManager.FilerTableElementsBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBManager.FilerTableElementsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBManager_DeleteElementFromTable(t *testing.T) {
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
			if err := tt.database.DeleteElementFromTable(tt.args.tableName, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.DeleteElementFromTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBManager_DeleteAllElementsFromTable(t *testing.T) {
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
			if err := tt.database.DeleteAllElementsFromTable(tt.args.tableName); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.DeleteAllElementsFromTable() error = %v, wantErr %v", err, tt.wantErr)
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

func TestDBManager_updateElementFromTable(t *testing.T) {
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
			if err := tt.database.updateElementFromTable(tt.args.tableName, tt.args.filter, tt.args.elem); (err != nil) != tt.wantErr {
				t.Errorf("DBManager.updateElementFromTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
