package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 int64 }
func genFn0() int8 {
    if true {
        var v5 int8 = -125
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: ("" + ("a" + "\n\t")), F1: 2147483648, F2: 0}
            var v8 uint = 754
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    var v9 int = -204
    var v10 int64 = -9223372036854775808
    v10 = -9223372036854775808
    v9 -= -9223372036854775808
    _ = v10
    _ = v9
    return -40
}
func genFn1(p0 int8, p1 int16, p2 int) bool {
    v11 := GenS1{F0: "\n\t", F1: 65535, F2: -32768}
    v12 := GenS1{F0: "\n\t", F1: 201, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return false
}

func main() {
    v1 := []uint32{2147483648, 509, 2147483648}
    var v2 int32 = 1
    v3 := []uint32{}
    v4 := make(map[int]bool)
    switch "a" {
    case "a":
        var v13 float64 = 0.1
        _ = v13
    default:
        if false {
            fmt.Println(v2)
            var v14 int16 = -385
            _ = v14
        } else {
            v15 := GenS2{F0: 18446744073709551615, F1: 1, F2: 18446744073709551615}
            v16 := v4[127]
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = v15
            _ = v16
        }
    }
    switch 9223372036854775807 {
    case -625:
        var v17 bool = false
        _ = v17
    case 2147483647:
        v2 = 42
    case 4294967295:
        for fi18 := 0; fi18 < 1; fi18++ {
            switch (fi18 >> ((65535) & 63)) {
            case -1:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v2 = (v2 * (v2 % ((1) | 1)))
            _ = fi18
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    v4[32768] = true
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v19 int32 = v2
    v2 = v2
    var v20 int = 2147483647
    v21 := GenS3{F0: "\n\t", F1: -838}
    var v22 float64 = 0.1
    fmt.Println(v22, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
