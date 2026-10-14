package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint32; F1 int }
func genFn0(p0 uint8) int8 {
    var v6 float64 = (0.1 * (((0.1 + (-2.25 - 1e10)) * 0.1) * 0.0))
    fmt.Println(p0, v6)
    v6 += float64((1.5 + float32((v6 * ((v6 / v6) + float64(float32(1.5)))))))
    _ = p0
    _ = v6
    return -1
}
func genFn1() int32 {
    var v7 float32 = 0.1
    v7 -= (v7 * (v7 + 1e10))
    var v8 uint = 829
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 1
    var v2 int32 = -165
    var v3 bool = ((32768 <= 258) || true)
    var v4 uint = 1
    var v5 int32 = (v2 / ((v2) | 1))
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v1, v3, fi9)
        var v10 uint8 = (v1 << ((0) & 63))
        _ = fi9
        _ = v10
    }
    var v11 float64 = float64(1.5)
    if (-2.25 >= (float32(float64(0.0)) / (1.5 * 0.1))) {
        if true {
            for fi12 := 0; fi12 < 9; fi12++ {
                switch -9223372036854775808 {
                case -1:
                    v1 = 42
                case 4294967295:
                    var v13 float32 = 1.5
                    _ = v13
                case -2147483648:
                    v14 := []int{(fi12 / ((fi12) | 1)), (fi12 - (fi12 % ((127) | 1))), (fi12 % ((-2147483648) | 1))}
                    _ = v14
                }
                _ = fi12
            }
        }
    }
    v2 -= (v5 - genFn1())
    switch 4294967295 {
    case 1:
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS1{F0: (-39 <= 9223372036854775807), F1: 255, F2: -1}
            for fi17 := 0; fi17 < 12; fi17++ {
                v4 -= v4
                v18 := []int{627}
                if ((110 != 267) && ((1 <= 255) || v3)) {
                    switch fi17 {
                    case 127:
                        if v3 {
                            var v19 uint64 = 385
                            fmt.Println(fi17, fi15, v2)
                            if ((!v3) || v3) {
                                v4 -= v4
                                var v20 int8 = int8(v2)
                                _ = v20
                            }
                            _ = v19
                        }
                    case 0:
                        v21 := GenS1{F0: v3, F1: 736, F2: int64(v4)}
                        _ = v21
                    case -9223372036854775808:
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    }
                } else {
                    fmt.Println(fi15, v5)
                    var v22 int8 = 42
                    var v23 int = 127
                    v24 := GenS2{F0: (v1 % (((^(v1))) | 1)), F1: 65535, F2: uint16(v1)}
                    var v25 int8 = v22
                    v26 := GenS2{F0: (v1 % ((v1) | 1)), F1: 1, F2: 1}
                    if (true && (0 == 2)) {
                        if true {
                            if true {
                                switch fi15 {
                                case 438:
                                    if v3 {
                                        if (!(!(v5 > (-(v2))))) {
                                            if len(v18) > 0 {
                                                v18[0] = (v23 + v23)
                                            }
                                        } else {
                                            fmt.Println(v2, v22, fi15)
                                            if false {
                                                if (v3 && true) {
                                                    var v27 float64 = v11
                                                    if len(v18) > 0 {
                                                        v18[0] = v23
                                                    }
                                                    _ = v27
                                                }
                                                var v28 uint16 = 42
                                                v3 = v3
                                                _ = v28
                                            }
                                            v29 := GenS3{F0: v1, F1: ""}
                                            _ = v29
                                        }
                                        switch ("foo bar" + ("\n\t" + "a")) {
                                        case "a":
                                            v30 := GenS1{F0: v3, F1: 65535, F2: -688}
                                            _ = v30
                                        case "foo bar":
                                            switch fi17 {
                                            case -9223372036854775808:
                                                v31 := GenS2{F0: (^(v1)), F1: 598, F2: 255}
                                                _ = v31
                                            case -392:
                                                var v32 float32 = float32(v11)
                                                _ = v32
                                            }
                                        default:
                                            switch (fi17 >> ((v4) & 63)) {
                                            case 1:
                                                v25 += (^(v22))
                                                if (v3 && false) {
                                                    v23 -= 127
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                } else {
                                                    var v33 float32 = 1e10
                                                    _ = v33
                                                }
                                                if len(v18) > 0 {
                                                    v18[0] = fi17
                                                }
                                            case 127:
                                                var v34 uint8 = (v1 + 128)
                                                _ = v34
                                            case 4294967295:
                                                switch fi17 {
                                                case 0:
                                                    if (false && false) {
                                                        var v35 float32 = 3.14159
                                                        _ = v35
                                                    } else {
                                                        v22 = genFn0(2)
                                                    }
                                                default:
                                                    v4 -= (v4 ^ 0)
                                                }
                                            }
                                        }
                                        v5 -= (v5 | int32(v23))
                                    } else {
                                        v36 := GenS1{F0: v3, F1: 255, F2: -9223372036854775808}
                                        _ = v36
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(v1, v4)
                                    case "foo bar":
                                        var v37 uint32 = 1
                                        _ = v37
                                    case "\n\t":
                                        fmt.Println(v25)
                                    default:
                                        var v38 uint64 = 18446744073709551615
                                        _ = v38
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                    default:
                                        v39 := GenS3{F0: 127, F1: "\n\t"}
                                        _ = v39
                                    }
                                default:
                                    v5 = (v5 * -394)
                                }
                                v3 = (((v1 | (v1 * 1)) != 128) && (534 > (v4 ^ v4)))
                                v22 += v22
                            } else {
                                if ((((v2 & 42) > v2) && (449 <= 0)) || v3) {
                                    if (!true) {
                                        v40 := GenS3{F0: 2, F1: ""}
                                        switch ("hello" + "a") {
                                        case "hello":
                                            fmt.Println(v25)
                                        case "\n\t":
                                            if ((1000 >= 1) && ((genFn1() == genFn1()) && false)) {
                                                switch "\n\t" {
                                                case "hello":
                                                    var v41 int16 = 0
                                                    _ = v41
                                                case "":
                                                    v18 = append(v18, (fi15 % (((fi17 << ((v4) & 63))) | 1)))
                                                case "foo bar":
                                                    if (!(false && (-1 <= v22))) {
                                                        var v42 int64 = 404
                                                        _ = len(v18)
                                                        _ = v42
                                                    } else {
                                                        var v43 float32 = 1e10
                                                        v44 := GenS2{F0: 254, F1: 32767, F2: 520}
                                                        if true {
                                                            v25 -= genFn0(v1)
                                                            var v45 int = -499
                                                            _ = v45
                                                        } else {
                                                            v18 = append(v18, (fi15 << ((42) & 63)))
                                                            if (127 < fi15) {
                                                                v25 = v22
                                                            } else {
                                                                v2 = (-(v5))
                                                                var v46 uint = v4
                                                                v23 = (v23 * v23)
                                                                _ = v46
                                                            }
                                                            var v47 uint8 = v1
                                                            _ = v47
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                default:
                                                    v18 = append(v18, -2147483648)
                                                }
                                                fmt.Println(v1, fi17)
                                                var v48 uint32 = 4294967295
                                                _ = v48
                                            }
                                        case "a":
                                            var v49 int = (fi15 - (v23 / (((fi17 << (((v4 + (v4 | v4))) & 63))) | 1)))
                                            _ = v49
                                        }
                                        _ = v40
                                    }
                                    var v50 bool = (((1.5 - ((0.1 + 1e10) * 1.5)) / 1.0) <= (0.1 / 0.1))
                                    _ = v50
                                } else {
                                    var v51 float32 = 1e10
                                    fmt.Println(v2, fi15)
                                    if ("foo bar" == ("\n\t" + ((((("" + "hello") + (((((("foo bar" + ("" + "foo bar")) + "\n\t") + "hello") + "\n\t") + "hello") + "foo bar")) + "\n\t") + "") + ""))) {
                                        fmt.Println(fi15, fi15)
                                        v52 := GenS2{F0: v1, F1: 883, F2: 0}
                                        v1 -= v1
                                        if (!(0 <= -146)) {
                                            var v53 uint = 4294967295
                                            var v54 float32 = v51
                                            if len(v18) > 0 {
                                                v18[0] = fi17
                                            }
                                            v51 += float32(float64(v51))
                                            _ = v53
                                            _ = v54
                                        } else {
                                            if (436 != 933) {
                                                _ = len(v18)
                                                var v55 int16 = int16(v1)
                                                var v56 float64 = v11
                                                _ = v55
                                                _ = v56
                                            }
                                            v57 := GenS1{F0: v3, F1: 244, F2: 9223372036854775807}
                                            _ = v57
                                        }
                                        _ = v52
                                    }
                                    _ = len(v18)
                                    var v58 int16 = -32768
                                    _ = v51
                                    _ = v58
                                }
                                if len(v18) > 0 {
                                    v18[0] = -1
                                }
                                var v59 float64 = float64(float32((v11 / (((0.1 + ((0.1 + (1.5 / float64(1.5))) * (v11 - v11))) - (1.5 + 0.1)) * 3.14159))))
                                _ = v59
                            }
                            if (!true) {
                                var v60 int32 = 2147483647
                                var v61 uint64 = 2
                                v1 = 255
                                _ = v60
                                _ = v61
                            }
                            v11 = v11
                        }
                        v11 += v11
                        if v3 {
                            if len(v18) > 0 {
                                v18[0] = (^(fi15))
                            }
                        } else {
                            v1 = v1
                        }
                    }
                    var v62 int8 = v22
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v62
                }
                var v63 string = "a"
                if (!v3) {
                    switch int(v1) {
                    case 1:
                        if (v4 == v4) {
                            var v64 int64 = 0
                            var v65 string = v63
                            v66 := GenS1{F0: v3, F1: 218, F2: v64}
                            var v67 int8 = 42
                            if (-2.25 <= v11) {
                                if (9223372036854775807 == (fi15 / ((fi17) | 1))) {
                                    v63 = v63
                                    v1 = v1
                                }
                                var v68 bool = (!(42 != 916))
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                _ = v68
                            } else {
                                v69 := GenS4{F0: ("hello" + (v63 + v63))}
                                var v70 uint8 = v1
                                _ = v69
                                _ = v70
                            }
                            _ = v64
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    case 332:
                        switch fi15 {
                        case 0:
                            v71 := GenS2{F0: v1, F1: 332, F2: 2}
                            _ = v71
                        case 152:
                            v63 = ""
                        case 690:
                            var v72 int16 = 32767
                            var v73 float32 = float32(1.5)
                            if v3 {
                                v74 := GenS5{F0: fi17, F1: 75}
                                v75 := GenS1{F0: (1 > 32767), F1: 42, F2: -497}
                                _ = v74
                                _ = v75
                            }
                            _ = v72
                            _ = v73
                        }
                    case 9223372036854775807:
                        var v76 uint32 = 4294967295
                        _ = v76
                    }
                    fmt.Println(v3)
                    var v77 string = v63
                    _ = v77
                } else {
                    v78 := GenS1{F0: v3, F1: 468, F2: -509}
                    _ = v78
                }
                _ = fi17
                _ = v18
                _ = v63
            }
            _ = fi15
            _ = v16
        }
    case 42:
        v4 -= (v4 & v4)
    case 32767:
        var v79 uint64 = 18446744073709551615
        _ = v79
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        v1 -= (v1 >> ((v4) & 63))
        _ = fi80
    }
    for fi81 := 0; fi81 < 12; fi81++ {
        v5 -= (-(v5))
        var v82 bool = (!true)
        if v3 {
            v2 = (v5 & (v2 & v2))
        }
        _ = fi81
        _ = v82
    }
    v83 := []int{0, 144}
    v1 = 225
    var v84 uint = v4
    for fi85 := 0; fi85 < 8; fi85++ {
        v3 = (!((v1 / ((2) | 1)) < (v1 >> ((v4) & 63))))
        v86 := GenS6{F0: 0, F1: fi85}
        var v87 float64 = v11
        _ = fi85
        _ = v86
        _ = v87
    }
    if v3 {
        if len(v83) >= 2 {
            v83 = v83[1:len(v83)]
        }
    }
    fmt.Println(v5, v3)
    v83 = append(v83, 137)
    v3 = (32767 > 303)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
}
