package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int; F1 int64 }
func genFn0() uint8 {
    var v6 uint = 832
    v6 -= v6
    switch 0 {
    case -321:
        v6 += v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{0, (fi8 >> ((v6) & 63)), fi7, (fi7 / ((fi7) | 1))}
            if (34 == 1) {
                if (813 != 893) {
                    if (true || true) {
                        v6 += (v6 | (v6 + v6))
                        var v10 uint8 = 42
                        switch "" {
                        case "":
                            var v11 float32 = (-2.25 + (((float32(((float64((1e10 * (1e10 + 1e10))) * float64((((0.1 / 1e10) / 1.0) + 3.14159))) + 1.5)) * (-2.25 * 1.5)) / (0.1 - ((0.1 / 1.0) / float32(1.5)))) * (3.14159 / 3.14159)))
                            _ = v11
                        case "foo bar":
                            fmt.Println(v10, fi7)
                        }
                        _ = v10
                    }
                    v6 -= (v6 - (^(v6)))
                    v6 += (v6 % ((v6) | 1))
                }
                fmt.Println(fi8)
                var v12 int32 = 1
                _ = v12
            }
            var v13 bool = false
            _ = fi8
            _ = v13
            _ = v9
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int32 = 0
            var v16 int8 = 127
            switch fi7 {
            case 32767:
                var v17 float32 = (-2.25 + 3.14159)
                _ = v17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        v6 -= (v6 ^ (v6 ^ (v6 >> ((v6) & 63))))
        _ = fi7
    }
    _ = v6
    return 194
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    var v3 int16 = 32767
    var v4 uint64 = 4294967295
    var v5 uint = 488
    fmt.Println(v5, v4, v3)
    fmt.Println(v5, v4, v3)
    if false {
        v5 = 18446744073709551615
        if false {
            var v18 uint8 = 2
            _ = len(v1)
            if (!true) {
                fmt.Println(v4, v2)
                fmt.Println(v2, v2, v5)
                v19 := []int{}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS1{F0: 4294967295, F1: 2147483648, F2: false}
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v23 := GenS2{F0: 2147483648}
                        if ((-(v3)) >= v3) {
                            v24 := GenS3{F0: 457, F1: 1}
                            v25 := GenS4{F0: 0, F1: int64(v4)}
                            if true {
                                _ = len(v1)
                                switch "hello" {
                                case "":
                                    v5 -= uint(v4)
                                case "hello":
                                    switch ((v2 + (v2 + v2)) + "foo bar") {
                                    case "a":
                                        v2 = (v2 + ((v2 + v2) + v2))
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                            } else {
                                v4 = uint64(v5)
                            }
                            _ = v24
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                    for fi26 := 0; fi26 < 17; fi26++ {
                        var v27 float32 = (((0.0 + (float32((1e10 / 0.1)) + -2.25)) + 0.0) * (-2.25 * (((0.0 / float32(3.14159)) * 0.0) / (float32((0.1 + 0.0)) * 0.1))))
                        _ = fi26
                        _ = v27
                    }
                    _ = fi20
                    _ = v21
                }
            }
            _ = v18
        } else {
            v5 = 65535
        }
    } else {
        v28 := []int{2147483647}
        _ = v28
    }
    var v29 uint16 = 42
    fmt.Println(v29, v2)
    v30 := GenS4{F0: 206, F1: -1}
    switch ("hello" + (v2 + ("hello" + "\n\t"))) {
    case "foo bar":
        var v31 uint32 = 4294967295
        _ = v31
    }
    v32 := GenS2{F0: 65535}
    var v33 uint32 = 985
    switch 1 {
    case 9223372036854775807:
        v34 := []int{763, 9223372036854775807, int(v4), 4294967295}
        _ = v34
    default:
        fmt.Println(v29, v33)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
