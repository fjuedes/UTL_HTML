package UTL_HTML
//
// UTL_HTML_Form
// Version: $Id: UTL_HTML_Form.go 82 2025-05-02 19:59:27Z fjuedes $
//
import (
	"fmt"
	"reflect"
	"slices"
)

func (p_HTML *T_HTML) FormOpen(p_Action, p_Method string, p_Attributes ...string) *T_HTML {
  Arguments := make([]string,0,len(p_Attributes)+2)
	appendAttribute("action",p_Action,&Arguments)
	appendAttribute("method",p_Method,&Arguments)
	Arguments = append(Arguments, p_Attributes...)
	return p_HTML.TagOpen("form", Arguments...)
} // END OpenForm

func (p_HTML *T_HTML) HiddenField(p_name, p_value string, p_Attributes... string ) *T_HTML {
	Arguments := []string{"type", "hidden", "name", p_name, "value", p_value}
	Arguments = append(Arguments, p_Attributes...)
	return p_HTML.Tag("input", "", Arguments...)
} // END HiddenField

func (p_HTML *T_HTML) TextField(p_name, p_size, p_maxlen, p_value string, p_Attributes ...string) *T_HTML {
  Arguments := make([]string,0,len(p_Attributes)+4)
	appendAttribute("type","text",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	appendAttribute("size",p_size,&Arguments)
	appendAttribute("maxlen",p_maxlen,&Arguments)
	appendAttribute("value",p_value,&Arguments)
	Arguments = append(Arguments, p_Attributes...)
	return p_HTML.Tag("input", "", Arguments...)
} // END TextField

func TextField(p_name, p_size, p_maxlen, p_value string, p_Attributes ...string) string {
  Arguments := make([]string,0,len(p_Attributes)+4)
	appendAttribute("type","text",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	appendAttribute("size",p_size,&Arguments)
	appendAttribute("maxlen",p_maxlen,&Arguments)
	appendAttribute("value",p_value,&Arguments)
	Arguments = append(Arguments, p_Attributes...)
	return Tag("input", "", Arguments...)
} // END TextField

func (p_HTML *T_HTML) BoolField(p_name string, p_checked bool, p_Attributes ...string) *T_HTML {
  Arguments := make([]string,0,len(p_Attributes)+1)
	appendAttribute("type","checkbox",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	if p_checked {
		Arguments = append(Arguments,"checked","checked")
	} // END if
	Arguments = append(Arguments, p_Attributes...)
	return p_HTML.Tag("input", "", Arguments...)
} // END BoolField

func BoolField(p_name string, p_checked bool, p_Attributes ...string) string {
  Arguments := make([]string,0,len(p_Attributes)+3)
	appendAttribute("type","checkbox",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	if p_checked {
		Arguments = append(Arguments,"checked","checked")
	} // END if
	Arguments = append(Arguments, p_Attributes...)
	return Tag("input", "", Arguments...)
} // END BoolField

// Append a select menu, with the menu-items from a map[string]string
func (p_HTML *T_HTML) SelectMenu(p_FieldName, p_MenuClassName, p_ItemClassName, p_DefaultValue string, p_CompareFunc t_CompareFunc, p_MenuItems any, p_Attributes ...any) *T_HTML {
	DataItems := reflect.ValueOf(p_MenuItems)
	if DataItems.Kind() != reflect.Map { panic(fmt.Sprintf("Unknown datatype used in TrTdMapSimple: %s",DataItems)) }

	Arguments := make([]string,0,len(p_Attributes)+2)
	appendAttribute("class",p_MenuClassName,&Arguments)
	appendAttribute("name",p_FieldName,&Arguments)
	p_HTML.TagOpen("select",Arguments...)

	Keys := DataItems.MapKeys()
	if p_CompareFunc != nil {
		slices.SortFunc(Keys,p_CompareFunc)
	} // END if
	for _,Key := range Keys {
		Value := ""
		if DataItems.MapIndex(Key).Kind() == reflect.Ptr {
			Value = fmt.Sprint(DataItems.MapIndex(Key).Elem())
		} else {
			Value = fmt.Sprint(DataItems.MapIndex(Key))
		} // END if
		if Value == p_DefaultValue {
		  p_HTML.Tag("option",fmt.Sprint(Key),"value",Value,"class","selected","selected",p_ItemClassName)
		} else {
		  p_HTML.Tag("option",fmt.Sprint(Key),"value",Value,"class",p_ItemClassName)
		} // END if
  } // END for
	return p_HTML
} // END SelectMenu


func (p_HTML *T_HTML) SubmitButton(p_name, p_label, p_value string, p_Attributes ...string) *T_HTML {
  Arguments := make([]string,0,len(p_Attributes)+3)
	appendAttribute("type","submit",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	appendAttribute("value",p_value,&Arguments)
	Arguments = append(Arguments, p_Attributes...)
	return p_HTML.Tag("button",p_label , Arguments...)
} // END SubmitButton

func SubmitButton(p_name, p_label, p_value, p_Class string, p_Attributes ...string) string {
  Arguments := make([]string,0,len(p_Attributes)+3)
	appendAttribute("type","submit",&Arguments)
	appendAttribute("name",p_name,&Arguments)
	appendAttribute("value",p_value,&Arguments)
	Arguments = append(Arguments, p_Attributes...)
	return Tag("button", p_label, Arguments...)
} // END Button

