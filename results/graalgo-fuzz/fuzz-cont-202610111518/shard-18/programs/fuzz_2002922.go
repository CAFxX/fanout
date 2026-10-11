package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint8, p1 int32) uint64 {
    p1 -= (p1 + 2147483647)
    var v6 int16 = -32768
    p0 = p0
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 0
    var v3 int64 = -579
    v4 := []int{v2, (v2 + (v2 % ((-9) | 1))), (v2 << ((42) & 63))}
    v5 := []int{(^(v2)), v2, 742}
    switch (v2 ^ v2) {
    case 572:
        fmt.Println(v1, v3)
    case 32768:
        v2 += (-(v2))
    default:
        v1 = ("hello" + v1)
    }
    var v7 int32 = 861
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v3 += v3
        fmt.Println(fi8)
        v9 := GenS1{F0: 4294967295, F1: (v3 & (^(v3)))}
        _ = fi8
        _ = v9
    }
    v7 += (v7 << ((0) & 63))
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint = 1
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
        switch ((v1 + v1) + "\n\t") {
        case "\n\t":
            var v12 string = v1
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch v1 {
        case "a":
            fmt.Println(v7)
        case "hello":
            _ = len(v5)
        }
        var v14 uint16 = 42
        v15 := GenS1{F0: 0, F1: (v3 ^ (v3 * v3))}
        _ = fi13
        _ = v14
        _ = v15
    }
    v7 = v7
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v2)
        var v17 uint64 = genFn0(128, (v7 >> ((4294967295) & 63)))
        _ = fi16
        _ = v17
    }
    if len(v5) > 0 {
        v5[0] = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
