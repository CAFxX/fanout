package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int; F1 int64; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 int16 }
func genFn0() string {
    var v7 int = 42
    fmt.Println(v7)
    v8 := GenS1{F0: 18}
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return (("foo bar" + "") + ("hello" + ("" + "")))
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 251
    var v3 uint8 = v2
    var v4 int8 = 0
    v5 := make(map[string]int)
    var v6 int32 = -2147483648
    if ((0.1 > 0.0) && (true || (true && (!(false || (!((-208 >= -1) || (!(((3.14159 + 0.0) - (((-2.25 - ((1e10 / (1.5 - 1e10)) / ((0.0 - (0.1 / (((1.5 - ((-2.25 + ((3.14159 / (1.5 - 1e10)) * (float32(0.0) / 1.0))) * (-2.25 / (0.0 - (-2.25 / 1e10))))) / 3.14159) - 0.0))) + 3.14159))) + (-2.25 - 0.0)) * -2.25)) >= 0.0))))))))) {
        fmt.Println(v6, v1, v2)
        var v10 bool = (uint16(v2) != 944)
        delete(v5, (("foo bar" + "foo bar") + ("foo bar" + "foo bar")))
        _ = v10
    } else {
        fmt.Println(v6, v3)
        if false {
            if (-9223372036854775808 != 42) {
                switch ((("a" + ("foo bar" + "hello")) + "a") + "a") {
                case "foo bar":
                    v11 := []int{32767, -1, 42}
                    _ = v11
                case "":
                    switch ("a" + ("hello" + ("hello" + ""))) {
                    case "foo bar":
                        switch "foo bar" {
                        case "":
                            if true {
                                v1 = 18446744073709551615
                                var v12 float32 = 1.5
                                for fi13 := 0; fi13 < 9; fi13++ {
                                    v14 := []int{127, (fi13 ^ fi13)}
                                    _ = fi13
                                    _ = v14
                                }
                                _ = v12
                            }
                        }
                    case "":
                        var v15 uint = 42
                        _ = v15
                    case "\n\t":
                        switch -2147483648 {
                        case 877:
                            v4 = 41
                        case 4294967295:
                            var v16 int = -39
                            _ = v16
                        case 2147483647:
                            var v17 uint8 = (v3 / (((v3 | 1)) | 1))
                            _ = v17
                        }
                    }
                }
                var v18 uint16 = 42
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 uint8 = (v3 | 115)
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                v6 -= 2147483647
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS1{F0: (^(v4))}
                _ = fi21
                _ = v22
            }
            v23 := GenS1{F0: 127}
            _ = v23
        }
    }
    fmt.Println(v1)
    var v24 int32 = 2147483647
    if (0.1 != (0.1 * (-2.25 - (-2.25 * ((float32((((1e10 / 1.5) - (float64((0.0 + 1e10)) * float64((1.5 / 1.0)))) - (-2.25 * float64((3.14159 * ((-2.25 / 1e10) / 1.0)))))) - -2.25) - 1e10))))) {
        switch "a" {
        case "foo bar":
            if ((false || (v6 <= 1)) || true) {
                v25 := []int{331}
                if false {
                    if true {
                        var v26 string = ("hello" + "\n\t")
                        var v27 int16 = 472
                        _ = v26
                        _ = v27
                    }
                    switch 9223372036854775807 {
                    case -1:
                        v24 += (v24 >> ((2) & 63))
                    case 1:
                        for fi28 := 0; fi28 < 1; fi28++ {
                            _ = len(v5)
                            if (224 <= 32767) {
                                if ((485 < 255) && false) {
                                    for fi29 := 0; fi29 < 11; fi29++ {
                                        v2 -= v2
                                        v30 := GenS2{F0: (v6 | (^(v6)))}
                                        var v31 float64 = (0.1 + 0.0)
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                    v4 = (v4 - (v4 | v4))
                                }
                                fmt.Println(v3)
                                fmt.Println(v3, v3)
                            }
                            var v32 string = ("a" + (("hello" + "hello") + "hello"))
                            _ = fi28
                            _ = v32
                        }
                    default:
                        switch "\n\t" {
                        case "":
                            v33 := GenS1{F0: (v4 - (v4 & v4))}
                            _ = v33
                        case "\n\t":
                            var v34 bool = (1e10 == float32((((0.0 + 1.5) * 0.0) + (0.1 + 0.0))))
                            _ = v34
                        case "a":
                            var v35 int8 = v4
                            _ = v35
                        }
                    }
                } else {
                    var v36 float64 = ((1e10 / ((1.5 / 1.0) / 0.1)) / ((((-2.25 + 3.14159) / float64((1.5 * 1.5))) + 0.0) / 1.0))
                    v37 := v5[(("a" + "hello") + "hello")]
                    switch (v37 % ((v37) | 1)) {
                    case 832:
                        for fi38 := 0; fi38 < 19; fi38++ {
                            for fi39 := 0; fi39 < 2; fi39++ {
                                v3 -= (v3 % (((v3 >> ((2) & 63))) | 1))
                                v40 := GenS2{F0: -2147483648}
                                _ = fi39
                                _ = v40
                            }
                            var v41 int16 = 32767
                            _ = fi38
                            _ = v41
                        }
                    }
                    _ = v36
                    _ = v37
                }
                if ((226 <= 255) || (false && false)) {
                    if true {
                        if true {
                            var v42 float64 = -2.25
                            _ = v42
                        }
                        v4 = v4
                        if ((1e10 + -2.25) <= ((((((float32((((3.14159 * 1e10) - 1e10) / (3.14159 / 1.0))) * 3.14159) / (3.14159 / 0.1)) * 0.1) - (1.5 / float32((1e10 + 1.5)))) + 1e10) + 0.0)) {
                            v1 -= (v1 / (((v1 * v1)) | 1))
                            v3 = v3
                        }
                    } else {
                        var v43 uint32 = 85
                        v44 := GenS3{F0: 42}
                        var v45 uint64 = v1
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    var v46 float64 = -2.25
                    _ = v46
                } else {
                    _ = len(v25)
                    v6 = 2147483647
                }
                _ = v25
            } else {
                var v47 bool = true
                for fi48 := 0; fi48 < 6; fi48++ {
                    v4 = 127
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 float32 = 0.0
                        if (-1 != 1000) {
                            if ((v3 << ((4294967295) & 63)) != 255) {
                                var v51 int32 = (^(v6))
                                v52 := []int{}
                                if true {
                                    var v53 int64 = -125
                                    _ = v53
                                }
                                _ = v51
                                _ = v52
                            } else {
                                v4 += int8(v6)
                            }
                            v50 = float32(1e10)
                        }
                        if v47 {
                            fmt.Println(fi49, v47)
                        } else {
                            var v54 uint32 = uint32(v3)
                            v55 := []int{fi49, (fi49 * 79), fi49}
                            v24 = 0
                            _ = v54
                            _ = v55
                        }
                        _ = fi49
                        _ = v50
                    }
                    _ = fi48
                }
                switch -9223372036854775808 {
                case 127:
                    switch int(v6) {
                    case -1:
                        if false {
                            if (-9223372036854775808 == 967) {
                                switch 1 {
                                case -9223372036854775808:
                                    v6 -= (v6 % (((v24 | (v24 - 2147483647))) | 1))
                                case 42:
                                    var v56 uint16 = 2
                                    _ = v56
                                default:
                                    v57 := GenS4{F0: v1, F1: 65535}
                                    _ = v57
                                }
                            }
                            var v58 int8 = v4
                            _ = v58
                        } else {
                            for fi59 := 0; fi59 < 10; fi59++ {
                                var v60 int16 = 941
                                _ = fi59
                                _ = v60
                            }
                            if (2 != 42) {
                                var v61 uint16 = 65535
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    if (!((!(v47 || (!false))) && ((v24 ^ int32(v4)) == v6))) {
                                        var v63 float32 = (-2.25 / 3.14159)
                                        _ = v63
                                    }
                                    _ = fi62
                                }
                                if (!(v47 && (v47 || v47))) {
                                    if v47 {
                                        if v47 {
                                            v24 = v24
                                            v64 := GenS5{F0: int(v2), F1: -9223372036854775808, F2: "hello"}
                                            _ = v64
                                        }
                                        switch ("a" + "\n\t") {
                                        case "hello":
                                            for fi65 := 0; fi65 < 20; fi65++ {
                                                v1 = (^(v1))
                                                var v66 int64 = 225
                                                var v67 uint32 = 589
                                                _ = fi65
                                                _ = v66
                                                _ = v67
                                            }
                                        case "":
                                            var v68 uint16 = (v61 & uint16(v1))
                                            _ = v68
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "hello":
                                                var v69 float64 = 0.1
                                                _ = v69
                                            }
                                        }
                                        v61 = uint16(v3)
                                    }
                                    v4 = (v4 / (((v4 >> ((616) & 63))) | 1))
                                    var v70 uint8 = 255
                                    _ = v70
                                } else {
                                    if v47 {
                                        var v71 int64 = 42
                                        v5[(("a" + "a") + "a")] = -964
                                        _ = v71
                                    } else {
                                        var v72 uint = 53
                                        _ = v72
                                    }
                                    fmt.Println(v61, v1)
                                }
                                _ = v61
                            } else {
                                var v73 int8 = 42
                                fmt.Println(v3)
                                v74 := GenS2{F0: -1}
                                _ = v73
                                _ = v74
                            }
                            fmt.Println(v2, v4, v6)
                        }
                    }
                case 32767:
                    switch "foo bar" {
                    case "a":
                        v75 := GenS6{F0: (v24 - 1)}
                        _ = v75
                    case "":
                        switch "a" {
                        case "\n\t":
                            v76 := GenS7{F0: -1, F1: v47, F2: v1}
                            _ = v76
                        case "a":
                            var v77 int32 = (v24 ^ int32(v4))
                            _ = v77
                        default:
                            v78 := GenS2{F0: 0}
                            _ = v78
                        }
                    case "foo bar":
                        var v79 float32 = float32(0.0)
                        _ = v79
                    default:
                        if (42 <= 32767) {
                            for fi80 := 0; fi80 < 7; fi80++ {
                                for fi81 := 0; fi81 < 17; fi81++ {
                                    var v82 int64 = -9223372036854775808
                                    fmt.Println(fi80)
                                    v83 := []int{42, (-(fi80)), fi81, int(v6)}
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                }
                                _ = fi80
                            }
                            for fi84 := 0; fi84 < 3; fi84++ {
                                v24 = 2147483647
                                _ = fi84
                            }
                        }
                    }
                }
                _ = v47
            }
        default:
            fmt.Println(v2, v2)
        }
        var v85 uint = 271
        for fi86 := 0; fi86 < 8; fi86++ {
            var v87 float32 = 0.0
            for fi88 := 0; fi88 < 13; fi88++ {
                var v89 uint16 = 1
                _ = fi88
                _ = v89
            }
            var v90 uint = 153
            _ = fi86
            _ = v87
            _ = v90
        }
        _ = v85
    }
    for fi91 := 0; fi91 < 9; fi91++ {
        v92 := GenS7{F0: 1, F1: true, F2: 2}
        _ = fi91
        _ = v92
    }
    v2 += v2
    for fi93 := 0; fi93 < 16; fi93++ {
        if (uint64(v3) >= v1) {
            var v94 uint64 = v1
            _ = v94
        } else {
            switch ("a" + ("a" + "")) {
            case "hello":
                switch ((("" + "a") + (("foo bar" + ("\n\t" + "\n\t")) + "hello")) + ("a" + "hello")) {
                case "\n\t":
                    var v95 int = (fi93 + fi93)
                    _ = v95
                case "":
                    var v96 uint = 255
                    _ = v96
                default:
                    switch (-(fi93)) {
                    case 4294967295:
                        v97 := GenS3{F0: fi93}
                        _ = v97
                    case 1:
                        switch fi93 {
                        case 42:
                            var v98 int64 = 9223372036854775807
                            _ = v98
                        case 127:
                            var v99 uint32 = 4294967295
                            _ = v99
                        }
                    }
                }
            case "foo bar":
                v5[("foo bar" + "a")] = fi93
            case "":
                for fi100 := 0; fi100 < 4; fi100++ {
                    v101 := GenS8{F0: (^(v1)), F1: 0}
                    var v102 uint64 = v1
                    fmt.Println(v102, v102)
                    v6 = v6
                    if (!(346 != 319)) {
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                if (4294967295 == 2147483648) {
                                    var v103 uint64 = v1
                                    _ = v103
                                }
                                if false {
                                    if (false || false) {
                                        v5[("a" + ("a" + ("foo bar" + "foo bar")))] = (fi100 << ((415) & 63))
                                        var v104 float64 = 1.5
                                        fmt.Println(v102)
                                        v105 := GenS7{F0: 1, F1: (32767 == 32767), F2: (^(v102))}
                                        _ = v104
                                        _ = v105
                                    } else {
                                        var v106 int16 = 0
                                        fmt.Println(v3)
                                        var v107 int8 = 27
                                        _ = v106
                                        _ = v107
                                    }
                                    fmt.Println(v1, v4, v102)
                                    var v108 bool = (1000 != -1)
                                    _ = v108
                                } else {
                                    v24 = v6
                                    var v109 float64 = float64(((float32(0.0) + (1.5 * 1e10)) / 1e10))
                                    var v110 int = fi93
                                    _ = v109
                                    _ = v110
                                }
                            }
                        case "\n\t":
                            var v111 float64 = 0.1
                            _ = v111
                        }
                        if (246 < (v3 << ((4294967295) & 63))) {
                            v112 := v5[""]
                            v113 := GenS3{F0: 42}
                            _ = v112
                            _ = v113
                        } else {
                            if (2 == 0) {
                                v114 := []int{32768}
                                _ = v114
                            }
                            v115 := GenS9{F0: -404}
                            _ = v115
                        }
                    }
                    fmt.Println(v2)
                    v116 := GenS8{F0: v1, F1: 1}
                    v117 := GenS10{F0: v102, F1: int16(fi100)}
                    _ = fi100
                    _ = v101
                    _ = v102
                    _ = v116
                    _ = v117
                }
            default:
                if true {
                    fmt.Println(v4, fi93)
                    var v118 uint = 550
                    v119 := GenS7{F0: 32767, F1: false, F2: 721}
                    _ = v118
                    _ = v119
                }
            }
            if (fi93 <= (^(fi93))) {
                v120 := []int{(fi93 >> ((18446744073709551615) & 63)), fi93, fi93}
                fmt.Println(v4, v1)
                if false {
                    fmt.Println(fi93, v1, v4)
                    v6 += 2147483647
                }
                _ = v120
            } else {
                var v121 bool = true
                _ = v121
            }
            fmt.Println(fi93, v3, v2)
        }
        for fi122 := 0; fi122 < 9; fi122++ {
            fmt.Println(v3, v1, v4)
            _ = fi122
        }
        var v123 uint = 0
        _ = fi93
        _ = v123
    }
    var v124 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v124
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
