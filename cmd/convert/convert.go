package convert

import (
	"Md-converter/logger"
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"os"
	"strings"
)

var (
	inputFile  string
	outputFile string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert into other format",
	Long:  "Convert into other format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile = args[0]
		convertFunc()
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&outputFile, "output", "o", "markdown.pdf", "Name of file")
}

func convertFunc() {
	fileSplit := strings.Split(outputFile, ".")
	switch fileSplit[1] {
	case "html":
		err := os.WriteFile(outputFile, toHTML(), 0644)
		if err != nil {
			logger.HaltOnError(err, "dont write in file")
		}
	case "pdf":
		toPDF()
	}

}

func toHTML() []byte {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	source, err := os.ReadFile(inputFile)
	if err != nil {
		logger.HaltOnError(err, "dont read file")
	}
	var buf bytes.Buffer
	err = md.Convert(source, &buf)
	if err != nil {
		logger.HaltOnError(err, "dont convert")
	}

	return buf.Bytes()
}

func toPDF() {
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		logger.HaltOnError(err, "dont make pdf generator")
	}

	pdf.Dpi.Set(600)
	pdf.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdf.Grayscale.Set(true)

	err = os.WriteFile("pdf.html", toHTML(), 0644)
	if err != nil {
		logger.HaltOnError(err, "dont write in file")
	}
	defer os.Remove("pdf.html")

	page := wkhtmltopdf.NewPage("pdf.html")

	pdf.AddPage(page)
	err = pdf.Create()
	if err != nil {
		logger.HaltOnError(err, "dont create pdf buffer")
	}

	err = pdf.WriteFile(outputFile)
	if err != nil {
		logger.HaltOnError(err, "dont write in pdf file")
	}
}
