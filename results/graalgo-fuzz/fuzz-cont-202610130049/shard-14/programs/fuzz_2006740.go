package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 uint32, p1 int32) int16 {
    var v7 uint32 = 0
    p1 = p1
    if ((!true) || ((p0 >> ((255) & 63)) <= 712)) {
        if (!false) {
            var v8 int32 = (p1 << ((255) & 63))
            _ = v8
        } else {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: -1}
                v11 := []int{(fi9 + fi9)}
                _ = fi9
                _ = v10
                _ = v11
            }
            if false {
                v7 -= (p0 + 4294967295)
                v12 := []int{9223372036854775807, -2147483648}
                _ = v12
            }
            v13 := GenS2{F0: p1}
            _ = v13
        }
        var v14 int32 = (^(p1))
        v15 := []int{217, 0}
        _ = v14
        _ = v15
    }
    v16 := []int{32768, 32767, -540}
    v7 = v7
    _ = p0
    _ = p1
    _ = v16
    _ = v7
    return 32767
}

func main() {
    var v1 uint = 42
    var v2 float32 = float32(0.1)
    var v3 uint = (v1 & v1)
    var v4 int64 = -9223372036854775808
    v5 := []int{0, 127, 0}
    var v6 uint32 = 2147483648
    var v17 int32 = 2147483647
    _ = len(v5)
    switch ("a" + ("" + "")) {
    case "":
        v18 := GenS2{F0: (v17 | 42)}
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if false {
            if true {
                v1 = v3
                for fi20 := 0; fi20 < 7; fi20++ {
                    if false {
                        _ = len(v5)
                    } else {
                        if true {
                            if len(v5) > 0 {
                                v5[0] = fi19
                            }
                            if (3.14159 == 3.14159) {
                                var v21 uint32 = (^(v6))
                                fmt.Println(fi20, v17)
                                _ = v21
                            }
                        }
                        if len(v5) > 0 {
                            v5[0] = fi19
                        }
                        fmt.Println(fi20, v17)
                        v4 += (v4 % (((v4 * 9223372036854775807)) | 1))
                        var v22 string = ("" + ("" + "\n\t"))
                        _ = v22
                    }
                    _ = fi20
                }
                if false {
                    _ = len(v5)
                }
            }
            v2 = ((((v2 * (v2 - ((0.1 - -2.25) / 1.0))) / 3.14159) + v2) + -2.25)
            switch "hello" {
            case "a":
                v2 -= v2
            case "\n\t":
                if true {
                    v3 = (v1 % ((1) | 1))
                    _ = len(v5)
                    fmt.Println(v4, v4, v1)
                }
            }
        } else {
            v23 := GenS2{F0: v17}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 = 42
            _ = v23
        }
        v17 -= v17
        _ = fi19
    }
    var v24 bool = (((1 <= 1) && (v4 == v4)) || (!true))
    var v25 float32 = v2
    var v26 uint8 = 2
    switch "foo bar" {
    case "hello":
        v27 := GenS1{F0: 366}
        _ = v27
    case "\n\t":
        v24 = v24
    case "":
        fmt.Println(v4, v26)
    }
    switch "hello" {
    case "a":
        v17 -= (v17 | -631)
    default:
        if (true && (v26 > (v26 & v26))) {
            var v28 float64 = (1e10 / 0.1)
            var v29 int8 = -1
            _ = v28
            _ = v29
        } else {
            switch "\n\t" {
            case "hello":
                for fi30 := 0; fi30 < 20; fi30++ {
                    v2 = float32((((1e10 + ((0.0 - float64((v2 - v2))) / -2.25)) + ((float64((1e10 / v25)) - (-2.25 * -2.25)) / (((0.1 - float64(((float32(0.1) * (-2.25 * v25)) - v25))) + float64(v2)) - 1.5))) * 3.14159))
                    v31 := GenS2{F0: 2147483647}
                    _ = fi30
                    _ = v31
                }
            case "a":
                if v24 {
                    var v32 string = (("\n\t" + "hello") + ("" + "foo bar"))
                    _ = v32
                }
            case "foo bar":
                var v33 bool = (982 != 2)
                _ = v33
            }
        }
    }
    switch "a" {
    case "a":
        for fi34 := 0; fi34 < 18; fi34++ {
            switch ((("" + ("\n\t" + "\n\t")) + ("foo bar" + "hello")) + "") {
            case "a":
                for fi35 := 0; fi35 < 8; fi35++ {
                    v6 = (v6 + 290)
                    v36 := GenS3{F0: v4, F1: (-(v17)), F2: -44}
                    var v37 int = (fi35 / (((fi34 - 0)) | 1))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            case "\n\t":
                var v38 uint32 = (v6 - (v6 - 65535))
                _ = v38
            case "":
                if false {
                    var v39 bool = v24
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 uint = (^(v3))
                        v39 = v24
                        switch "" {
                        case "foo bar":
                            if (v39 && v39) {
                                v24 = (v39 || true)
                                if v39 {
                                    var v42 float32 = float32((float64(v25) + (-2.25 * (0.1 / (0.0 - 3.14159)))))
                                    _ = v42
                                } else {
                                    v3 = 1
                                    fmt.Println(v4, v4, v26)
                                    v2 += v2
                                }
                            } else {
                                if (uint8(v6) < 255) {
                                    v43 := GenS4{F0: "", F1: v17}
                                    _ = v43
                                } else {
                                    if (v39 || v39) {
                                        v25 = float32(float64(v25))
                                    } else {
                                        var v44 int8 = 127
                                        var v45 string = ("a" + ("\n\t" + "hello"))
                                        var v46 uint64 = 597
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    v47 := GenS5{F0: v1, F1: (v17 | (v17 >> (((^(v3))) & 63)))}
                                    v48 := GenS6{F0: (v6 + (v6 / (((v6 * v6)) | 1))), F1: ""}
                                    if v24 {
                                        var v49 uint64 = 1
                                        v50 := GenS7{F0: v26, F1: v49}
                                        _ = v49
                                        _ = v50
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        switch int(v26) {
                                        case -9223372036854775808:
                                            v25 = v25
                                        case 477:
                                            v39 = v24
                                        case -133:
                                            var v51 uint8 = v26
                                            v6 -= 4294967295
                                            if v39 {
                                                v52 := GenS6{F0: v6, F1: "hello"}
                                                _ = v52
                                            } else {
                                                if v24 {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v53 := GenS6{F0: (v6 + (v6 << (((v3 / ((255) | 1))) & 63))), F1: ""}
                                                    if (float32((float64(((float32(-2.25) * v2) + 1e10)) - (3.14159 * 1.5))) >= v25) {
                                                        var v54 float64 = (-2.25 / 1.0)
                                                        _ = v54
                                                    } else {
                                                        var v55 int16 = 1000
                                                        fmt.Println(v6, v6)
                                                        v24 = (-2.25 < 0.1)
                                                        _ = v55
                                                    }
                                                    v56 := GenS8{F0: fi40, F1: 100}
                                                    _ = v53
                                                    _ = v56
                                                }
                                            }
                                            v57 := GenS9{F0: v17, F1: (fi40 * fi40)}
                                            switch (("foo bar" + ("a" + "\n\t")) + "\n\t") {
                                            case "":
                                                v24 = v39
                                                if (!v24) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v58 float32 = ((3.14159 / -2.25) * v25)
                                                    switch (^(fi34)) {
                                                    case -352:
                                                        var v59 uint = (v1 / ((v41) | 1))
                                                        _ = v59
                                                    }
                                                    _ = v58
                                                }
                                            case "a":
                                                if (953 == 4294967295) {
                                                    v26 -= v26
                                                    switch "hello" {
                                                    case "hello":
                                                        v60 := GenS10{F0: 597, F1: 255}
                                                        _ = v60
                                                    default:
                                                        switch (("a" + "foo bar") + "a") {
                                                        case "":
                                                            var v61 uint16 = 65535
                                                            _ = v61
                                                        case "a":
                                                            if ((!v24) || ((v4 / ((v4) | 1)) != (v4 << ((v41) & 63)))) {
                                                                var v62 uint64 = 2
                                                                _ = v62
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v25 = v25
                                            _ = v51
                                            _ = v57
                                        }
                                    case "\n\t":
                                        var v63 int = fi34
                                        _ = v63
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                switch "\n\t" {
                                case "":
                                    if v24 {
                                        switch fi34 {
                                        case 4294967295:
                                            var v64 uint16 = 0
                                            _ = v64
                                        case -1:
                                            v5 = append(v5, (fi40 ^ fi40))
                                            fmt.Println(v41, v24, fi34)
                                        case 789:
                                            v25 = 0.0
                                            fmt.Println(v3, v4, v17)
                                            v26 = 2
                                            v26 += v26
                                            v65 := GenS1{F0: genFn0(v6, int32(v41))}
                                            var v66 uint8 = 128
                                            _ = v65
                                            _ = v66
                                        default:
                                            v3 = 18446744073709551615
                                        }
                                        var v67 int32 = v17
                                        v39 = true
                                        _ = v67
                                    } else {
                                        v3 += v1
                                        switch "hello" {
                                        case "a":
                                            v17 = (v17 - (v17 + (v17 << ((v3) & 63))))
                                        case "\n\t":
                                            v3 = (v41 % (((v41 + 255)) | 1))
                                        default:
                                            var v68 uint8 = v26
                                            _ = v68
                                        }
                                    }
                                case "\n\t":
                                    var v69 float32 = (1e10 - (v25 + v25))
                                    _ = v69
                                case "foo bar":
                                    v1 = v3
                                }
                                if v24 {
                                    if false {
                                        var v70 float32 = (1e10 + (v25 + v2))
                                        _ = v70
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    if (79 <= v26) {
                                        v17 += (-(v17))
                                    }
                                    fmt.Println(v26, v1, v2)
                                    if (false || v24) {
                                        v71 := GenS2{F0: 2147483647}
                                        v4 = (v4 | 9223372036854775807)
                                        switch "" {
                                        case "hello":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            fmt.Println(fi34)
                                        }
                                        _ = v71
                                    }
                                    var v72 bool = v24
                                    v3 += v1
                                    _ = v72
                                }
                                var v73 float64 = float64((1.5 + float32(0.1)))
                                _ = v73
                            }
                        }
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                } else {
                    var v74 int8 = 1
                    _ = v74
                }
            default:
                v75 := GenS7{F0: v26, F1: 0}
                _ = v75
            }
            _ = fi34
        }
    case "hello":
        v76 := GenS11{F0: 630, F1: (v6 / ((2147483648) | 1))}
        _ = v76
    case "foo bar":
        if v24 {
            var v77 int32 = 1
            _ = v77
        } else {
            var v78 float64 = 1e10
            switch -1 {
            case 1:
                if len(v5) > 0 {
                    v5[0] = 32767
                }
            case 657:
                v79 := GenS12{F0: (^(v1)), F1: 440, F2: v4}
                _ = v79
            }
            _ = v78
        }
    default:
        if false {
            var v80 int = 1
            var v81 int64 = (v4 + v4)
            _ = v80
            _ = v81
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
