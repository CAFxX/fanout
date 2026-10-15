package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    v4 := GenS1{F0: -6, F1: -276, F2: 2}
    v3 += (((0.0 - v3) / v3) - float32(((1.5 - ((0.0 / -2.25) + 0.0)) + ((-2.25 - (0.0 + 0.0)) / 1.5))))
    if true {
        v5 := GenS1{F0: -128, F1: 1000, F2: 73}
        v1 -= 2147483647
        if true {
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 int64 = -1
                var v8 uint16 = 722
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                v3 = (v3 * (v3 - v3))
            case "hello":
                v9 := []int{127}
                _ = v9
            default:
                v1 -= (v1 - 4294967295)
            }
            if (true || (0 == -1)) {
                v10 := []int{-764, -665}
                _ = len(v10)
                if (!(!(false && (0 > 32767)))) {
                    var v11 int8 = 15
                    _ = v11
                } else {
                    if (false || (1 <= 42)) {
                        for fi12 := 0; fi12 < 2; fi12++ {
                            var v13 int = (^(fi12))
                            v1 -= (^(v1))
                            _ = fi12
                            _ = v13
                        }
                        v3 = (v3 + ((1.5 * v3) * (float32(1e10) / v3)))
                    } else {
                        v10 = append(v10, 556)
                        v14 := GenS1{F0: 100, F1: -356, F2: 65535}
                        _ = v14
                    }
                    v15 := GenS1{F0: 0, F1: 0, F2: 0}
                    _ = v15
                }
                _ = v10
            } else {
                v16 := GenS1{F0: -125, F1: -32768, F2: 65535}
                v17 := []int{}
                switch "\n\t" {
                case "":
                    v3 -= 1.5
                case "hello":
                    var v18 float64 = 3.14159
                    _ = v18
                case "a":
                    v1 += v1
                }
                _ = v16
                _ = v17
            }
            v3 -= (v3 - float32((-2.25 / (1e10 * float64(v3)))))
        }
        _ = v5
    }
    v19 := GenS1{F0: 0, F1: 1000, F2: 660}
    v20 := GenS1{F0: 11, F1: 288, F2: 65535}
    if (65535 == 2) {
        v3 -= 1.5
    }
    v21 := []int{42, 566}
    v22 := GenS2{F0: false, F1: 291, F2: -466}
    var v23 uint16 = 567
    switch ("a" + "\n\t") {
    case "":
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(fi24, fi24)
            var v25 int64 = 42
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS2{F0: false, F1: -589, F2: 0}
    for fi27 := 0; fi27 < 3; fi27++ {
        _ = len(v2)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
