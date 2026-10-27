package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64) int8 {
    switch 1 {
    case -788:
        v4 := GenS1{F0: true, F1: p0}
        _ = v4
    }
    switch -1 {
    case 128:
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 = 1
            for fi6 := 0; fi6 < 8; fi6++ {
                p0 -= 18446744073709551615
                var v7 uint = 4294967295
                var v8 bool = true
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    case -2147483648:
        var v9 float32 = float32(1e10)
        _ = v9
    case 4294967295:
        p0 -= 4294967295
    default:
        p0 = (p0 ^ (p0 + (p0 ^ (p0 ^ 42))))
    }
    v10 := []int{127}
    p0 += (p0 / (((p0 - 4294967295)) | 1))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = v10
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483648, 4294967295, 1, 42}
    v11 := GenS1{F0: (!(-775 == -159)), F1: 20}
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint = 1
        v13 -= v13
        v13 = (v13 & (^(v13)))
        _ = fi12
        _ = v13
    }
    var v14 bool = (-224 <= -9223372036854775808)
    var v15 string = "foo bar"
    v14 = (!(v14 && (0.1 <= (float64(1.5) + 1.5))))
    v14 = false
    var v16 bool = v14
    switch (v15 + (v15 + v15)) {
    case "foo bar":
        v17 := GenS1{F0: v14, F1: 632}
        _ = v17
    default:
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: (!(1.5 != 0.0)), F1: 2}
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
