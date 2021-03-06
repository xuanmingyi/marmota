// Code generated by entc, DO NOT EDIT.

package node

import (
	"marmota/server/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Metadata applies equality check predicate on the "metadata" field. It's identical to MetadataEQ.
func Metadata(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// MetadataEQ applies the EQ predicate on the "metadata" field.
func MetadataEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetadata), v))
	})
}

// MetadataNEQ applies the NEQ predicate on the "metadata" field.
func MetadataNEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMetadata), v))
	})
}

// MetadataIn applies the In predicate on the "metadata" field.
func MetadataIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMetadata), v...))
	})
}

// MetadataNotIn applies the NotIn predicate on the "metadata" field.
func MetadataNotIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMetadata), v...))
	})
}

// MetadataGT applies the GT predicate on the "metadata" field.
func MetadataGT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMetadata), v))
	})
}

// MetadataGTE applies the GTE predicate on the "metadata" field.
func MetadataGTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMetadata), v))
	})
}

// MetadataLT applies the LT predicate on the "metadata" field.
func MetadataLT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMetadata), v))
	})
}

// MetadataLTE applies the LTE predicate on the "metadata" field.
func MetadataLTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMetadata), v))
	})
}

// MetadataContains applies the Contains predicate on the "metadata" field.
func MetadataContains(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMetadata), v))
	})
}

// MetadataHasPrefix applies the HasPrefix predicate on the "metadata" field.
func MetadataHasPrefix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMetadata), v))
	})
}

// MetadataHasSuffix applies the HasSuffix predicate on the "metadata" field.
func MetadataHasSuffix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMetadata), v))
	})
}

// MetadataEqualFold applies the EqualFold predicate on the "metadata" field.
func MetadataEqualFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMetadata), v))
	})
}

// MetadataContainsFold applies the ContainsFold predicate on the "metadata" field.
func MetadataContainsFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMetadata), v))
	})
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	})
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDesc), v))
	})
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDesc), v...))
	})
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDesc), v...))
	})
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDesc), v))
	})
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDesc), v))
	})
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDesc), v))
	})
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDesc), v))
	})
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDesc), v))
	})
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDesc), v))
	})
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDesc), v))
	})
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDesc), v))
	})
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDesc), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		p(s.Not())
	})
}
