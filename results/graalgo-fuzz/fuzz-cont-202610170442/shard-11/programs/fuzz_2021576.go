package main

import (
    "fmt"
)

func genFn0() uint8 {
    if (18446744073709551615 <= 429) {
        var v6 int8 = -128
        switch "\n\t" {
        case "\n\t":
            v7 := []int{-1, 2147483647, 42}
            _ = v7
        case "a":
            var v8 bool = true
            _ = v8
        case "hello":
            v6 = v6
        }
        v6 = (v6 | v6)
        _ = v6
    }
    var v9 int8 = 42
    _ = v9
    return 29
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 1
    var v3 float64 = 1.5
    v4 := []int{4294967295, 32767}
    v5 := make(map[string]uint32)
    var v10 bool = true
    fmt.Println(v3, v1, v1)
    for fi11 := 0; fi11 < 15; fi11++ {
        v3 = float64((1e10 + 0.0))
        _ = fi11
    }
    _ = len(v5)
    switch ("" + "") {
    case "hello":
        fmt.Println(v10, v10)
    }
    var v12 float32 = 1.5
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int8 = (v1 & (v1 << ((2) & 63)))
        var v15 uint64 = 0
        var v16 uint8 = 231
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int64 = 386
    if v10 {
        if false {
            v12 = 3.14159
        }
        var v18 int64 = 42
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                v1 = (v1 | (-(v1)))
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
