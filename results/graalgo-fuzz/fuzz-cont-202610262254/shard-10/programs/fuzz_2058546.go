package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 string }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar"}
    var v4 int8 = 100
    var v5 uint32 = 423
    var v6 uint8 = 42
    v7 := GenS1{F0: 18446744073709551615}
    switch 128 {
    case 127:
        _ = len(v3)
    default:
        v4 = (v4 / (((v4 ^ (-(v4)))) | 1))
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v4 += int8(fi8)
        if false {
            var v9 float32 = ((0.1 / 1.0) * 3.14159)
            v9 -= 0.0
            _ = v9
        }
        _ = fi8
    }
    delete(v2, 738)
    switch ((("foo bar" + "hello") + "") + ("hello" + "")) {
    case "hello":
        _ = len(v2)
    case "a":
        v5 -= 2147483647
    default:
        v10 := v1[(v5 * 0)]
        _ = v10
    }
    v5 += v5
    for fi11 := 0; fi11 < 1; fi11++ {
        v6 -= v6
        for fi12 := 0; fi12 < 20; fi12++ {
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "hello")
            }
            v4 += 127
            _ = len(v3)
            v13 := GenS1{F0: 18446744073709551615}
            switch ("foo bar" + "hello") {
            case "":
                v6 -= (v6 % (((v6 / (((v6 % ((v6) | 1))) | 1))) | 1))
            default:
                v14 := GenS1{F0: 2}
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        if ((float64(-2.25) - 0.0) < ((1e10 - 1.5) + 1e10)) {
            fmt.Println(v6, fi11)
            fmt.Println(fi11, v6, fi11)
            if (!(2147483647 != 0)) {
                var v15 uint = 1
                var v16 float32 = 1.5
                switch "hello" {
                case "a":
                    v4 += 127
                case "foo bar":
                    if ((v4 << ((v15) & 63)) < (v4 >> ((0) & 63))) {
                        v6 += 127
                        v17 := GenS1{F0: (v15 / ((v15) | 1))}
                        _ = v17
                    }
                }
                _ = v15
                _ = v16
            } else {
                if (-1 > 1) {
                    v18 := GenS2{F0: (18446744073709551615 < 4294967295)}
                    v5 = v5
                    fmt.Println(v5)
                    _ = v18
                } else {
                    var v19 int = 32767
                    _ = v19
                }
                _ = len(v3)
            }
        }
        _ = fi11
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        if len(v3) > 0 {
            v3[0] = ("hello" + "a")
        }
        v4 -= (-(v4))
        _ = fi20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        switch ("" + "foo bar") {
        case "foo bar":
            var v22 uint64 = 405
            _ = v22
        case "\n\t":
            fmt.Println(fi21, v4)
        case "a":
            v4 = (v4 * (v4 * v4))
        default:
            var v23 string = "\n\t"
            _ = v23
        }
        v4 += v4
        for fi24 := 0; fi24 < 10; fi24++ {
            if false {
                var v25 int8 = v4
                v4 += (v25 / ((v25) | 1))
                v4 = (-(v25))
                _ = v25
            } else {
                v4 += v4
                v4 = v4
                var v26 int8 = (v4 / ((v4) | 1))
                _ = v26
            }
            var v27 float32 = 0.1
            _ = fi24
            _ = v27
        }
        _ = fi21
    }
    switch 459 {
    case 128:
        switch 645 {
        case 32768:
            for fi28 := 0; fi28 < 7; fi28++ {
                for fi29 := 0; fi29 < 18; fi29++ {
                    switch fi28 {
                    case 42:
                        switch fi28 {
                        case 825:
                            if (!false) {
                                switch "hello" {
                                case "a":
                                    v30 := GenS3{F0: 219}
                                    _ = v30
                                case "\n\t":
                                    if false {
                                        var v31 uint64 = 1
                                        v4 = v4
                                        v6 = v6
                                        v4 += (v4 + v4)
                                        v2[42] = v5
                                        fmt.Println(v31, fi29)
                                        _ = v31
                                    } else {
                                        v4 -= (v4 ^ v4)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case "foo bar":
                                    var v32 int16 = -32768
                                    var v33 int8 = (v4 % (((v4 + v4)) | 1))
                                    _ = v32
                                    _ = v33
                                default:
                                    v6 = (v6 << ((uint(v6)) & 63))
                                }
                                fmt.Println(fi29, v6)
                            } else {
                                v34 := GenS4{F0: 46, F1: 2147483647, F2: "foo bar"}
                                _ = v34
                            }
                        case 817:
                            delete(v2, v5)
                        default:
                            if true {
                                v6 = 127
                                var v35 int16 = 125
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v36 := GenS5{F0: v6, F1: 40}
                                _ = v35
                                _ = v36
                            } else {
                                fmt.Println(v4, fi29)
                            }
                        }
                    case 32768:
                        if len(v3) > 0 {
                            v3[0] = ("hello" + "hello")
                        }
                    default:
                        fmt.Println(v5)
                    }
                    _ = fi29
                }
                fmt.Println(fi28, v6)
                _ = fi28
            }
        case 341:
            fmt.Println(v4)
        case -499:
            v37 := v2[(v5 - v5)]
            _ = v37
        }
    case 42:
        for fi38 := 0; fi38 < 14; fi38++ {
            for fi39 := 0; fi39 < 13; fi39++ {
                fmt.Println(fi38, fi38)
                _ = fi39
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                if (v5 == v5) {
                    v4 += v4
                    v5 -= (v5 << ((65535) & 63))
                    v4 = (-(v4))
                }
                _ = len(v3)
                fmt.Println(fi38)
                _ = fi40
            }
            _ = fi38
        }
    case 127:
        var v41 int = 32768
        _ = v41
    default:
        for fi42 := 0; fi42 < 4; fi42++ {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                if ((v6 * 1) != (^(v6))) {
                    v1[(v5 | 2)] = (v5 / ((v5) | 1))
                    for fi43 := 0; fi43 < 13; fi43++ {
                        if true {
                            var v44 int8 = (v4 - v4)
                            _ = v44
                        } else {
                            if (((("a" + "a") + "\n\t") != "hello") && (-9223372036854775808 >= 9223372036854775807)) {
                                var v45 uint32 = (v5 % ((728) | 1))
                                var v46 int16 = 0
                                _ = v45
                                _ = v46
                            } else {
                                v47 := GenS5{F0: v6, F1: 1}
                                var v48 float32 = (((0.0 + 1.5) * (-2.25 + (-2.25 - (0.0 - (1.5 - float32(-2.25)))))) - 0.1)
                                v6 += (v6 * 255)
                                _ = v47
                                _ = v48
                            }
                            v49 := GenS3{F0: (^(v5))}
                            var v50 float32 = 1.5
                            var v51 int = -921
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        v4 += -29
                        _ = fi43
                    }
                    fmt.Println(v4)
                }
            case "":
                var v52 int = -833
                _ = v52
            case "foo bar":
                var v53 uint32 = v5
                _ = v53
            }
            if (-431 != 645) {
                var v54 uint64 = 2
                for fi55 := 0; fi55 < 6; fi55++ {
                    if true {
                        var v56 int = fi42
                        var v57 bool = (((!(("foo bar" + "\n\t") < (("hello" + (("hello" + "a") + "\n\t")) + "foo bar"))) || false) || (42 == v4))
                        _ = v56
                        _ = v57
                    } else {
                        v3 = append(v3, "\n\t")
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if true {
                        v5 = (v5 & 448)
                        v4 = (v4 | (v4 & int8(v6)))
                        if (-905 > -506) {
                            v54 += 0
                            v6 = (v6 ^ (v6 / ((v6) | 1)))
                        } else {
                            if false {
                                v54 -= 347
                                if true {
                                    v54 += 676
                                    switch (^(fi55)) {
                                    case 91:
                                        v3 = append(v3, "hello")
                                    default:
                                        v5 += 1
                                    }
                                }
                                v6 = (^(v6))
                            } else {
                                var v58 int8 = v4
                                _ = v58
                            }
                            var v59 uint8 = (^(v6))
                            _ = v59
                        }
                    } else {
                        v5 = (^(v5))
                    }
                    _ = fi55
                }
                _ = v54
            }
            _ = fi42
        }
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v60 int8 = v4
    fmt.Println(v4, v4, v4)
    v6 = (v6 >> ((18446744073709551615) & 63))
    v5 = 475
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v7
}
