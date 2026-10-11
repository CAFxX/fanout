package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS8 struct { F0 int; F1 string }
func genFn0(p0 uint64, p1 int16, p2 string) uint16 {
    var v4 float32 = -2.25
    p1 = (p1 / ((p1) | 1))
    v5 := GenS1{F0: -973}
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int) int64 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int64 = -9223372036854775808
        v10 := []int{p0, (fi8 - p0)}
        if true {
            if len(v10) > 0 {
                v10[0] = -9223372036854775808
            }
            if ((fi8 / (((fi8 % (((fi8 % (((p0 ^ 32768)) | 1))) | 1))) | 1)) != (p0 ^ 993)) {
                if true {
                    v11 := GenS1{F0: (-(p0))}
                    _ = len(v10)
                    _ = v11
                }
                p0 -= (p0 * (fi8 >> ((975) & 63)))
            } else {
                v10 = append(v10, -1)
                v9 -= v9
            }
        } else {
            switch (-(p0)) {
            case -1:
                var v12 bool = true
                _ = v12
            case 32767:
                if len(v10) > 0 {
                    v10[0] = p0
                }
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 = (fi13 % (((fi13 / ((p0) | 1))) | 1))
        _ = fi13
    }
    if (("a" + "hello") >= "hello") {
        for fi14 := 0; fi14 < 11; fi14++ {
            p0 -= -1
            _ = fi14
        }
        if true {
            fmt.Println(p0)
            var v15 int64 = -962
            var v16 uint16 = 463
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: p0}
        _ = v17
    }
    _ = p0
    return -37
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v18 int64 = 9223372036854775807
    v19 := GenS2{F0: 42, F1: 446, F2: 2}
    var v20 float32 = (((1.5 - 3.14159) + float32(((float64(0.0) + float64(1e10)) / 1.0))) / (1e10 + 1e10))
    _ = len(v1)
    v20 = (0.0 - (((v20 - v20) * (0.1 - v20)) + 1.5))
    switch (("a" + ("foo bar" + (("a" + (("\n\t" + "\n\t") + "foo bar")) + ""))) + "foo bar") {
    case "a":
        v18 -= (v18 & (v18 & 839))
    case "hello":
        var v21 float32 = v20
        _ = v21
    }
    v18 = genFn1(-9223372036854775808)
    var v22 float32 = -2.25
    if (-622 <= -2147483648) {
        if true {
            v23 := GenS2{F0: 70, F1: 42, F2: 977}
            _ = v23
        } else {
            var v24 float32 = 3.14159
            _ = v24
        }
        v20 = (0.0 / (v20 - 0.0))
    } else {
        if true {
            var v25 int8 = -45
            v25 = v25
            fmt.Println(v22, v25)
            _ = v25
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            v20 -= 1.5
            _ = len(v3)
            v22 = 0.1
            _ = fi26
        }
        if true {
            fmt.Println(v18)
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 bool = (((v18 & (v18 * v18)) < genFn1(fi27)) && (v18 >= v18))
                _ = fi27
                _ = v28
            }
            switch 127 {
            case 0:
                v29 := GenS2{F0: 2, F1: 255, F2: 42}
                _ = v29
            case 4294967295:
                v30 := GenS3{F0: 2}
                _ = v30
            default:
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 float32 = (v20 * float32(0.1))
                    _ = fi31
                    _ = v32
                }
            }
        } else {
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS1{F0: 32767}
                fmt.Println(v18, v20, v22)
                v35 := []int{}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    }
    v22 = 1e10
    if (32768 != 32767) {
        v36 := GenS3{F0: 1}
        if false {
            switch "hello" {
            case "a":
                v37 := GenS1{F0: -635}
                _ = v37
            case "foo bar":
                var v38 uint32 = 1
                _ = v38
            }
            for fi39 := 0; fi39 < 3; fi39++ {
                switch ((("hello" + "hello") + "foo bar") + "\n\t") {
                case "a":
                    var v40 uint64 = 680
                    _ = v40
                case "":
                    if ((!true) && ((!true) && ((-2.25 * ((3.14159 - 1e10) - 0.0)) != (1e10 / 1.5)))) {
                        for fi41 := 0; fi41 < 16; fi41++ {
                            v42 := []int{}
                            v18 = (v18 / ((v18) | 1))
                            switch (fi41 - 0) {
                            case 0:
                                if ((!(!(617 < 566))) || (!false)) {
                                    v18 -= v18
                                    fmt.Println(v20, fi39)
                                } else {
                                    fmt.Println(fi39, fi39)
                                    v43 := GenS3{F0: 2}
                                    if len(v42) > 0 {
                                        v42[0] = (fi41 << ((18446744073709551615) & 63))
                                    }
                                    _ = v43
                                }
                            case -9223372036854775808:
                                switch -821 {
                                case -2147483648:
                                    var v44 uint64 = 882
                                    _ = v44
                                default:
                                    v20 = ((-2.25 * float32((((1.5 - (1e10 - 1.5)) + 0.0) + ((((float64(3.14159) + 0.0) / -2.25) * 3.14159) / -2.25)))) / float32((-2.25 + (float64((v22 + 1e10)) + float64(1e10)))))
                                }
                            case 1:
                                var v45 int16 = 1
                                _ = v45
                            }
                            _ = fi41
                            _ = v42
                        }
                        var v46 uint64 = 18446744073709551615
                        _ = v46
                    } else {
                        v1[(fi39 ^ fi39)] = ("\n\t" + "")
                        fmt.Println(fi39)
                        var v47 uint8 = 128
                        _ = v47
                    }
                default:
                    var v48 int32 = 317
                    _ = v48
                }
                _ = fi39
            }
        } else {
            var v49 int16 = 0
            _ = v49
        }
        _ = v36
    } else {
        var v50 float32 = (v20 / -2.25)
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 uint16 = 32767
            for fi53 := 0; fi53 < 13; fi53++ {
                v54 := GenS4{F0: 86, F1: 4294967295}
                _ = fi53
                _ = v54
            }
            v55 := GenS5{F0: -286, F1: "a"}
            _ = fi51
            _ = v52
            _ = v55
        }
        v20 += float32((0.1 - (1e10 * 1e10)))
        _ = v50
    }
    var v56 int64 = (v18 & (^(v18)))
    switch "foo bar" {
    case "foo bar":
        for fi57 := 0; fi57 < 7; fi57++ {
            for fi58 := 0; fi58 < 5; fi58++ {
                v59 := GenS6{F0: fi57, F1: 128, F2: 32767}
                if false {
                    v60 := GenS7{F0: 255, F1: 587, F2: v18}
                    v61 := GenS8{F0: 964, F1: ("" + "hello")}
                    _ = v60
                    _ = v61
                } else {
                    if (!false) {
                        var v62 float64 = (0.0 / 1.5)
                        fmt.Println(v62, v22)
                        _ = v62
                    } else {
                        v56 = (v18 | 255)
                        v20 -= v22
                        v63 := GenS3{F0: 42}
                        _ = v63
                    }
                }
                v64 := []int{fi57, 42}
                _ = fi58
                _ = v59
                _ = v64
            }
            var v65 float32 = (float32(((float64(0.1) + (float64(v20) - 1e10)) - (0.1 + float64(v20)))) / v22)
            _ = fi57
            _ = v65
        }
    case "a":
        if false {
            switch "foo bar" {
            case "a":
                var v66 uint64 = 42
                _ = v66
            default:
                v67 := []int{9223372036854775807, 42, 32768}
                _ = v67
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v56
}
