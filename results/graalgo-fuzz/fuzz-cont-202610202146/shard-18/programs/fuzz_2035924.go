package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 uint, p1 bool) uint {
    v6 := GenS1{F0: 0, F1: -1}
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            switch 924 {
            case 0:
                v8 := []int{(fi7 / ((fi7) | 1)), (^(fi7)), (fi7 * 1)}
                _ = v8
            case -1:
                v9 := []int{}
                _ = v9
            case 2147483647:
                if p1 {
                    v10 := GenS2{F0: 364, F1: p0}
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 float32 = (0.0 / (0.0 + ((1e10 - (1e10 / 3.14159)) * (1e10 / 1.0))))
                        var v13 int64 = -154
                        var v14 int = fi7
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    switch (fi7 * fi7) {
                    case 127:
                        v15 := []int{}
                        _ = v15
                    }
                }
            default:
                var v16 float64 = 3.14159
                _ = v16
            }
        }
        var v17 uint16 = 255
        if (p1 && true) {
            var v18 int8 = 1
            var v19 float64 = (float64(0.1) / float64((0.1 - (1.5 + (3.14159 + float32(-2.25))))))
            _ = v18
            _ = v19
        }
        _ = fi7
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1(p0 int64, p1 bool, p2 int8) uint32 {
    p2 -= (p2 & 0)
    if (!p1) {
        v20 := GenS2{F0: -252, F1: 364}
        fmt.Println(p0)
        var v21 bool = false
        _ = v20
        _ = v21
    } else {
        v22 := GenS2{F0: -70, F1: 1}
        _ = v22
    }
    if p1 {
        var v23 bool = ((32767 > 9223372036854775807) || ((0 < 700) || (4294967295 < 2)))
        var v24 int8 = p2
        _ = v23
        _ = v24
    }
    v25 := []int{-1, -817, -167, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 837
}

func main() {
    var v1 uint8 = 1
    var v2 int = 9223372036854775807
    v3 := make(map[int]string)
    var v4 uint8 = (v1 - v1)
    var v5 float32 = 0.1
    if true {
        var v26 int = (v2 + -9223372036854775808)
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 float64 = (0.0 - -2.25)
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    _ = len(v3)
    if false {
        var v29 float64 = 3.14159
        v30 := v3[-9223372036854775808]
        v31 := []int{v2, int(v1), -176, 486}
        _ = v29
        _ = v30
        _ = v31
    }
    v1 -= (^(v4))
    v32 := []int{(v2 % ((-2147483648) | 1)), v2, (v2 % ((v2) | 1))}
    var v33 int64 = 665
    var v34 uint16 = 42
    var v35 uint16 = v34
    v36 := v3[(v2 | -525)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
