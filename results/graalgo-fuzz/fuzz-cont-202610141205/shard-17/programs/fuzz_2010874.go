package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 uint64, p1 uint16) int64 {
    var v4 int16 = -32768
    var v5 int64 = 0
    v6 := []int{-1, -831, -2147483648}
    v4 -= v4
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int) uint {
    if ((0.0 + float64(float32(float64((((0.1 - float32((1e10 - 3.14159))) + 0.0) + (1e10 * (-2.25 / 3.14159))))))) != 1e10) {
        switch p0 {
        case -9223372036854775808:
            v7 := []int{(p0 + (p0 | (p0 | p0))), 127, p0}
            _ = v7
        case 504:
            if true {
                for fi8 := 0; fi8 < 10; fi8++ {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 bool = (0 > fi8)
                        fmt.Println(fi9, fi9)
                        if ((((float64((1.5 + (1e10 * 1.5))) >= 1e10) && (v10 && (v10 && (!v10)))) || v10) && (245 > -1)) {
                            v10 = (2147483647 != 202)
                        }
                        _ = fi9
                        _ = v10
                    }
                    p0 -= (fi8 * fi8)
                    var v11 int8 = int8(p0)
                    _ = fi8
                    _ = v11
                }
            }
        default:
            v12 := []int{(p0 | p0)}
            _ = v12
        }
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := []int{-1, 42, 1}
    v2 := make(map[string]uint32)
    var v3 uint32 = 2147483647
    v13 := GenS1{F0: -96, F1: 65535, F2: "hello"}
    v3 = 789
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int8 = -1
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 42
    fmt.Println(v14, v17, v3)
    fmt.Println(v14, v3, v17)
    var v18 uint = (v14 >> ((4294967295) & 63))
    _ = len(v1)
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float32 = 1.5
        _ = fi19
        _ = v20
    }
    var v21 float32 = (0.1 + (0.1 * 1.5))
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v22 float64 = 3.14159
    if true {
        v23 := GenS2{F0: v3}
        _ = v23
    } else {
        if (!(v22 > (-2.25 - 1.5))) {
            v3 += (^(v3))
            if (-3 != 127) {
                if len(v1) > 0 {
                    v1[0] = -497
                }
                v24 := GenS3{F0: -95, F1: 32767}
                var v25 int = 42
                _ = v24
                _ = v25
            } else {
                var v26 float32 = v21
                if len(v1) > 0 {
                    v1[0] = 32767
                }
                _ = v26
            }
        }
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
    }
    var v27 int = -2147483648
    switch (v27 - 812) {
    case 42:
        v18 = 42
    case -121:
        var v28 int = int(v17)
        _ = v28
    case -9223372036854775808:
        if (v21 < float32(((float64(v21) - 1.5) / float64(v21)))) {
            var v29 string = (("\n\t" + "hello") + "hello")
            var v30 uint32 = 2
            v27 += 968
            _ = v29
            _ = v30
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
}
