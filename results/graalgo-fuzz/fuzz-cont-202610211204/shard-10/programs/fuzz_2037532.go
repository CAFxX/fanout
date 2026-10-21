package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
func genFn0() uint {
    v4 := []int{}
    var v5 uint32 = 4294967295
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            v5 = v5
            _ = fi7
        }
        if true {
            v5 = v5
            var v8 uint64 = 324
            _ = v8
        } else {
            if len(v4) > 0 {
                v4[0] = (fi6 * 9223372036854775807)
            }
            var v9 int32 = 2147483647
            _ = v9
        }
        var v10 float64 = 0.1
        _ = fi6
        _ = v10
    }
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int16, p1 uint) string {
    for fi11 := 0; fi11 < 8; fi11++ {
        p0 -= p0
        for fi12 := 0; fi12 < 6; fi12++ {
            p0 = (p0 + (p0 | 1000))
            _ = fi12
        }
        v13 := GenS1{F0: 200, F1: "", F2: ""}
        _ = fi11
        _ = v13
    }
    if (!("a" <= (("" + ("hello" + "a")) + ""))) {
        for fi14 := 0; fi14 < 8; fi14++ {
            p0 = p0
            _ = fi14
        }
        v15 := []int{-9223372036854775808, 9223372036854775807, -962, 292}
        _ = v15
    } else {
        var v16 float64 = (0.1 * 1e10)
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint32 = 65535
            _ = fi17
            _ = v18
        }
        var v19 int8 = 0
        _ = v16
        _ = v19
    }
    if (!true) {
        v20 := GenS1{F0: 537, F1: "\n\t", F2: "\n\t"}
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 bool = false
            p1 = p1
            if (1e10 == 1.5) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    fmt.Println(fi23, p0, p1)
                    fmt.Println(v22)
                    fmt.Println(p0, v22, fi23)
                    p1 -= p1
                    _ = fi23
                }
            }
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        fmt.Println(p1, p1)
        if (0.0 < 1e10) {
            if (1e10 < -2.25) {
                v24 := GenS2{F0: 4294967295, F1: 1, F2: 1}
                _ = v24
            } else {
                switch ("" + "\n\t") {
                case "foo bar":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        p0 += p0
                        var v26 int8 = 42
                        var v27 bool = ((^(fi25)) < (^(fi25)))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case "hello":
                    for fi28 := 0; fi28 < 7; fi28++ {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            var v30 uint32 = 850
                            _ = fi29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 4; fi31++ {
                            var v32 uint64 = 18446744073709551615
                            p0 -= (p0 << (((^(p1))) & 63))
                            p0 += 334
                            _ = fi31
                            _ = v32
                        }
                        var v33 int64 = 9223372036854775807
                        _ = fi28
                        _ = v33
                    }
                case "\n\t":
                    if false {
                        fmt.Println(p1, p0)
                        if false {
                            for fi34 := 0; fi34 < 2; fi34++ {
                                v35 := GenS1{F0: 0, F1: "foo bar", F2: "hello"}
                                fmt.Println(p1)
                                if true {
                                    v36 := GenS1{F0: 1, F1: "\n\t", F2: "foo bar"}
                                    v37 := GenS3{F0: 65535}
                                    fmt.Println(fi34, p0, fi34)
                                    _ = v36
                                    _ = v37
                                } else {
                                    if true {
                                        fmt.Println(p0)
                                        switch "hello" {
                                        case "a":
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v38 bool = (!false)
                                                _ = v38
                                            case "a":
                                                v39 := GenS4{F0: 42}
                                                _ = v39
                                            default:
                                                v40 := []int{}
                                                _ = v40
                                            }
                                        default:
                                            v41 := []int{128, (^(fi34))}
                                            _ = v41
                                        }
                                        p0 = p0
                                    }
                                }
                                _ = fi34
                                _ = v35
                            }
                            var v42 uint32 = 6
                            for fi43 := 0; fi43 < 18; fi43++ {
                                p0 = (p0 - p0)
                                v44 := []int{(fi43 * fi43), int(v42), fi43}
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    _ = len(v44)
                                    if false {
                                        switch fi45 {
                                        case 1:
                                            var v46 uint16 = 896
                                            _ = v46
                                        }
                                        var v47 int = (-(fi43))
                                        fmt.Println(v47)
                                        _ = v47
                                    }
                                    switch (fi43 * (fi45 >> ((p1) & 63))) {
                                    case -813:
                                        p0 -= p0
                                    case -1:
                                        var v48 int = (fi43 % ((1) | 1))
                                        _ = v48
                                    }
                                    _ = fi45
                                }
                                _ = fi43
                                _ = v44
                            }
                            _ = v42
                        }
                    } else {
                        switch "\n\t" {
                        case "a":
                            for fi49 := 0; fi49 < 6; fi49++ {
                                for fi50 := 0; fi50 < 13; fi50++ {
                                    var v51 uint64 = 350
                                    _ = fi50
                                    _ = v51
                                }
                                v52 := []int{32767, fi49, -1, fi49}
                                _ = fi49
                                _ = v52
                            }
                        default:
                            p1 = p1
                        }
                        switch 128 {
                        case 358:
                            fmt.Println(p0)
                        case -2147483648:
                            fmt.Println(p1)
                        }
                    }
                default:
                    p0 = -489
                }
            }
        }
        v53 := []int{}
        _ = v53
    }
    v54 := GenS5{F0: 0, F1: uint32(p1), F2: 0}
    for fi55 := 0; fi55 < 11; fi55++ {
        var v56 int16 = p0
        if true {
            if ((p1 ^ genFn0()) == p1) {
                fmt.Println(v56, fi55)
                v57 := []int{(fi55 >> ((p1) & 63)), 4294967295}
                _ = len(v57)
                _ = v57
            } else {
                var v58 float64 = 1.5
                for fi59 := 0; fi59 < 11; fi59++ {
                    fmt.Println(v58, fi59, v58)
                    var v60 uint = p1
                    _ = fi59
                    _ = v60
                }
                _ = v58
            }
            switch (fi55 << (((p1 * p1)) & 63)) {
            case 0:
                fmt.Println(fi55, fi55)
            case 779:
                v61 := []int{(^(fi55)), (fi55 / ((fi55) | 1)), fi55, (fi55 | (fi55 & (fi55 >> ((p1) & 63))))}
                _ = v61
            default:
                for fi62 := 0; fi62 < 5; fi62++ {
                    p1 += 255
                    fmt.Println(fi62, p0)
                    _ = fi62
                }
            }
            p0 = (p0 << ((p1) & 63))
        } else {
            var v63 int8 = 0
            v64 := []int{fi55}
            var v65 uint8 = uint8(p1)
            _ = v63
            _ = v64
            _ = v65
        }
        _ = fi55
        _ = v56
    }
    _ = p0
    _ = p1
    _ = v54
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    v3 := []string{}
    v66 := GenS1{F0: 2, F1: "", F2: "hello"}
    var v67 int64 = 565
    v68 := v1[("hello" + "hello")]
    v69 := v2[(v68 + ("a" + ("" + v68)))]
    v69 = v69
    for fi70 := 0; fi70 < 4; fi70++ {
        v71 := GenS4{F0: 2147483647}
        _ = fi70
        _ = v71
    }
    v72 := GenS4{F0: 669}
    v73 := v1[v69]
    switch -970 {
    case 128:
        var v74 int8 = 100
        _ = v74
    case 4294967295:
        v67 = v67
    case 127:
        v68 = v68
    default:
        if true {
            _ = len(v1)
            if true {
                v1[v73] = ("\n\t" + "a")
                var v75 int32 = -354
                _ = v75
            } else {
                v76 := GenS5{F0: 2147483648, F1: 2147483647, F2: 368}
                _ = v76
            }
        } else {
            v69 = (v69 + (("hello" + (("\n\t" + v69) + "a")) + v69))
        }
    }
    fmt.Println(v69, v67, v73)
    var v77 int64 = -706
    var v78 float32 = 1e10
    for fi79 := 0; fi79 < 1; fi79++ {
        var v80 int32 = 2147483647
        v67 = -999
        if len(v3) > 0 {
            v3[0] = v69
        }
        _ = fi79
        _ = v80
    }
    var v81 uint16 = 65535
    if false {
        if false {
            v82 := GenS3{F0: 65535}
            var v83 bool = (!(true && ((!(float64(0.1) > ((0.1 * float64(1e10)) - (1e10 / 1.5)))) || (true || true))))
            _ = v82
            _ = v83
        }
        var v84 uint8 = 128
        _ = v84
    }
    v85 := GenS5{F0: 95, F1: 2147483648, F2: 798}
    var v86 uint = 697
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v73:", v73)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v72
    _ = v73
    _ = v77
    _ = v78
    _ = v81
    _ = v85
    _ = v86
}
