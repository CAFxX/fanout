package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0(p0 uint32, p1 bool) uint64 {
    var v4 int64 = 42
    fmt.Println(p1)
    fmt.Println(v4, p0)
    v4 += v4
    var v5 uint8 = 89
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 600
}

func main() {
    v1 := []int{}
    var v2 uint = 2
    var v3 uint16 = 0
    v3 -= (v3 | (v3 + v3))
    v3 = (v3 >> (((v2 - 0)) & 63))
    v6 := GenS1{F0: -1, F1: "", F2: (^(v2))}
    var v7 uint16 = v3
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v7 = (v7 >> ((v2) & 63))
            var v11 int = (fi10 ^ fi9)
            _ = fi10
            _ = v11
        }
        var v12 int16 = -669
        var v13 uint32 = 4294967295
        _ = fi9
        _ = v12
        _ = v13
    }
    v2 -= 42
    var v14 uint = (v2 << ((v2) & 63))
    v7 = uint16(v8)
    if true {
        var v15 int16 = -32768
        v16 := GenS2{F0: (9223372036854775807 < 32767)}
        _ = v15
        _ = v16
    } else {
        switch "\n\t" {
        case "\n\t":
            var v17 uint32 = 42
            _ = v17
        default:
            fmt.Println(v7, v14)
        }
        switch "\n\t" {
        case "":
            var v18 bool = true
            _ = v18
        default:
            v1 = append(v1, 810)
        }
        v7 += 1
    }
    if (-38 > 1) {
        v19 := GenS2{F0: (false && true)}
        var v20 uint64 = v8
        var v21 float64 = 0.0
        _ = v19
        _ = v20
        _ = v21
    } else {
        v2 = v2
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS1{F0: -515, F1: (("\n\t" + (("foo bar" + (("" + "a") + ("foo bar" + "hello"))) + "\n\t")) + ("foo bar" + "\n\t")), F2: 255}
        if ("\n\t" > "foo bar") {
            _ = len(v1)
            v2 = 136
        } else {
            switch ("foo bar" + "") {
            case "hello":
                if (!true) {
                    switch "" {
                    case "\n\t":
                        switch fi22 {
                        case 32767:
                            var v24 uint64 = v8
                            _ = v24
                        case -1:
                            var v25 string = "\n\t"
                            _ = v25
                        case 687:
                            for fi26 := 0; fi26 < 11; fi26++ {
                                var v27 uint = v2
                                switch "hello" {
                                case "a":
                                    fmt.Println(fi22)
                                case "\n\t":
                                    v28 := GenS3{F0: ""}
                                    _ = v28
                                case "hello":
                                    var v29 uint = 255
                                    _ = v29
                                }
                                if true {
                                    var v30 int64 = -1
                                    fmt.Println(v14, v27, v27)
                                    _ = v30
                                } else {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v2, v2, fi26)
                                    case "\n\t":
                                        v8 += genFn0(42, true)
                                    }
                                    if (65535 > 87) {
                                        v14 -= (v27 % ((uint(v8)) | 1))
                                        var v31 int8 = 0
                                        _ = v31
                                    } else {
                                        v1 = append(v1, fi26)
                                    }
                                    v2 += uint(v7)
                                }
                                if (!(int8(v14) > 127)) {
                                    var v32 uint16 = 756
                                    var v33 int64 = -578
                                    v2 -= v2
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 uint = (v27 + (v2 ^ v27))
                                    if (uint16(v8) <= (v3 | (v7 / (((v7 / ((v3) | 1))) | 1)))) {
                                        switch 32767 {
                                        case -2147483648:
                                            v34 = v34
                                        case 32768:
                                            v35 := GenS2{F0: true}
                                            _ = v35
                                        case -1:
                                            var v36 float32 = 1e10
                                            _ = v36
                                        default:
                                            switch fi26 {
                                            case -2147483648:
                                                v34 += v34
                                            case -675:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 32767:
                                                if true {
                                                    v8 = v8
                                                    v14 = v27
                                                }
                                            }
                                        }
                                    } else {
                                        if false {
                                            switch (fi22 ^ fi26) {
                                            case 4294967295:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 1:
                                                v3 = 255
                                            }
                                        } else {
                                            switch (fi22 & fi26) {
                                            case 231:
                                                if ((42 != uint(v7)) || (false || true)) {
                                                    v8 = 422
                                                    v34 += 159
                                                }
                                            case 42:
                                                var v37 uint32 = 2147483647
                                                _ = v37
                                            case -2147483648:
                                                var v38 float32 = 1e10
                                                _ = v38
                                            default:
                                                switch ((("a" + ("foo bar" + ("hello" + ""))) + ("hello" + ("hello" + "hello"))) + ("foo bar" + (((((("" + "") + (("\n\t" + "\n\t") + "")) + ("foo bar" + "")) + "hello") + "foo bar") + ""))) {
                                                case "":
                                                    _ = len(v1)
                                                case "a":
                                                    var v39 string = "hello"
                                                    _ = v39
                                                case "hello":
                                                    switch (fi22 | (fi22 / (((fi26 % (((fi26 % (((-(fi22))) | 1))) | 1))) | 1))) {
                                                    case 733:
                                                        fmt.Println(v7)
                                                    default:
                                                        var v40 int64 = 976
                                                        _ = v40
                                                    }
                                                default:
                                                    var v41 uint32 = uint32(v34)
                                                    _ = v41
                                                }
                                            }
                                            var v42 uint64 = v8
                                            _ = v42
                                        }
                                        v3 += (v3 & (v3 >> ((v14) & 63)))
                                        switch "foo bar" {
                                        case "a":
                                            v8 = 42
                                        case "":
                                            if (618 >= (fi26 % (((^(fi22))) | 1))) {
                                                switch "foo bar" {
                                                case "a":
                                                    fmt.Println(fi22, v27, fi22)
                                                    switch ("\n\t" + "hello") {
                                                    case "a":
                                                        var v43 uint16 = (v7 / (((^(v3))) | 1))
                                                        _ = v43
                                                    case "hello":
                                                        if true {
                                                            var v44 int32 = -2147483648
                                                            v8 = genFn0(2147483648, (false || true))
                                                            _ = v44
                                                        } else {
                                                            var v45 uint8 = 255
                                                            v8 -= genFn0(2147483648, (false && true))
                                                            _ = v45
                                                        }
                                                    }
                                                    v27 = (v2 << ((v2) & 63))
                                                case "":
                                                    v8 = v8
                                                default:
                                                    v3 += (v7 % ((v7) | 1))
                                                }
                                                var v46 int32 = 42
                                                v14 -= v34
                                                _ = v46
                                            } else {
                                                v47 := GenS3{F0: "hello"}
                                                fmt.Println(v14, v14)
                                                v48 := GenS2{F0: (!((^(v2)) < 2))}
                                                _ = v47
                                                _ = v48
                                            }
                                        case "foo bar":
                                            var v49 int64 = -9223372036854775808
                                            _ = v49
                                        default:
                                            v7 = (^(v7))
                                        }
                                    }
                                    _ = v34
                                }
                                v1 = append(v1, 0)
                                _ = fi26
                                _ = v27
                            }
                        default:
                            var v50 int32 = 1
                            _ = v50
                        }
                    case "a":
                        var v51 uint = v2
                        _ = v51
                    case "hello":
                        if (false || (v3 != 317)) {
                            v3 -= v7
                        } else {
                            fmt.Println(fi22, v8)
                            switch "hello" {
                            case "a":
                                fmt.Println(fi22)
                            case "\n\t":
                                if len(v1) > 0 {
                                    v1[0] = (-(fi22))
                                }
                            default:
                                switch fi22 {
                                case -724:
                                    v1 = append(v1, fi22)
                                }
                            }
                            for fi52 := 0; fi52 < 2; fi52++ {
                                v53 := GenS2{F0: false}
                                v7 += (v7 << (((^(v2))) & 63))
                                switch (fi22 | -687) {
                                case -763:
                                    switch (fi22 / ((32768) | 1)) {
                                    case 32767:
                                        v8 = (v8 * v8)
                                    case 44:
                                        _ = len(v1)
                                    default:
                                        var v54 int16 = 32767
                                        _ = v54
                                    }
                                    if (!false) {
                                        if false {
                                            v8 = genFn0(2147483647, false)
                                            var v55 bool = false
                                            fmt.Println(v3, v8, v7)
                                            _ = v55
                                        } else {
                                            var v56 int8 = 50
                                            v2 += v14
                                            _ = v56
                                        }
                                        v7 = 65535
                                        v2 -= 65535
                                    } else {
                                        var v57 int32 = -416
                                        var v58 int64 = -520
                                        v59 := GenS1{F0: v58, F1: "a", F2: (v2 >> (((v14 % ((v2) | 1))) & 63))}
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                    v60 := GenS4{F0: 911, F1: 172}
                                    if len(v1) > 0 {
                                        v1[0] = 127
                                    }
                                    v3 -= v3
                                    _ = v60
                                case 42:
                                    switch fi52 {
                                    case -693:
                                        v61 := GenS1{F0: 42, F1: ("foo bar" + ""), F2: (v14 / ((18446744073709551615) | 1))}
                                        _ = v61
                                    case -687:
                                        v14 += (^(v14))
                                    case 923:
                                        var v62 uint32 = 2
                                        _ = v62
                                    default:
                                        fmt.Println(fi22, v8)
                                    }
                                case -9223372036854775808:
                                    switch fi22 {
                                    case 618:
                                        v63 := GenS5{F0: 9223372036854775807, F1: uint32(v2), F2: v3}
                                        _ = v63
                                    case -1:
                                        v64 := GenS1{F0: -611, F1: "hello", F2: v14}
                                        _ = v64
                                    case 32768:
                                        var v65 float64 = 0.0
                                        _ = v65
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi52
                                _ = v53
                            }
                        }
                    }
                }
            case "foo bar":
                v66 := GenS2{F0: false}
                _ = v66
            case "\n\t":
                v67 := GenS6{F0: false, F1: true, F2: 42}
                _ = v67
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
