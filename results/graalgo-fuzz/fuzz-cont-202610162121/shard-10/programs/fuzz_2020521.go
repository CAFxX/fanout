package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint; F2 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 uint32; F2 bool }
func genFn0() int8 {
    var v6 uint32 = 4294967295
    v6 = (v6 % (((v6 ^ v6)) | 1))
    _ = v6
    return 127
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 799
    var v3 int16 = -224
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: (v3 << ((730) & 63)), F1: 2147483647, F2: ("foo bar" + "")}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: ((("" + "\n\t") + "a") + "hello")}
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 = 50
        _ = len(v5)
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int32 = -1
            v13 := GenS3{F0: "foo bar"}
            if true {
                if (false || true) {
                    v14 := GenS1{F0: v3, F1: 25, F2: (("\n\t" + "\n\t") + (("hello" + ("\n\t" + "a")) + ""))}
                    v2 = v2
                    _ = v14
                } else {
                    v2 = (v2 + 255)
                    var v15 uint16 = v2
                    fmt.Println(fi11)
                    _ = v15
                }
                var v16 bool = (false || false)
                var v17 uint = 0
                _ = v16
                _ = v17
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v18 string = "hello"
    _ = len(v4)
    v3 -= -32768
    var v19 uint8 = 129
    v19 -= (v19 & 228)
    v20 := GenS4{F0: 42}
    switch "\n\t" {
    case "":
        switch 685 {
        case 197:
            if (141 <= 0) {
                switch -673 {
                case 42:
                    if true {
                        for fi21 := 0; fi21 < 8; fi21++ {
                            v19 = v19
                            v2 = uint16(v19)
                            _ = fi21
                        }
                        v2 -= (v2 | (v2 ^ (v2 % ((v2) | 1))))
                        switch "foo bar" {
                        case "foo bar":
                            var v22 int32 = -2147483648
                            _ = v22
                        case "":
                            var v23 int = -645
                            _ = v23
                        }
                    } else {
                        v19 += 94
                        v1 = append(v1, 310)
                    }
                case 127:
                    var v24 bool = true
                    _ = v24
                case 0:
                    fmt.Println(v3)
                default:
                    var v25 uint64 = 18446744073709551615
                    _ = v25
                }
                if (1 != 1) {
                    if (!true) {
                        var v26 float64 = 1e10
                        v2 = (^(v2))
                        for fi27 := 0; fi27 < 10; fi27++ {
                            fmt.Println(v2, v18, v19)
                            _ = fi27
                        }
                        _ = v26
                    } else {
                        _ = len(v1)
                    }
                    var v28 int = 96
                    switch "\n\t" {
                    case "\n\t":
                        var v29 float32 = (1.5 / 1.5)
                        _ = v29
                    }
                    _ = v28
                } else {
                    v30 := GenS1{F0: v3, F1: 1, F2: (("foo bar" + v18) + (v18 + v18))}
                    _ = v30
                }
            }
        case 32768:
            switch "" {
            case "\n\t":
                if false {
                    v31 := GenS2{F0: "foo bar"}
                    _ = v31
                } else {
                    if (v19 > v19) {
                        v19 = 42
                    } else {
                        _ = len(v1)
                        switch "hello" {
                        case "a":
                            fmt.Println(v2, v19, v18)
                        }
                        switch 869 {
                        case 1:
                            if (1.5 > 1.5) {
                                var v32 uint = 0
                                switch v18 {
                                case "hello":
                                    v32 -= v32
                                case "":
                                    for fi33 := 0; fi33 < 16; fi33++ {
                                        switch fi33 {
                                        case 127:
                                            var v34 bool = ((v2 + v2) == 2)
                                            _ = v34
                                        case 1:
                                            _ = len(v1)
                                        case 212:
                                            if true {
                                                if (v18 >= "foo bar") {
                                                    var v35 int64 = -1
                                                    v36 := GenS3{F0: v18}
                                                    _ = v35
                                                    _ = v36
                                                }
                                                var v37 uint8 = 42
                                                _ = v37
                                            } else {
                                                v38 := v4[(fi33 + 128)]
                                                _ = v38
                                            }
                                        }
                                        switch fi33 {
                                        case 32768:
                                            v18 = ""
                                        case 127:
                                            for fi39 := 0; fi39 < 11; fi39++ {
                                                v19 = (v19 + (v19 | v19))
                                                fmt.Println(fi33, v32)
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = 535
                                                    }
                                                    var v40 float32 = (3.14159 / (-2.25 * 3.14159))
                                                    v41 := GenS1{F0: 1000, F1: 2147483648, F2: v18}
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v32 = (v32 - 255)
                                                var v42 string = (("\n\t" + ((((v18 + ("" + (v18 + (((v18 + v18) + v18) + v18)))) + "\n\t") + "hello") + "foo bar")) + "hello")
                                                var v43 int = 536
                                                v44 := GenS2{F0: ((v42 + "hello") + v42)}
                                                var v45 bool = (0.1 > 3.14159)
                                                _ = fi39
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                        case 4294967295:
                                            var v46 float64 = (3.14159 + 0.1)
                                            _ = v46
                                        default:
                                            _ = len(v5)
                                        }
                                        _ = fi33
                                    }
                                case "foo bar":
                                    v3 += (-(v3))
                                }
                                _ = v32
                            }
                        }
                    }
                    v18 = (("a" + v18) + v18)
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    case "hello":
        for fi47 := 0; fi47 < 2; fi47++ {
            v2 = (v2 % ((65535) | 1))
            v3 += (-(v3))
            _ = fi47
        }
    }
    for fi48 := 0; fi48 < 15; fi48++ {
        if (!true) {
            v18 = v18
        } else {
            var v49 uint32 = 2
            switch "foo bar" {
            case "":
                if (!false) {
                    fmt.Println(v18)
                    v18 = (("\n\t" + v18) + v18)
                } else {
                    var v50 float32 = float32(float64((float32(3.14159) + (0.0 - -2.25))))
                    _ = v50
                }
            case "a":
                var v51 float32 = 0.0
                _ = v51
            default:
                v18 = (("a" + "hello") + v18)
            }
            _ = v49
        }
        var v52 int16 = (v3 ^ -738)
        delete(v4, (fi48 / ((226) | 1)))
        _ = fi48
        _ = v52
    }
    for fi53 := 0; fi53 < 10; fi53++ {
        for fi54 := 0; fi54 < 8; fi54++ {
            v3 = (v3 / ((v3) | 1))
            switch (fi54 / ((fi54) | 1)) {
            case -827:
                v18 = "hello"
            default:
                v18 = (("\n\t" + "\n\t") + "hello")
            }
            _ = fi54
        }
        _ = fi53
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        switch "a" {
        case "":
            v56 := GenS5{F0: 1, F1: false}
            _ = v56
        case "foo bar":
            for fi57 := 0; fi57 < 8; fi57++ {
                v2 = v2
                var v58 string = "\n\t"
                _ = fi57
                _ = v58
            }
        default:
            var v59 float32 = 1.5
            _ = v59
        }
        var v60 int8 = 127
        var v61 int32 = 42
        _ = fi55
        _ = v60
        _ = v61
    }
    v18 = ""
    for fi62 := 0; fi62 < 9; fi62++ {
        if (!true) {
            if false {
                if (100 != 81) {
                    var v63 bool = (!true)
                    for fi64 := 0; fi64 < 7; fi64++ {
                        if v63 {
                            var v65 uint16 = 42
                            v4[(fi64 | 9223372036854775807)] = fi64
                            if (0 > 9223372036854775807) {
                                switch v18 {
                                case "\n\t":
                                    switch v18 {
                                    case "a":
                                        v18 = v18
                                    default:
                                        v66 := GenS4{F0: 432}
                                        _ = len(v1)
                                        v67 := GenS5{F0: -1, F1: ((1 != -128) || v63)}
                                        _ = v66
                                        _ = v67
                                    }
                                case "":
                                    v68 := GenS1{F0: v3, F1: 2147483648, F2: ("" + v18)}
                                    _ = v68
                                case "a":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v69 uint32 = 150
                                v3 = (v3 % (((-(v3))) | 1))
                                v70 := GenS6{F0: v69}
                                _ = v69
                                _ = v70
                            }
                            _ = v65
                        } else {
                            var v71 float32 = 3.14159
                            v63 = (1.5 == v71)
                            var v72 string = (v18 + (("" + v18) + "a"))
                            _ = v71
                            _ = v72
                        }
                        v3 = -432
                        _ = fi64
                    }
                    if v63 {
                        v73 := GenS7{F0: 2147483647, F1: 139, F2: 849}
                        v74 := GenS8{F0: fi62}
                        _ = v73
                        _ = v74
                    }
                    _ = v63
                } else {
                    v75 := GenS7{F0: 2147483648, F1: 42, F2: v3}
                    var v76 uint64 = 61
                    _ = v75
                    _ = v76
                }
                if false {
                    v77 := GenS9{F0: genFn0(), F1: 1, F2: false}
                    _ = v77
                }
                var v78 int = -2147483648
                _ = v78
            }
            var v79 int16 = 1
            _ = v79
        }
        fmt.Println(v18, v18)
        var v80 int8 = 127
        _ = fi62
        _ = v80
    }
    v19 = (v19 % ((v19) | 1))
    v81 := GenS8{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v9
}
