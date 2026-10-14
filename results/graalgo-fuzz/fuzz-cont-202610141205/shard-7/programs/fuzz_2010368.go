package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int64 = 963
    var v2 uint64 = 4294967295
    var v3 uint32 = 2
    var v4 float32 = (float32((float64(float32(-2.25)) * 3.14159)) * float32(0.1))
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: -128, F1: false, F2: 32767}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int32 = -1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    v5 -= 128
    v12 := GenS1{F0: 42, F1: false, F2: 1}
    switch 216 {
    case -2147483648:
        v13 := GenS2{F0: -1, F1: v1}
        _ = v13
    case 32768:
        v2 += v2
    case 766:
        v14 := GenS3{F0: 1, F1: (^(v3)), F2: v2}
        _ = v14
    default:
        var v15 string = "\n\t"
        _ = v15
    }
    if true {
        if (-43 > 127) {
            if (!true) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v1 += v1
                    v1 = (v1 / (((v1 - v1)) | 1))
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS3{F0: -359, F1: v3, F2: (v2 + v2)}
                        v3 = (v3 - v3)
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                if false {
                    var v19 int16 = -547
                    var v20 int8 = 127
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 float64 = ((float64(1e10) + 0.1) + ((0.0 - (0.0 + (-2.25 / ((0.1 * 3.14159) + float64(v4))))) - 1.5))
                        _ = fi21
                        _ = v22
                    }
                    _ = v19
                    _ = v20
                }
            }
        } else {
            v23 := []int{32768, 42, 0}
            fmt.Println(v3, v4, v1)
            v2 -= 18446744073709551615
            _ = v23
        }
    } else {
        v3 += (v3 + (v3 >> ((1) & 63)))
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        fmt.Println(v4, v5, v4)
        v3 = v3
        _ = fi24
    }
    if true {
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := []int{(fi25 ^ fi25), (fi25 ^ (fi25 * (fi25 & 2147483647))), (fi25 & (^(fi25))), 1}
            if false {
                switch "a" {
                case "a":
                    var v27 bool = (true && ((v1 | int64(v2)) != v1))
                    _ = v27
                case "foo bar":
                    v3 -= v3
                }
                v28 := GenS4{F0: -128, F1: false, F2: 0}
                v29 := GenS1{F0: 100, F1: true, F2: 0}
                _ = v28
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        fmt.Println(v4)
        v1 += (v1 * 790)
    } else {
        if false {
            v1 = v1
        }
    }
    v30 := GenS3{F0: -1, F1: v3, F2: v2}
    v1 = (v1 + -9223372036854775808)
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 int16 = 1
        if false {
            switch ("" + "hello") {
            case "hello":
                for fi33 := 0; fi33 < 4; fi33++ {
                    v1 -= (v1 << ((358) & 63))
                    var v34 bool = false
                    var v35 float64 = 0.1
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case "a":
                for fi36 := 0; fi36 < 8; fi36++ {
                    if (65535 >= 616) {
                        v37 := []int{127}
                        _ = v37
                    } else {
                        var v38 uint16 = 922
                        _ = v38
                    }
                    v39 := []int{}
                    _ = fi36
                    _ = v39
                }
            case "foo bar":
                var v40 uint32 = 4294967295
                _ = v40
            }
        } else {
            if (("foo bar" >= ("a" + ("hello" + (((("" + ("" + ("foo bar" + ("hello" + "hello")))) + "") + "") + ("" + ("foo bar" + ("" + ("hello" + "\n\t")))))))) && true) {
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 string = "\n\t"
                    v43 := []int{}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                v5 -= (v5 + 189)
            } else {
                switch "a" {
                case "foo bar":
                    if (65535 < 2) {
                        var v44 string = ""
                        v3 = (v3 - 42)
                        for fi45 := 0; fi45 < 9; fi45++ {
                            fmt.Println(v5, v32, v5)
                            var v46 uint8 = 1
                            _ = fi45
                            _ = v46
                        }
                        _ = v44
                    } else {
                        v47 := GenS1{F0: 127, F1: (!false), F2: 255}
                        v4 = v4
                        for fi48 := 0; fi48 < 1; fi48++ {
                            if (false && true) {
                                v1 -= (v1 << ((4294967295) & 63))
                                if true {
                                    if (!true) {
                                        v49 := []int{(fi48 % ((fi48) | 1)), (fi48 | fi48), (^(fi48))}
                                        fmt.Println(v2)
                                        _ = v49
                                    } else {
                                        fmt.Println(v4, fi31)
                                    }
                                } else {
                                    v5 -= v5
                                    v3 += v3
                                    var v50 uint8 = (v5 + (v5 >> ((247) & 63)))
                                    _ = v50
                                }
                                fmt.Println(v5, v4, fi48)
                            }
                            _ = fi48
                        }
                        _ = v47
                    }
                case "":
                    if (true && (824 > 1)) {
                        v51 := GenS1{F0: 39, F1: ("a" == (("a" + ("a" + "")) + "\n\t")), F2: 1}
                        _ = v51
                    } else {
                        var v52 float32 = float32((1.5 / 1e10))
                        for fi53 := 0; fi53 < 1; fi53++ {
                            var v54 int = fi31
                            v1 = (v1 & (v1 << ((18446744073709551615) & 63)))
                            _ = fi53
                            _ = v54
                        }
                        var v55 uint64 = (v2 + 1)
                        _ = v52
                        _ = v55
                    }
                case "hello":
                    for fi56 := 0; fi56 < 4; fi56++ {
                        v57 := []int{-257, (fi31 * fi31), (fi31 | fi56)}
                        v57 = append(v57, (fi56 >> ((42) & 63)))
                        v32 = (v32 << ((2) & 63))
                        _ = fi56
                        _ = v57
                    }
                default:
                    if ((42 > 941) && true) {
                        v58 := GenS3{F0: 42, F1: (v3 * 1), F2: v2}
                        _ = v58
                    }
                }
                var v59 int32 = -2147483648
                _ = v59
            }
            v60 := GenS5{F0: 653}
            _ = v60
        }
        _ = fi31
        _ = v32
    }
    v61 := GenS5{F0: 421}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v61
}
