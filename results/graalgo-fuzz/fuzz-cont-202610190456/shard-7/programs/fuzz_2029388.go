package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 int64; F1 string; F2 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int64 }
func genFn0() uint32 {
    if false {
        var v4 uint16 = 10
        var v5 uint = 42
        var v6 uint8 = 1
        _ = v4
        _ = v5
        _ = v6
    }
    switch 278 {
    case 0:
        var v7 int32 = 346
        _ = v7
    case 32768:
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 float64 = (0.0 - 1e10)
                v10 += v10
                _ = fi9
                _ = v10
            }
            var v11 int = fi8
            fmt.Println(v11, v11)
            _ = fi8
            _ = v11
        }
    case 9223372036854775807:
        var v12 uint64 = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int8 = 100
        fmt.Println(v14, v14)
        for fi15 := 0; fi15 < 20; fi15++ {
            v14 = -1
            var v16 int = 9223372036854775807
            if (2147483647 != 254) {
                v17 := GenS1{F0: 1, F1: 2}
                v18 := GenS2{F0: 0, F1: 42, F2: 18446744073709551615}
                _ = v17
                _ = v18
            } else {
                v19 := GenS2{F0: 65535, F1: 812, F2: 0}
                var v20 int32 = 103
                if ((fi15 + v16) != (fi15 >> ((0) & 63))) {
                    v21 := GenS2{F0: 65535, F1: 22, F2: 904}
                    if (v16 <= (fi15 | -1)) {
                        v22 := GenS3{F0: fi15, F1: "\n\t", F2: 101}
                        fmt.Println(v20, fi15, v14)
                        _ = v22
                    }
                    fmt.Println(v20)
                    _ = v21
                }
                _ = v19
                _ = v20
            }
            var v23 uint32 = 902
            _ = fi15
            _ = v16
            _ = v23
        }
        _ = fi13
        _ = v14
    }
    var v24 float64 = 0.0
    v25 := GenS4{F0: -471, F1: ("" + "hello"), F2: 2}
    _ = v24
    _ = v25
    return 42
}
func genFn1(p0 uint) int32 {
    if true {
        p0 = p0
    } else {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int64 = 0
            var v28 uint = p0
            v29 := []int{}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            p0 = 65535
            _ = fi30
        }
        if false {
            p0 -= 630
            if false {
                var v31 string = ("a" + ("hello" + "hello"))
                v32 := GenS4{F0: 0, F1: v31, F2: 42}
                var v33 int = 4294967295
                _ = v31
                _ = v32
                _ = v33
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 float64 = float64((float32(1.5) - 1e10))
                    for fi36 := 0; fi36 < 10; fi36++ {
                        fmt.Println(fi34)
                        switch fi34 {
                        case 9223372036854775807:
                            fmt.Println(fi36)
                        default:
                            var v37 uint16 = 65535
                            _ = v37
                        }
                        var v38 int64 = 580
                        _ = fi36
                        _ = v38
                    }
                    _ = fi34
                    _ = v35
                }
            }
        }
    }
    if true {
        v39 := []int{-269}
        var v40 uint64 = 2
        _ = v39
        _ = v40
    }
    v41 := []int{-991, -965, 128}
    _ = p0
    _ = v41
    return 42
}

func main() {
    v1 := []int{32768}
    v2 := make(map[int]int)
    v3 := []int{32768, 32768}
    switch "" {
    case "a":
        if true {
            switch "\n\t" {
            case "hello":
                for fi42 := 0; fi42 < 11; fi42++ {
                    v1 = append(v1, (fi42 * fi42))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v43 uint8 = 62
                    _ = fi42
                    _ = v43
                }
            case "foo bar":
                for fi44 := 0; fi44 < 2; fi44++ {
                    if true {
                        var v45 uint = 18446744073709551615
                        v46 := v2[fi44]
                        _ = v45
                        _ = v46
                    } else {
                        var v47 float32 = 0.0
                        v47 -= float32((-2.25 - -2.25))
                        var v48 int64 = 1
                        _ = v47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 20; fi49++ {
                        var v50 bool = (false && true)
                        _ = fi49
                        _ = v50
                    }
                    _ = fi44
                }
            default:
                var v51 uint8 = 89
                _ = v51
            }
        } else {
            for fi52 := 0; fi52 < 1; fi52++ {
                var v53 bool = false
                _ = fi52
                _ = v53
            }
            for fi54 := 0; fi54 < 15; fi54++ {
                var v55 int = 446
                _ = fi54
                _ = v55
            }
            var v56 float32 = float32(3.14159)
            _ = v56
        }
    case "\n\t":
        var v57 uint32 = 4294967295
        _ = v57
    }
    if ("hello" > "foo bar") {
        var v58 int = 343
        _ = v58
    } else {
        v59 := GenS4{F0: 749, F1: "foo bar", F2: 32767}
        _ = v59
    }
    v60 := GenS5{F0: 9, F1: 1}
    var v61 int64 = 0
    switch -9223372036854775808 {
    case 0:
        v61 = -1
    case 32767:
        v62 := GenS2{F0: 65535, F1: 1, F2: 299}
        _ = v62
    default:
        var v63 int16 = 1000
        _ = v63
    }
    switch "\n\t" {
    case "":
        for fi64 := 0; fi64 < 17; fi64++ {
            fmt.Println(v61, v61)
            v65 := v2[(fi64 / ((42) | 1))]
            var v66 string = "hello"
            _ = fi64
            _ = v65
            _ = v66
        }
    case "hello":
        for fi67 := 0; fi67 < 3; fi67++ {
            switch (fi67 - (fi67 ^ fi67)) {
            case 405:
                switch ((("\n\t" + "\n\t") + "a") + "a") {
                case "\n\t":
                    for fi68 := 0; fi68 < 10; fi68++ {
                        v69 := GenS6{F0: -128, F1: 2147483647}
                        v70 := GenS6{F0: int8(fi67), F1: 385}
                        if true {
                            var v71 int16 = 0
                            if (false || (genFn1(2) > genFn1(18446744073709551615))) {
                                var v72 string = "a"
                                var v73 int16 = v71
                                var v74 uint = 398
                                _ = v72
                                _ = v73
                                _ = v74
                            } else {
                                v61 = (^(v61))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v71)
                                switch ("a" + "") {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v75 bool = (true || (((!true) && (!(v61 >= (v61 / (((v61 / (((v61 / ((v61) | 1))) | 1))) | 1))))) || (!(!false))))
                                    v61 += 865
                                    _ = v75
                                case "foo bar":
                                    fmt.Println(v61)
                                    fmt.Println(v71, fi67, fi68)
                                default:
                                    fmt.Println(v61, v61)
                                }
                                var v76 int32 = 254
                                var v77 int32 = (v76 + (-(v76)))
                                fmt.Println(v76, fi68)
                                var v78 uint8 = 0
                                var v79 bool = (true || ((!false) && false))
                                v80 := GenS4{F0: v61, F1: "hello", F2: 32767}
                                _ = v76
                                _ = v77
                                _ = v78
                                _ = v79
                                _ = v80
                            }
                            _ = v71
                        }
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                case "a":
                    var v81 bool = true
                    _ = v81
                case "":
                    for fi82 := 0; fi82 < 15; fi82++ {
                        v61 = -570
                        fmt.Println(v61, fi67, fi67)
                        _ = fi82
                    }
                }
            case 0:
                if (0 == 811) {
                    v61 = v61
                    v61 += (-(v61))
                }
            }
            v61 += (v61 >> ((65535) & 63))
            _ = fi67
        }
    case "a":
        for fi83 := 0; fi83 < 19; fi83++ {
            fmt.Println(fi83)
            var v84 uint64 = 0
            if len(v1) > 0 {
                v1[0] = -2147483648
            }
            _ = fi83
            _ = v84
        }
    default:
        v2[1] = -9223372036854775808
    }
    v85 := GenS7{F0: v61}
    var v86 uint32 = 820
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v61:", v61)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v85
    _ = v86
}
