package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    var v3 uint = 65535
    v4 := GenS1{F0: 18446744073709551615, F1: false, F2: (128 < 1)}
    fmt.Println(v3, v2)
    var v5 uint64 = 18446744073709551615
    switch ("foo bar" + "") {
    case "\n\t":
        for fi6 := 0; fi6 < 9; fi6++ {
            v5 = (v5 & (v5 << ((v3) & 63)))
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{fi6}
                var v9 uint64 = v5
                _ = fi7
                _ = v8
                _ = v9
            }
            v5 = v5
            _ = fi6
        }
    case "a":
        fmt.Println(v5)
    default:
        v10 := []int{573, 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            var v12 int32 = 81
            v13 := []int{(fi11 & (fi11 | fi11)), fi11, (fi11 / (((^(fi11))) | 1)), fi11}
            v12 -= v12
            _ = v12
            _ = v13
        }
        if true {
            var v14 uint = (^(v3))
            v15 := []int{fi11, (fi11 & (fi11 & (^(fi11)))), fi11}
            v16 := GenS1{F0: (v5 | (v5 & (v5 % ((uint64(v2)) | 1)))), F1: ("foo bar" != (("" + "") + (("" + (("" + ("hello" + "foo bar")) + "foo bar")) + "hello"))), F2: true}
            _ = v14
            _ = v15
            _ = v16
        } else {
            v17 := GenS2{F0: 0, F1: (v5 | 18446744073709551615), F2: (fi11 >> ((255) & 63))}
            v5 += (v5 << ((v3) & 63))
            _ = v17
        }
        var v18 uint64 = v5
        _ = fi11
        _ = v18
    }
    if (false || ((65535 == 42) || true)) {
        fmt.Println(v2, v5)
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int64 = 1
            var v21 uint8 = uint8(v5)
            fmt.Println(v21, v3, fi19)
            _ = fi19
            _ = v20
            _ = v21
        }
        if (42 > 0) {
            fmt.Println(v2)
        }
    }
    var v22 float64 = 1e10
    v3 = v3
    var v23 int8 = -128
    _ = len(v1)
    if ("hello" < "") {
        fmt.Println(v2, v5, v23)
        v23 = (v23 + -95)
    }
    v2 += v2
    v24 := []int{-84, 2147483647}
    v3 = (v3 ^ (v3 % ((1) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
