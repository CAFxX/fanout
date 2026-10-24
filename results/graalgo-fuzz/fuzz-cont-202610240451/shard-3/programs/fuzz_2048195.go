package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int; F1 uint16; F2 string }
type GenS7 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int64 = 42
    v2 := []uint32{}
    var v3 uint32 = 2147483647
    v4 := []string{"a", "\n\t"}
    var v5 uint32 = v3
    if (128 < int(v1)) {
        v6 := GenS1{F0: "", F1: -84}
        _ = v6
    }
    if (!true) {
        fmt.Println(v3, v1, v5)
        if (0 >= 178) {
            v7 := GenS2{F0: "foo bar", F1: 922}
            var v8 uint16 = uint16(v3)
            _ = v7
            _ = v8
        } else {
            v9 := GenS3{F0: 4294967295}
            var v10 uint64 = 4294967295
            var v11 string = "foo bar"
            _ = v10
            _ = v11
            _ = v9
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v12 uint8 = 1
    v5 = v5
    if true {
        var v13 int = -253
        _ = v13
    } else {
        for fi14 := 0; fi14 < 5; fi14++ {
            v1 -= v1
            _ = fi14
        }
        v15 := GenS2{F0: "foo bar", F1: 976}
        var v16 bool = false
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: "foo bar", F1: 95}
    var v18 bool = true
    switch 261 {
    case 42:
        switch "a" {
        case "\n\t":
            fmt.Println(v1, v5)
        case "a":
            v3 -= v3
        }
    case 32768:
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v5 -= v3
                v12 += v12
                var v21 uint = 552
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    }
    switch 32767 {
    case 4294967295:
        var v22 int64 = (^(v1))
        _ = v22
    }
    switch ("" + "hello") {
    case "a":
        var v23 int = 55
        _ = v23
    default:
        v3 = 2147483648
    }
    v1 -= v1
    _ = len(v4)
    if false {
        v1 = v1
        var v24 int32 = 832
        switch 710 {
        case -32:
            switch 42 {
            case -681:
                var v25 uint8 = v12
                _ = v25
            case 4294967295:
                v18 = (!v18)
            case -604:
                for fi26 := 0; fi26 < 3; fi26++ {
                    v3 = v3
                    fmt.Println(v24)
                    _ = fi26
                }
            }
        case -555:
            if v18 {
                v27 := GenS4{F0: false}
                _ = v27
            } else {
                var v28 int32 = (^(v24))
                _ = v28
            }
        default:
            v3 = 2147483648
        }
        _ = v24
    } else {
        v18 = (v18 && v18)
    }
    v5 = (^(v3))
    if ((-2.25 != 0.1) || ("foo bar" <= "a")) {
        v29 := GenS3{F0: 1}
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 uint8 = (v12 - 244)
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        fmt.Println(v1)
        v32 := GenS5{F0: v1, F1: 42}
        switch "hello" {
        case "":
            if true {
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS6{F0: fi33, F1: 42, F2: "\n\t"}
                    var v35 float64 = 1e10
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 uint64 = 216
                _ = v36
            } else {
                if (738 >= 628) {
                    for fi37 := 0; fi37 < 15; fi37++ {
                        if v18 {
                            v2 = append(v2, (v3 - (^(v5))))
                            switch "hello" {
                            case "\n\t":
                                var v38 uint8 = (v12 * (v12 & v12))
                                _ = v38
                            case "foo bar":
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    if v18 {
                                        v5 += v3
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (v12 <= v12) {
                                        switch 431 {
                                        case 509:
                                            switch fi39 {
                                            case -739:
                                                if false {
                                                    switch "" {
                                                    case "foo bar":
                                                        var v40 int32 = 0
                                                        _ = v40
                                                    case "hello":
                                                        var v41 int64 = v1
                                                        _ = v41
                                                    }
                                                } else {
                                                    var v42 uint = 2
                                                    var v43 uint32 = (v5 ^ 2147483647)
                                                    if v18 {
                                                        v12 -= 42
                                                    } else {
                                                        if (false && v18) {
                                                            v1 = (v1 - (v1 / ((v1) | 1)))
                                                        }
                                                        v18 = (v18 || true)
                                                        var v44 int32 = -444
                                                        _ = v44
                                                    }
                                                    _ = v42
                                                    _ = v43
                                                }
                                            default:
                                                var v45 string = "a"
                                                _ = v45
                                            }
                                        default:
                                            switch fi37 {
                                            case -407:
                                                var v46 string = "foo bar"
                                                _ = v46
                                            case 4294967295:
                                                var v47 uint32 = v5
                                                _ = v47
                                            case 760:
                                                var v48 int = (fi37 >> ((4294967295) & 63))
                                                fmt.Println(fi39, v12)
                                                _ = v48
                                            default:
                                                var v49 string = "\n\t"
                                                _ = v49
                                            }
                                        }
                                        var v50 bool = ((v12 ^ (v12 | v12)) == (^(v12)))
                                        var v51 int16 = -32768
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = fi39
                                }
                            }
                        }
                        fmt.Println(fi37, v3, v5)
                        if false {
                            if false {
                                if (v18 && v18) {
                                    v52 := GenS5{F0: v1, F1: 979}
                                    switch (-(fi37)) {
                                    case 725:
                                        v5 = (v5 * v3)
                                    case 0:
                                        if (v18 && (!(18446744073709551615 < 18446744073709551615))) {
                                            for fi53 := 0; fi53 < 10; fi53++ {
                                                if v18 {
                                                    v54 := GenS5{F0: v1, F1: -1}
                                                    if true {
                                                        v55 := GenS2{F0: "", F1: 4294967295}
                                                        if ((v18 && true) && (42 != 127)) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if false {
                                                                switch (fi37 + -2147483648) {
                                                                case 2147483647:
                                                                    var v56 uint16 = 65535
                                                                    _ = v56
                                                                default:
                                                                    v18 = true
                                                                }
                                                                var v57 uint64 = 225
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v58 := GenS3{F0: 18446744073709551615}
                                                                    _ = v58
                                                                case "":
                                                                    _ = len(v4)
                                                                default:
                                                                    var v59 int64 = (^(v1))
                                                                    _ = v59
                                                                }
                                                                _ = v57
                                                            }
                                                        }
                                                        var v60 float64 = 3.14159
                                                        _ = v55
                                                        _ = v60
                                                    } else {
                                                        var v61 int32 = 292
                                                        if (((v3 % ((v3) | 1)) < uint32(v12)) || ((v1 ^ v1) != v1)) {
                                                            var v62 bool = (0 > 1)
                                                            fmt.Println(v61)
                                                            v63 := GenS1{F0: ("a" + "a"), F1: 100}
                                                            var v64 int16 = -32768
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        _ = v61
                                                    }
                                                    v65 := GenS2{F0: ("\n\t" + ((("" + "foo bar") + "a") + "hello")), F1: 42}
                                                    _ = v54
                                                    _ = v65
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = v3
                                                }
                                                fmt.Println(v1)
                                                _ = fi53
                                            }
                                        }
                                    default:
                                        v5 -= 746
                                    }
                                    v4 = append(v4, ("a" + ("\n\t" + ("\n\t" + ""))))
                                    _ = v52
                                } else {
                                    v66 := GenS4{F0: v18}
                                    _ = v66
                                }
                            }
                            if v18 {
                                if false {
                                    var v67 uint32 = v5
                                    _ = len(v4)
                                    for fi68 := 0; fi68 < 1; fi68++ {
                                        var v69 uint32 = (v3 & (v5 >> ((354) & 63)))
                                        _ = fi68
                                        _ = v69
                                    }
                                    _ = v67
                                }
                                for fi70 := 0; fi70 < 11; fi70++ {
                                    v4 = append(v4, (("" + "a") + "\n\t"))
                                    _ = fi70
                                }
                            }
                        } else {
                            if (65535 != 2) {
                                v71 := GenS7{F0: 42, F1: 18446744073709551615}
                                _ = v71
                            } else {
                                var v72 float32 = float32(3.14159)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v1 -= v1
                                _ = v72
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi37
                    }
                    if len(v2) > 0 {
                        v2[0] = (v5 * (v5 ^ v5))
                    }
                } else {
                    for fi73 := 0; fi73 < 7; fi73++ {
                        var v74 uint8 = 255
                        _ = fi73
                        _ = v74
                    }
                }
                fmt.Println(v5, v18)
            }
        case "hello":
            v1 = v1
        case "\n\t":
            for fi75 := 0; fi75 < 7; fi75++ {
                for fi76 := 0; fi76 < 13; fi76++ {
                    fmt.Println(v3, v18, v18)
                    v77 := GenS1{F0: "a", F1: -108}
                    var v78 uint = 0
                    _ = fi76
                    _ = v77
                    _ = v78
                }
                v1 = 9223372036854775807
                v79 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar"), F1: -30}
                _ = fi75
                _ = v79
            }
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
