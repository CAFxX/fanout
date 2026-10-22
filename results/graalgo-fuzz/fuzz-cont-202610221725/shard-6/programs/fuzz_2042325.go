package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint; F1 int32; F2 uint }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0(p0 bool, p1 int16) bool {
    v4 := GenS1{F0: 0, F1: 2}
    v5 := GenS1{F0: (p1 | (-(p1))), F1: 255}
    var v6 float64 = 0.1
    switch "\n\t" {
    case "hello":
        if p0 {
            var v7 int16 = p1
            var v8 bool = true
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(p1, fi9)
                p0 = v8
                v10 := GenS1{F0: 539, F1: 673}
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        } else {
            var v11 string = "\n\t"
            p1 = 1000
            var v12 uint = 18446744073709551615
            _ = v11
            _ = v12
        }
    }
    if false {
        var v13 uint64 = 42
        v14 := GenS1{F0: (-(p1)), F1: 620}
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS1{F0: 0, F1: 1}
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            switch 32768 {
            case 128:
                if p0 {
                    p1 = int16(fi17)
                    v18 := []int{(-(fi17))}
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 uint = 42
                    _ = fi19
                    _ = v20
                }
            }
            v21 := GenS2{F0: 716, F1: 0}
            v22 := []int{fi17, -1}
            _ = fi17
            _ = v21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int32, p1 uint8) int8 {
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            p1 += (p1 + (p1 * (p1 << ((4294967295) & 63))))
            if (2147483647 != p0) {
                switch ("hello" + (("foo bar" + "hello") + "a")) {
                case "hello":
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS3{F0: 4294967295, F1: false}
                        v26 := []int{fi23, fi24, -2147483648}
                        if ((!(true && (0 != 190))) || true) {
                            switch "hello" {
                            case "foo bar":
                                p1 = (p1 / ((126) | 1))
                            }
                            v27 := GenS2{F0: 825, F1: 1}
                            _ = v27
                        }
                        v28 := GenS1{F0: 0, F1: 255}
                        var v29 uint16 = 65535
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v28
                        _ = v29
                    }
                case "foo bar":
                    fmt.Println(p1, p0, p1)
                default:
                    fmt.Println(fi23, p0)
                }
                fmt.Println(p1, p0)
                switch "hello" {
                case "a":
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 bool = ((true && ((0.1 > 1.5) && true)) || false)
                        fmt.Println(v31)
                        v32 := []int{(fi23 & fi23)}
                        switch (fi30 * fi23) {
                        case 189:
                            v33 := GenS4{F0: 42, F1: p1}
                            _ = v33
                        case -901:
                            var v34 bool = v31
                            _ = v34
                        case -1:
                            fmt.Println(p0, fi23)
                        default:
                            v32 = append(v32, (fi23 / ((4294967295) | 1)))
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case "foo bar":
                    if false {
                        var v35 float32 = -2.25
                        var v36 float64 = 3.14159
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int16 = 32767
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v39 := []int{fi38}
                            switch (fi23 ^ fi23) {
                            case 0:
                                if false {
                                    p1 = p1
                                }
                            }
                            _ = fi38
                            _ = v39
                        }
                        if (!true) {
                            for fi40 := 0; fi40 < 4; fi40++ {
                                var v41 int32 = -2147483648
                                var v42 uint8 = p1
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                            switch "hello" {
                            case "foo bar":
                                var v43 bool = (1 != 32767)
                                _ = v43
                            case "hello":
                                v44 := GenS5{F0: 0, F1: -2147483648, F2: 2}
                                _ = v44
                            case "a":
                                fmt.Println(p0, p1, p1)
                            default:
                                v37 = (^(v37))
                            }
                            var v45 int64 = -348
                            _ = v45
                        } else {
                            v37 += (v37 % ((0) | 1))
                            v37 -= 42
                            var v46 uint32 = 42
                            _ = v46
                        }
                        _ = v37
                    }
                }
            } else {
                switch fi23 {
                case 32768:
                    v47 := []int{fi23, fi23, fi23, 485}
                    _ = v47
                case 464:
                    v48 := GenS3{F0: 42, F1: ((p1 & p1) == (p1 - (p1 * (^(p1)))))}
                    _ = v48
                case 42:
                    switch (fi23 * (-(fi23))) {
                    case -9223372036854775808:
                        switch ("foo bar" + ("foo bar" + "")) {
                        case "a":
                            var v49 int = fi23
                            _ = v49
                        case "\n\t":
                            fmt.Println(fi23)
                        case "hello":
                            var v50 int64 = 1
                            _ = v50
                        default:
                            v51 := []int{(-(fi23))}
                            _ = v51
                        }
                    case -178:
                        var v52 uint8 = 117
                        _ = v52
                    case 32767:
                        if (4294967295 > 317) {
                            if false {
                                for fi53 := 0; fi53 < 5; fi53++ {
                                    switch (fi23 ^ fi53) {
                                    case 128:
                                        var v54 uint8 = p1
                                        v55 := []int{(fi53 & (fi23 - fi23)), (fi23 - fi23)}
                                        var v56 int64 = 526
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    default:
                                        v57 := []int{fi53, 4294967295}
                                        _ = v57
                                    }
                                    _ = fi53
                                }
                                p0 += (p0 - (p0 + (p0 | (p0 & 42))))
                                fmt.Println(p0)
                            }
                            v58 := []int{fi23, fi23}
                            _ = v58
                        }
                    }
                default:
                    p1 += p1
                }
            }
            p0 -= p0
            _ = fi23
        }
        p1 = p1
    } else {
        fmt.Println(p0, p0)
    }
    if true {
        p1 += (p1 << ((uint(p1)) & 63))
    } else {
        v59 := []int{-408, 128, 2147483647, 32767}
        v60 := GenS1{F0: 1, F1: 1}
        v61 := GenS6{F0: 1000, F1: "foo bar"}
        _ = v59
        _ = v60
        _ = v61
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 int16 = 194
        v64 := []int{}
        _ = fi62
        _ = v63
        _ = v64
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 798
    v3 := make(map[string]uint32)
    v2 = 255
    v65 := []int{0, 4294967295}
    _ = len(v65)
    v2 = (v2 - 255)
    var v66 int8 = 0
    v2 = 113
    v2 = (v2 - (^(v2)))
    _ = len(v1)
    switch "a" {
    case "foo bar":
        switch 937 {
        case 601:
            var v67 int8 = (v66 >> ((v2) & 63))
            _ = v67
        case -9223372036854775808:
            if (4294967295 != 2147483647) {
                v1[4294967295] = 1
                if (!((-2.25 + 3.14159) < 1e10)) {
                    var v68 float64 = 3.14159
                    _ = v68
                }
            } else {
                if true {
                    v69 := GenS7{F0: -1}
                    v70 := GenS6{F0: -32768, F1: ("hello" + "hello")}
                    _ = v69
                    _ = v70
                } else {
                    switch ("" + ("\n\t" + "hello")) {
                    case "hello":
                        switch 4294967295 {
                        case 32768:
                            if false {
                                var v71 uint32 = 65535
                                if (848 == 2147483647) {
                                    v71 += 874
                                    fmt.Println(v66)
                                }
                                fmt.Println(v2, v66)
                                _ = v71
                            } else {
                                switch 32767 {
                                case 9223372036854775807:
                                    if true {
                                        v72 := GenS4{F0: v66, F1: 70}
                                        v2 = (^(v2))
                                        fmt.Println(v2, v2)
                                        _ = v72
                                    }
                                case 614:
                                    v73 := v3["hello"]
                                    _ = v73
                                }
                            }
                        default:
                            for fi74 := 0; fi74 < 14; fi74++ {
                                var v75 int8 = 1
                                v2 += v2
                                if false {
                                    v75 = (v66 | v75)
                                    var v76 int64 = 1
                                    v75 -= (^(v75))
                                    _ = v76
                                }
                                _ = fi74
                                _ = v75
                            }
                        }
                    case "foo bar":
                        v2 += v2
                    default:
                        for fi77 := 0; fi77 < 2; fi77++ {
                            for fi78 := 0; fi78 < 14; fi78++ {
                                var v79 int = fi77
                                var v80 int = fi77
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                            if len(v65) > 0 {
                                v65[0] = fi77
                            }
                            _ = fi77
                        }
                    }
                }
                switch "" {
                case "":
                    v81 := GenS5{F0: (v2 % ((v2) | 1)), F1: 359, F2: v2}
                    _ = v81
                case "a":
                    switch ("hello" + "a") {
                    case "hello":
                        var v82 bool = (128 >= 255)
                        _ = v82
                    }
                default:
                    var v83 bool = (32767 >= 42)
                    _ = v83
                }
                v84 := GenS8{F0: (float64(float32(((1e10 * 1e10) * -2.25))) > 1e10), F1: 127, F2: -9223372036854775808}
                _ = v84
            }
        }
    case "\n\t":
        var v85 uint8 = 217
        _ = v85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
}
