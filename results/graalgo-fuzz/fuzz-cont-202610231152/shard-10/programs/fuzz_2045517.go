package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int64; F2 int64 }
type GenS6 struct { F0 uint; F1 string; F2 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64) string {
    v7 := GenS1{F0: uint16(p0), F1: 128}
    v8 := []int{128, 42}
    var v9 uint64 = p0
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ("a" + "a")
}
func genFn1() int {
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := []int{(fi12 | fi12), (fi11 & fi11), (fi11 ^ (fi12 + (fi12 + (fi11 >> ((188) & 63))))), fi11}
            if (((0.1 == 3.14159) && (18446744073709551615 >= 2)) || (989 != 65535)) {
                var v14 int = 127
                v14 -= 383
                v13 = append(v13, (fi11 % ((fi12) | 1)))
                _ = v14
            } else {
                var v15 int32 = 2147483647
                _ = v15
            }
            v16 := GenS2{F0: -1}
            _ = fi12
            _ = v13
            _ = v16
        }
        _ = fi11
    }
    var v17 uint = 255
    if false {
        if true {
            switch -726 {
            case 821:
                var v18 uint32 = 788
                _ = v18
            case -9223372036854775808:
                for fi19 := 0; fi19 < 11; fi19++ {
                    v17 = (v17 * v17)
                    var v20 string = "hello"
                    _ = fi19
                    _ = v20
                }
            case 42:
                v21 := GenS1{F0: 104, F1: 334}
                _ = v21
            default:
                if false {
                    v22 := GenS3{F0: -368, F1: (v17 ^ (v17 * 0))}
                    for fi23 := 0; fi23 < 15; fi23++ {
                        if (115 >= 128) {
                            v17 = 0
                            if true {
                                v24 := GenS2{F0: -483}
                                if true {
                                    if false {
                                        for fi25 := 0; fi25 < 11; fi25++ {
                                            v26 := []int{(fi23 + fi25), fi25, 640, fi23}
                                            _ = fi25
                                            _ = v26
                                        }
                                        v17 += (v17 ^ (v17 & 18446744073709551615))
                                        if (fi23 >= (-(fi23))) {
                                            v27 := []int{(fi23 >> (((v17 & 65535)) & 63)), fi23, fi23}
                                            v17 += v17
                                            _ = v27
                                        } else {
                                            if (("foo bar" > (("a" + "a") + "a")) || (false || false)) {
                                                fmt.Println(v17, v17)
                                                if false {
                                                    var v28 uint = v17
                                                    if false {
                                                        v29 := GenS1{F0: 2, F1: (fi23 | fi23)}
                                                        _ = v29
                                                    } else {
                                                        if true {
                                                            v30 := GenS4{F0: 176}
                                                            if false {
                                                                var v31 int = (-(fi23))
                                                                switch -914 {
                                                                case 32767:
                                                                    v31 = 4294967295
                                                                case -1:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        var v32 float64 = 1.5
                                                                        _ = v32
                                                                    default:
                                                                        fmt.Println(v31)
                                                                    }
                                                                case 9223372036854775807:
                                                                    v33 := []int{fi23, (fi23 >> (((^(v28))) & 63)), 42}
                                                                    _ = v33
                                                                default:
                                                                    fmt.Println(v31, v28, v28)
                                                                }
                                                                if (((-2.25 + (3.14159 + 1.5)) / 1.0) == -2.25) {
                                                                    var v34 float64 = (-2.25 * float64(0.1))
                                                                    var v35 int32 = -438
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                _ = v31
                                                            } else {
                                                                for fi36 := 0; fi36 < 2; fi36++ {
                                                                    v17 -= v17
                                                                    if (false && true) {
                                                                        v37 := GenS1{F0: 42, F1: -1}
                                                                        _ = v37
                                                                    } else {
                                                                        v38 := GenS1{F0: 780, F1: fi23}
                                                                        if (915 != 1) {
                                                                            if ((v17 | (v28 & 42)) != v17) {
                                                                                var v39 int64 = 677
                                                                                var v40 uint32 = 0
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            v28 += (v28 % ((18446744073709551615) | 1))
                                                                            v41 := GenS5{F0: -469, F1: -578, F2: -135}
                                                                            _ = v41
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if false {
                                                                                v42 := GenS6{F0: v17, F1: ("foo bar" + "hello"), F2: 42}
                                                                                fmt.Println(fi36, v28, fi23)
                                                                                var v43 uint64 = 544
                                                                                _ = v42
                                                                                _ = v43
                                                                            }
                                                                        default:
                                                                            v17 += v28
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                    _ = fi36
                                                                }
                                                            }
                                                            _ = v30
                                                        }
                                                        var v44 int16 = 618
                                                        v28 += v28
                                                        _ = v44
                                                    }
                                                    var v45 int16 = -626
                                                    _ = v28
                                                    _ = v45
                                                }
                                                v17 += (^(v17))
                                            } else {
                                                for fi46 := 0; fi46 < 15; fi46++ {
                                                    v47 := []int{(fi23 | (fi23 & (fi46 ^ fi23))), fi46, 32768, (fi46 / ((fi23) | 1))}
                                                    _ = fi46
                                                    _ = v47
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if true {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v48 float64 = (-2.25 - 0.1)
                                                            _ = v48
                                                        }
                                                        var v49 int = fi23
                                                        var v50 uint16 = 1
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS3{F0: 510, F1: 4294967295}
                                                        v52 := GenS7{F0: 147}
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                case "":
                                                    if (2 < 42) {
                                                        v17 = 2
                                                        v17 = v17
                                                    }
                                                }
                                            }
                                            var v53 bool = false
                                            var v54 int32 = 2147483647
                                            _ = v53
                                            _ = v54
                                        }
                                    } else {
                                        if true {
                                            for fi55 := 0; fi55 < 13; fi55++ {
                                                v56 := []int{32767, fi55, 4294967295, 412}
                                                _ = fi55
                                                _ = v56
                                            }
                                            if false {
                                                v57 := GenS8{F0: -913}
                                                var v58 int16 = 627
                                                v59 := []int{fi23, (^(fi23)), -29}
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            } else {
                                                var v60 string = "a"
                                                var v61 float64 = -2.25
                                                if false {
                                                    v61 += 0.1
                                                    v61 += (v61 + v61)
                                                    fmt.Println(v17, fi23)
                                                }
                                                _ = v60
                                                _ = v61
                                            }
                                            v17 = (^(v17))
                                        } else {
                                            for fi62 := 0; fi62 < 6; fi62++ {
                                                if true {
                                                    v17 -= (^(v17))
                                                    fmt.Println(v17, fi62)
                                                    switch "hello" {
                                                    case "a":
                                                        v63 := []int{fi23, (fi23 ^ (fi62 * -1))}
                                                        _ = v63
                                                    }
                                                    v17 = (v17 & v17)
                                                    var v64 int = (^(fi23))
                                                    _ = v64
                                                }
                                                v65 := GenS9{F0: -32768, F1: int(v17)}
                                                _ = fi62
                                                _ = v65
                                            }
                                            v17 = (v17 + v17)
                                            var v66 uint64 = 42
                                            _ = v66
                                        }
                                        fmt.Println(v17, fi23)
                                    }
                                    v67 := []int{-935, fi23, (fi23 + (fi23 >> ((534) & 63)))}
                                    if true {
                                        v17 = (v17 >> (((v17 | 971)) & 63))
                                        v67 = append(v67, fi23)
                                    } else {
                                        v17 = v17
                                        v68 := GenS10{F0: ("\n\t" + ("" + "hello"))}
                                        _ = v68
                                    }
                                    _ = v67
                                }
                                _ = v24
                            }
                            for fi69 := 0; fi69 < 4; fi69++ {
                                v70 := []int{(fi69 + (fi23 << ((42) & 63))), fi69}
                                if len(v70) > 0 {
                                    v70[0] = (fi23 % ((fi69) | 1))
                                }
                                _ = fi69
                                _ = v70
                            }
                        }
                        if (!true) {
                            if false {
                                fmt.Println(v17)
                            }
                            fmt.Println(v17, v17)
                        } else {
                            v71 := GenS8{F0: -1}
                            _ = v71
                        }
                        if (9223372036854775807 != 42) {
                            for fi72 := 0; fi72 < 12; fi72++ {
                                var v73 int = 32767
                                var v74 int64 = 1
                                if false {
                                    v75 := []int{32767, -212}
                                    if (false || true) {
                                        if (!(200 <= 42)) {
                                            var v76 int = v73
                                            _ = v76
                                        } else {
                                            v17 += v17
                                            v73 -= fi72
                                            v73 += (v73 & v73)
                                        }
                                        var v77 int64 = (v74 + -1)
                                        switch "" {
                                        case "foo bar":
                                            v17 = v17
                                        }
                                        fmt.Println(v74, v74)
                                        v74 -= (v74 * 0)
                                        switch v73 {
                                        case 32768:
                                            switch 2147483647 {
                                            case 32768:
                                                v74 = v74
                                            case 812:
                                                v77 = v77
                                            default:
                                                if false {
                                                    v75 = append(v75, 0)
                                                    _ = len(v75)
                                                } else {
                                                    var v78 string = "hello"
                                                    _ = v78
                                                }
                                            }
                                        default:
                                            v79 := GenS2{F0: -292}
                                            _ = v79
                                        }
                                        v80 := GenS2{F0: -674}
                                        v17 = 65535
                                        _ = v77
                                        _ = v80
                                    }
                                    v74 += (-(v74))
                                    _ = v75
                                } else {
                                    switch (fi23 % ((9223372036854775807) | 1)) {
                                    case 1:
                                        v81 := GenS6{F0: v17, F1: "hello", F2: 127}
                                        _ = v81
                                    case 42:
                                        v82 := GenS9{F0: 0, F1: (fi23 + fi72)}
                                        _ = v82
                                    }
                                    v83 := []int{fi72, fi23, fi72, 9223372036854775807}
                                    _ = v83
                                }
                                v74 += (v74 | (v74 << (((v17 + v17)) & 63)))
                                v73 += v73
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            var v84 string = ""
                            _ = v84
                        } else {
                            v85 := GenS4{F0: -1}
                            if false {
                                fmt.Println(v17, fi23)
                            }
                            v86 := GenS10{F0: ("" + (((("foo bar" + "hello") + "\n\t") + "\n\t") + ""))}
                            _ = v85
                            _ = v86
                        }
                        _ = fi23
                    }
                    v87 := []int{}
                    _ = v22
                    _ = v87
                } else {
                    v88 := GenS11{F0: -597}
                    v89 := GenS7{F0: 4294967295}
                    _ = v88
                    _ = v89
                }
            }
            for fi90 := 0; fi90 < 7; fi90++ {
                for fi91 := 0; fi91 < 16; fi91++ {
                    v92 := []int{(^(fi90)), (fi91 >> ((534) & 63))}
                    v92 = append(v92, fi91)
                    _ = fi91
                    _ = v92
                }
                switch fi90 {
                case -1:
                    v93 := GenS12{F0: 0}
                    _ = v93
                case 128:
                    v17 += 1
                case 32768:
                    v17 -= (^(v17))
                }
                _ = fi90
            }
        }
        fmt.Println(v17)
        v94 := GenS9{F0: 42, F1: 127}
        _ = v94
    }
    _ = v17
    return 1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 42
    v3 := []uint32{v2, v2}
    v4 := make(map[uint32]uint32)
    var v5 uint = 508
    var v6 bool = false
    v1 -= (v1 | 1)
    switch 572 {
    case 32768:
        var v95 int64 = 0
        _ = v95
    case 2147483647:
        v6 = v6
    case -1:
        v5 = (v5 / ((0) | 1))
    }
    if (true && v6) {
        for fi96 := 0; fi96 < 16; fi96++ {
            v2 = v2
            var v97 float32 = 0.1
            v98 := GenS12{F0: uint16(v5)}
            _ = fi96
            _ = v97
            _ = v98
        }
    }
    v1 = v1
    for fi99 := 0; fi99 < 2; fi99++ {
        var v100 uint = (v5 - 65535)
        _ = fi99
        _ = v100
    }
    if (!(1e10 == float64(((-2.25 / float32(3.14159)) / 1e10)))) {
        v6 = v6
        var v101 uint8 = 1
        for fi102 := 0; fi102 < 20; fi102++ {
            var v103 float32 = 3.14159
            v4[65535] = 396
            for fi104 := 0; fi104 < 15; fi104++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v105 int = (^(fi102))
                switch "\n\t" {
                case "hello":
                    var v106 bool = true
                    _ = v106
                case "foo bar":
                    var v107 uint8 = v101
                    _ = v107
                case "a":
                    fmt.Println(v1, fi102)
                default:
                    var v108 float32 = v103
                    _ = v108
                }
                _ = fi104
                _ = v105
            }
            _ = fi102
            _ = v103
        }
        _ = v101
    }
    fmt.Println(v2, v6, v5)
    switch "" {
    case "":
        v5 += 0
    case "foo bar":
        if false {
            var v109 bool = v6
            _ = v109
        }
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
