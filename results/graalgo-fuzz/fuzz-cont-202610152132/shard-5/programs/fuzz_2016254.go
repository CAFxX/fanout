package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int32; F1 uint8; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 int16; F1 int64; F2 uint8 }
type GenS9 struct { F0 bool; F1 int32; F2 int32 }
type GenS10 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 uint, p1 uint) int {
    p0 = (p1 << ((p0) & 63))
    p0 = p1
    v5 := []int{128, 32767}
    _ = p0
    _ = p1
    _ = v5
    return 397
}
func genFn1(p0 int32, p1 int64, p2 uint8) int8 {
    var v6 float64 = 0.0
    v7 := []int{genFn0(1, 152), -957}
    v6 = v6
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint32 = 2
    var v2 int8 = -128
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 int64 = -602
                switch ((("foo bar" + (("a" + "hello") + "hello")) + "") + "") {
                case "hello":
                    var v11 int = fi9
                    _ = v11
                case "a":
                    v10 = (v10 - 573)
                case "":
                    if ((v1 <= v1) && true) {
                        if false {
                            var v12 bool = (!true)
                            switch (fi8 >> ((18446744073709551615) & 63)) {
                            case 240:
                                var v13 bool = (42 == -1)
                                _ = v13
                            case -1:
                                v14 := GenS1{F0: 36, F1: 0, F2: "hello"}
                                fmt.Println(v10, v12, v12)
                                switch ("" + "") {
                                case "\n\t":
                                    var v15 int = fi8
                                    _ = v15
                                case "hello":
                                    switch (fi9 >> ((203) & 63)) {
                                    case 128:
                                        v10 -= (v10 << ((4294967295) & 63))
                                    case 1:
                                        fmt.Println(v2)
                                    }
                                }
                                _ = v14
                            case 32767:
                                v3["hello"] = false
                            default:
                                v16 := []int{(fi8 / (((fi8 & fi8)) | 1))}
                                _ = v16
                            }
                            v17 := GenS2{F0: "hello"}
                            _ = v12
                            _ = v17
                        }
                    } else {
                        v18 := GenS3{F0: "foo bar", F1: 0, F2: false}
                        v19 := GenS4{F0: 4294967295, F1: (fi8 | 42)}
                        v2 = (-(v2))
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi9
                _ = v10
            }
            var v20 uint16 = 421
            var v21 uint32 = v1
            _ = fi8
            _ = v20
            _ = v21
        }
    case "":
        v22 := GenS5{F0: 42, F1: 83, F2: (!false)}
        _ = v22
    }
    var v23 uint8 = 1
    v2 += 42
    fmt.Println(v1, v1, v1)
    if (-599 <= -150) {
        for fi24 := 0; fi24 < 18; fi24++ {
            fmt.Println(v23, fi24, v23)
            v25 := GenS6{F0: 32767}
            v26 := GenS1{F0: 65535, F1: 1, F2: "foo bar"}
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS7{F0: 9223372036854775807, F1: 0}
                v4["\n\t"] = (255 > 32767)
                v23 = 100
                _ = fi28
                _ = v29
            }
            v30 := GenS7{F0: 735, F1: 296}
            if (genFn0(2, 65535) == fi27) {
                var v31 int = genFn0(1, 18446744073709551615)
                _ = v31
            } else {
                for fi32 := 0; fi32 < 1; fi32++ {
                    switch (("a" + ((("a" + ("hello" + "hello")) + "a") + "hello")) + ("hello" + "\n\t")) {
                    case "hello":
                        v1 += (v1 | v1)
                    case "\n\t":
                        fmt.Println(v2)
                    case "a":
                        v2 += genFn1(225, int64(v1), v23)
                    }
                    v33 := []int{(fi27 + genFn0(0, 65535)), fi27}
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi27
            _ = v30
        }
        var v34 int8 = v2
        _ = v34
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 int8 = v2
            _ = fi35
            _ = v36
        }
        switch 128 {
        case 942:
            var v37 float32 = 0.1
            _ = v37
        default:
            v38 := GenS8{F0: -210, F1: int64(v2), F2: v23}
            _ = v38
        }
        fmt.Println(v23)
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            v41 := GenS8{F0: 1, F1: 1, F2: (v23 | v23)}
            _ = fi40
            _ = v41
        }
        v1 = v1
        for fi42 := 0; fi42 < 3; fi42++ {
            switch "hello" {
            case "hello":
                v2 = v2
            case "":
                switch ("a" + ("a" + "foo bar")) {
                case "foo bar":
                    v43 := []int{genFn0(255, 544), fi42, -1, 402}
                    _ = v43
                }
            case "a":
                fmt.Println(v2)
            default:
                v44 := []int{fi39, fi42, (fi42 ^ 743)}
                _ = v44
            }
            v23 = 1
            v45 := []int{(fi42 - genFn0(814, 874))}
            _ = fi42
            _ = v45
        }
        _ = fi39
    }
    v1 = v1
    for fi46 := 0; fi46 < 11; fi46++ {
        if ((-2147483648 != -528) || (531 <= 2)) {
            for fi47 := 0; fi47 < 8; fi47++ {
                if (!((v1 / (((^(v1))) | 1)) < v1)) {
                    if ((^(v23)) != v23) {
                        var v48 float32 = 0.0
                        v49 := GenS9{F0: true, F1: 42, F2: -657}
                        _ = v48
                        _ = v49
                    }
                    v50 := GenS10{F0: -560, F1: 2, F2: v2}
                    v3["foo bar"] = (2 != v23)
                    _ = v50
                }
                _ = fi47
            }
        } else {
            switch "foo bar" {
            case "a":
                for fi51 := 0; fi51 < 10; fi51++ {
                    if (!((((((1e10 * 1e10) * (1e10 / 1e10)) * 3.14159) - 0.1) < 1.5) && (!(genFn1(-568, 1, v23) != v2)))) {
                        var v52 uint32 = 2
                        _ = v52
                    }
                    var v53 uint64 = 42
                    var v54 uint8 = v23
                    switch "a" {
                    case "a":
                        var v55 uint = 2
                        if (-1 <= 1) {
                            v56 := []int{fi51, genFn0((v55 | (v55 * 255)), (v55 - (v55 + v55))), (^(fi51))}
                            v1 = (v1 + 2147483648)
                            fmt.Println(v1, v53, v1)
                            _ = v56
                        } else {
                            _ = len(v3)
                            v1 -= (^(v1))
                            var v57 uint64 = (v53 * 1)
                            _ = v57
                        }
                        _ = v55
                    case "hello":
                        var v58 uint64 = v53
                        _ = v58
                    case "\n\t":
                        fmt.Println(v54, v54)
                    }
                    _ = fi51
                    _ = v53
                    _ = v54
                }
            case "":
                if true {
                    for fi59 := 0; fi59 < 3; fi59++ {
                        v60 := GenS10{F0: int64(v1), F1: 0, F2: (-(v2))}
                        v1 -= v1
                        _ = fi59
                        _ = v60
                    }
                }
            case "foo bar":
                switch "" {
                case "hello":
                    var v61 uint32 = 0
                    _ = v61
                case "":
                    v2 = (v2 / (((v2 - (^(v2)))) | 1))
                }
            default:
                var v62 float32 = (-2.25 / -2.25)
                _ = v62
            }
        }
        _ = fi46
    }
    fmt.Println(v23, v2, v2)
    if true {
        if (551 <= 4294967295) {
            if (genFn1(698, 69, v23) > -1) {
                v23 = 128
            } else {
                var v63 bool = true
                _ = v63
            }
            var v64 uint32 = v1
            _ = v64
        }
        fmt.Println(v23, v1, v1)
    } else {
        v65 := []int{1, -1, 902, genFn0(255, 1)}
        v2 -= (v2 * -1)
        _ = v65
    }
    if true {
        if (false || true) {
            v23 -= 84
            v66 := GenS6{F0: 32767}
            _ = v66
        }
        var v67 int16 = 42
        _ = v67
    } else {
        v68 := GenS3{F0: ("hello" + "\n\t"), F1: 0, F2: (54 > int8(v23))}
        v69 := GenS2{F0: "hello"}
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 3; fi70++ {
        fmt.Println(fi70, v2, v23)
        _ = fi70
    }
    var v71 int16 = 560
    for fi72 := 0; fi72 < 2; fi72++ {
        var v73 int8 = v2
        v74 := []int{}
        _ = fi72
        _ = v73
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v71
}
