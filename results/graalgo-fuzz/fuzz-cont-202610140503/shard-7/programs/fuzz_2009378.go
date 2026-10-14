package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 bool; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0(p0 uint, p1 uint, p2 int8) uint32 {
    p0 = (p1 & (p1 % ((p1) | 1)))
    p0 = p1
    v7 := GenS1{F0: 158, F1: 410, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 154
}

func main() {
    var v1 uint64 = 0
    var v2 int = 4294967295
    v3 := []uint32{}
    var v4 int64 = -338
    v5 := []string{"hello", "foo bar"}
    v6 := make(map[uint32]int)
    for fi8 := 0; fi8 < 11; fi8++ {
        v4 -= v4
        for fi9 := 0; fi9 < 12; fi9++ {
            if (!true) {
                if false {
                    if (((((float64(1e10) + 0.0) * 1e10) + 0.0) <= ((0.0 - 1e10) - 0.1)) && (true && (881 != 778))) {
                        v4 = (-(v4))
                        if (false && false) {
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    v4 += (v4 ^ (v4 % ((v4) | 1)))
                                    var v10 uint8 = 255
                                    _ = v10
                                } else {
                                    v4 = v4
                                    v11 := GenS2{F0: genFn0(42, 1, -118), F1: 622}
                                    _ = v11
                                }
                            case "\n\t":
                                v12 := v6[546]
                                _ = v12
                            case "hello":
                                var v13 float64 = float64((0.1 - (float32(((1e10 / (3.14159 - (0.0 / 1e10))) * (0.1 * float64(0.1)))) / 3.14159)))
                                _ = v13
                            }
                            var v14 uint8 = 42
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                            fmt.Println(fi9)
                            v15 := GenS3{F0: 0, F1: v14, F2: 116}
                            v6[162] = fi8
                            if true {
                                var v16 bool = (false && (false || true))
                                if ((32767 > 373) && true) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v4 += int64(v14)
                                    case "hello":
                                        fmt.Println(v4, fi8, v4)
                                    default:
                                        var v17 int = (-(v2))
                                        _ = v17
                                    }
                                    v18 := GenS1{F0: v14, F1: fi9, F2: v1}
                                    if false {
                                        var v19 bool = (-41 > -1)
                                        _ = v19
                                    } else {
                                        var v20 int64 = v4
                                        if (!false) {
                                            switch ((("" + "foo bar") + "\n\t") + "") {
                                            case "\n\t":
                                                v5 = append(v5, ("hello" + "hello"))
                                            default:
                                                v2 = fi8
                                            }
                                        }
                                        fmt.Println(v1, v4, v2)
                                        v21 := GenS3{F0: -32768, F1: (v14 >> ((4294967295) & 63)), F2: 82}
                                        var v22 int32 = 42
                                        if v16 {
                                            var v23 uint32 = genFn0(480, 65535, -128)
                                            var v24 uint64 = 18446744073709551615
                                            var v25 uint = 1
                                            v26 := GenS1{F0: uint8(v24), F1: 654, F2: v24}
                                            fmt.Println(v2)
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        } else {
                                            v27 := v6[124]
                                            v28 := GenS2{F0: 42, F1: 42}
                                            var v29 uint16 = 42
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        if (1.5 != -2.25) {
                                            fmt.Println(fi8, v22)
                                            var v30 uint64 = v1
                                            if (v30 > v30) {
                                                switch (v2 | v2) {
                                                case 0:
                                                    var v31 int8 = 100
                                                    v2 -= fi9
                                                    var v32 int8 = int8(v20)
                                                    v20 += 1
                                                    v32 -= v31
                                                    _ = v31
                                                    _ = v32
                                                case 2147483647:
                                                    var v33 int64 = v4
                                                    _ = v33
                                                }
                                            }
                                            _ = v30
                                        }
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    fmt.Println(fi9)
                                    if false {
                                        v34 := GenS1{F0: 95, F1: v2, F2: (^(v1))}
                                        v2 = (v2 * fi9)
                                        v1 -= 0
                                        _ = v34
                                    } else {
                                        fmt.Println(v14)
                                        var v35 uint = 582
                                        if v16 {
                                            v36 := GenS4{F0: "hello"}
                                            var v37 uint8 = (^(v14))
                                            v38 := GenS4{F0: "foo bar"}
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        } else {
                                            if v16 {
                                                v2 = fi8
                                            }
                                            v35 = (v35 >> ((v35) & 63))
                                            switch fi8 {
                                            case 32767:
                                                delete(v6, 42)
                                            case 721:
                                                var v39 uint16 = 32767
                                                _ = v39
                                            }
                                            if ((42 >= -28) && v16) {
                                                var v40 uint8 = (v14 + v14)
                                                fmt.Println(v16)
                                                v40 -= (v40 * 11)
                                                _ = v40
                                            }
                                        }
                                        _ = v35
                                    }
                                    _ = v18
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v16 = v16
                                case "":
                                    fmt.Println(v4, v1, v1)
                                }
                                _ = v16
                            } else {
                                v4 = (v4 + 1)
                                if false {
                                    _ = len(v6)
                                    v4 = (v4 * -1)
                                    var v41 float64 = ((1e10 * 1e10) - 0.1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v41 -= (float64(1e10) - 3.14159)
                                    _ = v41
                                }
                            }
                            _ = v14
                            _ = v15
                        } else {
                            v2 += fi8
                            switch "foo bar" {
                            case "a":
                                v6[42] = (fi9 * v2)
                            case "foo bar":
                                v2 = fi9
                            case "\n\t":
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v4, fi8)
                                case "":
                                    var v42 float64 = 1e10
                                    _ = v42
                                }
                            }
                        }
                    } else {
                        v2 = (fi8 / ((164) | 1))
                    }
                    var v43 bool = true
                    v4 -= (v4 * (v4 ^ (v4 + v4)))
                    var v44 int64 = (^(v4))
                    _ = v43
                    _ = v44
                }
                var v45 float64 = 0.0
                if (false && true) {
                    v46 := GenS5{F0: 1, F1: (genFn0(18446744073709551615, 255, 118) <= 800), F2: (v4 | 9223372036854775807)}
                    _ = v46
                } else {
                    var v47 uint16 = 32767
                    _ = v47
                }
                _ = v45
            } else {
                var v48 bool = true
                _ = v48
            }
            var v49 uint = 255
            var v50 int8 = 1
            _ = fi9
            _ = v49
            _ = v50
        }
        _ = fi8
    }
    fmt.Println(v2, v1, v2)
    switch v2 {
    case 4294967295:
        for fi51 := 0; fi51 < 13; fi51++ {
            for fi52 := 0; fi52 < 1; fi52++ {
                fmt.Println(fi51)
                v53 := GenS5{F0: 1, F1: true, F2: v4}
                _ = fi52
                _ = v53
            }
            _ = len(v6)
            _ = fi51
        }
    case -2147483648:
        fmt.Println(v1, v4)
    case -9223372036854775808:
        delete(v6, genFn0(42, 757, 42))
    default:
        var v54 int8 = -1
        _ = v54
    }
    v55 := GenS6{F0: genFn0(2, 65535, 127), F1: 826}
    v2 -= (^(v2))
    v1 = (v1 + v1)
    v56 := v6[4294967295]
    var v57 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
