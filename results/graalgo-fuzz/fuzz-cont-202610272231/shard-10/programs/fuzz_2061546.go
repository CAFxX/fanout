package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int64 = -503
    v2 := []string{}
    var v3 bool = true
    var v4 int8 = 1
    switch 86 {
    case 32767:
        v3 = false
    case 914:
        v5 := GenS1{F0: "a", F1: -1, F2: 255}
        _ = v5
    case -9223372036854775808:
        v1 = -375
    default:
        switch 2147483647 {
        case 42:
            var v6 uint16 = 606
            _ = v6
        case -399:
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS2{F0: fi7, F1: 445, F2: 32767}
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 string = (("a" + "") + "foo bar")
                    switch v10 {
                    case "a":
                        var v11 int32 = -2147483648
                        _ = v11
                    case "\n\t":
                        fmt.Println(fi7)
                    }
                    if (42 >= 1000) {
                        var v12 float64 = (-2.25 / (3.14159 * 3.14159))
                        _ = v12
                    } else {
                        var v13 uint8 = 2
                        var v14 uint8 = v13
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                }
                for fi15 := 0; fi15 < 9; fi15++ {
                    fmt.Println(fi7, v1)
                    v4 = v4
                    v4 += v4
                    _ = fi15
                }
                _ = fi7
                _ = v8
            }
        case 314:
            for fi16 := 0; fi16 < 7; fi16++ {
                switch "a" {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                default:
                    switch fi16 {
                    case -217:
                        var v17 float32 = (((((-2.25 + ((((1.5 - 0.0) - 0.1) - (float32(3.14159) * (1e10 / 1.5))) * 1.5)) * ((1.5 + ((1.5 * float32(float64(float32(0.0)))) - -2.25)) * 3.14159)) + -2.25) * ((1.5 * (1.5 - float32(1e10))) / (1.5 * 0.1))) * 0.0)
                        _ = v17
                    }
                }
                if (!false) {
                    var v18 int64 = v1
                    if (2 < 1) {
                        switch int(v1) {
                        case 1:
                            v19 := GenS1{F0: "a", F1: (fi16 << ((65535) & 63)), F2: 104}
                            _ = v19
                        case 127:
                            for fi20 := 0; fi20 < 2; fi20++ {
                                var v21 uint32 = 65535
                                var v22 uint = 65535
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case 962:
                            if v3 {
                                switch "foo bar" {
                                case "foo bar":
                                    v3 = v3
                                }
                            } else {
                                v1 += v1
                            }
                        default:
                            fmt.Println(v3, v18, v4)
                        }
                        switch (fi16 | 9223372036854775807) {
                        case 4294967295:
                            var v23 uint8 = 1
                            _ = v23
                        }
                    } else {
                        v18 += -9223372036854775808
                    }
                    if ((false || ((fi16 - (fi16 & fi16)) >= (fi16 / ((-630) | 1)))) || (v3 || false)) {
                        v24 := GenS2{F0: (fi16 % (((fi16 - (fi16 % (((fi16 % ((fi16) | 1))) | 1)))) | 1)), F1: 2, F2: -32768}
                        v25 := GenS1{F0: "foo bar", F1: -2147483648, F2: 2}
                        v26 := GenS1{F0: ("hello" + "foo bar"), F1: fi16, F2: 128}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = v18
                } else {
                    for fi27 := 0; fi27 < 3; fi27++ {
                        var v28 float32 = (((1.5 / 3.14159) / 3.14159) * 1.5)
                        if v3 {
                            v29 := GenS3{F0: (v4 / ((-48) | 1)), F1: 164}
                            _ = v29
                        } else {
                            var v30 int32 = 0
                            _ = v30
                        }
                        fmt.Println(v28, v3)
                        _ = fi27
                        _ = v28
                    }
                    for fi31 := 0; fi31 < 4; fi31++ {
                        switch ("\n\t" + "") {
                        case "hello":
                            fmt.Println(fi16)
                        default:
                            var v32 bool = (1e10 == 3.14159)
                            _ = v32
                        }
                        v33 := GenS2{F0: fi16, F1: 1, F2: 0}
                        _ = fi31
                        _ = v33
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi16
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v3 = false
    v3 = v3
    var v34 uint64 = 42
    var v35 int64 = (v1 * v1)
    switch "a" {
    case "foo bar":
        if false {
            v36 := GenS1{F0: "a", F1: 32768, F2: 0}
            _ = v36
        } else {
            var v37 int32 = 42
            switch 1 {
            case 127:
                var v38 float32 = 0.1
                _ = v38
            case 5:
                var v39 bool = false
                _ = v39
            }
            if true {
                v2 = append(v2, ("" + "hello"))
                if (42 > 477) {
                    v40 := GenS1{F0: ("\n\t" + "foo bar"), F1: 32767, F2: 138}
                    var v41 int16 = 369
                    var v42 uint8 = 127
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            } else {
                if (0 != v37) {
                    v35 = v35
                    if v3 {
                        var v43 uint16 = 0
                        for fi44 := 0; fi44 < 5; fi44++ {
                            var v45 uint = 65535
                            _ = fi44
                            _ = v45
                        }
                        var v46 int8 = 1
                        _ = v43
                        _ = v46
                    }
                }
                v1 = (v35 + v1)
                _ = len(v2)
            }
            _ = v37
        }
    case "":
        v3 = (1e10 < ((1e10 * (((0.0 * 1.5) - float32(((0.1 + 3.14159) * 1e10))) / 3.14159)) * 0.0))
    case "a":
        fmt.Println(v4, v3, v35)
    default:
        v1 += v1
    }
    switch "hello" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = (("" + "a") + "foo bar")
            }
        case "hello":
            v47 := GenS2{F0: -1, F1: uint32(v34), F2: 1000}
            _ = v47
        case "a":
            var v48 float32 = float32(3.14159)
            _ = v48
        }
    case "foo bar":
        switch "\n\t" {
        case "a":
            v3 = v3
        case "":
            v49 := GenS3{F0: (v4 % ((33) | 1)), F1: 0}
            _ = v49
        default:
            if ("" <= "hello") {
                switch 4294967295 {
                case 9223372036854775807:
                    v50 := GenS3{F0: 42, F1: 65535}
                    _ = v50
                case 127:
                    for fi51 := 0; fi51 < 8; fi51++ {
                        var v52 uint8 = 42
                        fmt.Println(v1, v1)
                        v34 = (^(v34))
                        _ = fi51
                        _ = v52
                    }
                }
            } else {
                var v53 uint8 = 70
                _ = v53
            }
        }
    }
    if v3 {
        switch "\n\t" {
        case "a":
            for fi54 := 0; fi54 < 1; fi54++ {
                for fi55 := 0; fi55 < 15; fi55++ {
                    v34 -= v34
                    var v56 float32 = 0.0
                    if v3 {
                        v35 = v1
                        v57 := GenS4{F0: -741, F1: v4, F2: 86}
                        var v58 uint64 = v34
                        _ = v57
                        _ = v58
                    } else {
                        if true {
                            var v59 float32 = v56
                            v34 = (^(v34))
                            fmt.Println(fi54)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v59
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    }
                    _ = fi55
                    _ = v56
                }
                fmt.Println(fi54, v35, v1)
                _ = fi54
            }
        case "\n\t":
            var v60 int16 = -1
            _ = v60
        }
        var v61 uint32 = 1
        _ = v61
    } else {
        v62 := GenS5{F0: 1, F1: ""}
        v1 -= (^(v35))
        _ = v62
    }
    switch 2147483647 {
    case -9223372036854775808:
        switch 9223372036854775807 {
        case 0:
            v63 := GenS6{F0: 1, F1: 65535}
            _ = v63
        case 2147483647:
            v64 := GenS3{F0: v4, F1: 658}
            _ = v64
        case 20:
            fmt.Println(v34, v1)
        default:
            if v3 {
                var v65 bool = (((!(true && (v3 && false))) && v3) && v3)
                var v66 int16 = 32767
                _ = v65
                _ = v66
            }
        }
    case 128:
        var v67 uint64 = 805
        _ = v67
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        v4 = v4
        var v69 uint32 = 2147483648
        _ = fi68
        _ = v69
    }
    switch 32767 {
    case 127:
        v70 := GenS1{F0: "a", F1: 9223372036854775807, F2: 255}
        _ = v70
    default:
        v71 := GenS6{F0: 0, F1: 741}
        _ = v71
    }
    var v72 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v72
}
