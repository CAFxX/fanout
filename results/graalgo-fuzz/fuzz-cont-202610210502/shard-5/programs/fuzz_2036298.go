package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() uint64 {
    v5 := GenS1{F0: 0, F1: 0}
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 42
        _ = fi6
        _ = v7
    }
    var v8 int8 = 0
    _ = v5
    _ = v8
    return 39
}

func main() {
    var v1 float32 = 0.0
    var v2 int8 = -1
    var v3 int16 = -32768
    var v4 int = 289
    switch 32768 {
    case -2147483648:
        switch "a" {
        case "\n\t":
            v9 := []int{127, -236, (^(v4)), (v4 / ((-1) | 1))}
            _ = v9
        case "foo bar":
            var v10 uint64 = 2
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 2; fi11++ {
                switch (v4 & v4) {
                case 9223372036854775807:
                    v12 := GenS1{F0: 733, F1: v2}
                    _ = v12
                case -2147483648:
                    v13 := GenS1{F0: 1, F1: (v2 ^ v2)}
                    _ = v13
                default:
                    if (2147483647 > 42) {
                        v14 := []int{(fi11 | (v4 << ((255) & 63))), (fi11 << ((845) & 63)), int(v2)}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        v3 -= int16(v4)
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    switch (fi11 - v4) {
                    case 128:
                        v16 := GenS1{F0: 1, F1: 1}
                        _ = v16
                    case -2147483648:
                        fmt.Println(v4, v4, v1)
                    case -27:
                        var v17 uint64 = 905
                        _ = v17
                    }
                    _ = fi15
                }
                fmt.Println(v2, v1)
                _ = fi11
            }
        default:
            v18 := []int{(v4 % ((v4) | 1)), v4, (^(v4)), v4}
            _ = v18
        }
    case 127:
        v19 := []int{-2147483648, int(v3), v4, v4}
        _ = v19
    }
    var v20 uint8 = 0
    v21 := []int{}
    fmt.Println(v4, v3, v4)
    var v22 int = int(v20)
    switch "" {
    case "":
        for fi23 := 0; fi23 < 5; fi23++ {
            v2 = (v2 % ((v2) | 1))
            _ = fi23
        }
    case "foo bar":
        v21 = append(v21, v4)
    default:
        v24 := GenS2{F0: 1, F1: 42}
        _ = v24
    }
    switch v4 {
    case 0:
        v1 = v1
    }
    fmt.Println(v2)
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v25 int64 = 1
    var v26 float64 = ((float64((v1 - v1)) / 3.14159) / float64((v1 + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
