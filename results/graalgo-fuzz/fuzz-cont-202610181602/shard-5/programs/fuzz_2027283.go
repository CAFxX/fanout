package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }
func genFn0(p0 int32) int32 {
    var v6 uint64 = 2
    p0 = p0
    p0 -= 42
    v7 := []int{-461, 32767, -9223372036854775808}
    for fi8 := 0; fi8 < 10; fi8++ {
        switch (fi8 | (fi8 + 2147483647)) {
        case 9223372036854775807:
            v9 := GenS1{F0: 76, F1: 0, F2: (v6 | 1)}
            _ = v9
        case -219:
            var v10 int64 = 63
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return (p0 >> ((42) & 63))
}
func genFn1(p0 uint8, p1 uint8, p2 string) uint32 {
    v11 := GenS1{F0: 42, F1: 42, F2: 580}
    var v12 float64 = 1.5
    switch "\n\t" {
    case "\n\t":
        v13 := []int{-2147483648, -9223372036854775808, -2147483648}
        _ = v13
    case "foo bar":
        p0 = (p0 & (^(p0)))
    default:
        p2 = p2
    }
    v12 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 71
    var v3 string = "\n\t"
    var v4 int8 = -1
    v5 := make(map[string]uint32)
    v4 = v4
    v3 = ("" + v3)
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float64 = (0.0 + (-2.25 * 1e10))
        for fi16 := 0; fi16 < 6; fi16++ {
            v3 = "foo bar"
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v4 += (v4 & v4)
    var v17 uint = v2
    var v18 int64 = -860
    fmt.Println(v3, v17, v4)
    v17 = (v17 | v2)
    var v19 bool = (-1 != 1000)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
