package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 bool) int16 {
    if p0 {
        v4 := GenS1{F0: "", F1: 2}
        var v5 float64 = 1.5
        _ = v4
        _ = v5
    } else {
        if (2 < 2) {
            var v6 int8 = 119
            var v7 uint = 255
            var v8 uint32 = uint32(v7)
            _ = v6
            _ = v7
            _ = v8
        } else {
            switch "\n\t" {
            case "\n\t":
                switch "hello" {
                case "":
                    if (193 == 0) {
                        var v9 int64 = -684
                        if (!(-48 != -886)) {
                            v10 := GenS1{F0: "\n\t", F1: 212}
                            switch -729 {
                            case 963:
                                v11 := []int{128, -2147483648, -53, 0}
                                _ = v11
                            }
                            _ = v10
                        } else {
                            var v12 bool = p0
                            v13 := GenS1{F0: "a", F1: 42}
                            fmt.Println(p0, v12, p0)
                            _ = v12
                            _ = v13
                        }
                        var v14 float64 = (-2.25 + (-2.25 / (1.5 + 0.0)))
                        _ = v14
                        _ = v9
                    }
                case "hello":
                    switch "foo bar" {
                    case "a":
                        for fi15 := 0; fi15 < 10; fi15++ {
                            fmt.Println(p0, p0)
                            v16 := GenS1{F0: "hello", F1: 135}
                            var v17 uint64 = 42
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        var v18 uint64 = 4294967295
                        _ = v18
                    default:
                        v19 := []int{-376, -9223372036854775808, 1}
                        _ = v19
                    }
                default:
                    p0 = true
                }
            case "hello":
                var v20 uint8 = 85
                _ = v20
            case "foo bar":
                v21 := []int{32768, 0}
                _ = v21
            default:
                var v22 int = 42
                _ = v22
            }
        }
        p0 = false
        if false {
            switch 128 {
            case -661:
                v23 := []int{9223372036854775807, 42, -528}
                _ = v23
            }
            if (p0 && false) {
                v24 := []int{32767, -9223372036854775808, 1, -420}
                if (!p0) {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        switch ("a" + ("\n\t" + (("" + "foo bar") + ""))) {
                        case "foo bar":
                            if (-119 >= -1) {
                                var v26 int = fi25
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    var v28 uint = 1
                                    var v29 int16 = 32767
                                    v29 -= (v29 % (((-(v29))) | 1))
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                p0 = p0
                                _ = v26
                            }
                        default:
                            var v30 float64 = (0.1 + (0.1 / 1.5))
                            _ = v30
                        }
                        _ = fi25
                    }
                    p0 = false
                }
                p0 = p0
                _ = v24
            } else {
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 uint32 = 2
                    switch "" {
                    case "\n\t":
                        switch ("hello" + "hello") {
                        case "foo bar":
                            v33 := []int{}
                            _ = v33
                        }
                    case "a":
                        v34 := []int{(-(fi31))}
                        _ = v34
                    case "hello":
                        v35 := GenS2{F0: 32767, F1: "a", F2: "hello"}
                        _ = v35
                    default:
                        if false {
                            if (-2147483648 < -9223372036854775808) {
                                for fi36 := 0; fi36 < 7; fi36++ {
                                    v32 += (v32 ^ v32)
                                    v37 := GenS1{F0: "a", F1: 0}
                                    _ = fi36
                                    _ = v37
                                }
                            } else {
                                fmt.Println(v32, v32)
                                for fi38 := 0; fi38 < 17; fi38++ {
                                    v39 := []int{fi31, fi38, 4294967295}
                                    var v40 float64 = 1.5
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                var v41 uint16 = 0
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 12; fi42++ {
                                var v43 string = (("\n\t" + "") + "a")
                                v32 = v32
                                if true {
                                    fmt.Println(v43, v43, p0)
                                    var v44 uint64 = 1
                                    var v45 float32 = (float32(0.0) + (3.14159 + 1.5))
                                    _ = v44
                                    _ = v45
                                } else {
                                    var v46 uint64 = 813
                                    _ = v46
                                }
                                _ = fi42
                                _ = v43
                            }
                        }
                    }
                    v47 := []int{fi31}
                    _ = fi31
                    _ = v32
                    _ = v47
                }
                if false {
                    var v48 string = ("foo bar" + "")
                    fmt.Println(p0)
                    if false {
                        var v49 bool = p0
                        v49 = v49
                        _ = v49
                    } else {
                        p0 = p0
                        v50 := []int{255, -524, -70, -1}
                        _ = v50
                    }
                    _ = v48
                }
            }
        }
    }
    p0 = p0
    v51 := GenS2{F0: 42, F1: "foo bar", F2: "hello"}
    fmt.Println(p0)
    _ = p0
    _ = v51
    return 130
}

func main() {
    var v1 string = "hello"
    var v2 int8 = 0
    var v3 int8 = 31
    v52 := GenS3{F0: 124, F1: 127}
    if (-907 != -863) {
        v3 = 77
        v53 := []int{}
        var v54 uint = 0
        _ = v53
        _ = v54
    } else {
        fmt.Println(v1, v2, v3)
        v2 = (-(v3))
    }
    var v55 uint32 = 0
    fmt.Println(v2, v2, v2)
    switch v1 {
    case "":
        var v56 float32 = 0.0
        _ = v56
    case "a":
        var v57 float64 = (float64((0.1 * 1e10)) / float64(float32((1e10 - 0.0))))
        _ = v57
    default:
        v2 = -110
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        fmt.Println(v3, v1)
        var v59 int8 = (v3 % ((64) | 1))
        _ = fi58
        _ = v59
    }
    var v60 bool = false
    var v61 int64 = -9223372036854775808
    fmt.Println(v55)
    for fi62 := 0; fi62 < 18; fi62++ {
        var v63 float64 = -2.25
        if v60 {
            var v64 uint16 = 255
            _ = v64
        } else {
            v65 := GenS2{F0: genFn0((!((-(v61)) > v61))), F1: "\n\t", F2: v1}
            for fi66 := 0; fi66 < 13; fi66++ {
                v67 := GenS3{F0: 2147483647, F1: (fi66 - (fi62 + (fi66 ^ fi66)))}
                _ = fi66
                _ = v67
            }
            _ = v65
        }
        _ = fi62
        _ = v63
    }
    var v68 int = -328
    v3 = (v3 + (v2 >> ((0) & 63)))
    v69 := GenS2{F0: 42, F1: v1, F2: (("hello" + ((v1 + v1) + (v1 + ((v1 + (v1 + v1)) + (v1 + v1))))) + (("foo bar" + (v1 + v1)) + v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v55
    _ = v60
    _ = v61
    _ = v68
    _ = v69
}
