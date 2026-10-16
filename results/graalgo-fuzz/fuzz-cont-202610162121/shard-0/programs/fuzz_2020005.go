package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 bool) int8 {
    var v7 int32 = 42
    var v8 uint = 42
    var v9 int32 = v7
    v10 := GenS1{F0: -1, F1: -9223372036854775808, F2: "a"}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 100
}
func genFn1(p0 uint16, p1 int8) int16 {
    switch "hello" {
    case "":
        var v11 string = "foo bar"
        _ = v11
    default:
        p0 = 768
    }
    fmt.Println(p1)
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 -= (p0 & 2)
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if false {
            p0 += 670
        } else {
            if true {
                var v14 bool = (!true)
                p0 = p0
                _ = v14
            } else {
                p1 = (p1 | p1)
            }
            var v15 int64 = 9223372036854775807
            _ = v15
        }
        p0 += p0
        _ = fi13
    }
    p1 += p1
    _ = p0
    _ = p1
    return 371
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 float64 = float64(0.1)
    var v5 int8 = 1
    var v6 uint64 = 855
    var v16 int16 = -19
    v3 = append(v3, "hello")
    for fi17 := 0; fi17 < 6; fi17++ {
        switch "\n\t" {
        case "":
            if true {
                var v18 string = ("hello" + ("\n\t" + "foo bar"))
                switch fi17 {
                case 128:
                    var v19 uint = 0
                    _ = v19
                case 42:
                    if (191 < 255) {
                        v20 := GenS2{F0: int32(v16), F1: fi17, F2: (v5 - (v5 & (v5 << ((694) & 63))))}
                        v21 := GenS1{F0: int32(v5), F1: fi17, F2: "hello"}
                        _ = v20
                        _ = v21
                    }
                case 127:
                    v22 := GenS3{F0: 34}
                    _ = v22
                default:
                    v5 += v5
                }
                _ = v18
            } else {
                for fi23 := 0; fi23 < 5; fi23++ {
                    v6 = 792
                    switch fi17 {
                    case -406:
                        _ = len(v3)
                    case 32767:
                        v24 := GenS1{F0: -2147483648, F1: (fi17 / ((-9223372036854775808) | 1)), F2: "hello"}
                        _ = v24
                    }
                    _ = fi23
                }
            }
        case "hello":
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(v6, v4, v16)
                _ = fi25
            }
        case "\n\t":
            var v26 uint16 = 0
            _ = v26
        default:
            var v27 float32 = 3.14159
            _ = v27
        }
        _ = fi17
    }
    v4 = v4
    var v28 bool = true
    v2[""] = uint32(v6)
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 float32 = (3.14159 * ((3.14159 * -2.25) + float32((v4 * 1.5))))
            _ = fi30
            _ = v31
        }
        var v32 uint16 = 999
        fmt.Println(fi29, v6)
        _ = fi29
        _ = v32
    }
    v16 = 1
    delete(v2, "hello")
    v4 = float64(1.5)
    v33 := GenS1{F0: 1, F1: 32768, F2: "foo bar"}
    switch 127 {
    case -940:
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 uint8 = 128
            for fi36 := 0; fi36 < 10; fi36++ {
                v28 = false
                _ = fi36
            }
            switch fi34 {
            case -1:
                var v37 int = fi34
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    case -473:
        for fi38 := 0; fi38 < 14; fi38++ {
            if (v16 != 42) {
                v39 := GenS2{F0: -2147483648, F1: (fi38 | fi38), F2: (v5 + (v5 / ((v5) | 1)))}
                _ = v39
            } else {
                var v40 int = (^(fi38))
                _ = v40
            }
            switch (fi38 & 625) {
            case -272:
                v41 := GenS4{F0: (v16 << ((4294967295) & 63))}
                _ = v41
            case 0:
                for fi42 := 0; fi42 < 14; fi42++ {
                    if v28 {
                        if true {
                            v28 = v28
                            if (738 != 0) {
                                v4 += 0.1
                                if true {
                                    if (v28 && false) {
                                        v28 = v28
                                        if (!false) {
                                            switch int(v5) {
                                            case 2147483647:
                                                if (v28 && v28) {
                                                    v16 = v16
                                                }
                                            case -9223372036854775808:
                                                var v43 uint16 = 65535
                                                _ = v43
                                            case 9223372036854775807:
                                                var v44 int8 = (v5 + v5)
                                                _ = v44
                                            }
                                            var v45 uint32 = 4294967295
                                            v28 = v28
                                            _ = v45
                                        } else {
                                            v46 := GenS3{F0: 974}
                                            _ = v46
                                        }
                                    } else {
                                        v28 = (!(((false && false) && (v28 && true)) && v28))
                                        if false {
                                            var v47 int64 = 1
                                            _ = v47
                                        } else {
                                            v5 = (v5 / ((int8(fi42)) | 1))
                                            switch fi38 {
                                            case 128:
                                                var v48 int32 = int32(v16)
                                                _ = v48
                                            case 4294967295:
                                                fmt.Println(fi42)
                                            case 9223372036854775807:
                                                var v49 uint64 = (^(v6))
                                                _ = v49
                                            }
                                        }
                                        v50 := v1[2147483647]
                                        _ = v50
                                    }
                                    _ = len(v3)
                                } else {
                                    delete(v2, "a")
                                    if v28 {
                                        var v51 uint64 = 42
                                        v4 = v4
                                        var v52 uint32 = 682
                                        fmt.Println(v52)
                                        switch "a" {
                                        case "a":
                                            if (-366 > (fi42 / (((-(fi42))) | 1))) {
                                                v53 := GenS5{F0: 0, F1: 9223372036854775807, F2: 2}
                                                _ = v53
                                            } else {
                                                v3 = append(v3, (("a" + "hello") + "hello"))
                                                var v54 uint32 = 2147483648
                                                v54 += v54
                                                _ = v54
                                            }
                                        case "\n\t":
                                            var v55 uint16 = 8
                                            _ = v55
                                        case "hello":
                                            var v56 string = "hello"
                                            _ = v56
                                        }
                                        v57 := GenS5{F0: -1, F1: -1, F2: 32767}
                                        delete(v1, fi42)
                                        switch fi42 {
                                        case 9223372036854775807:
                                            _ = len(v2)
                                            var v58 int32 = 0
                                            var v59 uint64 = 1
                                            _ = v58
                                            _ = v59
                                        case 1:
                                            var v60 int64 = -831
                                            _ = v60
                                        }
                                        v61 := v2["\n\t"]
                                        _ = v51
                                        _ = v52
                                        _ = v57
                                        _ = v61
                                    }
                                }
                            } else {
                                v16 = v16
                                v16 = v16
                                if ("foo bar" == "foo bar") {
                                    var v62 int = fi38
                                    v6 = (^(v6))
                                    fmt.Println(v16)
                                    _ = v62
                                } else {
                                    _ = len(v2)
                                    if (((("\n\t" + ("foo bar" + ("" + (("foo bar" + "foo bar") + (("\n\t" + "") + "foo bar"))))) + "hello") + "") < "a") {
                                        var v63 uint8 = 204
                                        _ = v63
                                    } else {
                                        fmt.Println(v5, v4, v6)
                                        v64 := GenS1{F0: 138, F1: (fi38 * fi38), F2: "a"}
                                        var v65 int32 = 710
                                        _ = v64
                                        _ = v65
                                    }
                                    v28 = v28
                                }
                            }
                            var v66 int = 127
                            _ = v66
                        } else {
                            fmt.Println(v16)
                        }
                        var v67 bool = v28
                        if true {
                            if (((0.1 + 1.5) / (((0.0 + v4) + 3.14159) / 1.5)) > -2.25) {
                                var v68 int32 = 710
                                _ = v68
                            } else {
                                switch "a" {
                                case "foo bar":
                                    if v67 {
                                        switch "" {
                                        case "":
                                            fmt.Println(v28, v16)
                                        case "a":
                                            if ((v5 >> ((uint(v6)) & 63)) != (v5 & genFn0(v67))) {
                                                var v69 uint = 18446744073709551615
                                                v16 = v16
                                                v70 := GenS6{F0: 42, F1: 718}
                                                _ = v69
                                                _ = v70
                                            } else {
                                                v28 = (0 > 1)
                                                v71 := GenS3{F0: 9223372036854775807}
                                                _ = v71
                                            }
                                        case "\n\t":
                                            switch ("a" + (("a" + ("a" + ((("" + "foo bar") + "a") + "\n\t"))) + ("hello" + "foo bar"))) {
                                            case "foo bar":
                                                var v72 int32 = -2147483648
                                                _ = v72
                                            case "\n\t":
                                                var v73 float64 = float64((1.5 + 0.0))
                                                if (!((-2.25 < float32(v73)) || v67)) {
                                                    _ = len(v3)
                                                    var v74 uint = 19
                                                    _ = v74
                                                }
                                                v67 = v28
                                                _ = v73
                                            }
                                        }
                                    } else {
                                        v28 = v28
                                        if (!(!(v28 && (v28 || true)))) {
                                            if (!false) {
                                                v4 = v4
                                            } else {
                                                var v75 uint64 = 4294967295
                                                if len(v3) > 0 {
                                                    v3[0] = ("" + "hello")
                                                }
                                                var v76 string = (("foo bar" + "foo bar") + ("hello" + "a"))
                                                _ = v75
                                                _ = v76
                                            }
                                            fmt.Println(v16, fi38)
                                        } else {
                                            if v67 {
                                                var v77 bool = (42 > 975)
                                                var v78 uint = 179
                                                if (!(v78 <= (v78 >> (((v78 + (v78 / ((v78) | 1)))) & 63)))) {
                                                    if false {
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                        v16 -= (v16 & (v16 & (-(v16))))
                                                    } else {
                                                        if false {
                                                            delete(v2, ("" + ("hello" + "\n\t")))
                                                        }
                                                        var v79 uint64 = v6
                                                        v6 -= (^(v79))
                                                        _ = v79
                                                    }
                                                } else {
                                                    fmt.Println(v78)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v6 = v6
                                                }
                                                v67 = true
                                                _ = v77
                                                _ = v78
                                            } else {
                                                _ = len(v1)
                                                v80 := GenS2{F0: 0, F1: fi42, F2: v5}
                                                if (-2147483648 > 1) {
                                                    var v81 string = ("\n\t" + ("" + "foo bar"))
                                                    v28 = (0 > 1)
                                                    _ = v81
                                                } else {
                                                    v82 := GenS4{F0: genFn1(547, v5)}
                                                    v6 -= (v6 + v6)
                                                    _ = v82
                                                }
                                                _ = v80
                                            }
                                            var v83 int16 = v16
                                            _ = v83
                                        }
                                    }
                                default:
                                    if v67 {
                                        v4 -= 1e10
                                        v6 = (^(v6))
                                        v84 := GenS4{F0: (v16 * (v16 | (v16 / ((v16) | 1))))}
                                        _ = v84
                                    } else {
                                        var v85 int = (^(fi38))
                                        fmt.Println(v5, v5, v6)
                                        _ = v85
                                    }
                                }
                                v86 := GenS4{F0: -1}
                                v5 -= (-(v5))
                                var v87 bool = (!false)
                                v6 -= (v6 + 224)
                                _ = v86
                                _ = v87
                            }
                        } else {
                            v88 := GenS1{F0: -1, F1: (^(fi42)), F2: "\n\t"}
                            _ = v88
                        }
                        _ = v67
                    } else {
                        v16 = (v16 & genFn1(255, 109))
                        v4 = (v4 / v4)
                        if (!(1e10 < (((0.1 + 1e10) + 3.14159) - 1.5))) {
                            v89 := GenS1{F0: -1, F1: (fi42 / ((128) | 1)), F2: "foo bar"}
                            if true {
                                if true {
                                    var v90 string = ("foo bar" + ("a" + "\n\t"))
                                    var v91 int8 = -93
                                    _ = v90
                                    _ = v91
                                } else {
                                    v4 -= -2.25
                                    v92 := GenS4{F0: v16}
                                    _ = v92
                                }
                                var v93 uint32 = 2
                                v94 := GenS7{F0: false, F1: 0, F2: "hello"}
                                v5 += v5
                                _ = v93
                                _ = v94
                            } else {
                                var v95 float32 = 0.0
                                var v96 uint8 = 0
                                fmt.Println(v5, v28)
                                fmt.Println(v16, v6)
                                _ = v95
                                _ = v96
                            }
                            _ = v89
                        } else {
                            _ = len(v3)
                            if true {
                                if v28 {
                                    v97 := v2[(("foo bar" + ("foo bar" + "")) + "a")]
                                    v98 := GenS1{F0: -2147483648, F1: 42, F2: ("hello" + "foo bar")}
                                    _ = v97
                                    _ = v98
                                } else {
                                    var v99 int16 = -235
                                    v4 = v4
                                    _ = v99
                                }
                                var v100 float32 = 3.14159
                                var v101 int8 = 42
                                _ = v100
                                _ = v101
                            } else {
                                var v102 float64 = v4
                                _ = v102
                            }
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        }
                    }
                    if ((v5 & v5) > v5) {
                        switch int(v16) {
                        case -299:
                            var v103 int16 = v16
                            _ = v103
                        case -921:
                            v104 := v1[int(v16)]
                            _ = v104
                        }
                        v2["hello"] = 42
                    }
                    var v105 uint16 = 2
                    _ = fi42
                    _ = v105
                }
            case 127:
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi38
        }
    case 189:
        var v106 int = -9223372036854775808
        _ = v106
    default:
        fmt.Println(v16)
    }
    v5 = (v5 / ((v5) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
