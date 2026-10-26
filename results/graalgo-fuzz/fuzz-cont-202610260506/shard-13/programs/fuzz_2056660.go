package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int; F2 int }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int64; F1 int8; F2 string }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 int32 }
func genFn0(p0 int32) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        fmt.Println(p0, p0)
        p0 = (p0 % ((p0) | 1))
        if false {
            var v5 uint16 = 32767
            _ = v5
        } else {
            v6 := []int{}
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 314
    var v2 string = "\n\t"
    var v3 bool = (true || true)
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 = v2
        v1 -= (v1 & v1)
        v9 := GenS1{F0: fi8}
        _ = fi8
        _ = v9
    }
    if v3 {
        if v3 {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint64 = 4294967295
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := []int{(fi12 | (fi12 % (((-(fi12))) | 1))), fi12, (fi12 / (((-(fi12))) | 1))}
                fmt.Println(fi12, v2, v1)
                fmt.Println(v2, v1, v3)
                _ = fi12
                _ = v13
            }
            v3 = v3
        } else {
            v14 := []int{32767}
            _ = v14
        }
    }
    v3 = v3
    if (4294967295 >= 6) {
        v15 := []int{}
        v2 = v2
        if v3 {
            if ((v3 && (!v3)) && v3) {
                v2 = (("foo bar" + ((v2 + v2) + v2)) + (("hello" + "hello") + ""))
            }
            v2 = (v2 + (v2 + v2))
        }
        _ = v15
    } else {
        var v16 uint32 = 2147483648
        v17 := GenS1{F0: 9223372036854775807}
        v18 := GenS1{F0: 32767}
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(v2, v2, v2)
        if (128 != uint8(v1)) {
            for fi20 := 0; fi20 < 11; fi20++ {
                switch (^(fi19)) {
                case 127:
                    v21 := GenS1{F0: fi20}
                    _ = v21
                }
                switch fi20 {
                case 0:
                    v22 := GenS1{F0: fi20}
                    _ = v22
                }
                _ = fi20
            }
            var v23 int16 = 0
            _ = v23
        }
        v3 = v3
        _ = fi19
    }
    var v24 uint8 = 128
    v3 = (!false)
    var v25 int32 = 292
    v2 = "foo bar"
    var v26 bool = v3
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 string = (v2 + v2)
        _ = fi27
        _ = v28
    }
    if false {
        switch v2 {
        case "hello":
            switch 9223372036854775807 {
            case -1:
                var v29 float32 = (-2.25 + ((float32(1.5) / (-2.25 * ((0.1 / -2.25) - 0.1))) / (0.1 + 3.14159)))
                _ = v29
            case 0:
                v30 := GenS2{F0: ("foo bar" + "\n\t")}
                _ = v30
            case 9223372036854775807:
                v31 := []int{}
                _ = v31
            default:
                v32 := GenS3{F0: v24}
                _ = v32
            }
        case "foo bar":
            if (-32768 < 1) {
                v25 += (v25 | v25)
                switch 312 {
                case 824:
                    if ((v3 && (1 >= (v25 | 801))) || true) {
                        switch (v2 + "foo bar") {
                        case "a":
                            var v33 int32 = (v25 << ((42) & 63))
                            _ = v33
                        default:
                            if (!(1 < 32767)) {
                                v34 := []int{}
                                _ = v34
                            } else {
                                if v3 {
                                    fmt.Println(v1, v25, v24)
                                } else {
                                    v35 := []int{127, int(v24)}
                                    v35 = append(v35, 32767)
                                    _ = v35
                                }
                                switch -2147483648 {
                                case 128:
                                    switch 127 {
                                    case 42:
                                        v36 := GenS4{F0: -347}
                                        _ = v36
                                    case -2147483648:
                                        if ((v24 % ((v24) | 1)) <= uint8(v1)) {
                                            if (1.5 > float64(0.1)) {
                                                if ((-32768 < 1) && (-1 != 42)) {
                                                    var v37 uint32 = 65535
                                                    var v38 int = 9223372036854775807
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v39 := GenS1{F0: 37}
                                                _ = v39
                                            } else {
                                                switch v2 {
                                                case "foo bar":
                                                    var v40 float32 = 0.1
                                                    _ = v40
                                                case "\n\t":
                                                    v41 := GenS4{F0: 0}
                                                    _ = v41
                                                case "a":
                                                    v2 = v2
                                                }
                                                var v42 float32 = ((3.14159 + (((((3.14159 / 1e10) + (float32((((3.14159 - 1.5) + 1.5) - float64(((0.1 / 1.0) + 1.5)))) / 1e10)) / (float32((1e10 + 1.5)) - 3.14159)) + 0.1) * (-2.25 - -2.25))) / 1e10)
                                                switch 32768 {
                                                case 462:
                                                    fmt.Println(v1, v42, v42)
                                                case 2147483647:
                                                    v43 := []int{-832}
                                                    _ = v43
                                                case 42:
                                                    var v44 int = 882
                                                    _ = v44
                                                }
                                                _ = v42
                                            }
                                        }
                                    }
                                case 32768:
                                    switch ((v2 + v2) + v2) {
                                    case "hello":
                                        v45 := []int{32767, 0}
                                        _ = v45
                                    case "a":
                                        for fi46 := 0; fi46 < 17; fi46++ {
                                            var v47 uint32 = 14
                                            _ = fi46
                                            _ = v47
                                        }
                                    case "\n\t":
                                        for fi48 := 0; fi48 < 17; fi48++ {
                                            v2 = (v2 + "hello")
                                            _ = fi48
                                        }
                                    }
                                case -811:
                                    v49 := GenS3{F0: (v24 >> ((65535) & 63))}
                                    _ = v49
                                }
                            }
                        }
                        for fi50 := 0; fi50 < 20; fi50++ {
                            fmt.Println(v25, v24)
                            var v51 float64 = 0.0
                            if true {
                                v24 = (v24 | (^(v24)))
                                var v52 float64 = float64(float32(v51))
                                _ = v52
                            } else {
                                v53 := []int{782, fi50, 0, 128}
                                if len(v53) >= 2 {
                                    v53 = v53[1:len(v53)]
                                }
                                _ = v53
                            }
                            _ = fi50
                            _ = v51
                        }
                        v54 := []int{-2147483648, -2147483648, 2147483647, -7}
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 7; fi55++ {
                            switch fi55 {
                            case 6:
                                for fi56 := 0; fi56 < 3; fi56++ {
                                    v24 -= v24
                                    if (true && true) {
                                        fmt.Println(fi56)
                                        v57 := GenS5{F0: v1, F1: fi55, F2: (fi55 % ((fi55) | 1))}
                                        _ = v57
                                    }
                                    if v26 {
                                        v2 = v2
                                        switch "a" {
                                        case "":
                                            var v58 int64 = -9223372036854775808
                                            _ = v58
                                        default:
                                            if true {
                                                if false {
                                                    if (!v3) {
                                                        v59 := GenS6{F0: 0, F1: (((v2 + ((v2 + "a") + v2)) + ("a" + "")) + v2)}
                                                        _ = v59
                                                    }
                                                    if v3 {
                                                        var v60 int32 = (v25 + v25)
                                                        v61 := []int{(fi56 & (fi56 * (fi56 / (((fi55 % ((4294967295) | 1))) | 1)))), (fi56 / ((fi56) | 1)), (^(fi55))}
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        if v3 {
                                                            var v62 string = v2
                                                            _ = v62
                                                        }
                                                        fmt.Println(fi56)
                                                    }
                                                    fmt.Println(v24, v24)
                                                }
                                                switch (fi56 | (fi56 << ((914) & 63))) {
                                                case 928:
                                                    if v3 {
                                                        var v63 string = v2
                                                        _ = v63
                                                    }
                                                default:
                                                    if v3 {
                                                        fmt.Println(fi56, v24)
                                                        var v64 int64 = 1
                                                        _ = v64
                                                    } else {
                                                        switch 127 {
                                                        case 2147483647:
                                                            v65 := GenS6{F0: 1, F1: (v2 + "a")}
                                                            _ = v65
                                                        default:
                                                            if (!(1e10 >= 1e10)) {
                                                                fmt.Println(v26, v24, v2)
                                                            }
                                                        }
                                                        v66 := GenS7{F0: (v24 & v24), F1: fi55}
                                                        _ = v66
                                                    }
                                                }
                                                v1 += 4294967295
                                            } else {
                                                var v67 int32 = (-(v25))
                                                v68 := []int{}
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    }
                                    _ = fi56
                                }
                            case -5:
                                fmt.Println(v3, v1)
                            }
                            v69 := GenS2{F0: "hello"}
                            v2 = v2
                            _ = fi55
                            _ = v69
                        }
                        v25 -= 2147483647
                    }
                case -948:
                    var v70 uint = 0
                    _ = v70
                case -2147483648:
                    for fi71 := 0; fi71 < 8; fi71++ {
                        for fi72 := 0; fi72 < 10; fi72++ {
                            switch v2 {
                            case "":
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    switch "a" {
                                    case "":
                                        var v73 int = (-(fi71))
                                        _ = v73
                                    case "foo bar":
                                        v74 := GenS8{F0: 686, F1: 100, F2: (((v2 + ("a" + v2)) + "hello") + "a")}
                                        _ = v74
                                    case "\n\t":
                                        var v75 bool = (!(127 > 42))
                                        _ = v75
                                    default:
                                        if (!v26) {
                                            v76 := []int{9223372036854775807, fi72, (fi71 * (fi71 ^ -9223372036854775808))}
                                            _ = v76
                                        }
                                    }
                                case 392:
                                    var v77 int8 = 100
                                    _ = v77
                                }
                            case "a":
                                v78 := GenS3{F0: v24}
                                _ = v78
                            default:
                                v79 := GenS2{F0: v2}
                                _ = v79
                            }
                            v80 := GenS9{F0: -93, F1: (("" + ("\n\t" + v2)) <= ("foo bar" + (((v2 + ((("" + (v2 + ((("hello" + ("hello" + v2)) + ((("a" + v2) + "\n\t") + (((v2 + "\n\t") + "hello") + "a"))) + ((v2 + "hello") + v2)))) + v2) + v2)) + v2) + (v2 + (("a" + "\n\t") + v2)))))}
                            fmt.Println(v2, v25)
                            if v3 {
                                v2 = "a"
                            }
                            _ = fi72
                            _ = v80
                        }
                        if (((v25 & v25) >= v25) || v26) {
                            v24 = (v24 - 14)
                            var v81 int16 = -1
                            _ = v81
                        }
                        switch -2147483648 {
                        case -417:
                            var v82 int32 = v25
                            _ = v82
                        case -2147483648:
                            fmt.Println(v2, v24)
                        case 32767:
                            fmt.Println(v24, fi71)
                        }
                        _ = fi71
                    }
                default:
                    v1 = (^(v1))
                }
                v83 := []int{}
                _ = v83
            } else {
                var v84 int8 = 1
                var v85 uint64 = 223
                _ = v84
                _ = v85
            }
        case "\n\t":
            var v86 string = v2
            _ = v86
        default:
            for fi87 := 0; fi87 < 2; fi87++ {
                v3 = ((0 >= 851) || (true && false))
                if (!true) {
                    var v88 float64 = float64(-2.25)
                    for fi89 := 0; fi89 < 13; fi89++ {
                        var v90 int32 = v25
                        v91 := []int{fi89, 0}
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    v92 := []int{(fi87 / (((fi87 / (((fi87 >> ((uint(v1)) & 63))) | 1))) | 1)), -474, (fi87 << ((1) & 63)), 42}
                    _ = v88
                    _ = v92
                } else {
                    for fi93 := 0; fi93 < 14; fi93++ {
                        fmt.Println(v2, v3)
                        fmt.Println(fi93, v24, fi87)
                        _ = fi93
                    }
                    v94 := []int{2147483647, (fi87 | (fi87 * -33)), fi87, fi87}
                    _ = v94
                }
                var v95 uint8 = v24
                _ = fi87
                _ = v95
            }
        }
        if true {
            if (v26 || ((!(591 != 14)) && (562 <= 194))) {
                v25 = v25
            }
        } else {
            v3 = (true && true)
            v96 := []int{4294967295, 128, 1}
            _ = v96
        }
        v97 := GenS10{F0: (v25 - (v25 / ((-606) | 1)))}
        _ = v97
    }
    for fi98 := 0; fi98 < 17; fi98++ {
        v99 := []int{int(v25), fi98}
        fmt.Println(v3, v1, v3)
        if true {
            v24 += v24
            var v100 float64 = -2.25
            fmt.Println(v100)
            _ = v100
        } else {
            v3 = v3
        }
        _ = fi98
        _ = v99
    }
    v3 = v26
    v24 -= v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
