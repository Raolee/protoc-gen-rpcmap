// Code generated by protoc-gen-custom. DO NOT EDIT.
// versions:
// - protoc-gen-custom vv1.0.0
// source: 2pc.proto
package service

import (
	"context"
	"google.golang.org/protobuf/proto"
)

var RPCMap = map[string]func(context.Context, TwoPhaseCommitServiceServer, proto.Message) (proto.Message, error){
	"twophasecommit.PrepareRequest": func(ctx context.Context, s TwoPhaseCommitServiceServer, req proto.Message) (proto.Message, error) {
		msg, _ := req.(*PrepareRequest)
		return s.Prepare(ctx, msg)
	},
	"twophasecommit.CommitRequest": func(ctx context.Context, s TwoPhaseCommitServiceServer, req proto.Message) (proto.Message, error) {
		msg, _ := req.(*CommitRequest)
		return s.Commit(ctx, msg)
	},
	"twophasecommit.RollbackRequest": func(ctx context.Context, s TwoPhaseCommitServiceServer, req proto.Message) (proto.Message, error) {
		msg, _ := req.(*RollbackRequest)
		return s.Rollback(ctx, msg)
	},
	"twophasecommit.ReportRequest": func(ctx context.Context, s TwoPhaseCommitServiceServer, req proto.Message) (proto.Message, error) {
		msg, _ := req.(*ReportRequest)
		return s.Report(ctx, msg)
	},
}
