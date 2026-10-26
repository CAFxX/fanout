package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 string) int32 {
    var v5 int32 = -750
    var v6 uint16 = 0
    v7 := GenS1{F0: 128}
    var v8 uint64 = 2
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (-(v5))
}
func genFn1(p0 uint16, p1 uint64, p2 uint8) int {
    var v9 uint16 = (p0 | p0)
    fmt.Println(v9, v9, p1)
    v10 := []int{}
    v11 := GenS2{F0: 1, F1: (p2 >> ((18446744073709551615) & 63)), F2: 70}
    fmt.Println(v9, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 43
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 866
    var v3 uint32 = 397
    v4 := []string{}
    v2 += v2
    v3 -= (v3 / ((v3) | 1))
    _ = len(v4)
    switch 32767 {
    case -9223372036854775808:
        var v12 uint8 = 145
        _ = v12
    case 1:
        v2 += (v2 ^ (v2 / (((v2 + 42)) | 1)))
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int16 = 1
        switch "hello" {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            v1["a"] = (2 >= 0)
        case "foo bar":
            var v15 uint16 = 255
            _ = v15
        default:
            v16 := GenS1{F0: 42}
            _ = v16
        }
        var v17 int16 = 1
        _ = fi13
        _ = v14
        _ = v17
    }
    var v18 uint32 = 735
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int16 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int32 = genFn0("\n\t")
        v18 = 0
        _ = fi21
        _ = v22
    }
    var v23 float64 = (0.1 * -2.25)
    var v24 bool = (true || (!false))
    v18 += (v18 / (((^(v3))) | 1))
    v18 += v18
    v25 := v1[""]
    var v26 bool = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
