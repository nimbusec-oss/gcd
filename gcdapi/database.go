// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Database functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/nimbusec-oss/gcd/gcdmessage"
)

// Database object.
type DatabaseDatabase struct {
	Id      string `json:"id"`      // Database ID.
	Domain  string `json:"domain"`  // Database domain.
	Name    string `json:"name"`    // Database name.
	Version string `json:"version"` // Database version.
}

// Database error.
type DatabaseError struct {
	Message string `json:"message"` // Error message.
	Code    int    `json:"code"`    // Error code.
}

//
type DatabaseAddDatabaseEvent struct {
	Method string `json:"method"`
	Params struct {
		Database *DatabaseDatabase `json:"database"` //
	} `json:"Params,omitempty"`
}

type Database struct {
	target gcdmessage.ChromeTargeter
}

func NewDatabase(target gcdmessage.ChromeTargeter) *Database {
	c := &Database{target: target}
	return c
}

// Enables database tracking, database events will now be delivered to the client.
func (c *Database) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.enable"})
}

// Disables database tracking, prevents database events from being sent to the client.
func (c *Database) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.disable"})
}

// GetDatabaseTableNames -
// databaseId -
// Returns -  tableNames -
func (c *Database) GetDatabaseTableNames(databaseId string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["databaseId"] = databaseId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.getDatabaseTableNames", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TableNames []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.TableNames, nil
}

// ExecuteSQL -
// databaseId -
// query -
// Returns -  columnNames -  values -  sqlError -
func (c *Database) ExecuteSQL(databaseId string, query string) ([]string, []interface{}, *DatabaseError, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["databaseId"] = databaseId
	paramRequest["query"] = query
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.executeSQL", Params: paramRequest})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			ColumnNames []string
			Values      []interface{}
			SqlError    *DatabaseError
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.ColumnNames, chromeData.Result.Values, chromeData.Result.SqlError, nil
}
