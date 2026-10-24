package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int; F2 int64 }
type GenS9 struct { F0 uint; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 int32; F1 int32; F2 uint }
type GenS12 struct { F0 int; F1 bool; F2 int16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int; F1 int64; F2 int32 }
type GenS15 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint8, p1 string) int32 {
    p1 = "a"
    if true {
        v4 := GenS1{F0: -124, F1: 1, F2: 32767}
        v5 := []int{74, -918, -1, -9223372036854775808}
        _ = v4
        _ = v5
    } else {
        p0 = (p0 % ((p0) | 1))
        for fi6 := 0; fi6 < 13; fi6++ {
            if true {
                p1 = ""
            } else {
                switch (p1 + ("foo bar" + p1)) {
                case "a":
                    switch 32768 {
                    case 128:
                        p0 = p0
                    case -1:
                        if (42 < 657) {
                            p1 = p1
                            if false {
                                v7 := []int{fi6}
                                for fi8 := 0; fi8 < 16; fi8++ {
                                    if true {
                                        var v9 uint16 = 42
                                        _ = v9
                                    } else {
                                        p0 = (p0 & p0)
                                    }
                                    fmt.Println(p0, fi8, fi8)
                                    _ = fi8
                                }
                                _ = v7
                            } else {
                                if true {
                                    for fi10 := 0; fi10 < 16; fi10++ {
                                        var v11 int8 = 127
                                        if false {
                                            fmt.Println(p0, fi10)
                                            var v12 bool = false
                                            _ = v12
                                        } else {
                                            var v13 string = p1
                                            v14 := []int{(fi6 - fi10), fi6, -2147483648}
                                            _ = v13
                                            _ = v14
                                        }
                                        _ = fi10
                                        _ = v11
                                    }
                                    switch 128 {
                                    case 4294967295:
                                        var v15 uint = 24
                                        _ = v15
                                    case 0:
                                        var v16 int16 = 1000
                                        _ = v16
                                    case -9223372036854775808:
                                        var v17 int16 = 32767
                                        _ = v17
                                    }
                                } else {
                                    var v18 string = ("\n\t" + "\n\t")
                                    v19 := GenS1{F0: -1, F1: 2, F2: 1}
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                    default:
                        v20 := []int{(fi6 | fi6), 128}
                        _ = v20
                    }
                case "":
                    if false {
                        var v21 uint16 = 0
                        for fi22 := 0; fi22 < 11; fi22++ {
                            p0 = p0
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 20; fi23++ {
                            switch ("a" + p1) {
                            case "\n\t":
                                v24 := GenS2{F0: (215 < (p0 | (^(p0)))), F1: (v21 & (v21 * (^(v21)))), F2: 952}
                                _ = v24
                            default:
                                v25 := GenS3{F0: v21}
                                _ = v25
                            }
                            switch p1 {
                            case "a":
                                v21 += (^(v21))
                            }
                            _ = fi23
                        }
                        _ = v21
                    }
                }
                var v26 int = (fi6 * 42)
                p0 = (p0 % (((p0 ^ p0)) | 1))
                _ = v26
            }
            _ = fi6
        }
    }
    if true {
        var v27 string = ("hello" + "foo bar")
        var v28 int8 = -102
        _ = v27
        _ = v28
    } else {
        v29 := []int{-537, 1, 907, -18}
        switch ("foo bar" + p1) {
        case "":
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        case "foo bar":
            v30 := GenS4{F0: 1}
            _ = v30
        default:
            v31 := GenS3{F0: 65535}
            _ = v31
        }
        _ = v29
    }
    var v32 int32 = -1
    _ = p0
    _ = p1
    _ = v32
    return v32
}

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 int8 = 100
    for fi33 := 0; fi33 < 6; fi33++ {
        if (!false) {
            switch (fi33 ^ fi33) {
            case 2147483647:
                var v34 bool = (!true)
                _ = v34
            case -2147483648:
                v3 += v3
            }
            v35 := GenS4{F0: 9223372036854775807}
            _ = v35
        }
        var v36 int = (-(fi33))
        var v37 float32 = 0.1
        _ = fi33
        _ = v36
        _ = v37
    }
    var v38 uint8 = 42
    for fi39 := 0; fi39 < 11; fi39++ {
        if (917 != -32768) {
            fmt.Println(v3, fi39)
            fmt.Println(fi39, v3)
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        } else {
            v38 -= 128
            if (false && false) {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                v38 = v38
                v40 := GenS4{F0: 42}
                _ = v40
            }
        }
        fmt.Println(fi39, v3)
        if (-9223372036854775808 <= -1) {
            switch "foo bar" {
            case "foo bar":
                v38 = v38
            case "hello":
                v41 := GenS5{F0: 928, F1: 808}
                _ = v41
            default:
                fmt.Println(v38, v3)
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                v43 := GenS6{F0: "\n\t"}
                _ = fi42
                _ = v43
            }
        } else {
            fmt.Println(v3, fi39)
            var v44 int = fi39
            _ = v44
        }
        _ = fi39
    }
    var v45 int8 = v3
    v46 := GenS7{F0: 525}
    if ((uint64(v38) < 240) || true) {
        for fi47 := 0; fi47 < 17; fi47++ {
            var v48 uint8 = (^(v38))
            _ = fi47
            _ = v48
        }
        fmt.Println(v38, v3, v38)
        if true {
            for fi49 := 0; fi49 < 19; fi49++ {
                v45 = (v3 & v3)
                var v50 int = (-(fi49))
                v50 = -2147483648
                _ = fi49
                _ = v50
            }
        }
    } else {
        if true {
            var v51 uint8 = 128
            switch 4294967295 {
            case -880:
                for fi52 := 0; fi52 < 12; fi52++ {
                    delete(v2, "\n\t")
                    if false {
                        v1 = append(v1, "a")
                        if true {
                            var v53 bool = ((v3 > (v3 >> ((255) & 63))) || false)
                            for fi54 := 0; fi54 < 11; fi54++ {
                                if v53 {
                                    v38 = (v38 + (v38 & (v38 - v38)))
                                }
                                fmt.Println(fi52, v38)
                                _ = fi54
                            }
                            v53 = v53
                            _ = v53
                        } else {
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v56 := GenS7{F0: 1}
                                _ = fi55
                                _ = v56
                            }
                            v57 := GenS2{F0: true, F1: 851, F2: 442}
                            _ = v57
                        }
                    }
                    _ = fi52
                }
            default:
                if (!false) {
                    v58 := GenS8{F0: -2147483648, F1: 42, F2: 691}
                    if (1.5 <= -2.25) {
                        v1 = append(v1, ("a" + "a"))
                        fmt.Println(v3)
                        var v59 uint64 = 1
                        _ = v59
                    }
                    var v60 int32 = 1
                    _ = v58
                    _ = v60
                } else {
                    var v61 int = -2147483648
                    var v62 int16 = 42
                    _ = v61
                    _ = v62
                }
            }
            _ = v51
        }
        var v63 int64 = -867
        v2[("foo bar" + "")] = (("hello" + ("hello" + "")) + ("\n\t" + "hello"))
        _ = v63
    }
    _ = len(v1)
    for fi64 := 0; fi64 < 13; fi64++ {
        switch fi64 {
        case 128:
            if (false && (!((-(v45)) != v45))) {
                var v65 int64 = 782
                var v66 int = 32767
                switch (fi64 / ((v66) | 1)) {
                case 854:
                    if (18446744073709551615 != 18446744073709551615) {
                        var v67 int = (fi64 >> ((18446744073709551615) & 63))
                        v38 = (v38 ^ (v38 | 125))
                        var v68 int32 = -2147483648
                        _ = v67
                        _ = v68
                    }
                case 128:
                    var v69 float32 = -2.25
                    _ = v69
                case 0:
                    var v70 int32 = -2147483648
                    _ = v70
                default:
                    if (!true) {
                        var v71 int16 = -32768
                        for fi72 := 0; fi72 < 16; fi72++ {
                            var v73 int64 = 0
                            fmt.Println(v66)
                            _ = fi72
                            _ = v73
                        }
                        _ = v71
                    }
                }
                _ = v65
                _ = v66
            }
        case -1:
            switch 32768 {
            case -2147483648:
                switch ("hello" + "a") {
                case "a":
                    if true {
                        for fi74 := 0; fi74 < 11; fi74++ {
                            v75 := v2[("hello" + "a")]
                            v76 := GenS9{F0: 2, F1: -2147483648, F2: 169}
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        if (((1e10 - 0.0) > 0.0) && (((0.0 + float64(float32(0.0))) / 0.1) > (3.14159 / 1e10))) {
                            if false {
                                for fi77 := 0; fi77 < 19; fi77++ {
                                    if ("\n\t" <= "\n\t") {
                                        if true {
                                            if (1000 == 32767) {
                                                v45 -= (v3 / (((v3 % ((16) | 1))) | 1))
                                                v78 := GenS3{F0: uint16(v38)}
                                                var v79 int64 = -541
                                                _ = v78
                                                _ = v79
                                            } else {
                                                var v80 int = fi64
                                                v38 += (v38 - v38)
                                                if false {
                                                    if false {
                                                        var v81 uint32 = 4294967295
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v82 float64 = ((1.5 + (-2.25 - (float64(0.1) * 0.0))) - 0.0)
                                                        _ = v81
                                                        _ = v82
                                                    } else {
                                                        var v83 int8 = -1
                                                        v84 := v2["hello"]
                                                        if false {
                                                            v84 = v84
                                                        } else {
                                                            fmt.Println(v38, v84, v3)
                                                            v2[v84] = (v84 + (v84 + ((v84 + v84) + "foo bar")))
                                                        }
                                                        var v85 int64 = 1
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    if ((fi64 + fi64) < 2147483647) {
                                                        if (false || (!(false || false))) {
                                                            v45 -= (v45 | (v3 % ((v3) | 1)))
                                                        }
                                                    }
                                                    v3 = v3
                                                    if (true && true) {
                                                        if (0 == 9223372036854775807) {
                                                            var v86 uint16 = 255
                                                            _ = v86
                                                        } else {
                                                            switch ("a" + "foo bar") {
                                                            case "\n\t":
                                                                v38 += 128
                                                            case "hello":
                                                                v80 -= (fi64 << ((4294967295) & 63))
                                                            case "":
                                                                if false {
                                                                    switch -904 {
                                                                    case 128:
                                                                        v87 := GenS10{F0: 1, F1: 18446744073709551615}
                                                                        _ = v87
                                                                    }
                                                                    var v88 int16 = 1000
                                                                    v89 := GenS11{F0: -1, F1: 1, F2: 42}
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                            default:
                                                                v45 = (v3 * v45)
                                                                if ((v3 << ((65535) & 63)) <= (v45 % ((42) | 1))) {
                                                                    v3 -= v45
                                                                    v2[(("" + "hello") + ("\n\t" + "a"))] = ""
                                                                    v38 -= 251
                                                                }
                                                            }
                                                            var v90 int64 = 283
                                                            var v91 int = (fi64 | 128)
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        if ((v3 << ((471) & 63)) < -128) {
                                                            if false {
                                                                var v92 uint32 = 0
                                                                switch ("hello" + "foo bar") {
                                                                case "\n\t":
                                                                    fmt.Println(v92, v3)
                                                                }
                                                                _ = v92
                                                            }
                                                            v80 = 2147483647
                                                        }
                                                        v93 := GenS12{F0: 2147483647, F1: (!(1e10 == (0.1 + float32(0.0)))), F2: -1}
                                                        var v94 string = "hello"
                                                        if false {
                                                            if (true || false) {
                                                                delete(v2, v94)
                                                                v2[v94] = "a"
                                                                fmt.Println(fi64, v80)
                                                                v80 = fi64
                                                            }
                                                            if true {
                                                                v94 = "a"
                                                                fmt.Println(v94, v38)
                                                            }
                                                        }
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        var v95 uint64 = 2
                                                        v2["\n\t"] = (("hello" + "foo bar") + "")
                                                        _ = v95
                                                    }
                                                } else {
                                                    v45 = (v45 * int8(v38))
                                                }
                                                _ = v80
                                            }
                                        }
                                        v45 = (-(v45))
                                        if false {
                                            v45 = v3
                                        }
                                    } else {
                                        var v96 bool = ((("" + "hello") + ((("foo bar" + ("hello" + ("a" + "foo bar"))) + "\n\t") + "hello")) > "\n\t")
                                        var v97 uint32 = 42
                                        var v98 float64 = 0.0
                                        var v99 uint = 351
                                        v98 -= (((1.5 - v98) + v98) - (float64(-2.25) + v98))
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    v100 := GenS13{F0: "\n\t"}
                                    v45 = v3
                                    switch "hello" {
                                    case "foo bar":
                                        v38 = 42
                                        fmt.Println(v45, v45, fi64)
                                    case "a":
                                        v101 := GenS14{F0: fi77, F1: 1, F2: 2147483647}
                                        _ = v101
                                    }
                                    _ = fi77
                                    _ = v100
                                }
                                v102 := GenS8{F0: (fi64 % ((fi64) | 1)), F1: fi64, F2: int64(v45)}
                                _ = v102
                            } else {
                                fmt.Println(v3, fi64)
                            }
                        } else {
                            v103 := GenS15{F0: 18446744073709551615, F1: -339}
                            for fi104 := 0; fi104 < 14; fi104++ {
                                var v105 int = (^(fi104))
                                _ = fi104
                                _ = v105
                            }
                            var v106 int8 = 100
                            _ = v103
                            _ = v106
                        }
                        v3 = (v3 | (-(v3)))
                    }
                case "\n\t":
                    var v107 int = 320
                    _ = v107
                }
            }
        case 42:
            var v108 uint32 = 65535
            _ = v108
        }
        var v109 int64 = 1
        _ = fi64
        _ = v109
    }
    v110 := v2["\n\t"]
    v111 := v2[""]
    for fi112 := 0; fi112 < 11; fi112++ {
        var v113 uint16 = 914
        fmt.Println(v110, v38)
        if (!(!(true && (1 == 42)))) {
            v113 -= v113
            v45 = (v45 % (((v3 & 0)) | 1))
        } else {
            for fi114 := 0; fi114 < 2; fi114++ {
                v113 = (v113 & 0)
                v45 -= (v3 >> ((uint(v38)) & 63))
                var v115 int16 = 0
                if (0 < 42) {
                    v115 = -1
                    v45 = (v45 * v45)
                }
                delete(v2, v111)
                v38 = v38
                v45 += v45
                _ = fi114
                _ = v115
            }
            for fi116 := 0; fi116 < 4; fi116++ {
                v111 = v110
                _ = fi116
            }
        }
        _ = fi112
        _ = v113
    }
    v2["hello"] = (v110 + v111)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v38
    _ = v45
    _ = v46
}
