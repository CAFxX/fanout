package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 uint; F1 int8; F2 int8 }
type GenS9 struct { F0 uint; F1 uint32; F2 uint8 }

func main() {
    v1 := []string{""}
    var v2 uint = 42
    v3 := make(map[int]string)
    v4 := GenS1{F0: -2147483648, F1: "hello", F2: 0}
    if (-9223372036854775808 != 274) {
        var v5 uint64 = uint64(v2)
        _ = v5
    } else {
        var v6 int16 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 0, F1: "foo bar", F2: 65535}
        for fi9 := 0; fi9 < 11; fi9++ {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            v10 := GenS1{F0: 2147483647, F1: "\n\t", F2: 332}
            v2 -= v2
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS2{F0: 353, F1: 541, F2: 65535}
    v2 = v2
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v2)
        for fi13 := 0; fi13 < 11; fi13++ {
            v2 = v2
            var v14 float32 = ((0.1 + ((0.0 * 0.0) * 0.0)) / 0.1)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if (!(true || false)) {
        if false {
            _ = len(v3)
            if false {
                var v15 string = ((("hello" + "") + "") + "\n\t")
                _ = v15
            }
            if false {
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS3{F0: 0, F1: 255}
                    for fi18 := 0; fi18 < 19; fi18++ {
                        fmt.Println(fi18, fi18, fi18)
                        _ = fi18
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
    } else {
        if (!false) {
            if (-9223372036854775808 == -740) {
                fmt.Println(v2)
            } else {
                v19 := GenS4{F0: 42, F1: 2}
                fmt.Println(v2)
                _ = v19
            }
            v2 -= v2
            v20 := GenS5{F0: -256}
            _ = v20
        } else {
            var v21 uint32 = 0
            _ = v21
        }
        var v22 uint32 = 1
        switch "" {
        case "hello":
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        default:
            if true {
                v22 -= (v22 ^ (v22 ^ v22))
                var v23 uint16 = uint16(v22)
                switch ((("\n\t" + "\n\t") + "\n\t") + "a") {
                case "":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        _ = len(v3)
                        if (!(3.14159 > ((0.1 + ((0.1 + (float64(-2.25) - (1.5 * (float64(((0.0 * (0.0 / (1.5 * 1.5))) / 1.5)) - 3.14159)))) - (3.14159 / float64(0.1)))) + ((1e10 * float64(((((((0.1 / 1.0) - ((1.5 - -2.25) * (0.1 - 3.14159))) - 1e10) - ((1e10 * 0.0) - 1.5)) * float32(1.5)) / (0.1 - 1e10)))) * (1.5 + 3.14159))))) {
                            v25 := GenS5{F0: 962}
                            var v26 int16 = -281
                            _ = v25
                            _ = v26
                        } else {
                            fmt.Println(v2)
                            var v27 float32 = 1e10
                            _ = v27
                        }
                        fmt.Println(v22, v23)
                        _ = fi24
                    }
                case "\n\t":
                    var v28 uint32 = v22
                    _ = v28
                default:
                    v29 := GenS6{F0: 9223372036854775807, F1: 2147483647, F2: (v22 | 4294967295)}
                    _ = v29
                }
                _ = v23
            } else {
                switch "hello" {
                case "":
                    var v30 float64 = 3.14159
                    _ = v30
                default:
                    switch ("a" + "a") {
                    case "a":
                        switch "hello" {
                        case "":
                            v22 -= v22
                        case "a":
                            var v31 int8 = -1
                            _ = v31
                        case "\n\t":
                            fmt.Println(v22, v22)
                        default:
                            v32 := GenS7{F0: (-641 <= -32768), F1: 32767}
                            _ = v32
                        }
                    case "foo bar":
                        fmt.Println(v2, v22)
                    case "hello":
                        if false {
                            fmt.Println(v2, v2)
                        } else {
                            v22 = (^(v22))
                            var v33 bool = false
                            for fi34 := 0; fi34 < 11; fi34++ {
                                switch "hello" {
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "foo bar":
                                    switch fi34 {
                                    case 2147483647:
                                        v22 += (v22 >> (((v2 - v2)) & 63))
                                    case -1:
                                        for fi35 := 0; fi35 < 19; fi35++ {
                                            if (!false) {
                                                fmt.Println(v33)
                                            }
                                            if v33 {
                                                var v36 uint16 = 0
                                                v22 = v22
                                                var v37 uint64 = 835
                                                _ = v36
                                                _ = v37
                                            }
                                            _ = fi35
                                        }
                                    case 310:
                                        fmt.Println(fi34)
                                    }
                                }
                                switch "hello" {
                                case "":
                                    switch "\n\t" {
                                    case "a":
                                        var v38 float32 = 3.14159
                                        _ = v38
                                    default:
                                        for fi39 := 0; fi39 < 4; fi39++ {
                                            v33 = v33
                                            v40 := GenS2{F0: 42, F1: 4294967295, F2: 32767}
                                            if false {
                                                var v41 uint32 = (v22 / ((v22) | 1))
                                                if ((79 <= 2) || (v33 || v33)) {
                                                    switch "foo bar" {
                                                    case "":
                                                        if (1 < 255) {
                                                            v42 := GenS1{F0: -230, F1: "a", F2: (^(v41))}
                                                            fmt.Println(fi39, v2)
                                                            var v43 int = (fi39 << (((v2 + v2)) & 63))
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    }
                                                    if v33 {
                                                        fmt.Println(v41, v33)
                                                        v44 := GenS1{F0: -2147483648, F1: "", F2: (v41 ^ (^(v22)))}
                                                        v33 = v33
                                                        _ = len(v1)
                                                        _ = v44
                                                    } else {
                                                        var v45 uint = (v2 * v2)
                                                        v46 := GenS8{F0: (v2 ^ (v2 & v2)), F1: 127, F2: -36}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    var v47 uint = v2
                                                    _ = v47
                                                }
                                                var v48 bool = v33
                                                _ = v41
                                                _ = v48
                                            } else {
                                                v49 := GenS5{F0: 552}
                                                if true {
                                                    v1 = append(v1, ("\n\t" + "\n\t"))
                                                    v2 += 1
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (127 >= 127) {
                                                            fmt.Println(fi39)
                                                        }
                                                    case "hello":
                                                        if v33 {
                                                            switch ("" + ("hello" + (((("a" + "") + ("\n\t" + ("foo bar" + "a"))) + ("hello" + "a")) + ("\n\t" + "")))) {
                                                            case "\n\t":
                                                                v50 := GenS5{F0: 532}
                                                                _ = v50
                                                            case "foo bar":
                                                                v1 = append(v1, "a")
                                                            case "a":
                                                                v33 = v33
                                                                v2 += v2
                                                                fmt.Println(v22)
                                                                switch "a" {
                                                                case "hello":
                                                                    fmt.Println(v33, fi39, v2)
                                                                case "\n\t":
                                                                    v33 = v33
                                                                case "foo bar":
                                                                    v51 := GenS6{F0: fi39, F1: -959, F2: (v22 | v22)}
                                                                    _ = v51
                                                                }
                                                            }
                                                        } else {
                                                            var v52 string = "foo bar"
                                                            var v53 int16 = 1000
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    case "foo bar":
                                                        var v54 uint32 = (v22 + v22)
                                                        _ = v54
                                                    default:
                                                        var v55 uint = (^(v2))
                                                        _ = v55
                                                    }
                                                }
                                                v2 = (v2 + v2)
                                                _ = v49
                                            }
                                            _ = fi39
                                            _ = v40
                                        }
                                    }
                                default:
                                    v22 += v22
                                }
                                _ = fi34
                            }
                            _ = v33
                        }
                    default:
                        v1 = append(v1, "hello")
                    }
                }
                if false {
                    var v56 int32 = 54
                    if false {
                        var v57 uint64 = 1
                        if (true && (!((v57 ^ uint64(v2)) == (v57 + (v57 ^ (^(v57))))))) {
                            var v58 int16 = 225
                            _ = v58
                        }
                        _ = v57
                    } else {
                        for fi59 := 0; fi59 < 19; fi59++ {
                            v1 = append(v1, "\n\t")
                            _ = fi59
                        }
                        if (-1 != -477) {
                            var v60 float64 = 3.14159
                            _ = v60
                        }
                        var v61 int64 = 817
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 15; fi62++ {
                        v63 := GenS4{F0: (v22 | 2), F1: v22}
                        _ = fi62
                        _ = v63
                    }
                    _ = v56
                } else {
                    v2 = v2
                    v64 := GenS9{F0: uint(v22), F1: 1, F2: 0}
                    v2 = (v2 * (v2 & 0))
                    _ = v64
                }
            }
        }
        _ = v22
    }
    var v65 uint64 = 18446744073709551615
    var v66 int = 2147483647
    for fi67 := 0; fi67 < 13; fi67++ {
        var v68 int16 = int16(v65)
        var v69 uint16 = 0
        for fi70 := 0; fi70 < 11; fi70++ {
            v1 = append(v1, "a")
            var v71 float64 = -2.25
            _ = fi70
            _ = v71
        }
        _ = fi67
        _ = v68
        _ = v69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
}
