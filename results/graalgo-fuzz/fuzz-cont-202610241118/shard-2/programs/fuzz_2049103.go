package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 string; F1 uint8; F2 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int64 }
func genFn0(p0 bool, p1 int8, p2 int8) int16 {
    fmt.Println(p1, p0, p1)
    var v6 int16 = -32768
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(v6))
}
func genFn1() uint64 {
    v8 := GenS2{F0: true}
    if (("a" + ("hello" + (("a" + "a") + ((("foo bar" + ("a" + "")) + "a") + "hello")))) > "") {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint64 = 751
            for fi11 := 0; fi11 < 12; fi11++ {
                switch (fi9 & (fi11 / ((fi9) | 1))) {
                case 0:
                    v12 := GenS3{F0: 18446744073709551615, F1: 2, F2: 2147483648}
                    _ = v12
                case 4294967295:
                    v13 := []int{(fi11 | (-(fi11))), (fi11 | 32767), 1}
                    _ = v13
                default:
                    v10 -= v10
                }
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    var v14 int8 = -1
    if (!(255 == 2)) {
        if (-109 > -44) {
            switch 32768 {
            case 2147483647:
                if (-1 == -1) {
                    v14 = (^(v14))
                    var v15 string = "\n\t"
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v15 = "\n\t"
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 int = (fi17 - (fi17 % ((fi16) | 1)))
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                    _ = v15
                }
            case 0:
                switch "a" {
                case "foo bar":
                    v14 -= (v14 % (((v14 ^ v14)) | 1))
                case "\n\t":
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v14 = -1
                        _ = fi19
                    }
                case "hello":
                    var v20 uint32 = 1
                    _ = v20
                }
            case 127:
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch (fi21 * (fi21 | (fi21 & (fi21 & fi21)))) {
                    case 1:
                        v22 := []int{fi21}
                        _ = v22
                    case -2147483648:
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v14 += v14
                            fmt.Println(v14, fi23)
                            _ = fi23
                        }
                    case -1:
                        switch "" {
                        case "a":
                            v14 += 1
                        case "hello":
                            if false {
                                if false {
                                    v14 = (v14 * 42)
                                    for fi24 := 0; fi24 < 18; fi24++ {
                                        if (genFn0(false, v14, v14) < -1) {
                                            fmt.Println(v14)
                                        }
                                        _ = fi24
                                    }
                                } else {
                                    var v25 uint = 255
                                    var v26 uint16 = 32767
                                    v27 := GenS3{F0: v25, F1: v25, F2: 395}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                var v28 int64 = 150
                                var v29 string = (("\n\t" + ("" + "\n\t")) + (("foo bar" + "hello") + ""))
                                _ = v28
                                _ = v29
                            }
                        case "foo bar":
                            v30 := []int{fi21}
                            _ = v30
                        default:
                            fmt.Println(fi21, fi21)
                        }
                    default:
                        v31 := []int{fi21}
                        _ = v31
                    }
                    _ = fi21
                }
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                for fi33 := 0; fi33 < 4; fi33++ {
                    var v34 string = "foo bar"
                    v35 := GenS4{F0: 588, F1: (1e10 < 0.1), F2: 766}
                    v14 -= v14
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                if (2147483647 >= 1) {
                    v36 := GenS4{F0: 343, F1: true, F2: 33}
                    switch ("" + "foo bar") {
                    case "hello":
                        var v37 uint = 478
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 12; fi38++ {
                            v39 := []int{fi38, fi32, fi38}
                            if true {
                                if (!true) {
                                    var v40 uint16 = 65535
                                    var v41 uint8 = 2
                                    _ = v40
                                    _ = v41
                                } else {
                                    v42 := GenS4{F0: 65535, F1: ((727 < 2147483648) || false), F2: 2}
                                    if ((34 < 65535) && (!true)) {
                                        var v43 int32 = 1
                                        var v44 int = 32768
                                        _ = v43
                                        _ = v44
                                    } else {
                                        v14 = (v14 >> ((0) & 63))
                                        var v45 int64 = 9223372036854775807
                                        if len(v39) >= 2 {
                                            v39 = v39[1:len(v39)]
                                        }
                                        _ = v45
                                    }
                                    _ = v42
                                }
                                v46 := GenS4{F0: 42, F1: true, F2: 247}
                                if true {
                                    v14 = (v14 | v14)
                                    v47 := GenS5{F0: 1, F1: (fi38 | fi32)}
                                    var v48 uint = 995
                                    _ = v47
                                    _ = v48
                                }
                                _ = v46
                            } else {
                                if len(v39) >= 2 {
                                    v39 = v39[1:len(v39)]
                                }
                            }
                            v49 := GenS1{F0: -708}
                            v14 = v14
                            _ = fi38
                            _ = v39
                            _ = v49
                        }
                    }
                    _ = v36
                } else {
                    fmt.Println(v14)
                }
                _ = fi32
            }
        }
        switch ("a" + ("" + ("\n\t" + "a"))) {
        case "a":
            if true {
                var v50 float32 = 0.1
                v14 -= -12
                var v51 uint16 = 973
                _ = v50
                _ = v51
            } else {
                v52 := []int{-9223372036854775808, -50, 4294967295, 652}
                _ = v52
            }
        default:
            v14 -= 42
        }
        if ((0 > 32767) || (true || (0 <= 42))) {
            v53 := []int{32767, -826, 38, 4294967295}
            if len(v53) >= 2 {
                v53 = v53[1:len(v53)]
            }
            _ = v53
        }
    } else {
        v54 := GenS4{F0: 2, F1: ((("" == "\n\t") && false) || true), F2: 255}
        _ = v54
    }
    for fi55 := 0; fi55 < 8; fi55++ {
        for fi56 := 0; fi56 < 6; fi56++ {
            v57 := GenS6{F0: "", F1: 128, F2: 42}
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    _ = v14
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 160
    var v3 int32 = 2147483647
    var v4 int8 = 42
    v5 := []int{-2147483648}
    v5 = append(v5, 9223372036854775807)
    v58 := GenS4{F0: 2, F1: (2147483647 != 128), F2: 4294967295}
    switch -9223372036854775808 {
    case 127:
        if true {
            var v59 float64 = (1e10 + -2.25)
            var v60 uint32 = 42
            _ = v59
            _ = v60
        } else {
            if (1000 < -95) {
                v61 := GenS7{F0: 9223372036854775807}
                _ = v61
            }
            for fi62 := 0; fi62 < 10; fi62++ {
                var v63 int8 = -41
                switch "a" {
                case "hello":
                    for fi64 := 0; fi64 < 4; fi64++ {
                        var v65 string = "a"
                        if len(v5) > 0 {
                            v5[0] = fi62
                        }
                        _ = fi64
                        _ = v65
                    }
                default:
                    switch fi62 {
                    case 1:
                        v4 += v4
                    case 32767:
                        v66 := v1["\n\t"]
                        _ = v66
                    }
                }
                var v67 float32 = 0.1
                _ = fi62
                _ = v63
                _ = v67
            }
        }
    default:
        v68 := GenS3{F0: 1, F1: 42, F2: 951}
        _ = v68
    }
    v5 = append(v5, 0)
    fmt.Println(v4, v4)
    switch "" {
    case "":
        var v69 uint = 255
        _ = v69
    case "a":
        if (!true) {
            for fi70 := 0; fi70 < 1; fi70++ {
                v71 := GenS8{F0: 255, F1: -820}
                _ = fi70
                _ = v71
            }
            var v72 uint = 0
            _ = v72
        }
    case "\n\t":
        if ((184 == genFn0(true, (v4 - (^(v4))), v4)) || true) {
            switch 9223372036854775807 {
            case 32768:
                if ((((0.1 != 0.0) || false) || false) || true) {
                    v2 = 65535
                    var v73 uint = 0
                    _ = v73
                }
            default:
                v4 = 1
            }
            v3 -= -1
            v3 = (v3 % (((v3 & v3)) | 1))
        } else {
            if false {
                switch "a" {
                case "a":
                    for fi74 := 0; fi74 < 3; fi74++ {
                        switch (^(fi74)) {
                        case 0:
                            var v75 int8 = (v4 % ((v4) | 1))
                            _ = v75
                        default:
                            switch fi74 {
                            case -662:
                                v2 = (v2 ^ (v2 & (v2 % ((646) | 1))))
                            }
                        }
                        fmt.Println(v3, fi74, v3)
                        _ = fi74
                    }
                }
                var v76 bool = false
                v4 = (v4 << ((18446744073709551615) & 63))
                _ = v76
            } else {
                var v77 uint = 0
                var v78 int = 1
                _ = v77
                _ = v78
            }
        }
    }
    fmt.Println(v3, v3)
    var v79 float64 = 1.5
    v80 := GenS9{F0: 9223372036854775807}
    v79 += v79
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v79
    _ = v80
}
