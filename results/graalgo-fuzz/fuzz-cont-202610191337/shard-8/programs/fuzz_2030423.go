package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int8, p1 int16, p2 int16) uint16 {
    v4 := []int{0, 461, 4294967295}
    switch ("" + ("\n\t" + ("" + ("hello" + ("foo bar" + ("hello" + "hello")))))) {
    case "":
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
    default:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint32, p1 int64) int64 {
    fmt.Println(p0)
    var v5 bool = (false && true)
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 uint16 = 65535
    v6 := GenS1{F0: (((114 >= -1) && true) || true), F1: 255, F2: "hello"}
    var v7 uint = 1
    var v8 int16 = 1000
    v9 := GenS2{F0: -22, F1: (v7 >> ((v7) & 63)), F2: 1}
    v3 -= genFn0(-106, v8, (^(v8)))
    var v10 uint = (v7 & (v7 / ((v7) | 1)))
    v10 += v7
    var v11 float64 = float64((3.14159 * 0.1))
    var v12 uint16 = genFn0(-128, (v8 >> ((v10) & 63)), -32768)
    v13 := GenS1{F0: true, F1: (v12 | 575), F2: "foo bar"}
    v14 := []int{}
    var v15 int16 = v8
    if (false || false) {
        v16 := v2[-485]
        var v17 float64 = v11
        v8 -= (^(v15))
        _ = v16
        _ = v17
    } else {
        fmt.Println(v12)
        v18 := GenS3{F0: (v12 <= genFn0(110, 32767, v8)), F1: (-(v15)), F2: (v7 / ((v7) | 1))}
        if false {
            v7 -= uint(v12)
            v10 = v10
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                v10 += 65535
                v12 -= (v3 | v12)
                v20 := GenS2{F0: -566, F1: v7, F2: 42}
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                if (2147483647 >= 42) {
                    var v22 int = 0
                    for fi23 := 0; fi23 < 17; fi23++ {
                        switch "hello" {
                        case "a":
                            var v24 string = "\n\t"
                            _ = v24
                        case "foo bar":
                            v7 = 65535
                            var v25 int32 = 852
                            v26 := GenS4{F0: 4294967295}
                            _ = v25
                            _ = v26
                        default:
                            var v27 uint16 = 32767
                            _ = v27
                        }
                        v12 += 997
                        _ = fi23
                    }
                    fmt.Println(fi21, v3, v22)
                    _ = v22
                }
                var v28 int = fi21
                _ = fi21
                _ = v28
            }
        }
        _ = v18
    }
    var v29 uint64 = 1
    v14 = append(v14, -560)
    for fi30 := 0; fi30 < 2; fi30++ {
        if ("hello" > "") {
            fmt.Println(fi30, v29, v29)
            v31 := GenS3{F0: (18446744073709551615 < 801), F1: v8, F2: (^(v10))}
            _ = v31
        }
        v29 = v29
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
