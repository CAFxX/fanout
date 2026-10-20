package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    v1 := []string{"", ""}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v4 := []uint32{2147483648}
    var v5 float64 = 0.0
    switch "foo bar" {
    case "":
        v6 := v3[("" + "")]
        _ = v6
    default:
        if (true || true) {
            fmt.Println(v5)
            var v7 int64 = 1
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v7 += (v7 % ((-1) | 1))
                    var v10 float64 = -2.25
                    _ = fi9
                    _ = v10
                }
                if true {
                    v11 := GenS1{F0: 1, F1: -976}
                    _ = v11
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    if (((((1.5 / ((float32(v5) - (0.0 - 0.1)) - 0.1)) * 3.14159) - 3.14159) >= 1.5) && false) {
                        var v13 int32 = -2147483648
                        _ = v13
                    }
                    if false {
                        var v14 uint16 = 255
                        if true {
                            switch 1 {
                            case -463:
                                fmt.Println(v14)
                            }
                        } else {
                            v7 = v7
                        }
                        _ = v14
                    }
                    _ = fi12
                }
                _ = fi8
            }
            _ = v7
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 bool = (!false)
        v17 := GenS1{F0: 701, F1: 1000}
        v5 -= v5
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 = (v5 + v5)
    v5 = 0.1
    var v18 int = 32767
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
