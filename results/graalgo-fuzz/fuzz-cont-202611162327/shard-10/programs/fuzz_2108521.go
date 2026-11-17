package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16 }
func genFn0() string {
    v6 := []int{1, 186, -1, 186}
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7)
        switch ("foo bar" + "a") {
        case "\n\t":
            for fi8 := 0; fi8 < 6; fi8++ {
                fmt.Println(fi8, fi8)
                switch "a" {
                case "":
                    var v9 uint16 = 42
                    _ = v9
                case "hello":
                    var v10 float32 = 0.1
                    _ = v10
                default:
                    fmt.Println(fi8)
                }
                _ = fi8
            }
        case "hello":
            v11 := GenS1{F0: false}
            _ = v11
        case "":
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := GenS2{F0: 65535}
                v6 = append(v6, -492)
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
    }
    _ = v6
    return "a"
}
func genFn1(p0 int8, p1 int8) int8 {
    p0 = p0
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            var v15 float64 = 3.14159
            switch (fi14 + fi14) {
            case 42:
                p0 = 1
            case 32768:
                v16 := []int{127}
                _ = v16
            case -107:
                if (!(true || true)) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v18 := GenS3{F0: 42}
                        _ = fi17
                        _ = v18
                    }
                    v19 := []int{-786, fi14, (fi14 | 128), fi14}
                    if true {
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v21 := GenS4{F0: 513}
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = v19
                }
            }
            if (false || false) {
                p0 = 100
            }
            _ = v15
        }
        v22 := []int{(fi14 + (fi14 / (((fi14 | 4294967295)) | 1))), fi14, fi14}
        _ = fi14
        _ = v22
    }
    var v23 int64 = 707
    _ = p0
    _ = p1
    _ = v23
    return p0
}

func main() {
    var v1 uint64 = 1
    var v2 uint64 = (v1 >> ((4294967295) & 63))
    v3 := []uint32{65535, 663, 0, 4294967295}
    v4 := make(map[uint32]string)
    var v5 bool = ((true || (v1 <= 4294967295)) && true)
    var v24 int16 = 1000
    var v25 float32 = 0.0
    switch 915 {
    case -439:
        if (4294967295 <= 760) {
            v4[386] = ("foo bar" + "a")
            for fi26 := 0; fi26 < 14; fi26++ {
                v25 = v25
                _ = fi26
            }
        } else {
            var v27 int = -9223372036854775808
            var v28 float32 = v25
            var v29 int32 = -1
            _ = v27
            _ = v28
            _ = v29
        }
    case 32767:
        switch 706 {
        case 9223372036854775807:
            if v5 {
                _ = len(v3)
                for fi30 := 0; fi30 < 6; fi30++ {
                    v31 := GenS4{F0: 0}
                    fmt.Println(v5)
                    fmt.Println(fi30, v2, v2)
                    _ = fi30
                    _ = v31
                }
            }
        case 467:
            if ((2147483648 >= 42) || (65535 != 1)) {
                for fi32 := 0; fi32 < 9; fi32++ {
                    delete(v4, 2147483648)
                    v33 := GenS4{F0: 1}
                    _ = fi32
                    _ = v33
                }
            }
        default:
            var v34 uint64 = v2
            _ = v34
        }
    case -87:
        v35 := GenS5{F0: 2147483647}
        _ = v35
    }
    if v5 {
        switch -193 {
        case 2147483647:
            if v5 {
                v24 = v24
                var v36 int16 = (v24 >> ((383) & 63))
                v5 = ("\n\t" == "a")
                _ = v36
            }
        case 127:
            if v5 {
                var v37 float32 = ((0.0 + v25) / 0.1)
                v5 = true
                for fi38 := 0; fi38 < 11; fi38++ {
                    if v5 {
                        var v39 bool = v5
                        _ = v39
                    } else {
                        switch 711 {
                        case 1:
                            var v40 uint = 18446744073709551615
                            _ = v40
                        default:
                            v5 = (0 > (v24 | (v24 >> ((42) & 63))))
                        }
                    }
                    v41 := GenS1{F0: false}
                    var v42 int64 = 42
                    _ = fi38
                    _ = v41
                    _ = v42
                }
                _ = v37
            } else {
                if (442 == 48) {
                    if ((v5 && ("" != ("\n\t" + "foo bar"))) && ((v24 ^ v24) == v24)) {
                        v25 -= v25
                        v43 := v4[2147483647]
                        _ = v43
                    }
                    if false {
                        if true {
                            var v44 int8 = -3
                            fmt.Println(v25, v24, v24)
                            var v45 uint64 = 4294967295
                            _ = v44
                            _ = v45
                        } else {
                            v5 = (32767 == (^(v24)))
                            var v46 int = 128
                            _ = v46
                        }
                        v2 = (v2 ^ (v1 | 0))
                        v5 = false
                    } else {
                        if (((("" + ("hello" + "hello")) + "") + "hello") != ("a" + "")) {
                            v47 := v4[2147483648]
                            var v48 int16 = -1
                            var v49 uint = 42
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            if v5 {
                                var v50 int = 1
                                var v51 int64 = 579
                                _ = v50
                                _ = v51
                            }
                        }
                        _ = len(v3)
                        v52 := GenS6{F0: 0}
                        _ = v52
                    }
                    _ = len(v3)
                } else {
                    var v53 uint8 = 119
                    var v54 bool = (-128 != 90)
                    switch "a" {
                    case "hello":
                        v5 = v5
                    case "\n\t":
                        switch 128 {
                        case 127:
                            for fi55 := 0; fi55 < 1; fi55++ {
                                var v56 int = fi55
                                _ = fi55
                                _ = v56
                            }
                        }
                    }
                    _ = v53
                    _ = v54
                }
            }
        }
        v24 -= v24
        v24 = (v24 - (v24 | v24))
    } else {
        v24 += v24
        for fi57 := 0; fi57 < 13; fi57++ {
            var v58 int = fi57
            _ = fi57
            _ = v58
        }
    }
    v1 = v1
    for fi59 := 0; fi59 < 3; fi59++ {
        if v5 {
            switch "hello" {
            case "hello":
                switch "\n\t" {
                case "foo bar":
                    for fi60 := 0; fi60 < 13; fi60++ {
                        switch (fi59 ^ fi59) {
                        case -2147483648:
                            if v5 {
                                v61 := GenS7{F0: 582}
                                var v62 float32 = (((v25 - (((v25 * (1.5 / 0.1)) / v25) - float32((3.14159 + (0.0 * 0.1))))) / v25) - float32((0.0 + float64(0.0))))
                                v3 = append(v3, 907)
                                _ = v61
                                _ = v62
                            }
                        case 967:
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                v63 := GenS8{F0: (fi59 | fi59), F1: 40, F2: 37}
                                _ = v63
                            case "hello":
                                v3 = append(v3, 915)
                            default:
                                delete(v4, 650)
                            }
                        case 1:
                            if (v5 && ((v24 - (v24 >> ((825) & 63))) > 0)) {
                                v3 = append(v3, 814)
                            }
                        }
                        _ = fi60
                    }
                case "\n\t":
                    var v64 float32 = v25
                    _ = v64
                default:
                    var v65 uint32 = 42
                    _ = v65
                }
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                fmt.Println(v5, v5, v24)
            }
        } else {
            v1 = v2
            var v66 int32 = 2147483647
            v4[4294967295] = ("" + "a")
            _ = v66
        }
        if v5 {
            if v5 {
                var v67 uint8 = 127
                _ = v67
            }
            switch "hello" {
            case "\n\t":
                v68 := GenS6{F0: 257}
                _ = v68
            case "a":
                switch ("foo bar" + ("" + ((("" + "a") + ("" + (("" + "foo bar") + "a"))) + "foo bar"))) {
                case "hello":
                    switch ((("a" + "\n\t") + "") + "a") {
                    case "a":
                        var v69 int64 = -9223372036854775808
                        _ = v69
                    case "\n\t":
                        v5 = (v5 || (false && ((fi59 ^ (fi59 / ((int(v24)) | 1))) <= (fi59 + fi59))))
                    case "":
                        v2 = v2
                    }
                case "a":
                    v70 := GenS1{F0: ("hello" > "foo bar")}
                    _ = v70
                default:
                    var v71 int8 = 100
                    _ = v71
                }
            default:
                for fi72 := 0; fi72 < 20; fi72++ {
                    switch "\n\t" {
                    case "\n\t":
                        v3 = append(v3, 2)
                    default:
                        var v73 int16 = (v24 & 319)
                        _ = v73
                    }
                    if v5 {
                        var v74 int8 = 42
                        if v5 {
                            v75 := v4[1]
                            _ = v75
                        }
                        _ = v74
                    }
                    if (1 < -1) {
                        switch fi59 {
                        case 128:
                            if (!("" <= "hello")) {
                                v76 := GenS9{F0: 42}
                                if (895 >= v24) {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                }
                                var v77 uint16 = 65535
                                _ = v76
                                _ = v77
                            }
                        case 4294967295:
                            if false {
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    if true {
                                        if false {
                                            v78 := v4[547]
                                            v79 := GenS3{F0: 0}
                                            if ((1.5 * v25) > float32(0.1)) {
                                                var v80 uint16 = 1
                                                var v81 uint64 = v1
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = v78
                                            _ = v79
                                        } else {
                                            v82 := GenS5{F0: 42}
                                            _ = v82
                                        }
                                    }
                                    v1 -= (v2 ^ (v1 * 91))
                                    switch "\n\t" {
                                    case "a":
                                        if true {
                                            v25 = v25
                                        }
                                    case "":
                                        fmt.Println(v5)
                                    case "\n\t":
                                        var v83 uint64 = (^(v2))
                                        _ = v83
                                    default:
                                        switch (("foo bar" + "foo bar") + "\n\t") {
                                        case "hello":
                                            var v84 int16 = v24
                                            _ = v84
                                        default:
                                            if true {
                                                v85 := v4[67]
                                                switch 128 {
                                                case 51:
                                                    if (65535 > 2147483647) {
                                                        v86 := v4[0]
                                                        _ = v86
                                                    } else {
                                                        if true {
                                                            if (!v5) {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v87 int32 = 1
                                                                _ = v87
                                                            }
                                                        } else {
                                                            v1 = (v1 | 42)
                                                            fmt.Println(fi59, v25, v25)
                                                        }
                                                    }
                                                case -1:
                                                    delete(v4, 2147483648)
                                                case -2:
                                                    var v88 string = ((v85 + v85) + "hello")
                                                    _ = v88
                                                }
                                                _ = v85
                                            } else {
                                                var v89 bool = (!(1 <= 2147483647))
                                                v2 = 2
                                                var v90 bool = (!(65535 != 2))
                                                v89 = (-226 >= 1)
                                                var v91 int64 = 0
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                    }
                                    var v92 uint = 4294967295
                                    v5 = true
                                    if (!true) {
                                        v2 -= v2
                                        if v5 {
                                            var v93 float64 = -2.25
                                            _ = v93
                                        }
                                        v24 -= (^(v24))
                                    }
                                    v2 = v2
                                    v24 -= v24
                                    _ = v92
                                }
                            }
                        case 931:
                            v2 -= (v2 ^ 0)
                        }
                    }
                    _ = fi72
                }
            }
        } else {
            v25 -= ((((float32(1e10) - ((((v25 / (v25 * 3.14159)) + (((v25 - (3.14159 * float32(0.1))) - (float32(float64((v25 * -2.25))) * (1.5 * v25))) / 0.1)) + float32((-2.25 * 1e10))) * v25)) + ((v25 + 1e10) * ((((((float32(0.1) + 3.14159) - 0.1) / -2.25) + ((v25 - v25) * float32(1.5))) * float32(-2.25)) / float32(1e10)))) / (((float32(1e10) / 1e10) * 0.1) - (((((float32(3.14159) / v25) * v25) / (v25 - v25)) - (0.0 * ((float32(3.14159) - ((v25 / float32(1.5)) * v25)) * 0.0))) * (-2.25 * ((3.14159 - 3.14159) + 3.14159))))) * (v25 / ((0.1 + (((v25 * (v25 + (((float32(0.1) - (v25 / 1.0)) * v25) - v25))) + v25) * 1e10)) * (0.0 * (0.1 / v25)))))
            v5 = (!(((v5 && (-128 > -1)) || (v2 <= (v2 & 255))) || v5))
        }
        _ = fi59
    }
    for fi94 := 0; fi94 < 1; fi94++ {
        v95 := GenS2{F0: 65535}
        v96 := GenS2{F0: 2}
        fmt.Println(v5)
        _ = fi94
        _ = v95
        _ = v96
    }
    v25 = 1.5
    fmt.Println(v2, v2)
    var v97 float32 = (v25 / (float32((((0.0 - (((float64((0.0 / v25)) + 3.14159) + 3.14159) + 1.5)) - (1e10 + -2.25)) - 3.14159)) + 0.0))
    v98 := GenS10{F0: v24}
    switch ("foo bar" + ("hello" + ("hello" + "hello"))) {
    case "\n\t":
        var v99 int16 = v24
        _ = v99
    case "hello":
        var v100 uint32 = 899
        _ = v100
    }
    fmt.Println(v5, v2)
    fmt.Println(v5)
    delete(v4, 2147483648)
    _ = len(v4)
    v25 = float32((-2.25 - float64(float32((-2.25 - (0.1 * (((1e10 * 0.0) / 0.1) / 1.0)))))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
}
