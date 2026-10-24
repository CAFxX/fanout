package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int32; F2 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 int32) string {
    var v7 int32 = (p0 / (((p0 ^ p0)) | 1))
    v8 := GenS1{F0: (255 == 12), F1: 2, F2: 65535}
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v7, fi9, fi9)
        v10 := []int{(fi9 >> ((42) & 63)), (^(fi9)), (fi9 ^ fi9)}
        var v11 int16 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 = p0
    for fi12 := 0; fi12 < 20; fi12++ {
        p0 -= (v7 * 1)
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int = -2147483648
            v7 += 42
            p0 += v7
            if true {
                switch "" {
                case "a":
                    v15 := []int{fi13, (^(v14))}
                    _ = v15
                }
                var v16 uint = 18446744073709551615
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(fi12, v7)
            v7 = 0
            _ = fi17
        }
        _ = fi12
    }
    _ = p0
    _ = v7
    _ = v8
    return ("foo bar" + ((((("a" + "") + "") + "\n\t") + "a") + "foo bar"))
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = (!false)
    v3 := make(map[string]string)
    v4 := []string{("foo bar" + "foo bar")}
    var v5 float32 = ((1.5 - (0.0 + (1.5 / v1))) - 0.0)
    v6 := []int{32767, 32767}
    v18 := GenS1{F0: v2, F1: 2, F2: 111}
    switch -9223372036854775808 {
    case -1:
        if (v2 || (((true || true) || (4294967295 >= 42)) && false)) {
            for fi19 := 0; fi19 < 11; fi19++ {
                switch fi19 {
                case 499:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if true {
                            v2 = true
                            if (false || false) {
                                delete(v3, ("foo bar" + "foo bar"))
                                fmt.Println(v1, v5)
                            } else {
                                v21 := GenS2{F0: 42, F1: -2147483648, F2: 127}
                                if (!(false && v2)) {
                                    var v22 int = fi19
                                    _ = v22
                                }
                                fmt.Println(fi19, fi20)
                                _ = v21
                            }
                        }
                        var v23 int16 = -1
                        _ = fi20
                        _ = v23
                    }
                case -46:
                    var v24 uint32 = 108
                    _ = v24
                case 490:
                    switch 4294967295 {
                    case 127:
                        if (!v2) {
                            var v25 int64 = 329
                            fmt.Println(v5)
                            _ = len(v6)
                            _ = v25
                        }
                    case 32768:
                        switch ("\n\t" + "") {
                        case "\n\t":
                            if (0 <= 0) {
                                v5 = 0.1
                                switch 32767 {
                                case 1:
                                    if (!v2) {
                                        var v26 int16 = 294
                                        v5 -= (1.5 - (v1 / (3.14159 / (float32(-2.25) * (((-2.25 * v1) - v5) + v5)))))
                                        _ = v26
                                    }
                                default:
                                    v27 := GenS2{F0: 43, F1: 0, F2: 42}
                                    _ = v27
                                }
                            }
                        default:
                            switch "" {
                            case "foo bar":
                                v28 := GenS2{F0: 4294967295, F1: 999, F2: 0}
                                _ = v28
                            case "hello":
                                var v29 uint8 = 0
                                _ = v29
                            case "a":
                                v30 := GenS3{F0: -199}
                                _ = v30
                            }
                        }
                    }
                }
                _ = fi19
            }
        }
    case 42:
        v31 := GenS3{F0: -330}
        _ = v31
    default:
        v5 += (v1 * v1)
    }
    switch "a" {
    case "hello":
        v32 := GenS1{F0: true, F1: 18446744073709551615, F2: 1}
        _ = v32
    case "foo bar":
        var v33 int8 = -80
        _ = v33
    case "a":
        v1 = -2.25
    default:
        switch 32768 {
        case 935:
            if v2 {
                var v34 float32 = 1.5
                _ = len(v4)
                _ = v34
            } else {
                v35 := GenS1{F0: (v2 || false), F1: 4294967295, F2: 1}
                switch "\n\t" {
                case "hello":
                    v1 = 1.5
                default:
                    if (!(128 >= 42)) {
                        v5 -= 1e10
                    }
                }
                _ = v35
            }
        default:
            switch 128 {
            case -9223372036854775808:
                v2 = v2
            case 108:
                switch "hello" {
                case "a":
                    v1 += float32(float64(v5))
                case "\n\t":
                    v5 = 1.5
                default:
                    v36 := GenS3{F0: -1}
                    _ = v36
                }
            case 1:
                v3["\n\t"] = "foo bar"
            default:
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1, v1)
                case "a":
                    fmt.Println(v5)
                case "foo bar":
                    var v37 uint64 = 58
                    _ = v37
                }
            }
        }
    }
    v38 := v3["hello"]
    for fi39 := 0; fi39 < 16; fi39++ {
        v38 = ((("hello" + v38) + v38) + "\n\t")
        _ = fi39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v2 = v2
        _ = fi40
    }
    if v2 {
        v41 := GenS2{F0: 722, F1: -1, F2: 42}
        _ = v41
    } else {
        var v42 uint16 = 255
        var v43 bool = (v2 && v2)
        _ = v42
        _ = v43
    }
    v1 = (((1.5 / ((v1 / ((3.14159 * (float32((-2.25 / (1e10 - 0.0))) - (v1 * (v1 - (v1 / v1))))) + v1)) - 1.5)) * (v1 + v1)) * (float32(float64((float32(3.14159) * v1))) - float32((1.5 / float64(v5)))))
    v44 := GenS4{F0: 1, F1: 695}
    for fi45 := 0; fi45 < 2; fi45++ {
        v46 := GenS5{F0: -9223372036854775808, F1: 65535}
        for fi47 := 0; fi47 < 2; fi47++ {
            v2 = v2
            _ = fi47
        }
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 float32 = float32(((-2.25 / 1e10) / (0.1 - (float64(-2.25) + 0.0))))
            _ = fi48
            _ = v49
        }
        _ = fi45
        _ = v46
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 int = fi50
        switch v38 {
        case "hello":
            if true {
                var v52 uint = 42
                if v2 {
                    switch (fi50 >> (((v52 + v52)) & 63)) {
                    case -1:
                        v53 := GenS5{F0: 0, F1: 42}
                        _ = v53
                    case 128:
                        var v54 uint64 = 42
                        _ = v54
                    case 0:
                        v38 = ((v38 + v38) + "")
                    }
                    fmt.Println(v1)
                } else {
                    if false {
                        v51 -= (fi50 ^ 2147483647)
                        var v55 uint16 = 32767
                        _ = v55
                    }
                    var v56 int8 = 42
                    _ = v56
                }
                v1 += v1
                _ = v52
            } else {
                if (1 > -1) {
                    if false {
                        v57 := GenS1{F0: false, F1: 978, F2: 77}
                        switch 945 {
                        case 0:
                            for fi58 := 0; fi58 < 2; fi58++ {
                                _ = len(v3)
                                v5 = 1e10
                                _ = fi58
                            }
                        default:
                            var v59 uint = 255
                            _ = v59
                        }
                        _ = v57
                    } else {
                        for fi60 := 0; fi60 < 15; fi60++ {
                            v61 := GenS6{F0: 681, F1: "a"}
                            _ = fi60
                            _ = v61
                        }
                    }
                    switch (v51 >> ((65535) & 63)) {
                    case 128:
                        v1 = float32((1.5 / (0.0 + (float64((v1 * (v5 - (-2.25 + (0.0 - v5))))) - ((float64(v1) + ((-2.25 - -2.25) * 1.5)) - 1e10)))))
                    case -793:
                        var v62 uint = 1
                        _ = v62
                    case 32767:
                        v63 := GenS6{F0: 1, F1: v38}
                        _ = v63
                    }
                    var v64 uint = 65535
                    _ = v64
                }
            }
        default:
            switch ((("a" + v38) + v38) + v38) {
            case "":
                v65 := GenS7{F0: -1}
                _ = v65
            case "foo bar":
                switch -563 {
                case -9223372036854775808:
                    if (v2 && (!(670 != 0))) {
                        for fi66 := 0; fi66 < 2; fi66++ {
                            _ = len(v4)
                            v1 = v1
                            _ = fi66
                        }
                        for fi67 := 0; fi67 < 19; fi67++ {
                            _ = len(v3)
                            _ = fi67
                        }
                    }
                case 1:
                    if (1 >= 2) {
                        for fi68 := 0; fi68 < 3; fi68++ {
                            var v69 float64 = (1e10 - 1e10)
                            if ((0.1 / 0.1) != v69) {
                                if len(v4) > 0 {
                                    v4[0] = ("a" + v38)
                                }
                                if (v38 <= (v38 + "\n\t")) {
                                    if true {
                                        if false {
                                            switch "hello" {
                                            case "a":
                                                var v70 uint16 = 32767
                                                _ = v70
                                            case "":
                                                var v71 int32 = 0
                                                _ = v71
                                            }
                                        }
                                        var v72 int8 = -49
                                        v69 += 1e10
                                        _ = v72
                                    }
                                    var v73 int = (fi50 << ((0) & 63))
                                    delete(v3, "\n\t")
                                    v38 = v38
                                    _ = v73
                                }
                            }
                            _ = fi68
                            _ = v69
                        }
                        if (32767 > -107) {
                            var v74 uint32 = 4294967295
                            if true {
                                v74 = v74
                                _ = len(v3)
                                var v75 int64 = 42
                                _ = v75
                            }
                            _ = v74
                        }
                    } else {
                        var v76 float32 = (v1 + -2.25)
                        v51 += -558
                        _ = v76
                    }
                default:
                    if true {
                        v1 = v1
                    } else {
                        var v77 uint32 = 2
                        for fi78 := 0; fi78 < 6; fi78++ {
                            v2 = (!(42 >= 0))
                            _ = fi78
                        }
                        _ = v77
                    }
                }
            }
        }
        _ = fi50
        _ = v51
    }
    var v79 uint16 = 42
    switch (v38 + (("\n\t" + v38) + (v38 + v38))) {
    case "a":
        fmt.Println(v38, v38, v2)
    case "hello":
        switch -1 {
        case -722:
            v1 = 0.1
        case 2147483647:
            for fi80 := 0; fi80 < 9; fi80++ {
                v2 = ((42 >= 701) && v2)
                _ = len(v3)
                _ = fi80
            }
        case 9223372036854775807:
            var v81 uint = 4294967295
            _ = v81
        default:
            if v2 {
                v1 = v5
                v82 := GenS8{F0: 1, F1: -1, F2: 127}
                for fi83 := 0; fi83 < 12; fi83++ {
                    v2 = (false && false)
                    _ = fi83
                }
                _ = v82
            } else {
                if v2 {
                    v2 = (58 > 135)
                }
            }
        }
    case "":
        for fi84 := 0; fi84 < 4; fi84++ {
            for fi85 := 0; fi85 < 16; fi85++ {
                switch v38 {
                case "\n\t":
                    var v86 int32 = 894
                    _ = v86
                }
                v4 = append(v4, v38)
                _ = fi85
            }
            _ = fi84
        }
    default:
        for fi87 := 0; fi87 < 19; fi87++ {
            var v88 bool = v2
            v89 := GenS6{F0: -118, F1: v38}
            v79 += v79
            _ = fi87
            _ = v88
            _ = v89
        }
    }
    v90 := GenS9{F0: "hello"}
    v91 := GenS10{F0: 97, F1: 2147483648, F2: 1}
    for fi92 := 0; fi92 < 16; fi92++ {
        fmt.Println(v79)
        var v93 uint64 = 42
        switch fi92 {
        case 32767:
            for fi94 := 0; fi94 < 3; fi94++ {
                v3[v38] = (v38 + ((v38 + "foo bar") + (v38 + "")))
                v93 = v93
                _ = fi94
            }
        case -292:
            for fi95 := 0; fi95 < 10; fi95++ {
                v3[("a" + "foo bar")] = v38
                v5 += (1.5 / (-2.25 + 1e10))
                v38 = "a"
                _ = fi95
            }
        default:
            fmt.Println(v1)
        }
        _ = fi92
        _ = v93
    }
    switch -2147483648 {
    case -1:
        v96 := GenS9{F0: v38}
        _ = v96
    default:
        var v97 uint32 = 0
        _ = v97
    }
    v2 = (-312 == 634)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v79
    _ = v90
    _ = v91
}
