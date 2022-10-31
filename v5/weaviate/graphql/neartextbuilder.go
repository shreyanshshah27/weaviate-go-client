package graphql

import (
	"fmt"
	"strings"
)

// fldMover is a type representing field names of a move sub query
type fldMover string

const (
	fldMoverConcepts fldMover = "concepts"
	fldMoverForce    fldMover = "force"
	fldMoverID       fldMover = "id"
	fldMoverBeacon   fldMover = "beacon"
	fldMoverObjects  fldMover = "objects"
)

// MoveParameters to fine tune Explore queries
type MoveParameters struct {
	// Concepts that should be used as base for the movement operation
	Concepts []string
	// Force to be applied in the movement operation
	Force float32
	// Objects used to adjust the serach direction
	Objects []MoverObject
}

func (m *MoveParameters) String() string {
	concepts := marshalStrings(m.Concepts)
	ms := make([]string, 0, len(m.Objects))
	for _, m := range m.Objects {
		if s := m.String(); s != EmptyObjectStr {
			ms = append(ms, s)
		}
	}
	if len(ms) < 1 {
		return fmt.Sprintf("{%s: %s %s: %v}", fldMoverConcepts, concepts, fldMoverForce, m.Force)
	}

	s := "{"
	if len(m.Concepts) > 0 {
		s = fmt.Sprintf("{%s: %s", fldMoverConcepts, concepts)
	}
	return fmt.Sprintf("%s %s: %v %s: %v}", s, fldMoverForce, m.Force, fldMoverObjects, ms)
}

// MoverObject is the object the search is supposed to move close to (or further away from) it.
type MoverObject struct {
	ID     string
	Beacon string
}

// String returns string representation of m as {"id": "value" beacon:"value"}.
// Empty fields are considered optional and are excluded.
// It returns EmptyObjectStr if both fields are empty
func (m *MoverObject) String() string {
	if m.ID != "" && m.Beacon != "" {
		return fmt.Sprintf(`{%s: "%s" %s: "%s"}`, fldMoverID, m.ID, fldMoverBeacon, m.Beacon)
	}
	if m.ID != "" {
		return fmt.Sprintf(`{%s: "%s"}`, fldMoverID, m.ID)
	}
	if m.Beacon != "" {
		return fmt.Sprintf(`{%s: "%s"}`, fldMoverBeacon, m.Beacon)
	}
	return EmptyObjectStr
}

type NearTextArgumentBuilder struct {
	concepts        []string
	withCertainty   bool
	certainty       float32
	withDistance    bool
	distance        float32
	moveTo          *MoveParameters
	moveAwayFrom    *MoveParameters
	withAutocorrect bool
	autocorrect     bool
}

// WithConcepts the result is based on
func (e *NearTextArgumentBuilder) WithConcepts(concepts []string) *NearTextArgumentBuilder {
	e.concepts = concepts
	return e
}

// WithCertainty that is minimally required for an object to be included in the result set
func (e *NearTextArgumentBuilder) WithCertainty(certainty float32) *NearTextArgumentBuilder {
	e.withCertainty = true
	e.certainty = certainty
	return e
}

// WithDistance that is minimally required for an object to be included in the result set
func (e *NearTextArgumentBuilder) WithDistance(distance float32) *NearTextArgumentBuilder {
	e.withDistance = true
	e.distance = distance
	return e
}

// WithMoveTo specific concept
func (e *NearTextArgumentBuilder) WithMoveTo(parameters *MoveParameters) *NearTextArgumentBuilder {
	e.moveTo = parameters
	return e
}

// WithMoveAwayFrom specific concept
func (e *NearTextArgumentBuilder) WithMoveAwayFrom(parameters *MoveParameters) *NearTextArgumentBuilder {
	e.moveAwayFrom = parameters
	return e
}

// WithAutocorrect this is a setting enabling autocorrect of the concepts texts
func (e *NearTextArgumentBuilder) WithAutocorrect(autocorrect bool) *NearTextArgumentBuilder {
	e.withAutocorrect = true
	e.autocorrect = autocorrect
	return e
}

// Build build the given clause
func (e *NearTextArgumentBuilder) build() string {
	clause := []string{}
	concepts := marshalStrings(e.concepts)

	clause = append(clause, fmt.Sprintf("concepts: %s", concepts))
	if e.withCertainty {
		clause = append(clause, fmt.Sprintf("certainty: %v", e.certainty))
	}
	if e.withDistance {
		clause = append(clause, fmt.Sprintf("distance: %v", e.distance))
	}
	if e.moveTo != nil {
		clause = append(clause, fmt.Sprintf("moveTo: %s", e.moveTo))
	}
	if e.moveAwayFrom != nil {
		clause = append(clause, fmt.Sprintf("moveAwayFrom: %s", e.moveAwayFrom))
	}
	if e.withAutocorrect {
		clause = append(clause, fmt.Sprintf("autocorrect: %v", e.autocorrect))
	}
	return fmt.Sprintf("nearText:{%v}", strings.Join(clause, " "))
}
