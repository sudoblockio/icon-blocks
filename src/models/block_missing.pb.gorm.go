// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block_missing.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type BlockMissingORM struct {
	Number uint32 `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (BlockMissingORM) TableName() string {
	return "block_missings"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *BlockMissing) ToORM(ctx context.Context) (BlockMissingORM, error) {
	to := BlockMissingORM{}
	var err error
	if prehook, ok := interface{}(m).(BlockMissingWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	if posthook, ok := interface{}(m).(BlockMissingWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BlockMissingORM) ToPB(ctx context.Context) (BlockMissing, error) {
	to := BlockMissing{}
	var err error
	if prehook, ok := interface{}(m).(BlockMissingWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	if posthook, ok := interface{}(m).(BlockMissingWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type BlockMissing the arg will be the target, the caller the one being converted from

// BlockMissingBeforeToORM called before default ToORM code
type BlockMissingWithBeforeToORM interface {
	BeforeToORM(context.Context, *BlockMissingORM) error
}

// BlockMissingAfterToORM called after default ToORM code
type BlockMissingWithAfterToORM interface {
	AfterToORM(context.Context, *BlockMissingORM) error
}

// BlockMissingBeforeToPB called before default ToPB code
type BlockMissingWithBeforeToPB interface {
	BeforeToPB(context.Context, *BlockMissing) error
}

// BlockMissingAfterToPB called after default ToPB code
type BlockMissingWithAfterToPB interface {
	AfterToPB(context.Context, *BlockMissing) error
}

// DefaultCreateBlockMissing executes a basic gorm create call
func DefaultCreateBlockMissing(ctx context.Context, in *BlockMissing, db *gorm1.DB) (*BlockMissing, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockMissingORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockMissingORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BlockMissingORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockMissingORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskBlockMissing patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBlockMissing(ctx context.Context, patchee *BlockMissing, patcher *BlockMissing, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*BlockMissing, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Number" {
			patchee.Number = patcher.Number
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBlockMissing executes a gorm list call
func DefaultListBlockMissing(ctx context.Context, db *gorm1.DB) ([]*BlockMissing, error) {
	in := BlockMissing{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockMissingORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &BlockMissingORM{}, &BlockMissing{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockMissingORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("number")
	ormResponse := []BlockMissingORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockMissingORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*BlockMissing{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BlockMissingORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockMissingORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockMissingORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]BlockMissingORM) error
}