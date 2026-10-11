package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 bool) int {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int8 = 127
        switch (fi4 + fi4) {
        case 32767:
            if ("a" > "") {
                v6 := GenS1{F0: 2}
                switch 32767 {
                case 128:
                    if p0 {
                        var v7 int32 = -2147483648
                        var v8 uint = 0
                        v9 := []int{9223372036854775807, fi4}
                        _ = v7
                        _ = v8
                        _ = v9
                    } else {
                        v10 := []int{}
                        for fi11 := 0; fi11 < 1; fi11++ {
                            p0 = ((v5 >> ((4294967295) & 63)) > (v5 % ((-128) | 1)))
                            switch (("foo bar" + "foo bar") + ("hello" + "\n\t")) {
                            case "\n\t":
                                var v12 uint8 = 137
                                _ = v12
                            case "foo bar":
                                v13 := GenS1{F0: 2}
                                v5 = (v5 / ((v5) | 1))
                                switch fi11 {
                                case 4294967295:
                                    v14 := GenS2{F0: 126, F1: 10}
                                    _ = v14
                                case 32768:
                                    v15 := GenS1{F0: 18446744073709551615}
                                    _ = v15
                                case -9223372036854775808:
                                    if len(v10) > 0 {
                                        v10[0] = -32
                                    }
                                }
                                var v16 uint = 603
                                var v17 bool = ((1.5 / 1.5) == float64(float32((float64(((1.5 * 0.1) / -2.25)) - 0.0))))
                                _ = v13
                                _ = v16
                                _ = v17
                            case "hello":
                                v5 = (v5 * v5)
                            }
                            v10 = append(v10, fi4)
                            _ = fi11
                        }
                        v5 = -128
                        _ = v10
                    }
                case -2147483648:
                    v5 = -99
                default:
                    fmt.Println(v5, v5, p0)
                }
                if (-2.25 >= (1.5 * 1e10)) {
                    if (!(!((((3.14159 - 3.14159) * (3.14159 - float32(((1.5 + ((1.5 + (0.0 * 3.14159)) - float64(((-2.25 + ((0.0 + 3.14159) / float32(1e10))) - 1.5)))) - 0.0)))) * 3.14159) != -2.25))) {
                        var v18 int64 = 42
                        _ = v18
                    }
                } else {
                    v19 := GenS1{F0: 1}
                    if ((!p0) && p0) {
                        p0 = p0
                    }
                    p0 = false
                    _ = v19
                }
                _ = v6
            }
        case -9223372036854775808:
            switch ("" + "a") {
            case "hello":
                for fi20 := 0; fi20 < 2; fi20++ {
                    if (!false) {
                        v5 += (v5 * v5)
                    } else {
                        v21 := []int{(fi20 << ((0) & 63)), (-(fi4))}
                        v21 = append(v21, fi4)
                        v21 = append(v21, fi4)
                        _ = v21
                    }
                    if p0 {
                        var v22 int64 = 0
                        _ = v22
                    } else {
                        var v23 int64 = 42
                        var v24 uint32 = 30
                        v25 := GenS1{F0: 633}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    v26 := []int{(fi20 - (fi4 * fi4)), fi4, (fi4 % ((538) | 1))}
                    _ = fi20
                    _ = v26
                }
            }
        case 32768:
            var v27 string = "\n\t"
            _ = v27
        }
        v28 := GenS3{F0: (0 == -288), F1: fi4}
        _ = fi4
        _ = v28
        _ = v5
    }
    p0 = (0 < 977)
    if false {
        var v29 string = ("hello" + "hello")
        if (831 <= 722) {
            var v30 int16 = 169
            switch 9223372036854775807 {
            case 270:
                v30 = 1000
            case 2147483647:
                p0 = false
            case 42:
                var v31 int = 1
                _ = v31
            }
            _ = v30
        }
        if p0 {
            if (65535 <= 1) {
                var v32 bool = (-2.25 < (((0.1 - float32(0.1)) / 1e10) * 0.0))
                _ = v32
            } else {
                v33 := GenS4{F0: 20, F1: 1, F2: p0}
                p0 = (v29 > (v29 + (v29 + (("hello" + ("a" + v29)) + v29))))
                v29 = v29
                _ = v33
            }
            var v34 int8 = -1
            v35 := []int{127, -9223372036854775808, -841, -9223372036854775808}
            _ = v34
            _ = v35
        }
        _ = v29
    } else {
        p0 = p0
        v36 := []int{0}
        _ = v36
    }
    _ = p0
    return 128
}
func genFn1() int64 {
    var v37 int8 = -79
    switch "foo bar" {
    case "":
        var v38 int8 = v37
        _ = v38
    case "\n\t":
        var v39 uint8 = 42
        _ = v39
    }
    switch 4294967295 {
    case 1:
        var v40 int = -600
        _ = v40
    case -324:
        v37 = v37
    }
    _ = v37
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    switch 9223372036854775807 {
    case 127:
        v3 = (v3 + (^(v3)))
    default:
        switch 2147483647 {
        case 42:
            switch 127 {
            case 1:
                switch "" {
                case "":
                    var v41 int8 = 100
                    _ = v41
                case "hello":
                    if ((-9223372036854775808 != -1) && ((-2.25 + 0.0) > 1e10)) {
                        v42 := GenS3{F0: (42 == -2147483648), F1: genFn0((!true))}
                        v3 = v3
                        _ = v42
                    } else {
                        v3 += v3
                    }
                case "\n\t":
                    var v43 uint16 = 42
                    _ = v43
                }
            }
        case -9223372036854775808:
            v3 = v3
        default:
            for fi44 := 0; fi44 < 8; fi44++ {
                switch "" {
                case "":
                    var v45 int64 = 42
                    _ = v45
                case "hello":
                    v46 := []int{}
                    _ = v46
                case "foo bar":
                    var v47 int8 = 6
                    _ = v47
                }
                _ = fi44
            }
        }
    }
    if true {
        var v48 uint16 = 65535
        switch genFn0(false) {
        case 128:
            for fi49 := 0; fi49 < 6; fi49++ {
                v50 := GenS4{F0: 127, F1: 1, F2: true}
                if false {
                    if false {
                        v51 := GenS5{F0: v48}
                        _ = v51
                    }
                    switch fi49 {
                    case 42:
                        v52 := GenS5{F0: v48}
                        _ = v52
                    case -9223372036854775808:
                        _ = len(v2)
                    }
                    var v53 uint8 = v3
                    _ = v53
                } else {
                    var v54 float64 = float64(1e10)
                    v55 := v2[("\n\t" + "")]
                    _ = v54
                    _ = v55
                }
                _ = fi49
                _ = v50
            }
        default:
            v56 := GenS6{F0: "a", F1: 255}
            _ = v56
        }
        _ = v48
    }
    v3 -= v3
    if ((9223372036854775807 == 1) && (955 > -9223372036854775808)) {
        var v57 float64 = 1.5
        fmt.Println(v3, v3)
        if true {
            var v58 bool = false
            _ = v58
        } else {
            v57 = v57
        }
        _ = v57
    }
    switch ("a" + "a") {
    case "foo bar":
        if (1 >= 127) {
            fmt.Println(v3)
            if false {
                var v59 int32 = 42
                switch "foo bar" {
                case "a":
                    var v60 int64 = genFn1()
                    _ = v60
                default:
                    v3 = (v3 | (v3 & 42))
                }
                _ = v59
            }
        }
    case "a":
        if (!true) {
            v61 := GenS7{F0: 2147483648, F1: 717}
            v3 -= (v3 | (v3 ^ (v3 & v3)))
            switch "a" {
            case "hello":
                v3 -= (v3 - (^(v3)))
            case "\n\t":
                if true {
                    v62 := v1[genFn0(true)]
                    var v63 int32 = -2147483648
                    _ = v62
                    _ = v63
                }
            default:
                v3 = v3
            }
            _ = v61
        }
    }
    v64 := GenS8{F0: -32768, F1: 2}
    v65 := []int{824, 32768}
    var v66 uint64 = 42
    var v67 int8 = 127
    fmt.Println(v67)
    for fi68 := 0; fi68 < 7; fi68++ {
        var v69 uint16 = 32767
        _ = fi68
        _ = v69
    }
    switch 843 {
    case 1:
        v70 := GenS6{F0: ("a" + "\n\t"), F1: 1}
        _ = v70
    case 4294967295:
        if ("hello" < (("" + "foo bar") + (("foo bar" + "") + ""))) {
            if false {
                v71 := v1[892]
                var v72 float32 = float32(((float64(float32(1e10)) / (1.5 + 1.5)) + 0.0))
                if (((0.0 * v72) - float32((0.1 + float64(float32(-2.25))))) != 0.0) {
                    var v73 uint = 255
                    v74 := GenS9{F0: v66}
                    _ = v73
                    _ = v74
                } else {
                    var v75 int8 = (v67 ^ (v67 + 100))
                    var v76 uint32 = 65535
                    for fi77 := 0; fi77 < 20; fi77++ {
                        var v78 int16 = 1000
                        v79 := GenS3{F0: (!false), F1: v71}
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                    _ = v75
                    _ = v76
                }
                _ = v71
                _ = v72
            }
        } else {
            switch 42 {
            case 1:
                for fi80 := 0; fi80 < 9; fi80++ {
                    if true {
                        v3 -= 127
                        if true {
                            if (!(true && true)) {
                                var v81 int = genFn0(true)
                                _ = v81
                            } else {
                                v67 = 1
                                switch "" {
                                case "a":
                                    fmt.Println(v3, fi80, v3)
                                default:
                                    v67 -= -48
                                }
                            }
                            if len(v65) >= 2 {
                                v65 = v65[1:len(v65)]
                            }
                        } else {
                            v3 = v3
                            for fi82 := 0; fi82 < 6; fi82++ {
                                fmt.Println(v3, fi82)
                                if false {
                                    v83 := GenS8{F0: -1, F1: 1}
                                    if (1e10 >= 0.1) {
                                        fmt.Println(fi82, v3, fi80)
                                    } else {
                                        v66 -= 0
                                        v65 = append(v65, fi80)
                                    }
                                    _ = v83
                                }
                                _ = fi82
                            }
                        }
                        v84 := GenS1{F0: 18446744073709551615}
                        _ = v84
                    } else {
                        var v85 int64 = genFn1()
                        fmt.Println(v85)
                        _ = v85
                    }
                    v67 -= 69
                    _ = fi80
                }
            case -326:
                if len(v65) > 0 {
                    v65[0] = 229
                }
            case -647:
                if (128 == -9223372036854775808) {
                    if true {
                        fmt.Println(v3, v3)
                    }
                }
            }
            v66 += v66
            if true {
                var v86 int = 9223372036854775807
                fmt.Println(v67, v3, v3)
                _ = v86
            } else {
                fmt.Println(v3, v66)
                var v87 uint16 = 947
                _ = v87
            }
        }
    default:
        fmt.Println(v3, v66)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
