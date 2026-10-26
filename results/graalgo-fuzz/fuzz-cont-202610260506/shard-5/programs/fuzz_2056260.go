package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint, p1 uint8) string {
    var v4 uint16 = 2
    v5 := GenS1{F0: p1, F1: true}
    var v6 bool = (!(int32(p0) != 2147483647))
    var v7 int16 = 1000
    v4 = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 int8 = 100
    var v2 int8 = (v1 ^ (v1 | v1))
    var v3 bool = (((v1 + v1) != v2) || (float64(float32(float64((((0.1 - 1e10) - 0.1) + 0.1)))) > 1.5))
    if false {
        switch "\n\t" {
        case "a":
            var v8 int32 = 1
            _ = v8
        default:
            for fi9 := 0; fi9 < 20; fi9++ {
                v1 -= 127
                var v10 int64 = 42
                _ = fi9
                _ = v10
            }
        }
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v1, fi11)
        if v3 {
            if v3 {
                var v12 float64 = 3.14159
                _ = v12
            }
            var v13 int8 = v1
            var v14 uint8 = 127
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 uint32 = 290
    for fi16 := 0; fi16 < 9; fi16++ {
        switch "" {
        case "hello":
            v2 = v1
        default:
            var v17 uint16 = 65535
            _ = v17
        }
        var v18 uint32 = v15
        fmt.Println(v18)
        _ = fi16
        _ = v18
    }
    v19 := GenS2{F0: 78, F1: v15}
    v1 = v2
    switch "a" {
    case "\n\t":
        v20 := []int{9223372036854775807}
        _ = v20
    case "hello":
        if v3 {
            switch "" {
            case "hello":
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int = (fi21 + -742)
                    _ = fi21
                    _ = v22
                }
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    if true {
                        v1 = (v2 + (^(v2)))
                        if true {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                var v25 uint16 = 32767
                                _ = fi24
                                _ = v25
                            }
                            switch fi23 {
                            case 755:
                                switch 127 {
                                case -9223372036854775808:
                                    v2 = v1
                                case 105:
                                    if ("a" != "foo bar") {
                                        v26 := GenS2{F0: 797, F1: v15}
                                        var v27 bool = (v3 && true)
                                        if v27 {
                                            var v28 uint16 = 497
                                            for fi29 := 0; fi29 < 20; fi29++ {
                                                var v30 int16 = 422
                                                var v31 string = "foo bar"
                                                var v32 int64 = 1
                                                v32 -= v32
                                                _ = fi29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = v28
                                        }
                                        _ = v26
                                        _ = v27
                                    } else {
                                        switch 913 {
                                        case 2147483647:
                                            var v33 int32 = 2147483647
                                            _ = v33
                                        case 4294967295:
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v2, fi23)
                                            default:
                                                var v34 int8 = v1
                                                _ = v34
                                            }
                                        }
                                        v2 = v2
                                    }
                                default:
                                    var v35 float64 = ((-2.25 + 0.0) * 3.14159)
                                    _ = v35
                                }
                            }
                            var v36 float64 = (((1e10 / -2.25) / (1.5 * 0.1)) + ((3.14159 / float64(-2.25)) - (1e10 + (float64(float32(0.0)) + 0.1))))
                            _ = v36
                        }
                    }
                    var v37 float64 = float64((float32(0.1) / (float32(0.1) - 1.5)))
                    for fi38 := 0; fi38 < 16; fi38++ {
                        var v39 uint16 = 255
                        _ = fi38
                        _ = v39
                    }
                    _ = fi23
                    _ = v37
                }
            }
            var v40 uint32 = v15
            _ = v40
        }
    default:
        if (1 <= 255) {
            v41 := []int{127}
            var v42 int64 = -9223372036854775808
            var v43 uint16 = 0
            _ = v41
            _ = v42
            _ = v43
        } else {
            fmt.Println(v15)
        }
    }
    v44 := []int{9223372036854775807, 888}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v44
}
