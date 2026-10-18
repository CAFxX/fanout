package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = (0.1 / float64(1e10))
    var v2 string = "foo bar"
    var v3 float64 = ((float64(1e10) / v1) * v1)
    var v4 uint64 = 397
    fmt.Println(v3, v2)
    v5 := []int{}
    v6 := GenS1{F0: 42}
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = 127
    }
    v7 := GenS2{F0: 4294967295}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
