package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
func genFn0(p0 int) uint8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = 42
        if true {
            if true {
                if false {
                    fmt.Println(p0)
                }
                v8 := []int{(fi6 / (((p0 + fi6)) | 1)), 127, 9223372036854775807}
                _ = v8
            } else {
                var v9 int16 = 0
                var v10 uint64 = 18446744073709551615
                _ = v10
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    if false {
        var v11 bool = true
        v12 := []int{-9223372036854775808, 42}
        _ = v11
        _ = v12
    } else {
        var v13 int16 = 32767
        _ = v13
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 float32 = (((-2.25 * 1.5) + 1.5) + 0.1)
    v5 := []int{}
    v4 -= float32((((1e10 * 0.0) + 1e10) / (((1.5 / 1.0) + (float64(v4) * float64((((v4 / (1e10 / (v4 * v4))) * v4) * float32(1e10))))) + ((3.14159 - (((1.5 + (1.5 * float64(v4))) * (0.0 - ((1e10 * 1e10) - 0.1))) * 1e10)) - -2.25))))
    v4 -= ((((((((0.0 / float32((1e10 / (1.5 - ((1.5 + ((0.0 / 1e10) + 1.5)) / 1e10))))) + (-2.25 / (0.1 + (float32(((1.5 - (0.1 * (float64((v4 * v4)) / 3.14159))) / 1.5)) + v4)))) / v4) + 0.1) - v4) * v4) / float32(1.5)) + v4)
    v4 -= v4
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float64 = (-2.25 - 1e10)
        v16 := v1[2147483647]
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint16 = 32767
    switch -745 {
    case 412:
        switch "\n\t" {
        case "":
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS1{F0: 0, F1: 2, F2: int32(v17)}
                if len(v5) > 0 {
                    v5[0] = (fi18 >> ((624) & 63))
                }
                v3[((("" + "a") + "") + "")] = 240
                _ = fi18
                _ = v19
            }
        case "foo bar":
            fmt.Println(v17, v4)
        }
    case 364:
        v3[("a" + "foo bar")] = 4294967295
    }
    fmt.Println(v17)
    v4 = v4
    _ = len(v3)
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
