package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 string; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 string }
func genFn0(p0 uint64, p1 uint32, p2 uint8) string {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint16 = 65535
        switch fi4 {
        case 740:
            if (!((("" + "foo bar") + "") == "\n\t")) {
                p0 = (p0 | uint64(v5))
            }
        case 1:
            switch "hello" {
            case "\n\t":
                var v6 uint32 = uint32(p0)
                _ = v6
            case "hello":
                v7 := []int{fi4}
                _ = v7
            default:
                var v8 float32 = float32(-2.25)
                _ = v8
            }
        }
        _ = fi4
        _ = v5
    }
    v9 := GenS1{F0: 18446744073709551615}
    p0 = (^(p0))
    p1 += (p1 & (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "foo bar"
}
func genFn1() uint8 {
    var v10 uint8 = 127
    v11 := []int{127, 42, -810}
    v12 := GenS1{F0: 2}
    _ = len(v11)
    _ = v10
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 int64 = 9223372036854775807
    var v3 bool = (-1 == 0)
    if (!false) {
        v13 := GenS2{F0: 255, F1: -860, F2: -973}
        switch -1 {
        case 32767:
            for fi14 := 0; fi14 < 11; fi14++ {
                switch "\n\t" {
                case "hello":
                    if v3 {
                        v2 -= v2
                        var v15 float32 = (3.14159 + ((0.1 - 0.1) / 3.14159))
                        _ = v15
                    } else {
                        v16 := GenS3{F0: 4294967295}
                        fmt.Println(v3, v1)
                        _ = v16
                    }
                }
                if (false || true) {
                    v17 := GenS4{F0: v2, F1: ("hello" + "a"), F2: (-(v2))}
                    v18 := GenS3{F0: int(v1)}
                    _ = v17
                    _ = v18
                }
                if true {
                    switch "\n\t" {
                    case "":
                        var v19 float32 = float32(3.14159)
                        _ = v19
                    case "\n\t":
                        v20 := GenS5{F0: -32768}
                        _ = v20
                    }
                    var v21 uint8 = genFn1()
                    _ = v21
                }
                _ = fi14
            }
        case 436:
            for fi22 := 0; fi22 < 16; fi22++ {
                switch "hello" {
                case "hello":
                    v23 := []int{(fi22 % (((^(fi22))) | 1)), (-(fi22))}
                    _ = v23
                case "":
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v1 = v1
                        var v25 int32 = v1
                        _ = fi24
                        _ = v25
                    }
                default:
                    switch "hello" {
                    case "a":
                        v2 += (v2 << ((42) & 63))
                    case "\n\t":
                        var v26 uint64 = 42
                        _ = v26
                    }
                }
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 int16 = 0
                    var v29 uint8 = genFn1()
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch 243 {
                case -1:
                    var v30 int = fi22
                    _ = v30
                case 32768:
                    v31 := GenS1{F0: 65535}
                    _ = v31
                default:
                    if (-954 <= 0) {
                        switch fi22 {
                        case -2147483648:
                            v1 -= v1
                        default:
                            v2 -= (-(v2))
                        }
                        v1 = (v1 % ((2147483647) | 1))
                        v3 = (123 <= 120)
                    } else {
                        v2 += -1
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v33 := GenS1{F0: 0}
                            switch ("hello" + "foo bar") {
                            case "hello":
                                v34 := []int{}
                                _ = v34
                            case "a":
                                if (v3 || true) {
                                    v35 := GenS6{F0: 1, F1: 0}
                                    var v36 float64 = float64(1e10)
                                    fmt.Println(fi22)
                                    _ = v35
                                    _ = v36
                                }
                            case "":
                                v37 := []int{fi22, fi22, (fi22 - fi32), fi32}
                                _ = v37
                            default:
                                var v38 int64 = v2
                                _ = v38
                            }
                            switch 2147483647 {
                            case -778:
                                if (!v3) {
                                    if false {
                                        var v39 uint16 = 2
                                        _ = v39
                                    } else {
                                        v3 = true
                                        v40 := []int{42}
                                        var v41 float64 = (float64(3.14159) - 0.0)
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 float32 = ((float32(0.1) + (3.14159 / (1.5 / (0.0 + 1e10)))) + 3.14159)
                                    v43 := GenS4{F0: v2, F1: "foo bar", F2: -1}
                                    if v3 {
                                        v2 = v2
                                        var v44 float64 = 0.1
                                        if v3 {
                                            if v3 {
                                                var v45 int = fi22
                                                if (1 < 0) {
                                                    v1 += -649
                                                } else {
                                                    v1 += (v1 << ((1) & 63))
                                                }
                                                _ = v45
                                            } else {
                                                v46 := GenS7{F0: (v2 & v2), F1: -32768}
                                                v47 := GenS2{F0: 206, F1: v1, F2: (fi22 + (^(fi22)))}
                                                _ = v46
                                                _ = v47
                                            }
                                            if (false || v3) {
                                                if (!v3) {
                                                    var v48 uint8 = 128
                                                    if (4294967295 == 869) {
                                                        v49 := []int{fi32, 0, fi32}
                                                        _ = v49
                                                    }
                                                    var v50 bool = v3
                                                    _ = v48
                                                    _ = v50
                                                } else {
                                                    if (569 <= 1000) {
                                                        switch (fi32 & (fi32 / (((-(fi22))) | 1))) {
                                                        case -9223372036854775808:
                                                            switch fi32 {
                                                            case -2147483648:
                                                                v3 = v3
                                                            case 2147483647:
                                                                v1 += -1
                                                            case 4294967295:
                                                                fmt.Println(v3, v2)
                                                            default:
                                                                v51 := []int{}
                                                                _ = v51
                                                            }
                                                        case 0:
                                                            var v52 string = ("a" + "a")
                                                            _ = v52
                                                        }
                                                        v53 := []int{-347, (^(fi22)), fi22, -2147483648}
                                                        v2 += -9223372036854775808
                                                        _ = v53
                                                    }
                                                    v54 := []int{(fi22 * fi22), 0, fi22}
                                                    switch ("a" + ("a" + "\n\t")) {
                                                    case "a":
                                                        v44 -= (0.1 - v44)
                                                    case "":
                                                        var v55 uint16 = 601
                                                        _ = v55
                                                    case "\n\t":
                                                        var v56 bool = (("\n\t" + "a") == "\n\t")
                                                        _ = v56
                                                    default:
                                                        fmt.Println(fi22, v44)
                                                    }
                                                    _ = v54
                                                }
                                                var v57 int64 = 0
                                                var v58 uint16 = 1
                                                _ = v57
                                                _ = v58
                                            }
                                        } else {
                                            v59 := []int{(fi32 & (fi22 | (-(fi22)))), (fi32 * 9223372036854775807)}
                                            var v60 int8 = 57
                                            v3 = v3
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v44
                                    } else {
                                        switch 1 {
                                        case -665:
                                            v1 -= (v1 << ((578) & 63))
                                        case 63:
                                            if true {
                                                v1 -= v1
                                            }
                                        default:
                                            if v3 {
                                                v61 := []int{32767, fi32}
                                                _ = v61
                                            }
                                        }
                                        var v62 int16 = 948
                                        v3 = (!v3)
                                        _ = v62
                                    }
                                    if (1 > 1) {
                                        v63 := []int{fi32, fi22, (fi22 ^ (fi32 / ((fi22) | 1))), 2147483647}
                                        var v64 int = (fi32 ^ -1)
                                        var v65 uint = 65535
                                        var v66 uint8 = 255
                                        v67 := GenS8{F0: ""}
                                        var v68 uint = (v65 - (v65 & 42))
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = v42
                                    _ = v43
                                } else {
                                    v3 = true
                                    v1 = (^(v1))
                                }
                            case 4294967295:
                                v69 := []int{9223372036854775807}
                                _ = v69
                            case 127:
                                v1 += v1
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                }
                _ = fi22
            }
        case -1:
            var v70 int = -435
            _ = v70
        }
        _ = v13
    } else {
        v71 := []int{4294967295, 32767, -378, 296}
        for fi72 := 0; fi72 < 19; fi72++ {
            if ("hello" != "\n\t") {
                if len(v71) >= 2 {
                    v71 = v71[1:len(v71)]
                }
                var v73 int8 = -60
                if true {
                    var v74 int = (fi72 + fi72)
                    _ = v74
                } else {
                    v75 := GenS6{F0: (v73 * (v73 | v73)), F1: v73}
                    v1 = v1
                    switch 9223372036854775807 {
                    case -1:
                        for fi76 := 0; fi76 < 10; fi76++ {
                            v2 = -220
                            v3 = v3
                            _ = fi76
                        }
                    case 128:
                        fmt.Println(v3, v3)
                    case -880:
                        fmt.Println(v2, fi72, v3)
                    default:
                        var v77 int = (fi72 ^ (fi72 / ((fi72) | 1)))
                        _ = v77
                    }
                    _ = v75
                }
                _ = v73
            }
            if (0.1 > (((((-2.25 + 3.14159) * 1.5) * float32(((0.1 - -2.25) * (((-2.25 / (0.1 - 1.5)) - 0.1) * 0.1)))) * (-2.25 + (3.14159 - 1.5))) / float32(((0.1 + 3.14159) + 0.1)))) {
                var v78 uint32 = 2
                v79 := GenS8{F0: "hello"}
                _ = v78
                _ = v79
            }
            for fi80 := 0; fi80 < 2; fi80++ {
                var v81 int64 = (v2 * int64(v1))
                _ = fi80
                _ = v81
            }
            _ = fi72
        }
        _ = v71
    }
    switch 401 {
    case 1:
        var v82 int64 = v2
        _ = v82
    case 42:
        var v83 float64 = (1e10 * (1e10 / 1.0))
        _ = v83
    case -9223372036854775808:
        for fi84 := 0; fi84 < 15; fi84++ {
            v3 = (v3 && false)
            v3 = false
            _ = fi84
        }
    default:
        v85 := []int{9223372036854775807}
        _ = v85
    }
    var v86 int = 42
    for fi87 := 0; fi87 < 5; fi87++ {
        fmt.Println(fi87)
        _ = fi87
    }
    v88 := GenS5{F0: 0}
    var v89 string = "a"
    fmt.Println(v1)
    fmt.Println(v89, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v86:", v86)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v86
    _ = v88
    _ = v89
}
