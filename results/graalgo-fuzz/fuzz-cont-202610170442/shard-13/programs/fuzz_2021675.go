package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int }
func genFn0(p0 uint64) int32 {
    p0 += 360
    fmt.Println(p0)
    _ = p0
    return 2147483647
}
func genFn1(p0 uint64, p1 int32, p2 bool) uint16 {
    v6 := []int{-9223372036854775808, -2147483648, -744, -9223372036854775808}
    if p2 {
        var v7 uint16 = 316
        _ = v7
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 591
}

func main() {
    var v1 bool = false
    var v2 int32 = 2147483647
    var v3 int16 = 661
    v4 := []uint32{1, 0, 503, 2147483648}
    v5 := []uint32{2147483648, 2147483647, 239, 4294967295}
    v3 += (v3 & v3)
    for fi8 := 0; fi8 < 2; fi8++ {
        if false {
            switch (-(fi8)) {
            case 127:
                var v9 int64 = 42
                _ = v9
            }
        } else {
            if true {
                v2 -= v2
            } else {
                if (((3.14159 - (0.1 / 1e10)) < 1.5) && (65535 == 42)) {
                    var v10 int = (fi8 >> ((500) & 63))
                    _ = v10
                } else {
                    var v11 uint16 = 255
                    var v12 uint32 = 940
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 float64 = 0.0
                    switch -9223372036854775808 {
                    case 0:
                        v2 -= v2
                    default:
                        var v15 int8 = 0
                        _ = v15
                    }
                    v16 := GenS1{F0: "hello"}
                    var v17 uint64 = 42
                    var v18 int8 = 1
                    if (255 != genFn1((v17 % (((v17 / (((^(v17))) | 1))) | 1)), (v2 % ((v2) | 1)), (42 == 2))) {
                        if true {
                            var v19 float32 = float32(v14)
                            if len(v4) > 0 {
                                v4[0] = 362
                            }
                            _ = v19
                        } else {
                            var v20 uint = 497
                            _ = v20
                        }
                        v21 := GenS1{F0: ""}
                        if false {
                            if ((-70 >= 1) && v1) {
                                if v1 {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    if false {
                                        fmt.Println(v1, v14, v17)
                                    }
                                }
                                if false {
                                    v3 = v3
                                }
                                v3 = (v3 >> ((18446744073709551615) & 63))
                            }
                            var v22 uint32 = 2
                            _ = v22
                        } else {
                            fmt.Println(v18, v17, v17)
                            fmt.Println(fi8, v1)
                            fmt.Println(fi8)
                        }
                        v17 -= 2
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v23 := GenS2{F0: (v2 % (((v2 >> ((0) & 63))) | 1)), F1: 379}
                        v24 := GenS2{F0: v2, F1: 18446744073709551615}
                        v25 := GenS3{F0: "foo bar", F1: (v17 + (v17 * v17))}
                        switch "" {
                        case "foo bar":
                            v17 += (v17 >> ((800) & 63))
                        case "":
                            var v26 int32 = 1
                            _ = v26
                        default:
                            v27 := GenS4{F0: "foo bar", F1: ""}
                            _ = v27
                        }
                        fmt.Println(v3, v2, v3)
                        switch fi13 {
                        case 463:
                            var v28 bool = true
                            _ = v28
                        case 32767:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v29 uint = 65535
                        _ = v21
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v29
                    }
                    if (0.1 == 1e10) {
                        var v30 float32 = float32(3.14159)
                        _ = v30
                    } else {
                        if (!true) {
                            v1 = v1
                        }
                        if v1 {
                            v31 := GenS2{F0: (-(v2)), F1: 18446744073709551615}
                            v32 := GenS2{F0: v2, F1: 0}
                            _ = len(v5)
                            _ = v31
                            _ = v32
                        } else {
                            if v1 {
                                fmt.Println(fi13)
                                v17 -= 142
                                v33 := GenS4{F0: "foo bar", F1: "\n\t"}
                                _ = v33
                            } else {
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                }
                                v17 -= (v17 >> ((869) & 63))
                                v17 = 685
                            }
                            v14 += float64(float32(((float64(float32((((v14 - 1.5) / 1.0) * float64(1.5)))) + v14) - (((float64(1e10) * ((0.1 / (0.1 / 3.14159)) + (float64(1e10) / v14))) * v14) - (((v14 + v14) * (float64(float32(1e10)) - (v14 / v14))) * v14)))))
                        }
                        v4 = append(v4, 65535)
                    }
                    _ = fi13
                    _ = v14
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            if (65535 < 42) {
                v1 = false
                var v34 int = 127
                v35 := GenS4{F0: "", F1: "hello"}
                _ = v34
                _ = v35
            } else {
                v36 := GenS4{F0: "", F1: "\n\t"}
                var v37 int64 = 305
                if (false || (!((!true) || false))) {
                    fmt.Println(v37, v2, v37)
                    v38 := GenS1{F0: ("a" + (("a" + ("hello" + "\n\t")) + "a"))}
                    if (!v1) {
                        switch "\n\t" {
                        case "a":
                            var v39 float64 = 1.5
                            _ = v39
                        case "\n\t":
                            if (127 != 255) {
                                var v40 int = (fi8 / (((fi8 % (((fi8 + (^(fi8)))) | 1))) | 1))
                                v3 = v3
                                _ = v40
                            } else {
                                if (false && (false || (v1 && (127 < 1)))) {
                                    for fi41 := 0; fi41 < 19; fi41++ {
                                        if ((-2.25 - (0.0 / float32(1e10))) <= 0.1) {
                                            if (!(!((v1 && v1) || v1))) {
                                                if v1 {
                                                    fmt.Println(v37)
                                                    if v1 {
                                                        var v42 int32 = 0
                                                        _ = v42
                                                    }
                                                    v2 -= v2
                                                } else {
                                                    var v43 int8 = -24
                                                    switch (("" + ("hello" + ("foo bar" + ("\n\t" + ("hello" + (("a" + "hello") + "")))))) + ("a" + "foo bar")) {
                                                    case "\n\t":
                                                        var v44 int32 = (^(v2))
                                                        _ = v44
                                                    default:
                                                        fmt.Println(v1, fi41, v37)
                                                    }
                                                    _ = v43
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = 630
                                                }
                                                v3 = (v3 >> ((0) & 63))
                                            } else {
                                                fmt.Println(v3)
                                                if len(v5) > 0 {
                                                    v5[0] = 0
                                                }
                                            }
                                        }
                                        var v45 uint32 = 741
                                        v46 := GenS5{F0: 152, F1: v45}
                                        _ = len(v4)
                                        switch "hello" {
                                        case "\n\t":
                                            switch (fi8 << ((0) & 63)) {
                                            case 362:
                                                fmt.Println(v2, fi8)
                                            case 9223372036854775807:
                                                fmt.Println(v1)
                                            case 686:
                                                fmt.Println(v1, fi8)
                                            default:
                                                v2 -= -2147483648
                                            }
                                        case "foo bar":
                                            var v47 int8 = 42
                                            v4 = append(v4, (^(v45)))
                                            switch "foo bar" {
                                            case "\n\t":
                                                v48 := GenS2{F0: (v2 + v2), F1: 1}
                                                _ = v48
                                            case "hello":
                                                switch ("foo bar" + ("\n\t" + "a")) {
                                                case "a":
                                                    v49 := GenS6{F0: 1, F1: 128}
                                                    v5 = append(v5, v45)
                                                    var v50 int = (fi8 & (fi41 + fi8))
                                                    v4 = append(v4, 4294967295)
                                                    v37 += 420
                                                    _ = v49
                                                    _ = v50
                                                default:
                                                    var v51 string = ((((("a" + "foo bar") + "\n\t") + (((("\n\t" + ("a" + ("foo bar" + (("hello" + "\n\t") + (("\n\t" + "foo bar") + "\n\t"))))) + "\n\t") + "") + (("foo bar" + "") + "hello"))) + "") + "hello")
                                                    switch 32768 {
                                                    case 128:
                                                        if (2 <= 18446744073709551615) {
                                                            var v52 uint8 = 1
                                                            var v53 bool = ("a" > "")
                                                            var v54 float32 = 1e10
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                    default:
                                                        var v55 int = fi41
                                                        _ = v55
                                                    }
                                                    v56 := GenS1{F0: v51}
                                                    _ = v51
                                                    _ = v56
                                                }
                                            default:
                                                v57 := GenS1{F0: "hello"}
                                                _ = v57
                                            }
                                            _ = v47
                                        case "hello":
                                            v45 = 2147483647
                                        default:
                                            var v58 uint = 255
                                            _ = v58
                                        }
                                        _ = fi41
                                        _ = v45
                                        _ = v46
                                    }
                                    var v59 uint32 = 2
                                    _ = v59
                                } else {
                                    v1 = (1 > 0)
                                }
                                for fi60 := 0; fi60 < 10; fi60++ {
                                    v61 := GenS4{F0: "", F1: "a"}
                                    _ = fi60
                                    _ = v61
                                }
                            }
                        default:
                            if v1 {
                                v37 = (^(v37))
                            } else {
                                _ = len(v4)
                                switch "foo bar" {
                                case "foo bar":
                                    for fi62 := 0; fi62 < 14; fi62++ {
                                        if true {
                                            _ = len(v5)
                                        } else {
                                            v63 := GenS2{F0: -528, F1: 2}
                                            fmt.Println(v37)
                                            switch fi62 {
                                            case 32767:
                                                fmt.Println(v3, v1, v37)
                                            }
                                            var v64 uint = 42
                                            _ = v63
                                            _ = v64
                                        }
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                v65 := GenS1{F0: ("a" + ("\n\t" + ""))}
                                                _ = v65
                                            default:
                                                v1 = v1
                                            }
                                            fmt.Println(fi8)
                                            var v66 int64 = -9223372036854775808
                                            var v67 float64 = 1e10
                                            _ = v66
                                            _ = v67
                                        } else {
                                            var v68 float64 = float64(((float32(-2.25) - 1.5) / 1.0))
                                            _ = v68
                                        }
                                        _ = fi62
                                    }
                                default:
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            v1 = (true && (!(!(v1 && (1 < 322)))))
                                        }
                                    }
                                }
                            }
                        }
                        v1 = (v1 && (v1 && v1))
                        v1 = v1
                    }
                    _ = v38
                }
                _ = v36
                _ = v37
            }
        }
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v1 = ((-2.25 > 3.14159) || v1)
    v69 := GenS2{F0: (v2 / (((v2 ^ (v2 ^ v2))) | 1)), F1: 65535}
    v3 = v3
    if true {
        var v70 uint64 = 18446744073709551615
        fmt.Println(v3, v2, v2)
        if true {
            v71 := GenS1{F0: ((("foo bar" + "hello") + ("hello" + ("\n\t" + "hello"))) + "a")}
            v72 := GenS7{F0: 32768}
            _ = v71
            _ = v72
        }
        _ = v70
    }
    v73 := GenS6{F0: 42, F1: 1}
    v3 += (v3 - (v3 | v3))
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v73
}
