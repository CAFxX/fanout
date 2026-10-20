package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (false || false)
    var v3 int64 = -9223372036854775808
    var v4 int8 = 1
    var v5 int8 = v4
    var v6 int8 = v5
    if v2 {
        v3 = (v3 & (-(v3)))
        v6 = v4
    } else {
        if (2 < 1) {
            fmt.Println(v3)
            v7 := v1[1]
            v8 := GenS1{F0: (v3 - v3), F1: 65535, F2: (!v2)}
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: (v3 - v3), F1: 581, F2: ((1 == -1) && true)}
            for fi10 := 0; fi10 < 16; fi10++ {
                if (v4 >= (-(v6))) {
                    switch "foo bar" {
                    case "\n\t":
                        v3 -= (v3 - (^(v3)))
                    case "hello":
                        v5 = v6
                    }
                } else {
                    v11 := GenS2{F0: (("" + "a") + ("" + ("hello" + "\n\t"))), F1: 502, F2: 805}
                    v12 := []int{4294967295}
                    if true {
                        v4 = 1
                        v13 := GenS3{F0: 1}
                        _ = v13
                    } else {
                        var v14 int32 = 42
                        switch ("" + (("foo bar" + ("hello" + "")) + "")) {
                        case "foo bar":
                            v15 := GenS4{F0: ""}
                            _ = v15
                        case "a":
                            var v16 bool = v2
                            _ = v16
                        case "hello":
                            var v17 uint16 = 65535
                            _ = v17
                        }
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
            _ = v9
        }
        fmt.Println(v3)
        var v18 string = "\n\t"
        _ = v18
    }
    if v2 {
        fmt.Println(v6, v4)
        var v19 int = -670
        var v20 uint16 = 32767
        _ = v19
        _ = v20
    }
    switch "a" {
    case "a":
        switch "\n\t" {
        case "\n\t":
            if (!v2) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 int32 = -727
                    _ = fi21
                    _ = v22
                }
            } else {
                _ = len(v1)
            }
        case "hello":
            v23 := GenS1{F0: v3, F1: 42, F2: (-2147483648 != -461)}
            _ = v23
        case "a":
            v24 := GenS5{F0: v2, F1: 901, F2: (v6 + -128)}
            _ = v24
        default:
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := GenS3{F0: 32767}
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := v1[58]
                    v4 = v4
                    if true {
                        fmt.Println(v6, v2)
                    } else {
                        var v29 uint64 = 18446744073709551615
                        _ = v29
                    }
                    v4 += (-(v6))
                    var v30 string = "a"
                    _ = fi27
                    _ = v28
                    _ = v30
                }
                switch "foo bar" {
                case "":
                    if false {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            var v32 uint = 238
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 10; fi33++ {
                            if (!v2) {
                                v5 = 127
                                switch ("foo bar" + "a") {
                                case "":
                                    switch ("hello" + ((("foo bar" + "") + ("\n\t" + "\n\t")) + ((("\n\t" + "foo bar") + ("\n\t" + ("a" + "foo bar"))) + ((("a" + ("foo bar" + "\n\t")) + ("\n\t" + (("\n\t" + "a") + (("hello" + ("foo bar" + ("a" + "a"))) + "a")))) + "")))) {
                                    case "hello":
                                        v34 := v1[104]
                                        _ = v34
                                    case "foo bar":
                                        var v35 uint64 = 42
                                        _ = v35
                                    case "\n\t":
                                        if (1 < (v6 % ((127) | 1))) {
                                            v6 += v4
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v6, v5)
                                }
                                if (!(878 == 2147483648)) {
                                    v36 := []int{fi25}
                                    var v37 int64 = v3
                                    v38 := GenS4{F0: "foo bar"}
                                    if v2 {
                                        if false {
                                            var v39 string = "hello"
                                            _ = v39
                                        } else {
                                            v3 = (-(v3))
                                        }
                                    }
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                v40 := GenS2{F0: "foo bar", F1: 1, F2: 65535}
                                var v41 int32 = -1
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS5{F0: v2, F1: 2147483647, F2: v5}
                                if (((0.1 / (1.5 + 0.0)) * (3.14159 - float64((1e10 - ((float32(0.0) / float32(1.5)) + 3.14159))))) >= (0.0 + ((3.14159 * -2.25) / float64(1.5)))) {
                                    v4 = 107
                                    if v2 {
                                        var v43 int64 = v3
                                        if (1 < -21) {
                                            v44 := GenS6{F0: -32768, F1: 128, F2: 2}
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch 127 {
                                                case 32768:
                                                    var v45 bool = ("\n\t" == "\n\t")
                                                    _ = v45
                                                case 127:
                                                    v46 := []int{fi33, fi33, fi25}
                                                    _ = v46
                                                default:
                                                    v3 -= (^(v3))
                                                }
                                            default:
                                                v47 := GenS4{F0: "hello"}
                                                _ = v47
                                            }
                                            switch "\n\t" {
                                            case "":
                                                if v2 {
                                                    if true {
                                                        var v48 int8 = (v5 >> ((4294967295) & 63))
                                                        v49 := GenS4{F0: "foo bar"}
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        v50 := GenS4{F0: "\n\t"}
                                                        if ((v43 ^ v3) >= v43) {
                                                            switch 1 {
                                                            case 9223372036854775807:
                                                                v51 := GenS7{F0: 2147483648}
                                                                _ = v51
                                                            case -1:
                                                                v52 := []int{(fi25 << ((18446744073709551615) & 63)), -9223372036854775808, (fi33 ^ fi33)}
                                                                _ = v52
                                                            case -2147483648:
                                                                var v53 int = fi25
                                                                _ = v53
                                                            }
                                                            v4 = (v6 / ((0) | 1))
                                                            if true {
                                                                var v54 uint = 2
                                                                _ = v54
                                                            }
                                                        }
                                                        _ = v50
                                                    }
                                                    v6 = v5
                                                    fmt.Println(v4)
                                                } else {
                                                    v3 = (v3 & v3)
                                                }
                                            }
                                            _ = v44
                                        }
                                        _ = v43
                                    } else {
                                        v4 += 1
                                        v55 := GenS6{F0: 1000, F1: 42, F2: 667}
                                        _ = v55
                                    }
                                    v3 -= -9223372036854775808
                                }
                                _ = v42
                            }
                            v2 = true
                            fmt.Println(v3, v6)
                            _ = fi33
                        }
                    } else {
                        var v56 int32 = 0
                        _ = v56
                    }
                case "hello":
                    var v57 float32 = (3.14159 * 3.14159)
                    _ = v57
                case "a":
                    var v58 string = "hello"
                    _ = v58
                }
                _ = fi25
                _ = v26
            }
        }
    }
    switch "hello" {
    case "":
        if (3.14159 > 1e10) {
            for fi59 := 0; fi59 < 10; fi59++ {
                fmt.Println(v6, v5, v5)
                v60 := []int{(fi59 + 20), (fi59 + fi59), fi59, (fi59 * fi59)}
                _ = fi59
                _ = v60
            }
            if (v2 && (!(127 == (-(v6))))) {
                var v61 uint64 = 586
                for fi62 := 0; fi62 < 10; fi62++ {
                    var v63 bool = (!(32767 >= 2))
                    v5 = v5
                    _ = fi62
                    _ = v63
                }
                var v64 int64 = v3
                _ = v61
                _ = v64
            } else {
                var v65 int8 = v4
                if v2 {
                    v3 = v3
                } else {
                    var v66 float64 = (3.14159 / (1.5 - float64(((((((1.5 * 1.5) + ((float32(1e10) + (-2.25 - 0.0)) / 1.5)) + ((float32(-2.25) / 3.14159) / 0.1)) * 3.14159) * (1.5 / -2.25)) / 3.14159))))
                    _ = v66
                }
                _ = v65
            }
            v67 := GenS5{F0: false, F1: 42, F2: 100}
            _ = v67
        }
    case "\n\t":
        fmt.Println(v6, v4, v6)
    default:
        for fi68 := 0; fi68 < 17; fi68++ {
            fmt.Println(v3, v3, fi68)
            fmt.Println(fi68, v6)
            _ = fi68
        }
    }
    for fi69 := 0; fi69 < 4; fi69++ {
        for fi70 := 0; fi70 < 9; fi70++ {
            switch ("\n\t" + "hello") {
            case "\n\t":
                if (3.14159 < ((((3.14159 + 0.0) / (1.5 - -2.25)) - -2.25) - 1e10)) {
                    v71 := GenS8{F0: 2}
                    v72 := GenS2{F0: ("" + "foo bar"), F1: -32768, F2: 1}
                    var v73 float64 = -2.25
                    _ = v71
                    _ = v72
                    _ = v73
                }
            case "hello":
                v5 = v5
            case "":
                var v74 uint = 1000
                _ = v74
            default:
                v2 = (!(-999 <= 1000))
            }
            _ = fi70
        }
        v75 := GenS9{F0: 72}
        _ = fi69
        _ = v75
    }
    for fi76 := 0; fi76 < 9; fi76++ {
        v5 -= v5
        _ = len(v1)
        fmt.Println(fi76)
        _ = fi76
    }
    v77 := GenS1{F0: v3, F1: 0, F2: (v2 || (2147483647 == 65535))}
    v4 -= v4
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v78 := GenS10{F0: 255, F1: (v4 | (v5 & -22))}
        _ = v78
    case "a":
        if (((0.0 + (3.14159 + ((0.0 * 1.5) - 1.5))) / 1.0) >= 3.14159) {
            var v79 bool = (false || false)
            _ = v79
        }
    case "hello":
        v80 := GenS7{F0: 0}
        _ = v80
    }
    v6 = -60
    var v81 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v81
}
