package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 uint16; F2 bool }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 uint8 }

func main() {
    var v1 bool = ((!(!(true && false))) || false)
    var v2 int = 32767
    var v3 uint8 = 255
    v4 := []string{("hello" + ""), (("\n\t" + "") + "hello"), "foo bar", "a"}
    var v5 string = "a"
    switch 32767 {
    case 32768:
        v1 = v1
    case -2147483648:
        if false {
            if ((!v1) && ((!v1) && true)) {
                var v6 int16 = -32768
                v3 = v3
                _ = v6
            }
        }
    case 1:
        if v1 {
            if len(v4) > 0 {
                v4[0] = (((v5 + "") + "a") + ("foo bar" + ("a" + v5)))
            }
            if true {
                var v7 float32 = 0.1
                var v8 uint16 = 735
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 2, F1: (v3 ^ 126), F2: v3}
            _ = v9
        }
    default:
        v10 := GenS2{F0: 149, F1: true}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v2 -= fi11
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int = (fi12 / (((fi12 >> ((42) & 63))) | 1))
            var v14 bool = v1
            var v15 uint8 = (^(v3))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    switch v2 {
    case 127:
        fmt.Println(v3)
    case 4294967295:
        v16 := GenS3{F0: -2147483648, F1: ("foo bar" + v5), F2: (v3 ^ 182)}
        _ = v16
    case 563:
        fmt.Println(v2, v1, v2)
    default:
        if ((1.5 == (1.5 / float64(0.1))) && (!v1)) {
            v17 := GenS4{F0: -975, F1: 1}
            _ = v17
        } else {
            if (v2 > v2) {
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch fi18 {
                    case -1:
                        fmt.Println(v1)
                    case 252:
                        var v19 uint = 65535
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 string = "\n\t"
                        var v22 uint16 = 255
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi18
                }
                switch v2 {
                case 32767:
                    var v23 uint = 984
                    _ = v23
                case -645:
                    var v24 bool = false
                    _ = v24
                case 42:
                    var v25 uint16 = 127
                    _ = v25
                default:
                    v5 = (v5 + "foo bar")
                }
            } else {
                v5 = v5
                for fi26 := 0; fi26 < 2; fi26++ {
                    for fi27 := 0; fi27 < 19; fi27++ {
                        fmt.Println(v5)
                        _ = fi27
                    }
                    _ = fi26
                }
            }
            var v28 uint16 = 0
            _ = v28
        }
    }
    if (65535 != 0) {
        fmt.Println(v5, v3, v5)
    } else {
        switch "a" {
        case "foo bar":
            v3 = (v3 * v3)
        case "":
            v2 = (v2 << ((42) & 63))
        case "\n\t":
            var v29 int8 = 100
            _ = v29
        }
    }
    fmt.Println(v2)
    v3 -= 81
    v1 = v1
    for fi30 := 0; fi30 < 3; fi30++ {
        v5 = (("foo bar" + (v5 + "")) + (v5 + (v5 + ("a" + v5))))
        fmt.Println(fi30, fi30, v3)
        _ = fi30
    }
    v2 += v2
    if (v5 <= v5) {
        if true {
            switch v5 {
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "a":
                switch "foo bar" {
                case "foo bar":
                    _ = len(v4)
                case "\n\t":
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v32 := GenS5{F0: 123}
                        _ = fi31
                        _ = v32
                    }
                case "":
                    for fi33 := 0; fi33 < 10; fi33++ {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v2 = fi34
                            _ = fi34
                        }
                        if v1 {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v35 uint32 = 65535
                            switch (v5 + (v5 + v5)) {
                            case "a":
                                var v36 int64 = 9223372036854775807
                                _ = v36
                            default:
                                v2 = (v2 ^ -973)
                            }
                            _ = v35
                        } else {
                            var v37 uint32 = 0
                            _ = v37
                        }
                        v3 = v3
                        _ = fi33
                    }
                }
            case "":
                switch v2 {
                case 0:
                    switch v5 {
                    case "foo bar":
                        v38 := GenS4{F0: -32768, F1: -31}
                        _ = v38
                    default:
                        var v39 bool = v1
                        _ = v39
                    }
                case 32768:
                    for fi40 := 0; fi40 < 6; fi40++ {
                        var v41 float64 = -2.25
                        _ = fi40
                        _ = v41
                    }
                }
            }
        } else {
            for fi42 := 0; fi42 < 20; fi42++ {
                var v43 uint = 42
                switch (v2 / ((-437) | 1)) {
                case 456:
                    var v44 int = fi42
                    _ = v44
                default:
                    if false {
                        v45 := GenS4{F0: -32768, F1: 0}
                        v46 := GenS2{F0: 42, F1: (!(false || false))}
                        _ = v45
                        _ = v46
                    } else {
                        var v47 float64 = 1e10
                        if (!v1) {
                            for fi48 := 0; fi48 < 11; fi48++ {
                                var v49 bool = v1
                                _ = fi48
                                _ = v49
                            }
                            if (true && (0.0 < float32(v47))) {
                                v3 = 127
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch v5 {
                                case "a":
                                    v50 := GenS4{F0: 42, F1: -1}
                                    _ = v50
                                case "":
                                    v4 = append(v4, (v5 + v5))
                                default:
                                    var v51 uint = v43
                                    _ = v51
                                }
                                v2 = fi42
                            }
                            for fi52 := 0; fi52 < 6; fi52++ {
                                fmt.Println(fi42, v47, v1)
                                v5 = "\n\t"
                                v4 = append(v4, v5)
                                fmt.Println(v47)
                                if len(v4) > 0 {
                                    v4[0] = (v5 + v5)
                                }
                                _ = fi52
                            }
                        } else {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                if (255 <= uint8(v43)) {
                                    switch ("" + "hello") {
                                    case "hello":
                                        var v54 int16 = 271
                                        _ = v54
                                    case "foo bar":
                                        var v55 float32 = 1e10
                                        var v56 string = "a"
                                        _ = v55
                                        _ = v56
                                    default:
                                        if (!(1 < 42)) {
                                            fmt.Println(v1, v5, v2)
                                            if ((-32768 > 1) && (!((v3 % ((42) | 1)) > v3))) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v3 -= (v3 % ((v3) | 1))
                                                if (v1 && true) {
                                                    if v1 {
                                                        switch (fi53 >> ((18446744073709551615) & 63)) {
                                                        case -461:
                                                            var v57 uint64 = 54
                                                            _ = v57
                                                        }
                                                        var v58 uint8 = 147
                                                        _ = v58
                                                    } else {
                                                        if v1 {
                                                            if v1 {
                                                                v4 = append(v4, ((v5 + v5) + v5))
                                                            } else {
                                                                var v59 int8 = -1
                                                                _ = v59
                                                            }
                                                        } else {
                                                            v60 := GenS3{F0: -2147483648, F1: ("" + ("" + v5)), F2: v3}
                                                            v61 := GenS1{F0: 311, F1: (^(v3)), F2: v3}
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        var v62 int = (fi53 >> ((v43) & 63))
                                                        _ = v62
                                                    }
                                                    if false {
                                                        v63 := GenS6{F0: 1, F1: 32767, F2: (!(!true))}
                                                        _ = v63
                                                    } else {
                                                        v1 = true
                                                        v2 += (v2 & 2147483647)
                                                        fmt.Println(v2, v2, v47)
                                                    }
                                                }
                                            } else {
                                                v64 := GenS1{F0: 724, F1: (v3 * v3), F2: (v3 & v3)}
                                                var v65 int64 = 722
                                                _ = v64
                                                _ = v65
                                            }
                                        } else {
                                            var v66 float64 = 3.14159
                                            _ = v66
                                        }
                                    }
                                    v67 := GenS6{F0: 1, F1: 32767, F2: v1}
                                    _ = v67
                                } else {
                                    if (v3 <= (v3 - (^(v3)))) {
                                        _ = len(v4)
                                        if (v1 && (v1 && (v1 && false))) {
                                            if v1 {
                                                v43 = uint(v3)
                                            } else {
                                                var v68 bool = ((1.5 - 0.0) >= float32(-2.25))
                                                var v69 uint64 = 4294967295
                                                _ = v68
                                                _ = v69
                                            }
                                            if (63 == 117) {
                                                v3 = 0
                                                v70 := GenS1{F0: 0, F1: (v3 >> ((v43) & 63)), F2: v3}
                                                _ = v70
                                            } else {
                                                v2 += fi53
                                            }
                                        } else {
                                            v71 := GenS5{F0: 2147483647}
                                            fmt.Println(fi53)
                                            _ = v71
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = v5
                                        }
                                    }
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v72 uint8 = (v3 % ((v3) | 1))
                                _ = fi53
                                _ = v72
                            }
                        }
                        fmt.Println(v5, v2, v47)
                        _ = v47
                    }
                }
                if v1 {
                    v43 = v43
                    if (true || false) {
                        if (1000 >= -1) {
                            var v73 uint32 = uint32(v43)
                            _ = v73
                        } else {
                            v74 := GenS7{F0: v1, F1: ((127 != 100) || true)}
                            fmt.Println(v1, v1)
                            v75 := GenS8{F0: (v2 & -1), F1: v3}
                            _ = v74
                            _ = v75
                        }
                    } else {
                        switch fi42 {
                        case 368:
                            fmt.Println(v5, v5, v2)
                        default:
                            if v1 {
                                v43 -= v43
                                v43 = (v43 % ((v43) | 1))
                            }
                        }
                    }
                }
                _ = fi42
                _ = v43
            }
            v76 := GenS9{F0: (v2 ^ int(v3)), F1: 987}
            _ = v76
        }
    }
    _ = len(v4)
    if v1 {
        var v77 int = (v2 % ((128) | 1))
        var v78 int64 = -9223372036854775808
        v1 = false
        _ = v77
        _ = v78
    } else {
        v1 = (1 > 101)
        var v79 int32 = 2147483647
        var v80 uint64 = 18446744073709551615
        _ = v79
        _ = v80
    }
    for fi81 := 0; fi81 < 6; fi81++ {
        v1 = true
        if (false || true) {
            var v82 uint = 255
            fmt.Println(v3)
            v4 = append(v4, v5)
            _ = v82
        } else {
            v5 = ("hello" + v5)
            if true {
                if v1 {
                    fmt.Println(v1)
                    switch v2 {
                    case -9223372036854775808:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 128:
                        var v83 int = fi81
                        _ = v83
                    default:
                        fmt.Println(v2, v5, v3)
                    }
                    v4 = append(v4, v5)
                } else {
                    _ = len(v4)
                    fmt.Println(v2, fi81, v2)
                }
            }
            var v84 bool = v1
            _ = v84
        }
        _ = fi81
    }
    if v1 {
        v85 := GenS10{F0: 0}
        if (9223372036854775807 <= 940) {
            var v86 uint16 = 740
            _ = v86
        } else {
            v87 := GenS3{F0: 1, F1: (v5 + v5), F2: v3}
            _ = v87
        }
        fmt.Println(v3, v3)
        _ = v85
    } else {
        v88 := GenS11{F0: 42, F1: (v3 * (^(v3)))}
        switch v5 {
        case "foo bar":
            _ = len(v4)
        }
        _ = v88
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
