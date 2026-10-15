package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 bool = (3.14159 != (-2.25 * (0.1 - 1e10)))
    v3 := []uint32{}
    v4 := make(map[string]string)
    v5 := []uint32{1}
    v6 := GenS1{F0: 898}
    if v2 {
        v7 := GenS2{F0: "\n\t", F1: (!(!v2)), F2: 364}
        v2 = (!v2)
        _ = len(v5)
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if v2 {
            switch "foo bar" {
            case "foo bar":
                for fi9 := 0; fi9 < 4; fi9++ {
                    v10 := GenS2{F0: "foo bar", F1: (107 <= 127), F2: 62}
                    var v11 uint64 = 42
                    v12 := v4[("" + "hello")]
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            default:
                v13 := v4["hello"]
                _ = v13
            }
            var v14 uint16 = (^(v1))
            _ = v14
        }
        _ = fi8
    }
    if v2 {
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := v4["\n\t"]
            v16 = "foo bar"
            var v17 int8 = -128
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int16 = 1000
        if ((0.0 / (1.5 + (0.0 * float64((1e10 / 0.1))))) != (0.0 * (3.14159 + 1e10))) {
            if v2 {
                var v19 string = (("" + "foo bar") + ((((("a" + "") + "hello") + "foo bar") + (("hello" + "foo bar") + (("a" + "") + "a"))) + ("hello" + ("a" + ("a" + "hello")))))
                _ = v19
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v2 = (!v2)
                    _ = fi20
                }
                v2 = v2
                if v2 {
                    if v2 {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            switch (("\n\t" + "foo bar") + "foo bar") {
                            case "foo bar":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "a":
                                _ = len(v4)
                            default:
                                fmt.Println(v18, v18)
                            }
                            switch fi21 {
                            case 32768:
                                if (!(-2147483648 <= -2147483648)) {
                                    var v22 int16 = 32767
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: "foo bar", F1: (v2 || (uint64(v1) < 42)), F2: -1}
                                    _ = v23
                                }
                            case 42:
                                v1 = 42
                            case 587:
                                switch "hello" {
                                case "\n\t":
                                    v24 := GenS3{F0: 528, F1: (1 != 124)}
                                    _ = v24
                                case "a":
                                    fmt.Println(v18, v1)
                                default:
                                    switch "" {
                                    case "\n\t":
                                        for fi25 := 0; fi25 < 2; fi25++ {
                                            v1 = v1
                                            _ = fi25
                                        }
                                    case "hello":
                                        var v26 int32 = 815
                                        _ = v26
                                    }
                                }
                            default:
                                v27 := GenS2{F0: "foo bar", F1: (154 != 1), F2: 42}
                                _ = v27
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi21
                        }
                        v1 = v1
                        v1 -= v1
                    } else {
                        v18 = (v18 - (v18 + v18))
                        v1 -= (^(v1))
                        v1 = v1
                    }
                }
            }
        }
        _ = v18
    }
    var v28 bool = false
    switch -2147483648 {
    case -754:
        var v29 float32 = 0.0
        _ = v29
    case 2147483647:
        var v30 float64 = -2.25
        _ = v30
    case 32768:
        for fi31 := 0; fi31 < 6; fi31++ {
            if true {
                if (!v28) {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        var v33 float64 = (1.5 - 1e10)
                        v34 := v4[((("\n\t" + "\n\t") + "a") + "")]
                        var v35 int16 = 42
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                } else {
                    if v2 {
                        fmt.Println(v1, v1)
                    }
                }
            }
            _ = fi31
        }
    }
    v28 = v28
    if (0.0 < 3.14159) {
        fmt.Println(v1, v2)
        for fi36 := 0; fi36 < 4; fi36++ {
            if false {
                var v37 uint32 = 42
                switch fi36 {
                case 981:
                    var v38 int32 = 42
                    _ = v38
                case 42:
                    if false {
                        fmt.Println(fi36, v37, fi36)
                        switch (fi36 ^ fi36) {
                        case 753:
                            for fi39 := 0; fi39 < 7; fi39++ {
                                if false {
                                    var v40 uint = 920
                                    var v41 int16 = -1
                                    _ = v40
                                    _ = v41
                                } else {
                                    switch fi39 {
                                    case -317:
                                        switch fi39 {
                                        case 128:
                                            v37 += v37
                                        case 32768:
                                            if v2 {
                                                v42 := GenS4{F0: v2, F1: 506}
                                                var v43 uint8 = uint8(v1)
                                                v28 = (v2 || v2)
                                                _ = v42
                                                _ = v43
                                            } else {
                                                _ = len(v4)
                                                switch fi36 {
                                                case 4294967295:
                                                    fmt.Println(v1, fi39)
                                                default:
                                                    switch (fi36 + (fi39 % ((fi39) | 1))) {
                                                    case 1:
                                                        switch fi36 {
                                                        case -835:
                                                            v44 := GenS2{F0: "foo bar", F1: v28, F2: -1}
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            }
                                        case 4294967295:
                                            v1 -= (v1 ^ (v1 ^ (v1 & v1)))
                                        }
                                        if ("foo bar" <= "hello") {
                                            var v45 int64 = -1
                                            v2 = ((fi36 & fi39) == (fi36 >> ((65535) & 63)))
                                            if (-1 > -106) {
                                                var v46 string = "hello"
                                                _ = v46
                                            }
                                            _ = v45
                                        }
                                        fmt.Println(v1, fi39)
                                    case 373:
                                        var v47 float64 = (float64((1.5 + 3.14159)) / 1.0)
                                        _ = v47
                                    default:
                                        fmt.Println(v28, fi39)
                                    }
                                    v4["hello"] = "a"
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                if false {
                                    v48 := GenS1{F0: 18446744073709551615}
                                    _ = v48
                                }
                                v5 = append(v5, (v37 << ((4294967295) & 63)))
                                switch ("" + "") {
                                case "a":
                                    switch ("hello" + "hello") {
                                    case "a":
                                        fmt.Println(fi39, fi36, v37)
                                    }
                                }
                                _ = fi39
                            }
                        case 42:
                            v1 += 0
                        default:
                            v3 = append(v3, v37)
                        }
                    }
                case 193:
                    var v49 uint = uint(v37)
                    _ = v49
                }
                _ = v37
            } else {
                for fi50 := 0; fi50 < 6; fi50++ {
                    v1 += v1
                    var v51 string = ((("foo bar" + "a") + "hello") + ("\n\t" + ""))
                    switch (v51 + ((v51 + (v51 + "hello")) + (v51 + (v51 + (v51 + v51))))) {
                    case "\n\t":
                        if false {
                            var v52 uint = 0
                            if v28 {
                                v5 = append(v5, 1)
                                v1 += 42
                                fmt.Println(fi50, v52, v1)
                            }
                            _ = v52
                        } else {
                            var v53 string = "foo bar"
                            _ = v53
                        }
                    }
                    v28 = false
                    _ = fi50
                    _ = v51
                }
                fmt.Println(v28, v1, fi36)
                v28 = v2
            }
            var v54 int64 = 1
            if ((0.1 * 0.0) != 0.0) {
                var v55 int64 = (v54 << ((331) & 63))
                v54 -= (v54 / ((v55) | 1))
                _ = v55
            }
            _ = fi36
            _ = v54
        }
        if true {
            var v56 uint16 = v1
            v1 = (v1 - (v56 << ((18446744073709551615) & 63)))
            _ = v56
        }
    } else {
        v57 := GenS5{F0: 42}
        var v58 bool = (42 < 208)
        v58 = ((42 < 65535) && (!(-855 != 1)))
        _ = v57
        _ = v58
    }
    switch -128 {
    case 791:
        switch "foo bar" {
        case "a":
            if v2 {
                fmt.Println(v28, v1)
                if len(v5) > 0 {
                    v5[0] = 367
                }
                var v59 int = -9223372036854775808
                _ = v59
            }
        case "\n\t":
            if true {
                v1 = v1
                _ = len(v5)
            } else {
                v1 = 1
            }
        }
    case 32767:
        if v28 {
            if false {
                fmt.Println(v2)
                fmt.Println(v28, v1, v1)
                fmt.Println(v28, v28)
            } else {
                for fi60 := 0; fi60 < 7; fi60++ {
                    delete(v4, "")
                    _ = fi60
                }
                switch 9223372036854775807 {
                case 42:
                    v61 := v4[""]
                    _ = v61
                default:
                    switch "\n\t" {
                    case "":
                        fmt.Println(v1, v1)
                    default:
                        v4["hello"] = "\n\t"
                    }
                }
                v1 = v1
            }
        } else {
            var v62 int = 1
            _ = v62
        }
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        v64 := GenS6{F0: -1, F1: 42, F2: 2147483648}
        if v28 {
            var v65 string = ("a" + ("\n\t" + "\n\t"))
            v2 = false
            for fi66 := 0; fi66 < 3; fi66++ {
                if v2 {
                    fmt.Println(v65, fi63)
                } else {
                    v3 = append(v3, 934)
                    var v67 int64 = 9223372036854775807
                    var v68 bool = (!(-88 > 100))
                    _ = v67
                    _ = v68
                }
                fmt.Println(v2)
                _ = fi66
            }
            _ = v65
        }
        v69 := v4["foo bar"]
        _ = fi63
        _ = v64
        _ = v69
    }
    switch 2147483647 {
    case 0:
        switch 897 {
        case -443:
            var v70 string = "hello"
            _ = v70
        case 128:
            v28 = v2
        default:
            var v71 uint16 = 0
            _ = v71
        }
    case -2147483648:
        var v72 uint64 = 278
        _ = v72
    case 32767:
        var v73 uint8 = 128
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
