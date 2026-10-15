package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint32, p1 int8, p2 uint32) uint64 {
    fmt.Println(p0, p2, p1)
    p2 = (p2 << ((255) & 63))
    v6 := []int{32767, 42, -2147483648, -2147483648}
    if (32767 <= 32767) {
        switch 568 {
        case -843:
            p2 += p2
        case 626:
            p1 += 6
        }
        var v7 uint64 = 368
        var v8 int8 = p1
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 float32 = 0.0
    var v5 int16 = 253
    var v9 uint16 = 929
    v10 := []int{128, -9223372036854775808, 2147483647}
    fmt.Println(v3, v4)
    if false {
        var v11 uint64 = v3
        v12 := v1[("" + "a")]
        _ = v11
        _ = v12
    } else {
        if false {
            fmt.Println(v4)
        }
        var v13 uint16 = (^(v9))
        var v14 uint64 = genFn0(0, 42, uint32(v3))
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: 0, F1: 9223372036854775807}
        _ = fi15
        _ = v16
    }
    var v17 float32 = v4
    var v18 uint8 = 42
    var v19 uint16 = 769
    var v20 int64 = int64(v9)
    var v21 float64 = 0.0
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: 250, F1: (-(v20))}
            fmt.Println(v5, v18, v21)
            var v25 int16 = v5
            switch 509 {
            case -946:
                var v26 int64 = v20
                _ = v26
            case 790:
                var v27 float32 = v17
                _ = v27
            }
            v17 = (((-2.25 * 1.5) + v17) * (-2.25 + v4))
            _ = fi23
            _ = v24
            _ = v25
        }
        v28 := GenS1{F0: 187, F1: 42}
        _ = fi22
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        if (v4 == -2.25) {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 float64 = float64(v4)
                var v32 bool = false
                _ = fi30
                _ = v31
                _ = v32
            }
            if ((fi29 % ((fi29) | 1)) < (-(fi29))) {
                var v33 uint32 = uint32(v3)
                _ = v33
            }
            v34 := GenS1{F0: 42, F1: int64(v3)}
            _ = v34
        }
        v21 += (float64(float32(v21)) * (-2.25 + ((0.0 + (1e10 - v21)) / v21)))
        _ = fi29
    }
    v35 := GenS1{F0: 990, F1: (v20 * (-(v20)))}
    var v36 bool = (!(9223372036854775807 == 4294967295))
    fmt.Println(v19, v19)
    v37 := GenS1{F0: 0, F1: (v20 / (((v20 * (v20 >> ((18446744073709551615) & 63)))) | 1))}
    v18 += (v18 - 101)
    if (float32((float64((0.0 / v4)) / (float64(v4) - (v21 / 0.1)))) >= 1.5) {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
