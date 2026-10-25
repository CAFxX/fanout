package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint32; F2 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0(p0 uint16) uint {
    v7 := GenS1{F0: "", F1: 38, F2: "\n\t"}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float32 = float32((0.1 / 1.5))
        switch "hello" {
        case "foo bar":
            v10 := []int{128}
            _ = v10
        case "":
            var v11 uint8 = 42
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                switch "foo bar" {
                case "hello":
                    p0 = p0
                case "\n\t":
                    if ((0.1 - 1e10) >= ((((v9 * (float32(-2.25) - (float32((3.14159 + 1e10)) / (0.0 + (float32(0.0) + (1.5 + 0.1)))))) * (float32(1e10) / (v9 * ((v9 - (v9 / 1.0)) + (v9 / (((v9 * v9) + v9) + (float32(1e10) + v9))))))) / (1.5 * v9)) - v9)) {
                        v13 := GenS1{F0: (("\n\t" + (("a" + ("\n\t" + "a")) + "")) + "hello"), F1: 251, F2: ""}
                        var v14 float64 = (0.1 / ((((3.14159 * 0.1) - 3.14159) * -2.25) / (1.5 / 3.14159)))
                        switch 128 {
                        case -265:
                            var v15 int16 = 1
                            _ = v15
                        case -242:
                            v16 := GenS2{F0: 566, F1: "\n\t"}
                            _ = v16
                        case 32767:
                            if (v9 == 1.5) {
                                v17 := GenS2{F0: 0, F1: (("hello" + "") + "\n\t")}
                                v18 := []int{fi8, (^(fi8)), -134}
                                fmt.Println(p0, v14, fi12)
                                _ = v17
                                _ = v18
                            } else {
                                p0 = 661
                            }
                        }
                        _ = v13
                        _ = v14
                    }
                case "":
                    var v19 uint32 = 1
                    _ = v19
                default:
                    var v20 float64 = (0.0 + ((((-2.25 + 3.14159) / 1.5) / (1.5 - 0.0)) - (((0.0 / 0.1) / 1.0) / ((1.5 + 3.14159) - -2.25))))
                    _ = v20
                }
                v21 := GenS3{F0: 127}
                if (false || true) {
                    if ((127 != fi12) && false) {
                        v22 := GenS2{F0: -142, F1: "\n\t"}
                        fmt.Println(v9, fi8)
                        var v23 bool = true
                        _ = v22
                        _ = v23
                    }
                    v9 = 3.14159
                    if (42 < 42) {
                        if (true && ((-2.25 - 0.0) > (float64(0.1) - (0.1 + 0.0)))) {
                            var v24 bool = false
                            v25 := []int{(^(fi8)), 32768, (fi8 | fi8), fi8}
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                _ = len(v25)
                            case "a":
                                if (0.0 >= -2.25) {
                                    if false {
                                        fmt.Println(v24, fi12)
                                    }
                                } else {
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    v9 -= v9
                                    v25 = append(v25, fi8)
                                    v26 := GenS4{F0: 128}
                                    _ = v26
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                    } else {
                        var v27 uint8 = 118
                        var v28 float64 = 1e10
                        switch fi12 {
                        case 32768:
                            fmt.Println(fi12, fi8)
                        case 42:
                            var v29 bool = false
                            _ = v29
                        default:
                            v30 := GenS2{F0: 420, F1: "foo bar"}
                            _ = v30
                        }
                        fmt.Println(fi8, v28)
                        v28 += v28
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi8, v9)
                    var v31 int16 = -157
                    var v32 uint = 2
                    _ = v31
                    _ = v32
                } else {
                    if false {
                        var v33 int16 = -32768
                        v34 := GenS5{F0: "\n\t"}
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(v9, p0, fi12)
                        }
                        var v35 int64 = 9223372036854775807
                        v36 := GenS6{F0: 65535, F1: 122, F2: -1}
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        v37 := []int{fi8, (fi8 & fi8)}
                        _ = v37
                    }
                    v38 := GenS1{F0: ("" + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t")))), F1: 1, F2: (("a" + (("" + "hello") + "\n\t")) + ("hello" + "\n\t"))}
                    var v39 int64 = 809
                    _ = v38
                    _ = v39
                }
                v40 := GenS7{F0: 0}
                v9 = (v9 / 1e10)
                _ = fi12
                _ = v21
                _ = v40
            }
        }
        v41 := []int{fi8, fi8, (fi8 << ((65535) & 63))}
        _ = fi8
        _ = v41
        _ = v9
    }
    var v42 uint64 = 2
    v43 := []int{127, 42, -273}
    if true {
        p0 -= 1
        var v44 uint8 = 128
        switch 775 {
        case 32768:
            fmt.Println(p0)
        case 128:
            var v45 bool = true
            _ = v45
        case 42:
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 float64 = (((float64(0.1) / 0.1) - float64(0.0)) - -2.25)
                    p0 -= (^(p0))
                    _ = fi46
                    _ = v47
                }
            case "\n\t":
                fmt.Println(v42, v42, v44)
            case "":
                v48 := GenS4{F0: 247}
                _ = v48
            }
        }
        _ = v44
    }
    _ = p0
    _ = v42
    _ = v43
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16, p1 int64, p2 uint64) int32 {
    v49 := []int{9223372036854775807}
    var v50 uint8 = 2
    v50 -= (v50 / (((v50 >> ((uint(p0)) & 63))) | 1))
    v51 := GenS2{F0: 0, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v49
    _ = v50
    _ = v51
    return 1
}

func main() {
    var v1 string = (("\n\t" + "foo bar") + ("a" + "\n\t"))
    var v2 uint16 = 42
    var v3 string = (((("hello" + (((v1 + v1) + v1) + "\n\t")) + (v1 + (v1 + ((("a" + v1) + "\n\t") + "hello")))) + ("a" + ((("" + "") + v1) + "\n\t"))) + ("a" + v1))
    var v4 uint64 = 714
    var v5 int64 = 875
    var v6 float64 = 1e10
    v3 = v3
    v5 = v5
    var v52 uint8 = 214
    v53 := []int{128, 628}
    v54 := GenS8{F0: 928, F1: true}
    v6 = (v6 - (v6 - v6))
    if (-2.25 <= 0.0) {
        switch 2147483647 {
        case -161:
            v55 := GenS6{F0: 42, F1: 2, F2: 0}
            _ = v55
        case 0:
            v52 = uint8(v4)
        case -9223372036854775808:
            for fi56 := 0; fi56 < 4; fi56++ {
                v57 := GenS3{F0: fi56}
                if len(v53) >= 2 {
                    v53 = v53[1:len(v53)]
                }
                _ = fi56
                _ = v57
            }
        default:
            for fi58 := 0; fi58 < 20; fi58++ {
                v53 = append(v53, (fi58 >> ((65535) & 63)))
                var v59 int16 = 32767
                if (false && (false || true)) {
                    v4 = uint64(v2)
                } else {
                    v52 -= v52
                    if (((v4 != (v4 & v4)) || false) && (!false)) {
                        if ((v4 - (v4 & 42)) >= v4) {
                            if len(v53) >= 2 {
                                v53 = v53[1:len(v53)]
                            }
                            var v60 uint8 = v52
                            _ = v60
                        }
                    }
                    for fi61 := 0; fi61 < 19; fi61++ {
                        v53 = append(v53, fi58)
                        var v62 uint16 = v2
                        _ = fi61
                        _ = v62
                    }
                }
                _ = fi58
                _ = v59
            }
        }
        var v63 int = -9223372036854775808
        var v64 int32 = 29
        _ = v63
        _ = v64
    }
    v6 += -2.25
    var v65 uint64 = 1
    for fi66 := 0; fi66 < 5; fi66++ {
        v2 = (^(v2))
        _ = fi66
    }
    v4 = 842
    var v67 uint = genFn0((v2 * (v2 ^ 42)))
    switch -9223372036854775808 {
    case -2147483648:
        var v68 bool = (!false)
        _ = v68
    default:
        v53 = append(v53, -328)
    }
    if false {
        v69 := GenS7{F0: v4}
        _ = v69
    }
    v65 += (^(v4))
    v1 = v3
    if (((32767 == (v2 >> ((4294967295) & 63))) || false) || (v52 <= uint8(v2))) {
        for fi70 := 0; fi70 < 4; fi70++ {
            var v71 float32 = (((((float32(-2.25) / (((-2.25 / 3.14159) * -2.25) * (0.1 * 3.14159))) - 1e10) - 1e10) / 3.14159) + 0.0)
            v72 := GenS8{F0: 2147483648, F1: false}
            _ = fi70
            _ = v71
            _ = v72
        }
    } else {
        for fi73 := 0; fi73 < 3; fi73++ {
            switch fi73 {
            case -798:
                for fi74 := 0; fi74 < 15; fi74++ {
                    _ = len(v53)
                    v75 := GenS1{F0: (v1 + "foo bar"), F1: (v52 / ((128) | 1)), F2: ((v3 + v3) + "foo bar")}
                    _ = fi74
                    _ = v75
                }
            case 42:
                v6 = -2.25
            case 32768:
                v5 += -589
            }
            if (!false) {
                v6 += 0.0
                v65 -= v65
            }
            switch 32768 {
            case 730:
                for fi76 := 0; fi76 < 17; fi76++ {
                    var v77 uint32 = 4294967295
                    _ = fi76
                    _ = v77
                }
            case 32767:
                v78 := GenS9{F0: int32(v52), F1: (v52 << (((v67 * genFn0(uint16(v65)))) & 63)), F2: 203}
                _ = v78
            case 782:
                for fi79 := 0; fi79 < 16; fi79++ {
                    v1 = ((v3 + (("\n\t" + (v1 + "foo bar")) + "foo bar")) + v3)
                    _ = fi79
                }
            default:
                v80 := GenS6{F0: (v67 >> ((genFn0(544)) & 63)), F1: 0, F2: 0}
                _ = v80
            }
            _ = fi73
        }
        if ("\n\t" <= v3) {
            v67 = (^(v67))
        }
    }
    v53 = append(v53, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v65
    _ = v67
}
