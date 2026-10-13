package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0() uint16 {
    var v7 bool = (1 == 0)
    var v8 int16 = 0
    fmt.Println(v7, v8)
    for fi9 := 0; fi9 < 11; fi9++ {
        if ((0 >= 265) && v7) {
            v10 := []int{(fi9 << ((65535) & 63)), (fi9 >> ((325) & 63)), fi9, (fi9 << ((0) & 63))}
            _ = v10
        } else {
            if false {
                if (!v7) {
                    v8 = 42
                } else {
                    var v11 float32 = float32(float64(-2.25))
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if false {
                            v13 := []int{}
                            v7 = (127 >= 255)
                            v14 := GenS1{F0: 255, F1: "a"}
                            _ = v13
                            _ = v14
                        } else {
                            switch "" {
                            case "foo bar":
                                var v15 string = ("a" + ("a" + "a"))
                                _ = v15
                            }
                            fmt.Println(fi12, v11)
                        }
                        fmt.Println(v11, v11)
                        var v16 int64 = 996
                        if (v7 || (!false)) {
                            v17 := []int{fi12, (fi12 >> ((2) & 63)), fi12, (fi12 % (((^(fi12))) | 1))}
                            if (v8 > (v8 & v8)) {
                                v18 := GenS1{F0: 2, F1: ""}
                                v7 = (!(!true))
                                _ = v18
                            } else {
                                if v7 {
                                    if false {
                                        if (255 == 1) {
                                            v19 := GenS1{F0: 32767, F1: ("a" + ("" + "a"))}
                                            _ = v19
                                        } else {
                                            var v20 uint16 = 65535
                                            _ = v20
                                        }
                                        v21 := GenS2{F0: 88, F1: 4294967295, F2: 80}
                                        if (0 != 808) {
                                            switch ("hello" + ("hello" + "\n\t")) {
                                            case "a":
                                                var v22 uint32 = 178
                                                _ = v22
                                            case "":
                                                v17 = append(v17, (fi12 + int(v16)))
                                            case "\n\t":
                                                fmt.Println(v11)
                                            }
                                            v23 := GenS1{F0: 0, F1: ("hello" + "foo bar")}
                                            _ = v23
                                        } else {
                                            _ = len(v17)
                                            if len(v17) > 0 {
                                                v17[0] = (-(fi9))
                                            }
                                            v8 = -526
                                            var v24 string = ""
                                            _ = v24
                                        }
                                        switch 2147483647 {
                                        case 1:
                                            fmt.Println(v16, v11)
                                        case 9223372036854775807:
                                            var v25 float32 = (1e10 / (0.0 + v11))
                                            _ = v25
                                        case 128:
                                            _ = len(v17)
                                        }
                                        var v26 float32 = (0.0 - v11)
                                        _ = v21
                                        _ = v26
                                    }
                                    var v27 string = "\n\t"
                                    if v7 {
                                        var v28 uint8 = 2
                                        _ = v28
                                    }
                                    _ = v27
                                }
                                var v29 uint16 = 2
                                var v30 int64 = v16
                                _ = v29
                                _ = v30
                            }
                            var v31 uint = 507
                            _ = v17
                            _ = v31
                        }
                        _ = fi12
                        _ = v16
                    }
                    _ = v11
                }
                var v32 string = ("\n\t" + "\n\t")
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 uint = 65535
                    fmt.Println(v8)
                    var v35 int = 2147483647
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                _ = v32
            }
            if (v8 != v8) {
                v36 := []int{fi9, (fi9 - (^(fi9)))}
                fmt.Println(v7)
                _ = v36
            } else {
                var v37 string = ("hello" + ("" + "\n\t"))
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 4; fi38++ {
            v8 = 0
            v39 := []int{42}
            if (143 <= 255) {
                v39 = append(v39, (fi38 << ((65535) & 63)))
                var v40 bool = (0 >= 75)
                _ = v40
            } else {
                switch "a" {
                case "\n\t":
                    v8 = v8
                }
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
            }
            _ = fi38
            _ = v39
        }
        for fi41 := 0; fi41 < 6; fi41++ {
            v42 := GenS1{F0: 42, F1: "a"}
            _ = fi41
            _ = v42
        }
        _ = fi9
    }
    _ = v7
    _ = v8
    return 1
}
func genFn1() int {
    var v43 float32 = 3.14159
    v44 := []int{-516, 9223372036854775807, -549, 32767}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    _ = v43
    _ = v44
    return -285
}

func main() {
    var v1 int16 = 1
    var v2 uint = 0
    v3 := []string{"foo bar", ("foo bar" + ("\n\t" + "a")), ("hello" + "hello"), "\n\t"}
    var v4 uint64 = 0
    var v5 int8 = 42
    var v6 int8 = (^(v5))
    switch "a" {
    case "\n\t":
        var v45 uint = (v2 / ((v2) | 1))
        _ = v45
    case "foo bar":
        var v46 uint32 = 0
        _ = v46
    case "":
        switch "" {
        case "\n\t":
            v3 = append(v3, "\n\t")
        case "":
            if true {
                var v47 string = ""
                v4 = 42
                _ = v47
            } else {
                var v48 float32 = (1.5 / 0.1)
                if (!false) {
                    v6 += (^(v5))
                    if (-1 == -1) {
                        for fi49 := 0; fi49 < 9; fi49++ {
                            var v50 uint = v2
                            var v51 string = ""
                            v52 := GenS3{F0: -1, F1: 255, F2: -2147483648}
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            for fi53 := 0; fi53 < 2; fi53++ {
                                var v54 string = "\n\t"
                                v55 := GenS1{F0: genFn0(), F1: (v54 + v54)}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        case "hello":
                            v1 -= (v1 | (v1 ^ v1))
                        default:
                            var v56 int16 = (-(v1))
                            _ = v56
                        }
                        var v57 int64 = 235
                        fmt.Println(v4)
                        _ = v57
                    }
                }
                _ = v48
            }
        case "foo bar":
            for fi58 := 0; fi58 < 16; fi58++ {
                switch (("" + ((("\n\t" + ("a" + (("" + "foo bar") + "\n\t"))) + "a") + "a")) + "\n\t") {
                case "":
                    for fi59 := 0; fi59 < 3; fi59++ {
                        switch "\n\t" {
                        case "\n\t":
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                    var v60 int32 = 2147483647
                                    v2 -= (v2 & v2)
                                    v61 := GenS2{F0: fi58, F1: v2, F2: 192}
                                    if false {
                                        v62 := GenS1{F0: 42, F1: "hello"}
                                        var v63 float32 = 1e10
                                        if true {
                                            var v64 float32 = 0.0
                                            _ = v64
                                        }
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v60
                                    _ = v61
                                }
                                var v65 uint = v2
                                _ = v65
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + "")
                                }
                                fmt.Println(v4, fi58)
                                v66 := GenS1{F0: genFn0(), F1: ("foo bar" + ("\n\t" + "foo bar"))}
                                _ = v66
                            }
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                        case "a":
                            v1 = 42
                        default:
                            v67 := GenS4{F0: (v5 - (v5 + v5))}
                            var v68 float64 = 0.1
                            switch "foo bar" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = ("" + "\n\t")
                                }
                                fmt.Println(fi58)
                                fmt.Println(v5, v4, v4)
                            default:
                                switch "hello" {
                                case "\n\t":
                                    v6 = (v5 - (v6 >> (((v2 - 279)) & 63)))
                                case "hello":
                                    v5 = 0
                                }
                            }
                            var v69 uint32 = 0
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        v70 := GenS5{F0: ("" + (((((("hello" + "") + "a") + ("hello" + ("" + ("hello" + ("foo bar" + "hello"))))) + (("hello" + "\n\t") + ("a" + "hello"))) + "\n\t") + ((("foo bar" + (("hello" + (("\n\t" + "hello") + "")) + "foo bar")) + ("" + "hello")) + ("hello" + "foo bar")))), F1: 18446744073709551615, F2: (v4 & v4)}
                        fmt.Println(v1)
                        var v71 bool = (42 > 1)
                        v72 := GenS2{F0: fi58, F1: 255, F2: uint8(v4)}
                        _ = fi59
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                }
                _ = fi58
            }
        default:
            v6 -= 0
        }
    default:
        v3 = append(v3, (("hello" + "hello") + "\n\t"))
    }
    _ = len(v3)
    v73 := GenS3{F0: 9223372036854775807, F1: 32767, F2: 1}
    if (false || (float64(((3.14159 - 3.14159) * (-2.25 * ((float32(-2.25) / 0.1) / 3.14159)))) <= (float64((0.1 - 0.1)) - 0.0))) {
        v74 := GenS3{F0: 708, F1: 804, F2: -1}
        _ = v74
    }
    var v75 int = -308
    var v76 uint16 = 2
    var v77 float64 = 3.14159
    _ = len(v3)
    v76 -= (v76 - genFn0())
    v78 := GenS3{F0: 930, F1: genFn0(), F2: 881}
    for fi79 := 0; fi79 < 18; fi79++ {
        if true {
            if true {
                v80 := GenS2{F0: fi79, F1: v2, F2: 255}
                _ = v80
            } else {
                v5 = (v6 | (-(v6)))
                switch "foo bar" {
                case "":
                    var v81 string = "\n\t"
                    _ = v81
                case "\n\t":
                    fmt.Println(v6, v77, v2)
                case "foo bar":
                    v2 += v2
                }
                v2 = (v2 / ((714) | 1))
            }
            if len(v3) > 0 {
                v3[0] = ("hello" + "")
            }
        } else {
            if true {
                fmt.Println(v4)
                for fi82 := 0; fi82 < 16; fi82++ {
                    var v83 uint8 = 196
                    var v84 int64 = 812
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                var v85 uint16 = genFn0()
                _ = v85
            }
            v2 = 255
        }
        _ = len(v3)
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi79
    }
    var v86 bool = true
    switch (v75 % ((v75) | 1)) {
    case 32767:
        v87 := GenS6{F0: v4, F1: 538, F2: -480}
        _ = v87
    case -369:
        for fi88 := 0; fi88 < 2; fi88++ {
            switch (^(fi88)) {
            case 32768:
                v3 = append(v3, ("" + ("hello" + "a")))
            case 127:
                for fi89 := 0; fi89 < 5; fi89++ {
                    if true {
                        if false {
                            var v90 uint64 = v4
                            _ = v90
                        }
                        v86 = true
                        v76 += v76
                    }
                    v86 = (672 <= 9223372036854775807)
                    _ = fi89
                }
            }
            if (v4 >= (v4 + v4)) {
                var v91 int = genFn1()
                _ = v91
            }
            var v92 float32 = 1e10
            _ = fi88
            _ = v92
        }
    }
    v93 := GenS4{F0: int8(v76)}
    fmt.Println(v4, v1)
    if (228 > 20) {
        fmt.Println(v6, v77)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v86
    _ = v93
}
