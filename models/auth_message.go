// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// AuthMessage is an object representing the database table.
type AuthMessage struct {
	ID      int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID  int    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Message string `boil:"message" json:"message" toml:"message" yaml:"message"`

	R *authMessageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authMessageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authMessageR is where relationships are stored.
type authMessageR struct {
	User *AuthUser
}

// authMessageL is where Load methods for each relationship are stored.
type authMessageL struct{}

var (
	authMessageColumns               = []string{"id", "user_id", "message"}
	authMessageColumnsWithoutDefault = []string{"user_id", "message"}
	authMessageColumnsWithDefault    = []string{"id"}
	authMessagePrimaryKeyColumns     = []string{"id"}
)

type (
	// AuthMessageSlice is an alias for a slice of pointers to AuthMessage.
	// This should generally be used opposed to []AuthMessage.
	AuthMessageSlice []*AuthMessage

	authMessageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authMessageType                 = reflect.TypeOf(&AuthMessage{})
	authMessageMapping              = queries.MakeStructMapping(authMessageType)
	authMessagePrimaryKeyMapping, _ = queries.BindMapping(authMessageType, authMessageMapping, authMessagePrimaryKeyColumns)
	authMessageInsertCacheMut       sync.RWMutex
	authMessageInsertCache          = make(map[string]insertCache)
	authMessageUpdateCacheMut       sync.RWMutex
	authMessageUpdateCache          = make(map[string]updateCache)
	authMessageUpsertCacheMut       sync.RWMutex
	authMessageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single authMessage record from the query, and panics on error.
func (q authMessageQuery) OneP() *AuthMessage {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authMessage record from the query.
func (q authMessageQuery) One() (*AuthMessage, error) {
	o := &AuthMessage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_message")
	}

	return o, nil
}

// AllP returns all AuthMessage records from the query, and panics on error.
func (q authMessageQuery) AllP() AuthMessageSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthMessage records from the query.
func (q authMessageQuery) All() (AuthMessageSlice, error) {
	var o AuthMessageSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthMessage slice")
	}

	return o, nil
}

// CountP returns the count of all AuthMessage records in the query, and panics on error.
func (q authMessageQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthMessage records in the query.
func (q authMessageQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_message rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authMessageQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authMessageQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_message exists")
	}

	return count > 0, nil
}

// UserG pointed to by the foreign key.
func (o *AuthMessage) UserG(mods ...qm.QueryMod) authUserQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *AuthMessage) User(exec boil.Executor, mods ...qm.QueryMod) authUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "`auth_user`")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authMessageL) LoadUser(e boil.Executor, singular bool, maybeAuthMessage interface{}) error {
	var slice []*AuthMessage
	var object *AuthMessage

	count := 1
	if singular {
		object = maybeAuthMessage.(*AuthMessage)
	} else {
		slice = *maybeAuthMessage.(*AuthMessageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &authMessageR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &authMessageR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from `auth_user` where `id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthUser")
	}
	defer results.Close()

	var resultSlice []*AuthUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthUser")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// SetUserG of the auth_message to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAuthMessages.
// Uses the global database handle.
func (o *AuthMessage) SetUserG(insert bool, related *AuthUser) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUserP of the auth_message to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAuthMessages.
// Panics on error.
func (o *AuthMessage) SetUserP(exec boil.Executor, insert bool, related *AuthUser) {
	if err := o.SetUser(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUserGP of the auth_message to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAuthMessages.
// Uses the global database handle and panics on error.
func (o *AuthMessage) SetUserGP(insert bool, related *AuthUser) {
	if err := o.SetUser(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetUser of the auth_message to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAuthMessages.
func (o *AuthMessage) SetUser(exec boil.Executor, insert bool, related *AuthUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `auth_message` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, authMessagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID

	if o.R == nil {
		o.R = &authMessageR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &authUserR{
			UserAuthMessages: AuthMessageSlice{o},
		}
	} else {
		related.R.UserAuthMessages = append(related.R.UserAuthMessages, o)
	}

	return nil
}

// AuthMessagesG retrieves all records.
func AuthMessagesG(mods ...qm.QueryMod) authMessageQuery {
	return AuthMessages(boil.GetDB(), mods...)
}

// AuthMessages retrieves all the records using an executor.
func AuthMessages(exec boil.Executor, mods ...qm.QueryMod) authMessageQuery {
	mods = append(mods, qm.From("`auth_message`"))
	return authMessageQuery{NewQuery(exec, mods...)}
}

// FindAuthMessageG retrieves a single record by ID.
func FindAuthMessageG(id int, selectCols ...string) (*AuthMessage, error) {
	return FindAuthMessage(boil.GetDB(), id, selectCols...)
}

// FindAuthMessageGP retrieves a single record by ID, and panics on error.
func FindAuthMessageGP(id int, selectCols ...string) *AuthMessage {
	retobj, err := FindAuthMessage(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthMessage(exec boil.Executor, id int, selectCols ...string) (*AuthMessage, error) {
	authMessageObj := &AuthMessage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `auth_message` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(authMessageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_message")
	}

	return authMessageObj, nil
}

// FindAuthMessageP retrieves a single record by ID with an executor, and panics on error.
func FindAuthMessageP(exec boil.Executor, id int, selectCols ...string) *AuthMessage {
	retobj, err := FindAuthMessage(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthMessage) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthMessage) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthMessage) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthMessage) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_message provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(authMessageColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authMessageInsertCacheMut.RLock()
	cache, cached := authMessageInsertCache[key]
	authMessageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authMessageColumns,
			authMessageColumnsWithDefault,
			authMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authMessageType, authMessageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authMessageType, authMessageMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `auth_message` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `auth_message` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, authMessagePrimaryKeyColumns))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into auth_message")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authMessageMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for auth_message")
	}

CacheNoHooks:
	if !cached {
		authMessageInsertCacheMut.Lock()
		authMessageInsertCache[key] = cache
		authMessageInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AuthMessage record. See Update for
// whitelist behavior description.
func (o *AuthMessage) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthMessage record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthMessage) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthMessage, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthMessage) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthMessage.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthMessage) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	authMessageUpdateCacheMut.RLock()
	cache, cached := authMessageUpdateCache[key]
	authMessageUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authMessageColumns, authMessagePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_message, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `auth_message` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, authMessagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authMessageType, authMessageMapping, append(wl, authMessagePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update auth_message row")
	}

	if !cached {
		authMessageUpdateCacheMut.Lock()
		authMessageUpdateCache[key] = cache
		authMessageUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authMessageQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authMessageQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_message")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthMessageSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthMessageSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthMessageSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthMessageSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `auth_message` SET %s WHERE (`id`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authMessagePrimaryKeyColumns), len(colNames)+1, len(authMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authMessage slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthMessage) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthMessage) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthMessage) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthMessage) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_message provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(authMessageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	authMessageUpsertCacheMut.RLock()
	cache, cached := authMessageUpsertCache[key]
	authMessageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authMessageColumns,
			authMessageColumnsWithDefault,
			authMessageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authMessageColumns,
			authMessagePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_message, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "auth_message", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `auth_message` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(authMessageType, authMessageMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authMessageType, authMessageMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for auth_message")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authMessageMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for auth_message")
	}

CacheNoHooks:
	if !cached {
		authMessageUpsertCacheMut.Lock()
		authMessageUpsertCache[key] = cache
		authMessageUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single AuthMessage record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthMessage) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthMessage record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthMessage) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthMessage provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthMessage record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthMessage) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthMessage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthMessage) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthMessage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authMessagePrimaryKeyMapping)
	sql := "DELETE FROM `auth_message` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_message")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authMessageQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authMessageQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authMessageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_message")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthMessageSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthMessageSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthMessage slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthMessageSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthMessageSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthMessage slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `auth_message` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authMessagePrimaryKeyColumns), 1, len(authMessagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authMessage slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthMessage) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthMessage) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthMessage) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthMessage provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthMessage) Reload(exec boil.Executor) error {
	ret, err := FindAuthMessage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthMessageSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthMessageSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthMessageSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthMessageSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthMessageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authMessages := AuthMessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `auth_message`.* FROM `auth_message` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authMessagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authMessagePrimaryKeyColumns), 1, len(authMessagePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authMessages)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthMessageSlice")
	}

	*o = authMessages

	return nil
}

// AuthMessageExists checks if the AuthMessage row exists.
func AuthMessageExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `auth_message` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_message exists")
	}

	return exists, nil
}

// AuthMessageExistsG checks if the AuthMessage row exists.
func AuthMessageExistsG(id int) (bool, error) {
	return AuthMessageExists(boil.GetDB(), id)
}

// AuthMessageExistsGP checks if the AuthMessage row exists. Panics on error.
func AuthMessageExistsGP(id int) bool {
	e, err := AuthMessageExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthMessageExistsP checks if the AuthMessage row exists. Panics on error.
func AuthMessageExistsP(exec boil.Executor, id int) bool {
	e, err := AuthMessageExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}