package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint, p2 int32) int {
    for fi4 := 0; fi4 < 14; fi4++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(p0, p2)
        case "a":
            v5 := []int{fi4}
            _ = v5
        default:
            var v6 uint = p1
            _ = v6
        }
        _ = fi4
    }
    p1 = p1
    v7 := []int{}
    v8 := GenS1{F0: 0, F1: (p1 + (p1 ^ (p1 * (^(p1)))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = (1e10 - 1.5)
    v3 := make(map[string]uint32)
    v9 := GenS2{F0: 923}
    v10 := []int{-9223372036854775808, genFn0(989, 2, 2147483647), 0, 42}
    for fi11 := 0; fi11 < 14; fi11++ {
        if ((float64(float32(((1.5 + (v2 / v2)) + ((v2 / v2) - float64(-2.25))))) / 1.5) >= v2) {
            v2 = v2
            var v12 bool = (65535 <= 18446744073709551615)
            _ = v12
        } else {
            if (-247 < 831) {
                v2 = (v2 - float64(0.0))
                v2 -= -2.25
                switch ("" + ("hello" + "\n\t")) {
                case "hello":
                    var v13 float64 = v2
                    _ = v13
                case "foo bar":
                    v1 = (v1 % (((v1 | v1)) | 1))
                default:
                    for fi14 := 0; fi14 < 4; fi14++ {
                        if true {
                            var v15 int64 = v1
                            _ = v15
                        } else {
                            var v16 int64 = (^(v1))
                            var v17 float32 = 1e10
                            var v18 int64 = v16
                            v19 := GenS2{F0: 65535}
                            if (1 > -128) {
                                v17 -= v17
                                var v20 int16 = 360
                                v17 = (((float32(v2) - (((0.1 - v17) / (v17 - -2.25)) / (v17 / v17))) / -2.25) - float32(v2))
                                var v21 int16 = 42
                                var v22 int32 = 2147483647
                                fmt.Println(v18, v16)
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = fi14
                    }
                }
            }
            v1 -= (v1 + v1)
        }
        if false {
            v1 = v1
            var v23 string = "a"
            _ = v23
        }
        v2 += v2
        _ = fi11
    }
    v1 += (v1 % (((^(v1))) | 1))
    var v24 float32 = (float32(v2) - (float32(0.0) * -2.25))
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(fi25, v2)
        var v26 uint16 = 1
        _ = fi25
        _ = v26
    }
    if false {
        var v27 float64 = v2
        fmt.Println(v2, v1)
        v27 = (float64(1.5) * v2)
        _ = v27
    } else {
        var v28 uint64 = 4294967295
        v29 := GenS2{F0: 32767}
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 int = (-(fi30))
                var v33 uint16 = 65535
                _ = fi31
                _ = v32
                _ = v33
            }
            v1 = (^(v1))
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    v34 := GenS2{F0: 32767}
    var v35 int64 = v1
    if false {
        v3["foo bar"] = 65535
        fmt.Println(v24, v35)
        v36 := GenS2{F0: 319}
        _ = v36
    } else {
        if false {
            v2 = float64((float32(float64(v24)) / 1.5))
        }
        v2 = v2
    }
    var v37 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v9
}
