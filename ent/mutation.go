// Code generated by ent, DO NOT EDIT.

package ent

import (
	"FiberNewBie/ent/account"
	"FiberNewBie/ent/predicate"
	"FiberNewBie/ent/user"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccount = "Account"
	TypeUser    = "User"
)

// AccountMutation represents an operation that mutates the Account nodes in the graph.
type AccountMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	password      *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Account, error)
	predicates    []predicate.Account
}

var _ ent.Mutation = (*AccountMutation)(nil)

// accountOption allows management of the mutation configuration using functional options.
type accountOption func(*AccountMutation)

// newAccountMutation creates new mutation for the Account entity.
func newAccountMutation(c config, op Op, opts ...accountOption) *AccountMutation {
	m := &AccountMutation{
		config:        c,
		op:            op,
		typ:           TypeAccount,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountID sets the ID field of the mutation.
func withAccountID(id int) accountOption {
	return func(m *AccountMutation) {
		var (
			err   error
			once  sync.Once
			value *Account
		)
		m.oldValue = func(ctx context.Context) (*Account, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Account.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccount sets the old Account of the mutation.
func withAccount(node *Account) accountOption {
	return func(m *AccountMutation) {
		m.oldValue = func(context.Context) (*Account, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccountMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Account.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *AccountMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *AccountMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *AccountMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the "password" field.
func (m *AccountMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *AccountMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Account entity.
// If the Account object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *AccountMutation) ResetPassword() {
	m.password = nil
}

// SetOwnerID sets the "owner" edge to the User entity by id.
func (m *AccountMutation) SetOwnerID(id int) {
	m.owner = &id
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *AccountMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *AccountMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerID returns the "owner" edge ID in the mutation.
func (m *AccountMutation) OwnerID() (id int, exists bool) {
	if m.owner != nil {
		return *m.owner, true
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *AccountMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *AccountMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the AccountMutation builder.
func (m *AccountMutation) Where(ps ...predicate.Account) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AccountMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AccountMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Account, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AccountMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AccountMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Account).
func (m *AccountMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.username != nil {
		fields = append(fields, account.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, account.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case account.FieldUsername:
		return m.Username()
	case account.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case account.FieldUsername:
		return m.OldUsername(ctx)
	case account.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown Account field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) SetField(name string, value ent.Value) error {
	switch name {
	case account.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case account.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Account numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Account nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountMutation) ResetField(name string) error {
	switch name {
	case account.FieldUsername:
		m.ResetUsername()
		return nil
	case account.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown Account field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, account.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case account.EdgeOwner:
		if id := m.owner; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, account.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountMutation) EdgeCleared(name string) bool {
	switch name {
	case account.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountMutation) ClearEdge(name string) error {
	switch name {
	case account.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Account unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountMutation) ResetEdge(name string) error {
	switch name {
	case account.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Account edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op             Op
	typ            string
	id             *int
	age            *int
	addage         *int
	name           *string
	clearedFields  map[string]struct{}
	account        *int
	clearedaccount bool
	done           bool
	oldValue       func(context.Context) (*User, error)
	predicates     []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetAccountID sets the "account" edge to the Account entity by id.
func (m *UserMutation) SetAccountID(id int) {
	m.account = &id
}

// ClearAccount clears the "account" edge to the Account entity.
func (m *UserMutation) ClearAccount() {
	m.clearedaccount = true
}

// AccountCleared reports if the "account" edge to the Account entity was cleared.
func (m *UserMutation) AccountCleared() bool {
	return m.clearedaccount
}

// AccountID returns the "account" edge ID in the mutation.
func (m *UserMutation) AccountID() (id int, exists bool) {
	if m.account != nil {
		return *m.account, true
	}
	return
}

// AccountIDs returns the "account" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AccountID instead. It exists only for internal usage by the builders.
func (m *UserMutation) AccountIDs() (ids []int) {
	if id := m.account; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAccount resets all changes to the "account" edge.
func (m *UserMutation) ResetAccount() {
	m.account = nil
	m.clearedaccount = false
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.account != nil {
		edges = append(edges, user.EdgeAccount)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeAccount:
		if id := m.account; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedaccount {
		edges = append(edges, user.EdgeAccount)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeAccount:
		return m.clearedaccount
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeAccount:
		m.ClearAccount()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeAccount:
		m.ResetAccount()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
