package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }
type GenS5 struct { F0 bool; F1 string; F2 int8 }
type GenS6 struct { F0 string; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int8; F1 int32; F2 string }
type GenS9 struct { F0 int; F1 uint8 }
func genFn0(p0 uint32) uint8 {
    switch "" {
    case "hello":
        if true {
            p0 = (p0 << ((255) & 63))
            fmt.Println(p0)
        } else {
            var v7 int64 = 1
            p0 -= p0
            _ = v7
        }
    case "foo bar":
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 = 42
            fmt.Println(p0, p0)
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 float32 = (3.14159 * 0.1)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case "a":
        fmt.Println(p0)
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 17; fi11++ {
        p0 += 2
        if (!(2147483647 > -2147483648)) {
            v12 := []int{}
            var v13 int8 = 1
            _ = v12
            _ = v13
        } else {
            var v14 uint32 = p0
            v14 += 1
            v14 = p0
            _ = v14
        }
        _ = fi11
    }
    _ = p0
    return 1
}
func genFn1(p0 bool) int {
    v15 := GenS1{F0: 4294967295, F1: 86}
    switch -33 {
    case -423:
        var v16 uint8 = 0
        _ = v16
    case 32767:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint = 255
            v19 := []int{(fi17 - fi17), fi17, (fi17 >> ((v18) & 63))}
            if (169 > genFn0(259)) {
                if (-32768 > 808) {
                    var v20 float64 = (0.1 + 1e10)
                    var v21 uint32 = 42
                    var v22 int8 = 127
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    switch fi17 {
                    case 1:
                        v23 := GenS1{F0: v18, F1: 42}
                        _ = v23
                    case 0:
                        for fi24 := 0; fi24 < 5; fi24++ {
                            fmt.Println(p0, p0, fi17)
                            fmt.Println(fi24, fi17)
                            _ = fi24
                        }
                    default:
                        for fi25 := 0; fi25 < 14; fi25++ {
                            switch fi17 {
                            case 2147483647:
                                if p0 {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                }
                            case -435:
                                var v26 int8 = 1
                                _ = v26
                            case 506:
                                var v27 string = "hello"
                                _ = v27
                            }
                            var v28 int32 = 1
                            _ = fi25
                            _ = v28
                        }
                    }
                    var v29 uint = (v18 & (v18 - (v18 - (v18 & v18))))
                    switch (fi17 / ((42) | 1)) {
                    case 1:
                        fmt.Println(v18)
                    case 907:
                        v30 := GenS2{F0: 42, F1: 612}
                        _ = v30
                    case 127:
                        v19 = append(v19, (fi17 & fi17))
                    }
                    _ = v29
                }
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 int8 = 127
                    var v33 string = "hello"
                    p0 = false
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
    case -26:
        if (1 == 32767) {
            if ((42 != 4294967295) || p0) {
                v34 := GenS3{F0: true, F1: -446, F2: p0}
                if (-771 == 9223372036854775807) {
                    for fi35 := 0; fi35 < 19; fi35++ {
                        var v36 uint64 = 807
                        _ = fi35
                        _ = v36
                    }
                    p0 = p0
                    for fi37 := 0; fi37 < 2; fi37++ {
                        var v38 int = (fi37 << ((42) & 63))
                        if ("foo bar" == "hello") {
                            v39 := []int{fi37, fi37, 251}
                            _ = v39
                        } else {
                            v40 := GenS1{F0: 2, F1: 128}
                            for fi41 := 0; fi41 < 19; fi41++ {
                                switch 418 {
                                case 0:
                                    v38 -= (fi41 + fi41)
                                case 35:
                                    var v42 uint = 42
                                    _ = v42
                                case 2147483647:
                                    switch "a" {
                                    case "foo bar":
                                        var v43 int = v38
                                        _ = v43
                                    }
                                default:
                                    v44 := []int{}
                                    _ = v44
                                }
                                p0 = (2 == 109)
                                _ = fi41
                            }
                            v45 := GenS4{F0: 128, F1: 1, F2: 1}
                            _ = v40
                            _ = v45
                        }
                        p0 = (322 < 2)
                        _ = fi37
                        _ = v38
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        var v46 int64 = 0
                        _ = v46
                    case "\n\t":
                        p0 = ("a" > ("foo bar" + ("" + (((("\n\t" + ((("a" + "foo bar") + (("foo bar" + "\n\t") + "foo bar")) + "\n\t")) + (("\n\t" + "") + "")) + (("a" + "hello") + "foo bar")) + "a"))))
                    default:
                        var v47 int = 32767
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 7; fi48++ {
                        switch (fi48 & 128) {
                        case 0:
                            for fi49 := 0; fi49 < 15; fi49++ {
                                v50 := GenS5{F0: p0, F1: "a", F2: 127}
                                switch ("foo bar" + "a") {
                                case "hello":
                                    var v51 float32 = -2.25
                                    _ = v51
                                default:
                                    var v52 int64 = 42
                                    _ = v52
                                }
                                v53 := GenS6{F0: "foo bar", F1: 2, F2: 32767}
                                _ = fi49
                                _ = v50
                                _ = v53
                            }
                        default:
                            var v54 uint16 = 777
                            _ = v54
                        }
                        p0 = false
                        switch "" {
                        case "\n\t":
                            switch "foo bar" {
                            case "":
                                v55 := []int{}
                                _ = v55
                            }
                        default:
                            v56 := []int{fi48}
                            _ = v56
                        }
                        _ = fi48
                    }
                }
                fmt.Println(p0)
                _ = v34
            } else {
                var v57 uint64 = 2
                _ = v57
            }
        }
    }
    var v58 uint64 = 18446744073709551615
    _ = p0
    _ = v15
    _ = v58
    return -742
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 69}
    var v3 int16 = 32767
    var v4 uint = 4294967295
    var v5 float32 = (-2.25 + (((3.14159 + 0.1) / 1.5) + (0.0 - ((float32(0.0) + float32(0.0)) / 3.14159))))
    var v6 uint16 = 571
    if true {
        v4 = v4
    } else {
        if (553 != 2147483647) {
            v5 = float32((1.5 - 0.1))
        } else {
            _ = len(v2)
            if true {
                v59 := GenS4{F0: 61, F1: 65535, F2: v4}
                if true {
                    v6 = 42
                    for fi60 := 0; fi60 < 12; fi60++ {
                        var v61 uint = (^(v4))
                        for fi62 := 0; fi62 < 3; fi62++ {
                            if true {
                                v63 := GenS7{F0: -731, F1: 65535}
                                switch "\n\t" {
                                case "hello":
                                    switch 32767 {
                                    case 9223372036854775807:
                                        if false {
                                            if (!true) {
                                                var v64 uint32 = 242
                                                _ = v64
                                            } else {
                                                var v65 int8 = -42
                                                _ = v65
                                            }
                                            if true {
                                                v6 = v6
                                            }
                                            if ("a" < "") {
                                                fmt.Println(v61, fi62, fi60)
                                                fmt.Println(v5)
                                            }
                                        }
                                    case 32767:
                                        var v66 bool = (genFn0(1) != 99)
                                        _ = v66
                                    case 1:
                                        var v67 string = "a"
                                        _ = v67
                                    }
                                }
                                v61 -= v4
                                _ = v63
                            }
                            switch ("" + "hello") {
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = 0
                                }
                            case "hello":
                                v61 = (^(v61))
                            default:
                                v3 -= (^(v3))
                                if false {
                                    var v68 bool = false
                                    var v69 uint16 = (v6 / ((v6) | 1))
                                    v2 = append(v2, 2)
                                    _ = v68
                                    _ = v69
                                }
                            }
                            _ = fi62
                        }
                        v70 := GenS8{F0: 33, F1: -559, F2: "hello"}
                        _ = fi60
                        _ = v61
                        _ = v70
                    }
                }
                if (true && (1e10 < -2.25)) {
                    v71 := GenS3{F0: true, F1: 0, F2: (0 != 43)}
                    _ = v71
                }
                _ = v59
            } else {
                for fi72 := 0; fi72 < 7; fi72++ {
                    for fi73 := 0; fi73 < 18; fi73++ {
                        if (18446744073709551615 >= 592) {
                            var v74 uint64 = 838
                            if true {
                                switch 890 {
                                case 9223372036854775807:
                                    switch (("hello" + "") + "foo bar") {
                                    case "a":
                                        _ = len(v2)
                                    }
                                case -1:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 704:
                                    v74 -= 18446744073709551615
                                default:
                                    fmt.Println(v3)
                                }
                                switch genFn1((0 < -9223372036854775808)) {
                                case 0:
                                    v75 := GenS9{F0: (fi72 - 9223372036854775807), F1: 255}
                                    _ = v75
                                case -1:
                                    if true {
                                        var v76 uint16 = uint16(v74)
                                        if (((-178 < -1) && (v3 > 42)) && false) {
                                            var v77 bool = false
                                            _ = v77
                                        } else {
                                            var v78 uint32 = uint32(v74)
                                            _ = v78
                                        }
                                        var v79 int16 = -1
                                        var v80 uint = (v4 & uint(v74))
                                        _ = v76
                                        _ = v79
                                        _ = v80
                                    } else {
                                        var v81 int8 = 42
                                        switch (("hello" + "a") + "a") {
                                        case "hello":
                                            fmt.Println(v81, v3, fi72)
                                        }
                                        _ = v81
                                    }
                                case 127:
                                    v6 = (v6 / (((v6 % ((v6) | 1))) | 1))
                                default:
                                    v82 := v1[(^(fi72))]
                                    _ = v82
                                }
                                var v83 uint8 = 127
                                _ = v83
                            } else {
                                var v84 float64 = float64((((0.1 / 1.5) - 0.0) - (v5 + 1.5)))
                                _ = v84
                            }
                            _ = v74
                        }
                        var v85 bool = false
                        _ = fi73
                        _ = v85
                    }
                    v3 -= -32768
                    _ = fi72
                }
            }
        }
    }
    var v86 uint32 = 876
    if (1.5 < (((1.5 - v5) * 0.1) + (v5 * 1.5))) {
        if (v4 >= 42) {
            fmt.Println(v86, v6)
            switch "\n\t" {
            case "\n\t":
                if (0.0 <= 1.5) {
                    for fi87 := 0; fi87 < 17; fi87++ {
                        v3 = (v3 - v3)
                        _ = len(v1)
                        v88 := GenS6{F0: ("" + (("foo bar" + "hello") + "a")), F1: 707, F2: 2}
                        _ = fi87
                        _ = v88
                    }
                    v3 = 754
                    v4 = (v4 + 4294967295)
                }
            default:
                fmt.Println(v86, v3)
            }
        }
        var v89 int8 = 42
        var v90 string = "foo bar"
        _ = v89
        _ = v90
    } else {
        var v91 bool = true
        if true {
            for fi92 := 0; fi92 < 16; fi92++ {
                for fi93 := 0; fi93 < 17; fi93++ {
                    var v94 float32 = 1.5
                    var v95 uint16 = (v6 / ((v6) | 1))
                    _ = fi93
                    _ = v94
                    _ = v95
                }
                _ = fi92
            }
            v96 := GenS2{F0: 0, F1: (v6 | v6)}
            _ = v96
        } else {
            var v97 uint32 = v86
            fmt.Println(v97, v91, v91)
            fmt.Println(v86)
            _ = v97
        }
        v86 = 42
        _ = v91
    }
    for fi98 := 0; fi98 < 3; fi98++ {
        fmt.Println(v3, v86, v6)
        for fi99 := 0; fi99 < 8; fi99++ {
            var v100 uint64 = 18446744073709551615
            _ = fi99
            _ = v100
        }
        _ = fi98
    }
    switch 1 {
    case 42:
        fmt.Println(v5, v86)
    default:
        var v101 int = 4294967295
        _ = v101
    }
    switch "\n\t" {
    case "a":
        for fi102 := 0; fi102 < 13; fi102++ {
            var v103 bool = ((105 >= 214) && ((128 >= 215) && (true || true)))
            _ = fi102
            _ = v103
        }
    case "":
        var v104 int32 = 42
        _ = v104
    case "hello":
        v5 -= (1e10 * (float32(1e10) - (v5 * 0.0)))
    }
    v105 := GenS8{F0: 127, F1: 675, F2: "\n\t"}
    var v106 uint8 = 255
    var v107 bool = false
    var v108 uint8 = v106
    v6 -= v6
    v108 = uint8(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v86
}
