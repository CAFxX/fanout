package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 int16) int64 {
    if (967 >= 1) {
        v6 := GenS1{F0: -197}
        _ = v6
    } else {
        var v7 int8 = 127
        _ = v7
    }
    var v8 float64 = 3.14159
    var v9 uint64 = 4294967295
    fmt.Println(p0)
    _ = p0
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint = 42
    v2 := []int{-2147483648, 983, 2147483647}
    var v3 uint64 = 2
    var v4 float64 = 1e10
    v5 := []uint32{}
    fmt.Println(v3, v1)
    var v10 float64 = 0.0
    switch 32768 {
    case -830:
        v11 := GenS1{F0: -422}
        _ = v11
    case 9223372036854775807:
        v3 += v3
    case -536:
        var v12 float32 = 1e10
        _ = v12
    default:
        switch ("foo bar" + ((("\n\t" + "foo bar") + "foo bar") + "\n\t")) {
        case "hello":
            if (32767 >= 464) {
                for fi13 := 0; fi13 < 19; fi13++ {
                    if (v3 != 4294967295) {
                        var v14 uint8 = 23
                        _ = v14
                    } else {
                        v10 = ((v10 - float64(float32((v10 / 1.5)))) - (v10 * 3.14159))
                        if false {
                            var v15 uint = (^(v1))
                            switch 286 {
                            case 4294967295:
                                v2 = append(v2, (fi13 % ((int(v3)) | 1)))
                            case -52:
                                v1 = v1
                            case 127:
                                var v16 int8 = 1
                                _ = v16
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 376
                                }
                            }
                            _ = v15
                        } else {
                            v17 := GenS2{F0: 541, F1: 552, F2: -576}
                            _ = v17
                        }
                        var v18 uint8 = 127
                        _ = v18
                    }
                    _ = fi13
                }
                var v19 float64 = v4
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v19
            }
        case "foo bar":
            fmt.Println(v4, v4)
        case "a":
            var v20 uint = v1
            _ = v20
        }
    }
    v10 -= 3.14159
    for fi21 := 0; fi21 < 8; fi21++ {
        _ = len(v5)
        _ = len(v5)
        _ = fi21
    }
    v1 = v1
    switch "a" {
    case "\n\t":
        var v22 uint32 = 556
        _ = v22
    }
    var v23 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
