package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 int, p1 string) int {
    switch p0 {
    case -1:
        var v5 int = (p0 + p0)
        _ = v5
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint16 = 333
            var v8 uint64 = 42
            v7 -= (v7 * (v7 << ((402) & 63)))
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    var v9 uint32 = 1
    var v10 uint32 = v9
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: 4294967295, F1: 42, F2: -521}
        if (3.14159 == 1e10) {
            fmt.Println(v10)
        }
        var v13 uint = 177
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p0
}
func genFn1(p0 int) uint {
    p0 = (p0 & (p0 ^ p0))
    switch "hello" {
    case "hello":
        var v14 int8 = -100
        _ = v14
    case "foo bar":
        p0 -= p0
    case "\n\t":
        v15 := []int{}
        _ = v15
    default:
        fmt.Println(p0)
    }
    _ = p0
    return 218
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 517
    var v3 uint8 = 127
    var v4 uint = 2
    var v16 uint = uint(v2)
    var v17 uint16 = 962
    fmt.Println(v17, v3)
    var v18 int64 = 9223372036854775807
    v19 := v1[1]
    var v20 int64 = (-(v18))
    v17 = v17
    var v21 int8 = 1
    v17 -= (v17 >> ((632) & 63))
    v22 := GenS1{F0: (v16 | (v16 | (v4 / (((^(v16))) | 1)))), F1: -554, F2: int64(v21)}
    var v23 int16 = 0
    if (18446744073709551615 >= 0) {
        var v24 string = ("hello" + v19)
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v16 = v16
        v26 := v1[v2]
        fmt.Println(v23, v4)
        _ = fi25
        _ = v26
    }
    v2 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
