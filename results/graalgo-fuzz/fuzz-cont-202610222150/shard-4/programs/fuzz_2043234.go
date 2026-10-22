package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int) int32 {
    v4 := []int{p0, 128, -304}
    var v5 uint = 255
    for fi6 := 0; fi6 < 19; fi6++ {
        if len(v4) > 0 {
            v4[0] = (p0 >> ((0) & 63))
        }
        _ = fi6
    }
    v5 = (^(v5))
    _ = p0
    _ = v4
    _ = v5
    return int32(p0)
}
func genFn1(p0 uint8) uint16 {
    if (!(((true && true) && (uint32(p0) == 65535)) || (true || true))) {
        if (974 >= 42) {
            for fi7 := 0; fi7 < 6; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    if false {
                        fmt.Println(fi8, fi8, fi7)
                        var v9 float64 = ((float64(float32((1e10 * 3.14159))) / 3.14159) * (1.5 / 1.0))
                        var v10 int = fi7
                        _ = v10
                        _ = v9
                    } else {
                        var v11 uint64 = 18446744073709551615
                        v12 := GenS1{F0: 1, F1: uint32(v11)}
                        var v13 int32 = 0
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(fi8, fi7, p0)
                    if (42 > 0) {
                        switch (fi7 | 856) {
                        case 2147483647:
                            fmt.Println(fi7, fi7)
                        case 819:
                            p0 = p0
                            switch "a" {
                            case "":
                                p0 += p0
                            case "hello":
                                v14 := GenS1{F0: 0, F1: 2147483648}
                                _ = v14
                            case "\n\t":
                                p0 -= 0
                            default:
                                v15 := GenS1{F0: 2, F1: 2147483647}
                                _ = v15
                            }
                            p0 += 175
                        }
                        v16 := []int{}
                        _ = v16
                    }
                    _ = fi8
                }
                _ = fi7
            }
        }
        switch 42 {
        case 735:
            var v17 int32 = 42
            _ = v17
        default:
            var v18 uint32 = 993
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (!false) {
            v20 := GenS1{F0: 255, F1: 42}
            p0 = (p0 - p0)
            switch "hello" {
            case "":
                fmt.Println(fi19)
            case "hello":
                p0 += 1
            default:
                p0 -= p0
            }
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        p0 = p0
        p0 += (p0 + (p0 - p0))
        v22 := GenS1{F0: 205, F1: 65535}
        _ = fi21
        _ = v22
    }
    p0 = p0
    p0 += p0
    _ = p0
    return 2
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]string)
    var v3 int8 = -10
    v2[2] = "foo bar"
    if (!true) {
        v23 := []int{v1}
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := v2[65535]
        v3 = int8(v1)
        switch ("hello" + ("a" + v25)) {
        case "a":
            if false {
                for fi26 := 0; fi26 < 14; fi26++ {
                    v1 = 578
                    _ = fi26
                }
                v1 += (-(fi24))
            } else {
                v2[236] = v25
            }
        case "foo bar":
            if (float32((0.0 + 0.0)) == float32((-2.25 / ((((1.5 + (float64(((1.5 * (0.1 + 3.14159)) - (1e10 + ((float32((0.0 * 3.14159)) * 0.1) + 1e10)))) - (1.5 * 1e10))) * 3.14159) * (3.14159 / 1.0)) + (0.0 * -2.25))))) {
                if true {
                    var v27 int16 = 0
                    _ = v27
                } else {
                    _ = len(v2)
                }
                var v28 uint8 = 2
                _ = v28
            } else {
                v29 := GenS2{F0: (v3 << ((2) & 63)), F1: 65535, F2: 0}
                var v30 uint32 = 2147483648
                _ = v29
                _ = v30
            }
        case "hello":
            for fi31 := 0; fi31 < 3; fi31++ {
                fmt.Println(v3)
                v32 := GenS3{F0: 1, F1: 45}
                fmt.Println(v3)
                var v33 int32 = -2147483648
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi24
        _ = v25
    }
    var v34 int8 = v3
    v3 -= v34
    v3 += -128
    switch "a" {
    case "foo bar":
        if (!true) {
            for fi35 := 0; fi35 < 17; fi35++ {
                if (!(-2.25 == 1.5)) {
                    v3 = v34
                } else {
                    v34 -= -111
                    v36 := []int{761, fi35}
                    _ = v36
                }
                v37 := []int{}
                _ = fi35
                _ = v37
            }
            v38 := GenS3{F0: 720, F1: 42}
            _ = v38
        } else {
            v39 := GenS1{F0: 1, F1: 175}
            var v40 uint16 = 65535
            v41 := GenS4{F0: (v40 - (v40 >> ((4294967295) & 63)))}
            _ = v39
            _ = v40
            _ = v41
        }
    default:
        if true {
            for fi42 := 0; fi42 < 18; fi42++ {
                for fi43 := 0; fi43 < 8; fi43++ {
                    var v44 float32 = float32(((1e10 + (((1e10 * 0.1) - (((1.5 / (3.14159 + 0.0)) - (float64(1e10) - (3.14159 * (0.0 - 3.14159)))) + 1e10)) * 0.0)) * ((0.0 * 1.5) - 1.5)))
                    v34 += v34
                    _ = fi43
                    _ = v44
                }
                var v45 uint16 = 1
                v46 := []int{v1, (v1 & (-(fi42))), v1}
                _ = fi42
                _ = v45
                _ = v46
            }
            var v47 int64 = 1
            var v48 int = -930
            _ = v47
            _ = v48
        } else {
            if true {
                switch 32768 {
                case 42:
                    if (125 != 42) {
                        switch "" {
                        case "":
                            var v49 int8 = 1
                            _ = v49
                        case "\n\t":
                            v2[2] = "foo bar"
                        default:
                            fmt.Println(v1, v3, v3)
                        }
                    } else {
                        if true {
                            for fi50 := 0; fi50 < 9; fi50++ {
                                v3 = 1
                                _ = fi50
                            }
                        }
                        var v51 int8 = v3
                        var v52 int64 = 9223372036854775807
                        _ = v51
                        _ = v52
                    }
                case -172:
                    var v53 int16 = -478
                    _ = v53
                }
                v54 := []int{}
                var v55 uint = 947
                _ = v54
                _ = v55
            } else {
                v56 := []int{-1, v1, v1}
                _ = v56
            }
        }
    }
    v57 := v2[202]
    v58 := []int{v1, (v1 - v1), v1, v1}
    v59 := GenS5{F0: 42, F1: -2147483648, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v57
    _ = v58
    _ = v59
}
