// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/enthistory"

	"github.com/datumforge/entx"

	"time"
)

// GroupSettingHistory holds the schema definition for the GroupSettingHistory entity.
type GroupSettingHistory struct {
	ent.Schema
}

// Annotations of the GroupSettingHistory.
func (GroupSettingHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "group_setting_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

// Fields of the GroupSettingHistory.
func (GroupSettingHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.String("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
	}

	// get the fields from the mixins
	// we only want to include mixin fields, not edges
	// so this prevents FKs back to the main tables
	mixins := GroupSetting{}.Mixin()
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			historyFields = append(historyFields, field)
		}
	}

	original := GroupSetting{}
	for _, field := range original.Fields() {
		historyFields = append(historyFields, field)
	}

	return historyFields
}
func (GroupSettingHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}
