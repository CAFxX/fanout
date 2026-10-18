package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int }

func main() {
    var v1 uint32 = 1
    var v2 int16 = 1000
    var v3 int64 = 9223372036854775807
    var v4 uint16 = 32767
    v5 := make(map[string]uint32)
    var v6 float64 = 3.14159
    fmt.Println(v4)
    v1 += v1
    switch -2147483648 {
    case 0:
        v6 += 1.5
    }
    v3 += -28
    v7 := []int{-157, 32767}
    v4 -= 1
    if (false && (88 != 1)) {
        switch 1 {
        case -9223372036854775808:
            _ = len(v7)
        case 653:
            if (857 <= 42) {
                if ((v1 - v1) >= (^(v1))) {
                    v3 = (v3 | (-(v3)))
                    if len(v7) > 0 {
                        v7[0] = 195
                    }
                }
                var v8 float64 = v6
                _ = v8
            } else {
                for fi9 := 0; fi9 < 8; fi9++ {
                    switch fi9 {
                    case 42:
                        var v10 uint = 0
                        _ = v10
                    case 2147483647:
                        v1 = v1
                    case -890:
                        _ = len(v5)
                    }
                    v3 += (v3 << ((42) & 63))
                    v4 = 65535
                    _ = fi9
                }
                var v11 uint16 = v4
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            if true {
                var v13 uint = 0
                _ = v13
            } else {
                var v14 uint64 = 4294967295
                var v15 int64 = 9223372036854775807
                _ = v14
                _ = v15
            }
            var v16 string = ""
            if false {
                var v17 bool = ((v6 >= v6) || false)
                v18 := GenS1{F0: (-(v3))}
                _ = v17
                _ = v18
            }
            _ = fi12
            _ = v16
        }
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS2{F0: "foo bar", F1: uint16(v1)}
            for fi21 := 0; fi21 < 16; fi21++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v3 = (v3 - v3)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        if ((-(v2)) >= (v2 << ((65535) & 63))) {
            var v22 int64 = v3
            for fi23 := 0; fi23 < 13; fi23++ {
                switch -9223372036854775808 {
                case 202:
                    v5[(("a" + "") + ("foo bar" + ("" + "\n\t")))] = 42
                case 127:
                    if (1 < 2) {
                        var v24 int16 = v2
                        for fi25 := 0; fi25 < 5; fi25++ {
                            var v26 int = (fi23 & -9223372036854775808)
                            _ = fi25
                            _ = v26
                        }
                        v6 -= v6
                        _ = v24
                    }
                case 298:
                    var v27 uint = 65535
                    _ = v27
                }
                var v28 float32 = (((((3.14159 - 0.0) / (3.14159 + float32((0.1 / v6)))) + 0.1) / (0.1 + 0.1)) * 0.0)
                switch (fi23 + fi23) {
                case 9223372036854775807:
                    v6 = v6
                default:
                    if false {
                        var v29 int64 = v22
                        switch (fi23 >> ((2) & 63)) {
                        case 9223372036854775807:
                            for fi30 := 0; fi30 < 4; fi30++ {
                                v1 -= v1
                                v22 += v29
                                switch fi30 {
                                case 0:
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            v2 = (-(v2))
                                            var v31 uint16 = (v4 & v4)
                                            _ = v31
                                        }
                                    default:
                                        switch ("hello" + ("\n\t" + "a")) {
                                        case "a":
                                            v5[""] = 4294967295
                                        case "\n\t":
                                            v32 := GenS1{F0: (^(v3))}
                                            _ = v32
                                        case "":
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        }
                                    }
                                case 127:
                                    v5["\n\t"] = v1
                                default:
                                    var v33 string = (("foo bar" + "") + "")
                                    _ = v33
                                }
                                _ = fi30
                            }
                        case -101:
                            for fi34 := 0; fi34 < 4; fi34++ {
                                if true {
                                    _ = len(v7)
                                    fmt.Println(fi23, fi34)
                                }
                                v1 = (^(v1))
                                _ = fi34
                            }
                        case 0:
                            v29 += (-(v29))
                        default:
                            if false {
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    if (369 > 0) {
                                        v36 := GenS3{F0: 127}
                                        if (-2147483648 != fi35) {
                                            v4 += v4
                                        }
                                        if false {
                                            if (!true) {
                                                v1 = 65535
                                                var v37 bool = true
                                                _ = v37
                                            } else {
                                                v29 = v22
                                                if true {
                                                    v38 := GenS2{F0: "", F1: (v4 | v4)}
                                                    _ = v38
                                                } else {
                                                    v7 = append(v7, fi23)
                                                    v39 := GenS2{F0: ("a" + ""), F1: (^(v4))}
                                                    _ = v39
                                                }
                                                v5["hello"] = 65535
                                                var v40 int8 = 100
                                                _ = v40
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v41 int = int(v22)
                                                _ = v41
                                            case "a":
                                                fmt.Println(v2)
                                            case "\n\t":
                                                fmt.Println(fi23, v2)
                                            }
                                        }
                                        _ = v36
                                    } else {
                                        delete(v5, (("a" + "") + ("\n\t" + ("a" + ""))))
                                    }
                                    _ = fi35
                                }
                                fmt.Println(v28, v22, v1)
                                for fi42 := 0; fi42 < 3; fi42++ {
                                    v22 = 408
                                    switch "\n\t" {
                                    case "\n\t":
                                        v43 := GenS4{F0: ((^(v1)) <= 841)}
                                        _ = v43
                                    default:
                                        if ((!(!false)) || true) {
                                            delete(v5, "foo bar")
                                            _ = len(v7)
                                            v22 -= 230
                                            var v44 uint16 = v4
                                            _ = v44
                                        } else {
                                            var v45 uint8 = 75
                                            v1 += (v1 >> ((203) & 63))
                                            _ = v45
                                        }
                                    }
                                    switch (fi42 & fi23) {
                                    case 128:
                                        v28 = float32(1.5)
                                    case 257:
                                        switch 4294967295 {
                                        case 32768:
                                            v46 := v5["hello"]
                                            _ = v46
                                        case 9223372036854775807:
                                            if (!true) {
                                                _ = len(v5)
                                                var v47 uint32 = 2
                                                _ = v47
                                            } else {
                                                v29 += v29
                                                var v48 int8 = 95
                                                _ = v48
                                            }
                                        case -424:
                                            switch "" {
                                            case "a":
                                                if ((!(false && false)) || (true || true)) {
                                                    v3 = v22
                                                }
                                            case "hello":
                                                if false {
                                                    var v49 uint32 = v1
                                                    v22 = v29
                                                    var v50 float32 = -2.25
                                                    _ = v49
                                                    _ = v50
                                                }
                                            case "":
                                                v6 -= v6
                                            }
                                            v28 = (v28 * ((v28 - 1.5) - 3.14159))
                                            var v51 int = fi23
                                            _ = v51
                                        default:
                                            v7 = append(v7, 2147483647)
                                        }
                                    }
                                    var v52 bool = true
                                    var v53 int32 = int32(v3)
                                    _ = fi42
                                    _ = v52
                                    _ = v53
                                }
                            } else {
                                _ = len(v5)
                                for fi54 := 0; fi54 < 2; fi54++ {
                                    fmt.Println(v1)
                                    v6 = v6
                                    _ = fi54
                                }
                            }
                        }
                        _ = v29
                    }
                }
                _ = fi23
                _ = v28
            }
            _ = len(v7)
            _ = v22
        }
        var v55 int64 = 0
        _ = v55
    }
    if true {
        for fi56 := 0; fi56 < 8; fi56++ {
            v57 := GenS4{F0: true}
            var v58 uint8 = 57
            _ = fi56
            _ = v57
            _ = v58
        }
        v59 := GenS5{F0: uint8(v4), F1: 176}
        var v60 bool = (v3 <= v3)
        _ = v59
        _ = v60
    }
    v3 = (v3 >> ((619) & 63))
    switch -2147483648 {
    case 687:
        switch -128 {
        case -246:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v3, v2)
            case "a":
                switch -9223372036854775808 {
                case 0:
                    switch "a" {
                    case "\n\t":
                        for fi61 := 0; fi61 < 8; fi61++ {
                            v2 += int16(v3)
                            fmt.Println(v1)
                            _ = fi61
                        }
                    case "a":
                        var v62 bool = (!(9223372036854775807 <= v3))
                        _ = v62
                    case "hello":
                        v5["\n\t"] = (v1 - v1)
                    }
                case -763:
                    v3 += v3
                case 127:
                    switch ("" + "") {
                    case "\n\t":
                        switch "\n\t" {
                        case "\n\t":
                            var v63 float32 = (3.14159 + 1.5)
                            _ = v63
                        case "hello":
                            if false {
                                v64 := v5["hello"]
                                if true {
                                    if (((0.1 / 1.5) >= 1e10) && true) {
                                        var v65 uint = 42
                                        _ = v65
                                    }
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        fmt.Println(fi66, v2)
                                        v2 += (v2 - v2)
                                        v1 = (^(v64))
                                        _ = fi66
                                    }
                                } else {
                                    var v67 uint16 = v4
                                    if (!true) {
                                        if len(v7) > 0 {
                                            v7[0] = 307
                                        }
                                    }
                                    _ = v67
                                }
                                _ = v64
                            } else {
                                v2 = v2
                                switch "a" {
                                case "\n\t":
                                    if len(v7) > 0 {
                                        v7[0] = 59
                                    }
                                case "a":
                                    var v68 float32 = -2.25
                                    _ = v68
                                }
                            }
                        case "a":
                            var v69 uint8 = 7
                            _ = v69
                        default:
                            v1 = v1
                        }
                    case "":
                        var v70 uint8 = 128
                        _ = v70
                    case "a":
                        var v71 int16 = v2
                        _ = v71
                    }
                default:
                    for fi72 := 0; fi72 < 1; fi72++ {
                        for fi73 := 0; fi73 < 13; fi73++ {
                            _ = len(v5)
                            var v74 int32 = 881
                            if true {
                                fmt.Println(v6)
                                v4 = (v4 & (v4 >> ((2) & 63)))
                                if (!false) {
                                    switch 4294967295 {
                                    case 174:
                                        var v75 uint8 = uint8(v4)
                                        _ = v75
                                    case 127:
                                        if false {
                                            var v76 int16 = (v2 - (-(v2)))
                                            v2 = (^(v2))
                                            _ = v76
                                        }
                                    default:
                                        v1 = v1
                                    }
                                    v3 = 1
                                    v77 := GenS3{F0: int8(v2)}
                                    _ = v77
                                } else {
                                    v78 := GenS4{F0: (!(!(float32((1.5 * 1e10)) >= (float32((3.14159 - v6)) * float32(v6)))))}
                                    fmt.Println(v4, v3)
                                    _ = v78
                                }
                                v4 = 255
                            } else {
                                v2 = (v2 ^ (^(v2)))
                            }
                            _ = fi73
                            _ = v74
                        }
                        var v79 uint = 2
                        _ = fi72
                        _ = v79
                    }
                }
            }
        case 32767:
            v4 -= (v4 * v4)
        case 42:
            for fi80 := 0; fi80 < 15; fi80++ {
                if true {
                    v1 += v1
                    var v81 bool = (int64(v1) <= int64(v4))
                    v1 = (^(v1))
                    _ = v81
                }
                v1 = (v1 ^ (^(v1)))
                _ = fi80
            }
        }
    case -1:
        var v82 string = "\n\t"
        _ = v82
    case 1:
        _ = len(v7)
    default:
        var v83 int = -1
        _ = v83
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
