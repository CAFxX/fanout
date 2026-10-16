package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 bool, p1 int) bool {
    p1 = -9223372036854775808
    p1 = -2147483648
    _ = p0
    _ = p1
    return false
}
func genFn1(p0 bool, p1 bool, p2 uint32) uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 646, F1: uint(p2), F2: 0}
        fmt.Println(fi6, p1, p1)
        v8 := GenS1{F0: 1, F1: 2, F2: -2147483648}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p1, p0)
    v9 := []int{-9223372036854775808, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 2147483647
    var v3 uint8 = 255
    var v4 bool = false
    var v5 int16 = 1000
    var v10 int = 127
    v11 := GenS2{F0: 255, F1: 32767}
    fmt.Println(v2, v4, v10)
    var v12 uint16 = 1
    var v13 string = ""
    var v14 uint = 65535
    v15 := []int{int(v14), (v10 << (((^(v14))) & 63)), (^(v10)), 323}
    if (v4 || v4) {
        v1["foo bar"] = 2147483647
        delete(v1, "foo bar")
        if true {
            for fi16 := 0; fi16 < 19; fi16++ {
                for fi17 := 0; fi17 < 19; fi17++ {
                    v18 := v1[v13]
                    _ = fi17
                    _ = v18
                }
                v1[v13] = 2147483648
                fmt.Println(v3)
                _ = fi16
            }
            v13 = v13
        }
    } else {
        var v19 int16 = (v5 >> ((v14) & 63))
        fmt.Println(v2, v12, v12)
        fmt.Println(v10, v19)
        _ = v19
    }
    v2 = v2
    v2 = int32(v5)
    v12 = v12
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(v5, v5)
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS3{F0: 0, F1: 2}
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    if (v3 < v3) {
        var v23 bool = v4
        switch (v10 * (v10 * -250)) {
        case 127:
            v24 := GenS3{F0: 0, F1: 19}
            _ = v24
        case -221:
            _ = len(v15)
        default:
            v25 := GenS4{F0: 4294967295, F1: (1 == 65535)}
            _ = v25
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 float64 = 0.0
        _ = fi26
        _ = v27
    }
    if (v13 != v13) {
        v13 = v13
        v28 := GenS2{F0: 4294967295, F1: (v12 ^ v12)}
        var v29 uint32 = 2
        _ = v28
        _ = v29
    }
    switch "\n\t" {
    case "":
        v10 += -853
    default:
        v30 := GenS4{F0: 0, F1: (-2147483648 == v2)}
        _ = v30
    }
    v10 += (^(v10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
