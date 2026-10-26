package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    var v5 float64 = float64(3.14159)
    v5 = (v5 + (0.1 * (((v5 * v5) * v5) * 0.0)))
    p0 = (p0 ^ (p0 & (p0 / (((p0 ^ 32768)) | 1))))
    _ = p0
    _ = v5
    return false
}
func genFn1() int32 {
    v6 := []int{2147483647, 655, 4294967295, -684}
    var v7 float32 = float32((float64((float32(-2.25) + -2.25)) + 1e10))
    _ = v6
    _ = v7
    return 922
}

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32(float64(-2.25))
    var v8 float64 = (float64(((-2.25 * -2.25) * v4)) - 1.5)
    v9 := []int{42, -103, -187}
    v10 := v2[2147483648]
    for fi11 := 0; fi11 < 19; fi11++ {
        if (float64(1.5) <= -2.25) {
            var v12 bool = true
            _ = v12
        } else {
            var v13 int32 = -99
            _ = v13
        }
        var v14 bool = ((657 >= uint16(v10)) || true)
        v1 = (^(v1))
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 float32 = v4
        switch 1 {
        case 42:
            v17 := v3[872]
            _ = v17
        case 32768:
            switch (fi15 * fi15) {
            case 32767:
                v18 := v3[(v10 << ((v1) & 63))]
                _ = v18
            case 1:
                var v19 uint16 = 255
                _ = v19
            default:
                switch (fi15 / ((fi15) | 1)) {
                case 198:
                    var v20 int = fi15
                    _ = v20
                default:
                    if true {
                        v21 := GenS1{F0: 0, F1: ("" + ("\n\t" + ""))}
                        if ((!false) || false) {
                            for fi22 := 0; fi22 < 18; fi22++ {
                                var v23 uint64 = uint64(v10)
                                var v24 string = ("foo bar" + ("foo bar" + "a"))
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            if (v8 == (0.1 * v8)) {
                                switch (fi15 - (^(fi15))) {
                                case -1:
                                    _ = len(v9)
                                default:
                                    v16 = (v4 * (v16 / -2.25))
                                }
                                fmt.Println(fi15)
                                switch -9223372036854775808 {
                                case 887:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                case 42:
                                    if false {
                                        v25 := GenS1{F0: 1, F1: ("" + "a")}
                                        var v26 float64 = v8
                                        _ = v25
                                        _ = v26
                                    }
                                }
                            } else {
                                v8 -= ((3.14159 - v8) + v8)
                                fmt.Println(v16)
                            }
                        } else {
                            var v27 int8 = int8(fi15)
                            v28 := GenS2{F0: 2, F1: 400}
                            _ = v27
                            _ = v28
                        }
                        _ = v21
                    } else {
                        fmt.Println(v16, fi15)
                        if false {
                            for fi29 := 0; fi29 < 18; fi29++ {
                                _ = len(v3)
                                _ = fi29
                            }
                            if true {
                                if ((65535 <= 2) && (true && true)) {
                                    v30 := GenS3{F0: fi15}
                                    _ = v30
                                } else {
                                    if (42 == 1) {
                                        v31 := GenS4{F0: "hello"}
                                        v4 += 1e10
                                        _ = v31
                                    }
                                    v8 += v8
                                }
                            }
                        }
                        v32 := GenS1{F0: 100, F1: ""}
                        _ = v32
                    }
                }
            }
        default:
            v33 := GenS2{F0: fi15, F1: fi15}
            _ = v33
        }
        _ = fi15
        _ = v16
    }
    if (0.0 >= float64(v4)) {
        v34 := GenS4{F0: ("" + "hello")}
        v1 = v1
        var v35 uint = v1
        _ = v34
        _ = v35
    }
    var v36 uint64 = 1
    if (!true) {
        v2[670] = 2147483648
        _ = len(v3)
        _ = len(v2)
    } else {
        switch ("a" + "a") {
        case "hello":
            v37 := GenS4{F0: "foo bar"}
            _ = v37
        case "foo bar":
            for fi38 := 0; fi38 < 19; fi38++ {
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 int = fi39
                    if true {
                        if true {
                            if (false && true) {
                                v40 = fi38
                                var v41 int32 = -2147483648
                                _ = v41
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = (fi38 + (v40 * fi38))
                                }
                                v42 := v2[v10]
                                _ = v42
                            }
                            v1 = 255
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v43 uint32 = (v10 - 42)
                            _ = v43
                        }
                    } else {
                        v44 := GenS2{F0: 67, F1: fi38}
                        v36 = (v36 % ((18446744073709551615) | 1))
                        v1 = 255
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v9 = append(v9, 725)
                        _ = v44
                    }
                    if (1 < 0) {
                        var v45 uint16 = 255
                        v46 := GenS5{F0: 42}
                        _ = v45
                        _ = v46
                    } else {
                        v36 = v36
                    }
                    _ = fi39
                    _ = v40
                }
                v47 := GenS1{F0: -128, F1: (("" + "\n\t") + (("a" + ("a" + ("a" + "foo bar"))) + ("foo bar" + "\n\t")))}
                _ = fi38
                _ = v47
            }
        default:
            v36 += (v36 >> ((v1) & 63))
        }
        for fi48 := 0; fi48 < 16; fi48++ {
            v8 += ((1.5 + -2.25) / (v8 - v8))
            switch -2147483648 {
            case 32768:
                if false {
                    fmt.Println(v1)
                    fmt.Println(fi48)
                    switch fi48 {
                    case -2147483648:
                        switch "\n\t" {
                        case "\n\t":
                            v49 := GenS5{F0: -1}
                            _ = v49
                        default:
                            v50 := GenS6{F0: 9223372036854775807}
                            _ = v50
                        }
                    }
                }
            case 9223372036854775807:
                for fi51 := 0; fi51 < 19; fi51++ {
                    if true {
                        switch ("" + "hello") {
                        case "hello":
                            if (2 >= (v1 & 948)) {
                                var v52 string = "foo bar"
                                _ = v52
                            }
                        }
                        switch ("" + "\n\t") {
                        case "foo bar":
                            v36 = 0
                        default:
                            if (-356 == 32767) {
                                if ((("hello" + "foo bar") + "a") >= "") {
                                    if (!((fi48 | (fi51 / ((fi51) | 1))) <= (fi48 % ((fi48) | 1)))) {
                                        v4 += v4
                                        v10 -= (^(v10))
                                        var v53 uint32 = (^(v10))
                                        _ = v53
                                    } else {
                                        fmt.Println(fi51)
                                        var v54 int16 = 1000
                                        _ = v54
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                } else {
                                    v4 = v4
                                    v55 := GenS4{F0: "\n\t"}
                                    _ = v55
                                }
                                var v56 uint16 = 174
                                delete(v2, v10)
                                _ = v56
                            }
                        }
                    }
                    v8 += v8
                    _ = fi51
                }
            case 1:
                var v57 bool = (!true)
                _ = v57
            default:
                if true {
                    v58 := GenS4{F0: (("" + "\n\t") + "foo bar")}
                    _ = v58
                }
            }
            _ = fi48
        }
        if false {
            fmt.Println(v8)
            v1 -= (v1 * (v1 | 1))
            if false {
                v59 := GenS5{F0: 1}
                var v60 uint = v1
                for fi61 := 0; fi61 < 13; fi61++ {
                    switch "foo bar" {
                    case "a":
                        v1 = v60
                    case "\n\t":
                        if (false || (!true)) {
                            var v62 uint8 = 128
                            var v63 float32 = 3.14159
                            var v64 int = int(v36)
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            switch fi61 {
                            case -339:
                                v4 += float32(v8)
                            case 9223372036854775807:
                                if ((true || false) || ((((1e10 + -2.25) - (-2.25 + 1.5)) - v8) == 1e10)) {
                                    var v65 uint16 = 2
                                    var v66 float32 = 0.0
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    _ = v65
                                    _ = v66
                                } else {
                                    for fi67 := 0; fi67 < 19; fi67++ {
                                        if ("a" < "a") {
                                            delete(v2, 4294967295)
                                        }
                                        _ = len(v9)
                                        if false {
                                            var v68 bool = ((fi61 * fi67) == fi67)
                                            _ = v68
                                        }
                                        _ = fi67
                                    }
                                    if false {
                                        var v69 bool = false
                                        _ = v69
                                    }
                                    fmt.Println(fi61, v1)
                                }
                            default:
                                if false {
                                    v10 = (^(v10))
                                    v60 = v60
                                }
                            }
                            for fi70 := 0; fi70 < 7; fi70++ {
                                v71 := GenS7{F0: -1, F1: 255, F2: 1}
                                fmt.Println(v1)
                                _ = fi70
                                _ = v71
                            }
                            v72 := GenS3{F0: 127}
                            _ = v72
                        }
                    default:
                        if (!(((float64((v4 * v4)) / v8) - (v8 * 1e10)) >= v8)) {
                            v60 += v60
                            _ = len(v3)
                            if len(v9) > 0 {
                                v9[0] = (fi61 / ((fi61) | 1))
                            }
                        } else {
                            var v73 uint32 = v10
                            var v74 int16 = 906
                            _ = v73
                            _ = v74
                        }
                    }
                    fmt.Println(v60)
                    _ = fi61
                }
                _ = v59
                _ = v60
            }
        } else {
            v75 := GenS4{F0: "hello"}
            delete(v3, v10)
            fmt.Println(v4)
            _ = v75
        }
    }
    var v76 int = -568
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v76
    _ = v8
    _ = v9
}
