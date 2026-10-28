package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 bool }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
func genFn0() uint {
    v5 := []int{32767, 9223372036854775807, 0}
    var v6 bool = (0 < 762)
    v7 := GenS1{F0: 18446744073709551615, F1: 721, F2: v6}
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1() int {
    var v8 uint64 = 2
    v9 := GenS1{F0: v8, F1: 2, F2: (true && (-707 != 42))}
    v8 = (^(v8))
    v10 := GenS1{F0: (v8 & (v8 / (((v8 << ((0) & 63))) | 1))), F1: 65535, F2: (803 == uint16(v8))}
    var v11 int = 42
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 128
}

func main() {
    var v1 uint16 = 85
    var v2 float32 = 1e10
    var v3 uint8 = 127
    v4 := make(map[string]uint32)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int16 = -97
        var v14 uint8 = 30
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := []int{4294967295, -2147483648, 0}
    for fi16 := 0; fi16 < 2; fi16++ {
        if (false && ((v1 >> ((137) & 63)) == v1)) {
            v2 = v2
            v17 := GenS1{F0: 837, F1: (v1 | v1), F2: false}
            v3 += v3
            _ = v17
        } else {
            var v18 int8 = 0
            _ = v18
        }
        if (fi16 == genFn1()) {
            switch ("" + "hello") {
            case "\n\t":
                v19 := GenS2{F0: 1, F1: -219, F2: 2}
                _ = v19
            case "hello":
                fmt.Println(v2)
            case "a":
                v20 := GenS3{F0: 32767, F1: -128, F2: -1}
                _ = v20
            default:
                fmt.Println(v2, v2)
            }
        } else {
            fmt.Println(v3, v1)
            var v21 uint8 = v3
            var v22 uint8 = (^(v21))
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    v15 = append(v15, 0)
    v23 := GenS3{F0: 32767, F1: 80, F2: 1}
    var v24 int8 = 42
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    fmt.Println(v3, v24, v24)
    v1 -= (^(v1))
    var v25 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
