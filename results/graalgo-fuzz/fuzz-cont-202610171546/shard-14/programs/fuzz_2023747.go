package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 string) bool {
    v7 := GenS1{F0: 0, F1: true}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return (p1 != p1)
}
func genFn1(p0 bool, p1 int64, p2 bool) string {
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{}
        v10 := GenS1{F0: 42, F1: (p2 && true)}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return ("" + ("" + ""))
}

func main() {
    var v1 uint8 = 2
    var v2 int16 = 488
    v3 := []uint32{0, 693, 42, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 1
    var v6 int64 = -1
    for fi11 := 0; fi11 < 3; fi11++ {
        _ = len(v4)
        _ = fi11
    }
    _ = len(v3)
    fmt.Println(v1, v6, v2)
    switch 128 {
    case -368:
        v5 -= (v5 << ((409) & 63))
    case -1:
        v5 = 65535
    }
    v3 = append(v3, 42)
    switch "foo bar" {
    case "foo bar":
        var v12 int32 = 42
        _ = v12
    case "\n\t":
        v13 := v4[2147483647]
        _ = v13
    }
    switch 2147483647 {
    case -2147483648:
        var v14 int = 357
        _ = v14
    case 2147483647:
        switch "\n\t" {
        case "foo bar":
            var v15 uint8 = v1
            _ = v15
        default:
            fmt.Println(v1, v5, v6)
        }
    case 4294967295:
        v4[2] = (^(v5))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        _ = len(v3)
        fmt.Println(v2, v5)
        if ((v1 ^ 2) < 1) {
            var v16 uint = 875
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
