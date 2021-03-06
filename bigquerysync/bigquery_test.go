// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0

package bigquerysync_test

import (
	"testing"

	"google.golang.org/appengine/datastore"

	"ronoaldo.gopkg.net/aetools"
	"ronoaldo.gopkg.net/aetools/bigquerysync"
)

func TestIngestToBigQuery(t *testing.T) {
	c := SetupEnv(t)
	defer c.Close()

	e := &aetools.Entity{
		Key: datastore.NewKey(c, "Sample", "", 1, nil),
		Properties: []datastore.Property{
			datastore.Property{Name: "Name", Value: "Test value"},
		},
	}

	err := bigquerysync.IngestToBigQuery(c, "project", "dataset", []*aetools.Entity{e}, "")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
