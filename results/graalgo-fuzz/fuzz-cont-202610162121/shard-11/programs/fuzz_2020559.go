package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint32 = 282
    var v2 uint64 = 100
    var v3 uint8 = 1
    var v4 int8 = int8(v2)
    var v5 float64 = ((1.5 - (0.1 + (float64(1.5) * (1.5 / 1e10)))) - 0.1)
    v6 := make(map[uint32]int)
    fmt.Println(v5, v4)
    v4 -= v4
    if true {
        v7 := GenS1{F0: "foo bar", F1: 1, F2: v2}
        for fi8 := 0; fi8 < 19; fi8++ {
            if (9223372036854775807 > 9223372036854775807) {
                v9 := GenS2{F0: 1, F1: (v3 + (^(v3)))}
                _ = v9
            } else {
                var v10 uint = 4294967295
                switch int(v10) {
                case 760:
                    var v11 int32 = 1
                    _ = v11
                default:
                    v3 += (v3 * v3)
                }
                if false {
                    v12 := GenS2{F0: v10, F1: 0}
                    _ = v12
                }
                _ = v10
            }
            v13 := GenS3{F0: 795}
            fmt.Println(v4, fi8, v2)
            _ = fi8
            _ = v13
        }
        _ = v7
    } else {
        v6[(v1 % ((2147483648) | 1))] = 32768
        v6[v1] = -504
        switch -927 {
        case 32767:
            var v14 string = "\n\t"
            _ = v14
        case 4294967295:
            for fi15 := 0; fi15 < 9; fi15++ {
                v3 = (v3 & v3)
                _ = fi15
            }
        case 2147483647:
            v16 := GenS3{F0: 1}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := v6[0]
        for fi19 := 0; fi19 < 2; fi19++ {
            switch (fi19 - 32768) {
            case -606:
                if (!true) {
                    v18 = 4294967295
                    if true {
                        if false {
                            var v20 uint32 = v1
                            switch (-(fi19)) {
                            case 32767:
                                if (!false) {
                                    fmt.Println(v4, v4)
                                }
                            case -9223372036854775808:
                                switch (-(v18)) {
                                case 32767:
                                    v1 = (^(v20))
                                case 246:
                                    if false {
                                        v21 := v6[(v1 % (((v20 * 2147483648)) | 1))]
                                        v3 = (v3 >> ((65535) & 63))
                                        var v22 bool = (v5 <= 1e10)
                                        _ = v21
                                        _ = v22
                                    }
                                case 128:
                                    v23 := GenS2{F0: 2, F1: 42}
                                    var v24 uint32 = v1
                                    _ = v23
                                    _ = v24
                                default:
                                    v3 = v3
                                }
                            default:
                                _ = len(v6)
                            }
                            v2 -= (^(v2))
                            _ = v20
                        } else {
                            var v25 bool = true
                            v25 = ((v3 + uint8(v2)) == (v3 & (^(v3))))
                            var v26 bool = (!true)
                            v1 = 2147483647
                            v6[v1] = v18
                            switch "" {
                            case "":
                                fmt.Println(v18, v1)
                            }
                            v27 := []int{}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    } else {
                        v6[v1] = fi17
                        v3 += v3
                        if false {
                            var v28 float64 = v5
                            _ = v28
                        }
                    }
                    switch ("foo bar" + "hello") {
                    case "":
                        var v29 string = "\n\t"
                        _ = v29
                    case "hello":
                        v1 -= v1
                    case "a":
                        v30 := GenS2{F0: 65535, F1: 70}
                        _ = v30
                    }
                }
            case -818:
                var v31 uint8 = (v3 * v3)
                _ = v31
            case 128:
                switch (v18 << ((933) & 63)) {
                case -2147483648:
                    v32 := GenS1{F0: "foo bar", F1: 29, F2: 18446744073709551615}
                    _ = v32
                case -1:
                    var v33 string = "foo bar"
                    _ = v33
                case -514:
                    var v34 string = ("foo bar" + ("" + ""))
                    _ = v34
                default:
                    v4 = v4
                }
                v35 := GenS3{F0: 65535}
                var v36 int8 = int8(v1)
                _ = v35
                _ = v36
            default:
                v37 := []int{fi17}
                _ = v37
            }
            v38 := GenS3{F0: 65535}
            var v39 uint8 = v3
            _ = fi19
            _ = v38
            _ = v39
        }
        var v40 float32 = ((0.1 + (0.1 * -2.25)) / (1e10 - 1.5))
        _ = fi17
        _ = v18
        _ = v40
    }
    fmt.Println(v2, v1, v2)
    v41 := GenS1{F0: "", F1: 633, F2: uint64(v3)}
    var v42 float32 = 0.1
    v4 = int8(v1)
    if (758 <= 93) {
        if true {
            if (true || (true || (((((("foo bar" + "hello") + ("hello" + "")) + "hello") + "hello") + ("\n\t" + "\n\t")) != ""))) {
                v43 := GenS2{F0: 2, F1: (v3 * v3)}
                _ = v43
            } else {
                v44 := []int{-108}
                for fi45 := 0; fi45 < 6; fi45++ {
                    if false {
                        var v46 uint16 = 255
                        _ = v46
                    }
                    v3 += (v3 >> ((261) & 63))
                    switch fi45 {
                    case 1:
                        var v47 int8 = (v4 ^ v4)
                        _ = v47
                    case 32767:
                        var v48 int = fi45
                        _ = v48
                    case -1:
                        var v49 float64 = v5
                        _ = v49
                    default:
                        var v50 uint16 = 32767
                        _ = v50
                    }
                    _ = fi45
                }
                _ = v44
            }
            v51 := GenS3{F0: 880}
            var v52 float64 = v5
            _ = v51
            _ = v52
        }
    } else {
        v2 = 566
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 int64 = 404
            var v55 int = 1
            fmt.Println(v4, v55)
            _ = fi53
            _ = v54
            _ = v55
        }
        v4 += -36
    }
    var v56 int16 = -32768
    if (2147483647 >= 42) {
        v2 = (v2 >> ((0) & 63))
        fmt.Println(v3)
    }
    v57 := GenS3{F0: 42}
    fmt.Println(v3, v56, v42)
    v56 = v56
    if (1e10 <= v5) {
        if (false && (int(v2) > 977)) {
            var v58 uint = 42
            v59 := []int{int(v2)}
            _ = v58
            _ = v59
        } else {
            var v60 string = "foo bar"
            v61 := GenS1{F0: v60, F1: (-(v56)), F2: v2}
            v62 := []int{4294967295, 1, -416, 4294967295}
            _ = v60
            _ = v61
            _ = v62
        }
    }
    if (true && true) {
        if (554 < 2) {
            v5 -= v5
        } else {
            if false {
                var v63 int32 = 0
                if (false || true) {
                    var v64 bool = (v42 == v42)
                    fmt.Println(v56, v1)
                    _ = v64
                }
                _ = v63
            }
            var v65 float32 = v42
            var v66 string = "hello"
            _ = v65
            _ = v66
        }
    } else {
        v67 := GenS1{F0: "a", F1: (-(v56)), F2: v2}
        if true {
            for fi68 := 0; fi68 < 16; fi68++ {
                v56 += (v56 << ((2) & 63))
                _ = fi68
            }
        }
        for fi69 := 0; fi69 < 3; fi69++ {
            if (v5 > (-2.25 + ((-2.25 - v5) / (-2.25 - (v5 - 1.5))))) {
                fmt.Println(fi69, v3)
                v70 := GenS2{F0: 0, F1: v3}
                switch ("a" + "hello") {
                case "foo bar":
                    var v71 uint8 = v3
                    _ = v71
                default:
                    switch "foo bar" {
                    case "foo bar":
                        v72 := GenS3{F0: 4294967295}
                        _ = v72
                    case "":
                        v3 += v3
                    default:
                        var v73 float64 = (float64(v42) + v5)
                        _ = v73
                    }
                }
                _ = v70
            } else {
                if false {
                    if (-9223372036854775808 >= 1) {
                        var v74 float32 = v42
                        _ = v74
                    }
                    v75 := GenS4{F0: 2, F1: v4, F2: v3}
                    for fi76 := 0; fi76 < 1; fi76++ {
                        v42 += v42
                        if false {
                            switch 0 {
                            case 4294967295:
                                var v77 uint32 = uint32(v3)
                                _ = v77
                            case 42:
                                if true {
                                    if (!true) {
                                        var v78 uint32 = (v1 & 42)
                                        var v79 uint64 = (v2 & (v2 - (v2 % (((^(v2))) | 1))))
                                        v42 += ((((v42 - (v42 + v42)) - -2.25) + (v42 - (3.14159 - v42))) * 3.14159)
                                        if false {
                                            v80 := []int{(fi69 / ((127) | 1)), fi76}
                                            _ = v80
                                        } else {
                                            var v81 float32 = v42
                                            v82 := GenS3{F0: 18446744073709551615}
                                            _ = v81
                                            _ = v82
                                        }
                                        _ = v78
                                        _ = v79
                                    } else {
                                        var v83 string = "hello"
                                        _ = v83
                                    }
                                    v84 := GenS2{F0: uint(v3), F1: v3}
                                    fmt.Println(v5, v4, v3)
                                    _ = v84
                                }
                            default:
                                if (v3 != v3) {
                                    fmt.Println(fi69, v42)
                                    if ((^(v4)) != -80) {
                                        v85 := []int{fi76}
                                        v5 = ((1e10 + v5) * (0.1 - v5))
                                        var v86 string = ("hello" + ("hello" + (("hello" + ("hello" + "foo bar")) + "hello")))
                                        _ = v85
                                        _ = v86
                                    } else {
                                        v3 += v3
                                        if false {
                                            v2 -= (v2 % (((v2 ^ (v2 ^ 126))) | 1))
                                            v42 += 0.1
                                        } else {
                                            v87 := GenS5{F0: 1}
                                            v88 := []int{fi76}
                                            _ = v87
                                            _ = v88
                                        }
                                        switch fi69 {
                                        case 2147483647:
                                            var v89 float64 = float64(0.1)
                                            v90 := v6[uint32(v2)]
                                            _ = v89
                                            _ = v90
                                        }
                                    }
                                }
                            }
                            v91 := GenS3{F0: uint(v3)}
                            switch ("" + ("hello" + "a")) {
                            case "a":
                                var v92 uint = 976
                                _ = v92
                            case "\n\t":
                                v93 := GenS6{F0: "hello", F1: 442}
                                _ = v93
                            case "hello":
                                var v94 int32 = 999
                                v1 += v1
                                v95 := []int{2147483647, (fi76 / ((fi69) | 1)), fi69}
                                _ = v94
                                _ = v95
                            }
                            _ = v91
                        }
                        if true {
                            v5 = -2.25
                            switch fi69 {
                            case -9223372036854775808:
                                fmt.Println(v42, v56, v4)
                            case 160:
                                v4 += v4
                            case 32767:
                                var v96 int16 = (v56 * (v56 % (((v56 << ((4294967295) & 63))) | 1)))
                                _ = v96
                            default:
                                switch int(v3) {
                                case 354:
                                    var v97 uint8 = 42
                                    _ = v97
                                case 9223372036854775807:
                                    switch 127 {
                                    case 128:
                                        if false {
                                            var v98 int16 = (v56 & v56)
                                            v1 += 65535
                                            if true {
                                                switch "\n\t" {
                                                case "":
                                                    var v99 int16 = v56
                                                    _ = v99
                                                case "a":
                                                    var v100 int = fi76
                                                    _ = v100
                                                case "foo bar":
                                                    v4 = v4
                                                    v101 := GenS6{F0: ("a" + ""), F1: 255}
                                                    _ = len(v6)
                                                    _ = v101
                                                default:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v4 = 0
                                                    case "\n\t":
                                                        fmt.Println(v3)
                                                    }
                                                }
                                            } else {
                                                var v102 bool = ((v98 & v98) >= v56)
                                                var v103 int = fi69
                                                _ = v102
                                                _ = v103
                                            }
                                            _ = v98
                                        }
                                    case 4294967295:
                                        _ = len(v6)
                                    case -772:
                                        if ((^(fi76)) > fi76) {
                                            v104 := []int{42}
                                            _ = v104
                                        }
                                    }
                                case 1:
                                    if true {
                                        var v105 int16 = 0
                                        fmt.Println(v2, v56)
                                        _ = v105
                                    } else {
                                        if (fi69 >= (fi76 | (fi69 * (-(fi69))))) {
                                            var v106 float64 = (v5 + v5)
                                            v107 := []int{913}
                                            _ = v106
                                            _ = v107
                                        }
                                    }
                                default:
                                    if true {
                                        fmt.Println(v42, v3)
                                    }
                                }
                            }
                        }
                        _ = fi76
                    }
                    _ = v75
                } else {
                    var v108 uint32 = (v1 % ((65535) | 1))
                    var v109 int8 = (v4 & 42)
                    var v110 uint8 = (v3 / ((v3) | 1))
                    _ = v108
                    _ = v109
                    _ = v110
                }
                if false {
                    if (127 <= v3) {
                        if false {
                            fmt.Println(v4)
                        } else {
                            var v111 uint16 = 65535
                            _ = v111
                        }
                        var v112 uint8 = (v3 << ((192) & 63))
                        v42 = v42
                        _ = v112
                    } else {
                        var v113 int64 = 0
                        var v114 uint64 = (v2 * 42)
                        v115 := GenS6{F0: "", F1: 0}
                        _ = v113
                        _ = v114
                        _ = v115
                    }
                    v116 := GenS7{F0: (v3 & 1), F1: v3}
                    v117 := v6[v1]
                    _ = v116
                    _ = v117
                }
                var v118 int64 = 1
                _ = v118
            }
            v119 := v6[v1]
            fmt.Println(v56, fi69, fi69)
            _ = fi69
            _ = v119
        }
        _ = v67
    }
    v42 -= (((v42 + v42) + v42) - v42)
    v120 := GenS6{F0: ("a" + "foo bar"), F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
