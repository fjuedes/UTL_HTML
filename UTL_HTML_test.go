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

// Some little helper functions
func bool2any(b bool)any { return any(b) }
func bool2ptrAny(b bool)*any { a := any(b); return &a }
func float2any(f float32)any { return any(f) }
func float2ptrAny(f float32)*any { a := any(f); return &a }
func int2any(i int)any { return any(i) }
func int2ptrAny(i int)*any { a := any(i); return &a }
func string2any(s string)any { return any(s) }
func string2ptrAny(s string)*any { a := any(s); return &a }


/* */

// *******************************
// ***** Hello World in HTML *****
// *******************************
func Test_HelloWorld(t *testing.T) {
  FileName := "Test_HelloWorld.html"
	
  v_Doc := New(GC_DocTypeHTML5,0x07).
           HtmlOpen().
             HeadOpen().
               Title("Hello World in HTML").
             TagCloseTop().
             BodyOpen().
               Pf("","%s %s",I("Hello"),B("World!")).
           TagCloseAll()
					 
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_HelloWorld

/* */

// ***************************************
// ***** Example for an XML-Document *****
// ***************************************
func Test_XML(t *testing.T) {
  FileName := "Test_XML.svg"

  v_Doc := New(GC_DocTypeNONE,0x07).
	         AppendString(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>`).NL().
					 AppendString(`<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`).NL().
					 TagOpen("svg","width","512", "height","512", "viewBox","-70.5 -70.5 391 391", "xmlns","http://www.w3.org/2000/svg", "xmlns:xlink","http://www.w3.org/1999/xlink").
					   Tag("rect","","fill","#fff","stroke","#000","x","-70","y","-70","width","390","height","390").
						 TagOpen("g","opacity","0.8").
						   Tag("rect","","x","25","y","25","width","200","height","200","fill","lime","stroke-width","4","stroke","magenta").
							 Tag("circle","","cx","125","cy","125","r","75","fill","orange").
							 Tag("polyline","","points","50,150 50,200 200,200 200,100","stroke","red","stroke-width","4","fill","none").
							 Tag("line","", "x1","50", "y1","50" ,"x2","200" ,"y2","200", "stroke","blue", "stroke-width","4").
           TagCloseAll()

  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_HelloWorld

/* */

// **************************************************
// Testing Static Table generation with TrTh and TrTd
// **************************************************
func Test_Table_Static(t *testing.T) {
  FileName := "Test_Table_Static.html"
  Caption  := "Static Table Test performed at %s."
  FranksAge := 61
  HeaderCol1 := "Name"
  v_Doc := New(GC_DocTypeHTML5,0x02).
           HtmlOpen().
             HeadOpen("id","4711").
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
                   TrTh("class4tr","class4th",&HeaderCol1,"Age","City","Occupation","Alien").
                 TagCloseTop().
                 TbodyOpen().
                   TrTd("class4tr","class4td","Alice",28,"Old York","Software Engineer",false).
                   TrTd("class4tr","class4td","Bob",42,"Lost Angeles","Data Scientist",false).
                   TrTd("class4tr","class4td","Frank",&FranksAge,"Charleston","Unemployed",true).
           TagCloseAll()
  if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
    t.Errorf(e.Error())
  } else {
    fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
  } // END if
} // END Test_Table_Static

/* */
