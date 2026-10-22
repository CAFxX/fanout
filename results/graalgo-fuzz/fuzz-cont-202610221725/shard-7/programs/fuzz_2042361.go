package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int32; F1 int32; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint16, p1 bool, p2 bool) uint64 {
    switch "a" {
    case "":
        for fi4 := 0; fi4 < 17; fi4++ {
            for fi5 := 0; fi5 < 17; fi5++ {
                var v6 uint8 = 2
                v7 := []int{(fi5 + -9223372036854775808), fi4}
                var v8 int8 = 1
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            fmt.Println(fi4)
            _ = fi4
        }
    case "\n\t":
        fmt.Println(p2)
    case "hello":
        if (!(uint8(p0) > 187)) {
            v9 := GenS1{F0: p2}
            var v10 float32 = -2.25
            _ = v10
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 14; fi11++ {
            switch "foo bar" {
            case "\n\t":
                p0 = p0
            case "foo bar":
                v12 := []int{1, (^(fi11))}
                _ = v12
            case "":
                v13 := []int{fi11, (-(fi11)), (fi11 - fi11), 0}
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                if p1 {
                    var v15 int8 = 127
                    var v16 uint16 = p0
                    var v17 float64 = (1e10 * 0.1)
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    p1 = (100 == -2)
                }
                p0 -= 164
                p2 = p1
                _ = fi14
            }
            v18 := GenS2{F0: 1, F1: 289, F2: false}
            _ = fi11
            _ = v18
        }
    }
    v19 := []int{-395, 32768}
    if p2 {
        var v20 uint = 2
        fmt.Println(p0, p1)
        for fi21 := 0; fi21 < 4; fi21++ {
            if len(v19) > 0 {
                v19[0] = (^(fi21))
            }
            switch (fi21 / ((-2147483648) | 1)) {
            case 32768:
                for fi22 := 0; fi22 < 1; fi22++ {
                    if (p0 >= 65535) {
                        v23 := GenS2{F0: 0, F1: 0, F2: true}
                        v24 := GenS3{F0: true, F1: fi21, F2: (v20 / (((v20 >> ((v20) & 63))) | 1))}
                        _ = v23
                        _ = v24
                    } else {
                        var v25 string = (("foo bar" + "hello") + ("a" + (("foo bar" + ("\n\t" + "\n\t")) + "a")))
                        fmt.Println(v25)
                        var v26 uint8 = 38
                        _ = v25
                        _ = v26
                    }
                    if (1 < -299) {
                        var v27 uint8 = 255
                        if len(v19) > 0 {
                            v19[0] = (fi21 | (fi21 ^ (fi21 - fi21)))
                        }
                        _ = v27
                    } else {
                        switch fi21 {
                        case 127:
                            v19 = append(v19, fi22)
                            v19 = append(v19, 32767)
                            switch "hello" {
                            case "hello":
                                if len(v19) > 0 {
                                    v19[0] = 128
                                }
                            case "foo bar":
                                p0 = (p0 + 15)
                            case "a":
                                v28 := GenS2{F0: -225, F1: -926, F2: p1}
                                _ = v28
                            default:
                                fmt.Println(p1)
                            }
                        case 128:
                            fmt.Println(v20, p0)
                        }
                        var v29 bool = true
                        switch (("hello" + "foo bar") + "hello") {
                        case "\n\t":
                            var v30 float32 = 3.14159
                            _ = v30
                        case "":
                            switch -637 {
                            case -2147483648:
                                var v31 uint8 = 0
                                _ = v31
                            }
                        }
                        _ = v29
                    }
                    fmt.Println(v20, fi21, p1)
                    _ = fi22
                }
            case -458:
                if (!p2) {
                    var v32 int = fi21
                    if (v20 <= (v20 | 4294967295)) {
                        v20 = v20
                    } else {
                        var v33 float64 = ((0.1 + 1.5) + -2.25)
                        for fi34 := 0; fi34 < 6; fi34++ {
                            var v35 uint32 = 487
                            var v36 int = 42
                            fmt.Println(fi21)
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        if (!p1) {
                            var v37 uint64 = 0
                            _ = v37
                        }
                        _ = v33
                    }
                    _ = v32
                }
            }
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 int32 = 2147483647
                v40 := GenS1{F0: ("foo bar" > ("a" + ""))}
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi21
        }
        _ = v20
    } else {
        v41 := GenS2{F0: -534, F1: 42, F2: p1}
        if p2 {
            for fi42 := 0; fi42 < 7; fi42++ {
                var v43 uint64 = 18446744073709551615
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi42
                _ = v43
            }
        } else {
            p1 = (!(((255 >= 161) && ("\n\t" >= "\n\t")) && p2))
            var v44 uint32 = 2
            _ = v44
        }
        p2 = p2
        _ = v41
    }
    var v45 uint = 635
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v45
    return 2
}
func genFn1(p0 bool) uint {
    v46 := GenS4{F0: 42, F1: -128, F2: 216}
    for fi47 := 0; fi47 < 8; fi47++ {
        switch fi47 {
        case 551:
            p0 = false
        default:
            p0 = (-2147483648 != 0)
        }
        _ = fi47
    }
    switch -1 {
    case -217:
        v48 := GenS4{F0: 65535, F1: 127, F2: 0}
        _ = v48
    default:
        if false {
            var v49 float32 = 3.14159
            v49 = v49
            if p0 {
                var v50 uint = 18446744073709551615
                _ = v50
            } else {
                var v51 float64 = (float64(v49) + 0.0)
                if p0 {
                    var v52 int64 = 9223372036854775807
                    v53 := GenS2{F0: 42, F1: 0, F2: p0}
                    var v54 int = 128
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    switch 784 {
                    case 0:
                        v55 := []int{9223372036854775807, -2147483648}
                        _ = v55
                    case 127:
                        v56 := []int{0, 760}
                        _ = v56
                    }
                    var v57 float64 = (((1.5 / v51) / 1e10) - v51)
                    v51 -= ((v51 + ((float64(3.14159) + (v57 / v57)) - 1e10)) / v51)
                    _ = v57
                }
                var v58 uint = 0
                _ = v51
                _ = v58
            }
            _ = v49
        } else {
            v59 := GenS4{F0: 1, F1: -128, F2: 0}
            if p0 {
                switch 32767 {
                case 998:
                    for fi60 := 0; fi60 < 10; fi60++ {
                        fmt.Println(p0, fi60)
                        _ = fi60
                    }
                case 32767:
                    for fi61 := 0; fi61 < 20; fi61++ {
                        p0 = (!(!p0))
                        _ = fi61
                    }
                default:
                    if false {
                        p0 = p0
                    }
                }
            }
            _ = v59
        }
    }
    var v62 int64 = -1
    if (!(2 != 845)) {
        v63 := []int{128, 32767, 9223372036854775807, -2147483648}
        _ = v63
    }
    _ = p0
    _ = v46
    _ = v62
    return 18446744073709551615
}

func main() {
    var v1 int = 127
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    fmt.Println(v1)
    v1 += (v1 & (v1 >> ((1) & 63)))
    var v64 uint32 = 2
    var v65 string = "\n\t"
    if false {
        v1 = (v1 % (((v1 ^ (v1 | v1))) | 1))
        if true {
            v1 -= (v1 + v1)
            var v66 int8 = -53
            _ = v66
        } else {
            var v67 int8 = -1
            v64 = (v64 / (((v64 + 2147483647)) | 1))
            if false {
                if (32767 == -1) {
                    var v68 int32 = -2147483648
                    switch (v1 + (v1 & 881)) {
                    case 128:
                        if false {
                            var v69 int16 = 1
                            _ = v69
                        }
                    case 4294967295:
                        v70 := GenS2{F0: 293, F1: v68, F2: true}
                        _ = v70
                    case -2147483648:
                        v71 := []int{(v1 ^ 32767), 1}
                        _ = v71
                    default:
                        for fi72 := 0; fi72 < 3; fi72++ {
                            switch v65 {
                            case "":
                                v3[(v65 + (v65 + "a"))] = ((v1 | (fi72 * (fi72 % ((569) | 1)))) == v1)
                            }
                            _ = fi72
                        }
                    }
                    _ = v68
                }
                v73 := []int{(v1 / (((v1 % ((-167) | 1))) | 1)), v1}
                _ = v73
            } else {
                var v74 int32 = -2147483648
                for fi75 := 0; fi75 < 17; fi75++ {
                    var v76 string = (v65 + "")
                    _ = fi75
                    _ = v76
                }
                var v77 bool = true
                _ = v74
                _ = v77
            }
            _ = v67
        }
    }
    v78 := GenS5{F0: v64, F1: -917, F2: v65}
    fmt.Println(v1, v1)
    if (32767 < 1) {
        v1 += v1
    } else {
        v65 = v65
        var v79 int = v1
        var v80 uint16 = 0
        _ = v79
        _ = v80
    }
    var v81 uint = 65535
    var v82 bool = (!false)
    switch v1 {
    case 0:
        var v83 int16 = 1
        _ = v83
    case -598:
        switch (v1 + (v1 >> ((18446744073709551615) & 63))) {
        case 1:
            v64 = v64
        default:
            v84 := []int{v1, v1, 128, 0}
            _ = v84
        }
    case 32767:
        v82 = v82
    }
    var v85 bool = v82
    v86 := v2[1]
    var v87 int = 0
    if (genFn1((-2147483648 == -213)) > (v81 % ((65535) | 1))) {
        var v88 float32 = 0.0
        v89 := GenS6{F0: v85, F1: 32767}
        _ = v88
        _ = v89
    }
    if true {
        if true {
            var v90 string = v65
            _ = v90
        } else {
            var v91 float32 = (float32(0.1) - -2.25)
            if ((!v82) && (-1 < -1)) {
                v92 := GenS7{F0: 42, F1: 856, F2: (^(v81))}
                for fi93 := 0; fi93 < 9; fi93++ {
                    fmt.Println(v86)
                    for fi94 := 0; fi94 < 17; fi94++ {
                        var v95 bool = v86
                        _ = fi94
                        _ = v95
                    }
                    _ = fi93
                }
                var v96 uint8 = 233
                _ = v92
                _ = v96
            } else {
                for fi97 := 0; fi97 < 16; fi97++ {
                    for fi98 := 0; fi98 < 14; fi98++ {
                        _ = len(v2)
                        if ((("hello" + v65) + v65) <= v65) {
                            v1 = (fi97 % ((v1) | 1))
                            var v99 uint8 = 128
                            v99 = (v99 >> ((42) & 63))
                            if v86 {
                                v100 := GenS5{F0: v64, F1: -530, F2: (v65 + ((v65 + "hello") + (v65 + (v65 + "\n\t"))))}
                                v85 = v86
                                _ = v100
                            } else {
                                var v101 uint = 42
                                v87 += (fi97 - (^(fi98)))
                                _ = v101
                            }
                            if (!(42 < uint64(v81))) {
                                var v102 bool = ((2 >= 646) || false)
                                var v103 uint8 = 255
                                _ = v102
                                _ = v103
                            }
                            v104 := []int{}
                            _ = v104
                            _ = v99
                        }
                        var v105 int16 = 754
                        _ = fi98
                        _ = v105
                    }
                    delete(v2, (fi97 ^ fi97))
                    _ = fi97
                }
                fmt.Println(v87, v82)
                v106 := GenS8{F0: 0}
                _ = v106
            }
            _ = v91
        }
    } else {
        v107 := GenS9{F0: 1, F1: 128}
        _ = v107
    }
    fmt.Println(v87)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v64
    _ = v65
    _ = v78
    _ = v81
    _ = v82
    _ = v85
    _ = v86
    _ = v87
}
