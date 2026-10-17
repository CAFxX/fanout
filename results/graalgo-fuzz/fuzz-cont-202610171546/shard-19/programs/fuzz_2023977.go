package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int16; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
func genFn0() uint {
    var v4 float64 = 3.14159
    switch -1 {
    case 0:
        fmt.Println(v4)
    case -456:
        v5 := GenS1{F0: -2147483648, F1: -650}
        _ = v5
    case -2147483648:
        v4 = ((v4 * (1e10 + (1e10 / float64(3.14159)))) - -2.25)
    }
    _ = v4
    return 2
}
func genFn1(p0 int64) int32 {
    if ((42 >= -77) && true) {
        if true {
            if (false || (-618 <= -137)) {
                var v6 uint16 = 2
                _ = v6
            }
            p0 = p0
            var v7 int16 = 966
            _ = v7
        }
        p0 -= (p0 ^ -1)
        v8 := []int{-9223372036854775808, 9223372036854775807}
        _ = v8
    } else {
        p0 = (p0 & 42)
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{(fi9 * (fi9 + (fi9 ^ (fi9 & (fi9 | 32767))))), (fi9 ^ fi9), (fi9 & fi9)}
                p0 += 185
                if (true && false) {
                    v10 = append(v10, 42)
                    p0 = p0
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                _ = fi9
                _ = v10
            }
            v11 := []int{4294967295, -2147483648, 127, 1}
            _ = v11
        } else {
            fmt.Println(p0)
            if ((-1 < -320) || (1.5 > (1e10 + 0.1))) {
                p0 -= p0
                p0 -= p0
            }
            v12 := []int{-2147483648}
            _ = v12
        }
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 uint8 = (^(v1))
    v3 := make(map[string]string)
    delete(v3, "a")
    var v13 float64 = 0.1
    v14 := GenS2{F0: 32767, F1: 0}
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint = 123
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(fi17, fi15)
            v18 := GenS3{F0: ("a" + "hello")}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 float32 = -2.25
            if false {
                _ = len(v3)
                v21 := []int{(fi15 * fi19), (fi19 ^ -9223372036854775808), fi15}
                _ = v21
            }
            v16 += genFn0()
            if false {
                var v22 bool = (!true)
                v23 := GenS4{F0: 1, F1: -786, F2: v22}
                _ = v22
                _ = v23
            } else {
                v24 := GenS5{F0: 65535}
                _ = v24
            }
            _ = fi19
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    v25 := []int{-9223372036854775808, -907, -448, 127}
    fmt.Println(v1, v13)
    v2 -= v2
    v13 += (float64(0.1) / v13)
    v26 := GenS3{F0: (("foo bar" + ("" + (("" + "foo bar") + ""))) + "")}
    if false {
        v13 = v13
        v13 = v13
        fmt.Println(v13)
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS6{F0: true}
        var v29 uint32 = 2
        for fi30 := 0; fi30 < 16; fi30++ {
            v13 = (v13 + (((0.1 + ((v13 / (v13 - -2.25)) + (v13 - (0.0 - (1.5 - v13))))) / float64(-2.25)) * v13))
            v31 := GenS2{F0: 255, F1: int16(v1)}
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
