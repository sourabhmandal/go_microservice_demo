// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gomicro/ent/schema"
	"gomicro/ent/task"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescName is the schema descriptor for name field.
	taskDescName := taskFields[1].Descriptor()
	// task.NameValidator is a validator for the "name" field. It is called by the builders before save.
	task.NameValidator = taskDescName.Validators[0].(func(string) error)
	// taskDescIsDone is the schema descriptor for is_done field.
	taskDescIsDone := taskFields[2].Descriptor()
	// task.DefaultIsDone holds the default value on creation for the is_done field.
	task.DefaultIsDone = taskDescIsDone.Default.(bool)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[3].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[4].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
}
