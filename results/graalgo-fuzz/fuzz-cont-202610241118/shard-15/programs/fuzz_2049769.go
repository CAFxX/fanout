package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0() int {
    var v6 string = "foo bar"
    switch 880 {
    case 319:
        var v7 float32 = 1.5
        _ = v7
    default:
        for fi8 := 0; fi8 < 20; fi8++ {
            if true {
                v6 = (v6 + (v6 + (v6 + v6)))
            }
            v6 = (v6 + v6)
            _ = fi8
        }
    }
    _ = v6
    return 758
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    var v3 float32 = ((v1 - (v1 + ((((v1 * (float32((-2.25 / 3.14159)) - v1)) + v1) + (v1 - v1)) * ((1.5 * 3.14159) + v1)))) * float32(-2.25))
    var v4 int = 128
    var v5 float64 = 0.1
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int64 = -482
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int = v4
            _ = fi11
            _ = v12
        }
        v10 -= v10
        _ = fi9
        _ = v10
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: 4294967295, F1: ((18446744073709551615 != 101) && (genFn0() != (v4 << ((2) & 63)))), F2: 2}
        v15 := GenS1{F0: 748, F1: (-32768 <= 836), F2: 827}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 2, F1: true, F2: 497}
    var v17 bool = true
    var v18 float32 = 1e10
    v1 = ((v1 - v3) + ((v18 + v1) * (v18 * 0.0)))
    v19 := GenS1{F0: 65535, F1: v17, F2: 65535}
    v20 := []int{v4, (v4 & v4)}
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 string = "\n\t"
        if true {
            var v23 int64 = -810
            fmt.Println(v3)
            if v17 {
                v24 := GenS2{F0: 268}
                _ = v24
            } else {
                var v25 uint = 1
                _ = v25
            }
            _ = v23
        } else {
            fmt.Println(fi21, v1, fi21)
            if len(v20) > 0 {
                v20[0] = (fi21 ^ 42)
            }
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS1{F0: 0, F1: v17, F2: 207}
            v17 = false
            var v28 uint32 = 4294967295
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi21
        _ = v22
    }
    v5 += v5
    var v29 int16 = -32768
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
