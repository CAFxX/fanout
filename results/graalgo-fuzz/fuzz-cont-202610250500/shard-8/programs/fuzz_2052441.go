package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint = 4294967295
    v5 := v3["hello"]
    for fi6 := 0; fi6 < 11; fi6++ {
        v5 -= -9223372036854775808
        _ = fi6
    }
    v1 = v1
    v7 := GenS1{F0: -85}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 33}
        var v10 uint = 4294967295
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "foo bar" {
    case "":
        delete(v2, 32767)
    case "hello":
        v1 = -903
    }
    var v11 int8 = 0
    fmt.Println(v4, v11, v5)
    var v12 float64 = 0.0
    switch "\n\t" {
    case "a":
        v13 := GenS2{F0: -58}
        _ = v13
    case "foo bar":
        var v14 uint32 = 195
        _ = v14
    }
    var v15 int = v5
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 string = "\n\t"
        v12 -= 1.5
        _ = fi16
        _ = v17
    }
    v18 := []int{-1, (v5 ^ 4294967295), 2147483647, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
