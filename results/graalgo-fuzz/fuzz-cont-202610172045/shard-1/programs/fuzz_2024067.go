package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 int8, p2 uint64) int8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        switch fi7 {
        case 32767:
            p1 -= (p1 * p1)
        default:
            var v8 int8 = (p1 + p1)
            _ = v8
        }
        v9 := []int{(fi7 | -197), 0, 2147483647}
        _ = fi7
        _ = v9
    }
    p2 = p2
    switch (("" + "\n\t") + "a") {
    case "\n\t":
        var v10 int64 = -1
        _ = v10
    case "foo bar":
        v11 := []int{-858, p0, 4294967295, 32767}
        _ = v11
    case "a":
        var v12 int = p0
        _ = v12
    default:
        p2 += (^(p2))
    }
    v13 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    v4 := []uint32{}
    var v5 float64 = (float64(1e10) + 0.1)
    var v6 bool = (true && (127 == 176))
    v6 = v6
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v2, v5, v2)
        _ = fi14
    }
    _ = len(v4)
    if (42 <= 42) {
        v6 = (!v6)
        v15 := GenS1{F0: 127}
        v16 := GenS1{F0: -707}
        _ = v15
        _ = v16
    }
    var v17 int64 = -1
    v2 = v2
    fmt.Println(v17)
    var v18 float64 = v5
    switch v2 {
    case "":
        v19 := GenS2{F0: (v17 % (((v17 >> ((694) & 63))) | 1)), F1: 323}
        _ = v19
    default:
        if v6 {
            v2 = v2
            var v20 uint64 = 518
            _ = v20
        } else {
            if v6 {
                v17 = v17
                switch "\n\t" {
                case "a":
                    var v21 int = 128
                    _ = v21
                default:
                    switch 4294967295 {
                    case -1:
                        delete(v3, 32768)
                    }
                }
            }
            var v22 float32 = (0.0 + -2.25)
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 uint16 = 65535
                var v25 int16 = int16(v17)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v22
        }
    }
    var v26 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
