package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint64 }
func genFn0() int64 {
    if (91 != 411) {
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint16 = 42
                v8 := []int{(fi5 >> ((414) & 63))}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
        var v9 int16 = -539
        _ = v9
    } else {
        var v10 uint64 = 18446744073709551615
        var v11 int16 = 245
        v11 = (v11 & -1)
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: -1, F1: 9223372036854775807, F2: 4294967295}
    if false {
        var v13 uint32 = 652
        _ = v13
    }
    var v14 uint32 = 65535
    _ = v12
    _ = v14
    return 1
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    switch "" {
    case "a":
        var v15 int8 = 100
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint16 = 2
            for fi18 := 0; fi18 < 19; fi18++ {
                if (2 == 580) {
                    var v19 bool = false
                    _ = v19
                }
                fmt.Println(fi18, v17)
                v20 := GenS1{F0: 42, F1: -672, F2: 18446744073709551615}
                _ = fi18
                _ = v20
            }
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v1)
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(fi21)
        _ = fi21
    }
    v1 = v1
    v22 := GenS1{F0: 1, F1: -298, F2: 2}
    if false {
        var v23 float64 = 1e10
        _ = v23
    }
    v1 = v1
    v1 = -330
    if (-128 > 42) {
        v1 = 127
        if false {
            v1 -= v1
            if ((false && (-2147483648 == 42)) || false) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := GenS2{F0: -1}
                    v2[2147483647] = (v1 / (((fi24 & fi24)) | 1))
                    _ = fi24
                    _ = v25
                }
            } else {
                var v26 uint64 = 18446744073709551615
                _ = v26
            }
        }
    } else {
        fmt.Println(v1)
    }
    if false {
        var v27 string = "foo bar"
        var v28 int32 = 1
        switch (^(v1)) {
        case 0:
            v29 := GenS3{F0: 65535}
            _ = v29
        default:
            for fi30 := 0; fi30 < 16; fi30++ {
                v31 := GenS4{F0: 1, F1: 65535, F2: 0}
                _ = fi30
                _ = v31
            }
        }
        _ = v27
        _ = v28
    } else {
        var v32 int16 = -32768
        if false {
            var v33 int = -448
            _ = v33
        } else {
            var v34 string = (("foo bar" + ("foo bar" + "a")) + (("hello" + ("\n\t" + "a")) + ""))
            fmt.Println(v34, v1)
            delete(v2, 2147483648)
            _ = v34
        }
        _ = v32
    }
    var v35 int16 = -1
    if false {
        switch "a" {
        case "\n\t":
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 string = ""
                if false {
                    var v38 uint32 = 64
                    var v39 uint16 = 1
                    _ = v38
                    _ = v39
                } else {
                    v40 := GenS5{F0: 42}
                    v41 := []int{fi36}
                    for fi42 := 0; fi42 < 20; fi42++ {
                        _ = len(v41)
                        if false {
                            if len(v41) > 0 {
                                v41[0] = (^(v1))
                            }
                            fmt.Println(v35)
                            switch (v37 + "") {
                            case "":
                                var v43 int64 = -1
                                _ = v43
                            case "\n\t":
                                var v44 int16 = 1
                                _ = v44
                            case "hello":
                                var v45 int = (-(fi42))
                                _ = v45
                            }
                        }
                        _ = fi42
                    }
                    _ = v40
                    _ = v41
                }
                _ = fi36
                _ = v37
            }
        }
    } else {
        v35 += 819
    }
    switch (v1 << ((65535) & 63)) {
    case -9223372036854775808:
        v35 = 32767
    case 42:
        v46 := []int{v1, 2147483647, 4294967295, v1}
        _ = v46
    case 32767:
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 bool = (-989 >= 992)
            v35 += v35
            for fi49 := 0; fi49 < 10; fi49++ {
                if (0 != 528) {
                    if v48 {
                        if (genFn0() != 0) {
                            v50 := GenS6{F0: (fi49 & (fi47 + 0)), F1: 873, F2: v48}
                            if (!v48) {
                                v1 = 127
                            }
                            v3[""] = v48
                            _ = v50
                        }
                        v1 += -1
                        if (9223372036854775807 < 527) {
                            v35 = v35
                        }
                    } else {
                        v51 := GenS2{F0: 9223372036854775807}
                        switch 1 {
                        case 876:
                            fmt.Println(fi47, v48)
                        case 4294967295:
                            switch "hello" {
                            case "foo bar":
                                var v52 int = -1
                                _ = v52
                            }
                        case -9223372036854775808:
                            var v53 bool = (-2.25 != ((1.5 - (3.14159 - 3.14159)) * 0.0))
                            _ = v53
                        }
                        v54 := GenS7{F0: 255}
                        _ = v51
                        _ = v54
                    }
                    v55 := GenS2{F0: -158}
                    _ = v55
                }
                if ((("" + ("a" + "hello")) + ("foo bar" + "a")) > ("" + "\n\t")) {
                    if v48 {
                        switch "\n\t" {
                        case "\n\t":
                            v56 := v2[1]
                            _ = v56
                        case "hello":
                            var v57 int16 = -534
                            _ = v57
                        case "foo bar":
                            var v58 uint = 1
                            _ = v58
                        default:
                            fmt.Println(v35)
                            v59 := GenS1{F0: v35, F1: genFn0(), F2: 0}
                            if (("\n\t" + "a") != "") {
                                fmt.Println(fi49)
                                if v48 {
                                    var v60 int32 = 1
                                    v61 := GenS8{F0: fi49, F1: 42}
                                    if ((float64(3.14159) < ((1e10 * 0.0) + 0.1)) && v48) {
                                        fmt.Println(v60, v35)
                                        var v62 uint = 18446744073709551615
                                        _ = v62
                                    } else {
                                        v63 := GenS9{F0: fi47}
                                        v64 := GenS2{F0: -9223372036854775808}
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = v60
                                    _ = v61
                                }
                                v35 = v35
                            } else {
                                v65 := GenS1{F0: (^(v35)), F1: 348, F2: 92}
                                v48 = v48
                                v35 -= (v35 * -1)
                                _ = v65
                            }
                            _ = v59
                        }
                        v35 = v35
                        if ((354 == 846) || v48) {
                            switch -9223372036854775808 {
                            case 32768:
                                switch 711 {
                                case -171:
                                    var v66 uint16 = 65535
                                    _ = v66
                                case 9223372036854775807:
                                    var v67 int32 = 1
                                    _ = v67
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v35, v35, fi49)
                                    }
                                }
                            case 0:
                                switch "a" {
                                case "hello":
                                    v68 := GenS9{F0: v1}
                                    v69 := []int{(fi49 >> ((2) & 63)), (v1 & fi47), 128, (v1 * 1)}
                                    var v70 bool = (!v48)
                                    var v71 uint32 = 712
                                    v70 = (42 >= 70)
                                    if v70 {
                                        if v70 {
                                            var v72 uint8 = 0
                                            v69 = append(v69, fi49)
                                            v48 = (!(!true))
                                            _ = v72
                                        }
                                    } else {
                                        v73 := GenS3{F0: 2}
                                        if (true || (((-2.25 / -2.25) / 0.1) <= float64(0.0))) {
                                            v74 := GenS8{F0: (v1 | fi49), F1: 0}
                                            _ = v74
                                        }
                                        _ = v73
                                    }
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        if ((fi49 << ((0) & 63)) >= v1) {
                                            v75 := []int{}
                                            _ = v75
                                        }
                                    case "hello":
                                        var v76 float32 = float32(((0.1 + 1e10) / (-2.25 / float64(1.5))))
                                        _ = v76
                                    default:
                                        switch (("" + (((("a" + "a") + ("a" + ((("\n\t" + "hello") + ("" + "")) + ("hello" + ("" + "a"))))) + "\n\t") + "hello")) + "a") {
                                        case "hello":
                                            v35 -= (v35 + v35)
                                            if v48 {
                                                if (4294967295 <= 18446744073709551615) {
                                                    var v77 bool = (!v48)
                                                    v48 = (!(false || v48))
                                                    v78 := GenS3{F0: 439}
                                                    var v79 uint16 = 65535
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(fi47, v48, v1)
                                        }
                                        if ((127 != 100) || true) {
                                            var v80 int32 = 0
                                            v81 := GenS10{F0: 358, F1: -128, F2: 0}
                                            if (!false) {
                                                fmt.Println(fi47)
                                                var v82 uint = 255
                                                _ = v82
                                            }
                                            var v83 uint = 18446744073709551615
                                            var v84 uint16 = 32767
                                            fmt.Println(v35, v48, v83)
                                            _ = v80
                                            _ = v81
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                }
                            default:
                                var v85 int8 = 59
                                _ = v85
                            }
                            v86 := v2[586]
                            v87 := GenS11{F0: 595}
                            _ = v86
                            _ = v87
                        }
                    } else {
                        var v88 int64 = -403
                        var v89 uint8 = 42
                        if v48 {
                            fmt.Println(fi47, v89, v48)
                        }
                        _ = v88
                        _ = v89
                    }
                }
                _ = fi49
            }
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
}
