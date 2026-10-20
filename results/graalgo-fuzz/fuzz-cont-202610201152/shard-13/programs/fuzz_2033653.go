package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int64; F2 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint16; F1 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint16; F1 bool }
type GenS14 struct { F0 int64; F1 uint32 }
func genFn0(p0 int8, p1 uint64, p2 uint32) int32 {
    fmt.Println(p0)
    p1 += 867
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() int8 {
    var v4 float64 = (1e10 - float64(((1.5 / 3.14159) - (1.5 + 3.14159))))
    fmt.Println(v4)
    var v5 int = 127
    var v6 int8 = 42
    v4 = (0.1 + v4)
    _ = v4
    _ = v5
    _ = v6
    return 96
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = ((((0.1 - -2.25) - (((3.14159 + -2.25) / 0.1) + float64((0.1 + 1.5)))) / 1e10) / 3.14159)
    var v3 float64 = v2
    if ((true && false) || true) {
        v3 = 3.14159
        if (!(2147483647 >= 0)) {
            var v7 bool = (!false)
            for fi8 := 0; fi8 < 12; fi8++ {
                switch "a" {
                case "a":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        if v7 {
                            v2 = ((v3 - (0.1 - 1e10)) / -2.25)
                            v3 = v3
                            v7 = (true || false)
                            v10 := GenS1{F0: 9223372036854775807}
                            v11 := []int{2147483647, (fi8 ^ (fi8 / ((fi9) | 1)))}
                            _ = v10
                            _ = v11
                        }
                        _ = fi9
                    }
                }
                if false {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        switch 128 {
                        case -2147483648:
                            var v13 uint = 971
                            _ = v13
                        case 1:
                            v14 := v1["foo bar"]
                            _ = v14
                        case 128:
                            delete(v1, "hello")
                        }
                        _ = fi12
                    }
                } else {
                    fmt.Println(v3)
                }
                if (!v7) {
                    v15 := GenS2{F0: 1}
                    _ = v15
                } else {
                    var v16 uint8 = 42
                    _ = v16
                }
                _ = fi8
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := []int{fi17, fi17}
                if (!(false && (-1 <= -2147483648))) {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := GenS1{F0: 1}
                        _ = fi19
                        _ = v20
                    }
                    if (566 <= 812) {
                        if v7 {
                            v2 = ((v2 / v2) / v3)
                            v18 = append(v18, fi17)
                        } else {
                            v7 = ((float64(0.1) / ((0.1 - v2) + float64(1.5))) >= v3)
                        }
                        var v21 uint16 = 881
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v23 := GenS3{F0: 32767, F1: 4294967295}
                            v7 = (v2 >= v2)
                            var v24 uint = 901
                            if len(v18) > 0 {
                                v18[0] = 32768
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    } else {
                        v25 := GenS4{F0: 2, F1: 2}
                        v26 := GenS4{F0: 0, F1: 605}
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS3{F0: -32768, F1: 1}
                    _ = v27
                } else {
                    if v7 {
                        for fi28 := 0; fi28 < 3; fi28++ {
                            if ((("\n\t" + (("foo bar" + "hello") + (("" + "foo bar") + ""))) + "a") >= "a") {
                                var v29 uint16 = 65535
                                if v7 {
                                    v30 := v1["hello"]
                                    _ = len(v18)
                                    v31 := GenS4{F0: v30, F1: (v30 << ((1) & 63))}
                                    var v32 string = ("hello" + "")
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    if false {
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        v1["hello"] = 0
                                        if false {
                                            if false {
                                                _ = len(v18)
                                                v2 += v3
                                                v33 := GenS3{F0: 0, F1: 42}
                                                _ = v33
                                            }
                                            switch 0 {
                                            case -2147483648:
                                                v7 = false
                                                if v7 {
                                                    if ((1 == 1000) || v7) {
                                                        fmt.Println(v3, fi17, v3)
                                                        switch ("" + "a") {
                                                        case "":
                                                            v7 = (!(42 <= 42))
                                                        }
                                                        var v34 float64 = (v3 + (v2 / (v2 + float64(((((float32(1e10) * 0.1) * 3.14159) / 3.14159) * ((0.1 - -2.25) - 0.0))))))
                                                        switch ((("a" + "a") + ("" + "hello")) + "") {
                                                        case "a":
                                                            if (v7 && true) {
                                                                _ = len(v1)
                                                                switch int(v29) {
                                                                case 538:
                                                                    var v35 int = 14
                                                                    _ = v35
                                                                case -210:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v36 string = "foo bar"
                                                                        _ = v36
                                                                    case "a":
                                                                        var v37 uint64 = 0
                                                                        _ = v37
                                                                    default:
                                                                        var v38 int8 = -94
                                                                        _ = v38
                                                                    }
                                                                case -2147483648:
                                                                    fmt.Println(v34)
                                                                }
                                                                if (!v7) {
                                                                    if true {
                                                                        var v39 string = "\n\t"
                                                                        _ = v39
                                                                    } else {
                                                                        var v40 uint8 = 1
                                                                        var v41 int16 = 993
                                                                        if true {
                                                                            v42 := GenS1{F0: 42}
                                                                            v43 := GenS4{F0: 65535, F1: 65535}
                                                                            v44 := GenS1{F0: 0}
                                                                            _ = v42
                                                                            _ = v43
                                                                            _ = v44
                                                                        } else {
                                                                            v45 := GenS5{F0: 0}
                                                                            v46 := GenS4{F0: 2147483648, F1: 938}
                                                                            var v47 uint32 = 1
                                                                            v40 += v40
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                    v18 = append(v18, (-(fi17)))
                                                                    v34 -= 1e10
                                                                    fmt.Println(v3)
                                                                    if ((42 < 0) || false) {
                                                                        v7 = v7
                                                                        fmt.Println(v7, v2)
                                                                        switch (fi28 ^ 2147483647) {
                                                                        case -1:
                                                                            v48 := GenS2{F0: 2147483648}
                                                                            v2 = v2
                                                                            _ = v48
                                                                        default:
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                fmt.Println(v2)
                                                                            case "a":
                                                                                v34 = v3
                                                                            }
                                                                        }
                                                                        v2 = float64(-2.25)
                                                                    } else {
                                                                        var v49 int16 = -808
                                                                        var v50 string = "\n\t"
                                                                        if v7 {
                                                                            v34 -= 0.0
                                                                        } else {
                                                                            if v7 {
                                                                                fmt.Println(v7, v2, v49)
                                                                                var v51 float32 = 0.0
                                                                                _ = v51
                                                                            }
                                                                        }
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(fi17)
                                                                v29 = (v29 - 2)
                                                                var v52 int = (fi28 * fi28)
                                                                if len(v18) > 0 {
                                                                    v18[0] = (-(fi28))
                                                                }
                                                                v52 += v52
                                                                v7 = v7
                                                                _ = v52
                                                            }
                                                        case "\n\t":
                                                            v2 -= 3.14159
                                                        }
                                                        _ = v34
                                                    }
                                                } else {
                                                    if v7 {
                                                        var v53 int8 = genFn1()
                                                        if (!true) {
                                                            var v54 uint = 82
                                                            switch (fi17 ^ fi17) {
                                                            case -9223372036854775808:
                                                                v55 := GenS1{F0: -867}
                                                                _ = v55
                                                            case -1:
                                                                var v56 float64 = v2
                                                                _ = v56
                                                            }
                                                            _ = v54
                                                        } else {
                                                            if (v29 <= (^(v29))) {
                                                                v53 += (-(v53))
                                                            }
                                                            fmt.Println(v2)
                                                            fmt.Println(v3, v7, fi17)
                                                            v2 = ((((v3 - (v3 - v3)) + (3.14159 - (v3 + (float64(float32(float64(3.14159))) / v3)))) * (3.14159 * ((float64(1e10) + v3) / (float64(0.1) + ((1e10 / (v2 - v3)) - 1e10))))) + v2)
                                                        }
                                                        _ = v53
                                                    }
                                                    fmt.Println(v3)
                                                }
                                            default:
                                                v7 = true
                                            }
                                            v29 += v29
                                        } else {
                                            if len(v18) > 0 {
                                                v18[0] = fi28
                                            }
                                            if false {
                                                fmt.Println(v2, fi17, v29)
                                                v7 = v7
                                                fmt.Println(v3, fi17)
                                                var v57 string = "hello"
                                                _ = v57
                                            } else {
                                                fmt.Println(fi28)
                                                var v58 int64 = 9223372036854775807
                                                var v59 string = "a"
                                                _ = v58
                                                _ = v59
                                            }
                                            delete(v1, ("foo bar" + ("a" + "foo bar")))
                                        }
                                    } else {
                                        if (v7 || v7) {
                                            var v60 int32 = -2147483648
                                            var v61 int32 = 2147483647
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    if v7 {
                                        if ((-(fi28)) <= fi17) {
                                            delete(v1, "a")
                                            if (!(0 != 65535)) {
                                                var v62 bool = v7
                                                var v63 int8 = 57
                                                _ = v62
                                                _ = v63
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v64 int32 = 0
                                                _ = v64
                                            }
                                        } else {
                                            v29 = (^(v29))
                                            fmt.Println(v3, v7)
                                        }
                                        if v7 {
                                            v2 = (1.5 / (3.14159 / (1.5 * float64((0.1 * 0.1)))))
                                            v65 := GenS4{F0: 2147483648, F1: 2147483647}
                                            _ = v65
                                        } else {
                                            switch (fi28 / ((fi17) | 1)) {
                                            case 1:
                                                if ((v3 / v2) == 1e10) {
                                                    var v66 int16 = 1000
                                                    _ = len(v1)
                                                    var v67 int8 = 42
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    _ = len(v1)
                                                    var v68 uint = 0
                                                    _ = v68
                                                }
                                            }
                                            v2 += float64(float32(v2))
                                            v7 = (!false)
                                        }
                                        if ((^(v29)) >= v29) {
                                            switch fi28 {
                                            case -162:
                                                switch (fi28 ^ (fi28 | (fi28 & fi17))) {
                                                case 2147483647:
                                                    _ = len(v1)
                                                case 4294967295:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v69 uint8 = 127
                                                        _ = v69
                                                    default:
                                                        v70 := v1[((("hello" + ("" + "")) + (("\n\t" + "foo bar") + "\n\t")) + "\n\t")]
                                                        _ = v70
                                                    }
                                                case -9223372036854775808:
                                                    v18 = append(v18, 128)
                                                default:
                                                    if (1 <= 127) {
                                                        v71 := GenS2{F0: 4294967295}
                                                        v72 := GenS6{F0: 0, F1: 98, F2: 439}
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        if v7 {
                                                            v73 := GenS7{F0: 0}
                                                            if v7 {
                                                                v2 = (-2.25 * (v2 + (0.0 + 0.1)))
                                                                fmt.Println(v7, v7)
                                                                if v7 {
                                                                    fmt.Println(fi17, fi28)
                                                                    fmt.Println(fi17)
                                                                    v74 := GenS7{F0: 42}
                                                                    _ = len(v18)
                                                                    v7 = (42 < 208)
                                                                    if (42 <= 32767) {
                                                                        if (!((v29 * (v29 / ((v29) | 1))) >= v29)) {
                                                                            _ = len(v1)
                                                                            switch (-(fi28)) {
                                                                            case -2147483648:
                                                                                fmt.Println(v29)
                                                                            case 32767:
                                                                                if v7 {
                                                                                    if v7 {
                                                                                        if true {
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(fi17, v7, v7)
                                                                                            }
                                                                                        } else {
                                                                                            var v75 uint32 = 1
                                                                                            v76 := v1["a"]
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        }
                                                                                        if v7 {
                                                                                            v3 = v2
                                                                                            if len(v18) >= 2 {
                                                                                                v18 = v18[1:len(v18)]
                                                                                            }
                                                                                            var v77 int16 = 32767
                                                                                            _ = v77
                                                                                        } else {
                                                                                            var v78 bool = true
                                                                                            var v79 int64 = 1000
                                                                                            v2 -= (v3 / 3.14159)
                                                                                            var v80 int8 = 42
                                                                                            v81 := GenS8{F0: v79, F1: v29}
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        }
                                                                                        v2 += float64(0.0)
                                                                                    }
                                                                                    v29 += v29
                                                                                } else {
                                                                                    fmt.Println(fi28, v7)
                                                                                }
                                                                            case 2147483647:
                                                                                var v82 string = (("\n\t" + "\n\t") + "foo bar")
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        fmt.Println(fi17)
                                                                    }
                                                                    v7 = v7
                                                                    _ = v74
                                                                }
                                                                fmt.Println(fi28, v29)
                                                                v7 = (true || (!v7))
                                                                switch fi28 {
                                                                case -456:
                                                                    if v7 {
                                                                        v83 := GenS9{F0: 1, F1: 100}
                                                                        var v84 int64 = 1
                                                                        delete(v1, ("hello" + "foo bar"))
                                                                        _ = len(v1)
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                case -1:
                                                                    if (222 != 2) {
                                                                        if len(v18) > 0 {
                                                                            v18[0] = 128
                                                                        }
                                                                        v2 = v3
                                                                        v1["hello"] = 2147483647
                                                                    } else {
                                                                        var v85 int16 = -646
                                                                        delete(v1, ("a" + ("hello" + "foo bar")))
                                                                        var v86 uint8 = 1
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                case 128:
                                                                    var v87 uint64 = 42
                                                                    _ = v87
                                                                default:
                                                                    var v88 int = (fi28 | fi17)
                                                                    var v89 uint8 = 255
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v7 = v7
                                                                case "\n\t":
                                                                    if v7 {
                                                                        var v90 uint64 = 18446744073709551615
                                                                        v3 -= v2
                                                                        if (!false) {
                                                                            v1["hello"] = 2147483647
                                                                            var v91 uint8 = 167
                                                                            switch (fi17 + -956) {
                                                                            case -2147483648:
                                                                                var v92 int16 = 1
                                                                                _ = v92
                                                                            case 742:
                                                                                if v7 {
                                                                                    fmt.Println(v7, v91)
                                                                                    fmt.Println(v90, v3)
                                                                                    var v93 bool = v7
                                                                                    _ = v93
                                                                                }
                                                                            case 0:
                                                                                v1[("" + "foo bar")] = 2147483648
                                                                            default:
                                                                                v7 = v7
                                                                            }
                                                                            switch fi28 {
                                                                            case 431:
                                                                                var v94 int16 = -32768
                                                                                _ = v94
                                                                            }
                                                                            fmt.Println(v90)
                                                                            _ = v91
                                                                        }
                                                                        _ = v90
                                                                    }
                                                                case "a":
                                                                    fmt.Println(v7, v29, fi28)
                                                                }
                                                            }
                                                            _ = v73
                                                        }
                                                        if len(v18) >= 2 {
                                                            v18 = v18[1:len(v18)]
                                                        }
                                                        v95 := GenS10{F0: (fi28 - fi28), F1: 2}
                                                        v2 = (v3 / 1.5)
                                                        v96 := GenS8{F0: 9223372036854775807, F1: v29}
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                }
                                            case -1:
                                                var v97 float32 = 0.0
                                                _ = v97
                                            case 4294967295:
                                                if false {
                                                    fmt.Println(fi17, v3)
                                                }
                                            default:
                                                var v98 uint64 = 61
                                                _ = v98
                                            }
                                        } else {
                                            v3 = v2
                                        }
                                    }
                                }
                                v99 := GenS11{F0: 65535, F1: 0}
                                _ = v29
                                _ = v99
                            } else {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                            v100 := GenS6{F0: 311, F1: 9223372036854775807, F2: 255}
                            _ = fi28
                            _ = v100
                        }
                        switch fi17 {
                        case -2147483648:
                            var v101 int64 = 4
                            _ = v101
                        default:
                            _ = len(v1)
                        }
                        v7 = v7
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = v7
        }
    } else {
        fmt.Println(v2, v3)
        v2 = float64((float32((v3 * float64(-2.25))) * (1e10 / (float32(float64(1e10)) + float32(1.5)))))
        for fi102 := 0; fi102 < 5; fi102++ {
            v3 = 0.0
            var v103 uint16 = 65535
            _ = fi102
            _ = v103
        }
    }
    v104 := []int{}
    _ = len(v1)
    v3 = 3.14159
    for fi105 := 0; fi105 < 20; fi105++ {
        v106 := GenS12{F0: -988}
        _ = fi105
        _ = v106
    }
    v3 = (v2 * -2.25)
    v3 -= ((((((float64((1.5 + 0.0)) + (v2 - v3)) / ((v3 / v2) / v2)) * (1.5 / 1e10)) / float64((float32(float64((((1.5 * 1e10) * 1.5) + (1.5 + float32(((((v3 / v3) * v3) / v3) + (v3 + v3))))))) - (0.0 + (0.1 * 1.5))))) - (((1e10 + (3.14159 + v2)) * v2) + v2)) - 1e10)
    if false {
        if true {
            var v107 bool = true
            _ = v107
        } else {
            if (!(true && (42 != 42))) {
                for fi108 := 0; fi108 < 5; fi108++ {
                    v109 := GenS5{F0: 635}
                    _ = fi108
                    _ = v109
                }
                for fi110 := 0; fi110 < 18; fi110++ {
                    for fi111 := 0; fi111 < 7; fi111++ {
                        switch -9223372036854775808 {
                        case -127:
                            fmt.Println(v2, fi111, fi110)
                        default:
                            var v112 int = fi111
                            _ = v112
                        }
                        v2 -= -2.25
                        _ = fi111
                    }
                    v2 = (v2 * -2.25)
                    v3 = (v2 * (float64((-2.25 * 3.14159)) * (0.1 - (v3 - (v3 + 1e10)))))
                    _ = fi110
                }
                v2 = ((1e10 * 0.0) * v3)
            }
        }
        if true {
            switch 2147483647 {
            case 2147483647:
                fmt.Println(v2)
            case -458:
                v3 += -2.25
            case -747:
                for fi113 := 0; fi113 < 9; fi113++ {
                    if len(v104) > 0 {
                        v104[0] = fi113
                    }
                    var v114 string = "\n\t"
                    v114 = v114
                    _ = fi113
                    _ = v114
                }
            default:
                if false {
                    switch 2147483647 {
                    case 32768:
                        v115 := GenS10{F0: 127, F1: 246}
                        _ = v115
                    default:
                        fmt.Println(v3, v2)
                    }
                    if true {
                        fmt.Println(v2, v2)
                    }
                    switch "" {
                    case "hello":
                        var v116 float32 = 3.14159
                        _ = v116
                    }
                } else {
                    if false {
                        var v117 uint64 = 4294967295
                        var v118 uint64 = v117
                        v117 = v117
                        _ = v117
                        _ = v118
                    } else {
                        for fi119 := 0; fi119 < 11; fi119++ {
                            v3 = v3
                            v2 = 1.5
                            delete(v1, ("foo bar" + ("a" + ((("a" + "a") + "a") + ""))))
                            _ = fi119
                        }
                    }
                    var v120 uint16 = 0
                    _ = v120
                }
            }
        } else {
            if (-506 > 68) {
                fmt.Println(v3, v2)
                v121 := GenS11{F0: 702, F1: 914}
                var v122 int64 = 1
                _ = v121
                _ = v122
            }
            fmt.Println(v2)
            if true {
                v2 -= 1.5
                fmt.Println(v2)
            } else {
                if (2147483648 >= 799) {
                    v123 := GenS13{F0: 32767, F1: false}
                    delete(v1, ("\n\t" + ""))
                    _ = v123
                } else {
                    if true {
                        v104 = append(v104, -231)
                        fmt.Println(v3, v3)
                    } else {
                        switch "foo bar" {
                        case "a":
                            v3 = v2
                        case "foo bar":
                            fmt.Println(v3)
                        }
                        if (65535 <= 2) {
                            v2 = ((1.5 * v2) - v3)
                            switch 2147483647 {
                            case 1:
                                fmt.Println(v3)
                            case 2147483647:
                                v3 -= (v2 + v3)
                            case 4294967295:
                                var v124 float64 = v3
                                _ = v124
                            default:
                                switch -9223372036854775808 {
                                case -9223372036854775808:
                                    v3 += v3
                                case 631:
                                    for fi125 := 0; fi125 < 7; fi125++ {
                                        for fi126 := 0; fi126 < 10; fi126++ {
                                            var v127 uint16 = 900
                                            if true {
                                                v128 := GenS8{F0: 228, F1: (v127 * v127)}
                                                var v129 int32 = genFn0(-47, 42, 580)
                                                v2 = (v2 - (((((v2 * v2) / 1.5) / ((((v2 - float64(1.5)) * -2.25) / 0.1) / v3)) - 1e10) / -2.25))
                                                _ = v128
                                                _ = v129
                                            }
                                            _ = fi126
                                            _ = v127
                                        }
                                        _ = fi125
                                    }
                                }
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            if false {
                                var v130 float64 = v3
                                _ = v130
                            }
                        }
                    }
                    var v131 bool = (2 > 1)
                    _ = v131
                }
                for fi132 := 0; fi132 < 12; fi132++ {
                    v2 -= v2
                    v133 := v1["a"]
                    _ = fi132
                    _ = v133
                }
                if (32768 >= 1) {
                    var v134 string = ("a" + ("foo bar" + "\n\t"))
                    _ = v134
                }
            }
        }
    }
    switch 9223372036854775807 {
    case 128:
        var v135 uint32 = 1
        _ = v135
    default:
        v136 := GenS2{F0: 2}
        _ = v136
    }
    fmt.Println(v3, v3)
    var v137 uint64 = 32
    var v138 uint8 = 0
    if false {
        for fi139 := 0; fi139 < 1; fi139++ {
            for fi140 := 0; fi140 < 1; fi140++ {
                var v141 int64 = 9223372036854775807
                var v142 int64 = v141
                _ = fi140
                _ = v141
                _ = v142
            }
            v2 = float64(3.14159)
            _ = fi139
        }
        switch ("foo bar" + "a") {
        case "hello":
            v143 := GenS9{F0: v137, F1: 127}
            _ = v143
        default:
            v144 := GenS14{F0: -694, F1: 150}
            _ = v144
        }
        fmt.Println(v137)
    }
    var v145 uint16 = 580
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v104
    _ = v137
    _ = v138
    _ = v145
    _ = v2
    _ = v3
}
