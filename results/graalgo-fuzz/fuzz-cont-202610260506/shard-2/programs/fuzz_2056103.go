package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int8; F2 int }
type GenS6 struct { F0 uint8; F1 int8; F2 int32 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 bool; F2 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16; F1 int32; F2 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 uint; F1 int32; F2 int64 }
type GenS16 struct { F0 string; F1 int8; F2 string }
type GenS17 struct { F0 int8; F1 int }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]uint32)
    var v3 string = ("a" + "hello")
    var v4 int = 9223372036854775807
    var v5 uint32 = 2
    v6 := v2[("a" + v3)]
    v4 = v4
    var v7 float64 = -2.25
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: (^(v1)), F1: 0}
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := []int{fi8, -1}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v4, v5, v6)
        _ = fi12
    }
    fmt.Println(v7, v5)
    v13 := GenS1{F0: (v1 % ((207) | 1)), F1: 127}
    v14 := GenS2{F0: (v4 / (((v4 & (-(v4)))) | 1)), F1: v6}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint64 = 42
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS3{F0: (((false || (true && (-349 > 0))) || false) && (float32(v7) != ((3.14159 / 3.14159) / 3.14159))), F1: 788}
        switch v4 {
        case 29:
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint8 = (^(v1))
                if true {
                    var v21 bool = false
                    v22 := []int{(v4 % (((^(fi19))) | 1)), (fi19 ^ fi17)}
                    _ = v21
                    _ = v22
                } else {
                    v4 = fi17
                    if true {
                        if (1.5 <= 1e10) {
                            v23 := GenS2{F0: (-(fi19)), F1: v5}
                            if (v3 <= "") {
                                v24 := GenS3{F0: true, F1: 641}
                                v6 += v5
                                _ = v24
                            }
                            _ = v23
                        } else {
                            var v25 uint32 = 0
                            if true {
                                v20 = 220
                            }
                            _ = v25
                        }
                    }
                }
                _ = fi19
                _ = v20
            }
        default:
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := GenS4{F0: true}
                fmt.Println(fi26, fi17)
                fmt.Println(fi26, v3, v4)
                if false {
                    v1 = (v1 ^ (v1 * (v1 >> ((0) & 63))))
                    fmt.Println(v1)
                    var v28 float32 = (float32(-2.25) * 1e10)
                    _ = v28
                } else {
                    v6 -= v5
                    if false {
                        var v29 uint = 0
                        v3 = (v3 + ("foo bar" + "\n\t"))
                        _ = v29
                    } else {
                        v30 := []int{}
                        var v31 float64 = ((v7 - ((v7 * float64(1e10)) * 0.0)) - v7)
                        v1 += v1
                        _ = v30
                        _ = v31
                    }
                    switch ((v3 + ("" + (v3 + ("\n\t" + v3)))) + (v3 + (v3 + v3))) {
                    case "\n\t":
                        switch v3 {
                        case "foo bar":
                            var v32 uint16 = 32767
                            v33 := GenS5{F0: (v32 << ((2) & 63)), F1: -128, F2: v4}
                            _ = v32
                            _ = v33
                        }
                    }
                }
                v34 := GenS6{F0: 0, F1: -23, F2: -69}
                v35 := GenS4{F0: (false || (-1 <= -668))}
                var v36 int16 = 42
                if true {
                    v37 := []int{-502}
                    if len(v37) > 0 {
                        v37[0] = -1
                    }
                    _ = v37
                } else {
                    var v38 uint = 65535
                    if true {
                        var v39 int8 = 9
                        v6 += 420
                        v40 := GenS7{F0: true, F1: 1}
                        _ = v39
                        _ = v40
                    } else {
                        if (v7 < (v7 * (v7 - v7))) {
                            if true {
                                v36 = (v36 * 0)
                                v41 := GenS2{F0: (fi17 % (((-(fi26))) | 1)), F1: uint32(v38)}
                                _ = v41
                            } else {
                                v42 := GenS8{F0: v1}
                                v43 := GenS8{F0: v1}
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS3{F0: (1 != 42), F1: 0}
                            _ = v44
                        } else {
                            v45 := GenS2{F0: (fi17 & fi17), F1: v6}
                            v36 = -28
                            _ = v45
                        }
                        v46 := GenS9{F0: (v36 / ((v36) | 1))}
                        _ = v46
                    }
                    if (false && true) {
                        v38 = v38
                    } else {
                        v7 = v7
                        switch v3 {
                        case "a":
                            switch "a" {
                            case "hello":
                                if true {
                                    var v47 int64 = 253
                                    _ = v47
                                }
                            }
                        default:
                            var v48 bool = ((v1 & v1) >= v1)
                            _ = v48
                        }
                    }
                    _ = v38
                }
                _ = fi26
                _ = v27
                _ = v34
                _ = v35
                _ = v36
            }
        }
        var v49 float64 = float64(0.0)
        _ = fi17
        _ = v18
        _ = v49
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        switch (fi50 >> ((2) & 63)) {
        case 9223372036854775807:
            if (!true) {
                switch 0 {
                case -595:
                    v5 -= uint32(v1)
                }
            } else {
                v51 := GenS10{F0: v4, F1: true, F2: -707}
                _ = v51
            }
        case -9223372036854775808:
            fmt.Println(v4, fi50)
        case 0:
            v2["\n\t"] = v5
        }
        v52 := []int{(fi50 ^ v4), fi50, (fi50 / (((fi50 | (fi50 - v4))) | 1)), fi50}
        _ = fi50
        _ = v52
    }
    if true {
        for fi53 := 0; fi53 < 9; fi53++ {
            switch "hello" {
            case "foo bar":
                switch 875 {
                case -781:
                    switch 106 {
                    case 32767:
                        var v54 int16 = -811
                        _ = v54
                    case 127:
                        for fi55 := 0; fi55 < 20; fi55++ {
                            var v56 int32 = -1
                            v6 -= v6
                            if true {
                                if true {
                                    v57 := []int{}
                                    v5 -= (v5 | v6)
                                    var v58 uint = 255
                                    _ = v57
                                    _ = v58
                                }
                                switch -506 {
                                case -1:
                                    if true {
                                        _ = len(v2)
                                        if false {
                                            switch v3 {
                                            case "":
                                                v56 = v56
                                            default:
                                                var v59 uint8 = (^(v1))
                                                _ = v59
                                            }
                                            var v60 int64 = 0
                                            switch 127 {
                                            case 127:
                                                if (false || true) {
                                                    v4 -= 0
                                                    v61 := GenS11{F0: v1}
                                                    if ((v1 - 180) > v1) {
                                                        v62 := GenS12{F0: -845, F1: (^(v56)), F2: v1}
                                                        _ = v62
                                                    } else {
                                                        var v63 uint32 = (v5 | (^(v6)))
                                                        _ = v63
                                                    }
                                                    _ = v61
                                                } else {
                                                    var v64 int32 = v56
                                                    if (2147483647 != 657) {
                                                        if (!true) {
                                                            var v65 uint16 = 1
                                                            v3 = ((v3 + (v3 + v3)) + ("" + v3))
                                                            if (!(v65 != (v65 & v65))) {
                                                                v66 := GenS13{F0: (v56 * 42)}
                                                                _ = v66
                                                            } else {
                                                                v67 := []int{fi53, (fi55 * -529)}
                                                                v7 += v7
                                                                v68 := GenS5{F0: v65, F1: 13, F2: -273}
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            _ = v65
                                                        } else {
                                                            delete(v2, v3)
                                                            var v69 int = 1
                                                            v56 += v64
                                                            _ = v69
                                                        }
                                                        switch v3 {
                                                        case "hello":
                                                            v70 := GenS7{F0: false, F1: 358}
                                                            _ = v70
                                                        case "":
                                                            v2[v3] = (v5 ^ (v6 - 0))
                                                        default:
                                                            switch -2147483648 {
                                                            case 4294967295:
                                                                v7 = (-2.25 + v7)
                                                            }
                                                        }
                                                        var v71 uint32 = 65535
                                                        _ = v71
                                                    } else {
                                                        switch v4 {
                                                        case 127:
                                                            var v72 int16 = 1000
                                                            _ = v72
                                                        case 486:
                                                            v73 := GenS14{F0: 717, F1: (v1 << ((511) & 63))}
                                                            _ = v73
                                                        default:
                                                            v7 += v7
                                                        }
                                                    }
                                                    v1 = v1
                                                    _ = v64
                                                }
                                            default:
                                                v74 := []int{(fi55 >> ((65535) & 63)), 42}
                                                _ = v74
                                            }
                                            _ = v60
                                        } else {
                                            v75 := GenS2{F0: (v4 + 128), F1: v5}
                                            v76 := GenS9{F0: 104}
                                            v4 = int(v1)
                                            _ = v75
                                            _ = v76
                                        }
                                        v77 := GenS13{F0: (^(v56))}
                                        _ = v77
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            var v78 int = fi53
                                            _ = v78
                                        case "a":
                                            if false {
                                                v7 = (float64(3.14159) * (v7 - 1e10))
                                            } else {
                                                var v79 int8 = -44
                                                switch (fi55 >> ((4294967295) & 63)) {
                                                case 127:
                                                    if false {
                                                        v5 = 2147483647
                                                        var v80 uint32 = (v5 ^ v6)
                                                        v81 := []int{(fi55 * v4), fi53}
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        fmt.Println(v1, v1, v5)
                                                        var v82 uint16 = 1
                                                        v56 -= -281
                                                        if (!false) {
                                                            _ = len(v2)
                                                        }
                                                        v7 -= (float64(-2.25) / -2.25)
                                                        fmt.Println(fi53, v56)
                                                        v3 = (v3 + (v3 + v3))
                                                        _ = v82
                                                    }
                                                default:
                                                    if (!false) {
                                                        v1 += v1
                                                        v83 := []int{-2147483648}
                                                        var v84 int16 = -1
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                }
                                                if (v56 != 42) {
                                                    fmt.Println(v7)
                                                    switch (((v3 + (v3 + v3)) + ("a" + v3)) + v3) {
                                                    case "":
                                                        v79 = -128
                                                    }
                                                }
                                                _ = v79
                                            }
                                        }
                                        var v85 float64 = 1.5
                                        _ = v85
                                    }
                                case -2147483648:
                                    var v86 uint16 = 233
                                    _ = v86
                                default:
                                    var v87 uint32 = 100
                                    _ = v87
                                }
                                if ((121 < -128) && (9223372036854775807 > -1)) {
                                    v88 := []int{}
                                    _ = v88
                                }
                                v89 := []int{}
                                _ = v89
                            } else {
                                v90 := GenS9{F0: 32767}
                                _ = v90
                            }
                            _ = fi55
                            _ = v56
                        }
                    }
                case 2147483647:
                    for fi91 := 0; fi91 < 5; fi91++ {
                        switch "foo bar" {
                        case "a":
                            v6 = (v5 << ((65535) & 63))
                        case "\n\t":
                            v92 := []int{}
                            _ = v92
                        case "hello":
                            v6 += uint32(v1)
                        default:
                            v4 += 4294967295
                        }
                        v93 := []int{fi53, fi53, (fi91 >> ((4294967295) & 63))}
                        _ = fi91
                        _ = v93
                    }
                default:
                    v94 := GenS15{F0: uint(v6), F1: -627, F2: -1}
                    _ = v94
                }
            case "a":
                var v95 string = (v3 + "\n\t")
                _ = v95
            case "hello":
                delete(v2, "a")
            }
            v1 += uint8(v6)
            _ = fi53
        }
        var v96 int16 = -621
        delete(v2, "")
        _ = v96
    }
    v97 := GenS3{F0: true, F1: 0}
    if true {
        v7 += -2.25
        var v98 uint32 = (v5 << ((1) & 63))
        _ = v98
    } else {
        for fi99 := 0; fi99 < 9; fi99++ {
            var v100 uint = 4294967295
            for fi101 := 0; fi101 < 1; fi101++ {
                delete(v2, ("foo bar" + ("hello" + "a")))
                v102 := GenS1{F0: uint8(v100), F1: 42}
                _ = fi101
                _ = v102
            }
            v2[v3] = (v6 ^ v5)
            _ = fi99
            _ = v100
        }
        v3 = "a"
        switch (v4 * v4) {
        case 881:
            if (660 != 2147483648) {
                v103 := []int{v4, -9223372036854775808}
                if (false || true) {
                    v6 = (v6 << ((18446744073709551615) & 63))
                    var v104 uint = 93
                    _ = v104
                } else {
                    for fi105 := 0; fi105 < 6; fi105++ {
                        v106 := GenS16{F0: (v3 + (v3 + ("a" + "hello"))), F1: -128, F2: "\n\t"}
                        v3 = v3
                        _ = fi105
                        _ = v106
                    }
                    switch "" {
                    case "\n\t":
                        for fi107 := 0; fi107 < 4; fi107++ {
                            switch (fi107 | (^(v4))) {
                            case 128:
                                v5 += 0
                            default:
                                for fi108 := 0; fi108 < 18; fi108++ {
                                    fmt.Println(fi108, v6)
                                    v3 = v3
                                    if (100 < 1) {
                                        var v109 uint64 = 273
                                        _ = v109
                                    }
                                    _ = fi108
                                }
                            }
                            _ = fi107
                        }
                    default:
                        fmt.Println(v3, v1)
                    }
                }
                _ = v103
            }
        case 32768:
            v110 := GenS2{F0: 1, F1: v5}
            _ = v110
        default:
            if false {
                if (18446744073709551615 >= 433) {
                    for fi111 := 0; fi111 < 18; fi111++ {
                        v112 := []int{-2147483648, v4, (fi111 / ((42) | 1))}
                        _ = fi111
                        _ = v112
                    }
                }
                v113 := GenS17{F0: -128, F1: v4}
                _ = v113
            }
        }
    }
    v7 = 1.5
    var v114 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v114
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v97
}
