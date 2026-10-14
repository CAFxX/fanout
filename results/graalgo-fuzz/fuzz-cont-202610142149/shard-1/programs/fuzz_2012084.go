package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int16 = 42
    var v2 int8 = 42
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint8 = 42
        _ = fi4
        _ = v5
    }
    v6 := v3[""]
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{32768, -1}
        v8 = append(v8, fi7)
        _ = fi7
        _ = v8
    }
    _ = len(v3)
    fmt.Println(v1, v6, v6)
    var v9 int8 = int8(v1)
    var v10 float64 = 3.14159
    if (false && true) {
        if false {
            for fi11 := 0; fi11 < 4; fi11++ {
                switch "" {
                case "\n\t":
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 int32 = 0
                        var v14 int8 = 127
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "":
                    var v15 int32 = 1
                    _ = v15
                case "a":
                    v1 = (v1 + (v1 | (v1 ^ v1)))
                default:
                    if false {
                        v1 = (v1 / ((v1) | 1))
                    } else {
                        v10 += 3.14159
                        v10 -= v10
                        var v16 uint8 = 91
                        _ = v16
                    }
                }
                _ = fi11
            }
        }
        v3[((("foo bar" + (v6 + "\n\t")) + v6) + v6)] = v6
        fmt.Println(v2)
    }
    if (!((606 == 767) && false)) {
        v17 := GenS1{F0: false, F1: 0}
        v10 -= ((v10 - 1.5) + (((v10 + (float64((-2.25 - (1e10 - 3.14159))) / 1e10)) + (3.14159 - v10)) + (v10 * 3.14159)))
        _ = v17
    } else {
        var v18 int = 4294967295
        v2 = (v2 % ((v2) | 1))
        fmt.Println(v9)
        _ = v18
    }
    v19 := []int{-1, -2147483648, 32767, 9223372036854775807}
    v20 := v3[v6]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v9
}
