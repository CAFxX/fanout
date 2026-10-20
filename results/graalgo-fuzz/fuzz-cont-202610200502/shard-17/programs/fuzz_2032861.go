package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 int16, p2 int64) int64 {
    v5 := GenS1{F0: -2147483648, F1: 18446744073709551615}
    p2 = (p2 >> ((1) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -315
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    var v4 float64 = (1.5 / 3.14159)
    var v6 int = -9223372036854775808
    v7 := []int{v6}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            if (!(!(!true))) {
                if true {
                    if (false && false) {
                        _ = len(v7)
                    }
                }
            }
            var v10 uint8 = 42
            v7 = append(v7, v6)
            _ = fi9
            _ = v10
        }
        if (v4 >= ((0.1 + v4) - v4)) {
            v11 := GenS1{F0: v2, F1: 1}
            if true {
                if false {
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                }
                v12 := GenS1{F0: v2, F1: 0}
                var v13 uint8 = 1
                _ = v12
                _ = v13
            }
            _ = v11
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 int16 = 641
        v2 = (v2 / ((262) | 1))
        if len(v7) > 0 {
            v7[0] = 0
        }
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 18446744073709551615}
    v17 := GenS1{F0: (^(v2)), F1: 42}
    v18 := GenS1{F0: 11, F1: 0}
    var v19 uint32 = 494
    var v20 uint32 = 2147483648
    if (v4 == ((v4 - v4) - 1.5)) {
        v21 := GenS2{F0: 2}
        _ = v21
    } else {
        if (false || true) {
            for fi22 := 0; fi22 < 3; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    fmt.Println(v2, v19)
                    var v24 string = ("\n\t" + "foo bar")
                    v20 = (v19 & (v19 * (^(v20))))
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v2, v6)
                _ = fi22
            }
        }
    }
    fmt.Println(v19, v2, v4)
    switch "" {
    case "":
        v4 = 1.5
    case "\n\t":
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
