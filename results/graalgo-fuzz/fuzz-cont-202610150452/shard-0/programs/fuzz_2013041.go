package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint64, p1 uint8, p2 int32) uint64 {
    p0 -= (^(p0))
    if false {
        switch 9223372036854775807 {
        case 301:
            fmt.Println(p0)
        case 989:
            p0 = 2
        default:
            v4 := []int{-1, -9223372036854775808, -667}
            _ = v4
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]uint32)
    var v3 uint32 = 0
    v3 += (v3 >> ((0) & 63))
    fmt.Println(v1, v1)
    fmt.Println(v1, v3)
    v5 := GenS1{F0: 1}
    fmt.Println(v3, v1)
    var v6 uint8 = 1
    var v7 int64 = int64(v1)
    v8 := GenS1{F0: 42}
    if ((-1 <= -1) || (((true && (!((!false) || (0.1 <= -2.25)))) && false) && true)) {
        if true {
            if true {
                fmt.Println(v3, v7)
                if (!true) {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        if false {
                            v10 := GenS1{F0: 689}
                            if true {
                                v11 := GenS2{F0: 42, F1: 0}
                                _ = v11
                            }
                            v12 := GenS2{F0: 0, F1: 42}
                            _ = v10
                            _ = v12
                        }
                        v13 := GenS1{F0: -1}
                        v7 += v7
                        _ = fi9
                        _ = v13
                    }
                    switch "foo bar" {
                    case "":
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v15 := GenS1{F0: -190}
                            var v16 bool = (float64(1e10) > ((((0.1 * float64((0.1 + 3.14159))) / 1.5) * 0.1) - 0.1))
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case "\n\t":
                        var v17 int16 = 0
                        _ = v17
                    }
                    v18 := GenS1{F0: 799}
                    _ = v18
                }
                v19 := []int{}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    if false {
                        if (float32(1.5) <= 1e10) {
                            v21 := []int{(fi20 % ((fi20) | 1))}
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = v21
                        }
                        v22 := GenS3{F0: (v6 | 255)}
                        _ = v22
                    }
                    v3 -= (^(v3))
                    _ = fi20
                }
            }
            v23 := GenS3{F0: v6}
            _ = v23
        }
    } else {
        fmt.Println(v6)
        var v24 uint = 18446744073709551615
        _ = v24
    }
    if (607 < 22) {
        v3 -= 42
    } else {
        v25 := v2["hello"]
        _ = v25
    }
    fmt.Println(v6, v7, v7)
    v7 = 42
    v26 := GenS4{F0: (^(v6)), F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
