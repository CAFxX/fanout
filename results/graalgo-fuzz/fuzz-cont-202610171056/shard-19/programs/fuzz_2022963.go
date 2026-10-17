package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint64, p1 uint8, p2 string) uint32 {
    switch p2 {
    case "\n\t":
        var v7 int16 = -32768
        _ = v7
    case "a":
        var v8 string = "a"
        _ = v8
    case "":
        var v9 int8 = 2
        _ = v9
    default:
        if (!(3.14159 < float32((0.0 + 0.0)))) {
            for fi10 := 0; fi10 < 11; fi10++ {
                if (((^(p0)) < p0) || true) {
                    var v11 int32 = 0
                    _ = v11
                }
                if ((-1 >= -298) && true) {
                    if false {
                        v12 := []int{(^(fi10))}
                        _ = v12
                    }
                } else {
                    var v13 int32 = 0
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v15 := []int{(^(fi10)), (fi14 >> ((65535) & 63)), (fi14 / ((fi10) | 1)), (fi10 ^ (fi10 ^ fi14))}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        v16 := GenS1{F0: 13, F1: 583}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                }
                var v17 int32 = 342
                _ = fi10
                _ = v17
            }
            var v18 int64 = -1
            var v19 float32 = 1.5
            _ = v18
            _ = v19
        } else {
            fmt.Println(p0)
            p1 += p1
        }
    }
    var v20 int8 = 127
    if false {
        var v21 uint8 = (p1 % ((42) | 1))
        _ = v21
    }
    fmt.Println(p0, p1)
    if (127 < -933) {
        var v22 float32 = 0.0
        _ = v22
    } else {
        v20 = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (-2.25 / 1.5)
    var v3 uint16 = 42
    v4 := []string{"hello"}
    var v5 uint8 = 2
    var v6 int32 = 42
    v5 = (v5 * v5)
    for fi23 := 0; fi23 < 3; fi23++ {
        if ((-128 >= 1) || true) {
            switch (fi23 >> ((42) & 63)) {
            case 878:
                var v24 uint = 42
                _ = v24
            case 0:
                _ = len(v1)
            }
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                if (255 != 1) {
                    var v26 float32 = ((v2 * v2) * v2)
                    fmt.Println(v6, fi23)
                    v27 := GenS2{F0: (v6 % ((42) | 1))}
                    _ = v26
                    _ = v27
                }
                v4 = append(v4, "")
                v4 = append(v4, "a")
                fmt.Println(v6, v2, fi23)
                if len(v4) > 0 {
                    v4[0] = ((("a" + "foo bar") + "") + "a")
                }
                _ = fi25
            }
            if (!(true && (v6 >= (v6 >> ((421) & 63))))) {
                var v28 uint8 = 42
                if false {
                    switch 32767 {
                    case -195:
                        v28 = v5
                    case 2147483647:
                        var v29 int32 = (v6 / ((420) | 1))
                        _ = v29
                    }
                }
                _ = v28
            } else {
                switch "foo bar" {
                case "hello":
                    for fi30 := 0; fi30 < 7; fi30++ {
                        fmt.Println(v2)
                        _ = fi30
                    }
                case "a":
                    var v31 int8 = 0
                    _ = v31
                }
                if (0 != (v5 >> ((296) & 63))) {
                    fmt.Println(v5)
                    fmt.Println(v6, v6)
                } else {
                    switch (fi23 & fi23) {
                    case 128:
                        if (-128 > 100) {
                            for fi32 := 0; fi32 < 12; fi32++ {
                                v33 := GenS3{F0: ("hello" + ("foo bar" + "hello")), F1: ("\n\t" + ("foo bar" + "hello"))}
                                if false {
                                    var v34 int16 = 1000
                                    _ = v34
                                } else {
                                    v5 = (v5 & v5)
                                    v3 = v3
                                }
                                v5 = (v5 & (^(v5)))
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                }
                if true {
                    var v35 bool = ((v6 ^ v6) != (v6 >> ((255) & 63)))
                    _ = v35
                } else {
                    var v36 int64 = -1
                    fmt.Println(fi23, v2)
                    _ = v36
                }
            }
            v6 += -2147483648
        }
        var v37 int8 = 100
        v37 = v37
        _ = fi23
        _ = v37
    }
    v4 = append(v4, ((("a" + ("foo bar" + ((("\n\t" + ("foo bar" + ("hello" + "foo bar"))) + "hello") + "\n\t"))) + "foo bar") + "hello"))
    if len(v4) > 0 {
        v4[0] = ("\n\t" + ("a" + ("hello" + "")))
    }
    var v38 int64 = -1
    if len(v4) > 0 {
        v4[0] = ("\n\t" + "\n\t")
    }
    if (!(true && true)) {
        v5 -= 1
        for fi39 := 0; fi39 < 17; fi39++ {
            v2 -= 0.0
            _ = fi39
        }
    }
    switch ("" + "foo bar") {
    case "":
        v40 := GenS1{F0: -1, F1: 2}
        _ = v40
    case "a":
        if false {
            fmt.Println(v38)
        } else {
            v41 := GenS1{F0: 15, F1: 65535}
            v42 := GenS4{F0: 2, F1: v3, F2: v6}
            if false {
                v43 := GenS1{F0: 100, F1: 42}
                _ = v43
            }
            _ = v41
            _ = v42
        }
    default:
        if len(v4) > 0 {
            v4[0] = ""
        }
    }
    var v44 int32 = 0
    v4 = append(v4, "foo bar")
    v45 := GenS2{F0: v44}
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 int64 = (^(v38))
        fmt.Println(v6)
        _ = fi46
        _ = v47
    }
    switch -338 {
    case -879:
        if ((v2 + (-2.25 * float32((3.14159 - -2.25)))) < 1.5) {
            switch "hello" {
            case "foo bar":
                var v48 int32 = v6
                _ = v48
            case "\n\t":
                for fi49 := 0; fi49 < 18; fi49++ {
                    var v50 float64 = float64((v2 - v2))
                    if false {
                        var v51 bool = (v50 == float64(v2))
                        _ = len(v1)
                        var v52 uint8 = v5
                        _ = v51
                        _ = v52
                    }
                    _ = fi49
                    _ = v50
                }
            default:
                if (890 == 42) {
                    switch -2147483648 {
                    case 2147483647:
                        for fi53 := 0; fi53 < 7; fi53++ {
                            switch fi53 {
                            case -1:
                                v5 -= 181
                            }
                            var v54 float64 = (0.1 / 0.1)
                            _ = fi53
                            _ = v54
                        }
                    case 743:
                        if true {
                            var v55 string = ""
                            v1[9223372036854775807] = true
                            if false {
                                _ = len(v1)
                                switch -9223372036854775808 {
                                case 42:
                                    var v56 uint32 = 42
                                    _ = v56
                                case 32767:
                                    for fi57 := 0; fi57 < 5; fi57++ {
                                        for fi58 := 0; fi58 < 4; fi58++ {
                                            var v59 int64 = (^(v38))
                                            v3 += (v3 & 213)
                                            fmt.Println(fi58, v59, v6)
                                            _ = fi58
                                            _ = v59
                                        }
                                        if true {
                                            for fi60 := 0; fi60 < 5; fi60++ {
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                                _ = fi60
                                            }
                                            if (!true) {
                                                var v61 int64 = (v38 | (v38 * (-(v38))))
                                                _ = v61
                                            } else {
                                                if true {
                                                    var v62 uint32 = 2147483647
                                                    var v63 float32 = v2
                                                    _ = len(v1)
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    v1[fi57] = (2 >= 42)
                                                    fmt.Println(v55, v38)
                                                }
                                            }
                                            switch fi57 {
                                            case -1:
                                                var v64 uint8 = v5
                                                _ = v64
                                            case 1:
                                                switch v55 {
                                                case "foo bar":
                                                    var v65 int8 = 77
                                                    _ = v65
                                                default:
                                                    fmt.Println(fi57, v2)
                                                }
                                            }
                                        }
                                        for fi66 := 0; fi66 < 10; fi66++ {
                                            if false {
                                                if (v44 >= (v44 | int32(v38))) {
                                                    if ((v3 & (v3 ^ v3)) > (v3 << ((255) & 63))) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v67 := GenS4{F0: 1, F1: (v3 | v3), F2: 2147483647}
                                                            _ = v67
                                                        case "foo bar":
                                                            var v68 float64 = 0.0
                                                            _ = v68
                                                        default:
                                                            var v69 uint16 = 321
                                                            v6 = (v44 & v44)
                                                            _ = v69
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        switch (fi57 << ((2) & 63)) {
                                                        case 738:
                                                            var v70 uint32 = genFn0(1, 183, (((v55 + "") + v55) + v55))
                                                            _ = v70
                                                        case 32768:
                                                            v4 = append(v4, "")
                                                        case 4294967295:
                                                            switch (("foo bar" + ("hello" + v55)) + "hello") {
                                                            case "a":
                                                                if false {
                                                                    v71 := v1[128]
                                                                    var v72 uint64 = 2
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            }
                                                        }
                                                        switch v55 {
                                                        case "hello":
                                                            var v73 float32 = v2
                                                            _ = v73
                                                        default:
                                                            v44 = v44
                                                        }
                                                    }
                                                    var v74 uint16 = v3
                                                    _ = v74
                                                }
                                                var v75 uint32 = genFn0(321, (v5 >> ((18446744073709551615) & 63)), v55)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v75
                                            }
                                            fmt.Println(v6, v6)
                                            _ = fi66
                                        }
                                        _ = fi57
                                    }
                                }
                                switch -798 {
                                case 816:
                                    v44 = 195
                                case 1:
                                    switch 1 {
                                    case 2147483647:
                                        v76 := GenS3{F0: v55, F1: "foo bar"}
                                        _ = v76
                                    case 164:
                                        var v77 int = -9223372036854775808
                                        _ = v77
                                    case 9223372036854775807:
                                        v3 = v3
                                    default:
                                        v1[-683] = false
                                    }
                                case -2147483648:
                                    v78 := GenS3{F0: "", F1: ("" + (v55 + v55))}
                                    _ = v78
                                }
                            } else {
                                var v79 int32 = v6
                                switch v55 {
                                case "\n\t":
                                    v80 := GenS4{F0: 2147483647, F1: (v3 ^ (v3 ^ 255)), F2: (v79 | (v79 << ((42) & 63)))}
                                    _ = v80
                                case "hello":
                                    fmt.Println(v2, v44, v3)
                                default:
                                    if (true || ((3.14159 - 1.5) > 0.1)) {
                                        var v81 int = -1
                                        if false {
                                            var v82 uint32 = 65535
                                            _ = v82
                                        }
                                        _ = v81
                                    } else {
                                        switch -1 {
                                        case -305:
                                            if false {
                                                for fi83 := 0; fi83 < 7; fi83++ {
                                                    v84 := GenS5{F0: 416, F1: uint8(v3)}
                                                    var v85 bool = true
                                                    var v86 int32 = (v79 ^ (v79 * v6))
                                                    _ = fi83
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                }
                                                if (32767 > -2147483648) {
                                                    var v87 uint8 = (v5 + (v5 | 1))
                                                    var v88 bool = (18446744073709551615 > 448)
                                                    var v89 int16 = -1
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    var v90 int32 = v6
                                                    _ = v90
                                                }
                                                switch 32768 {
                                                case 4294967295:
                                                    switch 42 {
                                                    case 428:
                                                        fmt.Println(v6, v44, v79)
                                                    case 1:
                                                        fmt.Println(v44)
                                                    }
                                                case -1:
                                                    v5 = v5
                                                default:
                                                    for fi91 := 0; fi91 < 11; fi91++ {
                                                        var v92 uint8 = 127
                                                        v79 += v6
                                                        fmt.Println(v5)
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                }
                                            } else {
                                                v2 += 1.5
                                                v44 -= v44
                                            }
                                        case 42:
                                            switch 4294967295 {
                                            case 319:
                                                fmt.Println(v44)
                                            }
                                        case 127:
                                            v79 -= (v44 + v6)
                                        default:
                                            v93 := GenS3{F0: "foo bar", F1: (v55 + v55)}
                                            _ = v93
                                        }
                                    }
                                }
                                _ = len(v4)
                                _ = v79
                            }
                            _ = v55
                        } else {
                            var v94 uint64 = 18446744073709551615
                            if len(v4) > 0 {
                                v4[0] = ("\n\t" + "a")
                            }
                            fmt.Println(v2)
                            _ = v94
                        }
                    case 42:
                        v6 = (v44 - (v6 - v6))
                    }
                }
            }
        }
    case 9223372036854775807:
        v38 += (-(v38))
    }
    fmt.Println(v2, v38, v6)
    v6 -= (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
