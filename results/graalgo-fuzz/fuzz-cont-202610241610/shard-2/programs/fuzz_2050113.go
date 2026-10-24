package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 625
    v3 := make(map[uint32]bool)
    var v4 uint = 656
    v5 := make(map[string]uint32)
    v6 := []uint32{42, 2, uint32(v4), 168}
    v4 -= 18446744073709551615
    v3[65535] = (!(!false))
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = -9223372036854775808
        v4 += v4
        var v10 uint = (v4 * v4)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint64 = (^(v2))
    switch ("foo bar" + ("" + "")) {
    case "":
        var v12 uint = v4
        _ = v12
    case "\n\t":
        var v13 bool = (3.14159 >= 0.1)
        _ = v13
    case "foo bar":
        v14 := GenS1{F0: v11}
        _ = v14
    }
    fmt.Println(v11, v11)
    if (-31 > v1) {
        v1 += v1
        v4 = 121
        var v15 float64 = 1.5
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint32 = 42
        var v18 int32 = -1
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int32 = 812
    switch 32767 {
    case 1:
        var v20 float32 = (float32(float64(-2.25)) * -2.25)
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
