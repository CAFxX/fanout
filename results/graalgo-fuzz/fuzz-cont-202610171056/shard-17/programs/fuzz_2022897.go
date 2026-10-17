package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int8; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int8; F2 uint }
type GenS9 struct { F0 string }
func genFn0() int8 {
    var v5 uint16 = 32767
    v6 := []int{9223372036854775807, 4294967295, 97, 9223372036854775807}
    if true {
        var v7 bool = false
        var v8 int8 = -68
        var v9 uint32 = 2
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 2
        var v11 int32 = 1
        for fi12 := 0; fi12 < 9; fi12++ {
            if (uint64(v5) <= 18446744073709551615) {
                var v13 int = -464
                var v14 int64 = 0
                _ = v13
                _ = v14
            }
            if false {
                v11 += (v11 >> ((42) & 63))
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 uint16 = v5
                    _ = len(v6)
                    if (!true) {
                        var v17 int64 = -1
                        v11 = (v11 - int32(fi15))
                        v18 := GenS1{F0: v11, F1: (v10 << ((65535) & 63))}
                        _ = v17
                        _ = v18
                    } else {
                        switch (^(fi15)) {
                        case 127:
                            v10 = v10
                        case -2147483648:
                            if (-2.25 != 1e10) {
                                v5 = v5
                                if false {
                                    v19 := GenS1{F0: v11, F1: v10}
                                    _ = v19
                                } else {
                                    var v20 uint32 = 7
                                    if true {
                                        v11 -= 42
                                        v21 := GenS1{F0: -457, F1: (v10 % ((v10) | 1))}
                                        v16 = (v5 | v5)
                                        _ = v21
                                    } else {
                                        var v22 int16 = 1000
                                        v10 = (^(v10))
                                        v22 += v22
                                        _ = v22
                                    }
                                    switch (((("a" + ("foo bar" + "foo bar")) + ("" + "foo bar")) + "foo bar") + ((("\n\t" + "") + "\n\t") + ("\n\t" + ("a" + "\n\t")))) {
                                    case "a":
                                        if (v16 >= v5) {
                                            v20 = (v20 & 2147483648)
                                            var v23 int32 = (v11 | (-(v11)))
                                            _ = v23
                                        } else {
                                            var v24 float32 = 0.0
                                            _ = v24
                                        }
                                    case "":
                                        if len(v6) > 0 {
                                            v6[0] = 9223372036854775807
                                        }
                                    }
                                    v25 := GenS1{F0: v11, F1: (v10 * (^(v10)))}
                                    var v26 uint32 = v20
                                    var v27 string = "hello"
                                    _ = v20
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                if (32767 == 42) {
                                    var v28 uint32 = 4294967295
                                    _ = v28
                                } else {
                                    v29 := GenS2{F0: (fi12 | 32767), F1: true}
                                    v10 = uint8(v5)
                                    _ = v29
                                }
                                var v30 int16 = -32768
                                _ = v30
                            }
                        case 288:
                            var v31 int32 = (v11 % (((v11 | v11)) | 1))
                            _ = v31
                        }
                        if true {
                            if true {
                                switch (("hello" + "hello") + "hello") {
                                case "\n\t":
                                    v5 += v16
                                }
                                v32 := GenS1{F0: -100, F1: 127}
                                _ = v32
                            }
                        }
                        v33 := GenS3{F0: -803}
                        _ = v33
                    }
                    v11 = (^(v11))
                    fmt.Println(v11, v16)
                    _ = fi15
                    _ = v16
                }
                if len(v6) > 0 {
                    v6[0] = fi12
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    if (v5 < v5) {
        if (793 > 32767) {
            var v34 int32 = -838
            _ = v34
        }
    } else {
        var v35 uint64 = 42
        _ = v35
    }
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 24
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 int8 = 100
    var v3 uint = 2
    var v4 uint32 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS4{F0: 0, F1: v2}
            var v39 uint64 = 18446744073709551615
            _ = fi37
            _ = v38
            _ = v39
        }
        v40 := GenS5{F0: 2}
        if (!false) {
            if true {
                if (!false) {
                    v3 = 4294967295
                    for fi41 := 0; fi41 < 18; fi41++ {
                        var v42 uint = (v3 / ((v3) | 1))
                        _ = fi41
                        _ = v42
                    }
                    fmt.Println(v2, v4)
                } else {
                    if (!(0.0 <= 1.5)) {
                        var v43 int16 = 0
                        _ = v43
                    } else {
                        if false {
                            var v44 bool = true
                            if false {
                                var v45 uint = (^(v3))
                                var v46 string = "hello"
                                for fi47 := 0; fi47 < 5; fi47++ {
                                    switch v46 {
                                    case "\n\t":
                                        v45 += 221
                                    case "a":
                                        var v48 uint8 = 42
                                        _ = v48
                                    }
                                    _ = fi47
                                }
                                _ = v45
                                _ = v46
                            } else {
                                var v49 int16 = -1
                                v2 = v2
                                _ = v49
                            }
                            _ = v44
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 -= (v2 * (v2 * v2))
                    }
                    v4 = 792
                    v50 := GenS4{F0: 4294967295, F1: -117}
                    _ = v50
                }
                var v51 bool = ("a" > ("\n\t" + "a"))
                _ = v51
            } else {
                if (!true) {
                    _ = len(v1)
                }
            }
            v4 -= v4
            if false {
                var v52 int = (fi36 >> (((v3 >> (((^(v3))) & 63))) & 63))
                if len(v1) > 0 {
                    v1[0] = ("" + "\n\t")
                }
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = v52
            } else {
                switch "hello" {
                case "\n\t":
                    switch (fi36 - fi36) {
                    case 0:
                        var v53 int = (fi36 * 1)
                        _ = v53
                    case 190:
                        var v54 uint8 = 127
                        _ = v54
                    case 32767:
                        v3 = v3
                    }
                case "":
                    fmt.Println(v4, v4, fi36)
                default:
                    switch 448 {
                    case 781:
                        v2 = -128
                    default:
                        var v55 float64 = -2.25
                        _ = v55
                    }
                }
                if true {
                    var v56 int = 0
                    switch fi36 {
                    case -2147483648:
                        var v57 uint32 = uint32(v3)
                        _ = v57
                    case 729:
                        v56 = (-(v56))
                    case 128:
                        for fi58 := 0; fi58 < 17; fi58++ {
                            var v59 uint8 = 196
                            fmt.Println(fi36, v59, v3)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi58
                            _ = v59
                        }
                    }
                    var v60 int = fi36
                    _ = v56
                    _ = v60
                } else {
                    var v61 float64 = (1e10 + (float64(1e10) / ((1e10 / float64((1.5 / 0.1))) - ((0.0 * float64((1e10 + 0.1))) - 0.0))))
                    var v62 string = ("\n\t" + "foo bar")
                    _ = v61
                    _ = v62
                }
            }
        } else {
            v2 -= v2
            fmt.Println(v2, fi36)
        }
        _ = fi36
        _ = v40
    }
    var v63 string = "hello"
    if false {
        fmt.Println(v4)
        var v64 uint8 = 255
        _ = v64
    }
    var v65 uint64 = 2
    v65 = (^(v65))
    var v66 float64 = 1.5
    for fi67 := 0; fi67 < 6; fi67++ {
        if (127 > 127) {
            v63 = "hello"
            for fi68 := 0; fi68 < 16; fi68++ {
                v69 := GenS6{F0: false, F1: (v2 | (v2 * (v2 ^ genFn0()))), F2: (((((((((((("\n\t" + "hello") + v63) + v63) + (v63 + v63)) + v63) + v63) + v63) + "") + ((v63 + ((v63 + (v63 + v63)) + ("a" + v63))) + "\n\t")) + "\n\t") + v63) + v63)}
                var v70 int64 = -97
                _ = fi68
                _ = v69
                _ = v70
            }
        } else {
            v3 -= uint(v65)
            if (17 <= 2) {
                v1 = append(v1, "")
            }
        }
        _ = fi67
    }
    for fi71 := 0; fi71 < 19; fi71++ {
        var v72 uint32 = (^(v4))
        v73 := GenS7{F0: 255}
        _ = fi71
        _ = v72
        _ = v73
    }
    var v74 bool = false
    switch -2147483648 {
    case 42:
        v66 -= 1e10
    case 252:
        v75 := GenS6{F0: v74, F1: genFn0(), F2: "a"}
        _ = v75
    case -692:
        if (-889 != 42) {
            if false {
                var v76 uint16 = 466
                _ = v76
            }
        } else {
            switch v63 {
            case "":
                v77 := GenS8{F0: v65, F1: (v2 << ((v3) & 63)), F2: v3}
                _ = v77
            case "hello":
                for fi78 := 0; fi78 < 8; fi78++ {
                    for fi79 := 0; fi79 < 1; fi79++ {
                        if v74 {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            fmt.Println(fi78, v2)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (!(-367 == 42)) {
                            v80 := GenS7{F0: 255}
                            v2 = -128
                            if (!(((v74 || v74) && true) || true)) {
                                if v74 {
                                    v66 -= (v66 + v66)
                                } else {
                                    fmt.Println(v65)
                                    var v81 float64 = ((((float64((1e10 * 0.0)) - 3.14159) * ((v66 + 0.0) / v66)) * v66) + v66)
                                    _ = v81
                                }
                                v1 = append(v1, v63)
                            } else {
                                var v82 float64 = 0.0
                                if false {
                                    v83 := GenS6{F0: (22 < uint8(v4)), F1: v2, F2: v63}
                                    if v74 {
                                        v66 -= 1.5
                                    }
                                    _ = v83
                                } else {
                                    v82 = float64((1e10 + 3.14159))
                                    v84 := GenS6{F0: v74, F1: v2, F2: "a"}
                                    v74 = (((!v74) || v74) && v74)
                                    if (!(-234 != 1)) {
                                        var v85 float64 = v82
                                        v86 := GenS4{F0: v65, F1: (v2 >> ((v3) & 63))}
                                        if len(v1) > 0 {
                                            v1[0] = v63
                                        }
                                        _ = v85
                                        _ = v86
                                    } else {
                                        var v87 int = -1
                                        switch int(v4) {
                                        case 0:
                                            v66 = 3.14159
                                        }
                                        _ = v87
                                    }
                                    if ((v65 & v65) <= (v65 - (v65 & (v65 ^ v65)))) {
                                        v88 := GenS2{F0: (fi78 - (-(fi78))), F1: v74}
                                        _ = v88
                                    }
                                    _ = v84
                                }
                                v89 := GenS9{F0: ("foo bar" + v63)}
                                var v90 bool = false
                                _ = v82
                                _ = v89
                                _ = v90
                            }
                            _ = v80
                        }
                        _ = fi79
                    }
                    v66 += v66
                    _ = fi78
                }
            default:
                var v91 int16 = 1
                _ = v91
            }
            var v92 uint32 = uint32(v65)
            _ = v92
        }
    }
    if (((v65 >> (((^(v3))) & 63)) < (v65 % ((v65) | 1))) && true) {
        var v93 uint16 = 44
        if v74 {
            var v94 int16 = 1
            if v74 {
                var v95 uint8 = 127
                _ = v95
            }
            _ = v94
        } else {
            var v96 int = -377
            v93 = (v93 % ((uint16(v3)) | 1))
            _ = v96
        }
        if (((!(v74 && (42 >= -1))) && true) || ((v93 < (v93 >> (((v3 ^ (v3 / ((v3) | 1)))) & 63))) && (v74 && false))) {
            if len(v1) > 0 {
                v1[0] = v63
            }
            _ = len(v1)
        }
        _ = v93
    }
    v3 = (v3 | 65535)
    v63 = (v63 + "foo bar")
    if false {
        v3 -= v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v65
    _ = v66
    _ = v74
}
