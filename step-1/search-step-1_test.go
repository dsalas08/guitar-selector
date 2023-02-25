/**
 * Copyright 2023 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import  (
	"testing"

)

func TestSearch(t *testing.T) {

	t.Run("returns all available guitar makes when called with no filters", func(t *testing.T) {

		got := Search()
		want := "Fender"

		if want != got {
			t.Errorf("/ngot %s /nwant %s", got, want)
		}
	})

}