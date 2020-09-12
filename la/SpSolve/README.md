# SpSolve: Sparse-solve a linear system of equations

We solve a small linear system (find vector x such that A.x = b; given matrix A and vector b) where the system matrix A is sparse, i.e. matrix A has too many zeros (zero-valued components). Therefore, only the non-zero values in A need to be input and considered.

We represent the sparse matrix A using a structure called Triplet, where we store the indices and values of the non-zeros. The Triplet is a very convenient object for data input and even allows duplicated entries (automatically added).

In full form, matrix A looks like:

```
      _                      _
     |  2   3    0    0    0  |
     |  3   0    4    0    6  |
 A = |  0  -1   -3    2    0  |
     |  0   0    1    0    0  |
     |_ 0   4    2    0    1 _|
```

But we only need to input the non-zero values:

```go
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
```

The dense vector b (we call "dense" because we don't worry about zero values) is:

```
      _      _
     |    8   |
     |   45   |
 b = |   -3   |
     |    3   |
     |_  19  _|
```

Now, we calculate the solution (vector x) using Gosl (indirectly using UMFPACK):

```go
    b := []float64{8.0, 45.0, -3.0, 3.0, 19.0}
    x := la.SpSolve(A, b)
```

Output (execute `go run main.go`):

```
x = [1.000000 2.000000 3.000000 4.000000 5.000000]
xCorrect = [1 2 3 4 5]
```

See [UMFPACK](https://people.engr.tamu.edu/davis/suitesparse.html) for more information.
