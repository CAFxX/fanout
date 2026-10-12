package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []string{"foo bar", ""}
    var v2 float32 = -2.25
    var v3 uint64 = 894
    var v4 int8 = -128
    var v5 string = ("hello" + (("\n\t" + ("\n\t" + ("" + ("a" + "foo bar")))) + "a"))
    var v6 int64 = -324
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (v3 / ((v3) | 1))
    _ = len(v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int = (^(fi7))
        v8 -= v8
        _ = fi7
        _ = v8
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        switch (v5 + v5) {
        case "a":
            v9 := GenS1{F0: true, F1: 0}
            _ = v9
        case "\n\t":
            fmt.Println(v2, v2, v4)
        default:
            v5 = (v5 + (v5 + v5))
        }
        switch -9223372036854775808 {
        case 100:
            var v10 int8 = 0
            _ = v10
        default:
            var v11 bool = ((((v2 / (v2 - 1e10)) - ((((v2 / 1.5) + float32((3.14159 - 1.5))) - ((0.1 - v2) + v2)) * (1.5 * 1.5))) - (((v2 - ((1.5 - 3.14159) / (0.0 / (float32((-2.25 / 0.1)) + v2)))) / v2) - 1.5)) != 1e10)
            _ = v11
        }
    }
    if true {
        if true {
            for fi12 := 0; fi12 < 16; fi12++ {
                switch 128 {
                case 743:
                    if len(v1) > 0 {
                        v1[0] = v5
                    }
                case 1:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        if true {
                            v1 = append(v1, v5)
                        }
                        v14 := GenS1{F0: true, F1: 396}
                        var v15 uint = 18446744073709551615
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    v1 = append(v1, v5)
                }
                _ = fi12
            }
        }
    }
    var v16 float32 = (0.0 + float32(0.0))
    switch 32768 {
    case 828:
        v6 = (-(v6))
    }
    var v17 uint8 = 0
    var v18 bool = true
    v6 += v6
    v16 = ((1.5 + v2) + v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
