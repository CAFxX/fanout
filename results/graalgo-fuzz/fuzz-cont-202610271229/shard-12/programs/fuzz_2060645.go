package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS5 struct { F0 int32; F1 string; F2 string }
type GenS6 struct { F0 int32; F1 uint64; F2 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0(p0 uint, p1 int64) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(fi4, p1)
        _ = fi4
    }
    switch -1 {
    case 127:
        v5 := GenS1{F0: -32768, F1: 42}
        _ = v5
    default:
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{fi6, (fi6 + fi6), fi6}
            var v8 int64 = (p1 * p1)
            v9 := GenS1{F0: -795, F1: 4294967295}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int64 = -1
    var v2 int64 = (v1 + (^(v1)))
    v3 := make(map[int]string)
    if (4294967295 <= 2) {
        v1 -= (v1 & (v2 / ((9223372036854775807) | 1)))
        v1 = (v1 & (-(v1)))
        var v10 int8 = 42
        _ = v10
    } else {
        if (true && false) {
            if ((v1 & v2) > 9223372036854775807) {
                v3[4294967295] = "\n\t"
            }
            switch 9223372036854775807 {
            case -228:
                v2 = (^(v2))
            case 32768:
                v11 := GenS1{F0: 406, F1: 2}
                _ = v11
            case 9223372036854775807:
                fmt.Println(v2)
            default:
                v12 := []int{int(v2), -334, 4294967295}
                _ = v12
            }
        }
        switch "hello" {
        case "\n\t":
            switch "" {
            case "hello":
                switch "" {
                case "\n\t":
                    var v13 string = "foo bar"
                    _ = v13
                case "hello":
                    v14 := GenS2{F0: 65535, F1: "", F2: 404}
                    _ = v14
                default:
                    v15 := GenS3{F0: 966, F1: 188}
                    _ = v15
                }
            case "\n\t":
                fmt.Println(v1, v1)
            case "a":
                v1 = -1
            }
        case "hello":
            if (2 > 0) {
                switch ("" + "hello") {
                case "a":
                    v3[428] = "a"
                case "":
                    v16 := v3[2147483647]
                    _ = v16
                default:
                    var v17 float32 = 1.5
                    _ = v17
                }
            } else {
                v3[32768] = "hello"
                fmt.Println(v1, v1)
                v1 = v1
            }
        }
        fmt.Println(v1, v1)
    }
    var v18 bool = true
    switch ((("hello" + "foo bar") + ("\n\t" + "a")) + "hello") {
    case "":
        v2 = v2
    case "hello":
        var v19 int = 4294967295
        _ = v19
    }
    fmt.Println(v18, v1, v1)
    v1 -= v2
    var v20 uint32 = 0
    var v21 int64 = 42
    if v18 {
        v22 := GenS4{F0: 210, F1: 0, F2: 42}
        v23 := []int{128, -544, 354, 32768}
        _ = v22
        _ = v23
    } else {
        v18 = v18
    }
    v21 -= v2
    switch "\n\t" {
    case "foo bar":
        v18 = true
    case "":
        if v18 {
            v3[32767] = "\n\t"
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS2{F0: 1, F1: ("" + "\n\t"), F2: 42}
                if ((v20 * v20) == (v20 / ((956) | 1))) {
                    v26 := GenS5{F0: 42, F1: "\n\t", F2: ("\n\t" + "")}
                    for fi27 := 0; fi27 < 1; fi27++ {
                        fmt.Println(fi27, v1, fi24)
                        if v18 {
                            if (593 < -1) {
                                var v28 uint32 = v20
                                _ = v28
                            } else {
                                if (!(!(v2 >= (v1 | v1)))) {
                                    if true {
                                        fmt.Println(fi24, fi27, v2)
                                        if (((("foo bar" + "\n\t") + "a") + "a") <= "") {
                                            var v29 uint8 = 1
                                            _ = v29
                                        } else {
                                            v18 = true
                                            var v30 uint16 = 480
                                            _ = v30
                                        }
                                        v18 = (604 > 0)
                                    } else {
                                        var v31 uint32 = (^(v20))
                                        _ = v31
                                    }
                                    var v32 int8 = -128
                                    var v33 int32 = 1
                                    _ = v32
                                    _ = v33
                                }
                            }
                        } else {
                            v34 := GenS3{F0: 107, F1: uint64(v20)}
                            if false {
                                v2 = v2
                            } else {
                                var v35 uint32 = (^(v20))
                                _ = v35
                            }
                            var v36 int8 = 0
                            var v37 uint32 = (v20 + (^(v20)))
                            v20 -= (v37 | v20)
                            v38 := []int{fi24, (^(fi27))}
                            _ = v34
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        _ = fi27
                    }
                    _ = v26
                }
                v39 := GenS6{F0: 830, F1: 790, F2: fi24}
                _ = fi24
                _ = v25
                _ = v39
            }
        } else {
            switch "a" {
            case "":
                if (false && v18) {
                    v40 := []int{127}
                    _ = v40
                } else {
                    fmt.Println(v20)
                    v41 := GenS4{F0: 615, F1: 0, F2: 32767}
                    _ = v41
                }
            case "\n\t":
                for fi42 := 0; fi42 < 17; fi42++ {
                    for fi43 := 0; fi43 < 9; fi43++ {
                        if (!(-95 > -162)) {
                            v44 := GenS7{F0: v18}
                            var v45 float64 = 0.1
                            _ = v44
                            _ = v45
                        } else {
                            if (false || (-2.25 < 1.5)) {
                                var v46 uint64 = 2
                                v46 = 354
                                v2 += v1
                                _ = v46
                            } else {
                                if v18 {
                                    if ("\n\t" >= "hello") {
                                        v47 := []int{fi42}
                                        switch fi42 {
                                        case -2147483648:
                                            v3[(fi42 - (fi42 - (^(fi43))))] = ("" + "a")
                                        case 0:
                                            v48 := GenS2{F0: 742, F1: ("hello" + "foo bar"), F2: 977}
                                            _ = v48
                                        default:
                                            v18 = false
                                        }
                                        _ = v47
                                    }
                                } else {
                                    fmt.Println(v21)
                                }
                                var v49 uint16 = 1
                                switch "a" {
                                case "hello":
                                    if ((^(v2)) == 9223372036854775807) {
                                        if v18 {
                                            v50 := []int{(fi43 + fi42)}
                                            var v51 float32 = (0.0 * (0.1 + 1e10))
                                            if v18 {
                                                v51 = (((v51 / (float32(((0.1 / (-2.25 / float64(1e10))) / ((0.1 * -2.25) / (1.5 + (3.14159 * 0.0))))) + float32((((float64(1e10) / 1.0) - 1.5) - (0.1 - 0.1))))) - ((float32(1e10) * v51) + (v51 - v51))) - float32((float64(v51) / 1e10)))
                                                v18 = false
                                                if v18 {
                                                    v20 = v20
                                                    var v52 int = (fi42 - (^(fi42)))
                                                    v21 += 1
                                                    v51 = 0.1
                                                    v53 := GenS5{F0: -910, F1: (("a" + "foo bar") + "foo bar"), F2: ("a" + (("hello" + "hello") + ""))}
                                                    var v54 int16 = 32767
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            } else {
                                                v1 += v21
                                                if true {
                                                    v20 -= uint32(v49)
                                                }
                                                _ = len(v3)
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                        if true {
                                            var v55 float32 = 0.0
                                            _ = v55
                                        }
                                    }
                                }
                                _ = v49
                            }
                        }
                        _ = fi43
                    }
                    v56 := []int{fi42, fi42, (fi42 % (((fi42 / (((^(fi42))) | 1))) | 1)), (^(fi42))}
                    _ = fi42
                    _ = v56
                }
            }
        }
    case "hello":
        v18 = false
    }
    v57 := GenS6{F0: -989, F1: 311, F2: 42}
    for fi58 := 0; fi58 < 12; fi58++ {
        v59 := []int{fi58, (^(fi58)), fi58, (fi58 << ((genFn0(uint(v20), (v1 - -9223372036854775808))) & 63))}
        for fi60 := 0; fi60 < 17; fi60++ {
            if false {
                var v61 uint32 = 4294967295
                _ = v61
            }
            fmt.Println(v21, v20, fi58)
            delete(v3, fi58)
            _ = fi60
        }
        if (1 > 42) {
            var v62 uint32 = 204
            _ = v62
        } else {
            var v63 float32 = 3.14159
            v64 := GenS2{F0: 353, F1: ("" + "foo bar"), F2: 4294967295}
            _ = v63
            _ = v64
        }
        _ = fi58
        _ = v59
    }
    v65 := GenS8{F0: 42, F1: ("hello" + ""), F2: 1}
    v20 -= (v20 >> ((764) & 63))
    v2 = (v2 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v57
    _ = v65
}
