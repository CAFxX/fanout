package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 int32 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 string; F1 string }

func main() {
    var v1 uint32 = 756
    v2 := []int{9223372036854775807, 127, -2147483648, 42}
    var v3 int64 = 0
    v4 := make(map[string]int)
    var v5 float32 = (1.5 - 1.5)
    v6 := []string{"a", "", "\n\t", "hello"}
    switch ("foo bar" + "hello") {
    case "":
        switch 238 {
        case -431:
            if (!true) {
                for fi7 := 0; fi7 < 5; fi7++ {
                    v8 := GenS1{F0: fi7, F1: 237, F2: 42}
                    var v9 float64 = 1.5
                    for fi10 := 0; fi10 < 15; fi10++ {
                        if (-2147483648 != 1) {
                            v3 = 42
                        } else {
                            switch ("\n\t" + ("" + ("\n\t" + ("\n\t" + "hello")))) {
                            case "a":
                                switch fi7 {
                                case 32767:
                                    v11 := GenS2{F0: 18446744073709551615, F1: ("" + "")}
                                    _ = v11
                                }
                            case "\n\t":
                                v12 := GenS2{F0: 740, F1: ""}
                                _ = v12
                            default:
                                v13 := GenS3{F0: (-94 <= -232), F1: 964}
                                _ = v13
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        _ = fi10
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                for fi14 := 0; fi14 < 20; fi14++ {
                    if false {
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                    v16 := GenS3{F0: (650 >= 255), F1: 65535}
                    _ = fi14
                    _ = v16
                }
            } else {
                switch "a" {
                case "a":
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "\n\t":
                        v17 := GenS1{F0: 42, F1: uint8(v1), F2: -1}
                        _ = v17
                    default:
                        if (32767 >= -1) {
                            v18 := GenS3{F0: true, F1: 65535}
                            fmt.Println(v1)
                            _ = v18
                        }
                    }
                case "hello":
                    v19 := GenS4{F0: 9223372036854775807, F1: 42, F2: 1}
                    _ = v19
                case "foo bar":
                    var v20 uint = 1
                    _ = v20
                }
                if true {
                    v21 := GenS1{F0: 4294967295, F1: 0, F2: 0}
                    _ = v21
                }
                var v22 float64 = (((0.1 - 3.14159) - (0.0 * 0.1)) - 0.1)
                _ = v22
            }
        }
    case "\n\t":
        fmt.Println(v5, v1, v5)
    default:
        v23 := GenS2{F0: 0, F1: ("foo bar" + "foo bar")}
        _ = v23
    }
    fmt.Println(v3)
    fmt.Println(v1)
    if true {
        if false {
            switch -2147483648 {
            case 127:
                v24 := GenS5{F0: -1, F1: 32767}
                _ = v24
            case -2147483648:
                switch "" {
                case "\n\t":
                    fmt.Println(v1, v1)
                case "hello":
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 float32 = 0.0
                        _ = len(v2)
                        _ = fi25
                        _ = v26
                    }
                default:
                    v3 = (v3 / ((v3) | 1))
                }
            case -533:
                var v27 int64 = 134
                _ = v27
            default:
                var v28 string = ("foo bar" + "\n\t")
                _ = v28
            }
            switch 428 {
            case 32767:
                var v29 int = -9223372036854775808
                _ = v29
            default:
                v30 := v4["a"]
                _ = v30
            }
            for fi31 := 0; fi31 < 1; fi31++ {
                if false {
                    v1 = (v1 | 65)
                    var v32 uint = 139
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 bool = ((^(v1)) > v1)
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    var v35 uint64 = 310
                    for fi36 := 0; fi36 < 13; fi36++ {
                        var v37 string = (("foo bar" + ("a" + (("hello" + ("hello" + "foo bar")) + "hello"))) + (("" + ("\n\t" + ("hello" + "hello"))) + ("foo bar" + ((("\n\t" + "") + ("hello" + ("hello" + "a"))) + ("hello" + "foo bar")))))
                        v5 += v5
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
                var v38 float32 = v5
                for fi39 := 0; fi39 < 12; fi39++ {
                    v38 -= (((3.14159 * v38) + -2.25) + v5)
                    var v40 bool = true
                    if v40 {
                        if v40 {
                            _ = len(v6)
                        }
                    } else {
                        _ = len(v4)
                    }
                    if true {
                        fmt.Println(v1)
                    }
                    switch ("foo bar" + (("a" + "a") + "a")) {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        _ = len(v4)
                        var v41 bool = v40
                        _ = v41
                    case "foo bar":
                        var v42 float64 = 1e10
                        _ = v42
                    }
                    var v43 int8 = -128
                    var v44 uint64 = 2
                    _ = fi39
                    _ = v40
                    _ = v43
                    _ = v44
                }
                _ = fi31
                _ = v38
            }
        } else {
            v3 = v3
            if (-4 < 1) {
                v45 := GenS4{F0: 4294967295, F1: 2, F2: 0}
                var v46 uint16 = 2
                var v47 uint = 2
                _ = v45
                _ = v46
                _ = v47
            } else {
                _ = len(v2)
                for fi48 := 0; fi48 < 6; fi48++ {
                    var v49 bool = (!(!true))
                    var v50 float64 = (1e10 - (0.0 - 3.14159))
                    v3 -= v3
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                fmt.Println(v3, v5)
            }
            v5 = v5
        }
        v1 += v1
        for fi51 := 0; fi51 < 17; fi51++ {
            var v52 int64 = v3
            _ = fi51
            _ = v52
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            switch 528 {
            case 32767:
                for fi53 := 0; fi53 < 17; fi53++ {
                    for fi54 := 0; fi54 < 14; fi54++ {
                        v5 = v5
                        _ = fi54
                    }
                    switch ("\n\t" + ("\n\t" + "a")) {
                    case "":
                        var v55 uint64 = 2
                        _ = v55
                    case "\n\t":
                        _ = len(v4)
                    }
                    var v56 uint32 = v1
                    _ = fi53
                    _ = v56
                }
            case -1:
                for fi57 := 0; fi57 < 7; fi57++ {
                    fmt.Println(fi57)
                    if (18446744073709551615 < 255) {
                        var v58 uint16 = 1
                        _ = v58
                    } else {
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            for fi59 := 0; fi59 < 18; fi59++ {
                                v3 = int64(fi59)
                                _ = len(v4)
                                if false {
                                    var v60 string = "foo bar"
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v60
                                }
                                _ = fi59
                            }
                        }
                    }
                    v4[("a" + "a")] = 32767
                    _ = fi57
                }
            case 32768:
                v61 := GenS6{F0: "", F1: "a"}
                _ = v61
            default:
                v62 := GenS4{F0: 9223372036854775807, F1: 0, F2: -608}
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 5; fi63++ {
            for fi64 := 0; fi64 < 1; fi64++ {
                v1 += (v1 & v1)
                _ = fi64
            }
            _ = fi63
        }
        if true {
            var v65 int16 = -537
            var v66 int8 = 127
            for fi67 := 0; fi67 < 1; fi67++ {
                for fi68 := 0; fi68 < 10; fi68++ {
                    var v69 int16 = v65
                    var v70 uint16 = 255
                    v71 := GenS6{F0: "hello", F1: "\n\t"}
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
                var v72 uint16 = 1
                switch "hello" {
                case "":
                    var v73 float64 = 0.0
                    _ = v73
                case "foo bar":
                    var v74 uint16 = 1
                    _ = v74
                default:
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v1, v65)
                    case "a":
                        if false {
                            var v75 int8 = v66
                            v5 -= v5
                            _ = v75
                        }
                    case "hello":
                        _ = len(v4)
                    default:
                        v3 -= -9223372036854775808
                    }
                }
                _ = fi67
                _ = v72
            }
            _ = v65
            _ = v66
        } else {
            switch "foo bar" {
            case "":
                var v76 int64 = v3
                _ = v76
            default:
                var v77 uint = 2
                _ = v77
            }
            v5 -= float32(((1e10 / 1e10) / 1.5))
        }
    }
    for fi78 := 0; fi78 < 11; fi78++ {
        var v79 uint32 = (v1 | v1)
        v80 := GenS6{F0: "hello", F1: "foo bar"}
        v1 = v1
        _ = fi78
        _ = v79
        _ = v80
    }
    var v81 float64 = 1.5
    fmt.Println(v3, v1, v3)
    switch "foo bar" {
    case "a":
        v1 += (^(v1))
    case "foo bar":
        switch 127 {
        case 796:
            var v82 int64 = (v3 % ((v3) | 1))
            _ = v82
        default:
            fmt.Println(v81, v3, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
}
