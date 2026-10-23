package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 int64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int16; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int32; F1 string }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 int32; F1 uint16; F2 int32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int64; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 32767
    v3 := make(map[string]int)
    var v4 int32 = -963
    switch "foo bar" {
    case "":
        v4 = -1
    case "hello":
        fmt.Println(v4, v2)
    default:
        switch ((("\n\t" + "foo bar") + "hello") + ("a" + "hello")) {
        case "hello":
            for fi5 := 0; fi5 < 12; fi5++ {
                switch "a" {
                case "foo bar":
                    var v6 uint8 = 29
                    _ = v6
                case "hello":
                    var v7 uint = 420
                    _ = v7
                case "\n\t":
                    v8 := []int{42, 0, (fi5 | fi5), 2147483647}
                    _ = v8
                default:
                    var v9 int8 = 42
                    _ = v9
                }
                _ = fi5
            }
        case "":
            if true {
                fmt.Println(v2, v2)
                switch 32767 {
                case 2147483647:
                    if false {
                        v10 := v3[("\n\t" + (("hello" + ("" + "hello")) + "hello"))]
                        v10 = (v10 | (v10 >> ((1) & 63)))
                        v4 += 172
                        _ = v10
                    }
                case 32768:
                    v11 := GenS1{F0: 1, F1: ("foo bar" + "")}
                    _ = v11
                default:
                    v2 = (v2 >> ((1) & 63))
                }
                if (20 < 127) {
                    v12 := GenS2{F0: "", F1: (^(v2))}
                    _ = v12
                }
            }
        case "\n\t":
            fmt.Println(v4, v2)
        }
    }
    if ("a" <= "\n\t") {
        var v13 int = int(v4)
        var v14 uint64 = 573
        v3["\n\t"] = int(v14)
        _ = v13
        _ = v14
    }
    var v15 int = 128
    if false {
        var v16 float32 = 1.5
        if true {
            switch "foo bar" {
            case "foo bar":
                v17 := GenS3{F0: 0, F1: v4}
                _ = v17
            case "\n\t":
                v18 := GenS3{F0: -22, F1: 499}
                _ = v18
            }
            v19 := []int{(v15 % ((v15) | 1)), -1, (v15 % ((-502) | 1))}
            if len(v19) > 0 {
                v19[0] = v15
            }
            _ = v19
        }
        _ = v16
    } else {
        v20 := []int{-1}
        for fi21 := 0; fi21 < 18; fi21++ {
            v20 = append(v20, (fi21 & (^(fi21))))
            v15 -= 0
            _ = fi21
        }
        _ = v20
    }
    v22 := v3[("hello" + ("foo bar" + "hello"))]
    v22 = (v15 ^ -536)
    if (42 <= 65535) {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int = (^(fi23))
            var v25 int64 = -703
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        if (-1 == v4) {
            var v26 int64 = 990
            switch v15 {
            case 42:
                v27 := GenS4{F0: (v2 + 42), F1: -1}
                _ = v27
            case 9223372036854775807:
                delete(v1, v15)
            }
            if false {
                v15 += (-(v15))
                _ = len(v3)
                if false {
                    v28 := GenS5{F0: v2, F1: 9223372036854775807, F2: 128}
                    v29 := GenS6{F0: int32(v2), F1: v2, F2: 628}
                    for fi30 := 0; fi30 < 6; fi30++ {
                        if (65535 < 0) {
                            if false {
                                var v31 float32 = 0.0
                                v32 := []int{v22, fi30, v22, (^(v15))}
                                fmt.Println(v15)
                                _ = v31
                                _ = v32
                            }
                            var v33 uint32 = 2
                            _ = v33
                        } else {
                            var v34 int8 = 100
                            var v35 int16 = 42
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                    }
                    _ = v28
                    _ = v29
                } else {
                    v36 := GenS7{F0: "hello"}
                    if true {
                        v37 := GenS2{F0: "", F1: 1}
                        if (4294967295 == 1) {
                            v38 := []int{(v15 ^ (-(v22))), (v22 / ((v15) | 1)), v22, v22}
                            fmt.Println(v15, v22, v4)
                            for fi39 := 0; fi39 < 20; fi39++ {
                                v40 := GenS6{F0: (v4 ^ (v4 + (-(v4)))), F1: (v2 ^ (v2 * (-(v2)))), F2: 2}
                                switch ("a" + "") {
                                case "":
                                    _ = len(v38)
                                case "a":
                                    v41 := GenS8{F0: ("\n\t" + "hello"), F1: (960 <= 242)}
                                    _ = v41
                                default:
                                    if (-2.25 <= (0.1 / float64((((3.14159 / ((3.14159 + 0.0) * (3.14159 - 1e10))) / 1.0) + 0.0)))) {
                                        v42 := v1[v22]
                                        if (false || (fi39 < v22)) {
                                            var v43 float64 = 1e10
                                            _ = v43
                                        }
                                        _ = v42
                                    }
                                }
                                v4 += v4
                                _ = fi39
                                _ = v40
                            }
                            _ = v38
                        } else {
                            v2 = v2
                            fmt.Println(v2)
                        }
                        _ = v37
                    }
                    _ = v36
                }
            } else {
                var v44 uint = 65535
                _ = v44
            }
            _ = v26
        }
        var v45 float64 = (-2.25 - 0.0)
        v45 = v45
        _ = v45
    }
    if true {
        v46 := GenS2{F0: "", F1: (v2 ^ (v2 << ((0) & 63)))}
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint16 = 310
            _ = fi47
            _ = v48
        }
        _ = v46
    } else {
        v49 := GenS5{F0: v2, F1: 1, F2: 127}
        var v50 uint32 = 65535
        _ = v49
        _ = v50
    }
    v51 := GenS4{F0: -170, F1: 1}
    switch v22 {
    case 1:
        if true {
            v22 -= v22
            v15 = (^(v22))
        }
    }
    var v52 int32 = v4
    switch v15 {
    case 9223372036854775807:
        fmt.Println(v22)
    case 128:
        v53 := GenS3{F0: 9223372036854775807, F1: (^(v52))}
        _ = v53
    case 364:
        v15 = (v15 - v15)
    default:
        v22 = (-(v22))
    }
    var v54 bool = (!("foo bar" <= ("a" + "")))
    switch -9223372036854775808 {
    case -2147483648:
        v15 = v22
    case 32768:
        if false {
            fmt.Println(v22, v15)
            if ((-(v52)) >= v4) {
                var v55 int = v22
                if (!v54) {
                    fmt.Println(v54, v15, v52)
                } else {
                    var v56 int = v55
                    var v57 uint8 = 88
                    if (((1e10 * 3.14159) * 1e10) != 1e10) {
                        var v58 bool = false
                        _ = v58
                    } else {
                        fmt.Println(v54, v4, v57)
                    }
                    _ = v56
                    _ = v57
                }
                if (3.14159 == 0.1) {
                    fmt.Println(v55, v54, v54)
                    v55 = (-(v55))
                    var v59 float32 = 0.0
                    _ = v59
                } else {
                    if v54 {
                        v60 := []int{v55, (^(v55)), v22}
                        for fi61 := 0; fi61 < 1; fi61++ {
                            _ = len(v60)
                            fmt.Println(v2, v54)
                            for fi62 := 0; fi62 < 15; fi62++ {
                                if v54 {
                                    if false {
                                        fmt.Println(v55)
                                        var v63 uint64 = 42
                                        if true {
                                            if (2 < 65535) {
                                                var v64 uint = 18446744073709551615
                                                if v54 {
                                                    fmt.Println(v54, v52, v2)
                                                    var v65 string = "hello"
                                                    _ = v65
                                                }
                                                v64 = (v64 & (v64 >> ((v64) & 63)))
                                                _ = v64
                                            } else {
                                                v66 := GenS4{F0: 130, F1: 100}
                                                fmt.Println(v63, v22)
                                                var v67 int8 = -23
                                                _ = v66
                                                _ = v67
                                            }
                                        } else {
                                            if (float32(((3.14159 * 1e10) / 1.0)) >= (1.5 + (1.5 - (0.0 - (-2.25 / -2.25))))) {
                                                v3[""] = (fi61 - (fi62 + 32767))
                                                v2 -= v2
                                                var v68 uint = 741
                                                v68 = 255
                                                delete(v1, (v15 - (fi61 ^ v55)))
                                                _ = v68
                                            } else {
                                                switch 4294967295 {
                                                case 4294967295:
                                                    v55 -= fi61
                                                case 127:
                                                    v63 = 2
                                                case 881:
                                                    v69 := GenS9{F0: 786, F1: -128}
                                                    _ = v69
                                                }
                                            }
                                        }
                                        _ = v63
                                    } else {
                                        if (!false) {
                                            v54 = v54
                                            v15 = fi62
                                            var v70 int = -1
                                            _ = v70
                                        } else {
                                            if len(v60) > 0 {
                                                v60[0] = (v22 & fi61)
                                            }
                                            switch ((("" + ("foo bar" + "")) + "\n\t") + (("foo bar" + "a") + "hello")) {
                                            case "":
                                                var v71 int16 = v2
                                                _ = v71
                                            }
                                            delete(v3, ("a" + ""))
                                            var v72 int64 = 42
                                            _ = v72
                                        }
                                        var v73 int = v22
                                        _ = v73
                                    }
                                    if (18446744073709551615 >= 255) {
                                        v52 += -119
                                    } else {
                                        v60 = append(v60, fi61)
                                        if len(v60) > 0 {
                                            v60[0] = 32767
                                        }
                                    }
                                } else {
                                    if v54 {
                                        if len(v60) >= 2 {
                                            v60 = v60[1:len(v60)]
                                        }
                                        var v74 float64 = (1e10 * ((1e10 - 3.14159) - float64(1e10)))
                                        if (-377 >= (v2 | v2)) {
                                            var v75 float64 = v74
                                            if (!true) {
                                                if v54 {
                                                    var v76 int8 = 127
                                                    switch v55 {
                                                    case 0:
                                                        var v77 uint = 18446744073709551615
                                                        _ = v77
                                                    case 618:
                                                        v1[v55] = (("\n\t" + "hello") + "")
                                                    case 32768:
                                                        if ((!(-252 >= (v4 * -2147483648))) || false) {
                                                            v78 := GenS3{F0: 0, F1: 42}
                                                            var v79 int8 = -128
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                    }
                                                    _ = v76
                                                } else {
                                                    var v80 string = ""
                                                    v81 := GenS4{F0: (v2 / ((v2) | 1)), F1: -1}
                                                    switch ((v80 + v80) + v80) {
                                                    case "a":
                                                        _ = len(v60)
                                                        v82 := v3[(((v80 + v80) + v80) + (v80 + "\n\t"))]
                                                        v52 += (v52 - v52)
                                                        _ = v82
                                                    default:
                                                        switch int(v2) {
                                                        case 128:
                                                            v22 = fi62
                                                        case 4294967295:
                                                            var v83 uint32 = 2147483648
                                                            _ = v83
                                                        }
                                                        v74 = v74
                                                        if (255 >= 127) {
                                                            var v84 int16 = -1
                                                            if false {
                                                                var v85 float32 = 0.0
                                                                _ = v85
                                                            }
                                                            v52 -= v4
                                                            v84 += (v2 * int16(v4))
                                                            _ = v84
                                                        } else {
                                                            v54 = (-454 == (v2 - 42))
                                                            fmt.Println(v75, fi61, v2)
                                                        }
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                }
                                                fmt.Println(v15)
                                                v86 := GenS9{F0: 32767, F1: 108}
                                                _ = v86
                                            }
                                            _ = v75
                                        }
                                        _ = v74
                                    }
                                }
                                switch v55 {
                                case 32768:
                                    fmt.Println(v4, v15, v52)
                                case 127:
                                    v54 = (true || (false || true))
                                default:
                                    var v87 uint8 = 1
                                    _ = v87
                                }
                                _ = fi62
                            }
                            _ = fi61
                        }
                        var v88 int64 = 9223372036854775807
                        _ = v60
                        _ = v88
                    }
                    var v89 uint8 = 127
                    _ = v89
                }
                _ = v55
            } else {
                if ("foo bar" > ("foo bar" + "\n\t")) {
                    if v54 {
                        if (!(v2 >= (v2 + (v2 - (v2 - (v2 >> ((191) & 63))))))) {
                            for fi90 := 0; fi90 < 15; fi90++ {
                                v91 := GenS10{F0: (v4 & (-(v52))), F1: "\n\t"}
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    v15 = 800
                                    v52 -= (v4 + 0)
                                    v2 = (^(v2))
                                    _ = fi92
                                }
                                _ = fi90
                                _ = v91
                            }
                            var v93 uint32 = 2
                            var v94 int32 = (^(v52))
                            _ = v93
                            _ = v94
                        } else {
                            v54 = true
                        }
                        var v95 float32 = 1.5
                        fmt.Println(v4)
                        _ = v95
                    }
                } else {
                    v22 += (v22 - 127)
                    v96 := GenS11{F0: v2, F1: (v22 << ((4294967295) & 63))}
                    _ = v96
                }
            }
            if (2 >= 4294967295) {
                switch (v15 % (((v15 ^ v22)) | 1)) {
                case -2147483648:
                    var v97 int8 = 42
                    _ = v97
                }
                v98 := []int{v15}
                switch (v15 & 273) {
                case 1:
                    for fi99 := 0; fi99 < 20; fi99++ {
                        if v54 {
                            v52 = -2147483648
                            if true {
                                var v100 int64 = -1
                                var v101 uint = 433
                                _ = v100
                                _ = v101
                            } else {
                                if ((1.5 < 0.1) || (1 > 127)) {
                                    v22 = (v22 / ((-623) | 1))
                                    if len(v98) > 0 {
                                        v98[0] = -9223372036854775808
                                    }
                                }
                            }
                        } else {
                            switch ("" + "") {
                            case "foo bar":
                                if ((!(true || false)) && v54) {
                                    v1[(v15 + v15)] = "foo bar"
                                    var v102 string = ("a" + "")
                                    _ = v102
                                } else {
                                    var v103 int64 = 660
                                    var v104 int = (v22 >> ((0) & 63))
                                    for fi105 := 0; fi105 < 13; fi105++ {
                                        v106 := GenS4{F0: v2, F1: 0}
                                        v2 = 0
                                        v104 = (^(v15))
                                        fmt.Println(v4)
                                        _ = fi105
                                        _ = v106
                                    }
                                    _ = v103
                                    _ = v104
                                }
                            }
                            for fi107 := 0; fi107 < 19; fi107++ {
                                fmt.Println(v54, fi99)
                                v108 := GenS12{F0: (v4 % ((42) | 1)), F1: 65535, F2: (v52 % ((int32(v2)) | 1))}
                                v109 := v3["hello"]
                                if (9223372036854775807 < 0) {
                                    fmt.Println(v22)
                                    v110 := GenS7{F0: ""}
                                    v1[-1] = (("foo bar" + ("hello" + "a")) + "\n\t")
                                    _ = v110
                                }
                                _ = fi107
                                _ = v108
                                _ = v109
                            }
                            for fi111 := 0; fi111 < 6; fi111++ {
                                v112 := GenS13{F0: 128}
                                v113 := GenS14{F0: -1, F1: 255, F2: 2}
                                _ = fi111
                                _ = v112
                                _ = v113
                            }
                        }
                        var v114 int16 = v2
                        _ = fi99
                        _ = v114
                    }
                case 4294967295:
                    var v115 float64 = (((0.0 / 1.0) * (-2.25 - (-2.25 + 1.5))) - 1.5)
                    _ = v115
                default:
                    fmt.Println(v22)
                }
                _ = v98
            } else {
                v4 = (v4 / (((v52 % ((v52) | 1))) | 1))
                v116 := []int{v15}
                if len(v116) > 0 {
                    v116[0] = (v22 - v22)
                }
                _ = v116
            }
        } else {
            v2 = v2
            var v117 int32 = (-(v4))
            _ = v117
        }
    case 42:
        if false {
            var v118 float64 = (-2.25 * 0.0)
            var v119 uint8 = 1
            _ = v118
            _ = v119
        } else {
            for fi120 := 0; fi120 < 17; fi120++ {
                v121 := GenS2{F0: "\n\t", F1: (^(v2))}
                v1[-104] = ("a" + "\n\t")
                _ = fi120
                _ = v121
            }
            v122 := GenS1{F0: -84, F1: ("" + ("\n\t" + ((((("hello" + "foo bar") + "\n\t") + ("a" + "foo bar")) + "hello") + "hello")))}
            _ = v122
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v54
}
