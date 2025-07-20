package grpc

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errToGRPC(err error, code codes.Code) error {
	switch e := err.(type) {
	case mongo.WriteException:
		for _, we := range e.WriteErrors {
			if we.Code == 11000 { // Duplicate key error
				return status.Error(codes.AlreadyExists, "duplicate key error: "+we.Message)
			}
		}
		return status.Error(codes.Internal, e.Error())
	case mongo.CommandError:
		if e.HasErrorCode(11000) {
			return status.Error(codes.AlreadyExists, "duplicate key error: "+e.Message)
		}
		return status.Error(codes.Internal, e.Error())
	case mongo.WriteConcernError:
		return status.Error(codes.Unavailable, e.Error())
	case mongo.BulkWriteException:
		for _, bwe := range e.WriteErrors {
			if bwe.Code == 11000 {
				return status.Error(codes.AlreadyExists, "duplicate key error: "+bwe.Message)
			}
		}
		return status.Error(codes.Internal, e.Error())
	}

	switch {
	case err == mongo.ErrNoDocuments:
		return status.Error(codes.NotFound, "document not found")
	case code != codes.Internal:
		return status.Error(code, err.Error())
	}

	return status.Error(codes.Internal, err.Error())
}
