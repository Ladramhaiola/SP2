package main

import "fmt"

func testTable(t *Table) {
	// fmt.Printf("Starting table:\n\n")
	// t.print()

	// fmt.Printf("\nTesting direct functions:\n\n")
	// fmt.Printf("Select 4: %s", t.selectDirect(4).tostring())
	fmt.Printf("Insert 1: %s", t.insertDirect(newRecord("Keldman", 0, 89.0), 1).tostring())
	// fmt.Printf("Update 7: %s", t.updateDirect(newRecord("Aksyonenko", 1, 3.0), 7).tostring())
	// fmt.Printf("Delete 8\n\n")
	// t.deleteDirect(9)

	// fmt.Printf("Resulting Table:\n\n")
	// t.print()

	// fmt.Printf("\nTesting linear functions:\n\n")
	// fmt.Printf("Select Key{Burbil, 9}: %s", t.selectLinear(Key{"Burbil", 9}).tostring())
	fmt.Printf("Insert Key{Karagan, 90}: %s", t.insertLinear(newRecord("Karagan", 90, 13.3)).tostring())
	// fmt.Printf("Update Key{Kurbil, 9}: %s", t.updateLinear(Key{"Burbil", 9}, newRecord("Vel", 13, 0.3)).tostring())
	// fmt.Printf("Delete Key{Daragan, 90}\n\n")
	// t.deleteLinear(Key{"Daragan", 90})

	// fmt.Printf("Resulting Table:\n\n")
	// t.print()

	// fmt.Printf("\nTesting Binary functions:\n\n")
	// fmt.Printf("Select Key{Feldman, 0}: %s", t.selectBinary(Key{"Feldman", 0}).tostring())
	// fmt.Printf("Update Key{Feldman, 0}: %s", t.updateBinary(newRecord("Feldman", 15, 3), Key{"Feldman", 0}).tostring())
	fmt.Printf("Insert Key{Vel, 134}:   %s", t.insertBinary(newRecord("Vel", 134, 98)).tostring())
	// fmt.Printf("Delete Key{Mirchuk, 7}\n\n")
	// t.deleteBinary(Key{"Mirchuk", 7})

	// fmt.Printf("Resulting Table:\n\n")
	// t.print()

	// t.insertDirect(newRecord("Kza", 7, 13), 4)
	t.print()
	fmt.Printf("\nTesting closest search:\n")
	fmt.Printf("Search 'Kay':\n")
	res, err := t.searchClosest("K")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, r := range res {
			fmt.Println(r.tostring())
		}
	}
}
