package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 uint32; F1 string; F2 uint }
type GenS10 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS11 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0() int32 {
    v7 := GenS1{F0: 128}
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch fi9 {
            case 32767:
                fmt.Println(fi10)
            case 0:
                v11 := []int{(fi10 + fi9), 175, 4294967295}
                _ = v11
            }
            _ = fi10
        }
        v12 := GenS2{F0: "a", F1: 65535}
        v8 -= (v8 | 2)
        _ = fi9
        _ = v12
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint32 = 153
    v3 := []uint32{(v2 % ((v2) | 1))}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 223
    v6 := []uint32{}
    var v13 uint8 = 42
    var v14 int = -2147483648
    if true {
        fmt.Println(v13)
        v15 := GenS3{F0: 18446744073709551615}
        var v16 uint64 = 1
        _ = v15
        _ = v16
    } else {
        if ((false && true) || (true && (572 != 9223372036854775807))) {
            switch "" {
            case "hello":
                if true {
                    v1 = v1
                    var v17 int8 = int8(v2)
                    if (false || true) {
                        var v18 int32 = -2147483648
                        if (32767 >= (v14 - 9223372036854775807)) {
                            var v19 bool = false
                            _ = v19
                        }
                        fmt.Println(v18)
                        _ = v18
                    }
                    _ = v17
                } else {
                    v20 := GenS3{F0: 801}
                    _ = v20
                }
            }
            fmt.Println(v14, v1, v13)
        }
        var v21 int8 = 127
        _ = v21
    }
    switch 378 {
    case 9223372036854775807:
        if false {
            fmt.Println(v1, v2)
            v22 := GenS4{F0: 1}
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 string = (("a" + "") + ("a" + "hello"))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
    default:
        v5 = 260
    }
    v25 := GenS5{F0: 121, F1: 985}
    v4[v2] = true
    if ((("" == ("hello" + "a")) && false) || (!true)) {
        if len(v6) > 0 {
            v6[0] = 42
        }
        fmt.Println(v13)
        for fi26 := 0; fi26 < 12; fi26++ {
            if false {
                if (v14 == 127) {
                    if false {
                        if false {
                            if ((!(1000 > -74)) || (!true)) {
                                v27 := v4[(^(v2))]
                                _ = v27
                            } else {
                                if true {
                                    var v28 float64 = (0.0 / 0.1)
                                    if (!true) {
                                        fmt.Println(v14)
                                    }
                                    _ = v28
                                }
                                fmt.Println(v5, v1)
                                for fi29 := 0; fi29 < 3; fi29++ {
                                    var v30 uint32 = (^(v2))
                                    v31 := GenS6{F0: (true && true), F1: "foo bar"}
                                    v3 = append(v3, v2)
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            for fi32 := 0; fi32 < 15; fi32++ {
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    v33 := GenS1{F0: v13}
                                    _ = v33
                                }
                                _ = fi32
                            }
                        }
                        v34 := GenS1{F0: 48}
                        v35 := GenS3{F0: 2}
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int64 = 0
                        switch "foo bar" {
                        case "a":
                            v37 := GenS7{F0: "a", F1: false}
                            _ = v37
                        case "hello":
                            var v38 uint32 = v2
                            _ = v38
                        case "foo bar":
                            switch (("foo bar" + "foo bar") + ("hello" + "\n\t")) {
                            case "hello":
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    v1 = genFn0()
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi39
                                }
                            default:
                                v40 := GenS7{F0: ("" + ("a" + "\n\t")), F1: false}
                                _ = v40
                            }
                        }
                        if (!false) {
                            if false {
                                if (false || (v36 < (v36 | int64(v5)))) {
                                    delete(v4, (v2 & (^(v2))))
                                    v36 -= v36
                                }
                                for fi41 := 0; fi41 < 16; fi41++ {
                                    fmt.Println(v5)
                                    if (!((((((float64(1e10) - float64((1e10 * float32(float64(-2.25))))) - 1e10) + 0.1) + 3.14159) / ((3.14159 * 0.0) - (((0.0 / -2.25) - ((1.5 + 3.14159) + (1.5 + float64(-2.25)))) + (3.14159 - 1.5)))) == 0.1)) {
                                        fmt.Println(v13, v14)
                                        var v42 float64 = ((-2.25 / float64((1e10 - (-2.25 * (1e10 - 0.0))))) * (1.5 - (float64(float32((3.14159 / ((0.1 - float64((-2.25 * 3.14159))) + 3.14159)))) * (1e10 - (((1e10 + 0.0) / 1.5) / (0.0 + 0.1))))))
                                        if false {
                                            var v43 bool = false
                                            if (v43 && ((!(!false)) || (!v43))) {
                                                v4[(v2 ^ (v2 << ((0) & 63)))] = (((v2 - 369) >= (v2 % ((v2) | 1))) || ((32767 <= 32767) && true))
                                                var v44 uint32 = uint32(v13)
                                                _ = v44
                                            } else {
                                                var v45 int16 = -32768
                                                v2 = 42
                                                _ = v45
                                            }
                                            var v46 int16 = 32767
                                            _ = v43
                                            _ = v46
                                        }
                                        switch "" {
                                        case "":
                                            if len(v6) > 0 {
                                                v6[0] = uint32(v13)
                                            }
                                        case "a":
                                            v47 := GenS1{F0: (v13 & (^(v13)))}
                                            _ = v47
                                        case "foo bar":
                                            var v48 int64 = (^(v36))
                                            _ = v48
                                        }
                                        if true {
                                            v49 := GenS6{F0: true, F1: "\n\t"}
                                            _ = v49
                                        } else {
                                            var v50 float32 = (((0.0 + (3.14159 - (((0.0 * (1.5 / (1e10 + 1.5))) + 0.0) + (0.0 - (0.0 - 1.5))))) * 0.0) / 3.14159)
                                            _ = v50
                                        }
                                        _ = v42
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    _ = fi41
                                }
                            }
                            if true {
                                var v51 int16 = 1000
                                _ = v51
                            }
                        }
                        _ = v36
                    }
                    if len(v3) > 0 {
                        v3[0] = (v2 | (v2 * 2147483648))
                    }
                    v13 += (v13 / ((v13) | 1))
                } else {
                    var v52 float32 = ((3.14159 - (3.14159 * float32(0.0))) + 0.0)
                    _ = v52
                }
                switch (v14 - 42) {
                case -2147483648:
                    v13 += 2
                case 2147483647:
                    for fi53 := 0; fi53 < 9; fi53++ {
                        v6 = append(v6, (v2 / (((^(v2))) | 1)))
                        v54 := GenS3{F0: 40}
                        v55 := GenS3{F0: 18446744073709551615}
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                case 32767:
                    var v56 int = v14
                    _ = v56
                default:
                    if (((((1.5 / 1.0) + 1e10) / 3.14159) <= (0.0 / 1e10)) || (!false)) {
                        v2 -= (v2 / ((v2) | 1))
                        _ = len(v4)
                    }
                }
                var v57 uint = 2
                _ = v57
            } else {
                if false {
                    v6 = append(v6, 42)
                    v2 = v2
                }
            }
            _ = fi26
        }
    }
    var v58 uint32 = uint32(v5)
    v59 := GenS8{F0: v13, F1: 2}
    for fi60 := 0; fi60 < 2; fi60++ {
        if (4294967295 <= 4294967295) {
            for fi61 := 0; fi61 < 12; fi61++ {
                var v62 uint = 2
                _ = fi61
                _ = v62
            }
        } else {
            for fi63 := 0; fi63 < 7; fi63++ {
                var v64 float64 = -2.25
                if false {
                    switch -1 {
                    case -596:
                        var v65 int64 = 672
                        _ = v65
                    case -615:
                        v13 -= (v13 + v13)
                    case 2147483647:
                        v2 -= (v58 * (v2 - 65535))
                    }
                } else {
                    if false {
                        v66 := GenS7{F0: "a", F1: (true && ("hello" <= ""))}
                        v14 = -9223372036854775808
                        v64 = (v64 * (float64(1e10) + float64(1.5)))
                        _ = v66
                    }
                    v67 := GenS9{F0: uint32(v13), F1: ("\n\t" + ""), F2: 870}
                    v68 := GenS7{F0: "foo bar", F1: (!(-32768 == 420))}
                    _ = v67
                    _ = v68
                }
                fmt.Println(v5, v1, fi60)
                _ = fi63
                _ = v64
            }
            v13 += 1
        }
        v14 = (fi60 * int(v13))
        var v69 int32 = (v1 ^ (v1 >> ((1) & 63)))
        _ = fi60
        _ = v69
    }
    for fi70 := 0; fi70 < 20; fi70++ {
        switch v14 {
        case -924:
            v1 += (v1 | v1)
        }
        var v71 int8 = 42
        v72 := GenS7{F0: "", F1: true}
        _ = fi70
        _ = v71
        _ = v72
    }
    var v73 uint16 = 1
    if (!(v13 < (v13 - v13))) {
        switch (v14 * v14) {
        case 469:
            v73 = (v5 | (^(v73)))
        case 127:
            if true {
                var v74 uint16 = 130
                for fi75 := 0; fi75 < 16; fi75++ {
                    switch v14 {
                    case -9223372036854775808:
                        if true {
                            v4[v2] = false
                            v1 += (v1 | (v1 / (((v1 + v1)) | 1)))
                            var v76 uint32 = v2
                            _ = v76
                        }
                    case -2147483648:
                        v77 := v4[uint32(v74)]
                        _ = v77
                    case 247:
                        v78 := GenS10{F0: 0, F1: 722, F2: v13}
                        _ = v78
                    }
                    v79 := GenS11{F0: 255, F1: -1, F2: genFn0()}
                    _ = fi75
                    _ = v79
                }
                v80 := GenS1{F0: v13}
                _ = v74
                _ = v80
            } else {
                switch (v14 | v14) {
                case 2147483647:
                    v2 += v2
                case -9223372036854775808:
                    if (!(!(938 != int64(v58)))) {
                        v13 += (v13 / (((v13 ^ uint8(v5))) | 1))
                        fmt.Println(v2)
                        v81 := GenS3{F0: 0}
                        _ = v81
                    }
                case -2147483648:
                    v82 := GenS10{F0: 42, F1: 4294967295, F2: (v13 / (((v13 + v13)) | 1))}
                    _ = v82
                }
                v1 = (v1 | genFn0())
                v58 += (v2 | 224)
            }
        case 0:
            var v83 int8 = 127
            _ = v83
        default:
            v2 -= (v2 % (((v58 + v58)) | 1))
        }
    }
    switch ("a" + (("" + ("hello" + "hello")) + "hello")) {
    case "a":
        v6 = append(v6, v2)
    case "":
        delete(v4, (v2 >> ((832) & 63)))
    case "hello":
        v84 := GenS3{F0: 18446744073709551615}
        _ = v84
    default:
        var v85 uint16 = (v73 << ((1) & 63))
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v73
}
