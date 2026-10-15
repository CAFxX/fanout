package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 float32 = 1e10
    var v2 int16 = 42
    v3 := make(map[int]uint32)
    var v4 uint32 = 2147483647
    v2 -= v2
    var v5 uint64 = 2
    fmt.Println(v2)
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: true, F1: fi6, F2: 260}
        v8 := GenS2{F0: 648}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := v3[0]
    for fi10 := 0; fi10 < 3; fi10++ {
        v2 = int16(v9)
        v2 += (v2 % ((v2) | 1))
        v5 -= 2
        _ = fi10
    }
    if false {
        _ = len(v3)
    }
    delete(v3, -308)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 bool = ("foo bar" <= ("\n\t" + "a"))
        _ = fi11
        _ = v12
    }
    fmt.Println(v9, v9, v5)
    switch ("foo bar" + "\n\t") {
    case "":
        v13 := GenS3{F0: 4294967295, F1: 9223372036854775807}
        _ = v13
    case "a":
        var v14 int16 = (-(v2))
        _ = v14
    case "\n\t":
        v1 += (0.0 - 3.14159)
    }
    delete(v3, -989)
    var v15 uint32 = (v4 % ((2147483648) | 1))
    v16 := GenS2{F0: uint(v15)}
    v17 := []int{750}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
