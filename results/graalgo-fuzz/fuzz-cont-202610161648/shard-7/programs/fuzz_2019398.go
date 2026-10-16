package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int64; F2 int32 }
type GenS5 struct { F0 uint8; F1 bool; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int8; F1 int8; F2 string }
func genFn0(p0 uint, p1 uint, p2 int64) int8 {
    v4 := GenS1{F0: 42, F1: 42}
    for fi5 := 0; fi5 < 3; fi5++ {
        if false {
            v6 := []int{-387, 32768, (fi5 - fi5), (fi5 & fi5)}
            v7 := GenS1{F0: 251, F1: 0}
            if len(v6) > 0 {
                v6[0] = (fi5 << ((p1) & 63))
            }
            _ = v6
            _ = v7
        }
        v8 := []int{32767}
        _ = fi5
        _ = v8
    }
    var v9 float32 = 0.1
    p2 += p2
    for fi10 := 0; fi10 < 11; fi10++ {
        p1 += (p1 & (p1 | (p0 | p1)))
        var v11 uint8 = 2
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return -77
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{737, -388, 32768}
    v3 := make(map[int]uint32)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch (^(fi13)) {
            case 42:
                if true {
                    switch "" {
                    case "\n\t":
                        v14 := GenS1{F0: 282, F1: -2147483648}
                        v15 := GenS2{F0: "\n\t", F1: 1, F2: 481}
                        _ = v14
                        _ = v15
                    case "hello":
                        switch ("" + "foo bar") {
                        case "":
                            var v16 string = "\n\t"
                            _ = v16
                        case "foo bar":
                            var v17 int = 731
                            var v18 uint = 42
                            _ = v17
                            _ = v18
                        default:
                            var v19 uint64 = 18446744073709551615
                            _ = v19
                        }
                    }
                } else {
                    if true {
                        var v20 string = ""
                        v20 = "a"
                        _ = v20
                    } else {
                        if (true || false) {
                            var v21 int64 = 722
                            fmt.Println(fi13, v21, v21)
                            switch "" {
                            case "hello":
                                v22 := GenS3{F0: -1}
                                _ = v22
                            case "foo bar":
                                var v23 bool = ((fi12 >> ((78) & 63)) >= fi12)
                                _ = v23
                            }
                            _ = v21
                        } else {
                            var v24 string = "hello"
                            if (v24 > (((("a" + (v24 + (v24 + v24))) + v24) + v24) + "hello")) {
                                v24 = v24
                                var v25 float64 = 1.5
                                _ = v25
                            }
                            _ = v24
                        }
                        switch "a" {
                        case "a":
                            switch ("" + ("a" + "foo bar")) {
                            case "":
                                v26 := GenS2{F0: ((("\n\t" + ("" + ("foo bar" + "\n\t"))) + "a") + ""), F1: 255, F2: 42}
                                _ = v26
                            }
                        case "\n\t":
                            fmt.Println(fi13, fi13)
                        default:
                            var v27 float32 = 0.1
                            _ = v27
                        }
                        var v28 uint64 = 18446744073709551615
                        _ = v28
                    }
                }
            case -1:
                var v29 int64 = 9223372036854775807
                _ = v29
            case 4294967295:
                if true {
                    var v30 int16 = 1000
                    fmt.Println(fi12)
                    fmt.Println(fi13)
                    v30 -= (-(v30))
                    _ = v30
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        var v31 string = "hello"
                        _ = v31
                    case "":
                        var v32 bool = (true || ((!((1 < 127) && (false || (0.1 <= 0.1)))) && (-49 >= -32768)))
                        _ = v32
                    case "a":
                        var v33 int = (fi12 & 32768)
                        _ = v33
                    default:
                        var v34 uint16 = 42
                        _ = v34
                    }
                    var v35 uint64 = 0
                    _ = v35
                }
            }
            v36 := GenS1{F0: 322, F1: -472}
            _ = fi13
            _ = v36
        }
        _ = fi12
    }
    v37 := GenS4{F0: 0, F1: 1, F2: -1}
    v38 := GenS5{F0: 166, F1: false, F2: false}
    var v39 float64 = 1.5
    var v40 float32 = (1.5 + float32(((((0.0 - float64(1.5)) * v39) * v39) - 1.5)))
    var v41 int = 127
    v40 = v40
    v42 := GenS6{F0: 0}
    v39 += (((((3.14159 + (v39 + v39)) + v39) * -2.25) / ((v39 * ((v39 + ((((v39 / v39) * ((0.0 * float64(v40)) * ((float64(v40) - (v39 - v39)) + float64(((v40 / 3.14159) * ((1.5 / v40) + 1.5)))))) - (0.1 - (-2.25 / 1e10))) + v39)) / v39)) + float64(v40))) / v39)
    v2 = append(v2, (v41 << ((255) & 63)))
    fmt.Println(v41, v40)
    if (0 >= 0) {
        _ = len(v3)
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "a":
            switch v41 {
            case -218:
                var v43 string = "foo bar"
                _ = v43
            default:
                v44 := v3[v41]
                _ = v44
            }
        case "":
            var v45 int32 = -282
            _ = v45
        case "\n\t":
            switch "a" {
            case "a":
                var v46 uint = 2
                _ = v46
            case "hello":
                v47 := GenS3{F0: v41}
                _ = v47
            case "":
                if true {
                    v48 := GenS4{F0: 79, F1: 538, F2: 1}
                    v40 += float32(v39)
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        if false {
                            fmt.Println(fi49, v41)
                            if (255 <= 32767) {
                                _ = len(v1)
                                fmt.Println(v39)
                            }
                            if false {
                                v50 := GenS7{F0: 18446744073709551615, F1: -405}
                                if (255 < 183) {
                                    var v51 uint8 = 0
                                    v40 = (float32(0.0) - v40)
                                    _ = v51
                                } else {
                                    if (0 < 255) {
                                        v41 = 128
                                    } else {
                                        var v52 bool = false
                                        v53 := GenS8{F0: 127, F1: 0, F2: "foo bar"}
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 uint16 = 32767
                                    _ = v54
                                }
                                _ = v50
                            } else {
                                v55 := GenS6{F0: 0}
                                _ = v55
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (v41 | (v41 << ((1) & 63)))
                        }
                        _ = fi49
                    }
                    v56 := GenS3{F0: (v41 * -214)}
                    _ = v56
                }
            default:
                if (false || (42 == 794)) {
                    for fi57 := 0; fi57 < 16; fi57++ {
                        v39 = (v39 - float64(-2.25))
                        _ = fi57
                    }
                    _ = len(v2)
                }
            }
        }
        var v58 bool = true
        _ = v58
    }
    v40 = (3.14159 - ((((v40 / v40) - v40) * v40) / (v40 * -2.25)))
    if (128 <= 42) {
        for fi59 := 0; fi59 < 1; fi59++ {
            v39 -= (float64(-2.25) * v39)
            v41 -= -731
            var v60 int8 = -1
            _ = fi59
            _ = v60
        }
        var v61 int = 4294967295
        _ = v61
    } else {
        v62 := GenS2{F0: ((((("foo bar" + ("hello" + "\n\t")) + "\n\t") + "\n\t") + ("a" + "hello")) + ""), F1: 2, F2: 18446744073709551615}
        for fi63 := 0; fi63 < 12; fi63++ {
            if len(v2) > 0 {
                v2[0] = (v41 & fi63)
            }
            _ = fi63
        }
        if (255 <= 255) {
            _ = len(v2)
            v39 -= v39
        }
        _ = v62
    }
    _ = len(v2)
    fmt.Println(v39, v41, v39)
    var v64 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v64
}
