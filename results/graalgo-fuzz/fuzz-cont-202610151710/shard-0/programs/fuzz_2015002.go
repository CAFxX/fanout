package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 uint, p1 uint, p2 int) int64 {
    p2 = p2
    var v7 float64 = 0.1
    v8 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -78
}
func genFn1(p0 uint64, p1 uint8, p2 uint64) uint8 {
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float64 = (0.1 + (((1.5 + (-2.25 + -2.25)) / 1.0) / 1.0))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p1
}

func main() {
    v1 := []string{("\n\t" + ""), "foo bar", "hello", "a"}
    var v2 uint32 = 104
    v3 := []string{"\n\t", "foo bar"}
    v4 := []uint32{v2}
    var v5 int8 = 1
    v6 := make(map[int]uint32)
    var v12 uint32 = v2
    var v13 uint64 = 0
    switch 42 {
    case -971:
        fmt.Println(v12, v5, v13)
    case 128:
        v14 := GenS1{F0: 255}
        _ = v14
    }
    v15 := GenS1{F0: 4294967295}
    v5 += v5
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            fmt.Println(v5, fi16, v13)
            v13 += 2
            v17 := GenS2{F0: "\n\t", F1: 1}
            _ = fi16
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 0}
        _ = v18
    }
    if true {
        if (!false) {
            if false {
                var v19 float32 = (float32(1.5) / 3.14159)
                v20 := GenS1{F0: 0}
                _ = v19
                _ = v20
            }
            if ((("hello" + "foo bar") + "hello") < ("" + "\n\t")) {
                delete(v6, 127)
                v21 := GenS2{F0: "\n\t", F1: 255}
                v6[1] = 220
                _ = v21
            } else {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint = 972
                    v5 = v5
                    _ = fi22
                    _ = v23
                }
                v2 = v12
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            fmt.Println(v13, v2, v5)
        } else {
            v4 = append(v4, (v12 * v2))
            var v24 int32 = 2147483647
            _ = v24
        }
    } else {
        v25 := v6[-738]
        _ = v25
    }
    v6[-174] = (v2 * v2)
    fmt.Println(v2, v2, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
