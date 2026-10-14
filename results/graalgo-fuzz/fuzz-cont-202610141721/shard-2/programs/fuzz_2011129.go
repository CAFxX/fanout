package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 bool }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 string; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint64 }
func genFn0(p0 bool) int8 {
    v4 := GenS1{F0: 0, F1: 255, F2: p0}
    var v5 int16 = 1
    v6 := GenS1{F0: 2147483647, F1: 262, F2: (!(1e10 < 1e10))}
    if p0 {
        var v7 bool = ((v5 | 0) != v5)
        if (!true) {
            if (!false) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v5 = (^(v5))
                    var v9 uint8 = 128
                    v10 := []int{(fi8 % (((-(fi8))) | 1)), fi8}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            } else {
                v11 := GenS1{F0: 1, F1: 2, F2: (-2.25 != 1.5)}
                if v7 {
                    var v12 int64 = -560
                    _ = v12
                } else {
                    v5 = (v5 << ((65535) & 63))
                    v13 := GenS1{F0: 489, F1: 869, F2: false}
                    var v14 uint64 = 695
                    _ = v13
                    _ = v14
                }
                p0 = true
                _ = v11
            }
            fmt.Println(p0, v5, v5)
        }
        _ = v7
    }
    if (0 >= 127) {
        v15 := []int{}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                v18 := GenS1{F0: 1, F1: 2, F2: (p0 || false)}
                v19 := []int{fi16, (fi17 % (((^(fi17))) | 1)), 42}
                p0 = ((p0 && false) || p0)
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int16 = v5
            _ = fi16
            _ = v20
        }
        if (!((^(v5)) != v5)) {
            p0 = true
            v21 := []int{32767, -2147483648, 232, 1}
            v5 -= v5
            _ = v21
        }
        if (!true) {
            v22 := []int{-9223372036854775808, 32767, -447}
            v23 := GenS1{F0: 1, F1: 701, F2: p0}
            switch "\n\t" {
            case "\n\t":
                if (-338 > -842) {
                    if p0 {
                        var v24 string = ("\n\t" + (("hello" + "a") + "foo bar"))
                        if p0 {
                            switch -589 {
                            case 2147483647:
                                if true {
                                    if ((-2.25 * float32(3.14159)) != 0.0) {
                                        var v25 float32 = 3.14159
                                        _ = v25
                                    } else {
                                        switch 127 {
                                        case 944:
                                            var v26 int64 = 9223372036854775807
                                            _ = v26
                                        case 32768:
                                            p0 = (32768 <= 42)
                                        case 9223372036854775807:
                                            var v27 uint16 = 146
                                            _ = v27
                                        default:
                                            v24 = v24
                                        }
                                        _ = len(v22)
                                        var v28 int32 = 1
                                        _ = v28
                                    }
                                    fmt.Println(v24, v5)
                                }
                            default:
                                for fi29 := 0; fi29 < 1; fi29++ {
                                    fmt.Println(v5, v5)
                                    _ = fi29
                                }
                            }
                        }
                        _ = v24
                    } else {
                        fmt.Println(v5, v5)
                        for fi30 := 0; fi30 < 3; fi30++ {
                            p0 = (2 > 127)
                            _ = fi30
                        }
                        _ = len(v22)
                    }
                    var v31 uint64 = 18446744073709551615
                    v5 -= v5
                    _ = v31
                }
            case "a":
                for fi32 := 0; fi32 < 13; fi32++ {
                    p0 = p0
                    var v33 int16 = v5
                    _ = fi32
                    _ = v33
                }
            case "":
                var v34 int8 = 42
                _ = v34
            default:
                if false {
                    var v35 bool = (((0.0 - 1.5) - 0.0) == (1.5 - 1e10))
                    v36 := GenS2{F0: (p0 || false), F1: 2147483647}
                    _ = v35
                    _ = v36
                }
            }
            _ = v22
            _ = v23
        }
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -128
}
func genFn1() uint8 {
    if (127 < 100) {
        var v37 uint64 = 0
        fmt.Println(v37)
        _ = v37
    }
    switch 201 {
    case 127:
        var v38 uint8 = 0
        _ = v38
    }
    if ((float32(1e10) - 0.0) >= 3.14159) {
        var v39 int64 = -1
        _ = v39
    }
    return 136
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = (float64(float32(0.1)) * (1e10 * 3.14159))
    var v3 int16 = -1
    v40 := GenS2{F0: (!false), F1: 151}
    v41 := []int{1, 32767, 42}
    switch -9223372036854775808 {
    case -268:
        var v42 uint64 = 18446744073709551615
        _ = v42
    case -579:
        var v43 float32 = (0.1 * float32((v2 + (((v2 / v2) * v2) + v2))))
        _ = v43
    default:
        var v44 int = 42
        _ = v44
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 float32 = -2.25
        _ = fi45
        _ = v46
    }
    v47 := GenS3{F0: 42}
    var v48 int16 = v3
    switch 2147483647 {
    case -58:
        for fi49 := 0; fi49 < 9; fi49++ {
            v50 := GenS4{F0: 768, F1: ((("foo bar" + ("foo bar" + "hello")) + ("foo bar" + "")) + "\n\t"), F2: "hello"}
            v51 := GenS1{F0: 2147483647, F1: 32767, F2: true}
            var v52 float64 = v2
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    default:
        v53 := GenS2{F0: ("foo bar" <= ""), F1: 1}
        _ = v53
    }
    v48 = (v3 * v3)
    var v54 int32 = 1
    switch "" {
    case "\n\t":
        if len(v41) > 0 {
            v41[0] = -797
        }
    }
    fmt.Println(v2, v1)
    v55 := GenS5{F0: ((42 > 1) && true)}
    v56 := GenS6{F0: ("foo bar" + ""), F1: 42}
    var v57 int16 = v3
    v2 += 1.5
    for fi58 := 0; fi58 < 8; fi58++ {
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
        _ = len(v41)
        if (((v2 / v2) * -2.25) != (v2 - (((v2 / (v2 * v2)) / ((float64(float32(v2)) + (((v2 - ((v2 * (float64(1.5) - (v2 * 1e10))) + v2)) - 0.1) - v2)) / float64(((3.14159 - 1e10) / ((0.0 + 1.5) / ((1e10 - -2.25) - (-2.25 / 1e10))))))) + float64(-2.25)))) {
            fmt.Println(v54)
        } else {
            var v59 uint64 = 0
            var v60 uint8 = 2
            _ = v59
            _ = v60
        }
        _ = fi58
    }
    for fi61 := 0; fi61 < 17; fi61++ {
        var v62 float64 = (float64(1e10) / (3.14159 / (v2 + v2)))
        v48 = v3
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v48:", v48)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v47
    _ = v48
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
