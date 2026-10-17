package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16; F1 int16; F2 int32 }
type GenS11 struct { F0 string; F1 uint16; F2 uint64 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint, p1 uint, p2 int32) string {
    p1 += (p0 + p1)
    var v7 uint8 = 164
    var v8 int8 = 1
    var v9 int8 = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint8) uint32 {
    if (!(100 >= 20)) {
        if (1.5 < (0.1 + 3.14159)) {
            switch -367 {
            case -1:
                v10 := GenS1{F0: 4294967295, F1: (!(false || (977 > -32768))), F2: p0}
                _ = v10
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                switch (fi11 / ((fi11) | 1)) {
                case 545:
                    v12 := []int{}
                    _ = v12
                default:
                    v13 := GenS2{F0: -32768, F1: (fi11 ^ (fi11 - int(p0))), F2: fi11}
                    _ = v13
                }
                v14 := GenS1{F0: 818, F1: true, F2: (p0 % ((p0) | 1))}
                switch fi11 {
                case 42:
                    var v15 bool = ((p0 < p0) && false)
                    _ = v15
                case -456:
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v17 := []int{(fi11 / ((-910) | 1)), 503}
                        _ = fi16
                        _ = v17
                    }
                default:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        if true {
                            p0 = (p0 ^ (p0 << ((451) & 63)))
                            if (0 > 1) {
                                switch fi18 {
                                case 4294967295:
                                    switch 9223372036854775807 {
                                    case -821:
                                        var v19 float32 = (3.14159 / 1.5)
                                        _ = v19
                                    default:
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                fmt.Println(fi11, fi18)
                                                var v20 uint16 = 65535
                                                v20 -= (v20 | 73)
                                                _ = v20
                                            }
                                        }
                                    }
                                case -9223372036854775808:
                                    v21 := []int{(^(fi18))}
                                    _ = v21
                                case -259:
                                    if false {
                                        v22 := []int{(fi18 << ((1) & 63)), (fi11 << ((1) & 63)), fi18, (fi18 & (-(fi11)))}
                                        _ = v22
                                    }
                                default:
                                    v23 := []int{fi11, 32768, (fi18 >> ((65535) & 63)), fi18}
                                    _ = v23
                                }
                                p0 = p0
                            } else {
                                v24 := []int{fi11, fi18, (fi18 | -2147483648)}
                                _ = v24
                            }
                        } else {
                            if (false || (false && (!true))) {
                                var v25 uint = 0
                                if false {
                                    var v26 float64 = (0.0 + (-2.25 - float64(((3.14159 + ((1.5 - 0.0) - -2.25)) * -2.25))))
                                    _ = v26
                                } else {
                                    p0 = (^(p0))
                                    v27 := []int{fi11, (fi18 - (fi18 | (fi11 + fi18))), fi11, (fi11 & (fi18 + (fi18 >> ((v25) & 63))))}
                                    _ = v27
                                }
                                switch "hello" {
                                case "hello":
                                    fmt.Println(p0)
                                default:
                                    if (0.1 != 0.0) {
                                        if false {
                                            if false {
                                                fmt.Println(p0, fi18, fi18)
                                                var v28 string = "hello"
                                                _ = v28
                                            } else {
                                                p0 = p0
                                                v29 := []int{fi11, (fi11 ^ (fi11 % (((fi11 - fi11)) | 1))), (fi18 * (^(fi11)))}
                                                v30 := GenS3{F0: true}
                                                if len(v29) >= 2 {
                                                    v29 = v29[1:len(v29)]
                                                }
                                                _ = v29
                                                _ = v30
                                            }
                                            v31 := GenS4{F0: 18446744073709551615, F1: 1}
                                            fmt.Println(fi18, fi11)
                                            var v32 int = (fi11 & 949)
                                            _ = v31
                                            _ = v32
                                        } else {
                                            switch (("foo bar" + (((("a" + "\n\t") + "") + "hello") + "a")) + "\n\t") {
                                            case "foo bar":
                                                v25 = v25
                                            default:
                                                fmt.Println(v25, p0)
                                            }
                                        }
                                        p0 -= p0
                                    } else {
                                        v33 := GenS4{F0: 503, F1: v25}
                                        var v34 int64 = 9223372036854775807
                                        var v35 int8 = 42
                                        switch (fi11 >> (((^(v25))) & 63)) {
                                        case 127:
                                            switch ("\n\t" + "\n\t") {
                                            case "foo bar":
                                                switch fi11 {
                                                case -1:
                                                    var v36 uint = (v25 / (((^(v25))) | 1))
                                                    _ = v36
                                                case 1:
                                                    var v37 float64 = 0.0
                                                    _ = v37
                                                case 127:
                                                    v38 := GenS2{F0: -310, F1: fi11, F2: fi11}
                                                    _ = v38
                                                default:
                                                    v34 = v34
                                                }
                                            case "hello":
                                                fmt.Println(fi11, v25, v35)
                                                p0 += (^(p0))
                                            case "a":
                                                v39 := GenS3{F0: (((float32((-2.25 + 3.14159)) * 0.0) - (float32(float64(1.5)) + 0.0)) != float32(3.14159))}
                                                _ = v39
                                            }
                                            var v40 int64 = v34
                                            _ = v40
                                        case -2147483648:
                                            v41 := []int{(fi18 % ((42) | 1))}
                                            _ = v41
                                        default:
                                            var v42 bool = false
                                            _ = v42
                                        }
                                        fmt.Println(p0)
                                        v43 := GenS5{F0: uint32(v25)}
                                        var v44 float32 = 0.1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                _ = v25
                            } else {
                                fmt.Println(p0)
                                v45 := GenS4{F0: 922, F1: 867}
                                v46 := GenS6{F0: -1}
                                _ = v45
                                _ = v46
                            }
                            v47 := []int{fi18}
                            fmt.Println(fi18, p0, fi11)
                            _ = v47
                        }
                        p0 = (p0 | 2)
                        _ = fi18
                    }
                }
                _ = fi11
                _ = v14
            }
            if (!true) {
                p0 = p0
                if (("hello" <= ("\n\t" + "\n\t")) && false) {
                    switch "foo bar" {
                    case "hello":
                        v48 := GenS7{F0: 9223372036854775807, F1: uint(p0), F2: 4294967295}
                        _ = v48
                    }
                    if ((((2 > 4294967295) || ("foo bar" <= ("a" + "\n\t"))) || ((-9223372036854775808 < -1) || true)) || false) {
                        switch -96 {
                        case 128:
                            v49 := GenS5{F0: 937}
                            _ = v49
                        case 32768:
                            var v50 int = 0
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 1; fi51++ {
                                for fi52 := 0; fi52 < 17; fi52++ {
                                    p0 = (p0 - 73)
                                    p0 += p0
                                    _ = fi52
                                }
                                if (65535 >= 406) {
                                    p0 -= p0
                                    p0 += 127
                                } else {
                                    if false {
                                        p0 = 127
                                        p0 = (^(p0))
                                    }
                                    if (!(78 >= 912)) {
                                        p0 -= (p0 - (p0 << ((2) & 63)))
                                        for fi53 := 0; fi53 < 3; fi53++ {
                                            var v54 int = -9223372036854775808
                                            p0 = p0
                                            v55 := []int{-284, 32768, (v54 << ((4294967295) & 63)), 1}
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                    } else {
                                        p0 = 1
                                        var v56 string = (("\n\t" + ("foo bar" + "\n\t")) + ((("a" + "foo bar") + "\n\t") + ("hello" + "hello")))
                                        _ = v56
                                    }
                                }
                                _ = fi51
                            }
                        }
                        fmt.Println(p0)
                    } else {
                        v57 := GenS8{F0: 96, F1: 4294967295}
                        fmt.Println(p0)
                        _ = v57
                    }
                    p0 = 23
                } else {
                    switch 42 {
                    case -492:
                        var v58 int32 = 2147483647
                        _ = v58
                    }
                    var v59 uint8 = p0
                    _ = v59
                }
                v60 := []int{-87, 1, 9223372036854775807}
                _ = v60
            } else {
                v61 := GenS8{F0: 4294967295, F1: -9223372036854775808}
                _ = v61
            }
        } else {
            p0 = (p0 << ((670) & 63))
            var v62 int32 = 42
            for fi63 := 0; fi63 < 4; fi63++ {
                switch "hello" {
                case "foo bar":
                    var v64 int8 = 0
                    _ = v64
                case "hello":
                    v65 := GenS3{F0: (!(!false))}
                    _ = v65
                default:
                    p0 -= p0
                }
                switch 4294967295 {
                case 945:
                    if (!(((("" + "\n\t") + "") <= ("a" + "")) || ("\n\t" > ("" + "")))) {
                        v62 = v62
                        v62 = (v62 * (v62 << ((145) & 63)))
                    } else {
                        v66 := []int{fi63, 459}
                        _ = v66
                    }
                case 32768:
                    v67 := GenS1{F0: 42, F1: (false || false), F2: (^(p0))}
                    _ = v67
                case 32767:
                    v62 -= (v62 << ((0) & 63))
                }
                _ = fi63
            }
            _ = v62
        }
        v68 := []int{-614, -2147483648, 1, 963}
        _ = v68
    }
    p0 -= p0
    var v69 float64 = (((0.0 / 1e10) / float64(-2.25)) * (1e10 + float64(1.5)))
    v70 := GenS9{F0: 42}
    _ = p0
    _ = v69
    _ = v70
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    v3 := []uint32{728, 42, 1}
    var v4 string = "\n\t"
    var v5 int = 0
    var v6 float64 = 1e10
    fmt.Println(v2)
    for fi71 := 0; fi71 < 15; fi71++ {
        if true {
            for fi72 := 0; fi72 < 5; fi72++ {
                if (0 >= -11) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch (v2 + "a") {
                    case "a":
                        fmt.Println(fi72)
                    default:
                        fmt.Println(v4, fi71, v6)
                    }
                    fmt.Println(v5)
                    v5 += fi71
                    var v73 int16 = 0
                    _ = v73
                } else {
                    var v74 uint16 = 5
                    var v75 uint16 = 32767
                    _ = v74
                    _ = v75
                }
                if (false && true) {
                    fmt.Println(fi71, fi71)
                    fmt.Println(v2)
                    if (v2 < (v4 + ((v4 + (v2 + v2)) + v4))) {
                        v76 := v1[fi72]
                        v77 := GenS10{F0: 1, F1: -700, F2: 1}
                        var v78 uint32 = 2147483647
                        _ = v76
                        _ = v77
                        _ = v78
                    } else {
                        v1[-1] = 32767
                        var v79 uint8 = 2
                        if (true && (-35 > int8(fi72))) {
                            switch (-(fi72)) {
                            case 2147483647:
                                fmt.Println(v2, v79, v79)
                                if false {
                                    var v80 string = v4
                                    v81 := v1[fi72]
                                    if (0 <= 4294967295) {
                                        if false {
                                            fmt.Println(fi72)
                                            v82 := GenS5{F0: 42}
                                            var v83 uint32 = 811
                                            var v84 uint64 = 1
                                            _ = v82
                                            _ = v83
                                            _ = v84
                                        } else {
                                            v3 = append(v3, 65535)
                                            switch v4 {
                                            case "":
                                                var v85 float64 = ((v6 + v6) / 1.5)
                                                var v86 int64 = 2
                                                _ = v85
                                                _ = v86
                                            case "\n\t":
                                                v79 -= v79
                                            default:
                                                var v87 int = (v81 * v81)
                                                _ = v87
                                            }
                                        }
                                        var v88 uint = 826
                                        _ = v88
                                    } else {
                                        v89 := GenS11{F0: "", F1: 1, F2: 4294967295}
                                        v81 = (fi72 + 42)
                                        _ = v89
                                    }
                                    _ = v80
                                    _ = v81
                                }
                                v4 = v2
                                var v90 int16 = 0
                                _ = v90
                            }
                        }
                        _ = v79
                    }
                } else {
                    fmt.Println(v4, fi71)
                }
                _ = fi72
            }
            var v91 int32 = -742
            v92 := GenS6{F0: 471}
            _ = v91
            _ = v92
        }
        switch v4 {
        case "\n\t":
            var v93 int64 = 9223372036854775807
            _ = v93
        case "":
            for fi94 := 0; fi94 < 10; fi94++ {
                v5 += (v5 + v5)
                if (!(true && true)) {
                    var v95 float32 = 1e10
                    if ((false && (true && false)) || false) {
                        v96 := GenS12{F0: (!(42 < 0)), F1: 4294967295}
                        var v97 string = v4
                        _ = v96
                        _ = v97
                    } else {
                        v98 := v1[fi71]
                        switch (v2 + v4) {
                        case "a":
                            fmt.Println(v2, fi94, v4)
                        case "hello":
                            if true {
                                var v99 int8 = 1
                                _ = v99
                            } else {
                                v6 = v6
                                if (0 >= 0) {
                                    v100 := GenS2{F0: -116, F1: -466, F2: fi71}
                                    var v101 string = ((v4 + (((("hello" + "a") + (v4 + v4)) + ("\n\t" + "\n\t")) + v2)) + v2)
                                    _ = v100
                                    _ = v101
                                }
                                var v102 int = 222
                                _ = v102
                            }
                        default:
                            v103 := GenS1{F0: genFn1(146), F1: ((-24 >= -1) || false), F2: 42}
                            _ = v103
                        }
                        v95 = (float32(((v6 - ((v6 * (1e10 - 1e10)) + -2.25)) - (((0.1 / (v6 * v6)) * (1.5 + v6)) - 0.1))) - v95)
                        _ = v98
                    }
                    _ = v95
                }
                _ = fi94
            }
        case "hello":
            v104 := GenS8{F0: 904, F1: 127}
            _ = v104
        }
        if false {
            v105 := GenS13{F0: 435, F1: 480}
            var v106 uint64 = 4294967295
            var v107 float64 = v6
            _ = v105
            _ = v106
            _ = v107
        }
        _ = fi71
    }
    var v108 int = (v5 + v5)
    switch ("" + (v4 + ((("\n\t" + (v4 + v2)) + v4) + v2))) {
    case "foo bar":
        v3 = append(v3, 2147483647)
    }
    fmt.Println(v4, v5, v108)
    v109 := GenS14{F0: 2, F1: 127}
    v3 = append(v3, 922)
    v110 := GenS3{F0: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
