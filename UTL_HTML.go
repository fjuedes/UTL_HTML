// # UTL_HTML - Programmatically generate HTML documents
//
// This package defines a type T_HTML, with functions/methods to build an HTML document from the ground up. For each HTML-tag there is a corresponding function/method with the same name, just with the first letter capitalized. The Div() function/method for example generates complete <div> tags with opening and closing tags, content and attributes. Simple HTML-documents can be constructed by calling the tag-methods in the order the tags should appear in the document and using method-chaining the go code will represent the structure of the HTML document rather well.
//
// # Version
//
// $Id: UTL_HTML.go 79 2025-04-29 20:33:24Z fjuedes $
//
// # Example: (Hello World)
//
//  func HelloWorld() {
//    FileName := "Test_HelloWorld.html"
//    
//    v_Doc := New(GC_DocTypeHTML,0x07).
//             HtmlOpen().
//               HeadOpen().
//                 Title("Hello World in HTML").
//               TagCloseTop().
//               BodyOpen().
//                 POpen().I("Hello ").B("World!").
//             TagCloseAll()
//             
//    if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
//      panic(e.Error())
//    } else {
//      fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
//    } // END if
//  } // END HelloWorld
//
//
// # Overview
//
// Please find a thematically ordered list of the exported functions and methods of the T_HTML object below. Usually the name of the function/method is the name of the HTML-element, just with the first letter capitalized, for example the <title> tag is created through the method Title(). Some tags will never be generated with their complete content, for example the <body> tag. So there is no Body() method defined, but a BodyOpen() method which generates the opening tag only. These tags are tracked in an internal tag-stack and will be closed semi-automagically when one of the TagClose...() methods is called. 
// 
// Funtions whose name end with the letter "f" work similar to fmt.Printf, accepting a string as a format-mask and a variable number of data-items of any type. For example the B(string) method encapsulates content in <b> tags: <b>bold text</b>. The simliar named method Bf(string, any...) method formats data-items according to the format-mask and encapsulates the resulting string into the <b> tag. Example: Bf("Parametervalue is '%d'.",p_Value) will append <b>Parametervalue is '4711'.</b> to the HTML document if the value of p_Value is the number 4711.
//
// Some functions have been also implemented to return the generated tags as strings which sometimes this is a useful hack… 
//
//
// # Basic functions
//
// The list below contains the basic functions to handle the T_HTML struct: Creation of a new document-variable, appending strings and NewLine, the Stringer interface and write functions.
//  - func New(p_DocType string) *T_HTML
//  - func (p_HTML *T_HTML) String()string
//  - func (p_HTML *T_HTML) AS(p_Content string) *T_HTML // shortcut for AppendString
//  - func (p_HTML *T_HTML) AppendString(p_Content string) *T_HTML
//  - func (p_HTML *T_HTML) AppendStringf(p_Format string, p_Data ... any) *T_HTML
//  - func (p_HTML *T_HTML) NL() *T_HTML // NewLine
//  - func (p_HTML *T_HTML) Write(w http.ResponseWriter)
//  - func (p_HTML *T_HTML) CloseTagsAndWrite(w http.ResponseWriter)
//
//
// # Functions for generic tags
//
// The list below contains generic functions to build tags for any MarkUp language: From building complete tags with or without content, with or without attributes to opening and closing tags. 
// By using these functions any type of MarkUp language document can be built.
//  - func (p_HTML *T_HTML) Tag(p_Name, p_Content string, p_Attributes... string) *T_HTML
//  - func (p_HTML *T_HTML) Tagf(p_TagName, p_Class, p_Format string, p_Data ... any) *T_HTML
//  - func (p_HTML *T_HTML) TagOpen(p_Name string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TagCloseTop() *T_HTML
//  - func (p_HTML *T_HTML) TagCloseAll() *T_HTML
//  - func (p_HTML *T_HTML) TagCloseUntil(p_Name string) *T_HTML
//  - func Tag(p_Name, p_Content string, p_Attributes ...string) string
//
// # Example (SVG document)
//
// func Test_XML(t *testing.T) {
//   FileName := "Test_XML.svg"
// 
//   v_Doc := New(GC_DocTypeNONE,0x07).
//            AppendString(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>`).NL().
//            AppendString(`<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`).NL().
//            TagOpen("svg","width","512", "height","512", "viewBox","-70.5 -70.5 391 391", "xmlns","http://www.w3.org/2000/svg", "xmlns:xlink","http://www.w3.org/1999/xlink").
//              Tag("rect","","fill","#fff","stroke","#000","x","-70","y","-70","width","390","height","390").
//              TagOpen("g","opacity","0.8").
//                Tag("rect","","x","25","y","25","width","200","height","200","fill","lime","stroke-width","4","stroke","magenta").
//                Tag("circle","","cx","125","cy","125","r","75","fill","orange").
//                Tag("polyline","","points","50,150 50,200 200,200 200,100","stroke","red","stroke-width","4","fill","none").
//                Tag("line","", "x1","50", "y1","50" ,"x2","200" ,"y2","200", "stroke","blue", "stroke-width","4").
//             TagCloseAll()
// 
//   if e := os.WriteFile(FileName,[]byte(v_Doc.String()),0644); e != nil {
//     t.Errorf(e.Error())
//   } else {
//     fmt.Printf("***** VERIFY file »%s« in a browser.\n",FileName)
//   } // END if
// } // END Test_HelloWorld
// 
//
// # Document structure
//
// Valid HTML documents must contain at least three tags: The <html>tag encloses the entire document. The <head> tag contains information about the text written in the body. The <body> tag contains the information that is displayed by the browser. 
//  - func (p_HTML *T_HTML) HtmlOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) BodyOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) HeadOpen(p_Attributes ...string) *T_HTML
//
//
// # Document header
//
// The document header contains (meta-)information about the HTML-document, defined using the tags <base>, <link>, <meta>, <style> and <title>. - And yes, the <script> tag can be used in HTML-document headers, but that's so 1990's… just don't do this anymore!
//  - func (p_HTML *T_HTML) Base(p_Href, p_target string) *T_HTML 
//  - func (p_HTML *T_HTML) Link(p_Rel, p_Type, p_URL string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Meta(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Style(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Title(p_Content string, p_Attributes ...string) *T_HTML 
//
//
// # Content structuring 
//
// The list below contains basic tags to structure the document-content:
// Headers, line breaks, horizontal rulers, paragraphs, divisions and inline-containers. 
//  - func (p_HTML *T_HTML) Br(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Div(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) DivOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Header(p_Grade, p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Hr(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) P(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Pf(p_Class, p_Format string, p_Data ... any) *T_HTML
//  - func (p_HTML *T_HTML) POpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Span(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Spanf(p_Class, p_Format string, p_Data ... any) *T_HTML {
//  - func (p_HTML *T_HTML) SpanOpen(p_Attributes ...string) *T_HTML
//  - func Span(p_Content string, p_Attributes ...string) string
//
//
// # Basic content formatting
//
// Below list contains direct formatting HTML tags in all four flavors, for example: B() and Bf() for the T_HTML Type and B() and Bf() returning a string. The use of these »ancient« formatting tags is controversial as formatting with CSS is more flexible and separates the data (HTML) from the presentation (CSS).
//  - func (p_HTML *T_HTML) B(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) Bf(p_Class, p_Format string, p_Data... any) *T_HTML 
//  - func (p_HTML *T_HTML) Em(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) Emf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) I(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) If(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) Q(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) Qf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) S(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) Sf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) Strong(p_Content string) *T_HTML
//  - func (p_HTML *T_HTML) Strongf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) Sub(p_Content string) *T_HTML
//  - func (p_HTML *T_HTML) Subf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) Sup(p_Content string) *T_HTML
//  - func (p_HTML *T_HTML) Supf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func (p_HTML *T_HTML) U(p_Content string) *T_HTML 
//  - func (p_HTML *T_HTML) Uf(p_Class, p_Format string, p_Data... any) *T_HTML
//  - func B(p_Content string) string
//  - func Bf(p_Class, p_Format string, p_Data... any) string 
//  - func Em(p_Content string) string 
//  - func Emf(p_Class, p_Format string, p_Data... any) string
//  - func I(p_Content string) string
//  - func If(p_Class, p_Format string, p_Data... any) string
//  - func Q(p_Content string) string
//  - func Qf(p_Class, p_Format string, p_Data... any) string
//  - func S(p_Content string) string
//  - func Sf(p_Class, p_Format string, p_Data... any) string
//  - func Strong(p_Content string) string
//  - func Strongf(p_Class, p_Format string, p_Data... any) string
//  - func Sub(p_Content string) string
//  - func Subf(p_Class, p_Format string, p_Data... any) string
//  - func Sup(p_Content string) string
//  - func Supf(p_Class, p_Format string, p_Data... any) string
//  - func U(p_Content string) string
//  - func Uf(p_Class, p_Format string, p_Data... any) string
//
//
// # Forms [file: UTL_HTML_Form]
//
// 🚧 Currently under construction, TODO: More functions for other input-types, maybe standard form-structures, … 🚧
//  - func (p_HTML *T_HTML) FormOpen(p_action, p_method string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) BoolField(p_name string, p_checked bool, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) HiddenField(p_name, p_value string, p_Attributes... string ) *T_HTML
//  - func (p_HTML *T_HTML) SubmitButton(p_name, p_label, p_value, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TextField(p_name, p_size, p_maxlen, p_value, p_Attributes ...string) *T_HTML
//  - func BoolField(p_name string, p_checked bool, p_Attributes ...string) string
//  - func SubmitButton(p_name, p_label, p_value, p_Attributes ...string) string
//  - func TextField(p_name, p_size, p_maxlen, p_value, p_Attributes ...string) string
//  - func (p_HTML *T_HTML) SelectMenu(p_FieldName, p_MenuClassName, p_ItemClassName, p_DefaultValue string, p_CompareFunc t_CompareFunc, p_MenuItems any, p_Attributes ...any) *T_HTML
//
//
// # Lists [File: UTL_HTML_List]
//
// 🚧 Currently under construction, TODO: more functions for complex data-types 🚧
//  - func (p_HTML *T_HTML) OlOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) UlOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) LiOpen(p_Attributes ...string) *T_HTML {
//  - func (p_HTML *T_HTML) Li(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Lif(p_Class, p_Format string, p_Data ...any) *T_HTML
//
//
// # Tables [File: UTL_HTML_Table]
// 
// Methods to support the creation of HTML-Tables, starting with static tables to the generation of header- and data-rows from structures, maps and slices.
//  - func (p_HTML *T_HTML) TableOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Caption(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Captionf(p_Format string, p_data ...any) *T_HTML
//  - func (p_HTML *T_HTML) TheadOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TbodyOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TfootOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Th(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Thf(p_Class, p_Content string, p_Data ...any) *T_HTML
//  - func (p_HTML *T_HTML) ThOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Td(p_Content string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Tdf(p_Class, p_Content string, p_Data ...any) *T_HTML
//  - func (p_HTML *T_HTML) TdOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TrOpen(p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) TrTh(p_TrClass, p_ThClass string, p_DataItems ...any) *T_HTML
//  - func (p_HTML *T_HTML) TrTd(p_TrClass, p_TdClass string, p_DataItems ...any) *T_HTML
//  - func (p_HTML *T_HTML) TrThStruct(p_TrClass, p_ThClass, p_KeyColHeader string, p_DataItem any) *T_HTML
//  - func (p_HTML *T_HTML) TrTdStruct(p_TrClass, p_TdClass, p_KeyColValue string, p_DataItem any) *T_HTML
//  - func (p_HTML *T_HTML) TrTdMap(p_TrClass, p_TdClass string, p_CompareFunc t_CompareFunc, p_DataItems any) *T_HTML
//  - func (p_HTML *T_HTML) TrTdSlice(p_TrClass, p_TdClass string, p_DataRows any) *T_HTML
//  - func (p_HTML *T_HTML) TrThSqlRows(p_TrClass, p_ThClass string, p_DataRows *sql.Rows) *T_HTML
//  - func (p_HTML *T_HTML) TrTdSqlRows(p_TrClass, p_TdClass string, p_DataRows *sql.Rows) *T_HTML
//
// Example for a simple static table:
//
//  v_Doc := New(GC_DocTypeHTML,0x03).
//           HtmlOpen().
//             HeadOpen("id","4711").
//               Base("/test/","_self").
//               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
//               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
//               Link("stylesheet","text/css","static/Default.css").
//               Link("stylesheet","text/css","static/Bookmarks.css").
//               Title("TableTest").
//             TagCloseUntil("head").
//             BodyOpen().
//               TableOpen("class","class4table").
//                 Captionf("class4caption","Static Table Test performed at %s.",time.Now().Format("2006-01-02T15:04:05")).
//                 TheadOpen().
//                   TrTh("class4tr","class4th","Name","Age","City","Occupation").
//                 TagCloseTop(). // thead
//                 TbodyOpen().
//                   TrTd("class4tr","class4td","Alice","28","Old York","Software Engineer").
//                   TrTd("class4tr","class4td","Bob","42","Lost Angeles","Data Scientist").
//                   TrTd("class4tr","class4td","Charlie","61","Charleston","Freelancer").
//           TagCloseAll()
//  fmt.Println(v_Doc)
//  
// The variable v_Doc contains the whole HTML document, which can be extracted as a string and, printed to an http.ResponseWriter for example.
//
// Beyond just simple static tables there is a group of methods implemented that handle structs, maps and slices. The methods TrThStruct() and TrTdStruct() generate header- and data rows from structures. For example we have a simplified structure for stock symbols (i will explain the html struct-tags later):
//
//  type t_TickerSymbol struct {
//    Name       string   `html:"ColHeader='Full Name' HeaderClass='AdditionalHeaderClass'"`
//    Type       string   `html:"DataClass='AdditionalDataClass'"`
//    LastPrice  float32  `html:"Style='text-align: right; color: green;'"`
//    LastVolume int      `html:"Style='text-align: right; color: blue;'"`
//  } // END t_TickerSymbol 
//
// To generate the table-header row the following snippet can be used:
//
//  TableOpen("class","class4table").
//    Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
//    TheadOpen().
//      TrThStruct("class4tr","class4th","Symbol",t_TickerSymbol{}).
//    TagCloseTop().
//    …
//
// The method TrThStruct() takes CSS class-names for the <tr> and the <th> tag, one additional column-name (we see later for what that is being used) and an empty t_TickerSymbol to generate a <tr><th>…</th></tr> structure. No loop necessary!
// 
// Data-rows are handled by the method TrTDStruct(), for example:
//    … 
//    TbodyOpen().
//      TrTdStruct("class4tr","class4td","YHOO",MyTickerSymbol).
//    …
// 
// However that does not make much sense, as this takes care of only a single row of data. You will rarely use the TrTdStruct method directly, but it is used by the methods that handle entire maps and slices. If we have the following data:
//
//  var TickerTable = map[string]t_TickerSymbol{
//    "XOMO"  : {Name: "YieldMax XOM Option Income Strategy ETF", Type: "ETF", LastPrice: 14.48, LastVolume: 46165},
//    "EURUSD": {Name: "Euro US Dollar", Type: "Currency", LastPrice: 1.0813},
//    "TB91D" : {Name: "91 Day Treasury Bill", Type: "FixedIncome", LastPrice: 4.1900},
//    "GSHRU" : {Name: "Gesher Acquisition Corp. II", Type: "IPO", LastPrice: 10.02,  LastVolume: 6324701},
//    "EACAX" : {Name: "Eaton Vance California Municipal Opportunities Fund Class A", Type: "Mutual Fund", LastPrice: 10.99},
//    "CNP"   : {Name: "CenterPoint Energy, Inc (Holding Co) Common Stock", Type: "Stock", LastPrice: 35.77, LastVolume: 8615303},
//    "SPX"   : {Name: "S&P 500", Type: "Index", LastPrice: 5667.5600},
//  }
// 
// The whole ticker-table can be generated with this code:
//   …
//   TableOpen("class","class4table").
//     Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
//     TheadOpen().
//       TrThStruct("class4tr","class4th","Symbol",t_TickerSymbol{}).
//     TagCloseTop().
//     TbodyOpen().
//       TrTdMap("class4tr","class4td",CmpAsc,TickerTable).
//   …
// 
// See the full example in the file UTL_HTML_test as func Test_Table_MapStrStruct(t *testing.T).
//
// The method TrThStruct accepts any struct-type and non-exported fields of the struct{} are skipped.
//
// The method TrTdStruct accepts any struct-type and non-exported fields of the struct{} are skipped. Pointer fields are dereferenced and nil-values are replaced with "&nbsp;"
//
// The method TrTdMap loops through all the keys of the map - sorted or unsorted - and generates the data-rows. This method can process the following types of maps:
//  - map[KeyType]any → Table with two columns: Key, Value
//  - map[KeyType]*any → Table with two columns: Key, Value
//  - map[KeyType]struct{} → Table with one column for the map-key and one column per exported struct-field
//  - map[KeyType]*struct{} → Table with one column for the map-key and one column per exported struct-field
//  - map[KeyType][]any → Table with one column for the map-key and one column per slice element
//  - map[KeyType][]*any → Table with one column for the map-key and one column per slice element
//
// Keys and values are being converted into strings using fmt.Sprint, composite types will therefore be converted into strings according to their Stringer interface which might not be the desired result.
//
// In a very similar way, the method TrTdSlice() is processes slices into table data-rows. The following slice-types are supported:
//  - []any → one row with n columns
//  - []*any → one row with n columns
//  - []struct → Table with one column per struct-field
//  - []*struct → Table with one column per struct-field
//  - [][]any → two dimensional table
//  - [][]*any → two dimensional table
//
// Again, values are being converted into strings using fmt.Sprint, so composite types will be converted into strings according to their Stringer interface.
//
// # The html struct-tag
//
// Now finally, about the struct-tags in the example t_TickerSymbol structure: When you create your own data-structures, you can add additional information to each field of your structure to control how HTML is generated. The syntax of the html struct-tag:
//  - `html:"Skip"`                         - Skip this field
//  - `html:"ColHeader='HeaderText'"`       - Use 'Headertext' instead of field-name
//  - `html:"HeaderClass='CSS class-name'"` - add class-name to <th> tag
//  - `html:"DataClass='CSS class-name'"`   - add class-name to <td> tag
//  - `html:"Style="CSS Declaration"`       - add style attribute to <th> or <td> tag
//
// All combinations are supported, for example:
// `html:"ColHeader='PrimaryKey'" HeaderClass='Centered' DataClass='Centered' Style='color: red;'`
//
// The keywords "ColHeader", "HeaderClass", "DataClass", "Skip" and "Style" are case insensitive, the tag-name "html" is case-sensitive.
//
// Please use the html struct-tag only as last resort: Remember the the »Separation of Concerns«:
//
// - HTML serves as the structure of the content. It provides the semantic markup for the information being displayed on the web page.
// - CSS is used for the presentation or design of the HTML elements. It controls how the elements are displayed in terms of layout, colors, fonts, and other visual aspects. 
// - JavaScript is responsible for behavior and/or interactivity of the web page. It enables dynamic content updates, event handling, and user interactions.
// 
// Embedding style information into the HTML code violates this design philosophy!
//
// 
// # One more thing.
//
// Usually data is stored in some back-end systems like relationbal Databases, organized in tables. UTL_HTML directly supports SQL rowsets; the following table is defined in the attached SQLite database:
//
//  CREATE TABLE "DuckBreeds" (
//      "Breed"         TEXT NOT NULL,
//      "EggColor"      TEXT NOT NULL,
//      "EggSize"       TEXT NOT NULL,
//      "EggProduction" TEXT NOT NULL,
//      "Class"         TEXT NOT NULL,
//      "WeightRange"   TEXT NOT NULL,
//      "Mothering"     TEXT NOT NULL,
//      "Foraging"      TEXT NOT NULL,
//      "Hardiness"     TEXT NOT NULL,
//      "Personality"   TEXT NOT NULL,
//      "Flying"        INTEGER NOT NULL,
//      CONSTRAINT "pk_DuckBreeds" PRIMARY KEY("Breed")
//  );
//      
//
// Only a few lines of code are necessary to display the content of this table as an HTM document:
//
//  // Open Database
//  dbh,err := sql.Open("sqlite","Ducks.sqlite3")
//  if err != nil { panic(err) }
//  defer dbh.Close()
//  
//  // Perform SQL-Query
//  Rows,err := dbh.Query("SELECT * FROM DuckBreeds ORDER BY 1")
//  if err != nil { panic(err) }
//  defer Rows.Close()  
//  
//  // Generate HTML-Table
//  v_Doc := New(GC_DocTypeHTML5,0x02).
//           HtmlOpen().
//             HeadOpen("id","0815").
//               Base("/test/","_self").
//               Meta("http-equiv","content-type","content","text/html; charset=UTF-8").
//               Meta("name","viewport", "content","width=device-width, initial-scale=0.9").
//               Link("stylesheet","text/css","static/Default.css").
//               Link("stylesheet","text/css","static/Bookmarks.css").
//               Title("TableTest").
//             TagCloseUntil("head").
//             BodyOpen().
//               TableOpen("class","class4table").
//                 Captionf("class4caption",Caption,time.Now().Format("2006-01-02T15:04:05")).
//                 TheadOpen().
//                   TrThSqlRows("class4tr","class4th",Rows).
//                 TagCloseTop().
//                 TbodyOpen().
//                   TrTdSqlRows("class4tr","class4td",Rows).
//           TagCloseAll()
//
// The methods TrTHSqlRows and TrTdSqlRows retrieve the column-headers and the data-rows from the result-set of the query and generate the column-headers and the data-rows.
//
// For more examples see the files: »UTL_HTML_*test.go«
//
//
// # CGI functions
//  - func ReadReqParameter(r *http.Request) map[string]string
//
//
// # Other Functions
//  - func (p_HTML *T_HTML) A(p_Content, p_Href, p_Title string, p_Attributes ...string) *T_HTML
//  - func (p_HTML *T_HTML) Af(p_Class, p_Href, p_Title, p_Format string, p_Data ...any) *T_HTML
//  - func (p_HTML *T_HTML) Comment(p_Content string) *T_HTML
//  - func (p_HTML *T_HTML) Commentf(p_Content string) *T_HTML
//  - func CmpAsc(a,b reflect.Value) (Result int) 
//  - func CmpDesc(a,b reflect.Value) (Result int) 
// 
//
package UTL_HTML

import (
  "strings"
  "encoding/base64"
  "fmt"
  "cmp"
  "net/http"
  "bytes"
  "regexp"
  "reflect"
)

const (
  GC_DocTypeHTML5    string = `html`
  GC_DocTypeXML      string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`
  GC_DocTypeMathML20 string = `math PUBLIC "-//W3C//DTD MathML 2.0//EN" "http://www.w3.org/Math/DTD/mathml2/mathml2.dtd"`
  GC_DocTypeMathML10 string = `math SYSTEM "http://www.w3.org/Math/DTD/mathml1/mathml.dtd"`
  GC_DocTypeSVG11    string = `SVG 1.1//EN "http://www.w3.org/2002/04/xhtml-math-svg/xhtml-math-svg.dtd"`
  GC_DocTypeNONE     string = ``
  
  gc_Version     string = "$Id: UTL_HTML.go 79 2025-04-29 20:33:24Z fjuedes $"
) // END const

// Structure to build and manage html content
type T_HTML struct {
  content  bytes.Buffer // contains the generated HTML document
  tagStack []string     // the stack for the opened tags
  
  if_active    bool     // true we're in an IF condition
  if_condition bool     // result of the active condition
  
  nlMode   byte         // 0x00: No NL Mode
                        // 0x01: NL after every closing tag
                        // 0x02: NL after every complete tag
                        // 0x04: NL after every opening tag
} // END T_HTML


// -------------------------------------------------------------------
// ********** Manage the tagStack with push, pop and peek **********
// all internal, nothing is exported
// - push: Push a TagName to the top of the stack
// - pop : Remove the top Element from the stack and return the TagName
// - peek: Return the top Element from the stack without removing it
// -------------------------------------------------------------------

// push a tagname to the stack
func (p_HTML *T_HTML) pushTag(p_TagName string) {
  p_HTML.tagStack = append(p_HTML.tagStack, p_TagName)
} // END pushTag

// pop the top tagname from the stack
func (p_HTML *T_HTML) popTag() string {
  if len(p_HTML.tagStack) == 0 {
    return ""
  } // END if
  LastIndex := len(p_HTML.tagStack) - 1
  TagName := p_HTML.tagStack[LastIndex]
  p_HTML.tagStack = p_HTML.tagStack[:LastIndex]
  return TagName
} // END popTag

// peek at the top tagname on the stack
func (p_HTML *T_HTML) peekTag() string {
  if len(p_HTML.tagStack) == 0 {
    return ""
  } // END if
  LastIndex := len(p_HTML.tagStack) - 1
  TagName := p_HTML.tagStack[LastIndex]
  return TagName
} // END peekTag


// -------------------------------------------------
// ********** HTML object basic functions **********
// -------------------------------------------------

// Return a new HTML Document with doc-type
// availablae docu-types in this package are
//  - GC_DocTypeHTML for HTML documents
//  - GC_DocTypeNONE for HTML snippets - basically an empty string
//
func New(p_DocType string, p_NLMode byte) *T_HTML {
  v_HTML := new(T_HTML)
  v_HTML.tagStack = make([]string, 0, 10)
  v_HTML.nlMode = p_NLMode
	v_HTML.if_active = false
	v_HTML.if_condition = true
  if p_DocType != "" {
    v_HTML.AppendStringf("<!DOCTYPE %s>",p_DocType).NL().
           AppendStringf("<!-- %s  -->",gc_Version).NL()
  } // END if
  return v_HTML
} // END New_HTML

// Return the whole HTML-document as a single string
// Implementation of the Stringer interface
func (p_HTML *T_HTML) String() string {
  return p_HTML.content.String()
} // END String

// Append string data to the HTML-content
// Alias for AppendString
// ¡This method implements the IF ELSE ENDIF functionality!
func (p_HTML *T_HTML) AS(p_Content string) *T_HTML {
	if !p_HTML.if_active || p_HTML.if_condition {
    p_HTML.content.WriteString(p_Content)
	} // END if
  return p_HTML
} // END AS

// Append string data to the HTML-content
// Alias for AS
func (p_HTML *T_HTML) AppendString(p_Content string) *T_HTML { 
  return p_HTML.AS(p_Content)
} // END AppendString

// Append string data to the HTML-content
// Parameters are identical to fmt.Printf
func (p_HTML *T_HTML) AppendStringf(p_Format string, p_Data ... any) *T_HTML {
  return p_HTML.AS(fmt.Sprintf(p_Format, p_Data...))
} // END AppendStringf

// Append CR/NL data to the content
func (p_HTML *T_HTML) NL() *T_HTML { 
  return p_HTML.AS("\r\n") 
} // END NL

// Write HTML-document to a ResponseWriter
func (p_HTML *T_HTML) Write(w http.ResponseWriter) {
  w.Write(p_HTML.content.Bytes())
} // END Write

// Close all remaining tags and write conten to the ResponseWriter
func (p_HTML *T_HTML) CloseTagsAndWrite(w http.ResponseWriter) {
  p_HTML.TagCloseAll()
  w.Write(p_HTML.TagCloseAll().content.Bytes())
} // END Write

// Return the length of the HTML-Buffer
func (p_HTML *T_HTML) Len()int {
  return p_HTML.content.Len()
} // END Len

// -------------------------------------------
// ********** Generic tag functions **********
// -------------------------------------------

// Append an Attribute with value if Name and Value to a slice of strings when both are not empty
// not exported
func appendAttribute(p_Name, p_Value string, p_AttrList *[]string)  { 
  if p_Name != "" && p_Value != "" {
    *p_AttrList = append(*p_AttrList,p_Name,p_Value)
  } // END if
} // END appendAttribute

// Build a tag with or without content and with or without attributes and return it as a string.
// Can generate tags for any kind of markup language.
func Tag(p_Name, p_Content string, p_Attributes ...string) string {
  Result := "<" + p_Name;
  
  if len(p_Attributes) == 0 { // tag without any attributes
    if len(p_Content) == 0 {
      Result += " />" // tag without content
    } else {
      Result += ">" + p_Content + "</" + p_Name + ">"
    } // END if
    return Result
  } // END if

  // process attributes
  for index := 0; index < len(p_Attributes); index++ {
    if index == (len(p_Attributes) - 1) { // no more value = tag without value
      Result += " " + p_Attributes[index]
    } else {
      Result += " " + p_Attributes[index] + "=\"" + p_Attributes[index+1] + "\""
      index++
    } // END if
  } // END for

  if len(p_Content) == 0 {
    Result += " />"
  } else {
    Result += ">" + p_Content + "</" + p_Name + ">"
  } // END if

  return Result
} // END Tag

// Return a complete tag with class-attribute and content as a string
// The content is composed from data-elements, formatted through fmt.Sprintf.
func Tagf(p_TagName, p_Class, p_Format string, p_Data ... any) string {
  if p_Class != "" {
    return Tag(p_TagName,fmt.Sprintf(p_Format,p_Data...),"class",p_Class)
  } else {
    return Tag(p_TagName,fmt.Sprintf(p_Format,p_Data...))
  } // END if
} // END Tagf

// Return complete tag with or without content and with or without attributes to the HTML-document
// Can generate tags for any kind of markup language.
func (p_HTML *T_HTML) Tag(p_TagName, p_Content string, p_Attributes ...string) *T_HTML {
  p_HTML.AS("<" + p_TagName)

  if len(p_Attributes) == 0 { // tag without any attributes
    if len(p_Content) == 0 {
      p_HTML.AS(" />") // tag without content
    } else {
      p_HTML.AS(">" + p_Content + "</" + p_TagName + ">")
    } // END if
    if (p_HTML.nlMode & 0x01) != 0 {
      p_HTML.NL()
    } // END if
    return p_HTML
  } // END if

  for index := 0; index < len(p_Attributes); index++ {
    if index == (len(p_Attributes) - 1) {
      p_HTML.AS(" " + p_Attributes[index])
    } else {
      p_HTML.AS(" " + p_Attributes[index] + "=\"" + p_Attributes[index+1] + "\"")
      index++
    } // END if
  } // END for

  if len(p_Content) == 0 {
    p_HTML.AS(" />")
  } else {
    p_HTML.AS(">" + p_Content + "</" + p_TagName + ">")
  } // END if

  if (p_HTML.nlMode & 0x02) != 0 {
    p_HTML.NL()
  } // END if
  return p_HTML
} // END Tag

// Append a complete tag with class-attribute and content to the HTML-document.
// The content is composed from data-elements, formatted through fmt.Sprintf.
func (p_HTML *T_HTML) Tagf(p_TagName, p_Class, p_Format string, p_Data ... any) *T_HTML {
  if p_Class != "" {
    return p_HTML.Tag(p_TagName,fmt.Sprintf(p_Format,p_Data...),"class",p_Class)
  } else {
    return p_HTML.Tag(p_TagName,fmt.Sprintf(p_Format,p_Data...))
  } // END if
} // END Tagf

// Append an opening tag with attributes
func (p_HTML *T_HTML) TagOpen(p_TagName string, p_Attributes ...string) *T_HTML {
  // Print the opening tag name
  p_HTML.pushTag(p_TagName)
  p_HTML.AS("<" + p_TagName)

  // process the attributes - if any
  for index := 0; index < len(p_Attributes); index++ {
    if index == (len(p_Attributes) - 1) {
      p_HTML.AS(" " + p_Attributes[index])
    } else {
      // TODO: Think about how to treat empty attribute values: attr=""; attr="attr"; attr maybe an EmptyAttrMode?
      p_HTML.AS(" " + p_Attributes[index] + "=\"" + p_Attributes[index+1] + "\"")
      index++
    } // END if
  } // END for

  p_HTML.AS(">")
  if (p_HTML.nlMode & 0x04) != 0 {
    p_HTML.NL()
  } // END if
  return p_HTML
} // END TagClose

// Close the tag on the top of the stack
func (p_HTML *T_HTML) TagCloseTop() *T_HTML {
  if TagName := p_HTML.popTag(); TagName != "" {
    p_HTML.AS("</" + TagName + ">")
  } // END if
  if (p_HTML.nlMode & 0x01) != 0 {
    p_HTML.NL()
  } // nlMode on closing tag
  return p_HTML
} // END TagCloseTop

// Close all remaining tags
func (p_HTML *T_HTML) TagCloseAll() *T_HTML {
  for {
    if TagName := p_HTML.peekTag(); TagName != "" {
      p_HTML.TagCloseTop()
    } else {
      break
    } // END if
  } // END for
  return p_HTML
} // END TagCloseAll

// Close all remaining tags until p_Name is found
func (p_HTML *T_HTML) TagCloseUntil(p_Name string) *T_HTML {
  for {
    if TagName := p_HTML.popTag(); TagName == "" {
      break // stack is empty
    } else {
      p_HTML.AS("</" + TagName + ">")
      if (p_HTML.nlMode & 0x01) != 0 {
        p_HTML.NL()
      } // nlMode on closing tag
      if TagName == p_Name {
        break
      }
    } // END if
  } // END for
  return p_HTML
} // END TagCloseUntil


// ------------------------------------------
// ********** Structural HTML tags **********
// ------------------------------------------

func (p_HTML *T_HTML) HtmlOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("html",p_Attributes...) }
func (p_HTML *T_HTML) BodyOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("body",p_Attributes...) }
func (p_HTML *T_HTML) HeadOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("head",p_Attributes...) }


// ----------------------------------------------------
// ********** HTML Tags for document headers **********
// ----------------------------------------------------

func (p_HTML *T_HTML) Base(p_Href, p_target string) *T_HTML { return p_HTML.Tag("base","","href",p_Href,"target",p_target) }
func (p_HTML *T_HTML) Link(p_Rel, p_Type, p_Href string, p_Attributes ...string) *T_HTML { 
  Arguments := make([]string,0,len(p_Attributes)+3)
  appendAttribute("type",p_Type,&Arguments)
  appendAttribute("rel",p_Rel,&Arguments)
  appendAttribute("href",p_Href,&Arguments)
  Arguments = append(Arguments, p_Attributes...)
  return p_HTML.Tag("link","",Arguments...) 
} // END Link
func (p_HTML *T_HTML) Meta(p_Attributes ...string) *T_HTML { return p_HTML.Tag("meta", "", p_Attributes...) }
func (p_HTML *T_HTML) Style(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("style",p_Content,p_Attributes...) }
func (p_HTML *T_HTML) Title(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("title", p_Content, p_Attributes...) }


// ----------------------------------------------------
// ********** Document structuring HTML Tags **********
// ----------------------------------------------------

func (p_HTML *T_HTML) Header(p_Grade, p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("h"+p_Grade, p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Br(p_Attributes ...string) *T_HTML { return p_HTML.Tag("br", "", p_Attributes...) }
func (p_HTML *T_HTML) Hr(p_Attributes ...string) *T_HTML { return p_HTML.Tag("hr", "", p_Attributes...) }
func (p_HTML *T_HTML) Div(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("div", p_Content, p_Attributes...) }
func (p_HTML *T_HTML) DivOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("div", p_Attributes...) }
func (p_HTML *T_HTML) P(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("p", p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Pf(p_Class, p_Format string, p_Data ... any) *T_HTML { return p_HTML.Tagf("p",p_Class,p_Format,p_Data...) }
func (p_HTML *T_HTML) POpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("p", p_Attributes...) }
func (p_HTML *T_HTML) Span(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("span", p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Spanf(p_Class, p_Format string, p_Data ... any) *T_HTML { return p_HTML.Tagf("span", p_Class, p_Format,p_Data...) }
func (p_HTML *T_HTML) SpanOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("span", p_Attributes...) }
func Span(p_Content string, p_Attributes ...string) string { return Tag("span",p_Content,p_Attributes...) }
func Spanf(p_Class, p_Format string, p_Data ...any) string { return Tagf("span",p_Class,p_Format,p_Data...) }


// ---------------------------------------------
// ********** Content Formatting Tags **********
// ---------------------------------------------

func (p_HTML *T_HTML) B(p_Content string) *T_HTML { return p_HTML.Tag("b", p_Content) }
func (p_HTML *T_HTML) Bf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("b",p_Class,p_Format,p_Data...) }
func (p_HTML *T_HTML) Em(p_Content string) *T_HTML { return p_HTML.Tag("em", p_Content) }
func (p_HTML *T_HTML) Emf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("em", p_Class, p_Format,p_Data...) }
func (p_HTML *T_HTML) I(p_Content string) *T_HTML { return p_HTML.Tag("i", p_Content) }
func (p_HTML *T_HTML) If(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("i", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) Q(p_Content string) *T_HTML { return p_HTML.Tag("u", p_Content) }
func (p_HTML *T_HTML) Qf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("u", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) S(p_Content string) *T_HTML { return p_HTML.Tag("u", p_Content) }
func (p_HTML *T_HTML) Sf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("u", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) Strong(p_Content string) *T_HTML { return p_HTML.Tag("strong", p_Content) }
func (p_HTML *T_HTML) Strongf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("strong", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) Sub(p_Content string) *T_HTML { return p_HTML.Tag("sub", p_Content) }
func (p_HTML *T_HTML) Subf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("sub", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) Sup(p_Content string) *T_HTML { return p_HTML.Tag("sup", p_Content) }
func (p_HTML *T_HTML) Supf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("sup", p_Class, p_Format, p_Data...) }
func (p_HTML *T_HTML) U(p_Content string) *T_HTML { return p_HTML.Tag("u", p_Content) }
func (p_HTML *T_HTML) Uf(p_Class, p_Format string, p_Data... any) *T_HTML { return p_HTML.Tagf("u", p_Class, p_Format, p_Data...) }

func B(p_Content string) string { return Tag("b", p_Content) }
func Bf(p_Class, p_Format string, p_Data... any) string { return Tagf("b",p_Class,p_Format,p_Data...) }
func Em(p_Content string) string { return Tag("em", p_Content) }
func Emf(p_Class, p_Format string, p_Data... any) string { return Tagf("em", p_Class, p_Format,p_Data...) }
func I(p_Content string) string { return Tag("i", p_Content) }
func If(p_Class, p_Format string, p_Data... any) string { return Tagf("i", p_Class, p_Format, p_Data...) }
func Q(p_Content string) string { return Tag("u", p_Content) }
func Qf(p_Class, p_Format string, p_Data... any) string { return Tagf("u", p_Class, p_Format, p_Data...) }
func S(p_Content string) string { return Tag("u", p_Content) }
func Sf(p_Class, p_Format string, p_Data... any) string { return Tagf("u", p_Class, p_Format, p_Data...) }
func Strong(p_Content string) string { return Tag("strong", p_Content) }
func Strongf(p_Class, p_Format string, p_Data... any) string { return Tagf("strong", p_Class, p_Format, p_Data...) }
func Sub(p_Content string) string { return Tag("sub", p_Content) }
func Subf(p_Class, p_Format string, p_Data... any) string { return Tagf("sub", p_Class, p_Format, p_Data...) }
func Sup(p_Content string) string { return Tag("sup", p_Content) }
func Supf(p_Class, p_Format string, p_Data... any) string { return Tagf("sup", p_Class, p_Format, p_Data...) }
func U(p_Content string) string { return Tag("u", p_Content) }
func Uf(p_Class, p_Format string, p_Data... any) string { return Tagf("u", p_Class, p_Format, p_Data...) }

// -------------------------------------
// ********** Other HTML-Tags **********
// -------------------------------------

// Append Hyper-Link
func (p_HTML *T_HTML) A(p_Content, p_Href, p_Title string, p_Attributes ...string) *T_HTML {
  Arguments := make([]string,0,len(p_Attributes)+2)
  appendAttribute("href",p_Href,&Arguments)
  appendAttribute("title",p_Title,&Arguments)
  Arguments = append(Arguments, p_Attributes...)
  return p_HTML.Tag("a", p_Content, Arguments...)
} // END A

// Append Hyper-Link with formatted content
func (p_HTML *T_HTML) Af(p_Class, p_Href, p_Title, p_Format string, p_Data ...any) *T_HTML {
  return p_HTML.A(fmt.Sprintf(p_Format,p_Data...),p_Href,p_Title,"class",p_Class)
} // END A

// Append comment
func (p_HTML *T_HTML) Comment(p_Content string) *T_HTML {
  return p_HTML.AS("<!-- " + p_Content + " -->\n")
} // END Comment

// Append formatted comment
func (p_HTML *T_HTML) Commentf(p_format string, p_Data ... any) *T_HTML {
  return p_HTML.AS("<!-- " + fmt.Sprintf(p_format, p_Data...) + " -->\n")
} // END Comment


func (p_HTML *T_HTML) WHEN(p_Condition bool) *T_HTML {
	if p_HTML.if_active { panic("Nested IF is not implemented!") }
	p_HTML.if_active = true
	p_HTML.if_condition = p_Condition
	return p_HTML
} // END WHEN

func (p_HTML *T_HTML) OTHERWISE() *T_HTML {
	if !p_HTML.if_active { panic("ELSE without IF!") }
	p_HTML.if_condition = !p_HTML.if_condition
	return p_HTML
} // END OTHERWISE

func (p_HTML *T_HTML) ENDWHEN() *T_HTML {
	if !p_HTML.if_active { panic("ENDIF without IF!") }
	p_HTML.if_active = false
	p_HTML.if_condition = true
	return p_HTML
} // END ENDWHEN


// -----------------------------------------
// ********** Other CGI functions **********
// -----------------------------------------

// Copy the headers, url-parameters and post-data into a string-map
func ReadReqParameter(r *http.Request) map[string]string {
  // create new parameter-list
  Result := make(map[string]string)
  
  // copy the header
  for Name,Values := range(r.Header) {
    Result[Name] = strings.Join(Values,";") // KISS: Multiple values are separated by ";"
    if Name == "Authorization" {
      // Try to decode the basic authentication header and store resulty into parameterlist
      if  Credentials, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(Result[Name], "Basic ")); err == nil {
        Result["UserCred"] = string(Credentials)
        if parts := strings.SplitN(string(Credentials), ":", 2); len(parts) == 2 {
          Result["UserName"] = parts[0]
          Result["UserAuth"] = parts[1]
        } // END if
      } // END if
    } // END if
  } // END for
  
  // URL and POST Data
  if Error := r.ParseForm(); Error != nil { panic(Error) }
  for Name,Values := range(r.Form) {
    Result[Name] = strings.Join(Values,";") // KISS: Multiple values are separated by ";"
  } // END for

  return Result
} // END ReadReqParameter

// -----------------------------------------------------------------------------
// ********** helper function for the analysis of the html struct-tag **********
// -----------------------------------------------------------------------------
// Syntax of the html struct-tag:
// `html:"Skip"`                         - Skip this field
// `html:"ColHeader='HeaderText'"`       - use 'Headertext' instead of field-name
// `html:"HeaderClass='CSS class-name'"` - add class-name to <th> tag
// `html:"DataClass='CSS class-name'"`   - add class-name to <td> tag
// `html:"Style="CSS Declaration"`       - add style attribute to <th> or <td> tag
//
// All combinations are supported, for example:
// `html:"ColHeader='PrimaryKey'" HeaderClass='Centered' DataClass='Centered' Style='color: red;'`
//

var regExp_ColHeader = regexp.MustCompile(`(?i)ColHeader='(.*?)'`)
var regExp_HeaderClass = regexp.MustCompile(`(?i)HeaderClass='(.*?)'`)
var regExp_DataClass = regexp.MustCompile(`(?i)DataClass='(.*?)'`)
var regExp_Style = regexp.MustCompile(`(?i)Style='(.*?)'`)
var regExp_Skip = regexp.MustCompile(`(?i)Skip`)

func analyzeHtmlStructTag(p_HtmlStructTag string) (HeaderText, HeaderClass, DataClass, Style string, Skip bool) {
  if Skip = regExp_Skip.MatchString(p_HtmlStructTag); Skip {
    return //
  } // END if
  if match := regExp_ColHeader.FindStringSubmatch(p_HtmlStructTag); len(match) > 1 {
    HeaderText = match[1]
  } // END if
  if match := regExp_HeaderClass.FindStringSubmatch(p_HtmlStructTag); len(match) > 1 {
    HeaderClass = match[1]
  } // END if
  if match := regExp_DataClass.FindStringSubmatch(p_HtmlStructTag); len(match) > 1 {
    DataClass = match[1]
  } // END if
  if match := regExp_Style.FindStringSubmatch(p_HtmlStructTag); len(match) > 1 {
    Style = match[1]
  } // END if
  return
} // END analyzeHtmlStructTag

// ---------------------------------------------------------------
// ********** sort functions for reflect.Value elements **********
// ---------------------------------------------------------------
type t_CompareFunc func(a,b reflect.Value)int

// Compare function for ascending sort-order
func CmpAsc(a,b reflect.Value) (Result int) {
  Result = 0
  if a.Kind() != b.Kind() { 
    Result = cmp.Compare(fmt.Sprint(a.Interface()),fmt.Sprint(b.Interface())) // different types: compare the Stringer value
  } // END if
  
  switch a.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,reflect.Int64: Result = cmp.Compare(a.Int(),b.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64: Result =  cmp.Compare(a.Uint(),b.Uint())
    case reflect.Float32, reflect.Float64: Result = cmp.Compare(a.Float(),b.Float()) 
    default: Result = cmp.Compare(fmt.Sprint(a.Interface()),fmt.Sprint(b.Interface())) // everything else will be compared as string
  } // END switch
  return Result
} // END CopAsc

// Compare function for descending sort-order
func CmpDesc(a,b reflect.Value) (Result int) { return -CmpAsc(a,b) }
