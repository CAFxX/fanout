package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS10 struct { F0 string }
func genFn0(p0 string, p1 int64, p2 bool) bool {
    var v6 int = 42
    var v7 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}
func genFn1(p0 int64) int {
    var v8 bool = false
    v9 := GenS1{F0: -32768}
    var v10 uint64 = 4294967295
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := []int{32767, fi11}
        var v13 uint32 = 42
        if false {
            if (v13 >= (v13 * v13)) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v8 = (42 > 162)
                    v15 := GenS2{F0: 1, F1: ("\n\t" + "")}
                    if (!false) {
                        if v8 {
                            var v16 uint32 = v13
                            _ = v16
                        } else {
                            v17 := GenS2{F0: 32767, F1: (("foo bar" + "hello") + "a")}
                            v8 = false
                            _ = v17
                        }
                        v13 = 65535
                    }
                    _ = fi14
                    _ = v15
                }
                v18 := GenS3{F0: 4294967295, F1: "", F2: 398}
                _ = v18
            } else {
                var v19 uint64 = (^(v10))
                _ = len(v12)
                _ = v19
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                if false {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    var v20 int16 = 666
                    _ = v20
                } else {
                    switch fi11 {
                    case 0:
                        if v8 {
                            p0 = (^(p0))
                        } else {
                            if v8 {
                                v21 := GenS1{F0: 0}
                                p0 = -1
                                v22 := GenS1{F0: 1}
                                _ = v21
                                _ = v22
                            }
                        }
                    case 961:
                        for fi23 := 0; fi23 < 11; fi23++ {
                            if (42 > 2) {
                                var v24 uint = 161
                                _ = v24
                            } else {
                                v25 := GenS4{F0: "hello", F1: v13}
                                _ = v25
                            }
                            fmt.Println(fi23, fi11)
                            _ = fi23
                        }
                    }
                }
            case "":
                if (v13 != 2147483648) {
                    fmt.Println(v10, v13)
                    switch (fi11 % (((fi11 + fi11)) | 1)) {
                    case 170:
                        var v26 uint64 = 4294967295
                        _ = v26
                    }
                } else {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        fmt.Println(p0, fi27, fi27)
                        if (!v8) {
                            switch ((((("\n\t" + "") + "hello") + "foo bar") + "") + "\n\t") {
                            case "hello":
                                if v8 {
                                    v28 := GenS2{F0: 1, F1: "foo bar"}
                                    switch 42 {
                                    case 904:
                                        var v29 int = fi27
                                        _ = v29
                                    case 32767:
                                        v13 = 2147483647
                                    }
                                    v8 = true
                                    _ = v28
                                } else {
                                    if v8 {
                                        var v30 float64 = ((((3.14159 - (float64(0.1) / float64(1.5))) + 0.0) + -2.25) / 1.0)
                                        var v31 uint32 = v13
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if (!(!(!false))) {
                                            var v32 string = "foo bar"
                                            p0 -= int64(v10)
                                            fmt.Println(fi27, p0)
                                            var v33 uint32 = (v13 - (v13 % ((1) | 1)))
                                            v13 = 243
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    p0 = p0
                                }
                            case "a":
                                fmt.Println(p0, p0)
                            case "foo bar":
                                switch "a" {
                                case "a":
                                    if v8 {
                                        switch "" {
                                        case "a":
                                            v34 := GenS1{F0: -117}
                                            _ = v34
                                        case "hello":
                                            if true {
                                                var v35 int8 = 42
                                                _ = v35
                                            } else {
                                                if v8 {
                                                    var v36 uint8 = 2
                                                    _ = v36
                                                }
                                                v37 := GenS3{F0: v10, F1: (("" + ("a" + "a")) + "foo bar"), F2: 65535}
                                                fmt.Println(fi11)
                                                fmt.Println(fi11, p0, fi27)
                                                _ = v37
                                            }
                                        }
                                        v38 := GenS1{F0: int16(fi27)}
                                        if v8 {
                                            switch ("" + "foo bar") {
                                            case "":
                                                if v8 {
                                                    if (fi11 < (fi11 & (fi27 << ((65535) & 63)))) {
                                                        fmt.Println(fi27, fi27)
                                                        fmt.Println(v13, v13, v8)
                                                    } else {
                                                        fmt.Println(v10, fi27, v8)
                                                        var v39 bool = true
                                                        if (!(!(0 != 1))) {
                                                            switch (fi11 & fi27) {
                                                            case 171:
                                                                v40 := GenS1{F0: -32768}
                                                                _ = v40
                                                            default:
                                                                switch "foo bar" {
                                                                case "":
                                                                    if true {
                                                                        if (false && (!false)) {
                                                                            var v41 string = "foo bar"
                                                                            _ = v41
                                                                        }
                                                                        fmt.Println(p0, v39, v13)
                                                                        v39 = v8
                                                                        fmt.Println(fi27)
                                                                    }
                                                                case "hello":
                                                                    if (2 != v10) {
                                                                        v42 := GenS5{F0: 1, F1: v13}
                                                                        v10 -= 1
                                                                        switch 42 {
                                                                        case 127:
                                                                            var v43 uint8 = 201
                                                                            _ = v43
                                                                        case 12:
                                                                            fmt.Println(fi11)
                                                                        case -2147483648:
                                                                            fmt.Println(v39, p0, v39)
                                                                        }
                                                                        _ = v42
                                                                    }
                                                                case "a":
                                                                    fmt.Println(fi27, v13)
                                                                }
                                                                var v44 int = 1
                                                                v13 = v13
                                                                _ = v44
                                                            }
                                                            var v45 uint8 = 127
                                                            v46 := GenS3{F0: v10, F1: "", F2: 0}
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            if false {
                                                                v47 := GenS1{F0: -32768}
                                                                _ = v47
                                                            } else {
                                                                v10 = 2
                                                                v48 := GenS5{F0: 32767, F1: uint32(v10)}
                                                                switch "" {
                                                                case "\n\t":
                                                                    v10 -= uint64(v13)
                                                                case "hello":
                                                                    p0 = (-(p0))
                                                                default:
                                                                    switch "" {
                                                                    case "":
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v49 uint32 = 89
                                                                            _ = v49
                                                                        case "":
                                                                            var v50 int8 = -1
                                                                            _ = v50
                                                                        case "a":
                                                                            v8 = ((((((-2.25 - float32((float64(0.0) * 1e10))) / 3.14159) * (1.5 + 3.14159)) - float32(((1e10 / (1.5 + (-2.25 - float64(1.5)))) - (1e10 * 1.5)))) == (3.14159 - (3.14159 * float32(1.5)))) && ((!(0 > 65535)) || (!false)))
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi27, p0)
                                                                    }
                                                                }
                                                                _ = v48
                                                            }
                                                            var v51 float32 = 0.1
                                                            p0 = (^(p0))
                                                            _ = v51
                                                        }
                                                        var v52 uint8 = 42
                                                        _ = v39
                                                        _ = v52
                                                    }
                                                }
                                            case "hello":
                                                if len(v12) > 0 {
                                                    v12[0] = fi11
                                                }
                                            case "\n\t":
                                                fmt.Println(v10, p0, fi11)
                                            default:
                                                v8 = true
                                            }
                                            _ = len(v12)
                                        } else {
                                            var v53 float32 = ((0.0 * 1e10) - 1.5)
                                            if v8 {
                                                v53 -= (float32(((3.14159 - -2.25) + (float64(v53) * 0.1))) * 3.14159)
                                                _ = len(v12)
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v54 string = ("foo bar" + ("foo bar" + "a"))
                                                    _ = v54
                                                }
                                            }
                                            fmt.Println(fi11, v13, fi27)
                                            _ = v53
                                        }
                                        _ = v38
                                    }
                                case "foo bar":
                                    v55 := GenS3{F0: v10, F1: "", F2: 2}
                                    _ = v55
                                case "":
                                    if (!true) {
                                        v56 := GenS2{F0: -956, F1: "foo bar"}
                                        var v57 uint16 = uint16(v10)
                                        _ = v56
                                        _ = v57
                                    }
                                default:
                                    v58 := GenS6{F0: 255}
                                    _ = v58
                                }
                            }
                            v12 = append(v12, fi11)
                            switch "a" {
                            case "foo bar":
                                switch "" {
                                case "foo bar":
                                    var v59 float32 = 1.5
                                    _ = v59
                                default:
                                    var v60 uint16 = 422
                                    v13 += uint32(v10)
                                    fmt.Println(v13)
                                    _ = v60
                                }
                            default:
                                var v61 int32 = -1
                                _ = v61
                            }
                            var v62 int64 = 457
                            _ = v62
                        }
                        _ = fi27
                    }
                    if (436 == v10) {
                        if (((v13 ^ 2147483647) > (v13 << ((65535) & 63))) && false) {
                            fmt.Println(v10, fi11)
                            v63 := GenS3{F0: (v10 | v10), F1: "\n\t", F2: 668}
                            _ = v63
                        }
                        if v8 {
                            v13 -= 0
                            fmt.Println(v8, p0)
                            var v64 string = ""
                            _ = v64
                        } else {
                            for fi65 := 0; fi65 < 2; fi65++ {
                                var v66 int16 = 0
                                v8 = (2147483647 > 1)
                                p0 += -4
                                v10 = v10
                                _ = fi65
                                _ = v66
                            }
                            v12 = append(v12, 2147483647)
                            var v67 int64 = -9223372036854775808
                            _ = v67
                        }
                    } else {
                        v8 = (4294967295 == 255)
                        v8 = ((v13 * 42) <= 2147483648)
                    }
                }
            default:
                v68 := GenS6{F0: 255}
                _ = v68
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 554
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -891
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 float64 = v1
    v69 := GenS7{F0: -1, F1: 2147483647}
    v70 := GenS8{F0: 42}
    _ = len(v4)
    _ = len(v3)
    fmt.Println(v5, v5, v2)
    v1 = ((0.1 * -2.25) + (v1 + v5))
    v71 := []int{-9223372036854775808, 135, -9223372036854775808, 32767}
    var v72 bool = (!true)
    v73 := GenS8{F0: 2147483647}
    if len(v71) > 0 {
        v71[0] = 936
    }
    var v74 uint16 = 255
    fmt.Println(v1, v2, v1)
    var v75 uint = 1
    v76 := GenS8{F0: 9223372036854775807}
    switch (("" + ("foo bar" + "")) + "hello") {
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v77 uint = (v75 | (v75 / ((v75) | 1)))
            _ = v77
        case "foo bar":
            delete(v4, 0)
        case "hello":
            for fi78 := 0; fi78 < 16; fi78++ {
                v79 := GenS4{F0: ((("a" + "hello") + "a") + ("hello" + ("hello" + ("\n\t" + (("foo bar" + ("foo bar" + "foo bar")) + "foo bar"))))), F1: 531}
                _ = fi78
                _ = v79
            }
        }
    case "":
        for fi80 := 0; fi80 < 13; fi80++ {
            v75 = (^(v75))
            _ = fi80
        }
    case "foo bar":
        v2 -= 0
    }
    if true {
        v81 := GenS9{F0: (!(-1 >= 9)), F1: v74, F2: 213}
        v82 := GenS10{F0: ("foo bar" + "foo bar")}
        fmt.Println(v72)
        _ = v81
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
