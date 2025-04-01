package pg

import (
	"github.com/cxnub/fas-mgmt-system/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// helper to convert nullable pgtype.Timestamp to time.Time
func toTime(valid *pgtype.Timestamp) *time.Time {
	if valid != nil && valid.Valid {
		return &valid.Time
	}
	return nil
}

// helper to convert nullable pgtype.Date to time.Time
func toDate(valid *pgtype.Date) *time.Time {
	if valid != nil && valid.Valid {
		return &valid.Time
	}
	return nil
}

// helper to convert time.Time to nullable pgtype.Timestamp
func fromTime(t *time.Time) *pgtype.Timestamp {
	if t != nil {
		return &pgtype.Timestamp{Time: *t, Valid: true}
	}
	return &pgtype.Timestamp{Valid: false}
}

// helper to convert time.Time to nullable pgtype.Date
func fromDate(t *time.Time) *pgtype.Date {
	if t != nil {
		return &pgtype.Date{Time: *t, Valid: true}
	}
	return &pgtype.Date{Valid: false}
}

func safeUUID(id *uuid.UUID) uuid.UUID {
	if id == nil {
		return uuid.Nil // Return an empty UUID
	}
	return *id
}

func safeString(s *string) string {
	if s == nil {
		return "" // Return empty string instead of nil
	}
	return *s
}

func safeEmploymentStatus(es *domain.EmploymentStatus) EmploymentStatus {
	if es == nil {
		return "" // Default to empty status
	}
	return EmploymentStatus(*es)
}

func safeMaritalStatus(ms *domain.MaritalStatus) MaritalStatus {
	if ms == nil {
		return "" // Default to empty status
	}
	return MaritalStatus(*ms)
}

func safeSex(s *domain.Sex) Sex {
	if s == nil {
		return "" // Default to empty sex
	}
	return Sex(*s)
}

// ==================== Applicant Conversions ====================

func (a *Applicant) ToEntity() *domain.Applicant {
	if a == nil {
		return nil
	}
	return &domain.Applicant{
		ID:               &a.ID,
		Name:             &a.Name,
		EmploymentStatus: (*domain.EmploymentStatus)(&a.EmploymentStatus),
		MaritalStatus:    (*domain.MaritalStatus)(&a.MaritalStatus),
		Sex:              (*domain.Sex)(&a.Sex),
		DateOfBirth:      toDate(&a.DateOfBirth),
		CreatedAt:        toTime(&a.CreatedAt),
		UpdatedAt:        toTime(&a.UpdatedAt),
	}
}

func ApplicantFromEntity(e *domain.Applicant) *Applicant {
	if e == nil {
		return nil
	}
	return &Applicant{
		ID:               safeUUID(e.ID),
		Name:             safeString(e.Name),
		EmploymentStatus: safeEmploymentStatus(e.EmploymentStatus),
		MaritalStatus:    safeMaritalStatus(e.MaritalStatus),
		Sex:              safeSex(e.Sex),
		DateOfBirth:      *fromDate(e.DateOfBirth),
		CreatedAt:        *fromTime(e.CreatedAt),
		UpdatedAt:        *fromTime(e.UpdatedAt),
	}
}

// ==================== Application Conversions ====================

func (a *Application) ToEntity() *domain.Application {
	if a == nil {
		return nil
	}
	return &domain.Application{
		ID:          &a.ID,
		ApplicantID: &a.ApplicantID,
		SchemeID:    &a.SchemeID,
		CreatedAt:   toTime(&a.CreatedAt),
		UpdatedAt:   toTime(&a.UpdatedAt),
	}
}

func ApplicationFromEntity(e *domain.Application) *Application {
	if e == nil {
		return nil
	}
	return &Application{
		ID:          safeUUID(e.ID),
		ApplicantID: safeUUID(e.ApplicantID),
		SchemeID:    safeUUID(e.SchemeID),
		CreatedAt:   *fromTime(e.CreatedAt),
		UpdatedAt:   *fromTime(e.UpdatedAt),
	}
}

// ==================== Benefit Conversions ====================

func (b *Benefit) ToEntity() *domain.Benefit {
	if b == nil {
		return nil
	}
	return &domain.Benefit{
		ID:        &b.ID,
		SchemeID:  &b.SchemeID,
		Name:      &b.Name,
		Amount:    &b.Amount.Float64,
		CreatedAt: toTime(&b.CreatedAt),
		UpdatedAt: toTime(&b.UpdatedAt),
	}
}

func BenefitFromEntity(e *domain.Benefit) *Benefit {
	if e == nil {
		return nil
	}
	return &Benefit{
		ID:        safeUUID(e.ID),
		SchemeID:  safeUUID(e.SchemeID),
		Name:      safeString(e.Name),
		Amount:    pgtype.Float8{Float64: *e.Amount, Valid: true},
		CreatedAt: *fromTime(e.CreatedAt),
		UpdatedAt: *fromTime(e.UpdatedAt),
	}
}

// ==================== BenefitCriterium Conversions ====================

func (bc *BenefitCriterium) ToEntity() *domain.BenefitCriteria {
	if bc == nil {
		return nil
	}
	return &domain.BenefitCriteria{
		ID:        &bc.ID,
		BenefitID: &bc.BenefitID,
		Name:      &bc.Name,
		Value:     &bc.Value.String,
		CreatedAt: toTime(&bc.CreatedAt),
		UpdatedAt: toTime(&bc.UpdatedAt),
	}
}

func BenefitCriteriumFromEntity(e *domain.BenefitCriteria) *BenefitCriterium {
	if e == nil {
		return nil
	}
	return &BenefitCriterium{
		ID:        safeUUID(e.ID),
		BenefitID: safeUUID(e.BenefitID),
		Name:      safeString(e.Name),
		Value:     pgtype.Text{String: safeString(e.Value), Valid: *e.Value != ""},
		CreatedAt: *fromTime(e.CreatedAt),
		UpdatedAt: *fromTime(e.UpdatedAt),
	}
}

// ==================== Relationship Conversions ====================

func (r *Relationship) ToEntity() *domain.Relationship {
	if r == nil {
		return nil
	}
	return &domain.Relationship{
		ID:               &r.ID,
		ApplicantAID:     &r.ApplicantAID,
		ApplicantBID:     &r.ApplicantBID,
		RelationshipType: (*domain.RelationshipType)(&r.RelationshipType),
		CreatedAt:        toTime(&r.CreatedAt),
		UpdatedAt:        toTime(&r.UpdatedAt),
	}
}

func RelationshipFromEntity(e *domain.Relationship) *Relationship {
	if e == nil {
		return nil
	}
	return &Relationship{
		ID:               safeUUID(e.ID),
		ApplicantAID:     safeUUID(e.ApplicantAID),
		ApplicantBID:     safeUUID(e.ApplicantBID),
		RelationshipType: RelationshipType(safeString((*string)(e.RelationshipType))),
		CreatedAt:        *fromTime(e.CreatedAt),
		UpdatedAt:        *fromTime(e.UpdatedAt),
	}
}

// ==================== Scheme Conversions ====================

func (s *Scheme) ToEntity() *domain.Scheme {
	if s == nil {
		return nil
	}
	return &domain.Scheme{
		ID:        &s.ID,
		Name:      &s.Name,
		CreatedAt: toTime(&s.CreatedAt),
		UpdatedAt: toTime(&s.UpdatedAt),
	}
}

func SchemeFromEntity(e *domain.Scheme) *Scheme {
	if e == nil {
		return nil
	}
	return &Scheme{
		ID:        safeUUID(e.ID),
		Name:      safeString(e.Name),
		CreatedAt: *fromTime(e.CreatedAt),
		UpdatedAt: *fromTime(e.UpdatedAt),
	}
}

// ==================== SchemeCriterium Conversions ====================

func (sc *SchemeCriterium) ToEntity() *domain.SchemeCriteria {
	if sc == nil {
		return nil
	}
	return &domain.SchemeCriteria{
		ID:        &sc.ID,
		SchemeID:  &sc.SchemeID,
		Name:      &sc.Name,
		Value:     &sc.Value.String,
		CreatedAt: toTime(&sc.CreatedAt),
		UpdatedAt: toTime(&sc.UpdatedAt),
	}
}

func SchemeCriteriumFromEntity(e *domain.SchemeCriteria) *SchemeCriterium {
	if e == nil {
		return nil
	}
	return &SchemeCriterium{
		ID:        safeUUID(e.ID),
		SchemeID:  safeUUID(e.SchemeID),
		Name:      safeString(e.Name),
		Value:     pgtype.Text{String: safeString(e.Value), Valid: *e.Value != ""},
		CreatedAt: *fromTime(e.CreatedAt),
		UpdatedAt: *fromTime(e.UpdatedAt),
	}
}
