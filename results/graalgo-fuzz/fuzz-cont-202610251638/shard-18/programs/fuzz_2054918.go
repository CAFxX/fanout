package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 string; F2 uint16 }
type GenS8 struct { F0 uint; F1 int32; F2 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint; F1 string; F2 bool }
type GenS11 struct { F0 uint64; F1 bool; F2 bool }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint8 }
func genFn0(p0 string) uint8 {
    var v4 uint64 = 4294967295
    var v5 string = ("" + p0)
    p0 = ("" + (((p0 + (v5 + p0)) + p0) + (p0 + ("foo bar" + p0))))
    var v6 uint64 = (^(v4))
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string, p1 int8, p2 int16) bool {
    if ((0.0 <= (0.0 * 3.14159)) && ((0.1 * (0.0 + (0.0 - (0.1 / 0.1)))) >= 0.1)) {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 14; fi8++ {
                fmt.Println(p1)
                var v9 string = (("a" + (p0 + p0)) + p0)
                _ = fi8
                _ = v9
            }
            var v10 int8 = -4
            _ = fi7
            _ = v10
        }
        switch "hello" {
        case "hello":
            if false {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if (42 <= 0) {
                        v12 := []int{}
                        _ = v12
                    }
                    _ = fi11
                }
                fmt.Println(p1)
                p1 = 100
            } else {
                v13 := []int{}
                for fi14 := 0; fi14 < 20; fi14++ {
                    if (128 <= 0) {
                        v15 := GenS1{F0: (-9223372036854775808 < 67), F1: true}
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            if (true || (255 == 2)) {
                                fmt.Println(p1, p1, p0)
                                var v17 int16 = (p2 / (((p2 % ((-32768) | 1))) | 1))
                                if (false || (1 != 695)) {
                                    var v18 string = "foo bar"
                                    _ = v18
                                } else {
                                    fmt.Println(v17, fi14, p1)
                                    v19 := GenS2{F0: 1}
                                    _ = v19
                                }
                                _ = v17
                            }
                            switch p0 {
                            case "foo bar":
                                if true {
                                    switch int(p1) {
                                    case -1:
                                        v20 := GenS3{F0: -1, F1: 1}
                                        _ = v20
                                    default:
                                        v21 := GenS4{F0: -1, F1: p2, F2: 2147483647}
                                        _ = v21
                                    }
                                    var v22 float64 = (float64(0.0) * 0.1)
                                    _ = v22
                                }
                            }
                            var v23 uint64 = 42
                            _ = fi16
                            _ = v23
                        }
                        var v24 uint32 = 2147483647
                        var v25 uint16 = 278
                        _ = v24
                        _ = v25
                    }
                    if (((p0 + ("foo bar" + "a")) + "a") > p0) {
                        for fi26 := 0; fi26 < 13; fi26++ {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            v27 := GenS2{F0: 0}
                            fmt.Println(p0, p0, fi14)
                            _ = fi26
                            _ = v27
                        }
                        if true {
                            v28 := GenS5{F0: p0}
                            _ = v28
                        } else {
                            p0 = "\n\t"
                            if true {
                                var v29 float64 = ((1.5 * -2.25) + (1e10 / (((float64(0.1) * 0.0) - 0.1) + 0.0)))
                                _ = v29
                            }
                        }
                    }
                    _ = fi14
                }
                _ = v13
            }
        case "\n\t":
            v30 := GenS6{F0: p0}
            _ = v30
        }
    } else {
        fmt.Println(p0, p2)
        if true {
            var v31 int32 = 1
            var v32 bool = (0 >= 0)
            _ = v31
            _ = v32
        }
    }
    switch 127 {
    case 1:
        fmt.Println(p0)
    }
    var v33 int16 = (p2 ^ -978)
    if (4294967295 >= 0) {
        v34 := GenS7{F0: int(v33), F1: "foo bar", F2: 590}
        _ = v34
    } else {
        if ((p1 | p1) < p1) {
            p2 -= (p2 >> ((4294967295) & 63))
            for fi35 := 0; fi35 < 12; fi35++ {
                p0 = p0
                fmt.Println(p0, v33, p2)
                _ = fi35
            }
        } else {
            p2 = (-(v33))
        }
    }
    v36 := []int{941}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v36
    return (!(0 < 1))
}

func main() {
    var v1 uint = 42
    v2 := []int{-9223372036854775808, 2147483647, 1}
    v3 := []string{("" + (("foo bar" + (("foo bar" + "foo bar") + "\n\t")) + "hello")), "a"}
    v1 -= (v1 + 65535)
    var v37 int16 = 1
    v37 += -459
    _ = len(v3)
    var v38 uint32 = 97
    v38 += v38
    for fi39 := 0; fi39 < 20; fi39++ {
        if (!((0 <= v38) || true)) {
            fmt.Println(v38, fi39)
            fmt.Println(v37, fi39, fi39)
            _ = len(v2)
        }
        var v40 bool = false
        if (v1 == v1) {
            var v41 int8 = 127
            if (((((-2.25 / (((((-2.25 - (float64((1.5 - float32(-2.25))) * 1e10)) - (float64(1e10) + ((float64(0.0) / 1.5) / 1.0))) + 1e10) * 1e10) * 1e10)) + (1.5 + (3.14159 / 1.0))) + (1.5 + 0.0)) * 1.5) >= -2.25) {
                v42 := GenS2{F0: -791}
                v1 -= 42
                _ = v42
            }
            v40 = (!false)
            _ = v41
        } else {
            v37 = (-(v37))
            v37 = (v37 >> ((924) & 63))
            if true {
                v40 = false
                fmt.Println(v40, v37, v1)
            } else {
                var v43 uint16 = 963
                _ = v43
            }
        }
        _ = fi39
        _ = v40
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v37 = 1
    default:
        var v44 uint = (v1 - 18446744073709551615)
        _ = v44
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 32768 {
    case 9223372036854775807:
        v37 = v37
    case 4294967295:
        switch "a" {
        case "\n\t":
            v37 = -1
        case "foo bar":
            fmt.Println(v1)
        }
    case 311:
        v45 := GenS1{F0: (-2147483648 == 1), F1: true}
        _ = v45
    default:
        if (!((!false) && (4294967295 <= 42))) {
            var v46 string = "a"
            for fi47 := 0; fi47 < 5; fi47++ {
                switch (v46 + (((((("\n\t" + v46) + v46) + (v46 + v46)) + v46) + "a") + "hello")) {
                case "hello":
                    for fi48 := 0; fi48 < 10; fi48++ {
                        switch fi48 {
                        case 32768:
                            if true {
                                switch (fi47 | 42) {
                                case 9223372036854775807:
                                    if len(v2) > 0 {
                                        v2[0] = (fi48 & (fi47 << ((65535) & 63)))
                                    }
                                }
                                v49 := GenS4{F0: 438, F1: 42, F2: -670}
                                fmt.Println(v1)
                                _ = len(v3)
                                _ = v49
                            }
                        case 32767:
                            var v50 int = (fi47 % ((fi48) | 1))
                            _ = v50
                        case -1:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi48
                    }
                case "":
                    for fi51 := 0; fi51 < 7; fi51++ {
                        var v52 string = v46
                        var v53 uint64 = 0
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                default:
                    for fi54 := 0; fi54 < 8; fi54++ {
                        v55 := GenS2{F0: -1}
                        if (false || true) {
                            var v56 uint16 = 2
                            if true {
                                switch v46 {
                                case "a":
                                    var v57 string = ""
                                    _ = v57
                                case "\n\t":
                                    var v58 uint64 = 2
                                    _ = v58
                                case "hello":
                                    if ((0.0 - (3.14159 + (float32(((float64(0.0) / 3.14159) + (0.1 / 1.0))) + ((0.1 * 0.1) / -2.25)))) >= 1e10) {
                                        var v59 int = (fi54 ^ fi54)
                                        if (v38 == 929) {
                                            var v60 int16 = (v37 + v37)
                                            v37 = (v60 ^ -32768)
                                            _ = v60
                                        } else {
                                            var v61 int8 = 100
                                            _ = v61
                                        }
                                        _ = v59
                                    } else {
                                        switch fi54 {
                                        case 288:
                                            if (true || true) {
                                                if false {
                                                    v38 -= (v38 - (v38 * 0))
                                                    var v62 float64 = 1.5
                                                    v63 := GenS1{F0: (v46 < (v46 + v46)), F1: true}
                                                    _ = v62
                                                    _ = v63
                                                }
                                                if false {
                                                    if (-2.25 != 1.5) {
                                                        v64 := GenS5{F0: v46}
                                                        v65 := GenS2{F0: -393}
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    v66 := GenS3{F0: (v37 / ((v37) | 1)), F1: 32767}
                                                    v67 := GenS3{F0: (-(v37)), F1: 173}
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    v68 := GenS6{F0: ""}
                                                    if (-32768 == 1000) {
                                                        if (!true) {
                                                            if true {
                                                                var v69 int8 = -127
                                                                if ((fi54 | (fi47 & fi47)) < (fi54 << ((939) & 63))) {
                                                                    if (368 >= (v56 / ((v56) | 1))) {
                                                                        v70 := GenS8{F0: v1, F1: 325, F2: 42}
                                                                        v71 := GenS9{F0: (^(v56))}
                                                                        _ = v70
                                                                        _ = v71
                                                                    } else {
                                                                        v69 -= v69
                                                                        if (true && true) {
                                                                            v69 = 42
                                                                            fmt.Println(fi54)
                                                                            var v72 uint32 = (v38 + 2)
                                                                            v69 = int8(v56)
                                                                            v69 = -77
                                                                            _ = v72
                                                                        } else {
                                                                            fmt.Println(fi54, fi54, v69)
                                                                        }
                                                                        var v73 int8 = (v69 / ((1) | 1))
                                                                        _ = v73
                                                                    }
                                                                    var v74 float32 = 0.0
                                                                    v69 -= v69
                                                                    _ = v74
                                                                }
                                                                v75 := GenS4{F0: 2147483647, F1: v37, F2: 42}
                                                                _ = v69
                                                                _ = v75
                                                            }
                                                            var v76 bool = false
                                                            fmt.Println(v76, fi54, v56)
                                                            _ = v76
                                                        } else {
                                                            v46 = v46
                                                            v38 += (v38 + (v38 * v38))
                                                        }
                                                        v46 = v46
                                                        var v77 uint8 = 195
                                                        _ = v77
                                                    }
                                                    var v78 int16 = v37
                                                    _ = v68
                                                    _ = v78
                                                }
                                                switch (fi54 | -1) {
                                                case 4294967295:
                                                    if len(v2) > 0 {
                                                        v2[0] = int(v56)
                                                    }
                                                default:
                                                    switch v46 {
                                                    case "a":
                                                        v79 := GenS10{F0: 42, F1: v46, F2: false}
                                                        _ = v79
                                                    case "":
                                                        v1 = uint(v56)
                                                        var v80 uint32 = v38
                                                        v38 = 698
                                                        _ = v80
                                                    default:
                                                        switch (v46 + v46) {
                                                        case "a":
                                                            v1 += v1
                                                        }
                                                    }
                                                }
                                            }
                                        case 32767:
                                            v46 = v46
                                        case 4294967295:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v81 := GenS6{F0: v46}
                                        _ = v81
                                    }
                                }
                                if ((((false && ((v37 == v37) || true)) || true) && ((^(v1)) >= 255)) && (42 < -852)) {
                                    v82 := GenS7{F0: fi54, F1: v46, F2: v56}
                                    v46 = ((("foo bar" + v46) + v46) + ("foo bar" + ((("" + v46) + "\n\t") + "a")))
                                    _ = v82
                                } else {
                                    v38 -= (v38 << ((343) & 63))
                                    if (!(false && false)) {
                                        _ = len(v3)
                                        v83 := GenS11{F0: 18446744073709551615, F1: (192 >= uint32(v56)), F2: (!false)}
                                        _ = v83
                                    } else {
                                        var v84 int16 = (v37 % ((v37) | 1))
                                        fmt.Println(v84, v1)
                                        var v85 string = ("\n\t" + ((v46 + v46) + v46))
                                        _ = v84
                                        _ = v85
                                    }
                                }
                                if (1.5 != -2.25) {
                                    switch (fi47 | 246) {
                                    case -729:
                                        var v86 bool = true
                                        _ = v86
                                    case 0:
                                        v1 -= 18446744073709551615
                                    case -1:
                                        v87 := GenS12{F0: 2}
                                        _ = v87
                                    }
                                } else {
                                    v88 := GenS13{F0: 110}
                                    if (v56 == v56) {
                                        var v89 string = ((v46 + "hello") + v46)
                                        if len(v2) > 0 {
                                            v2[0] = (-(fi47))
                                        }
                                        _ = v89
                                    }
                                    _ = v88
                                }
                                fmt.Println(v1, fi54)
                                if (-760 == -2147483648) {
                                    var v90 uint8 = 255
                                    _ = v90
                                }
                                v1 = (v1 + 381)
                            }
                            _ = v56
                        }
                        v37 = 1000
                        _ = fi54
                        _ = v55
                    }
                }
                fmt.Println(v37)
                _ = fi47
            }
            var v91 float64 = ((3.14159 - (0.1 / (((3.14159 - 1e10) * (0.1 * -2.25)) + (float64(1e10) * -2.25)))) - 3.14159)
            _ = v46
            _ = v91
        }
    }
    v92 := GenS9{F0: 938}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v92
}
