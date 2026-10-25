package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []int{2147483647, 860, -9223372036854775808}
    v2 := []string{"", "hello"}
    v3 := make(map[uint32]string)
    var v4 float32 = 1e10
    var v5 int16 = -1
    v4 = v4
    v5 += v5
    if (2 <= 2) {
        fmt.Println(v5, v5)
    }
    switch ("hello" + "") {
    case "":
        var v6 uint32 = 587
        _ = v6
    case "a":
        var v7 bool = (241 == 1)
        _ = v7
    case "\n\t":
        var v8 string = "\n\t"
        _ = v8
    default:
        if true {
            for fi9 := 0; fi9 < 15; fi9++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi9
            }
            v4 -= ((1.5 * -2.25) * ((3.14159 / 1.5) * v4))
            if (false || (false && (42 > 2))) {
                switch 9223372036854775807 {
                case -9223372036854775808:
                    if (-71 != 100) {
                        v1 = append(v1, -1)
                        v4 -= 0.0
                    } else {
                        var v10 uint8 = 34
                        fmt.Println(v10, v4)
                        _ = v10
                    }
                case 127:
                    var v11 int32 = 2147483647
                    _ = v11
                case -990:
                    var v12 int = 4294967295
                    _ = v12
                }
            }
        } else {
            v4 -= v4
        }
    }
    v13 := GenS1{F0: 255, F1: (^(v5))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += v4
    var v14 uint32 = 1
    fmt.Println(v4, v14)
    if ((true && (float64(1e10) >= -2.25)) || (true && true)) {
        v4 = v4
        var v15 uint64 = 0
        for fi16 := 0; fi16 < 14; fi16++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v5 = v5
            _ = fi16
        }
        _ = v15
    } else {
        v1 = append(v1, 42)
        v14 -= v14
        fmt.Println(v5, v14)
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
