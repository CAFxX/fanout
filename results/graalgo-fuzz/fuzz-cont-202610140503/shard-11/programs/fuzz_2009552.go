package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint8; F2 bool }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int64; F1 string; F2 int64 }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16 }
func genFn0() int8 {
    switch 9223372036854775807 {
    case 9223372036854775807:
        v6 := GenS1{F0: 2147483648, F1: 42, F2: 127}
        _ = v6
    case -9223372036854775808:
        var v7 uint64 = 815
        _ = v7
    }
    var v8 int32 = 42
    _ = v8
    return 100
}
func genFn1(p0 int32, p1 int) int {
    v9 := GenS2{F0: 0}
    p1 += (p1 + p1)
    fmt.Println(p0, p0)
    switch (-(p1)) {
    case 32768:
        if (false || false) {
            var v10 string = ""
            var v11 int8 = genFn0()
            _ = v10
            _ = v11
        } else {
            if (!true) {
                var v12 uint64 = 2
                var v13 uint16 = 347
                v12 -= v12
                _ = v12
                _ = v13
            } else {
                v14 := GenS3{F0: 117}
                fmt.Println(p1)
                _ = v14
            }
            p0 = 2147483647
            p0 -= p0
        }
    case 0:
        p0 += 2147483647
    case -501:
        v15 := []int{}
        _ = v15
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            if (p0 > (p0 << ((4294967295) & 63))) {
                var v17 int16 = 42
                _ = v17
            } else {
                fmt.Println(p0, p0)
            }
            var v18 float64 = 1.5
            _ = fi16
            _ = v18
        }
    }
    if (false && false) {
        p0 -= -371
        if false {
            switch ((("foo bar" + "hello") + ("\n\t" + "")) + "hello") {
            case "\n\t":
                if ((0 < -128) && (("\n\t" == "foo bar") && (622 > 18446744073709551615))) {
                    v19 := GenS4{F0: ((!true) || true), F1: -9223372036854775808, F2: 32767}
                    v20 := GenS2{F0: int64(p0)}
                    _ = v19
                    _ = v20
                }
            }
            var v21 uint64 = 18446744073709551615
            p0 = (p0 << ((42) & 63))
            _ = v21
        } else {
            fmt.Println(p0)
            v22 := []int{(p1 * p1)}
            p1 = (p1 - 42)
            _ = v22
        }
        var v23 uint = 743
        _ = v23
    } else {
        switch p1 {
        case -242:
            var v24 float32 = 3.14159
            _ = v24
        }
        fmt.Println(p1)
        var v25 float64 = 1.5
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v9
    return (p1 >> ((2) & 63))
}

func main() {
    var v1 int8 = 0
    v2 := []int{-274, 128, 42, -162}
    v3 := []string{}
    var v4 string = ""
    var v5 float64 = 0.0
    switch -1 {
    case 2147483647:
        var v26 float64 = (v5 * v5)
        _ = v26
    }
    v27 := GenS2{F0: 1}
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS3{F0: 0}
        v4 = ("" + v4)
        v30 := GenS1{F0: 4294967295, F1: 13, F2: 2}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint8 = 0
    var v32 float32 = float32(float64(0.1))
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v5)
        _ = fi33
    }
    _ = len(v2)
    for fi34 := 0; fi34 < 5; fi34++ {
        switch fi34 {
        case 128:
            v1 += genFn0()
        case -845:
            v35 := GenS5{F0: (868 >= 2), F1: v31}
            _ = v35
        case 0:
            switch ((v4 + v4) + ((v4 + "hello") + v4)) {
            case "hello":
                v5 = float64(3.14159)
            case "a":
                v31 += v31
            default:
                for fi36 := 0; fi36 < 13; fi36++ {
                    v3 = append(v3, "hello")
                    v4 = ("foo bar" + v4)
                    _ = fi36
                }
            }
        default:
            if false {
                v37 := GenS4{F0: (!(!true)), F1: 42, F2: 255}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 14; fi38++ {
                    fmt.Println(v4, v1, fi34)
                    _ = fi38
                }
                if ((v1 + (v1 << ((626) & 63))) != v1) {
                    var v39 int64 = 0
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v40 float32 = float32(v5)
                    _ = v39
                    _ = v40
                } else {
                    if (2 == 670) {
                        var v41 int = fi34
                        _ = v41
                    } else {
                        v31 = (v31 * (v31 % (((v31 ^ 127)) | 1)))
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v42 := GenS5{F0: false, F1: 255}
                    _ = v42
                }
            }
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            switch fi34 {
            case -1:
                v44 := GenS6{F0: 1, F1: 127}
                _ = v44
            case 9223372036854775807:
                fmt.Println(v32, v5)
            case 32768:
                if false {
                    var v45 uint16 = 870
                    if ((fi34 + (fi43 - (fi34 & (fi34 + (fi34 >> ((4294967295) & 63)))))) > 32767) {
                        if true {
                            var v46 string = (v4 + "foo bar")
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = 32768
                                }
                            }
                            _ = v46
                        }
                        fmt.Println(v4, v4)
                        var v47 int16 = -32768
                        v5 = 3.14159
                        _ = v47
                    } else {
                        v32 -= (3.14159 / (1.5 - v32))
                        var v48 uint16 = 478
                        var v49 uint8 = v31
                        _ = v48
                        _ = v49
                    }
                    _ = v45
                }
            default:
                v50 := GenS1{F0: 48, F1: -927, F2: (v31 - 128)}
                _ = v50
            }
            _ = fi43
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi34
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        if false {
            if (true || ((v5 + ((float64(v32) + 0.1) / 0.1)) <= v5)) {
                for fi52 := 0; fi52 < 5; fi52++ {
                    var v53 bool = false
                    v54 := GenS7{F0: 938, F1: v31, F2: (true && v53)}
                    var v55 bool = (float32(v5) != 0.0)
                    v53 = v53
                    if true {
                        if (!(316 >= 663)) {
                            v56 := GenS7{F0: 4294967295, F1: (^(v31)), F2: (!(v1 == 7))}
                            var v57 uint8 = (^(v31))
                            switch "\n\t" {
                            case "\n\t":
                                switch genFn1(42, fi51) {
                                case 42:
                                    var v58 float32 = 0.1
                                    _ = v58
                                case 0:
                                    switch fi52 {
                                    case 127:
                                        var v59 uint64 = 18446744073709551615
                                        _ = v59
                                    default:
                                        switch int(v1) {
                                        case -1:
                                            if (2 > 65535) {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi51 << ((255) & 63))
                                                }
                                                if ("foo bar" <= "foo bar") {
                                                    v60 := GenS3{F0: v57}
                                                    switch fi52 {
                                                    case -9223372036854775808:
                                                        var v61 bool = (!true)
                                                        _ = v61
                                                    case 350:
                                                        if (!(false || ((v57 >> ((0) & 63)) > v57))) {
                                                            v55 = (v57 > 128)
                                                        }
                                                    default:
                                                        var v62 bool = true
                                                        _ = v62
                                                    }
                                                    if v53 {
                                                        var v63 int8 = -40
                                                        v64 := GenS6{F0: 18446744073709551615, F1: v1}
                                                        switch fi52 {
                                                        case 127:
                                                            var v65 float32 = v32
                                                            _ = v65
                                                        case 1:
                                                            v31 = (^(v31))
                                                        case -694:
                                                            _ = len(v2)
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 32767
                                                        }
                                                        if (!(!false)) {
                                                            var v66 uint64 = 4294967295
                                                            var v67 bool = false
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    _ = v60
                                                } else {
                                                    fmt.Println(v32)
                                                    var v68 int32 = 2147483647
                                                    _ = v68
                                                }
                                                var v69 bool = (!true)
                                                _ = len(v2)
                                                switch (fi52 + (fi52 << ((1) & 63))) {
                                                case 127:
                                                    var v70 uint = 878
                                                    _ = v70
                                                }
                                                _ = v69
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if false {
                                                    var v71 float64 = ((((v5 + (v5 * 1e10)) + 0.1) / 1e10) / v5)
                                                    fmt.Println(v5, fi52, v31)
                                                    _ = v71
                                                }
                                                if v53 {
                                                    var v72 string = v4
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v1 = (v1 | v1)
                                                    switch (fi51 << ((42) & 63)) {
                                                    case 0:
                                                        var v73 int64 = -1
                                                        _ = len(v2)
                                                        var v74 int32 = -1
                                                        _ = v73
                                                        _ = v74
                                                    case -2147483648:
                                                        if (1 >= -8) {
                                                            switch int(v1) {
                                                            case 32768:
                                                                switch (fi51 & fi52) {
                                                                case -2147483648:
                                                                    v55 = (286 != 1)
                                                                case 4294967295:
                                                                    var v75 float32 = 1.5
                                                                    _ = v75
                                                                case -91:
                                                                    v72 = "a"
                                                                }
                                                            default:
                                                                v1 += (v1 % ((genFn0()) | 1))
                                                            }
                                                            if v53 {
                                                                v4 = v72
                                                                v32 = float32(v5)
                                                            } else {
                                                                switch (fi51 << ((79) & 63)) {
                                                                case 127:
                                                                    if v55 {
                                                                        fmt.Println(v5, v4, v31)
                                                                    } else {
                                                                        v76 := GenS7{F0: 1, F1: (v57 - 0), F2: v53}
                                                                        v72 = (("foo bar" + "") + "a")
                                                                        fmt.Println(v53, v32)
                                                                        _ = v76
                                                                    }
                                                                    if v53 {
                                                                        var v77 uint64 = 57
                                                                        var v78 int64 = 9223372036854775807
                                                                        switch genFn1(42, (-(fi51))) {
                                                                        case 0:
                                                                            var v79 float32 = (float32(v5) * float32(1.5))
                                                                            _ = v79
                                                                        case 279:
                                                                            var v80 int64 = v78
                                                                            _ = v80
                                                                        default:
                                                                            fmt.Println(v1)
                                                                        }
                                                                        if (((v77 << ((717) & 63)) >= (v77 - (v77 & v77))) || (v77 >= (v77 / ((155) | 1)))) {
                                                                            var v81 uint8 = 1
                                                                            _ = v81
                                                                        }
                                                                        _ = v77
                                                                        _ = v78
                                                                    } else {
                                                                        if (!(v55 || v55)) {
                                                                            var v82 uint8 = 0
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (fi51 - (-(fi52)))
                                                                            }
                                                                            switch v72 {
                                                                            case "foo bar":
                                                                                var v83 uint = 1
                                                                                _ = v83
                                                                            default:
                                                                                var v84 uint32 = 2147483647
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v4
                                                                                }
                                                                                var v85 uint = 2
                                                                                _ = v84
                                                                                _ = v85
                                                                            }
                                                                            _ = v82
                                                                        } else {
                                                                            fmt.Println(v4)
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v72 + v4)
                                                                        }
                                                                    }
                                                                    var v86 uint64 = 18446744073709551615
                                                                    v1 += (v1 ^ 42)
                                                                    if false {
                                                                        var v87 uint64 = 2
                                                                        _ = len(v2)
                                                                        v88 := GenS8{F0: (^(fi51)), F1: (v57 / ((255) | 1))}
                                                                        _ = v87
                                                                        _ = v88
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 1
                                                                        }
                                                                        var v89 uint8 = v31
                                                                        if (-41 > 1) {
                                                                            v90 := GenS6{F0: (v86 % ((v86) | 1)), F1: 54}
                                                                            var v91 int64 = -29
                                                                            if (!(((0.1 / v32) / (-2.25 * v32)) >= v32)) {
                                                                                v91 -= v91
                                                                                var v92 int16 = -32768
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                _ = v92
                                                                            } else {
                                                                                if v53 {
                                                                                    var v93 int = fi51
                                                                                    v94 := GenS9{F0: v91, F1: (v72 + v4), F2: (v91 + v91)}
                                                                                    switch 32767 {
                                                                                    case 1:
                                                                                        v95 := GenS2{F0: v91}
                                                                                        _ = v95
                                                                                    case 128:
                                                                                        var v96 int16 = 959
                                                                                        _ = v96
                                                                                    default:
                                                                                        if v55 {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            var v97 int = (fi51 ^ (v93 ^ (v93 % ((fi51) | 1))))
                                                                                            v3 = append(v3, (v72 + v4))
                                                                                            _ = v97
                                                                                        }
                                                                                    }
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                var v98 int8 = int8(v91)
                                                                                v55 = (!((v4 + (v72 + v4)) <= "foo bar"))
                                                                                _ = v98
                                                                            }
                                                                            _ = v90
                                                                            _ = v91
                                                                        } else {
                                                                            switch fi52 {
                                                                            case -2147483648:
                                                                                var v99 uint32 = 37
                                                                                _ = v99
                                                                            case 1:
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v100 := GenS6{F0: v86, F1: v1}
                                                                                    _ = v100
                                                                                case "foo bar":
                                                                                    var v101 float64 = v5
                                                                                    _ = v101
                                                                                default:
                                                                                    if (v53 && false) {
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v32 = (v32 * -2.25)
                                                                                        case "a":
                                                                                            v4 = (("a" + ("foo bar" + "a")) + "")
                                                                                        case "\n\t":
                                                                                            v4 = (v72 + "\n\t")
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = v72
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v89
                                                                    }
                                                                    fmt.Println(v1, fi52)
                                                                    _ = v86
                                                                case -2147483648:
                                                                    v102 := GenS4{F0: false, F1: 1, F2: 255}
                                                                    _ = v102
                                                                default:
                                                                    fmt.Println(v32, v32)
                                                                }
                                                            }
                                                            fmt.Println(v32, v72, v1)
                                                        }
                                                    default:
                                                        fmt.Println(v4)
                                                    }
                                                    var v103 float32 = (float32(v5) + v32)
                                                    _ = v103
                                                    _ = v72
                                                }
                                            }
                                        }
                                    }
                                case 517:
                                    v5 = (float64(v32) + v5)
                                default:
                                    var v104 float32 = float32(((((((v5 / ((((v5 * v5) / v5) * v5) * v5)) * (float64(v32) + v5)) * 3.14159) + v5) - v5) * -2.25))
                                    _ = v104
                                }
                            case "foo bar":
                                var v105 int8 = genFn0()
                                _ = v105
                            default:
                                v55 = (true && v55)
                            }
                            fmt.Println(fi52, fi51)
                            _ = v56
                            _ = v57
                        } else {
                            var v106 bool = v53
                            v4 = "foo bar"
                            _ = v106
                        }
                        v53 = v53
                        if v55 {
                            switch "" {
                            case "\n\t":
                                var v107 uint = 255
                                _ = v107
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = -524
                                }
                            case "hello":
                                v108 := GenS10{F0: v1, F1: -1}
                                _ = v108
                            }
                            v3 = append(v3, "hello")
                        }
                    } else {
                        switch v4 {
                        case "a":
                            var v109 uint = 0
                            _ = v109
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            v31 += (^(v31))
                        }
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                    _ = v55
                }
                switch (v4 + "hello") {
                case "":
                    v110 := GenS4{F0: true, F1: -304, F2: 65535}
                    _ = v110
                }
            }
        }
        v4 = v4
        v4 = (("hello" + "foo bar") + "")
        _ = fi51
    }
    v31 = v31
    if (true && false) {
        _ = len(v3)
    }
    switch 9223372036854775807 {
    case 32767:
        var v111 uint64 = 1
        _ = v111
    }
    var v112 uint64 = 42
    for fi113 := 0; fi113 < 5; fi113++ {
        v32 = float32(v5)
        if false {
            if len(v2) > 0 {
                v2[0] = genFn1(int32(v1), -1)
            }
            v31 = v31
        }
        _ = fi113
    }
    v114 := GenS11{F0: (^(v31))}
    v115 := GenS12{F0: 32767}
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v116 int = 248
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v112
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
