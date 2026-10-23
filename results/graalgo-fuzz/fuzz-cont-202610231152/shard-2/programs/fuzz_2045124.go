package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 string, p1 int32, p2 int8) uint8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{(fi6 | (^(fi6))), (fi6 & fi6)}
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 42
    var v9 int32 = 42
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 float32 = (3.14159 + 0.0)
    var v2 uint = 65535
    var v3 float64 = (0.1 * (((1.5 * 0.0) - (3.14159 - 0.0)) * -2.25))
    var v4 float64 = (v3 - float64(v1))
    v5 := make(map[uint32]bool)
    var v11 uint = v2
    v12 := GenS1{F0: 120, F1: 1, F2: 142}
    var v13 int = 32767
    if false {
        for fi14 := 0; fi14 < 13; fi14++ {
            if ((!(!(!false))) || true) {
                if true {
                    switch (fi14 ^ 752) {
                    case 9223372036854775807:
                        switch v13 {
                        case -1:
                            switch "" {
                            case "\n\t":
                                for fi15 := 0; fi15 < 4; fi15++ {
                                    v16 := GenS2{F0: 127, F1: 1}
                                    var v17 float64 = (1e10 / v4)
                                    var v18 int32 = -609
                                    var v19 uint8 = 42
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            case "":
                                fmt.Println(fi14, v4)
                            case "foo bar":
                                v20 := GenS2{F0: 128, F1: -128}
                                _ = v20
                            default:
                                v11 = v11
                            }
                        case -599:
                            var v21 string = "\n\t"
                            _ = v21
                        case 42:
                            v22 := []int{(v13 % ((v13) | 1))}
                            _ = v22
                        }
                    case -476:
                        fmt.Println(v3)
                    }
                }
                v3 -= v3
            } else {
                v5[2147483647] = false
            }
            _ = fi14
        }
        var v23 int16 = 1
        fmt.Println(v11)
        _ = v23
    } else {
        if true {
            v3 = float64(v1)
            fmt.Println(v13, v11, v4)
        } else {
            var v24 int8 = 0
            var v25 float64 = v3
            _ = v24
            _ = v25
        }
        var v26 uint64 = 1
        _ = v26
    }
    v3 = 3.14159
    var v27 uint64 = 2
    v2 -= v2
    fmt.Println(v1)
    v28 := []int{v13, v13, v13, v13}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
