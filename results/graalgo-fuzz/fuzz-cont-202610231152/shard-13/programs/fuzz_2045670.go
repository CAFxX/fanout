package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 int = -379
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 int64 = 34
    var v5 uint64 = (v3 * 42)
    var v6 int16 = -597
    v7 := GenS1{F0: 1}
    v6 += v6
    v3 -= v5
    var v8 int8 = 1
    var v9 int16 = v6
    if (-1 > 348) {
        v10 := GenS1{F0: 693}
        v6 += (v6 - (v6 / ((v6) | 1)))
        for fi11 := 0; fi11 < 10; fi11++ {
            if true {
                v8 += (-(v8))
                _ = len(v2)
            } else {
                var v12 int16 = -174
                var v13 uint16 = 580
                if (v4 == (v4 << ((724) & 63))) {
                    var v14 float64 = ((1e10 - -2.25) / ((0.1 * ((0.1 - (1.5 - (0.0 - 1.5))) + ((((-2.25 / (1e10 * (0.0 - 0.1))) / 1.5) + ((3.14159 / (-2.25 + 0.0)) * 3.14159)) * ((float64(3.14159) / (-2.25 / 3.14159)) - float64(3.14159))))) * -2.25))
                    v15 := GenS1{F0: v13}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int = v1
                    v5 = v3
                    for fi17 := 0; fi17 < 18; fi17++ {
                        fmt.Println(v5, fi17)
                        v9 = v9
                        _ = fi17
                    }
                    _ = v16
                }
                _ = v12
                _ = v13
            }
            v18 := GenS1{F0: 305}
            _ = fi11
            _ = v18
        }
        _ = v10
    } else {
        var v19 bool = (0.0 < 1.5)
        v20 := []int{v1, (v1 | v1)}
        fmt.Println(v9, v5, v5)
        _ = v19
        _ = v20
    }
    v21 := []int{}
    v22 := v2[830]
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint32 = 0
        if true {
            v25 := GenS2{F0: 65535, F1: v24, F2: 1}
            _ = v25
        } else {
            var v26 uint = 18446744073709551615
            _ = v26
        }
        v4 -= (v4 % (((^(v4))) | 1))
        _ = fi23
        _ = v24
    }
    v6 -= v9
    for fi27 := 0; fi27 < 4; fi27++ {
        v4 = (v4 * v4)
        v28 := GenS3{F0: false, F1: 1}
        _ = fi27
        _ = v28
    }
    v1 = (v22 & -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
