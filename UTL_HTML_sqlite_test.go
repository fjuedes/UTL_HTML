package UTL_HTML
// -------------------------------------------------------------------------------------------------

//
// ------------

/* */
import (
  "os"
  "fmt"
  "time"
  "testing"
	"database/sql"
	_ "modernc.org/sqlite"
)

/* */

// ********************************************
// Testing table creation from a database table
// ********************************************

func TestTable_Sqlite3DuckBreeds(t *testing.T) {
  FileName := "Test_Table_SQLite3DuckBreeds.html"
  Caption  := "Table from Database-Query performed at %s."
	
	dbh,err := sql.Open("sqlite","Ducks.sqlite3")
	if err != nil {
    t.Errorf(err.Error())
		return
	} // END if
	defer dbh.Close()
	
	Rows,err := dbh.Query("SELECT * FROM DuckBreeds ORDER BY 1")
	if err != nil {
    t.Errorf(err.Error())
		return
	} // END if
  defer Rows.Close()	
	
	
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","0815").
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
                   TrThSqlRows("class4tr","class4th",Rows).
                 TagCloseTop().
                 TbodyOpen().
                   TrTdSqlRows("class4tr","class4td",Rows).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END TestTable_Sqlite3DuckBreeds

/* */
