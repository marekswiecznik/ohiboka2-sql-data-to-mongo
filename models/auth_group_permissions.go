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

// AuthGroupPermission is an object representing the database table.
type AuthGroupPermission struct {
	ID           int `boil:"id" json:"id" toml:"id" yaml:"id"`
	GroupID      int `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	PermissionID int `boil:"permission_id" json:"permission_id" toml:"permission_id" yaml:"permission_id"`

	R *authGroupPermissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authGroupPermissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authGroupPermissionR is where relationships are stored.
type authGroupPermissionR struct {
	Group      *AuthGroup
	Permission *AuthPermission
}

// authGroupPermissionL is where Load methods for each relationship are stored.
type authGroupPermissionL struct{}

var (
	authGroupPermissionColumns               = []string{"id", "group_id", "permission_id"}
	authGroupPermissionColumnsWithoutDefault = []string{"group_id", "permission_id"}
	authGroupPermissionColumnsWithDefault    = []string{"id"}
	authGroupPermissionPrimaryKeyColumns     = []string{"id"}
)

type (
	// AuthGroupPermissionSlice is an alias for a slice of pointers to AuthGroupPermission.
	// This should generally be used opposed to []AuthGroupPermission.
	AuthGroupPermissionSlice []*AuthGroupPermission

	authGroupPermissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authGroupPermissionType                 = reflect.TypeOf(&AuthGroupPermission{})
	authGroupPermissionMapping              = queries.MakeStructMapping(authGroupPermissionType)
	authGroupPermissionPrimaryKeyMapping, _ = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, authGroupPermissionPrimaryKeyColumns)
	authGroupPermissionInsertCacheMut       sync.RWMutex
	authGroupPermissionInsertCache          = make(map[string]insertCache)
	authGroupPermissionUpdateCacheMut       sync.RWMutex
	authGroupPermissionUpdateCache          = make(map[string]updateCache)
	authGroupPermissionUpsertCacheMut       sync.RWMutex
	authGroupPermissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single authGroupPermission record from the query, and panics on error.
func (q authGroupPermissionQuery) OneP() *AuthGroupPermission {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authGroupPermission record from the query.
func (q authGroupPermissionQuery) One() (*AuthGroupPermission, error) {
	o := &AuthGroupPermission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_group_permissions")
	}

	return o, nil
}

// AllP returns all AuthGroupPermission records from the query, and panics on error.
func (q authGroupPermissionQuery) AllP() AuthGroupPermissionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthGroupPermission records from the query.
func (q authGroupPermissionQuery) All() (AuthGroupPermissionSlice, error) {
	var o AuthGroupPermissionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthGroupPermission slice")
	}

	return o, nil
}

// CountP returns the count of all AuthGroupPermission records in the query, and panics on error.
func (q authGroupPermissionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthGroupPermission records in the query.
func (q authGroupPermissionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_group_permissions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authGroupPermissionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authGroupPermissionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_group_permissions exists")
	}

	return count > 0, nil
}

// GroupG pointed to by the foreign key.
func (o *AuthGroupPermission) GroupG(mods ...qm.QueryMod) authGroupQuery {
	return o.Group(boil.GetDB(), mods...)
}

// Group pointed to by the foreign key.
func (o *AuthGroupPermission) Group(exec boil.Executor, mods ...qm.QueryMod) authGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthGroups(exec, queryMods...)
	queries.SetFrom(query.Query, "`auth_group`")

	return query
}

// PermissionG pointed to by the foreign key.
func (o *AuthGroupPermission) PermissionG(mods ...qm.QueryMod) authPermissionQuery {
	return o.Permission(boil.GetDB(), mods...)
}

// Permission pointed to by the foreign key.
func (o *AuthGroupPermission) Permission(exec boil.Executor, mods ...qm.QueryMod) authPermissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.PermissionID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthPermissions(exec, queryMods...)
	queries.SetFrom(query.Query, "`auth_permission`")

	return query
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authGroupPermissionL) LoadGroup(e boil.Executor, singular bool, maybeAuthGroupPermission interface{}) error {
	var slice []*AuthGroupPermission
	var object *AuthGroupPermission

	count := 1
	if singular {
		object = maybeAuthGroupPermission.(*AuthGroupPermission)
	} else {
		slice = *maybeAuthGroupPermission.(*AuthGroupPermissionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &authGroupPermissionR{}
		}
		args[0] = object.GroupID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &authGroupPermissionR{}
			}
			args[i] = obj.GroupID
		}
	}

	query := fmt.Sprintf(
		"select * from `auth_group` where `id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthGroup")
	}
	defer results.Close()

	var resultSlice []*AuthGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthGroup")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Group = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				break
			}
		}
	}

	return nil
}

// LoadPermission allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authGroupPermissionL) LoadPermission(e boil.Executor, singular bool, maybeAuthGroupPermission interface{}) error {
	var slice []*AuthGroupPermission
	var object *AuthGroupPermission

	count := 1
	if singular {
		object = maybeAuthGroupPermission.(*AuthGroupPermission)
	} else {
		slice = *maybeAuthGroupPermission.(*AuthGroupPermissionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &authGroupPermissionR{}
		}
		args[0] = object.PermissionID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &authGroupPermissionR{}
			}
			args[i] = obj.PermissionID
		}
	}

	query := fmt.Sprintf(
		"select * from `auth_permission` where `id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthPermission")
	}
	defer results.Close()

	var resultSlice []*AuthPermission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthPermission")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Permission = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PermissionID == foreign.ID {
				local.R.Permission = foreign
				break
			}
		}
	}

	return nil
}

// SetGroupG of the auth_group_permission to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthGroupPermissions.
// Uses the global database handle.
func (o *AuthGroupPermission) SetGroupG(insert bool, related *AuthGroup) error {
	return o.SetGroup(boil.GetDB(), insert, related)
}

// SetGroupP of the auth_group_permission to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthGroupPermissions.
// Panics on error.
func (o *AuthGroupPermission) SetGroupP(exec boil.Executor, insert bool, related *AuthGroup) {
	if err := o.SetGroup(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGroupGP of the auth_group_permission to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthGroupPermissions.
// Uses the global database handle and panics on error.
func (o *AuthGroupPermission) SetGroupGP(insert bool, related *AuthGroup) {
	if err := o.SetGroup(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetGroup of the auth_group_permission to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthGroupPermissions.
func (o *AuthGroupPermission) SetGroup(exec boil.Executor, insert bool, related *AuthGroup) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `auth_group_permissions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
		strmangle.WhereClause("`", "`", 0, authGroupPermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID

	if o.R == nil {
		o.R = &authGroupPermissionR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &authGroupR{
			GroupAuthGroupPermissions: AuthGroupPermissionSlice{o},
		}
	} else {
		related.R.GroupAuthGroupPermissions = append(related.R.GroupAuthGroupPermissions, o)
	}

	return nil
}

// SetPermissionG of the auth_group_permission to the related item.
// Sets o.R.Permission to related.
// Adds o to related.R.PermissionAuthGroupPermissions.
// Uses the global database handle.
func (o *AuthGroupPermission) SetPermissionG(insert bool, related *AuthPermission) error {
	return o.SetPermission(boil.GetDB(), insert, related)
}

// SetPermissionP of the auth_group_permission to the related item.
// Sets o.R.Permission to related.
// Adds o to related.R.PermissionAuthGroupPermissions.
// Panics on error.
func (o *AuthGroupPermission) SetPermissionP(exec boil.Executor, insert bool, related *AuthPermission) {
	if err := o.SetPermission(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPermissionGP of the auth_group_permission to the related item.
// Sets o.R.Permission to related.
// Adds o to related.R.PermissionAuthGroupPermissions.
// Uses the global database handle and panics on error.
func (o *AuthGroupPermission) SetPermissionGP(insert bool, related *AuthPermission) {
	if err := o.SetPermission(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPermission of the auth_group_permission to the related item.
// Sets o.R.Permission to related.
// Adds o to related.R.PermissionAuthGroupPermissions.
func (o *AuthGroupPermission) SetPermission(exec boil.Executor, insert bool, related *AuthPermission) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `auth_group_permissions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"permission_id"}),
		strmangle.WhereClause("`", "`", 0, authGroupPermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PermissionID = related.ID

	if o.R == nil {
		o.R = &authGroupPermissionR{
			Permission: related,
		}
	} else {
		o.R.Permission = related
	}

	if related.R == nil {
		related.R = &authPermissionR{
			PermissionAuthGroupPermissions: AuthGroupPermissionSlice{o},
		}
	} else {
		related.R.PermissionAuthGroupPermissions = append(related.R.PermissionAuthGroupPermissions, o)
	}

	return nil
}

// AuthGroupPermissionsG retrieves all records.
func AuthGroupPermissionsG(mods ...qm.QueryMod) authGroupPermissionQuery {
	return AuthGroupPermissions(boil.GetDB(), mods...)
}

// AuthGroupPermissions retrieves all the records using an executor.
func AuthGroupPermissions(exec boil.Executor, mods ...qm.QueryMod) authGroupPermissionQuery {
	mods = append(mods, qm.From("`auth_group_permissions`"))
	return authGroupPermissionQuery{NewQuery(exec, mods...)}
}

// FindAuthGroupPermissionG retrieves a single record by ID.
func FindAuthGroupPermissionG(id int, selectCols ...string) (*AuthGroupPermission, error) {
	return FindAuthGroupPermission(boil.GetDB(), id, selectCols...)
}

// FindAuthGroupPermissionGP retrieves a single record by ID, and panics on error.
func FindAuthGroupPermissionGP(id int, selectCols ...string) *AuthGroupPermission {
	retobj, err := FindAuthGroupPermission(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthGroupPermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthGroupPermission(exec boil.Executor, id int, selectCols ...string) (*AuthGroupPermission, error) {
	authGroupPermissionObj := &AuthGroupPermission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `auth_group_permissions` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(authGroupPermissionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_group_permissions")
	}

	return authGroupPermissionObj, nil
}

// FindAuthGroupPermissionP retrieves a single record by ID with an executor, and panics on error.
func FindAuthGroupPermissionP(exec boil.Executor, id int, selectCols ...string) *AuthGroupPermission {
	retobj, err := FindAuthGroupPermission(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthGroupPermission) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthGroupPermission) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthGroupPermission) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthGroupPermission) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_group_permissions provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(authGroupPermissionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authGroupPermissionInsertCacheMut.RLock()
	cache, cached := authGroupPermissionInsertCache[key]
	authGroupPermissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authGroupPermissionColumns,
			authGroupPermissionColumnsWithDefault,
			authGroupPermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `auth_group_permissions` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `auth_group_permissions` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, authGroupPermissionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into auth_group_permissions")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authGroupPermissionMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for auth_group_permissions")
	}

CacheNoHooks:
	if !cached {
		authGroupPermissionInsertCacheMut.Lock()
		authGroupPermissionInsertCache[key] = cache
		authGroupPermissionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AuthGroupPermission record. See Update for
// whitelist behavior description.
func (o *AuthGroupPermission) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthGroupPermission record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthGroupPermission) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthGroupPermission, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthGroupPermission) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthGroupPermission.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthGroupPermission) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	authGroupPermissionUpdateCacheMut.RLock()
	cache, cached := authGroupPermissionUpdateCache[key]
	authGroupPermissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authGroupPermissionColumns, authGroupPermissionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_group_permissions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `auth_group_permissions` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, authGroupPermissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, append(wl, authGroupPermissionPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update auth_group_permissions row")
	}

	if !cached {
		authGroupPermissionUpdateCacheMut.Lock()
		authGroupPermissionUpdateCache[key] = cache
		authGroupPermissionUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authGroupPermissionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authGroupPermissionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_group_permissions")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthGroupPermissionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthGroupPermissionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthGroupPermissionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthGroupPermissionSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `auth_group_permissions` SET %s WHERE (`id`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authGroupPermissionPrimaryKeyColumns), len(colNames)+1, len(authGroupPermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authGroupPermission slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthGroupPermission) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthGroupPermission) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthGroupPermission) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthGroupPermission) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_group_permissions provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(authGroupPermissionColumnsWithDefault, o)

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

	authGroupPermissionUpsertCacheMut.RLock()
	cache, cached := authGroupPermissionUpsertCache[key]
	authGroupPermissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authGroupPermissionColumns,
			authGroupPermissionColumnsWithDefault,
			authGroupPermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authGroupPermissionColumns,
			authGroupPermissionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_group_permissions, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "auth_group_permissions", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `auth_group_permissions` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_group_permissions")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == authGroupPermissionMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for auth_group_permissions")
	}

CacheNoHooks:
	if !cached {
		authGroupPermissionUpsertCacheMut.Lock()
		authGroupPermissionUpsertCache[key] = cache
		authGroupPermissionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single AuthGroupPermission record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthGroupPermission) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthGroupPermission record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthGroupPermission) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthGroupPermission provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthGroupPermission record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthGroupPermission) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthGroupPermission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthGroupPermission) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthGroupPermission provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authGroupPermissionPrimaryKeyMapping)
	sql := "DELETE FROM `auth_group_permissions` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_group_permissions")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authGroupPermissionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authGroupPermissionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authGroupPermissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_group_permissions")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthGroupPermissionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthGroupPermissionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthGroupPermission slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthGroupPermissionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthGroupPermissionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthGroupPermission slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `auth_group_permissions` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authGroupPermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authGroupPermissionPrimaryKeyColumns), 1, len(authGroupPermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authGroupPermission slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthGroupPermission) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthGroupPermission) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthGroupPermission) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthGroupPermission provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthGroupPermission) Reload(exec boil.Executor) error {
	ret, err := FindAuthGroupPermission(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthGroupPermissionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthGroupPermissionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthGroupPermissionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthGroupPermissionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthGroupPermissionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authGroupPermissions := AuthGroupPermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `auth_group_permissions`.* FROM `auth_group_permissions` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authGroupPermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authGroupPermissionPrimaryKeyColumns), 1, len(authGroupPermissionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authGroupPermissions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthGroupPermissionSlice")
	}

	*o = authGroupPermissions

	return nil
}

// AuthGroupPermissionExists checks if the AuthGroupPermission row exists.
func AuthGroupPermissionExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `auth_group_permissions` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_group_permissions exists")
	}

	return exists, nil
}

// AuthGroupPermissionExistsG checks if the AuthGroupPermission row exists.
func AuthGroupPermissionExistsG(id int) (bool, error) {
	return AuthGroupPermissionExists(boil.GetDB(), id)
}

// AuthGroupPermissionExistsGP checks if the AuthGroupPermission row exists. Panics on error.
func AuthGroupPermissionExistsGP(id int) bool {
	e, err := AuthGroupPermissionExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthGroupPermissionExistsP checks if the AuthGroupPermission row exists. Panics on error.
func AuthGroupPermissionExistsP(exec boil.Executor, id int) bool {
	e, err := AuthGroupPermissionExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}