package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int8, p1 uint16) int8 {
    var v7 uint = 65535
    v8 := []int{729}
    fmt.Println(p0)
    v7 += (^(v7))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 1
    var v3 uint = 18446744073709551615
    var v4 int16 = 0
    v5 := []uint32{671}
    var v6 uint64 = 130
    v1 += (v1 * (v1 % ((v1) | 1)))
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int = fi9
        var v11 uint = 928
        for fi12 := 0; fi12 < 11; fi12++ {
            v10 = fi12
            switch "\n\t" {
            case "hello":
                fmt.Println(fi9, v1)
            case "foo bar":
                v6 = (v6 ^ 820)
            }
            var v13 int8 = -128
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v14 := GenS1{F0: v2, F1: v6}
    v5 = append(v5, (v2 - 2))
    var v15 uint64 = v6
    fmt.Println(v1, v15)
    v2 = 2147483647
    v5 = append(v5, 963)
    var v16 uint64 = v6
    if (!(1e10 <= 1e10)) {
        if (1 < 0) {
            _ = len(v5)
            if ((^(v2)) != v2) {
                fmt.Println(v1, v15, v6)
                v17 := GenS1{F0: v2, F1: v15}
                _ = v17
            } else {
                v16 = (v16 >> ((65535) & 63))
            }
            fmt.Println(v2)
        }
        v4 = v4
        _ = len(v5)
    } else {
        var v18 int32 = 853
        v19 := GenS1{F0: (v2 + (^(v2))), F1: (^(v16))}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        v5 = append(v5, (v2 + 209))
        _ = fi20
    }
    v15 = v6
    v21 := GenS1{F0: (^(v2)), F1: v6}
    var v22 int64 = v1
    var v23 int32 = -15
    v15 = (v6 ^ uint64(v3))
    var v24 uint16 = 585
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
