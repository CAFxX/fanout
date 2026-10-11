package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int32; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint; F2 int16 }
type GenS8 struct { F0 uint16; F1 int16; F2 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 uint8; F2 int }
type GenS12 struct { F0 string; F1 uint64; F2 int8 }
type GenS13 struct { F0 int16; F1 int16 }
type GenS14 struct { F0 int8; F1 string }
type GenS15 struct { F0 uint64; F1 int16 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint32; F1 uint16 }
type GenS18 struct { F0 int; F1 int32 }
type GenS19 struct { F0 int32; F1 int; F2 bool }
type GenS20 struct { F0 int; F1 string; F2 bool }
type GenS21 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int8) uint64 {
    v7 := []int{}
    p0 = p0
    switch (("a" + "a") + "a") {
    case "a":
        var v8 uint8 = 196
        _ = v8
    case "":
        var v9 uint8 = 143
        _ = v9
    default:
        if true {
            var v10 int8 = p0
            v11 := GenS1{F0: 255, F1: 4294967295, F2: 32768}
            for fi12 := 0; fi12 < 1; fi12++ {
                p0 = v10
                _ = fi12
            }
            _ = v10
            _ = v11
        } else {
            p0 += (p0 ^ -128)
        }
    }
    if (float64(-2.25) != -2.25) {
        fmt.Println(p0)
        _ = len(v7)
        var v13 int8 = p0
        _ = v13
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float32 = 1.5
            v16 := GenS2{F0: 2147483648}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint32 = 491
        p0 = p0
        _ = v17
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    var v3 float32 = 1e10
    var v4 uint32 = 65535
    var v5 int64 = 9223372036854775807
    v6 := []string{}
    var v18 int16 = 0
    if (!(false || (!false))) {
        if true {
            v19 := GenS2{F0: v4}
            switch 127 {
            case 42:
                v5 = v5
            }
            var v20 int = -9223372036854775808
            _ = v19
            _ = v20
        } else {
            if false {
                if false {
                    var v21 int64 = (v5 ^ (v5 % (((-(v5))) | 1)))
                    v1["a"] = (65 == 127)
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v4 -= v4
                        fmt.Println(v18)
                        _ = fi22
                    }
                    _ = v21
                } else {
                    v23 := GenS2{F0: (v4 / ((v4) | 1))}
                    var v24 uint64 = 42
                    if true {
                        switch (("" + "hello") + "\n\t") {
                        case "hello":
                            var v25 float32 = ((((v3 * v3) * v3) - float32(0.0)) - v3)
                            _ = v25
                        case "\n\t":
                            if (v2 >= (v2 << ((72) & 63))) {
                                v24 = (v24 / ((v24) | 1))
                            } else {
                                v18 = v18
                                var v26 uint8 = 144
                                _ = v26
                            }
                        }
                        v27 := GenS3{F0: 32767}
                        _ = v27
                    } else {
                        fmt.Println(v5, v2, v4)
                        for fi28 := 0; fi28 < 9; fi28++ {
                            for fi29 := 0; fi29 < 15; fi29++ {
                                if (v5 != v5) {
                                    var v30 float32 = (v3 + v3)
                                    var v31 int8 = 127
                                    var v32 uint = 1
                                    var v33 float64 = (3.14159 - (0.1 + -2.25))
                                    if true {
                                        if ((^(v5)) == v5) {
                                            v2 = int32(v32)
                                            v24 -= genFn0((v31 << ((uint(v4)) & 63)))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            if (111 == v4) {
                                                if false {
                                                    if (false || false) {
                                                        v34 := GenS3{F0: (fi29 & fi29)}
                                                        v33 -= -2.25
                                                        switch fi28 {
                                                        case -502:
                                                            v3 += v30
                                                        case 43:
                                                            if true {
                                                                fmt.Println(v3, v24)
                                                                var v35 float64 = 1e10
                                                                _ = v35
                                                            }
                                                        }
                                                        if (true && (-13 < v18)) {
                                                            _ = len(v1)
                                                            var v36 uint8 = 1
                                                            _ = v36
                                                        }
                                                        _ = v34
                                                    }
                                                    if ((v32 & 0) <= (v32 % ((v32) | 1))) {
                                                        var v37 int16 = (v18 ^ v18)
                                                        v4 = (v4 & (^(v4)))
                                                        _ = v37
                                                    }
                                                    var v38 float64 = v33
                                                    v2 = (v2 << ((v32) & 63))
                                                    fmt.Println(v30, v32)
                                                    _ = v38
                                                }
                                                var v39 uint64 = v24
                                                var v40 int64 = v5
                                                v32 += (^(v32))
                                                v3 = (v3 / float32(((v33 * v33) + (v33 - ((v33 + -2.25) / (0.0 * float64(float32(v33))))))))
                                                _ = v39
                                                _ = v40
                                            }
                                        } else {
                                            _ = len(v6)
                                            var v41 uint = v32
                                            if ("a" != "a") {
                                                v33 -= v33
                                                var v42 float32 = v30
                                                _ = v42
                                            }
                                            _ = v41
                                        }
                                        var v43 uint32 = v4
                                        if (fi29 > (fi28 - int(v32))) {
                                            var v44 int = (fi28 % ((fi29) | 1))
                                            _ = v44
                                        } else {
                                            v4 = v4
                                            var v45 int8 = v31
                                            switch int(v45) {
                                            case 42:
                                                v5 += -980
                                            case 965:
                                                if false {
                                                    v46 := GenS4{F0: 1, F1: 65535, F2: v24}
                                                    _ = v46
                                                } else {
                                                    switch fi29 {
                                                    case 4294967295:
                                                        _ = len(v6)
                                                    default:
                                                        switch (("hello" + "foo bar") + "foo bar") {
                                                        case "a":
                                                            var v47 float32 = ((float32((v33 - (float64((float32((v33 + (v33 * v33))) / ((0.0 + (1.5 / (v30 - v3))) * float32(v33)))) / (((v33 - ((v33 / 1e10) + v33)) + (v33 - (v33 * 0.1))) - ((0.0 * (3.14159 / v33)) + ((v33 * (v33 - v33)) / v33)))))) / 3.14159) * float32(v33))
                                                            _ = v47
                                                        case "foo bar":
                                                            fmt.Println(v31, v31)
                                                        }
                                                        var v48 uint16 = 57
                                                        switch "" {
                                                        case "a":
                                                            v5 += -578
                                                        default:
                                                            v6 = append(v6, "a")
                                                        }
                                                        switch 2147483647 {
                                                        case -877:
                                                            if (128 > (fi28 + int(v24))) {
                                                                _ = len(v1)
                                                                v18 += v18
                                                            } else {
                                                                v18 -= v18
                                                                if true {
                                                                    var v49 uint = (v32 / ((v32) | 1))
                                                                    v50 := GenS5{F0: (v45 / (((v31 ^ (^(v31)))) | 1)), F1: (^(v5))}
                                                                    if (false && (!(v48 < (v48 + 65535)))) {
                                                                        v33 -= float64((v30 + (v30 - (v30 - ((v30 / ((v3 + v30) - float32(3.14159))) - v30)))))
                                                                        v51 := GenS6{F0: 520, F1: v2, F2: (v24 / (((v24 ^ v24)) | 1))}
                                                                        var v52 int = (fi29 >> (((v49 - v32)) & 63))
                                                                        _ = v51
                                                                        _ = v52
                                                                    } else {
                                                                        if (v49 != v49) {
                                                                            var v53 uint32 = v4
                                                                            switch fi29 {
                                                                            case 32768:
                                                                                if false {
                                                                                    var v54 uint = (v32 * (^(v32)))
                                                                                    var v55 int16 = 0
                                                                                    v24 = v24
                                                                                    if (-273 > (v5 << (((v49 - v32)) & 63))) {
                                                                                        delete(v1, "a")
                                                                                        if (!(v3 != -2.25)) {
                                                                                            _ = len(v6)
                                                                                            v33 += 0.1
                                                                                        }
                                                                                        v18 += v18
                                                                                        var v56 uint16 = (v48 ^ v48)
                                                                                        _ = v56
                                                                                    } else {
                                                                                        var v57 string = (("\n\t" + "\n\t") + ("" + ("foo bar" + "\n\t")))
                                                                                        var v58 uint16 = (v48 / (((^(v48))) | 1))
                                                                                        var v59 int16 = (v18 ^ v55)
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                    switch ("foo bar" + (("a" + ("\n\t" + "a")) + "a")) {
                                                                                    case "":
                                                                                        var v60 float64 = (((((v33 / 0.1) / v33) + 1e10) * v33) + -2.25)
                                                                                        _ = v60
                                                                                    case "a":
                                                                                        if true {
                                                                                            v61 := GenS5{F0: int8(v48), F1: (v5 + (-(v5)))}
                                                                                            v45 -= (-(v45))
                                                                                            _ = v61
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        var v62 uint32 = 386
                                                                                        if (127 >= fi29) {
                                                                                            v6 = append(v6, "\n\t")
                                                                                            v4 -= (v4 / ((v4) | 1))
                                                                                        }
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = v62
                                                                                    } else {
                                                                                        var v63 string = ""
                                                                                        _ = v63
                                                                                    }
                                                                                    v18 -= int16(v4)
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                } else {
                                                                                    var v64 int16 = 0
                                                                                    v33 += 1e10
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                            _ = v53
                                                                        } else {
                                                                            var v65 int8 = 77
                                                                            var v66 string = "a"
                                                                            v67 := GenS3{F0: 9223372036854775807}
                                                                            _ = v65
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        v3 -= v3
                                                                        var v68 uint16 = (v48 >> ((v49) & 63))
                                                                        _ = v68
                                                                    }
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            }
                                                        case -153:
                                                            v18 -= 0
                                                        case 32767:
                                                            var v69 int = fi29
                                                            if (479 == (v2 + v2)) {
                                                                if (!((0.1 * v3) < v30)) {
                                                                    v70 := GenS7{F0: (v32 << ((v32) & 63)), F1: (v32 / (((v32 * v32)) | 1)), F2: 0}
                                                                    var v71 int = (fi29 + -9223372036854775808)
                                                                    v72 := GenS8{F0: (v48 * v48), F1: (-(v18)), F2: v69}
                                                                    _ = v70
                                                                    _ = v71
                                                                    _ = v72
                                                                } else {
                                                                    v73 := GenS9{F0: true}
                                                                    _ = v73
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    v33 -= -2.25
                                                                case "hello":
                                                                    v30 = v30
                                                                case "foo bar":
                                                                    if (-2.25 == v3) {
                                                                        var v74 uint16 = (v48 / ((uint16(v4)) | 1))
                                                                        fmt.Println(v48, fi28)
                                                                        _ = v74
                                                                    } else {
                                                                        var v75 int64 = v5
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            if ((v43 % (((^(v43))) | 1)) >= (v4 + v43)) {
                                                                                var v76 float64 = v33
                                                                                v33 += v76
                                                                                _ = v76
                                                                            } else {
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = ""
                                                                                }
                                                                                v31 += (v45 % (((-(v31))) | 1))
                                                                                switch "a" {
                                                                                case "":
                                                                                    v48 += v48
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v75
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch fi29 {
                                                                        case 0:
                                                                            _ = len(v1)
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v4, v33, v69)
                                                                        } else {
                                                                            v3 += v3
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v77 := GenS8{F0: v48, F1: (v18 & -1), F2: (v69 << (((v32 >> ((v32) & 63))) & 63))}
                                                                                _ = v77
                                                                            case "hello":
                                                                                var v78 uint64 = 4294967295
                                                                                _ = v78
                                                                            case "":
                                                                                var v79 uint = v32
                                                                                _ = v79
                                                                            default:
                                                                                var v80 uint32 = v43
                                                                                _ = v80
                                                                            }
                                                                            v81 := GenS8{F0: 2, F1: v18, F2: fi29}
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if ((float32((1.5 + v33)) / v3) < v30) {
                                                                v82 := GenS7{F0: uint(v48), F1: v32, F2: (v18 + 566)}
                                                                if (false || true) {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = "foo bar"
                                                                    }
                                                                    v83 := GenS7{F0: uint(v24), F1: (v32 * 362), F2: v18}
                                                                    var v84 float32 = float32(v33)
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                _ = v82
                                                            } else {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v85 uint8 = 2
                                                                    _ = v85
                                                                case "\n\t":
                                                                    v5 += v5
                                                                }
                                                                v45 -= (v45 * v45)
                                                            }
                                                            v86 := GenS10{F0: (v69 + (fi28 & int(v31)))}
                                                            _ = v69
                                                            _ = v86
                                                        }
                                                        _ = v48
                                                    }
                                                }
                                            case 942:
                                                v87 := GenS11{F0: "", F1: 255, F2: int(v18)}
                                                _ = v87
                                            default:
                                                v43 = (v43 * v4)
                                            }
                                            _ = v45
                                        }
                                        if false {
                                            v88 := GenS1{F0: (^(v32)), F1: (v4 * uint32(v32)), F2: (fi28 & (fi28 + fi28))}
                                            _ = v88
                                        }
                                        _ = v43
                                    }
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                _ = len(v6)
                                v89 := GenS12{F0: "foo bar", F1: v24, F2: 71}
                                var v90 string = "\n\t"
                                if false {
                                    v1[""] = false
                                    switch v90 {
                                    case "\n\t":
                                        v91 := GenS13{F0: (-(v18)), F1: (^(v18))}
                                        _ = v91
                                    case "foo bar":
                                        var v92 float64 = float64(0.1)
                                        _ = v92
                                    default:
                                        if (v4 < uint32(v24)) {
                                            v93 := GenS14{F0: int8(v18), F1: (((("\n\t" + ("a" + ("hello" + "a"))) + v90) + "") + v90)}
                                            _ = v93
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        if false {
                                            v94 := GenS15{F0: 405, F1: (v18 + -32768)}
                                            _ = v94
                                        } else {
                                            switch 128 {
                                            case 2147483647:
                                                if false {
                                                    if ((((0.1 + (1e10 + (1.5 * 3.14159))) * -2.25) * 0.1) >= (0.1 + 1.5)) {
                                                        v90 = v90
                                                    }
                                                } else {
                                                    v2 = (v2 + (v2 ^ (^(v2))))
                                                    v95 := GenS3{F0: (fi28 - (fi29 >> ((583) & 63)))}
                                                    var v96 uint8 = 2
                                                    var v97 int32 = v2
                                                    _ = v95
                                                    _ = v96
                                                    _ = v97
                                                }
                                            default:
                                                var v98 bool = false
                                                _ = v98
                                            }
                                            var v99 uint16 = 0
                                            _ = v99
                                        }
                                        v3 = 3.14159
                                    }
                                }
                                _ = fi29
                                _ = v89
                                _ = v90
                            }
                            var v100 int64 = (v5 / ((v5) | 1))
                            _ = fi28
                            _ = v100
                        }
                        v101 := GenS14{F0: 56, F1: "\n\t"}
                        _ = v101
                    }
                    _ = v23
                    _ = v24
                }
                v102 := v1[("hello" + "\n\t")]
                var v103 int8 = -36
                _ = v102
                _ = v103
            }
        }
        for fi104 := 0; fi104 < 17; fi104++ {
            if (("foo bar" <= "foo bar") || (!((v18 + v18) >= v18))) {
                var v105 float32 = v3
                fmt.Println(v2, v18)
                switch fi104 {
                case -235:
                    for fi106 := 0; fi106 < 19; fi106++ {
                        var v107 float32 = -2.25
                        if (!true) {
                            if true {
                                v6 = append(v6, ("a" + ""))
                            }
                        } else {
                            var v108 uint32 = v4
                            _ = v108
                        }
                        v109 := GenS1{F0: 31, F1: (v4 >> ((65535) & 63)), F2: -1}
                        fmt.Println(fi104, v5)
                        _ = fi106
                        _ = v107
                        _ = v109
                    }
                case -164:
                    delete(v1, "")
                default:
                    fmt.Println(v3, v3)
                }
                _ = v105
            } else {
                switch fi104 {
                case 9223372036854775807:
                    var v110 float32 = 0.0
                    _ = v110
                case 1:
                    var v111 bool = false
                    _ = v111
                case 109:
                    var v112 bool = (!(v4 != v4))
                    _ = v112
                }
                fmt.Println(v4, v2)
                v113 := GenS16{F0: "hello"}
                _ = v113
            }
            if (623 >= 0) {
                v6 = append(v6, "hello")
                if true {
                    v114 := GenS11{F0: ("hello" + (("" + "\n\t") + "")), F1: 0, F2: int(v2)}
                    _ = v114
                } else {
                    switch (fi104 ^ (^(fi104))) {
                    case 128:
                        for fi115 := 0; fi115 < 20; fi115++ {
                            v2 -= v2
                            _ = fi115
                        }
                    }
                    delete(v1, "a")
                    _ = len(v6)
                }
                if ("" < "") {
                    for fi116 := 0; fi116 < 5; fi116++ {
                        var v117 int64 = (v5 / ((v5) | 1))
                        _ = fi116
                        _ = v117
                    }
                    v18 = v18
                } else {
                    if false {
                        fmt.Println(v2, v2)
                        for fi118 := 0; fi118 < 19; fi118++ {
                            if (42 != 242) {
                                if (!true) {
                                    v119 := GenS13{F0: -254, F1: (v18 + int16(v2))}
                                    var v120 string = ""
                                    var v121 float64 = float64((v3 * float32(1e10)))
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                } else {
                                    fmt.Println(v5, v5, fi104)
                                    if (18446744073709551615 != 147) {
                                        v122 := GenS12{F0: ("hello" + "a"), F1: 42, F2: 127}
                                        if len(v6) > 0 {
                                            v6[0] = "\n\t"
                                        }
                                        switch -596 {
                                        case 32768:
                                            fmt.Println(fi104, v18)
                                        }
                                        _ = v122
                                    } else {
                                        v4 -= (v4 | v4)
                                    }
                                }
                                var v123 string = ""
                                var v124 string = v123
                                var v125 int32 = v2
                                v18 -= 1
                                _ = v123
                                _ = v124
                                _ = v125
                            }
                            _ = fi118
                        }
                    }
                    if ((-2.25 / 1e10) <= 1.5) {
                        _ = len(v6)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v126 float32 = float32((0.0 * (0.1 * -2.25)))
                        _ = v126
                    }
                    switch 32767 {
                    case 0:
                        var v127 uint64 = 1
                        _ = v127
                    case 32768:
                        v128 := GenS17{F0: 42, F1: 682}
                        _ = v128
                    case 32767:
                        v18 -= v18
                    default:
                        v2 = (-(v2))
                    }
                }
            }
            _ = fi104
        }
        v3 -= (v3 * v3)
    } else {
        v4 += 4294967295
    }
    fmt.Println(v3)
    _ = len(v1)
    if false {
        switch "hello" {
        case "hello":
            v18 = (v18 << ((0) & 63))
        case "foo bar":
            var v129 string = ""
            _ = v129
        default:
            v130 := GenS8{F0: 632, F1: 1, F2: 127}
            _ = v130
        }
        for fi131 := 0; fi131 < 15; fi131++ {
            var v132 uint8 = 42
            _ = fi131
            _ = v132
        }
    } else {
        if ((v2 / ((42) | 1)) < 42) {
            v2 = -614
            var v133 int = 2147483647
            _ = v133
        }
        v18 -= (^(v18))
    }
    fmt.Println(v4, v3)
    v134 := v1[""]
    v135 := GenS18{F0: 32768, F1: v2}
    for fi136 := 0; fi136 < 8; fi136++ {
        for fi137 := 0; fi137 < 14; fi137++ {
            delete(v1, ("foo bar" + "a"))
            switch "\n\t" {
            case "foo bar":
                v138 := v1[("hello" + (("a" + ("\n\t" + "hello")) + ((("" + "a") + ("a" + ("foo bar" + "a"))) + (("foo bar" + ("hello" + (("a" + "\n\t") + "hello"))) + ""))))]
                _ = v138
            default:
                var v139 uint64 = 952
                _ = v139
            }
            switch ("\n\t" + "hello") {
            case "hello":
                switch "hello" {
                case "\n\t":
                    fmt.Println(v18, v4, fi137)
                case "":
                    v140 := GenS4{F0: (^(v5)), F1: 1, F2: 1}
                    _ = v140
                }
            case "":
                if (v134 && false) {
                    var v141 bool = v134
                    _ = v141
                } else {
                    v18 = v18
                    if v134 {
                        v142 := GenS19{F0: (v2 ^ (^(v2))), F1: (fi136 / ((-1) | 1)), F2: v134}
                        _ = v142
                    }
                }
            case "foo bar":
                var v143 uint16 = 65535
                var v144 float64 = 0.0
                v143 = v143
                _ = v143
                _ = v144
            default:
                v145 := GenS11{F0: ("" + ""), F1: 13, F2: (fi137 >> ((1) & 63))}
                _ = v145
            }
            _ = fi137
        }
        v134 = (v2 >= -2147483648)
        _ = fi136
    }
    v134 = false
    v1["a"] = v134
    switch 1 {
    case -83:
        fmt.Println(v134)
    case 127:
        v146 := GenS9{F0: (255 != 32)}
        _ = v146
    case -729:
        for fi147 := 0; fi147 < 7; fi147++ {
            if v134 {
                for fi148 := 0; fi148 < 17; fi148++ {
                    v149 := GenS10{F0: (fi147 | (fi148 & 9223372036854775807))}
                    v150 := GenS4{F0: -323, F1: 32767, F2: 39}
                    v134 = v134
                    _ = fi148
                    _ = v149
                    _ = v150
                }
                var v151 int16 = v18
                _ = v151
            } else {
                if (230 != 649) {
                    fmt.Println(v18)
                    v152 := GenS5{F0: 1, F1: 232}
                    _ = v152
                } else {
                    v6 = append(v6, "hello")
                    var v153 int = fi147
                    v154 := GenS20{F0: fi147, F1: ("hello" + "hello"), F2: (!(-265 > (v5 ^ v5)))}
                    _ = v153
                    _ = v154
                }
            }
            switch "\n\t" {
            case "a":
                v155 := GenS17{F0: (v4 / ((911) | 1)), F1: 2}
                _ = v155
            case "":
                v5 += (v5 << ((255) & 63))
            default:
                if v134 {
                    fmt.Println(v2, v3, v134)
                    v156 := GenS11{F0: "", F1: 1, F2: fi147}
                    _ = v156
                } else {
                    var v157 uint32 = v4
                    for fi158 := 0; fi158 < 7; fi158++ {
                        if true {
                            var v159 uint16 = 65535
                            v157 -= 216
                            if v134 {
                                v134 = (210 < 222)
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v2, v5)
                                default:
                                    v160 := GenS21{F0: 2, F1: (v157 << ((119) & 63))}
                                    _ = v160
                                }
                            }
                            _ = v159
                        }
                        _ = fi158
                    }
                    _ = v157
                }
            }
            _ = fi147
        }
    default:
        if v134 {
            for fi161 := 0; fi161 < 11; fi161++ {
                v134 = false
                _ = fi161
            }
            fmt.Println(v4, v18, v3)
            v18 = v18
        }
    }
    fmt.Println(v134, v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v134
    _ = v135
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
