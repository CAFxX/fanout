package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 string; F1 uint8 }
func genFn0() int32 {
    v7 := []int{0, 128}
    var v8 float32 = (1.5 / 1.0)
    switch -1 {
    case 0:
        var v9 int8 = -119
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if true {
        v10 := GenS1{F0: 127}
        if len(v7) > 0 {
            v7[0] = 777
        }
        fmt.Println(v8)
        _ = v10
    }
    _ = v7
    _ = v8
    return -402
}

func main() {
    var v1 float64 = (-2.25 - 0.0)
    var v2 int8 = -128
    var v3 int64 = 128
    v4 := []string{(((("" + "hello") + "\n\t") + "a") + "")}
    v5 := []string{""}
    v6 := make(map[uint32]bool)
    v11 := GenS1{F0: 2}
    switch "hello" {
    case "hello":
        v12 := GenS1{F0: 0}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int32 = 2147483647
            if ((!true) || (-106 <= -1)) {
                if true {
                    v14 = v14
                    v3 = int64(v14)
                } else {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 float32 = 0.0
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 bool = true
                v18 := GenS2{F0: 100}
                _ = v17
                _ = v18
            } else {
                if (false && (420 < 42)) {
                    v3 -= (-(v3))
                    if false {
                        fmt.Println(v14, v1)
                        var v19 bool = (123 <= 65535)
                        _ = v19
                    } else {
                        switch "hello" {
                        case "":
                            if len(v4) > 0 {
                                v4[0] = (("foo bar" + "hello") + "")
                            }
                        }
                        fmt.Println(v14)
                    }
                } else {
                    var v20 int64 = (v3 / ((-767) | 1))
                    v3 += (^(v20))
                    var v21 int = fi13
                    _ = v20
                    _ = v21
                }
                var v22 uint16 = 42
                fmt.Println(fi13)
                _ = v22
            }
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v23 := GenS2{F0: v2}
        _ = v23
    }
    var v24 float32 = ((-2.25 - (-2.25 * float32(float64((1.5 * (1.5 / 3.14159)))))) + 0.1)
    v1 -= ((v1 + v1) * float64(v24))
    if (2147483647 < -1) {
        switch ("foo bar" + ("a" + ("\n\t" + "\n\t"))) {
        case "\n\t":
            v25 := GenS1{F0: 0}
            _ = v25
        case "a":
            switch 4294967295 {
            case 1:
                for fi26 := 0; fi26 < 1; fi26++ {
                    var v27 bool = (-1 < 1000)
                    fmt.Println(v24)
                    _ = len(v6)
                    _ = fi26
                    _ = v27
                }
            case 99:
                var v28 int64 = 1
                _ = v28
            }
        case "hello":
            if false {
                v2 = -1
                v2 = (v2 | v2)
            }
        default:
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := GenS3{F0: 979, F1: 2, F2: 0}
                var v31 float32 = v24
                v32 := GenS1{F0: 1}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
    }
    if true {
        switch "hello" {
        case "foo bar":
            v3 -= (v3 - v3)
        case "a":
            switch ("foo bar" + "\n\t") {
            case "a":
                if false {
                    v24 = (v24 - 0.0)
                } else {
                    var v33 int64 = v3
                    v24 += v24
                    var v34 string = "foo bar"
                    _ = v33
                    _ = v34
                }
            case "foo bar":
                var v35 uint16 = 42
                _ = v35
            default:
                if (65535 <= 2147483647) {
                    var v36 uint8 = 128
                    if true {
                        v1 += v1
                        if true {
                            v37 := GenS2{F0: (v2 + v2)}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v37
                        }
                    }
                    _ = v36
                } else {
                    v2 -= 42
                }
            }
        case "\n\t":
            v1 = v1
        }
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 uint8 = 163
        for fi40 := 0; fi40 < 1; fi40++ {
            v41 := GenS4{F0: (v3 / ((1) | 1))}
            if (50 < 255) {
                var v42 float32 = float32(v1)
                switch (fi38 ^ (fi40 / ((fi38) | 1))) {
                case 1:
                    v1 -= v1
                }
                v1 -= (v1 - float64((((0.1 / 1e10) - 1e10) / v24)))
                _ = v42
            } else {
                if false {
                    var v43 uint64 = 1
                    var v44 uint16 = 65535
                    v45 := GenS5{F0: 4294967295}
                    delete(v6, 2147483648)
                    if (int8(v3) >= (v2 ^ (v2 + 100))) {
                        var v46 uint64 = v43
                        if false {
                            fmt.Println(fi38)
                        } else {
                            var v47 int32 = 2147483647
                            var v48 int8 = v2
                            fmt.Println(v47, v3, v47)
                            if ((42 < 565) || (969 <= 2)) {
                                fmt.Println(v2)
                                v46 -= uint64(v39)
                                v6[2147483648] = (v48 < (^(v48)))
                            } else {
                                v39 -= (^(v39))
                                var v49 float32 = float32(v1)
                                var v50 uint8 = uint8(v43)
                                _ = v49
                                _ = v50
                            }
                            _ = v47
                            _ = v48
                        }
                        _ = v46
                    } else {
                        if (true && true) {
                            v51 := GenS5{F0: 42}
                            _ = v51
                        } else {
                            var v52 bool = ((!true) && true)
                            var v53 uint16 = uint16(v39)
                            _ = v52
                            _ = v53
                        }
                    }
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    fmt.Println(fi38)
                }
            }
            _ = fi40
            _ = v41
        }
        var v54 uint32 = 0
        _ = fi38
        _ = v39
        _ = v54
    }
    if (121 <= 1) {
        for fi55 := 0; fi55 < 1; fi55++ {
            switch "hello" {
            case "a":
                v56 := GenS6{F0: 42}
                _ = v56
            case "foo bar":
                v24 = ((0.0 * float32(v1)) / ((v24 + 0.0) * float32(v1)))
            case "\n\t":
                switch fi55 {
                case 32768:
                    var v57 float64 = 0.0
                    _ = v57
                case -1:
                    switch ((("foo bar" + ("a" + "a")) + "foo bar") + "a") {
                    case "":
                        _ = len(v5)
                    case "foo bar":
                        v6[486] = ((float64(v24) + float64(v24)) != (float64((float32(v1) * v24)) + 0.0))
                    case "\n\t":
                        switch 32768 {
                        case 0:
                            v1 += -2.25
                        default:
                            v58 := GenS7{F0: 4294967295}
                            _ = v58
                        }
                    default:
                        v24 -= v24
                    }
                case 9223372036854775807:
                    var v59 uint = 18446744073709551615
                    _ = v59
                }
            }
            _ = fi55
        }
        fmt.Println(v3)
    } else {
        if false {
            if false {
                v6[4294967295] = (3.14159 != -2.25)
                if (238 >= 9) {
                    var v60 int16 = 42
                    _ = v60
                } else {
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        var v61 float32 = (((v24 - -2.25) * ((((0.0 / v24) - v24) / 0.1) * -2.25)) * (0.1 + v24))
                        _ = v61
                    default:
                        for fi62 := 0; fi62 < 12; fi62++ {
                            v63 := GenS4{F0: (v3 | v3)}
                            v24 = 0.0
                            v1 = (v1 + float64((v24 * 3.14159)))
                            _ = fi62
                            _ = v63
                        }
                    }
                    switch 9223372036854775807 {
                    case 127:
                        for fi64 := 0; fi64 < 6; fi64++ {
                            v5 = append(v5, ("a" + "\n\t"))
                            _ = fi64
                        }
                    case 32767:
                        var v65 int = 4294967295
                        _ = v65
                    }
                }
            }
            v4 = append(v4, (("\n\t" + "hello") + "hello"))
        }
    }
    for fi66 := 0; fi66 < 18; fi66++ {
        for fi67 := 0; fi67 < 6; fi67++ {
            var v68 uint64 = 2
            if true {
                if (v1 == 1e10) {
                    switch (fi66 << ((0) & 63)) {
                    case -1:
                        switch 586 {
                        case 387:
                            v69 := GenS1{F0: 0}
                            _ = v69
                        case 0:
                            var v70 uint64 = v68
                            _ = v70
                        case -2147483648:
                            v24 = v24
                        }
                    }
                    v1 += (0.0 / float64(v24))
                    if (182 <= 2) {
                        if ((("a" + "") + "\n\t") < "foo bar") {
                            v6[448] = (((!((v68 ^ v68) <= v68)) || true) && false)
                            v24 += v24
                            var v71 int16 = 875
                            v72 := GenS7{F0: 0}
                            _ = v71
                            _ = v72
                        } else {
                            fmt.Println(fi67, v2)
                            _ = len(v5)
                        }
                    } else {
                        fmt.Println(v24)
                        _ = len(v6)
                        if (!false) {
                            var v73 string = ((("hello" + ("foo bar" + "a")) + ("hello" + "")) + "hello")
                            v68 -= (v68 / ((v68) | 1))
                            var v74 uint64 = (v68 & (v68 & 397))
                            _ = v73
                            _ = v74
                        } else {
                            v68 += (v68 * v68)
                            v2 += (v2 << ((445) & 63))
                            v75 := GenS7{F0: v68}
                            var v76 int = fi67
                            _ = v75
                            _ = v76
                        }
                        _ = len(v6)
                    }
                } else {
                    switch (fi66 ^ fi67) {
                    case 127:
                        var v77 int = -2147483648
                        if (v1 >= v1) {
                            v78 := GenS6{F0: 128}
                            var v79 float64 = 0.1
                            var v80 uint32 = 521
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            if true {
                                fmt.Println(v77, v1, fi66)
                            } else {
                                v3 = (v3 & 1)
                            }
                            v81 := GenS2{F0: v2}
                            var v82 uint = 65535
                            _ = v81
                            _ = v82
                        }
                        v24 = v24
                        _ = v77
                    }
                    var v83 uint32 = 42
                    if (false && true) {
                        var v84 bool = false
                        delete(v6, 248)
                        _ = v84
                    } else {
                        if false {
                            v85 := GenS8{F0: ((!(v68 == v68)) && true), F1: 1}
                            v24 += (v24 + v24)
                            v3 -= (v3 - (-(v3)))
                            _ = v85
                        }
                        v86 := GenS9{F0: v83, F1: 32767}
                        _ = v86
                    }
                    _ = v83
                }
            }
            if ((^(v2)) >= (v2 + 42)) {
                var v87 string = "a"
                _ = v87
            } else {
                v68 = v68
            }
            _ = fi67
            _ = v68
        }
        var v88 int8 = 42
        if false {
            switch "hello" {
            case "a":
                switch "hello" {
                case "":
                    var v89 uint64 = 2
                    _ = v89
                case "\n\t":
                    v3 += (-(v3))
                default:
                    if (65535 < 65535) {
                        var v90 int16 = 0
                        if (v3 <= -573) {
                            v91 := GenS10{F0: 42, F1: fi66}
                            _ = v91
                        } else {
                            for fi92 := 0; fi92 < 11; fi92++ {
                                v90 += -263
                                v90 = v90
                                v93 := GenS11{F0: "a", F1: 128}
                                _ = fi92
                                _ = v93
                            }
                            v24 += (float32((v1 * v1)) + (0.1 * ((0.1 + -2.25) + v24)))
                        }
                        _ = v90
                    }
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            var v94 uint64 = 561
            v94 = (v94 >> ((236) & 63))
            _ = v94
        } else {
            v88 = (v2 | (v88 | (^(v88))))
        }
        _ = fi66
        _ = v88
    }
    var v95 int32 = 653
    v96 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v95:", v95)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
    _ = v96
}
