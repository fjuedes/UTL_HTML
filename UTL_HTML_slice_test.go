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

// *****************************************
// Testing table creation from a []*struct{}
// *****************************************
type t_CityRec struct {
  Name        string  `html:"ColHeader='City Name'"`
  Population  int     `html:"ColHeader='Population in 2023' style='text-align: right'"`
  Latitude    float32 `html:"ColHeader='Latitude' style='text-align: right'"`
  Visited     bool    `html:"colheader='Visited' style='text-align: center'"`
  private     string
} // END t_CityRec

var v_CityList = []*t_CityRec{
	&t_CityRec{Name: "Houston", Population: 2314157, Latitude: 29.762778, Visited: true, private: "hidden"},
  &t_CityRec{Name: "Phoenix", Population: 1650070, Latitude: 33.448333, Visited: false, private: "secret"},
  &t_CityRec{Name: "Charleston", Population: 46838, Latitude: 38.349819, Visited: true, private: "not to be displayed"},
}
// Table from Slice of struct
func TestTable_SlicePtrStruct(t *testing.T) {
  FileName := "Test_Table_[]PtrStruct.html"
  Caption  := "Table from []*struct{} performed at %s."
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
                   TrThStruct("class4tr","class4th","",t_CityRec{}).
                 TagCloseTop().
                 TbodyOpen().
//                   TrTdStruct("class4tr","class4td","",*v_CityList[0]).
                   TrTdSlice("class4tr","class4td",v_CityList).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END TestTable_SlicePtrStruct

/* */

// *****************************************
// Testing table creation from a [][]any
// *****************************************
var v_CityTable = [][]any{
	{"Houston",   2314157, 29.762778,  true},
  {"Phoenix",   1650070, 33.448333, false},
  {"Charleston",  46838, 38.349819,  true},
}
// Table from Slice of slice
func TestTable_SliceSlice(t *testing.T) {
  FileName := "Test_Table_[][]any.html"
  Caption  := "Table from [][]any performed at %s."
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
                   TrThStruct("class4tr","class4th","",t_CityRec{}).
                 TagCloseTop().
                 TbodyOpen().
//                   TrTdSlice("class4tr","class4td",v_CityTable[0]).
                   TrTdSlice("class4tr","class4td",v_CityTable).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END TestTable_SliceSlice

/* */
