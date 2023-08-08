package models

import (
	"github.com/jung-kurt/gofpdf"
)

type PDFFileBuilder struct {
	Title           string
	Content         string
	Author          string
	Font            FontType
	FontSize        int16
	Header          HeaderType
	Footer          FooterType
	PageOrientation bool
	Margings        float32
	Image           string
}

func NewPDFFileBuilder() *PDFFileBuilder {
	return &PDFFileBuilder{
		Font:            Arial,
		Header:          NormalHeader,
		Footer:          NormalFooter,
		PageOrientation: true,
	}
}

func (b *PDFFileBuilder) SetTitle(title string) *PDFFileBuilder {
	b.Title = title
	return b
}

func (b *PDFFileBuilder) SetContent(content string) *PDFFileBuilder {
	b.Content = content
	return b
}

func (b *PDFFileBuilder) SetAuthor(author string) *PDFFileBuilder {
	b.Author = author
	return b
}

func (b *PDFFileBuilder) SetFont(font FontType) *PDFFileBuilder {
	b.Font = font
	return b
}

func (b *PDFFileBuilder) SetFontSize(fontSize int16) *PDFFileBuilder {
	b.FontSize = fontSize
	return b
}

func (b *PDFFileBuilder) SetHeader(header HeaderType) *PDFFileBuilder {
	b.Header = header
	return b
}

func (b *PDFFileBuilder) SetFooter(footer FooterType) *PDFFileBuilder {
	b.Footer = footer
	return b
}

func (b *PDFFileBuilder) SetPageOrientation(pageOrientation bool) *PDFFileBuilder {
	b.PageOrientation = pageOrientation
	return b
}

func (b *PDFFileBuilder) SetMargings(margings float32) *PDFFileBuilder {
	b.Margings = margings
	return b
}

func (b *PDFFileBuilder) Build() error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, b.Title)

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, b.Content, "", "", false)

	err := pdf.OutputFileAndClose("output.pdf")
	return err
}
