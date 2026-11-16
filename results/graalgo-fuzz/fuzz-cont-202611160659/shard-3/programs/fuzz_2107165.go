package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
type GenS6 struct { F0 uint; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint64 = 498
    var v2 float64 = 0.0
    v3 := []string{(("foo bar" + "hello") + "a")}
    var v4 uint32 = 364
    v5 := []string{(("foo bar" + ("hello" + "foo bar")) + ((("" + "\n\t") + "foo bar") + ("" + "a"))), "hello", "foo bar"}
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 7; fi7++ {
        v1 = (v1 & v1)
        var v8 bool = ((("hello" + "hello") + ("a" + ("" + "foo bar"))) < "\n\t")
        var v9 int16 = -739
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = (("a" + ("foo bar" + "hello")) < ("foo bar" + "a"))
    if v10 {
        fmt.Println(v1, v10)
    }
    var v11 uint64 = v1
    v10 = (166 <= 913)
    var v12 int8 = -128
    var v13 int32 = 0
    var v14 int16 = -813
    switch 720 {
    case 0:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -712:
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 int64 = -9223372036854775808
            var v17 int16 = v14
            var v18 uint16 = 42
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    default:
        v19 := GenS1{F0: 2, F1: 350}
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint64 = (v11 | 914)
        _ = fi20
        _ = v21
    }
    var v22 int32 = (v13 >> ((2) & 63))
    if (2 >= 104) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v12 -= int8(v4)
    } else {
        var v23 int8 = (v12 * (v12 | (v12 << ((255) & 63))))
        _ = v23
    }
    if (!(v12 >= v12)) {
        var v24 uint64 = v11
        if v10 {
            v3 = append(v3, ("a" + ("\n\t" + "a")))
        } else {
            if v10 {
                v2 = 0.0
            }
        }
        if (9223372036854775807 == -1) {
            v25 := GenS1{F0: 255, F1: 65535}
            fmt.Println(v22, v13)
            v11 = 749
            _ = v25
        }
        _ = v24
    } else {
        for fi26 := 0; fi26 < 1; fi26++ {
            switch fi26 {
            case -907:
                var v27 int8 = 127
                _ = v27
            default:
                switch ("\n\t" + ("" + ("foo bar" + ("" + ("hello" + "a"))))) {
                case "foo bar":
                    v1 = v11
                }
            }
            if v10 {
                switch int(v22) {
                case 7:
                    v28 := GenS2{F0: v10, F1: 0}
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v10 = (81 > 236)
                        v30 := GenS3{F0: (v14 ^ (v14 << ((1) & 63))), F1: 4294967295}
                        if (false && ("a" <= "")) {
                            switch ("hello" + ("" + ("a" + ""))) {
                            case "hello":
                                switch (("foo bar" + "") + "") {
                                case "":
                                    var v31 int8 = (v12 + (v12 | v12))
                                    _ = v31
                                case "a":
                                    v32 := GenS4{F0: "a", F1: (-(v13)), F2: v12}
                                    _ = v32
                                default:
                                    _ = len(v5)
                                }
                            default:
                                fmt.Println(v12, v1, v22)
                            }
                            fmt.Println(v2, v22)
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                var v33 uint32 = 2147483648
                _ = v33
            }
            _ = fi26
        }
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        fmt.Println(v13, v22, v11)
        v10 = (v13 == v22)
        var v35 uint32 = (^(v4))
        _ = fi34
        _ = v35
    }
    v36 := v6["hello"]
    if v10 {
        v4 = v4
        if (v4 == v36) {
            v1 -= (v1 + 4294967295)
            v10 = (796 != 127)
            v37 := GenS1{F0: uint(v11), F1: 785}
            _ = v37
        } else {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            var v38 int64 = 0
            switch -2147483648 {
            case 0:
                for fi39 := 0; fi39 < 18; fi39++ {
                    v11 = v1
                    v40 := v6[("foo bar" + ("a" + ("hello" + "")))]
                    if v10 {
                        v2 = 3.14159
                    } else {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            switch (fi41 % (((fi41 ^ (fi39 * fi41))) | 1)) {
                            case 128:
                                var v42 uint = 65535
                                _ = v42
                            }
                            _ = fi41
                        }
                        fmt.Println(v4)
                    }
                    _ = fi39
                    _ = v40
                }
            case -9223372036854775808:
                if true {
                    v43 := GenS1{F0: 320, F1: 164}
                    if (v10 || v10) {
                        if false {
                            for fi44 := 0; fi44 < 12; fi44++ {
                                if false {
                                    if v10 {
                                        var v45 uint16 = 0
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 14; fi46++ {
                                        var v47 uint32 = v4
                                        _ = fi46
                                        _ = v47
                                    }
                                    switch (-(fi44)) {
                                    case 32767:
                                        for fi48 := 0; fi48 < 1; fi48++ {
                                            v49 := GenS2{F0: v10, F1: uint16(v11)}
                                            switch (fi44 / ((32768) | 1)) {
                                            case 32768:
                                                v50 := GenS5{F0: v12, F1: v10, F2: 42}
                                                if (v11 == (v1 | (v11 >> ((206) & 63)))) {
                                                    switch 9223372036854775807 {
                                                    case -329:
                                                        v4 -= 2147483648
                                                        if false {
                                                            fmt.Println(v2, v38, v4)
                                                        } else {
                                                            fmt.Println(v1)
                                                            v51 := GenS6{F0: 4294967295, F1: (v11 | 1), F2: (v38 >> ((1) & 63))}
                                                            v14 = 294
                                                            _ = v51
                                                        }
                                                    case 9223372036854775807:
                                                        if len(v3) > 0 {
                                                            v3[0] = ("\n\t" + "foo bar")
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v52 := GenS1{F0: 180, F1: 518}
                                                    var v53 uint16 = 0
                                                    var v54 bool = (!v10)
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v50
                                            default:
                                                var v55 float64 = 1e10
                                                var v56 int32 = 327
                                                v57 := GenS5{F0: v12, F1: (172 >= v38), F2: (v14 * 32767)}
                                                fmt.Println(v55, fi44)
                                                if v10 {
                                                    var v58 uint32 = v36
                                                    fmt.Println(v10, v1)
                                                    var v59 int64 = -11
                                                    if (false || (v36 >= v4)) {
                                                        switch (fi48 ^ (fi48 & (fi44 / (((fi48 >> ((4294967295) & 63))) | 1)))) {
                                                        case -2147483648:
                                                            var v60 float32 = 1.5
                                                            _ = v60
                                                        }
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    if (2 < 65535) {
                                                        var v61 bool = v10
                                                        _ = v61
                                                    } else {
                                                        v6[("" + "hello")] = (v36 << ((710) & 63))
                                                        v62 := GenS7{F0: 128, F1: 42}
                                                        fmt.Println(v14)
                                                        _ = v62
                                                    }
                                                    v63 := v6["a"]
                                                    v2 = float64(1.5)
                                                    _ = v63
                                                }
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            }
                                            v4 -= (v4 ^ (v36 % (((v4 % ((65535) | 1))) | 1)))
                                            _ = fi48
                                            _ = v49
                                        }
                                    }
                                } else {
                                    switch (fi44 >> ((18446744073709551615) & 63)) {
                                    case 45:
                                        for fi64 := 0; fi64 < 12; fi64++ {
                                            if v10 {
                                                v11 = (v11 - (v11 / ((v1) | 1)))
                                            } else {
                                                var v65 uint = 2
                                                _ = v65
                                            }
                                            _ = fi64
                                        }
                                    case -849:
                                        var v66 int16 = 42
                                        _ = v66
                                    case -2147483648:
                                        v5 = append(v5, "foo bar")
                                    }
                                    for fi67 := 0; fi67 < 8; fi67++ {
                                        fmt.Println(v22, v22)
                                        if v10 {
                                            fmt.Println(v2, v14)
                                            if v10 {
                                                if ((v12 << ((889) & 63)) < (v12 & v12)) {
                                                    switch (fi67 * (fi67 % (((fi44 >> ((1) & 63))) | 1))) {
                                                    case -1:
                                                        fmt.Println(v1, v13, fi44)
                                                    default:
                                                        if v10 {
                                                            v13 -= (v13 - v13)
                                                            v68 := v6[(("" + "hello") + ("a" + "\n\t"))]
                                                            var v69 bool = false
                                                            _ = v68
                                                            _ = v69
                                                        } else {
                                                            var v70 uint8 = 1
                                                            _ = v70
                                                        }
                                                    }
                                                    switch int(v14) {
                                                    case 959:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    default:
                                                        _ = len(v5)
                                                    }
                                                } else {
                                                    v36 += v4
                                                    var v71 string = "a"
                                                    _ = v71
                                                }
                                                var v72 bool = v10
                                                switch (fi44 ^ 2147483647) {
                                                case -1:
                                                    var v73 int8 = 0
                                                    _ = v73
                                                }
                                                _ = v72
                                            }
                                        }
                                        var v74 uint32 = v36
                                        _ = fi67
                                        _ = v74
                                    }
                                }
                                _ = fi44
                            }
                        }
                    }
                    _ = v43
                } else {
                    if (1 > 1) {
                        v75 := GenS2{F0: v10, F1: 1}
                        v1 -= (v1 - v1)
                        _ = v75
                    }
                    v76 := GenS7{F0: 32, F1: 0}
                    _ = v76
                }
            case -644:
                v6["\n\t"] = 4294967295
            default:
                if v10 {
                    v77 := GenS4{F0: "foo bar", F1: v22, F2: 121}
                    switch "a" {
                    case "a":
                        fmt.Println(v2, v13)
                    default:
                        v78 := GenS8{F0: (("\n\t" + "") + "a"), F1: v1}
                        _ = v78
                    }
                    _ = v77
                } else {
                    v79 := GenS6{F0: 82, F1: v11, F2: 9223372036854775807}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if v10 {
                        if v10 {
                            var v80 string = ""
                            delete(v6, (((v80 + v80) + v80) + "foo bar"))
                            fmt.Println(v13, v13)
                            _ = v80
                        }
                    } else {
                        _ = len(v6)
                        v81 := GenS5{F0: int8(v38), F1: (!v10), F2: 32767}
                        _ = v81
                    }
                    _ = v79
                }
            }
            _ = v38
        }
        v6[("" + "a")] = 2147483647
    } else {
        v4 = 246
        for fi82 := 0; fi82 < 5; fi82++ {
            fmt.Println(v2, v11, v22)
            switch "\n\t" {
            case "foo bar":
                for fi83 := 0; fi83 < 1; fi83++ {
                    var v84 string = "hello"
                    _ = fi83
                    _ = v84
                }
            case "hello":
                for fi85 := 0; fi85 < 1; fi85++ {
                    fmt.Println(v4, v13)
                    _ = fi85
                }
            default:
                v13 -= (v22 * (v22 * int32(v12)))
            }
            v86 := GenS6{F0: 2, F1: v11, F2: 0}
            _ = fi82
            _ = v86
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
