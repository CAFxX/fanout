package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483648
    var v5 int16 = 707
    if true {
        v1 = v1
    } else {
        v1 += v1
        delete(v3, 544)
        switch "a" {
        case "":
            switch -1 {
            case -141:
                switch "hello" {
                case "\n\t":
                    v6 := []int{1, -102}
                    _ = v6
                }
            }
        case "a":
            if ((1000 == 1000) || ((v1 + 1) == 356)) {
                v7 := v3[v4]
                switch 166 {
                case 42:
                    v8 := v2[(v7 + v7)]
                    _ = v8
                case 722:
                    for fi9 := 0; fi9 < 9; fi9++ {
                        var v10 string = v7
                        v10 = "a"
                        v11 := GenS1{F0: -32768, F1: 42}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 9223372036854775807:
                    if (v4 <= uint32(v1)) {
                        if ((1e10 + 0.0) <= -2.25) {
                            v7 = v7
                            var v12 uint16 = uint16(v4)
                            _ = v12
                        } else {
                            v4 += v4
                            switch -1 {
                            case 42:
                                v7 = v7
                            }
                        }
                        var v13 float32 = 3.14159
                        var v14 float64 = -2.25
                        _ = v13
                        _ = v14
                    }
                }
                v15 := GenS1{F0: (v5 | (v5 << ((18446744073709551615) & 63))), F1: (v1 - (v1 | v1))}
                _ = v15
                _ = v7
            } else {
                fmt.Println(v4)
            }
        case "foo bar":
            var v16 float64 = 3.14159
            _ = v16
        default:
            v17 := GenS1{F0: v5, F1: (v1 | v1)}
            _ = v17
        }
    }
    _ = len(v3)
    switch 128 {
    case 1:
        var v18 uint32 = v4
        _ = v18
    case 240:
        v4 = (v4 - v4)
    case 9223372036854775807:
        v19 := GenS1{F0: -163, F1: uint16(v4)}
        _ = v19
    default:
        v20 := GenS1{F0: v5, F1: (v1 / ((695) | 1))}
        _ = v20
    }
    v21 := v3[v4]
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            v21 = (v21 + "a")
            v23 := GenS1{F0: v5, F1: v1}
            _ = v23
        }
        _ = fi22
    }
    var v24 uint = 1
    var v25 bool = (v21 < v21)
    v26 := []int{-1, -1}
    var v27 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
