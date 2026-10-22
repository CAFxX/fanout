package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int; F2 int16 }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint64; F1 string; F2 string }
type GenS9 struct { F0 int8; F1 string; F2 int8 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 bool }
func genFn0(p0 int) uint64 {
    var v7 uint32 = 125
    v8 := []int{(p0 - p0), (p0 + (p0 - -9223372036854775808)), 127}
    if true {
        var v9 float64 = ((0.1 - 3.14159) * -2.25)
        var v10 int16 = -506
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS1{F0: 1, F1: 943, F2: (("hello" + "\n\t") + "hello")}
                if true {
                    var v14 int = fi12
                    _ = v14
                } else {
                    switch fi11 {
                    case 4294967295:
                        switch "hello" {
                        case "\n\t":
                            p0 -= 2147483647
                        }
                    case 42:
                        var v15 int16 = 1000
                        _ = v15
                    case 2147483647:
                        fmt.Println(fi11)
                    }
                    v7 = v7
                }
                var v16 int32 = 2147483647
                _ = fi12
                _ = v13
                _ = v16
            }
            var v17 uint8 = 2
            _ = fi11
            _ = v17
        }
    }
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "\n\t"
    v3 := []string{v2, "", v2}
    v4 := []string{}
    var v5 int64 = 171
    var v6 uint64 = 2
    var v18 int32 = 2147483647
    var v19 uint8 = 1
    v20 := GenS2{F0: -94, F1: ((!true) || ((!false) || false))}
    delete(v1, 2147483648)
    switch -9223372036854775808 {
    case 588:
        switch 720 {
        case 0:
            if (false || (true && false)) {
                var v21 string = v2
                switch int(v6) {
                case 0:
                    if (4294967295 != 2147483648) {
                        var v22 uint = 0
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v22
                    } else {
                        var v23 bool = false
                        _ = v23
                    }
                default:
                    v5 = (v5 + 781)
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v18)
                    v21 = ((("\n\t" + ((("a" + v21) + ("foo bar" + "\n\t")) + "a")) + v21) + v21)
                    _ = fi24
                }
                _ = v21
            } else {
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint8 = (v19 * v19)
                    var v27 uint16 = 255
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                var v28 float32 = 1e10
                var v29 uint16 = 65535
                _ = v28
                _ = v29
            }
        }
    case 42:
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int32 = v18
            for fi32 := 0; fi32 < 11; fi32++ {
                switch "\n\t" {
                case "\n\t":
                    v5 = (v5 >> ((4294967295) & 63))
                case "hello":
                    var v33 uint = 1
                    _ = v33
                }
                v34 := GenS2{F0: 1, F1: (10 <= -128)}
                v35 := GenS3{F0: int64(fi30), F1: v2}
                _ = fi32
                _ = v34
                _ = v35
            }
            switch v2 {
            case "a":
                v36 := GenS4{F0: -360, F1: v6}
                _ = v36
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi30
            _ = v31
        }
    default:
        for fi37 := 0; fi37 < 20; fi37++ {
            v18 -= 2147483647
            if (!(v19 >= (^(v19)))) {
                for fi38 := 0; fi38 < 8; fi38++ {
                    v39 := GenS5{F0: 1000}
                    if false {
                        switch fi38 {
                        case -9223372036854775808:
                            v2 = ""
                        case 42:
                            switch (fi38 * fi38) {
                            case 32767:
                                if false {
                                    if (true && false) {
                                        var v40 float32 = (((3.14159 / -2.25) / (((0.0 / -2.25) + 1e10) - 1.5)) + (-2.25 - 1.5))
                                        v41 := GenS6{F0: 282, F1: (fi38 ^ fi38), F2: 32767}
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                var v42 int64 = v5
                                v43 := GenS7{F0: false, F1: v6}
                                v44 := GenS1{F0: 349, F1: 32767, F2: v2}
                                _ = v42
                                _ = v43
                                _ = v44
                            case 810:
                                v6 = 738
                            case 127:
                                if ((false || (false || ((255 < 32767) && ((v6 + 18446744073709551615) < (v6 % (((^(v6))) | 1)))))) && ((v2 + "foo bar") <= "\n\t")) {
                                    v45 := GenS7{F0: ((!((^(v6)) < (v6 * (v6 & (v6 - v6))))) && true), F1: 0}
                                    var v46 float32 = 1e10
                                    _ = v45
                                    _ = v46
                                } else {
                                    fmt.Println(v2, fi38, fi38)
                                }
                                v18 = -823
                                var v47 float64 = 0.1
                                _ = v47
                            default:
                                var v48 int8 = 127
                                _ = v48
                            }
                        case 771:
                            delete(v1, 65535)
                        default:
                            var v49 uint64 = 2
                            _ = v49
                        }
                        v50 := GenS8{F0: v6, F1: v2, F2: v2}
                        _ = v50
                    } else {
                        switch (fi38 << ((165) & 63)) {
                        case -26:
                            var v51 uint16 = 1
                            switch "a" {
                            case "":
                                var v52 float32 = -2.25
                                _ = v52
                            default:
                                v4 = append(v4, "foo bar")
                            }
                            _ = v51
                        case 1:
                            v53 := GenS9{F0: 127, F1: v2, F2: 100}
                            _ = v53
                        default:
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                        }
                        var v54 uint32 = 4294967295
                        var v55 int16 = 0
                        _ = v54
                        _ = v55
                    }
                    if (true || true) {
                        switch fi37 {
                        case -2147483648:
                            if (("foo bar" > (v2 + "hello")) || true) {
                                _ = len(v4)
                                var v56 uint32 = 4294967295
                                _ = v56
                            } else {
                                v4 = append(v4, v2)
                                switch v2 {
                                case "a":
                                    if ((v19 - v19) <= 255) {
                                        var v57 uint16 = 42
                                        v57 -= 772
                                        _ = v57
                                    } else {
                                        fmt.Println(v2)
                                    }
                                default:
                                    var v58 int32 = (^(v18))
                                    _ = v58
                                }
                            }
                        }
                        fmt.Println(fi38)
                        if true {
                            v59 := v1[2147483647]
                            switch (fi38 * v59) {
                            case -869:
                                delete(v1, 370)
                            case 407:
                                if (v2 <= v2) {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                }
                            default:
                                switch (("foo bar" + ("\n\t" + (((((v2 + "foo bar") + (v2 + ("" + ((v2 + ("hello" + ("" + v2))) + v2)))) + ("a" + "\n\t")) + (("" + ((v2 + (v2 + "\n\t")) + "")) + v2)) + v2))) + v2) {
                                case "a":
                                    v60 := v1[1]
                                    _ = v60
                                }
                            }
                            if (!true) {
                                fmt.Println(v2, fi37)
                                var v61 int8 = -112
                                var v62 float64 = 0.1
                                _ = v61
                                _ = v62
                            } else {
                                if (v2 >= v2) {
                                    v2 = (v2 + ("" + "foo bar"))
                                } else {
                                    v63 := GenS7{F0: true, F1: (v6 ^ v6)}
                                    _ = len(v3)
                                    _ = v63
                                }
                            }
                            switch (fi37 & v59) {
                            case 0:
                                var v64 int16 = -1
                                _ = v64
                            case 128:
                                switch (fi38 + v59) {
                                case 42:
                                    if false {
                                        v59 -= (-(v59))
                                        if false {
                                            var v65 uint = 42
                                            _ = v65
                                        } else {
                                            fmt.Println(fi38)
                                            var v66 int64 = int64(v6)
                                            var v67 uint8 = (v19 - v19)
                                            var v68 uint16 = 0
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                case 2147483647:
                                    var v69 int64 = (v5 * (v5 * v5))
                                    _ = v69
                                }
                                v4 = append(v4, "")
                                if (("hello" + "foo bar") < v2) {
                                    switch fi38 {
                                    case 4294967295:
                                        var v70 int = int(v6)
                                        _ = v70
                                    }
                                    if true {
                                        var v71 bool = (!(65535 == 4294967295))
                                        if (!v71) {
                                            switch 614 {
                                            case 2147483647:
                                                switch (v59 & v59) {
                                                case 32768:
                                                    if v71 {
                                                        v2 = v2
                                                    } else {
                                                        v72 := GenS5{F0: -166}
                                                        delete(v1, 698)
                                                        _ = v72
                                                    }
                                                case 4294967295:
                                                    var v73 float64 = 3.14159
                                                    _ = v73
                                                }
                                            default:
                                                fmt.Println(v19)
                                                v2 = (("hello" + (v2 + (v2 + v2))) + (v2 + v2))
                                            }
                                            v74 := GenS10{F0: 42, F1: (v19 | v19)}
                                            v71 = v71
                                            var v75 uint16 = 1
                                            _ = v74
                                            _ = v75
                                        } else {
                                            _ = len(v4)
                                            if (!(2147483648 < 2)) {
                                                var v76 float64 = 1e10
                                                _ = v76
                                            }
                                        }
                                        _ = v71
                                    } else {
                                        var v77 uint8 = v19
                                        _ = len(v1)
                                        _ = v77
                                    }
                                } else {
                                    var v78 int32 = 463
                                    if true {
                                        if ((^(v78)) != (v78 & -2147483648)) {
                                            var v79 int = fi38
                                            _ = len(v3)
                                            var v80 uint = 255
                                            _ = v79
                                            _ = v80
                                        } else {
                                            v81 := v1[uint32(v19)]
                                            fmt.Println(v5, v18)
                                            v78 = (v78 & 366)
                                            _ = v81
                                        }
                                    } else {
                                        if true {
                                            if false {
                                                if (v2 < v2) {
                                                    var v82 float32 = 0.0
                                                    switch (v59 / ((fi38) | 1)) {
                                                    case 32767:
                                                        var v83 int32 = (^(v78))
                                                        _ = v83
                                                    case 32768:
                                                        v84 := GenS1{F0: uint32(v19), F1: 1, F2: "a"}
                                                        _ = v84
                                                    case -2147483648:
                                                        v85 := GenS7{F0: (18446744073709551615 != 2), F1: (v6 - 2)}
                                                        _ = v85
                                                    default:
                                                        v19 -= (v19 | 247)
                                                    }
                                                    _ = v82
                                                }
                                            }
                                            var v86 float64 = 1e10
                                            v78 += -1
                                            _ = v86
                                        }
                                        v59 += v59
                                    }
                                    _ = v78
                                }
                            case 291:
                                if false {
                                    var v87 int8 = -1
                                    v18 += (^(v18))
                                    _ = v87
                                }
                                switch v59 {
                                case 127:
                                    var v88 float32 = (1.5 - 0.0)
                                    _ = v88
                                default:
                                    var v89 int8 = 42
                                    _ = v89
                                }
                                if false {
                                    if false {
                                        var v90 uint16 = 441
                                        var v91 uint64 = v6
                                        var v92 uint32 = 537
                                        v91 -= v91
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    } else {
                                        var v93 string = v2
                                        var v94 uint32 = 160
                                        _ = v93
                                        _ = v94
                                    }
                                }
                            default:
                                if true {
                                    switch fi38 {
                                    case -2147483648:
                                        fmt.Println(fi37, fi38, fi38)
                                    }
                                    var v95 uint = 742
                                    _ = v95
                                } else {
                                    v96 := v1[663]
                                    v97 := GenS11{F0: false}
                                    var v98 uint16 = 1
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                            }
                            v18 = -1
                            v2 = "\n\t"
                            if (("" < "hello") || true) {
                                v1[807] = (fi37 ^ fi37)
                                fmt.Println(fi38)
                            } else {
                                v6 = v6
                                var v99 int32 = v18
                                _ = v99
                            }
                            _ = v59
                        } else {
                            v6 += (v6 - (v6 / (((^(v6))) | 1)))
                        }
                    } else {
                        var v100 bool = false
                        if (2147483647 < 42) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v18 -= v18
                        }
                        delete(v1, 1)
                        _ = v100
                    }
                    _ = fi38
                    _ = v39
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = ("\n\t" + ("\n\t" + v2))
            }
            _ = fi37
        }
    }
    if true {
        for fi101 := 0; fi101 < 11; fi101++ {
            fmt.Println(fi101)
            var v102 uint64 = 624
            if (((v2 + v2) >= "") && (1000 == 42)) {
                if true {
                    _ = len(v4)
                    var v103 uint8 = 1
                    _ = v103
                } else {
                    v102 += 18446744073709551615
                    v6 = genFn0((fi101 >> ((1) & 63)))
                }
                v2 = v2
            }
            _ = fi101
            _ = v102
        }
        fmt.Println(v2, v2)
    }
    fmt.Println(v6, v6, v19)
    v104 := GenS5{F0: 1000}
    var v105 float64 = 0.1
    fmt.Println(v5, v6, v18)
    switch "a" {
    case "":
        if (!(-1 <= 1)) {
            var v106 float64 = v105
            _ = v106
        }
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = v2
        }
    case "\n\t":
        var v107 uint = 955
        _ = v107
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v104
    _ = v105
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
