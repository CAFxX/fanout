package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    var v1 bool = true
    var v2 uint8 = 129
    v3 := []uint32{2, 2147483647, 2}
    var v4 int32 = 2147483647
    v5 := []int{1, 610, 42}
    v6 := []uint32{}
    for fi7 := 0; fi7 < 9; fi7++ {
        v4 = v4
        _ = fi7
    }
    v4 = v4
    var v8 float32 = ((1e10 + (0.0 - float32((0.1 * (3.14159 * (float64((0.0 / 1.0)) * 1.5)))))) - 3.14159)
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    var v11 uint = 255
    switch "a" {
    case "foo bar":
        if (!false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case "hello":
        v1 = ((^(v11)) < v11)
    case "\n\t":
        if (false && v1) {
            if v1 {
                switch 42 {
                case 1:
                    v1 = ((v4 ^ (v4 & (v4 / (((^(v4))) | 1)))) <= -367)
                case 4294967295:
                    if ((!((v4 ^ v4) > v4)) || ((!(!true)) && (!(!(1 < 512))))) {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            if true {
                                var v13 uint16 = 255
                                _ = v13
                            }
                            _ = fi12
                        }
                    } else {
                        v14 := GenS1{F0: (("a" + "a") + ""), F1: 32767}
                        v2 += v2
                        _ = v14
                    }
                case 484:
                    fmt.Println(v11)
                }
            }
        }
    default:
        v2 = v2
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
