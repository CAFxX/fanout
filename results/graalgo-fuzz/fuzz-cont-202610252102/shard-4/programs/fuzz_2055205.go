package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 uint; F1 int32; F2 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 uint8; F1 int8 }
func genFn0(p0 int) uint8 {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 float32 = 0.0
        if (!false) {
            var v9 uint16 = 255
            var v10 uint32 = 1
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v11 bool = false
    _ = p0
    _ = v11
    return 217
}
func genFn1(p0 uint, p1 int16) string {
    if (-2.25 == 1e10) {
        var v12 uint32 = 2
        fmt.Println(p0, p1)
        p1 += -1
        _ = v12
    } else {
        var v13 string = (("hello" + "\n\t") + (("foo bar" + "\n\t") + ("foo bar" + (((("a" + "hello") + "foo bar") + "") + ("hello" + "a")))))
        v14 := GenS1{F0: -71, F1: ("a" + (v13 + v13)), F2: 32767}
        p0 += 531
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        p0 = p0
        for fi16 := 0; fi16 < 13; fi16++ {
            switch "" {
            case "hello":
                v17 := []int{(fi16 | fi15)}
                _ = v17
            case "":
                if (-1 >= -1) {
                    if true {
                        var v18 uint64 = 2
                        _ = v18
                    }
                    v19 := []int{(fi15 + (fi15 * fi16)), fi16, fi15, 4294967295}
                    if len(v19) > 0 {
                        v19[0] = (fi15 + fi16)
                    }
                    p0 = 1
                    _ = v19
                } else {
                    v20 := GenS2{F0: (p0 << ((p0) & 63)), F1: 1, F2: 32767}
                    p0 -= p0
                    p1 -= (p1 - p1)
                    v21 := GenS1{F0: 0, F1: "", F2: fi15}
                    _ = v20
                    _ = v21
                }
            case "\n\t":
                if true {
                    if true {
                        v22 := []int{(-(fi16))}
                        _ = v22
                    }
                    p0 += p0
                } else {
                    if (127 >= 2) {
                        var v23 int = (fi15 / (((fi15 ^ -172)) | 1))
                        _ = v23
                    } else {
                        var v24 uint = (^(p0))
                        _ = v24
                    }
                }
            default:
                v25 := []int{fi16, -1, fi15, fi16}
                p1 = p1
                _ = v25
            }
            p1 -= (p1 | 632)
            var v26 uint8 = genFn0((fi16 + (fi15 ^ fi15)))
            _ = fi16
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := []int{1, fi15}
            if len(v28) > 0 {
                v28[0] = 128
            }
            _ = fi27
            _ = v28
        }
        _ = fi15
    }
    if (2 <= 2) {
        switch 32768 {
        case -1000:
            p1 = (-(p1))
        default:
            var v29 float64 = -2.25
            _ = v29
        }
        fmt.Println(p1, p0)
    } else {
        if false {
            fmt.Println(p1)
            if true {
                for fi30 := 0; fi30 < 10; fi30++ {
                    switch (-(fi30)) {
                    case 128:
                        var v31 int32 = 2147483647
                        _ = v31
                    case 167:
                        v32 := []int{fi30, (fi30 / (((fi30 - fi30)) | 1))}
                        _ = v32
                    default:
                        var v33 int32 = -1
                        _ = v33
                    }
                    p0 = 65535
                    _ = fi30
                }
                v34 := GenS3{F0: 2147483647, F1: p0}
                _ = v34
            } else {
                p1 += p1
                v35 := []int{-744}
                for fi36 := 0; fi36 < 20; fi36++ {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        var v38 uint = p0
                        p0 -= p0
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        _ = fi37
                        _ = v38
                    }
                    var v39 float32 = ((1e10 + 3.14159) * 3.14159)
                    switch "hello" {
                    case "hello":
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v41 := GenS4{F0: 2, F1: (-(fi36))}
                            fmt.Println(fi36, fi36, fi40)
                            v42 := GenS3{F0: (fi40 / (((^(fi36))) | 1)), F1: 18446744073709551615}
                            fmt.Println(p0, v39)
                            fmt.Println(p1)
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    case "foo bar":
                        v43 := GenS5{F0: -131}
                        _ = v43
                    }
                    _ = fi36
                    _ = v39
                }
                _ = v35
            }
        }
        p1 -= 1000
    }
    var v44 float64 = (((3.14159 - -2.25) + (1.5 + 1e10)) + 1e10)
    var v45 float32 = -2.25
    _ = p0
    _ = p1
    _ = v44
    _ = v45
    return (("a" + (("a" + "\n\t") + "")) + "")
}

func main() {
    var v1 uint32 = 346
    var v2 uint64 = 0
    var v3 uint64 = (v2 % (((v2 / (((v2 + (^(v2)))) | 1))) | 1))
    var v4 string = (("foo bar" + "hello") + ("foo bar" + "\n\t"))
    var v5 int = -383
    v6 := make(map[uint32]bool)
    v46 := []int{}
    for fi47 := 0; fi47 < 9; fi47++ {
        for fi48 := 0; fi48 < 4; fi48++ {
            v4 = ((v4 + "hello") + "foo bar")
            _ = fi48
        }
        _ = fi47
    }
    v49 := GenS4{F0: 2147483648, F1: v5}
    switch 42 {
    case -717:
        v50 := GenS5{F0: v5}
        _ = v50
    case 4294967295:
        var v51 uint8 = 127
        _ = v51
    case 1:
        v5 -= 737
    default:
        if true {
            v2 = (v3 << ((65535) & 63))
            v3 = v3
        } else {
            v3 = (^(v3))
            if (0.1 <= float64((-2.25 / float32(float64(3.14159))))) {
                v52 := GenS4{F0: (v1 + (v1 + v1)), F1: 9223372036854775807}
                fmt.Println(v3, v2, v4)
                v53 := GenS6{F0: 0, F1: v1, F2: 2147483647}
                _ = v52
                _ = v53
            }
        }
    }
    var v54 bool = false
    var v55 uint8 = 225
    _ = len(v6)
    v56 := GenS7{F0: (v55 + (v55 / (((^(v55))) | 1))), F1: -794}
    if true {
        v1 = 2
    } else {
        for fi57 := 0; fi57 < 11; fi57++ {
            v55 = (v55 - 127)
            var v58 int8 = -109
            fmt.Println(v54, v58)
            _ = fi57
            _ = v58
        }
        v4 = v4
    }
    v4 = v4
    v59 := v6[2147483647]
    var v60 int = v5
    if false {
        switch (^(v60)) {
        case -2147483648:
            var v61 int64 = 618
            _ = v61
        case 127:
            if (!false) {
                var v62 float64 = float64(0.0)
                for fi63 := 0; fi63 < 10; fi63++ {
                    switch v60 {
                    case -9223372036854775808:
                        var v64 int64 = int64(v55)
                        _ = v64
                    case 352:
                        for fi65 := 0; fi65 < 2; fi65++ {
                            var v66 int64 = 9223372036854775807
                            _ = fi65
                            _ = v66
                        }
                    }
                    _ = fi63
                }
                _ = v62
            } else {
                for fi67 := 0; fi67 < 20; fi67++ {
                    var v68 string = "foo bar"
                    v69 := v6[(v1 * 42)]
                    fmt.Println(v60, v55)
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                if ((((1e10 - ((0.0 * 1e10) / (-2.25 / (1e10 + (1e10 - 1e10))))) * 3.14159) / ((0.1 - ((0.1 + (1.5 / 1e10)) * -2.25)) - (0.0 + 1e10))) > float64(1e10)) {
                    fmt.Println(v55, v3, v3)
                    var v70 int64 = 1
                    _ = v70
                } else {
                    var v71 uint8 = (^(v55))
                    v3 += uint64(v1)
                    _ = v71
                }
            }
        default:
            var v72 int8 = -1
            _ = v72
        }
    } else {
        v73 := GenS8{F0: 1}
        for fi74 := 0; fi74 < 18; fi74++ {
            if v54 {
                if (-265 <= -1) {
                    for fi75 := 0; fi75 < 7; fi75++ {
                        if false {
                            var v76 int = fi75
                            _ = v76
                        } else {
                            switch "" {
                            case "a":
                                var v77 uint8 = (v55 | 1)
                                _ = v77
                            }
                            v60 = 32768
                        }
                        fmt.Println(v54, v54, fi74)
                        var v78 int16 = -1
                        _ = fi75
                        _ = v78
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v6[0] = (-532 == -582)
                    default:
                        v79 := GenS7{F0: 255, F1: -929}
                        _ = v79
                    }
                    v80 := GenS8{F0: 1}
                    _ = v80
                }
                v5 += 1
                var v81 float32 = (3.14159 * ((((0.1 * 3.14159) * 1e10) / 0.1) + -2.25))
                _ = v81
            }
            for fi82 := 0; fi82 < 3; fi82++ {
                switch v4 {
                case "hello":
                    delete(v6, v1)
                case "\n\t":
                    var v83 bool = ((((v4 + v4) + "foo bar") + ("foo bar" + "hello")) > (v4 + "hello"))
                    _ = v83
                default:
                    if len(v46) > 0 {
                        v46[0] = 4294967295
                    }
                }
                var v84 uint64 = (v2 * 4294967295)
                if (2 < v1) {
                    var v85 int32 = -2147483648
                    _ = v85
                } else {
                    var v86 bool = v59
                    _ = v86
                }
                v4 = "foo bar"
                _ = fi82
                _ = v84
            }
            _ = fi74
        }
        v87 := GenS9{F0: -1, F1: v55}
        _ = v73
        _ = v87
    }
    v3 = (v3 & uint64(v55))
    v88 := GenS10{F0: (v55 - (v55 ^ (v55 | (v55 * (v55 << ((412) & 63)))))), F1: 28}
    v59 = (!v59)
    var v89 int8 = -92
    switch 9223372036854775807 {
    case 682:
        var v90 uint32 = 1
        _ = v90
    case 32768:
        fmt.Println(v3, v55, v5)
    case 731:
        var v91 int8 = v89
        _ = v91
    default:
        v89 = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v59
    _ = v6
    _ = v60
    _ = v88
    _ = v89
}
