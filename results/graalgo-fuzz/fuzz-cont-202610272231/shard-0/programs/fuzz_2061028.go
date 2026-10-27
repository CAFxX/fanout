package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint) uint8 {
    if (1 > 815) {
        for fi4 := 0; fi4 < 19; fi4++ {
            p0 = (p0 / (((p0 & p0)) | 1))
            _ = fi4
        }
    } else {
        fmt.Println(p0)
        fmt.Println(p0)
        switch -9223372036854775808 {
        case 32768:
            var v5 uint16 = 2
            _ = v5
        case 1:
            var v6 uint64 = 18446744073709551615
            _ = v6
        case 2147483647:
            v7 := []int{9223372036854775807, -438, 9223372036854775807}
            _ = v7
        default:
            fmt.Println(p0)
        }
    }
    v8 := []int{-9223372036854775808, -1, 0, 4294967295}
    var v9 uint64 = 42
    v10 := GenS1{F0: 128}
    var v11 int16 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint16 = 0
    var v2 int64 = 9223372036854775807
    var v3 int = 9223372036854775807
    v12 := []int{(-(v3)), 1}
    v2 += (v2 % (((v2 & 1)) | 1))
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v1 += 65535
    default:
        v1 -= (v1 - 317)
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v2 += v2
        var v14 uint32 = 2147483648
        for fi15 := 0; fi15 < 15; fi15++ {
            v1 = v1
            v16 := GenS2{F0: -1}
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    if (18446744073709551615 > 2) {
        var v17 uint64 = 597
        if false {
            v18 := GenS3{F0: (v2 << ((193) & 63)), F1: -32768, F2: 116}
            var v19 uint16 = v1
            _ = v18
            _ = v19
        } else {
            fmt.Println(v17)
            fmt.Println(v3)
            switch 198 {
            case 1:
                var v20 bool = (1 == v17)
                _ = v20
            case 9223372036854775807:
                var v21 float32 = (3.14159 + (1.5 / 0.1))
                _ = v21
            default:
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int16 = 1000
                    fmt.Println(fi22)
                    _ = fi22
                    _ = v23
                }
            }
        }
        _ = v17
    }
    v24 := GenS4{F0: 556, F1: 240}
    v2 += (^(v2))
    switch v3 {
    case 392:
        _ = len(v12)
    default:
        _ = len(v12)
    }
    v25 := GenS2{F0: -1}
    if (42 == 2) {
        var v26 float64 = float64(0.0)
        _ = v26
    }
    var v27 int64 = v2
    v3 -= v3
    var v28 int32 = 0
    var v29 int8 = 1
    var v30 uint64 = 33
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
