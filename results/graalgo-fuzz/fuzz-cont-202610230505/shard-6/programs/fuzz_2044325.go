package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 int32, p2 int32) uint64 {
    var v5 int8 = 42
    p0 += 4294967295
    v6 := GenS1{F0: 65535, F1: 0}
    var v7 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (^(p0))
}
func genFn1(p0 int8) int16 {
    fmt.Println(p0)
    var v8 uint = 65535
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := []int{128, 9223372036854775807, -574}
    var v2 uint16 = 596
    var v3 int64 = 0
    var v4 uint32 = 2147483648
    switch -486 {
    case -9223372036854775808:
        v9 := GenS1{F0: (v2 >> ((1) & 63)), F1: (v2 ^ (v2 >> ((18446744073709551615) & 63)))}
        _ = v9
    case 9223372036854775807:
        fmt.Println(v3, v3)
    case 127:
        v2 -= 2
    default:
        var v10 string = (("\n\t" + "a") + "a")
        _ = v10
    }
    v11 := GenS1{F0: v2, F1: (v2 + (v2 | v2))}
    v3 = (v3 >> ((865) & 63))
    v4 = (v4 >> ((2) & 63))
    v12 := GenS1{F0: v2, F1: (v2 >> ((190) & 63))}
    v4 = (v4 & (^(v4)))
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v3, v3, v3)
        if (true || true) {
            v1 = append(v1, -1)
        }
        _ = fi13
    }
    v14 := GenS1{F0: (^(v2)), F1: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
