package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0(p0 uint, p1 uint8) bool {
    var v4 int16 = 1000
    if (-9223372036854775808 > 127) {
        p1 = (p1 ^ 128)
        v5 := GenS1{F0: (v4 << ((42) & 63)), F1: 1}
        v6 := []int{4294967295}
        _ = v5
        _ = v6
    } else {
        if (false || (-1 != -744)) {
            var v7 float64 = 1.5
            if (42 <= -1) {
                v8 := []int{4294967295, -537, 846}
                _ = v8
            } else {
                v9 := GenS1{F0: v4, F1: 0}
                v10 := GenS2{F0: 4294967295, F1: 123, F2: -128}
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            p0 -= p0
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return ((v4 * (^(v4))) < 1000)
}

func main() {
    var v1 float64 = (((float64((1e10 * -2.25)) * 0.1) / (1e10 / (0.1 / 1.0))) * 1e10)
    var v2 float64 = v1
    var v3 int16 = 0
    if (-9223372036854775808 < 9223372036854775807) {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint64 = 1
            fmt.Println(v12, v3)
            fmt.Println(fi11)
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS2{F0: 564, F1: 0, F2: -26}
            var v15 bool = (!(3.14159 > -2.25))
            fmt.Println(v15, v3)
            _ = fi13
            _ = v14
            _ = v15
        }
        v2 = 3.14159
    }
    v16 := GenS2{F0: int(v3), F1: 0, F2: -17}
    v17 := GenS3{F0: 0, F1: 0}
    if (-1 > 0) {
        if true {
            v18 := []int{-449, 32767, 127, 127}
            _ = v18
        }
    }
    fmt.Println(v1, v2)
    v19 := []int{4294967295, 314, 2147483647}
    var v20 uint = 2
    var v21 int = 0
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    if false {
        if (0 > -1) {
            v22 := GenS4{F0: 194, F1: 127}
            fmt.Println(v1, v1)
            _ = v22
        } else {
            var v23 float32 = (1e10 / (1.5 / (((1e10 - ((1e10 / 0.1) + ((0.1 + ((float32(v1) - (1e10 - (1e10 + (float32(float64(0.0)) - float32((v1 - float64(3.14159))))))) + 0.1)) - 1e10))) / 0.1) + float32(v1))))
            _ = v23
        }
        if false {
            v1 -= (0.1 + float64(0.1))
        }
    }
    v1 += v1
    var v24 bool = false
    v3 += v3
    if (v21 >= 239) {
        for fi25 := 0; fi25 < 6; fi25++ {
            fmt.Println(v3, v1, v2)
            v20 += v20
            _ = fi25
        }
        v19 = append(v19, (^(v21)))
        if false {
            if len(v19) > 0 {
                v19[0] = (v21 - -1)
            }
            for fi26 := 0; fi26 < 11; fi26++ {
                v20 = v20
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 float32 = ((0.1 * (-2.25 - 1e10)) + 0.1)
                    if v24 {
                        var v29 int64 = 9223372036854775807
                        v30 := GenS1{F0: (v3 >> (((v20 & (v20 - v20))) & 63)), F1: 0}
                        fmt.Println(v1, v21, v24)
                        if (v28 > float32(float64(float32((v2 * v2))))) {
                            v1 = 0.1
                        } else {
                            v2 += v2
                        }
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                v19 = append(v19, (v21 * fi26))
                _ = fi26
            }
            switch "a" {
            case "foo bar":
                v20 += v20
            case "\n\t":
                v31 := GenS2{F0: v21, F1: 42, F2: 127}
                _ = v31
            default:
                var v32 int8 = 42
                _ = v32
            }
        }
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v1 = v2
        v2 = 3.14159
        _ = fi33
    }
    if v24 {
        var v34 uint64 = 2
        if v24 {
            v21 = 9223372036854775807
        }
        if (v24 && (v24 && (!false))) {
            if len(v19) > 0 {
                v19[0] = v21
            }
            var v35 bool = ((false && false) || v24)
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = v35
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
}
