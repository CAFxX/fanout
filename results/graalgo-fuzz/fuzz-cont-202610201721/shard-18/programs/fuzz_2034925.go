package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 string }
func genFn0(p0 int16, p1 uint32, p2 uint) uint64 {
    var v7 int64 = -1
    v8 := GenS1{F0: 0, F1: ("foo bar" + (("\n\t" + "\n\t") + ("foo bar" + (("foo bar" + "hello") + "a"))))}
    var v9 int32 = -2147483648
    v10 := GenS2{F0: 0}
    for fi11 := 0; fi11 < 14; fi11++ {
        if true {
            v12 := []int{2147483647, fi11}
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 float64 = (((0.0 + (1.5 / 3.14159)) / (((0.0 - ((1.5 - 1e10) / 1.5)) - (3.14159 + -2.25)) - 1.5)) * ((1e10 + 0.1) / (0.1 + float64(1e10))))
                if false {
                    fmt.Println(fi11, p1, fi11)
                    fmt.Println(fi11, fi11)
                    var v15 float64 = v14
                    _ = v15
                } else {
                    v16 := GenS3{F0: uint64(p1), F1: p2}
                    v14 -= (float64((0.0 - (0.1 + 0.1))) * float64(0.1))
                    fmt.Println(p0)
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            if true {
                var v17 int8 = -1
                var v18 int = (fi11 + 42)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v17
                _ = v18
            } else {
                v12 = append(v12, fi11)
                var v19 uint = (p2 & (p2 & (p2 | p2)))
                v12 = append(v12, -2147483648)
                _ = v19
            }
            _ = v12
        } else {
            v7 += (v7 % ((912) | 1))
            var v20 uint64 = 42
            _ = v20
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 string) uint64 {
    fmt.Println(p0)
    p0 = (p0 + ("" + p0))
    _ = p0
    return 294
}

func main() {
    v1 := []uint32{853, 24, 65535, 425}
    v2 := make(map[uint32]bool)
    v3 := []string{"a", "foo bar"}
    var v4 int64 = -9223372036854775808
    var v5 uint64 = 973
    v6 := make(map[int]int)
    switch -492 {
    case 127:
        var v21 bool = true
        _ = v21
    }
    v22 := GenS2{F0: 53}
    if true {
        v4 = -883
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(v5, v5, fi23)
            _ = fi23
        }
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v4)
            switch "hello" {
            case "foo bar":
                if true {
                    if (2147483647 != -891) {
                        var v25 uint = 18446744073709551615
                        v4 += (v4 + (v4 + -9223372036854775808))
                        v6[(fi24 >> (((v25 / (((v25 * uint(v5))) | 1))) & 63))] = (fi24 + fi24)
                        _ = v25
                    }
                }
            case "":
                v6[131] = fi24
            case "\n\t":
                var v26 string = "foo bar"
                _ = v26
            default:
                if (0 < 255) {
                    if true {
                        _ = len(v3)
                        v2[42] = (533 >= 0)
                        v5 = v5
                    }
                    if false {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            var v28 int64 = (v4 >> ((4294967295) & 63))
                            _ = fi27
                            _ = v28
                        }
                        switch (fi24 | (fi24 / (((fi24 << ((255) & 63))) | 1))) {
                        case 1:
                            v4 -= (v4 - (v4 ^ 1))
                        case -1:
                            var v29 uint8 = 127
                            _ = v29
                        case 329:
                            switch fi24 {
                            case 127:
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    var v31 float32 = -2.25
                                    var v32 int64 = 99
                                    v32 = (^(v32))
                                    var v33 float32 = float32(0.1)
                                    var v34 int8 = 127
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            case -1:
                                var v35 uint64 = (v5 % ((923) | 1))
                                _ = v35
                            case 4294967295:
                                switch "\n\t" {
                                case "\n\t":
                                    v5 -= (^(v5))
                                case "hello":
                                    var v36 uint = 42
                                    _ = v36
                                case "":
                                    v3 = append(v3, "\n\t")
                                default:
                                    switch (fi24 * fi24) {
                                    case 751:
                                        switch 993 {
                                        case 1:
                                            v5 -= v5
                                        case -2147483648:
                                            var v37 uint64 = (v5 | v5)
                                            _ = v37
                                        }
                                    }
                                }
                            }
                        default:
                            var v38 int16 = 705
                            _ = v38
                        }
                        switch fi24 {
                        case 127:
                            v5 += (v5 >> ((42) & 63))
                        case -9223372036854775808:
                            v39 := GenS1{F0: 42, F1: "\n\t"}
                            _ = v39
                        case 2147483647:
                            var v40 int8 = 0
                            _ = v40
                        }
                    }
                } else {
                    var v41 uint64 = (v5 + v5)
                    _ = v41
                }
            }
            var v42 int16 = 32767
            _ = fi24
            _ = v42
        }
        for fi43 := 0; fi43 < 1; fi43++ {
            switch "foo bar" {
            case "\n\t":
                for fi44 := 0; fi44 < 15; fi44++ {
                    v45 := GenS2{F0: 145}
                    switch 2147483647 {
                    case -633:
                        var v46 bool = ((((0.1 / 1.0) / ((((1.5 + ((-2.25 - 1.5) - (((3.14159 - float32((0.0 + 1.5))) * ((3.14159 / 1.5) / (-2.25 / (-2.25 * (float32(0.1) / 3.14159))))) * -2.25))) * -2.25) - (0.1 - 3.14159)) - 0.0)) >= (float32(((float64(((0.0 - 1e10) - ((0.0 / (float32((3.14159 / (float64(-2.25) / 1.0))) - float32(((3.14159 / 1.0) / ((-2.25 - 3.14159) + -2.25))))) + 0.1))) - 3.14159) + (((1.5 * 0.1) * 1.5) - -2.25))) + 0.1)) || (!false))
                        _ = v46
                    case 32767:
                        if (true && false) {
                            var v47 int = 9223372036854775807
                            var v48 int8 = -128
                            _ = v47
                            _ = v48
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            v5 = (^(v5))
                        case "\n\t":
                            v49 := GenS1{F0: 1, F1: ""}
                            _ = v49
                        case "":
                            var v50 float32 = float32(float64(0.0))
                            _ = v50
                        }
                    }
                    v4 = (^(v4))
                    v51 := GenS1{F0: 255, F1: "hello"}
                    v52 := GenS1{F0: 2, F1: "a"}
                    _ = fi44
                    _ = v45
                    _ = v51
                    _ = v52
                }
            case "a":
                fmt.Println(fi43, v4, fi43)
            default:
                v5 -= genFn0(-142, 972, 65535)
            }
            for fi53 := 0; fi53 < 3; fi53++ {
                switch (("" + "") + "foo bar") {
                case "":
                    v5 -= (v5 >> ((729) & 63))
                }
                _ = fi53
            }
            var v54 int8 = 0
            _ = fi43
            _ = v54
        }
        for fi55 := 0; fi55 < 3; fi55++ {
            fmt.Println(v5, v4)
            for fi56 := 0; fi56 < 18; fi56++ {
                v57 := GenS4{F0: (("hello" + "\n\t") + "hello"), F1: 49}
                var v58 float64 = float64(1e10)
                fmt.Println(fi55, v58, fi56)
                _ = fi56
                _ = v57
                _ = v58
            }
            _ = fi55
        }
    }
    v5 = v5
    var v59 int16 = 42
    v4 = 684
    var v60 uint16 = 2
    switch -9223372036854775808 {
    case 32768:
        var v61 int8 = 111
        _ = v61
    case 0:
        v62 := GenS5{F0: ("a" + ("\n\t" + ("" + "a")))}
        _ = v62
    case 9223372036854775807:
        v63 := GenS3{F0: genFn0(v59, 0, 529), F1: 171}
        _ = v63
    }
    switch "a" {
    case "\n\t":
        v4 -= 42
    case "foo bar":
        switch ("\n\t" + "foo bar") {
        case "a":
            var v64 uint32 = 42
            _ = v64
        default:
            if false {
                v65 := GenS5{F0: "\n\t"}
                _ = v65
            }
        }
    case "hello":
        v66 := v6[0]
        _ = v66
    }
    var v67 float32 = 0.0
    for fi68 := 0; fi68 < 8; fi68++ {
        var v69 string = ("\n\t" + "")
        _ = fi68
        _ = v69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v67
}
