package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int16; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16; F1 string; F2 int8 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int64; F1 int16 }
func genFn0(p0 string, p1 bool) bool {
    switch "" {
    case "foo bar":
        v4 := GenS1{F0: 2147483647, F1: 42, F2: 2147483647}
        _ = v4
    case "\n\t":
        if p1 {
            p0 = p0
        }
    }
    if p1 {
        v5 := []int{128, 717, -95, 32768}
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 0, F1: 562, F2: 2147483647}
            if len(v5) > 0 {
                v5[0] = (fi6 - fi6)
            }
            switch p0 {
            case "a":
                p1 = p1
            case "hello":
                v8 := GenS1{F0: 0, F1: -163, F2: 834}
                _ = v8
            default:
                p0 = p0
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        p1 = (0.0 < 3.14159)
        v9 := GenS1{F0: 589, F1: 9223372036854775807, F2: 65535}
        v10 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 70}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return (0 < 126)
}
func genFn1(p0 string) uint32 {
    if true {
        var v11 uint16 = 1
        var v12 int16 = 1
        var v13 uint16 = (v11 >> ((93) & 63))
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := GenS2{F0: 4294967295, F1: 32768, F2: 2}
        for fi15 := 0; fi15 < 17; fi15++ {
            if false {
                p0 = "hello"
            } else {
                p0 = "\n\t"
            }
            fmt.Println(fi15)
            var v16 string = "a"
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v17 := []int{}
    p0 = ((("\n\t" + p0) + "hello") + "")
    var v18 bool = (0 <= -32768)
    _ = p0
    _ = v17
    _ = v18
    return 330
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 42
    v3 := []int{-2147483648, 0, -644, 32767}
    fmt.Println(v2)
    v2 = v2
    if false {
        v2 -= (v2 & v2)
        for fi19 := 0; fi19 < 7; fi19++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            if (42 >= -32768) {
                if len(v3) > 0 {
                    v3[0] = (^(fi20))
                }
                v2 += (-(v2))
            }
            _ = fi20
        }
    } else {
        var v21 uint = 901
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v22 := GenS2{F0: (^(v21)), F1: 2147483647, F2: uint16(v21)}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        switch -1 {
        case 128:
            var v24 bool = true
            _ = v24
        case 0:
            if ((574 == 37) && false) {
                var v25 uint16 = 1
                _ = v25
            }
        default:
            v2 = 127
        }
        _ = fi23
    }
    switch ("a" + ("" + "\n\t")) {
    case "\n\t":
        switch -9223372036854775808 {
        case 4294967295:
            var v26 float32 = (-2.25 * float32(1e10))
            _ = v26
        }
    case "foo bar":
        if true {
            if len(v3) > 0 {
                v3[0] = 45
            }
            v27 := GenS3{F0: 42}
            _ = v27
        }
    default:
        if false {
            switch ("hello" + "") {
            case "foo bar":
                v2 -= -128
            default:
                fmt.Println(v2)
            }
            v2 = v2
        } else {
            switch "foo bar" {
            case "a":
                v28 := GenS3{F0: -9223372036854775808}
                _ = v28
            case "\n\t":
                var v29 int = 32767
                _ = v29
            case "hello":
                var v30 bool = (true && false)
                _ = v30
            default:
                switch "hello" {
                case "foo bar":
                    var v31 float32 = (1e10 + (1.5 / (1.5 + float32(1.5))))
                    _ = v31
                default:
                    v2 = v2
                }
            }
            v32 := GenS3{F0: -1}
            if true {
                if false {
                    v2 = 127
                }
                v33 := GenS4{F0: 1, F1: -68, F2: "\n\t"}
                _ = v33
            } else {
                for fi34 := 0; fi34 < 15; fi34++ {
                    v35 := GenS5{F0: 42}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi34
                    _ = v35
                }
                if false {
                    v2 += -1
                    v2 += v2
                } else {
                    v36 := GenS5{F0: 25}
                    fmt.Println(v2)
                    if true {
                        if false {
                            switch -2147483648 {
                            case 32768:
                                for fi37 := 0; fi37 < 10; fi37++ {
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = (-(fi37))
                                        }
                                        var v38 bool = ((1 >= 0) || false)
                                        v38 = (307 < 65535)
                                        _ = v38
                                    }
                                    v2 = v2
                                    switch fi37 {
                                    case -1:
                                        for fi39 := 0; fi39 < 16; fi39++ {
                                            v40 := GenS2{F0: 378, F1: (fi37 >> ((255) & 63)), F2: 255}
                                            v2 -= -3
                                            _ = fi39
                                            _ = v40
                                        }
                                    case -9223372036854775808:
                                        if (((-2.25 + 1e10) / 1.0) == 1e10) {
                                            if (!(!(-1 < -773))) {
                                                for fi41 := 0; fi41 < 4; fi41++ {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v42 uint = 363
                                                        _ = v42
                                                    case "":
                                                        var v43 uint16 = 333
                                                        _ = v43
                                                    case "\n\t":
                                                        if false {
                                                            v1[fi41] = false
                                                            v44 := GenS5{F0: 42}
                                                            _ = v44
                                                        }
                                                    }
                                                    v2 = v2
                                                    v45 := GenS1{F0: 556, F1: 9223372036854775807, F2: 2147483648}
                                                    _ = fi41
                                                    _ = v45
                                                }
                                                v46 := GenS6{F0: 32767, F1: true, F2: v2}
                                                fmt.Println(v2)
                                                _ = v46
                                            } else {
                                                v2 = 42
                                                fmt.Println(fi37)
                                                v2 += int8(fi37)
                                            }
                                            v2 -= 9
                                            fmt.Println(v2)
                                        } else {
                                            if ((v2 - v2) != v2) {
                                                v2 = (v2 & (v2 / (((v2 % (((v2 | v2)) | 1))) | 1)))
                                                v47 := GenS5{F0: 930}
                                                v2 = (v2 & v2)
                                                _ = v47
                                            } else {
                                                v2 = (v2 / (((v2 ^ (v2 * v2))) | 1))
                                                var v48 float64 = 1.5
                                                for fi49 := 0; fi49 < 14; fi49++ {
                                                    v2 += v2
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = 348
                                                        }
                                                    }
                                                    _ = fi49
                                                }
                                                _ = v48
                                            }
                                            v2 = v2
                                            switch "a" {
                                            case "foo bar":
                                                v2 -= 127
                                            case "":
                                                switch 127 {
                                                case 4294967295:
                                                    if (!(0.1 != 3.14159)) {
                                                        for fi50 := 0; fi50 < 12; fi50++ {
                                                            v3 = append(v3, (^(fi50)))
                                                            var v51 string = "\n\t"
                                                            if false {
                                                                _ = len(v1)
                                                                var v52 uint32 = 2147483648
                                                                _ = v52
                                                            } else {
                                                                v53 := GenS5{F0: 2}
                                                                _ = v53
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v51 = ("foo bar" + "")
                                                            _ = fi50
                                                            _ = v51
                                                        }
                                                        if (!(1 == 1)) {
                                                            if (255 > 255) {
                                                                v2 = v2
                                                            }
                                                            v1[-1] = (!false)
                                                            v54 := GenS3{F0: 42}
                                                            _ = v54
                                                        } else {
                                                            v55 := GenS7{F0: 1, F1: ("a" + ("foo bar" + "")), F2: v2}
                                                            _ = len(v1)
                                                            fmt.Println(v2, fi37)
                                                            _ = v55
                                                        }
                                                        for fi56 := 0; fi56 < 6; fi56++ {
                                                            if ((true || (false && (0.1 != 0.0))) && true) {
                                                                var v57 float32 = (((((0.0 - 0.0) - (float32(-2.25) + 0.1)) - -2.25) * (((-2.25 / (1.5 / (1e10 - 1.5))) * 0.0) * -2.25)) - (float32(((3.14159 + -2.25) / 1.0)) * -2.25))
                                                                if (!false) {
                                                                    v58 := v1[128]
                                                                    fmt.Println(fi37, v58)
                                                                    _ = v58
                                                                } else {
                                                                    _ = len(v1)
                                                                    if true {
                                                                        var v59 int8 = -1
                                                                        v60 := GenS8{F0: (("a" + "\n\t") + "hello"), F1: 12}
                                                                        v57 += v57
                                                                        var v61 int32 = -290
                                                                        fmt.Println(v61, v61, v61)
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                    v1[(fi56 ^ (fi56 + 42))] = true
                                                                }
                                                                v57 = (v57 / (((((1.5 / (v57 * ((0.1 + (float32(-2.25) + (3.14159 + (0.1 + (1.5 + float32(1e10)))))) + v57))) * v57) + v57) + 0.0) * v57))
                                                                var v62 string = "hello"
                                                                var v63 int64 = 0
                                                                _ = v57
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            _ = fi56
                                                        }
                                                    } else {
                                                        for fi64 := 0; fi64 < 19; fi64++ {
                                                            if ((0 < 622) || (!((!true) || false))) {
                                                                v65 := GenS9{F0: 42, F1: (fi64 + fi64)}
                                                                v2 = 100
                                                                _ = v65
                                                            }
                                                            var v66 uint64 = 18446744073709551615
                                                            _ = fi64
                                                            _ = v66
                                                        }
                                                        v1[(fi37 % ((42) | 1))] = (!false)
                                                    }
                                                case 0:
                                                    if false {
                                                        for fi67 := 0; fi67 < 12; fi67++ {
                                                            var v68 int64 = -9223372036854775808
                                                            v69 := GenS8{F0: "hello", F1: 1}
                                                            switch ("hello" + "a") {
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v68 = (-(v68))
                                                                default:
                                                                    var v70 float64 = 0.0
                                                                    _ = v70
                                                                }
                                                            case "a":
                                                                if (!false) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    fmt.Println(fi67, fi67, fi37)
                                                                    v71 := GenS8{F0: "a", F1: 301}
                                                                    _ = v71
                                                                }
                                                            default:
                                                                v2 += v2
                                                            }
                                                            _ = fi67
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        if false {
                                                            v2 = (-(v2))
                                                            v72 := GenS10{F0: ((0.0 + float64(1.5)) != 3.14159)}
                                                            var v73 int8 = (v2 << ((859) & 63))
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            v74 := GenS4{F0: -518, F1: 1000, F2: ("a" + (("\n\t" + ((("foo bar" + ("" + "hello")) + "") + ("a" + "foo bar"))) + "\n\t"))}
                                                            _ = v74
                                                        case "":
                                                            v75 := GenS11{F0: 18446744073709551615, F1: "hello"}
                                                            _ = v75
                                                        }
                                                    } else {
                                                        if false {
                                                            var v76 bool = (!true)
                                                            _ = v76
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    case 587:
                                        for fi77 := 0; fi77 < 1; fi77++ {
                                            var v78 uint16 = 2
                                            fmt.Println(fi37, v2)
                                            if false {
                                                var v79 uint = 2
                                                v79 = (v79 / ((v79) | 1))
                                                fmt.Println(v78)
                                                fmt.Println(fi77, v79, v2)
                                                _ = v79
                                            } else {
                                                v80 := GenS4{F0: 819, F1: -766, F2: "hello"}
                                                v78 -= (v78 | v78)
                                                var v81 bool = false
                                                _ = v80
                                                _ = v81
                                            }
                                            var v82 uint = 833
                                            _ = fi77
                                            _ = v78
                                            _ = v82
                                        }
                                    default:
                                        var v83 int64 = 0
                                        _ = v83
                                    }
                                    _ = fi37
                                }
                            case 2147483647:
                                v84 := GenS12{F0: (v2 | -1)}
                                _ = v84
                            case -2147483648:
                                if (1 <= 4294967295) {
                                    _ = len(v3)
                                    v85 := GenS13{F0: 42, F1: 1}
                                    _ = v85
                                } else {
                                    var v86 uint32 = 65535
                                    fmt.Println(v86)
                                    _ = v86
                                }
                            default:
                                fmt.Println(v2)
                            }
                            var v87 uint8 = 255
                            v88 := GenS14{F0: v2}
                            _ = v87
                            _ = v88
                        }
                    } else {
                        fmt.Println(v2)
                    }
                    _ = v36
                }
                v2 += (-(v2))
            }
            _ = v32
        }
    }
    v89 := GenS7{F0: 255, F1: "", F2: -107}
    fmt.Println(v2)
    for fi90 := 0; fi90 < 10; fi90++ {
        v2 = (v2 - (v2 * 78))
        _ = fi90
    }
    v91 := v1[127]
    _ = len(v3)
    v92 := GenS15{F0: 42}
    v93 := GenS16{F0: 0, F1: 1000}
    var v94 int16 = 1000
    switch -137 {
    case -807:
        switch 0 {
        case -2147483648:
            var v95 int = 597
            _ = v95
        case 0:
            fmt.Println(v2, v2)
        case 1:
            var v96 int64 = 42
            _ = v96
        }
    case 9223372036854775807:
        v91 = (396 <= 0)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v91:", v91)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v89
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
