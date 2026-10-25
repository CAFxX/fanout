package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int64; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0() uint32 {
    var v5 int16 = 994
    v5 = v5
    v6 := GenS1{F0: 793, F1: -2147483648, F2: 1}
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 uint32 = 303
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "hello", "foo bar"}
    var v4 int32 = 0
    switch "a" {
    case "hello":
        if (true && ((3.14159 <= 3.14159) || (42 > 202))) {
            for fi7 := 0; fi7 < 16; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    if (true || (!true)) {
                        v4 -= v4
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        v4 += v4
                    }
                    var v9 int64 = 42
                    _ = fi8
                    _ = v9
                }
                switch fi7 {
                case -1:
                    switch ("a" + "") {
                    case "":
                        v1 = genFn0()
                    case "hello":
                        if ((((3.14159 / ((float32(0.1) + 1e10) / 3.14159)) - float32(3.14159)) / 1.0) >= float32((float64(1.5) * 1e10))) {
                            for fi10 := 0; fi10 < 20; fi10++ {
                                v11 := GenS2{F0: 2147483647, F1: 230, F2: 42}
                                v3 = append(v3, "foo bar")
                                _ = fi10
                                _ = v11
                            }
                        }
                    case "a":
                        var v12 int = fi7
                        _ = v12
                    default:
                        v13 := GenS2{F0: fi7, F1: 0, F2: 877}
                        _ = v13
                    }
                case 32767:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS3{F0: (v4 | v4), F1: 686, F2: -56}
                        _ = len(v2)
                        v1 -= (v1 + v1)
                        _ = fi14
                        _ = v15
                    }
                case 9223372036854775807:
                    delete(v2, (("a" + "a") + ""))
                }
                v1 = (v1 + v1)
                _ = fi7
            }
            fmt.Println(v1)
            v1 -= genFn0()
        } else {
            if true {
                var v16 float32 = -2.25
                _ = v16
            }
            delete(v2, "\n\t")
            var v17 bool = (127 >= -37)
            _ = v17
        }
    default:
        delete(v2, "")
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 string = ("\n\t" + "\n\t")
        if ((fi18 / ((fi18) | 1)) > fi18) {
            var v20 string = "foo bar"
            v19 = "foo bar"
            fmt.Println(fi18, v1)
            _ = v20
        } else {
            var v21 float64 = 1.5
            var v22 uint16 = 255
            _ = v21
            _ = v22
        }
        v23 := GenS3{F0: (v4 << ((uint(v1)) & 63)), F1: uint(v1), F2: 100}
        _ = fi18
        _ = v19
        _ = v23
    }
    var v24 int = 329
    var v25 uint32 = genFn0()
    var v26 string = "foo bar"
    if (genFn0() != genFn0()) {
        var v27 int8 = -1
        var v28 float64 = 0.0
        _ = v27
        _ = v28
    } else {
        if (0 <= -786) {
            for fi29 := 0; fi29 < 6; fi29++ {
                switch "foo bar" {
                case "\n\t":
                    v30 := GenS1{F0: 1, F1: (v4 / ((v4) | 1)), F2: 2}
                    _ = v30
                default:
                    v31 := GenS3{F0: 1, F1: 496, F2: -128}
                    _ = v31
                }
                if ((v25 - genFn0()) > v1) {
                    _ = len(v2)
                    var v32 uint16 = 255
                    _ = v32
                }
                _ = fi29
            }
            v33 := GenS4{F0: 121}
            _ = v33
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            v1 -= v1
            _ = fi34
        }
    }
    v26 = "\n\t"
    var v35 string = v26
    switch "a" {
    case "":
        v36 := GenS2{F0: 9223372036854775807, F1: -484, F2: 42}
        _ = v36
    case "hello":
        switch v35 {
        case "\n\t":
            var v37 float64 = (3.14159 / 3.14159)
            _ = v37
        case "hello":
            if (((0.0 - (1.5 + 3.14159)) * ((0.1 - 3.14159) * (-2.25 + 1e10))) <= 0.1) {
                if false {
                    if len(v3) > 0 {
                        v3[0] = ((v26 + "a") + "foo bar")
                    }
                    delete(v2, "a")
                    delete(v2, (v35 + "\n\t"))
                }
                fmt.Println(v1, v35, v1)
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    switch (-(fi38)) {
                    case 2147483647:
                        v24 = (fi38 | fi38)
                    }
                    var v39 bool = true
                    fmt.Println(v25, v25, v1)
                    _ = fi38
                    _ = v39
                }
                v2["hello"] = v25
            }
        default:
            var v40 uint64 = 18446744073709551615
            _ = v40
        }
    case "foo bar":
        if false {
            fmt.Println(v25, v4)
            v41 := GenS5{F0: 2, F1: (-624 > 0)}
            _ = v41
        }
    }
    switch v35 {
    case "foo bar":
        var v42 int = 321
        _ = v42
    case "":
        v43 := v2["\n\t"]
        _ = v43
    case "hello":
        for fi44 := 0; fi44 < 10; fi44++ {
            v35 = v35
            _ = fi44
        }
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        _ = len(v3)
        _ = fi45
    }
    switch 2147483647 {
    case 1:
        if true {
            var v46 float32 = float32((-2.25 + 1e10))
            _ = v46
        } else {
            for fi47 := 0; fi47 < 13; fi47++ {
                v2[v26] = v25
                var v48 uint16 = 65535
                _ = fi47
                _ = v48
            }
        }
    case -9223372036854775808:
        if true {
            switch (v35 + (v26 + (v35 + ""))) {
            case "hello":
                for fi49 := 0; fi49 < 8; fi49++ {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    if (false || (18446744073709551615 <= 18446744073709551615)) {
                        v50 := v2[(v26 + v35)]
                        for fi51 := 0; fi51 < 11; fi51++ {
                            var v52 uint16 = 255
                            _ = fi51
                            _ = v52
                        }
                        var v53 uint64 = 4294967295
                        _ = v50
                        _ = v53
                    } else {
                        switch v24 {
                        case 42:
                            var v54 uint32 = genFn0()
                            _ = v54
                        case -9223372036854775808:
                            v55 := GenS6{F0: 9223372036854775807, F1: 2, F2: 0}
                            _ = v55
                        case -2147483648:
                            var v56 uint8 = 0
                            _ = v56
                        default:
                            v24 = (v24 * (v24 / ((127) | 1)))
                        }
                        for fi57 := 0; fi57 < 3; fi57++ {
                            if ((17 > 42) && ((false || false) || true)) {
                                v58 := GenS4{F0: 0}
                                if (989 <= 2) {
                                    _ = len(v3)
                                    if true {
                                        var v59 uint8 = 255
                                        v60 := v2[("hello" + v26)]
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 uint32 = (v25 - 2147483648)
                                        if (!(-1 == -1)) {
                                            if false {
                                                fmt.Println(v61, fi57)
                                                var v62 bool = false
                                                _ = v62
                                            } else {
                                                fmt.Println(v24, fi57, v25)
                                                var v63 string = v26
                                                _ = v63
                                            }
                                        }
                                        var v64 uint64 = 370
                                        _ = v61
                                        _ = v64
                                    }
                                }
                                _ = v58
                            }
                            _ = fi57
                        }
                        var v65 int8 = -1
                        _ = v65
                    }
                    if true {
                        v35 = ((v26 + "hello") + v35)
                        v66 := GenS2{F0: 42, F1: 0, F2: 252}
                        _ = v66
                    } else {
                        v4 = (v4 - -404)
                        v25 += (v25 >> ((42) & 63))
                        v67 := GenS7{F0: 65535, F1: v25}
                        _ = v67
                    }
                    _ = fi49
                }
            default:
                v68 := GenS7{F0: 914, F1: (v1 / (((v1 >> ((42) & 63))) | 1))}
                _ = v68
            }
            v69 := v2[v26]
            fmt.Println(v24)
            _ = v69
        }
    case -917:
        fmt.Println(v35)
    default:
        if false {
            var v70 uint64 = 4294967295
            _ = v70
        } else {
            v4 -= v4
        }
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v4
}
