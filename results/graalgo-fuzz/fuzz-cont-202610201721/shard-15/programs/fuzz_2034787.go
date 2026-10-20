package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
func genFn0() uint8 {
    if false {
        v4 := []int{}
        var v5 uint = 65535
        v5 = v5
        _ = v4
        _ = v5
    }
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 92, F1: 2147483647}
    _ = v6
    _ = v7
    return 214
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v8 := []int{42, -73}
    var v9 float64 = 0.1
    var v10 int = 32767
    v11 := GenS2{F0: 65535, F1: -1}
    if false {
        var v12 uint16 = 142
        _ = v12
    } else {
        v10 = v10
        var v13 uint64 = 0
        _ = v13
    }
    v10 -= (-(v10))
    v10 += v10
    var v14 uint64 = 18446744073709551615
    if false {
        if (2 > 308) {
            v9 -= v9
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint8 = 127
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS3{F0: "", F1: 18, F2: int8(v10)}
        for fi18 := 0; fi18 < 4; fi18++ {
            if (617 != 2) {
                v10 -= (fi18 | fi18)
                if false {
                    var v19 int64 = 41
                    _ = v19
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := GenS1{F0: -128, F1: -2147483648}
                    _ = fi20
                    _ = v21
                }
            } else {
                v14 -= 4294967295
                if true {
                    fmt.Println(v9, v14)
                    for fi22 := 0; fi22 < 11; fi22++ {
                        if false {
                            var v23 float64 = (v9 - ((float64(3.14159) / v9) / float64(float32((v9 / ((v9 - (float64(((((((float32((3.14159 + v9)) / 0.1) / 1e10) * (0.0 / 1.0)) / 1e10) + ((0.1 - 0.1) - (((1e10 - -2.25) + 0.1) + 1e10))) + 0.1)) + 0.1)) - 1e10))))))
                            _ = v23
                        }
                        v24 := GenS2{F0: 255, F1: -556}
                        v10 -= v10
                        _ = fi22
                        _ = v24
                    }
                    var v25 int = v10
                    _ = v25
                }
                v14 -= 1
            }
            v10 = fi18
            _ = len(v3)
            _ = fi18
        }
        _ = v17
    }
    switch "a" {
    case "hello":
        v14 += v14
    case "\n\t":
        v26 := GenS2{F0: 0, F1: 522}
        _ = v26
    default:
        var v27 int = v10
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
