package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint }
type GenS5 struct { F0 int64; F1 uint; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int16; F2 int }
type GenS8 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 uint16, p1 uint8) uint64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint8 = 135
        _ = fi6
        _ = v7
    }
    v8 := []int{0}
    for fi9 := 0; fi9 < 15; fi9++ {
        _ = len(v8)
        var v10 uint32 = 42
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 4294967295
}
func genFn1() uint {
    var v11 float32 = 3.14159
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int16 = 32767
        v11 = v11
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := []int{fi12, (-(fi12)), (fi12 >> ((18446744073709551615) & 63)), 4294967295}
            v15 = append(v15, fi14)
            fmt.Println(v11, fi14, fi12)
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 float32 = ((v11 - float32(-2.25)) + (float32((((0.0 + 1.5) * 1.5) / 1.5)) / (v11 - v11)))
        _ = fi16
        _ = v17
    }
    v11 = ((v11 / 1.0) - v11)
    v18 := GenS1{F0: -32768}
    _ = v11
    _ = v18
    return 65535
}

func main() {
    var v1 uint = 1
    var v2 bool = true
    v3 := make(map[string]int)
    var v4 int32 = -324
    v5 := make(map[int]int)
    if v2 {
        v19 := GenS1{F0: 1}
        var v20 int16 = 941
        switch "" {
        case "a":
            var v21 int64 = 144
            _ = v21
        case "hello":
            v2 = (v4 < (v4 ^ -1))
        case "foo bar":
            fmt.Println(v20, v2)
        }
        _ = v19
        _ = v20
    }
    v4 = (^(v4))
    if (((-1 != 9223372036854775807) || true) && v2) {
        for fi22 := 0; fi22 < 2; fi22++ {
            if true {
                var v23 uint = (v1 >> ((v1) & 63))
                var v24 int32 = -2147483648
                if ((!(42 <= 2)) || v2) {
                    switch "foo bar" {
                    case "":
                        v25 := GenS1{F0: 1000}
                        _ = v25
                    case "\n\t":
                        v26 := GenS1{F0: -32768}
                        _ = v26
                    case "hello":
                        if (0 >= 768) {
                            v27 := GenS2{F0: -9223372036854775808, F1: "hello", F2: (^(v1))}
                            _ = v27
                        }
                    default:
                        switch (fi22 | fi22) {
                        case -1:
                            var v28 int = fi22
                            _ = v28
                        case -9223372036854775808:
                            if ((3.14159 - 1e10) < 3.14159) {
                                switch fi22 {
                                case 1:
                                    for fi29 := 0; fi29 < 18; fi29++ {
                                        v23 -= v1
                                        fmt.Println(v4, v1)
                                        _ = fi29
                                    }
                                case 4294967295:
                                    v4 += (v24 >> (((v23 - v1)) & 63))
                                case 9223372036854775807:
                                    v30 := GenS1{F0: -32768}
                                    _ = v30
                                default:
                                    v31 := GenS3{F0: 2}
                                    _ = v31
                                }
                                fmt.Println(v4, v24, v1)
                            } else {
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    var v33 string = "foo bar"
                                    v33 = v33
                                    _ = fi32
                                    _ = v33
                                }
                            }
                        default:
                            delete(v5, fi22)
                        }
                    }
                    v2 = (true && (2 >= 42))
                } else {
                    var v34 int8 = 127
                    _ = v34
                }
                _ = v23
                _ = v24
            }
            v35 := []int{fi22, (fi22 & (-(fi22))), fi22}
            fmt.Println(v4, v2)
            _ = fi22
            _ = v35
        }
        v2 = (!false)
        var v36 uint32 = 4
        _ = v36
    }
    var v37 uint16 = 606
    var v38 int = 9223372036854775807
    switch "" {
    case "\n\t":
        for fi39 := 0; fi39 < 7; fi39++ {
            v40 := v5[239]
            _ = fi39
            _ = v40
        }
    case "foo bar":
        switch "hello" {
        case "\n\t":
            v4 = (^(v4))
        case "foo bar":
            if v2 {
                fmt.Println(v2, v4)
            }
        case "a":
            if v2 {
                v1 = 543
                switch (v38 * (^(v38))) {
                case 1:
                    if (!false) {
                        if (v2 || (v2 && true)) {
                            v41 := []int{v38, v38}
                            if ((v37 - v37) <= (^(v37))) {
                                fmt.Println(v2, v37)
                                v42 := GenS4{F0: 1, F1: 62, F2: (^(v1))}
                                v43 := GenS1{F0: 32767}
                                _ = v42
                                _ = v43
                            }
                            _ = v41
                        } else {
                            switch v38 {
                            case 157:
                                switch ("" + "a") {
                                case "a":
                                    v44 := GenS5{F0: -109, F1: 18446744073709551615, F2: ((((("hello" + "hello") + "hello") + "foo bar") + ("hello" + ("a" + "\n\t"))) + "a")}
                                    _ = v44
                                case "foo bar":
                                    v37 = (v37 - v37)
                                case "":
                                    for fi45 := 0; fi45 < 16; fi45++ {
                                        v46 := GenS3{F0: v37}
                                        fmt.Println(v1, v37, v38)
                                        _ = fi45
                                        _ = v46
                                    }
                                }
                            case 917:
                                if (!(v2 || false)) {
                                    switch 4294967295 {
                                    case 883:
                                        switch ("" + "foo bar") {
                                        case "foo bar":
                                            v47 := GenS4{F0: -32768, F1: 127, F2: genFn1()}
                                            _ = v47
                                        case "hello":
                                            for fi48 := 0; fi48 < 11; fi48++ {
                                                v49 := []int{427}
                                                var v50 float64 = ((((1e10 / (1e10 * float64((0.1 + 3.14159)))) / (-2.25 - 0.0)) * float64(0.1)) / 1.0)
                                                v37 -= 0
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                    default:
                                        var v51 int8 = 100
                                        _ = v51
                                    }
                                    fmt.Println(v37, v1, v38)
                                } else {
                                    for fi52 := 0; fi52 < 15; fi52++ {
                                        v2 = ("" != ("hello" + "foo bar"))
                                        if (127 == -85) {
                                            v53 := GenS2{F0: -1, F1: "foo bar", F2: (v1 >> ((v1) & 63))}
                                            _ = v53
                                        }
                                        _ = fi52
                                    }
                                    fmt.Println(v38, v1, v4)
                                }
                            default:
                                switch ("\n\t" + ("\n\t" + "")) {
                                case "":
                                    v4 += v4
                                default:
                                    if v2 {
                                        var v54 uint8 = 1
                                        v55 := GenS1{F0: -32768}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 uint8 = 0
                                        for fi57 := 0; fi57 < 18; fi57++ {
                                            v4 = (v4 & v4)
                                            v56 = (v56 * (v56 & v56))
                                            for fi58 := 0; fi58 < 16; fi58++ {
                                                if false {
                                                    var v59 uint32 = 2
                                                    switch v38 {
                                                    case 941:
                                                        v4 = -2147483648
                                                    default:
                                                        var v60 uint64 = 1
                                                        _ = v60
                                                    }
                                                    v61 := GenS6{F0: ("\n\t" + "")}
                                                    v62 := []int{(fi57 | fi58), fi58}
                                                    if v2 {
                                                        v63 := GenS7{F0: (1 != 2), F1: -384, F2: fi58}
                                                        var v64 int64 = 0
                                                        var v65 uint16 = 65535
                                                        fmt.Println(v1, v64, v59)
                                                        var v66 uint32 = (v59 ^ v59)
                                                        var v67 int = 9223372036854775807
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        if len(v62) >= 2 {
                                                            v62 = v62[1:len(v62)]
                                                        }
                                                        v59 = v59
                                                        if (!v2) {
                                                            var v68 uint16 = v37
                                                            fmt.Println(v56)
                                                            _ = v68
                                                        } else {
                                                            var v69 float32 = 0.0
                                                            _ = v69
                                                        }
                                                        v4 = (v4 ^ v4)
                                                    }
                                                    switch (("foo bar" + "\n\t") + "\n\t") {
                                                    case "hello":
                                                        v59 = (v59 << ((1) & 63))
                                                    case "\n\t":
                                                        v70 := GenS8{F0: false, F1: 2, F2: (v59 / (((v59 >> ((77) & 63))) | 1))}
                                                        _ = v70
                                                    case "":
                                                        v4 = (v4 + v4)
                                                        v2 = v2
                                                        if len(v62) > 0 {
                                                            v62[0] = -9223372036854775808
                                                        }
                                                    }
                                                    _ = v59
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = fi58
                                            }
                                            _ = fi57
                                        }
                                        _ = v56
                                    }
                                }
                            }
                        }
                        for fi71 := 0; fi71 < 14; fi71++ {
                            switch 2147483647 {
                            case -2147483648:
                                switch (v38 % (((fi71 << ((833) & 63))) | 1)) {
                                case -38:
                                    switch 127 {
                                    case -1:
                                        fmt.Println(v2, v2)
                                    case 127:
                                        v1 = genFn1()
                                    case 1:
                                        var v72 int16 = -32768
                                        _ = v72
                                    default:
                                        v73 := []int{32767, fi71, (-(v38)), fi71}
                                        _ = v73
                                    }
                                case 9223372036854775807:
                                    v74 := []int{v38, 42, v38, (fi71 & int(v1))}
                                    _ = v74
                                case -2147483648:
                                    for fi75 := 0; fi75 < 15; fi75++ {
                                        var v76 uint32 = 4294967295
                                        v1 = v1
                                        if false {
                                            v77 := []int{712, fi71}
                                            var v78 string = "hello"
                                            var v79 uint32 = v76
                                            v38 -= (v38 + 9223372036854775807)
                                            if v2 {
                                                if v2 {
                                                    var v80 uint64 = uint64(v76)
                                                    v3["a"] = v38
                                                    _ = v80
                                                }
                                                _ = len(v77)
                                            }
                                            _ = v77
                                            _ = v78
                                            _ = v79
                                        } else {
                                            if v2 {
                                                v81 := GenS1{F0: 1}
                                                v2 = ((v2 || v2) && (!(v2 && (((!v2) || ("hello" > "hello")) || ((-128 == 119) || v2)))))
                                                var v82 int16 = 209
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                        if false {
                                            if ((!(v2 && false)) || v2) {
                                                v83 := []int{}
                                                if len(v83) > 0 {
                                                    v83[0] = (v38 >> (((v1 >> ((v1) & 63))) & 63))
                                                }
                                                _ = v83
                                            }
                                            v84 := GenS9{F0: (v37 >> (((^(v1))) & 63))}
                                            switch "" {
                                            case "hello":
                                                v38 = fi75
                                            default:
                                                v85 := []int{fi75, (fi71 ^ 32768), fi71, (-(v38))}
                                                _ = v85
                                            }
                                            _ = v84
                                        } else {
                                            v86 := v5[(v38 ^ v38)]
                                            var v87 string = "hello"
                                            _ = v86
                                            _ = v87
                                        }
                                        _ = fi75
                                        _ = v76
                                    }
                                }
                            case 4294967295:
                                for fi88 := 0; fi88 < 7; fi88++ {
                                    _ = len(v5)
                                    _ = fi88
                                }
                            default:
                                v89 := GenS10{F0: 184}
                                _ = v89
                            }
                            _ = fi71
                        }
                        v37 = 65535
                    }
                case 127:
                    var v90 int = v38
                    _ = v90
                case 484:
                    if v2 {
                        v4 = v4
                    } else {
                        v91 := GenS11{F0: (v2 && v2), F1: v37, F2: 29}
                        v92 := GenS7{F0: v2, F1: -438, F2: (v38 | (v38 | (^(v38))))}
                        v93 := v3[("hello" + ("a" + "foo bar"))]
                        _ = v91
                        _ = v92
                        _ = v93
                    }
                default:
                    if true {
                        for fi94 := 0; fi94 < 20; fi94++ {
                            v37 -= v37
                            fmt.Println(v37, v4)
                            _ = fi94
                        }
                        var v95 string = "hello"
                        _ = v95
                    }
                }
            }
        }
    case "hello":
        if (true || false) {
            v96 := []int{}
            switch (v38 | v38) {
            case -9223372036854775808:
                delete(v5, (v38 - (v38 >> (((^(v1))) & 63))))
            case 32767:
                v1 = (^(v1))
            case 4294967295:
                v2 = (v4 < v4)
            default:
                v4 += v4
            }
            _ = v96
        }
    }
    var v97 int8 = 56
    v98 := GenS5{F0: -9223372036854775808, F1: v1, F2: (("foo bar" + "a") + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v97
    _ = v98
}
