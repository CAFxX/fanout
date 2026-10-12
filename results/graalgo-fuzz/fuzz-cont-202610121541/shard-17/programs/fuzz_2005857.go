package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 string; F1 uint64; F2 int16 }
func genFn0(p0 int8, p1 uint16) int {
    switch "foo bar" {
    case "\n\t":
        p1 += 65535
    case "foo bar":
        switch -1 {
        case 32767:
            for fi4 := 0; fi4 < 7; fi4++ {
                switch (((("" + (((("\n\t" + "foo bar") + "\n\t") + ("a" + "a")) + "foo bar")) + "") + "a") + "foo bar") {
                case "\n\t":
                    fmt.Println(p0, p1)
                case "foo bar":
                    fmt.Println(fi4, fi4)
                case "":
                    p1 = (p1 ^ (p1 ^ (p1 ^ p1)))
                default:
                    p1 = p1
                }
                v5 := []int{}
                for fi6 := 0; fi6 < 5; fi6++ {
                    var v7 uint = 4294967295
                    _ = len(v5)
                    v8 := GenS1{F0: v7, F1: 2147483647, F2: 4294967295}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                _ = fi4
                _ = v5
            }
        default:
            var v9 bool = false
            _ = v9
        }
    default:
        var v10 int8 = (p0 & p0)
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    return -912
}
func genFn1(p0 bool) int8 {
    var v12 int32 = 2147483647
    fmt.Println(v12, p0)
    if false {
        if false {
            fmt.Println(v12)
            var v13 uint64 = 366
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            p0 = p0
            if p0 {
                if (!false) {
                    v15 := GenS1{F0: 475, F1: fi14, F2: 1}
                    _ = v15
                }
                var v16 bool = p0
                _ = v16
            } else {
                var v17 uint16 = 42
                v18 := []int{fi14, 42, (fi14 | (fi14 << ((2) & 63)))}
                _ = v17
                _ = v18
            }
            p0 = p0
            _ = fi14
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            if (42 >= 119) {
                v20 := GenS2{F0: (("a" + ("foo bar" + "")) + "hello"), F1: p0}
                if true {
                    v21 := GenS1{F0: 1, F1: fi19, F2: 651}
                    _ = v21
                }
                v12 -= (v12 | (v12 * (v12 << ((539) & 63))))
                _ = v20
            } else {
                var v22 int64 = 690
                var v23 int16 = 32767
                _ = v22
                _ = v23
            }
            v12 += 547
            _ = fi19
        }
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 string = "hello"
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS3{F0: 72}
            if (((!false) && true) && (2 != 32767)) {
                if true {
                    fmt.Println(v3)
                    if true {
                        var v27 int8 = -1
                        v28 := GenS1{F0: 42, F1: fi25, F2: 2147483648}
                        fmt.Println(fi24)
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint8 = 0
                        _ = v29
                    }
                    var v30 int = fi24
                    if true {
                        if (!false) {
                            var v31 int = (fi24 | fi25)
                            var v32 uint = 268
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS1{F0: 65535, F1: fi24, F2: 65535}
                            var v34 uint64 = 42
                            _ = v33
                            _ = v34
                        }
                        var v35 uint16 = 2
                        _ = v35
                    } else {
                        if false {
                            v30 -= (v30 + v30)
                            switch (v30 << ((42) & 63)) {
                            case 1:
                                v30 = (-(fi24))
                            case -116:
                                switch 0 {
                                case -9223372036854775808:
                                    if false {
                                        var v36 uint16 = 388
                                        switch -413 {
                                        case 9223372036854775807:
                                            v37 := GenS1{F0: 255, F1: (fi25 + genFn0(100, (v36 / ((v36) | 1)))), F2: 1}
                                            _ = v37
                                        default:
                                            switch (v30 - v30) {
                                            case -751:
                                                v30 = (-(fi24))
                                            case 0:
                                                var v38 int8 = -39
                                                _ = v38
                                            case 415:
                                                var v39 uint8 = 2
                                                _ = v39
                                            }
                                        }
                                        v40 := v1[513]
                                        _ = v36
                                        _ = v40
                                    } else {
                                        var v41 uint32 = 2147483648
                                        _ = v41
                                    }
                                }
                            }
                        } else {
                            switch v3 {
                            case "foo bar":
                                switch (v30 & v30) {
                                case 128:
                                    fmt.Println(v3)
                                case -285:
                                    fmt.Println(fi24)
                                }
                                switch "foo bar" {
                                case "a":
                                    v30 = (fi24 | -9223372036854775808)
                                default:
                                    v30 = fi24
                                }
                                v1[-310] = 2147483647
                            case "":
                                v30 = fi25
                            case "hello":
                                v3 = v3
                            }
                            var v42 uint = 255
                            v43 := GenS1{F0: 514, F1: 0, F2: uint32(v42)}
                            _ = v42
                            _ = v43
                        }
                        var v44 uint32 = 4294967295
                        var v45 uint64 = 42
                        v44 -= v44
                        _ = v44
                        _ = v45
                    }
                    if false {
                        v46 := GenS4{F0: "\n\t"}
                        _ = v46
                    } else {
                        if (-128 >= genFn1(true)) {
                            fmt.Println(fi25, fi24, v3)
                        } else {
                            v30 = (fi25 | 128)
                            var v47 uint = 1
                            v48 := []int{(-(fi24)), (fi24 - 32767), (fi25 ^ (-(fi25))), -816}
                            _ = v47
                            _ = v48
                        }
                        fmt.Println(v3, v3, v3)
                        v49 := GenS4{F0: "\n\t"}
                        _ = v49
                    }
                    _ = v30
                } else {
                    var v50 uint32 = 285
                    _ = v50
                }
                v3 = (v3 + v3)
            }
            v51 := GenS2{F0: ((v3 + v3) + v3), F1: true}
            _ = fi25
            _ = v26
            _ = v51
        }
        for fi52 := 0; fi52 < 14; fi52++ {
            switch (fi24 / ((32768) | 1)) {
            case -651:
                if false {
                    var v53 int16 = 396
                    v53 -= 990
                    v54 := []int{}
                    _ = v53
                    _ = v54
                } else {
                    if false {
                        v3 = v3
                        var v55 bool = true
                        _ = v55
                    }
                }
            case 32767:
                v3 = v3
            default:
                fmt.Println(v3)
            }
            _ = fi52
        }
        _ = fi24
    }
    if true {
        if (3 == 495) {
            if ((((v3 + v3) + "hello") + (v3 + ("\n\t" + v3))) == "hello") {
                v56 := GenS5{F0: 405}
                _ = v56
            } else {
                v57 := GenS4{F0: ""}
                var v58 int32 = -697
                _ = v57
                _ = v58
            }
            if false {
                var v59 uint64 = 1
                if (0.0 == (1.5 * (float32(-2.25) * (0.0 * (float32((((((0.1 - float64(3.14159)) * -2.25) - 0.1) + -2.25) * 3.14159)) - 0.1))))) {
                    v59 = v59
                    v60 := GenS2{F0: v3, F1: (false || ((42 >= 436) && (false || true)))}
                    _ = v60
                }
                _ = v59
            }
            var v61 uint8 = 0
            _ = v61
        }
        fmt.Println(v3)
    }
    for fi62 := 0; fi62 < 3; fi62++ {
        v63 := []int{fi62, 0, (fi62 << ((65535) & 63))}
        _ = len(v1)
        _ = fi62
        _ = v63
    }
    var v64 int32 = 1
    v3 = v3
    if true {
        switch genFn0(-94, 38) {
        case 128:
            switch v3 {
            case "a":
                if ((!(true && (1 == 720))) || (!(true || true))) {
                    fmt.Println(v3, v3)
                    v64 = (v64 + 0)
                    var v65 int16 = -599
                    _ = v65
                }
            }
        case -31:
            fmt.Println(v3, v64)
        default:
            var v66 int32 = -362
            _ = v66
        }
        var v67 uint32 = 65535
        switch "hello" {
        case "a":
            var v68 int64 = 653
            _ = v68
        case "\n\t":
            if false {
                fmt.Println(v64)
                if true {
                    v69 := []int{4294967295, 127, 32767, 4294967295}
                    _ = v69
                } else {
                    if true {
                        for fi70 := 0; fi70 < 9; fi70++ {
                            for fi71 := 0; fi71 < 14; fi71++ {
                                var v72 int64 = 975
                                _ = fi71
                                _ = v72
                            }
                            for fi73 := 0; fi73 < 11; fi73++ {
                                v74 := GenS4{F0: ((v3 + v3) + "\n\t")}
                                var v75 string = (("a" + v3) + (((v3 + ("" + ("\n\t" + "hello"))) + v3) + ((v3 + ("a" + "hello")) + v3)))
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                            v67 += (v67 / ((v67) | 1))
                            _ = fi70
                        }
                        var v76 uint32 = v67
                        var v77 bool = (0.1 < ((3.14159 - 0.0) * (((0.0 / 1.5) / (1.5 + 1.5)) / (1.5 * ((0.1 - (-2.25 / 1.5)) - 0.0)))))
                        _ = v76
                        _ = v77
                    }
                    v67 = v67
                    if false {
                        for fi78 := 0; fi78 < 7; fi78++ {
                            for fi79 := 0; fi79 < 4; fi79++ {
                                v64 = -2147483648
                                var v80 string = (v3 + ((v3 + v3) + (("\n\t" + v3) + "foo bar")))
                                delete(v2, 328)
                                v2[(^(v67))] = v67
                                if (false || true) {
                                    var v81 int64 = 9223372036854775807
                                    _ = v81
                                }
                                _ = fi79
                                _ = v80
                            }
                            for fi82 := 0; fi82 < 8; fi82++ {
                                var v83 uint32 = v67
                                v2[(^(v83))] = (^(v67))
                                switch 128 {
                                case 523:
                                    var v84 float64 = -2.25
                                    _ = v84
                                case 275:
                                    v67 += v83
                                    var v85 uint16 = 1
                                    switch fi82 {
                                    case 4294967295:
                                        v86 := []int{1, (fi78 - (^(fi78)))}
                                        _ = v86
                                    case 32768:
                                        v87 := GenS6{F0: 841, F1: v67}
                                        _ = v87
                                    }
                                    _ = v85
                                default:
                                    var v88 string = "a"
                                    _ = v88
                                }
                                _ = fi82
                                _ = v83
                            }
                            v67 = (v67 & v67)
                            _ = fi78
                        }
                        switch -1 {
                        case -2147483648:
                            var v89 int8 = -6
                            _ = v89
                        }
                        v90 := []int{}
                        _ = v90
                    }
                }
                fmt.Println(v3, v3, v64)
            }
        case "hello":
            var v91 float32 = -2.25
            _ = v91
        }
        _ = v67
    } else {
        v64 = -2147483648
        v92 := GenS5{F0: 42}
        _ = v92
    }
    v3 = "foo bar"
    if ((-2.25 - 0.1) < 1.5) {
        var v93 int8 = 0
        fmt.Println(v93)
        _ = v93
    } else {
        var v94 bool = ("foo bar" > (v3 + v3))
        switch 32768 {
        case 512:
            var v95 int16 = 547
            _ = v95
        case 9223372036854775807:
            var v96 int16 = -32768
            _ = v96
        case -789:
            _ = len(v1)
        default:
            v97 := v2[2]
            _ = v97
        }
        _ = v94
    }
    v3 = ""
    for fi98 := 0; fi98 < 17; fi98++ {
        if false {
            for fi99 := 0; fi99 < 10; fi99++ {
                v100 := GenS3{F0: 705}
                _ = fi99
                _ = v100
            }
        } else {
            _ = len(v2)
            for fi101 := 0; fi101 < 1; fi101++ {
                v102 := GenS7{F0: (v3 + v3), F1: 192, F2: 32767}
                v64 = (-(v64))
                _ = fi101
                _ = v102
            }
            for fi103 := 0; fi103 < 12; fi103++ {
                v3 = v3
                v104 := []int{(fi103 >> ((786) & 63)), (fi98 - fi98)}
                _ = fi103
                _ = v104
            }
        }
        switch v3 {
        case "hello":
            v105 := v1[genFn0(127, 2)]
            _ = v105
        }
        v64 = (v64 + -1)
        _ = fi98
    }
    var v106 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v64
}
