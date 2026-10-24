package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint32; F1 int32; F2 uint64 }
func genFn0(p0 string) uint8 {
    switch 128 {
    case 0:
        var v7 float32 = 0.1
        _ = v7
    case 4294967295:
        for fi8 := 0; fi8 < 1; fi8++ {
            if true {
                v9 := []int{}
                fmt.Println(fi8, fi8)
                _ = v9
            } else {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint8 = 1
                    p0 = "hello"
                    if (0.1 >= 1.5) {
                        var v12 bool = false
                        switch ((("a" + p0) + (((p0 + (p0 + "hello")) + (p0 + p0)) + "")) + (p0 + p0)) {
                        case "a":
                            var v13 float64 = (3.14159 * (3.14159 / ((((-2.25 + 3.14159) + 3.14159) - (0.0 + 3.14159)) / ((1.5 / 1.0) * 3.14159))))
                            _ = v13
                        case "":
                            var v14 int16 = 0
                            _ = v14
                        }
                        if v12 {
                            v15 := []int{(fi10 | (fi8 / ((-9223372036854775808) | 1)))}
                            var v16 uint64 = 968
                            _ = v15
                            _ = v16
                        }
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                fmt.Println(p0, fi8)
                for fi17 := 0; fi17 < 8; fi17++ {
                    p0 = p0
                    if ((!true) || false) {
                        p0 = ("a" + ((p0 + ((((p0 + p0) + p0) + ((p0 + "foo bar") + (("" + "foo bar") + p0))) + "")) + ("a" + p0)))
                    }
                    _ = fi17
                }
            }
            fmt.Println(fi8, fi8)
            if true {
                v18 := []int{}
                switch fi8 {
                case 1:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        _ = len(v18)
                        switch (("\n\t" + p0) + p0) {
                        case "foo bar":
                            switch ((p0 + p0) + p0) {
                            case "\n\t":
                                p0 = p0
                            case "foo bar":
                                if true {
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                } else {
                                    v18 = append(v18, (fi8 % ((fi19) | 1)))
                                    v20 := GenS1{F0: (p0 + (((((p0 + p0) + ("" + "\n\t")) + "\n\t") + p0) + p0))}
                                    _ = v20
                                }
                            case "a":
                                var v21 uint64 = 245
                                _ = v21
                            }
                            v22 := GenS1{F0: "hello"}
                            p0 = p0
                            _ = v22
                        case "a":
                            v23 := GenS1{F0: (("\n\t" + p0) + ("\n\t" + p0))}
                            _ = v23
                        }
                        var v24 int16 = 32767
                        _ = fi19
                        _ = v24
                    }
                default:
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if true {
                            if len(v18) > 0 {
                                v18[0] = 32767
                            }
                        } else {
                            var v26 float64 = 3.14159
                            switch 0 {
                            case 32767:
                                fmt.Println(fi8)
                            case 42:
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            case 32768:
                                v27 := GenS1{F0: "hello"}
                                _ = v27
                            }
                            _ = v26
                        }
                        switch p0 {
                        case "a":
                            var v28 uint = 255
                            _ = v28
                        case "foo bar":
                            var v29 bool = false
                            _ = v29
                        case "\n\t":
                            fmt.Println(p0)
                        default:
                            p0 = "foo bar"
                        }
                        if (!(42 > 39)) {
                            var v30 string = p0
                            p0 = ("foo bar" + p0)
                            _ = v30
                        }
                        _ = fi25
                    }
                }
                _ = v18
            }
            _ = fi8
        }
    case -1:
        p0 = ((p0 + ((p0 + (((p0 + "hello") + p0) + p0)) + ("a" + "a"))) + "a")
    }
    fmt.Println(p0)
    v31 := GenS1{F0: p0}
    var v32 int8 = 20
    fmt.Println(p0, p0)
    _ = p0
    _ = v31
    _ = v32
    return 0
}
func genFn1(p0 int16) uint32 {
    v33 := GenS1{F0: "a"}
    v34 := GenS2{F0: 29}
    fmt.Println(p0)
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 int16 = -1
        _ = fi35
        _ = v36
    }
    switch -1 {
    case 1:
        for fi37 := 0; fi37 < 5; fi37++ {
            for fi38 := 0; fi38 < 13; fi38++ {
                v39 := GenS3{F0: (fi37 % ((fi37) | 1)), F1: true, F2: 1}
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    case 0:
        p0 -= p0
    }
    _ = p0
    _ = v33
    _ = v34
    return 0
}

func main() {
    var v1 int16 = -130
    var v2 uint = 255
    v3 := make(map[int]bool)
    var v4 uint8 = 2
    v5 := []string{"hello"}
    var v6 float32 = ((1e10 - 0.1) - 1e10)
    switch 496 {
    case 9223372036854775807:
        var v40 string = "foo bar"
        _ = v40
    case 127:
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    case -344:
        if ((v1 / ((v1) | 1)) == 32767) {
            v41 := GenS4{F0: -729}
            switch 32768 {
            case 9223372036854775807:
                var v42 float64 = float64(0.1)
                _ = v42
            case 919:
                fmt.Println(v4)
            default:
                v5 = append(v5, "\n\t")
            }
            _ = v41
        } else {
            switch -32 {
            case 4294967295:
                _ = len(v3)
            case 0:
                for fi43 := 0; fi43 < 1; fi43++ {
                    v5 = append(v5, "\n\t")
                    var v44 int8 = 42
                    _ = fi43
                    _ = v44
                }
            case 393:
                fmt.Println(v4, v2)
            default:
                if true {
                    _ = len(v5)
                    if (true && true) {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v46 := GenS2{F0: v2}
                            fmt.Println(v4, v4)
                            _ = fi45
                            _ = v46
                        }
                        if ("foo bar" >= "a") {
                            v47 := GenS3{F0: -2147483648, F1: (!false), F2: 236}
                            _ = v47
                        }
                        v48 := GenS5{F0: 0, F1: "\n\t"}
                        _ = v48
                    } else {
                        var v49 string = (((("\n\t" + "hello") + "a") + (("\n\t" + "hello") + ("a" + (("\n\t" + ("" + "foo bar")) + "foo bar")))) + "a")
                        switch -2147483648 {
                        case 128:
                            for fi50 := 0; fi50 < 2; fi50++ {
                                for fi51 := 0; fi51 < 15; fi51++ {
                                    if false {
                                        switch "a" {
                                        case "a":
                                            switch v49 {
                                            case "hello":
                                                v49 = ((v49 + v49) + (v49 + v49))
                                            default:
                                                v49 = (v49 + v49)
                                            }
                                        }
                                        fmt.Println(v1, v6, v4)
                                        fmt.Println(fi50)
                                    } else {
                                        var v52 uint32 = 715
                                        var v53 float32 = (v6 - v6)
                                        switch (fi51 ^ 32767) {
                                        case 9223372036854775807:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 127:
                                            switch "hello" {
                                            case "a":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        default:
                                            var v54 float64 = float64(v6)
                                            _ = v54
                                        }
                                        if (788 >= 764) {
                                            v55 := GenS6{F0: 1, F1: false}
                                            var v56 int = -1
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v4 += genFn0("foo bar")
                                            v6 = ((0.1 + (v53 / float32(3.14159))) + (v53 - v6))
                                        }
                                        v52 += 929
                                        var v57 uint32 = (v52 ^ (v52 & genFn1((v1 & v1))))
                                        _ = v52
                                        _ = v53
                                        _ = v57
                                    }
                                    var v58 float32 = float32(-2.25)
                                    _ = fi51
                                    _ = v58
                                }
                                var v59 int8 = -128
                                var v60 int64 = 0
                                _ = fi50
                                _ = v59
                                _ = v60
                            }
                        case -841:
                            if (!true) {
                                for fi61 := 0; fi61 < 13; fi61++ {
                                    v49 = "hello"
                                    v4 += (v4 ^ (v4 + (^(v4))))
                                    _ = fi61
                                }
                            }
                        case 2147483647:
                            v62 := GenS3{F0: 32767, F1: true, F2: 412}
                            _ = v62
                        default:
                            v3[1] = (((0.1 - -2.25) <= float64(v6)) || true)
                        }
                        _ = v49
                    }
                }
            }
            switch 127 {
            case 9223372036854775807:
                v2 += 307
            case 2147483647:
                for fi63 := 0; fi63 < 11; fi63++ {
                    var v64 int64 = 0
                    fmt.Println(v6)
                    v4 -= v4
                    _ = fi63
                    _ = v64
                }
            case -2147483648:
                var v65 int16 = -1
                _ = v65
            default:
                v66 := GenS3{F0: 2147483647, F1: false, F2: 18446744073709551615}
                _ = v66
            }
        }
    default:
        if true {
            var v67 int8 = 11
            var v68 uint = v2
            if false {
                var v69 int8 = -43
                _ = v69
            }
            _ = v67
            _ = v68
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v70 := GenS4{F0: 42}
    var v71 uint8 = v4
    if ((v71 / (((v71 & v71)) | 1)) > v4) {
        switch -612 {
        case -329:
            v72 := GenS1{F0: "\n\t"}
            _ = v72
        case 304:
            v2 = 1
        case -752:
            if len(v5) > 0 {
                v5[0] = "hello"
            }
        }
        if true {
            if (95 <= 1) {
                switch "" {
                case "\n\t":
                    if false {
                        for fi73 := 0; fi73 < 14; fi73++ {
                            v74 := GenS1{F0: "foo bar"}
                            _ = fi73
                            _ = v74
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v71 = (^(v71))
                        case "foo bar":
                            var v75 uint = 4294967295
                            _ = v75
                        default:
                            for fi76 := 0; fi76 < 1; fi76++ {
                                v77 := GenS4{F0: 148}
                                fmt.Println(v1, fi76, v2)
                                fmt.Println(v4)
                                _ = fi76
                                _ = v77
                            }
                        }
                        if true {
                            v78 := GenS5{F0: 0, F1: ""}
                            _ = v78
                        } else {
                            var v79 uint8 = (v71 ^ (v71 + v71))
                            var v80 float64 = 0.0
                            switch -923 {
                            case 128:
                                v6 = (1e10 + ((float32(v80) - 3.14159) / 1e10))
                            case 460:
                                var v81 int64 = 9223372036854775807
                                _ = v81
                            case 9223372036854775807:
                                v82 := GenS5{F0: 1, F1: "a"}
                                _ = v82
                            default:
                                v80 += v80
                            }
                            _ = v79
                            _ = v80
                        }
                    } else {
                        var v83 string = "\n\t"
                        _ = v83
                    }
                }
                fmt.Println(v2, v71)
                var v84 uint32 = genFn1(v1)
                _ = v84
            }
            if true {
                v71 += (v71 | (^(v4)))
                for fi85 := 0; fi85 < 14; fi85++ {
                    v4 += (v4 - 146)
                    var v86 uint8 = v71
                    var v87 int64 = 779
                    _ = fi85
                    _ = v86
                    _ = v87
                }
                var v88 int32 = 0
                _ = v88
            }
        }
        delete(v3, 32768)
    } else {
        v89 := GenS2{F0: 602}
        v90 := GenS7{F0: v4, F1: 1}
        _ = v89
        _ = v90
    }
    fmt.Println(v1, v6, v71)
    var v91 bool = false
    v92 := GenS3{F0: 127, F1: true, F2: 2}
    var v93 uint64 = 373
    v2 += 42
    if (!(v91 || (!(!((^(v1)) < v1))))) {
        if v91 {
            switch "\n\t" {
            case "hello":
                var v94 bool = (2147483647 != 42)
                _ = v94
            case "foo bar":
                var v95 int = 4294967295
                _ = v95
            case "a":
                for fi96 := 0; fi96 < 12; fi96++ {
                    var v97 bool = v91
                    var v98 int32 = 0
                    _ = len(v3)
                    _ = fi96
                    _ = v97
                    _ = v98
                }
            }
            switch 32767 {
            case 127:
                var v99 uint8 = (v71 << ((18446744073709551615) & 63))
                _ = v99
            }
        } else {
            v100 := GenS7{F0: (v4 + (v71 * genFn0((("a" + (((("\n\t" + "\n\t") + "") + "foo bar") + ("\n\t" + ""))) + (("\n\t" + ("a" + "a")) + (("\n\t" + "a") + "")))))), F1: v1}
            if (0.1 == float64(0.0)) {
                v101 := GenS5{F0: v93, F1: ("foo bar" + ("\n\t" + (("hello" + "hello") + (("\n\t" + ("a" + ((("\n\t" + ("\n\t" + ("\n\t" + "hello"))) + ("foo bar" + "foo bar")) + "hello"))) + "a"))))}
                if false {
                    v4 += v71
                    var v102 int8 = -127
                    switch "\n\t" {
                    case "foo bar":
                        v103 := GenS2{F0: (v2 % ((v2) | 1))}
                        _ = v103
                    default:
                        var v104 uint16 = 0
                        _ = v104
                    }
                    _ = v102
                }
                _ = v101
            }
            _ = len(v5)
            _ = v100
        }
        for fi105 := 0; fi105 < 4; fi105++ {
            v4 = (v71 ^ 128)
            v4 -= v71
            var v106 uint16 = 189
            _ = fi105
            _ = v106
        }
    } else {
        for fi107 := 0; fi107 < 4; fi107++ {
            var v108 int32 = 669
            for fi109 := 0; fi109 < 2; fi109++ {
                if true {
                    var v110 uint = v2
                    _ = v110
                }
                _ = fi109
            }
            v1 -= 1
            _ = fi107
            _ = v108
        }
        for fi111 := 0; fi111 < 20; fi111++ {
            switch (fi111 - 870) {
            case -2147483648:
                v112 := GenS4{F0: -1}
                _ = v112
            }
            fmt.Println(v71, v2)
            var v113 int32 = 552
            _ = fi111
            _ = v113
        }
    }
    _ = len(v5)
    fmt.Println(v2)
    v114 := GenS8{F0: genFn1(v1), F1: int32(v4), F2: uint64(v2)}
    var v115 float32 = v6
    v116 := GenS6{F0: 42, F1: v91}
    v115 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v91
    _ = v92
    _ = v93
}
