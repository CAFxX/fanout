package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8) int64 {
    v5 := GenS1{F0: "", F1: 441, F2: (p0 | p1)}
    var v6 uint16 = 0
    var v7 uint8 = p1
    v7 = 228
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 799
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    var v3 int8 = 72
    var v4 bool = true
    var v9 float32 = 1.5
    v10 := []int{9223372036854775807, -9223372036854775808, -9223372036854775808}
    if v4 {
        if (-519 != 32767) {
            _ = len(v1)
        }
        var v11 uint8 = 106
        v4 = (((v9 / v9) + v9) < -2.25)
        _ = v11
    }
    if true {
        if false {
            v4 = true
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(v3, v9)
                if (-1 == 0) {
                    v13 := v1[2]
                    _ = v13
                } else {
                    fmt.Println(fi12)
                }
                _ = fi12
            }
            _ = len(v10)
        }
    } else {
        fmt.Println(v3)
        v14 := GenS2{F0: 0}
        _ = v14
    }
    v2 = v2
    switch 389 {
    case 9223372036854775807:
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 string = "\n\t"
                if v4 {
                    fmt.Println(v3)
                } else {
                    v18 := GenS1{F0: v17, F1: 135, F2: 2}
                    v4 = ((((!(0 <= 2)) && (-2147483648 == 42)) || (9223372036854775807 >= 42)) && v4)
                    v3 -= (v3 << ((65535) & 63))
                    _ = v18
                }
                var v19 uint32 = 0
                _ = fi16
                _ = v17
                _ = v19
            }
            var v20 string = "hello"
            v2 = v2
            _ = fi15
            _ = v20
        }
    case 32767:
        switch -711 {
        case -1:
            var v21 uint = 1
            _ = v21
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 19; fi23++ {
                    v4 = (!((v3 * 100) <= v3))
                    if (((0 != -1) || v4) || false) {
                        var v24 uint16 = 255
                        fmt.Println(v9, v9)
                        var v25 int32 = -2147483648
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
                v9 = 1.5
                _ = fi22
            }
        }
    case 1:
        v1[2147483648] = 2147483648
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            var v27 uint8 = 42
            delete(v1, 14)
            v4 = false
            _ = fi26
            _ = v27
        }
    }
    v9 += (v9 * (1e10 * (0.0 / 1.0)))
    v10 = append(v10, 0)
    v4 = false
    v28 := GenS2{F0: -32768}
    for fi29 := 0; fi29 < 18; fi29++ {
        if true {
            fmt.Println(fi29)
            fmt.Println(v9, v4)
        } else {
            if ((fi29 ^ (fi29 * fi29)) > 127) {
                v3 += (v3 + 127)
            }
            switch (fi29 * fi29) {
            case 128:
                fmt.Println(v2, fi29, v4)
            case -9223372036854775808:
                fmt.Println(v2, v9, v3)
            case 674:
                fmt.Println(fi29)
            }
        }
        v2 += (v2 & 0)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi29
    }
    var v30 uint8 = 127
    v4 = v4
    var v31 float64 = (float64(((v9 + -2.25) / (1e10 - v9))) / 1.5)
    var v32 uint16 = (v2 & uint16(v30))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
