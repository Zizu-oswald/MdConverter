package convert

import (
	"bytes"
	"context"

	"path/filepath"

	"github.com/Zizu-oswald/MdConverter/logger"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

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

	convertCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Name of file")

	convertCmd.PreRun = func(cmd *cobra.Command, args []string) { // Если имя не задано то использует имя входного файла + .html
		if outputFile == "" {
			inputFile = args[0]
			outputFile = filepath.Base(inputFile)
			outputFile = strings.Split(outputFile, ".")[0] + ".pdf"
		}
	}
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
	err := os.WriteFile("pdf.html", toHTML(), 0644)
	if err != nil {
		logger.HaltOnError(err, "dont write in file")
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	url, _ := filepath.Abs("pdf.html")
	url = "file://" + url
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().Do(ctx)
			if err != nil {
				return err
			}
			return os.WriteFile(outputFile, buf, 0644)
		}),
	)
	if err != nil {
		logger.Warn(err, "error in chromedp.Run")
	}
}
