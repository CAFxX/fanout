package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint64; F1 uint8; F2 string }
type GenS11 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 int8) bool {
    var v7 uint64 = 18446744073709551615
    var v8 int8 = p0
    if true {
        for fi9 := 0; fi9 < 10; fi9++ {
            v8 += 42
            for fi10 := 0; fi10 < 3; fi10++ {
                if false {
                    p0 = (p0 >> ((418) & 63))
                    if false {
                        v8 += 100
                        var v11 uint = 4294967295
                        v12 := GenS1{F0: -841, F1: p0, F2: 42}
                        _ = v11
                        _ = v12
                    }
                    var v13 float32 = float32(float64(1.5))
                    _ = v13
                }
                var v14 uint32 = 42
                _ = fi10
                _ = v14
            }
            _ = fi9
        }
        fmt.Println(p0, v8, v7)
        var v15 int64 = 42
        _ = v15
    } else {
        var v16 uint64 = v7
        if false {
            p0 = 127
        }
        _ = v16
    }
    if true {
        if (!false) {
            for fi17 := 0; fi17 < 13; fi17++ {
                if (v7 <= v7) {
                    switch (fi17 >> ((2) & 63)) {
                    case -376:
                        v18 := []int{(fi17 << ((849) & 63)), fi17, (fi17 >> ((255) & 63)), (fi17 >> ((65535) & 63))}
                        _ = v18
                    case 2147483647:
                        v8 = -128
                    case -9223372036854775808:
                        v19 := GenS2{F0: "foo bar", F1: -9223372036854775808, F2: 100}
                        _ = v19
                    }
                }
                v20 := GenS1{F0: (fi17 % ((2147483647) | 1)), F1: p0, F2: 2147483647}
                if (false || (true && true)) {
                    if (false || true) {
                        v21 := GenS3{F0: uint(v7), F1: 0, F2: "\n\t"}
                        v22 := GenS2{F0: ("hello" + ((("hello" + "") + ("foo bar" + ("hello" + ("\n\t" + ("foo bar" + ""))))) + "hello")), F1: 887, F2: 127}
                        _ = v21
                        _ = v22
                    } else {
                        if ((1 > uint32(v7)) && true) {
                            if true {
                                fmt.Println(v7, fi17, v7)
                                var v23 uint64 = 434
                                if false {
                                    for fi24 := 0; fi24 < 10; fi24++ {
                                        if true {
                                            p0 = (p0 * 42)
                                            v7 = v7
                                        } else {
                                            v25 := GenS3{F0: 255, F1: 0, F2: "\n\t"}
                                            if true {
                                                v26 := GenS4{F0: (-(p0))}
                                                v27 := []int{(fi17 >> ((42) & 63)), (fi24 | (fi24 >> ((506) & 63)))}
                                                var v28 float32 = (float32((-2.25 * 0.1)) + 1e10)
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := GenS4{F0: (v8 << ((4294967295) & 63))}
                                            _ = v25
                                            _ = v29
                                        }
                                        v30 := []int{}
                                        _ = fi24
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 17; fi31++ {
                                        fmt.Println(fi31)
                                        v32 := []int{}
                                        _ = fi31
                                        _ = v32
                                    }
                                } else {
                                    switch (("a" + "\n\t") + ((((("hello" + "hello") + "\n\t") + "foo bar") + "hello") + ("hello" + "a"))) {
                                    case "\n\t":
                                        if true {
                                            var v33 int8 = (^(p0))
                                            _ = v33
                                        } else {
                                            v34 := []int{(fi17 << ((290) & 63)), (fi17 >> ((0) & 63)), -955}
                                            v35 := GenS2{F0: "foo bar", F1: -316, F2: 50}
                                            fmt.Println(v23, v23)
                                            _ = v34
                                            _ = v35
                                        }
                                    case "foo bar":
                                        var v36 string = "a"
                                        _ = v36
                                    case "a":
                                        var v37 bool = (2 >= 42)
                                        _ = v37
                                    }
                                    if (!(-1 <= -2147483648)) {
                                        var v38 uint8 = 42
                                        var v39 uint8 = (^(v38))
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = v23
                            }
                        } else {
                            switch 127 {
                            case 4294967295:
                                v40 := GenS1{F0: fi17, F1: p0, F2: -683}
                                _ = v40
                            case 32767:
                                for fi41 := 0; fi41 < 1; fi41++ {
                                    fmt.Println(fi41, v7, fi17)
                                    v42 := GenS3{F0: 2, F1: 653, F2: ("a" + ("\n\t" + "\n\t"))}
                                    fmt.Println(fi41, v8)
                                    _ = fi41
                                    _ = v42
                                }
                            default:
                                var v43 string = "hello"
                                _ = v43
                            }
                        }
                        if true {
                            p0 += int8(v7)
                            p0 = (v8 ^ v8)
                            for fi44 := 0; fi44 < 14; fi44++ {
                                if false {
                                    var v45 float32 = 1.5
                                    switch (fi17 / ((fi17) | 1)) {
                                    case 261:
                                        if (false || true) {
                                            v7 += (^(v7))
                                            v8 -= (p0 << ((65535) & 63))
                                            p0 -= p0
                                            v46 := GenS5{F0: 458, F1: -1}
                                            _ = v46
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                if (670 < 0) {
                                                    v47 := GenS6{F0: ((("hello" < "\n\t") || true) && true)}
                                                    if (937 >= 0) {
                                                        var v48 int32 = 622
                                                        v49 := GenS1{F0: -431, F1: v8, F2: v48}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    v45 = v45
                                                    _ = v47
                                                } else {
                                                    var v50 int = -9223372036854775808
                                                    _ = v50
                                                }
                                            default:
                                                var v51 uint8 = 0
                                                _ = v51
                                            }
                                        }
                                    case 1:
                                        var v52 bool = false
                                        _ = v52
                                    default:
                                        v53 := []int{1, 0, fi44}
                                        _ = v53
                                    }
                                    var v54 int32 = -733
                                    _ = v45
                                    _ = v54
                                } else {
                                    v7 += (v7 << ((255) & 63))
                                    switch "\n\t" {
                                    case "\n\t":
                                        v55 := []int{}
                                        _ = v55
                                    default:
                                        v8 = p0
                                    }
                                    fmt.Println(v8, p0, v8)
                                }
                                _ = fi44
                            }
                        } else {
                            v56 := GenS7{F0: "foo bar", F1: fi17}
                            for fi57 := 0; fi57 < 2; fi57++ {
                                var v58 int = fi17
                                var v59 int32 = 327
                                v60 := []int{fi17}
                                _ = fi57
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                            _ = v56
                        }
                    }
                    if false {
                        p0 += p0
                    }
                    p0 -= (p0 - p0)
                } else {
                    v61 := GenS4{F0: (v8 + v8)}
                    _ = v61
                }
                _ = fi17
                _ = v20
            }
            v62 := GenS8{F0: -1}
            _ = v62
        }
        for fi63 := 0; fi63 < 17; fi63++ {
            for fi64 := 0; fi64 < 16; fi64++ {
                v65 := GenS9{F0: 127, F1: 128}
                var v66 uint = 65535
                fmt.Println(fi64, p0, p0)
                fmt.Println(fi64, fi63)
                _ = fi64
                _ = v65
                _ = v66
            }
            var v67 int32 = 2147483647
            _ = fi63
            _ = v67
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return (v7 >= v7)
}
func genFn1(p0 int64, p1 int16, p2 uint16) int32 {
    p1 += (p1 >> ((4294967295) & 63))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 570
    v4 := make(map[string]uint32)
    var v5 int32 = 0
    var v6 float32 = float32(3.14159)
    for fi68 := 0; fi68 < 6; fi68++ {
        v6 = -2.25
        var v69 uint64 = 751
        for fi70 := 0; fi70 < 5; fi70++ {
            if (fi68 != (fi68 + fi68)) {
                var v71 int16 = v3
                var v72 uint16 = 14
                var v73 int8 = 0
                _ = v71
                _ = v72
                _ = v73
            } else {
                var v74 bool = (!false)
                v74 = v74
                var v75 int32 = (v5 * (v5 - (v5 & genFn1(int64(v5), -32768, 32767))))
                _ = v74
                _ = v75
            }
            var v76 uint16 = 1
            _ = len(v4)
            _ = fi70
            _ = v76
        }
        _ = fi68
        _ = v69
    }
    v77 := v1["hello"]
    var v78 int64 = 1
    var v79 uint32 = 1
    v80 := GenS10{F0: 4294967295, F1: 255, F2: "foo bar"}
    switch -1 {
    case -484:
        switch 32768 {
        case -9223372036854775808:
            v5 -= (v5 | (^(v5)))
        case 42:
            v3 = (v3 | (v3 / ((v3) | 1)))
        case 0:
            if (131 == 193) {
                v5 = (v5 & v5)
                if (!true) {
                    v78 = (v78 << ((uint(v79)) & 63))
                    v4["\n\t"] = v79
                    v81 := []int{-2147483648, 553, -486, 127}
                    _ = v81
                }
                switch 9223372036854775807 {
                case 42:
                    v82 := GenS7{F0: (("" + "a") + "foo bar"), F1: 4294967295}
                    _ = v82
                case 32768:
                    v83 := v2[42]
                    _ = v83
                case -98:
                    var v84 int = 42
                    _ = v84
                default:
                    v79 = 4294967295
                }
            }
        }
    case -9223372036854775808:
        var v85 int16 = 42
        _ = v85
    }
    v2[(v79 ^ v79)] = v79
    switch "foo bar" {
    case "foo bar":
        switch 741 {
        case 0:
            fmt.Println(v3, v5)
        case 823:
            if false {
                fmt.Println(v6, v3)
                var v86 uint = 726
                if v77 {
                    var v87 uint32 = (v79 & v79)
                    _ = v87
                } else {
                    var v88 int64 = v78
                    _ = v88
                }
                _ = v86
            } else {
                for fi89 := 0; fi89 < 3; fi89++ {
                    if v77 {
                        for fi90 := 0; fi90 < 14; fi90++ {
                            var v91 int16 = 42
                            if false {
                                v92 := GenS7{F0: "hello", F1: 42}
                                _ = v92
                            } else {
                                if ((0 == -128) && ((((3.14159 - 1.5) * v6) + -2.25) >= (0.0 / (1e10 + (v6 - v6))))) {
                                    var v93 uint16 = 625
                                    if false {
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            v94 := GenS11{F0: -32768, F1: (3.14159 == (0.0 / 1.5)), F2: v93}
                                            _ = v94
                                        case "hello":
                                            var v95 uint32 = (v79 << ((2) & 63))
                                            _ = v95
                                        default:
                                            if (v77 && ((((("\n\t" + "a") + "hello") + "\n\t") + "foo bar") < "\n\t")) {
                                                v96 := []int{(fi90 % (((^(fi89))) | 1)), fi90, fi90}
                                                fmt.Println(v6, v79)
                                                if len(v96) >= 2 {
                                                    v96 = v96[1:len(v96)]
                                                }
                                                _ = v96
                                            } else {
                                                v97 := []int{(fi89 / ((fi89) | 1)), 9223372036854775807, (fi90 % (((^(fi89))) | 1)), fi90}
                                                v98 := GenS2{F0: "", F1: 42, F2: 233}
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                        v79 += 271
                                        switch fi90 {
                                        case -629:
                                            fmt.Println(v3, v77, v91)
                                        case -9223372036854775808:
                                            var v99 int64 = (v78 + (v78 % ((v78) | 1)))
                                            _ = v99
                                        default:
                                            fmt.Println(v93, v3)
                                        }
                                    } else {
                                        v79 -= v79
                                        switch ("foo bar" + "") {
                                        case "":
                                            v79 = v79
                                        case "\n\t":
                                            fmt.Println(v93, v79)
                                        }
                                        v100 := []int{(-(fi90)), fi90}
                                        _ = v100
                                    }
                                    v91 = v91
                                    _ = v93
                                }
                            }
                            fmt.Println(v77, v3)
                            _ = fi90
                            _ = v91
                        }
                        v6 = -2.25
                    } else {
                        if false {
                            var v101 uint32 = v79
                            v101 = (v101 * (^(v79)))
                            _ = v101
                        }
                    }
                    v5 = v5
                    _ = fi89
                }
                v78 -= 9223372036854775807
                if false {
                    v6 = v6
                }
            }
        default:
            v79 -= (v79 >> ((390) & 63))
        }
    case "":
        _ = len(v1)
    case "\n\t":
        var v102 bool = ("foo bar" <= ("" + ("foo bar" + ((("foo bar" + "") + "") + "foo bar"))))
        _ = v102
    }
    fmt.Println(v79, v6, v5)
    var v103 float64 = float64(v6)
    var v104 uint32 = (v79 - 0)
    if v77 {
        for fi105 := 0; fi105 < 10; fi105++ {
            var v106 uint64 = 520
            _ = fi105
            _ = v106
        }
        v6 = v6
        v107 := v1["foo bar"]
        _ = v107
    }
    v108 := []int{}
    switch 736 {
    case 9223372036854775807:
        fmt.Println(v6)
    case 32768:
        v109 := GenS9{F0: 2147483647, F1: 127}
        _ = v109
    }
    var v110 uint64 = 18446744073709551615
    switch "" {
    case "hello":
        v5 += (v5 / (((^(v5))) | 1))
    case "foo bar":
        fmt.Println(v110, v79)
    default:
        v111 := GenS3{F0: 75, F1: (v104 & (v79 & v104)), F2: (("foo bar" + ((("" + "") + (("a" + "hello") + "foo bar")) + ("foo bar" + ""))) + ("a" + ("foo bar" + "hello")))}
        _ = v111
    }
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v110:", v110)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v103
    _ = v104
    _ = v108
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
