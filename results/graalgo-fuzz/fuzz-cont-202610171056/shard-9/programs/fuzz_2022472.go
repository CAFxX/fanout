package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint32; F2 int16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS9 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint32) uint16 {
    p1 = p0
    var v5 string = ((("foo bar" + ((("a" + ("a" + "foo bar")) + "hello") + "\n\t")) + (("\n\t" + ((("a" + "foo bar") + "foo bar") + (((("\n\t" + ((("a" + "\n\t") + "") + "foo bar")) + ("hello" + ("hello" + "a"))) + (("" + "hello") + "")) + "foo bar"))) + "foo bar")) + ("hello" + "a"))
    v5 = v5
    v6 := []int{128, 978, 128, -37}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 833
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := []string{"", ("" + "\n\t")}
    if true {
        if true {
            v7 := GenS1{F0: 469, F1: 18446744073709551615, F2: 578}
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 4294967295, F1: 898, F2: 2147483648}
                v1 = v1
                switch "" {
                case "\n\t":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 int32 = 2147483647
                        _ = fi10
                        _ = v11
                    }
                case "":
                    if (v1 < (v1 | (v1 * 1))) {
                        v12 := GenS1{F0: 42, F1: 268, F2: 42}
                        var v13 float64 = ((1.5 + 3.14159) - 1e10)
                        v13 += v13
                        _ = v12
                        _ = v13
                    } else {
                        var v14 float32 = (((1e10 / float32(-2.25)) / 3.14159) - (float32(((0.0 - 3.14159) * 0.0)) * 1.5))
                        v15 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 0}
                        _ = v14
                        _ = v15
                    }
                case "hello":
                    v1 = v1
                }
                _ = fi8
                _ = v9
            }
            if false {
                switch 119 {
                case 73:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v1 += v1
                        _ = fi16
                    }
                case -9223372036854775808:
                    var v17 float32 = 3.14159
                    _ = v17
                }
            }
            _ = v7
        }
        v1 -= 1
        delete(v3, 1)
    } else {
        v1 += v1
        v18 := GenS1{F0: 255, F1: 369, F2: 42}
        _ = v18
    }
    v19 := GenS2{F0: 4294967295, F1: 2}
    v1 = (v1 - (v1 + v1))
    for fi20 := 0; fi20 < 15; fi20++ {
        if false {
            for fi21 := 0; fi21 < 3; fi21++ {
                fmt.Println(fi21, fi20)
                switch ("" + "a") {
                case "":
                    v22 := GenS1{F0: 942, F1: 18446744073709551615, F2: 2147483648}
                    _ = v22
                }
                var v23 int = fi20
                _ = fi21
                _ = v23
            }
            if true {
                var v24 uint32 = 2
                fmt.Println(v1, fi20)
                _ = v24
            }
            var v25 uint64 = 4294967295
            _ = v25
        }
        if false {
            fmt.Println(fi20, v1)
        } else {
            fmt.Println(v1, v1)
            var v26 int = 42
            switch (("hello" + "a") + "foo bar") {
            case "\n\t":
                switch (v26 ^ v26) {
                case -689:
                    v27 := GenS3{F0: false, F1: 2147483647, F2: 187}
                    _ = v27
                case 1:
                    for fi28 := 0; fi28 < 1; fi28++ {
                        if true {
                            if false {
                                v26 += 0
                                var v29 bool = ((fi28 >> ((151) & 63)) == fi20)
                                _ = v29
                            }
                            if (("hello" + "foo bar") != "\n\t") {
                                var v30 float32 = -2.25
                                v26 -= 42
                                _ = v30
                            }
                        } else {
                            v26 = 9223372036854775807
                            if true {
                                fmt.Println(fi20, fi20)
                                v31 := GenS2{F0: 995, F1: 42}
                                _ = v31
                            }
                            var v32 uint32 = 0
                            _ = v32
                        }
                        var v33 int64 = v1
                        _ = fi28
                        _ = v33
                    }
                case -729:
                    var v34 uint16 = 234
                    _ = v34
                }
            case "":
                v1 = (v1 << ((4294967295) & 63))
            }
            _ = v26
        }
        _ = fi20
    }
    fmt.Println(v1)
    var v35 string = (("a" + ("foo bar" + "a")) + "")
    var v36 float64 = 0.1
    for fi37 := 0; fi37 < 14; fi37++ {
        if true {
            v38 := GenS4{F0: v35, F1: 0, F2: 0}
            _ = len(v3)
            _ = v38
        } else {
            for fi39 := 0; fi39 < 12; fi39++ {
                switch (fi37 * fi39) {
                case 427:
                    v40 := GenS1{F0: 65535, F1: 996, F2: 65535}
                    v36 -= v36
                    _ = v40
                case 9223372036854775807:
                    switch v35 {
                    case "":
                        fmt.Println(fi39, fi39, fi39)
                    case "\n\t":
                        var v41 float64 = float64(-2.25)
                        _ = v41
                    default:
                        if (false && false) {
                            if (!((^(v1)) == -625)) {
                                if (272 <= 255) {
                                    if true {
                                        switch (fi39 & 42) {
                                        case 1:
                                            v42 := GenS1{F0: 65535, F1: 2, F2: 4294967295}
                                            _ = v42
                                        case 596:
                                            if false {
                                                if true {
                                                    v43 := GenS2{F0: 832, F1: 8}
                                                    v44 := GenS5{F0: true, F1: 332, F2: 565}
                                                    var v45 uint = 1
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                                v35 = "\n\t"
                                            } else {
                                                v36 += -2.25
                                                v46 := GenS2{F0: -9223372036854775808, F1: 128}
                                                v47 := GenS6{F0: 83}
                                                _ = v46
                                                _ = v47
                                            }
                                        case 564:
                                            _ = len(v3)
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = v35
                                            }
                                        }
                                        var v48 int16 = -475
                                        var v49 uint8 = 1
                                        _ = v48
                                        _ = v49
                                    }
                                    v35 = (v35 + v35)
                                    var v50 uint8 = 196
                                    _ = v50
                                }
                                v51 := GenS7{F0: 65535, F1: 18446744073709551615, F2: 1}
                                var v52 uint8 = 189
                                switch fi37 {
                                case 0:
                                    v2[65535] = v35
                                case 9223372036854775807:
                                    switch (fi37 & (fi37 << ((856) & 63))) {
                                    case 0:
                                        if len(v4) > 0 {
                                            v4[0] = (((v35 + v35) + "\n\t") + v35)
                                        }
                                    case 675:
                                        switch ((v35 + v35) + v35) {
                                        case "foo bar":
                                            fmt.Println(fi39, fi39, v52)
                                        case "":
                                            _ = len(v4)
                                        case "a":
                                            switch v35 {
                                            case "hello":
                                                v53 := GenS1{F0: 650, F1: 4294967295, F2: 35}
                                                _ = v53
                                            case "a":
                                                fmt.Println(fi37, v1, v52)
                                            case "":
                                                if true {
                                                    fmt.Println(fi37)
                                                    fmt.Println(fi37, v36)
                                                    var v54 uint16 = 32767
                                                    v52 += v52
                                                    fmt.Println(v52, v35)
                                                    v55 := GenS8{F0: 0, F1: 2, F2: 1}
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if (0.0 <= 0.0) {
                                                    var v56 string = v35
                                                    var v57 bool = false
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    var v58 bool = (false && (!false))
                                                    v52 -= (^(v52))
                                                    fmt.Println(fi37, v58)
                                                    _ = v58
                                                }
                                            }
                                        default:
                                            if false {
                                                var v59 uint32 = 2
                                                v1 = 1
                                                var v60 int16 = -1
                                                _ = v59
                                                _ = v60
                                            } else {
                                                var v61 int = fi37
                                                _ = v61
                                            }
                                        }
                                        var v62 int = (fi39 << ((uint(v52)) & 63))
                                        delete(v3, 2147483648)
                                        _ = v62
                                    case 638:
                                        var v63 float64 = (v36 + (-2.25 - v36))
                                        _ = v63
                                    }
                                case 63:
                                    var v64 uint8 = (v52 >> ((65535) & 63))
                                    _ = v64
                                default:
                                    _ = len(v3)
                                }
                                _ = v51
                                _ = v52
                            }
                        }
                    }
                    if ((("hello" + (v35 + (v35 + ""))) + (((("hello" + (("hello" + "a") + v35)) + ((("a" + v35) + v35) + v35)) + (v35 + (v35 + "a"))) + v35)) >= v35) {
                        v35 = v35
                        var v65 uint = 529
                        switch v35 {
                        case "\n\t":
                            v35 = "foo bar"
                        default:
                            var v66 uint32 = 42
                            _ = v66
                        }
                        _ = v65
                    } else {
                        if false {
                            v67 := GenS1{F0: 2, F1: 104, F2: 26}
                            v35 = "\n\t"
                            v1 += (v1 % (((v1 * v1)) | 1))
                            _ = v67
                        }
                    }
                    fmt.Println(v35)
                    if false {
                        v1 = v1
                    }
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    v68 := GenS9{F0: 0, F1: 7}
                    var v69 float64 = 0.1
                    switch 32767 {
                    case -500:
                        fmt.Println(v35, v69)
                    case 32767:
                        var v70 float32 = float32((v69 * (v69 + (v69 / 1.5))))
                        _ = v70
                    }
                    var v71 float64 = v69
                    _ = v68
                    _ = v69
                    _ = v71
                default:
                    v1 = (v1 * (-(v1)))
                }
                v1 -= (v1 << ((18446744073709551615) & 63))
                _ = fi39
            }
            v1 += v1
            var v72 uint = 709
            _ = v72
        }
        _ = fi37
    }
    var v73 uint = 309
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v73
}
