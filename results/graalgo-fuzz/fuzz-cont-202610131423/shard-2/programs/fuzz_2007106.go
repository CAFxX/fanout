package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    var v1 int8 = -128
    var v2 float64 = (float64((1.5 / (1e10 - -2.25))) - 3.14159)
    v3 := []int{127, 128, 1}
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    if true {
        v2 -= 1e10
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := v4[(-(fi6))]
            _ = fi6
            _ = v7
        }
        var v8 string = "hello"
        _ = v8
    } else {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: "a", F1: 149, F2: 42}
            _ = fi9
            _ = v10
        }
        if false {
            for fi11 := 0; fi11 < 8; fi11++ {
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 float32 = (-2.25 / 1e10)
                    switch fi12 {
                    case 32768:
                        var v14 int8 = (v1 ^ 42)
                        _ = v14
                    case -674:
                        var v15 float64 = (v2 / 3.14159)
                        _ = v15
                    default:
                        fmt.Println(fi12, fi11, fi12)
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        } else {
            switch 127 {
            case 32768:
                for fi16 := 0; fi16 < 16; fi16++ {
                    _ = len(v3)
                    _ = fi16
                }
            case 638:
                for fi17 := 0; fi17 < 17; fi17++ {
                    fmt.Println(v1, v1, v1)
                    var v18 uint = 42
                    _ = fi17
                    _ = v18
                }
            case 9223372036854775807:
                switch "" {
                case "\n\t":
                    fmt.Println(v1, v1)
                default:
                    if false {
                        var v19 int16 = 32767
                        var v20 int64 = -109
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 float64 = 1e10
                            v3 = append(v3, 32767)
                            _ = fi21
                            _ = v22
                        }
                        _ = v19
                        _ = v20
                    }
                }
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v4, 32767)
    var v23 int16 = 355
    if ((false && (42 > 0)) || (!true)) {
        fmt.Println(v1, v2, v23)
        v5[-2147483648] = (110 >= -9223372036854775808)
        var v24 int16 = 42
        _ = v24
    }
    var v25 uint32 = 323
    if true {
        var v26 int8 = (v1 * v1)
        _ = v26
    }
    v1 = 77
    v27 := GenS1{F0: "", F1: 255, F2: 0}
    v1 = (v1 - 42)
    var v28 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
