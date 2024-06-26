package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

// CreateDocumentData is the resolver for the createDocumentData field.
func (r *mutationResolver) CreateDocumentData(ctx context.Context, input generated.CreateDocumentDataInput) (*DocumentDataCreatePayload, error) {
	data, err := withTransactionalMutation(ctx).DocumentData.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			validationError := err.(*generated.ValidationError)

			r.logger.Debugw("validation error", "field", validationError.Name, "error", validationError.Error())

			return nil, validationError
		}

		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			r.logger.Debugw("constraint error", "error", constraintError.Error())

			return nil, constraintError
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionCreate, "document data")
		}

		r.logger.Errorw("failed to create document data", "error", err)

		return nil, err
	}

	return &DocumentDataCreatePayload{DocumentData: data}, nil
}

// CreateBulkDocumentData is the resolver for the createBulkDocumentData field.
func (r *mutationResolver) CreateBulkDocumentData(ctx context.Context, input []*generated.CreateDocumentDataInput) (*DocumentDataBulkCreatePayload, error) {
	return r.bulkCreateDocumentData(ctx, input)
}

// CreateBulkCSVDocumentData is the resolver for the createBulkCSVDocumentData field.
func (r *mutationResolver) CreateBulkCSVDocumentData(ctx context.Context, input graphql.Upload) (*DocumentDataBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateDocumentDataInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateDocumentData(ctx, data)
}

// UpdateDocumentData is the resolver for the updateDocumentData field.
func (r *mutationResolver) UpdateDocumentData(ctx context.Context, id string, input generated.UpdateDocumentDataInput) (*DocumentDataUpdatePayload, error) {
	data, err := withTransactionalMutation(ctx).DocumentData.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get document data on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "document data")
		}

		r.logger.Errorw("failed to get document data", "error", err)
		return nil, ErrInternalServerError
	}

	data, err = data.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update document data", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "document data")
		}

		r.logger.Errorw("failed to update document data", "error", err)
		return nil, ErrInternalServerError
	}

	return &DocumentDataUpdatePayload{DocumentData: data}, nil
}

// DeleteDocumentData is the resolver for the deleteDocumentData field.
func (r *mutationResolver) DeleteDocumentData(ctx context.Context, id string) (*DocumentDataDeletePayload, error) {
	if err := withTransactionalMutation(ctx).DocumentData.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "document data")
		}

		r.logger.Errorw("failed to delete document data", "error", err)
		return nil, err
	}

	if err := generated.DocumentDataEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &DocumentDataDeletePayload{DeletedID: id}, nil
}

// DocumentData is the resolver for the documentData field.
func (r *queryResolver) DocumentData(ctx context.Context, id string) (*generated.DocumentData, error) {
	data, err := withTransactionalMutation(ctx).DocumentData.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get document data", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "document data")
		}

		return nil, ErrInternalServerError
	}

	return data, nil
}
