package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 3.14159
    var v3 uint8 = 255
    var v4 int16 = -32768
    v5 := []int{}
    fmt.Println(v4)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 bool = false
        fmt.Println(v2, v3)
        _ = fi6
        _ = v7
    }
    switch 203 {
    case 32767:
        var v8 uint32 = 254
        _ = v8
    default:
        if true {
            if ((true && (0.0 < 1.5)) || ((!false) && (float32(1e10) >= -2.25))) {
                switch "hello" {
                case "foo bar":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        switch fi9 {
                        case -9223372036854775808:
                            v10 := GenS1{F0: 18446744073709551615, F1: fi9}
                            _ = v10
                        case 0:
                            var v11 int32 = -540
                            _ = v11
                        }
                        v4 += v4
                        _ = fi9
                    }
                case "\n\t":
                    v2 = ((((v2 - ((((v2 - float64(((-2.25 * 0.1) + (1e10 + float32((float64((1e10 - 1.5)) + -2.25)))))) * v2) - (1e10 - ((-2.25 + float64(float32(v2))) / 1e10))) * (v2 / ((v2 + ((v2 * -2.25) / v2)) - (1e10 - 1.5))))) / v2) * v2) / ((1e10 / v2) * (v2 + v2)))
                }
            }
            v12 := GenS1{F0: uint64(v3), F1: 758}
            var v13 uint8 = v3
            _ = v12
            _ = v13
        }
    }
    if (!false) {
        var v14 int16 = -32768
        _ = v14
    }
    var v15 int64 = -1
    _ = len(v1)
    v15 = (-(v15))
    var v16 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
