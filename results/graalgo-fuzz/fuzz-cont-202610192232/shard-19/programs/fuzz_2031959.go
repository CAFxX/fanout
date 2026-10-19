package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
type GenS3 struct { F0 bool }
func genFn0() int {
    v7 := GenS1{F0: 127, F1: 100}
    var v8 int = 42
    var v9 int32 = 2147483647
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS2{F0: 65535, F1: 5, F2: 0}
        var v12 uint = 18446744073709551615
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: -128, F1: 42}
    _ = v13
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 4294967295
    var v3 int8 = 127
    var v4 bool = false
    v5 := make(map[int]int)
    v6 := []uint32{1, 4294967295}
    v2 = (v2 - 0)
    if (!true) {
        v3 = -128
        v3 = v3
    }
    v4 = v4
    v14 := GenS3{F0: true}
    v4 = v4
    if (1 > 2) {
        if (-9223372036854775808 > -616) {
            v3 += (v3 - (v3 - v3))
            for fi15 := 0; fi15 < 11; fi15++ {
                v4 = (!(125 > 166))
                v4 = true
                for fi16 := 0; fi16 < 6; fi16++ {
                    if true {
                        v4 = (42 == 2)
                        v2 -= (fi15 * (fi15 % ((genFn0()) | 1)))
                        switch fi15 {
                        case 255:
                            var v17 bool = (!v4)
                            _ = v17
                        case 128:
                            v4 = true
                        case 1:
                            var v18 int64 = -927
                            _ = v18
                        default:
                            fmt.Println(fi16)
                        }
                        var v19 bool = ((2 >= 0) && ((4294967295 < 361) || true))
                        _ = v19
                    } else {
                        var v20 uint16 = 122
                        if (!v4) {
                            var v21 int16 = 1
                            _ = v21
                        }
                        if ((1e10 * (3.14159 + 0.1)) > 3.14159) {
                            if (false && false) {
                                v20 = (v20 % ((793) | 1))
                            } else {
                                var v22 int16 = -116
                                var v23 uint8 = 255
                                var v24 float64 = (3.14159 + (float64(3.14159) / 3.14159))
                                fmt.Println(v22, v2, v2)
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        } else {
                            v25 := GenS3{F0: (!v4)}
                            v4 = v4
                            _ = v25
                        }
                        v26 := GenS2{F0: v20, F1: int16(v3), F2: 42}
                        _ = len(v5)
                        _ = v20
                        _ = v26
                    }
                    _ = fi16
                }
                _ = fi15
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS3{F0: v4}
                _ = fi27
                _ = v28
            }
        } else {
            v4 = ((-9223372036854775808 >= (v2 * (v2 ^ (v2 * v2)))) || v4)
        }
    }
    v3 += (v3 % ((v3) | 1))
    v29 := GenS3{F0: true}
    var v30 float64 = 0.0
    var v31 int = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
