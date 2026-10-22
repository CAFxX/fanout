package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 int = 1
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{}
        v4 += (-(v4))
        _ = fi5
        _ = v6
    }
    v4 -= v4
    for fi7 := 0; fi7 < 3; fi7++ {
        if (false || false) {
            v8 := GenS1{F0: -128, F1: -128, F2: 61}
            if (!true) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v4 -= 1
                    var v10 uint8 = 73
                    var v11 uint16 = 255
                    v12 := GenS1{F0: -51, F1: 105, F2: -73}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v1 -= v1
            } else {
                if false {
                    switch "hello" {
                    case "a":
                        v13 := v3[(^(fi7))]
                        _ = v13
                    case "foo bar":
                        if false {
                            v1 = v1
                            for fi14 := 0; fi14 < 15; fi14++ {
                                v1 -= v1
                                var v15 int64 = 9223372036854775807
                                _ = fi14
                                _ = v15
                            }
                        }
                    case "\n\t":
                        v1 += v1
                    default:
                        v16 := GenS1{F0: -41, F1: 0, F2: -1}
                        _ = v16
                    }
                    var v17 uint32 = 850
                    _ = v17
                } else {
                    var v18 bool = (-89 >= 0)
                    _ = v18
                }
                fmt.Println(v1)
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                if false {
                    v1 = (3.14159 * 0.0)
                } else {
                    switch v4 {
                    case 4294967295:
                        fmt.Println(v4, fi7)
                    case 1:
                        fmt.Println(v1, fi19)
                        var v20 int16 = 1000
                        v20 -= 144
                        _ = v20
                    }
                }
                var v21 int16 = 0
                if (0 != 42) {
                    v1 = 3.14159
                    v22 := GenS1{F0: -1, F1: 100, F2: 41}
                    _ = v22
                }
                var v23 uint8 = 2
                delete(v2, 1)
                _ = fi19
                _ = v21
                _ = v23
            }
            _ = v8
        }
        var v24 int8 = -128
        v25 := GenS2{F0: 32767, F1: 128}
        _ = fi7
        _ = v24
        _ = v25
    }
    var v26 uint8 = 69
    var v27 uint64 = 42
    v28 := []int{v4, (v4 - -1), v4}
    var v29 float64 = (1.5 - v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
