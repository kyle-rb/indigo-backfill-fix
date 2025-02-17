// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.repo.applyWrites

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// RepoApplyWrites_Create is a "create" in the com.atproto.repo.applyWrites schema.
//
// Create a new record.
//
// RECORDTYPE: RepoApplyWrites_Create
type RepoApplyWrites_Create struct {
	LexiconTypeID string                   `json:"$type,const=com.atproto.repo.applyWrites#create" cborgen:"$type,const=com.atproto.repo.applyWrites#create"`
	Collection    string                   `json:"collection" cborgen:"collection"`
	Rkey          *string                  `json:"rkey,omitempty" cborgen:"rkey,omitempty"`
	Value         *util.LexiconTypeDecoder `json:"value" cborgen:"value"`
}

// RepoApplyWrites_Delete is a "delete" in the com.atproto.repo.applyWrites schema.
//
// Delete an existing record.
//
// RECORDTYPE: RepoApplyWrites_Delete
type RepoApplyWrites_Delete struct {
	LexiconTypeID string `json:"$type,const=com.atproto.repo.applyWrites#delete" cborgen:"$type,const=com.atproto.repo.applyWrites#delete"`
	Collection    string `json:"collection" cborgen:"collection"`
	Rkey          string `json:"rkey" cborgen:"rkey"`
}

// RepoApplyWrites_Input is the input argument to a com.atproto.repo.applyWrites call.
type RepoApplyWrites_Input struct {
	// repo: The handle or DID of the repo.
	Repo       string  `json:"repo" cborgen:"repo"`
	SwapCommit *string `json:"swapCommit,omitempty" cborgen:"swapCommit,omitempty"`
	// validate: Flag for validating the records.
	Validate *bool                                `json:"validate,omitempty" cborgen:"validate,omitempty"`
	Writes   []*RepoApplyWrites_Input_Writes_Elem `json:"writes" cborgen:"writes"`
}

type RepoApplyWrites_Input_Writes_Elem struct {
	RepoApplyWrites_Create *RepoApplyWrites_Create
	RepoApplyWrites_Update *RepoApplyWrites_Update
	RepoApplyWrites_Delete *RepoApplyWrites_Delete
}

func (t *RepoApplyWrites_Input_Writes_Elem) MarshalJSON() ([]byte, error) {
	if t.RepoApplyWrites_Create != nil {
		t.RepoApplyWrites_Create.LexiconTypeID = "com.atproto.repo.applyWrites#create"
		return json.Marshal(t.RepoApplyWrites_Create)
	}
	if t.RepoApplyWrites_Update != nil {
		t.RepoApplyWrites_Update.LexiconTypeID = "com.atproto.repo.applyWrites#update"
		return json.Marshal(t.RepoApplyWrites_Update)
	}
	if t.RepoApplyWrites_Delete != nil {
		t.RepoApplyWrites_Delete.LexiconTypeID = "com.atproto.repo.applyWrites#delete"
		return json.Marshal(t.RepoApplyWrites_Delete)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *RepoApplyWrites_Input_Writes_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "com.atproto.repo.applyWrites#create":
		t.RepoApplyWrites_Create = new(RepoApplyWrites_Create)
		return json.Unmarshal(b, t.RepoApplyWrites_Create)
	case "com.atproto.repo.applyWrites#update":
		t.RepoApplyWrites_Update = new(RepoApplyWrites_Update)
		return json.Unmarshal(b, t.RepoApplyWrites_Update)
	case "com.atproto.repo.applyWrites#delete":
		t.RepoApplyWrites_Delete = new(RepoApplyWrites_Delete)
		return json.Unmarshal(b, t.RepoApplyWrites_Delete)

	default:
		return fmt.Errorf("closed enums must have a matching value")
	}
}

// RepoApplyWrites_Update is a "update" in the com.atproto.repo.applyWrites schema.
//
// Update an existing record.
//
// RECORDTYPE: RepoApplyWrites_Update
type RepoApplyWrites_Update struct {
	LexiconTypeID string                   `json:"$type,const=com.atproto.repo.applyWrites#update" cborgen:"$type,const=com.atproto.repo.applyWrites#update"`
	Collection    string                   `json:"collection" cborgen:"collection"`
	Rkey          string                   `json:"rkey" cborgen:"rkey"`
	Value         *util.LexiconTypeDecoder `json:"value" cborgen:"value"`
}

// RepoApplyWrites calls the XRPC method "com.atproto.repo.applyWrites".
func RepoApplyWrites(ctx context.Context, c *xrpc.Client, input *RepoApplyWrites_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.repo.applyWrites", nil, input, nil); err != nil {
		return err
	}

	return nil
}
