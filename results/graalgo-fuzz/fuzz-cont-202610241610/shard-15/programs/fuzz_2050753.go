package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint, p1 uint) int32 {
    v4 := GenS1{F0: -726, F1: 9223372036854775807}
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = 0
            p1 += (p0 ^ (p1 - 65535))
            v7 += 32767
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 2
        fmt.Println(fi5, v8)
        _ = fi5
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint8, p1 int16) int64 {
    fmt.Println(p1, p0)
    var v9 int64 = 9223372036854775807
    v9 -= v9
    _ = p0
    _ = p1
    _ = v9
    return 434
}

func main() {
    v1 := []int{305, 32767, 0, 55}
    v2 := []uint32{}
    v3 := make(map[string]int)
    switch 136 {
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if true {
        if (-1 < 100) {
            for fi10 := 0; fi10 < 9; fi10++ {
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (25 != 237) {
                        var v12 uint32 = 4294967295
                        _ = v12
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch 9223372036854775807 {
                    case 128:
                        var v13 float32 = 1.5
                        _ = v13
                    case 4294967295:
                        var v14 string = "a"
                        _ = v14
                    case -624:
                        v15 := GenS1{F0: 690, F1: (fi11 >> ((18446744073709551615) & 63))}
                        _ = v15
                    default:
                        var v16 float64 = -2.25
                        _ = v16
                    }
                    if (-24 < -12) {
                        var v17 uint16 = 2
                        v18 := GenS2{F0: fi11, F1: 924, F2: 509}
                        if (0 > 854) {
                            var v19 float64 = (0.0 * 0.0)
                            if (!true) {
                                v19 += 1e10
                                v17 -= (v17 >> ((573) & 63))
                            } else {
                                var v20 float32 = 3.14159
                                _ = len(v2)
                                _ = v20
                            }
                            v17 += v17
                            _ = v19
                        } else {
                            fmt.Println(fi11, v17)
                            var v21 int32 = 2147483647
                            var v22 int32 = v21
                            _ = v21
                            _ = v22
                        }
                        v17 += (v17 - v17)
                        _ = v17
                        _ = v18
                    }
                    _ = fi11
                }
                var v23 string = "a"
                _ = fi10
                _ = v23
            }
            var v24 int = 9223372036854775807
            var v25 int16 = 1000
            _ = v24
            _ = v25
        }
        var v26 float32 = 0.1
        v27 := v3[""]
        _ = v26
        _ = v27
    }
    switch 9223372036854775807 {
    case -2147483648:
        var v28 int8 = 127
        _ = v28
    case 1:
        var v29 float64 = (1e10 - (1.5 * -2.25))
        _ = v29
    default:
        var v30 int16 = -578
        _ = v30
    }
    var v31 int32 = -977
    v32 := GenS3{F0: 9223372036854775807}
    v33 := GenS4{F0: -167}
    fmt.Println(v31)
    var v34 int = 847
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
