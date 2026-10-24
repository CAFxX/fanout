package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 int; F2 int8 }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 uint16; F1 string; F2 uint32 }
type GenS11 struct { F0 int8 }
func genFn0(p0 uint32, p1 bool) string {
    fmt.Println(p1, p0)
    fmt.Println(p0)
    fmt.Println(p1)
    var v4 float32 = -2.25
    _ = p0
    _ = p1
    _ = v4
    return "foo bar"
}
func genFn1(p0 bool, p1 uint, p2 uint64) int32 {
    if p0 {
        var v5 uint = p1
        for fi6 := 0; fi6 < 13; fi6++ {
            switch "foo bar" {
            case "":
                if (1000 != 0) {
                    if (!true) {
                        var v7 float64 = (-2.25 + (1e10 / 1.0))
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 20; fi8++ {
                            p0 = (p0 || false)
                            _ = fi8
                        }
                        switch "foo bar" {
                        case "hello":
                            if p0 {
                                p0 = (!true)
                            } else {
                                fmt.Println(fi6)
                            }
                        case "":
                            for fi9 := 0; fi9 < 19; fi9++ {
                                v10 := []int{fi9}
                                p0 = p0
                                _ = fi9
                                _ = v10
                            }
                        }
                    }
                    v11 := GenS1{F0: (p0 || (p0 && p0))}
                    _ = v11
                } else {
                    v12 := GenS2{F0: "a", F1: 65535}
                    _ = v12
                }
            case "\n\t":
                var v13 bool = p0
                _ = v13
            case "a":
                var v14 int = (fi6 << ((v5) & 63))
                _ = v14
            }
            var v15 bool = ((fi6 & (-(fi6))) != fi6)
            if v15 {
                var v16 float64 = (-2.25 - (3.14159 + 0.0))
                var v17 float32 = (0.1 / (3.14159 - 1.5))
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS2{F0: "hello", F1: 2}
                    switch 2147483647 {
                    case 689:
                        var v20 int64 = -704
                        _ = v20
                    case 70:
                        var v21 int32 = 0
                        _ = v21
                    case 32767:
                        var v22 string = ("foo bar" + ("foo bar" + "\n\t"))
                        _ = v22
                    default:
                        var v23 float32 = -2.25
                        _ = v23
                    }
                    v24 := []int{32767, (fi18 + fi6), (fi6 % ((fi18) | 1)), 9223372036854775807}
                    _ = fi18
                    _ = v19
                    _ = v24
                }
                _ = v16
                _ = v17
            }
            _ = fi6
            _ = v15
        }
        _ = v5
    }
    if true {
        switch "" {
        case "foo bar":
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 float64 = 1e10
                p0 = p0
                var v27 int64 = 0
                _ = fi25
                _ = v26
                _ = v27
            }
        case "hello":
            var v28 uint16 = 0
            _ = v28
        case "":
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 float64 = -2.25
                _ = fi29
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS2{F0: ("hello" + (("foo bar" + "a") + "\n\t")), F1: 65535}
            fmt.Println(p1, p1, p2)
            p0 = (!(!(4294967295 != 4294967295)))
            _ = fi31
            _ = v32
        }
        p0 = false
    }
    var v33 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int16 = -1
    v34 := GenS3{F0: (v3 % ((v3) | 1)), F1: true}
    for fi35 := 0; fi35 < 17; fi35++ {
        fmt.Println(fi35, v3)
        var v36 uint64 = 18446744073709551615
        if (-658 < -9223372036854775808) {
            v37 := GenS2{F0: "", F1: 161}
            v3 = v3
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            if (!true) {
                var v40 uint32 = 85
                _ = v40
            }
            var v41 float64 = (3.14159 * ((float64(3.14159) / (3.14159 - (0.1 + ((float64(float32(1e10)) * (-2.25 - (1.5 * (-2.25 * 1.5)))) + (-2.25 - (3.14159 * -2.25)))))) - 3.14159))
            _ = fi39
            _ = v41
        }
        v42 := GenS4{F0: "\n\t", F1: 42}
        _ = fi38
        _ = v42
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        if false {
            v3 = v3
        }
        if false {
            var v44 uint16 = 65535
            if false {
                v45 := GenS5{F0: 0, F1: 255}
                _ = v45
            } else {
                for fi46 := 0; fi46 < 8; fi46++ {
                    var v47 uint16 = (v44 >> ((0) & 63))
                    v47 = v44
                    _ = fi46
                    _ = v47
                }
            }
            if (!(65535 >= 594)) {
                v1[(fi43 ^ fi43)] = 0
            }
            _ = v44
        }
        _ = fi43
    }
    v48 := GenS1{F0: false}
    for fi49 := 0; fi49 < 11; fi49++ {
        v50 := GenS6{F0: 255, F1: 1}
        var v51 int8 = -1
        v3 = 32767
        _ = fi49
        _ = v50
        _ = v51
    }
    fmt.Println(v3)
    switch "foo bar" {
    case "hello":
        if true {
            var v52 uint32 = 42
            _ = v52
        } else {
            if false {
                for fi53 := 0; fi53 < 4; fi53++ {
                    if false {
                        v3 += (v3 >> ((903) & 63))
                        v54 := GenS4{F0: "\n\t", F1: 65535}
                        var v55 int16 = 200
                        _ = v54
                        _ = v55
                    }
                    _ = fi53
                }
            }
            if ((359 == 652) && true) {
                if ("a" == "foo bar") {
                    if false {
                        fmt.Println(v3)
                        var v56 uint64 = 18446744073709551615
                        _ = v56
                    } else {
                        if false {
                            v57 := GenS7{F0: (v3 | v3)}
                            _ = v57
                        }
                        v58 := GenS8{F0: "a", F1: 9223372036854775807, F2: 0}
                        v3 = v3
                        _ = v58
                    }
                }
                var v59 int32 = -2147483648
                _ = v59
            } else {
                _ = len(v1)
                for fi60 := 0; fi60 < 17; fi60++ {
                    v61 := v1[2147483647]
                    if true {
                        var v62 uint8 = 0
                        v63 := GenS4{F0: ("a" + ("" + ((((((("a" + "a") + "\n\t") + "\n\t") + "hello") + "foo bar") + "\n\t") + "foo bar"))), F1: 42}
                        v61 = 4294967295
                        _ = v62
                        _ = v63
                    } else {
                        var v64 uint = 42
                        _ = v64
                    }
                    var v65 bool = ((false || true) || false)
                    _ = fi60
                    _ = v61
                    _ = v65
                }
            }
            if true {
                switch "" {
                case "":
                    v66 := GenS2{F0: "", F1: 42}
                    _ = v66
                }
                var v67 int32 = 1
                _ = v67
            } else {
                switch "" {
                case "hello":
                    v68 := GenS6{F0: 32767, F1: 2147483647}
                    _ = v68
                case "a":
                    for fi69 := 0; fi69 < 19; fi69++ {
                        for fi70 := 0; fi70 < 4; fi70++ {
                            v3 = (v3 >> ((18446744073709551615) & 63))
                            var v71 string = "a"
                            var v72 int8 = -81
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        _ = fi69
                    }
                default:
                    var v73 int64 = 9223372036854775807
                    _ = v73
                }
            }
        }
    default:
        for fi74 := 0; fi74 < 8; fi74++ {
            for fi75 := 0; fi75 < 11; fi75++ {
                if false {
                    fmt.Println(fi75, fi75, v3)
                    v3 = (v3 << ((65535) & 63))
                    v76 := GenS2{F0: "", F1: 255}
                    _ = v76
                }
                if true {
                    if ((-1 <= 127) && true) {
                        if true {
                            switch fi75 {
                            case 42:
                                switch (fi75 << ((42) & 63)) {
                                case 86:
                                    v77 := GenS7{F0: (v3 - -195)}
                                    _ = v77
                                case -737:
                                    if true {
                                        switch "\n\t" {
                                        case "a":
                                            var v78 int64 = 42
                                            _ = v78
                                        case "hello":
                                            if ((1 == 1) && false) {
                                                v79 := GenS1{F0: (127 != fi74)}
                                                if true {
                                                    v80 := []int{(fi74 & fi74)}
                                                    v3 -= 1000
                                                    _ = v80
                                                }
                                                if (!(!(0 < 0))) {
                                                    var v81 int16 = 32767
                                                    v82 := []int{fi74, (fi74 / (((-(fi74))) | 1)), fi74}
                                                    var v83 uint16 = 32767
                                                    _ = v81
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    if true {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v84 := GenS7{F0: (v3 & v3)}
                                                            _ = v84
                                                        case "\n\t":
                                                            v3 = v3
                                                        default:
                                                            fmt.Println(fi74)
                                                        }
                                                        var v85 float64 = 1.5
                                                        _ = v85
                                                    }
                                                }
                                                _ = v79
                                            } else {
                                                if true {
                                                    var v86 string = ""
                                                    v86 = v86
                                                    switch (("a" + v86) + ("\n\t" + v86)) {
                                                    case "a":
                                                        var v87 uint32 = 2147483647
                                                        _ = v87
                                                    case "":
                                                        v88 := GenS9{F0: (18446744073709551615 != 42), F1: 42}
                                                        _ = v88
                                                    }
                                                    _ = v86
                                                }
                                                v89 := GenS6{F0: 169, F1: 1}
                                                _ = v89
                                            }
                                        }
                                        var v90 uint16 = 32767
                                        _ = v90
                                    } else {
                                        fmt.Println(fi74)
                                        v3 = -823
                                        var v91 float64 = 0.0
                                        _ = v91
                                    }
                                    v3 += (v3 % ((v3) | 1))
                                default:
                                    var v92 uint8 = 255
                                    _ = v92
                                }
                            case 128:
                                switch (^(fi75)) {
                                case -1:
                                    var v93 int16 = (-(v3))
                                    _ = v93
                                }
                            case -643:
                                fmt.Println(fi74, v3, v3)
                            }
                            var v94 bool = false
                            _ = v94
                        }
                        v3 = v3
                    }
                    v95 := []int{-960, 9223372036854775807, (fi75 | fi75), (^(fi75))}
                    _ = v95
                }
                switch fi75 {
                case 9223372036854775807:
                    var v96 uint = 2
                    _ = v96
                }
                _ = fi75
            }
            switch "hello" {
            case "foo bar":
                v97 := GenS10{F0: 427, F1: "foo bar", F2: 42}
                _ = v97
            case "":
                v98 := GenS11{F0: 1}
                _ = v98
            }
            for fi99 := 0; fi99 < 4; fi99++ {
                fmt.Println(fi99, v3, fi74)
                var v100 float32 = 0.1
                _ = fi99
                _ = v100
            }
            _ = fi74
        }
    }
    if (42 != 4294967295) {
        for fi101 := 0; fi101 < 10; fi101++ {
            v3 = 1000
            for fi102 := 0; fi102 < 19; fi102++ {
                var v103 uint32 = 2
                var v104 int64 = -1
                var v105 uint = 18446744073709551615
                _ = fi102
                _ = v103
                _ = v104
                _ = v105
            }
            var v106 int = 167
            _ = fi101
            _ = v106
        }
        if (!(false && true)) {
            fmt.Println(v3)
            v3 += (^(v3))
        } else {
            fmt.Println(v3)
            v3 -= v3
        }
        v107 := GenS6{F0: 32767, F1: 42}
        _ = v107
    } else {
        v108 := GenS9{F0: (110 > 18446744073709551615), F1: 0}
        for fi109 := 0; fi109 < 17; fi109++ {
            v110 := GenS10{F0: 65535, F1: "hello", F2: 25}
            fmt.Println(fi109, v3)
            _ = fi109
            _ = v110
        }
        _ = v108
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v48
}
