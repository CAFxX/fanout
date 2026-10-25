package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8; F1 string; F2 uint8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint; F1 uint64; F2 int32 }
type GenS15 struct { F0 int64; F1 string; F2 uint }
type GenS16 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int16 }
func genFn0() uint64 {
    var v7 int8 = 127
    var v8 float64 = 1e10
    var v9 int64 = -1
    switch "\n\t" {
    case "foo bar":
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(fi10, v8)
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 float32 = 1.5
                v9 += v9
                if true {
                    v13 := GenS1{F0: 0, F1: 1, F2: 127}
                    _ = v13
                } else {
                    fmt.Println(fi11)
                    fmt.Println(v8, v9)
                }
                v14 := GenS1{F0: 4294967295, F1: 1, F2: 57}
                _ = fi11
                _ = v12
                _ = v14
            }
            _ = fi10
        }
    case "\n\t":
        v15 := []int{325, 0, 1, -986}
        _ = v15
    case "":
        switch -885 {
        case 128:
            fmt.Println(v9, v7, v8)
        case 2147483647:
            v16 := []int{-284}
            _ = v16
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: 4294967295, F1: 0, F2: 0}
                _ = fi17
                _ = v18
            }
        }
    default:
        v19 := []int{32768}
        _ = v19
    }
    v20 := GenS2{F0: 0}
    _ = v20
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    var v4 string = "hello"
    var v5 bool = (104 != 42)
    var v6 uint64 = 1
    if (0 >= 286) {
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS1{F0: (v6 % (((^(v6))) | 1)), F1: 2147483647, F2: 42}
            _ = fi21
            _ = v22
        }
        var v23 float32 = 1e10
        _ = v23
    } else {
        var v24 int = -2147483648
        _ = v24
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint64 = 1
            var v28 uint = 2
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS3{F0: 128, F1: v2}
        _ = fi25
        _ = v29
    }
    var v30 bool = (!(false || (-128 >= 1)))
    v2 = (v2 | (v2 * -59))
    if true {
        switch (v4 + "") {
        case "a":
            v31 := GenS1{F0: (v6 + (^(v6))), F1: 0, F2: 42}
            _ = v31
        }
        if v30 {
            var v32 int64 = 9223372036854775807
            v6 = (^(v6))
            _ = v32
        } else {
            v33 := GenS2{F0: 128}
            v6 = (v6 | genFn0())
            _ = v33
        }
        if (v30 && v5) {
            for fi34 := 0; fi34 < 13; fi34++ {
                delete(v3, v4)
                _ = fi34
            }
            if v30 {
                _ = len(v3)
                var v35 bool = (1.5 == 0.0)
                var v36 string = "\n\t"
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 11; fi37++ {
                    fmt.Println(fi37, v4, v4)
                    v38 := GenS4{F0: v4}
                    for fi39 := 0; fi39 < 14; fi39++ {
                        v2 = (v2 << ((65535) & 63))
                        _ = fi39
                    }
                    _ = fi37
                    _ = v38
                }
            }
            for fi40 := 0; fi40 < 12; fi40++ {
                for fi41 := 0; fi41 < 2; fi41++ {
                    if (v5 && (v30 || (true || (float32(0.0) > 0.1)))) {
                        switch fi40 {
                        case -2147483648:
                            var v42 int32 = -2147483648
                            _ = v42
                        case -1:
                            fmt.Println(v5, v5, v6)
                        default:
                            var v43 bool = v5
                            _ = v43
                        }
                    } else {
                        v4 = ("a" + (((((v4 + "\n\t") + "") + "\n\t") + v4) + ("a" + (v4 + ("\n\t" + v4)))))
                        v44 := GenS5{F0: genFn0(), F1: 2, F2: 127}
                        switch (fi41 + (fi40 << ((2) & 63))) {
                        case 32768:
                            if (9223372036854775807 > -9223372036854775808) {
                                v45 := GenS6{F0: 127, F1: v30}
                                if (genFn0() < (v6 + (v6 ^ (v6 % (((^(v6))) | 1))))) {
                                    v46 := []int{fi41, 649, 954}
                                    v47 := GenS4{F0: (v4 + (v4 + v4))}
                                    _ = v46
                                    _ = v47
                                } else {
                                    switch fi40 {
                                    case -1:
                                        var v48 uint32 = v1
                                        _ = v48
                                    }
                                }
                                _ = v45
                            } else {
                                if (false && (fi41 == fi40)) {
                                    v49 := GenS7{F0: (v4 + (v4 + ""))}
                                    switch (fi41 << ((1) & 63)) {
                                    case 0:
                                        if (-9223372036854775808 >= fi40) {
                                            v4 = ("a" + "\n\t")
                                            switch ("" + v4) {
                                            case "":
                                                switch fi40 {
                                                case 32768:
                                                    fmt.Println(v2)
                                                }
                                            case "a":
                                                fmt.Println(v4)
                                            case "foo bar":
                                                v50 := GenS8{F0: v6, F1: v1, F2: 2147483648}
                                                _ = v50
                                            }
                                            if false {
                                                if v5 {
                                                    v51 := GenS8{F0: 0, F1: (v1 / ((42) | 1)), F2: (^(v1))}
                                                    if (fi41 < fi40) {
                                                        v5 = v30
                                                        v5 = (false && v30)
                                                    }
                                                    switch fi41 {
                                                    case -647:
                                                        v1 = v1
                                                    }
                                                    _ = v51
                                                }
                                                if v30 {
                                                    v2 = v2
                                                    delete(v3, "a")
                                                }
                                            } else {
                                                delete(v3, (((v4 + v4) + "\n\t") + v4))
                                                if v5 {
                                                    if true {
                                                        v52 := GenS9{F0: 520, F1: 1}
                                                        v6 += v6
                                                        fmt.Println(v1, v1, v30)
                                                        _ = v52
                                                    }
                                                }
                                                var v53 int64 = -805
                                                _ = v53
                                            }
                                        }
                                    case -2147483648:
                                        delete(v3, v4)
                                    default:
                                        v6 = genFn0()
                                    }
                                    if (int64(v2) < 9223372036854775807) {
                                        if true {
                                            switch (v4 + "\n\t") {
                                            case "hello":
                                                v1 -= 2147483648
                                            case "\n\t":
                                                var v54 int8 = 100
                                                _ = v54
                                            case "foo bar":
                                                fmt.Println(v2)
                                                var v55 string = v4
                                                if v5 {
                                                    v6 -= (v6 * v6)
                                                    delete(v3, v55)
                                                    var v56 bool = true
                                                    var v57 uint32 = (^(v1))
                                                    _ = v56
                                                    _ = v57
                                                }
                                                fmt.Println(v5)
                                                v58 := []int{fi40}
                                                _ = v55
                                                _ = v58
                                            default:
                                                v59 := GenS6{F0: 7, F1: (true || v5)}
                                                _ = v59
                                            }
                                        } else {
                                            if v30 {
                                                var v60 float32 = 1e10
                                                v6 = v6
                                                _ = v60
                                            } else {
                                                fmt.Println(v2, v5, v1)
                                                v4 = v4
                                                var v61 float64 = 0.1
                                                v30 = (-1 >= 275)
                                                _ = v61
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                var v62 int64 = 138
                                                v5 = (-2.25 >= 3.14159)
                                                switch ("foo bar" + (v4 + (((v4 + ("a" + (("hello" + ((v4 + (v4 + v4)) + v4)) + v4))) + (("\n\t" + ((v4 + v4) + "\n\t")) + "")) + (((v4 + v4) + v4) + "\n\t")))) {
                                                case "foo bar":
                                                    if true {
                                                        v1 = (v1 | (v1 >> ((663) & 63)))
                                                        var v63 uint16 = 65535
                                                        var v64 uint8 = 147
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                case "":
                                                    fmt.Println(v62, v2)
                                                case "\n\t":
                                                    delete(v3, ("a" + v4))
                                                    var v65 int = 2147483647
                                                    if v5 {
                                                        v66 := GenS10{F0: 0}
                                                        v30 = ((v6 / ((v6) | 1)) > v6)
                                                        _ = v66
                                                    } else {
                                                        if false {
                                                            _ = len(v3)
                                                            var v67 int16 = v2
                                                            fmt.Println(v1, v67, v30)
                                                            _ = v67
                                                        }
                                                        v5 = true
                                                    }
                                                    if false {
                                                        if v30 {
                                                            var v68 int8 = 127
                                                            v3[((v4 + v4) + v4)] = v5
                                                            _ = v68
                                                        } else {
                                                            v69 := GenS3{F0: (-(v65)), F1: 473}
                                                            var v70 uint64 = v6
                                                            var v71 int32 = 1
                                                            v72 := []int{-874, (fi41 & (^(fi40)))}
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        if (v30 || (127 < 127)) {
                                                            switch 9223372036854775807 {
                                                            case -9223372036854775808:
                                                                v73 := []int{fi41, int(v2)}
                                                                _ = v73
                                                            }
                                                        }
                                                        v74 := GenS6{F0: 127, F1: ("foo bar" >= v4)}
                                                        _ = v74
                                                    }
                                                    _ = v65
                                                }
                                                v75 := GenS11{F0: v2}
                                                _ = v62
                                                _ = v75
                                            case "\n\t":
                                                var v76 string = v4
                                                _ = v76
                                            default:
                                                v4 = v4
                                            }
                                        }
                                        v77 := GenS12{F0: 42, F1: "hello", F2: 0}
                                        v2 = (v2 & v2)
                                        v5 = true
                                        switch -2147483648 {
                                        case 32768:
                                            v78 := GenS13{F0: 42}
                                            _ = v78
                                        case -825:
                                            fmt.Println(v6)
                                        }
                                        fmt.Println(v5, v4, fi40)
                                        _ = v77
                                    }
                                    _ = v49
                                } else {
                                    if v5 {
                                        fmt.Println(v6)
                                        if true {
                                            var v79 string = "a"
                                            _ = v79
                                        } else {
                                            if (1 >= 676) {
                                                if v5 {
                                                    v3["hello"] = true
                                                }
                                            }
                                            var v80 uint8 = 236
                                            _ = v80
                                        }
                                        var v81 string = ("a" + v4)
                                        _ = v81
                                    }
                                    v82 := v3[v4]
                                    _ = v82
                                }
                                v83 := GenS6{F0: -113, F1: false}
                                _ = v83
                            }
                        default:
                            v84 := GenS8{F0: 566, F1: v1, F2: 732}
                            var v85 uint = 2
                            switch (fi41 & fi40) {
                            case -9223372036854775808:
                                fmt.Println(v30, v30)
                            case -1:
                                v86 := []int{}
                                _ = v86
                            case 32767:
                                var v87 uint64 = genFn0()
                                v87 = (v87 << (((v85 | (^(v85)))) & 63))
                                v85 += 674
                                if (93 <= 88) {
                                    v88 := GenS11{F0: v2}
                                    fmt.Println(v85, v1, v30)
                                    var v89 uint16 = 32767
                                    _ = v88
                                    _ = v89
                                }
                                _ = v87
                            }
                            _ = v84
                            _ = v85
                        }
                        delete(v3, ((v4 + v4) + ((v4 + (v4 + ("foo bar" + v4))) + ("a" + v4))))
                        switch fi41 {
                        case 1:
                            v90 := GenS12{F0: 1, F1: v4, F2: 209}
                            _ = v90
                        case 32767:
                            v91 := GenS14{F0: 4294967295, F1: (v6 >> ((255) & 63)), F2: -2147483648}
                            _ = v91
                        case 42:
                            v6 = v6
                        default:
                            v92 := GenS15{F0: 0, F1: (v4 + ((v4 + (("hello" + (("foo bar" + "a") + v4)) + "\n\t")) + "")), F2: 65535}
                            _ = v92
                        }
                        v6 -= 188
                        _ = v44
                    }
                    var v93 uint64 = v6
                    var v94 string = ("hello" + v4)
                    _ = fi41
                    _ = v93
                    _ = v94
                }
                _ = fi40
            }
        } else {
            fmt.Println(v4, v1)
        }
    } else {
        if v5 {
            v6 = (v6 | v6)
            for fi95 := 0; fi95 < 18; fi95++ {
                if v5 {
                    if v30 {
                        for fi96 := 0; fi96 < 1; fi96++ {
                            if v30 {
                                if v30 {
                                    switch (fi96 * fi96) {
                                    case 0:
                                        v97 := GenS16{F0: v6, F1: 254, F2: 65535}
                                        _ = v97
                                    case -354:
                                        v3[("\n\t" + v4)] = true
                                    default:
                                        switch (^(fi95)) {
                                        case 0:
                                            v1 += (v1 / ((110) | 1))
                                        default:
                                            v98 := GenS7{F0: (v4 + ("a" + v4))}
                                            _ = v98
                                        }
                                    }
                                }
                                var v99 uint64 = v6
                                _ = v99
                            } else {
                                var v100 string = v4
                                _ = v100
                            }
                            switch (("\n\t" + "") + v4) {
                            case "":
                                var v101 float32 = (1e10 / ((0.0 + ((0.1 / ((((3.14159 + 1e10) - 1e10) + (0.1 + (0.1 * 0.1))) + 0.1)) * 1e10)) + 0.0))
                                _ = v101
                            case "foo bar":
                                if v5 {
                                    var v102 float32 = 1.5
                                    if (1 == 1) {
                                        var v103 string = v4
                                        v3[v103] = (1 <= int64(fi95))
                                        var v104 uint32 = 42
                                        _ = len(v3)
                                        v105 := GenS17{F0: 1}
                                        _ = v103
                                        _ = v104
                                        _ = v105
                                    } else {
                                        var v106 uint8 = 255
                                        fmt.Println(v4, v6)
                                        v4 = ((("hello" + v4) + v4) + "hello")
                                        _ = v106
                                    }
                                    if false {
                                        var v107 bool = true
                                        v2 = v2
                                        _ = v107
                                    }
                                    _ = v102
                                } else {
                                    if true {
                                        v5 = (0.0 > 0.0)
                                    }
                                    var v108 uint8 = 42
                                    _ = v108
                                }
                            }
                            fmt.Println(v1)
                            v109 := GenS18{F0: 65535}
                            _ = fi96
                            _ = v109
                        }
                    }
                }
                fmt.Println(v2, v6)
                _ = fi95
            }
            v110 := v3["foo bar"]
            _ = v110
        } else {
            v2 += (^(v2))
            var v111 uint64 = uint64(v1)
            _ = v111
        }
        v1 -= 198
    }
    var v112 string = ("" + v4)
    fmt.Println(v5, v6)
    if true {
        for fi113 := 0; fi113 < 3; fi113++ {
            switch (^(fi113)) {
            case -2147483648:
                var v114 int16 = v2
                _ = v114
            case 4294967295:
                var v115 int8 = 1
                _ = v115
            case -9223372036854775808:
                var v116 int64 = -9223372036854775808
                _ = v116
            default:
                switch ("\n\t" + v4) {
                case "":
                    v117 := GenS18{F0: 65535}
                    _ = v117
                default:
                    for fi118 := 0; fi118 < 13; fi118++ {
                        v119 := GenS19{F0: (v2 - (-(v2)))}
                        _ = fi118
                        _ = v119
                    }
                }
            }
            var v120 int8 = -1
            v4 = v112
            _ = fi113
            _ = v120
        }
        fmt.Println(v30, v5, v1)
        var v121 uint8 = 128
        _ = v121
    }
    if v30 {
        if (42 < 4294967295) {
            var v122 int16 = (v2 % (((v2 & v2)) | 1))
            _ = v122
        }
    }
    var v123 uint16 = 65535
    v112 = v112
    var v124 float32 = (0.1 + 1e10)
    fmt.Println(v112, v123)
    v5 = false
    for fi125 := 0; fi125 < 19; fi125++ {
        v126 := GenS7{F0: v4}
        fmt.Println(v1, v112, v30)
        if (!v30) {
            var v127 uint32 = v1
            _ = v127
        }
        _ = fi125
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
