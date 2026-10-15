package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint32; F2 bool }
type GenS5 struct { F0 uint; F1 bool; F2 uint }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int; F1 int8; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 int16; F2 uint32 }
func genFn0() int8 {
    switch 0 {
    case 422:
        v4 := GenS1{F0: 1, F1: "", F2: (("" + "hello") + "foo bar")}
        _ = v4
    default:
        v5 := []int{-79}
        _ = v5
    }
    var v6 uint16 = 32767
    _ = v6
    return 127
}
func genFn1() int32 {
    var v7 uint64 = 430
    switch 0 {
    case 1:
        fmt.Println(v7)
    case 0:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int64 = -17
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint32 = uint32(v7)
                if true {
                    var v12 bool = false
                    _ = v12
                } else {
                    v9 = (v9 << ((255) & 63))
                }
                var v13 bool = (0 > 802)
                _ = fi10
                _ = v11
                _ = v13
            }
            if ("foo bar" < "\n\t") {
                var v14 uint64 = (v7 / ((1) | 1))
                switch "\n\t" {
                case "hello":
                    if false {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v9 = (^(v9))
                            v16 := []int{(fi8 % ((fi15) | 1)), -1}
                            v7 = (v7 + 1)
                            _ = fi15
                            _ = v16
                        }
                    }
                case "":
                    if (!false) {
                        if false {
                            fmt.Println(v14, fi8, v7)
                            v17 := GenS1{F0: 0, F1: (("a" + "a") + "\n\t"), F2: "hello"}
                            v18 := []int{int(v14), (fi8 / ((fi8) | 1))}
                            _ = v17
                            _ = v18
                        }
                    }
                case "\n\t":
                    v19 := GenS2{F0: 255, F1: 0, F2: v9}
                    _ = v19
                default:
                    if ("hello" >= "foo bar") {
                        v14 += (v14 % ((v7) | 1))
                    } else {
                        var v20 uint32 = 4294967295
                        var v21 int16 = -32768
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v9 -= v9
                            v23 := []int{1, fi22, fi22}
                            var v24 int64 = (v9 - v9)
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v20
                        _ = v21
                    }
                }
                _ = v14
            }
            _ = fi8
            _ = v9
        }
    case -818:
        fmt.Println(v7)
    default:
        if (1 <= 0) {
            v7 = (v7 / ((889) | 1))
        } else {
            var v25 int16 = 42
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v25, v25, v25)
                for fi27 := 0; fi27 < 18; fi27++ {
                    if false {
                        if (((v25 % ((v25) | 1)) == (v25 * (v25 | (v25 & v25)))) || (true && ((!true) || (true || (42 < 662))))) {
                            if true {
                                var v28 int16 = (v25 - (v25 >> ((65535) & 63)))
                                if false {
                                    v29 := GenS2{F0: 128, F1: (v28 & v28), F2: 0}
                                    _ = v29
                                } else {
                                    var v30 int64 = int64(v7)
                                    v31 := GenS1{F0: -1, F1: "foo bar", F2: "foo bar"}
                                    v28 = v25
                                    v28 += v28
                                    _ = v30
                                    _ = v31
                                }
                                _ = v28
                            } else {
                                v25 += 1
                            }
                        } else {
                            v25 += v25
                        }
                        var v32 uint = 18446744073709551615
                        v25 -= int16(fi27)
                        v33 := []int{(fi26 >> ((94) & 63)), (fi27 + (fi26 * (fi27 / ((1) | 1))))}
                        if (!true) {
                            v25 = v25
                            v34 := GenS3{F0: false}
                            var v35 uint = v32
                            _ = v34
                            _ = v35
                        } else {
                            var v36 int32 = -2147483648
                            v25 -= (v25 | (v25 * -1))
                            _ = v36
                        }
                        _ = v32
                        _ = v33
                    } else {
                        if true {
                            v37 := GenS1{F0: -128, F1: (("\n\t" + "") + "hello"), F2: "\n\t"}
                            v38 := GenS4{F0: 42, F1: 65535, F2: (65535 > 875)}
                            var v39 int32 = -171
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                    v7 = v7
                    switch fi27 {
                    case 68:
                        fmt.Println(fi26, v7, fi26)
                    case -9223372036854775808:
                        var v40 uint8 = 42
                        _ = v40
                    default:
                        var v41 bool = (!true)
                        v42 := GenS4{F0: 0, F1: 65535, F2: (100 == 121)}
                        v43 := GenS5{F0: 2, F1: (v25 > (v25 - 550)), F2: 0}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi27
                }
                var v44 float32 = 0.1
                _ = fi26
                _ = v44
            }
            _ = v25
        }
    }
    if true {
        v7 = v7
        v45 := GenS3{F0: (-1 >= 22)}
        v46 := []int{}
        _ = v45
        _ = v46
    } else {
        for fi47 := 0; fi47 < 20; fi47++ {
            if (!(v7 < (v7 | (v7 / ((4294967295) | 1))))) {
                v48 := GenS6{F0: 42, F1: -1}
                var v49 int16 = 1000
                v7 = (v7 * v7)
                _ = v48
                _ = v49
            }
            fmt.Println(v7, v7)
            _ = fi47
        }
        switch "a" {
        case "hello":
            v50 := []int{4294967295}
            _ = v50
        case "\n\t":
            var v51 uint32 = 0
            _ = v51
        case "a":
            v52 := GenS1{F0: genFn0(), F1: "\n\t", F2: "\n\t"}
            _ = v52
        }
    }
    switch 303 {
    case -242:
        switch -9223372036854775808 {
        case 564:
            var v53 int = -152
            _ = v53
        case 752:
            var v54 uint64 = 1
            _ = v54
        }
    case 207:
        if (255 < 83) {
            var v55 float32 = (-2.25 - 1.5)
            _ = v55
        } else {
            v56 := GenS7{F0: 771, F1: 127, F2: -278}
            v57 := []int{-9223372036854775808, 2147483647}
            _ = v56
            _ = v57
        }
    case -698:
        fmt.Println(v7)
    default:
        if (!((3.14159 / (3.14159 / (1.5 + -2.25))) >= 1.5)) {
            switch -2147483648 {
            case -643:
                v58 := GenS4{F0: 42, F1: 655, F2: true}
                _ = v58
            case -2147483648:
                v59 := GenS4{F0: 42, F1: 885, F2: (true && true)}
                _ = v59
            default:
                for fi60 := 0; fi60 < 10; fi60++ {
                    var v61 uint16 = 65535
                    var v62 bool = false
                    var v63 uint8 = 128
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
            var v64 float32 = (float32(((float64(1.5) / float64((3.14159 * -2.25))) / 0.1)) - (3.14159 * 0.0))
            for fi65 := 0; fi65 < 11; fi65++ {
                if true {
                    if (494 > 0) {
                        var v66 int16 = 32767
                        v67 := []int{128, fi65, 32767}
                        _ = v66
                        _ = v67
                    } else {
                        var v68 int32 = -1
                        v69 := []int{(^(fi65)), fi65, fi65, (fi65 & (^(fi65)))}
                        if true {
                            for fi70 := 0; fi70 < 16; fi70++ {
                                v68 -= (v68 << ((42) & 63))
                                _ = len(v69)
                                var v71 int32 = v68
                                _ = fi70
                                _ = v71
                            }
                            var v72 int64 = -1
                            _ = v72
                        }
                        _ = v68
                        _ = v69
                    }
                }
                fmt.Println(v7, v7)
                _ = fi65
            }
            _ = v64
        } else {
            switch -467 {
            case 4294967295:
                v7 += 0
            case 132:
                if false {
                    switch -9223372036854775808 {
                    case 194:
                        if (false || (65535 > 2)) {
                            fmt.Println(v7)
                            v73 := GenS6{F0: 65535, F1: 1}
                            var v74 uint = 0
                            _ = v73
                            _ = v74
                        } else {
                            for fi75 := 0; fi75 < 12; fi75++ {
                                switch (fi75 | fi75) {
                                case -319:
                                    if false {
                                        v7 = (v7 * 2)
                                        v76 := []int{}
                                        _ = v76
                                    }
                                }
                                var v77 uint8 = 127
                                v7 = v7
                                _ = fi75
                                _ = v77
                            }
                            v7 -= 2
                        }
                    case -9223372036854775808:
                        var v78 int = 32768
                        _ = v78
                    case 242:
                        if (!true) {
                            for fi79 := 0; fi79 < 15; fi79++ {
                                var v80 int32 = 42
                                _ = fi79
                                _ = v80
                            }
                            v7 -= v7
                            var v81 bool = (18446744073709551615 < 255)
                            _ = v81
                        } else {
                            for fi82 := 0; fi82 < 19; fi82++ {
                                var v83 uint32 = 1
                                _ = fi82
                                _ = v83
                            }
                        }
                    default:
                        var v84 bool = true
                        _ = v84
                    }
                    fmt.Println(v7)
                } else {
                    var v85 int16 = 1000
                    fmt.Println(v85, v85)
                    v86 := GenS3{F0: (false || true)}
                    _ = v85
                    _ = v86
                }
            }
            var v87 int16 = 32767
            _ = v87
        }
    }
    _ = v7
    return -2147483648
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = 453
    var v3 int8 = 0
    fmt.Println(v2, v1, v2)
    v88 := GenS8{F0: -2147483648}
    switch -2147483648 {
    case -2147483648:
        var v89 uint16 = 32767
        _ = v89
    }
    v90 := GenS9{F0: 2147483647, F1: (v2 & (v2 | 552)), F2: 547}
    v2 += v2
    switch 32768 {
    case -9223372036854775808:
        for fi91 := 0; fi91 < 17; fi91++ {
            if true {
                v92 := []int{0}
                var v93 bool = (2 <= 255)
                fmt.Println(v2, v1)
                _ = v92
                _ = v93
            }
            v94 := []int{fi91, 2147483647}
            v94 = append(v94, fi91)
            _ = fi91
            _ = v94
        }
    case 42:
        for fi95 := 0; fi95 < 7; fi95++ {
            for fi96 := 0; fi96 < 7; fi96++ {
                var v97 uint = 1
                if true {
                    var v98 uint64 = 2
                    v97 = (v97 * (^(v97)))
                    v99 := []int{(fi95 ^ 2147483647), fi96, fi95, (-(fi95))}
                    _ = v98
                    _ = v99
                } else {
                    v100 := []int{(-(fi96)), 1}
                    v101 := GenS5{F0: 42, F1: false, F2: v97}
                    _ = v100
                    _ = v101
                }
                _ = fi96
                _ = v97
            }
            _ = fi95
        }
    }
    v102 := GenS9{F0: 42, F1: (v2 ^ -947), F2: 42}
    var v103 float32 = 1e10
    v103 -= 0.1
    v104 := GenS4{F0: -9223372036854775808, F1: 2, F2: (0.1 <= v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v88
    _ = v90
}
