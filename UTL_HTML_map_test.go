package UTL_HTML
// -------------------------------------------------------------------------------------------------

//
// -------------------------------------------------------------------------------------------------

import (
  "os"
  "fmt"
  "time"
  "testing"
)

/* */

// ***********************************************
// Testing table creation from a map[string]string
// ***********************************************
var Environment = map[string]string{
	"ALLUSERSPROFILE"        : `C:\ProgramData`,
  "COMMONPROGRAMFILES"     : `C:\Program Files (x86)\Common Files`,
  "CommonProgramFiles(x86)": `C:\Program Files (x86)\Common Files`,
  "CommonProgramW6432"     : `C:\Program Files\Common Files`,
  "COMSPEC"                : `C:\Windows\system32\cmd.exe`,
  "DISPLAY"                : `127.0.0.1:0.0`,
  "DriverData"             : `C:\Windows\System32\Drivers\DriverData`,
  "PROCESSOR_ARCHITECTURE" : `x86`,
  "PROCESSOR_ARCHITEW6432" : `AMD64`,
  "ProgramData"            : `C:\ProgramData`,
  "PROGRAMFILES"           : `C:\Program Files (x86)`,
  "ProgramFiles(x86)"      : `C:\Program Files (x86)`,
  "ProgramW6432"           : `C:\Program Files`,
}
func Test_Table_MapStrStr(t *testing.T) {
  FileName := "Test_Table_Map[string]string.html"
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","1103").
               Base("/test/","_self").
               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
               Link("stylesheet","text/css","static/Default.css").
               Link("stylesheet","text/css","static/Bookmarks.css").
               Title("TableTest").
             TagCloseUntil("head").
             BodyOpen().
               TableOpen("class","class4table").
                 Captionf("class4caption","Table from Map[string]*string performed at %s.",time.Now().Format("2006-01-02T15:04:05")).
                 TheadOpen().
                   TrTh("class4tr","class4th","Name","Value").
                 TagCloseTop().
                 TbodyOpen().
                   TrTdMap("class4tr","class4td",CmpAsc,Environment).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_MapStrStr

/* */


// *********************************************
// Testing table creation from a map[string]*any
// *********************************************
var Parameter = map[string]*any{
	"DB_BLOCK_SIZE"          : int2ptrAny(8192),
  "PROCESSES"              : int2ptrAny(192),
  "DB_NAME"                : string2ptrAny("example_database"),
  "NLS_DATE_FORMAT"        : string2ptrAny("YYYY-MM-DD"),
  "PLSQL_DEBUG"            : bool2ptrAny(true),
  "DO_WHAT_CUSTOMER_WANTS" : bool2ptrAny(true),
  "RESOURCE_LIMIT"         : bool2ptrAny(false),
}
func Test_Table_MapStrPtrAny(t *testing.T) {
  FileName := "Test_Table_Map[string]PtrAny.html"
  Caption  := "Table from Map[string]*any performed at %s."
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","1103").
               Base("/test/","_self").
               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
               Link("stylesheet","text/css","static/Default.css").
               Link("stylesheet","text/css","static/Bookmarks.css").
               Title("TableTest").
             TagCloseUntil("head").
             BodyOpen().
               TableOpen("class","class4table").
                 Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
                 TheadOpen().
                   TrTh("class4tr","class4th","Name","Value").
                 TagCloseTop().
                 TbodyOpen().
                   TrTdMap("class4tr","class4td",CmpAsc,Parameter).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_MapStrPtrAny

/* */

// *************************************************
// Testing table creation from a map[string]struct{}
// *************************************************
type t_TickerSymbol struct {
  Name       string   `html:"ColHeader='Fullname' HeaderClass='AdditionalHeaderClass'"`
  Type       string   `html:"DataClass='AdditionalDataClass' Style='text-align: center;'"`
  LastPrice  float32  `html:"Style='text-align: right; color: green;'"`
  LastVolume *any     `html:"Style='text-align: right; color: blue;'"`
} // END t_TickerSymbol
var TickerTable = map[string]t_TickerSymbol{
	"XOMO"  : {Name: "YieldMax XOM Option Income Strategy ETF", Type: "ETF", LastPrice: 14.48, LastVolume: int2ptrAny(46165)},
  "EURUSD": {Name: "Euro US Dollar", Type: "Currency", LastPrice: 1.0813},
  "TB91D" : {Name: "91 Day Treasury Bill", Type: "FixedIncome", LastPrice: 4.1900},
  "GSHRU" : {Name: "Gesher Acquisition Corp. II", Type: "IPO", LastPrice: 10.02,  LastVolume: int2ptrAny(6324701)},
  "EACAX" : {Name: "Eaton Vance California Municipal Opportunities Fund Class A", Type: "Mutual Fund", LastPrice: 10.99},
  "CNP"   : {Name: "CenterPoint Energy, Inc (Holding Co) Common Stock", Type: "Stock", LastPrice: 35.77, LastVolume: int2ptrAny(8615303)},
  "SPX"   : {Name: "S&P 500", Type: "Index", LastPrice: 5667.5600},
}
func Test_Table_MapStrStruct(t *testing.T) {
  FileName := "Test_Table_Map[string]struct.html"
  Caption  := "Table from Map[string]struct{} performed at %s."
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","1103").
               Base("/test/","_self").
               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
               Link("stylesheet","text/css","static/Default.css").
               Link("stylesheet","text/css","static/Bookmarks.css").
               Title("TableTest").
             TagCloseUntil("head").
             BodyOpen().
               TableOpen("class","class4table").
                 Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
                 TheadOpen().
                   TrThStruct("class4tr","class4th","Symbol",t_TickerSymbol{}).
                 TagCloseTop().
                 TbodyOpen().
                   TrTdMap("class4tr","class4td",CmpAsc,TickerTable).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_MapStrStruct


/* */

// **************************************************
// Testing table creation from a map[string]*struct{}
// **************************************************
type t_Employee struct {
  FirstName     string
  LastName      string
  Salary        int     `html:"Style='text-align: right;'"`
  PctCommission float32 `html:"Style='text-align: right;'"`
  Retired       bool    `html:"Skip"`
} // END t_Employee
var EmployeeMap = map[int32]*t_Employee{
	1: &t_Employee{FirstName: "John",    LastName: "Doe",      Salary: 1234, PctCommission: 1.2, Retired: false},
  2: &t_Employee{FirstName: "Jane",    LastName: "Smith",    Salary: 2345, PctCommission: 2.3, Retired: false},
  3: &t_Employee{FirstName: "Bob",     LastName: "Johnson",  Salary: 3456, PctCommission: 3.4, Retired: false},
  4: &t_Employee{FirstName: "Alice",   LastName: "Williams", Salary: 4567, PctCommission: 4.5, Retired: false},
  5: &t_Employee{FirstName: "Charlie", LastName: "Brown",    Salary: 5678, PctCommission: 5.6, Retired: false},
  6: &t_Employee{FirstName: "Wile",    LastName: "Coyote",   Salary: 0,    PctCommission:   0, Retired: true},
} // END EmployeeMap
func Test_Table_MapIntPtrStruct(t *testing.T) {
  FileName := "Test_Table_Map[int]PtrStruct.html"
  Caption  := "Table from Map[int]*struct{} performed at %s."
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","1103").
               Base("/test/","_self").
               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
               Link("stylesheet","text/css","static/Default.css").
               Link("stylesheet","text/css","static/Bookmarks.css").
               Title("TableTest").
             TagCloseUntil("head").
             BodyOpen().
               TableOpen("class","class4table").
                 Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
                 TheadOpen().
                   TrThStruct("class4tr","class4th","Employee Number",t_Employee{}).
                 TagCloseTop().
                 TbodyOpen().
                   TrTdMap("class4tr","class4td",CmpDesc,EmployeeMap).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_MapStrPrtStruct

/* */

// *****************************************************
// Testing table creation from a map[string][]SimpleType
// *****************************************************
var EmployeeTable = map[int32][]any{
	1: {string2ptrAny("John"),    string2ptrAny("Doe"),      int2any(1234), float2any(1.2), bool2any(false)},
  2: {string2ptrAny("Jane"),    string2ptrAny("Smith"),    int2any(2345), float2any(2.3), bool2any(false)},
  4: {string2ptrAny("Alice"),   string2ptrAny("Williams"), int2any(4567), float2any(4.5), bool2any(false)},
  5: {string2ptrAny("Charlie"), string2ptrAny("Brown"),    int2any(5678), float2any(5.6), bool2any(false)},
  3: {string2ptrAny("Bob"),     string2ptrAny("Johnson"),  int2any(3456), float2any(3.4), bool2any(false)},
  6: {string2ptrAny("Wile"),    string2ptrAny("Coyote"),   int2any(0),    float2any(0),   bool2any(true)},
} // END EmployeeTable
func Test_Table_MapIntSlice(t *testing.T) {
  FileName := "Test_Table_Map[int]Slice.html"
  Caption  := "Table from Map[int]Slice performed at %s."
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","1103").
               Base("/test/","_self").
               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
               Link("stylesheet","text/css","static/Default.css").
               Link("stylesheet","text/css","static/Bookmarks.css").
               Title("TableTest").
             TagCloseUntil("head").
             BodyOpen().
               TableOpen("class","class4table").
                 Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
                 TheadOpen().
                   TrTh("class4tr","class4th","EmpNum","Firstname","Lastname","Salary","Commission","Retired").
                 TagCloseTop().
                 TbodyOpen().
                   TrTdMap("class4tr","class4td",nil,EmployeeTable).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_MapIntSlice

/* */
