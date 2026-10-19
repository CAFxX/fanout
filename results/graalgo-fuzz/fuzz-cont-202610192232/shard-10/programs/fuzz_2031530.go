package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 string; F1 bool; F2 int64 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int8; F1 int; F2 uint8 }
func genFn0() int8 {
    var v6 uint8 = 15
    if false {
        for fi7 := 0; fi7 < 20; fi7++ {
            v6 += 1
            v8 := []int{fi7, (fi7 * (fi7 >> ((18446744073709551615) & 63)))}
            _ = fi7
            _ = v8
        }
    } else {
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint8 = 0
                _ = fi9
                _ = v10
            }
        } else {
            v6 = (v6 * v6)
        }
        if false {
            var v11 int16 = -272
            _ = v11
        } else {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 float64 = 1e10
                v6 = (v6 + v6)
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 0, F1: 129}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if ((v6 - 0) > v6) {
            v6 = (^(v6))
            v16 := GenS2{F0: ("hello" + "a"), F1: 42, F2: 16}
            for fi17 := 0; fi17 < 17; fi17++ {
                v6 += 180
                v18 := GenS3{F0: 2147483647}
                v6 = 255
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = fi15
    }
    if ((!((0.1 >= float32(0.1)) && ((!(42 > -675)) && (65535 == 4294967295)))) && true) {
        var v19 int32 = -2147483648
        _ = v19
    }
    _ = v6
    return -5
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = (1e10 * 0.1)
    v3 := []uint32{42, 729, 2147483647}
    var v4 uint16 = 1
    v5 := make(map[uint32]bool)
    var v20 string = ("" + "hello")
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 = v1
        v4 += (v4 + 31)
        if true {
            for fi22 := 0; fi22 < 15; fi22++ {
                if false {
                    var v23 float64 = ((1.5 / (0.1 * (0.1 / 0.1))) / ((float64((v2 * v2)) * 1e10) - float64(float32(float64(((v2 + (((float32((((1e10 / float64(0.1)) * 1.5) * 1e10)) + v2) + v2) - (((v2 + 3.14159) + v2) - v2))) * 0.1))))))
                    v1 -= v1
                    _ = v23
                } else {
                    if false {
                        v2 -= ((-2.25 / v2) - (-2.25 / v2))
                        v20 = ("" + "\n\t")
                        var v24 int8 = 1
                        _ = v24
                    } else {
                        v25 := GenS4{F0: -1}
                        if true {
                            switch fi21 {
                            case 127:
                                v1 += 233
                            case 555:
                                v26 := GenS5{F0: 255, F1: 606}
                                _ = v26
                            case 2147483647:
                                v2 -= (v2 + (1e10 * v2))
                            }
                        } else {
                            v27 := GenS6{F0: v20, F1: ("a" != "a"), F2: 1}
                            var v28 int = (^(fi21))
                            var v29 uint16 = v4
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = v25
                    }
                    switch (v20 + v20) {
                    case "\n\t":
                        var v30 int8 = 36
                        _ = v30
                    case "hello":
                        var v31 int32 = -221
                        _ = v31
                    case "foo bar":
                        fmt.Println(fi21, fi22, fi21)
                    default:
                        if true {
                            fmt.Println(v1, v4)
                            switch fi22 {
                            case 993:
                                delete(v5, 42)
                            case -1:
                                switch (^(fi21)) {
                                case 127:
                                    fmt.Println(v2)
                                    var v32 int64 = 886
                                    _ = v32
                                case -541:
                                    _ = len(v5)
                                case 9223372036854775807:
                                    if (!(!false)) {
                                        if (4294967295 != 2) {
                                            if false {
                                                if len(v3) > 0 {
                                                    v3[0] = 76
                                                }
                                                fmt.Println(v20, v20, v20)
                                                v1 += -32768
                                                if len(v3) > 0 {
                                                    v3[0] = 42
                                                }
                                                var v33 int32 = -360
                                                switch "hello" {
                                                case "\n\t":
                                                    var v34 int8 = genFn0()
                                                    _ = v34
                                                case "a":
                                                    var v35 uint = 75
                                                    _ = v35
                                                }
                                                _ = v33
                                            } else {
                                                _ = len(v3)
                                            }
                                        }
                                    }
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 0
                                }
                            }
                            fmt.Println(fi22)
                            fmt.Println(fi21, v1)
                        } else {
                            var v36 uint16 = (v4 & v4)
                            _ = v36
                        }
                    }
                    var v37 float64 = 0.0
                    _ = v37
                }
                v2 += v2
                _ = fi22
            }
        } else {
            if (!(v1 > (^(v1)))) {
                if false {
                    for fi38 := 0; fi38 < 19; fi38++ {
                        v39 := v5[2147483647]
                        _ = fi38
                        _ = v39
                    }
                    if (true || (15 == -2147483648)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (95 >= 2147483647) {
                            fmt.Println(v4)
                            var v40 float64 = 1.5
                            _ = len(v3)
                            _ = v40
                        }
                    } else {
                        if false {
                            var v41 int = (^(fi21))
                            var v42 uint = 2
                            _ = v41
                            _ = v42
                        } else {
                            var v43 uint32 = 1
                            for fi44 := 0; fi44 < 12; fi44++ {
                                v43 = 2
                                switch v20 {
                                case "\n\t":
                                    v45 := GenS7{F0: -362, F1: 127}
                                    _ = v45
                                }
                                _ = fi44
                            }
                            _ = v43
                        }
                    }
                    var v46 string = "foo bar"
                    _ = v46
                } else {
                    v3 = append(v3, 2147483648)
                    switch fi21 {
                    case 2147483647:
                        fmt.Println(fi21)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = len(v5)
                }
            } else {
                v2 += (v2 / 1.0)
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                if (v20 <= "foo bar") {
                    v20 = v20
                } else {
                    if false {
                        if true {
                            v20 = "\n\t"
                        } else {
                            v2 += (v2 * v2)
                        }
                    }
                    v4 = v4
                }
                _ = fi47
            }
            for fi48 := 0; fi48 < 20; fi48++ {
                var v49 uint32 = 1
                _ = len(v3)
                _ = fi48
                _ = v49
            }
        }
        _ = fi21
    }
    fmt.Println(v4, v20)
    v50 := GenS8{F0: 128, F1: 0}
    v51 := GenS9{F0: 127, F1: -279, F2: 1}
    fmt.Println(v2, v4)
    var v52 int16 = (v1 | int16(v4))
    v20 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
