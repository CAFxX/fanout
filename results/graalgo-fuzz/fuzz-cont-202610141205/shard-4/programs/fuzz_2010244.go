package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int16; F1 string; F2 bool }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8; F1 uint16; F2 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16; F1 int8 }
type GenS12 struct { F0 int8; F1 int64; F2 uint8 }
type GenS13 struct { F0 uint16; F1 uint8; F2 string }
type GenS14 struct { F0 int32; F1 uint16; F2 bool }
type GenS15 struct { F0 uint; F1 uint8 }
type GenS16 struct { F0 bool }

func main() {
    var v1 uint = 255
    v2 := make(map[int]string)
    var v3 int = 42
    var v4 uint8 = 42
    var v5 float64 = 1.5
    var v6 int = v3
    switch (v6 * -9223372036854775808) {
    case -9223372036854775808:
        switch ("\n\t" + "") {
        case "a":
            var v7 uint16 = 255
            _ = v7
        case "hello":
            v4 = 42
        default:
            v8 := GenS1{F0: -32768}
            _ = v8
        }
    case 2147483647:
        v5 = 1e10
    }
    switch 128 {
    case 32768:
        switch (v6 << (((v1 * v1)) & 63)) {
        case 32768:
            if (0 <= 1) {
                if (876 < 2) {
                    var v9 int16 = -32768
                    _ = v9
                }
                v10 := GenS1{F0: 654}
                if (-9223372036854775808 > 2147483647) {
                    var v11 string = "\n\t"
                    _ = v11
                }
                _ = v10
            } else {
                v12 := GenS2{F0: (v1 << ((v1) & 63)), F1: 648}
                v13 := GenS2{F0: v1, F1: 231}
                _ = v12
                _ = v13
            }
        case 322:
            v14 := GenS3{F0: (v4 - (v4 | 44)), F1: 255, F2: -428}
            _ = v14
        case 842:
            var v15 float64 = (v5 * v5)
            _ = v15
        default:
            switch "hello" {
            case "hello":
                switch v6 {
                case -2147483648:
                    v16 := GenS4{F0: 606}
                    _ = v16
                case 0:
                    fmt.Println(v5, v4, v4)
                default:
                    for fi17 := 0; fi17 < 1; fi17++ {
                        var v18 string = (("" + "hello") + "\n\t")
                        _ = fi17
                        _ = v18
                    }
                }
            case "\n\t":
                if (127 != v3) {
                    if (42 < -2147483648) {
                        v3 = (v6 + v6)
                        v19 := []int{(v3 + (v6 + (v6 & (^(v3))))), (v3 * v6)}
                        _ = v19
                    } else {
                        if (128 > v4) {
                            if true {
                                v20 := GenS5{F0: 4294967295, F1: 100}
                                switch "" {
                                case "\n\t":
                                    var v21 int16 = -32768
                                    _ = v21
                                default:
                                    if (9223372036854775807 == 0) {
                                        if false {
                                            var v22 uint32 = 42
                                            _ = v22
                                        }
                                        v23 := []int{}
                                        switch "foo bar" {
                                        case "\n\t":
                                            v24 := GenS6{F0: -32768, F1: "", F2: false}
                                            _ = v24
                                        }
                                        _ = v23
                                    } else {
                                        switch (v6 >> ((v1) & 63)) {
                                        case 128:
                                            var v25 uint8 = (^(v4))
                                            _ = v25
                                        case 42:
                                            var v26 int16 = 32767
                                            _ = v26
                                        default:
                                            v27 := []int{v3, v6, v3}
                                            _ = v27
                                        }
                                    }
                                }
                                v3 -= v6
                                _ = v20
                            }
                            if (((v5 + (0.1 * 3.14159)) / v5) >= v5) {
                                for fi28 := 0; fi28 < 4; fi28++ {
                                    var v29 int64 = 0
                                    v29 -= v29
                                    if (!false) {
                                        switch fi28 {
                                        case 0:
                                            v1 = v1
                                        case -2147483648:
                                            v3 += v3
                                        }
                                        switch fi28 {
                                        case -2147483648:
                                            for fi30 := 0; fi30 < 4; fi30++ {
                                                var v31 uint64 = 797
                                                var v32 int64 = (v29 >> ((v1) & 63))
                                                v33 := GenS7{F0: 989, F1: (v6 << ((42) & 63))}
                                                var v34 string = ("foo bar" + ("" + "a"))
                                                v35 := []int{(^(fi28)), v6, v3, 9223372036854775807}
                                                _ = fi30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                        case 969:
                                            fmt.Println(fi28, v6, v29)
                                        case 178:
                                            if false {
                                                var v36 int8 = 127
                                                if (v1 < (v1 / ((v1) | 1))) {
                                                    switch (((("hello" + ("" + "")) + "foo bar") + "hello") + "\n\t") {
                                                    case "foo bar":
                                                        for fi37 := 0; fi37 < 17; fi37++ {
                                                            var v38 uint8 = v4
                                                            _ = fi37
                                                            _ = v38
                                                        }
                                                    }
                                                    if false {
                                                        v39 := GenS8{F0: 0}
                                                        v36 = 40
                                                        fmt.Println(v1)
                                                        _ = v39
                                                    }
                                                    v2[(v3 % ((0) | 1))] = ("hello" + "a")
                                                }
                                                _ = v36
                                            } else {
                                                fmt.Println(v6, fi28)
                                                switch (v6 | (fi28 - fi28)) {
                                                case 691:
                                                    if (v5 <= float64(0.0)) {
                                                        if ((v29 % ((v29) | 1)) >= (v29 / ((9223372036854775807) | 1))) {
                                                            v40 := GenS5{F0: 2, F1: 1}
                                                            v6 += (fi28 + v6)
                                                            v2[v3] = "foo bar"
                                                            _ = v40
                                                        }
                                                        v41 := []int{}
                                                        _ = v41
                                                    }
                                                case -53:
                                                    if true {
                                                        var v42 int8 = -128
                                                        switch "foo bar" {
                                                        case "":
                                                            for fi43 := 0; fi43 < 17; fi43++ {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v44 uint16 = 65535
                                                                    v4 -= (v4 * v4)
                                                                    _ = v44
                                                                default:
                                                                    v45 := []int{(fi43 << (((^(v1))) & 63))}
                                                                    _ = v45
                                                                }
                                                                v46 := GenS9{F0: 50, F1: 255, F2: 1}
                                                                _ = fi43
                                                                _ = v46
                                                            }
                                                        case "\n\t":
                                                            v47 := v2[(^(fi28))]
                                                            _ = v47
                                                        case "a":
                                                            if true {
                                                                for fi48 := 0; fi48 < 13; fi48++ {
                                                                    var v49 int16 = -962
                                                                    v50 := GenS10{F0: ((!false) || ((18446744073709551615 <= 418) || true))}
                                                                    _ = fi48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                fmt.Println(v5)
                                                                v51 := []int{127, v6, v3, v6}
                                                                _ = v51
                                                            } else {
                                                                for fi52 := 0; fi52 < 1; fi52++ {
                                                                    if (("" + "a") != (("\n\t" + "\n\t") + "foo bar")) {
                                                                        v53 := GenS11{F0: 65535, F1: 100}
                                                                        _ = v53
                                                                    } else {
                                                                        fmt.Println(v29)
                                                                        var v54 float32 = 1e10
                                                                        switch (("" + "\n\t") + "a") {
                                                                        case "":
                                                                            v5 = (1e10 - v5)
                                                                            var v55 string = ("\n\t" + ("" + ("\n\t" + (("" + "\n\t") + "foo bar"))))
                                                                            v54 += -2.25
                                                                            if (790 >= 255) {
                                                                                var v56 string = v55
                                                                                if (0 >= 1) {
                                                                                    var v57 int8 = (v42 + 0)
                                                                                    var v58 string = "a"
                                                                                    v55 = (v55 + v58)
                                                                                    if false {
                                                                                        switch ("foo bar" + ("hello" + "hello")) {
                                                                                        case "\n\t":
                                                                                            v59 := GenS6{F0: 1, F1: "a", F2: (141 != v4)}
                                                                                            _ = v59
                                                                                        case "":
                                                                                            switch (-(fi28)) {
                                                                                            case -1:
                                                                                                v42 = v42
                                                                                            }
                                                                                        case "hello":
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v6)
                                                                                            case "":
                                                                                                fmt.Println(fi28, v5, v1)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (1 <= v4) {
                                                                                        v3 = v6
                                                                                    } else {
                                                                                        v29 += (v29 & v29)
                                                                                    }
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                if (int16(v4) == -1) {
                                                                                    if true {
                                                                                        v60 := GenS12{F0: v42, F1: -9223372036854775808, F2: (^(v4))}
                                                                                        switch v56 {
                                                                                        case "foo bar":
                                                                                            var v61 uint8 = (v4 * v4)
                                                                                            _ = v61
                                                                                        case "":
                                                                                            var v62 int64 = -372
                                                                                            _ = v62
                                                                                        default:
                                                                                            fmt.Println(v6)
                                                                                        }
                                                                                        _ = v60
                                                                                    }
                                                                                    v1 = v1
                                                                                    v63 := GenS13{F0: 1, F1: v4, F2: v55}
                                                                                    _ = v63
                                                                                }
                                                                                _ = v56
                                                                            } else {
                                                                                if false {
                                                                                    v5 += v5
                                                                                }
                                                                            }
                                                                            v64 := []int{(fi52 ^ fi28), (fi52 >> ((v1) & 63))}
                                                                            fmt.Println(v29, v6)
                                                                            var v65 uint32 = 2147483648
                                                                            _ = v55
                                                                            _ = v64
                                                                            _ = v65
                                                                        case "a":
                                                                            if true {
                                                                                if (!(uint32(v4) > 2147483648)) {
                                                                                    if true {
                                                                                        switch fi28 {
                                                                                        case 2147483647:
                                                                                            v66 := GenS14{F0: 1, F1: 1, F2: (42 <= 2)}
                                                                                            _ = v66
                                                                                        case 127:
                                                                                            if false {
                                                                                                v1 = 2
                                                                                            }
                                                                                        case -2147483648:
                                                                                            var v67 int = v6
                                                                                            _ = v67
                                                                                        }
                                                                                    }
                                                                                    v68 := []int{(v6 ^ -1), int(v4), v6, fi28}
                                                                                    _ = v68
                                                                                } else {
                                                                                    v69 := GenS9{F0: 62, F1: 0, F2: (v3 * (-(fi28)))}
                                                                                    var v70 int8 = v42
                                                                                    v3 += fi52
                                                                                    var v71 float64 = v5
                                                                                    var v72 int16 = -32768
                                                                                    var v73 int64 = int64(v1)
                                                                                    var v74 float64 = v5
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v54
                                                                    }
                                                                    var v75 int16 = int16(v4)
                                                                    _ = fi52
                                                                    _ = v75
                                                                }
                                                            }
                                                        default:
                                                            var v76 int64 = v29
                                                            _ = v76
                                                        }
                                                        var v77 int64 = (v29 - int64(v1))
                                                        _ = v42
                                                        _ = v77
                                                    }
                                                case 0:
                                                    v4 -= v4
                                                default:
                                                    for fi78 := 0; fi78 < 14; fi78++ {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v79 int16 = 1000
                                                            _ = v79
                                                        default:
                                                            v6 = fi78
                                                            fmt.Println(v6)
                                                        }
                                                        _ = fi78
                                                    }
                                                }
                                                var v80 float64 = v5
                                                _ = v80
                                            }
                                        }
                                        v5 -= (v5 * (v5 / (((3.14159 * (1.5 + (v5 + (((((-2.25 * 0.1) / ((float64(float32(-2.25)) + (v5 / ((1e10 * v5) - (v5 * -2.25)))) / (v5 + 3.14159))) + v5) * (v5 * (v5 / 0.1))) / 0.1)))) + (v5 * v5)) + 0.0)))
                                    } else {
                                        for fi81 := 0; fi81 < 15; fi81++ {
                                            var v82 float64 = (v5 + (float64(0.1) / v5))
                                            _ = fi81
                                            _ = v82
                                        }
                                        v29 = v29
                                        for fi83 := 0; fi83 < 13; fi83++ {
                                            v5 += (-2.25 / ((3.14159 + ((0.0 - (float64(((3.14159 * float32(((float64(((float32(1e10) - ((-2.25 / (0.1 / float32(v5))) * 0.0)) / 1.0)) + v5) * ((v5 * float64(((float32(-2.25) + 0.1) - (0.0 * (float32(0.1) + 3.14159))))) * float64(float32((v5 * 3.14159))))))) + 0.1)) - ((v5 / 3.14159) + v5))) * (((0.0 + float64((0.1 + 0.0))) * (1e10 / v5)) * ((v5 - 0.0) + v5)))) * ((1.5 - (0.1 - -2.25)) + -2.25)))
                                            var v84 uint = 255
                                            if ((v5 > ((float64(1.5) + ((v5 * ((v5 / v5) - v5)) - 1e10)) - ((1e10 - v5) - v5))) || (-1 != -1)) {
                                                v1 = v84
                                                v6 -= 0
                                            } else {
                                                if false {
                                                    var v85 uint16 = 1
                                                    var v86 int = 317
                                                    v87 := GenS15{F0: v1, F1: v4}
                                                    var v88 int8 = 100
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                    _ = v88
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    v89 := GenS13{F0: 42, F1: v4, F2: ""}
                                                    _ = v89
                                                default:
                                                    v29 = 776
                                                }
                                                switch 741 {
                                                case 32768:
                                                    v90 := GenS9{F0: 42, F1: 575, F2: -909}
                                                    _ = v90
                                                case 32767:
                                                    v91 := GenS16{F0: (32767 > 2)}
                                                    _ = v91
                                                case 4294967295:
                                                    v4 -= uint8(v84)
                                                }
                                            }
                                            _ = fi83
                                            _ = v84
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                var v92 bool = true
                                if (42 <= 887) {
                                    v93 := []int{}
                                    v6 += v3
                                    _ = v93
                                } else {
                                    var v94 bool = (587 == 9223372036854775807)
                                    var v95 bool = ((!true) || v94)
                                    v96 := []int{(v6 ^ v6), 1, v6}
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                                _ = v92
                            } else {
                                var v97 string = (("\n\t" + "hello") + ("hello" + "\n\t"))
                                v5 += (v5 - ((float64(-2.25) * 0.1) + v5))
                                _ = v97
                            }
                            v5 = 3.14159
                        }
                        v1 -= (v1 - 255)
                        v6 = v6
                    }
                } else {
                    v98 := GenS13{F0: 255, F1: v4, F2: ("hello" + "hello")}
                    v99 := []int{(v3 | (^(v3))), (v6 * 127), v6}
                    _ = v98
                    _ = v99
                }
            case "foo bar":
                var v100 float64 = 0.1
                _ = v100
            default:
                var v101 int32 = 534
                _ = v101
            }
        }
    }
    var v102 string = ("" + ((("\n\t" + "") + "hello") + ("a" + "\n\t")))
    for fi103 := 0; fi103 < 10; fi103++ {
        v5 += v5
        var v104 bool = (v1 == (v1 / (((v1 << (((v1 % ((4294967295) | 1))) & 63))) | 1)))
        v104 = false
        _ = fi103
        _ = v104
    }
    for fi105 := 0; fi105 < 7; fi105++ {
        var v106 uint16 = 0
        _ = fi105
        _ = v106
    }
    v5 += v5
    for fi107 := 0; fi107 < 2; fi107++ {
        var v108 uint = 1
        v109 := []int{4294967295}
        v110 := GenS14{F0: 2147483647, F1: 65535, F2: (!false)}
        _ = fi107
        _ = v108
        _ = v109
        _ = v110
    }
    if ((-2.25 == 0.1) || false) {
        v6 += v6
        var v111 string = v102
        _ = v111
    }
    if false {
        v4 -= (v4 | v4)
    }
    v112 := GenS5{F0: 262, F1: 127}
    switch v3 {
    case -2147483648:
        if (!((v102 + v102) == "foo bar")) {
            var v113 uint16 = 369
            var v114 float32 = (1.5 * 3.14159)
            _ = v113
            _ = v114
        } else {
            v5 = float64(0.1)
            v5 += v5
        }
    case 1:
        delete(v2, (v3 / (((v3 / (((v3 >> (((v1 << ((1) & 63))) & 63))) | 1))) | 1)))
    default:
        v5 -= (v5 * v5)
    }
    v115 := []int{32767, int(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v115:", len(v115), v115)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v112
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
