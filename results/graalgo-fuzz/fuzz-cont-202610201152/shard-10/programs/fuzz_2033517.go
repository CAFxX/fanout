package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int32; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int16; F1 int; F2 int }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = 42
        if false {
            var v8 string = ("foo bar" + "hello")
            v9 := []int{718, fi6, fi6}
            v7 = 1
            _ = v8
            _ = v9
        } else {
            if false {
                var v10 bool = (2 <= 32767)
                var v11 uint8 = 128
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 18; fi12++ {
                v7 -= v7
                _ = fi12
            }
        }
        v7 += 2
        _ = fi6
        _ = v7
    }
    switch -675 {
    case -2147483648:
        var v13 float32 = 1e10
        _ = v13
    case 9223372036854775807:
        var v14 bool = (4 == 42)
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        switch "hello" {
        case "":
            if false {
                v16 := GenS1{F0: 1}
                _ = v16
            } else {
                var v17 string = ("a" + "a")
                v18 := []int{fi15}
                for fi19 := 0; fi19 < 9; fi19++ {
                    if (422 <= 1) {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        if (-394 <= -1) {
                            if true {
                                if false {
                                    v17 = "foo bar"
                                    if len(v18) > 0 {
                                        v18[0] = fi15
                                    }
                                    if ((((float64(3.14159) - 3.14159) - 1.5) + float64(-2.25)) >= (1e10 - float64(1.5))) {
                                        v17 = v17
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        var v20 uint = 42
                                        _ = v20
                                    } else {
                                        if true {
                                            v21 := GenS2{F0: ((v17 + ((((v17 + v17) + (("hello" + (v17 + "\n\t")) + v17)) + "") + "foo bar")) + "a"), F1: 42, F2: 0}
                                            var v22 int16 = -1
                                            _ = v21
                                            _ = v22
                                        }
                                        var v23 uint32 = 780
                                        if true {
                                            var v24 bool = false
                                            _ = v24
                                        } else {
                                            switch ((v17 + v17) + (("foo bar" + "foo bar") + v17)) {
                                            case "a":
                                                fmt.Println(fi19, fi15, fi15)
                                            case "":
                                                v25 := GenS1{F0: 540}
                                                _ = v25
                                            default:
                                                var v26 uint64 = 834
                                                _ = v26
                                            }
                                            v17 = ""
                                            if false {
                                                if true {
                                                    if len(v18) > 0 {
                                                        v18[0] = 945
                                                    }
                                                    v23 = (v23 << ((255) & 63))
                                                }
                                                switch "hello" {
                                                case "a":
                                                    v17 = (("" + v17) + "")
                                                default:
                                                    var v27 float64 = 1e10
                                                    _ = v27
                                                }
                                            } else {
                                                var v28 uint64 = 18446744073709551615
                                                _ = v28
                                            }
                                            v18 = append(v18, fi15)
                                        }
                                        _ = v23
                                    }
                                }
                                if (("a" + "a") < v17) {
                                    fmt.Println(fi15)
                                    v29 := GenS3{F0: 514}
                                    _ = v29
                                } else {
                                    v30 := GenS4{F0: 2147483647, F1: -966, F2: fi19}
                                    if (-950 < 0) {
                                        var v31 int64 = 1
                                        var v32 string = v17
                                        if (true || false) {
                                            var v33 uint16 = 32767
                                            _ = v33
                                        }
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v34 int64 = 0
                                        _ = v34
                                    }
                                    var v35 bool = (!(3.14159 <= -2.25))
                                    _ = v30
                                    _ = v35
                                }
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            _ = len(v18)
                            switch fi15 {
                            case 128:
                                v36 := GenS2{F0: v17, F1: 127, F2: 1}
                                var v37 uint = 1
                                _ = v36
                                _ = v37
                            case -9223372036854775808:
                                if len(v18) > 0 {
                                    v18[0] = fi15
                                }
                            case 42:
                                var v38 float64 = float64(float32(-2.25))
                                _ = v38
                            default:
                                v17 = v17
                            }
                            v17 = ""
                        } else {
                            if false {
                                fmt.Println(fi15, fi19, v17)
                                v18 = append(v18, (-(fi15)))
                                if true {
                                    fmt.Println(v17, fi19)
                                    fmt.Println(v17, fi19, v17)
                                    v17 = v17
                                    switch "\n\t" {
                                    case "hello":
                                        if (true && ((42 >= -32768) && (!false))) {
                                            var v39 string = ("a" + v17)
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            _ = v39
                                        } else {
                                            v17 = v17
                                            v17 = v17
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            fmt.Println(fi19)
                                            if true {
                                                var v40 bool = (false || ((!false) || (1 < 127)))
                                                v17 = v17
                                                fmt.Println(fi15, v17, fi19)
                                                fmt.Println(v40, v40)
                                                var v41 bool = (890 > 2)
                                                _ = v40
                                                _ = v41
                                            }
                                        }
                                    default:
                                        _ = len(v18)
                                    }
                                }
                            } else {
                                fmt.Println(fi15, fi19, v17)
                                var v42 int32 = 2147483647
                                fmt.Println(v42, fi19)
                                _ = v42
                            }
                            _ = len(v18)
                            v17 = ""
                        }
                    } else {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        switch -910 {
                        case -8:
                            v43 := GenS5{F0: 57}
                            _ = v43
                        case 32768:
                            fmt.Println(fi15, fi15)
                        case 860:
                            v44 := GenS3{F0: fi19}
                            var v45 uint16 = 2
                            var v46 int8 = 127
                            _ = v44
                            _ = v45
                            _ = v46
                        default:
                            _ = len(v18)
                        }
                    }
                    var v47 uint16 = 42
                    _ = fi19
                    _ = v47
                }
                _ = v17
                _ = v18
            }
        case "foo bar":
            fmt.Println(fi15)
        }
        var v48 int32 = 0
        _ = fi15
        _ = v48
    }
    return 1
}
func genFn1(p0 uint32, p1 uint16) uint {
    p0 = (^(p0))
    for fi49 := 0; fi49 < 18; fi49++ {
        if (4294967295 < 4294967295) {
            fmt.Println(p1, fi49, fi49)
        } else {
            var v50 float32 = 1.5
            v50 = (v50 / (v50 * v50))
            var v51 int32 = -1
            _ = v50
            _ = v51
        }
        var v52 uint = 2
        switch 9223372036854775807 {
        case 742:
            if (!true) {
                v53 := GenS2{F0: "\n\t", F1: genFn0(), F2: p1}
                v54 := []int{fi49, (fi49 + -1), -142}
                _ = v53
                _ = v54
            }
        case 2147483647:
            if false {
                v52 += v52
                var v55 int = fi49
                _ = v55
            }
        }
        _ = fi49
        _ = v52
    }
    v56 := GenS6{F0: 255, F1: 1}
    var v57 bool = true
    for fi58 := 0; fi58 < 11; fi58++ {
        var v59 bool = v57
        v59 = false
        _ = fi58
        _ = v59
    }
    _ = p0
    _ = p1
    _ = v56
    _ = v57
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 61
    var v3 float64 = ((((3.14159 - ((((1e10 * float64(3.14159)) * (-2.25 - 3.14159)) + (-2.25 + -2.25)) / 3.14159)) / 0.1) - 0.0) + float64(0.0))
    v4 := []int{-9223372036854775808, 127, 1}
    v5 := make(map[int]int)
    for fi60 := 0; fi60 < 9; fi60++ {
        if false {
            for fi61 := 0; fi61 < 18; fi61++ {
                fmt.Println(fi61, fi60, v2)
                v5[(-(fi60))] = (fi60 << ((0) & 63))
                fmt.Println(fi60)
                _ = fi61
            }
            fmt.Println(v3)
            if false {
                v62 := GenS6{F0: 1, F1: 1}
                v3 += -2.25
                v63 := GenS5{F0: v2}
                _ = v62
                _ = v63
            } else {
                var v64 bool = true
                v64 = (!(42 == -119))
                _ = v64
            }
        }
        switch fi60 {
        case 127:
            v65 := GenS3{F0: (fi60 & fi60)}
            _ = v65
        case 128:
            if true {
                if false {
                    if (0 == -620) {
                        delete(v1, fi60)
                        fmt.Println(v2, fi60)
                        v2 = genFn0()
                    } else {
                        for fi66 := 0; fi66 < 17; fi66++ {
                            if (true && true) {
                                if false {
                                    v2 -= (^(v2))
                                }
                                v2 += v2
                                var v67 int8 = 100
                                _ = v67
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = fi66
                                }
                                delete(v5, fi60)
                            }
                            if (42 < int8(fi60)) {
                                var v68 uint16 = 727
                                v69 := GenS1{F0: 1}
                                var v70 int8 = 1
                                _ = v68
                                _ = v69
                                _ = v70
                            } else {
                                fmt.Println(fi66)
                                switch -474 {
                                case -1:
                                    v3 += v3
                                    v3 -= 0.0
                                }
                            }
                            v71 := GenS7{F0: -32768, F1: 0, F2: 9223372036854775807}
                            _ = fi66
                            _ = v71
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v72 := GenS3{F0: fi60}
                            _ = v72
                        case "":
                            var v73 uint16 = 255
                            _ = v73
                        case "a":
                            v74 := GenS3{F0: 128}
                            _ = v74
                        default:
                            fmt.Println(v2, v2, v3)
                        }
                        v75 := GenS8{F0: (1 >= 42), F1: 12}
                        _ = v75
                    }
                }
                v2 += 128
            } else {
                v76 := GenS2{F0: "hello", F1: v2, F2: 65535}
                _ = v76
            }
        case 9223372036854775807:
            if false {
                v3 += (v3 + 3.14159)
            }
        default:
            fmt.Println(v2)
        }
        var v77 uint32 = 4294967295
        _ = fi60
        _ = v77
    }
    var v78 int = 2147483647
    if (727 <= 0) {
        switch (v78 | v78) {
        case 0:
            var v79 uint64 = 226
            _ = v79
        case 32768:
            for fi80 := 0; fi80 < 5; fi80++ {
                v81 := GenS7{F0: -359, F1: fi80, F2: 4294967295}
                if (true || true) {
                    var v82 float32 = (((((3.14159 - 1.5) + (float32((v3 * (float64(0.1) - v3))) - -2.25)) + (-2.25 + float32(3.14159))) * 0.0) + 0.1)
                    for fi83 := 0; fi83 < 10; fi83++ {
                        v84 := GenS9{F0: true}
                        fmt.Println(v82, fi83)
                        _ = fi83
                        _ = v84
                    }
                    var v85 bool = false
                    _ = v82
                    _ = v85
                }
                v86 := GenS4{F0: 1, F1: 1, F2: (fi80 + -9223372036854775808)}
                _ = fi80
                _ = v81
                _ = v86
            }
        default:
            for fi87 := 0; fi87 < 18; fi87++ {
                switch (v78 ^ v78) {
                case 2147483647:
                    fmt.Println(v3, v2)
                case 4294967295:
                    var v88 string = ("foo bar" + ("a" + ("" + ("\n\t" + "a"))))
                    _ = v88
                default:
                    for fi89 := 0; fi89 < 10; fi89++ {
                        if false {
                            fmt.Println(v2, v2, v3)
                            var v90 int16 = 182
                            switch (-(v78)) {
                            case 128:
                                v91 := GenS4{F0: 1, F1: -1, F2: 958}
                                _ = v91
                            default:
                                v90 = v90
                            }
                            v2 += v2
                            fmt.Println(v3)
                            _ = v90
                        }
                        var v92 string = "a"
                        fmt.Println(v2)
                        _ = fi89
                        _ = v92
                    }
                }
                _ = fi87
            }
        }
        if (!true) {
            if (false && (4294967295 <= 65535)) {
                v93 := GenS10{F0: -1}
                switch ("\n\t" + (("a" + ("\n\t" + ("foo bar" + (("hello" + ("a" + ("hello" + "hello"))) + "hello")))) + "")) {
                case "a":
                    switch v78 {
                    case 128:
                        var v94 string = "foo bar"
                        _ = v94
                    }
                case "hello":
                    fmt.Println(v3, v2, v78)
                default:
                    v3 -= v3
                }
                fmt.Println(v2, v78)
                _ = v93
            }
        } else {
            delete(v5, (^(v78)))
            var v95 uint16 = 65535
            if false {
                v3 -= v3
            } else {
                v2 = v2
                v3 -= (1e10 - v3)
                v2 -= v2
            }
            _ = v95
        }
        fmt.Println(v78, v78)
    } else {
        if (true && (v3 != float64((1.5 / float32((0.0 - (float64(float32((((-2.25 / v3) * (3.14159 * (float64(1.5) / float64((1e10 + 3.14159))))) + v3))) + -2.25))))))) {
            v5[v78] = -2147483648
            v4 = append(v4, 1)
        } else {
            if true {
                v2 = 127
                v2 -= (v2 + v2)
                v3 += (v3 * ((v3 / float64(float32(v3))) + 1e10))
            }
            var v96 int8 = -69
            _ = v96
        }
        v97 := GenS11{F0: 519}
        _ = v97
    }
    var v98 float32 = 1.5
    switch (-(v78)) {
    case 42:
        var v99 int8 = -126
        _ = v99
    default:
        v100 := GenS9{F0: false}
        _ = v100
    }
    fmt.Println(v3, v78)
    fmt.Println(v78, v98)
    fmt.Println(v78, v2, v3)
    var v101 float32 = v98
    var v102 int16 = 32767
    fmt.Println(v78, v3, v2)
    switch v78 {
    case -128:
        v2 -= 2
    case 9223372036854775807:
        v2 = (v2 / ((v2) | 1))
    case 136:
        fmt.Println(v2, v101, v78)
    }
    switch ((("hello" + "foo bar") + ("a" + "a")) + ((("" + "hello") + "a") + "")) {
    case "\n\t":
        for fi103 := 0; fi103 < 9; fi103++ {
            var v104 string = "\n\t"
            var v105 int32 = 48
            _ = fi103
            _ = v104
            _ = v105
        }
    case "":
        for fi106 := 0; fi106 < 8; fi106++ {
            v3 = v3
            var v107 int16 = v102
            _ = fi106
            _ = v107
        }
    case "a":
        if true {
            var v108 string = "hello"
            v108 = ("hello" + v108)
            _ = v108
        } else {
            v102 = v102
        }
    }
    v3 += float64(-2.25)
    var v109 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v78:", v78)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v101
    _ = v102
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v98
}
