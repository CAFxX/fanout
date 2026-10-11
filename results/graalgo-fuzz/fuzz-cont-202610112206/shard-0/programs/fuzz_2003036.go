package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 int32) uint16 {
    var v4 float32 = 0.0
    v5 := []int{42}
    _ = p0
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 int16 = 1
    v2 := []int{}
    var v3 int64 = 9223372036854775807
    switch -855 {
    case 32768:
        var v6 bool = (!(!false))
        _ = v6
    default:
        var v7 uint64 = 2
        _ = v7
    }
    _ = len(v2)
    if true {
        v1 = v1
        var v8 uint8 = 128
        _ = v8
    }
    v1 += v1
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v3)
        _ = fi9
    }
    v10 := GenS1{F0: 42, F1: 2, F2: 802}
    switch -255 {
    case -119:
        for fi11 := 0; fi11 < 5; fi11++ {
            if (0 <= 0) {
                v12 := GenS1{F0: 0, F1: 2, F2: 42}
                fmt.Println(v3)
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                if (false && true) {
                    fmt.Println(fi11, v3, v1)
                    var v14 uint64 = 1
                    _ = v14
                } else {
                    var v15 bool = (-2.25 >= (((float32(0.0) - 1.5) * ((0.0 + -2.25) * ((1e10 + 3.14159) / (float32((1e10 - -2.25)) - ((float32(1e10) * 1.5) - 3.14159))))) * -2.25))
                    _ = v15
                }
                if (v3 != v3) {
                    v3 -= v3
                } else {
                    if (!true) {
                        v2 = append(v2, -9223372036854775808)
                        var v16 uint64 = 0
                        if (42 > 787) {
                            if true {
                                v16 = 472
                            }
                            var v17 int64 = (v3 * v3)
                            _ = v17
                        } else {
                            if (42 > -128) {
                                var v18 int16 = v1
                                v18 = 42
                                _ = v18
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v19 uint64 = v16
                            _ = v19
                        }
                        _ = v16
                    }
                    var v20 int64 = (v3 / ((v3) | 1))
                    _ = v20
                }
                v1 = v1
                _ = fi13
            }
            _ = len(v2)
            _ = fi11
        }
    default:
        fmt.Println(v3)
    }
    if true {
        if (false && (((0.0 - 0.1) - (-2.25 - (0.0 * 1e10))) <= ((float32(0.1) + 0.1) / -2.25))) {
            var v21 int8 = 127
            _ = v21
        }
    } else {
        var v22 int8 = 100
        v3 -= (^(v3))
        for fi23 := 0; fi23 < 9; fi23++ {
            switch fi23 {
            case 592:
                if false {
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if len(v2) > 0 {
                            v2[0] = (fi24 - 0)
                        }
                        fmt.Println(v1, v3, v3)
                        var v25 float64 = float64(0.1)
                        _ = fi24
                        _ = v25
                    }
                }
            case 32768:
                for fi26 := 0; fi26 < 2; fi26++ {
                    v22 += v22
                    var v27 uint = 0
                    var v28 uint64 = 228
                    var v29 int8 = 127
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            fmt.Println(v22)
            _ = fi23
        }
        _ = v22
    }
    if false {
        var v30 int = 9223372036854775807
        var v31 int = v30
        var v32 int16 = v1
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 4294967295, F1: 97, F2: 0}
    v2 = append(v2, -2147483648)
    var v34 bool = true
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 uint16 = 303
        v37 := GenS1{F0: 2, F1: v36, F2: 1}
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
