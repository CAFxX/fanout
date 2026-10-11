package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
type GenS3 struct { F0 int8; F1 string; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint8; F1 int64; F2 bool }
type GenS9 struct { F0 int32; F1 int64; F2 bool }
type GenS10 struct { F0 uint64; F1 int64; F2 int32 }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 uint8; F1 uint16 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 uint64; F1 bool }
type GenS18 struct { F0 int }
func genFn0(p0 string, p1 uint, p2 uint16) string {
    v4 := GenS1{F0: p1, F1: -126, F2: false}
    v5 := []int{516, 4294967295, -1, int(p2)}
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return "a"
}
func genFn1() bool {
    v7 := GenS1{F0: 2, F1: 100, F2: (true && true)}
    v8 := GenS1{F0: 888, F1: 0, F2: true}
    var v9 int32 = -2147483648
    _ = v7
    _ = v8
    _ = v9
    return (3.14159 != 0.1)
}

func main() {
    var v1 float64 = ((-2.25 * -2.25) / -2.25)
    var v2 string = ""
    v3 := make(map[uint32]int)
    v10 := GenS1{F0: 2, F1: 127, F2: false}
    var v11 bool = (true || (!false))
    var v12 float64 = v1
    for fi13 := 0; fi13 < 5; fi13++ {
        v2 = "foo bar"
        var v14 int64 = -9223372036854775808
        _ = fi13
        _ = v14
    }
    if false {
        v12 += float64(-2.25)
        var v15 bool = (-128 != 0)
        _ = v15
    }
    var v16 uint16 = 255
    var v17 uint64 = 2
    v1 += (1e10 * (1.5 + 0.1))
    v11 = false
    _ = len(v3)
    if (false && (v11 || (("foo bar" <= "\n\t") || v11))) {
        fmt.Println(v17, v17)
        for fi18 := 0; fi18 < 5; fi18++ {
            if ((!v11) || (-965 == fi18)) {
                if true {
                    if v11 {
                        if ("a" <= v2) {
                            v19 := []int{(fi18 % ((9223372036854775807) | 1)), fi18, 9223372036854775807}
                            v16 = v16
                            v11 = ((!(615 != 2147483648)) || true)
                            _ = v19
                        } else {
                            v1 = v1
                            delete(v3, 441)
                            var v20 int16 = 762
                            _ = v20
                        }
                        if v11 {
                            v2 = v2
                            fmt.Println(v12, v16)
                        } else {
                            switch "" {
                            case "":
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    var v22 uint32 = 573
                                    if v11 {
                                        v16 = v16
                                    }
                                    _ = fi21
                                    _ = v22
                                }
                            default:
                                v23 := GenS1{F0: 42, F1: int8(v17), F2: (-701 > 781)}
                                _ = v23
                            }
                            switch fi18 {
                            case 0:
                                switch fi18 {
                                case 2147483647:
                                    for fi24 := 0; fi24 < 4; fi24++ {
                                        if v11 {
                                            fmt.Println(v11, v17)
                                        } else {
                                            if true {
                                                v25 := GenS1{F0: 274, F1: -1, F2: v11}
                                                var v26 int16 = 152
                                                if v11 {
                                                    var v27 int32 = -2147483648
                                                    v17 = (v17 & v17)
                                                    delete(v3, 2147483648)
                                                    _ = v27
                                                }
                                                _ = v25
                                                _ = v26
                                            }
                                            switch (v2 + v2) {
                                            case "foo bar":
                                                fmt.Println(v1)
                                            }
                                            if v11 {
                                                switch fi24 {
                                                case -249:
                                                    if true {
                                                        fmt.Println(fi18, v16)
                                                        v28 := GenS1{F0: 0, F1: 100, F2: (("foo bar" < ((v2 + (v2 + v2)) + v2)) || v11)}
                                                        fmt.Println(v2, fi24)
                                                        _ = v28
                                                    } else {
                                                        fmt.Println(v17, fi18, fi24)
                                                        var v29 int = (fi24 << ((135) & 63))
                                                        v12 = 1e10
                                                        if v11 {
                                                            v1 = v12
                                                            var v30 string = v2
                                                            v31 := GenS1{F0: 0, F1: 66, F2: ((1 > fi18) && ((fi18 >= (fi18 * (^(v29)))) || true))}
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        fmt.Println(v11)
                                                        var v32 float64 = (3.14159 / 0.1)
                                                        _ = v29
                                                        _ = v32
                                                    }
                                                    _ = len(v3)
                                                    if v11 {
                                                        var v33 bool = v11
                                                        _ = v33
                                                    } else {
                                                        var v34 int8 = int8(fi24)
                                                        _ = v34
                                                    }
                                                default:
                                                    var v35 int = fi18
                                                    switch 42 {
                                                    case 2147483647:
                                                        if true {
                                                            if (0 <= 42) {
                                                                switch (v35 << ((42) & 63)) {
                                                                case 32768:
                                                                    var v36 uint8 = 42
                                                                    _ = v36
                                                                case -294:
                                                                    v37 := GenS2{F0: 18446744073709551615, F1: 42, F2: int16(v16)}
                                                                    _ = v37
                                                                default:
                                                                    var v38 float32 = 0.1
                                                                    _ = v38
                                                                }
                                                                var v39 uint16 = v16
                                                                _ = len(v3)
                                                                _ = v39
                                                            } else {
                                                                var v40 int = 32768
                                                                _ = len(v3)
                                                                v41 := GenS3{F0: 42, F1: (v2 + ((v2 + v2) + v2)), F2: v2}
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        } else {
                                                            if true {
                                                                var v42 float64 = v1
                                                                v3[2] = fi18
                                                                _ = v42
                                                            }
                                                            v17 = v17
                                                            var v43 bool = (!(1 <= 41))
                                                            _ = len(v3)
                                                            v3[2] = 414
                                                            _ = v43
                                                        }
                                                    case 9223372036854775807:
                                                        var v44 bool = (!v11)
                                                        v45 := GenS1{F0: 65535, F1: 69, F2: ((true || (((-(fi24)) != fi18) || (v12 <= v12))) || v44)}
                                                        var v46 float32 = 0.0
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    default:
                                                        var v47 uint32 = 65535
                                                        _ = v47
                                                    }
                                                    var v48 float64 = 3.14159
                                                    _ = v35
                                                    _ = v48
                                                }
                                                switch -1 {
                                                case 128:
                                                    var v49 uint = 65535
                                                    _ = v49
                                                case 32768:
                                                    var v50 int64 = 42
                                                    _ = v50
                                                case 9223372036854775807:
                                                    _ = len(v3)
                                                }
                                                if ((v16 * v16) <= v16) {
                                                    if (0 <= 2147483647) {
                                                        if (!(108 > 428)) {
                                                            var v51 string = v2
                                                            v2 = v51
                                                            var v52 bool = v11
                                                            _ = v51
                                                            _ = v52
                                                        } else {
                                                            v53 := GenS4{F0: 0}
                                                            _ = v53
                                                        }
                                                    } else {
                                                        v2 = "a"
                                                    }
                                                    var v54 uint32 = 65535
                                                    _ = v54
                                                }
                                                if false {
                                                    var v55 uint32 = 2
                                                    v17 = v17
                                                    var v56 string = v2
                                                    v56 = "foo bar"
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    v57 := []int{fi18}
                                                    if v11 {
                                                        if true {
                                                            var v58 float32 = 1e10
                                                            _ = v58
                                                        }
                                                        var v59 int = 1
                                                        v60 := GenS5{F0: 32767, F1: -128}
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        var v61 uint8 = 127
                                                        _ = v61
                                                    }
                                                    fmt.Println(v2, v12, v17)
                                                    _ = v57
                                                }
                                                if (-32768 != -32768) {
                                                    v16 = v16
                                                } else {
                                                    switch v2 {
                                                    case "":
                                                        if (!v11) {
                                                            v17 -= (v17 << ((65535) & 63))
                                                            var v62 int16 = 42
                                                            _ = v62
                                                        }
                                                    case "foo bar":
                                                        v63 := []int{fi18, -9223372036854775808, fi24}
                                                        _ = v63
                                                    }
                                                }
                                                var v64 uint16 = (^(v16))
                                                v3[uint32(v17)] = (^(fi24))
                                                _ = v64
                                            } else {
                                                switch fi18 {
                                                case 32767:
                                                    v17 += v17
                                                case -430:
                                                    var v65 uint64 = (v17 >> ((940) & 63))
                                                    _ = v65
                                                case 0:
                                                    v66 := GenS6{F0: (v16 - v16), F1: 1}
                                                    _ = v66
                                                }
                                                v3[287] = (fi24 % (((fi24 >> ((65535) & 63))) | 1))
                                                v11 = true
                                            }
                                        }
                                        _ = fi24
                                    }
                                case 128:
                                    for fi67 := 0; fi67 < 10; fi67++ {
                                        if (true && (2147483647 > 2)) {
                                            v68 := []int{0, 0}
                                            fmt.Println(v1, fi67)
                                            _ = v68
                                        } else {
                                            if (!(-9223372036854775808 < 386)) {
                                                if (!(fi67 <= (^(fi18)))) {
                                                    if false {
                                                        if v11 {
                                                            fmt.Println(v17)
                                                            v69 := GenS7{F0: 757, F1: (v17 >> ((212) & 63))}
                                                            _ = v69
                                                        } else {
                                                            var v70 uint8 = 1
                                                            var v71 int8 = 127
                                                            fmt.Println(v70, v71)
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        if (1000 > 32767) {
                                                            if true {
                                                                v72 := GenS8{F0: 0, F1: 0, F2: v11}
                                                                v73 := []int{}
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                var v74 bool = (42 != 750)
                                                                if false {
                                                                    fmt.Println(v2)
                                                                    fmt.Println(v12)
                                                                    v75 := []int{(fi18 << ((4294967295) & 63))}
                                                                    _ = v75
                                                                }
                                                                v76 := GenS8{F0: 7, F1: -9223372036854775808, F2: v11}
                                                                _ = v74
                                                                _ = v76
                                                            }
                                                            v77 := []int{-737, (^(fi18)), fi67}
                                                            var v78 int8 = -1
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            v12 = (0.1 - v1)
                                                            v11 = false
                                                        }
                                                        if (252 >= 1) {
                                                            v79 := GenS9{F0: -1, F1: 0, F2: v11}
                                                            v16 = (v16 * v16)
                                                            v80 := GenS2{F0: uint64(v16), F1: 111, F2: 369}
                                                            if (-546 <= -1) {
                                                                v2 = v2
                                                                fmt.Println(v11, v17, v11)
                                                                v81 := GenS10{F0: 521, F1: -171, F2: 631}
                                                                switch (^(fi18)) {
                                                                case -1:
                                                                    v82 := v3[325]
                                                                    _ = v82
                                                                case 127:
                                                                    switch 2147483647 {
                                                                    case 2147483647:
                                                                        v83 := []int{fi67, (fi18 << ((65535) & 63)), fi18}
                                                                        _ = v83
                                                                    case 127:
                                                                        if (-9223372036854775808 == 0) {
                                                                            if v11 {
                                                                                var v84 uint32 = uint32(v17)
                                                                                _ = v84
                                                                            }
                                                                            if (2 == 68) {
                                                                                switch v2 {
                                                                                case "a":
                                                                                    v11 = v11
                                                                                }
                                                                                if v11 {
                                                                                    var v85 int32 = 2147483647
                                                                                    var v86 int32 = -1
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                } else {
                                                                                    _ = len(v3)
                                                                                }
                                                                            }
                                                                        }
                                                                    case -2147483648:
                                                                        fmt.Println(v11, v1, v17)
                                                                    }
                                                                case 42:
                                                                    if v11 {
                                                                        fmt.Println(v12, v1, v1)
                                                                        var v87 uint8 = 42
                                                                        _ = v87
                                                                    } else {
                                                                        v12 -= (float64(3.14159) / v12)
                                                                        if ("\n\t" != (v2 + ("\n\t" + (v2 + (v2 + v2))))) {
                                                                            fmt.Println(fi18, v11, v17)
                                                                            v88 := GenS8{F0: 127, F1: -240, F2: v11}
                                                                            _ = v88
                                                                        } else {
                                                                            v89 := v3[227]
                                                                            _ = v89
                                                                        }
                                                                        switch ("hello" + v2) {
                                                                        case "foo bar":
                                                                            v90 := GenS8{F0: 11, F1: -1, F2: v11}
                                                                            _ = v90
                                                                        case "":
                                                                            var v91 int32 = 0
                                                                            _ = v91
                                                                        default:
                                                                            v1 = (1e10 - 0.1)
                                                                        }
                                                                        v92 := []int{fi67, fi18}
                                                                        _ = v92
                                                                    }
                                                                }
                                                                _ = v81
                                                            } else {
                                                                v17 = (v17 | v17)
                                                                if false {
                                                                    var v93 uint8 = 82
                                                                    if (-2147483648 >= -825) {
                                                                        var v94 uint8 = (v93 | (v93 % ((255) | 1)))
                                                                        v1 -= 1.5
                                                                        v12 = (0.1 - 1e10)
                                                                        _ = v94
                                                                    } else {
                                                                        var v95 int16 = 0
                                                                        v96 := GenS7{F0: v17, F1: (v17 | (v17 << ((255) & 63)))}
                                                                        v12 -= ((v1 / 1.5) * 1.5)
                                                                        _ = v95
                                                                        _ = v96
                                                                    }
                                                                    _ = v93
                                                                } else {
                                                                    v1 -= v1
                                                                    if (988 < -1) {
                                                                        switch v2 {
                                                                        case "hello":
                                                                            var v97 int = (fi18 << ((65535) & 63))
                                                                            _ = v97
                                                                        default:
                                                                            if (v11 && v11) {
                                                                                switch fi67 {
                                                                                case 32767:
                                                                                    v12 -= 0.1
                                                                                case 32768:
                                                                                    var v98 uint8 = 100
                                                                                    _ = v98
                                                                                case 4294967295:
                                                                                    if (!(42 <= 42)) {
                                                                                        fmt.Println(v11)
                                                                                    } else {
                                                                                        switch v2 {
                                                                                        case "\n\t":
                                                                                            switch fi67 {
                                                                                            case 0:
                                                                                                var v99 uint16 = (^(v16))
                                                                                                _ = v99
                                                                                            case -2147483648:
                                                                                                if v11 {
                                                                                                    switch -573 {
                                                                                                    case 32768:
                                                                                                        v12 -= (v12 - v1)
                                                                                                    default:
                                                                                                        if (0 != 65535) {
                                                                                                            v100 := GenS8{F0: 0, F1: 1, F2: v11}
                                                                                                            fmt.Println(v1)
                                                                                                            v101 := GenS4{F0: 0}
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                        } else {
                                                                                                            v102 := GenS5{F0: -32768, F1: -1}
                                                                                                            v12 += 1.5
                                                                                                            var v103 float64 = (v12 / float64(1.5))
                                                                                                            _ = v102
                                                                                                            _ = v103
                                                                                                        }
                                                                                                        var v104 int16 = 26
                                                                                                        switch fi18 {
                                                                                                        case -235:
                                                                                                            switch v2 {
                                                                                                            case "hello":
                                                                                                                if (3.14159 < 3.14159) {
                                                                                                                    v105 := GenS8{F0: 2, F1: 692, F2: (false || true)}
                                                                                                                    v11 = (v11 && v11)
                                                                                                                    _ = v105
                                                                                                                }
                                                                                                            default:
                                                                                                                v1 += v12
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    v106 := GenS11{F0: fi18, F1: 42}
                                                                                                    _ = v106
                                                                                                }
                                                                                            default:
                                                                                                v11 = false
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v107 := GenS8{F0: 128, F1: -9223372036854775808, F2: v11}
                                                                                            _ = v107
                                                                                        case "":
                                                                                            switch ((((v2 + v2) + ("hello" + "a")) + ((v2 + "a") + "foo bar")) + (v2 + "\n\t")) {
                                                                                            case "\n\t":
                                                                                                if v11 {
                                                                                                    v108 := GenS7{F0: (v17 << ((789) & 63)), F1: (v17 ^ v17)}
                                                                                                    if (v17 > (v17 % ((v17) | 1))) {
                                                                                                        if v11 {
                                                                                                            switch v2 {
                                                                                                            case "foo bar":
                                                                                                                fmt.Println(v12, fi18, v16)
                                                                                                            case "\n\t":
                                                                                                                var v109 uint = 0
                                                                                                                _ = v109
                                                                                                            case "":
                                                                                                                if v11 {
                                                                                                                    if false {
                                                                                                                        var v110 uint16 = 950
                                                                                                                        _ = v110
                                                                                                                    }
                                                                                                                    v111 := GenS12{F0: 14, F1: v16}
                                                                                                                    v1 = v12
                                                                                                                    _ = v111
                                                                                                                }
                                                                                                            }
                                                                                                            var v112 int8 = -1
                                                                                                            _ = v112
                                                                                                        }
                                                                                                        if (true && (!(-943 < -32768))) {
                                                                                                            fmt.Println(v12, v16)
                                                                                                        } else {
                                                                                                            fmt.Println(v16)
                                                                                                        }
                                                                                                        var v113 string = ""
                                                                                                        _ = v113
                                                                                                    }
                                                                                                    if false {
                                                                                                        v12 = float64((1e10 - float32(v1)))
                                                                                                        v17 += 279
                                                                                                        if v11 {
                                                                                                            switch v2 {
                                                                                                            case "foo bar":
                                                                                                                v3[2147483647] = (fi18 * fi18)
                                                                                                            case "hello":
                                                                                                                v12 += (-2.25 * v12)
                                                                                                            case "":
                                                                                                                if (!v11) {
                                                                                                                    v1 = v12
                                                                                                                }
                                                                                                            }
                                                                                                            var v114 int64 = -9223372036854775808
                                                                                                            var v115 int32 = 444
                                                                                                            _ = v114
                                                                                                            _ = v115
                                                                                                        } else {
                                                                                                            v116 := []int{(^(fi18)), 422, fi18, (fi67 >> ((515) & 63))}
                                                                                                            v2 = v2
                                                                                                            if (!true) {
                                                                                                                if (!(v11 || v11)) {
                                                                                                                    fmt.Println(v17)
                                                                                                                    if ((("a" + v2) + v2) != "") {
                                                                                                                        v117 := GenS13{F0: (fi67 >> ((2) & 63)), F1: 956}
                                                                                                                        v11 = (163 > 2147483647)
                                                                                                                        v16 = 2
                                                                                                                        var v118 int8 = -62
                                                                                                                        _ = v117
                                                                                                                        _ = v118
                                                                                                                    } else {
                                                                                                                        v17 -= v17
                                                                                                                        fmt.Println(v11, v11, v2)
                                                                                                                        v12 = ((v1 * (v12 - 3.14159)) * 0.1)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v119 := GenS11{F0: fi18, F1: 32767}
                                                                                                                    _ = v119
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v116
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v108
                                                                                                }
                                                                                            case "a":
                                                                                                var v120 int = fi18
                                                                                                _ = v120
                                                                                            default:
                                                                                                switch (fi67 - (fi67 & 331)) {
                                                                                                case 1:
                                                                                                    v16 += (v16 ^ (^(v16)))
                                                                                                case 42:
                                                                                                    v16 -= (v16 / (((^(v16))) | 1))
                                                                                                    v121 := GenS14{F0: 127, F1: 65, F2: 65535}
                                                                                                    fmt.Println(fi18)
                                                                                                    _ = v121
                                                                                                case -382:
                                                                                                    if (((!v11) || (v11 || v11)) && (!v11)) {
                                                                                                        v12 = -2.25
                                                                                                        var v122 bool = true
                                                                                                        _ = v122
                                                                                                    }
                                                                                                default:
                                                                                                    var v123 uint32 = 2
                                                                                                    _ = v123
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v1)
                                                                                            if (1e10 <= float32(1e10)) {
                                                                                                if (v11 || (!false)) {
                                                                                                    if (!(!v11)) {
                                                                                                        v12 += 3.14159
                                                                                                        v12 -= (3.14159 * (((v12 + v12) + 3.14159) / v12))
                                                                                                        fmt.Println(fi18)
                                                                                                    } else {
                                                                                                        v124 := []int{627}
                                                                                                        var v125 int = fi18
                                                                                                        var v126 uint32 = 4294967295
                                                                                                        _ = v124
                                                                                                        _ = v125
                                                                                                        _ = v126
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    var v127 int16 = 798
                                                                                                    if (100 <= -128) {
                                                                                                        v12 = v1
                                                                                                        _ = len(v3)
                                                                                                        var v128 bool = (fi67 > fi18)
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    if (1 <= (^(fi67))) {
                                                                                                        var v129 int8 = 127
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            v130 := []int{(fi18 << ((281) & 63)), (fi67 / ((int(v17)) | 1))}
                                                                                                            _ = v130
                                                                                                        case "foo bar":
                                                                                                            v131 := GenS15{F0: v17, F1: -1}
                                                                                                            switch "foo bar" {
                                                                                                            case "foo bar":
                                                                                                                var v132 int32 = 2147483647
                                                                                                                _ = v132
                                                                                                            default:
                                                                                                                fmt.Println(v17, v2)
                                                                                                            }
                                                                                                            var v133 int16 = -32768
                                                                                                            _ = v131
                                                                                                            _ = v133
                                                                                                        default:
                                                                                                            var v134 uint32 = uint32(v16)
                                                                                                            _ = v134
                                                                                                        }
                                                                                                        var v135 int8 = -128
                                                                                                        _ = v129
                                                                                                        _ = v135
                                                                                                    } else {
                                                                                                        var v136 int64 = 0
                                                                                                        switch fi18 {
                                                                                                        case -191:
                                                                                                            v1 += 1e10
                                                                                                        default:
                                                                                                            v137 := []int{}
                                                                                                            _ = v137
                                                                                                        }
                                                                                                        _ = v136
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(fi67)
                                                                                                var v138 bool = v11
                                                                                                v139 := GenS16{F0: -2147483648, F1: v17}
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v140 bool = ((((1.5 + 1.5) + 3.14159) / 3.14159) >= (((v1 + (v12 + -2.25)) / float64(3.14159)) * v12))
                                                                                    _ = v140
                                                                                default:
                                                                                    v141 := []int{2147483647, fi67, (fi67 | 42)}
                                                                                    _ = v141
                                                                                }
                                                                                switch v2 {
                                                                                case "":
                                                                                    switch v2 {
                                                                                    case "":
                                                                                        v17 = 0
                                                                                    case "a":
                                                                                        v142 := GenS7{F0: (v17 / (((^(v17))) | 1)), F1: v17}
                                                                                        _ = v142
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            if true {
                                                                                                var v143 int32 = 1
                                                                                                _ = v143
                                                                                            }
                                                                                            v144 := []int{fi18, fi18}
                                                                                            _ = v144
                                                                                        }
                                                                                        v145 := []int{fi18, 1, fi67}
                                                                                        switch v2 {
                                                                                        case "hello":
                                                                                            v16 = 213
                                                                                        }
                                                                                        _ = v145
                                                                                    default:
                                                                                        fmt.Println(v17, v12)
                                                                                        if v11 {
                                                                                            switch (fi18 % ((fi18) | 1)) {
                                                                                            case -2147483648:
                                                                                                v146 := []int{fi18}
                                                                                                _ = v146
                                                                                            }
                                                                                            var v147 uint64 = (v17 / ((v17) | 1))
                                                                                            v12 += v1
                                                                                            _ = v147
                                                                                        }
                                                                                        switch v2 {
                                                                                        case "\n\t":
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                v148 := GenS16{F0: -1, F1: 248}
                                                                                                _ = v148
                                                                                            case "hello":
                                                                                                v149 := []int{(fi67 - fi67), fi18, fi67}
                                                                                                _ = v149
                                                                                            }
                                                                                        case "":
                                                                                            v16 = 405
                                                                                        case "foo bar":
                                                                                            var v150 int8 = 42
                                                                                            _ = v150
                                                                                        default:
                                                                                            var v151 uint32 = 4294967295
                                                                                            _ = v151
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch (fi67 / (((fi67 << ((65535) & 63))) | 1)) {
                                                                                case 128:
                                                                                    if true {
                                                                                        var v152 float64 = (-2.25 - (1.5 - ((((1e10 + v12) - (0.0 / v12)) + v12) / (v12 * (0.1 * v12)))))
                                                                                        var v153 int8 = -128
                                                                                        v154 := GenS7{F0: v17, F1: (v17 % (((v17 ^ (v17 >> ((255) & 63)))) | 1))}
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                        _ = v154
                                                                                    }
                                                                                case -841:
                                                                                    var v155 uint8 = 2
                                                                                    _ = v155
                                                                                case 0:
                                                                                    v11 = (42 != 203)
                                                                                    switch (v2 + "") {
                                                                                    case "a":
                                                                                        v1 -= v1
                                                                                    default:
                                                                                        v11 = true
                                                                                    }
                                                                                    v156 := GenS17{F0: 896, F1: true}
                                                                                    v12 += 0.1
                                                                                    _ = v156
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v2, v17)
                                                                }
                                                            }
                                                            v1 -= (0.1 * v1)
                                                            _ = v79
                                                            _ = v80
                                                        } else {
                                                            v157 := GenS7{F0: (v17 % ((v17) | 1)), F1: 1}
                                                            var v158 int64 = int64(fi67)
                                                            _ = len(v3)
                                                            _ = v157
                                                            _ = v158
                                                        }
                                                    }
                                                    var v159 bool = true
                                                    var v160 int32 = -777
                                                    _ = v159
                                                    _ = v160
                                                } else {
                                                    switch (fi18 / ((0) | 1)) {
                                                    case 2147483647:
                                                        if (v16 >= 2) {
                                                            delete(v3, 42)
                                                        }
                                                    default:
                                                        if (float64(0.1) == v1) {
                                                            fmt.Println(v17, v12)
                                                            v2 = "hello"
                                                            switch (fi18 + -2147483648) {
                                                            case 42:
                                                                fmt.Println(v17)
                                                            case 0:
                                                                v11 = (!v11)
                                                            case 392:
                                                                v2 = ((v2 + (v2 + v2)) + "a")
                                                            }
                                                        }
                                                    }
                                                    var v161 uint16 = 32767
                                                    v16 -= 2
                                                    var v162 uint = 0
                                                    if (!(v17 == (v17 << (((v162 ^ 2)) & 63)))) {
                                                        if v11 {
                                                            v162 -= v162
                                                            v163 := GenS6{F0: (v161 + (v161 * uint16(v162))), F1: 316}
                                                            var v164 int16 = -824
                                                            _ = v163
                                                            _ = v164
                                                        }
                                                        v161 = (v161 ^ 29)
                                                    }
                                                    _ = v161
                                                    _ = v162
                                                }
                                                var v165 bool = v11
                                                _ = v165
                                            }
                                            switch (fi18 / (((fi18 ^ fi67)) | 1)) {
                                            case 42:
                                                fmt.Println(v12)
                                                v166 := v3[65535]
                                                fmt.Println(v2)
                                                var v167 uint64 = 220
                                                if (!v11) {
                                                    var v168 int = 127
                                                    var v169 bool = (-2147483648 != 42)
                                                    delete(v3, 923)
                                                    _ = v168
                                                    _ = v169
                                                }
                                                switch fi67 {
                                                case 483:
                                                    fmt.Println(v17)
                                                }
                                                _ = v166
                                                _ = v167
                                            case -128:
                                                var v170 int8 = 95
                                                _ = v170
                                            default:
                                                _ = len(v3)
                                            }
                                            fmt.Println(v2, v11)
                                        }
                                        _ = fi67
                                    }
                                case 9223372036854775807:
                                    v17 = v17
                                }
                            case -1:
                                switch v2 {
                                case "foo bar":
                                    if v11 {
                                        if v11 {
                                            if v11 {
                                                switch (fi18 + (fi18 | fi18)) {
                                                case 128:
                                                    var v171 int16 = int16(fi18)
                                                    _ = v171
                                                case 2147483647:
                                                    v2 = "a"
                                                default:
                                                    var v172 int32 = 2147483647
                                                    _ = v172
                                                }
                                            }
                                            if (546 <= 671) {
                                                v173 := v3[1]
                                                v2 = "foo bar"
                                                _ = v173
                                            } else {
                                                v2 = (v2 + v2)
                                                var v174 int32 = -1
                                                v1 -= (((v1 * -2.25) * ((((3.14159 + (3.14159 * ((v12 + (v12 * v1)) * float64(1.5)))) - float64((0.0 / 1e10))) + v12) / (float64((1e10 / 1.5)) - ((v12 - v1) + float64((0.0 + 3.14159)))))) * ((0.1 - v1) + 0.0))
                                                _ = v174
                                            }
                                            var v175 uint8 = 206
                                            _ = v175
                                        } else {
                                            if ((v2 > (v2 + v2)) || (159 == 42)) {
                                                v176 := v3[42]
                                                v177 := v3[563]
                                                _ = v176
                                                _ = v177
                                            } else {
                                                if true {
                                                    var v178 float64 = (v12 * (v1 * -2.25))
                                                    _ = len(v3)
                                                    fmt.Println(v1, v17)
                                                    _ = v178
                                                }
                                                v12 = (((-2.25 - float64((3.14159 + 1e10))) + float64((((0.1 - float32(v12)) / 3.14159) * 3.14159))) * ((0.0 + v12) * (float64(0.0) / ((v12 - ((v12 - float64(-2.25)) * v12)) / (((1e10 - (((v12 + (0.1 + ((0.1 * 1e10) - 0.0))) * 1.5) - (v12 + -2.25))) * v12) + (3.14159 - v12))))))
                                                switch 0 {
                                                case -2147483648:
                                                    v179 := []int{int(v17)}
                                                    _ = v179
                                                case -197:
                                                    if false {
                                                        fmt.Println(v11)
                                                        var v180 int64 = 8
                                                        _ = v180
                                                    }
                                                case 2147483647:
                                                    v16 -= v16
                                                }
                                            }
                                        }
                                        if ((0 == 214) || v11) {
                                            v181 := GenS18{F0: fi18}
                                            v182 := []int{}
                                            _ = v181
                                            _ = v182
                                        }
                                        v183 := GenS15{F0: 4294967295, F1: -1}
                                        _ = v183
                                    }
                                default:
                                    v184 := GenS10{F0: v17, F1: 268, F2: -711}
                                    _ = v184
                                }
                            default:
                                switch (v2 + "hello") {
                                case "hello":
                                    v185 := GenS16{F0: -2147483648, F1: 1}
                                    _ = v185
                                case "a":
                                    var v186 uint64 = 356
                                    _ = v186
                                default:
                                    _ = len(v3)
                                }
                            }
                        }
                    }
                    var v187 int16 = 42
                    v187 = int16(v17)
                    _ = v187
                }
            }
            _ = fi18
        }
    } else {
        var v188 int = 42
        v11 = v11
        _ = v188
    }
    var v189 uint16 = 32767
    var v190 uint16 = (v16 + v16)
    v1 -= (v1 - v1)
    switch v2 {
    case "":
        v12 -= 1e10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v189:", v189)
    fmt.Println("v190:", v190)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v189
    _ = v190
    _ = v2
    _ = v3
}
