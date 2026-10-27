package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint64; F2 bool }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int16; F1 uint8; F2 bool }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 int8 = (-(v1))
    delete(v3, ("a" + "hello"))
    if (false && true) {
        if (!false) {
            v5 := v3["hello"]
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 int8 = -1
                var v8 uint8 = 196
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        }
        if false {
            if true {
                if false {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        v10 := GenS1{F0: 97, F1: 948}
                        _ = fi9
                        _ = v10
                    }
                } else {
                    if ((128 < -1) && false) {
                        if true {
                            switch "hello" {
                            case "a":
                                if (1 >= -32768) {
                                    if true {
                                        for fi11 := 0; fi11 < 13; fi11++ {
                                            fmt.Println(fi11, v4)
                                            _ = fi11
                                        }
                                    }
                                } else {
                                    var v12 int16 = -32768
                                    _ = v12
                                }
                            case "\n\t":
                                v4 = v4
                            default:
                                v13 := GenS2{F0: 128, F1: 4294967295}
                                _ = v13
                            }
                        } else {
                            for fi14 := 0; fi14 < 2; fi14++ {
                                for fi15 := 0; fi15 < 15; fi15++ {
                                    if (4294967295 <= 2) {
                                        fmt.Println(fi15, v4, v4)
                                    }
                                    _ = fi15
                                }
                                _ = fi14
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            switch 9223372036854775807 {
                            case 4294967295:
                                if (0 >= 1) {
                                    if true {
                                        v1 = (^(v1))
                                    }
                                    v2[65535] = 128
                                    var v16 uint16 = 1
                                    _ = v16
                                }
                            case 905:
                                var v17 string = "\n\t"
                                _ = v17
                            }
                        case "foo bar":
                            v18 := GenS3{F0: ("" + "\n\t"), F1: 0}
                            _ = v18
                        default:
                            switch 32767 {
                            case -1:
                                v4 -= v1
                            case -102:
                                v1 -= (v1 << ((939) & 63))
                            }
                        }
                    }
                    v19 := GenS2{F0: 42, F1: 42}
                    _ = v19
                }
            } else {
                if true {
                    v1 = -1
                } else {
                    v1 = (v4 & (v4 * (v1 * (v1 * v4))))
                    for fi20 := 0; fi20 < 5; fi20++ {
                        fmt.Println(v4)
                        var v21 string = "hello"
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 uint8 = 1
                _ = v22
            }
        }
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            if (v1 <= (v1 + v4)) {
                var v24 float64 = 0.1
                if (!true) {
                    var v25 int64 = -1
                    for fi26 := 0; fi26 < 19; fi26++ {
                        switch "foo bar" {
                        case "hello":
                            v27 := GenS4{F0: 2147483647, F1: "", F2: 382}
                            _ = v27
                        case "\n\t":
                            v28 := GenS5{F0: true, F1: 166, F2: (4294967295 != 148)}
                            _ = v28
                        case "foo bar":
                            if false {
                                v24 -= (float64(((0.1 + (-2.25 / 3.14159)) / 0.1)) * v24)
                                fmt.Println(v4)
                                if (!true) {
                                    var v29 uint = 691
                                    fmt.Println(v24, fi23, v24)
                                    v30 := []int{fi26, (fi23 >> (((v29 | (v29 << ((v29) & 63)))) & 63)), (fi26 >> ((v29) & 63)), -9223372036854775808}
                                    _ = v29
                                    _ = v30
                                } else {
                                    switch ((("" + ((("a" + "") + "a") + "hello")) + "") + ("" + "foo bar")) {
                                    case "foo bar":
                                        v25 = -683
                                        v24 = (v24 - ((((3.14159 - 0.0) - v24) + 1.5) - v24))
                                    default:
                                        fmt.Println(fi26, v1, fi23)
                                    }
                                    fmt.Println(fi26, fi26, v24)
                                    var v31 uint8 = 22
                                    _ = v31
                                }
                            }
                        default:
                            v32 := []int{2147483647, fi26, fi23, fi26}
                            _ = v32
                        }
                        var v33 uint = 65535
                        var v34 int16 = -32768
                        var v35 bool = true
                        _ = fi26
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    _ = v25
                } else {
                    v24 -= float64(0.0)
                }
                _ = v24
            }
            for fi36 := 0; fi36 < 6; fi36++ {
                if true {
                    fmt.Println(fi23, v4, v1)
                    if false {
                        switch (fi23 ^ (-(fi23))) {
                        case 183:
                            fmt.Println(v1)
                        case 0:
                            v4 -= 0
                        case 42:
                            switch fi36 {
                            case 32767:
                                switch fi23 {
                                case 572:
                                    _ = len(v3)
                                }
                            case -1:
                                v1 -= 45
                            default:
                                if (!(float32(3.14159) != (1e10 / 1.0))) {
                                    switch "hello" {
                                    case "\n\t":
                                        v1 += (v4 + (v4 / ((v1) | 1)))
                                    case "":
                                        v4 -= v1
                                    }
                                } else {
                                    var v37 uint64 = 1
                                    _ = v37
                                }
                            }
                        }
                    }
                    var v38 uint = 42
                    _ = v38
                }
                v39 := GenS1{F0: 128, F1: 4294967295}
                v40 := []int{}
                _ = fi36
                _ = v39
                _ = v40
            }
            v41 := GenS3{F0: "hello", F1: 570}
            _ = fi23
            _ = v41
        }
    }
    v42 := []int{1, -740, 128, 1}
    if false {
        v4 = (v4 << ((65535) & 63))
        var v43 uint8 = 0
        v3["hello"] = 2
        _ = v43
    } else {
        var v44 int = 9223372036854775807
        _ = v44
    }
    if (851 != 0) {
        switch -645 {
        case 32767:
            if false {
                switch "foo bar" {
                case "hello":
                    var v45 int64 = 1
                    _ = v45
                }
                v1 = (v1 - v1)
                fmt.Println(v1, v4)
            }
        case 1:
            _ = len(v42)
        case -796:
            v1 = (v1 << ((646) & 63))
        }
    }
    switch "\n\t" {
    case "\n\t":
        v4 -= v4
    case "hello":
        var v46 int = 0
        _ = v46
    case "foo bar":
        switch -9223372036854775808 {
        case 0:
            v4 += (-(v1))
        }
    }
    v4 += (-(v4))
    v4 = (v1 << ((2) & 63))
    if (2 >= 461) {
        fmt.Println(v4)
        _ = len(v42)
    }
    if false {
        _ = len(v2)
    } else {
        _ = len(v2)
        for fi47 := 0; fi47 < 17; fi47++ {
            v48 := GenS5{F0: (true && true), F1: 833, F2: (0 < -897)}
            v42 = append(v42, fi47)
            for fi49 := 0; fi49 < 1; fi49++ {
                if (!true) {
                    var v50 uint = 538
                    switch (fi49 ^ fi49) {
                    case 270:
                        if len(v42) >= 2 {
                            v42 = v42[1:len(v42)]
                        }
                    }
                    _ = v50
                } else {
                    var v51 int64 = 9223372036854775807
                    var v52 uint = 65535
                    v53 := GenS3{F0: "foo bar", F1: 32767}
                    _ = v51
                    _ = v52
                    _ = v53
                }
                if ((fi49 - (fi47 + (fi49 - (fi49 & -9223372036854775808)))) == fi47) {
                    switch fi47 {
                    case -914:
                        var v54 bool = false
                        _ = v54
                    case 4294967295:
                        if false {
                            if (!(true && (true && (!(42 > 62))))) {
                                v55 := GenS2{F0: 395, F1: 0}
                                v4 -= 100
                                var v56 float32 = (1.5 / 1.5)
                                _ = len(v2)
                                fmt.Println(v56, v4, fi49)
                                fmt.Println(fi47)
                                v57 := GenS6{F0: "", F1: (("a" <= ("" + "foo bar")) || true)}
                                _ = v55
                                _ = v56
                                _ = v57
                            } else {
                                var v58 uint8 = 1
                                v42 = append(v42, fi47)
                                switch (fi47 - 39) {
                                case 697:
                                    switch -2147483648 {
                                    case -1:
                                        var v59 int8 = (v4 - (v1 | v1))
                                        _ = v59
                                    default:
                                        v60 := GenS2{F0: (fi47 ^ fi49), F1: 255}
                                        _ = v60
                                    }
                                case 4294967295:
                                    fmt.Println(fi49)
                                    var v61 uint8 = (v58 / ((v58) | 1))
                                    var v62 string = ""
                                    _ = v61
                                    _ = v62
                                default:
                                    if true {
                                        switch fi49 {
                                        case 4294967295:
                                            if (2147483647 == 884) {
                                                var v63 string = ""
                                                v64 := GenS4{F0: 2147483647, F1: v63, F2: 774}
                                                _ = v63
                                                _ = v64
                                            } else {
                                                var v65 uint16 = 2
                                                if true {
                                                    fmt.Println(v1, v4)
                                                    v66 := GenS5{F0: true, F1: uint64(v58), F2: false}
                                                    if ((-(v1)) == (v1 & (^(v1)))) {
                                                        v67 := GenS3{F0: "", F1: 1000}
                                                        if (18446744073709551615 == 0) {
                                                            switch "" {
                                                            case "\n\t":
                                                                v68 := GenS7{F0: 330, F1: fi47}
                                                                _ = v68
                                                            case "hello":
                                                                _ = len(v2)
                                                            }
                                                            var v69 uint64 = 4294967295
                                                            _ = v69
                                                        } else {
                                                            var v70 int = (fi49 % ((fi49) | 1))
                                                            _ = v70
                                                        }
                                                        _ = v67
                                                    } else {
                                                        if false {
                                                            delete(v2, 2)
                                                            if len(v42) >= 2 {
                                                                v42 = v42[1:len(v42)]
                                                            }
                                                            v71 := GenS7{F0: 2147483647, F1: (^(fi47))}
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v66
                                                }
                                                _ = v65
                                            }
                                        }
                                        v4 += (v4 >> ((0) & 63))
                                    }
                                }
                                _ = v58
                            }
                        } else {
                            if false {
                                var v72 bool = false
                                _ = v72
                            } else {
                                v73 := GenS6{F0: ("hello" + ("\n\t" + ("foo bar" + ""))), F1: ((false || ("hello" > ("foo bar" + ""))) || false)}
                                v74 := GenS4{F0: -652, F1: "foo bar", F2: 2}
                                v1 = -4
                                _ = v73
                                _ = v74
                            }
                            var v75 string = "hello"
                            _ = v75
                        }
                    case 13:
                        var v76 int32 = 42
                        _ = v76
                    default:
                        if true {
                            _ = len(v42)
                        } else {
                            var v77 uint16 = 42
                            _ = len(v2)
                            v78 := v2[2]
                            var v79 uint = 1
                            v4 += v4
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                    }
                    v4 = 100
                } else {
                    switch "hello" {
                    case "a":
                        v4 += 116
                    case "":
                        if (2 >= 1) {
                            var v80 uint64 = 42
                            v81 := GenS8{F0: 1, F1: -9223372036854775808}
                            switch 2147483647 {
                            case 0:
                                v1 = v1
                            case 515:
                                if ((v4 % ((v1) | 1)) != (-(v1))) {
                                    v82 := GenS4{F0: 979, F1: "", F2: 42}
                                    _ = v82
                                }
                            case 42:
                                var v83 float64 = (-2.25 + (-2.25 * 1e10))
                                _ = v83
                            }
                            _ = v80
                            _ = v81
                        }
                    }
                }
                _ = fi49
            }
            _ = fi47
            _ = v48
        }
    }
    if (!true) {
        v42 = append(v42, 32768)
        var v84 uint16 = 2
        v85 := GenS9{F0: -459, F1: 127, F2: (1000 != -1)}
        _ = v84
        _ = v85
    }
    fmt.Println(v4, v1)
    var v86 int16 = 711
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v86
}
