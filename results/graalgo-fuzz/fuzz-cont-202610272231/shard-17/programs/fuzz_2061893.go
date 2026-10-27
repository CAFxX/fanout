package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 bool; F2 int }
type GenS5 struct { F0 bool; F1 string; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int64; F2 string }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0() bool {
    if false {
        switch "hello" {
        case "":
            switch 32768 {
            case 0:
                var v4 uint32 = 2147483648
                _ = v4
            case 793:
                v5 := []int{-2147483648, 2147483647}
                _ = v5
            }
        case "a":
            var v6 int64 = -481
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 bool = false
                switch "" {
                case "":
                    var v9 int8 = -24
                    _ = v9
                case "\n\t":
                    v8 = (1 < 42)
                case "hello":
                    v10 := GenS1{F0: 660, F1: "a", F2: 882}
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 int32 = -1
                        var v13 uint16 = 42
                        if false {
                            v8 = v8
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                _ = fi7
                _ = v8
            }
        }
        var v14 float32 = -2.25
        var v15 int8 = 14
        _ = v14
        _ = v15
    }
    var v16 uint32 = 611
    v17 := GenS1{F0: 887, F1: "hello", F2: 470}
    _ = v16
    _ = v17
    return (127 >= 42)
}
func genFn1(p0 bool) uint64 {
    if true {
        p0 = p0
        p0 = (75 == 758)
    } else {
        switch 0 {
        case 2147483647:
            p0 = (-2147483648 == 873)
        case -970:
            var v18 int8 = -128
            _ = v18
        case 32768:
            var v19 int = -9223372036854775808
            _ = v19
        }
        fmt.Println(p0)
    }
    p0 = false
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(p0)
        _ = fi20
    }
    _ = p0
    return 790
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 0
    v3 := []int{992, 32768, -9223372036854775808}
    if len(v3) > 0 {
        v3[0] = 42
    }
    v2 += v2
    var v21 float32 = (-2.25 + -2.25)
    if false {
        for fi22 := 0; fi22 < 1; fi22++ {
            v2 -= v2
            v23 := GenS1{F0: v1, F1: "foo bar", F2: (^(v1))}
            v24 := GenS1{F0: v1, F1: ("a" + ("hello" + "foo bar")), F2: 161}
            _ = fi22
            _ = v23
            _ = v24
        }
    } else {
        var v25 int = -51
        _ = v25
    }
    v3 = append(v3, 32767)
    for fi26 := 0; fi26 < 16; fi26++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (!((v2 - v2) <= -1)) {
            v2 = (-(v2))
        } else {
            if true {
                var v27 uint64 = 151
                v1 = v1
                v28 := GenS1{F0: v1, F1: "\n\t", F2: (v1 | (v1 + 65535))}
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 18; fi29++ {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = (fi29 | 4294967295)
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (fi26 | (fi29 >> ((677) & 63)))
                        }
                        if true {
                            v2 += 42
                            fmt.Println(v2)
                            v21 += v21
                        }
                    }
                    v3 = append(v3, (fi29 ^ fi29))
                    _ = fi29
                }
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                if (340 < 32767) {
                    fmt.Println(v2, fi26)
                } else {
                    switch (fi26 & fi30) {
                    case 32768:
                        var v31 bool = false
                        _ = v31
                    }
                    v2 = v2
                }
                _ = fi30
            }
            var v32 float32 = v21
            _ = v32
        }
        _ = fi26
    }
    v33 := GenS1{F0: v1, F1: "hello", F2: v1}
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            v2 += v2
            _ = len(v3)
            v1 = (^(v1))
            _ = fi35
        }
        _ = fi34
    }
    if false {
        var v36 uint = 65535
        v1 += v1
        v37 := GenS1{F0: v1, F1: "\n\t", F2: v1}
        _ = v36
        _ = v37
    } else {
        if false {
            v38 := GenS2{F0: (!(!true)), F1: 42, F2: (v2 << ((255) & 63))}
            _ = v38
        } else {
            var v39 bool = (-2.25 < 1e10)
            var v40 uint32 = 0
            var v41 int64 = 1
            _ = v39
            _ = v40
            _ = v41
        }
        if len(v3) > 0 {
            v3[0] = 127
        }
        var v42 float32 = v21
        _ = v42
    }
    var v43 string = "foo bar"
    if (false || (42 < 1000)) {
        for fi44 := 0; fi44 < 12; fi44++ {
            var v45 int32 = (v2 * 2147483647)
            _ = fi44
            _ = v45
        }
        switch (v43 + ("foo bar" + v43)) {
        case "a":
            fmt.Println(v21, v21)
        default:
            v43 = ((v43 + (v43 + "")) + "")
        }
        v46 := GenS1{F0: (v1 & 778), F1: v43, F2: v1}
        _ = v46
    } else {
        if (!true) {
            var v47 uint = 65535
            _ = v47
        } else {
            for fi48 := 0; fi48 < 6; fi48++ {
                var v49 int = (^(fi48))
                _ = fi48
                _ = v49
            }
        }
    }
    for fi50 := 0; fi50 < 14; fi50++ {
        v51 := GenS3{F0: false}
        switch v43 {
        case "\n\t":
            fmt.Println(v21)
        case "":
            v52 := GenS1{F0: (v1 << ((1) & 63)), F1: "\n\t", F2: v1}
            _ = v52
        case "a":
            fmt.Println(v21, v2, fi50)
        default:
            if ((v2 << ((4294967295) & 63)) < v2) {
                v53 := GenS4{F0: 42, F1: true, F2: fi50}
                v54 := GenS3{F0: false}
                if (false || (42 == 42)) {
                    v2 = v2
                    for fi55 := 0; fi55 < 2; fi55++ {
                        var v56 int64 = -1
                        v57 := GenS5{F0: ((v56 & 0) != v56), F1: "a", F2: (v2 & 0)}
                        fmt.Println(v43)
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    if true {
                        for fi58 := 0; fi58 < 14; fi58++ {
                            if true {
                                if true {
                                    var v59 int = fi50
                                    fmt.Println(fi58)
                                    _ = v59
                                }
                                v60 := GenS1{F0: 0, F1: (v43 + v43), F2: (v1 >> ((1) & 63))}
                                switch "foo bar" {
                                case "":
                                    fmt.Println(v43)
                                default:
                                    _ = len(v3)
                                }
                                var v61 int16 = -1
                                _ = v60
                                _ = v61
                            }
                            _ = fi58
                        }
                        for fi62 := 0; fi62 < 20; fi62++ {
                            switch v43 {
                            case "":
                                fmt.Println(v43, v2)
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = (^(fi62))
                                }
                            default:
                                v63 := GenS6{F0: 19}
                                _ = v63
                            }
                            _ = len(v3)
                            _ = fi62
                        }
                    }
                }
                _ = v53
                _ = v54
            } else {
                var v64 uint16 = v1
                if false {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        var v66 uint64 = genFn1((true && true))
                        fmt.Println(v64, v1, fi50)
                        fmt.Println(v66)
                        _ = fi65
                        _ = v66
                    }
                    switch (fi50 % (((fi50 % ((fi50) | 1))) | 1)) {
                    case -904:
                        v3 = append(v3, (fi50 ^ fi50))
                    case -954:
                        v64 -= v64
                    }
                } else {
                    if (109 >= 1) {
                        var v67 uint32 = 0
                        for fi68 := 0; fi68 < 18; fi68++ {
                            v69 := GenS7{F0: 53, F1: 1, F2: v43}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi68
                            _ = v69
                        }
                        _ = v67
                    }
                    v3 = append(v3, -1)
                }
                _ = v64
            }
        }
        _ = fi50
        _ = v51
    }
    if false {
        fmt.Println(v43, v43, v43)
        _ = len(v3)
    }
    switch 9223372036854775807 {
    case -949:
        v70 := GenS8{F0: 0, F1: 42, F2: 127}
        _ = v70
    case -2147483648:
        fmt.Println(v43, v43)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v43
}
