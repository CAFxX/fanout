package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
type GenS5 struct { F0 uint; F1 int64; F2 int }
type GenS6 struct { F0 uint64; F1 string; F2 uint16 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0(p0 string) uint16 {
    if false {
        v6 := []int{}
        p0 = p0
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        if (2 >= 32767) {
            v8 := []int{fi7, fi7, (^(fi7)), (fi7 + (-(fi7)))}
            p0 = ""
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            switch p0 {
            case "":
                fmt.Println(fi7, p0)
            case "foo bar":
                v10 := []int{fi7, 128, (fi7 | fi9), 996}
                _ = v10
            default:
                fmt.Println(fi7, p0, fi7)
            }
            fmt.Println(fi9)
            _ = fi9
        }
        _ = fi7
    }
    _ = p0
    return 42
}
func genFn1(p0 bool) uint16 {
    v11 := []int{2147483647}
    v11 = append(v11, 32768)
    if (!true) {
        v12 := GenS1{F0: 1, F1: 622}
        p0 = (p0 || p0)
        var v13 uint16 = 42
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v11
    return 462
}

func main() {
    var v1 uint32 = 2
    var v2 uint32 = (v1 ^ v1)
    v3 := make(map[int]int)
    var v4 bool = (-190 == -832)
    var v5 bool = v4
    fmt.Println(v4)
    v14 := GenS1{F0: 4, F1: 420}
    if ((!true) && (2147483647 <= 42)) {
        v4 = false
        fmt.Println(v2, v1, v5)
    }
    if v4 {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS1{F0: 42, F1: 65535}
            for fi17 := 0; fi17 < 4; fi17++ {
                if v5 {
                    var v18 int64 = 584
                    var v19 string = "\n\t"
                    var v20 int64 = v18
                    v19 = v19
                    _ = v18
                    _ = v19
                    _ = v20
                }
                v21 := []int{(fi15 | (^(fi17))), -217, -387}
                _ = fi17
                _ = v21
            }
            _ = fi15
            _ = v16
        }
        if ((!("a" <= "a")) || false) {
            switch "\n\t" {
            case "foo bar":
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 int = fi22
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 bool = (false || (42 >= 0))
                        v2 = 968
                        if v5 {
                            v26 := []int{fi24, -766, 42}
                            var v27 uint64 = 700
                            _ = v26
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
            case "hello":
                v2 = (^(v1))
            case "a":
                var v28 string = (("a" + "foo bar") + ("a" + "foo bar"))
                _ = v28
            default:
                v29 := GenS2{F0: v2, F1: (true || ((3.14159 * (0.0 + 1e10)) < (3.14159 + 3.14159))), F2: 42}
                _ = v29
            }
            v1 = v1
            v30 := GenS2{F0: (^(v2)), F1: v4, F2: 56}
            _ = v30
        } else {
            v31 := GenS2{F0: v1, F1: true, F2: -1}
            _ = v31
        }
        if ((!true) || (true && (-634 <= 862))) {
            var v32 string = (("hello" + ("foo bar" + (("\n\t" + "\n\t") + ""))) + ((("\n\t" + ("\n\t" + "foo bar")) + "\n\t") + ("" + "a")))
            for fi33 := 0; fi33 < 15; fi33++ {
                v3[(fi33 | (fi33 & (fi33 + (fi33 << ((1) & 63)))))] = fi33
                if v4 {
                    v34 := v3[(fi33 * fi33)]
                    var v35 uint8 = 2
                    v34 -= (fi33 * -767)
                    _ = v34
                    _ = v35
                }
                fmt.Println(v1, v32, v1)
                _ = fi33
            }
            _ = v32
        } else {
            if v5 {
                v36 := GenS3{F0: v5, F1: genFn0("\n\t"), F2: "a"}
                switch "" {
                case "foo bar":
                    fmt.Println(v1)
                case "a":
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v38 := GenS2{F0: v2, F1: (-439 >= 1000), F2: 19}
                        v5 = ((-705 < 893) || (!(2 <= 255)))
                        v39 := v3[(fi37 % (((fi37 >> ((0) & 63))) | 1))]
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
                _ = v36
            }
            v40 := []int{-409, 127, 648, -2147483648}
            _ = v40
        }
    } else {
        var v41 string = ("hello" + "hello")
        fmt.Println(v1, v41, v41)
        _ = v41
    }
    if (9223372036854775807 != 2147483647) {
        var v42 int64 = 1
        if v4 {
            v1 -= 42
            if true {
                if false {
                    var v43 float32 = ((((1.5 * ((0.0 / 3.14159) + 0.0)) - 0.0) - 1e10) / ((1e10 + 1e10) - 0.1))
                    var v44 int16 = 1000
                    if v5 {
                        if v5 {
                            v45 := []int{37}
                            v4 = (v5 || v5)
                            var v46 string = "\n\t"
                            _ = v45
                            _ = v46
                        }
                    } else {
                        v47 := GenS4{F0: 912, F1: 4294967295, F2: ""}
                        _ = v47
                    }
                    _ = v43
                    _ = v44
                } else {
                    v48 := GenS3{F0: (!true), F1: 1, F2: "a"}
                    v49 := []int{32768, 954, 127}
                    for fi50 := 0; fi50 < 19; fi50++ {
                        switch "a" {
                        case "\n\t":
                            for fi51 := 0; fi51 < 12; fi51++ {
                                v52 := GenS1{F0: 1, F1: 520}
                                _ = fi51
                                _ = v52
                            }
                        case "hello":
                            var v53 int = -718
                            _ = v53
                        case "":
                            var v54 uint16 = 1
                            _ = v54
                        default:
                            for fi55 := 0; fi55 < 7; fi55++ {
                                v56 := GenS5{F0: 255, F1: v42, F2: fi55}
                                v57 := v3[9223372036854775807]
                                if (!v5) {
                                    fmt.Println(v2)
                                    var v58 uint = 0
                                    var v59 int64 = -64
                                    fmt.Println(v57)
                                    v42 = 42
                                    _ = v58
                                    _ = v59
                                }
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        }
                        switch fi50 {
                        case 32768:
                            v60 := GenS6{F0: 105, F1: "\n\t", F2: uint16(v1)}
                            _ = v60
                        }
                        _ = fi50
                    }
                    _ = v48
                    _ = v49
                }
            } else {
                switch "a" {
                case "foo bar":
                    v61 := GenS3{F0: false, F1: 65535, F2: ("\n\t" + "foo bar")}
                    _ = v61
                default:
                    v62 := []int{971}
                    _ = v62
                }
                v63 := []int{}
                var v64 string = "\n\t"
                _ = v63
                _ = v64
            }
        } else {
            var v65 int64 = (v42 << ((255) & 63))
            var v66 uint16 = 255
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 15; fi67++ {
            for fi68 := 0; fi68 < 8; fi68++ {
                switch fi67 {
                case 128:
                    v5 = ((341 != 65535) || v5)
                case 656:
                    fmt.Println(v4, v2)
                }
                _ = fi68
            }
            switch fi67 {
            case 32767:
                var v69 int = fi67
                _ = v69
            case -9223372036854775808:
                if ((!v5) || false) {
                    if false {
                        v70 := GenS7{F0: 662, F1: (fi67 & fi67)}
                        _ = v70
                    }
                } else {
                    v71 := GenS5{F0: 18446744073709551615, F1: (v42 - v42), F2: fi67}
                    _ = v71
                }
            case 127:
                fmt.Println(v2)
            default:
                if (-1 != 1) {
                    v42 = v42
                }
            }
            fmt.Println(v2)
            _ = fi67
        }
        _ = v42
    } else {
        for fi72 := 0; fi72 < 5; fi72++ {
            for fi73 := 0; fi73 < 12; fi73++ {
                var v74 uint8 = 42
                _ = fi73
                _ = v74
            }
            v5 = v4
            v1 = (v1 ^ v2)
            _ = fi72
        }
    }
    v75 := GenS8{F0: v1}
    var v76 uint8 = 42
    for fi77 := 0; fi77 < 9; fi77++ {
        for fi78 := 0; fi78 < 20; fi78++ {
            v4 = v5
            v79 := GenS9{F0: 101, F1: 1, F2: v1}
            _ = fi78
            _ = v79
        }
        _ = fi77
    }
    for fi80 := 0; fi80 < 11; fi80++ {
        var v81 int = (fi80 | (fi80 % ((fi80) | 1)))
        _ = fi80
        _ = v81
    }
    switch ("" + "\n\t") {
    case "\n\t":
        for fi82 := 0; fi82 < 6; fi82++ {
            _ = len(v3)
            v2 -= 373
            _ = len(v3)
            _ = fi82
        }
    case "foo bar":
        if (v5 || v5) {
            v83 := GenS8{F0: (v1 + (v2 & (v2 | v1)))}
            _ = v83
        }
    case "a":
        fmt.Println(v2, v4)
    }
    if (false && false) {
        switch (("foo bar" + (("foo bar" + "") + ("\n\t" + "a"))) + "hello") {
        case "foo bar":
            switch ("" + "\n\t") {
            case "\n\t":
                var v84 int32 = -344
                _ = v84
            }
        default:
            _ = len(v3)
        }
        for fi85 := 0; fi85 < 18; fi85++ {
            var v86 uint64 = 2
            _ = fi85
            _ = v86
        }
        if true {
            v76 = v76
            v76 = (v76 ^ 42)
            var v87 uint64 = 42
            _ = v87
        }
    }
    for fi88 := 0; fi88 < 20; fi88++ {
        for fi89 := 0; fi89 < 8; fi89++ {
            v90 := GenS7{F0: -144, F1: (fi88 & 4294967295)}
            v5 = ((fi89 + (fi89 + (-(fi88)))) == 32767)
            v3[fi89] = fi88
            _ = fi89
            _ = v90
        }
        _ = fi88
    }
    _ = len(v3)
    v91 := GenS9{F0: 1, F1: 766, F2: 1}
    var v92 uint8 = 1
    v93 := GenS7{F0: 1, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v91
    _ = v92
    _ = v93
}
