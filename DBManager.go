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

type Field struct {
	Name string
	Type string
}
type Value map[string]string

type DBManager struct {
	db           *sqlx.DB
	errManager   error
	tableNames   []string
	tableSchemas map[string][]Field
	connected    bool
}

func (database *DBManager) tableExists(tableName string) bool {
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

func (database *DBManager) GetAllTableNames() []string {
	elem, _ := database.GetAllTableElements(SCHEMA_MANAGER)

	names := make([]string, len(elem))
	for i, el := range elem {
		names[i] = el["tablename"]
	}
	return names
}

func (database *DBManager) Connect(dbType string, dbName string, dbUser string, dbPassword string) error {

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

func (database *DBManager) addToSchemaManager(tableName string, fields Value) {
	tableName = strings.ToLower(tableName)
	if !database.tableExists(SCHEMA_MANAGER) {
		tableSchema := fmt.Sprintf("CREATE TABLE %s (table_name text, table_schema text)", SCHEMA_MANAGER)
		log.Println(tableSchema)
		database.db.MustExec(tableSchema)
	}
	database.InsertElement(SCHEMA_MANAGER, fields)
}

func (database *DBManager) CreateTable(tableName string, fields []Field) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if database.tableExists(tableName) {
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

func (database *DBManager) GetAllTableElements(tableName string) ([]Value, error) {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return nil, errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return nil, errors.New("DBManager - Table doesn't exist!")
	}

	finalMap := []Value{}

	result := make(map[string]interface{})
	rows, err := database.db.Queryx("SELECT * FROM " + tableName)

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
			//fmt.Printf("%s: %s |", k, v)
		}
		//fmt.Printf("%#v\n", result)
		finalMap = append(finalMap, temp)
		rows.Close()
	}

	return finalMap, nil
}

func (database *DBManager) FilerTableElementsBy(tableName string, filterBy Value, orderBy ...string) ([]Value, error) {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return nil, errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return nil, errors.New("DBManager - Table doesn't exist!")
	}

	finalMap := []Value{}

	filters := ""
	i := 0
	for k, v := range filterBy {
		filters += k + "='" + v + "'"
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
			//fmt.Printf("%s: %s |", k, v)
		}
		//fmt.Printf("%#v\n", result)
		finalMap = append(finalMap, temp)
		rows.Close()
	}

	return finalMap, nil
}

func (database *DBManager) DeleteElementFromTable(tableName string, element Value) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	if len(element) <= 0 {
		return errors.New("DBManager - No filter provided")
	}
	filters := ""
	i := 0
	for name, value := range element {

		filters += name + "='" + value + "'"

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

func (database *DBManager) DeleteAllElementsFromTable(tableName string) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	deleteString := fmt.Sprintf("DELETE FROM %s ", tableName)
	log.Println(deleteString)
	database.db.MustExec(deleteString)
	return nil
}

func (database *DBManager) InsertElement(tableName string, element Value) error {
	tableName = strings.ToLower(tableName)

	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
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

	v := Value{}
	v["hello"] = "ala"

	return nil
}

func (database *DBManager) DropTable(tableName string) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
		log.Println("Table doesn't exist!!!")
		return errors.New("DBManager - Table doesn't exist!")
	}

	database.DeleteElementFromTable(SCHEMA_MANAGER, Value{"table_name": tableName})
	deleteString := fmt.Sprintf("DROP TABLE %s ", tableName)

	log.Println(deleteString)
	database.db.MustExec(deleteString)
	return nil
}

func (database *DBManager) updateElementFromTable(tableName string, filter Value, elem Value) error {
	tableName = strings.ToLower(tableName)
	if !database.connected {
		log.Println("Database Not Connected!!!")
		return errors.New("DBManager - Database Not Connected!")
	}

	if !database.tableExists(tableName) {
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
			filterFields += colName + "='" + colValue + "' AND "
		} else {
			filterFields += colName + "='" + colValue + "'"
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
			updateFields += colName + "='" + colValue + "', "
		} else {
			updateFields += colName + "='" + colValue + "'"
		}
		i++
	}

	updateString := fmt.Sprintf("UPDATE %s SET %s WHERE %s;", tableName, updateFields, filterFields)
	log.Print(updateString)
	database.db.MustExec(updateString)
	return nil
}
