package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 bool; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 uint16; F1 string; F2 int }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 uint64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(fi8)
            default:
                var v9 int64 = 1
                _ = v9
            }
            var v10 uint32 = 42
            v10 = v10
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    v11 := []int{-915}
    var v12 float32 = 1.5
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int32 = -2147483648
        var v15 uint8 = 2
        _ = len(v11)
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if true {
            switch 128 {
            case 32768:
                v12 += (float32(3.14159) - v12)
            case 4294967295:
                v12 += v12
            default:
                v12 = v12
            }
            v12 += v12
        } else {
            if (127 == 127) {
                if true {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v12 = (0.1 + 3.14159)
                        var v17 int8 = -38
                        _ = fi16
                        _ = v17
                    }
                } else {
                    v18 := GenS1{F0: 1}
                    fmt.Println(v12)
                    _ = v18
                }
                switch 42 {
                case -1:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v11 = append(v11, 43)
                        _ = fi19
                    }
                }
                switch "" {
                case "a":
                    v20 := GenS1{F0: -32768}
                    _ = v20
                }
            }
            var v21 float32 = v12
            var v22 int16 = -32768
            _ = v21
            _ = v22
        }
    }
    _ = v11
    _ = v12
    return 32767
}
func genFn1(p0 uint64, p1 int64, p2 uint64) uint32 {
    switch 32768 {
    case -2147483648:
        switch "hello" {
        case "hello":
            v23 := GenS2{F0: 753, F1: 2147483647}
            _ = v23
        default:
            fmt.Println(p1)
        }
    case -9223372036854775808:
        if true {
            var v24 uint8 = 1
            var v25 uint32 = 2147483647
            _ = v24
            _ = v25
        }
    case 32767:
        var v26 uint8 = 106
        _ = v26
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        p1 += (p1 * p1)
        for fi28 := 0; fi28 < 20; fi28++ {
            if (true || true) {
                v29 := []int{2147483647}
                var v30 int64 = (p1 << ((42) & 63))
                var v31 int64 = 735
                _ = v29
                _ = v30
                _ = v31
            } else {
                fmt.Println(p0, p1, p2)
            }
            _ = fi28
        }
        var v32 int = fi27
        _ = fi27
        _ = v32
    }
    if (890 > 0) {
        var v33 bool = (65535 > 65535)
        fmt.Println(p1, p0)
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 uint16 = 65535
        p1 = 405
        _ = fi34
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    var v5 uint64 = 4294967295
    var v6 string = (v3 + (v3 + "foo bar"))
    if false {
        v36 := GenS3{F0: v5}
        _ = v36
    } else {
        delete(v1, 4294967295)
        fmt.Println(v6)
        var v37 int16 = 42
        _ = v37
    }
    var v38 string = v6
    var v39 int32 = -885
    fmt.Println(v39, v39, v6)
    var v40 float32 = 3.14159
    for fi41 := 0; fi41 < 12; fi41++ {
        if true {
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 int64 = 0
                var v44 uint = 1
                if false {
                    switch ("a" + ((v3 + (v6 + v3)) + (v6 + v6))) {
                    case "\n\t":
                        v3 = (v3 + v6)
                    case "hello":
                        v45 := []int{}
                        v44 = (^(v44))
                        _ = v45
                    }
                }
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        _ = fi41
    }
    if false {
        v40 = v40
        switch "a" {
        case "\n\t":
            fmt.Println(v39, v3, v38)
        case "a":
            delete(v4, "")
        }
        v46 := []int{1, 127, 127}
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v3 = v6
        delete(v4, v6)
        _ = fi47
    }
    var v48 int8 = 100
    if ((v6 + v6) == v3) {
        v5 += (v5 / ((0) | 1))
        if true {
            if true {
                v39 = (v39 * -639)
                for fi49 := 0; fi49 < 3; fi49++ {
                    v38 = (((v38 + (v3 + ("a" + v38))) + v6) + ((v6 + v3) + "a"))
                    v50 := GenS1{F0: 32767}
                    _ = fi49
                    _ = v50
                }
            } else {
                if (true || ((!((-2.25 / float64(1.5)) == 1.5)) || false)) {
                    v6 = "hello"
                    switch "foo bar" {
                    case "":
                        var v51 uint32 = uint32(v5)
                        _ = v51
                    case "a":
                        switch -489 {
                        case -1:
                            if false {
                                var v52 int8 = v48
                                v53 := GenS2{F0: 42, F1: 65535}
                                _ = v52
                                _ = v53
                            }
                        case 128:
                            v54 := GenS4{F0: (v5 | v5), F1: false, F2: "a"}
                            _ = v54
                        }
                    case "hello":
                        if (-1 > genFn0()) {
                            v55 := v4[v6]
                            for fi56 := 0; fi56 < 3; fi56++ {
                                if (true || true) {
                                    var v57 uint32 = (v55 - v55)
                                    v39 = v39
                                    delete(v4, v38)
                                    _ = v57
                                } else {
                                    var v58 uint8 = 127
                                    v59 := GenS1{F0: genFn0()}
                                    v60 := GenS5{F0: 32767}
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                v61 := GenS6{F0: uint16(v5)}
                                for fi62 := 0; fi62 < 10; fi62++ {
                                    v5 = v5
                                    v63 := GenS1{F0: 367}
                                    v64 := GenS6{F0: 458}
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                _ = fi56
                                _ = v61
                            }
                            _ = v55
                        } else {
                            for fi65 := 0; fi65 < 3; fi65++ {
                                if (-2.25 < float32(-2.25)) {
                                    fmt.Println(v5, v5)
                                    v4[v3] = genFn1(4294967295, 9223372036854775807, v5)
                                } else {
                                    v66 := GenS7{F0: 42, F1: 32767}
                                    switch v38 {
                                    case "foo bar":
                                        fmt.Println(v40, fi65, v39)
                                    }
                                    _ = v66
                                }
                                if (0 != 255) {
                                    var v67 bool = (0 > 32767)
                                    _ = v67
                                }
                                _ = fi65
                            }
                        }
                    }
                } else {
                    v39 = (v39 & int32(v48))
                }
                _ = len(v2)
            }
        }
        var v68 uint = 1
        _ = v68
    }
    if (!((false && (2 > 255)) && ((1 != v39) || (0 < genFn0())))) {
        v69 := GenS5{F0: 767}
        if true {
            if (!(true && (!true))) {
                for fi70 := 0; fi70 < 16; fi70++ {
                    var v71 int8 = v48
                    var v72 float64 = ((0.1 * -2.25) * float64(v40))
                    v48 = 65
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                v38 = v38
                if (!(((v40 / (v40 / v40)) / 0.1) < v40)) {
                    v3 = "\n\t"
                    if (-531 <= 1000) {
                        var v73 float32 = v40
                        v40 -= 3.14159
                        switch v3 {
                        case "\n\t":
                            var v74 int = 127
                            _ = v74
                        default:
                            v39 += (v39 % ((v39) | 1))
                        }
                        _ = v73
                    } else {
                        for fi75 := 0; fi75 < 16; fi75++ {
                            delete(v1, 4294967295)
                            var v76 uint32 = 446
                            _ = fi75
                            _ = v76
                        }
                        if ((v5 / (((v5 & 18446744073709551615)) | 1)) < (v5 * v5)) {
                            for fi77 := 0; fi77 < 12; fi77++ {
                                var v78 uint16 = 255
                                if (-294 > 42) {
                                    var v79 uint8 = 2
                                    v80 := GenS8{F0: (((v3 + "a") + ("" + (v38 + (v38 + ("\n\t" + "foo bar"))))) + (("foo bar" + (v38 + ((v38 + v6) + (((v6 + ("a" + v6)) + v3) + ("a" + (((v6 + v6) + v38) + "hello")))))) + "foo bar")), F1: (v39 / ((-1) | 1))}
                                    _ = v79
                                    _ = v80
                                } else {
                                    for fi81 := 0; fi81 < 1; fi81++ {
                                        if true {
                                            switch -2147483648 {
                                            case -1:
                                                v82 := GenS9{F0: v6, F1: 65535}
                                                _ = v82
                                            case -2147483648:
                                                var v83 uint64 = (v5 / (((v5 * (v5 % ((v5) | 1)))) | 1))
                                                _ = v83
                                            case 128:
                                                v84 := GenS1{F0: 1000}
                                                _ = v84
                                            default:
                                                fmt.Println(v6, v78, v5)
                                            }
                                            var v85 int16 = -305
                                            _ = v85
                                        }
                                        v86 := v2[(fi81 >> ((1) & 63))]
                                        _ = fi81
                                        _ = v86
                                    }
                                    v87 := []int{}
                                    _ = v87
                                }
                                fmt.Println(v78, v6, v3)
                                _ = fi77
                                _ = v78
                            }
                        } else {
                            v48 = -100
                        }
                    }
                } else {
                    if ((!true) || true) {
                        if (42 == 2147483647) {
                            switch 128 {
                            case -2147483648:
                                v88 := []int{-611, int(v48), 127, 128}
                                _ = v88
                            }
                            v89 := []int{429, -9223372036854775808, 32767, 1}
                            if len(v89) > 0 {
                                v89[0] = 0
                            }
                            _ = v89
                        } else {
                            switch -51 {
                            case -9223372036854775808:
                                v90 := GenS10{F0: (v48 * (^(v48)))}
                                _ = v90
                            case 32768:
                                var v91 uint16 = 2
                                _ = v91
                            case 42:
                                v92 := GenS2{F0: 2147483648, F1: 19}
                                _ = v92
                            }
                        }
                        fmt.Println(v5)
                    } else {
                        if ((!(true && true)) && true) {
                            v93 := v1[885]
                            _ = v93
                        } else {
                            if (v40 == 3.14159) {
                                var v94 int8 = (v48 / (((v48 + (v48 + v48))) | 1))
                                fmt.Println(v40, v48, v6)
                                _ = v94
                            }
                        }
                    }
                    var v95 int32 = (v39 - v39)
                    _ = v95
                }
            } else {
                fmt.Println(v48, v48)
                for fi96 := 0; fi96 < 15; fi96++ {
                    v39 += (v39 | 0)
                    v3 = ""
                    v1[1] = false
                    _ = fi96
                }
                for fi97 := 0; fi97 < 2; fi97++ {
                    switch v3 {
                    case "":
                        for fi98 := 0; fi98 < 13; fi98++ {
                            switch fi98 {
                            case 9223372036854775807:
                                v99 := GenS10{F0: 0}
                                _ = v99
                            default:
                                v100 := GenS11{F0: -1, F1: (-(v48))}
                                v6 = (("\n\t" + (((v6 + v38) + "a") + v3)) + ("" + "hello"))
                                _ = v100
                            }
                            v101 := []int{(fi98 >> ((18446744073709551615) & 63))}
                            if ((((1e10 - 3.14159) + 1.5) * v40) > v40) {
                                switch (v6 + v38) {
                                case "hello":
                                    if true {
                                        var v102 int64 = 57
                                        v3 = ""
                                        _ = v102
                                    }
                                    v103 := GenS12{F0: 2, F1: v38, F2: fi97}
                                    _ = len(v101)
                                    _ = v103
                                case "\n\t":
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        _ = len(v4)
                                    default:
                                        var v104 int64 = -506
                                        if (1 <= 255) {
                                            var v105 int32 = v39
                                            v106 := GenS13{F0: 582, F1: (^(v104))}
                                            v107 := GenS14{F0: (v5 * 384)}
                                            _ = v105
                                            _ = v106
                                            _ = v107
                                        } else {
                                            v108 := GenS14{F0: 42}
                                            fmt.Println(fi97)
                                            _ = v108
                                        }
                                        _ = v104
                                    }
                                }
                                switch (v38 + (((v6 + v38) + v6) + v38)) {
                                case "a":
                                    var v109 uint8 = 80
                                    _ = v109
                                default:
                                    if (((v38 + ("" + v6)) + v6) <= (("hello" + ((v3 + "foo bar") + ((v3 + (((v6 + "a") + ("\n\t" + ("\n\t" + "\n\t"))) + v38)) + v6))) + v38)) {
                                        v38 = v6
                                        v110 := GenS7{F0: 42, F1: 519}
                                        _ = v110
                                    }
                                }
                            } else {
                                v111 := GenS7{F0: uint(v5), F1: genFn0()}
                                v112 := v4["hello"]
                                var v113 int8 = v48
                                _ = v111
                                _ = v112
                                _ = v113
                            }
                            v1[463] = false
                            _ = fi98
                            _ = v101
                        }
                    case "a":
                        for fi114 := 0; fi114 < 2; fi114++ {
                            v3 = v6
                            if (32767 <= 42) {
                                if (v39 > (-(v39))) {
                                    v115 := []int{fi97, -787}
                                    if (2147483647 <= (v39 + (v39 + (v39 - v39)))) {
                                        var v116 string = v6
                                        _ = v116
                                    }
                                    var v117 float64 = 1e10
                                    var v118 uint8 = 127
                                    if ((!(0 >= uint64(v118))) && true) {
                                        v6 = (v3 + "")
                                        v38 = v6
                                        switch -1 {
                                        case 4294967295:
                                            v119 := GenS4{F0: (v5 & (^(v5))), F1: false, F2: (v3 + ((("\n\t" + "") + "") + (v3 + (v3 + "hello"))))}
                                            _ = v119
                                        default:
                                            v5 -= v5
                                        }
                                    } else {
                                        switch v38 {
                                        case "hello":
                                            if len(v115) >= 2 {
                                                v115 = v115[1:len(v115)]
                                            }
                                        default:
                                            if (true && false) {
                                                v120 := v4[v3]
                                                _ = v120
                                            }
                                        }
                                        var v121 uint8 = (v118 & v118)
                                        v3 = v3
                                        _ = v121
                                    }
                                    _ = v115
                                    _ = v117
                                    _ = v118
                                }
                                delete(v2, 2147483647)
                                var v122 int8 = (^(v48))
                                v123 := GenS4{F0: v5, F1: false, F2: (v6 + v6)}
                                switch (fi114 | fi97) {
                                case 128:
                                    if true {
                                        delete(v1, uint32(v5))
                                        var v124 bool = (248 > 1)
                                        _ = v124
                                    }
                                default:
                                    if (!(!((fi97 % (((fi114 * fi114)) | 1)) > (fi114 << ((4294967295) & 63))))) {
                                        if true {
                                            var v125 int64 = 42
                                            _ = v125
                                        } else {
                                            delete(v2, 0)
                                            v6 = (v6 + v38)
                                        }
                                    } else {
                                        var v126 string = v3
                                        var v127 uint16 = 151
                                        _ = v126
                                        _ = v127
                                    }
                                }
                                _ = v122
                                _ = v123
                            }
                            _ = fi114
                        }
                    case "foo bar":
                        v5 -= 4294967295
                    }
                    _ = fi97
                }
            }
        }
        _ = v69
    } else {
        v128 := GenS12{F0: 255, F1: v38, F2: 9223372036854775807}
        if (442 > 32767) {
            v129 := []int{-727}
            _ = v129
        }
        _ = v128
    }
    fmt.Println(v5, v38, v39)
    var v130 uint = 2
    v38 = (("foo bar" + ("\n\t" + (v6 + (v6 + (v3 + ("hello" + v6)))))) + (v38 + (v38 + ((("a" + ((v3 + "foo bar") + (("hello" + "a") + v38))) + v3) + "a"))))
    v2[127] = 650
    delete(v4, "a")
    v48 -= (v48 & 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v130
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v48
    _ = v5
    _ = v6
}
