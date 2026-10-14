package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int, p1 int) int {
    v4 := GenS1{F0: p0}
    var v5 int16 = -757
    v6 := GenS2{F0: 2, F1: p1, F2: 4294967295}
    var v7 int64 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p1
}
func genFn1() int64 {
    var v8 bool = (!((true || (!(false && true))) && (3.14159 == 0.1)))
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = (fi9 * (fi9 | fi9))
        var v11 int = genFn0(v10, v10)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    var v3 uint8 = 42
    v12 := GenS1{F0: 0}
    v13 := GenS3{F0: 1}
    v14 := GenS4{F0: 9223372036854775807, F1: ("\n\t" + "foo bar"), F2: 0}
    v15 := v1[2147483647]
    v2 = v2
    v16 := v1[9223372036854775807]
    v2 = v2
    v17 := GenS5{F0: 1, F1: uint64(v3)}
    switch 42 {
    case 128:
        var v18 int8 = 100
        _ = v18
    default:
        if (!true) {
            v3 = v3
            var v19 int64 = 9223372036854775807
            _ = v19
        }
    }
    switch -575 {
    case 32768:
        var v20 uint16 = 772
        _ = v20
    }
    v2 -= v2
    v21 := v1[-170]
    for fi22 := 0; fi22 < 13; fi22++ {
        switch (fi22 & fi22) {
        case -213:
            fmt.Println(v15, v16, v2)
        case -796:
            if (true && false) {
                v2 -= v2
            } else {
                v23 := GenS6{F0: 817}
                var v24 float32 = 0.1
                _ = v23
                _ = v24
            }
        case -9223372036854775808:
            v25 := GenS6{F0: (v2 >> ((42) & 63))}
            _ = v25
        }
        fmt.Println(v16, v15)
        _ = fi22
    }
    var v26 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v26
    _ = v3
}
