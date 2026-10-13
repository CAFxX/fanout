package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8 }
func genFn0(p0 int, p1 bool, p2 uint32) int64 {
    if p1 {
        var v4 uint16 = 0
        if true {
            p2 += 1
        } else {
            fmt.Println(p0, p1, v4)
            v5 := GenS1{F0: 42, F1: 660}
            var v6 uint = 0
            _ = v5
            _ = v6
        }
        _ = v4
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 165, F1: 4294967295}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint64, p1 int32, p2 int64) uint16 {
    var v9 int16 = -32768
    switch "\n\t" {
    case "foo bar":
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint16 = 65535
            v12 := []int{(fi10 >> ((194) & 63)), 167}
            v13 := GenS1{F0: 0, F1: 18446744073709551615}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    case "":
        if (42 > 60) {
            v14 := GenS2{F0: false, F1: 1, F2: -128}
            v15 := GenS3{F0: 64}
            switch 32768 {
            case 105:
                v16 := []int{1, 9223372036854775807, -9223372036854775808, 727}
                _ = v16
            case 0:
                var v17 uint32 = uint32(p0)
                _ = v17
            case -460:
                var v18 float64 = -2.25
                _ = v18
            default:
                var v19 bool = ((!true) && true)
                _ = v19
            }
            _ = v14
            _ = v15
        }
    case "\n\t":
        p1 += -2147483648
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    if (true || true) {
        var v20 string = "a"
        v20 = "\n\t"
        v21 := GenS4{F0: (550 < 4294967295), F1: 42}
        _ = v20
        _ = v21
    }
    var v22 bool = true
    if false {
        var v23 string = "hello"
        for fi24 := 0; fi24 < 12; fi24++ {
            if false {
                v25 := GenS3{F0: 4294967295}
                switch v23 {
                case "foo bar":
                    var v26 int = (fi24 >> ((65535) & 63))
                    _ = v26
                case "\n\t":
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if true {
                            var v28 string = v23
                            v29 := GenS5{F0: v22, F1: (v23 + v23)}
                            _ = v28
                            _ = v29
                        }
                        v22 = (true || v22)
                        v30 := []int{(fi27 ^ 9223372036854775807)}
                        if (true && true) {
                            v31 := GenS2{F0: v22, F1: -32768, F2: 0}
                            _ = v31
                        } else {
                            v32 := v1[(fi27 ^ -9223372036854775808)]
                            var v33 float32 = float32(1e10)
                            var v34 int16 = -32768
                            switch ((v23 + v23) + v23) {
                            case "foo bar":
                                var v35 uint16 = 898
                                _ = v35
                            case "\n\t":
                                var v36 float32 = v33
                                _ = v36
                            case "a":
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    var v37 int8 = -1
                                    _ = v37
                                }
                            default:
                                v32 = v32
                            }
                            v23 = ((((v23 + (v23 + "\n\t")) + v23) + v23) + "foo bar")
                            if true {
                                var v38 int32 = 180
                                fmt.Println(v23, v22)
                                var v39 uint64 = 1
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    v1[(-(fi27))] = v32
                                case "":
                                    if (!(v22 && false)) {
                                        switch 127 {
                                        case 127:
                                            v40 := GenS2{F0: (!(fi24 != fi24)), F1: v34, F2: 1}
                                            _ = v40
                                        }
                                        v41 := GenS4{F0: false, F1: v32}
                                        _ = v41
                                    } else {
                                        _ = len(v30)
                                        var v42 int8 = -128
                                        if v22 {
                                            var v43 uint8 = 148
                                            _ = v43
                                        } else {
                                            v30 = append(v30, (fi24 / (((fi24 >> ((2) & 63))) | 1)))
                                            var v44 int32 = (^(v38))
                                            if v22 {
                                                v42 -= (-(v42))
                                            } else {
                                                v39 = v39
                                                v45 := GenS6{F0: 0, F1: 4294967295}
                                                _ = len(v30)
                                                _ = v45
                                            }
                                            _ = v44
                                        }
                                        _ = v42
                                    }
                                }
                                switch "foo bar" {
                                case "hello":
                                    switch "" {
                                    case "foo bar":
                                        v2[(-(fi24))] = v22
                                    default:
                                        v46 := GenS5{F0: v22, F1: v23}
                                        _ = v46
                                    }
                                case "":
                                    _ = len(v3)
                                }
                                _ = v38
                                _ = v39
                            } else {
                                v30 = append(v30, (^(fi27)))
                                v32 += (v32 % (((v32 | (v32 + (v32 >> ((42) & 63))))) | 1))
                            }
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        _ = fi27
                        _ = v30
                    }
                case "a":
                    var v47 int = fi24
                    _ = v47
                default:
                    v48 := GenS2{F0: v22, F1: 1000, F2: -128}
                    _ = v48
                }
                _ = v25
            } else {
                for fi49 := 0; fi49 < 14; fi49++ {
                    v50 := GenS7{F0: 27}
                    v51 := GenS8{F0: (((float32(1e10) * ((0.1 * ((((((0.0 + (float32(float64(3.14159)) + 0.1)) / 1.0) / 3.14159) - (float32((((1.5 + ((float64(float32(0.0)) - 1e10) / 0.1)) * float64(1.5)) + ((-2.25 * float64((-2.25 + 1e10))) * 1.5))) / (0.0 - 1.5))) + 1e10) / -2.25)) - 1e10)) * 1.5) == (1e10 / -2.25))}
                    v52 := []int{(-(fi49)), fi49}
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
                for fi53 := 0; fi53 < 7; fi53++ {
                    var v54 uint32 = 1
                    v2[fi53] = v22
                    if (v22 && (!true)) {
                        fmt.Println(v23)
                    } else {
                        if false {
                            v55 := GenS7{F0: 77}
                            v22 = (177 <= (^(v54)))
                            v56 := []int{402, -9223372036854775808}
                            _ = v55
                            _ = v56
                        }
                        v57 := GenS7{F0: 1}
                        switch v23 {
                        case "hello":
                            var v58 bool = (!v22)
                            _ = v58
                        case "foo bar":
                            if ((!((3.14159 / (float64(3.14159) - 1e10)) >= 0.0)) || (42 <= -1)) {
                                v59 := GenS7{F0: 42}
                                _ = v59
                            }
                        }
                        var v60 uint16 = 43
                        v61 := GenS9{F0: 42}
                        v23 = v23
                        _ = v57
                        _ = v60
                        _ = v61
                    }
                    _ = fi53
                    _ = v54
                }
            }
            _ = len(v3)
            _ = len(v1)
            _ = fi24
        }
        _ = v23
    }
    if v22 {
        fmt.Println(v22)
        v62 := []int{42, 301, 501, -9223372036854775808}
        fmt.Println(v22)
        _ = v62
    } else {
        var v63 int8 = 0
        _ = v63
    }
    v22 = v22
    var v64 int16 = 560
    if v22 {
        var v65 int64 = -9223372036854775808
        fmt.Println(v64, v64)
        _ = v65
    } else {
        fmt.Println(v64, v64)
        v22 = true
    }
    v66 := GenS7{F0: 2}
    if v22 {
        if (-1 > genFn0(-305, false, 42)) {
            switch "hello" {
            case "\n\t":
                v22 = ((1e10 == 0.1) || ((!v22) && v22))
            case "":
                v64 -= (v64 ^ v64)
            default:
                var v67 int16 = (v64 + v64)
                _ = v67
            }
            v64 = (v64 * 157)
            var v68 uint16 = 557
            _ = v68
        }
    } else {
        var v69 uint = 0
        _ = v69
    }
    v2[4294967295] = ((1.5 * 1e10) < 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v64
    _ = v66
}
