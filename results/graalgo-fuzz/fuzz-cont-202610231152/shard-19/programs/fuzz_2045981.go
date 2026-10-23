package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0() int32 {
    var v7 int32 = -1
    var v8 int = 440
    var v9 bool = ((-1 > 127) && (42 > 227))
    v10 := GenS1{F0: v8, F1: 47, F2: 42}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 2, 1}
    var v2 int8 = 100
    var v3 float64 = -2.25
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 uint64 = 315
    switch "foo bar" {
    case "\n\t":
        v2 += (-(v2))
    case "hello":
        v6 = 1
    }
    var v11 uint64 = v6
    v1 = append(v1, 2147483647)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v6)
        switch -9223372036854775808 {
        case 994:
            fmt.Println(fi12)
        }
        _ = fi12
    }
    fmt.Println(v3, v11)
    var v13 string = "a"
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 float64 = (float64(0.1) - (v3 + (v3 - v3)))
        var v16 int32 = 489
        if (((128 <= 0) && (365 >= 533)) || (!true)) {
            var v17 uint64 = v11
            _ = v17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = "a"
    v18 := v5[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
