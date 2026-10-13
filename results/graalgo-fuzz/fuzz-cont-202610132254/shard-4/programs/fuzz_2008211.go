package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0(p0 int16) int32 {
    var v4 uint = 413
    var v5 int8 = 100
    if (-800 <= 412) {
        v6 := []int{-9223372036854775808, -837, -380}
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int8 = (v5 ^ (v5 & v5))
            if true {
                if (-1 < 463) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v9 int = -155
                    v5 -= 42
                    _ = v9
                }
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                var v10 int = fi7
                _ = v10
            }
            fmt.Println(v8, v5, p0)
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v11 float32 = (3.14159 + (0.1 / 1e10))
        _ = v11
    }
    if (229 >= 313) {
        var v12 uint16 = 955
        var v13 int16 = (^(p0))
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}
func genFn1() uint16 {
    var v14 uint = 945
    var v15 bool = false
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v14, v15)
        fmt.Println(v15)
        v17 := []int{2147483647}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: (1e10 != 1.5), F1: -505}
    for fi19 := 0; fi19 < 12; fi19++ {
        v14 = 984
        _ = fi19
    }
    _ = v14
    _ = v15
    _ = v18
    return 0
}

func main() {
    var v1 uint8 = 0
    var v2 int32 = -2147483648
    v3 := make(map[int]uint32)
    if true {
        v20 := GenS2{F0: 42, F1: v1}
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS3{F0: v2, F1: 1}
            v23 := GenS3{F0: genFn0(-635), F1: 1}
            fmt.Println(v2, fi21)
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    }
    switch 457 {
    case -9223372036854775808:
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 int16 = 1
            var v26 uint8 = v1
            _ = fi24
            _ = v25
            _ = v26
        }
    case 4294967295:
        v1 += (v1 * v1)
    default:
        var v27 int = 32767
        _ = v27
    }
    v1 -= (v1 >> ((255) & 63))
    var v28 int32 = (v2 >> ((255) & 63))
    var v29 int64 = 42
    for fi30 := 0; fi30 < 15; fi30++ {
        v28 = (-(v2))
        fmt.Println(v28, v2)
        _ = fi30
    }
    if true {
        v2 += (v2 + (-(v2)))
        var v31 int64 = (v29 + v29)
        fmt.Println(v31, v2)
        _ = v31
    }
    var v32 string = "\n\t"
    v33 := GenS4{F0: 2, F1: ""}
    var v34 int32 = -437
    switch "a" {
    case "hello":
        v29 = (v29 / ((42) | 1))
    case "\n\t":
        v34 += genFn0(1)
    default:
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 float64 = ((1e10 * (((0.1 * float64(0.0)) - 1e10) + (3.14159 * 0.0))) / 3.14159)
            if true {
                fmt.Println(v1, v29)
            } else {
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 uint32 = 316
                    _ = fi37
                    _ = v38
                }
                var v39 int8 = 42
                _ = v39
            }
            fmt.Println(v36)
            _ = fi35
            _ = v36
        }
    }
    var v40 uint64 = 2
    v41 := []int{9223372036854775807, 2147483647, 724, 575}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v40
    _ = v41
}
