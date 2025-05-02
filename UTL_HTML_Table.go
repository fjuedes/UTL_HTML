package UTL_HTML
//
// UTL_HTML_Table
// Version: $Id: UTL_HTML_Table.go 78 2025-04-22 00:43:36Z fjuedes $
//
import (
  "os"
	"fmt"
	"slices"
	"reflect"
	"database/sql"
)

// *************************************************************************************
// A couple of not exported little helper functions to handle empty parameters correctly
func (p_HTML *T_HTML) helper_TrOpen(p_TrClass string) *T_HTML {
	if p_TrClass != "" {
		p_HTML.TrOpen("class",p_TrClass)
	} else {
		p_HTML.TrOpen()
	} // END if
	return p_HTML
} // END helper_TrOpen

func (p_HTML *T_HTML) helper_Tag(p_TagName, p_Content, p_TagClass, p_AdditionalClass, p_Style string) *T_HTML {
	TagStr := "<" + p_TagName
	if p_TagClass != "" && p_AdditionalClass != "" {
		TagStr += fmt.Sprintf(` class="%s %s"`,p_TagClass,p_AdditionalClass)
	} else if p_TagClass != "" {
		TagStr += fmt.Sprintf(` class="%s"`,p_TagClass)
	} else {
		TagStr += fmt.Sprintf(` class="%s"`,p_AdditionalClass)
	} // END if
	
	if p_Style != "" {
		TagStr += fmt.Sprintf(` style="%s"`,p_Style)
	} // END if
		
	if p_Content != "" {
		TagStr += fmt.Sprintf(`>%s</%s>`,p_Content,p_TagName)
	} else {
		TagStr += `/>`
	} // END if

	if (p_HTML.nlMode & 0x02) != 0 {
		TagStr += "\r\n"
	} // END if

	return p_HTML.AS(TagStr)
} // END helper_Tag
	
func (p_HTML *T_HTML) helper_Th(p_Content, p_ThClass, p_HeaderClass, p_Style string) *T_HTML {
	return p_HTML.helper_Tag("th",p_Content, p_ThClass, p_HeaderClass, p_Style)
} // END helper_Th

func (p_HTML *T_HTML) helper_Td(p_Content, p_TdClass, p_DataClass, p_Style string) *T_HTML {
	return p_HTML.helper_Tag("td",p_Content, p_TdClass, p_DataClass, p_Style)
} // END helper_Td

// *************************************************************************************
//


func (p_HTML *T_HTML) TableOpen(p_Attributes ...string) *T_HTML {	return p_HTML.TagOpen("table", p_Attributes...) }
func (p_HTML *T_HTML) TheadOpen(p_Attributes ...string) *T_HTML {	return p_HTML.TagOpen("thead", p_Attributes...) }
func (p_HTML *T_HTML) TbodyOpen(p_Attributes ...string) *T_HTML {	return p_HTML.TagOpen("tbody", p_Attributes...) }
func (p_HTML *T_HTML) TfootOpen(p_Attributes ...string) *T_HTML {	return p_HTML.TagOpen("tfoot", p_Attributes...) }
func (p_HTML *T_HTML) TrOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("tr", p_Attributes...) }
func (p_HTML *T_HTML) Caption(p_Content string, p_Attributes ...string) *T_HTML {	return p_HTML.Tag("caption",p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Captionf(p_Class, p_Format string, p_Data ...any) *T_HTML {	return p_HTML.Tagf("caption",p_Class,p_Format, p_Data...) }
func (p_HTML *T_HTML) Th(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("th",p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Thf(p_Class, p_Content string, p_Data ...any) *T_HTML { return p_HTML.Tagf("th",p_Class,p_Content,p_Data...) }
func (p_HTML *T_HTML) Td(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("td",p_Content, p_Attributes...) }
func (p_HTML *T_HTML) Tdf(p_Class, p_Content string, p_Data ...any) *T_HTML { return p_HTML.Tagf("td",p_Class,p_Content,p_Data...) }
func (p_HTML *T_HTML) ThOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("th",p_Attributes...) }
func (p_HTML *T_HTML) TdOpen(p_Attributes ...string) *T_HTML { 	return p_HTML.TagOpen("td",p_Attributes...) }

// Create table header-row from a variable number of DataItems:
//  - p_TrClass - CSS classname for <tr>
//  - p_ThClass - CSS classname for <th>
//  - p_DataItems - Any data-type that is printable with fmt.Sprint
//  - Pointer-types will be dereferenced and nil values replaced with "&nbsp;"
//
func (p_HTML *T_HTML) TrTh(p_TrClass, p_ThClass string, p_DataItems ...any) *T_HTML {
	p_HTML.helper_TrOpen(p_TrClass)
	for _,DataItem := range(p_DataItems) {
		if reflect.TypeOf(DataItem).Kind() != reflect.Ptr {
	    p_HTML.helper_Th(fmt.Sprint(DataItem),p_ThClass,"","")
		} else {
			if reflect.ValueOf(DataItem).IsNil() {
				p_HTML.helper_Th("&nbsp;",p_ThClass,"","")
			} else {
	      p_HTML.helper_Th(fmt.Sprint(reflect.ValueOf(DataItem).Elem().Interface()),p_ThClass,"","")
			} // END if
		} // END if
	} // END for
	return p_HTML.TagCloseTop() // tr
} // END TrTh

// Create table data-row from a variable number of DataItems:
//  - p_TrClass - CSS classname for <tr>
//  - p_TdClass - CSS classname for <td>
//  - p_DataItems - Any data-type that is printable with fmt.Sprint; pointer-types will be dereferenced
//
func (p_HTML *T_HTML) TrTd(p_TrClass, p_TdClass string, p_DataItems ...any) *T_HTML {
	p_HTML.helper_TrOpen(p_TrClass)
	for _,DataItem := range(p_DataItems) {
		if reflect.TypeOf(DataItem).Kind() != reflect.Ptr {
 		  p_HTML.helper_Td(fmt.Sprint(DataItem),p_TdClass,"","")
		} else {
 		  p_HTML.helper_Td(fmt.Sprint(reflect.ValueOf(DataItem).Elem().Interface()),p_TdClass,"","")
		} // END if
	} // END for
	return p_HTML.TagCloseTop() // tr
} // END TrTd

// Create table header-row from the field-names of a struct{}
//  - p_TrClass - CSS classname for <tr>
//  - p_ThClass - CSS classname for <th>
//  - p_KeyColHeader - The name of an additional key column when working with map[]struct{}
//  - p_DataItem - Any struct{}-type
//
// - Non exported fields of the struct{} are skipped
// - html struct-tags are interpreted
func (p_HTML *T_HTML) TrThStruct(p_TrClass, p_ThClass, p_KeyColHeader string, p_DataItem any) *T_HTML {
	SType := reflect.TypeOf(p_DataItem)
	if SType.Kind() != reflect.Struct { panic(fmt.Sprintf("Unknown datatype used in TrThStruct: %s",SType)) }
	
	p_HTML.helper_TrOpen(p_TrClass)
	if p_KeyColHeader != "" {
		p_HTML.helper_Th(p_KeyColHeader,p_ThClass,"","") // Add key-column header for map[]struct{}
	} // END if
	for index := 0; index < SType.NumField(); index++ {
		FType := SType.Field(index)
		if FType.Name[0] < 'A' || FType.Name[0] >'Z' { continue } 												// skip private field
 		ColHeader,HeaderClass,_,_,Skip := analyzeHtmlStructTag(FType.Tag.Get("html")) // analyze html struct-tag
	  if Skip { continue } // skip field
	  if ColHeader != "" {
	    p_HTML.helper_Th(ColHeader,p_ThClass,HeaderClass,"")
		} else {
	    p_HTML.helper_Th(FType.Name,p_ThClass,HeaderClass,"")
		} // END if
	} // END for index
	
	return p_HTML.TagCloseTop() // tr
} // END TrThStruct

// Create table header-row from the field-names of a struct{}
//  - p_TrClass - CSS classname for <tr>
//  - p_ThClass - CSS classname for <th>
//  - p_KeyColHeader - The name of an additional key column when working with map[]struct{}
//  - p_DataItem - Any struct{}-type
//
// Non exported fields of the structure are skipped
// - html struct-tags are interpreted
// - Pointer fields are dereferenced and nil is replaced with "&nbsp;"
//
func (p_HTML *T_HTML) TrTdStruct(p_TrClass, p_TdClass, p_KeyColValue string, p_DataItem any) *T_HTML {
	SType := reflect.TypeOf(p_DataItem)
	if SType.Kind() != reflect.Struct { panic(fmt.Sprintf("Unknown datatype used in TrThStruct: %s",SType)) }
	
	p_HTML.helper_TrOpen(p_TrClass)
	if p_KeyColValue != "" {
		p_HTML.helper_Td(p_KeyColValue,p_TdClass,"","") // Add key-column value for map[]struct{}
	} // END if
	
	SValues := reflect.ValueOf(p_DataItem)
	for index := 0; index < SType.NumField(); index++ {
		FType := SType.Field(index)
 	  if FType.Name[0] < 'A' || FType.Name[0] >'Z' { continue } 							// skip private field
		_,_,DataClass,Style,Skip := analyzeHtmlStructTag(FType.Tag.Get("html")) // analyze html struct-tag
		//fmt.Printf("Fieldname='%s'; DataClass='%s',Style='%s'; Skip='%t'\n",FType.Name,DataClass,Style,Skip)
		//fmt.Printf("htmltag=»%s«\n",FType.Tag.Get("html"))
		if Skip { continue } // skip field
		Field := SValues.Field(index)
		if Field.Kind() == reflect.Ptr {
			if Field.IsNil() {
			  p_HTML.helper_Td("&nbsp;",p_TdClass,DataClass,Style)
			} else {
			  p_HTML.helper_Td(fmt.Sprint(Field.Elem().Interface()),p_TdClass,DataClass,Style)
			} // END if
    } else {
		  p_HTML.helper_Td(fmt.Sprint(Field.Interface()),p_TdClass,DataClass,Style)
    } // END if
	} // END for index
	
	return p_HTML.TagCloseTop() // tr
} // END TrTdStruct

// Convert a map into HTML-table rows.
// The following map types can be processed:
//  - map[KeyType]any → two columns: Key, Value
//  - map[KeyType]*any → two columns: Key, Value
//  - map[KeyType]struct{} → Key-Column and one column per exported struct-field
//  - map[KeyType]*struct{} → Key-Column and one column per exported struct-field
//  - map[KeyType][]any → Key-Column and one column per slice element
//  - map[KeyType][]*any → Key-Column and one column per slice element
//
// Keys and values are being converted into strings using fmt.Sprint. 
// Composite types will be converted into strings according to their Stringer interface which might not be the desired result.
//
func (p_HTML *T_HTML) TrTdMap(p_TrClass, p_TdClass string, p_CompareFunc t_CompareFunc, p_DataItems any) *T_HTML {
	DataItems := reflect.ValueOf(p_DataItems)
	if DataItems.Kind() != reflect.Map { panic(fmt.Sprintf("Unknown datatype used in TrTdMapSimple: %s",DataItems)) }
	
	Keys := DataItems.MapKeys()
	if p_CompareFunc != nil {
		slices.SortFunc(Keys,p_CompareFunc)
	} // END if
	for _,key := range Keys {
		if DataItems.MapIndex(key).Kind() == reflect.Ptr {
			if DataItems.MapIndex(key).Elem().Kind() == reflect.Struct {
			  // map[KeyType]*struct{}
        p_HTML.TrTdStruct(p_TrClass, p_TdClass,fmt.Sprint(key.Interface()), DataItems.MapIndex(key).Elem().Interface()) // print struct{} fields as columns
			} else {
 				// Assumption: map[KeyType]*SimpleType
	      p_HTML.helper_TrOpen(p_TrClass).helper_Td(fmt.Sprint(key),p_TdClass,"","").helper_Td(fmt.Sprint(DataItems.MapIndex(key).Elem()),p_TdClass,"","").TagCloseTop() 
			} // END if
		} else {
			switch DataItems.MapIndex(key).Kind() {
				case reflect.Struct: { // map[KeyType]struct{}
          p_HTML.TrTdStruct(p_TrClass, p_TdClass,fmt.Sprint(key.Interface()), DataItems.MapIndex(key).Interface()) // print struct{} into columns
				} // END case
				case reflect.Slice: { // map[KeyType][]any
					RowLen := DataItems.MapIndex(key).Len()
					Arguments := make([]any,RowLen+1,RowLen+1)
					Arguments[0] = key.Interface()
					for column := 0; column < RowLen; column ++ {
						if DataItems.MapIndex(key).Index(column).Kind() != reflect.Ptr {
						  Arguments[column+1] = DataItems.MapIndex(key).Index(column).Interface()
						} else {
						  Arguments[column+1] = DataItems.MapIndex(key).Index(column).Elem().Interface()
						} // END if
					} // END if
					p_HTML.TrTd(p_TrClass,p_TdClass,Arguments...)
				} // END case
				default: { // Assumption: map[KeyType]SimpleType
					p_HTML.helper_TrOpen(p_TrClass).helper_Td(fmt.Sprint(key),p_TdClass,"","").helper_Td(fmt.Sprint(DataItems.MapIndex(key)),p_TdClass,"","").TagCloseTop() 
				} // END default:
			} // END switch
		} // END if
	} // END for

	return p_HTML
} // END TrTdMap

	
// Convert a slice into HTML-table row(s).
// The following slice types can be processed:
//  - []any → one row with a column for each element of the slice
//  - []*any → one row with a column for each element of the slice
//  - [][]any → two dimensional table
//  - [][]*any → two dimensional table
//  - []*[]*any →  two dimensional table
//  - []struct → Table with one column per struct-field
//  - []*struct → Table with one column per struct-field
//
// Values are being converted into strings using fmt.Sprint, so composite types will be converted into strings according to their Stringer interface.
//
func (p_HTML *T_HTML) TrTdSlice(p_TrClass, p_TdClass string, p_DataRows any) *T_HTML {
	DataRows := reflect.ValueOf(p_DataRows)
	if DataRows.Kind() != reflect.Slice { panic(fmt.Sprintf("Unknown datatype used in TrTdSliceStruct: %T",DataRows)) }
	
	NumRows := DataRows.Len() 
	if NumRows == 0 { 
	  fmt.Fprintln(os.Stderr,"TrTdSlice: Warning slice has no elements.")
	  return  p_HTML 
	} // empty slice
	
	EleType := DataRows.Index(0).Kind() 
	if EleType == reflect.Ptr {
		EleType = DataRows.Index(0).Elem().Kind()
	} // END if
  switch EleType {
		case reflect.Struct: { // []struct or []*struct
	  fmt.Println("TrTdSlice: []struct or []*struct")
			for row := 0; row < NumRows; row++ {
				if DataRows.Index(row).Kind() == reflect.Ptr {
					if !DataRows.Index(row).IsNil() {
			      p_HTML.TrTdStruct(p_TrClass,p_TdClass,"",DataRows.Index(row).Elem().Interface())
					} // END if
				} else {
			    p_HTML.TrTdStruct(p_TrClass,p_TdClass,"",DataRows.Index(row).Interface())
				} // END if
			} // END for
		} // END case
		case reflect.Slice: { // [][]any or []*[]any ¡recursion!
      for row := 0; row < NumRows; row++ {
				if DataRows.Index(row).Kind() == reflect.Ptr {
		      p_HTML.TrTdSlice(p_TrClass,p_TdClass,DataRows.Index(row).Elem().Interface())
				} else {
			    p_HTML.TrTdSlice(p_TrClass,p_TdClass,DataRows.Index(row).Interface())
				} // END if
			} // END for
		} // END case
		default: { // assuming []any
			p_HTML.helper_TrOpen(p_TrClass)
			for column := 0; column < NumRows; column++  {
				if DataRows.Index(column).Kind() == reflect.Ptr {
					if DataRows.Index(column).IsNil() {
					  p_HTML.helper_Td("&nbsp;",p_TdClass,"","")
					} else {
					  p_HTML.helper_Td(fmt.Sprint(DataRows.Index(column).Interface()),p_TdClass,"","")
					} // END if
				} else {
					p_HTML.helper_Td(fmt.Sprint(DataRows.Index(column).Interface()),p_TdClass,"","")
				} // END if
			} // END for
	  } // END default
	} // END switch		

	return p_HTML
} // END TrTdSlice

//
// Convert the resultset of an SQL-Query into HTML-table header-row(s).
//
func (p_HTML *T_HTML) TrThSqlRows(p_TrClass, p_ThClass string, p_DataRows *sql.Rows) *T_HTML {
	if ColumnNames, err := p_DataRows.Columns(); err != nil {
		panic(err)
	} else {
		ColumnHeaders := make([]any,len(ColumnNames),len(ColumnNames))
		for i := range ColumnNames {
			ColumnHeaders[i] = any(ColumnNames[i])
		} // END for
		return p_HTML.TrTh(p_TrClass,p_ThClass,ColumnHeaders...)
	} // END if
} // END TrTdSqlRows

//
// Convert the resultset of an SQL-Query into HTML-table row(s).
//
func (p_HTML *T_HTML) TrTdSqlRows(p_TrClass, p_TdClass string, p_DataRows *sql.Rows) *T_HTML {
	ColumnNames, err := p_DataRows.Columns()
	if err != nil { panic(err ) }
	
	NumberOfColumns := len(ColumnNames)
	for p_DataRows.Next() {
		RowPointers := make([]any,NumberOfColumns)
		RowValues   := make([]any,NumberOfColumns)
		for index := range RowValues {
			RowPointers[index] = &RowValues[index]
		} // END for index
		
		if err := p_DataRows.Scan(RowPointers...); err != nil { panic(err) }
		p_HTML.TrTdSlice(p_TrClass,p_TdClass,RowValues)
		
	} // END for
	return p_HTML
} // END TrTdSqlRows
