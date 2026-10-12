package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
func genFn0() uint {
    var v4 bool = true
    var v5 int16 = 1000
    v5 += (v5 >> ((0) & 63))
    _ = v4
    _ = v5
    return 992
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "hello"
    var v3 int64 = 891
    v3 = -9223372036854775808
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = -485
            v9 := []int{-2147483648, (fi6 ^ 32768)}
            var v10 uint = 0
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    var v11 int64 = (v3 & (v3 * v3))
    v12 := []int{32768, 9223372036854775807}
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v11, v2)
        v12 = append(v12, (fi13 % ((32767) | 1)))
        _ = fi13
    }
    _ = len(v12)
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS1{F0: v2, F1: 42, F2: v2}
        var v17 int8 = -67
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int8 = 100
        var v20 float64 = ((1.5 * 1e10) + ((1.5 / 3.14159) - -2.25))
        _ = fi18
        _ = v19
        _ = v20
    }
    if len(v12) > 0 {
        v12[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
