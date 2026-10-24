package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 bool }
func genFn0(p0 int, p1 int64) int16 {
    var v5 float64 = 1e10
    v6 := GenS1{F0: 100}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1000
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{}
    switch -746 {
    case 32767:
        v7 := GenS2{F0: -873, F1: 4294967295, F2: 394}
        _ = v7
    case -401:
        v8 := GenS1{F0: 100}
        _ = v8
    case 9223372036854775807:
        fmt.Println(v2)
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 float32 = v2
        _ = fi9
        _ = v10
    }
    if (false || (float32(-2.25) < (1.5 - v2))) {
        var v11 float64 = 0.0
        switch -346 {
        case 1:
            if (true && false) {
                var v12 int = -100
                if (("hello" + "a") != "a") {
                    var v13 bool = false
                    for fi14 := 0; fi14 < 17; fi14++ {
                        if false {
                            fmt.Println(v11, v2)
                            if v13 {
                                v13 = v13
                            }
                            delete(v3, 65535)
                        } else {
                            for fi15 := 0; fi15 < 14; fi15++ {
                                if v13 {
                                    var v16 int16 = 151
                                    _ = v16
                                } else {
                                    v17 := GenS2{F0: (fi15 * fi15), F1: 42, F2: 65535}
                                    if v13 {
                                        switch "hello" {
                                        case "":
                                            if v13 {
                                                v12 -= 1
                                            }
                                        case "a":
                                            var v18 int32 = -2147483648
                                            _ = v18
                                        case "\n\t":
                                            v19 := GenS2{F0: v12, F1: 997, F2: 2147483648}
                                            _ = v19
                                        }
                                        var v20 int16 = 0
                                        v2 -= 0.1
                                        _ = v20
                                    }
                                    _ = v17
                                }
                                v2 += v2
                                v21 := GenS3{F0: 127, F1: 32768, F2: (!false)}
                                _ = fi15
                                _ = v21
                            }
                        }
                        for fi22 := 0; fi22 < 7; fi22++ {
                            if v13 {
                                _ = len(v3)
                                delete(v3, 4294967295)
                            }
                            v4 = append(v4, 781)
                            v11 -= float64(v2)
                            _ = fi22
                        }
                        _ = fi14
                    }
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 uint = 4294967295
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v25 := GenS2{F0: (v12 / ((-2147483648) | 1)), F1: v24, F2: 65535}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v13
                }
                _ = v12
            }
        default:
            if len(v4) > 0 {
                v4[0] = 656
            }
        }
        _ = v11
    } else {
        if (1000 != 262) {
            var v26 int8 = 1
            v2 = (0.1 + (((1e10 - v2) - v2) * ((3.14159 * float32(1.5)) - v2)))
            for fi27 := 0; fi27 < 19; fi27++ {
                v26 -= v26
                _ = fi27
            }
            _ = v26
        } else {
            var v28 int32 = -1
            _ = v28
        }
    }
    v2 = 3.14159
    v29 := GenS3{F0: 127, F1: 127, F2: false}
    v4 = append(v4, 251)
    v30 := GenS3{F0: 234, F1: 640, F2: (2147483647 <= 1)}
    fmt.Println(v2)
    v2 += (v2 * 1e10)
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
