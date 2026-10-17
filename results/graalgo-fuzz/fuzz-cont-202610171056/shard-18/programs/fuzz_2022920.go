package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 string; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 uint16) int {
    v5 := GenS1{F0: 42}
    var v6 int = -9223372036854775808
    fmt.Println(v6)
    v7 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v6 ^ (v6 * v6))
}
func genFn1() int32 {
    v8 := GenS1{F0: 0}
    if (114 >= 127) {
        v9 := []int{-938, 2147483647}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int64 = 1
            if (!true) {
                var v12 int = 831
                var v13 int32 = 0
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    if (255 < 588) {
        v14 := []int{128, 128, 2147483647, 527}
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int32 = -1
            switch ("" + "a") {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    v17 := GenS2{F0: 69, F1: "foo bar", F2: true}
                    _ = v17
                default:
                    switch "hello" {
                    case "\n\t":
                        v18 := GenS3{F0: 255}
                        _ = v18
                    case "":
                        var v19 uint8 = 42
                        _ = v19
                    default:
                        v16 = -2147483648
                    }
                }
            case "\n\t":
                switch "" {
                case "hello":
                    switch "a" {
                    case "\n\t":
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            for fi20 := 0; fi20 < 12; fi20++ {
                                var v21 float32 = 0.1
                                _ = fi20
                                _ = v21
                            }
                        case "a":
                            v16 = v16
                        default:
                            if len(v14) > 0 {
                                v14[0] = -180
                            }
                        }
                    case "a":
                        for fi22 := 0; fi22 < 11; fi22++ {
                            if (-43 != 1) {
                                v16 += 1
                                v16 += v16
                                v23 := GenS4{F0: 4294967295, F1: 1, F2: false}
                                _ = v23
                            } else {
                                v16 = (v16 - 2147483647)
                            }
                            _ = fi22
                        }
                    case "foo bar":
                        fmt.Println(v16)
                    }
                case "":
                    if false {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 int32 = (v16 << ((4294967295) & 63))
                            var v26 uint32 = 65535
                            v27 := GenS5{F0: -128}
                            if (208 >= 1) {
                                fmt.Println(v26, fi15, v25)
                                var v28 int64 = 9223372036854775807
                                var v29 float64 = -2.25
                                _ = v28
                                _ = v29
                            }
                            v25 -= v16
                            _ = fi24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        v16 = v16
                    }
                }
            default:
                v30 := GenS4{F0: -695, F1: 252, F2: (42 > 127)}
                _ = v30
            }
            if (!false) {
                if (!(9223372036854775807 == fi15)) {
                    var v31 int32 = -185
                    v31 += (v31 % (((v31 / ((v16) | 1))) | 1))
                    if (!false) {
                        var v32 int = (fi15 | (fi15 | fi15))
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v32 -= (v32 >> ((0) & 63))
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            _ = fi33
                        }
                        _ = v32
                    }
                    _ = v31
                } else {
                    var v34 float32 = float32((1.5 + (float64(1.5) + 1.5)))
                    var v35 uint16 = 255
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 uint32 = 1
                var v37 float32 = 0.0
                _ = v36
                _ = v37
            }
            _ = fi15
            _ = v16
        }
        var v38 float64 = 0.0
        _ = v14
        _ = v38
    }
    v39 := []int{364}
    _ = v39
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (float32(0.0) / (3.14159 + float32(-2.25)))
    var v2 float64 = 0.1
    v3 := []uint32{65535, 65535, 432, 1}
    v4 := make(map[string]string)
    fmt.Println(v1, v1)
    v1 = v1
    var v40 uint = 4294967295
    v3 = append(v3, 42)
    switch "" {
    case "hello":
        switch "\n\t" {
        case "":
            v40 -= v40
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            _ = len(v4)
        default:
            if (true || true) {
                v41 := GenS6{F0: 199, F1: 104}
                var v42 int32 = 0
                _ = v41
                _ = v42
            } else {
                var v43 int64 = 42
                for fi44 := 0; fi44 < 14; fi44++ {
                    var v45 float64 = ((v2 + 0.0) + v2)
                    for fi46 := 0; fi46 < 4; fi46++ {
                        var v47 int = fi44
                        v40 -= 18446744073709551615
                        _ = fi46
                        _ = v47
                    }
                    var v48 uint32 = 93
                    _ = fi44
                    _ = v45
                    _ = v48
                }
                _ = v43
            }
        }
    case "\n\t":
        v40 = 4294967295
    }
    v40 -= v40
    v49 := GenS6{F0: 128, F1: 42}
    v50 := GenS4{F0: 4294967295, F1: 0, F2: true}
    v2 -= v2
    v2 = (0.1 - (float64(3.14159) - -2.25))
    fmt.Println(v40, v40)
    var v51 int = genFn0(v40, 565)
    v52 := GenS6{F0: 2, F1: 288}
    var v53 float64 = (float64(v1) - ((((-2.25 * (v2 / v2)) + 3.14159) - (v2 - (v2 / 1e10))) / float64(1.5)))
    v54 := GenS6{F0: 38, F1: 42}
    var v55 float32 = float32(float64(v1))
    delete(v4, "hello")
    fmt.Println(v53)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
