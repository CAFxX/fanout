package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (float64(float32(-2.25)) / 1.5)
    var v3 uint32 = 2147483647
    var v4 bool = false
    var v5 int8 = 1
    v6 := []int{42, 1, -1}
    if v4 {
        if true {
            switch -698 {
            case -200:
                if (!v4) {
                    fmt.Println(v2, v4, v5)
                    v5 = (v5 / ((v5) | 1))
                } else {
                    v7 := GenS1{F0: 2147483647, F1: v5, F2: 1000}
                    _ = v7
                }
            case -695:
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 int8 = v5
                    _ = fi8
                    _ = v9
                }
            case 1:
                var v10 uint64 = 4294967295
                _ = v10
            default:
                v1[v3] = ""
            }
            v5 = (-(v5))
            _ = len(v6)
        }
        var v11 uint = 65535
        _ = v11
    }
    v3 += 308
    v2 = v2
    var v12 int32 = -2147483648
    var v13 int16 = 922
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 float64 = (float64(float32(v2)) * v2)
        _ = fi14
        _ = v15
    }
    v2 += ((0.1 + ((1.5 - v2) + float64(((1.5 * float32((3.14159 - v2))) * 3.14159)))) + float64(0.0))
    if v4 {
        fmt.Println(v5, v12)
    }
    var v16 float64 = (3.14159 - v2)
    switch 9223372036854775807 {
    case -841:
        for fi17 := 0; fi17 < 19; fi17++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v18 int16 = (-(v13))
            _ = fi17
            _ = v18
        }
    case 776:
        var v19 float64 = ((v2 - v16) + (((v2 * v16) - v2) * v2))
        _ = v19
    }
    var v20 uint = 667
    switch "foo bar" {
    case "hello":
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int64 = 31
            _ = fi21
            _ = v22
        }
    case "\n\t":
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint16 = 2
            var v25 uint = v20
            _ = fi23
            _ = v24
            _ = v25
        }
    default:
        v2 = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
