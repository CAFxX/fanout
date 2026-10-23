package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 uint16 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint32; F1 uint32; F2 string }
type GenS8 struct { F0 int; F1 uint64; F2 int32 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint8 }
func genFn0(p0 int16, p1 int8) int64 {
    v7 := []int{-2147483648}
    var v8 int = 32768
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 573
}
func genFn1(p0 int16, p1 uint64, p2 uint64) int32 {
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint16 = 65535
        for fi11 := 0; fi11 < 20; fi11++ {
            if (127 < -128) {
                v12 := GenS1{F0: 127, F1: 200}
                _ = v12
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if false {
        v13 := []int{}
        switch "\n\t" {
        case "a":
            p0 = 1000
        case "\n\t":
            var v14 int32 = 33
            _ = v14
        case "foo bar":
            if false {
                if false {
                    switch -496 {
                    case -9223372036854775808:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case -2147483648:
                        if false {
                            p1 = p2
                            if false {
                                if false {
                                    p1 = 2
                                }
                            }
                            v15 := GenS1{F0: 221, F1: 4294967295}
                            _ = v15
                        }
                    case 42:
                        var v16 int = -2147483648
                        _ = v16
                    }
                    fmt.Println(p2, p0)
                } else {
                    if true {
                        var v17 int8 = 42
                        for fi18 := 0; fi18 < 18; fi18++ {
                            var v19 uint64 = (p2 | (^(p1)))
                            _ = fi18
                            _ = v19
                        }
                        p1 += p1
                        _ = v17
                    } else {
                        var v20 string = ("hello" + "a")
                        _ = len(v13)
                        for fi21 := 0; fi21 < 2; fi21++ {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                v23 := GenS1{F0: 223, F1: 255}
                                if false {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    fmt.Println(p0)
                                } else {
                                    fmt.Println(p1, fi22, v20)
                                    v24 := GenS1{F0: 2, F1: 42}
                                    var v25 int8 = 127
                                    _ = v24
                                    _ = v25
                                }
                                p2 = 1
                                switch 9223372036854775807 {
                                case 93:
                                    fmt.Println(fi21)
                                case 773:
                                    p1 += (p1 << ((18446744073709551615) & 63))
                                default:
                                    if (0.0 >= 0.1) {
                                        switch ("foo bar" + (v20 + v20)) {
                                        case "":
                                            var v26 int = fi22
                                            _ = v26
                                        }
                                        fmt.Println(fi21)
                                        fmt.Println(p0)
                                    }
                                }
                                v27 := GenS1{F0: 255, F1: 2}
                                _ = fi22
                                _ = v23
                                _ = v27
                            }
                            v28 := GenS2{F0: p2, F1: 1, F2: -32768}
                            for fi29 := 0; fi29 < 11; fi29++ {
                                v13 = append(v13, fi29)
                                v30 := GenS2{F0: 751, F1: 1, F2: (p0 - 32767)}
                                var v31 int64 = -1
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = fi21
                            _ = v28
                        }
                        _ = v20
                    }
                    if (((((1e10 - 1.5) / 1.0) / 1.0) * (((-2.25 * ((3.14159 - 0.1) / float32(((0.0 + (((((1.5 - (float64(3.14159) / 3.14159)) + (-2.25 - float64(1e10))) + (3.14159 + 0.0)) - ((0.1 - 0.0) - 3.14159)) / (-2.25 - 1.5))) - 1e10)))) * ((1.5 * -2.25) + (0.1 - 3.14159))) - 1.5)) < 3.14159) {
                        var v32 int8 = 42
                        _ = v32
                    } else {
                        var v33 uint = 42
                        var v34 int32 = 42
                        v35 := GenS3{F0: 923, F1: 146}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                var v36 uint16 = 0
                p2 -= p2
                _ = v36
            }
        default:
            switch ("hello" + "\n\t") {
            case "a":
                p0 = (p0 | (p0 ^ (p0 / ((p0) | 1))))
            case "":
                v13 = append(v13, 473)
            case "\n\t":
                v37 := GenS4{F0: 278, F1: 642}
                _ = v37
            }
        }
        v38 := GenS2{F0: (p1 | 42), F1: 65535, F2: p0}
        _ = v13
        _ = v38
    } else {
        p0 = 972
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 uint64 = 785
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    var v4 uint64 = v2
    v5 := make(map[string]string)
    var v6 uint = 390
    _ = len(v3)
    v39 := GenS5{F0: 1, F1: -650, F2: 42}
    for fi40 := 0; fi40 < 19; fi40++ {
        if true {
            for fi41 := 0; fi41 < 5; fi41++ {
                fmt.Println(fi41, v4, fi40)
                v42 := GenS3{F0: 4294967295, F1: 42}
                switch (fi41 + 2147483647) {
                case 4294967295:
                    v43 := GenS6{F0: -1, F1: 728}
                    _ = v43
                case 32767:
                    v44 := GenS7{F0: 2147483648, F1: 1, F2: "\n\t"}
                    _ = v44
                case -724:
                    v45 := []int{4294967295}
                    _ = v45
                default:
                    if false {
                        if false {
                            if ((2147483647 == 739) && false) {
                                var v46 int16 = -109
                                _ = v46
                            } else {
                                var v47 uint = (v6 & (v6 ^ (^(v6))))
                                switch "hello" {
                                case "a":
                                    var v48 uint64 = (v2 - 18446744073709551615)
                                    _ = v48
                                case "foo bar":
                                    var v49 int64 = -9223372036854775808
                                    _ = v49
                                case "":
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        switch (fi41 ^ (fi41 + fi40)) {
                                        case -1:
                                            if (32767 > -1) {
                                                v50 := GenS8{F0: fi41, F1: v2, F2: -2147483648}
                                                v2 = uint64(v6)
                                                var v51 int32 = 1
                                                _ = v50
                                                _ = v51
                                            }
                                        default:
                                            var v52 float64 = -2.25
                                            _ = v52
                                        }
                                    case "foo bar":
                                        var v53 uint = (^(v47))
                                        _ = v53
                                    }
                                }
                                _ = v47
                            }
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v6 += (^(v6))
                        case "hello":
                            v1 = (v1 * (v1 * 93))
                        case "\n\t":
                            fmt.Println(fi40, v4, v1)
                            if (float64(1.5) > 1.5) {
                                v54 := GenS4{F0: (v6 ^ v6), F1: 0}
                                _ = v54
                            }
                        }
                        v4 -= v4
                    }
                }
                _ = fi41
                _ = v42
            }
            var v55 uint8 = 2
            var v56 int64 = -658
            _ = v55
            _ = v56
        }
        _ = fi40
    }
    v6 += (v6 * 200)
    if true {
        var v57 int64 = 566
        var v58 int64 = (v57 * v57)
        _ = v57
        _ = v58
    }
    if false {
        v59 := v5["foo bar"]
        for fi60 := 0; fi60 < 1; fi60++ {
            var v61 int64 = -1
            for fi62 := 0; fi62 < 13; fi62++ {
                fmt.Println(v2, v2, v61)
                var v63 int8 = -1
                _ = fi62
                _ = v63
            }
            if ((v6 % (((v6 / ((613) | 1))) | 1)) > (^(v6))) {
                if true {
                    for fi64 := 0; fi64 < 18; fi64++ {
                        switch v59 {
                        case "a":
                            v59 = "\n\t"
                        case "hello":
                            v65 := GenS1{F0: 255, F1: (v6 / ((v6) | 1))}
                            _ = v65
                        default:
                            var v66 string = (v59 + v59)
                            _ = v66
                        }
                        _ = fi64
                    }
                } else {
                    v1 += 2
                    for fi67 := 0; fi67 < 6; fi67++ {
                        if false {
                            v4 = 1
                            if ((false || (true && true)) || (false || (true && true))) {
                                var v68 int64 = -1
                                if ((!(v61 == (-(v61)))) || true) {
                                    if (((2 == 255) || false) && false) {
                                        v5[v59] = "hello"
                                        var v69 uint = 130
                                        _ = v69
                                    }
                                }
                                v61 = v68
                                v70 := []int{(fi67 * (fi60 / ((364) | 1)))}
                                _ = v68
                                _ = v70
                            }
                        }
                        _ = fi67
                    }
                    for fi71 := 0; fi71 < 17; fi71++ {
                        v72 := GenS4{F0: v6, F1: 42}
                        _ = fi71
                        _ = v72
                    }
                }
            } else {
                if (0 < 0) {
                    v73 := v5[(v59 + ("foo bar" + v59))]
                    v4 = v1
                    _ = v73
                } else {
                    fmt.Println(v6)
                    delete(v3, "foo bar")
                    switch "hello" {
                    case "a":
                        fmt.Println(v61, v2, v2)
                    }
                }
                for fi74 := 0; fi74 < 6; fi74++ {
                    var v75 uint8 = 213
                    _ = fi74
                    _ = v75
                }
            }
            _ = fi60
            _ = v61
        }
        _ = v59
    }
    v3["hello"] = 692
    fmt.Println(v6)
    var v76 int8 = -93
    v1 -= (^(v4))
    var v77 int64 = -9223372036854775808
    v78 := GenS4{F0: (v6 / (((v6 | (^(v6)))) | 1)), F1: 65535}
    if false {
        switch -655 {
        case 379:
            var v79 string = "a"
            _ = v79
        case 127:
            v80 := GenS8{F0: int(v76), F1: v1, F2: -2147483648}
            _ = v80
        default:
            switch "foo bar" {
            case "foo bar":
                _ = len(v3)
            case "a":
                var v81 bool = false
                _ = v81
            default:
                v6 = 255
            }
        }
        switch "\n\t" {
        case "":
            v82 := GenS9{F0: 1, F1: (v1 & 587)}
            _ = v82
        }
        v83 := GenS2{F0: 536, F1: 2, F2: 0}
        _ = v83
    }
    v84 := GenS6{F0: (v76 - v76), F1: 305}
    v85 := GenS4{F0: v6, F1: 1}
    if true {
        if false {
            var v86 uint16 = 585
            for fi87 := 0; fi87 < 3; fi87++ {
                var v88 uint8 = 2
                for fi89 := 0; fi89 < 20; fi89++ {
                    var v90 float32 = ((0.1 * 1e10) + ((1e10 * ((-2.25 * -2.25) * 1.5)) * 0.1))
                    v91 := GenS10{F0: -1, F1: v88}
                    _ = fi89
                    _ = v90
                    _ = v91
                }
                _ = fi87
                _ = v88
            }
            _ = v86
        }
    } else {
        if (("\n\t" + ("foo bar" + "hello")) != "\n\t") {
            if false {
                var v92 int64 = (v77 - v77)
                v1 = v1
                _ = v92
            }
        } else {
            var v93 uint8 = 42
            _ = v93
        }
    }
    var v94 uint8 = 110
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v84
    _ = v85
    _ = v94
}
