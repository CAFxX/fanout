package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    v3 := []int{-495, 206, -213, 271}
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v4, v1, v4)
        if (128 >= 255) {
            v4 = v4
            switch "foo bar" {
            case "a":
                for fi6 := 0; fi6 < 10; fi6++ {
                    v1 = (!(v1 || ((fi5 * (fi6 << ((18446744073709551615) & 63))) != 128)))
                    v7 := v2[(("" + "foo bar") + "a")]
                    _ = fi6
                    _ = v7
                }
            case "\n\t":
                v3 = append(v3, (fi5 & fi5))
            default:
                v8 := GenS1{F0: 2147483648, F1: v1}
                _ = v8
            }
        }
        _ = fi5
    }
    v9 := GenS1{F0: 0, F1: (!v1)}
    var v10 uint16 = 1
    v11 := GenS1{F0: 515, F1: v1}
    var v12 int8 = 42
    v10 -= v10
    fmt.Println(v1, v4)
    switch 32767 {
    case 1:
        var v13 uint8 = 42
        _ = v13
    case -674:
        v10 += 32767
    default:
        for fi14 := 0; fi14 < 2; fi14++ {
            if v1 {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if len(v3) > 0 {
                        v3[0] = fi14
                    }
                    var v16 int8 = (v12 << ((2) & 63))
                    _ = fi15
                    _ = v16
                }
            }
            var v17 uint64 = 18446744073709551615
            _ = fi14
            _ = v17
        }
    }
    v4 += v4
    fmt.Println(v4)
    var v18 int64 = 9223372036854775807
    v1 = (42 > (v12 + (v12 + (v12 / (((v12 & v12)) | 1)))))
    fmt.Println(v12)
    var v19 int8 = (v12 - v12)
    v1 = (!v1)
    v19 -= v19
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 773, F1: (v4 > v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
