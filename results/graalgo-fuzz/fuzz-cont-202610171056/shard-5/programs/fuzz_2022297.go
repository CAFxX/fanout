package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 bool; F1 int; F2 int8 }
func genFn0(p0 uint32, p1 int16) uint64 {
    p0 += p0
    v7 := GenS1{F0: 695, F1: 42}
    var v8 int16 = p1
    var v9 float64 = 3.14159
    v10 := GenS1{F0: 4294967295, F1: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 914
}

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar"}
    v2 := []string{}
    var v3 bool = true
    var v4 uint = 2
    var v5 uint8 = 127
    var v6 bool = ((9223372036854775807 != -1) || (9223372036854775807 == 42))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    if true {
        v11 := GenS1{F0: 0, F1: 2147483647}
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if false {
        v12 := GenS2{F0: 128, F1: 42}
        _ = v12
    }
    var v13 float64 = ((-2.25 + 3.14159) * 3.14159)
    fmt.Println(v6, v13, v13)
    if ((v5 / ((243) | 1)) > v5) {
        var v14 float64 = -2.25
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int32 = 0
                if len(v2) > 0 {
                    v2[0] = ""
                }
                if true {
                    if v6 {
                        fmt.Println(fi15, fi16, v6)
                        var v18 float64 = v14
                        _ = v18
                    } else {
                        v13 = v14
                        v5 = (v5 | 1)
                    }
                    v3 = v6
                }
                _ = fi16
                _ = v17
            }
            var v19 uint = v4
            v20 := GenS3{F0: "hello", F1: v19}
            _ = fi15
            _ = v19
            _ = v20
        }
        switch ("hello" + "a") {
        case "hello":
            v6 = v3
        case "":
            for fi21 := 0; fi21 < 5; fi21++ {
                for fi22 := 0; fi22 < 7; fi22++ {
                    switch "a" {
                    case "a":
                        if true {
                            switch (^(fi21)) {
                            case 32768:
                                var v23 uint = 0
                                _ = v23
                            case 0:
                                v24 := GenS1{F0: 18446744073709551615, F1: -1}
                                _ = v24
                            case 992:
                                switch "foo bar" {
                                case "\n\t":
                                    if (-1 <= 2147483647) {
                                        if (0 != 1) {
                                            fmt.Println(v3)
                                        }
                                        v6 = v3
                                    } else {
                                        v25 := GenS2{F0: 0, F1: 42}
                                        var v26 int16 = 42
                                        v27 := GenS1{F0: 262, F1: -2147483648}
                                        var v28 int16 = (v26 ^ (v26 >> ((v4) & 63)))
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                case "foo bar":
                                    var v29 uint16 = 1
                                    _ = v29
                                }
                            default:
                                v30 := GenS2{F0: (^(fi21)), F1: 0}
                                _ = v30
                            }
                            v31 := GenS3{F0: ("hello" + "hello"), F1: (v4 << ((694) & 63))}
                            _ = v31
                        } else {
                            var v32 bool = (v4 == v4)
                            v13 = float64(1.5)
                            _ = v32
                        }
                    case "\n\t":
                        v33 := GenS1{F0: 680, F1: 868}
                        _ = v33
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v34 float64 = 1.5
                        _ = v34
                    default:
                        var v35 bool = v6
                        _ = v35
                    }
                    if true {
                        if v6 {
                            switch "a" {
                            case "hello":
                                fmt.Println(fi22, fi21)
                                var v36 uint8 = (^(v5))
                                _ = v36
                            case "":
                                v2 = append(v2, ("hello" + "a"))
                            }
                            var v37 uint64 = 1
                            v14 += float64(3.14159)
                            _ = v37
                        } else {
                            switch 0 {
                            case 349:
                                _ = len(v1)
                            case 32767:
                                var v38 bool = v3
                                _ = v38
                            case -9223372036854775808:
                                fmt.Println(v14, v14, v14)
                            default:
                                if v6 {
                                    switch fi21 {
                                    case 2147483647:
                                        if ((!v6) || (v6 && (v3 && (-2.25 == ((3.14159 * 0.0) / 0.1))))) {
                                            fmt.Println(v5)
                                            v14 = 0.0
                                            fmt.Println(v5, v3)
                                        }
                                    case 426:
                                        v39 := GenS2{F0: fi21, F1: 1}
                                        _ = v39
                                    default:
                                        if ((3.14159 >= (float32((3.14159 * v14)) - 0.1)) || (2 == 437)) {
                                            var v40 int16 = 0
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(fi21, v40, fi22)
                                            case "foo bar":
                                                var v41 bool = (!(!(!false)))
                                                _ = v41
                                            case "\n\t":
                                                var v42 int = (-(fi22))
                                                _ = v42
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    if true {
                                                        v5 = (v5 % (((v5 / ((v5) | 1))) | 1))
                                                        v43 := GenS4{F0: 91}
                                                        _ = v43
                                                    }
                                                default:
                                                    _ = len(v2)
                                                }
                                            }
                                            _ = v40
                                        } else {
                                            var v44 bool = (!(1 <= 42))
                                            _ = v44
                                        }
                                    }
                                    v4 += 2
                                }
                            }
                            _ = len(v1)
                            v45 := GenS5{F0: 499, F1: int(v4)}
                            _ = v45
                        }
                        v13 = 3.14159
                        v5 += uint8(v4)
                    }
                    switch "hello" {
                    case "":
                        var v46 int64 = -729
                        _ = v46
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    _ = fi22
                }
                _ = fi21
            }
        }
        _ = v14
    }
    v13 = v13
    v47 := GenS3{F0: "foo bar", F1: v4}
    for fi48 := 0; fi48 < 7; fi48++ {
        if ((uint16(v5) != 864) || (v3 && (0 != -596))) {
            v1 = append(v1, (("hello" + (("\n\t" + ((("\n\t" + "foo bar") + "\n\t") + "\n\t")) + "")) + "a"))
        }
        var v49 uint16 = 1
        _ = fi48
        _ = v49
    }
    var v50 uint16 = 0
    fmt.Println(v3, v4, v3)
    if v3 {
        v3 = v6
        v50 += (v50 - (v50 - v50))
        for fi51 := 0; fi51 < 6; fi51++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if v3 {
                v3 = true
                v1 = append(v1, "\n\t")
            } else {
                var v52 int32 = 1
                v53 := GenS6{F0: true, F1: fi51, F2: 100}
                _ = v52
                _ = v53
            }
            _ = fi51
        }
    }
    switch 127 {
    case 42:
        for fi54 := 0; fi54 < 12; fi54++ {
            if false {
                for fi55 := 0; fi55 < 14; fi55++ {
                    v5 = 101
                    _ = fi55
                }
            } else {
                v3 = v3
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi54
        }
    case -57:
        v5 = v5
    case 928:
        v56 := GenS3{F0: "foo bar", F1: uint(v50)}
        _ = v56
    default:
        for fi57 := 0; fi57 < 13; fi57++ {
            for fi58 := 0; fi58 < 3; fi58++ {
                fmt.Println(v3, fi57, v4)
                if true {
                    var v59 float32 = 0.0
                    _ = v59
                } else {
                    v1 = append(v1, "")
                    var v60 uint64 = 1
                    _ = v60
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = fi58
            }
            if (-1 >= -1) {
                switch "a" {
                case "\n\t":
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v6)
                    default:
                        v3 = v6
                    }
                }
            }
            _ = fi57
        }
    }
    switch "\n\t" {
    case "":
        var v61 uint8 = v5
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
