package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int8 }
func genFn0(p0 string, p1 int16, p2 string) bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        p2 = (p0 + ("foo bar" + "foo bar"))
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        switch (p2 + (p0 + "a")) {
        case "a":
            switch p2 {
            case "a":
                var v10 uint8 = 127
                _ = v10
            case "foo bar":
                if (float64(1.5) >= 0.1) {
                    var v11 int = (fi9 * fi9)
                    _ = v11
                }
            case "":
                fmt.Println(p0)
            }
        case "":
            for fi12 := 0; fi12 < 16; fi12++ {
                switch "" {
                case "a":
                    var v13 int32 = -1
                    _ = v13
                default:
                    v14 := []int{}
                    _ = v14
                }
                v15 := GenS1{F0: -1, F1: 1}
                p2 = ("a" + ("a" + ""))
                var v16 string = p0
                var v17 uint16 = 0
                _ = fi12
                _ = v15
                _ = v16
                _ = v17
            }
        case "hello":
            for fi18 := 0; fi18 < 20; fi18++ {
                if (0 > 42) {
                    p0 = p0
                    var v19 float32 = 1.5
                    v20 := GenS1{F0: -9223372036854775808, F1: 42}
                    _ = v19
                    _ = v20
                }
                var v21 int64 = 42
                fmt.Println(v21, fi9, fi9)
                _ = fi18
                _ = v21
            }
        }
        var v22 int16 = (p1 - p1)
        if (((((0.1 + (((((1.5 - 1e10) / float32(1e10)) * 0.1) - float32(0.1)) * -2.25)) - float32(3.14159)) - ((0.1 + (0.1 + (0.1 * -2.25))) * (float32(0.0) * 1.5))) / 1e10) == 0.1) {
            p2 = p0
        }
        _ = fi9
        _ = v22
    }
    v23 := GenS1{F0: 9223372036854775807, F1: 649}
    p2 = p0
    switch -472 {
    case 42:
        switch p2 {
        case "foo bar":
            var v24 bool = true
            _ = v24
        case "a":
            fmt.Println(p2)
        }
    case 4294967295:
        if false {
            p1 = (p1 << ((566) & 63))
            var v25 float32 = (float32(((0.0 - (-2.25 * float64(float32((0.1 * 1.5))))) - -2.25)) + ((float32(1.5) * 3.14159) * 0.0))
            _ = v25
        } else {
            var v26 uint32 = 2
            fmt.Println(p2, p2)
            if false {
                var v27 int16 = (p1 >> ((65535) & 63))
                for fi28 := 0; fi28 < 5; fi28++ {
                    for fi29 := 0; fi29 < 7; fi29++ {
                        p0 = p0
                        _ = fi29
                    }
                    p1 = 32767
                    _ = fi28
                }
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 string = p0
                    var v32 uint64 = 4294967295
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v27
            }
            _ = v26
        }
    case 32768:
        if true {
            var v33 uint = 65535
            v33 = v33
            _ = v33
        } else {
            v34 := GenS1{F0: -997, F1: 4294967295}
            _ = v34
        }
    default:
        p2 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return false
}

func main() {
    var v1 bool = true
    var v2 int = -904
    v3 := []string{"a", "\n\t"}
    var v4 uint = 4294967295
    v5 := []uint32{4294967295, 4294967295, 2147483647}
    v6 := make(map[uint32]string)
    switch 1 {
    case -2147483648:
        v35 := GenS2{F0: 0}
        _ = v35
    case -912:
        var v36 int8 = -1
        _ = v36
    default:
        for fi37 := 0; fi37 < 11; fi37++ {
            for fi38 := 0; fi38 < 6; fi38++ {
                switch fi38 {
                case 2147483647:
                    var v39 float64 = (1e10 * 3.14159)
                    _ = v39
                case 42:
                    var v40 float32 = 0.0
                    _ = v40
                case -9223372036854775808:
                    v4 = (v4 % (((^(v4))) | 1))
                default:
                    var v41 uint = (v4 * 42)
                    _ = v41
                }
                fmt.Println(v1)
                _ = fi38
            }
            _ = fi37
        }
    }
    v2 -= v2
    v42 := GenS3{F0: 0, F1: (v4 - (v4 >> ((v4) & 63)))}
    var v43 float32 = 1e10
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 string = "\n\t"
        if false {
            if (v45 >= ("\n\t" + (v45 + ((v45 + "") + v45)))) {
                if (!false) {
                    var v46 int32 = -747
                    switch 127 {
                    case -2147483648:
                        v45 = (v45 + "foo bar")
                    }
                    var v47 bool = v1
                    _ = v46
                    _ = v47
                } else {
                    v4 += (v4 << ((v4) & 63))
                    v48 := GenS3{F0: 264, F1: v4}
                    delete(v6, 42)
                    _ = v48
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = 385
                }
            }
            v49 := GenS4{F0: (v45 + "")}
            if (1 < -989) {
                for fi50 := 0; fi50 < 16; fi50++ {
                    v45 = v45
                    var v51 bool = false
                    v52 := GenS1{F0: 1, F1: 0}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            } else {
                fmt.Println(v45, v45, v45)
                var v53 uint = 554
                v4 += (v4 ^ (v53 ^ 255))
                _ = v53
            }
            _ = v49
        } else {
            if (!(-624 > 1)) {
                v2 += fi44
                switch (((((v45 + "foo bar") + ((v45 + v45) + ((v45 + (v45 + v45)) + "foo bar"))) + "foo bar") + v45) + "hello") {
                case "foo bar":
                    v2 = (fi44 << ((v4) & 63))
                case "":
                    if false {
                        v1 = true
                        for fi54 := 0; fi54 < 3; fi54++ {
                            v55 := v6[42]
                            switch v2 {
                            case 42:
                                if ((v4 >> ((v4) & 63)) >= 18) {
                                    var v56 int32 = 0
                                    _ = v56
                                } else {
                                    if ((fi44 & (fi44 % (((fi54 >> ((969) & 63))) | 1))) > (fi54 - fi44)) {
                                        if (v1 || (1e10 == (3.14159 / (1e10 * (0.0 + (-2.25 - (0.0 / 1.0))))))) {
                                            v45 = v45
                                            v55 = v55
                                        } else {
                                            var v57 string = v55
                                            v55 = v45
                                            _ = v57
                                        }
                                        v2 -= -507
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v58 int8 = 81
                                        _ = v58
                                    } else {
                                        fmt.Println(fi44)
                                        var v59 int64 = -9223372036854775808
                                        _ = v59
                                    }
                                }
                            case 234:
                                v55 = "foo bar"
                            case 2147483647:
                                _ = len(v5)
                            }
                            v1 = (857 >= 4294967295)
                            fmt.Println(v4, v4, v2)
                            v55 = v55
                            _ = fi54
                            _ = v55
                        }
                    }
                case "hello":
                    var v60 uint32 = uint32(v4)
                    _ = v60
                default:
                    if v1 {
                        for fi61 := 0; fi61 < 8; fi61++ {
                            v1 = (989 == 42)
                            if v1 {
                                v6[2147483648] = (v45 + ("hello" + v45))
                                if true {
                                    v62 := v6[565]
                                    v63 := GenS5{F0: 722}
                                    _ = len(v6)
                                    _ = v62
                                    _ = v63
                                } else {
                                    if (!v1) {
                                        if false {
                                            v64 := GenS1{F0: 1, F1: 922}
                                            if (((-1 != -128) && v1) || true) {
                                                fmt.Println(fi44, v45, v43)
                                                fmt.Println(fi44, v4)
                                                v43 -= ((1.5 + 0.0) + (v43 * (v43 - ((v43 + v43) - v43))))
                                                var v65 uint64 = 2
                                                if v1 {
                                                    v65 = (v65 / (((v65 + (^(v65)))) | 1))
                                                    _ = len(v6)
                                                    var v66 float32 = v43
                                                    fmt.Println(v65, v65, v1)
                                                    if v1 {
                                                        var v67 int8 = -21
                                                        var v68 int8 = int8(v65)
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    switch (v2 ^ fi44) {
                                                    case -1:
                                                        delete(v6, uint32(v65))
                                                    }
                                                    if v1 {
                                                        delete(v6, 2147483648)
                                                        var v69 uint16 = 42
                                                        v70 := GenS3{F0: 65535, F1: v4}
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    _ = v66
                                                }
                                                _ = v65
                                            }
                                            _ = v64
                                        }
                                        switch (v45 + "") {
                                        case "hello":
                                            var v71 float32 = v43
                                            _ = v71
                                        }
                                        if ((1 >= 529) || ((true && v1) || (0 == -9223372036854775808))) {
                                            if (-640 == -535) {
                                                switch "a" {
                                                case "":
                                                    v1 = false
                                                default:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                delete(v6, 2147483647)
                                            }
                                            if (0 < -698) {
                                                fmt.Println(fi44, v45)
                                                fmt.Println(v45)
                                                v72 := GenS4{F0: "foo bar"}
                                                _ = v72
                                            }
                                        } else {
                                            var v73 int = fi44
                                            v45 = (v45 + v45)
                                            fmt.Println(fi44, v45)
                                            _ = v73
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    if false {
                                        switch ((v45 + "a") + (v45 + v45)) {
                                        case "":
                                            if true {
                                                if v1 {
                                                    if (!(42 > 1)) {
                                                        v45 = ""
                                                        var v74 int32 = 2147483647
                                                        fmt.Println(v74, v74)
                                                        _ = v74
                                                    }
                                                    v5 = append(v5, 1)
                                                    delete(v6, 65535)
                                                    var v75 int8 = -21
                                                    _ = v75
                                                } else {
                                                    var v76 int8 = -61
                                                    v77 := GenS3{F0: 65535, F1: (v4 | (v4 << ((171) & 63)))}
                                                    _ = v76
                                                    _ = v77
                                                }
                                                fmt.Println(fi44, v45)
                                                v78 := GenS5{F0: 0}
                                                _ = v78
                                            } else {
                                                v79 := GenS6{F0: fi44, F1: -2147483648}
                                                v2 = (fi61 - 2147483647)
                                                v4 = 1
                                                _ = v79
                                            }
                                        default:
                                            var v80 int64 = 9223372036854775807
                                            _ = v80
                                        }
                                    }
                                    if v1 {
                                        if (v1 && (true && (2 == 18))) {
                                            fmt.Println(v43, v2, v45)
                                        }
                                        var v81 uint16 = 573
                                        switch "hello" {
                                        case "hello":
                                            v82 := GenS2{F0: 65535}
                                            _ = v82
                                        case "":
                                            if v1 {
                                                _ = len(v3)
                                            }
                                        case "a":
                                            v5 = append(v5, 2147483648)
                                        }
                                        _ = v81
                                    } else {
                                        v83 := GenS6{F0: fi61, F1: 42}
                                        _ = v83
                                    }
                                    v84 := GenS7{F0: 29}
                                    _ = v84
                                }
                                var v85 bool = (!(((1 == 817) && v1) || v1))
                                _ = v85
                            }
                            _ = fi61
                        }
                        fmt.Println(v2, v45, fi44)
                    }
                }
            }
        }
        _ = fi44
        _ = v45
    }
    v5 = append(v5, 891)
    var v86 uint8 = 127
    _ = len(v6)
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v4)
        var v87 uint64 = 18446744073709551615
        _ = v87
    }
    v88 := GenS2{F0: 4294967295}
    v89 := GenS1{F0: 1, F1: 4294967295}
    if false {
        if (v1 || true) {
            var v90 bool = (648 >= -629)
            _ = v90
        } else {
            var v91 float32 = v43
            v2 -= 42
            _ = v91
        }
        var v92 uint16 = 32767
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v2 -= (-(v2))
        case "\n\t":
            for fi93 := 0; fi93 < 20; fi93++ {
                v86 += (v86 & v86)
                _ = fi93
            }
        default:
            fmt.Println(v43, v43, v86)
        }
        _ = v92
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1, v1, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v86
    _ = v88
    _ = v89
}
