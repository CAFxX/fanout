package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint, p2 uint64) bool {
    if (1.5 >= (-2.25 - 3.14159)) {
        if (42 < 0) {
            var v5 uint16 = 747
            p0 = p0
            _ = v5
        }
        var v6 int8 = 42
        _ = v6
    } else {
        v7 := []int{0}
        fmt.Println(p1, p0)
        _ = v7
    }
    var v8 uint = uint(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return true
}
func genFn1(p0 uint16, p1 string) uint8 {
    var v9 int32 = -941
    v10 := GenS1{F0: (0.0 == -2.25)}
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    v3 := []int{}
    v4 := []string{"a"}
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int16 = -32768
        fmt.Println(v12)
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 128, F1: 9, F2: -9223372036854775808}
    _ = len(v3)
    if (127 > 191) {
        v1 += v1
        var v14 uint = 18446744073709551615
        _ = v14
    } else {
        switch "hello" {
        case "a":
            if (-412 < -1) {
                v1 += (v1 / (((v1 ^ (v1 + 127))) | 1))
            } else {
                if false {
                    if false {
                        _ = len(v3)
                    } else {
                        v15 := GenS3{F0: 1, F1: 42}
                        for fi16 := 0; fi16 < 5; fi16++ {
                            if ((1.5 <= 1e10) || (42 == 1)) {
                                var v17 float64 = (3.14159 - 1.5)
                                switch fi16 {
                                case -235:
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v17, v17)
                                    case "foo bar":
                                        v18 := GenS4{F0: "a", F1: false}
                                        _ = v18
                                    case "hello":
                                        var v19 int16 = 1
                                        _ = v19
                                    }
                                default:
                                    v20 := GenS1{F0: true}
                                    _ = v20
                                }
                                fmt.Println(v1, fi16)
                                _ = v17
                            }
                            _ = len(v2)
                            v21 := GenS3{F0: 903, F1: 2147483648}
                            _ = fi16
                            _ = v21
                        }
                        _ = v15
                    }
                    v1 = -41
                }
                v22 := GenS5{F0: (^(v1))}
                v1 = v1
                _ = v22
            }
        case "\n\t":
            var v23 int = -1
            _ = v23
        default:
            fmt.Println(v1)
        }
        var v24 float64 = -2.25
        var v25 float64 = v24
        _ = v24
        _ = v25
    }
    var v26 bool = false
    v1 -= v1
    v1 += v1
    var v27 bool = v26
    v26 = false
    var v28 uint64 = 18446744073709551615
    fmt.Println(v27, v27, v28)
    if false {
        fmt.Println(v1, v1)
        v26 = v27
    }
    if true {
        v29 := v2[""]
        _ = v29
    } else {
        v30 := GenS5{F0: (^(v1))}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
