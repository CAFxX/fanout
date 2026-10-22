package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint8, p1 uint64, p2 int16) uint64 {
    if false {
        p1 += p1
        p1 = (p1 & (^(p1)))
    }
    var v6 bool = true
    v7 := []int{2147483647}
    p1 = 1
    if len(v7) > 0 {
        v7[0] = -378
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    var v3 uint = 973
    v4 := []string{"\n\t"}
    var v5 int8 = -7
    switch "\n\t" {
    case "a":
        v8 := GenS1{F0: v5}
        _ = v8
    case "":
        var v9 string = ("\n\t" + "a")
        _ = v9
    default:
        v1 -= v3
    }
    switch 0 {
    case 1:
        v10 := GenS1{F0: 42}
        _ = v10
    case 236:
        switch ("a" + ("hello" + "")) {
        case "hello":
            v5 = (v5 ^ (^(v5)))
        case "\n\t":
            v5 = (v5 & (v5 ^ (v5 | 0)))
        default:
            v5 -= 127
        }
    case -158:
        for fi11 := 0; fi11 < 6; fi11++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v5)
            case "foo bar":
                switch ("a" + ("a" + "a")) {
                case "foo bar":
                    var v12 int16 = -32768
                    _ = v12
                default:
                    var v13 int = 606
                    _ = v13
                }
            case "hello":
                if false {
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v14 int64 = 1
                        _ = v14
                    case "":
                        v5 += v5
                    default:
                        if (3.14159 != float32(float64((float32((float64(1.5) + float64((((3.14159 / (1.5 + 0.1)) + 1e10) * ((0.0 / 1.5) - (1.5 + 1e10)))))) * 1.5)))) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    v1 = (v1 - (v1 | v3))
                    v15 := v2[(fi11 / ((fi11) | 1))]
                    _ = v15
                } else {
                    v1 -= v3
                }
            default:
                v3 += (v3 << ((v3) & 63))
            }
            fmt.Println(v5, v1)
            _ = fi11
        }
    default:
        v5 -= v5
    }
    v16 := GenS1{F0: v5}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS1{F0: (v5 - -1)}
    v5 = (v5 ^ v5)
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 bool = true
        v5 -= 100
        _ = fi18
        _ = v19
    }
    v3 += (v1 / (((v1 | v1)) | 1))
    v20 := GenS2{F0: -9223372036854775808, F1: v5}
    delete(v2, 9223372036854775807)
    v1 -= v3
    v21 := v2[-1]
    v21 -= v21
    v1 = (v3 - 0)
    switch "foo bar" {
    case "foo bar":
        v3 = (v3 - uint(v21))
    case "":
        fmt.Println(v5, v1)
    case "\n\t":
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v1)
                _ = fi23
            }
            v24 := GenS2{F0: 0, F1: (v5 | v5)}
            v1 -= 4294967295
            _ = fi22
            _ = v24
        }
    default:
        var v25 uint8 = 2
        _ = v25
    }
    _ = len(v2)
    var v26 int64 = -9223372036854775808
    if len(v4) > 0 {
        v4[0] = (("hello" + ("foo bar" + ("a" + ""))) + "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
