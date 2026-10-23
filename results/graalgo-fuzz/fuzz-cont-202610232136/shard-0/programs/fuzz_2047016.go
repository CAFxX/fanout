package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int; F2 string }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 int = 42
    var v2 int16 = 709
    v3 := make(map[int]string)
    var v4 string = "a"
    v1 += v1
    if true {
        v5 := []int{}
        var v6 uint32 = 456
        _ = v5
        _ = v6
    } else {
        var v7 uint64 = 18446744073709551615
        fmt.Println(v2, v2, v2)
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: -1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := v3[v1]
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int32 = -1
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := []int{}
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: 2, F1: 795, F2: 65535}
        _ = fi12
        _ = v13
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    v3[v1] = (v4 + (("\n\t" + v4) + v4))
    v2 = v2
    switch (v1 - 32767) {
    case -1:
        if true {
            v2 = (^(v2))
        } else {
            var v18 int = v1
            for fi19 := 0; fi19 < 19; fi19++ {
                v17 = (v17 ^ v17)
                if (42 <= 255) {
                    v20 := GenS2{F0: 2, F1: 503, F2: 65535}
                    if true {
                        var v21 int8 = 100
                        delete(v3, 4294967295)
                        fmt.Println(v18, fi19)
                        _ = v21
                    } else {
                        v2 -= (-(v2))
                        v18 += (v18 & -1)
                    }
                    if (32767 >= 105) {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if (2 < 266) {
                                switch v1 {
                                case -2147483648:
                                    var v23 float32 = 0.0
                                    _ = v23
                                case -253:
                                    var v24 uint64 = 0
                                    _ = v24
                                }
                            } else {
                                v2 = (^(v2))
                                switch v4 {
                                case "hello":
                                    v2 = (v2 * 42)
                                case "a":
                                    v1 += (fi19 / (((v1 + -2147483648)) | 1))
                                case "\n\t":
                                    fmt.Println(fi19, fi19)
                                default:
                                    var v25 int32 = 2147483647
                                    _ = v25
                                }
                            }
                            fmt.Println(fi22, v17, fi19)
                            v1 = (-(v1))
                            var v26 uint16 = 46
                            _ = fi22
                            _ = v26
                        }
                        if true {
                            fmt.Println(v18, v18)
                        }
                        var v27 uint = 1
                        _ = v27
                    }
                    _ = v20
                }
                var v28 int16 = 42
                _ = fi19
                _ = v28
            }
            _ = v18
        }
    default:
        switch v4 {
        case "hello":
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 int32 = 1
                v4 = ""
                _ = fi29
                _ = v30
            }
        case "":
            var v31 float32 = ((1.5 + float32((1.5 - 1e10))) + ((3.14159 * -2.25) + 1e10))
            _ = v31
        case "\n\t":
            v32 := GenS3{F0: 128, F1: 32767}
            _ = v32
        default:
            var v33 uint8 = 79
            _ = v33
        }
    }
    if false {
        v4 = "a"
        switch 32767 {
        case -2147483648:
            if (false || (!((1e10 - (-2.25 - 0.0)) <= float32(-2.25)))) {
                for fi34 := 0; fi34 < 8; fi34++ {
                    fmt.Println(v17)
                    v35 := GenS4{F0: 42}
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 3; fi36++ {
                    fmt.Println(fi36, v17, v4)
                    _ = fi36
                }
                if (!(65535 >= 2)) {
                    v17 -= -9223372036854775808
                    switch (v1 - int(v17)) {
                    case -1:
                        var v37 float64 = -2.25
                        _ = v37
                    case 0:
                        v2 += int16(v1)
                    case 32768:
                        var v38 int = v1
                        _ = v38
                    }
                    v39 := GenS5{F0: 911, F1: v1, F2: ((((v4 + ("a" + v4)) + "a") + "\n\t") + ((((v4 + ("foo bar" + ("\n\t" + ""))) + "foo bar") + "foo bar") + (v4 + "")))}
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 6; fi40++ {
                    v41 := []int{v1, (fi40 + fi40)}
                    for fi42 := 0; fi42 < 18; fi42++ {
                        var v43 uint32 = 2147483648
                        if true {
                            var v44 uint32 = 65535
                            var v45 bool = ((3.14159 * 1e10) >= float64(-2.25))
                            fmt.Println(fi40, v43, v2)
                            var v46 uint16 = 0
                            switch "" {
                            case "a":
                                v2 += 32767
                            case "foo bar":
                                v47 := GenS4{F0: 32767}
                                _ = v47
                            case "hello":
                                v48 := GenS6{F0: 42, F1: v4}
                                _ = v48
                            }
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            v4 = "\n\t"
                            var v49 uint = 509
                            var v50 int8 = -1
                            _ = v49
                            _ = v50
                        }
                        _ = fi42
                        _ = v43
                    }
                    var v51 int8 = 39
                    _ = fi40
                    _ = v41
                    _ = v51
                }
                v17 = v17
                v4 = ("hello" + ("foo bar" + (v4 + ((((v4 + ("a" + ((v4 + (v4 + "\n\t")) + v4))) + (v4 + "foo bar")) + "\n\t") + "a"))))
            }
        }
        v2 = v2
    } else {
        v52 := v3[v1]
        for fi53 := 0; fi53 < 13; fi53++ {
            v4 = "hello"
            if true {
                fmt.Println(fi53)
                var v54 int16 = -1
                _ = v54
            }
            var v55 float64 = -2.25
            _ = fi53
            _ = v55
        }
        v52 = v52
        _ = v52
    }
    var v56 int32 = -2147483648
    fmt.Println(v4, v56, v2)
    v2 -= v2
    v57 := GenS7{F0: (v17 | v17), F1: false, F2: v4}
    var v58 uint64 = 922
    switch "\n\t" {
    case "a":
        for fi59 := 0; fi59 < 2; fi59++ {
            var v60 int = (v1 + (v1 * v1))
            v58 = (v58 % ((v58) | 1))
            _ = fi59
            _ = v60
        }
    case "\n\t":
        v61 := []int{(v1 & 250), (^(v1)), (v1 * (v1 | -9223372036854775808))}
        _ = v61
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
}
