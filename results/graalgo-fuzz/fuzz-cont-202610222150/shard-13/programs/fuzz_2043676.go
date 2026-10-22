package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 string; F2 uint64 }
type GenS7 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 bool) uint16 {
    v6 := []int{32767}
    p0 = true
    _ = p0
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 519
    var v3 uint8 = 1
    v4 := []string{"foo bar", (("" + "a") + "a"), "hello"}
    var v5 float32 = (((1.5 / -2.25) * (-2.25 / (0.1 - (-2.25 - -2.25)))) * float32(1.5))
    _ = len(v1)
    var v7 uint64 = 2
    var v8 int = -1
    var v9 uint64 = uint64(v3)
    switch 353 {
    case -1:
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: (!true), F1: uint8(v2), F2: 2}
            if (2 < 42) {
                switch "hello" {
                case "hello":
                    var v12 int8 = 0
                    _ = v12
                default:
                    v13 := GenS1{F0: true, F1: v3, F2: 410}
                    _ = v13
                }
                var v14 float32 = 1e10
                _ = v14
            } else {
                v15 := GenS2{F0: -490, F1: 128}
                _ = v15
            }
            switch fi10 {
            case 32767:
                v2 = v2
            case -959:
                switch v8 {
                case 107:
                    if false {
                        fmt.Println(v3)
                    }
                case -2147483648:
                    v16 := GenS3{F0: 1, F1: ((^(v2)) >= (v2 * v2)), F2: -30}
                    _ = v16
                default:
                    v17 := GenS1{F0: true, F1: 62, F2: 1}
                    _ = v17
                }
            case -9223372036854775808:
                if ((1000 >= 99) && false) {
                    switch (v8 << ((865) & 63)) {
                    case 1:
                        v7 = (v9 ^ (^(v7)))
                    case 42:
                        var v18 uint16 = v2
                        _ = v18
                    case -875:
                        fmt.Println(fi10, v7)
                    }
                    if (-1 == 100) {
                        if (-9223372036854775808 != (v8 & (-(fi10)))) {
                            for fi19 := 0; fi19 < 17; fi19++ {
                                switch (fi19 - (^(v8))) {
                                case 1:
                                    var v20 uint64 = v9
                                    _ = v20
                                case 128:
                                    switch v8 {
                                    case 368:
                                        switch (fi10 >> ((255) & 63)) {
                                        case 0:
                                            var v21 string = (("a" + (((("\n\t" + ("" + "")) + "a") + "foo bar") + ("" + "foo bar"))) + "hello")
                                            _ = v21
                                        }
                                    case 1:
                                        fmt.Println(v2, v7, v9)
                                    case 32768:
                                        if (v3 == (v3 / ((128) | 1))) {
                                            var v22 float32 = v5
                                            v23 := GenS3{F0: 9223372036854775807, F1: ((-32768 > 1) && (false && (!(-2.25 < 1.5)))), F2: 1}
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                case 32768:
                                    if true {
                                        fmt.Println(v7)
                                        _ = len(v4)
                                    }
                                }
                                _ = fi19
                            }
                        }
                        fmt.Println(fi10, fi10, v2)
                        var v24 uint64 = (v7 << ((0) & 63))
                        _ = v24
                    } else {
                        var v25 uint8 = (v3 << ((1) & 63))
                        switch "hello" {
                        case "hello":
                            var v26 string = ("\n\t" + ("foo bar" + "\n\t"))
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 8; fi27++ {
                                if false {
                                    if (false || (int8(v2) > -128)) {
                                        var v28 int = -951
                                        _ = v28
                                    } else {
                                        fmt.Println(v7, fi10)
                                        v4 = append(v4, "a")
                                    }
                                    v2 = v2
                                } else {
                                    v29 := GenS4{F0: 1}
                                    v9 += v9
                                    var v30 int64 = 1
                                    if (!(!(42 < 0))) {
                                        var v31 uint64 = (v9 / ((1) | 1))
                                        _ = v31
                                    } else {
                                        var v32 int32 = 689
                                        v1[v8] = (true || (false || (-32768 >= 42)))
                                        v3 -= 0
                                        _ = v32
                                    }
                                    if (2 >= 65535) {
                                        if false {
                                            v33 := GenS1{F0: false, F1: v3, F2: 2}
                                            v7 = v9
                                            _ = v33
                                        }
                                        if (2147483647 != 2) {
                                            v5 -= 3.14159
                                        }
                                        switch (fi10 % ((fi10) | 1)) {
                                        case 0:
                                            var v34 uint8 = (v3 | (v25 << ((0) & 63)))
                                            _ = v34
                                        case 281:
                                            delete(v1, (fi10 & (fi10 ^ 9223372036854775807)))
                                        case -1:
                                            var v35 string = ("a" + "a")
                                            _ = v35
                                        default:
                                            if (!true) {
                                                v1[(v8 - 0)] = (v5 <= v5)
                                                _ = len(v1)
                                                fmt.Println(fi27)
                                            }
                                        }
                                    } else {
                                        var v36 int32 = 771
                                        var v37 string = ""
                                        fmt.Println(fi10, v36)
                                        _ = v36
                                        _ = v37
                                    }
                                    v2 -= 42
                                    _ = v29
                                    _ = v30
                                }
                                switch 460 {
                                case 42:
                                    var v38 uint16 = v2
                                    _ = v38
                                case 4294967295:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v3 = v25
                                    case "hello":
                                        var v39 uint64 = v9
                                        _ = v39
                                    }
                                case -46:
                                    var v40 int8 = -1
                                    fmt.Println(v3, fi27, v9)
                                    var v41 uint32 = 1
                                    _ = v40
                                    _ = v41
                                default:
                                    delete(v1, (v8 >> ((65535) & 63)))
                                    var v42 float32 = v5
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v25 -= 42
                                        case "foo bar":
                                            var v43 string = ("hello" + ("\n\t" + ""))
                                            v7 -= uint64(v2)
                                            _ = v43
                                        }
                                        fmt.Println(fi27, v7, fi10)
                                        var v44 int16 = -1
                                        _ = v44
                                    } else {
                                        switch int(v2) {
                                        case 127:
                                            if true {
                                                if (!true) {
                                                    fmt.Println(v5, v2)
                                                    v9 = (^(v7))
                                                    var v45 uint32 = uint32(v2)
                                                    fmt.Println(v2)
                                                    v46 := GenS1{F0: (uint16(v45) <= (^(v2))), F1: v25, F2: (v45 * v45)}
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v47 := GenS5{F0: false}
                                                    _ = v47
                                                }
                                                var v48 int32 = -2147483648
                                                var v49 int64 = 1
                                                _ = v48
                                                _ = v49
                                            } else {
                                                _ = len(v4)
                                                var v50 string = ("\n\t" + (("\n\t" + ("foo bar" + "hello")) + (("foo bar" + "") + "hello")))
                                                fmt.Println(v9, v50)
                                                var v51 uint8 = 235
                                                v1[(fi27 ^ fi10)] = ((1e10 > v42) && (((^(v2)) >= (v2 + genFn0(true))) && (-1 == 1000)))
                                                _ = v50
                                                _ = v51
                                            }
                                        }
                                    }
                                    _ = v42
                                }
                                _ = fi27
                            }
                        }
                        _ = v25
                    }
                }
            }
            _ = fi10
            _ = v11
        }
    default:
        fmt.Println(v7)
    }
    for fi52 := 0; fi52 < 8; fi52++ {
        v3 = (v3 ^ v3)
        var v53 uint16 = 32767
        v8 = (v8 + (fi52 * 9223372036854775807))
        _ = fi52
        _ = v53
    }
    v54 := GenS6{F0: 1, F1: "foo bar", F2: (v7 + v9)}
    switch (v8 << ((1) & 63)) {
    case -1:
        _ = len(v1)
    case 9223372036854775807:
        switch v8 {
        case 2147483647:
            v55 := GenS7{F0: 42, F1: 180, F2: (v8 >> ((4294967295) & 63))}
            _ = v55
        default:
            fmt.Println(v9, v9)
        }
    case 2147483647:
        switch int(v9) {
        case 64:
            if true {
                switch (("" + "foo bar") + "") {
                case "a":
                    fmt.Println(v3, v5)
                case "":
                    if (!true) {
                        switch v8 {
                        case -2147483648:
                            if len(v4) > 0 {
                                v4[0] = ("\n\t" + ("\n\t" + "\n\t"))
                            }
                        case 9223372036854775807:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v2 = 975
                    }
                case "foo bar":
                    v9 -= uint64(v2)
                default:
                    for fi56 := 0; fi56 < 14; fi56++ {
                        switch ((("foo bar" + "hello") + "\n\t") + ("" + "\n\t")) {
                        case "\n\t":
                            v57 := GenS4{F0: 1}
                            _ = v57
                        case "foo bar":
                            if (90 != v3) {
                                if (0 >= 65535) {
                                    var v58 float64 = float64((v5 - 1e10))
                                    _ = v58
                                } else {
                                    var v59 int = fi56
                                    _ = v59
                                }
                            } else {
                                if true {
                                    if (false || true) {
                                        if true {
                                            v2 -= genFn0(true)
                                            var v60 int32 = -2147483648
                                            _ = v60
                                        }
                                    } else {
                                        var v61 int8 = -108
                                        var v62 bool = (412 == 1)
                                        _ = v61
                                        _ = v62
                                    }
                                    v5 = float32((0.1 - 0.0))
                                    var v63 bool = false
                                    _ = v63
                                } else {
                                    if false {
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi64 := 0; fi64 < 8; fi64++ {
                                                fmt.Println(fi56, v8, v8)
                                                if ("hello" < ("hello" + "foo bar")) {
                                                    var v65 bool = true
                                                    _ = v65
                                                }
                                                _ = fi64
                                            }
                                        case "a":
                                            if ((v9 | 18446744073709551615) < 544) {
                                                v5 -= v5
                                                _ = len(v4)
                                                for fi66 := 0; fi66 < 13; fi66++ {
                                                    fmt.Println(v2, v7, v8)
                                                    v7 -= (^(v7))
                                                    _ = fi66
                                                }
                                            }
                                        case "foo bar":
                                            fmt.Println(fi56, v2, v8)
                                        }
                                        v2 = v2
                                        v3 = v3
                                    }
                                    v67 := GenS2{F0: int32(v7), F1: (v3 - 85)}
                                    for fi68 := 0; fi68 < 9; fi68++ {
                                        if false {
                                            var v69 int = (fi68 | fi68)
                                            _ = v69
                                        } else {
                                            if ((!(-775 == 1)) || true) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        }
                                        if (fi56 != 127) {
                                            var v70 int = fi56
                                            _ = v70
                                        } else {
                                            v71 := GenS1{F0: true, F1: v3, F2: 2147483648}
                                            if true {
                                                v9 += (v9 % ((v7) | 1))
                                                v72 := v1[42]
                                                var v73 float32 = v5
                                                _ = v72
                                                _ = v73
                                            } else {
                                                fmt.Println(fi56, fi68)
                                                var v74 int = 0
                                                _ = v74
                                            }
                                            _ = v71
                                        }
                                        var v75 int32 = 1
                                        v2 += (^(v2))
                                        _ = fi68
                                        _ = v75
                                    }
                                    _ = v67
                                }
                                v8 -= (v8 ^ (v8 % ((128) | 1)))
                            }
                        default:
                            for fi76 := 0; fi76 < 20; fi76++ {
                                switch v8 {
                                case 2147483647:
                                    if (-128 != 1) {
                                        switch (fi56 << ((171) & 63)) {
                                        case 128:
                                            switch (fi76 ^ (^(fi56))) {
                                            case -771:
                                                v8 = v8
                                            case 32767:
                                                fmt.Println(v2, v7, fi76)
                                                switch "foo bar" {
                                                case "hello":
                                                    var v77 uint64 = (v9 * 42)
                                                    _ = v77
                                                case "foo bar":
                                                    var v78 uint16 = uint16(v7)
                                                    _ = v78
                                                case "\n\t":
                                                    v2 -= v2
                                                }
                                                if true {
                                                    v1[(fi76 / ((127) | 1))] = ((!((false && (!true)) || (0 <= 65535))) || true)
                                                    fmt.Println(v3, v8, v3)
                                                    delete(v1, v8)
                                                    if false {
                                                        if ((v5 < ((float32((0.0 + (float64(v5) / float64(v5)))) - (float32(float64(v5)) + ((((v5 - (1e10 * 3.14159)) * 0.1) + 1e10) / (v5 / (v5 * v5))))) + (v5 + 3.14159))) && false) {
                                                            switch "hello" {
                                                            case "hello":
                                                                v79 := GenS2{F0: -2147483648, F1: 42}
                                                                _ = v79
                                                            case "foo bar":
                                                                var v80 int = (-(fi76))
                                                                _ = v80
                                                            }
                                                            v81 := GenS7{F0: -128, F1: -32768, F2: (fi76 & v8)}
                                                            var v82 uint16 = genFn0(true)
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    } else {
                                                        if true {
                                                            v3 = 127
                                                            var v83 float64 = 0.0
                                                            fmt.Println(v5, fi56, v5)
                                                            _ = v83
                                                        } else {
                                                            var v84 int8 = -128
                                                            var v85 uint16 = v2
                                                            v86 := GenS7{F0: v84, F1: 1, F2: (v8 - fi76)}
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        switch (v8 / ((fi56) | 1)) {
                                                        case 4294967295:
                                                            fmt.Println(v7, v3)
                                                        default:
                                                            v5 = -2.25
                                                        }
                                                        var v87 uint64 = v9
                                                        var v88 float32 = v5
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                }
                                            }
                                        default:
                                            var v89 float32 = v5
                                            _ = v89
                                        }
                                        if false {
                                            fmt.Println(v5, fi56, v8)
                                            var v90 int = fi56
                                            _ = v90
                                        } else {
                                            v4 = append(v4, "a")
                                            v5 = 1.5
                                        }
                                    } else {
                                        var v91 bool = ((!((128 >= 152) && true)) && (!true))
                                        var v92 bool = ((v2 >> ((316) & 63)) != v2)
                                        _ = v91
                                        _ = v92
                                    }
                                }
                                v2 -= uint16(v3)
                                var v93 string = ("\n\t" + ("\n\t" + "\n\t"))
                                _ = fi76
                                _ = v93
                            }
                        }
                        v94 := GenS4{F0: -310}
                        var v95 int16 = 32767
                        _ = fi56
                        _ = v94
                        _ = v95
                    }
                }
            } else {
                v4 = append(v4, "a")
            }
        case 42:
            for fi96 := 0; fi96 < 15; fi96++ {
                v2 = v2
                v8 = v8
                _ = fi96
            }
        case 350:
            v1[v8] = (!(!false))
        default:
            v97 := GenS1{F0: (!true), F1: v3, F2: 4294967295}
            _ = v97
        }
    }
    v7 += v9
    if true {
        var v98 uint32 = 982
        var v99 int64 = -9223372036854775808
        _ = v98
        _ = v99
    }
    var v100 int32 = 386
    fmt.Println(v5, v9)
    v100 -= v100
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v7
    _ = v8
    _ = v9
}
