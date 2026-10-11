package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 string; F1 int16; F2 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 int8; F2 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 int; F1 uint16; F2 int64 }
type GenS14 struct { F0 uint; F1 int16 }
func genFn0(p0 uint8) uint32 {
    v7 := []int{42}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int8 = 57
        for fi10 := 0; fi10 < 1; fi10++ {
            if false {
                var v11 uint64 = 18446744073709551615
                v12 := GenS1{F0: v11}
                _ = v11
                _ = v12
            } else {
                fmt.Println(p0, fi10)
            }
            _ = fi10
        }
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    switch "\n\t" {
    case "":
        var v13 uint = 0
        _ = v13
    case "a":
        var v14 uint32 = uint32(p0)
        _ = v14
    default:
        p0 = (^(p0))
    }
    v15 := GenS1{F0: 90}
    if false {
        switch ((("" + "hello") + "\n\t") + "a") {
        case "\n\t":
            for fi16 := 0; fi16 < 9; fi16++ {
                _ = len(v7)
                var v17 string = "\n\t"
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS2{F0: 18446744073709551615, F1: 487}
        _ = v18
    }
    _ = p0
    _ = v15
    _ = v7
    return 42
}

func main() {
    var v1 float32 = (3.14159 + 0.1)
    var v2 int16 = 679
    var v3 float32 = (v1 / v1)
    v4 := make(map[uint32]uint32)
    var v5 uint = 65535
    var v6 uint32 = 2
    v19 := GenS3{F0: 128, F1: 311}
    if false {
        if (((!true) && true) || (536 > 1)) {
            switch "hello" {
            case "":
                if true {
                    if true {
                        fmt.Println(v2, v1, v1)
                        v1 -= v1
                    } else {
                        if (true || (78 >= 9223372036854775807)) {
                            v20 := GenS2{F0: (^(v5)), F1: -233}
                            v21 := []int{127}
                            _ = len(v4)
                            _ = v20
                            _ = v21
                        }
                        switch ("a" + (("foo bar" + ("" + ((("\n\t" + ("\n\t" + "foo bar")) + "foo bar") + ((("hello" + "") + "hello") + ("foo bar" + ("foo bar" + "a")))))) + (("foo bar" + "a") + ("foo bar" + "a")))) {
                        case "hello":
                            switch "" {
                            case "a":
                                v2 = (v2 * -32768)
                            case "\n\t":
                                v2 += (v2 + v2)
                            case "foo bar":
                                if ((55 > 255) || false) {
                                    var v22 float32 = float32(3.14159)
                                    _ = v22
                                }
                            }
                        }
                        var v23 int64 = 42
                        _ = v23
                    }
                    if (0 >= 493) {
                        v24 := GenS4{F0: "", F1: (^(v2)), F2: -357}
                        _ = v24
                    }
                    v5 += v5
                } else {
                    if false {
                        v6 = 2
                        v25 := []int{42, 97}
                        for fi26 := 0; fi26 < 6; fi26++ {
                            if false {
                                var v27 uint = (^(v5))
                                _ = v27
                            } else {
                                v28 := GenS1{F0: uint64(v5)}
                                _ = v28
                            }
                            v29 := GenS3{F0: -826, F1: -2147483648}
                            _ = fi26
                            _ = v29
                        }
                        _ = v25
                    } else {
                        delete(v4, (v6 % (((^(v6))) | 1)))
                        v30 := []int{128, 0, 21, 919}
                        v1 += (1e10 + v3)
                        _ = v30
                    }
                }
            case "a":
                var v31 int16 = v2
                _ = v31
            case "foo bar":
                var v32 int = -2147483648
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                switch "" {
                case "foo bar":
                    if (v5 <= v5) {
                        v3 = 1.5
                        var v34 float32 = -2.25
                        _ = v34
                    }
                }
                if (42 == -1) {
                    v35 := v4[genFn0(147)]
                    var v36 uint8 = 2
                    v5 += v5
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        v37 := []int{(fi33 + -291), fi33}
                        for fi38 := 0; fi38 < 12; fi38++ {
                            _ = len(v4)
                            if len(v37) > 0 {
                                v37[0] = -9223372036854775808
                            }
                            v1 = 1e10
                            _ = fi38
                        }
                        _ = v37
                    }
                    v6 -= 2
                }
                v5 += v5
                _ = fi33
            }
            for fi39 := 0; fi39 < 16; fi39++ {
                v3 -= 1.5
                switch fi39 {
                case -2147483648:
                    if (1e10 >= (0.1 + 3.14159)) {
                        for fi40 := 0; fi40 < 5; fi40++ {
                            var v41 int64 = -1
                            var v42 uint16 = 0
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        if (true || false) {
                            if (("hello" + ("" + "a")) > "\n\t") {
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    var v44 int32 = -2147483648
                                    v5 = v5
                                    v45 := GenS5{F0: uint16(v5), F1: fi43}
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                                v46 := GenS6{F0: 4294967295}
                                for fi47 := 0; fi47 < 18; fi47++ {
                                    v48 := []int{fi39, (fi47 ^ fi47), (^(fi39)), -425}
                                    if false {
                                        switch fi47 {
                                        case 9223372036854775807:
                                            var v49 uint = 1
                                            _ = v49
                                        }
                                        v50 := GenS7{F0: 1}
                                        _ = v50
                                    } else {
                                        v1 = ((float32(((float64(3.14159) / 1.0) + (-2.25 / (((0.0 - (float64(v3) / -2.25)) + 0.0) * (1.5 / (float64(0.0) + 1.5)))))) - v3) - float32(((3.14159 / float64(3.14159)) + ((1.5 - 1.5) - -2.25))))
                                        var v51 int = -1
                                        fmt.Println(fi47, v6)
                                        v3 = (0.1 / ((3.14159 / (v3 - 0.1)) / v3))
                                        _ = v51
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                                _ = v46
                            }
                            for fi52 := 0; fi52 < 16; fi52++ {
                                var v53 uint8 = uint8(v5)
                                _ = fi52
                                _ = v53
                            }
                        } else {
                            for fi54 := 0; fi54 < 4; fi54++ {
                                v6 = 0
                                _ = fi54
                            }
                            if (332 <= 0) {
                                if false {
                                    v55 := GenS8{F0: (fi39 & -892)}
                                    for fi56 := 0; fi56 < 5; fi56++ {
                                        v1 -= v1
                                        var v57 uint = (v5 ^ v5)
                                        v58 := GenS3{F0: -322, F1: 26}
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                    if true {
                                        var v59 uint16 = 47
                                        v60 := []int{4294967295, (fi39 | (fi39 & 0))}
                                        var v61 int = (fi39 << ((18446744073709551615) & 63))
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                    _ = v55
                                } else {
                                    var v62 uint = 899
                                    switch fi39 {
                                    case 127:
                                        v5 += (v5 | v5)
                                    }
                                    if (fi39 != 32767) {
                                        v62 = v62
                                    } else {
                                        if true {
                                            var v63 uint64 = 4294967295
                                            _ = v63
                                        }
                                        var v64 uint16 = 502
                                        _ = v64
                                    }
                                    _ = v62
                                }
                            } else {
                                switch -9223372036854775808 {
                                case 918:
                                    for fi65 := 0; fi65 < 19; fi65++ {
                                        if (-873 != -9223372036854775808) {
                                            if (false && ((-229 != 0) && (-128 < -104))) {
                                                fmt.Println(v6)
                                            }
                                            var v66 uint = (v5 % ((v5) | 1))
                                            v67 := v4[(v6 * v6)]
                                            var v68 string = "a"
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                        _ = fi65
                                    }
                                case 276:
                                    var v69 float64 = (float64(((((3.14159 / ((1.5 + (1e10 + v1)) - v3)) * ((v3 * v3) - 0.1)) - (v1 / (-2.25 * (-2.25 - 1e10)))) / v3)) - 0.0)
                                    _ = v69
                                default:
                                    var v70 int64 = -719
                                    _ = v70
                                }
                            }
                        }
                        v5 = (v5 >> ((v5) & 63))
                    } else {
                        var v71 int32 = 2147483647
                        fmt.Println(v71, v2, v1)
                        _ = v71
                    }
                case 2147483647:
                    v72 := GenS9{F0: 0, F1: -128, F2: (fi39 / (((fi39 * fi39)) | 1))}
                    _ = v72
                default:
                    var v73 uint32 = 2
                    _ = v73
                }
                _ = fi39
            }
        }
        var v74 uint = v5
        _ = v74
    } else {
        v1 = float32(1.5)
    }
    for fi75 := 0; fi75 < 6; fi75++ {
        if (!(!(!(false && false)))) {
            if (42 < 1) {
                v1 -= float32(0.0)
            }
            switch "hello" {
            case "hello":
                v76 := GenS10{F0: false}
                _ = v76
            case "":
                v77 := []int{127, (fi75 + fi75)}
                _ = v77
            default:
                v78 := []int{fi75}
                _ = v78
            }
            v79 := GenS6{F0: (^(v6))}
            _ = v79
        } else {
            for fi80 := 0; fi80 < 13; fi80++ {
                v81 := GenS11{F0: 190}
                v82 := GenS1{F0: 2}
                v2 += v2
                _ = fi80
                _ = v81
                _ = v82
            }
            v83 := v4[v6]
            _ = v83
        }
        _ = fi75
    }
    if true {
        var v84 float64 = 0.1
        _ = v84
    } else {
        v6 = (v6 - (^(v6)))
    }
    var v85 float32 = (v3 + (((float32(0.0) + 0.0) / (v1 * ((v3 * 0.1) * v1))) + v1))
    var v86 bool = ((((float32(0.1) - v85) - 0.0) / v1) <= -2.25)
    for fi87 := 0; fi87 < 18; fi87++ {
        fmt.Println(v6, v86, v1)
        _ = fi87
    }
    var v88 uint64 = 18446744073709551615
    var v89 int = -45
    v90 := GenS12{F0: (v89 % ((v89) | 1)), F1: ""}
    v2 += v2
    v91 := v4[65535]
    switch ("a" + "a") {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        var v92 int16 = v2
        _ = v92
    case "hello":
        for fi93 := 0; fi93 < 1; fi93++ {
            v94 := v4[v91]
            for fi95 := 0; fi95 < 3; fi95++ {
                switch fi93 {
                case -1:
                    var v96 int16 = -675
                    var v97 string = "hello"
                    _ = v96
                    _ = v97
                }
                v98 := []int{(^(fi93))}
                v99 := GenS7{F0: (^(v88))}
                if v86 {
                    var v100 int8 = 1
                    v101 := GenS13{F0: (fi95 / (((fi93 << (((v5 | uint(v88))) & 63))) | 1)), F1: 65535, F2: int64(v6)}
                    _ = v100
                    _ = v101
                } else {
                    switch "hello" {
                    case "foo bar":
                        v102 := GenS6{F0: (v94 | v94)}
                        _ = v102
                    case "":
                        v2 -= (^(v2))
                    case "\n\t":
                        v103 := GenS14{F0: uint(v88), F1: v2}
                        _ = v103
                    default:
                        var v104 uint16 = 255
                        _ = v104
                    }
                    var v105 uint32 = v6
                    var v106 int32 = int32(v5)
                    _ = v105
                    _ = v106
                }
                _ = fi95
                _ = v98
                _ = v99
            }
            _ = fi93
            _ = v94
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
    _ = v86
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
