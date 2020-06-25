//Package dbmanager Provides a postgres database manager using sqlx.
//It uses maps to create manage an save data in postgres database
package dbmanager

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

var SCHEMA_MANAGER = "schema_manager_db"

//Field - Struct of a database field, defines name and type of column
type Field struct {
	Name string
	Type string
}

//Value - Value extracted from database, map with colum as keys
type Value map[string]interface{}

//DBManager - Struct containig DB information
type DBManager struct {
	db           *sqlx.DB
	tableNames   []string
	tableSchemas map[string][]Field
	connected    bool
}

//TableExists - Check if table exists
//PARAMS:
//	tableName: name of the table to search for
//		returns: True if the table exists, False if not
func (database *DBManager) TableExists(tableName string) bool {
	rows, _ := database.db.Queryx(fmt.Sprintf("SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = '%s'", tableName))

	for rows.Next() {
		rows.Close()
		return true

	}
	return false
}

func (database *DBManager) columnFromTableExists(tableName string, columnName string) bool {
	colString := fmt.Sprintf("SELECT * FROM information_schema.columns "+
		"WHERE table_name='%s' AND column_name='%s';", tableName, columnName)
	//log.Println(colString)
	rows, _ := database.db.Queryx(colString)
	for rows.Next() {
		rows.Close()
		return true

	}
	return false
}

//getAllTableNames - Returns the name of all the tables managed
func (database *DBManager) getAllTableNames() []string {
	elem, _ := database.GetAllRows(SCHEMA_MANAGER)

	names := make([]string, len(elem))
	for i, el := range elem {
		names[i] = fmt.Sprint(el["tablename"])
	}
	return names
}

//Connect - connect to the database specified
//	PARAMS:
//	dbName: name of the database to connect to
//	dbUser: username of the postgres user
//	dbPassword: password of the postgres user
//		returns: error if any
func (database *DBManager) Connect(dbName string, dbUser string, dbPassword string) error {

	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	var err error
	database.db, err = sqlx.Connect("postgres", connString)

	if err != nil {
		return err
	}

	database.connected = true
	database.tableNames = []string{}
	database.tableSchemas = map[string][]Field{}

	return nil
}

//ConnectURL - connect to the database specified
//	PARAMS:
//	dbName: name of the database to connect to
//	dbUser: username of the postgres user
//	dbPassword: password of the postgres user
//		returns: error if any
func (database *DBManager) ConnectURL(url string) error {
	var err error

	database.db, err = sqlx.Open("postgres", url)

	if err != nil {
		return err
	}

	database.connected = true
	database.tableNames = []string{}
	database.tableSchemas = map[string][]Field{}
	defer database.db.Close()

	err = database.db.Ping()
	if err != nil {
		panic(err)
	}
	return nil
}

func (database *DBManager) addToSchemaManager(tableName string, fields Value) {
	tableName = strings.ToLower(tableName)
	if !database.TableExists(SCHEMA_MANAGER) {
		tableSchema := fmt.Sprintf("CREATE TABLE %s (table_name text, table_schema text)", SCHEMA_MANAGER)
		log.Println(tableSchema)
		database.db.MustExec(tableSchema)
	}
	database.InsertElement(SCHEMA_MANAGER, fields)
}

//CreateTable - creates a table with provided name and fields
//	PARAMS:
//	tableName: name of the table to be created
//	fields: slice of fields with field name and type
//		returns: error if any
func (database *DBManager) CreateTable(tableName string, fields []Field) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if database.TableExists(tableName) {
		log.Println("Table already exists!!!")
		return errors.New("DBManager - Table already exists!")
	}

	fieldString := ""

	for idx, field := range fields {
		if idx < len(fields)-1 {
			fieldString += strings.ToLower(field.Name) + " " + field.Type + ","
		} else {
			fieldString += strings.ToLower(field.Name) + " " + field.Type
		}
	}

	tableSchema := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, fieldString)
	log.Print(tableSchema)
	database.db.MustExec(tableSchema)

	database.tableNames = append(database.tableNames, tableName)

	database.tableSchemas[tableName] = fields

	database.addToSchemaManager(tableName, Value{"table_name": tableName, "table_schema": fieldString})

	return nil
}

//GetAllRows - returns all the rows of the table
//	PARAMS:
//	tableName: name of the table to be returned
//		returns: slice of Value, error if any
func (database *DBManager) GetAllRows(tableName string) ([]Value, error) {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return nil, errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return nil, errors.New("DBManager - Table doesn't exist!")
	}

	finalMap := []Value{}
	query := "SELECT * FROM " + tableName
	log.Println(query)
	result := make(map[string]interface{})
	rows, err := database.db.Queryx(query)

	if err != nil {
		return nil, errors.New("DBManager - Couldn't execute query")
	}

	for rows.Next() {
		err := rows.MapScan(result)
		if err != nil {
			return nil, errors.New("DBManager - Couldn't map values")
		}

		temp := make(Value)
		for k, v := range result {
			temp[k] = fmt.Sprint(v)
		}
		finalMap = append(finalMap, temp)
	}

	return finalMap, nil
}

//FilerRowsBy - Search by column containing exactly each column value
//	PARAMS:
//	tableName: name of the table where to search
//	filterBy:  a Value(key-value) map of the column to search by
//	orderBy: optional string of column name to sort by
// 		returns: Value slice of each row, error if any
func (database *DBManager) FilerRowsBy(tableName string, filterBy Value, orderBy ...string) ([]Value, error) {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return nil, errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return nil, errors.New("DBManager - Table doesn't exist!")
	}

	finalMap := []Value{}

	filters := ""
	i := 0
	for k, v := range filterBy {
		filters += k + "='" + fmt.Sprint(v) + "'"
		if i < len(filterBy)-1 {
			filters += " AND "
		}
		i++
	}

	query := "SELECT * FROM " + tableName

	if len(filterBy) > 0 {
		query += " WHERE " + filters
	}

	if len(orderBy) > 0 {
		query += " ORDER BY " + orderBy[0] + " ASC"
	}

	log.Println(query)

	result := make(map[string]interface{})
	rows, err := database.db.Queryx(query)
	if err != nil {
		return nil, errors.New("DBManager - Couldn't execute query")
	}
	for rows.Next() {
		err := rows.MapScan(result)
		if err != nil {
			return nil, errors.New("DBManager - Couldn't execute query")
		}

		temp := make(Value)
		for k, v := range result {
			temp[k] = fmt.Sprint(v)
		}
		finalMap = append(finalMap, temp)
	}

	return finalMap, nil
}

//SearchRowsBy - Search by column containing part of the string
//	PARAMS:
//	tableName: name of the table where to search
//	filterBy:  a Value(key-value) map of the column to search by
//	orderBy: optional string of column name to sort by
// 		returns: Value slice of each row, error if any
func (database *DBManager) SearchRowsBy(tableName string, filterBy Value, orderBy ...string) ([]Value, error) {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return nil, errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return nil, errors.New("DBManager - Table doesn't exist!")
	}

	finalMap := []Value{}

	filters := ""
	i := 0
	for k, v := range filterBy {
		filters += k + "::text LIKE '%" + fmt.Sprint(v) + "%'"
		if i < len(filterBy)-1 {
			filters += " AND "
		}
		i++
	}

	query := "SELECT * FROM " + tableName

	if len(filterBy) > 0 {
		query += " WHERE " + filters
	}

	if len(orderBy) > 0 {
		query += " ORDER BY " + orderBy[0] + " ASC"
	}

	log.Println(query)

	result := make(map[string]interface{})
	rows, err := database.db.Queryx(query)
	if err != nil {
		return nil, errors.New("DBManager - Couldn't execute query")
	}
	for rows.Next() {
		err := rows.MapScan(result)
		if err != nil {
			return nil, errors.New("DBManager - Couldn't execute query")
		}

		temp := make(Value)
		for k, v := range result {
			temp[k] = fmt.Sprint(v)
		}
		finalMap = append(finalMap, temp)
	}

	return finalMap, nil
}

//DeleteRowBy - Delete element from table
//	PARAMS:
//	tableName: name of the table where to delete
//	element:  a Value(key-value) map of the column(s) to search by
// 		returns: error if any
func (database *DBManager) DeleteRowBy(tableName string, element Value) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	if len(element) <= 0 {
		return errors.New("DBManager - No filter provided")
	}
	filters := ""
	i := 0
	for name, value := range element {

		filters += name + "='" + fmt.Sprint(value) + "'"

		if i < len(element)-1 {
			filters += " AND "
		}
		i++
	}

	deleteString := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, filters)
	log.Println(deleteString)
	database.db.MustExec(deleteString)
	return nil
}

//DeleteAllRows - deletes all the elements from a table
//	PARAMS:
//	tableName: name of the table where to delete
// 		returns: error if any
func (database *DBManager) DeleteAllRows(tableName string) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	deleteString := fmt.Sprintf("DELETE FROM %s ", tableName)
	log.Println(deleteString)
	database.db.MustExec(deleteString)
	return nil
}

//InsertElement - inserts element in a table
//	PARAMS:
//	tablename: name of the table where to instert
//	element: Value(Key-Value map) of the element to insert
//		returns: error if any
func (database *DBManager) InsertElement(tableName string, element Value) error {
	tableName = strings.ToLower(tableName)

	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	variableString := ""
	variableNumber := ""
	variableValue := make([]interface{}, len(element))
	idx := 1
	for name, value := range element {
		variableString += name + ","
		variableNumber += "$" + strconv.Itoa(idx) + ","
		variableValue[idx-1] = value
		idx++
	}
	variableString = variableString[:len(variableString)-1]
	variableNumber = variableNumber[:len(variableNumber)-1]

	insertString := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, variableString, variableNumber)
	log.Println(insertString)

	database.db.MustExec(insertString, variableValue...)

	return nil
}

//DropTable - Deletes table from system
//	PARAMS:
//	tableName: name of the table to search for
//		returns: error if any
func (database *DBManager) DropTable(tableName string) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	database.DeleteRowBy(SCHEMA_MANAGER, Value{"table_name": tableName})
	deleteString := fmt.Sprintf("DROP TABLE %s ", tableName)

	log.Println(deleteString)
	database.db.MustExec(deleteString)
	return nil
}

//UpdateRowBy - update all matching rows from a certain table
//PARAMS:
//	tablename: name of the table to update
//	filter: filter to search row to change
//	elem: map of the column(s) to update
//		returns: error if any
func (database *DBManager) UpdateRowBy(tableName string, filter Value, elem Value) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.TableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}
	j := 0
	filterFields := ""
	for colName, colValue := range filter {
		if !database.columnFromTableExists(tableName, colName) {
			log.Println("Column doesn't exist!!!")
			return errors.New("DBManager - Column doesn't exist!")
		}
		if j < len(filter)-1 {
			filterFields += colName + "='" + fmt.Sprint(colValue) + "' AND "
		} else {
			filterFields += colName + "='" + fmt.Sprint(colValue) + "'"
		}
		j++
	}
	i := 0
	updateFields := ""
	for colName, colValue := range elem {
		if !database.columnFromTableExists(tableName, colName) {
			log.Println("Column doesn't exist!!!")
			return errors.New("DBManager - Column doesn't exist!")
		}
		fmt.Println(colValue)
		if i < len(elem)-1 {
			updateFields += colName + "='" + fmt.Sprint(colValue) + "', "
		} else {
			updateFields += colName + "='" + fmt.Sprint(colValue) + "'"
		}
		i++
	}

	updateString := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", tableName, updateFields, filterFields)
	log.Print(updateString)
	database.db.MustExec(updateString)
	return nil
}
