package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int8; F2 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    var v3 float64 = 1.5
    v4 := make(map[int]uint32)
    v5 := []string{"hello", "a", "foo bar", ("\n\t" + "foo bar")}
    v2["\n\t"] = 9223372036854775807
    if false {
        if false {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                var v6 bool = false
                for fi7 := 0; fi7 < 7; fi7++ {
                    if v6 {
                        var v8 float64 = v3
                        _ = v8
                    }
                    _ = fi7
                }
                _ = v6
            }
            v1 += v1
            var v9 int = -9223372036854775808
            _ = v9
        }
        v3 = float64(float32(((0.1 * v3) / v3)))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if (731 < 94) {
        v10 := GenS1{F0: 42, F1: 1}
        delete(v2, "")
        _ = v10
    } else {
        if true {
            _ = len(v4)
            for fi11 := 0; fi11 < 4; fi11++ {
                switch (("hello" + "foo bar") + "hello") {
                case "":
                    _ = len(v4)
                case "hello":
                    for fi12 := 0; fi12 < 9; fi12++ {
                        if (-1 == -2147483648) {
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            v3 += v3
                            var v13 int16 = (v1 << ((990) & 63))
                            _ = v13
                        } else {
                            fmt.Println(v3, fi11)
                            v1 += v1
                        }
                        v1 = v1
                        _ = fi12
                    }
                default:
                    v14 := GenS2{F0: 335, F1: 162, F2: ""}
                    _ = v14
                }
                if false {
                    v15 := GenS3{F0: ("\n\t" + "foo bar")}
                    var v16 bool = false
                    var v17 float64 = 1.5
                    _ = v15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 float32 = (float32((v3 + 1.5)) - 0.0)
                    var v20 uint8 = 253
                    v21 := v2[("" + (("foo bar" + ("" + "a")) + ""))]
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi11
            }
            v3 = (v3 - (float64(-2.25) / 1.5))
        } else {
            for fi22 := 0; fi22 < 12; fi22++ {
                v1 -= (v1 * 102)
                var v23 uint16 = 527
                var v24 uint = 18446744073709551615
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    v3 = v3
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 bool = false
        var v27 float64 = (v3 - -2.25)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v3)
    if false {
        var v28 float32 = 0.0
        v4[934] = 520
        if (!(true && (!false))) {
            v3 = v3
        }
        _ = v28
    } else {
        v1 = 75
        v1 = v1
    }
    if true {
        var v29 float32 = 1.5
        _ = v29
    }
    if (0 == 18446744073709551615) {
        for fi30 := 0; fi30 < 5; fi30++ {
            switch fi30 {
            case -9223372036854775808:
                for fi31 := 0; fi31 < 18; fi31++ {
                    switch fi31 {
                    case 32768:
                        v3 = v3
                    }
                    v32 := GenS4{F0: 42, F1: 120, F2: 32767}
                    if (-1 <= -1) {
                        switch -9223372036854775808 {
                        case -1:
                            if (42 > 719) {
                                v1 = v1
                                switch (("" + "foo bar") + ((("hello" + "") + "a") + "hello")) {
                                case "\n\t":
                                    v3 = (3.14159 / v3)
                                case "foo bar":
                                    if (false || (-1 != -92)) {
                                        delete(v4, fi31)
                                    }
                                case "a":
                                    if false {
                                        v1 -= 602
                                    } else {
                                        var v33 uint16 = 997
                                        v5 = append(v5, "hello")
                                        switch "\n\t" {
                                        case "hello":
                                            var v34 float32 = 3.14159
                                            _ = v34
                                        default:
                                            var v35 uint = 0
                                            _ = v35
                                        }
                                        _ = v33
                                    }
                                }
                                v36 := GenS4{F0: 100, F1: 127, F2: 636}
                                _ = v36
                            } else {
                                var v37 float64 = ((((((0.0 / ((1e10 - 0.0) / v3)) + (((float64(float32((-2.25 * v3))) / (((v3 / ((float64((0.0 - (float32(float64(0.0)) - 3.14159))) / v3) - v3)) / -2.25) + 0.0)) * ((v3 * v3) - v3)) + 1.5)) + 1e10) / (v3 - (float64((1e10 / 1.0)) + ((-2.25 + v3) - v3)))) * v3) - ((v3 + 1e10) + -2.25))
                                _ = v37
                            }
                        }
                        v38 := GenS5{F0: 42, F1: (v1 << ((1) & 63))}
                        _ = v38
                    }
                    if true {
                        fmt.Println(fi31, fi31, fi31)
                    } else {
                        fmt.Println(v1, fi31, fi30)
                        if (false && true) {
                            v39 := GenS4{F0: 0, F1: -1, F2: 65535}
                            _ = v39
                        } else {
                            switch fi31 {
                            case 4294967295:
                                var v40 uint16 = 42
                                _ = v40
                            case 0:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 2147483647:
                                v5 = append(v5, "foo bar")
                            }
                            switch (fi31 & (fi30 << ((0) & 63))) {
                            case 9223372036854775807:
                                var v41 uint8 = 42
                                _ = v41
                            }
                        }
                    }
                    _ = fi31
                    _ = v32
                }
            case 32767:
                fmt.Println(v3)
            case -922:
                fmt.Println(v1, v1, v3)
            }
            _ = fi30
        }
    } else {
        if (-22 == 42) {
            var v42 string = ("hello" + "foo bar")
            _ = v42
        }
    }
    fmt.Println(v3, v3)
    var v43 bool = (!true)
    for fi44 := 0; fi44 < 13; fi44++ {
        if v43 {
            for fi45 := 0; fi45 < 13; fi45++ {
                fmt.Println(v3, fi45, v1)
                v1 = -67
                var v46 int16 = 32767
                _ = fi45
                _ = v46
            }
        } else {
            if (42 > 642) {
                for fi47 := 0; fi47 < 13; fi47++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (v43 || v43) {
                        switch fi44 {
                        case 0:
                            switch ("" + "a") {
                            case "":
                                if (!v43) {
                                    var v48 uint8 = 42
                                    v3 = v3
                                    if (v48 > v48) {
                                        var v49 float64 = 0.1
                                        switch (fi44 ^ 0) {
                                        case 4294967295:
                                            if (!v43) {
                                                if v43 {
                                                    v48 -= 42
                                                }
                                            }
                                        case -2147483648:
                                            v3 -= (1e10 - ((((float64(((0.0 * 3.14159) - (((3.14159 / ((1e10 * -2.25) / float32((1.5 * (((v49 + v3) - (3.14159 + (3.14159 + float64((float32(0.0) * 1.5))))) - 0.1))))) - -2.25) - -2.25))) * v3) / v49) - v49) / v3))
                                        case 1:
                                            if (((v48 | v48) < v48) && v43) {
                                                var v50 uint64 = 614
                                                _ = v50
                                            }
                                        }
                                        if ((v49 >= (v49 + (0.1 / v49))) && v43) {
                                            var v51 uint16 = 32767
                                            if (v3 <= -2.25) {
                                                v51 = (^(v51))
                                                var v52 uint64 = uint64(v51)
                                                v52 = 4294967295
                                                _ = v52
                                            }
                                            _ = v51
                                        } else {
                                            v4[fi44] = 1
                                        }
                                        _ = v49
                                    }
                                    _ = v48
                                }
                            default:
                                fmt.Println(v43, fi44)
                            }
                        case 32768:
                            fmt.Println(v3)
                        case 1:
                            var v53 int = (fi44 | fi44)
                            _ = v53
                        default:
                            v1 = (-(v1))
                        }
                        if ((v43 || (!(2 <= 42))) || (1 == 9223372036854775807)) {
                            v54 := GenS6{F0: 711, F1: 2}
                            fmt.Println(fi47, v3)
                            _ = v54
                        }
                        delete(v2, "a")
                    }
                    _ = fi47
                }
            }
            for fi55 := 0; fi55 < 15; fi55++ {
                if (v1 != (v1 - (^(v1)))) {
                    v56 := GenS1{F0: 0, F1: 0}
                    delete(v2, "")
                    _ = v56
                }
                _ = len(v2)
                switch (-(fi55)) {
                case -2147483648:
                    var v57 int32 = 42
                    _ = v57
                default:
                    v58 := GenS6{F0: 65535, F1: 2}
                    _ = v58
                }
                switch "foo bar" {
                case "foo bar":
                    v43 = ((2 > 109) && (!true))
                case "a":
                    fmt.Println(v1, v43)
                case "\n\t":
                    v3 += (float64(-2.25) / (float64((1.5 - (0.1 / -2.25))) / ((1.5 * 1.5) - (((v3 * v3) + -2.25) + 0.0))))
                    var v59 uint8 = 255
                    _ = v59
                }
                v60 := GenS7{F0: 127, F1: 892}
                delete(v2, ("a" + "hello"))
                _ = fi55
                _ = v60
            }
        }
        _ = fi44
    }
    v61 := GenS8{F0: -1}
    v62 := GenS4{F0: 54, F1: -128, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v61
    _ = v62
}
