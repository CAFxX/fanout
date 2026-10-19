package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }

func main() {
    var v1 uint16 = 32767
    v2 := []int{}
    var v3 float32 = -2.25
    v4 := make(map[uint32]bool)
    var v5 float64 = (1.5 / -2.25)
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 4; fi7++ {
        if (true || (!(!true))) {
            var v8 int8 = 0
            v1 -= 32767
            switch fi7 {
            case -130:
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 uint32 = 2
                    _ = len(v4)
                    var v11 float32 = (float32(1e10) - v3)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            _ = v8
        }
        v3 -= 1e10
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 float32 = (v3 + v3)
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    if false {
        v14 := v4[0]
        v15 := GenS1{F0: (-68 < 42), F1: -264}
        v1 = (v1 | (^(v1)))
        _ = v14
        _ = v15
    }
    v6 += 182
    fmt.Println(v1)
    v5 = float64(v3)
    if (!false) {
        var v16 uint16 = (v1 ^ v1)
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                delete(v4, 2147483648)
                v19 := GenS2{F0: 127, F1: (v6 + (v6 * v6)), F2: -2147483648}
                var v20 int64 = 9223372036854775807
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v3, fi17, v6)
            var v21 uint8 = 1
            _ = fi17
            _ = v21
        }
        _ = v16
    } else {
        if (!false) {
            switch "" {
            case "":
                var v22 uint8 = v6
                _ = v22
            default:
                if (-29 == 100) {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        if len(v2) > 0 {
                            v2[0] = fi23
                        }
                        v1 += (v1 | v1)
                        v5 += (v5 * (v5 / v5))
                        _ = fi23
                    }
                    v24 := GenS1{F0: true, F1: 42}
                    _ = v24
                } else {
                    var v25 uint32 = 265
                    v26 := v4[2147483648]
                    fmt.Println(v25, v26, v3)
                    _ = v25
                    _ = v26
                }
            }
        }
        var v27 float32 = 1e10
        _ = v27
    }
    fmt.Println(v3, v1)
    v3 = v3
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int32 = -1
        _ = fi28
        _ = v29
    }
    if (1 == 42) {
        var v30 int32 = int32(v1)
        _ = v30
    }
    v3 += ((((float32(v5) - v3) * v3) - 0.0) / v3)
    fmt.Println(v1, v3)
    var v31 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
