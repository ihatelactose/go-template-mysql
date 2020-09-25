// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testFollowers(t *testing.T) {
	t.Parallel()

	query := Followers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFollowersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Followers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFollowersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FollowerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Follower exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FollowerExists to return true, but got false.")
	}
}

func testFollowersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	followerFound, err := FindFollower(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if followerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFollowersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Followers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFollowersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Followers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFollowersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	followerOne := &Follower{}
	followerTwo := &Follower{}
	if err = randomize.Struct(seed, followerOne, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}
	if err = randomize.Struct(seed, followerTwo, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Followers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFollowersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	followerOne := &Follower{}
	followerTwo := &Follower{}
	if err = randomize.Struct(seed, followerOne, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}
	if err = randomize.Struct(seed, followerTwo, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = followerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = followerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testFollowersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(followerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFollowerToOneUserUsingFollowee(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Follower
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FolloweeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Followee().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FollowerSlice{&local}
	if err = local.L.LoadFollowee(ctx, tx, false, (*[]*Follower)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Followee == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Followee = nil
	if err = local.L.LoadFollowee(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Followee == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFollowerToOneUserUsingFollower(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Follower
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, followerDBTypes, false, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FollowerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Follower().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := FollowerSlice{&local}
	if err = local.L.LoadFollower(ctx, tx, false, (*[]*Follower)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Follower == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Follower = nil
	if err = local.L.LoadFollower(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Follower == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testFollowerToOneSetOpUserUsingFollowee(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Follower
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, followerDBTypes, false, strmangle.SetComplement(followerPrimaryKeyColumns, followerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetFollowee(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Followee != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FolloweeFollowers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FolloweeID != x.ID {
			t.Error("foreign key was wrong value", a.FolloweeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FolloweeID))
		reflect.Indirect(reflect.ValueOf(&a.FolloweeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FolloweeID != x.ID {
			t.Error("foreign key was wrong value", a.FolloweeID, x.ID)
		}
	}
}
func testFollowerToOneSetOpUserUsingFollower(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Follower
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, followerDBTypes, false, strmangle.SetComplement(followerPrimaryKeyColumns, followerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetFollower(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Follower != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FollowerFollowers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FollowerID != x.ID {
			t.Error("foreign key was wrong value", a.FollowerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FollowerID))
		reflect.Indirect(reflect.ValueOf(&a.FollowerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FollowerID != x.ID {
			t.Error("foreign key was wrong value", a.FollowerID, x.ID)
		}
	}
}

func testFollowersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFollowersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FollowerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFollowersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Followers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	followerDBTypes = map[string]string{`ID`: `integer`, `FollowerID`: `integer`, `FolloweeID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testFollowersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(followerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(followerAllColumns) == len(followerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followerDBTypes, true, followerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFollowersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(followerAllColumns) == len(followerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Follower{}
	if err = randomize.Struct(seed, o, followerDBTypes, true, followerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, followerDBTypes, true, followerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(followerAllColumns, followerPrimaryKeyColumns) {
		fields = followerAllColumns
	} else {
		fields = strmangle.SetComplement(
			followerAllColumns,
			followerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := FollowerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFollowersUpsert(t *testing.T) {
	t.Parallel()

	if len(followerAllColumns) == len(followerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Follower{}
	if err = randomize.Struct(seed, &o, followerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Follower: %s", err)
	}

	count, err := Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, followerDBTypes, false, followerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Follower struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Follower: %s", err)
	}

	count, err = Followers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}