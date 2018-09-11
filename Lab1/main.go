package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

// Key defines struct to contain information about table key
type Key struct {
	str string
	mod uint16
}

// Func represents functional field of Record
type Func struct {
	f float64
}

// Record represents data struct in table
type Record struct {
	key Key
	fnc Func
}

// Table -
type Table struct {
	records []*Record
}

// check if record is empty
func (r *Record) isEmpty() bool {
	if r.key.str == "" {
		return true
	}
	return false
}

// clear record
func (r *Record) clear() {
	r.key.str = ""
	r.key.mod = 0
	r.fnc.f = 0
}

// copy data from one record to other
func (r *Record) copy(other *Record) {
	r.key = other.key
	r.fnc = other.fnc
}

func (r *Record) tostring() string {
	if r.isEmpty() {
		return fmt.Sprintln("Record ( empty )")
	}
	return fmt.Sprintf("Record ( key: %s, mod: %d, functional: %0.2f)\n", r.key.str, r.key.mod, r.fnc.f)
}

func (r *Record) print() {
	fmt.Println(r.tostring())
}

func newRecord(keyStr string, keyMod uint16, funcF float64) *Record {
	return &Record{key: Key{keyStr, keyMod}, fnc: Func{funcF}}
}

func emptyRecord() *Record {
	return &Record{}
}

// get record by index
func (tb *Table) selectDirect(n int) *Record {
	if len(tb.records) <= n {
		log.Println("No element at such position")
		return emptyRecord()
	}
	return tb.records[n]
}

// insert record by index
func (tb *Table) insertDirect(rec *Record, n int) *Record {
	count := n - len(tb.records)
	if count > 0 {
		for count >= 0 {
			tb.records = append(tb.records, emptyRecord())
			count--
		}
	}
	tb.records[n] = rec
	return tb.records[n]
}

// delete record by index
func (tb *Table) deleteDirect(n int) {
	if n < len(tb.records) {
		tb.records[n].clear()
		return
	}
	log.Println("deleteDirect: index out of range")
}

// update record at given index
func (tb *Table) updateDirect(rec *Record, n int) *Record {
	if n < len(tb.records) {
		tb.records[n].copy(rec)
		return tb.records[n]
	}
	log.Println("updateDirect: index out of range")
	return emptyRecord()
}

// linear search implementation for table
func (tb *Table) selectLinear(key Key) *Record {
	for _, rec := range tb.records {
		if cmpKeys(key, rec.key) == 0 {
			return rec
		}
	}
	return emptyRecord()
}

func (tb *Table) insertLinear(rec *Record) *Record {
	for i, record := range tb.records {
		// if able to find emty space
		if record.key.str == "" {
			tb.records[i] = rec
			return tb.records[i]
		}
	}
	// else just append
	tb.records = append(tb.records, rec)
	return tb.records[len(tb.records)-1]
}

func (tb *Table) deleteLinear(key Key) {
	res := tb.selectLinear(key)
	if !res.isEmpty() {
		res.clear()
	}
}

// update record which already exists in table
func (tb *Table) updateLinear(key Key, rec *Record) *Record {
	res := tb.selectLinear(key)
	if !res.isEmpty() {
		res.copy(rec)
		return res
	}
	log.Println("updateLinear: record not found")
	return emptyRecord()
}

// delete all empty records from table
func (tb *Table) pack() {
	for i := len(tb.records) - 1; i > 0; i-- {
		if tb.records[i].isEmpty() {
			tb.records = append(tb.records[:i], tb.records[i+1:]...)
		}
	}
}

// binary search implementation for table
func (tb *Table) selectBinary(key Key) *Record {
	tb.sortByStr(key.str, 0)
	l, r := 0, len(tb.records)
	for l != r {
		mid := (l + r) / 2
		cmp := cmpKeys(tb.records[mid].key, key)
		if cmp == 0 {
			return tb.records[mid]
		} else if cmp < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return emptyRecord()
}

func (tb *Table) insertBinary(rec *Record) *Record {
	res := tb.selectBinary(rec.key)
	if res.isEmpty() {
		res.copy(rec)
		tb.records = append(tb.records, res)
	}
	return res
}

func (tb *Table) deleteBinary(key Key) {
	res := tb.selectBinary(key)
	if !res.isEmpty() {
		res.clear()
		return
	}
	fmt.Println("No records found")
}

func (tb *Table) updateBinary(rec *Record, key Key) *Record {
	res := tb.selectBinary(key)
	if !res.isEmpty() {
		res.copy(rec)
	}
	return res
}

func cmpKeys(k1 Key, k2 Key) uint16 {
	if cmp := strings.Compare(k1.str, k2.str); cmp != 0 {
		return uint16(cmp)
	}
	return k1.mod - k2.mod
}

func (tb *Table) sortByStr(str string, limit int) []*Record {
	By(func(r1, r2 *Record) bool {
		return cmpStr(str, r1.key.str) > cmpStr(str, r2.key.str)
	}).Sort(tb)
	return tb.records[:limit]
}

func (tb *Table) searchClosest(str string) ([]*Record, error) {
	result := []*Record{}
	for _, rec := range tb.records {
		if cmpStr(str, rec.key.str) > 0 {
			result = append(result, rec)
		}
	}
	if len(result) < 1 {
		return nil, errors.New("nothing found")
	}
	t := &Table{result}
	By(func(r1, r2 *Record) bool {
		return cmpStr(str, r1.key.str) > cmpStr(str, r2.key.str)
	}).Sort(t)
	return t.records, nil
}

func (tb *Table) print() {
	fmt.Println("Table:")
	for i, elem := range tb.records {
		fmt.Printf("\t%d %s", i, elem.tostring())
	}
}

func main() {
	t := &Table{records: []*Record{
		newRecord("Shapka", 8, 13.3),
		emptyRecord(),
		emptyRecord(),
		newRecord("Kuzmenko", 17, 3.38),
		newRecord("Mirchuk", 7, 14.9),
		emptyRecord(),
		newRecord("Kuzmenko", 14, 78.3),
		emptyRecord(),
		newRecord("Burbil", 9, 12.301),
		newRecord("Vel", 2, 0.56),
	}}
	testTable(t)
}

// By is a type of a "less" function that defines the ordering of its Record arguments
type By func(r1, r2 *Record) bool

type recordsSorter struct {
	records []*Record
	by      func(r1, r2 *Record) bool
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function
func (by By) Sort(t *Table) {
	rs := &recordsSorter{
		records: t.records,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order
	}
	sort.Sort(rs)
}

// Len is part of sort.Interface, so must be implemented
func (s *recordsSorter) Len() int {
	return len(s.records)
}

// Swap is also a part of sort interface
func (s *recordsSorter) Swap(i, j int) {
	s.records[i], s.records[j] = s.records[j], s.records[i]
}

// Less ...
func (s *recordsSorter) Less(i, j int) bool {
	return s.by(s.records[i], s.records[j])
}

// function to compare strings by rule from 1.3 table
func cmpStr(a, b string) int {
	if b == "" {
		return 0
	}
	n, limit := 0, int(math.Min(float64(len(a)), float64(len(b))))
	for i := 0; i < limit; i++ {
		if b[i] == a[i] {
			n++
		} else {
			break
		}
	}
	return n
}
