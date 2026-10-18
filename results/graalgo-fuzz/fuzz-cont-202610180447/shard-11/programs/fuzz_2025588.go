package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int) int {
    var v6 uint8 = 0
    v7 := GenS1{F0: 2147483647, F1: -128, F2: 0}
    v8 := []int{(p0 | p0), p0, 538, (p0 % ((2147483647) | 1))}
    switch p0 {
    case 2147483647:
        for fi9 := 0; fi9 < 12; fi9++ {
            if false {
                if ((uint32(v6) != 0) || (!true)) {
                    v8 = append(v8, (p0 * p0))
                } else {
                    p0 = p0
                    var v10 int = p0
                    _ = v10
                }
                p0 += fi9
                v6 = v6
            }
            _ = fi9
        }
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p0, p0)
    case "a":
        var v11 int8 = 0
        _ = v11
    default:
        var v12 float64 = float64(0.1)
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := []string{"\n\t", "a", ("" + "a"), "a"}
    v3 := []int{}
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    if true {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint = 459
            _ = fi13
            _ = v14
        }
    } else {
        v15 := GenS2{F0: 2}
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            v1 = 9223372036854775807
            v18 := GenS2{F0: 1}
            v1 = 9223372036854775807
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS2{F0: 215}
    fmt.Println(v1)
    switch "\n\t" {
    case "a":
        var v20 int32 = 1
        _ = v20
    case "\n\t":
        v21 := GenS3{F0: 42, F1: 192, F2: 4294967295}
        _ = v21
    default:
        var v22 int16 = -1
        _ = v22
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 19; fi23++ {
        switch fi23 {
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = (fi23 + fi23)
            }
        default:
            for fi24 := 0; fi24 < 20; fi24++ {
                if true {
                    fmt.Println(fi23, v1)
                    fmt.Println(fi23, fi23)
                    var v25 int = fi24
                    v25 = (fi24 / ((v25) | 1))
                    _ = v25
                }
                var v26 int = fi24
                _ = fi24
                _ = v26
            }
        }
        fmt.Println(v1)
        _ = fi23
    }
    var v27 uint32 = 4294967295
    var v28 uint64 = 42
    v28 = 1
    var v29 int8 = 0
    v5["\n\t"] = (4294967295 <= (v28 << ((42) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
