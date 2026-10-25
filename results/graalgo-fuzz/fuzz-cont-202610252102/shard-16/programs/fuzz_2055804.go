package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint }
func genFn0(p0 int, p1 uint32, p2 int8) uint32 {
    fmt.Println(p1)
    switch ("foo bar" + ("" + (("" + "hello") + "\n\t"))) {
    case "\n\t":
        v6 := GenS1{F0: 128, F1: 0}
        _ = v6
    case "a":
        v7 := GenS2{F0: 459}
        _ = v7
    default:
        v8 := GenS1{F0: 1, F1: 224}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 / ((42) | 1))
}

func main() {
    var v1 int = -371
    var v2 float32 = (0.1 - (0.0 / (float32(3.14159) + -2.25)))
    var v3 uint32 = 790
    var v4 uint64 = 0
    v5 := make(map[int]string)
    v9 := GenS1{F0: 127, F1: 2}
    var v10 int = (^(v1))
    var v11 int32 = 691
    if ((v2 * v2) > 0.1) {
        var v12 uint8 = 1
        _ = v12
    } else {
        var v13 string = "\n\t"
        var v14 int64 = 81
        fmt.Println(v14)
        _ = v13
        _ = v14
    }
    if ((9223372036854775807 == 9223372036854775807) || true) {
        var v15 uint64 = v4
        _ = v15
    } else {
        var v16 int32 = int32(v4)
        v17 := GenS2{F0: 4294967295}
        _ = v16
        _ = v17
    }
    fmt.Println(v1, v10, v2)
    v11 += v11
    _ = len(v5)
    if true {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 float64 = 3.14159
            _ = fi18
            _ = v19
        }
        var v20 uint64 = (v4 & uint64(v3))
        v21 := GenS3{F0: "\n\t", F1: (v1 + v1)}
        _ = v20
        _ = v21
    }
    if (18446744073709551615 != 423) {
        switch v1 {
        case 32768:
            if true {
                v22 := v5[(v1 + 9223372036854775807)]
                for fi23 := 0; fi23 < 3; fi23++ {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 float64 = (-2.25 + 3.14159)
                        if false {
                            fmt.Println(v1, v11, v11)
                            var v26 bool = ("" == v22)
                            _ = v26
                        } else {
                            v27 := []int{v1, (fi24 - -123), (^(fi23))}
                            _ = v27
                        }
                        var v28 uint8 = 0
                        _ = fi24
                        _ = v25
                        _ = v28
                    }
                    switch 4294967295 {
                    case 194:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 int = 128
                            if false {
                                var v31 uint = 357
                                _ = v31
                            }
                            v1 += (fi29 - 32767)
                            _ = fi29
                            _ = v30
                        }
                    case 42:
                        var v32 int32 = v11
                        _ = v32
                    case 4294967295:
                        var v33 int = fi23
                        _ = v33
                    default:
                        v34 := GenS2{F0: (v4 << ((943) & 63))}
                        _ = v34
                    }
                    v35 := GenS4{F0: -128}
                    _ = fi23
                    _ = v35
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    fmt.Println(fi36, fi36)
                    _ = fi36
                }
                _ = v22
            } else {
                v4 -= (v4 + 520)
                v37 := GenS5{F0: v10, F1: 4294967295}
                _ = v37
            }
        case -256:
            v1 += v1
        default:
            var v38 uint8 = 0
            _ = v38
        }
    } else {
        switch v10 {
        case -837:
            var v39 float32 = v2
            _ = v39
        case -1:
            var v40 uint16 = 32767
            _ = v40
        }
    }
    v41 := []int{int(v11), (v1 & v10), (^(v10))}
    v11 = (v11 - v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v9
}
