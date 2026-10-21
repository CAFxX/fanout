package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch fi6 {
        case 0:
            v7 := GenS1{F0: true, F1: "foo bar", F2: 612}
            _ = v7
        case 747:
            var v8 uint16 = 2
            _ = v8
        default:
            switch fi6 {
            case 4294967295:
                v9 := GenS2{F0: 0}
                _ = v9
            }
        }
        _ = fi6
    }
    if false {
        v10 := GenS3{F0: 509, F1: (false && (false && false))}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            switch (fi11 - fi11) {
            case 9223372036854775807:
                v12 := []int{}
                _ = v12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := []int{32768, 9223372036854775807, -22, 4294967295}
            for fi15 := 0; fi15 < 13; fi15++ {
                if len(v14) > 0 {
                    v14[0] = 876
                }
                if (!true) {
                    var v16 int8 = 127
                    v16 = (v16 & v16)
                    var v17 uint32 = 2147483647
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
            v14 = append(v14, fi13)
            _ = fi13
            _ = v14
        }
        var v18 float64 = (float64(((0.1 - (((1e10 + 3.14159) / (((0.0 + ((-2.25 / (((0.0 + (0.0 - 1e10)) - 1.5) / 1e10)) / 3.14159)) + 1e10) - 1.5)) / float32(-2.25))) * 0.0)) * float64((1e10 + float32(0.1))))
        _ = v18
    }
    v19 := GenS1{F0: false, F1: "hello", F2: 2147483647}
    var v20 float64 = 0.1
    v20 = 0.1
    _ = v19
    _ = v20
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    var v4 int16 = -32768
    v5 := []int{128, 32768, 42, -339}
    v3 -= int32(v4)
    var v21 uint = 0
    switch ("\n\t" + "a") {
    case "\n\t":
        if false {
            for fi22 := 0; fi22 < 3; fi22++ {
                switch fi22 {
                case -534:
                    var v23 uint = (v21 * (v21 * 804))
                    _ = v23
                default:
                    fmt.Println(v21)
                }
                v3 = genFn0()
                v21 += v21
                _ = fi22
            }
        }
    default:
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint64 = 0
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 uint16 = 42
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    v28 := GenS3{F0: v4, F1: true}
    v1["\n\t"] = (("foo bar" + "a") + "\n\t")
    if true {
        var v29 bool = (65535 != 4294967295)
        fmt.Println(v4)
        var v30 int16 = -942
        _ = v29
        _ = v30
    }
    var v31 float64 = (float64(0.0) / 3.14159)
    v3 = 1
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := GenS4{F0: (v21 / (((v21 & (v21 * v21))) | 1))}
        fmt.Println(v4, v21, v31)
        v21 = 65535
        _ = fi32
        _ = v33
    }
    v34 := GenS5{F0: 223}
    if (v21 > v21) {
        v4 -= (v4 | v4)
    } else {
        v4 = v4
    }
    v31 -= v31
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 int64 = 42
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
