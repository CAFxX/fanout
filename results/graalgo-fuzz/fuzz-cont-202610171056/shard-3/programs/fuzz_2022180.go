package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int8; F2 int16 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = 32768
    var v3 int64 = -9223372036854775808
    switch 128 {
    case -1:
        var v4 int16 = -32768
        _ = v4
    default:
        fmt.Println(v2, v3)
    }
    var v5 int64 = 412
    var v6 uint16 = 32767
    switch (v2 * -2147483648) {
    case 0:
        switch "" {
        case "hello":
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := []int{v2}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v9 int64 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
        case "\n\t":
            var v10 int8 = 127
            _ = v10
        default:
            v2 = 127
        }
    case -2147483648:
        v11 := []int{(v2 << ((2) & 63)), int(v6)}
        _ = v11
    default:
        v12 := []int{(^(v2)), v2}
        _ = v12
    }
    if (42 == -827) {
        v13 := []int{(v2 % ((v2) | 1)), v2}
        v3 -= v3
        var v14 string = ("a" + "\n\t")
        _ = v13
        _ = v14
    }
    v15 := []int{(v2 ^ (v2 % (((v2 - (v2 / ((v2) | 1)))) | 1))), v2}
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: 4294967295, F1: "\n\t"}
        var v18 bool = (false && (false && false))
        switch (-(fi16)) {
        case 2147483647:
            v15 = append(v15, (fi16 >> ((v1) & 63)))
        case 32768:
            var v19 int = (^(fi16))
            _ = v19
        case -2147483648:
            for fi20 := 0; fi20 < 1; fi20++ {
                if (!true) {
                    var v21 int = (v2 | v2)
                    fmt.Println(v2, v6, v6)
                    var v22 uint32 = 1
                    _ = v21
                    _ = v22
                }
                v15 = append(v15, v2)
                var v23 uint16 = (v6 * v6)
                switch v2 {
                case 436:
                    if v18 {
                        var v24 uint64 = 2
                        if (-59 >= 42) {
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        if true {
                            if len(v15) > 0 {
                                v15[0] = fi16
                            }
                            if v18 {
                                _ = len(v15)
                                var v25 int8 = -128
                                v26 := GenS2{F0: -188, F1: (^(v6))}
                                _ = v25
                                _ = v26
                            } else {
                                v3 = (v5 | 0)
                                switch (-(fi16)) {
                                case 32767:
                                    v27 := GenS3{F0: 795, F1: 18446744073709551615, F2: v3}
                                    _ = v27
                                case -1:
                                    if ((v18 || ((fi16 > (-(v2))) && false)) || v18) {
                                        switch fi20 {
                                        case -908:
                                            var v28 uint8 = 32
                                            _ = v28
                                        case 34:
                                            v6 -= v6
                                        default:
                                            if (199 == 255) {
                                                if len(v15) > 0 {
                                                    v15[0] = (fi16 & 42)
                                                }
                                                fmt.Println(fi16, v5)
                                                v3 -= v3
                                            }
                                        }
                                        v2 -= v2
                                        var v29 uint = (^(v1))
                                        _ = v29
                                    } else {
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        var v30 bool = (!(42 == 679))
                                        v15 = append(v15, (v2 << ((0) & 63)))
                                        v23 -= v23
                                        _ = v30
                                    }
                                case 37:
                                    if false {
                                        var v31 bool = v18
                                        var v32 uint32 = 1
                                        var v33 float32 = 0.0
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            switch (fi20 << ((42) & 63)) {
                                            case 127:
                                                fmt.Println(v6, v5, v5)
                                            }
                                        case "":
                                            switch fi16 {
                                            case 404:
                                                if false {
                                                    switch (fi16 + int(v23)) {
                                                    case 47:
                                                        var v34 int16 = 42
                                                        _ = v34
                                                    case -597:
                                                        if false {
                                                            var v35 uint = (v1 + (^(v1)))
                                                            if v18 {
                                                                _ = len(v15)
                                                            }
                                                            if v18 {
                                                                v24 = (^(v24))
                                                                switch "" {
                                                                case "":
                                                                    switch fi20 {
                                                                    case 127:
                                                                        var v36 bool = ((^(v23)) >= v23)
                                                                        _ = v36
                                                                    }
                                                                case "hello":
                                                                    v37 := GenS2{F0: 0, F1: v23}
                                                                    _ = v37
                                                                }
                                                            }
                                                            _ = v35
                                                        }
                                                    default:
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                    }
                                                    var v38 string = ("hello" + "foo bar")
                                                    if (!(0.1 < 1e10)) {
                                                        var v39 uint64 = (v24 & v24)
                                                        if false {
                                                            var v40 float64 = (0.0 / 1.0)
                                                            _ = v40
                                                        } else {
                                                            v15 = append(v15, (fi16 + (fi20 ^ int(v6))))
                                                            fmt.Println(v18)
                                                        }
                                                        var v41 float64 = (1e10 - -2.25)
                                                        _ = v39
                                                        _ = v41
                                                    } else {
                                                        var v42 float64 = float64(1e10)
                                                        _ = len(v15)
                                                        _ = v42
                                                    }
                                                    _ = v38
                                                } else {
                                                    v2 = (v2 << ((255) & 63))
                                                    fmt.Println(v2)
                                                }
                                            case 0:
                                                v3 = v3
                                            }
                                        }
                                        v43 := GenS4{F0: 643}
                                        v3 = -820
                                        v44 := GenS5{F0: "a", F1: 63}
                                        if len(v15) > 0 {
                                            v15[0] = (v2 * -724)
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            }
                        }
                        _ = v24
                    }
                case 0:
                    fmt.Println(v2, v3)
                default:
                    if (-540 > 0) {
                        var v45 uint64 = 4294967295
                        var v46 uint64 = v45
                        var v47 int16 = 0
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                }
                _ = len(v15)
                switch (fi16 ^ (v2 ^ fi16)) {
                case 9223372036854775807:
                    _ = len(v15)
                default:
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v48 uint32 = 65535
                        _ = v48
                    default:
                        fmt.Println(fi16)
                    }
                }
                _ = fi20
                _ = v23
            }
        default:
            if v18 {
                var v49 float64 = ((0.0 * 0.0) / 0.1)
                v50 := GenS4{F0: -2147483648}
                v51 := GenS5{F0: (("\n\t" + ((("a" + "hello") + "foo bar") + "foo bar")) + ("" + ("" + "hello"))), F1: 1}
                _ = v49
                _ = v50
                _ = v51
            } else {
                for fi52 := 0; fi52 < 17; fi52++ {
                    switch "" {
                    case "a":
                        v53 := GenS2{F0: -32768, F1: (v6 ^ (v6 >> ((0) & 63)))}
                        _ = v53
                    case "foo bar":
                        v15 = append(v15, (^(fi52)))
                    default:
                        v54 := GenS4{F0: 0}
                        _ = v54
                    }
                    v55 := GenS4{F0: int32(v1)}
                    _ = fi52
                    _ = v55
                }
                v56 := GenS5{F0: ("hello" + ((("a" + ("\n\t" + "")) + "") + ((("" + "") + "a") + ("a" + "")))), F1: 0}
                for fi57 := 0; fi57 < 11; fi57++ {
                    v58 := GenS6{F0: 0, F1: 127, F2: -874}
                    _ = fi57
                    _ = v58
                }
                _ = v56
            }
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v59 uint16 = (v6 ^ (v6 + v6))
    if len(v15) > 0 {
        v15[0] = (v2 >> ((v1) & 63))
    }
    var v60 float32 = ((1e10 + 1.5) / 1.0)
    v6 -= (v59 ^ 0)
    v61 := GenS7{F0: 852, F1: 2147483647}
    var v62 uint64 = 98
    switch "a" {
    case "\n\t":
        var v63 uint32 = 814
        _ = v63
    default:
        var v64 float32 = 1e10
        _ = v64
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        v15 = append(v15, fi65)
        for fi66 := 0; fi66 < 18; fi66++ {
            switch "foo bar" {
            case "\n\t":
                v67 := GenS4{F0: -2147483648}
                _ = v67
            case "":
                if (55 > 706) {
                    if len(v15) > 0 {
                        v15[0] = (v2 % ((-858) | 1))
                    }
                    v3 = v3
                }
            case "a":
                v6 = v59
            default:
                var v68 bool = ("a" <= "foo bar")
                _ = v68
            }
            _ = fi66
        }
        _ = fi65
    }
    if true {
        if (false || true) {
            v3 -= (-(v5))
        }
    } else {
        var v69 int32 = 42
        _ = v69
    }
    for fi70 := 0; fi70 < 5; fi70++ {
        for fi71 := 0; fi71 < 17; fi71++ {
            v5 = (v3 / ((v5) | 1))
            if (false || (false || false)) {
                var v72 int8 = 0
                fmt.Println(fi70, v6)
                _ = v72
            } else {
                var v73 int64 = v5
                if ((4294967295 <= int(v3)) && false) {
                    var v74 float64 = 3.14159
                    if (!true) {
                        var v75 int8 = -31
                        v15 = append(v15, v2)
                        _ = v75
                    } else {
                        switch (("hello" + ("\n\t" + "\n\t")) + "\n\t") {
                        case "a":
                            var v76 uint = 65535
                            _ = v76
                        default:
                            if false {
                                v60 += v60
                            }
                        }
                        if false {
                            if false {
                                var v77 int16 = 0
                                var v78 float32 = (v60 - (v60 / v60))
                                if false {
                                    var v79 uint64 = (v62 ^ v62)
                                    v80 := GenS8{F0: 101, F1: -118}
                                    _ = v79
                                    _ = v80
                                } else {
                                    if ("foo bar" < "foo bar") {
                                        var v81 int8 = 0
                                        v77 -= v77
                                        _ = v81
                                    }
                                    v15 = append(v15, (fi70 + 0))
                                    var v82 float64 = (0.1 + ((v74 + (0.1 - float64(v78))) + (((v74 * (v74 * (float64((v78 * v78)) / 1e10))) + (float64((((1.5 - (v60 - (v78 / -2.25))) - 0.1) / v78)) + float64(0.1))) + (float64(1e10) + 1e10))))
                                    _ = v82
                                }
                                v83 := GenS4{F0: 150}
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                _ = v77
                                _ = v78
                                _ = v83
                            }
                            v84 := GenS8{F0: 42, F1: -128}
                            var v85 bool = false
                            _ = v84
                            _ = v85
                        } else {
                            var v86 int64 = v5
                            var v87 float32 = 1.5
                            _ = v86
                            _ = v87
                        }
                        fmt.Println(v74, fi71)
                    }
                    _ = v74
                }
                var v88 uint32 = 42
                _ = v73
                _ = v88
            }
            fmt.Println(v2)
            _ = fi71
        }
        if true {
            for fi89 := 0; fi89 < 18; fi89++ {
                if len(v15) > 0 {
                    v15[0] = (fi89 & fi89)
                }
                _ = fi89
            }
            for fi90 := 0; fi90 < 15; fi90++ {
                v5 += v3
                var v91 int = fi70
                _ = fi90
                _ = v91
            }
        }
        fmt.Println(v6, v62)
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
