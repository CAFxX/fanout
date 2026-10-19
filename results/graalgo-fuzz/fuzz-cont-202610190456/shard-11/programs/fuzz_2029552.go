package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = ((-2.25 - -2.25) + 1e10)
    v2 := []int{202}
    var v3 uint16 = 32767
    fmt.Println(v1, v1)
    var v4 int8 = 42
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v4, v4, v3)
    case "":
        v5 := GenS1{F0: (v3 & (v3 | v3)), F1: (-(v4)), F2: 0}
        _ = v5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint = 2
        if ((v4 - v4) > v4) {
            var v8 int16 = -885
            _ = v8
        } else {
            var v9 int16 = 12
            for fi10 := 0; fi10 < 2; fi10++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 += (v4 & (v4 ^ v4))
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v2 = append(v2, fi6)
                }
                _ = fi10
            }
            switch ("hello" + "hello") {
            case "\n\t":
                var v11 uint64 = 1
                _ = v11
            case "":
                _ = len(v2)
            }
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v4 += v4
    v12 := GenS1{F0: (^(v3)), F1: v4, F2: 128}
    v1 += ((v1 / v1) * ((float32(((0.1 * ((-2.25 - (-2.25 / 1.0)) * (0.0 / 1e10))) + -2.25)) - v1) * (((((0.0 - v1) - v1) - v1) * 3.14159) - float32(1e10))))
    v4 -= (^(v4))
    _ = len(v2)
    var v13 int32 = 1
    var v14 bool = (2 <= (v3 / (((^(v3))) | 1)))
    v15 := GenS1{F0: (^(v3)), F1: (v4 ^ 127), F2: 1}
    if (!v14) {
        var v16 float64 = 1.5
        for fi17 := 0; fi17 < 15; fi17++ {
            if v14 {
                v18 := GenS2{F0: 0}
                _ = v18
            }
            fmt.Println(v16, v13)
            _ = fi17
        }
        v2 = append(v2, 578)
        _ = v16
    }
    var v19 int16 = 42
    v3 = v3
    v13 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
