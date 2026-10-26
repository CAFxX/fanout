package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int8; F1 string; F2 string }
func genFn0(p0 uint64, p1 int32, p2 uint16) uint {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 string = ("a" + ("foo bar" + "foo bar"))
        var v6 int32 = 0
        var v7 int8 = 100
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    switch 32767 {
    case 2147483647:
        switch 42 {
        case 32768:
            var v8 int64 = 0
            _ = v8
        default:
            for fi9 := 0; fi9 < 10; fi9++ {
                p1 += p1
                switch ("hello" + ("" + (("hello" + "hello") + "a"))) {
                case "":
                    if (42 >= -428) {
                        var v10 int64 = 42
                        p0 -= p0
                        _ = v10
                    }
                case "hello":
                    var v11 float32 = ((3.14159 - -2.25) + (3.14159 * (-2.25 + (((0.1 / (-2.25 * 1.5)) * (0.1 * (1.5 * 1.5))) * 0.0))))
                    _ = v11
                case "foo bar":
                    var v12 float32 = (0.0 - (0.0 - ((-2.25 * 1.5) + -2.25)))
                    _ = v12
                }
                _ = fi9
            }
        }
    default:
        if false {
            var v13 int64 = -9223372036854775808
            _ = v13
        } else {
            v14 := GenS1{F0: 0, F1: "a", F2: 1}
            if false {
                p2 = (p2 % ((p2) | 1))
            } else {
                v15 := GenS1{F0: 2147483648, F1: "foo bar", F2: 255}
                _ = v15
            }
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 278
    var v3 int32 = v1
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 bool = false
        if (18446744073709551615 <= 1) {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int32 = (v1 - v1)
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v17 = (v17 || (("\n\t" + "") >= "hello"))
                switch "a" {
                case "\n\t":
                    switch "" {
                    case "foo bar":
                        if false {
                            var v21 uint64 = (v2 + v2)
                            fmt.Println(v2, fi20)
                            if false {
                                v22 := GenS1{F0: 606, F1: "", F2: 2}
                                _ = v22
                            }
                            _ = v21
                        }
                    }
                case "a":
                    var v23 int = fi16
                    _ = v23
                case "hello":
                    if v17 {
                        if (42 >= 2) {
                            var v24 uint = 4294967295
                            _ = v24
                        } else {
                            switch fi16 {
                            case 127:
                                switch (fi20 | (fi20 | 228)) {
                                case -375:
                                    v17 = v17
                                case 1:
                                    fmt.Println(v1)
                                case 510:
                                    fmt.Println(v1, v17)
                                default:
                                    var v25 uint32 = 2147483648
                                    _ = v25
                                }
                            case 4294967295:
                                switch fi16 {
                                case 9223372036854775807:
                                    var v26 int64 = 9223372036854775807
                                    _ = v26
                                default:
                                    if (v17 || v17) {
                                        fmt.Println(v17, fi16, v1)
                                        if (!(!v17)) {
                                            if (!("foo bar" < "foo bar")) {
                                                v27 := GenS2{F0: (fi16 + fi20), F1: 0}
                                                if v17 {
                                                    v28 := []int{fi20}
                                                    _ = v28
                                                } else {
                                                    switch fi20 {
                                                    case 9223372036854775807:
                                                        var v29 int16 = 621
                                                        _ = v29
                                                    case 624:
                                                        var v30 float32 = float32((0.0 * float64(0.0)))
                                                        _ = v30
                                                    }
                                                    fmt.Println(fi20, v2, v17)
                                                }
                                                if v17 {
                                                    v31 := []int{fi20, (fi20 | fi16), 86, fi20}
                                                    _ = v31
                                                }
                                                _ = v27
                                            }
                                            var v32 uint64 = (v2 % (((v2 & (^(v2)))) | 1))
                                            v33 := GenS2{F0: fi16, F1: 100}
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := GenS3{F0: -232, F1: 127}
                                        _ = v34
                                    }
                                    v35 := []int{(fi16 | (fi16 % (((fi20 << ((65535) & 63))) | 1)))}
                                    _ = v35
                                }
                            case 32767:
                                v3 += v1
                            }
                            var v36 float64 = (3.14159 - ((1.5 + float64(3.14159)) - -2.25))
                            fmt.Println(v2, fi20, v3)
                            _ = v36
                        }
                        var v37 int = fi16
                        _ = v37
                    }
                default:
                    v38 := []int{fi16, fi20}
                    if v17 {
                        if len(v38) >= 2 {
                            v38 = v38[1:len(v38)]
                        }
                        var v39 uint8 = 97
                        v38 = append(v38, (fi16 + (-(fi20))))
                        _ = len(v38)
                        _ = v39
                    } else {
                        switch 4294967295 {
                        case -624:
                            v40 := GenS3{F0: 1, F1: 0}
                            _ = v40
                        case -1:
                            var v41 int32 = -1
                            v3 = (v1 * (v3 >> ((801) & 63)))
                            var v42 string = "foo bar"
                            v3 = 1
                            var v43 float64 = float64(0.1)
                            _ = v41
                            _ = v42
                            _ = v43
                        default:
                            fmt.Println(v1, fi16)
                        }
                    }
                    var v44 uint16 = 32767
                    switch (-(fi16)) {
                    case -9223372036854775808:
                        if (1.5 <= (0.0 + ((((-2.25 + 1e10) / ((0.0 + (0.1 / ((0.1 + -2.25) - 0.1))) - (0.0 * (1e10 + (1.5 - -2.25))))) / 3.14159) - ((1.5 - (float32(0.0) * 0.0)) * 1e10)))) {
                            var v45 float64 = ((((((0.0 + ((((1e10 / 1.0) + 1e10) / 1.5) + 0.1)) - 1.5) / float64(0.1)) / 1e10) * 1e10) + 1.5)
                            var v46 bool = (fi20 == 763)
                            _ = v45
                            _ = v46
                        }
                    default:
                        v17 = ((v2 ^ (v2 % ((v2) | 1))) > v2)
                    }
                    _ = v38
                    _ = v44
                }
                _ = fi20
            }
            var v47 bool = v17
            _ = v47
        } else {
            v48 := []int{fi16, -2147483648, fi16, (fi16 >> ((1) & 63))}
            if ((42 >= v1) || (32767 >= 32767)) {
                v49 := GenS4{F0: -1, F1: 255, F2: -1}
                _ = v49
            }
            _ = v48
        }
        _ = fi16
        _ = v17
    }
    switch "foo bar" {
    case "a":
        v50 := []int{1, -578}
        _ = v50
    default:
        var v51 uint = 42
        _ = v51
    }
    if (!true) {
        var v52 float32 = 3.14159
        _ = v52
    }
    if true {
        v53 := []int{831, 32767, -565}
        _ = v53
    }
    if true {
        for fi54 := 0; fi54 < 3; fi54++ {
            v55 := GenS5{F0: true}
            if (0.1 < 1.5) {
                v2 = (v2 >> ((42) & 63))
                for fi56 := 0; fi56 < 1; fi56++ {
                    var v57 uint32 = 1
                    v58 := GenS6{F0: (v2 - (v2 / ((v2) | 1))), F1: false}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            }
            _ = fi54
            _ = v55
        }
        v2 = 687
        switch 32767 {
        case -9223372036854775808:
            switch -273 {
            case -816:
                fmt.Println(v2, v1, v2)
            case 32767:
                var v59 float64 = 0.1
                _ = v59
            case 128:
                if true {
                    var v60 float64 = float64(0.0)
                    v61 := GenS6{F0: 4294967295, F1: false}
                    _ = v60
                    _ = v61
                }
            default:
                v62 := GenS7{F0: 127, F1: ("foo bar" + ""), F2: "hello"}
                _ = v62
            }
        case 128:
            if (false || (-128 == 28)) {
                switch (("a" + "\n\t") + ("\n\t" + ("hello" + ("a" + "\n\t")))) {
                case "hello":
                    v3 -= v1
                case "a":
                    var v63 int8 = 73
                    _ = v63
                case "foo bar":
                    v64 := []int{}
                    _ = v64
                default:
                    v1 -= v1
                }
                var v65 bool = true
                _ = v65
            }
        case -550:
            v3 -= v3
        }
    } else {
        var v66 uint64 = v2
        _ = v66
    }
    switch ("hello" + (("a" + ("" + "a")) + "foo bar")) {
    case "foo bar":
        var v67 int16 = -32768
        _ = v67
    }
    for fi68 := 0; fi68 < 14; fi68++ {
        fmt.Println(v3)
        switch (fi68 + (fi68 | (fi68 + fi68))) {
        case -1:
            if false {
                v2 -= (v2 & (^(v2)))
            } else {
                v69 := []int{fi68, fi68, fi68, fi68}
                _ = v69
            }
        case 1:
            v70 := []int{}
            _ = v70
        default:
            v1 += (-(v3))
        }
        fmt.Println(fi68)
        _ = fi68
    }
    v1 = -1
    var v71 int16 = -1
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v71
}
