package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 uint32; F1 int8; F2 int }
func genFn0(p0 int8, p1 int) string {
    if (0 > 278) {
        v5 := []int{}
        fmt.Println(p1)
        if ((false && false) && ((false || true) || (true || false))) {
            v6 := GenS1{F0: (p1 ^ p1), F1: 9223372036854775807, F2: 935}
            _ = v6
        } else {
            _ = len(v5)
            var v7 uint64 = 382
            _ = v7
        }
        _ = v5
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint16 = 122
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return (("foo bar" + "hello") + "a")
}
func genFn1(p0 uint64, p1 int, p2 int32) int {
    v10 := []int{}
    if false {
        p2 += p2
    } else {
        if (p1 != p1) {
            if (127 > 152) {
                p2 -= p2
            } else {
                if (255 < 127) {
                    switch "\n\t" {
                    case "\n\t":
                        if ((p0 % ((18446744073709551615) | 1)) < p0) {
                            p2 = p2
                            p0 = 0
                            switch (p1 ^ 1) {
                            case 2147483647:
                                switch p1 {
                                case -2147483648:
                                    if true {
                                        var v11 int32 = p2
                                        var v12 float64 = 0.1
                                        _ = v11
                                        _ = v12
                                    } else {
                                        var v13 uint64 = (p0 + 42)
                                        if false {
                                            v14 := GenS2{F0: uint16(p0)}
                                            _ = v14
                                        }
                                        v15 := GenS1{F0: p1, F1: 0, F2: 2}
                                        _ = v13
                                        _ = v15
                                    }
                                case 9223372036854775807:
                                    if true {
                                        var v16 int64 = 609
                                        var v17 int32 = p2
                                        var v18 uint16 = 255
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            default:
                                fmt.Println(p2, p0, p2)
                            }
                        }
                    default:
                        var v19 int = (p1 + 0)
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 18; fi20++ {
                        p0 += 18446744073709551615
                        var v21 bool = (false && false)
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 uint64 = p0
                _ = v22
            }
            if len(v10) > 0 {
                v10[0] = (p1 - 128)
            }
            var v23 uint16 = 0
            _ = v23
        } else {
            p2 = p2
            var v24 uint32 = 65535
            _ = v24
        }
        if ((-2.25 * (3.14159 + ((3.14159 - float64(-2.25)) + 1e10))) >= 1e10) {
            if ((1.5 < 3.14159) && (((!true) || (!false)) && (!false))) {
                for fi25 := 0; fi25 < 19; fi25++ {
                    fmt.Println(fi25, fi25, p0)
                    v26 := GenS3{F0: -445}
                    _ = fi25
                    _ = v26
                }
                for fi27 := 0; fi27 < 20; fi27++ {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = len(v10)
                        _ = fi28
                    }
                    _ = fi27
                }
                switch 0 {
                case 2147483647:
                    var v29 float32 = 1e10
                    _ = v29
                case 0:
                    v30 := GenS4{F0: -55}
                    _ = v30
                case -9223372036854775808:
                    var v31 int32 = 0
                    _ = v31
                default:
                    if (1.5 == (3.14159 * (3.14159 + 1e10))) {
                        p0 += (p0 + p0)
                        p1 -= p1
                        switch ("a" + ((("" + "hello") + "foo bar") + "hello")) {
                        case "a":
                            for fi32 := 0; fi32 < 13; fi32++ {
                                var v33 uint16 = 2
                                _ = fi32
                                _ = v33
                            }
                        default:
                            if len(v10) > 0 {
                                v10[0] = 4294967295
                            }
                        }
                    }
                }
            } else {
                for fi34 := 0; fi34 < 16; fi34++ {
                    v35 := GenS2{F0: 32767}
                    v36 := GenS4{F0: 42}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if len(v10) > 0 {
                    v10[0] = (p1 * (p1 << ((255) & 63)))
                }
            }
            p0 = (p0 >> ((1) & 63))
            v10 = append(v10, p1)
        } else {
            var v37 uint32 = 907
            v37 = 2147483647
            var v38 float64 = 1.5
            _ = v37
            _ = v38
        }
    }
    var v39 uint16 = 775
    var v40 int64 = 42
    var v41 int32 = -311
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v39
    _ = v40
    _ = v41
    return p1
}

func main() {
    var v1 string = ("hello" + "")
    v2 := []int{402, 1, 892}
    var v3 uint8 = 0
    var v4 bool = false
    for fi42 := 0; fi42 < 18; fi42++ {
        if true {
            v4 = v4
            var v43 uint32 = 42
            var v44 int32 = 0
            _ = v43
            _ = v44
        }
        switch (v1 + "a") {
        case "foo bar":
            var v45 int16 = 1000
            _ = v45
        case "a":
            _ = len(v2)
        case "":
            v1 = v1
        default:
            var v46 float64 = 1.5
            _ = v46
        }
        _ = fi42
    }
    if (1e10 <= float64(1e10)) {
        if len(v2) > 0 {
            v2[0] = 32768
        }
        var v47 uint32 = 2
        _ = v47
    }
    var v48 uint32 = 42
    fmt.Println(v4, v48)
    if ((float64(((3.14159 + (float32(3.14159) + 1.5)) * (0.1 / float32(1.5)))) > (float64(3.14159) / 3.14159)) && (!(!(!true)))) {
        if (v4 || (((-2.25 * 1e10) + 0.0) != float64((0.1 * 1.5)))) {
            fmt.Println(v48, v48, v3)
            fmt.Println(v48, v48, v4)
            v49 := GenS5{F0: 240, F1: 33}
            _ = v49
        } else {
            switch -381 {
            case 2147483647:
                fmt.Println(v1, v48, v3)
            case 128:
                fmt.Println(v4)
            case -689:
                v48 = 0
            }
            for fi50 := 0; fi50 < 1; fi50++ {
                v3 = v3
                var v51 bool = v4
                _ = fi50
                _ = v51
            }
            v52 := GenS6{F0: v4, F1: 987}
            _ = v52
        }
        v53 := GenS4{F0: -729}
        if (255 <= 458) {
            var v54 string = ""
            var v55 uint16 = 255
            _ = v54
            _ = v55
        }
        _ = v53
    } else {
        v56 := GenS7{F0: 987, F1: 1}
        var v57 int64 = 42
        _ = v56
        _ = v57
    }
    var v58 float64 = 1e10
    switch -115 {
    case 1:
        var v59 int16 = 42
        _ = v59
    case -1:
        v58 += v58
    default:
        if (((1e10 * (v58 - 3.14159)) + v58) > v58) {
            switch "hello" {
            case "":
                if (-178 > 127) {
                    var v60 int = 32768
                    if (!(100 > -128)) {
                        var v61 int = v60
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 19; fi62++ {
                        if true {
                            var v63 bool = true
                            if len(v2) > 0 {
                                v2[0] = (fi62 << ((1) & 63))
                            }
                            var v64 int16 = 1
                            _ = v63
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 4; fi65++ {
                            var v66 float64 = (float64(-2.25) - v58)
                            v3 += (v3 % (((^(v3))) | 1))
                            _ = fi65
                            _ = v66
                        }
                        _ = fi62
                    }
                    _ = v60
                }
            case "foo bar":
                if (true && true) {
                    for fi67 := 0; fi67 < 14; fi67++ {
                        switch (-(fi67)) {
                        case 127:
                            fmt.Println(fi67)
                        default:
                            var v68 string = (("foo bar" + "foo bar") + v1)
                            _ = v68
                        }
                        _ = len(v2)
                        var v69 int32 = -2147483648
                        _ = fi67
                        _ = v69
                    }
                } else {
                    switch 40 {
                    case 348:
                        v70 := GenS8{F0: 2, F1: ((((((v1 + v1) + "") + v1) + (v1 + "hello")) + ("\n\t" + "foo bar")) == v1), F2: (v48 << ((42) & 63))}
                        _ = v70
                    }
                }
            case "hello":
                v4 = (!(v4 || v4))
            }
            if v4 {
                v71 := GenS7{F0: 42, F1: 750}
                switch genFn1(887, -2147483648, 671) {
                case 377:
                    v58 = (((v58 * (v58 - 3.14159)) / -2.25) - ((v58 / v58) + (v58 + v58)))
                case -794:
                    var v72 int64 = -1
                    _ = v72
                case 9223372036854775807:
                    v73 := GenS8{F0: 117, F1: false, F2: v48}
                    _ = v73
                default:
                    v74 := GenS9{F0: 0, F1: 127, F2: 1}
                    _ = v74
                }
                var v75 float64 = ((0.0 + v58) + v58)
                _ = v71
                _ = v75
            }
            if v4 {
                fmt.Println(v58)
                var v76 bool = (0 >= -1)
                v76 = (!false)
                _ = v76
            }
        }
    }
    v3 -= (v3 % (((v3 / (((v3 / (((v3 - v3)) | 1))) | 1))) | 1))
    if (-9223372036854775808 != 42) {
        switch ("\n\t" + (v1 + ((v1 + v1) + (("\n\t" + "a") + (v1 + v1))))) {
        case "foo bar":
            v2 = append(v2, -9223372036854775808)
        case "hello":
            for fi77 := 0; fi77 < 16; fi77++ {
                v48 = v48
                for fi78 := 0; fi78 < 13; fi78++ {
                    var v79 int8 = 1
                    var v80 uint = 4294967295
                    var v81 uint8 = 42
                    _ = fi78
                    _ = v79
                    _ = v80
                    _ = v81
                }
                var v82 int32 = 1
                _ = fi77
                _ = v82
            }
        case "a":
            switch 779 {
            case 42:
                for fi83 := 0; fi83 < 9; fi83++ {
                    for fi84 := 0; fi84 < 17; fi84++ {
                        fmt.Println(v58)
                        var v85 uint = 65535
                        fmt.Println(v4, v3)
                        _ = fi84
                        _ = v85
                    }
                    _ = fi83
                }
            case 1:
                var v86 int16 = 0
                _ = v86
            default:
                if v4 {
                    v87 := GenS10{F0: 1, F1: 0}
                    v48 = (v48 / (((^(v48))) | 1))
                    _ = v87
                } else {
                    v48 += (^(v48))
                }
            }
        default:
            if (!v4) {
                var v88 int16 = 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v3, v88, v48)
                _ = v88
            }
        }
        var v89 uint = 42
        _ = v89
    }
    var v90 string = v1
    v91 := GenS11{F0: (v48 + uint32(v3)), F1: 1, F2: genFn1(0, 2147483647, -580)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v58:", v58)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v58
    _ = v90
    _ = v91
}
