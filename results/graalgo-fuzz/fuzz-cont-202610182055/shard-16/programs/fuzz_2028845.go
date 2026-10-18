package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int }
func genFn0(p0 string) uint64 {
    switch p0 {
    case "a":
        fmt.Println(p0)
    case "hello":
        p0 = ((p0 + "foo bar") + (((p0 + p0) + "a") + ("\n\t" + ("" + (p0 + "hello")))))
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = p0
        var v8 uint8 = 92
        _ = fi7
        _ = v8
    }
    if false {
        if (!false) {
            v9 := []int{42, -300, -1}
            switch "a" {
            case "":
                fmt.Println(p0)
            case "a":
                if (18446744073709551615 < 42) {
                    v9 = append(v9, -822)
                }
            case "\n\t":
                switch p0 {
                case "a":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case "":
                    p0 = "a"
                }
            }
            var v10 int = 9223372036854775807
            _ = v10
            _ = v9
        }
        if false {
            var v11 uint64 = 42
            _ = v11
        } else {
            if true {
                var v12 uint64 = 42
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 bool = true
                    _ = fi13
                    _ = v14
                }
                fmt.Println(p0, v12)
                _ = v12
            }
            var v15 uint = 0
            switch -2147483648 {
            case -589:
                v16 := []int{127, 128}
                _ = v16
            }
            _ = v15
        }
        if true {
            var v17 int16 = 0
            v18 := GenS1{F0: 307, F1: 255}
            _ = v17
            _ = v18
        }
    }
    v19 := []int{4294967295, -1}
    var v20 bool = true
    _ = p0
    _ = v19
    _ = v20
    return 925
}
func genFn1(p0 uint, p1 int64) int16 {
    p1 += p1
    var v21 bool = (genFn0(("a" + "a")) == 4294967295)
    var v22 float32 = ((3.14159 / 1e10) + 0.0)
    v23 := GenS2{F0: 42, F1: 4294967295}
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := []int{fi24, (fi24 >> (((p0 / ((p0) | 1))) & 63))}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 bool = (127 <= -59)
    var v2 float32 = (((0.0 * 1.5) - (((1.5 * float32(3.14159)) + -2.25) / 0.1)) / (1e10 + (-2.25 - -2.25)))
    v3 := []int{733, 127}
    v4 := []uint32{0, 2147483648}
    v5 := []string{"", "", "hello", "foo bar"}
    v6 := make(map[int]string)
    if (559 < 18446744073709551615) {
        fmt.Println(v2, v1)
        v26 := GenS1{F0: -1, F1: 2}
        var v27 float32 = v2
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            v1 = v1
            if v1 {
                if (2 == 0) {
                    v2 -= ((v2 - -2.25) / v2)
                    fmt.Println(fi28)
                }
            }
            _ = fi29
        }
        var v30 uint = 255
        if true {
            if ((fi28 | (fi28 * fi28)) <= fi28) {
                if (v1 || v1) {
                    if (-1 != -1) {
                        if (((-2.25 + v2) * 3.14159) != 1.5) {
                            v31 := GenS3{F0: 1, F1: 84, F2: 127}
                            var v32 uint = (v30 % ((42) | 1))
                            _ = v31
                            _ = v32
                        } else {
                            _ = len(v6)
                            v30 -= 18446744073709551615
                            var v33 int64 = 783
                            _ = v33
                        }
                        v1 = (1 > 91)
                        var v34 uint64 = 2
                        _ = v34
                    } else {
                        v2 = v2
                    }
                }
            } else {
                if v1 {
                    var v35 uint = 42
                    _ = v35
                } else {
                    v36 := GenS4{F0: 13}
                    for fi37 := 0; fi37 < 17; fi37++ {
                        if (255 != v30) {
                            v2 = (v2 * v2)
                        } else {
                            v6[699] = "foo bar"
                            if false {
                                if v1 {
                                    if true {
                                        v38 := GenS5{F0: 42, F1: -2147483648, F2: "a"}
                                        _ = v38
                                    }
                                } else {
                                    v39 := GenS4{F0: -128}
                                    var v40 int = fi28
                                    fmt.Println(v1)
                                    _ = v39
                                    _ = v40
                                }
                                fmt.Println(v1, v1, v1)
                            }
                            v1 = (392 <= -2147483648)
                        }
                        _ = fi37
                    }
                    var v41 string = "foo bar"
                    _ = v36
                    _ = v41
                }
                var v42 bool = (-313 <= 1000)
                _ = v42
            }
        } else {
            fmt.Println(v30)
        }
        _ = fi28
        _ = v30
    }
    var v43 bool = v1
    switch 42 {
    case 9223372036854775807:
        v44 := v6[-785]
        _ = v44
    case -2147483648:
        var v45 int = 9
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        v2 += v2
        var v47 uint16 = 2
        switch "" {
        case "\n\t":
            switch -734 {
            case -781:
                switch "a" {
                case "a":
                    fmt.Println(v2)
                default:
                    var v48 bool = (0 > 0)
                    _ = v48
                }
            }
        case "a":
            fmt.Println(fi46)
        default:
            if (!((true && v1) || v43)) {
                var v49 int8 = -33
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v49
            } else {
                var v50 float64 = 0.0
                switch "" {
                case "":
                    if v43 {
                        if v43 {
                            var v51 float64 = (float64(((-2.25 / (v2 * v2)) + (v2 * v2))) / ((v50 + ((v50 - 1e10) / float64(v2))) * 1e10))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v43 = v43
                            _ = v51
                        } else {
                            v43 = false
                            v4 = append(v4, 1)
                        }
                        v43 = v1
                        var v52 int = 4294967295
                        _ = v52
                    } else {
                        switch "a" {
                        case "foo bar":
                            var v53 int8 = 127
                            _ = v53
                        }
                    }
                default:
                    for fi54 := 0; fi54 < 4; fi54++ {
                        v47 += v47
                        fmt.Println(fi54, v43)
                        v3 = append(v3, (-(fi54)))
                        _ = fi54
                    }
                }
                _ = v50
            }
        }
        _ = fi46
        _ = v47
    }
    v2 += v2
    switch "" {
    case "":
        if v1 {
            switch -1 {
            case -1:
                if v1 {
                    var v55 uint8 = 128
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v1 = v1
                        v2 = float32(float64((v2 + v2)))
                        _ = fi56
                    }
                    _ = v55
                } else {
                    var v57 uint = 255
                    for fi58 := 0; fi58 < 16; fi58++ {
                        v57 = 914
                        var v59 uint8 = 255
                        _ = fi58
                        _ = v59
                    }
                    _ = v57
                }
            case -385:
                if (!true) {
                    var v60 uint = 4294967295
                    v1 = v1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v60
                }
            default:
                if (9223372036854775807 == -9223372036854775808) {
                    if v43 {
                        v61 := GenS6{F0: "a", F1: 2}
                        _ = len(v5)
                        for fi62 := 0; fi62 < 11; fi62++ {
                            if len(v3) > 0 {
                                v3[0] = (fi62 ^ fi62)
                            }
                            switch (^(fi62)) {
                            case 538:
                                for fi63 := 0; fi63 < 10; fi63++ {
                                    v2 = 1.5
                                    var v64 uint64 = 0
                                    v43 = (false || v1)
                                    _ = fi63
                                    _ = v64
                                }
                            case -422:
                                v65 := v6[fi62]
                                _ = v65
                            }
                            var v66 int = fi62
                            _ = fi62
                            _ = v66
                        }
                        _ = v61
                    }
                    v6[127] = "\n\t"
                } else {
                    v43 = v43
                    switch ("a" + "a") {
                    case "hello":
                        v1 = (9223372036854775807 < 42)
                    default:
                        v67 := GenS4{F0: -128}
                        _ = v67
                    }
                }
            }
            fmt.Println(v43, v43, v2)
        }
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = ""
        }
    }
    v68 := GenS7{F0: 137, F1: 24, F2: 127}
    for fi69 := 0; fi69 < 17; fi69++ {
        switch "\n\t" {
        case "\n\t":
            v70 := GenS2{F0: 42, F1: 4294967295}
            _ = v70
        case "a":
            delete(v6, (fi69 >> ((65535) & 63)))
        case "foo bar":
            v2 = (v2 / v2)
        }
        _ = fi69
    }
    var v71 uint32 = 42
    var v72 uint8 = 127
    switch "a" {
    case "\n\t":
        for fi73 := 0; fi73 < 11; fi73++ {
            v1 = v1
            fmt.Println(v72)
            _ = fi73
        }
    case "a":
        var v74 int32 = 1
        _ = v74
    case "hello":
        var v75 uint32 = uint32(v72)
        _ = v75
    }
    var v76 int = -614
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v68
    _ = v71
    _ = v72
    _ = v76
}
