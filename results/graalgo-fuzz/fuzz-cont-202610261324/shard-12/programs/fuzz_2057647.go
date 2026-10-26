package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 string }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int; F1 int; F2 int32 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint64; F1 int8; F2 int16 }
type GenS13 struct { F0 uint8; F1 int; F2 uint32 }
type GenS14 struct { F0 int64; F1 uint32; F2 uint64 }
func genFn0(p0 int, p1 uint32) int16 {
    p0 += (p0 | 683)
    v7 := GenS1{F0: 290}
    p1 += (p1 * 707)
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(p0, p0)
        if ((false || (false && true)) || false) {
            p0 = (-(fi8))
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{(p0 >> ((0) & 63)), 1, p0, fi9}
                v11 := GenS1{F0: 0}
                _ = fi9
                _ = v10
                _ = v11
            }
            p1 = (p1 % ((42) | 1))
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 985}
                if (true || false) {
                    v14 := GenS1{F0: 0}
                    _ = v14
                }
                v15 := []int{p0}
                _ = fi12
                _ = v13
                _ = v15
            }
            switch ((("\n\t" + "a") + "a") + ("foo bar" + ("" + ("\n\t" + "a")))) {
            case "\n\t":
                if false {
                    if true {
                        var v16 uint32 = (p1 ^ 4294967295)
                        switch (("" + ((("a" + "foo bar") + "") + "a")) + ("foo bar" + "a")) {
                        case "":
                            var v17 int = 9223372036854775807
                            _ = v17
                        case "foo bar":
                            v16 += p1
                        case "\n\t":
                            p0 -= p0
                        }
                        v18 := GenS2{F0: -54, F1: v16, F2: 100}
                        _ = v16
                        _ = v18
                    }
                }
            case "":
                if true {
                    var v19 string = ""
                    v20 := GenS3{F0: 50, F1: 1000}
                    fmt.Println(p0, p1, p0)
                    _ = v19
                    _ = v20
                }
            default:
                v21 := GenS4{F0: 999, F1: 32767, F2: (("foo bar" + "hello") + ("hello" + "foo bar"))}
                _ = v21
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                switch fi22 {
                case -2147483648:
                    v23 := GenS5{F0: -9223372036854775808, F1: 42}
                    _ = v23
                case 0:
                    var v24 float64 = 1.5
                    _ = v24
                case 4294967295:
                    var v25 uint = 18446744073709551615
                    _ = v25
                }
                _ = fi22
            }
        }
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return -32768
}
func genFn1(p0 uint64, p1 uint16) int8 {
    fmt.Println(p0, p1)
    var v26 uint32 = 1
    fmt.Println(v26)
    p0 = p0
    var v27 int = 832
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 0
}

func main() {
    var v1 bool = (false && (true || false))
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    v4 := []uint32{65535, 702}
    v5 := make(map[uint32]uint32)
    v6 := []string{("\n\t" + "\n\t"), "a", "\n\t"}
    if true {
        if ((!(!true)) || (0 != -9223372036854775808)) {
            switch 1 {
            case -1:
                v1 = v1
            case 42:
                delete(v3, 2147483647)
            case 128:
                var v28 int64 = (^(v2))
                _ = v28
            default:
                if (-1 >= 2147483647) {
                    if v1 {
                        v2 += v2
                        var v29 uint = 10
                        v30 := GenS6{F0: 501, F1: -1}
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS6{F0: -643, F1: 32767}
                    switch 0 {
                    case -28:
                        _ = len(v3)
                    default:
                        switch "hello" {
                        case "hello":
                            for fi32 := 0; fi32 < 7; fi32++ {
                                var v33 int = (fi32 - fi32)
                                _ = fi32
                                _ = v33
                            }
                        case "a":
                            if (3.14159 >= ((1e10 - 1e10) - 0.1)) {
                                var v34 int64 = (v2 + (v2 & v2))
                                var v35 uint8 = 23
                                _ = v34
                                _ = v35
                            } else {
                                if true {
                                    var v36 uint16 = 255
                                    for fi37 := 0; fi37 < 7; fi37++ {
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(fi37)
                                        }
                                        var v38 uint64 = 0
                                        if true {
                                            for fi39 := 0; fi39 < 13; fi39++ {
                                                switch "foo bar" {
                                                case "":
                                                    if (v38 <= v38) {
                                                        v40 := GenS7{F0: 2147483647}
                                                        if (v38 == v38) {
                                                            if (-1 < -1) {
                                                                delete(v3, 330)
                                                            } else {
                                                                var v41 int16 = -969
                                                                v5[2] = 0
                                                                switch -153 {
                                                                case 9223372036854775807:
                                                                    v42 := GenS8{F0: 0, F1: 127, F2: v36}
                                                                    _ = v42
                                                                }
                                                                _ = v41
                                                            }
                                                            v43 := GenS9{F0: fi39, F1: fi37, F2: -1}
                                                            _ = v43
                                                        } else {
                                                            fmt.Println(fi39)
                                                            switch ("hello" + "a") {
                                                            case "hello":
                                                                if ((fi39 + (fi37 & fi37)) == (fi39 ^ fi37)) {
                                                                    var v44 bool = false
                                                                    v2 += (v2 % ((v2) | 1))
                                                                    v45 := GenS9{F0: fi37, F1: 42, F2: 2147483647}
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(fi39)
                                                            case "":
                                                                var v46 uint = 42
                                                                _ = v46
                                                            }
                                                            v1 = (255 != 1)
                                                        }
                                                        v36 = (^(v36))
                                                        _ = v40
                                                    } else {
                                                        var v47 uint8 = 225
                                                        v1 = v1
                                                        v48 := GenS10{F0: 4294967295, F1: 273}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                }
                                                v38 = 2
                                                if v1 {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v49 int32 = -2147483648
                                                    delete(v5, 4294967295)
                                                    v50 := GenS3{F0: (v2 + (-(v2))), F1: -1}
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    var v51 uint32 = 2147483647
                                                    v2 += v2
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if (true || v1) {
                                                            var v52 uint = 19
                                                            if false {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v1 = v1
                                                                    var v53 float64 = 0.1
                                                                    _ = v53
                                                                }
                                                                if v1 {
                                                                    fmt.Println(fi39, v1)
                                                                    if true {
                                                                        var v54 int64 = v2
                                                                        _ = v54
                                                                    } else {
                                                                        if false {
                                                                            var v55 uint16 = 42
                                                                            v55 = uint16(v51)
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v56 uint = uint(v38)
                                                                                _ = v56
                                                                            }
                                                                            _ = v55
                                                                        } else {
                                                                            var v57 int32 = 2147483647
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                    v58 := v5[v51]
                                                                    _ = v58
                                                                } else {
                                                                    v52 -= 65535
                                                                }
                                                                v1 = (fi39 > (fi39 + (^(fi39))))
                                                            } else {
                                                                fmt.Println(v36, v1, v52)
                                                                v1 = v1
                                                                if true {
                                                                    var v59 int64 = 9223372036854775807
                                                                    if v1 {
                                                                        switch (fi39 * (fi39 * fi39)) {
                                                                        case 4294967295:
                                                                            var v60 uint8 = 0
                                                                            _ = v60
                                                                        case 311:
                                                                            var v61 bool = v1
                                                                            if (float64(1e10) == -2.25) {
                                                                                switch ("" + ("a" + "a")) {
                                                                                case "hello":
                                                                                    var v62 int = -356
                                                                                    _ = v62
                                                                                case "\n\t":
                                                                                    var v63 int = 2147483647
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            switch ((("foo bar" + ("" + ("foo bar" + "\n\t"))) + "\n\t") + "a") {
                                                                            case "a":
                                                                                v64 := v5[658]
                                                                                _ = v64
                                                                            case "hello":
                                                                                if (v2 >= v59) {
                                                                                    var v65 int8 = 1
                                                                                    _ = v65
                                                                                }
                                                                            case "foo bar":
                                                                                if false {
                                                                                    v59 = v59
                                                                                    v3[v51] = v51
                                                                                }
                                                                            default:
                                                                                v59 = (v2 * (v2 * (v59 >> (((v52 >> ((v52) & 63))) & 63))))
                                                                            }
                                                                            _ = v61
                                                                        case -697:
                                                                            if (v1 && (("" + "a") >= "\n\t")) {
                                                                                var v66 uint64 = (v38 | 42)
                                                                                v67 := GenS5{F0: v2, F1: 1000}
                                                                                var v68 string = (("\n\t" + "foo bar") + "foo bar")
                                                                                _ = v66
                                                                                _ = v67
                                                                                _ = v68
                                                                            } else {
                                                                                switch int(v51) {
                                                                                case 143:
                                                                                    var v69 int = -986
                                                                                    _ = v69
                                                                                }
                                                                                v2 = (v2 + v59)
                                                                                if ((v2 ^ (v59 * 9223372036854775807)) != (v59 - (v59 << (((^(v52))) & 63)))) {
                                                                                    fmt.Println(v51)
                                                                                    v70 := GenS11{F0: 42}
                                                                                    _ = v70
                                                                                }
                                                                                if (v59 != v2) {
                                                                                    v71 := GenS11{F0: 32767}
                                                                                    switch (fi39 ^ fi37) {
                                                                                    case -558:
                                                                                        v36 -= v36
                                                                                    case 0:
                                                                                        var v72 uint32 = 1
                                                                                        _ = v72
                                                                                    case -373:
                                                                                        var v73 float32 = float32((1e10 * 1.5))
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v71
                                                                                }
                                                                            }
                                                                        default:
                                                                            v74 := GenS12{F0: 18446744073709551615, F1: 0, F2: 125}
                                                                            _ = v74
                                                                        }
                                                                        if (0 != 2147483647) {
                                                                            var v75 int64 = int64(v36)
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 uint64 = v38
                                                                            var v77 uint64 = (v38 >> ((v52) & 63))
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        var v78 int16 = -32768
                                                                        _ = v78
                                                                    } else {
                                                                        if false {
                                                                            v79 := GenS11{F0: 1000}
                                                                            v1 = (!true)
                                                                            var v80 int16 = genFn0(-2147483648, v51)
                                                                            _ = v79
                                                                            _ = v80
                                                                        } else {
                                                                            var v81 int = fi37
                                                                            _ = v81
                                                                        }
                                                                        var v82 string = ("hello" + ("\n\t" + "foo bar"))
                                                                        _ = v82
                                                                    }
                                                                    _ = v59
                                                                }
                                                                v36 += (v36 + (v36 % ((724) | 1)))
                                                            }
                                                            _ = v52
                                                        }
                                                    case "a":
                                                        v83 := GenS6{F0: -1, F1: -32768}
                                                        _ = v83
                                                    default:
                                                        if true {
                                                            switch (("foo bar" + "\n\t") + "a") {
                                                            case "foo bar":
                                                                fmt.Println(fi37, v1)
                                                            case "\n\t":
                                                                var v84 int = -461
                                                                _ = v84
                                                            default:
                                                                v2 = 1
                                                            }
                                                            fmt.Println(fi37, v1)
                                                        } else {
                                                            _ = len(v5)
                                                            v85 := v3[0]
                                                            fmt.Println(fi39)
                                                            _ = v85
                                                        }
                                                    }
                                                    _ = v51
                                                }
                                                _ = fi39
                                            }
                                        }
                                        _ = fi37
                                        _ = v38
                                    }
                                    _ = v36
                                } else {
                                    for fi86 := 0; fi86 < 18; fi86++ {
                                        var v87 uint = 761
                                        v2 -= (v2 | int64(v87))
                                        _ = fi86
                                        _ = v87
                                    }
                                    if v1 {
                                        v4 = append(v4, 40)
                                        switch "a" {
                                        case "":
                                            for fi88 := 0; fi88 < 3; fi88++ {
                                                var v89 uint = 648
                                                if (!v1) {
                                                    v1 = (v1 && (3.14159 < 0.0))
                                                    v1 = true
                                                }
                                                v89 += v89
                                                _ = fi88
                                                _ = v89
                                            }
                                        }
                                        v90 := GenS13{F0: 127, F1: 781, F2: 4294967295}
                                        _ = v90
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = "a"
                                    }
                                }
                                var v91 int64 = (v2 >> ((527) & 63))
                                _ = v91
                            }
                        case "foo bar":
                            fmt.Println(v2, v2)
                        }
                    }
                    _ = v31
                } else {
                    v1 = (v2 >= (^(v2)))
                    fmt.Println(v1, v2)
                }
            }
            var v92 uint8 = 201
            _ = v92
        } else {
            v93 := GenS10{F0: 444, F1: 0}
            _ = v93
        }
        if ((2147483648 <= 574) || true) {
            v94 := GenS13{F0: 42, F1: -75, F2: 65535}
            v95 := GenS14{F0: -418, F1: 2147483648, F2: 2}
            _ = v94
            _ = v95
        }
        for fi96 := 0; fi96 < 16; fi96++ {
            var v97 int64 = v2
            var v98 bool = true
            var v99 int8 = 0
            _ = fi96
            _ = v97
            _ = v98
            _ = v99
        }
    }
    _ = len(v4)
    if ((286 > 32767) && false) {
        v1 = true
    }
    v1 = v1
    _ = len(v4)
    for fi100 := 0; fi100 < 12; fi100++ {
        _ = len(v3)
        v2 -= (v2 % (((v2 + v2)) | 1))
        _ = fi100
    }
    fmt.Println(v1)
    v1 = (!(1e10 > (float32(3.14159) - 1.5)))
    if (!(v2 > (^(v2)))) {
        v1 = false
        v2 = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
