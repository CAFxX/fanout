package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 bool; F2 int }
type GenS7 struct { F0 bool }
func genFn0(p0 uint64) int16 {
    fmt.Println(p0)
    if (255 == 766) {
        fmt.Println(p0)
    }
    _ = p0
    return 0
}
func genFn1(p0 int64) uint64 {
    v4 := GenS1{F0: -2147483648, F1: int(p0), F2: 167}
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 = 1
        v6 := GenS1{F0: fi5, F1: (-(fi5)), F2: 42}
        p0 = p0
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 2
}

func main() {
    var v1 string = ("" + ("\n\t" + ((("a" + ("foo bar" + "a")) + (((("hello" + "") + "a") + "foo bar") + "")) + "a")))
    var v2 int = 856
    var v3 uint64 = 2
    v1 = "\n\t"
    v7 := []int{(v2 >> ((255) & 63)), 309, (-(v2))}
    v2 -= -2147483648
    if (29 >= 18446744073709551615) {
        v7 = append(v7, (v2 >> ((42) & 63)))
        var v8 bool = true
        v9 := GenS2{F0: 0, F1: -443}
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 127
        _ = v10
    }
    var v11 int16 = -32768
    for fi12 := 0; fi12 < 15; fi12++ {
        v2 = fi12
        _ = fi12
    }
    v2 = (v2 - (^(v2)))
    var v13 int8 = -1
    var v14 int32 = 0
    if true {
        var v15 uint = 4294967295
        var v16 float32 = float32(-2.25)
        for fi17 := 0; fi17 < 16; fi17++ {
            if len(v7) > 0 {
                v7[0] = (v2 / ((fi17) | 1))
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    } else {
        var v18 uint8 = 41
        var v19 uint = 262
        fmt.Println(v2, v2)
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 float32 = ((0.1 - -2.25) / 1e10)
        if false {
            switch fi20 {
            case 42:
                v14 = (v14 + (v14 << ((312) & 63)))
            case 906:
                switch fi20 {
                case 159:
                    v22 := GenS3{F0: 0}
                    _ = v22
                case 182:
                    switch v2 {
                    case 0:
                        var v23 bool = true
                        _ = v23
                    case 127:
                        fmt.Println(v2)
                    default:
                        if (false || false) {
                            var v24 int16 = genFn0((v3 | 18446744073709551615))
                            switch v2 {
                            case 32768:
                                fmt.Println(fi20, v13)
                            case 2147483647:
                                v13 += v13
                            case 0:
                                for fi25 := 0; fi25 < 3; fi25++ {
                                    if true {
                                        if (((0.1 / 0.1) > 3.14159) && false) {
                                            v26 := GenS3{F0: 127}
                                            _ = v26
                                        } else {
                                            fmt.Println(fi25, fi20, v24)
                                            v7 = append(v7, 0)
                                        }
                                        if false {
                                            v14 -= (v14 % ((v14) | 1))
                                        }
                                    }
                                    if (255 < 127) {
                                        var v27 float32 = (v21 - v21)
                                        var v28 uint8 = 1
                                        _ = v27
                                        _ = v28
                                    } else {
                                        if (((65535 >= 65535) && false) && false) {
                                            if false {
                                                var v29 int64 = 1
                                                _ = v29
                                            } else {
                                                if true {
                                                    v11 = genFn0(18446744073709551615)
                                                    var v30 int = fi25
                                                    v14 = (v14 & (v14 - (v14 >> ((18446744073709551615) & 63))))
                                                    _ = v30
                                                }
                                                v31 := GenS1{F0: (fi25 | (v2 + -9223372036854775808)), F1: 9223372036854775807, F2: 217}
                                                fmt.Println(v13, v24)
                                                _ = len(v7)
                                                v32 := GenS3{F0: 42}
                                                _ = v31
                                                _ = v32
                                            }
                                            var v33 float32 = 1.5
                                            _ = v33
                                        } else {
                                            v24 = genFn0((v3 << ((42) & 63)))
                                            v11 = -1
                                            fmt.Println(v13, fi25)
                                        }
                                    }
                                    v34 := GenS1{F0: -115, F1: (fi20 ^ 0), F2: 232}
                                    var v35 string = "foo bar"
                                    if false {
                                        _ = len(v7)
                                    } else {
                                        v1 = "a"
                                        var v36 int16 = v24
                                        if false {
                                            v37 := GenS2{F0: 572, F1: 728}
                                            var v38 float64 = 1.5
                                            if (v38 > ((((0.1 / v38) + 0.1) + (float64((float32(float64((-2.25 - v21))) * v21)) + v38)) / (((3.14159 - 0.1) - float64(-2.25)) * 1.5))) {
                                                v39 := GenS4{F0: 0, F1: v14, F2: 65535}
                                                v40 := GenS5{F0: 127}
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v13 += (v13 | (v13 | 0))
                                                if len(v7) > 0 {
                                                    v7[0] = v2
                                                }
                                                if false {
                                                    if ((!true) || true) {
                                                        switch 32768 {
                                                        case 32768:
                                                            var v41 float32 = (0.1 * v21)
                                                            v21 = (0.1 - v41)
                                                            _ = v41
                                                        default:
                                                            v11 -= (v24 % (((v11 & (v36 & (v36 - genFn0(4294967295))))) | 1))
                                                            if (uint16(v3) <= 255) {
                                                                if false {
                                                                    var v42 uint = 913
                                                                    fmt.Println(v14, v11, v38)
                                                                    _ = v42
                                                                } else {
                                                                    if ((v21 * float32(v38)) == -2.25) {
                                                                        fmt.Println(v36, v24)
                                                                    } else {
                                                                        v36 = -144
                                                                    }
                                                                    v43 := GenS6{F0: (v2 ^ (fi25 << ((0) & 63))), F1: true, F2: (-(fi25))}
                                                                    v2 += (fi20 << ((452) & 63))
                                                                    v44 := GenS1{F0: (v2 * (^(v2))), F1: 32767, F2: 140}
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            } else {
                                                                var v45 int64 = 42
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                _ = v45
                                                            }
                                                            fmt.Println(v35, v21)
                                                        }
                                                        v46 := GenS1{F0: int(v3), F1: fi25, F2: 70}
                                                        _ = v46
                                                    }
                                                    var v47 int8 = 127
                                                    _ = v47
                                                }
                                            }
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v48 := GenS5{F0: 255}
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            if false {
                                                switch fi20 {
                                                case 32767:
                                                    v35 = v1
                                                case 0:
                                                    if true {
                                                        var v49 string = "\n\t"
                                                        v35 = v1
                                                        _ = v49
                                                    } else {
                                                        if true {
                                                            switch v35 {
                                                            case "foo bar":
                                                                v24 += (v24 ^ (^(v36)))
                                                            case "a":
                                                                fmt.Println(v13, v13, v21)
                                                            }
                                                        } else {
                                                            v50 := GenS3{F0: 194}
                                                            fmt.Println(fi20, v21, v14)
                                                            _ = v50
                                                        }
                                                    }
                                                case 127:
                                                    if true {
                                                        v51 := GenS1{F0: v2, F1: -730, F2: 247}
                                                        var v52 uint8 = 1
                                                        fmt.Println(v11)
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                                if true {
                                                    v14 += v14
                                                    v36 = (v36 & v36)
                                                }
                                                v36 += 1
                                            }
                                            v2 += v2
                                            var v53 uint32 = 42
                                            _ = v48
                                            _ = v53
                                        }
                                        _ = v36
                                    }
                                    _ = fi25
                                    _ = v34
                                    _ = v35
                                }
                            default:
                                v54 := GenS7{F0: (genFn0((^(v3))) >= (v24 * (v24 * v24)))}
                                _ = v54
                            }
                            v13 = v13
                            _ = v24
                        } else {
                            fmt.Println(v3, v21)
                        }
                    }
                default:
                    v13 -= 1
                }
            }
        } else {
            for fi55 := 0; fi55 < 14; fi55++ {
                v21 -= v21
                fmt.Println(fi55)
                _ = fi55
            }
        }
        v11 = (v11 * (v11 << ((42) & 63)))
        _ = fi20
        _ = v21
    }
    v11 = 0
    v3 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
