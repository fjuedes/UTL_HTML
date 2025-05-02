package UTL_HTML
//
// UTL_HTML_List
// Version: $Id: UTL_HTML_List.go 82 2025-05-02 19:59:27Z fjuedes $
//

import (
	//"fmt"
)

func (p_HTML *T_HTML) OlOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("ol", p_Attributes...) }
func (p_HTML *T_HTML) UlOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("ul", p_Attributes...) }
func (p_HTML *T_HTML) LiOpen(p_Attributes ...string) *T_HTML { return p_HTML.TagOpen("li", p_Attributes...) }
func (p_HTML *T_HTML) Li(p_Content string, p_Attributes ...string) *T_HTML { return p_HTML.Tag("li", p_Content, p_Attributes...) } 
func (p_HTML *T_HTML) Lif(p_Class, p_Format string, p_Data ...any) *T_HTML { return p_HTML.Tagf("li",p_Class,p_Format,p_Data...) }
