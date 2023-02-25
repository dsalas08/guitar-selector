/**
 * Copyright 2023 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import  (
	"testing"
	"reflect"
)

func TestSearch(t *testing.T) {
	t.Run("returns all available guitar makes when called with no filters", func(t *testing.T) {

		guitars := []Guitar{
			{
				Make: "Fender",
			},
			{
				Make: "Ibenez",
			},
		}

		got := Search(guitars)
		want := guitars

		if !reflect.DeepEqual(want, got) {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})
}