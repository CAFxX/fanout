package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS5 struct { F0 int; F1 int; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0(p0 int32) int64 {
    p0 += p0
    fmt.Println(p0)
    if false {
        var v5 uint64 = 442
        _ = v5
    }
    _ = p0
    return 1
}

func main() {
    var v1 int64 = -542
    v2 := make(map[int]string)
    v3 := []int{-9223372036854775808, -9223372036854775808}
    var v4 int16 = 1
    for fi6 := 0; fi6 < 14; fi6++ {
        delete(v2, (fi6 % ((fi6) | 1)))
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi6, fi7, fi6)
            var v8 float64 = (0.0 + (0.0 + 1.5))
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: 42, F1: 106, F2: "hello"}
    for fi10 := 0; fi10 < 16; fi10++ {
        delete(v2, (-(fi10)))
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint = 71
            if (float32((3.14159 + 1e10)) <= -2.25) {
                var v13 bool = true
                var v14 uint64 = 2
                if (248 < 206) {
                    switch "foo bar" {
                    case "":
                        v15 := GenS2{F0: (v4 * (v4 >> ((4294967295) & 63)))}
                        _ = v15
                    default:
                        v4 = v4
                    }
                } else {
                    if true {
                        v16 := GenS1{F0: v4, F1: 65535, F2: ((("foo bar" + ("\n\t" + "")) + (("foo bar" + "foo bar") + "")) + "")}
                        if len(v3) > 0 {
                            v3[0] = fi10
                        }
                        v12 -= v12
                        _ = v16
                    }
                }
                v17 := GenS1{F0: v4, F1: 1, F2: "\n\t"}
                var v18 int = (fi11 ^ (^(fi11)))
                _ = v13
                _ = v14
                _ = v17
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v19 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println(v1, v19, v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: -795, F1: 1, F2: ("foo bar" + "\n\t")}
        for fi22 := 0; fi22 < 6; fi22++ {
            fmt.Println(fi22, v1)
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    if true {
        if true {
            v23 := GenS3{F0: "hello", F1: 2147483647, F2: -1}
            v19 = v19
            if true {
                v4 = v4
                if true {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 int64 = 0
                        v26 := GenS4{F0: 42, F1: 128, F2: (^(v25))}
                        v19 = 0.0
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    var v27 uint = 0
                    _ = v27
                } else {
                    var v28 int32 = 2147483647
                    _ = v28
                }
            } else {
                var v29 float64 = 0.1
                for fi30 := 0; fi30 < 13; fi30++ {
                    fmt.Println(fi30)
                    v1 = (^(v1))
                    _ = fi30
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    fmt.Println(v4, v19)
                    v29 += 1e10
                    _ = fi31
                }
                _ = v29
            }
            _ = v23
        } else {
            for fi32 := 0; fi32 < 18; fi32++ {
                if (v1 >= genFn0(2147483647)) {
                    var v33 uint16 = 0
                    for fi34 := 0; fi34 < 9; fi34++ {
                        if (0 > 161) {
                            var v35 bool = (((128 >= uint8(v33)) && (!(v1 < (v1 / ((v1) | 1))))) || true)
                            v36 := GenS2{F0: (v4 & (v4 << ((65535) & 63)))}
                            _ = v35
                            _ = v36
                        } else {
                            v19 = v19
                            fmt.Println(v1, v1, v19)
                            switch (fi32 | (fi32 & (fi32 - (fi34 - (^(fi34)))))) {
                            case 32768:
                                v37 := GenS4{F0: 205, F1: 42, F2: (v1 - (v1 >> ((223) & 63)))}
                                _ = v37
                            case 276:
                                fmt.Println(fi34)
                            case -54:
                                var v38 float64 = 0.1
                                _ = v38
                            default:
                                var v39 float32 = (3.14159 / float32(((float64(3.14159) - 0.0) / ((float64(0.1) + 1.5) - v19))))
                                _ = v39
                            }
                        }
                        _ = fi34
                    }
                    fmt.Println(v4)
                    _ = v33
                }
                _ = fi32
            }
            switch -491 {
            case -9223372036854775808:
                if false {
                    v19 = v19
                    if true {
                        for fi40 := 0; fi40 < 20; fi40++ {
                            v41 := GenS1{F0: (v4 * (v4 % ((1) | 1))), F1: 255, F2: (("foo bar" + ("a" + ("" + (("a" + "hello") + "\n\t")))) + "a")}
                            v1 -= (^(v1))
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            fmt.Println(v19, v4, v1)
                            switch fi42 {
                            case 127:
                                for fi43 := 0; fi43 < 4; fi43++ {
                                    var v44 int64 = (v1 - v1)
                                    v1 -= (v44 & v1)
                                    _ = fi43
                                    _ = v44
                                }
                            }
                            var v45 float64 = (float64((((-2.25 + float32(0.1)) - 0.0) / 1.5)) / -2.25)
                            _ = fi42
                            _ = v45
                        }
                        fmt.Println(v1)
                        v19 -= float64(0.0)
                    }
                    v2[32768] = "hello"
                } else {
                    v4 = (v4 * v4)
                    v4 = 1000
                }
            case 1:
                v2[439] = ""
            default:
                for fi46 := 0; fi46 < 13; fi46++ {
                    if false {
                        v47 := GenS5{F0: fi46, F1: (^(fi46)), F2: -1}
                        var v48 int = (^(fi46))
                        for fi49 := 0; fi49 < 12; fi49++ {
                            if true {
                                _ = len(v3)
                                delete(v2, 32768)
                            } else {
                                var v50 float32 = 1e10
                                _ = len(v2)
                                _ = v50
                            }
                            v19 -= -2.25
                            if true {
                                var v51 int8 = -23
                                var v52 uint8 = 0
                                v51 -= (v51 & v51)
                                _ = v51
                                _ = v52
                            }
                            _ = fi49
                        }
                        _ = v47
                        _ = v48
                    } else {
                        v4 = (-(v4))
                        for fi53 := 0; fi53 < 20; fi53++ {
                            v54 := GenS6{F0: 4294967295, F1: 0}
                            switch int(v1) {
                            case 32767:
                                v55 := GenS7{F0: -1, F1: (("hello" + "a") + ((("hello" + "hello") + "foo bar") + "foo bar"))}
                                _ = v55
                            case 127:
                                switch fi53 {
                                case 32767:
                                    v1 += (v1 / ((v1) | 1))
                                default:
                                    var v56 float32 = 0.1
                                    _ = v56
                                }
                                v57 := GenS1{F0: (v4 >> ((42) & 63)), F1: 241, F2: "foo bar"}
                                var v58 int64 = 42
                                _ = v57
                                _ = v58
                            case -2147483648:
                                switch ("foo bar" + "\n\t") {
                                case "hello":
                                    v19 = (float64(3.14159) / (v19 * v19))
                                case "\n\t":
                                    if false {
                                        switch "" {
                                        case "a":
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                var v59 int16 = (v4 << ((1) & 63))
                                                v19 += v19
                                                _ = v59
                                            }
                                        case "foo bar":
                                            if true {
                                                if false {
                                                    v60 := v2[fi46]
                                                    v4 = (v4 ^ v4)
                                                    _ = v60
                                                } else {
                                                    var v61 uint16 = 0
                                                    if true {
                                                        if (1 == 1) {
                                                            v1 += -9223372036854775808
                                                            switch (fi53 % (((fi53 | (fi53 - (fi46 << ((726) & 63))))) | 1)) {
                                                            case 128:
                                                                if true {
                                                                    v62 := GenS8{F0: 2147483647, F1: 701, F2: 4294967295}
                                                                    var v63 int16 = 32767
                                                                    var v64 uint = 18446744073709551615
                                                                    _ = v62
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                var v65 int64 = v1
                                                                switch "a" {
                                                                case "hello":
                                                                    v19 = (float64(-2.25) + v19)
                                                                case "":
                                                                    v61 = (v61 >> ((132) & 63))
                                                                case "\n\t":
                                                                    v66 := GenS6{F0: 2147483647, F1: 18446744073709551615}
                                                                    _ = v66
                                                                default:
                                                                    v61 = v61
                                                                }
                                                                _ = v65
                                                            case 2147483647:
                                                                v19 = (v19 + v19)
                                                            default:
                                                                v19 = (((((v19 * (1.5 / v19)) / (v19 / (1.5 + v19))) - float64((1e10 + ((0.0 / -2.25) + 1.5)))) - 1e10) * ((((v19 - (v19 - v19)) / (v19 + v19)) + ((v19 * v19) + float64(float32(1.5)))) + ((0.1 - (1.5 + 0.1)) * (v19 / 1.5))))
                                                            }
                                                            switch fi46 {
                                                            case 206:
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi46
                                                                }
                                                            default:
                                                                if false {
                                                                    _ = len(v2)
                                                                    var v67 uint32 = 4294967295
                                                                    _ = v67
                                                                } else {
                                                                    v68 := GenS7{F0: -6, F1: "a"}
                                                                    v69 := GenS7{F0: 0, F1: "foo bar"}
                                                                    v1 = (v1 * (v1 << ((18446744073709551615) & 63)))
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                            }
                                                        } else {
                                                            if ((139 != 255) && false) {
                                                                v1 = genFn0(-2147483648)
                                                                v4 -= -1
                                                                delete(v2, 665)
                                                            }
                                                        }
                                                        if ((v4 == 376) && true) {
                                                            v70 := GenS9{F0: 2, F1: 4294967295}
                                                            v3 = append(v3, fi53)
                                                            _ = v70
                                                        } else {
                                                            switch (fi46 * 394) {
                                                            case 4294967295:
                                                                v71 := v2[(fi53 / ((fi53) | 1))]
                                                                _ = v71
                                                            case 2147483647:
                                                                v19 -= ((((1.5 * 0.0) + (float64((0.1 * -2.25)) * (((v19 - v19) - (1.5 + float64(float32((v19 - 0.1))))) - (1.5 - 0.0)))) / 1.0) + 0.1)
                                                            }
                                                            v4 -= (v4 & v4)
                                                        }
                                                        v2[(fi46 << ((633) & 63))] = ("hello" + "\n\t")
                                                        fmt.Println(fi53)
                                                    }
                                                    var v72 uint16 = v61
                                                    _ = v61
                                                    _ = v72
                                                }
                                                var v73 uint64 = 2
                                                v19 -= v19
                                                _ = v73
                                            } else {
                                                v74 := GenS6{F0: 429, F1: 0}
                                                v75 := GenS9{F0: 0, F1: 14}
                                                var v76 int16 = (v4 << ((1) & 63))
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                        }
                                    }
                                default:
                                    switch fi53 {
                                    case 4294967295:
                                        switch "a" {
                                        case "a":
                                            var v77 string = "hello"
                                            _ = v77
                                        }
                                    }
                                    var v78 int32 = -42
                                    switch ("" + ("a" + "foo bar")) {
                                    case "\n\t":
                                        fmt.Println(fi53, v1)
                                    }
                                    fmt.Println(v4, v1, fi46)
                                    _ = v78
                                }
                            }
                            _ = fi53
                            _ = v54
                        }
                    }
                    var v79 float32 = ((0.0 + 0.1) - -2.25)
                    _ = fi46
                    _ = v79
                }
            }
        }
        v80 := GenS5{F0: -497, F1: 32768, F2: 0}
        v81 := GenS4{F0: 2, F1: 92, F2: v1}
        _ = v80
        _ = v81
    } else {
        var v82 uint64 = 18446744073709551615
        fmt.Println(v19)
        _ = v82
    }
    var v83 int16 = 1
    var v84 uint32 = 257
    for fi85 := 0; fi85 < 10; fi85++ {
        v83 += (v4 >> ((42) & 63))
        var v86 int64 = 1
        if (!(false && (fi85 == 4294967295))) {
            for fi87 := 0; fi87 < 1; fi87++ {
                v88 := v2[(^(fi85))]
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) > 0 {
                    v3[0] = (fi87 / ((-660) | 1))
                }
                _ = fi87
                _ = v88
            }
            v89 := GenS7{F0: 90, F1: "foo bar"}
            _ = v89
        } else {
            fmt.Println(fi85, v86)
            var v90 uint32 = v84
            v91 := GenS10{F0: 255, F1: 1, F2: 42}
            _ = v90
            _ = v91
        }
        _ = fi85
        _ = v86
    }
    v83 -= v4
    v92 := v2[128]
    var v93 int8 = -23
    v19 -= float64(0.0)
    v84 += v84
    v93 = 127
    v3 = append(v3, 128)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v9
    _ = v92
    _ = v93
}
