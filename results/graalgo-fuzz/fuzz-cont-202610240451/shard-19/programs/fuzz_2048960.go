package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0() bool {
    var v4 int = -276
    fmt.Println(v4)
    _ = v4
    return false
}
func genFn1(p0 uint16, p1 string) uint64 {
    v5 := GenS1{F0: 0}
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 255}
        _ = fi6
        _ = v7
    }
    var v8 bool = ((18446744073709551615 == 255) || true)
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int32 = int32(p0)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 163
    v2 := make(map[string]string)
    var v3 uint32 = 1
    v1 += -9223372036854775808
    v11 := v2[("a" + "a")]
    v12 := []int{88, -1, 68}
    v13 := GenS1{F0: 1}
    v14 := GenS1{F0: 82}
    var v15 int = 371
    var v16 int = (^(v15))
    v3 = (v3 * 2147483648)
    if true {
        v1 = v1
        fmt.Println(v11)
    }
    var v17 string = v11
    for fi18 := 0; fi18 < 13; fi18++ {
        v2["\n\t"] = "a"
        _ = fi18
    }
    switch v16 {
    case 128:
        v3 = (v3 ^ 4294967295)
    default:
        v19 := GenS1{F0: uint8(v3)}
        _ = v19
    }
    fmt.Println(v1, v15)
    var v20 int8 = -126
    if true {
        v3 = v3
        var v21 string = v17
        _ = v21
    } else {
        if (!false) {
            var v22 float64 = (1.5 / 1.0)
            v23 := GenS2{F0: 1, F1: v3}
            _ = v22
            _ = v23
        } else {
            var v24 uint = 891
            _ = v24
        }
        fmt.Println(v3, v15)
        v25 := GenS3{F0: 42, F1: 830}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
