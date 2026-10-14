package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 bool; F1 uint; F2 bool }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint8; F1 string; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint; F2 int16 }
func genFn0(p0 int8, p1 uint32, p2 int16) uint8 {
    var v6 int = -663
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            if false {
                v8 := GenS1{F0: 656, F1: 65535}
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := GenS2{F0: "foo bar", F1: (("hello" + "") + "\n\t"), F2: (p0 - (p0 / (((p0 << ((194) & 63))) | 1)))}
                    if (true && true) {
                        v11 := GenS2{F0: "a", F1: "\n\t", F2: (p0 ^ (p0 | p0))}
                        v12 := GenS3{F0: 42, F1: 1, F2: fi9}
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS2{F0: ("" + ""), F1: "\n\t", F2: p0}
                    _ = fi9
                    _ = v10
                    _ = v13
                }
                var v14 float32 = float32(-2.25)
                _ = v14
                _ = v8
            }
            var v15 bool = false
            v16 := []int{int(p0), fi7, (fi7 % ((-1) | 1))}
            _ = fi7
            _ = v15
            _ = v16
        }
        v17 := []int{(^(v6))}
        _ = v17
    }
    p2 = int16(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := []string{"hello", ((("" + "foo bar") + "hello") + (("" + "a") + "")), "foo bar"}
    var v2 float32 = float32(0.0)
    var v3 uint32 = 0
    var v4 int64 = 182
    v5 := make(map[uint32]bool)
    v4 += -9223372036854775808
    fmt.Println(v4, v4, v2)
    v18 := GenS4{F0: 1, F1: 1}
    if (v4 >= v4) {
        fmt.Println(v3)
        switch 4294967295 {
        case 32767:
            v3 += v3
        case 32768:
            var v19 uint8 = 42
            _ = v19
        case 2147483647:
            v20 := GenS4{F0: 42, F1: 882}
            _ = v20
        default:
            switch 32768 {
            case -9223372036854775808:
                var v21 uint8 = 128
                _ = v21
            case -824:
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint32 = 2147483648
                    v24 := GenS3{F0: v4, F1: -1, F2: -695}
                    var v25 int64 = v4
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                v26 := GenS5{F0: (0 == -1), F1: uint(v3), F2: (genFn0(-67, (^(v3)), 1) > 2)}
                _ = v26
            }
        }
    } else {
        switch (("a" + "") + ("\n\t" + "")) {
        case "a":
            _ = len(v5)
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        }
        v4 += v4
        v3 += (v3 * 42)
    }
    var v27 uint32 = v3
    for fi28 := 0; fi28 < 1; fi28++ {
        if ((v27 << ((940) & 63)) < v3) {
            var v29 int = -1
            _ = v29
        }
        fmt.Println(v2, v4, v4)
        switch (fi28 & 9223372036854775807) {
        case 4294967295:
            switch (fi28 % ((128) | 1)) {
            case 42:
                v30 := GenS6{F0: ("a" + "hello"), F1: 0}
                _ = v30
            case 1:
                if false {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if (!false) {
                            v4 -= v4
                            switch "foo bar" {
                            case "hello":
                                v2 = (v2 * v2)
                            case "a":
                                var v32 int8 = -128
                                _ = v32
                            default:
                                var v33 uint = 1
                                _ = v33
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = fi31
                    }
                    if (false && true) {
                        if true {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v35 := GenS7{F0: 233, F1: ("hello" + ""), F2: 0}
                                _ = fi34
                                _ = v35
                            }
                            v36 := GenS5{F0: (4294967295 > 0), F1: 42, F2: (2 == 179)}
                            _ = v36
                        }
                    }
                } else {
                    var v37 int = fi28
                    for fi38 := 0; fi38 < 19; fi38++ {
                        fmt.Println(v27)
                        switch "\n\t" {
                        case "hello":
                            v27 -= v3
                        case "a":
                            v4 = 9223372036854775807
                        case "\n\t":
                            v27 = v27
                        default:
                            v37 += 127
                        }
                        if false {
                            fmt.Println(v37, fi28, fi28)
                            var v39 int32 = -2147483648
                            if true {
                                var v40 uint64 = 4294967295
                                var v41 float32 = ((1e10 * v2) / v2)
                                _ = v40
                                _ = v41
                            } else {
                                switch int(v4) {
                                case 1:
                                    var v42 uint8 = 127
                                    _ = v42
                                case 4294967295:
                                    if ((^(v27)) == (v3 ^ (v27 << ((42) & 63)))) {
                                        if (("foo bar" + "foo bar") >= ("a" + "hello")) {
                                            var v43 int16 = int16(v27)
                                            if false {
                                                v39 -= v39
                                            }
                                            _ = v43
                                        } else {
                                            v39 = v39
                                        }
                                        fmt.Println(fi28)
                                    } else {
                                        if false {
                                            if true {
                                                var v44 string = ("\n\t" + "foo bar")
                                                _ = v44
                                            }
                                            v2 = (v2 + 1e10)
                                        }
                                    }
                                }
                                v4 = v4
                                switch 772 {
                                case -530:
                                    v45 := GenS6{F0: "", F1: 0}
                                    _ = v45
                                default:
                                    fmt.Println(v39, fi38, fi38)
                                    var v46 float64 = (((3.14159 * (1.5 * 0.0)) + 0.1) + 3.14159)
                                    _ = v46
                                }
                                var v47 uint16 = 42
                                switch "" {
                                case "hello":
                                    var v48 int32 = 1
                                    _ = v48
                                case "\n\t":
                                    switch -868 {
                                    case -1:
                                        switch 32768 {
                                        case -9223372036854775808:
                                            var v49 uint16 = v47
                                            _ = v49
                                        case 9223372036854775807:
                                            v3 += v27
                                        case 127:
                                            delete(v5, v27)
                                        }
                                    case 2147483647:
                                        v50 := GenS3{F0: v4, F1: v39, F2: (v37 & 452)}
                                        _ = v50
                                    default:
                                        v51 := GenS5{F0: (true || true), F1: 2, F2: (42 < 2)}
                                        _ = v51
                                    }
                                default:
                                    v39 -= 1
                                }
                                _ = v47
                            }
                            switch fi28 {
                            case 9223372036854775807:
                                if ((fi28 % ((int(v4)) | 1)) > -36) {
                                    var v52 uint = 44
                                    v3 -= (v27 | (v3 & (^(v3))))
                                    var v53 uint32 = (v3 * 0)
                                    fmt.Println(v39)
                                    _ = v52
                                    _ = v53
                                }
                            case 128:
                                v2 += v2
                            default:
                                switch ("hello" + ("" + "")) {
                                case "foo bar":
                                    v54 := GenS8{F0: "hello"}
                                    v27 -= v27
                                    v37 -= (fi28 >> ((0) & 63))
                                    _ = v54
                                case "hello":
                                    v37 = 127
                                case "a":
                                    v55 := GenS8{F0: "\n\t"}
                                    _ = v55
                                }
                                var v56 bool = (!(-17 <= -66))
                                _ = v56
                            }
                            _ = v39
                        } else {
                            if true {
                                v57 := GenS3{F0: (v4 + -9223372036854775808), F1: -1, F2: -392}
                                var v58 int16 = 42
                                var v59 int64 = (v4 >> ((65535) & 63))
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            if true {
                                v4 -= 0
                                if (!false) {
                                    var v60 uint16 = 255
                                    v61 := GenS9{F0: "\n\t", F1: 2, F2: 1}
                                    _ = v60
                                    _ = v61
                                } else {
                                    if (!true) {
                                        if true {
                                            if (!(true || true)) {
                                                v27 -= (v27 ^ v27)
                                                fmt.Println(fi28)
                                            }
                                        } else {
                                            delete(v5, (v27 % (((v27 >> ((809) & 63))) | 1)))
                                        }
                                        v4 -= (v4 << ((279) & 63))
                                        v62 := GenS3{F0: 42, F1: -2147483648, F2: fi28}
                                        _ = v62
                                    } else {
                                        if false {
                                            v63 := GenS8{F0: ("\n\t" + "foo bar")}
                                            switch v37 {
                                            case 4294967295:
                                                switch (fi38 ^ (fi28 - v37)) {
                                                case 510:
                                                    fmt.Println(fi38, v4, v3)
                                                case -91:
                                                    v64 := GenS2{F0: ("a" + ("\n\t" + ("hello" + "\n\t"))), F1: "\n\t", F2: -21}
                                                    _ = v64
                                                default:
                                                    v37 -= (fi38 >> ((65535) & 63))
                                                }
                                            }
                                            _ = v63
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = (("\n\t" + "\n\t") + ("" + "foo bar"))
                                            }
                                            fmt.Println(fi38, v27)
                                            switch (fi28 & fi28) {
                                            case 127:
                                                fmt.Println(fi38)
                                            }
                                        }
                                        v27 -= 0
                                    }
                                    v65 := GenS2{F0: "foo bar", F1: "a", F2: -128}
                                    var v66 uint = 2
                                    var v67 int16 = -710
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                }
                                v3 = v27
                            }
                            v68 := GenS3{F0: v4, F1: 1, F2: fi28}
                            _ = v68
                        }
                        v3 += (v27 ^ (v3 & v3))
                        _ = fi38
                    }
                    _ = v37
                }
            case -885:
                v69 := GenS2{F0: "", F1: "a", F2: 42}
                _ = v69
            }
        case -2147483648:
            var v70 int64 = v4
            _ = v70
        }
        _ = fi28
    }
    if false {
        fmt.Println(v4, v4)
    }
    v71 := GenS1{F0: 0, F1: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
