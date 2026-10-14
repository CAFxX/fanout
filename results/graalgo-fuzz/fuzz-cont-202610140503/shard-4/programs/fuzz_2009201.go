package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0(p0 uint64, p1 int8) string {
    p1 = p1
    var v4 int32 = -1
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 1000, F1: p0, F2: v4}
            v4 = (v4 << ((0) & 63))
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 / ((42) | 1)), fi7, 128}
            p1 = (-(p1))
            _ = fi7
            _ = v8
        }
        var v9 int32 = v4
        _ = v9
    } else {
        v10 := []int{1, 9223372036854775807, 127, 728}
        var v11 bool = ((1e10 >= 1.5) || false)
        var v12 float32 = ((0.0 + -2.25) + (3.14159 + 3.14159))
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float64 = (((3.14159 * 1e10) / 0.1) - (0.0 + (1.5 + 3.14159)))
    _ = p0
    _ = p1
    _ = v13
    _ = v4
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint16 = 606
    var v14 float32 = 3.14159
    v3 -= (v3 | v3)
    v3 = 1
    if true {
        var v15 uint8 = 42
        var v16 uint16 = v3
        _ = v15
        _ = v16
    }
    var v17 int64 = 0
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int = (fi18 ^ 127)
        v19 = 9223372036854775807
        if (true && true) {
            v3 -= 255
        }
        _ = fi18
        _ = v19
    }
    v17 += (^(v17))
    fmt.Println(v17, v14, v3)
    v20 := []int{-9223372036854775808, 128, -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
