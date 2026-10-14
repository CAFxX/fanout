package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 int8) uint64 {
    switch 42 {
    case 128:
        if (2 <= 42) {
            var v6 uint64 = 4294967295
            for fi7 := 0; fi7 < 17; fi7++ {
                switch "hello" {
                case "":
                    v6 += v6
                case "\n\t":
                    var v8 uint64 = v6
                    _ = v8
                default:
                    fmt.Println(v6, v6)
                }
                _ = fi7
            }
            _ = v6
        } else {
            if false {
                if ((p2 - (p2 ^ p2)) == (p2 + (p2 % (((p2 / (((-(p2))) | 1))) | 1)))) {
                    var v9 string = "hello"
                    var v10 uint32 = 2
                    v11 := GenS1{F0: 907}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                var v12 string = "hello"
                var v13 int32 = -2147483648
                _ = v12
                _ = v13
            }
            var v14 int64 = p0
            v14 = 1
            _ = v14
        }
    case 693:
        var v15 int = 0
        _ = v15
    case 2147483647:
        v16 := GenS1{F0: 42}
        _ = v16
    default:
        var v17 int64 = p0
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS1{F0: 18446744073709551615}
        v20 := GenS1{F0: 42}
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 float64 = 1.5
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = 1.5
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    v5 := []string{"foo bar"}
    v1 = v1
    v2 -= ((0.0 - v2) - float64(0.0))
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int32 = -2147483648
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 18446744073709551615}
    if true {
        var v26 int32 = 1
        _ = v26
    } else {
        v1 = (^(v1))
        v2 = v2
        if (42 < 377) {
            v1 += (v1 - v1)
            if true {
                if ("" > "foo bar") {
                    if (!(42 != 0)) {
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                    } else {
                        var v27 uint16 = 100
                        _ = v27
                    }
                    var v28 uint8 = 2
                    var v29 int16 = -1
                    _ = v28
                    _ = v29
                }
                switch "\n\t" {
                case "":
                    v30 := GenS2{F0: 255, F1: 0}
                    _ = v30
                case "\n\t":
                    v4[v1] = 391
                case "a":
                    var v31 float64 = 3.14159
                    _ = v31
                default:
                    v2 += (3.14159 - 1.5)
                }
                v2 = 0.0
            } else {
                if true {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v33 := GenS1{F0: genFn0(1, -906, 74)}
                        switch 32768 {
                        case -425:
                            v34 := GenS2{F0: 2, F1: 2}
                            _ = v34
                        default:
                            _ = len(v5)
                        }
                        var v35 int16 = 32767
                        _ = fi32
                        _ = v33
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 2; fi36++ {
                        v37 := GenS3{F0: 127, F1: 42}
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v2 = v2
                            _ = fi38
                        }
                        _ = len(v5)
                        _ = fi36
                        _ = v37
                    }
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v2)
                    }
                }
                fmt.Println(v2, v1)
                v39 := GenS1{F0: genFn0(-935, -32768, 100)}
                _ = v39
            }
            _ = len(v5)
        }
    }
    var v40 float64 = (v2 - 3.14159)
    var v41 uint64 = 0
    v40 += ((((v40 / (v2 * 1.5)) * 0.0) - v2) * (v40 * v40))
    for fi42 := 0; fi42 < 11; fi42++ {
        v43 := v3[v1]
        _ = fi42
        _ = v43
    }
    var v44 int16 = -659
    v44 += v44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
}
