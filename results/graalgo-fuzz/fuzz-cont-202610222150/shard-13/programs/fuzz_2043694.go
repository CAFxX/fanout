package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 string; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 int32; F2 uint }
type GenS12 struct { F0 int8; F1 uint32; F2 int }
type GenS13 struct { F0 string; F1 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint16; F1 uint }
type GenS16 struct { F0 int32; F1 uint16; F2 bool }
func genFn0(p0 int32) int16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 = (p0 - p0)
        v7 := GenS1{F0: 309}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    v8 := GenS2{F0: 1, F1: 0}
    v9 := GenS3{F0: 42}
    switch "" {
    case "foo bar":
        p0 += 42
    case "":
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint32 = 80
            _ = fi10
            _ = v11
        }
    case "a":
        v12 := GenS1{F0: 1}
        _ = v12
    }
    _ = p0
    _ = v8
    _ = v9
    return 115
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{0}
    var v3 int = 9223372036854775807
    var v4 int16 = 398
    v5 := make(map[string]int)
    if true {
        fmt.Println(v4, v3, v3)
        switch "\n\t" {
        case "hello":
            v4 -= v4
        case "":
            fmt.Println(v1, v4, v1)
        case "a":
            fmt.Println(v1, v4, v3)
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint16 = 42
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS1{F0: v4}
            var v17 float32 = (-2.25 * (0.0 - (0.1 + float32(0.1))))
            _ = fi15
            _ = v16
            _ = v17
        }
        if (true || true) {
            v1 = 2
        } else {
            var v18 float32 = -2.25
            if true {
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch "" {
                    case "":
                        var v20 uint8 = 0
                        _ = v20
                    case "a":
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v1 = 745
                            switch (^(v3)) {
                            case 4294967295:
                                v18 = ((0.1 * (1e10 * v18)) - 0.0)
                            case 1:
                                v22 := GenS1{F0: genFn0(-2147483648)}
                                delete(v5, "")
                                var v23 string = "\n\t"
                                _ = v22
                                _ = v23
                            case 224:
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v4)
                                }
                            default:
                                var v24 int16 = 42
                                _ = v24
                            }
                            _ = fi21
                        }
                    }
                    var v25 uint16 = 857
                    _ = fi19
                    _ = v25
                }
                fmt.Println(v3, v1, v1)
            } else {
                v26 := GenS4{F0: 65535, F1: "", F2: 0}
                var v27 string = ("\n\t" + "a")
                v1 = (v1 - (v1 * (v1 >> ((368) & 63))))
                _ = v26
                _ = v27
            }
            _ = v18
        }
    }
    var v28 uint8 = 2
    for fi29 := 0; fi29 < 18; fi29++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v3, v4)
        case "":
            var v30 int = v3
            _ = v30
        }
        _ = len(v5)
        _ = fi29
    }
    var v31 uint32 = 65535
    switch "\n\t" {
    case "foo bar":
        if true {
            switch 156 {
            case 4294967295:
                switch (v3 / ((v3) | 1)) {
                case 261:
                    for fi32 := 0; fi32 < 9; fi32++ {
                        if true {
                            v33 := GenS2{F0: 255, F1: 118}
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v3 = (v3 << ((65535) & 63))
                                v5[("" + "")] = (^(fi34))
                                v1 -= v1
                                _ = fi34
                            }
                            if true {
                                var v35 int32 = -2147483648
                                _ = v35
                            }
                            switch int(v1) {
                            case 128:
                                v36 := GenS5{F0: 42}
                                _ = v36
                            case 32767:
                                v37 := GenS5{F0: -9223372036854775808}
                                _ = v37
                            case 127:
                                var v38 int8 = 42
                                _ = v38
                            }
                        }
                        var v39 float32 = ((0.1 - (float32(-2.25) - ((3.14159 / 1e10) - 3.14159))) / 3.14159)
                        v4 = v4
                        _ = fi32
                        _ = v39
                    }
                default:
                    v40 := v5["foo bar"]
                    _ = v40
                }
            case 457:
                v31 -= (v31 * uint32(v1))
            default:
                v4 += (v4 * (v4 - v4))
            }
        }
    default:
        var v41 uint = 302
        _ = v41
    }
    if ((false && true) && true) {
        for fi42 := 0; fi42 < 18; fi42++ {
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 uint64 = v1
                _ = fi43
                _ = v44
            }
            for fi45 := 0; fi45 < 13; fi45++ {
                if true {
                    var v46 int32 = 2147483647
                    var v47 string = (("\n\t" + "foo bar") + "foo bar")
                    v46 = v46
                    _ = v46
                    _ = v47
                }
                switch "hello" {
                case "a":
                    if true {
                        var v48 uint8 = 236
                        v49 := GenS4{F0: 373, F1: ((("foo bar" + "\n\t") + "foo bar") + "hello"), F2: -9223372036854775808}
                        var v50 int32 = -88
                        _ = v48
                        _ = v49
                        _ = v50
                    } else {
                        var v51 uint64 = v1
                        if (255 > 18446744073709551615) {
                            var v52 int64 = 0
                            _ = len(v2)
                            v53 := GenS3{F0: 42}
                            _ = v52
                            _ = v53
                        } else {
                            var v54 bool = true
                            v55 := GenS6{F0: "", F1: 32767}
                            v54 = false
                            v56 := GenS6{F0: ("foo bar" + "hello"), F1: 2}
                            v51 -= v51
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        fmt.Println(v4, fi45, v28)
                        _ = v51
                    }
                }
                if (((((0.0 + -2.25) + ((float32(-2.25) + -2.25) / 1.0)) - (((0.1 + float32(0.0)) + 0.0) * 0.0)) <= float32((1.5 - 0.0))) || (uint64(v28) >= (v1 - v1))) {
                    var v57 uint = 939
                    if (((0.0 + -2.25) * 0.1) >= (float32(float64((3.14159 / ((0.0 + float32(0.0)) - float32(1e10))))) + 0.1)) {
                        var v58 int64 = -9223372036854775808
                        v31 = 350
                        var v59 float32 = float32((-2.25 + -2.25))
                        if false {
                            v31 -= (v31 >> (((v57 | (v57 & (v57 - 4294967295)))) & 63))
                            v57 = 18446744073709551615
                        } else {
                            if true {
                                var v60 int16 = 106
                                v61 := GenS7{F0: fi45}
                                var v62 int = (fi42 ^ 32767)
                                _ = v60
                                _ = v61
                                _ = v62
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (fi45 & v3)
                                }
                                switch (v3 % (((-(fi42))) | 1)) {
                                case 128:
                                    switch "" {
                                    case "\n\t":
                                        if (536 != 42) {
                                            delete(v5, "foo bar")
                                        } else {
                                            if (!(!true)) {
                                                fmt.Println(v58)
                                                fmt.Println(v57, fi42, fi45)
                                                v28 += v28
                                            } else {
                                                v31 -= v31
                                                v31 = (v31 ^ 2147483647)
                                                v58 -= v58
                                            }
                                            v63 := GenS4{F0: v57, F1: (("foo bar" + "\n\t") + "a"), F2: v58}
                                            fmt.Println(v31, v31, v1)
                                            switch "" {
                                            case "foo bar":
                                                v64 := GenS1{F0: genFn0(int32(fi45))}
                                                _ = v64
                                            }
                                            _ = v63
                                        }
                                        fmt.Println(v28)
                                        v57 = v57
                                        v65 := GenS8{F0: (("\n\t" + (("a" + ("hello" + ("hello" + "hello"))) + "")) + "\n\t"), F1: 4294967295}
                                        _ = v65
                                    case "hello":
                                        var v66 uint16 = 32767
                                        _ = v66
                                    case "foo bar":
                                        if (143 > v28) {
                                            fmt.Println(v4)
                                            v67 := GenS4{F0: (v57 * 18446744073709551615), F1: "hello", F2: v58}
                                            v68 := GenS9{F0: v28}
                                            _ = v67
                                            _ = v68
                                        } else {
                                            if false {
                                                fmt.Println(v58, fi42, v58)
                                                switch 0 {
                                                case 4294967295:
                                                    switch "" {
                                                    case "hello":
                                                        if (false || false) {
                                                            var v69 uint8 = v28
                                                            v3 = (v3 >> ((v57) & 63))
                                                            _ = v69
                                                        } else {
                                                            v4 = genFn0(0)
                                                            v4 = v4
                                                        }
                                                    case "\n\t":
                                                        delete(v5, ("" + "foo bar"))
                                                    case "foo bar":
                                                        v70 := GenS2{F0: (v28 ^ v28), F1: 0}
                                                        _ = v70
                                                    default:
                                                        var v71 uint32 = v31
                                                        _ = v71
                                                    }
                                                    v2 = append(v2, int(v58))
                                                case 2147483647:
                                                    switch (("\n\t" + "hello") + "hello") {
                                                    case "hello":
                                                        var v72 int64 = -665
                                                        _ = v72
                                                    case "foo bar":
                                                        var v73 string = (("" + ("foo bar" + "foo bar")) + ("a" + ""))
                                                        _ = v73
                                                    default:
                                                        v74 := GenS10{F0: genFn0(-2147483648)}
                                                        _ = v74
                                                    }
                                                default:
                                                    switch ("a" + ("foo bar" + "foo bar")) {
                                                    case "foo bar":
                                                        v59 += (1e10 + (v59 * 3.14159))
                                                        if true {
                                                            fmt.Println(v31, v3, v58)
                                                            v3 -= 128
                                                        }
                                                        var v75 uint = 698
                                                        _ = v75
                                                    case "a":
                                                        v76 := GenS11{F0: true, F1: -2147483648, F2: (v57 / (((v57 * (^(v57)))) | 1))}
                                                        _ = v76
                                                    default:
                                                        if false {
                                                            if (v58 >= v58) {
                                                                var v77 int64 = int64(v1)
                                                                var v78 float32 = (v59 / (v59 + float32(((3.14159 + (1.5 + (0.1 / 1.0))) / float64((v59 * v59))))))
                                                                switch fi42 {
                                                                case 32768:
                                                                    _ = len(v5)
                                                                case 1:
                                                                    v79 := GenS12{F0: 1, F1: (^(v31)), F2: int(v28)}
                                                                    _ = v79
                                                                default:
                                                                    v80 := GenS13{F0: "\n\t", F1: true}
                                                                    _ = v80
                                                                }
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            switch fi45 {
                                                            case 779:
                                                                var v81 float64 = 0.1
                                                                _ = v81
                                                            case -988:
                                                                v58 = 42
                                                                var v82 uint32 = (v31 * (v31 % ((v31) | 1)))
                                                                _ = v82
                                                            default:
                                                                fmt.Println(v57, v1, v57)
                                                            }
                                                        }
                                                    }
                                                }
                                                v83 := GenS12{F0: 1, F1: v31, F2: fi42}
                                                _ = v83
                                            } else {
                                                fmt.Println(v28)
                                                var v84 bool = true
                                                _ = v84
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v85 int64 = v58
                                            v3 = (fi45 * (fi45 * 445))
                                            _ = v85
                                        }
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                                var v86 int64 = v58
                                _ = v86
                            }
                            switch (("\n\t" + "foo bar") + "") {
                            case "\n\t":
                                v59 = v59
                            case "a":
                                var v87 float64 = 1.5
                                _ = v87
                            case "foo bar":
                                v59 = (v59 + v59)
                            }
                            var v88 uint8 = v28
                            switch "" {
                            case "hello":
                                fmt.Println(v28)
                            }
                            v89 := GenS14{F0: 1}
                            _ = v88
                            _ = v89
                        }
                        if (uint(v28) < v57) {
                            var v90 uint16 = 42
                            var v91 uint32 = (v31 << ((0) & 63))
                            _ = v90
                            _ = v91
                        }
                        _ = v58
                        _ = v59
                    }
                    fmt.Println(v1, v57, v4)
                    if true {
                        v3 = fi45
                    }
                    _ = v57
                } else {
                    var v92 int8 = 110
                    v3 -= fi45
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                    _ = v92
                }
                _ = fi45
            }
            _ = fi42
        }
        v3 = v3
    }
    for fi93 := 0; fi93 < 9; fi93++ {
        var v94 bool = true
        _ = fi93
        _ = v94
    }
    var v95 uint16 = 645
    for fi96 := 0; fi96 < 7; fi96++ {
        _ = len(v2)
        fmt.Println(v4)
        _ = fi96
    }
    v97 := GenS15{F0: v95, F1: 1}
    fmt.Println(v1, v28, v28)
    var v98 string = "hello"
    for fi99 := 0; fi99 < 14; fi99++ {
        v3 += -2147483648
        for fi100 := 0; fi100 < 12; fi100++ {
            if true {
                var v101 int16 = (^(v4))
                var v102 uint = 4294967295
                _ = v101
                _ = v102
            }
            _ = fi100
        }
        v98 = "a"
        _ = fi99
    }
    if ((-(v3)) >= v3) {
        for fi103 := 0; fi103 < 8; fi103++ {
            if (295 != 893) {
                v104 := GenS4{F0: 1, F1: "\n\t", F2: 1}
                var v105 float32 = -2.25
                v106 := GenS1{F0: 0}
                _ = v104
                _ = v105
                _ = v106
            } else {
                if false {
                    v107 := GenS8{F0: v98, F1: (v1 % (((^(v1))) | 1))}
                    var v108 uint32 = v31
                    v98 = "foo bar"
                    _ = v107
                    _ = v108
                }
            }
            for fi109 := 0; fi109 < 17; fi109++ {
                if false {
                    fmt.Println(v98, fi103, fi103)
                    v98 = ("hello" + "\n\t")
                } else {
                    fmt.Println(fi109)
                    switch v3 {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = fi109
            }
            _ = fi103
        }
    }
    if (-54 > int8(v3)) {
        fmt.Println(v4)
        fmt.Println(v98, v31, v3)
        v110 := GenS16{F0: 2147483647, F1: (v95 * uint16(v1)), F2: (v31 > v31)}
        _ = v110
    }
    v111 := GenS8{F0: "\n\t", F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v95:", v95)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v111
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v95
    _ = v97
    _ = v98
}
