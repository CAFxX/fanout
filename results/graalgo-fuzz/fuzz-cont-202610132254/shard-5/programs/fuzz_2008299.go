package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int64; F2 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 int64; F1 int64; F2 uint16 }
type GenS11 struct { F0 uint64; F1 int16 }
type GenS12 struct { F0 uint32; F1 string; F2 int32 }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 int32; F1 int64; F2 int }
type GenS16 struct { F0 bool; F1 int32; F2 int8 }
type GenS17 struct { F0 uint32; F1 bool }
func genFn0() bool {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 float32 = (3.14159 + (1e10 * (3.14159 * (0.1 + 0.1))))
        var v7 int32 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi8, (fi9 % ((1) | 1)), fi8}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            if false {
                v12 := GenS1{F0: 4294967295}
                fmt.Println(fi11, fi8)
                _ = v12
            } else {
                switch "hello" {
                case "":
                    v13 := GenS2{F0: -14}
                    _ = v13
                case "a":
                    v14 := GenS2{F0: -577}
                    _ = v14
                default:
                    var v15 int16 = -32768
                    _ = v15
                }
                var v16 uint16 = 2
                v17 := []int{(fi11 >> ((255) & 63))}
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: -2147483648}
            _ = fi11
            _ = v18
        }
        var v19 uint32 = 2147483648
        _ = fi8
        _ = v19
    }
    return true
}
func genFn1() uint {
    var v20 float32 = 1.5
    v20 = v20
    for fi21 := 0; fi21 < 13; fi21++ {
        if true {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint64 = 2
                v20 -= float32(float64(0.0))
                if false {
                    if (false || (false && true)) {
                        if (0 <= 0) {
                            v24 := GenS3{F0: -2147483648}
                            v25 := GenS4{F0: 845, F1: 32767}
                            v26 := []int{(-(fi21)), (fi21 / ((fi22) | 1)), (fi21 | (fi21 % ((fi21) | 1)))}
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v27 := []int{}
                            v27 = append(v27, fi21)
                            if ((-128 <= -1) && (462 >= 32767)) {
                                v28 := GenS4{F0: -32768, F1: 33}
                                fmt.Println(fi22, v20)
                                _ = v28
                            } else {
                                var v29 int8 = -128
                                if true {
                                    var v30 string = ("\n\t" + "hello")
                                    var v31 float64 = float64(v20)
                                    var v32 uint = 855
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS5{F0: 32767}
                                    _ = v33
                                }
                                var v34 float32 = v20
                                v35 := GenS3{F0: 941}
                                if ((211 <= 255) && (true || (false || (true && ((128 >= 2) && true))))) {
                                    if true {
                                        v36 := GenS1{F0: (^(v23))}
                                        var v37 uint16 = 32767
                                        var v38 uint64 = 2
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 int8 = (v29 % (((v29 + (^(v29)))) | 1))
                                        _ = v39
                                    }
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    var v40 uint64 = v23
                                    switch fi22 {
                                    case -9223372036854775808:
                                        var v41 uint = 4294967295
                                        _ = v41
                                    case 298:
                                        fmt.Println(fi21)
                                    case -2147483648:
                                        var v42 int8 = v29
                                        _ = v42
                                    default:
                                        var v43 int64 = -659
                                        _ = v43
                                    }
                                    _ = v40
                                } else {
                                    var v44 float64 = float64((v20 / v34))
                                    if (0 != int64(v23)) {
                                        v45 := GenS6{F0: 42, F1: -90, F2: 206}
                                        _ = v45
                                    }
                                    var v46 int16 = int16(fi22)
                                    _ = v44
                                    _ = v46
                                }
                                _ = v29
                                _ = v34
                                _ = v35
                            }
                            _ = v27
                        }
                        v47 := GenS7{F0: -1}
                        var v48 int = 42
                        _ = v47
                        _ = v48
                    }
                }
                _ = fi22
                _ = v23
            }
            v20 += v20
        } else {
            v49 := GenS4{F0: -32768, F1: 865}
            v20 = 1.5
            v20 += v20
            _ = v49
        }
        fmt.Println(fi21, v20)
        _ = fi21
    }
    _ = v20
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := []string{((("hello" + "hello") + ("foo bar" + "foo bar")) + "\n\t"), ("foo bar" + "hello"), (("foo bar" + (("" + "hello") + "foo bar")) + "")}
    var v4 uint8 = 2
    v50 := GenS8{F0: 18, F1: 110, F2: 42}
    for fi51 := 0; fi51 < 10; fi51++ {
        var v52 float64 = (1.5 / (0.1 + float64(1e10)))
        for fi53 := 0; fi53 < 5; fi53++ {
            fmt.Println(fi51, fi53, v4)
            _ = fi53
        }
        _ = fi51
        _ = v52
    }
    if ((true || (!true)) || false) {
        for fi54 := 0; fi54 < 18; fi54++ {
            v55 := GenS4{F0: 1, F1: 2}
            _ = fi54
            _ = v55
        }
        var v56 int = -651
        _ = v56
    } else {
        fmt.Println(v4)
        for fi57 := 0; fi57 < 13; fi57++ {
            switch "foo bar" {
            case "hello":
                var v58 int16 = int16(v4)
                _ = v58
            case "\n\t":
                if true {
                    v59 := GenS9{F0: 0, F1: ""}
                    v4 = v4
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 1; fi60++ {
                        if true {
                            v4 = (v4 << ((246) & 63))
                            fmt.Println(fi57, fi60, fi57)
                        }
                        _ = fi60
                    }
                    var v61 int8 = 100
                    _ = v61
                }
            case "":
                for fi62 := 0; fi62 < 5; fi62++ {
                    v63 := GenS10{F0: 1, F1: 864, F2: 0}
                    _ = fi62
                    _ = v63
                }
            }
            _ = fi57
        }
        if (true && (!(!(447 > 4294967295)))) {
            v64 := GenS11{F0: 42, F1: 0}
            v65 := GenS12{F0: 42, F1: (((("foo bar" + "a") + "hello") + "\n\t") + ""), F2: -2147483648}
            switch int(v4) {
            case 0:
                fmt.Println(v4)
            case 32768:
                v4 -= (v4 ^ 0)
            case 32767:
                switch "\n\t" {
                case "a":
                    v4 -= v4
                case "foo bar":
                    for fi66 := 0; fi66 < 6; fi66++ {
                        if false {
                            v67 := GenS13{F0: 1, F1: 42}
                            _ = v67
                        }
                        v4 = v4
                        if true {
                            v68 := GenS14{F0: false, F1: 42}
                            if true {
                                v69 := GenS4{F0: 1, F1: 646}
                                v70 := v2[2]
                                _ = v69
                                _ = v70
                            }
                            _ = v68
                        } else {
                            v71 := v2[977]
                            for fi72 := 0; fi72 < 9; fi72++ {
                                v71 = (("" + (("a" + "") + ("a" + v71))) + v71)
                                _ = fi72
                            }
                            for fi73 := 0; fi73 < 7; fi73++ {
                                v74 := GenS1{F0: 0}
                                if ((!true) || (108 >= 4294967295)) {
                                    delete(v1, int(v4))
                                    var v75 float32 = 0.0
                                    v4 = 220
                                    _ = v75
                                } else {
                                    v3 = append(v3, "hello")
                                    fmt.Println(fi73, v71, v71)
                                    delete(v2, 4294967295)
                                }
                                _ = fi73
                                _ = v74
                            }
                            _ = v71
                        }
                        _ = fi66
                    }
                case "":
                    v76 := GenS3{F0: 1}
                    _ = v76
                default:
                    if (((false && (!false)) && false) || true) {
                        var v77 uint8 = 128
                        fmt.Println(v77, v77)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v77
                    } else {
                        switch (("a" + "hello") + "\n\t") {
                        case "\n\t":
                            v78 := GenS11{F0: 1, F1: 1}
                            _ = v78
                        case "":
                            v79 := GenS12{F0: 2147483648, F1: "\n\t", F2: 1}
                            _ = v79
                        }
                        if (false || ((true || false) && (!(2147483647 != 0)))) {
                            var v80 float64 = -2.25
                            delete(v2, 1)
                            _ = v80
                        } else {
                            var v81 int32 = -2147483648
                            _ = v81
                        }
                    }
                }
            }
            _ = v64
            _ = v65
        }
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v82 uint = 255
    for fi83 := 0; fi83 < 13; fi83++ {
        for fi84 := 0; fi84 < 3; fi84++ {
            switch "\n\t" {
            case "":
                v85 := GenS1{F0: 2}
                _ = v85
            }
            v86 := v1[fi83]
            if false {
                if true {
                    var v87 uint32 = 0
                    var v88 uint8 = (v4 << ((uint(v87)) & 63))
                    _ = v87
                    _ = v88
                } else {
                    var v89 bool = true
                    v4 = (v4 ^ (v4 & v4))
                    var v90 int8 = 0
                    _ = v89
                    _ = v90
                }
                v4 += 0
                v4 += (v4 * 2)
            }
            _ = fi84
            _ = v86
        }
        v91 := GenS15{F0: -1, F1: 0, F2: 127}
        for fi92 := 0; fi92 < 10; fi92++ {
            v93 := GenS16{F0: false, F1: 343, F2: -27}
            if (false || ((fi92 % (((fi92 + (fi83 / (((fi92 * (^(fi83)))) | 1)))) | 1)) != fi92)) {
                v94 := GenS17{F0: 0, F1: (32767 >= 2147483647)}
                _ = v94
            } else {
                var v95 int8 = 44
                if false {
                    if ((((fi92 ^ (fi83 - fi83)) >= -928) || ((^(v4)) == v4)) || false) {
                        switch (fi83 ^ (fi92 ^ (fi92 << ((65535) & 63)))) {
                        case 32767:
                            var v96 uint32 = 42
                            _ = v96
                        default:
                            v95 = (^(v95))
                        }
                        if true {
                            v97 := v1[-941]
                            if true {
                                v95 -= (-(v95))
                            } else {
                                v95 = v95
                            }
                            _ = len(v2)
                            _ = v97
                        }
                        var v98 bool = true
                        var v99 uint16 = 42
                        _ = v98
                        _ = v99
                    } else {
                        v100 := GenS11{F0: 0, F1: int16(fi92)}
                        v101 := GenS16{F0: (65535 >= uint(v4)), F1: 9, F2: v95}
                        v4 = 255
                        _ = v100
                        _ = v101
                    }
                    delete(v1, fi92)
                    v102 := GenS3{F0: 2147483647}
                    _ = v102
                }
                v95 += v95
                _ = v95
            }
            v82 -= (v82 << (((v82 | v82)) & 63))
            _ = fi92
            _ = v93
        }
        _ = fi83
        _ = v91
    }
    v82 += genFn1()
    var v103 uint64 = 479
    if false {
        v103 -= (^(v103))
        fmt.Println(v82)
    }
    switch ("a" + "a") {
    case "hello":
        var v104 uint = (v82 << ((v82) & 63))
        _ = v104
    }
    v82 += (v82 & (v82 | v82))
    fmt.Println(v4)
    v103 = 2
    _ = len(v2)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v82
}
