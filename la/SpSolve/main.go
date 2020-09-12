// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"gosl/io"
	"gosl/la"
)

func main() {

	// input matrix data into Triplet
	A := new(la.Triplet)
	A.Init(5, 5, 13)
	A.Put(0, 0, +1.0) // << note duplicated entry
	A.Put(0, 0, +1.0) // << duplicated
	A.Put(1, 0, +3.0)
	A.Put(0, 1, +3.0)
	A.Put(2, 1, -1.0)
	A.Put(4, 1, +4.0)
	A.Put(1, 2, +4.0)
	A.Put(2, 2, -3.0)
	A.Put(3, 2, +1.0)
	A.Put(4, 2, +2.0)
	A.Put(2, 3, +2.0)
	A.Put(1, 4, +6.0)
	A.Put(4, 4, +1.0)

	// solve
	b := []float64{8.0, 45.0, -3.0, 3.0, 19.0}
	x := la.SpSolve(A, b)

	// computed output
	io.Pf("x = %f\n", x)

	// show correct values
	io.Pf("xCorrect = %v\n", []float64{1, 2, 3, 4, 5})
}
