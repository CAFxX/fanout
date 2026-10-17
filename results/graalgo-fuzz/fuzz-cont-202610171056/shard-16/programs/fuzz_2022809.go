package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32, p1 uint) int {
    if ((1000 < 0) && false) {
        v4 := GenS1{F0: true}
        _ = v4
    } else {
        v5 := []int{32767, 127}
        p0 = p0
        var v6 int32 = -1
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        if (!(1 > 1)) {
            for fi8 := 0; fi8 < 17; fi8++ {
                p0 = p0
                var v9 int16 = -32768
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: true}
            _ = v10
        }
        var v11 string = ""
        v12 := GenS2{F0: 254, F1: 398}
        _ = fi7
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: -513}
    p1 = p1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 65
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    v14 := v3["hello"]
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float32 = 1.5
        _ = fi15
        _ = v16
    }
    v2 -= v2
    v17 := GenS3{F0: 1, F1: v1}
    var v18 int = 128
    v2 = -707
    v18 = v18
    if (float32((0.0 * 0.0)) != -2.25) {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS3{F0: 42, F1: (^(v1))}
            v1 = 42
            v21 := []int{(v18 / ((v18) | 1)), (v18 % ((-566) | 1))}
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        v2 = (v2 % ((0) | 1))
        switch "" {
        case "hello":
            v18 = v18
        case "":
            for fi22 := 0; fi22 < 1; fi22++ {
                delete(v3, "\n\t")
                v23 := GenS3{F0: 18446744073709551615, F1: v1}
                _ = fi22
                _ = v23
            }
        default:
            fmt.Println(v18)
        }
        v24 := []int{}
        _ = v24
    }
    var v25 uint64 = 0
    if (!((18446744073709551615 != 42) || false)) {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
