package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// splitCSV splits the input CSV file into multiple parts with the given chunk size.
func splitCSV(inputFileName string, chunkSize int) error {
	// Open the input file.
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer inputFile.Close()

	// Create CSV reader.
	reader := csv.NewReader(inputFile)

	// Read the header.
	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("cannot read header: %w", err)
	}

	baseName := strings.TrimSuffix(filepath.Base(inputFileName), filepath.Ext(inputFileName))
	part := 1

	for {
		outputFileName := fmt.Sprintf("%s.part%d.csv", baseName, part)

		// Write the partial file with the given chunk size.
		err := writePart(outputFileName, header, reader, chunkSize)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("cannot write part: %w", err)
		}

		part++
	}

	return nil
}

// writePart writes a part of the CSV file with the given chunk size.
func writePart(outputFileName string, header []string, reader *csv.Reader, chunkSize int) error {
	firstRecord, err := reader.Read()
	if errors.Is(err, io.EOF) {
		// No more records.
		return err
	}
	if err != nil {
		return fmt.Errorf("cannot read first record: %w", err)
	}

	// Create new output file.
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer outputFile.Close()

	// Create CSV writer.
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header.
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("cannot write header: %w", err)
	}

	// Write the first record.
	if err := writer.Write(firstRecord); err != nil {
		return fmt.Errorf("cannot write first record: %w", err)
	}

	// Write the remaining records.
	for i := 1; i < chunkSize; i++ {
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return err
			}
			return fmt.Errorf("cannot read record: %w", err)
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("cannot write record: %w", err)
		}
	}

	return nil
}
