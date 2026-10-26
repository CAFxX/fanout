package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 uint8) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = ("hello" + "a")
        _ = fi7
        _ = v8
    }
    switch ("hello" + "hello") {
    case "a":
        v9 := []int{4294967295, 9223372036854775807}
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{fi10, 2147483647, fi10}
            _ = fi10
            _ = v11
        }
    case "foo bar":
        for fi12 := 0; fi12 < 10; fi12++ {
            if (!(-10 == 2147483647)) {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 string = "foo bar"
                    v15 := []int{int(p0), 128, (fi12 - fi13)}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := []int{}
                _ = v16
            } else {
                if false {
                    var v17 float32 = 0.0
                    v17 += 1e10
                    v18 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: 1}
                    _ = v17
                    _ = v18
                }
            }
            _ = fi12
        }
    }
    fmt.Println(p0)
    for fi19 := 0; fi19 < 1; fi19++ {
        switch "" {
        case "":
            p0 = 0
        case "foo bar":
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := []int{(fi20 & fi19), (fi20 | 1), (fi19 & (fi19 << ((18446744073709551615) & 63))), fi20}
                switch "" {
                case "hello":
                    var v22 int = (fi19 + 389)
                    _ = v22
                default:
                    var v23 uint32 = 65535
                    _ = v23
                }
                if false {
                    var v24 int64 = 42
                    _ = v24
                } else {
                    if (!((!(true && (0 <= (^(p0))))) || false)) {
                        p0 = (p0 / ((1) | 1))
                        var v25 float64 = 1e10
                        _ = v25
                    }
                    var v26 string = "a"
                    _ = len(v21)
                    v26 = "a"
                    if true {
                        var v27 uint8 = p0
                        _ = v27
                    }
                    _ = v26
                }
                _ = fi20
                _ = v21
            }
        default:
            v28 := GenS2{F0: 377, F1: -644}
            _ = v28
        }
        _ = fi19
    }
    if (-79 != 127) {
        if false {
            var v29 uint = 4294967295
            _ = v29
        }
    } else {
        var v30 uint64 = 2
        _ = v30
    }
    _ = p0
    return p0
}

func main() {
    v1 := []uint32{819, 0}
    v2 := []int{}
    v3 := make(map[int]string)
    v4 := []int{-9223372036854775808, -2147483648}
    v5 := make(map[string]bool)
    var v6 int16 = 256
    delete(v5, "foo bar")
    v5["hello"] = false
    v31 := GenS1{F0: 616, F1: 1, F2: 2}
    var v32 int = -2147483648
    switch v32 {
    case 128:
        for fi33 := 0; fi33 < 11; fi33++ {
            if false {
                v6 += v6
                switch "a" {
                case "":
                    v34 := GenS1{F0: -1, F1: 4294967295, F2: 0}
                    _ = v34
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        _ = len(v5)
                    default:
                        v6 -= 162
                    }
                case "foo bar":
                    v6 = v6
                }
                v32 = (-(v32))
            }
            var v35 uint64 = 1
            _ = fi33
            _ = v35
        }
    case 32767:
        var v36 bool = false
        _ = v36
    }
    v32 = v32
    for fi37 := 0; fi37 < 18; fi37++ {
        v6 += v6
        var v38 uint64 = 4294967295
        if true {
            for fi39 := 0; fi39 < 19; fi39++ {
                if (!((false || false) && (!(!false)))) {
                    var v40 int16 = 1
                    if true {
                        var v41 uint16 = 722
                        var v42 uint8 = 0
                        _ = v41
                        _ = v42
                    } else {
                        var v43 uint16 = 291
                        if false {
                            if ((v43 + 255) > v43) {
                                v32 += 1
                                if true {
                                    v44 := GenS1{F0: 0, F1: 255, F2: 255}
                                    v6 += (v40 << ((65535) & 63))
                                    var v45 bool = (1 > -595)
                                    _ = v44
                                    _ = v45
                                } else {
                                    delete(v3, 38)
                                    switch "" {
                                    case "":
                                        v43 = (v43 << ((17) & 63))
                                    }
                                    v46 := GenS2{F0: 330, F1: v40}
                                    v47 := GenS2{F0: v38, F1: 970}
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                var v48 uint64 = (^(v38))
                                v49 := GenS2{F0: v48, F1: -1}
                                _ = v48
                                _ = v49
                            }
                            v50 := GenS3{F0: v43}
                            v6 = v40
                            v6 = (v6 | v40)
                            v51 := GenS2{F0: 0, F1: v40}
                            _ = v50
                            _ = v51
                        } else {
                            switch v32 {
                            case 42:
                                fmt.Println(v43, v43, fi39)
                            case -2147483648:
                                if ((fi37 > fi37) && ((v40 + (-(v40))) < 1000)) {
                                    _ = len(v1)
                                    if len(v2) > 0 {
                                        v2[0] = 127
                                    }
                                    v52 := GenS4{F0: 4294967295, F1: 2147483648}
                                    _ = v52
                                } else {
                                    v32 -= 32768
                                }
                            case -1:
                                var v53 uint = 4294967295
                                _ = v53
                            }
                            v1 = append(v1, 2147483647)
                        }
                        var v54 float64 = (0.1 / 1e10)
                        _ = v43
                        _ = v54
                    }
                    _ = v40
                } else {
                    switch (("a" + "foo bar") + ((("\n\t" + "") + ("foo bar" + ("a" + "foo bar"))) + "\n\t")) {
                    case "foo bar":
                        var v55 float64 = 0.1
                        _ = v55
                    case "":
                        v38 = 2
                    case "hello":
                        if true {
                            var v56 uint16 = uint16(v38)
                            var v57 bool = true
                            fmt.Println(v38)
                            v56 += (^(v56))
                            _ = v56
                            _ = v57
                        } else {
                            if (("\n\t" + "") <= "hello") {
                                switch (fi39 % (((fi39 / ((fi37) | 1))) | 1)) {
                                case 127:
                                    switch (v32 << ((42) & 63)) {
                                    case -805:
                                        fmt.Println(v38, v38)
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            v38 = v38
                                        case "":
                                            v38 = 1
                                        default:
                                            v32 -= 32768
                                        }
                                    }
                                case -1:
                                    v5[("" + "")] = (true || (("foo bar" + ("" + "hello")) >= "hello"))
                                case 128:
                                    if false {
                                        if ((((("" + "") + "") > "hello") && (v38 < (v38 | 415))) && ((1e10 <= (3.14159 - 0.0)) || false)) {
                                            if ("" != "hello") {
                                                if ((!(true || false)) || ((109 > 42) || (float32(-2.25) <= 0.0))) {
                                                    var v58 string = "hello"
                                                    _ = v58
                                                }
                                                v59 := GenS4{F0: 2147483648, F1: 132}
                                                _ = v59
                                            }
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v60 int8 = 100
                                        v32 = (v32 >> ((255) & 63))
                                        _ = v60
                                    }
                                }
                                var v61 int64 = -1
                                _ = v61
                            }
                            v5[("a" + "hello")] = (255 <= 325)
                            switch (v32 ^ 0) {
                            case 289:
                                v62 := GenS4{F0: 805, F1: 65535}
                                _ = v62
                            }
                        }
                    }
                }
                _ = fi39
            }
        } else {
            for fi63 := 0; fi63 < 4; fi63++ {
                var v64 int = fi37
                var v65 uint16 = 0
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        _ = fi37
        _ = v38
    }
    if true {
        if (!(true || (338 > 641))) {
            fmt.Println(v6)
        } else {
            var v66 float32 = 3.14159
            switch (-(v32)) {
            case -9223372036854775808:
                v66 = v66
            case 2:
                v66 = (3.14159 * v66)
            default:
                fmt.Println(v66)
            }
            switch v32 {
            case 32768:
                if true {
                    v67 := GenS4{F0: 42, F1: 2}
                    for fi68 := 0; fi68 < 18; fi68++ {
                        var v69 int64 = 0
                        if (1 == 0) {
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = fi68
                                }
                                v70 := GenS4{F0: 1, F1: 407}
                                _ = v70
                            } else {
                                v71 := GenS5{F0: 640}
                                _ = v71
                            }
                        }
                        for fi72 := 0; fi72 < 13; fi72++ {
                            var v73 int64 = (v69 - (v69 % (((-(v69))) | 1)))
                            if false {
                                var v74 uint16 = 0
                                var v75 int = 291
                                var v76 uint16 = (v74 % ((v74) | 1))
                                switch (fi68 << ((42) & 63)) {
                                case 32767:
                                    delete(v3, v75)
                                }
                                _ = v74
                                _ = v75
                                _ = v76
                            } else {
                                var v77 uint8 = 255
                                var v78 bool = (("foo bar" + "") != ("foo bar" + "\n\t"))
                                v79 := GenS6{F0: 44, F1: 1, F2: 1}
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            if true {
                                v80 := GenS4{F0: 646, F1: 2147483647}
                                fmt.Println(v66, fi72)
                                var v81 int16 = 42
                                _ = v80
                                _ = v81
                            } else {
                                switch "\n\t" {
                                case "":
                                    if (114 < 42) {
                                        var v82 float32 = ((v66 * v66) - v66)
                                        v83 := GenS7{F0: -989, F1: -108}
                                        _ = v82
                                        _ = v83
                                    } else {
                                        var v84 int = fi68
                                        v85 := GenS2{F0: 2, F1: v6}
                                        if (0.0 != 0.0) {
                                            v86 := GenS8{F0: 840, F1: 70, F2: genFn0(2)}
                                            _ = v86
                                        } else {
                                            if (!(0 < 236)) {
                                                v87 := GenS8{F0: 1, F1: 4294967295, F2: 2}
                                                v88 := GenS9{F0: "a", F1: 2}
                                                switch v32 {
                                                case -9223372036854775808:
                                                    v73 += v69
                                                }
                                                _ = v87
                                                _ = v88
                                            }
                                            v4 = append(v4, 42)
                                        }
                                        _ = v84
                                        _ = v85
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch ("a" + "foo bar") {
                                    case "\n\t":
                                        var v89 bool = (!true)
                                        switch (v32 / ((fi68) | 1)) {
                                        case -9223372036854775808:
                                            v89 = (!(true || v89))
                                        case -258:
                                            var v90 float32 = v66
                                            v91 := GenS10{F0: 2}
                                            var v92 uint16 = 42
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        case -2147483648:
                                            v69 -= (v73 >> ((802) & 63))
                                        }
                                        v73 -= v73
                                        v73 -= v73
                                        if true {
                                            _ = len(v5)
                                            v93 := GenS5{F0: 401}
                                            v6 += (-(v6))
                                            _ = len(v5)
                                            _ = v93
                                        }
                                        v94 := GenS4{F0: 1, F1: 487}
                                        switch (fi72 - (fi68 >> ((4294967295) & 63))) {
                                        case -512:
                                            fmt.Println(v32, fi72)
                                        case 750:
                                            var v95 float32 = (3.14159 * (v66 * v66))
                                            _ = v95
                                        case 4294967295:
                                            v96 := GenS5{F0: 65535}
                                            _ = v96
                                        default:
                                            var v97 int16 = v6
                                            _ = v97
                                        }
                                        _ = v89
                                        _ = v94
                                    case "a":
                                        var v98 uint16 = 32767
                                        _ = v98
                                    case "foo bar":
                                        fmt.Println(v32, v66)
                                        v5["a"] = (-2147483648 != -2147483648)
                                    }
                                case "a":
                                    v32 = fi72
                                }
                                fmt.Println(v69, fi68)
                                v73 -= -52
                                if false {
                                    var v99 uint = 4294967295
                                    _ = v99
                                }
                                var v100 bool = (1 != 42)
                                fmt.Println(v66)
                                _ = v100
                            }
                            _ = fi72
                            _ = v73
                        }
                        _ = fi68
                        _ = v69
                    }
                    _ = v67
                }
            case 42:
                delete(v5, "foo bar")
            case 127:
                fmt.Println(v66)
            }
            _ = v66
        }
        var v101 uint = 255
        var v102 float32 = 0.0
        _ = v101
        _ = v102
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v6 = (^(v6))
    }
    var v103 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
