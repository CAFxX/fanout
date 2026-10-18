package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int64; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 string, p1 int64) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(fi7)
                var v9 int8 = -19
                fmt.Println(fi8, fi7, p1)
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 int64 = p1
                fmt.Println(v11, p1, fi10)
                v12 := GenS1{F0: -128}
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(fi13, p0, fi13)
                _ = fi13
            }
        }
        v14 := GenS1{F0: 42}
        if true {
            v15 := GenS2{F0: 128, F1: 127}
            _ = v15
        } else {
            var v16 uint8 = 97
            v17 := GenS3{F0: "hello"}
            var v18 float64 = float64((3.14159 + -2.25))
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi7
        _ = v14
    }
    v19 := []int{32768}
    fmt.Println(p1, p0)
    v20 := GenS1{F0: -1}
    var v21 string = ""
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 599
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 float64 = 0.1
    v6 := []int{-1}
    v4["\n\t"] = -1
    v5 = (v5 * (1.5 / v5))
    if (("\n\t" + ("\n\t" + ((("a" + "hello") + ((("\n\t" + "foo bar") + "") + "foo bar")) + "\n\t"))) <= ((("hello" + ("" + ("" + ("hello" + ("" + "foo bar"))))) + (("\n\t" + "a") + "hello")) + "a")) {
        switch "hello" {
        case "a":
            var v22 uint64 = 262
            _ = v22
        case "foo bar":
            var v23 int64 = 9223372036854775807
            _ = v23
        default:
            var v24 int = 1
            _ = v24
        }
    } else {
        switch "hello" {
        case "\n\t":
            var v25 int16 = 32767
            _ = v25
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = -1
            }
        case "a":
            var v26 int16 = -32768
            _ = v26
        default:
            if false {
                _ = len(v3)
                switch ("foo bar" + ((("" + "") + ("\n\t" + ("foo bar" + ("a" + "a")))) + "foo bar")) {
                case "\n\t":
                    v27 := GenS4{F0: ("\n\t" + "\n\t"), F1: -9223372036854775808, F2: 124}
                    _ = v27
                case "hello":
                    v28 := GenS1{F0: 10}
                    _ = v28
                case "foo bar":
                    switch -1 {
                    case -446:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            v5 = float64(1.5)
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                            _ = fi29
                        }
                    case -734:
                        v5 += ((0.0 * 1.5) * v5)
                    case 42:
                        var v30 uint32 = 1
                        _ = v30
                    default:
                        v2 = v2
                    }
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    if v2 {
                        v32 := GenS1{F0: 127}
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            fmt.Println(v5, v2, v5)
                            var v34 bool = (-914 < 0)
                            v2 = v2
                            _ = fi33
                            _ = v34
                        }
                    }
                    var v35 uint16 = 0
                    v36 := GenS5{F0: -208}
                    _ = fi31
                    _ = v35
                    _ = v36
                }
            } else {
                if (!(2 > 1)) {
                    v37 := GenS3{F0: ("hello" + "\n\t")}
                    if (true && (genFn0("a", 42) < genFn0("hello", -9223372036854775808))) {
                        _ = len(v6)
                        for fi38 := 0; fi38 < 1; fi38++ {
                            var v39 int64 = 0
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        var v40 int16 = 1
                        switch "a" {
                        case "a":
                            for fi41 := 0; fi41 < 4; fi41++ {
                                fmt.Println(v5)
                                v40 = (v40 + (v40 << ((4294967295) & 63)))
                                if (!v2) {
                                    _ = len(v6)
                                    v42 := GenS5{F0: fi41}
                                    var v43 int32 = -1
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi41
                            }
                        case "hello":
                            if (-977 > 1) {
                                if (0 >= -990) {
                                    if v2 {
                                        switch 1 {
                                        case 0:
                                            var v44 uint = 956
                                            _ = v44
                                        case 669:
                                            if ("foo bar" != ("" + "\n\t")) {
                                                v45 := GenS4{F0: "hello", F1: 9223372036854775807, F2: 255}
                                                v46 := GenS2{F0: 255, F1: 91}
                                                _ = v45
                                                _ = v46
                                            } else {
                                                v47 := GenS6{F0: "hello", F1: false}
                                                _ = v47
                                            }
                                        }
                                        var v48 bool = ((v40 | 1) <= (v40 % ((1000) | 1)))
                                        v48 = (1 <= 42)
                                        _ = v48
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v49 := GenS2{F0: 255, F1: 231}
                                        _ = v49
                                    default:
                                        if ((-1 >= 42) && ((("\n\t" + "") + ("a" + "foo bar")) > "\n\t")) {
                                            if v2 {
                                                var v50 uint = 0
                                                _ = v50
                                            } else {
                                                _ = len(v4)
                                                var v51 bool = true
                                                for fi52 := 0; fi52 < 8; fi52++ {
                                                    v53 := GenS7{F0: fi52}
                                                    var v54 uint64 = 2
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v51
                                            }
                                        }
                                    }
                                } else {
                                    if false {
                                        var v55 int16 = (v40 << ((42) & 63))
                                        for fi56 := 0; fi56 < 2; fi56++ {
                                            v2 = (true || (!true))
                                            _ = fi56
                                        }
                                        v57 := v1[-427]
                                        _ = v55
                                        _ = v57
                                    } else {
                                        if (-469 != 0) {
                                            if (float64((0.0 - (1.5 * (((1.5 + 0.0) / (1.5 / 1e10)) - 1e10)))) <= v5) {
                                                var v58 bool = false
                                                _ = v58
                                            }
                                            v2 = (711 < 154)
                                        }
                                    }
                                    var v59 int8 = 127
                                    _ = v59
                                }
                                v5 = v5
                                v2 = v2
                            }
                        case "":
                            v6 = append(v6, 9223372036854775807)
                        default:
                            _ = len(v6)
                        }
                        v1[4294967295] = 2147483647
                        _ = v40
                    }
                    fmt.Println(v5)
                    _ = v37
                } else {
                    for fi60 := 0; fi60 < 1; fi60++ {
                        var v61 int8 = -5
                        var v62 float64 = 1e10
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                }
                var v63 uint32 = 4294967295
                _ = v63
            }
        }
        var v64 int16 = 1
        var v65 uint32 = 0
        _ = v64
        _ = v65
    }
    var v66 int16 = 912
    var v67 int32 = 0
    v2 = (0 > 2)
    switch ("hello" + "hello") {
    case "hello":
        v5 = (0.1 * v5)
    default:
        v68 := GenS8{F0: 128, F1: int8(v67), F2: "foo bar"}
        _ = v68
    }
    var v69 uint64 = 0
    v2 = (0 <= 1)
    var v70 bool = ((v66 << ((953) & 63)) != 1)
    v69 = 18446744073709551615
    v6 = append(v6, 4294967295)
    v6 = append(v6, -906)
    v5 = 1.5
    for fi71 := 0; fi71 < 13; fi71++ {
        v69 -= v69
        v5 = (1e10 / -2.25)
        _ = fi71
    }
    var v72 uint32 = 292
    fmt.Println(v66, v72)
    switch "hello" {
    case "hello":
        delete(v4, "\n\t")
    case "foo bar":
        v69 -= v69
    case "":
        if false {
            if len(v6) > 0 {
                v6[0] = 127
            }
            if v2 {
                v69 = v69
            } else {
                var v73 uint16 = 543
                fmt.Println(v66, v66)
                fmt.Println(v67, v69, v70)
                _ = v73
            }
            v74 := GenS8{F0: 255, F1: -128, F2: "foo bar"}
            _ = v74
        } else {
            for fi75 := 0; fi75 < 2; fi75++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi75
            }
            for fi76 := 0; fi76 < 17; fi76++ {
                var v77 bool = (42 >= 2)
                fmt.Println(v77, fi76, v69)
                var v78 int64 = -1
                _ = fi76
                _ = v77
                _ = v78
            }
            v79 := GenS6{F0: "foo bar", F1: (!((^(v72)) >= 2))}
            _ = v79
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v69
    _ = v70
    _ = v72
}
