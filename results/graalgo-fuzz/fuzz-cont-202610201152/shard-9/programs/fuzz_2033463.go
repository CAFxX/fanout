package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0() string {
    var v5 uint16 = 191
    v5 = (^(v5))
    v6 := []int{4294967295}
    v5 = (v5 ^ v5)
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 int32, p1 int8) int64 {
    fmt.Println(p1)
    switch -1 {
    case -562:
        if (-1 < (p1 | 100)) {
            var v7 int64 = 1
            v8 := []int{}
            switch 32767 {
            case 160:
                v9 := GenS1{F0: 42}
                _ = v9
            case 42:
                if (!true) {
                    _ = len(v8)
                    v10 := GenS2{F0: 0}
                    switch 2147483647 {
                    case 993:
                        v7 -= (v7 | (-(v7)))
                    case 738:
                        if (1 == -890) {
                            v7 -= v7
                            fmt.Println(p1, v7, p1)
                        } else {
                            p1 = p1
                            v11 := GenS2{F0: (p0 / ((-222) | 1))}
                            if false {
                                for fi12 := 0; fi12 < 3; fi12++ {
                                    var v13 uint32 = 2
                                    var v14 string = "\n\t"
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                            } else {
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                            _ = v11
                        }
                    case 150:
                        fmt.Println(p1)
                    default:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    _ = v10
                }
            default:
                for fi15 := 0; fi15 < 10; fi15++ {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        fmt.Println(fi16, fi15, p0)
                        _ = fi16
                    }
                    p1 = 57
                    _ = fi15
                }
            }
            _ = v7
            _ = v8
        }
    case 32767:
        switch "" {
        case "":
            for fi17 := 0; fi17 < 5; fi17++ {
                p1 = -128
                var v18 uint64 = 1
                _ = fi17
                _ = v18
            }
        case "hello":
            if false {
                v19 := []int{42, 172, 880, 128}
                _ = v19
            }
        case "a":
            var v20 int16 = 360
            _ = v20
        }
    }
    v21 := []int{42, 42}
    if true {
        switch -2147483648 {
        case 9223372036854775807:
            v22 := GenS2{F0: 890}
            _ = v22
        case -119:
            var v23 int = 307
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v21
    return -1
}

func main() {
    var v1 int64 = 42
    var v2 int64 = 0
    var v3 bool = ((((1e10 / ((0.1 / float64(-2.25)) - 0.0)) * (1.5 / 3.14159)) * 1e10) == (0.0 - (-2.25 + 0.0)))
    var v4 uint8 = 1
    if ((true || v3) || (65535 > 0)) {
        for fi24 := 0; fi24 < 15; fi24++ {
            switch ("" + "hello") {
            case "hello":
                switch ("hello" + "hello") {
                case "foo bar":
                    v1 -= 0
                default:
                    v3 = true
                }
            }
            v25 := GenS3{F0: (fi24 | (-(fi24))), F1: "a"}
            var v26 uint64 = 606
            _ = fi24
            _ = v25
            _ = v26
        }
        switch 4294967295 {
        case 32768:
            var v27 uint32 = 63
            _ = v27
        case 42:
            if true {
                var v28 uint8 = v4
                var v29 uint = 18446744073709551615
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 3; fi30++ {
                    v4 = (v4 * (^(v4)))
                    _ = fi30
                }
                v2 += genFn1(-2147483648, 127)
                if v3 {
                    var v31 uint32 = uint32(v4)
                    _ = v31
                }
            }
        case 112:
            v2 = (-(v2))
        default:
            switch 898 {
            case 32767:
                var v32 string = "foo bar"
                _ = v32
            default:
                v33 := GenS1{F0: 0}
                _ = v33
            }
        }
    } else {
        var v34 float32 = (-2.25 - 1.5)
        _ = v34
    }
    if false {
        v35 := []int{493, 128}
        switch 46 {
        case 446:
            switch 1 {
            case 4294967295:
                switch ("\n\t" + "") {
                case "hello":
                    switch ("" + "\n\t") {
                    case "foo bar":
                        var v36 int = 4294967295
                        _ = v36
                    default:
                        fmt.Println(v3, v3, v3)
                    }
                case "foo bar":
                    var v37 float64 = 0.1
                    _ = v37
                case "\n\t":
                    fmt.Println(v2)
                default:
                    v35 = append(v35, -257)
                }
            default:
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 int64 = 42
                    v2 += 42
                    _ = fi38
                    _ = v39
                }
            }
        }
        if (!(((3.14159 + (-2.25 * float32(0.1))) - 0.1) <= (1.5 * 0.0))) {
            fmt.Println(v1, v1)
            if len(v35) >= 2 {
                v35 = v35[1:len(v35)]
            }
            fmt.Println(v1)
        } else {
            var v40 int32 = -2147483648
            if v3 {
                fmt.Println(v2)
            }
            switch 4294967295 {
            case 4294967295:
                switch "a" {
                case "foo bar":
                    fmt.Println(v4)
                case "hello":
                    if v3 {
                        switch -1 {
                        case 1:
                            v41 := GenS3{F0: 42, F1: "a"}
                            _ = v41
                        }
                        v40 = (-(v40))
                    }
                default:
                    var v42 int16 = 1
                    _ = v42
                }
            case -797:
                v43 := GenS4{F0: "hello"}
                _ = v43
            default:
                for fi44 := 0; fi44 < 6; fi44++ {
                    switch "\n\t" {
                    case "\n\t":
                        if len(v35) > 0 {
                            v35[0] = fi44
                        }
                    case "foo bar":
                        v45 := GenS5{F0: 0}
                        _ = v45
                    }
                    _ = fi44
                }
            }
            _ = v40
        }
        _ = v35
    }
    fmt.Println(v4, v1, v1)
    for fi46 := 0; fi46 < 5; fi46++ {
        if v3 {
            var v47 int = fi46
            v48 := []int{v47, v47, 1}
            _ = v47
            _ = v48
        } else {
            var v49 int16 = 573
            v49 -= 1
            _ = v49
        }
        switch -618 {
        case 612:
            v50 := GenS6{F0: 304}
            _ = v50
        case -9223372036854775808:
            switch "a" {
            case "foo bar":
                fmt.Println(v3)
            case "hello":
                if v3 {
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        switch -559 {
                        case 32767:
                            v2 += (v1 / ((0) | 1))
                        case 9223372036854775807:
                            v2 += v1
                        default:
                            v1 += 176
                        }
                    case "foo bar":
                        v51 := []int{(fi46 | fi46), fi46}
                        _ = v51
                    case "a":
                        v3 = (2 > 32767)
                    default:
                        if false {
                            v52 := GenS7{F0: -50, F1: (v2 * (^(v1))), F2: 2147483647}
                            fmt.Println(fi46, v1)
                            fmt.Println(v2)
                            _ = v52
                        }
                    }
                }
            case "":
                fmt.Println(v2, v3, v2)
            default:
                v53 := []int{fi46}
                _ = v53
            }
        default:
            var v54 int = fi46
            _ = v54
        }
        v2 += v2
        _ = fi46
    }
    var v55 uint8 = v4
    v4 = (v55 >> ((65535) & 63))
    v56 := []int{993, 850, 127, -167}
    var v57 int8 = 0
    switch ("foo bar" + "") {
    case "hello":
        for fi58 := 0; fi58 < 7; fi58++ {
            var v59 bool = false
            var v60 bool = (v3 || (1.5 != float32(((float64(0.1) / 1.0) + 1e10))))
            _ = fi58
            _ = v59
            _ = v60
        }
    case "\n\t":
        if true {
            if v3 {
                v3 = (42 == 22)
                v56 = append(v56, 128)
                if (248 >= 1) {
                    if (0 > 42) {
                        switch "a" {
                        case "a":
                            switch (("" + ((("a" + "foo bar") + "a") + "\n\t")) + "\n\t") {
                            case "":
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v4, v55, v55)
                                case "hello":
                                    var v61 uint8 = 2
                                    _ = v61
                                }
                            case "foo bar":
                                v56 = append(v56, 0)
                            }
                        case "":
                            if (-2.25 >= (0.0 / 1.0)) {
                                var v62 int8 = (-(v57))
                                _ = v62
                            }
                        }
                        fmt.Println(v4)
                    }
                    if true {
                        var v63 int16 = 42
                        var v64 uint8 = (^(v55))
                        var v65 string = (("" + "foo bar") + "a")
                        _ = v63
                        _ = v64
                        _ = v65
                    } else {
                        var v66 uint16 = 42
                        if v3 {
                            _ = len(v56)
                        } else {
                            if len(v56) > 0 {
                                v56[0] = 127
                            }
                        }
                        v67 := GenS7{F0: 100, F1: v2, F2: 9}
                        _ = v66
                        _ = v67
                    }
                    v4 -= 1
                }
            }
            v57 = -128
        } else {
            if v3 {
                if (1.5 > 0.1) {
                    for fi68 := 0; fi68 < 6; fi68++ {
                        v69 := GenS3{F0: -9223372036854775808, F1: ("hello" + ("a" + ""))}
                        _ = fi68
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 8; fi70++ {
                        switch "foo bar" {
                        case "hello":
                            if len(v56) > 0 {
                                v56[0] = fi70
                            }
                        case "a":
                            v2 = (v2 ^ (v2 / ((v1) | 1)))
                        }
                        if (2 != 281) {
                            var v71 uint64 = 42
                            v72 := GenS3{F0: (fi70 * (fi70 / (((fi70 * fi70)) | 1))), F1: "foo bar"}
                            v1 += (v1 >> ((699) & 63))
                            _ = v71
                            _ = v72
                        }
                        _ = fi70
                    }
                }
                v2 += v1
                v73 := GenS5{F0: -32768}
                _ = v73
            } else {
                var v74 int = 9223372036854775807
                v4 -= (v4 & (v55 & (v4 & 127)))
                if len(v56) >= 2 {
                    v56 = v56[1:len(v56)]
                }
                _ = v74
            }
        }
    case "":
        if false {
            fmt.Println(v1)
            var v75 uint16 = 794
            if (true || true) {
                var v76 uint = 4294967295
                v77 := GenS1{F0: -774}
                _ = v76
                _ = v77
            } else {
                for fi78 := 0; fi78 < 11; fi78++ {
                    for fi79 := 0; fi79 < 11; fi79++ {
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                        v80 := GenS3{F0: (fi79 + -316), F1: "a"}
                        _ = fi79
                        _ = v80
                    }
                    v56 = append(v56, (fi78 ^ -167))
                    var v81 float32 = 3.14159
                    _ = fi78
                    _ = v81
                }
                if len(v56) > 0 {
                    v56[0] = 32767
                }
            }
            _ = v75
        } else {
            var v82 string = "hello"
            v1 += v2
            _ = v82
        }
    default:
        fmt.Println(v4)
    }
    v57 = v57
    v56 = append(v56, 591)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
}
