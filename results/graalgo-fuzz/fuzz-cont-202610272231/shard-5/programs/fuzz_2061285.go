package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 uint, p2 string) int64 {
    fmt.Println(p1, p1, p2)
    if (1 > -359) {
        p1 = 363
        p1 = p1
        p0 = (-(p0))
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int32, p1 int8, p2 uint8) string {
    switch "hello" {
    case "\n\t":
        switch -57 {
        case 300:
            v6 := []int{}
            _ = v6
        case 4294967295:
            var v7 uint = 65535
            _ = v7
        default:
            p0 += p0
        }
    case "foo bar":
        v8 := GenS1{F0: 42, F1: 42}
        _ = v8
    case "a":
        var v9 uint64 = 453
        _ = v9
    default:
        var v10 string = ""
        _ = v10
    }
    if false {
        v11 := GenS2{F0: 90, F1: 2147483647}
        var v12 int32 = p0
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int8 = 127
            p1 += (p1 - 0)
            switch "" {
            case "a":
                v14 += p1
            case "\n\t":
                var v15 int32 = 0
                _ = v15
            case "hello":
                var v16 int = -597
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        v17 := GenS3{F0: 85, F1: (402 >= 42), F2: 1}
        v18 := []int{51, 42, 536}
        _ = v17
        _ = v18
    }
    p0 += (p0 * 0)
    var v19 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return "\n\t"
}

func main() {
    var v1 int64 = -699
    var v2 uint8 = 145
    var v3 int = 78
    var v4 uint = 65535
    var v5 uint16 = 177
    for fi20 := 0; fi20 < 9; fi20++ {
        fmt.Println(fi20, v4)
        _ = fi20
    }
    v2 = v2
    v21 := GenS4{F0: 1}
    var v22 uint8 = (v2 + (^(v2)))
    var v23 uint64 = 2
    var v24 uint8 = (^(v22))
    v24 += uint8(v23)
    v25 := []int{int(v2)}
    fmt.Println(v22)
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint64 = (v23 << (((v4 + 504)) & 63))
        v3 += -9223372036854775808
        _ = fi26
        _ = v27
    }
    v22 += v22
    var v28 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
