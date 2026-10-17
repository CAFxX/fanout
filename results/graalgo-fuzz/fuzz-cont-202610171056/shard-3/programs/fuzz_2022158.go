package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 uint16) int {
    var v5 uint32 = 565
    v6 := GenS1{F0: 2, F1: 150, F2: 18446744073709551615}
    v5 = 308
    var v7 int64 = 42
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return int(v5)
}
func genFn1(p0 string, p1 int, p2 int64) uint8 {
    switch p1 {
    case 42:
        p1 += p1
    case 749:
        p2 = p2
    case -1:
        var v8 uint64 = 137
        _ = v8
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{(fi9 >> ((42) & 63)), (fi9 ^ (fi9 * (fi9 + 32768))), (fi9 ^ (fi9 % ((fi9) | 1))), genFn0(1)}
            _ = fi9
            _ = v10
        }
    }
    var v11 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 128
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v12 int64 = -9223372036854775808
    v1 = (v1 + v1)
    fmt.Println(v1, v12)
    var v13 int64 = (v12 << ((18446744073709551615) & 63))
    var v14 float32 = 3.14159
    v15 := GenS1{F0: 69, F1: 4294967295, F2: 603}
    if true {
        var v16 int64 = (^(v12))
        _ = v16
    }
    if true {
        var v17 uint64 = (v1 / ((4294967295) | 1))
        var v18 int64 = 9223372036854775807
        _ = v17
        _ = v18
    }
    var v19 uint32 = 856
    fmt.Println(v12, v13)
    fmt.Println(v13)
    fmt.Println(v19)
    switch "hello" {
    case "\n\t":
        if (!(!((!false) || true))) {
            fmt.Println(v14, v12)
        }
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            switch "\n\t" {
            case "foo bar":
                var v21 string = "foo bar"
                _ = v21
            case "a":
                v22 := GenS2{F0: 0, F1: 690}
                _ = v22
            }
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
