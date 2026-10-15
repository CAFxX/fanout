package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int }
type GenS9 struct { F0 uint; F1 int16; F2 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool; F1 int16; F2 int8 }
type GenS12 struct { F0 int16; F1 int32 }
func genFn0(p0 bool) uint {
    v6 := []int{128, 235, 32768}
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 = append(v6, 4294967295)
        var v8 float32 = 1.5
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: p0, F1: 2147483647, F2: -9223372036854775808}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    return 65535
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = 694
    var v3 uint16 = (v1 * (^(v1)))
    var v4 uint8 = 168
    var v5 int64 = v2
    var v11 uint64 = 811
    fmt.Println(v5, v3)
    v12 := GenS1{F0: (false || (v2 <= v2)), F1: 716, F2: v5}
    v1 -= (v1 / ((v3) | 1))
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 = (v3 - (v1 % (((v1 - v1)) | 1)))
        v14 := []int{fi13}
        v15 := GenS1{F0: (false && false), F1: 65535, F2: 0}
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            v4 = 2
            for fi17 := 0; fi17 < 14; fi17++ {
                v2 += int64(v3)
                v18 := []int{fi16, fi17, (fi16 * (-(fi16))), -421}
                v19 := GenS1{F0: true, F1: 1, F2: 42}
                v11 = (v11 | (^(v11)))
                v20 := GenS1{F0: true, F1: 335, F2: (^(v5))}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                fmt.Println(v5, v2, fi21)
                _ = fi21
            }
            _ = fi16
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 float32 = ((float32(3.14159) - float32(1.5)) * (1e10 + 0.1))
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 int64 = (v2 - v2)
                if false {
                    if (-32768 == 953) {
                        switch ("foo bar" + "") {
                        case "foo bar":
                            var v26 uint = 42
                            _ = v26
                        case "":
                            v25 += v25
                        case "\n\t":
                            v5 -= 0
                        }
                        var v27 string = ((("a" + "\n\t") + "a") + "a")
                        _ = v27
                    }
                } else {
                    var v28 string = ""
                    _ = v28
                }
                _ = fi24
                _ = v25
            }
            _ = fi22
            _ = v23
        }
    } else {
        if true {
            v4 += v4
            var v29 string = ("a" + ("a" + ""))
            _ = v29
        }
    }
    var v30 int32 = 69
    if false {
        var v31 float64 = (1e10 * (-2.25 + 0.0))
        _ = v31
    }
    v32 := GenS2{F0: 65535}
    for fi33 := 0; fi33 < 7; fi33++ {
        fmt.Println(v30, v5, v1)
        switch fi33 {
        case 128:
            var v34 int8 = 0
            _ = v34
        default:
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := GenS2{F0: 65535}
                switch 42 {
                case 1:
                    var v37 uint = 535
                    _ = v37
                default:
                    v30 = (v30 & v30)
                }
                v38 := GenS3{F0: (("\n\t" + ((("hello" + ("" + "a")) + "foo bar") + "foo bar")) + "foo bar"), F1: (v30 % (((v30 | (-(v30)))) | 1))}
                _ = fi35
                _ = v36
                _ = v38
            }
        }
        fmt.Println(v1)
        _ = fi33
    }
    v39 := []int{1, 714}
    for fi40 := 0; fi40 < 19; fi40++ {
        if false {
            fmt.Println(v3)
        } else {
            fmt.Println(v5)
            var v41 uint16 = 65535
            _ = v41
        }
        _ = fi40
    }
    fmt.Println(v4, v5)
    fmt.Println(v1, v11, v4)
    if ((("foo bar" + "hello") + (("a" + ("foo bar" + (("" + ("a" + "\n\t")) + (("hello" + "") + "hello")))) + ("" + "foo bar"))) >= "a") {
        var v42 uint32 = 2
        if (!false) {
            if (((v4 & 255) >= v4) && false) {
                for fi43 := 0; fi43 < 13; fi43++ {
                    fmt.Println(v5, v2, v3)
                    _ = fi43
                }
                if (1 > 1) {
                    fmt.Println(v1)
                } else {
                    v5 -= 0
                    fmt.Println(v42)
                    var v44 uint64 = (v11 - (v11 / ((v11) | 1)))
                    _ = v44
                }
            } else {
                if true {
                    if (("hello" > "") || false) {
                        v1 -= (v1 << ((65535) & 63))
                        switch 32767 {
                        case -2147483648:
                            v3 = v1
                        case 127:
                            var v45 int = 0
                            _ = v45
                        case 433:
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                        }
                    } else {
                        var v46 int8 = 0
                        if true {
                            fmt.Println(v1, v4)
                            v47 := GenS4{F0: ((v42 & v42) <= 1), F1: (v42 | v42)}
                            v48 := GenS5{F0: -1}
                            _ = v47
                            _ = v48
                        } else {
                            var v49 float64 = (1.5 - -2.25)
                            if len(v39) > 0 {
                                v39[0] = 9223372036854775807
                            }
                            _ = v49
                        }
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v51 := GenS6{F0: 9223372036854775807, F1: 42, F2: (v3 % (((v3 - v1)) | 1))}
                            v39 = append(v39, fi50)
                            fmt.Println(v30)
                            _ = fi50
                            _ = v51
                        }
                        _ = v46
                    }
                }
                v39 = append(v39, -755)
            }
        }
        v39 = append(v39, 572)
        _ = v42
    } else {
        if true {
            var v52 float32 = 0.1
            var v53 float32 = v52
            fmt.Println(v52)
            _ = v52
            _ = v53
        }
        v39 = append(v39, 406)
    }
    for fi54 := 0; fi54 < 12; fi54++ {
        var v55 float64 = (float64(0.0) + -2.25)
        for fi56 := 0; fi56 < 4; fi56++ {
            if (!((v11 ^ 1) < v11)) {
                if (uint8(v11) == 128) {
                    switch "\n\t" {
                    case "\n\t":
                        switch "" {
                        case "":
                            v3 += (v3 - uint16(v4))
                        case "foo bar":
                            fmt.Println(v3, v4, v2)
                        case "a":
                            _ = len(v39)
                        default:
                            var v57 uint64 = v11
                            _ = v57
                        }
                    case "":
                        v58 := GenS1{F0: (false || (!(!false))), F1: 42, F2: 1}
                        _ = v58
                    default:
                        if (false && false) {
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                            v59 := GenS3{F0: "\n\t", F1: (^(v30))}
                            var v60 float32 = 3.14159
                            _ = v59
                            _ = v60
                        } else {
                            v61 := GenS7{F0: v11, F1: 42}
                            if (0.0 > (0.1 + (((v55 / v55) + v55) * (float64((1.5 * 3.14159)) * v55)))) {
                                var v62 float32 = float32(v55)
                                if (!false) {
                                    var v63 string = "a"
                                    v2 = (v2 * v5)
                                    if (3.14159 <= (v62 + ((((float32((v55 * v55)) * -2.25) / v62) * 0.1) / ((-2.25 / (v62 / v62)) * (v62 / 3.14159))))) {
                                        var v64 float64 = v55
                                        var v65 bool = false
                                        v65 = (17 <= 1000)
                                        v4 = (v4 - (v4 >> ((65535) & 63)))
                                        v66 := GenS8{F0: v4, F1: 4294967295, F2: -9223372036854775808}
                                        fmt.Println(v65)
                                        var v67 float64 = (0.0 / 1.5)
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    if (v4 <= 0) {
                                        v63 = v63
                                        if ((v3 << ((genFn0(false)) & 63)) >= (v3 & v3)) {
                                            switch "hello" {
                                            case "hello":
                                                v55 += ((0.1 * v55) + v55)
                                            }
                                        }
                                    }
                                    _ = v63
                                }
                                var v68 uint64 = uint64(v4)
                                _ = v62
                                _ = v68
                            } else {
                                v11 = v11
                            }
                            v69 := GenS3{F0: "\n\t", F1: (v30 ^ (v30 % ((v30) | 1)))}
                            _ = v61
                            _ = v69
                        }
                    }
                } else {
                    if (v30 > -939) {
                        switch (fi56 & fi54) {
                        case 4294967295:
                            v70 := GenS9{F0: 0, F1: -583, F2: (-(fi56))}
                            _ = v70
                        default:
                            v71 := GenS10{F0: (v30 << ((18446744073709551615) & 63))}
                            _ = v71
                        }
                        v39 = append(v39, 42)
                    } else {
                        var v72 int64 = (v2 & (-(v2)))
                        v30 = v30
                        v73 := GenS11{F0: ((0.1 - 0.1) > (1.5 * (-2.25 / 0.1))), F1: -32768, F2: 0}
                        switch "" {
                        case "foo bar":
                            v3 += 955
                        case "":
                            v55 += 1e10
                        case "a":
                            switch fi54 {
                            case 32767:
                                if false {
                                    v2 += -544
                                    v74 := GenS3{F0: "\n\t", F1: 1}
                                    v4 += v4
                                    _ = v74
                                }
                            }
                        }
                        _ = v72
                        _ = v73
                    }
                    v30 -= -512
                    if (false || true) {
                        if (((3.14159 * float32(0.1)) != 0.1) || false) {
                            v75 := GenS6{F0: 1, F1: 2, F2: (v1 | (v3 & (v3 << ((4294967295) & 63))))}
                            v76 := GenS12{F0: 1, F1: 1}
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v55, v30)
                            v2 -= v5
                        }
                        v39 = append(v39, -951)
                        var v77 int8 = -20
                        if (1 < 0) {
                            if (false || ((v30 / ((v30) | 1)) != (v30 * v30))) {
                                var v78 int32 = (^(v30))
                                _ = v78
                            }
                        }
                        _ = v77
                    }
                }
            }
            _ = fi56
        }
        _ = fi54
        _ = v55
    }
    for fi79 := 0; fi79 < 2; fi79++ {
        var v80 float32 = 0.0
        fmt.Println(fi79, v2)
        for fi81 := 0; fi81 < 8; fi81++ {
            if true {
                v30 -= 0
                var v82 int = -143
                v83 := GenS7{F0: 0, F1: 42}
                _ = v82
                _ = v83
            }
            if len(v39) > 0 {
                v39[0] = 128
            }
            _ = fi81
        }
        _ = fi79
        _ = v80
    }
    switch "a" {
    case "":
        var v84 uint8 = 35
        _ = v84
    case "a":
        var v85 float64 = 1e10
        _ = v85
    case "foo bar":
        var v86 float64 = 0.1
        _ = v86
    default:
        if false {
            _ = len(v39)
        } else {
            var v87 int64 = (v2 % (((^(v5))) | 1))
            _ = v87
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v39
    _ = v4
    _ = v5
}
