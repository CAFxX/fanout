package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint16 }
func genFn0() int8 {
    var v5 int8 = 1
    if true {
        var v6 int8 = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v5 += (v5 / ((v5) | 1))
        if (!true) {
            var v8 int = (fi7 >> ((42) & 63))
            v5 = v5
            v9 := GenS1{F0: 0, F1: -9223372036854775808}
            _ = v8
            _ = v9
        }
        switch (-(fi7)) {
        case -2147483648:
            fmt.Println(v5)
        case 0:
            var v10 uint32 = 42
            _ = v10
        default:
            var v11 uint32 = 0
            _ = v11
        }
        _ = fi7
    }
    v5 += v5
    _ = v5
    return v5
}

func main() {
    v1 := []int{4294967295}
    v2 := []int{-1, 32767, 0, 159}
    v3 := make(map[string]uint32)
    var v4 int = -2147483648
    switch 0 {
    case 0:
        if false {
            if len(v1) > 0 {
                v1[0] = (-(v4))
            }
        } else {
            for fi12 := 0; fi12 < 12; fi12++ {
                switch "\n\t" {
                case "hello":
                    if false {
                        if true {
                            fmt.Println(v4, v4)
                            var v13 bool = (100 >= genFn0())
                            _ = v13
                        } else {
                            switch "a" {
                            case "\n\t":
                                if (true && true) {
                                    var v14 int32 = -1
                                    if (((((0.0 - (3.14159 - 1e10)) / 1.0) / 1.0) / -2.25) != -2.25) {
                                        for fi15 := 0; fi15 < 20; fi15++ {
                                            if false {
                                                var v16 uint32 = 706
                                                _ = v16
                                            }
                                            if (true || true) {
                                                switch fi15 {
                                                case 1:
                                                    fmt.Println(fi15, v14)
                                                case 0:
                                                    fmt.Println(v4, v4, v14)
                                                case 4294967295:
                                                    var v17 float64 = ((0.1 - ((1.5 / (((1e10 - (((float64(1.5) + (3.14159 * (((3.14159 - 0.0) - 3.14159) / (1e10 + float64(0.0))))) * 1.5) / 3.14159)) * float64(1e10)) * (((0.1 + (1e10 / 1.5)) / 1.0) - -2.25))) * -2.25)) * 1e10)
                                                    _ = v17
                                                default:
                                                    fmt.Println(v4, fi15)
                                                }
                                            } else {
                                                if (v4 >= (fi12 * (v4 - (fi15 >> ((255) & 63))))) {
                                                    v18 := GenS2{F0: "foo bar"}
                                                    if (32767 > 1000) {
                                                        var v19 int32 = v14
                                                        fmt.Println(fi15, fi12, fi15)
                                                        var v20 int8 = -95
                                                        fmt.Println(v20)
                                                        switch fi15 {
                                                        case 2147483647:
                                                            switch (-(fi12)) {
                                                            case 89:
                                                                var v21 float64 = float64(1.5)
                                                                _ = v21
                                                            case -469:
                                                                v22 := GenS1{F0: 1, F1: -9223372036854775808}
                                                                _ = v22
                                                            }
                                                        case 282:
                                                            v4 -= v4
                                                        case 32767:
                                                            v23 := v3[("hello" + "\n\t")]
                                                            _ = v23
                                                        default:
                                                            var v24 uint = 607
                                                            _ = v24
                                                        }
                                                        var v25 uint = 65535
                                                        _ = v19
                                                        _ = v20
                                                        _ = v25
                                                    }
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            if (65 <= 748) {
                                                                _ = len(v2)
                                                            } else {
                                                                switch "foo bar" {
                                                                case "":
                                                                    v14 += (v14 - -1)
                                                                }
                                                                v3[("a" + "hello")] = 931
                                                            }
                                                        case "\n\t":
                                                            if ((2 < 830) || false) {
                                                                v26 := GenS2{F0: "foo bar"}
                                                                v14 = v14
                                                                _ = v26
                                                            } else {
                                                                var v27 int64 = -474
                                                                _ = v27
                                                            }
                                                        case "hello":
                                                            v4 += fi15
                                                        default:
                                                            v4 = 701
                                                        }
                                                    }
                                                    _ = v18
                                                }
                                            }
                                            var v28 uint8 = 127
                                            _ = fi15
                                            _ = v28
                                        }
                                    } else {
                                        var v29 string = "a"
                                        v30 := GenS3{F0: 0, F1: 255, F2: 20}
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v14
                                }
                            }
                        }
                        fmt.Println(fi12, fi12)
                        var v31 float32 = 3.14159
                        _ = v31
                    }
                case "foo bar":
                    fmt.Println(fi12, fi12)
                case "a":
                    v32 := GenS3{F0: 32767, F1: 459, F2: 0}
                    _ = v32
                }
                var v33 float64 = 0.0
                _ = fi12
                _ = v33
            }
        }
    case 128:
        _ = len(v3)
    }
    if (42 <= 42) {
        v4 -= (^(v4))
        fmt.Println(v4)
    } else {
        if true {
            v34 := GenS4{F0: -32768}
            _ = v34
        }
        fmt.Println(v4)
    }
    v4 = (v4 / (((v4 / ((4294967295) | 1))) | 1))
    v4 = v4
    switch v4 {
    case -2147483648:
        fmt.Println(v4)
    }
    v4 = v4
    v35 := GenS5{F0: 1, F1: ((0.0 - (float32(-2.25) / 1e10)) < 1e10), F2: 400}
    var v36 float32 = 1e10
    if (65535 != 172) {
        v36 = float32(0.0)
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        fmt.Println(v4)
        switch "foo bar" {
        case "hello":
            if (!true) {
                for fi38 := 0; fi38 < 4; fi38++ {
                    fmt.Println(v4)
                    var v39 int = -666
                    _ = fi38
                    _ = v39
                }
            }
        case "":
            if (-32768 <= 599) {
                switch "\n\t" {
                case "":
                    for fi40 := 0; fi40 < 7; fi40++ {
                        if ("hello" != "a") {
                            v41 := GenS6{F0: (-(fi37))}
                            var v42 uint8 = 45
                            var v43 int16 = -1
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        v36 = (v36 - -2.25)
                        _ = fi40
                    }
                case "\n\t":
                    switch (("a" + (("foo bar" + ((((("" + "hello") + "a") + (("hello" + "") + "\n\t")) + ("" + "hello")) + "foo bar")) + "")) + (("\n\t" + "a") + "\n\t")) {
                    case "hello":
                        v44 := GenS1{F0: 127, F1: 0}
                        _ = v44
                    case "foo bar":
                        var v45 float64 = (1e10 / -2.25)
                        _ = v45
                    case "a":
                        switch "a" {
                        case "":
                            v46 := GenS1{F0: 255, F1: 0}
                            _ = v46
                        case "\n\t":
                            for fi47 := 0; fi47 < 4; fi47++ {
                                v4 += (v4 % (((-(v4))) | 1))
                                fmt.Println(v4, v4, v4)
                                var v48 uint = 743
                                if true {
                                    v49 := GenS1{F0: 42, F1: -1}
                                    if (2147483647 == 42) {
                                        if (1 <= 128) {
                                            var v50 bool = ((-2147483648 < -1) && true)
                                            _ = v50
                                        } else {
                                            var v51 uint16 = 65535
                                            v36 -= 0.1
                                            if (732 != 2147483647) {
                                                var v52 uint64 = 4294967295
                                                fmt.Println(v51, v36)
                                                _ = v52
                                            }
                                            _ = v51
                                        }
                                        _ = len(v1)
                                    } else {
                                        var v53 int = (fi47 >> ((v48) & 63))
                                        _ = v53
                                    }
                                    _ = v49
                                } else {
                                    fmt.Println(v36, v36)
                                    switch "hello" {
                                    case "\n\t":
                                        v36 = -2.25
                                    case "a":
                                        v54 := GenS7{F0: 795}
                                        _ = v54
                                    default:
                                        v48 += v48
                                    }
                                }
                                if ((fi37 >> ((v48) & 63)) < (fi37 << ((v48) & 63))) {
                                    v36 -= ((v36 - -2.25) - float32(1.5))
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v55 int32 = 2147483647
                                        _ = v55
                                    case "a":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "":
                                        var v56 int16 = 1
                                        _ = v56
                                    }
                                }
                                v4 = fi37
                                fmt.Println(v48, v48, v36)
                                var v57 int32 = 42
                                var v58 uint32 = 65535
                                _ = fi47
                                _ = v48
                                _ = v57
                                _ = v58
                            }
                        case "a":
                            if true {
                                v59 := GenS1{F0: 1, F1: 1}
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    fmt.Println(v36, fi60)
                                    fmt.Println(fi37)
                                    _ = fi60
                                }
                                _ = v59
                            }
                        }
                    }
                }
                for fi61 := 0; fi61 < 13; fi61++ {
                    v62 := v3[("hello" + "foo bar")]
                    fmt.Println(v62, v62, v4)
                    _ = fi61
                    _ = v62
                }
                v4 = fi37
            } else {
                var v63 bool = (!(v4 != v4))
                if (-54 <= 1) {
                    var v64 string = "foo bar"
                    _ = v64
                }
                if (4294967295 < 1) {
                    var v65 uint8 = 0
                    _ = v65
                } else {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        if ((v36 / v36) <= v36) {
                            fmt.Println(v36, v63)
                            var v67 int32 = 1
                            fmt.Println(v36)
                            _ = v67
                        }
                        fmt.Println(v36, v36)
                        _ = fi66
                    }
                    fmt.Println(v63, v4)
                    v68 := GenS8{F0: false, F1: 0}
                    _ = v68
                }
                _ = v63
            }
        case "a":
            var v69 uint8 = 0
            _ = v69
        }
        if (42 >= -9223372036854775808) {
            var v70 int64 = -132
            for fi71 := 0; fi71 < 9; fi71++ {
                var v72 float64 = 0.1
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v70 = v70
                v73 := GenS9{F0: 0}
                if false {
                    var v74 string = ("hello" + "foo bar")
                    fmt.Println(v72)
                    v75 := v3[v74]
                    _ = v74
                    _ = v75
                } else {
                    v76 := GenS7{F0: 2}
                    _ = v76
                }
                v36 += (1e10 - v36)
                _ = fi71
                _ = v72
                _ = v73
            }
            _ = v70
        } else {
            for fi77 := 0; fi77 < 3; fi77++ {
                if false {
                    var v78 uint64 = 4294967295
                    _ = v78
                }
                var v79 bool = (42 >= 0)
                fmt.Println(v36, v36, v79)
                _ = fi77
                _ = v79
            }
            for fi80 := 0; fi80 < 14; fi80++ {
                v36 = v36
                switch fi80 {
                case -2147483648:
                    if (-1 > 100) {
                        var v81 int32 = -359
                        v36 -= 3.14159
                        v4 = (fi80 * (fi80 * (v4 ^ 293)))
                        _ = v81
                    }
                default:
                    var v82 uint = 65535
                    _ = v82
                }
                _ = fi80
            }
            v36 = (3.14159 + (-2.25 * 3.14159))
        }
        _ = fi37
    }
    var v83 int64 = 1
    switch (v4 * v4) {
    case -630:
        v3[(("a" + "\n\t") + ("" + (("hello" + ("" + ("hello" + (("foo bar" + "foo bar") + "hello")))) + "")))] = 42
    case -2147483648:
        if true {
            for fi84 := 0; fi84 < 10; fi84++ {
                v36 += 1.5
                _ = fi84
            }
            var v85 uint32 = 0
            _ = v85
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v83
}
