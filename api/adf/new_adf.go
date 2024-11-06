// Default constructors
package adf

func NewDocNode() *DocNode {
	node := DocNode{Type: "doc"}
	return &node
}

func (dn *DocNode) Add(node any) {
	dn.Content = append(dn.Content, node)
}

func NewParagraphNode() *ParagraphNode {
	node := ParagraphNode{Type: "paragraph"}
	return &node
}

func (pn *ParagraphNode) Add(node any) {
	pn.Content = append(pn.Content, node)
}

func NewTextNode(text string) *TextNode {
	textNode := TextNode{}
	textNode.Text = text
	textNode.Type = "text"
	return &textNode
}

func (hn *HeadingNode) Add(node any) {
	hn.Content = append(hn.Content, node)
}
func NewHeadingNode(level int, text string) *HeadingNode {
	node := HeadingNode{Type: "heading"}
	node.Attrs = HeadingNodeAttrs{Level: 3}
	node.Add(NewTextNode(text))
	return &node
}

func NewHardBreakNode() *HardBreakNode {
	node := HardBreakNode{Type: "hardBreak"}
	return &node
}

func NewTableNode() *TableNode {
	node := TableNode{Type: "table"}
	return &node
}

func (tn *TableNode) Add(node *TableRowNode) {
	tn.Content = append(tn.Content, node)
}

func NewTableRowNode() *TableRowNode {
	node := TableRowNode{Type: "tableRow"}
	return &node
}

func (tr *TableRowNode) Add(node any) {
	tr.Content = append(tr.Content, node)
}

func NewTableHeader(text string) *TableHeaderNode {
	node := TableHeaderNode{Type: "tableHeader"}
	node.Attrs = &TableHeaderNodeAttrs{}
	col, row := 1.0, 1.0
	node.Attrs.Colspan = &col
	node.Attrs.Rowspan = &row
	pn := NewParagraphNode()
	node.Add(pn)
	pn.Add(NewTextNode(text))
	return &node
}

func (tr *TableHeaderNode) Add(node any) {
	tr.Content = append(tr.Content, node)
}

func NewTableCell(text string) *TableCellNode {
	node := TableCellNode{Type: "tableCell"}
	node.Attrs = &TableCellNodeAttrs{}
	col, row := 1.0, 1.0
	node.Attrs.Colspan = &col
	node.Attrs.Rowspan = &row
	pn := NewParagraphNode()
	node.Add(pn)
	pn.Add(NewTextNode(text))
	return &node
}

func (tr *TableCellNode) Add(node any) {
	tr.Content = append(tr.Content, node)
}
