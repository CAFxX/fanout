package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 float32 = float32(3.14159)
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 230
    v5 := make(map[string]string)
    v2 = append(v2, -66)
    v1 = v1
    for fi6 := 0; fi6 < 13; fi6++ {
        v1 = v1
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v1 += v1
        _ = fi7
    }
    switch -295 {
    case -1:
        switch 128 {
        case -9223372036854775808:
            var v8 uint8 = (^(v4))
            _ = v8
        case -562:
            v9 := GenS1{F0: -345, F1: -641}
            _ = v9
        }
    }
    fmt.Println(v4, v1)
    v1 = float32(float64(((v1 * float32((-2.25 * 3.14159))) - ((float32(-2.25) * 1e10) - float32(float64(float32((1e10 - float64(v1)))))))))
    var v10 int16 = -32768
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            v10 = int16(fi11)
            _ = fi11
        }
        switch "" {
        case "a":
            for fi12 := 0; fi12 < 17; fi12++ {
                switch ("a" + "") {
                case "a":
                    var v13 uint64 = 2
                    _ = v13
                default:
                    if false {
                        fmt.Println(fi12, v4, v4)
                    }
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    switch "\n\t" {
                    case "\n\t":
                        if (false && (94 == 42)) {
                            v15 := GenS2{F0: 127, F1: "foo bar"}
                            v16 := GenS2{F0: -79, F1: "foo bar"}
                            var v17 uint64 = 4294967295
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        var v18 uint32 = 183
                        _ = v18
                    default:
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    }
                    fmt.Println(v10, v1)
                    var v20 int = fi14
                    _ = fi14
                    _ = v20
                }
                _ = fi12
            }
        }
        fmt.Println(v1, v10, v1)
    } else {
        v10 = 1000
    }
    fmt.Println(v4, v10)
    v1 = v1
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
