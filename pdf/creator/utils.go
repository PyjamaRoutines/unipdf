/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package creator

import (
	"math"
	"os"

	"github.com/unidoc/unidoc/pdf/model"
)

// Loads the template from path as a list of pages.
func loadPagesFromFile(path string) ([]*model.PdfPage, error) {
	// Read the input pdf file.
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}

	// Load the pages.
	pages := []*model.PdfPage{}
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	return pages, nil
}

func round(val float64, roundOn float64, places int) float64 {
	var round float64

	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	div = math.Copysign(div, val)

	roundOn = math.Copysign(roundOn, val)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return round / pow
}
