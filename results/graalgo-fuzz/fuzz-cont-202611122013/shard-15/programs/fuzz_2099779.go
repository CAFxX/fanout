package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{42, 2147483647}
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 int8 = 1
    switch 434 {
    case -9223372036854775808:
        v5 -= (v5 | 85)
    case 798:
        if (!false) {
            switch 1 {
            case 2147483647:
                v5 -= (v5 ^ 0)
            case 128:
                v6 := GenS1{F0: 491, F1: 18446744073709551615, F2: 127}
                _ = v6
            default:
                v7 := GenS1{F0: 4294967295, F1: 42, F2: 39}
                _ = v7
            }
            v2 = append(v2, 2)
            if true {
                var v8 uint32 = 555
                fmt.Println(v5, v5)
                var v9 int = -373
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 480, F1: 42, F2: 9223372036854775807}
                v5 += v5
                _ = v10
            }
        } else {
            var v11 int8 = (v5 | v5)
            v5 = -100
            _ = v11
        }
    }
    v5 += v5
    v4[-2147483648] = 2147483647
    if ((v5 * (v5 - 42)) > v5) {
        fmt.Println(v5)
        switch "foo bar" {
        case "hello":
            if ((!false) && true) {
                switch 2147483647 {
                case 1:
                    v12 := v1["foo bar"]
                    _ = v12
                default:
                    if false {
                        v2 = append(v2, 232)
                        var v13 uint16 = 133
                        _ = v13
                    }
                }
                v2 = append(v2, 2147483647)
                v5 = 42
            } else {
                v14 := GenS1{F0: 115, F1: 2, F2: -178}
                _ = v14
            }
        case "":
            v5 += (^(v5))
        default:
            if len(v2) > 0 {
                v2[0] = 400
            }
        }
    } else {
        v5 = (v5 ^ (v5 >> ((255) & 63)))
        var v15 int16 = -1
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 698, F2: 32767}
    var v17 int8 = v5
    for fi18 := 0; fi18 < 16; fi18++ {
        v5 -= v17
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS2{F0: int8(fi18), F1: 356, F2: -85}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    delete(v1, ("a" + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
