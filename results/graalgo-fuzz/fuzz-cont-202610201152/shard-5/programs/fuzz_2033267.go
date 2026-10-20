package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]bool)
    var v3 uint32 = 1
    var v4 int8 = 100
    v5 := []uint32{}
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 0, F1: 127, F2: 42}
        v1 += v1
        _ = fi7
        _ = v8
    }
    v2[128] = (-544 > 128)
    fmt.Println(v6, v4)
    var v9 int16 = -32768
    v10 := GenS1{F0: 18446744073709551615, F1: -128, F2: 2147483647}
    if (((3.14159 + 1.5) - (1e10 + -2.25)) > -2.25) {
        var v11 uint64 = 0
        _ = v11
    } else {
        delete(v2, 127)
        v6 += v6
        if ("foo bar" < ("foo bar" + ((((("\n\t" + "hello") + "a") + "a") + "a") + "\n\t"))) {
            if false {
                switch 4294967295 {
                case -1:
                    _ = len(v5)
                case 1:
                    fmt.Println(v1, v3, v1)
                default:
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v13 := GenS1{F0: 2, F1: (v4 / (((v4 / ((10) | 1))) | 1)), F2: (v3 & v3)}
                        fmt.Println(v1, v6)
                        _ = fi12
                        _ = v13
                    }
                }
            } else {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v3 = 42
                    _ = fi14
                }
                v15 := GenS2{F0: (v3 - (v3 * (v3 & (v3 | (v3 + (v3 ^ v3)))))), F1: (v1 & v1)}
                _ = v15
            }
        } else {
            switch 128 {
            case -615:
                var v16 int16 = (^(v9))
                _ = v16
            case 492:
                var v17 uint64 = v1
                _ = v17
            }
            _ = len(v5)
        }
    }
    fmt.Println(v1, v9, v9)
    v1 = (^(v1))
    var v18 int8 = -128
    fmt.Println(v4, v6, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
