package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 uint8; F2 int8 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32 }
func genFn0() bool {
    v6 := []int{270}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = 1.5
        fmt.Println(v8)
        var v9 int = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: -1}
    v11 := GenS2{F0: 32767, F1: "a"}
    switch "foo bar" {
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "hello":
        switch (("\n\t" + "") + "") {
        case "hello":
            var v12 uint32 = 897
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 12; fi13++ {
                _ = len(v6)
                var v14 float64 = -2.25
                v14 = v14
                _ = fi13
                _ = v14
            }
        case "a":
            var v15 int8 = 93
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint = 2
            _ = fi16
            _ = v17
        }
    }
    _ = v10
    _ = v11
    _ = v6
    return (0 != 251)
}
func genFn1(p0 int8) uint8 {
    v18 := GenS3{F0: "a", F1: (128 == 857), F2: 365}
    v19 := []int{}
    for fi20 := 0; fi20 < 15; fi20++ {
        v19 = append(v19, 32767)
        if false {
            switch "a" {
            case "\n\t":
                var v21 bool = true
                _ = v21
            case "a":
                var v22 int16 = -923
                _ = v22
            case "hello":
                for fi23 := 0; fi23 < 5; fi23++ {
                    if false {
                        fmt.Println(fi23)
                    }
                    p0 = int8(fi23)
                    _ = fi23
                }
            }
            var v24 uint32 = 2147483648
            _ = v24
        }
        fmt.Println(fi20)
        _ = fi20
    }
    _ = p0
    _ = v18
    _ = v19
    return 128
}

func main() {
    var v1 uint8 = 127
    var v2 float32 = (float32((3.14159 - 3.14159)) * (3.14159 / 1e10))
    var v3 uint64 = 867
    v4 := []string{(((("a" + "hello") + ("" + "foo bar")) + "\n\t") + "\n\t"), "hello", ("foo bar" + (("hello" + "a") + ""))}
    v5 := make(map[string]uint32)
    var v25 bool = ((1e10 >= float64(v2)) && (((!false) || (1e10 > (1e10 + 0.0))) && false))
    fmt.Println(v2, v25)
    if v25 {
        v1 -= 0
        var v26 uint8 = 2
        _ = v26
    } else {
        switch -700 {
        case 1:
            fmt.Println(v2)
        case 688:
            var v27 int8 = 127
            _ = v27
        default:
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := v5["foo bar"]
                v30 := GenS4{F0: v1}
                v2 += v2
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        var v31 float32 = (((v2 * v2) / 1.5) / v2)
        _ = v31
    }
    switch 42 {
    case 9223372036854775807:
        var v32 int64 = -9223372036854775808
        _ = v32
    case 32768:
        if (-116 >= int8(v1)) {
            fmt.Println(v25, v25)
            for fi33 := 0; fi33 < 1; fi33++ {
                fmt.Println(v2, v25)
                _ = fi33
            }
            if (!(1 < -110)) {
                for fi34 := 0; fi34 < 2; fi34++ {
                    var v35 int8 = 0
                    v35 += 42
                    var v36 uint = 255
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                v5["foo bar"] = uint32(v1)
                if len(v4) > 0 {
                    v4[0] = ("hello" + "\n\t")
                }
            } else {
                switch "\n\t" {
                case "hello":
                    v5["foo bar"] = 0
                case "":
                    if v25 {
                        v37 := GenS4{F0: v1}
                        switch -1 {
                        case 1:
                            if false {
                                for fi38 := 0; fi38 < 20; fi38++ {
                                    var v39 string = ""
                                    _ = len(v5)
                                    v1 -= (v1 ^ 127)
                                    _ = fi38
                                    _ = v39
                                }
                                v40 := GenS5{F0: v3, F1: genFn1(-97), F2: 675}
                                _ = v40
                            } else {
                                fmt.Println(v2)
                                var v41 uint8 = (^(v1))
                                var v42 uint8 = genFn1(-128)
                                _ = v41
                                _ = v42
                            }
                        case 0:
                            if (2 == 2) {
                                for fi43 := 0; fi43 < 18; fi43++ {
                                    v44 := v5["a"]
                                    for fi45 := 0; fi45 < 10; fi45++ {
                                        v4 = append(v4, "a")
                                        v46 := GenS6{F0: v44, F1: 2}
                                        _ = fi45
                                        _ = v46
                                    }
                                    _ = fi43
                                    _ = v44
                                }
                            } else {
                                v5["\n\t"] = 521
                                v47 := GenS1{F0: 738, F1: -1}
                                _ = v47
                            }
                        case 127:
                            _ = len(v4)
                        default:
                            v48 := GenS1{F0: 1, F1: -42}
                            _ = v48
                        }
                        if ((v25 && ((uint16(v3) >= 1) && v25)) && v25) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v1 -= 127
                            v2 = -2.25
                        }
                        _ = v37
                    } else {
                        fmt.Println(v1, v3, v1)
                        if (-9223372036854775808 != -503) {
                            var v49 bool = v25
                            var v50 uint8 = (^(v1))
                            _ = v49
                            _ = v50
                        }
                    }
                case "\n\t":
                    if false {
                        if v25 {
                            for fi51 := 0; fi51 < 6; fi51++ {
                                var v52 float64 = (3.14159 - 0.0)
                                switch fi51 {
                                case 4294967295:
                                    v2 = v2
                                default:
                                    if ((!(!v25)) || v25) {
                                        v5["foo bar"] = 2
                                        v53 := GenS7{F0: 42, F1: (v3 & v3)}
                                        var v54 uint64 = v3
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v25 = ((float64((3.14159 * (v2 + v2))) - (1.5 * ((3.14159 + v52) - 0.0))) <= v52)
                                        v55 := GenS8{F0: -128}
                                        var v56 uint64 = 18446744073709551615
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                v52 = ((float64(-2.25) * v52) - v52)
                                _ = fi51
                                _ = v52
                            }
                        }
                        var v57 int8 = -128
                        _ = v57
                    }
                }
                switch -959 {
                case 1:
                    v58 := GenS9{F0: -66, F1: v1, F2: -1}
                    _ = v58
                case 0:
                    v3 = v3
                case -184:
                    if (1 > 2) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        fmt.Println(v2)
                        if (("hello" != ("foo bar" + "foo bar")) || ((-32768 == -1) || v25)) {
                            switch 202 {
                            case -896:
                                switch 32768 {
                                case -1:
                                    var v59 uint64 = (v3 >> ((506) & 63))
                                    _ = v59
                                case -951:
                                    if (854 < 2) {
                                        delete(v5, "hello")
                                        switch -9223372036854775808 {
                                        case -2147483648:
                                            if v25 {
                                                v60 := GenS6{F0: 1, F1: 65535}
                                                _ = v60
                                            } else {
                                                v61 := GenS10{F0: 795, F1: 42}
                                                for fi62 := 0; fi62 < 5; fi62++ {
                                                    v3 += 0
                                                    _ = fi62
                                                }
                                                _ = v61
                                            }
                                        case 42:
                                            for fi63 := 0; fi63 < 1; fi63++ {
                                                var v64 uint16 = 298
                                                _ = fi63
                                                _ = v64
                                            }
                                        case -9223372036854775808:
                                            if ((!(!(!true))) || v25) {
                                                v4 = append(v4, (("\n\t" + "foo bar") + "foo bar"))
                                            }
                                        }
                                        fmt.Println(v2, v3)
                                    }
                                }
                            case 127:
                                var v65 uint16 = 32767
                                _ = v65
                            case 301:
                                if (!false) {
                                    v66 := v5["hello"]
                                    v67 := GenS2{F0: -151, F1: ""}
                                    _ = v66
                                    _ = v67
                                }
                            }
                            v68 := GenS11{F0: 65535}
                            _ = v68
                        } else {
                            v25 = (!(-1 <= 1))
                            v2 = ((v2 * v2) * v2)
                            var v69 uint64 = 1
                            _ = v69
                        }
                        fmt.Println(v1, v3, v3)
                    }
                }
            }
        }
    }
    v25 = v25
    var v70 uint8 = 1
    var v71 float32 = v2
    for fi72 := 0; fi72 < 20; fi72++ {
        if v25 {
            v73 := v5[("\n\t" + "hello")]
            _ = v73
        }
        v2 += v2
        _ = fi72
    }
    v74 := GenS12{F0: -980}
    for fi75 := 0; fi75 < 19; fi75++ {
        var v76 bool = (v25 && (genFn1(-105) == genFn1(-128)))
        _ = fi75
        _ = v76
    }
    if (false && v25) {
        var v77 int = 778
        v78 := GenS10{F0: 646, F1: 0}
        _ = v77
        _ = v78
    } else {
        if true {
            v2 = v71
            v70 = v70
        } else {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            var v79 float64 = (float64(float32((1.5 - (3.14159 / (float64(((0.0 * ((-2.25 * (v2 + 0.1)) - (float32(float64((float32(-2.25) * v2))) * v2))) + 1e10)) - ((3.14159 / (-2.25 * 0.1)) * 0.0)))))) - (0.0 - (-2.25 + (1e10 + 1.5))))
            _ = v79
        }
        var v80 uint64 = 4294967295
        _ = v80
    }
    v81 := GenS2{F0: 896, F1: ("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + "\n\t")))}
    var v82 int16 = 32767
    if (1 == 127) {
        v5["\n\t"] = 2147483648
        var v83 int16 = (v82 ^ v82)
        _ = len(v4)
        _ = v83
    } else {
        if true {
            if false {
                if false {
                    if false {
                        if true {
                            v2 = (v2 * (float32((float64(3.14159) - 1e10)) / (0.1 / (1e10 + (float32(-2.25) - v2)))))
                            v70 += 42
                            var v84 float64 = (1.5 - 3.14159)
                            _ = v84
                        }
                        if v25 {
                            switch 922 {
                            case -23:
                                var v85 float32 = v2
                                _ = v85
                            case 32768:
                                var v86 int8 = 42
                                _ = v86
                            }
                            v82 += (v82 / ((v82) | 1))
                        }
                    }
                    for fi87 := 0; fi87 < 2; fi87++ {
                        if len(v4) > 0 {
                            v4[0] = ("a" + "")
                        }
                        _ = fi87
                    }
                }
                v1 = v70
                _ = len(v4)
            } else {
                v88 := GenS12{F0: -446}
                v25 = (2 <= 2)
                var v89 uint = 420
                _ = v88
                _ = v89
            }
            switch "\n\t" {
            case "":
                v82 = (v82 | v82)
            case "a":
                v3 -= v3
            }
        }
    }
    var v90 uint = 0
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v3)
    }
    var v91 uint = v90
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v82:", v82)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
    _ = v74
    _ = v81
    _ = v82
    _ = v90
    _ = v91
}
