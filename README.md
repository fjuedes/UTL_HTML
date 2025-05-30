# UTL_HTML - Programmatically generate HTML documents

This go package defines a type T_HTML, with functions/methods to build an HTML document from the ground up. For the most used HTML-tags there is a corresponding function/method with the same name, just with the first letter capitalized. The Div() function/method for example generates complete \<div\> tags with opening and closing tags, content and attributes. Simple HTML-documents can be constructed by calling the tag-methods in the order the tags should appear in the document and using method-chaining the go code will represent the structure of the HTML document rather well.

# Version
```$Id: UTL_HTML.go 79 2025-04-29 20:33:24Z fjuedes $```

# Example, Hello World
```
 func HelloWorld() {
   FileName := "Test_HelloWorld.html"                                                
                                                                                     // Generates:
   v_Doc := New(GC_DocTypeHTML5).                                                    // <!DOCTYPE html>
            HtmlOpen().                                                              // <html>
              HeadOpen().                                                            // <head>
                Title("Hello World in HTML").                                        // <title>Hellow World in HTML</title>
              TagCloseTop().                                                         // </head>
              BodyOpen().                                                            // <body>
                POpen().I("Hello ").B("World!").                                     // <p><i>Hello</i> <b>World!</b></p>
            TagCloseAll()                                                            // </body></html>
            
   if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
     panic(e.Error())
   } else {
     fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
   } // END if
 } // END HelloWorld
```

# Overview
Please find a thematically ordered list of the exported functions and methods of the T_HTML object below. Usually the name of the function/method is the name of the HTML-element, just with the first letter capitalized, for example the \<title\> tag is created through the method Title(). Some tags will never be generated with their complete content, for example the \<body\> tag. So there is no Body() method defined, but a BodyOpen() method which generates the opening tag only. These tags are tracked in an internal tag-stack and will be closed semi-automagically when one of the TagClose...() methods is called. 

Funtions whose name end with the letter "f" work similar to fmt.Printf, accepting a string as a format-mask and a variable number of data-items of any type. For example the B(string) method encapsulates content in \<b\> tags: \<b\>bold text\</b\>. The simliar named method Bf(string, any...) method formats data-items according to the format-mask and encapsulates the resulting string into the \<b\> tag. Example: Bf("Parametervalue is '%d'.",p_Value) will append \<b\>Parametervalue is '4711'.\</b\> to the HTML document if the value of p_Value is the number 4711.

Some functions have been also implemented to return the generated tags as strings which sometimes this is a useful hack… 


# Basic functions
The list below contains the basic functions to handle the T_HTML struct: Creation of a new document-variable, appending strings and NewLine, the Stringer interface and write functions.
 - func New(p_DocType string) *T_HTML
 - func (p_HTML *T_HTML) AS(p_Content string) *T_HTML // shortcut for AppendString
 - func (p_HTML *T_HTML) AppendString(p_Content string) *T_HTML
 - func (p_HTML *T_HTML) AppendStringf(p_Format string, p_Data ... any) *T_HTML
 - func (p_HTML *T_HTML) NL() *T_HTML // Append an HTML-NewLine, always cr/lf regardless of the OS default
 - func (p_HTML *T_HTML) String()string
 - func (p_HTML *T_HTML) Write(w http.ResponseWriter)
 - func (p_HTML *T_HTML) CloseTagsAndWrite(w http.ResponseWriter)


# Functions for generic tags
The list below contains generic functions to build tags for any MarkUp language: From building complete tags with or without content, with or without attributes to opening and closing tags. 
By using these functions any type of MarkUp language document can be built.
 - func (p_HTML *T_HTML) Tag(p_Name, p_Content string, p_Attributes... string) *T_HTML
 - func (p_HTML *T_HTML) Tagf(p_TagName, p_Class, p_Format string, p_Data ... any) *T_HTML
 - func (p_HTML *T_HTML) TagOpen(p_Name string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TagCloseTop() *T_HTML
 - func (p_HTML *T_HTML) TagCloseAll() *T_HTML
 - func (p_HTML *T_HTML) TagCloseUntil(p_Name string) *T_HTML
 - func Tag(p_Name, p_Content string, p_Attributes ...string) string

# Example, SVG document:
```
func Test_XML(t *testing.T) {
  FileName := "Test_XML.svg"

  v_Doc := New(GC_DocTypeNONE).
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
```

# Document structure
Valid HTML documents must contain at least three tags: The <html>tag encloses the entire document. The <head> tag contains information about the text written in the body. The <body> tag contains the information that is displayed by the browser. 
 - func (p_HTML *T_HTML) HtmlOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) BodyOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) HeadOpen(p_Attributes ...string) *T_HTML

# Document header
The document header contains (meta-)information about the HTML-document, defined using the tags <base>, <link>, <meta>, <style> and <title>. - And yes, the <script> tag can still be used in HTML-document headers, but that's so 1990's… just don't do this anymore!
 - func (p_HTML *T_HTML) Base(p_Href, p_target string) *T_HTML 
 - func (p_HTML *T_HTML) Link(p_Rel, p_Type, p_URL string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Meta(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Style(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Title(p_Content string, p_Attributes ...string) *T_HTML 

# Content structuring 
The list below contains basic tags to structure the document-content:
Headers, line breaks, horizontal rulers, paragraphs, divisions and inline-containers. 
 - func (p_HTML *T_HTML) Br(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Div(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) DivOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Header(p_Grade, p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Hr(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) P(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Pf(p_Class, p_Format string, p_Data ... any) *T_HTML
 - func (p_HTML *T_HTML) POpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Span(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Spanf(p_Class, p_Format string, p_Data ... any) *T_HTML {
 - func (p_HTML *T_HTML) SpanOpen(p_Attributes ...string) *T_HTML
 - func Span(p_Content string, p_Attributes ...string) string


# Basic content formatting
Below list contains direct formatting HTML tags in all four flavors, for example: B() and Bf() for the T_HTML Type and B() and Bf() returning a string. The use of these »ancient« formatting tags is controversial as formatting with CSS is more flexible and separates the data (HTML) from the presentation (CSS).
 - func (p_HTML *T_HTML) B(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) Bf(p_Class, p_Format string, p_Data... any) *T_HTML 
 - func (p_HTML *T_HTML) Em(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) Emf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) I(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) If(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) Q(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) Qf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) S(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) Sf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) Strong(p_Content string) *T_HTML
 - func (p_HTML *T_HTML) Strongf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) Sub(p_Content string) *T_HTML
 - func (p_HTML *T_HTML) Subf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) Sup(p_Content string) *T_HTML
 - func (p_HTML *T_HTML) Supf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func (p_HTML *T_HTML) U(p_Content string) *T_HTML 
 - func (p_HTML *T_HTML) Uf(p_Class, p_Format string, p_Data... any) *T_HTML
 - func B(p_Content string) string
 - func Bf(p_Class, p_Format string, p_Data... any) string 
 - func Em(p_Content string) string 
 - func Emf(p_Class, p_Format string, p_Data... any) string
 - func I(p_Content string) string
 - func If(p_Class, p_Format string, p_Data... any) string
 - func Q(p_Content string) string
 - func Qf(p_Class, p_Format string, p_Data... any) string
 - func S(p_Content string) string
 - func Sf(p_Class, p_Format string, p_Data... any) string
 - func Strong(p_Content string) string
 - func Strongf(p_Class, p_Format string, p_Data... any) string
 - func Sub(p_Content string) string
 - func Subf(p_Class, p_Format string, p_Data... any) string
 - func Sup(p_Content string) string
 - func Supf(p_Class, p_Format string, p_Data... any) string
 - func U(p_Content string) string
 - func Uf(p_Class, p_Format string, p_Data... any) string


# Forms [file: UTL_HTML_Form.go]
🚧 Currently under construction 🚧 TODO: More functions for other input-types
 - func (p_HTML *T_HTML) FormOpen(p_action, p_method string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) BoolField(p_name string, p_checked bool, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) HiddenField(p_name, p_value string, p_Attributes... string ) *T_HTML
 - func (p_HTML *T_HTML) SubmitButton(p_name, p_label, p_value, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TextField(p_name, p_size, p_maxlen, p_value, p_Attributes ...string) *T_HTML
 - func BoolField(p_name string, p_checked bool, p_Attributes ...string) string
 - func SubmitButton(p_name, p_label, p_value, p_Attributes ...string) string
 - func TextField(p_name, p_size, p_maxlen, p_value, p_Attributes ...string) string
 - func (p_HTML *T_HTML) SelectMenu(p_FieldName, p_MenuClassName, p_ItemClassName, p_DefaultValue string, p_CompareFunc t_CompareFunc, p_MenuItems any, p_Attributes ...any) *T_HTML

# Lists [File: UTL_HTML_List.go]
🚧 Currently under construction 🚧 TODO: more functions for complex data-types
 - func (p_HTML *T_HTML) OlOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) UlOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) LiOpen(p_Attributes ...string) *T_HTML {
 - func (p_HTML *T_HTML) Li(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Lif(p_Class, p_Format string, p_Data ...any) *T_HTML


# Tables [File: UTL_HTML_Table.go]
Methods to support the creation of HTML-Tables, starting with static tables to the generation of header- and data-rows from structures, maps and slices.
 - func (p_HTML *T_HTML) TableOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Caption(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Captionf(p_Format string, p_data ...any) *T_HTML
 - func (p_HTML *T_HTML) TheadOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TbodyOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TfootOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Th(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Thf(p_Class, p_Content string, p_Data ...any) *T_HTML
 - func (p_HTML *T_HTML) ThOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Td(p_Content string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Tdf(p_Class, p_Content string, p_Data ...any) *T_HTML
 - func (p_HTML *T_HTML) TdOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TrOpen(p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) TrTh(p_TrClass, p_ThClass string, p_DataItems ...any) *T_HTML
 - func (p_HTML *T_HTML) TrTd(p_TrClass, p_TdClass string, p_DataItems ...any) *T_HTML
 - func (p_HTML *T_HTML) TrThStruct(p_TrClass, p_ThClass, p_KeyColHeader string, p_DataItem any) *T_HTML
 - func (p_HTML *T_HTML) TrTdStruct(p_TrClass, p_TdClass, p_KeyColValue string, p_DataItem any) *T_HTML
 - func (p_HTML *T_HTML) TrTdMap(p_TrClass, p_TdClass string, p_CompareFunc t_CompareFunc, p_DataItems any) *T_HTML
 - func (p_HTML *T_HTML) TrTdSlice(p_TrClass, p_TdClass string, p_DataRows any) *T_HTML
 - func (p_HTML *T_HTML) TrThSqlRows(p_TrClass, p_ThClass string, p_DataRows *sql.Rows) *T_HTML
 - func (p_HTML *T_HTML) TrTdSqlRows(p_TrClass, p_TdClass string, p_DataRows *sql.Rows) *T_HTML

Example for a simple static table:
```
 v_Doc := New(GC_DocTypeHTML5).
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
                Captionf("class4caption","Static Table Test performed at %s.",time.Now().Format("2006-01-02T15:04:05")).
                TheadOpen().
                  TrTh("class4tr","class4th","Name","Age","City","Occupation").
                TagCloseTop(). // thead
                TbodyOpen().
                  TrTd("class4tr","class4td","Alice","28","Old York","Software Engineer").
                  TrTd("class4tr","class4td","Bob","42","Lost Angeles","Data Scientist").
                  TrTd("class4tr","class4td","Charlie","61","Charleston","Freelancer").
          TagCloseAll()
 fmt.Println(v_Doc)
```
 
The variable v_Doc contains the whole HTML document, which can be extracted as a string and, printed to an http.ResponseWriter for example.

Beyond just simple static tables there is a group of methods implemented that handle structs, maps and slices. The methods TrThStruct() and TrTdStruct() generate header- and data rows from structures. For example we have a simplified structure for stock symbols (i will explain the html struct-tags later):
```
 type t_TickerSymbol struct {
   Name       string   `html:"ColHeader='Full Name' HeaderClass='AdditionalHeaderClass'"`
   Type       string   `html:"DataClass='AdditionalDataClass'"`
   LastPrice  float32  `html:"Style='text-align: right; color: green;'"`
   LastVolume int      `html:"Style='text-align: right; color: blue;'"`
 } // END t_TickerSymbol 
```

To generate the table-header row the following snippet can be used:
```
 TableOpen("class","class4table").
   Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
   TheadOpen().
     TrThStruct("class4tr","class4th","Symbol",t_TickerSymbol{}).
   TagCloseTop().
   …
```
The method TrThStruct() takes CSS class-names for the <tr> and the <th> tag, one additional column-name (we see later for what that is being used) and an empty t_TickerSymbol to generate a <tr><th>…</th></tr> structure. No loop necessary!

Data-rows are handled by the method TrTDStruct(), for example:
```
   … 
   TbodyOpen().
     TrTdStruct("class4tr","class4td","YHOO",MyTickerSymbol).
   …
```

However that does not make much sense, as this takes care of only a single row of data. You will rarely use the TrTdStruct method directly, but it is used by the methods that handle entire maps and slices. If we have the following data:
```
 var TickerTable = map[string]t_TickerSymbol{
   "XOMO"  : {Name: "YieldMax XOM Option Income Strategy ETF", Type: "ETF", LastPrice: 14.48, LastVolume: 46165},
   "EURUSD": {Name: "Euro US Dollar", Type: "Currency", LastPrice: 1.0813},
   "TB91D" : {Name: "91 Day Treasury Bill", Type: "FixedIncome", LastPrice: 4.1900},
   "GSHRU" : {Name: "Gesher Acquisition Corp. II", Type: "IPO", LastPrice: 10.02,  LastVolume: 6324701},
   "EACAX" : {Name: "Eaton Vance California Municipal Opportunities Fund Class A", Type: "Mutual Fund", LastPrice: 10.99},
   "CNP"   : {Name: "CenterPoint Energy, Inc (Holding Co) Common Stock", Type: "Stock", LastPrice: 35.77, LastVolume: 8615303},
   "SPX"   : {Name: "S&P 500", Type: "Index", LastPrice: 5667.5600},
 }
```
The whole ticker-table can be generated with this code:
```
  …
  TableOpen("class","class4table").
    Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
    TheadOpen().
      TrThStruct("class4tr","class4th","Symbol",t_TickerSymbol{}).
    TagCloseTop().
    TbodyOpen().
      TrTdMap("class4tr","class4td",CmpAsc,TickerTable).
  …
```
See the full example in the file UTL_HTML_test as func Test_Table_MapStrStruct(t *testing.T).

The method TrThStruct accepts any struct-type and non-exported fields of the struct{} are skipped.

The method TrTdStruct accepts any struct-type and non-exported fields of the struct{} are skipped. Pointer fields are dereferenced and nil-values are replaced with "&nbsp;"

The method TrTdMap loops through all the keys of the map - sorted or unsorted - and generates the data-rows. This method can process the following types of maps:
 - map[KeyType]any → Table with two columns: Key, Value
 - map[KeyType]*any → Table with two columns: Key, Value
 - map[KeyType]struct{} → Table with one column for the map-key and one column per exported struct-field
 - map[KeyType]*struct{} → Table with one column for the map-key and one column per exported struct-field
 - map[KeyType][]any → Table with one column for the map-key and one column per slice element
 - map[KeyType][]*any → Table with one column for the map-key and one column per slice element

Keys and values are being converted into strings using fmt.Sprint, composite types will therefore be converted into strings according to their Stringer interface which might not be the desired result.

In a very similar way, the method TrTdSlice() is processes slices into table data-rows. The following slice-types are supported:
 - []any → one row with n columns
 - []*any → one row with n columns
 - []struct → Table with one column per struct-field
 - []*struct → Table with one column per struct-field
 - [][]any → two dimensional table
 - [][]*any → two dimensional table

Again, values are being converted into strings using fmt.Sprint, so composite types will be converted into strings according to their Stringer interface.

# The html struct-tag

Now finally, about the struct-tags in the example t_TickerSymbol structure: When you create your own data-structures, you can add additional information to each field of your structure to control how HTML is generated. The syntax of the html struct-tag:
 - `html:"Skip"`                         - Skip this field
 - `html:"ColHeader='HeaderText'"`       - Use 'Headertext' instead of field-name
 - `html:"HeaderClass='CSS class-name'"` - add class-name to <th> tag
 - `html:"DataClass='CSS class-name'"`   - add class-name to <td> tag
 - `html:"Style="CSS Declaration"`       - add style attribute to <th> or <td> tag

All combinations are supported, for example:
`html:"ColHeader='PrimaryKey'" HeaderClass='Centered' DataClass='Centered' Style='color: red;'`

The keywords "ColHeader", "HeaderClass", "DataClass", "Skip" and "Style" are case insensitive, the tag-name "html" is case-sensitive.

Please use the html struct-tag only as last resort: Remember the the »Separation of Concerns«:

- HTML serves as the structure of the content. It provides the semantic markup for the information being displayed on the web page.
- CSS is used for the presentation or design of the HTML elements. It controls how the elements are displayed in terms of layout, colors, fonts, and other visual aspects. 
- JavaScript is responsible for behavior and/or interactivity of the web page. It enables dynamic content updates, event handling, and user interactions.

Embedding style information into the HTML code violates this design philosophy!


# One more thing.

Usually data is stored in some back-end systems like relationbal Databases, organized in tables. UTL_HTML directly supports SQL rowsets; the following table is defined in the attached SQLite database:
```
 CREATE TABLE "DuckBreeds" (
     "Breed"         TEXT NOT NULL,
     "EggColor"      TEXT NOT NULL,
     "EggSize"       TEXT NOT NULL,
     "EggProduction" TEXT NOT NULL,
     "Class"         TEXT NOT NULL,
     "WeightRange"   TEXT NOT NULL,
     "Mothering"     TEXT NOT NULL,
     "Foraging"      TEXT NOT NULL,
     "Hardiness"     TEXT NOT NULL,
     "Personality"   TEXT NOT NULL,
     "Flying"        INTEGER NOT NULL,
     CONSTRAINT "pk_DuckBreeds" PRIMARY KEY("Breed")
 );
```

Only a few lines of code are necessary to display the content of this table as an HTM document:
```
 // Open Database
 dbh,err := sql.Open("sqlite","Ducks.sqlite3")
 if err != nil { panic(err) }
 defer dbh.Close()
 
 // Perform SQL-Query
 Rows,err := dbh.Query("SELECT * FROM DuckBreeds ORDER BY 1")
 if err != nil { panic(err) }
 defer Rows.Close()  
 
 // Generate HTML-Table
 v_Doc := New(GC_DocTypeHTML5).
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
```

The methods TrTHSqlRows and TrTdSqlRows retrieve the column-headers and the data-rows from the result-set of the query and generate the column-headers and the data-rows.

For more examples see the files: »UTL_HTML_*test.go«

# Conditional HTML generation
 - func (p_HTML *T_HTML) WHEN(p_Condition bool) *T_HTML 
 - func (p_HTML *T_HTML) OTHERWISE() *T_HTML
 - func (p_HTML *T_HTML) ENDWHEN() *T_HTML

# Example, conditional HTML:

The code snippet below generates an HTML form with an additional field when the username is "admin"
```
  ...
  BMAppHeader().
    BodyOpen("id","ToP").
      FormOpen("formaction",http.MethodPost,"class","FormArea").
        TableOpen("class","FormTable").
          Captionf("FormTableCaption","Add new Bookmar/Folder in %s",v_Parent.NamePath).
          HiddenField("p_FormName","AddForm").
          HiddenField("p_ParentID",strconv.Itoa(v_Parent.ObjectID)).
          TrOpen("class","FormTableRow").
            Td(UTL_HTML.Span("Name","FormPromptField"),"class","FormTableCell").
            Td(UTL_HTML.TextField("p_Name","50","150","","class","FormTextField"),"class","FormTableCell").
          TagCloseUntil("tr").
          TrOpen("class","FormTableRow").
            Td(UTL_HTML.Span("Description","FormPromptField"),"class","FormTableCell").
            Td(UTL_HTML.TextField("p_Description","50","1000","","class","FormTextField"),"class","FormTableCell").
          TagCloseUntil("tr").
          TrOpen("class","FormTableRow").
            Td(UTL_HTML.Span("URL","FormPromptField"),"class","FormTableCell").
            Td(UTL_HTML.TextField("p_URL","50","1000","","class","FormTextField"),"class","FormTableCell").
          TagCloseUntil("tr").
          TrOpen("class","FormTableRow").
            Td(UTL_HTML.Span("Keywords","FormPromptField"),"class","FormTableCell").
            Td(UTL_HTML.TextField("p_Keywords","50","1000","","class","FormTextField"),"class","FormTableCell").
          TagCloseUntil("tr").
          WHEN(User == BMAppDB.GC_UserAdmin). // admin can change ownership, so add the Owner field to the form 
            TrOpen("class","FormTableRow").
              Td(UTL_HTML.Span("Owner","FormPromptField"),"class","FormTableCell").
              Td(UTL_HTML.TextField("p_Owner","50","100",User,"class","FormTextField"),"class","FormTableCell").
            TagCloseUntil("tr").
          ENDWHEN().
          TrOpen("class","FormTableRow").
            Td(UTL_HTML.Span("Public","FormPromptField"),"class","FormTableCell").
            Td(UTL_HTML.BoolField("p_Public",true,"class","FormBoolField"),"class","FormTableCell").
          TagCloseUntil("tr").
          TrOpen("class","FormTableButtonRow").
            TdOpen("class","FormTableButtonCell","colspan","2").
              SubmitButton("p_Submit", "Cancel", "Cancel", "class","FormSubmitButton").
              SubmitButton("p_Submit", "Save", "Save", "class","FormSubmitButton").
    CloseTagsAndWrite(w)
```

# CGI functions

🚧 Under Construction 🚧 TODO: Move into new package UTL_CGI
 - func ReadReqParameter(r *http.Request) map[string]string


# Other Functions
 - func (p_HTML *T_HTML) A(p_Content, p_Href, p_Title string, p_Attributes ...string) *T_HTML
 - func (p_HTML *T_HTML) Af(p_Class, p_Href, p_Title, p_Format string, p_Data ...any) *T_HTML
 - func (p_HTML *T_HTML) Comment(p_Content string) *T_HTML
 - func (p_HTML *T_HTML) Commentf(p_Content string) *T_HTML
 - func CmpAsc(a,b reflect.Value) (Result int) 
 - func CmpDesc(a,b reflect.Value) (Result int) 
