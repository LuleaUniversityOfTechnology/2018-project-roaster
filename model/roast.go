// Package model roast implementation.
package model

import (
	"log"
	"math"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	errorCost   = 0.8 // Errors accounts for 80 % of points penalty.
	warningCost = 0.2 // Warnings accounts for 20 % of points penalty.
)

// RoastMessage represents a general Roast message.
type RoastMessage struct {
	Hash        uuid.UUID `json:"hash"`
	Row         uint      `json:"row"`
	Column      uint      `json:"column"`
	Engine      string    `json:"engine"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// RoastError represents a Roast error message.
type RoastError struct {
	RoastMessage
}

// RoastWarning represents a Roast error message.
type RoastWarning struct {
	RoastMessage
}

// RoastStatistics holds statistics for a Roast.
type RoastStatistics struct {
	LinesOfCode      uint64 `json:"linesOfCode"`
	NumberOfErrors   uint64 `json:"numberOfErrors"`
	NumberOfWarnings uint64 `json:"numberOfWarnings"`
}

// RoastResult represent a Roast result.
type RoastResult struct {
	Username   string          `json:"username"`
	Code       string          `json:"code"`
	Score      uint            `json:"score"`
	Language   string          `json:"language"`
	Errors     []RoastError    `json:"errors"`
	Warnings   []RoastWarning  `json:"warnings"`
	CreateTime time.Time       `json:"createTime"`
	Statistics RoastStatistics `json:"statistics"`
}

// AddError adds an error to the RoastResult.
func (r *RoastResult) AddError(hash uuid.UUID, row, column uint, engine, name, description string) {
	r.Errors = append(r.Errors, RoastError{
		RoastMessage{
			Hash:        hash,
			Row:         row,
			Column:      column,
			Engine:      engine,
			Name:        name,
			Description: description,
		},
	})
	r.Statistics.NumberOfErrors++
	r.calculateScore()
}

// AddWarning adds an warning to the RoastResult.
func (r *RoastResult) AddWarning(hash uuid.UUID, row, column uint, engine, name, description string) {
	r.Warnings = append(r.Warnings, RoastWarning{
		RoastMessage{
			Hash:        hash,
			Row:         row,
			Column:      column,
			Engine:      engine,
			Name:        name,
			Description: description,
		},
	})
	r.Statistics.NumberOfWarnings++
	r.calculateScore()
}

// calculateSLOC implements a naïve line counter for code.
// All newlines are counted, so even empty rows and comments are counted.
//
// Note: Yes this will toast the CPU and fill the RAM for large code snippets -
// it should have been implemented using streams and stuff. But this would
// require some refactoring so the code is passed as a stream through out the
// system.
func (r *RoastResult) calculateSLOC() {
	n := uint64(strings.Count(r.Code, "\n"))
	if len(r.Code) > 0 && !strings.HasSuffix(r.Code, "\n") {
		n++
	}
	r.Statistics.LinesOfCode = n
	return
}

// calculateScore calculates the score according to a not-so-smart algorithm.
func (r *RoastResult) calculateScore() {
	numErrors := float64(len(r.Errors))
	numWarnings := float64(len(r.Warnings))

	r.Score = uint(math.Round(float64(r.Statistics.LinesOfCode) /
		(((errorCost * numErrors) + (warningCost * numWarnings)) + 1)))
}

// NewRoastResult creates a new RoastResult with username, language and code but
// without warning/error messages.
func NewRoastResult(username, language, code string) (result *RoastResult) {
	result = &RoastResult{
		Username: username,
		Language: language,
		Code:     code,
	}

	result.calculateSLOC()
	result.calculateScore()

	return
}

// GetRoast returns the RoastResult for the specific ID.
func GetRoast(id int) (roast *RoastResult, err error) {
	err = database.QueryRow(`
		SELECT username, score, language, create_time
		FROM "roaster"."roast"
		WHERE id=$1
	`, id).Scan(&roast.Username, &roast.Score, &roast.Language, &roast.CreateTime)
	if err != nil {
		return
	}

	errorRows, err := database.Query(`
		SELECT e.name, e.description
		FROM "roaster"."roast_has_errors" rhe
		INNER JOIN "roaster"."error" e
			ON rhe.error = e.id
		WHERE rhe.roast = $1
	`, id)
	if err != nil {
		return
	}
	defer errorRows.Close()

	for errorRows.Next() {
		msg := RoastError{}
		err = errorRows.Scan(&msg.Hash, &msg.Row, &msg.Column, &msg.Engine, &msg.Name, &msg.Description)
		roast.Errors = append(roast.Errors, msg)
		if err != nil {
			return
		}
	}

	warningRows, err := database.Query(`
		SELECT w.name, w.description
		FROM "roaster"."roast_has_warnings" rhw
		INNER JOIN "roaster"."warning" w
			ON rhw.warning = w.id
		WHERE rhw.roast = $1
	`, id)
	if err != nil {
		return
	}
	defer warningRows.Close()

	for warningRows.Next() {
		msg := RoastWarning{}
		err = warningRows.Scan(&msg.Hash, &msg.Row, &msg.Column, &msg.Engine, &msg.Name, &msg.Description)
		roast.Warnings = append(roast.Warnings, msg)
	}

	return
}

// PutRoast adds a RoastResult to the database.
func PutRoast(roast *RoastResult) (err error) {
	tx, err := database.Begin()
	if err != nil {
		return
	}
	defer func() { // Rollback transaction on error.
		if err != nil {
			rollBackErr := tx.Rollback()
			if rollBackErr != nil {
				log.Println(rollBackErr)
			}

			return
		}
		err = tx.Commit()
	}()

	var roastID int
	err = tx.QueryRow(`
		INSERT INTO "roaster"."roast"
		(username, code, score, language, create_time)
		VALUES
		(TRIM($1), $2, $3, $4, $5)
 		RETURNING id
		`,
		roast.Username,
		roast.Code,
		roast.Score,
		roast.Language,
		time.Now()).Scan(&roastID)
	if err != nil {
		return
	}

	_, err = tx.Exec(`
		INSERT INTO "roaster"."roast_statistics"
		(roast, lines_of_code, number_of_errors, number_of_warnings)
		VALUES
		($1, $2, $3, $4)
		`,
		roastID,
		roast.Statistics.LinesOfCode,
		roast.Statistics.NumberOfErrors,
		roast.Statistics.NumberOfWarnings)
	if err != nil {
		return
	}

	errorInsertStmt, err := tx.Prepare(`
		INSERT INTO "roaster"."error"
		(id, row, "column", engine, name, description)
		VALUES
		($1, $2, $3, $4, $5, $6)
		ON CONFLICT DO NOTHING`)
	if err != nil {
		return
	}
	defer errorInsertStmt.Close()

	roastHasErrorsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roaster"."roast_has_errors"
		(roast, error)
		VALUES
		($1, $2)`)
	if err != nil {
		return
	}
	defer roastHasErrorsInsertStmt.Close()

	for _, errorMessage := range roast.Errors {
		_, err := errorInsertStmt.Exec(
			errorMessage.Hash,
			errorMessage.Row,
			errorMessage.Column,
			errorMessage.Engine,
			errorMessage.Name,
			errorMessage.Description)
		if err != nil {
			return err
		}

		_, err = roastHasErrorsInsertStmt.Exec(roastID, errorMessage.Hash)
		if err != nil {
			return err
		}
	}

	warningInsertStmt, err := tx.Prepare(`
		INSERT INTO "roaster"."warning"
		(id, row, "column", engine, name, description)
		VALUES
		($1, $2, $3, $4, $5, $6)
		ON CONFLICT DO NOTHING`)
	if err != nil {
		return
	}
	defer warningInsertStmt.Close()

	roastHasWarningsInsertStmt, err := tx.Prepare(`
		INSERT INTO "roaster"."roast_has_warnings"
		(roast, warning)
		VALUES
		($1, $2)`)
	if err != nil {
		return
	}
	defer roastHasWarningsInsertStmt.Close()

	for _, warningMessage := range roast.Warnings {
		_, err := warningInsertStmt.Exec(
			warningMessage.Hash,
			warningMessage.Row,
			warningMessage.Column,
			warningMessage.Engine,
			warningMessage.Name,
			warningMessage.Description)
		if err != nil {
			return err
		}

		_, err = roastHasWarningsInsertStmt.Exec(roastID, warningMessage.Hash)
		if err != nil {
			return err
		}
	}

	return
}
