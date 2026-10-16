package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 497
    v3 := []int{}
    var v4 int8 = -1
    for fi5 := 0; fi5 < 11; fi5++ {
        if true {
            var v6 float32 = 3.14159
            v2 = (v2 & v2)
            switch -707 {
            case 9223372036854775807:
                if true {
                    v6 -= float32(-2.25)
                    v7 := GenS1{F0: 23, F1: false}
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        var v8 int32 = (^(v2))
                        _ = v8
                    case -986:
                        v9 := GenS1{F0: 42, F1: false}
                        _ = v9
                    case 0:
                        v6 = v6
                    }
                    _ = v7
                } else {
                    v6 += v6
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case 32767:
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS1{F0: 1, F1: false}
                    var v12 uint64 = 869
                    var v13 int = fi10
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case -575:
                fmt.Println(v4)
            }
            _ = v6
        }
        v14 := GenS1{F0: 2, F1: ((!(-2.25 < 3.14159)) || true)}
        _ = fi5
        _ = v14
    }
    v4 = -1
    var v15 float64 = 0.1
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS1{F0: 128, F1: false}
            v19 := GenS1{F0: 42, F1: ((-801 < -22) && (42 >= 42))}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = len(v1)
        _ = fi16
    }
    if true {
        v20 := GenS1{F0: 128, F1: (18446744073709551615 > 2)}
        v21 := GenS2{F0: 0}
        _ = v20
        _ = v21
    } else {
        v2 = -2147483648
        v22 := GenS3{F0: 63, F1: 919, F2: false}
        var v23 uint16 = 824
        _ = v22
        _ = v23
    }
    switch 2147483647 {
    case 128:
        v24 := GenS3{F0: 0, F1: 1, F2: true}
        _ = v24
    case -1:
        v2 -= (v2 & 941)
    default:
        if ((-471 != -1) || true) {
            fmt.Println(v4, v4, v4)
            v25 := v1[42]
            _ = v25
        }
    }
    var v26 string = "foo bar"
    v2 = (v2 % ((v2) | 1))
    if false {
        var v27 float64 = v15
        _ = v27
    } else {
        var v28 int = -1
        fmt.Println(v4)
        v4 -= (^(v4))
        _ = v28
    }
    v26 = (v26 + v26)
    v2 -= v2
    var v29 float32 = (0.0 / (1.5 + (((float32((((float64(-2.25) + 0.0) * float64(1e10)) / (((0.0 * (0.0 + float64(3.14159))) / (0.0 - v15)) + (((v15 * 0.1) / (v15 / (v15 - float64((0.0 * (0.1 + (1e10 - (0.0 * ((((0.1 + 0.1) * (3.14159 + 3.14159)) * (1e10 / -2.25)) / -2.25))))))))) / ((((v15 / 1.5) / v15) * v15) / v15))))) / 0.1) / 1.5) + 1.5)))
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
