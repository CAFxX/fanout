package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0() int64 {
    var v5 int32 = -168
    for fi6 := 0; fi6 < 8; fi6++ {
        if (((2147483647 != -1) || true) || true) {
            var v7 uint64 = 0
            for fi8 := 0; fi8 < 9; fi8++ {
                v5 = v5
                _ = fi8
            }
            _ = v7
        }
        _ = fi6
    }
    if (("hello" + ("foo bar" + "")) >= "foo bar") {
        fmt.Println(v5)
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                v10 := GenS1{F0: -1, F1: -24, F2: 4294967295}
                v11 := GenS1{F0: 9223372036854775807, F1: 44, F2: 42}
                var v12 int64 = 9223372036854775807
                _ = v10
                _ = v11
                _ = v12
            }
            var v13 bool = (false && (361 == 221))
            _ = fi9
            _ = v13
        }
    }
    v5 = v5
    _ = v5
    return 42
}
func genFn1(p0 int64) int32 {
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := []int{}
        var v16 uint64 = 2
        v17 := GenS1{F0: -9223372036854775808, F1: -31, F2: v16}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := []int{}
    v19 := GenS1{F0: (p0 - p0), F1: -60, F2: 2}
    _ = p0
    _ = v18
    _ = v19
    return -1
}

func main() {
    var v1 int = -901
    var v2 float32 = 0.0
    var v3 int = (v1 / ((v1) | 1))
    var v4 uint = 65535
    v2 = v2
    v4 -= v4
    var v20 int16 = 42
    switch v1 {
    case 127:
        var v21 int32 = 2147483647
        _ = v21
    case 32767:
        fmt.Println(v20, v3, v20)
    case 688:
        v3 += 42
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v4 += v4
        _ = fi22
    }
    v23 := GenS1{F0: -1, F1: -128, F2: 4294967295}
    v4 = v4
    v4 = v4
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS1{F0: 0, F1: -128, F2: 2}
        if false {
            if (63 >= 9223372036854775807) {
                fmt.Println(v3, v4, v2)
            } else {
                v26 := GenS1{F0: 42, F1: 42, F2: 4294967295}
                if (867 > 32767) {
                    var v27 uint32 = 42
                    v1 -= v1
                    var v28 int32 = -129
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint16 = 255
                    v30 := GenS2{F0: 0}
                    _ = v29
                    _ = v30
                }
                switch 494 {
                case -9223372036854775808:
                    switch ("hello" + ("" + "\n\t")) {
                    case "hello":
                        if (693 > 2) {
                            switch "foo bar" {
                            case "a":
                                if false {
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        v32 := []int{(fi24 * -2147483648), (fi24 / (((-(fi31))) | 1))}
                                        var v33 uint8 = 255
                                        var v34 uint = (^(v4))
                                        if ((0.1 == ((-2.25 - 1e10) / 3.14159)) && ("foo bar" > "a")) {
                                            _ = len(v32)
                                            v20 -= v20
                                            v2 -= v2
                                        }
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                } else {
                                    v4 += 5
                                }
                            case "\n\t":
                                switch "hello" {
                                case "foo bar":
                                    switch "hello" {
                                    case "hello":
                                        for fi35 := 0; fi35 < 8; fi35++ {
                                            if (!true) {
                                                v36 := GenS2{F0: (v4 % (((v4 & v4)) | 1))}
                                                v37 := GenS3{F0: "hello", F1: 255}
                                                v38 := []int{fi35, 128, fi35}
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                v39 := []int{}
                                                v4 = 65535
                                                if len(v39) > 0 {
                                                    v39[0] = (-(v3))
                                                }
                                                var v40 uint = 910
                                                v41 := GenS3{F0: "a", F1: 369}
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            }
                                            var v42 uint = v4
                                            v1 = -194
                                            _ = fi35
                                            _ = v42
                                        }
                                    case "foo bar":
                                        if (true || true) {
                                            var v43 uint = v4
                                            v44 := []int{}
                                            var v45 int32 = genFn1(-753)
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        } else {
                                            v2 += 0.1
                                            fmt.Println(v2, v3)
                                            v46 := []int{-984, fi24, v1}
                                            _ = v46
                                        }
                                    }
                                case "a":
                                    v47 := []int{2147483647}
                                    _ = v47
                                default:
                                    v2 = v2
                                }
                            case "foo bar":
                                v1 = v3
                            default:
                                v48 := GenS3{F0: "hello", F1: 255}
                                _ = v48
                            }
                        }
                    case "foo bar":
                        if (12 != 69) {
                            if true {
                                switch "hello" {
                                case "foo bar":
                                    v49 := []int{(v1 - 42)}
                                    _ = v49
                                case "hello":
                                    var v50 int = (v3 << (((^(v4))) & 63))
                                    _ = v50
                                default:
                                    if ((v3 | (fi24 % (((-(fi24))) | 1))) <= fi24) {
                                        if true {
                                            v2 = float32(1e10)
                                            var v51 uint16 = 65535
                                            _ = v51
                                        }
                                        if (584 <= 489) {
                                            var v52 uint16 = 1
                                            _ = v52
                                        } else {
                                            if (true && true) {
                                                v2 -= v2
                                                if true {
                                                    fmt.Println(v3)
                                                    v53 := []int{(v1 & v3)}
                                                    v54 := GenS1{F0: 194, F1: -128, F2: 0}
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    v55 := []int{}
                                                    v55 = append(v55, fi24)
                                                    _ = v55
                                                }
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v56 bool = false
                                                    _ = v56
                                                case "a":
                                                    if (false && false) {
                                                        var v57 uint32 = 65535
                                                        _ = v57
                                                    }
                                                }
                                                var v58 uint64 = 1
                                                _ = v58
                                            }
                                            for fi59 := 0; fi59 < 13; fi59++ {
                                                var v60 uint = v4
                                                _ = fi59
                                                _ = v60
                                            }
                                            v61 := GenS2{F0: (^(v4))}
                                            _ = v61
                                        }
                                    } else {
                                        switch "a" {
                                        case "a":
                                            var v62 int = (v3 / ((int(v4)) | 1))
                                            _ = v62
                                        case "hello":
                                            var v63 int64 = 1
                                            _ = v63
                                        }
                                    }
                                }
                                v64 := []int{-854, (v1 + fi24), fi24}
                                v65 := GenS4{F0: 42, F1: ((1 <= 2147483647) && false)}
                                _ = v64
                                _ = v65
                            }
                            v66 := GenS3{F0: ((("hello" + ("foo bar" + ("" + (("hello" + "") + "a")))) + "a") + "\n\t"), F1: 65535}
                            _ = v66
                        }
                    case "a":
                        v67 := []int{fi24}
                        _ = v67
                    }
                default:
                    v2 = 3.14159
                }
                _ = v26
            }
            for fi68 := 0; fi68 < 16; fi68++ {
                fmt.Println(v20, v20, v2)
                _ = fi68
            }
            v69 := GenS5{F0: 4294967295}
            _ = v69
        } else {
            var v70 int = -9223372036854775808
            _ = v70
        }
        _ = fi24
        _ = v25
    }
    var v71 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v71
}
