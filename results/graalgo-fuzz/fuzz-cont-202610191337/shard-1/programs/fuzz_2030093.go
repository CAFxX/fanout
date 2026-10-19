package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint16; F1 int64; F2 int32 }
type GenS9 struct { F0 int32; F1 int8; F2 int64 }
type GenS10 struct { F0 uint32; F1 int16; F2 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint64; F1 uint8 }
type GenS14 struct { F0 string; F1 int16; F2 uint64 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    var v3 int16 = v2
    v4 := GenS1{F0: -9223372036854775808, F1: (v2 * 42), F2: 255}
    var v5 uint64 = 0
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int8 = -1
        _ = fi6
        _ = v7
    }
    v5 += v5
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            v5 += 42
            v9 := GenS1{F0: 42, F1: v2, F2: 181}
            for fi10 := 0; fi10 < 3; fi10++ {
                if true {
                    fmt.Println(fi8, v3, fi10)
                    _ = len(v1)
                    v3 += (v2 * v3)
                }
                if false {
                    var v11 uint16 = 42
                    _ = v11
                } else {
                    v12 := GenS1{F0: 83, F1: v3, F2: 42}
                    _ = v12
                }
                v2 = 1
                if true {
                    if false {
                        if ((123 > -112) || false) {
                            if (!false) {
                                var v13 int16 = 927
                                if false {
                                    if true {
                                        switch fi8 {
                                        case 0:
                                            var v14 int16 = -1
                                            _ = v14
                                        default:
                                            var v15 uint16 = 42
                                            _ = v15
                                        }
                                        var v16 uint = 2
                                        v13 = -1
                                        _ = v16
                                    } else {
                                        var v17 bool = true
                                        switch "a" {
                                        case "":
                                            if true {
                                                if v17 {
                                                    if v17 {
                                                        if false {
                                                            if v17 {
                                                                v2 = (v2 | (v13 - (v3 + (-(v3)))))
                                                                var v18 int64 = -9223372036854775808
                                                                _ = v18
                                                            }
                                                            var v19 float32 = 0.0
                                                            _ = v19
                                                        }
                                                        _ = len(v1)
                                                    } else {
                                                        var v20 int = (^(fi8))
                                                        _ = v20
                                                    }
                                                    if true {
                                                        var v21 float32 = (1e10 + 0.0)
                                                        _ = v21
                                                    } else {
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v22 uint16 = 1
                                                            _ = v22
                                                        }
                                                        fmt.Println(v3, v2, v13)
                                                    }
                                                    v13 = (v2 + int16(v5))
                                                }
                                                v23 := GenS2{F0: 100}
                                                _ = v23
                                            }
                                        default:
                                            var v24 bool = false
                                            _ = v24
                                        }
                                        var v25 bool = v17
                                        var v26 float32 = (1e10 + (3.14159 * 1.5))
                                        _ = v17
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                if false {
                                    v5 = (^(v5))
                                    v1[4294967295] = (!false)
                                }
                                _ = v13
                            }
                        }
                        v2 = (v2 & v3)
                        if (!true) {
                            if true {
                                v27 := GenS3{F0: 1, F1: ("hello" + "a")}
                                _ = v27
                            }
                            if false {
                                var v28 uint8 = 128
                                fmt.Println(v2)
                                fmt.Println(fi10, v5, v3)
                                switch "hello" {
                                case "\n\t":
                                    var v29 uint16 = 2
                                    _ = v29
                                case "foo bar":
                                    v5 -= (v5 - (v5 << ((65535) & 63)))
                                case "a":
                                    var v30 int16 = (v2 | v2)
                                    _ = v30
                                default:
                                    var v31 uint64 = v5
                                    _ = v31
                                }
                                switch fi10 {
                                case 127:
                                    v32 := []int{fi8, 42}
                                    _ = v32
                                case 963:
                                    if ((!true) || (!false)) {
                                        fmt.Println(v3)
                                        if (2 > 42) {
                                            v5 += (^(v5))
                                            var v33 uint = 65535
                                            fmt.Println(v33, v28, v3)
                                            _ = v33
                                        }
                                        switch (fi10 + 2147483647) {
                                        case 2147483647:
                                            if (v5 < 18446744073709551615) {
                                                switch ("hello" + "foo bar") {
                                                case "\n\t":
                                                    fmt.Println(v2, fi10, v5)
                                                case "":
                                                    v34 := []int{fi10, -1, 9223372036854775807}
                                                    _ = v34
                                                case "foo bar":
                                                    var v35 uint16 = 65535
                                                    _ = v35
                                                default:
                                                    fmt.Println(v5, v2, fi8)
                                                }
                                                if false {
                                                    v28 = v28
                                                }
                                            }
                                        case -208:
                                            fmt.Println(v5, v5, v3)
                                        }
                                    }
                                }
                                v36 := GenS4{F0: (((("\n\t" + "\n\t") + "\n\t") + "a") + (("hello" + "hello") + "")), F1: 42}
                                _ = v28
                                _ = v36
                            } else {
                                var v37 uint8 = 1
                                _ = v37
                            }
                        } else {
                            if (!true) {
                                var v38 float64 = 1.5
                                delete(v1, 2147483647)
                                var v39 uint8 = 86
                                _ = v38
                                _ = v39
                            } else {
                                if ((!true) || true) {
                                    v40 := GenS5{F0: 92}
                                    fmt.Println(fi10, v5)
                                    _ = v40
                                }
                                v3 = (^(v3))
                                v41 := []int{(fi10 / ((fi10) | 1)), -1, (fi8 + fi10)}
                                _ = v41
                            }
                            v3 = v2
                            if true {
                                if (0 <= 32767) {
                                    if true {
                                        if (((((uint32(v5) >= 2147483648) || (!true)) && (1.5 <= 1e10)) && (0 <= 100)) || (true && (1 <= 140))) {
                                            v2 = (v2 ^ v3)
                                            v3 += (^(v2))
                                            var v42 int8 = -40
                                            fmt.Println(fi8, v42)
                                            if true {
                                                delete(v1, 2)
                                            } else {
                                                v43 := GenS6{F0: 42, F1: 0, F2: v3}
                                                _ = v43
                                            }
                                            var v44 float32 = 3.14159
                                            v44 = 0.1
                                            var v45 float64 = 1.5
                                            var v46 float64 = float64((((v44 / -2.25) * (1e10 / v44)) + (v44 * float32(v45))))
                                            _ = v42
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        v5 = v5
                                        v47 := []int{2147483647, fi8, 128, -2147483648}
                                        var v48 int16 = v2
                                        v49 := GenS7{F0: 2147483647, F1: 128}
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    } else {
                                        var v50 int = (fi8 * fi10)
                                        var v51 float32 = (1.5 - 0.0)
                                        v52 := GenS8{F0: 42, F1: 42, F2: 2147483647}
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 uint8 = 21
                                    var v54 int16 = (v2 - (-(v2)))
                                    var v55 float64 = 1e10
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                v56 := GenS9{F0: 0, F1: 42, F2: 0}
                                fmt.Println(fi10)
                                _ = v56
                            } else {
                                var v57 int64 = -598
                                v58 := []int{fi8}
                                _ = v57
                                _ = v58
                            }
                        }
                        fmt.Println(v3, fi8)
                        if (!(!true)) {
                            v59 := GenS10{F0: 761, F1: (v2 % ((v2) | 1)), F2: ""}
                            if false {
                                if (false || (0 == 4294967295)) {
                                    var v60 int16 = v2
                                    var v61 float32 = (-2.25 * 1e10)
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v62 int8 = 100
                                    _ = v62
                                }
                            } else {
                                fmt.Println(v3)
                                if (9223372036854775807 < int64(v2)) {
                                    fmt.Println(fi10, fi10)
                                }
                                var v63 int8 = 22
                                _ = v63
                            }
                            _ = v59
                        } else {
                            if true {
                                switch fi10 {
                                case 32767:
                                    if true {
                                        v3 -= v2
                                    }
                                case 694:
                                    fmt.Println(fi10, fi8, fi10)
                                case -589:
                                    switch fi8 {
                                    case 42:
                                        v64 := GenS11{F0: 838}
                                        _ = v64
                                    case 0:
                                        var v65 uint = 4294967295
                                        _ = v65
                                    default:
                                        var v66 uint64 = v5
                                        _ = v66
                                    }
                                }
                                if (1 <= 128) {
                                    var v67 bool = true
                                    var v68 int16 = -1
                                    _ = v67
                                    _ = v68
                                }
                                var v69 string = "hello"
                                var v70 float32 = (0.1 - 0.1)
                                if true {
                                    v70 = 1.5
                                    v69 = (((v69 + (v69 + "foo bar")) + ((v69 + ((((v69 + "\n\t") + v69) + v69) + v69)) + v69)) + "hello")
                                    v71 := GenS1{F0: -423, F1: int16(v5), F2: 118}
                                    _ = v71
                                } else {
                                    if true {
                                        v72 := GenS12{F0: (v5 & v5), F1: 32767}
                                        if (42 > 127) {
                                            if false {
                                                var v73 uint64 = v5
                                                v74 := GenS13{F0: v5, F1: 1}
                                                if true {
                                                    var v75 int = fi8
                                                    _ = v75
                                                }
                                                _ = v73
                                                _ = v74
                                            }
                                        } else {
                                            var v76 int16 = 32767
                                            _ = v76
                                        }
                                        _ = v72
                                    } else {
                                        if (!true) {
                                            v3 = (v2 ^ v3)
                                            v77 := GenS13{F0: (v5 * 42), F1: 0}
                                            _ = v77
                                        }
                                    }
                                }
                                var v78 int = fi8
                                v5 = 595
                                var v79 int32 = 70
                                _ = v69
                                _ = v70
                                _ = v78
                                _ = v79
                            } else {
                                v80 := []int{(fi8 ^ fi10), fi8, (fi8 >> ((255) & 63)), (fi10 ^ 9223372036854775807)}
                                var v81 string = "a"
                                _ = v80
                                _ = v81
                            }
                            v82 := v1[563]
                            var v83 uint = 2
                            if (1e10 <= 1.5) {
                                var v84 uint = v83
                                v5 = (v5 & (v5 - (v5 << ((255) & 63))))
                                _ = v84
                            }
                            _ = v82
                            _ = v83
                        }
                        switch 1 {
                        case -9223372036854775808:
                            fmt.Println(fi10)
                        }
                        if true {
                            v3 = (^(v2))
                        }
                        v85 := GenS14{F0: "a", F1: 212, F2: v5}
                        _ = v85
                    }
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v86 int8 = -116
        if false {
            fmt.Println(v2, v2)
            switch -546 {
            case 42:
                if ("foo bar" == ("hello" + "a")) {
                    v2 += -32768
                    v1[65535] = (true && false)
                } else {
                    var v87 int = 1
                    _ = v87
                }
            case 127:
                if false {
                    v88 := []int{402, -1, 128, 4294967295}
                    var v89 int16 = (v2 | -32768)
                    _ = v88
                    _ = v89
                }
            }
        } else {
            v5 = (v5 / (((v5 / ((0) | 1))) | 1))
            v86 += v86
        }
        _ = v86
    }
    var v90 uint = 0
    v3 = v3
    if true {
        v91 := GenS10{F0: 677, F1: (v2 & (v3 + v3)), F2: "foo bar"}
        if false {
            if true {
                var v92 float64 = 1e10
                var v93 string = "foo bar"
                _ = v92
                _ = v93
            } else {
                v90 = (v90 >> ((v90) & 63))
            }
            var v94 bool = (0 == 32767)
            _ = v94
        } else {
            for fi95 := 0; fi95 < 8; fi95++ {
                v5 = v5
                v96 := GenS15{F0: (!true), F1: uint16(v5)}
                _ = fi95
                _ = v96
            }
            v2 = (v3 / (((^(v2))) | 1))
            fmt.Println(v5, v90)
        }
        var v97 uint64 = v5
        _ = v91
        _ = v97
    }
    v98 := GenS5{F0: (v5 >> ((uint(v5)) & 63))}
    for fi99 := 0; fi99 < 1; fi99++ {
        v3 += (v3 / (((^(v2))) | 1))
        switch "a" {
        case "foo bar":
            v100 := GenS2{F0: 82}
            _ = v100
        case "\n\t":
            v5 -= (v5 + 18446744073709551615)
        }
        var v101 string = "a"
        _ = fi99
        _ = v101
    }
    v102 := GenS16{F0: 0}
    switch "hello" {
    case "foo bar":
        v2 += (v2 | v2)
    }
    fmt.Println(v2, v3)
    var v103 string = ((("hello" + (("hello" + "foo bar") + ("a" + "a"))) + "\n\t") + "")
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v98
}
