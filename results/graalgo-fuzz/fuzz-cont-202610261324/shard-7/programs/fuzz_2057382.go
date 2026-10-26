package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    var v3 int = -816
    var v4 uint8 = 127
    var v5 uint32 = 132
    if true {
        v6 := GenS1{F0: 42, F1: (v3 | v3)}
        _ = v6
    }
    if (436 > 269) {
        v7 := GenS1{F0: 18446744073709551615, F1: (v3 & (v3 % ((v3) | 1)))}
        switch -973 {
        case 42:
            switch (v3 % ((42) | 1)) {
            case 121:
                v3 = 1
            case 42:
                v8 := v2[v3]
                _ = v8
            case 99:
                var v9 uint32 = v5
                _ = v9
            default:
                var v10 int32 = -968
                _ = v10
            }
        case 128:
            if (739 < 1) {
                v3 = (^(v3))
            }
        case 838:
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v3 = -9223372036854775808
                    v13 := []int{649, fi11, 127, fi11}
                    _ = fi12
                    _ = v13
                }
                var v14 int = (v3 * v3)
                _ = fi11
                _ = v14
            }
        default:
            fmt.Println(v1)
        }
        _ = v7
    }
    fmt.Println(v5, v4, v1)
    v5 -= v5
    for fi15 := 0; fi15 < 5; fi15++ {
        v4 = 157
        var v16 bool = ((0.1 / 1.0) < 3.14159)
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 2, F1: v3}
    v5 += v5
    var v18 int8 = 127
    switch "foo bar" {
    case "a":
        v19 := GenS2{F0: 690, F1: -1, F2: (v4 / ((v4) | 1))}
        _ = v19
    default:
        fmt.Println(v3)
    }
    var v20 int32 = 0
    v3 -= v3
    if true {
        for fi21 := 0; fi21 < 13; fi21++ {
            _ = len(v2)
            _ = fi21
        }
        switch (^(v3)) {
        case 4294967295:
            var v22 uint16 = 779
            _ = v22
        case -608:
            for fi23 := 0; fi23 < 10; fi23++ {
                _ = len(v2)
                _ = fi23
            }
        case -1:
            var v24 int64 = -1
            _ = v24
        default:
            v25 := []int{v3, v3}
            _ = v25
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 float64 = float64(float32(0.0))
            _ = fi26
            _ = v27
        }
    } else {
        v1 += v1
        v20 -= 1
    }
    var v28 string = "a"
    var v29 uint16 = 0
    for fi30 := 0; fi30 < 13; fi30++ {
        v5 = (v5 * v5)
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
