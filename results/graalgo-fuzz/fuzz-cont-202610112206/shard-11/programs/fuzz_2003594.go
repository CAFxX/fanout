package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int32, p1 int64, p2 int8) uint64 {
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{fi7, (fi7 ^ 42), (^(fi7))}
        if (0 < 42) {
            p2 -= (p2 * 100)
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (v6 & (v6 | (v6 | (v6 >> ((0) & 63)))))
}

func main() {
    v1 := []string{"a", "hello", ("" + ("" + ("foo bar" + "a")))}
    v2 := []string{}
    var v3 float32 = 3.14159
    var v4 float64 = 1e10
    v5 := make(map[uint32]string)
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint32 = 538
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 255}
            v1 = append(v1, "\n\t")
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = len(v5)
    _ = len(v1)
    fmt.Println(v3)
    if false {
        switch ("\n\t" + ("foo bar" + "hello")) {
        case "hello":
            v4 += 3.14159
        }
    } else {
        fmt.Println(v3, v4)
    }
    var v13 uint64 = genFn0(1, 42, 42)
    for fi14 := 0; fi14 < 18; fi14++ {
        if true {
            var v15 uint = 0
            var v16 uint = v15
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(fi14)
                _ = len(v5)
                fmt.Println(v3, fi17, v4)
                var v18 int8 = -1
                delete(v5, 42)
                _ = fi17
                _ = v18
            }
        }
        switch 128 {
        case 127:
            var v19 bool = (0 <= 442)
            _ = v19
        default:
            v13 += genFn0(0, -384, -1)
        }
        _ = fi14
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + ("a" + ("a" + "")))
    }
    var v20 int16 = 32767
    if true {
        v20 -= 32767
    } else {
        v1 = append(v1, ("foo bar" + "a"))
    }
    if false {
        v21 := GenS2{F0: int(v20), F1: 42}
        v20 = -850
        var v22 int32 = 0
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        switch "hello" {
        case "foo bar":
            var v24 uint8 = 127
            _ = v24
        case "a":
            var v25 int = (^(fi23))
            _ = v25
        default:
            var v26 uint32 = 0
            _ = v26
        }
        var v27 bool = false
        for fi28 := 0; fi28 < 10; fi28++ {
            if ((0.0 * (v4 / v4)) == (v4 / 3.14159)) {
                var v29 uint16 = 564
                _ = len(v5)
                _ = v29
            }
            var v30 float64 = float64((v3 + v3))
            v30 -= v4
            var v31 int32 = -2147483648
            if true {
                v31 += v31
                if (false || (((v30 + v30) == v4) && true)) {
                    v20 += (v20 | -886)
                    if v27 {
                        var v32 int16 = 1
                        fmt.Println(fi28)
                        if true {
                            var v33 uint8 = 0
                            if (!false) {
                                delete(v5, 42)
                                v3 -= float32(3.14159)
                                var v34 int = 1
                                switch 32768 {
                                case 403:
                                    if (0 >= 42) {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        v35 := GenS1{F0: 2}
                                        if v27 {
                                            switch ("\n\t" + "foo bar") {
                                            case "foo bar":
                                                v33 += 255
                                            case "":
                                                var v36 int = (fi28 + (^(v34)))
                                                _ = v36
                                            case "hello":
                                                var v37 int32 = v31
                                                _ = v37
                                            }
                                            if (v32 >= (v20 - 0)) {
                                                if (255 != uint16(v33)) {
                                                    v20 = (v32 ^ (v20 << ((178) & 63)))
                                                }
                                                switch fi23 {
                                                case 965:
                                                    var v38 int16 = v20
                                                    _ = v38
                                                case 127:
                                                    fmt.Println(v3, v27, v20)
                                                default:
                                                    if false {
                                                        fmt.Println(fi28)
                                                        v32 += (v32 % (((v32 - (^(v20)))) | 1))
                                                        var v39 int32 = (v31 << ((18446744073709551615) & 63))
                                                        var v40 uint16 = 65535
                                                        v34 = (^(fi28))
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        var v41 int32 = (-(v31))
                                                        v4 += v4
                                                        if (false || (("" != ("hello" + ("" + ((("" + (((("" + (((("a" + ("hello" + "hello")) + ("foo bar" + "a")) + ("hello" + "")) + "foo bar")) + "hello") + "hello") + "foo bar")) + "foo bar") + "\n\t")))) && true)) {
                                                            if (!true) {
                                                                fmt.Println(v30)
                                                                v42 := GenS1{F0: 73}
                                                                _ = v42
                                                            }
                                                            var v43 uint16 = 301
                                                            _ = v43
                                                        }
                                                        var v44 float64 = v4
                                                        _ = v41
                                                        _ = v44
                                                    }
                                                }
                                            } else {
                                                v3 = 1e10
                                                var v45 int32 = v31
                                                _ = v45
                                            }
                                            var v46 float64 = ((float64(v3) * 0.0) - (3.14159 - v4))
                                            if (1 <= 2147483648) {
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(v31)
                                                case "":
                                                    var v47 uint = 255
                                                    _ = v47
                                                case "foo bar":
                                                    if (!((false || ((!v27) && v27)) || ((2147483647 >= uint32(v13)) || ("hello" <= "\n\t")))) {
                                                        v5[2147483647] = ""
                                                        v48 := GenS1{F0: 0}
                                                        _ = v48
                                                    }
                                                default:
                                                    fmt.Println(v46)
                                                }
                                                fmt.Println(v32, v34)
                                            }
                                            if (v34 == (-(fi23))) {
                                                var v49 uint8 = 108
                                                switch fi28 {
                                                case -840:
                                                    v50 := GenS3{F0: 1, F1: ""}
                                                    _ = v50
                                                case 32768:
                                                    var v51 float32 = v3
                                                    _ = v51
                                                }
                                                var v52 int32 = (v31 + v31)
                                                _ = v49
                                                _ = v52
                                            } else {
                                                if (156 != 65535) {
                                                    if false {
                                                        v53 := GenS2{F0: int(v13), F1: 115}
                                                        v46 = v4
                                                        switch -35 {
                                                        case 4294967295:
                                                            v13 = 2
                                                        case -2147483648:
                                                            fmt.Println(v46)
                                                        default:
                                                            var v54 bool = (!(!false))
                                                            _ = v54
                                                        }
                                                        var v55 bool = v27
                                                        if (true && (0 < 936)) {
                                                            v56 := GenS2{F0: (fi28 >> ((65535) & 63)), F1: 1}
                                                            v57 := GenS1{F0: 4294967295}
                                                            var v58 string = ("foo bar" + ("" + "foo bar"))
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            v59 := GenS1{F0: 0}
                                                            _ = v59
                                                        }
                                                        _ = v53
                                                        _ = v55
                                                    }
                                                    v30 -= float64((float32(v30) - 1e10))
                                                    v60 := GenS3{F0: int8(v31), F1: "foo bar"}
                                                    _ = v60
                                                }
                                            }
                                            _ = v46
                                        }
                                        _ = v35
                                    }
                                case 32767:
                                    v61 := GenS2{F0: -9223372036854775808, F1: 5}
                                    _ = v61
                                default:
                                    if false {
                                        fmt.Println(fi23, v4)
                                    } else {
                                        var v62 uint64 = v13
                                        var v63 uint = 1
                                        v3 -= v3
                                        v64 := GenS3{F0: 0, F1: "\n\t"}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                switch ("foo bar" + ("hello" + "a")) {
                                case "a":
                                    var v65 int = (v34 * -76)
                                    _ = v65
                                case "hello":
                                    var v66 int32 = (v31 + (v31 >> ((uint(v13)) & 63)))
                                    _ = v66
                                case "":
                                    var v67 uint32 = uint32(v13)
                                    _ = v67
                                default:
                                    v34 = fi28
                                }
                                _ = v34
                            } else {
                                _ = len(v5)
                                if (111 > 42) {
                                    if v27 {
                                        var v68 uint64 = (v13 - genFn0((v31 & (^(v31))), 9223372036854775807, -45))
                                        var v69 float64 = 1e10
                                        v70 := GenS2{F0: (fi28 - fi28), F1: -128}
                                        fmt.Println(v3, v69)
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    } else {
                                        var v71 uint = 18446744073709551615
                                        _ = v71
                                    }
                                    v30 = 1.5
                                    _ = len(v5)
                                } else {
                                    v32 = 413
                                    v72 := GenS2{F0: fi28, F1: -77}
                                    v4 += (((v30 + v30) / (v30 - v4)) + 1e10)
                                    _ = v72
                                }
                                switch 602 {
                                case 127:
                                    v20 -= 1
                                case 32767:
                                    if v27 {
                                        v3 = v3
                                        if (((("hello" + "\n\t") + "hello") + "\n\t") != "hello") {
                                            var v73 string = "hello"
                                            var v74 uint16 = 439
                                            _ = v73
                                            _ = v74
                                        } else {
                                            _ = len(v5)
                                            if (832 > 2147483648) {
                                                var v75 uint = 65535
                                                var v76 uint32 = 0
                                                v13 = 0
                                                _ = v75
                                                _ = v76
                                            }
                                            if true {
                                                var v77 float64 = v4
                                                _ = v77
                                            }
                                            var v78 string = "hello"
                                            _ = v78
                                        }
                                    } else {
                                        v79 := GenS4{F0: 9223372036854775807, F1: 811}
                                        v3 = v3
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        _ = v79
                                    }
                                case 32768:
                                    v80 := GenS5{F0: 4294967295}
                                    _ = v80
                                default:
                                    v81 := GenS5{F0: 12}
                                    _ = v81
                                }
                            }
                            fmt.Println(v31, v33)
                            fmt.Println(fi28)
                            _ = v33
                        } else {
                            v82 := GenS6{F0: 4294967295}
                            if true {
                                v4 += v30
                            } else {
                                _ = len(v5)
                            }
                            v13 -= genFn0((-(v31)), 9223372036854775807, -63)
                            if ((!(!false)) && (((((((float32(0.0) + 1e10) / v3) / float32(-2.25)) * 1e10) * v3) / -2.25) < v3)) {
                                v83 := GenS7{F0: v13}
                                v20 = (v20 / ((177) | 1))
                                _ = v83
                            }
                            var v84 bool = v27
                            _ = v82
                            _ = v84
                        }
                        _ = v32
                    }
                } else {
                    v3 = float32(1.5)
                    v85 := GenS6{F0: 2}
                    _ = v85
                }
                var v86 float32 = float32(v4)
                _ = v86
            }
            _ = fi28
            _ = v30
            _ = v31
        }
        _ = fi23
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
