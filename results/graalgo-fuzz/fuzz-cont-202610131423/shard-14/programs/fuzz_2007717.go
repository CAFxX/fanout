package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int8; F1 bool; F2 bool }
type GenS9 struct { F0 int; F1 string; F2 bool }
type GenS10 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int64; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int) int64 {
    v7 := []int{p1, p1}
    v8 := GenS1{F0: ((p1 < 32768) || (65535 <= 2)), F1: 65535}
    var v9 uint16 = 65535
    var v10 uint = 783
    p0 += p0
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float32 = (-2.25 / 3.14159)
    var v2 int8 = -1
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    for fi11 := 0; fi11 < 16; fi11++ {
        v1 += (-2.25 / v1)
        _ = fi11
    }
    v12 := GenS1{F0: false, F1: 30}
    fmt.Println(v3, v2, v2)
    for fi13 := 0; fi13 < 2; fi13++ {
        if false {
            v4[0] = 42
            v14 := GenS2{F0: 65535, F1: 128}
            _ = v14
        }
        var v15 string = v3
        v16 := GenS1{F0: false, F1: 32767}
        _ = fi13
        _ = v15
        _ = v16
    }
    fmt.Println(v2)
    v17 := GenS3{F0: 1, F1: 345, F2: -128}
    var v18 uint = 18446744073709551615
    delete(v5, 1)
    if true {
        var v19 int = 2147483647
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint8 = 255
            v19 = 127
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        switch 42 {
        case 0:
            for fi22 := 0; fi22 < 13; fi22++ {
                v18 -= 0
                _ = fi22
            }
        case 806:
            var v23 uint32 = 1
            _ = v23
        case 66:
            switch 32767 {
            case 381:
                v24 := GenS4{F0: 0, F1: 2147483647}
                _ = v24
            case 127:
                for fi25 := 0; fi25 < 19; fi25++ {
                    if false {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := []int{fi26, fi26, (^(fi25)), fi25}
                            var v28 int16 = -1
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        delete(v6, "")
                    } else {
                        fmt.Println(v18)
                        switch (-(fi25)) {
                        case 127:
                            v29 := GenS4{F0: 137, F1: (fi25 >> (((v18 / (((v18 % ((v18) | 1))) | 1))) & 63))}
                            _ = v29
                        case -197:
                            v2 -= v2
                        case 128:
                            switch "" {
                            case "\n\t":
                                v30 := GenS5{F0: 323, F1: 0, F2: -969}
                                _ = v30
                            case "foo bar":
                                fmt.Println(v1, v2, v18)
                            }
                        }
                        v31 := []int{(fi25 % ((128) | 1)), (-(fi25))}
                        _ = v31
                    }
                    _ = fi25
                }
            default:
                switch 128 {
                case -9223372036854775808:
                    switch 42 {
                    case 32767:
                        if (true && (42 != 255)) {
                            v2 = (v2 % ((v2) | 1))
                            if (false || ((v18 * 0) >= v18)) {
                                if true {
                                    if true {
                                        switch -687 {
                                        case -2147483648:
                                            fmt.Println(v18, v18)
                                        case -493:
                                            v32 := GenS6{F0: 2147483647, F1: -1, F2: 18446744073709551615}
                                            _ = v32
                                        case -748:
                                            if (!(42 != 2147483647)) {
                                                if true {
                                                    v3 = (v3 + ("\n\t" + v3))
                                                } else {
                                                    v33 := GenS3{F0: 0, F1: 0, F2: (v2 & (^(v2)))}
                                                    if (42 == 9223372036854775807) {
                                                        var v34 uint16 = 1000
                                                        if true {
                                                            v2 += 1
                                                            var v35 float32 = float32(1.5)
                                                            _ = v35
                                                        } else {
                                                            var v36 uint16 = (v34 >> (((v18 & v18)) & 63))
                                                            for fi37 := 0; fi37 < 13; fi37++ {
                                                                for fi38 := 0; fi38 < 7; fi38++ {
                                                                    if true {
                                                                        if (0 >= 140) {
                                                                            v18 -= (v18 % (((v18 << (((v18 | v18)) & 63))) | 1))
                                                                            if (true || (127 != 2)) {
                                                                                var v39 uint64 = 4294967295
                                                                                v40 := []int{(fi38 & 717), fi38, fi38, fi38}
                                                                                _ = v39
                                                                                _ = v40
                                                                            } else {
                                                                                if false {
                                                                                    v34 = 1
                                                                                    var v41 uint16 = 489
                                                                                    _ = v41
                                                                                } else {
                                                                                    v42 := GenS1{F0: false, F1: v34}
                                                                                    var v43 int = (fi38 + (fi38 ^ (fi38 * (fi38 / ((fi38) | 1)))))
                                                                                    var v44 int64 = -9223372036854775808
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                    _ = v44
                                                                                }
                                                                                switch ((v3 + (v3 + "")) + ("hello" + v3)) {
                                                                                case "foo bar":
                                                                                    v45 := GenS4{F0: 127, F1: (fi37 + fi37)}
                                                                                    _ = v45
                                                                                default:
                                                                                    switch (fi37 ^ fi38) {
                                                                                    case 2147483647:
                                                                                        v46 := GenS2{F0: 2147483647, F1: 19}
                                                                                        _ = v46
                                                                                    case 670:
                                                                                        fmt.Println(v34, v34, v1)
                                                                                    case 1:
                                                                                        if (-2.25 <= (1.5 - (1e10 - 1e10))) {
                                                                                            var v47 bool = true
                                                                                            var v48 uint16 = 189
                                                                                            v49 := GenS7{F0: (v18 % (((v18 >> ((18446744073709551615) & 63))) | 1)), F1: 157}
                                                                                            _ = v47
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch ("a" + "foo bar") {
                                                                                case "hello":
                                                                                    v50 := []int{fi38, fi37}
                                                                                    _ = v50
                                                                                }
                                                                                fmt.Println(v36, v2)
                                                                                v1 = ((v1 + float32(3.14159)) - float32(0.1))
                                                                            }
                                                                            var v51 float32 = (0.0 * (v1 - v1))
                                                                            _ = v51
                                                                        } else {
                                                                            v52 := GenS3{F0: 4294967295, F1: 4294967295, F2: (v2 >> ((18446744073709551615) & 63))}
                                                                            if false {
                                                                                v34 = v34
                                                                                if true {
                                                                                    if true {
                                                                                        v53 := []int{fi38, fi37, (fi38 / (((fi38 + 0)) | 1))}
                                                                                        v6[(v3 + (("\n\t" + "") + ""))] = (fi37 + fi37)
                                                                                        _ = v53
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v6)
                                                                                    var v54 bool = true
                                                                                    _ = v54
                                                                                }
                                                                                v36 += 829
                                                                            } else {
                                                                                fmt.Println(v2, v18)
                                                                                v36 += v36
                                                                            }
                                                                            _ = v52
                                                                        }
                                                                        fmt.Println(v34, v34)
                                                                    }
                                                                    _ = fi38
                                                                }
                                                                v55 := GenS1{F0: false, F1: (v36 - (^(v36)))}
                                                                v56 := []int{fi37, int(v36)}
                                                                _ = fi37
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                            _ = v36
                                                        }
                                                        for fi57 := 0; fi57 < 19; fi57++ {
                                                            var v58 int32 = 944
                                                            v59 := GenS1{F0: (true && (238 != 183)), F1: (v34 % ((65535) | 1))}
                                                            _ = fi57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v34
                                                    } else {
                                                        var v60 int8 = v2
                                                        if (!(!true)) {
                                                            if (!true) {
                                                                v3 = "\n\t"
                                                            } else {
                                                                v6[v3] = -671
                                                                v18 += (v18 * (v18 & (v18 + (^(v18)))))
                                                            }
                                                            var v61 int32 = 592
                                                            v62 := GenS7{F0: v18, F1: 4294967295}
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            if false {
                                                                var v63 uint = v18
                                                                if true {
                                                                    fmt.Println(v2, v2)
                                                                    v1 -= float32(3.14159)
                                                                }
                                                                _ = v63
                                                            }
                                                            var v64 uint32 = 463
                                                            v65 := GenS8{F0: (v60 % ((v2) | 1)), F1: true, F2: (!true)}
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        var v66 int32 = -892
                                                        _ = v60
                                                        _ = v66
                                                    }
                                                    _ = v33
                                                }
                                                if true {
                                                    delete(v5, 545)
                                                }
                                            }
                                        }
                                        fmt.Println(v2)
                                    }
                                } else {
                                    if false {
                                        if false {
                                            var v67 int32 = 2147483647
                                            if false {
                                                v68 := GenS8{F0: 1, F1: true, F2: ((!false) || true)}
                                                _ = v68
                                            } else {
                                                v69 := GenS3{F0: 0, F1: 711, F2: -83}
                                                v70 := GenS4{F0: 66, F1: 128}
                                                switch v3 {
                                                case "hello":
                                                    v18 = 451
                                                case "a":
                                                    v1 = 0.0
                                                }
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = v67
                                        }
                                        var v71 uint64 = 1
                                        _ = v71
                                    } else {
                                        if ((2 <= 57) && (2 > 145)) {
                                            delete(v6, "hello")
                                            v72 := []int{0, 127}
                                            _ = v72
                                        } else {
                                            v3 = ((v3 + "") + v3)
                                            v3 = ("" + (v3 + ((((v3 + v3) + v3) + (v3 + ("a" + (v3 + v3)))) + "")))
                                            v73 := GenS9{F0: 1, F1: "foo bar", F2: false}
                                            _ = v73
                                        }
                                    }
                                    switch v3 {
                                    case "\n\t":
                                        v74 := GenS6{F0: 2147483648, F1: 2147483647, F2: (v18 & 18446744073709551615)}
                                        _ = v74
                                    case "":
                                        for fi75 := 0; fi75 < 10; fi75++ {
                                            fmt.Println(v18, v1, v18)
                                            if (false && (false || (!true))) {
                                                var v76 uint8 = 2
                                                v76 = v76
                                                _ = v76
                                            } else {
                                                if (((!true) || (!(v2 != v2))) && true) {
                                                    var v77 uint16 = 255
                                                    if true {
                                                        for fi78 := 0; fi78 < 11; fi78++ {
                                                            if ("foo bar" < v3) {
                                                                fmt.Println(fi78)
                                                            }
                                                            _ = fi78
                                                        }
                                                        v79 := v6["a"]
                                                        _ = v79
                                                    }
                                                    var v80 int8 = v2
                                                    _ = v77
                                                    _ = v80
                                                } else {
                                                    v2 = (v2 * (v2 << ((v18) & 63)))
                                                    v81 := []int{(^(fi75)), fi75}
                                                    if (!true) {
                                                        v82 := GenS10{F0: 2, F1: 163, F2: 32767}
                                                        v1 = -2.25
                                                        v18 = v18
                                                        _ = v82
                                                    }
                                                    _ = v81
                                                }
                                            }
                                            _ = fi75
                                        }
                                    case "foo bar":
                                        for fi83 := 0; fi83 < 8; fi83++ {
                                            if (!true) {
                                                v2 = v2
                                                v1 += float32(1.5)
                                                switch ("a" + v3) {
                                                case "":
                                                    for fi84 := 0; fi84 < 15; fi84++ {
                                                        if (false || (!false)) {
                                                            v85 := v6["a"]
                                                            _ = v85
                                                        } else {
                                                            _ = len(v6)
                                                        }
                                                        v18 = v18
                                                        if (111 <= 1) {
                                                            if (true && ("a" <= v3)) {
                                                                var v86 uint64 = 732
                                                                _ = v86
                                                            }
                                                            if (genFn0((v2 - (v2 + (^(v2)))), 42) < 9223372036854775807) {
                                                                delete(v5, 65535)
                                                            } else {
                                                                if (!((-537 != -1) || true)) {
                                                                    fmt.Println(fi84)
                                                                    var v87 uint64 = 1
                                                                    if ((fi84 % ((32768) | 1)) <= fi83) {
                                                                        if true {
                                                                            if false {
                                                                                v88 := GenS4{F0: 128, F1: 4294967295}
                                                                                _ = v88
                                                                            }
                                                                            v89 := GenS4{F0: 93, F1: -565}
                                                                            v1 = v1
                                                                            v2 = (v2 % (((v2 % ((-128) | 1))) | 1))
                                                                            _ = v89
                                                                        }
                                                                    } else {
                                                                        v1 = float32(0.1)
                                                                        var v90 bool = ((v2 != 100) || true)
                                                                        var v91 uint32 = 0
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    if false {
                                                                        var v92 float32 = float32(0.0)
                                                                        _ = v92
                                                                    } else {
                                                                        v18 = (^(v18))
                                                                        v93 := v5[653]
                                                                        switch int(v2) {
                                                                        case 32767:
                                                                            switch fi83 {
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v2, v87)
                                                                            }
                                                                        case 127:
                                                                            v87 = uint64(v18)
                                                                        case 32768:
                                                                            v94 := []int{4294967295}
                                                                            var v95 bool = false
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                        if false {
                                                                            v2 = v2
                                                                        }
                                                                        delete(v4, (v93 << (((^(v18))) & 63)))
                                                                        v96 := GenS11{F0: v2}
                                                                        _ = v93
                                                                        _ = v96
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    var v97 uint16 = 950
                                                                    _ = v97
                                                                }
                                                            }
                                                        } else {
                                                            var v98 uint16 = 984
                                                            _ = v98
                                                        }
                                                        _ = fi84
                                                    }
                                                default:
                                                    fmt.Println(v1, v3, v18)
                                                }
                                            } else {
                                                v18 -= v18
                                                var v99 bool = (65535 == 371)
                                                for fi100 := 0; fi100 < 12; fi100++ {
                                                    var v101 float32 = v1
                                                    v102 := v4[2147483648]
                                                    var v103 float32 = v101
                                                    switch v3 {
                                                    case "hello":
                                                        fmt.Println(fi100)
                                                    default:
                                                        switch ((v3 + "foo bar") + "a") {
                                                        case "":
                                                            var v104 float32 = v1
                                                            _ = v104
                                                        }
                                                    }
                                                    _ = fi100
                                                    _ = v101
                                                    _ = v102
                                                    _ = v103
                                                }
                                                _ = v99
                                            }
                                            if true {
                                                for fi105 := 0; fi105 < 10; fi105++ {
                                                    var v106 uint8 = 255
                                                    v1 += (v1 * 1.5)
                                                    _ = fi105
                                                    _ = v106
                                                }
                                            }
                                            var v107 uint = (v18 + v18)
                                            _ = fi83
                                            _ = v107
                                        }
                                    }
                                    v108 := []int{-1}
                                    _ = v108
                                }
                            } else {
                                var v109 string = "a"
                                _ = v109
                            }
                        }
                    default:
                        fmt.Println(v1)
                    }
                case 1:
                    var v110 float32 = v1
                    _ = v110
                }
            }
        }
        for fi111 := 0; fi111 < 20; fi111++ {
            var v112 int32 = -2147483648
            if (((0.0 - 3.14159) - 1.5) >= 0.1) {
                var v113 string = v3
                _ = v113
            } else {
                var v114 float64 = (0.0 + 0.1)
                v112 = (v112 - v112)
                switch v3 {
                case "foo bar":
                    fmt.Println(v3)
                case "\n\t":
                    fmt.Println(v3, v3)
                default:
                    if (true || (9223372036854775807 > 392)) {
                        v18 -= 0
                    }
                }
                _ = v114
            }
            fmt.Println(v1)
            _ = fi111
            _ = v112
        }
    }
    if true {
        v115 := GenS12{F0: 127}
        v116 := GenS13{F0: 1, F1: -1, F2: 42}
        v2 = v2
        _ = v115
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
