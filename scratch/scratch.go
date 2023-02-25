// Copyright 2023 Charge Net Stations and Contributors.
// SPDX-License-Identifier: Apache-2.0
package main

/* 

-- Step One --
Determoine user stories and create them using the following format
*** As a Persona, I want to Action, so that I can Result ***

1. List all guitars by make that are available in the store.
	As a Musician, I want to see all guitars available in the store so that I can see which are avilable.
2. I want to be able to filter by model ie. When i give Fender, I get back Strat, Squire, etc.
	As a musician, I want to be able to filter results by Make so thst I can see guitars that I am interested in.
3. I want to know how many guitars I have in stock
	As a store owner, I want to be able to see my guitar inventory so that I can know what is in stock and shippable.
4. I want to filter by acoustic guitars vs. electric
	As a musician, I want to filter guitars by acoustic and electric so that I can look at guitar types I'm interested in.
5. I want to be able to add new guitar inventory
	As a store owner I want to be able to add new guitars as I need so that I can place them in the store for sale.
6. I want pricing information.
	As a musician I want to be able to see pricing information so that I can tell if I want to purchase a guitar.
7. I want to filter by color, wood type, and other physical properties.
	As a musicain, I want to be able to filter guitars by attributes such as color, wood type, etc. so that I can browse guitars I'm interested in.
8. I want to be able to CRUD guitars (Create, Read, Upadate, and Delete)
	As a store owner I want to be able to create new guitar types so that I can add new guitars to my store inventory.
	As a store owner I want to be able to list guitars so that I can determine what guitars I have in stock.
	As a store owner I want to be able to update exsiting guitar types i case there is a misttake or change in the offering so that my store accuately reflects the guitars I have in stock.
9. Inventory reporting - I want to know when a guitar that is not in stock will be available.
	As a store owner I want to be able to list guitars that I offer but are not in stock so that I can stock my store with needed inventory.
	As a musician I want to see guitars that are in stock vs. out of stock so that I can determine if I want to wait for a guitar to be in stock.
10. Guitar types (Bass, normal)
	See #7

-- Step Two --
Get started and create the first test without writing any code.  First we need to just be able to list a guitat or two.

func TestSearch(t *testing.T) {

	t.Run("returns all available guitar makes when called with no filters", func(t *testing.T) {

		got := Search()
		want := "Fender"

		if want != got {
			t.Errorf("/ngot %s /nwant %s", got, want)
		}
	})
}

-- Step 2.1 --
Write the least amount of code to pass the test

func Search() []string {
	return []string{"Fender", "Ibenez"}
}

-- Step 3 --
guitars := []string{"Fender", "Ibenez"}

	t.Run("returns all available guitar makes when called with no filters", func(t *testing.T) {

		got := Search()
		want := guitars

		if !reflect.DeepEqual(want, got) {
			t.Errorf("/ngot %s /nwant %s", got, want)
		}
	})

	-- Step Three --

	















	*/