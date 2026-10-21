package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 string }
func genFn0(p0 int64, p1 uint32) int8 {
    if (127 >= 1) {
        var v6 uint32 = p1
        v6 = (v6 & v6)
        _ = v6
    }
    v7 := []int{0, 966}
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := []uint32{}
    v4 := []string{"foo bar", "foo bar", (("\n\t" + "hello") + "hello")}
    v5 := []uint32{341, 110}
    if true {
        var v8 int16 = 788
        var v9 float64 = -2.25
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 0
    var v12 int64 = 9223372036854775807
    v3 = append(v3, 2147483648)
    if len(v5) > 0 {
        v5[0] = 2
    }
    v13 := v1[""]
    switch "hello" {
    case "\n\t":
        v12 = v12
    case "hello":
        for fi14 := 0; fi14 < 16; fi14++ {
            if v13 {
                fmt.Println(fi14)
            } else {
                var v15 int = -9223372036854775808
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                switch "\n\t" {
                case "a":
                    switch fi14 {
                    case 127:
                        if v13 {
                            var v17 float64 = -2.25
                            _ = v17
                        }
                    case 128:
                        var v18 uint = 1
                        _ = v18
                    case 32767:
                        var v19 uint64 = 42
                        _ = v19
                    default:
                        if len(v4) > 0 {
                            v4[0] = ("foo bar" + "hello")
                        }
                    }
                    fmt.Println(fi14, v13, v13)
                    switch fi16 {
                    case 32767:
                        if true {
                            fmt.Println(v12, v11)
                        } else {
                            switch (fi14 & fi14) {
                            case 1:
                                switch "\n\t" {
                                case "hello":
                                    switch "" {
                                    case "":
                                        var v20 int16 = -32768
                                        _ = v20
                                    default:
                                        v13 = (true || ((^(v11)) < 1))
                                    }
                                default:
                                    v21 := GenS1{F0: 65535, F1: 2147483647, F2: 32767}
                                    _ = v21
                                }
                            case -600:
                                v11 -= 137
                            default:
                                fmt.Println(fi14, fi14, v12)
                            }
                            v22 := GenS1{F0: 843, F1: 42, F2: -234}
                            _ = v22
                        }
                    }
                    switch fi14 {
                    case -842:
                        switch (("\n\t" + "foo bar") + (("" + "hello") + "\n\t")) {
                        case "foo bar":
                            if true {
                                if false {
                                    v23 := GenS1{F0: 42, F1: 42, F2: 1000}
                                    v3 = append(v3, 179)
                                    v11 = v11
                                    _ = v23
                                } else {
                                    v24 := v1["\n\t"]
                                    _ = v24
                                }
                                switch "a" {
                                case "hello":
                                    v25 := GenS1{F0: 65535, F1: 0, F2: 732}
                                    _ = v25
                                case "":
                                    if false {
                                        var v26 uint = 42
                                        v27 := v1[""]
                                        v27 = ((v11 ^ v11) >= 245)
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                if (!v13) {
                                    v28 := GenS2{F0: (((!(v13 && v13)) || true) || (v13 || (!(v13 && v13)))), F1: 465, F2: -1}
                                    v12 = (v12 / (((^(v12))) | 1))
                                    _ = v28
                                } else {
                                    v12 = 42
                                    _ = len(v1)
                                    var v29 int16 = 823
                                    delete(v2, (fi14 + -9223372036854775808))
                                    _ = v29
                                }
                            }
                        case "a":
                            var v30 float64 = 0.0
                            _ = v30
                        }
                    case 816:
                        var v31 uint8 = (^(v11))
                        _ = v31
                    case -1:
                        v1[""] = (!(42 < 1))
                    }
                    if (2147483647 == 42) {
                        switch fi14 {
                        case -9223372036854775808:
                            var v32 string = ("foo bar" + ("foo bar" + ""))
                            _ = v32
                        }
                        var v33 int8 = -128
                        if len(v4) > 0 {
                            v4[0] = (("hello" + ("\n\t" + "")) + "")
                        }
                        _ = v33
                    }
                    var v34 int = 0
                    var v35 uint16 = 711
                    _ = v34
                    _ = v35
                case "hello":
                    if v13 {
                        var v36 uint32 = 4294967295
                        _ = v36
                    }
                }
                _ = fi16
            }
            if v13 {
                for fi37 := 0; fi37 < 8; fi37++ {
                    _ = len(v5)
                    _ = fi37
                }
                if v13 {
                    switch "" {
                    case "a":
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v39 := GenS3{F0: (-(v12))}
                            var v40 uint32 = 0
                            var v41 string = "a"
                            v42 := v1[v41]
                            var v43 uint = 225
                            if (966 <= 109) {
                                if ((v41 + ((v41 + ("\n\t" + v41)) + v41)) > v41) {
                                    var v44 int8 = 127
                                    v45 := GenS2{F0: (!true), F1: v43, F2: v12}
                                    var v46 bool = (v42 || v42)
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    switch v41 {
                                    case "":
                                        v43 = 255
                                    case "a":
                                        switch (((v41 + "a") + (v41 + v41)) + v41) {
                                        case "a":
                                            fmt.Println(v13)
                                            if v13 {
                                                switch (v41 + v41) {
                                                case "hello":
                                                    var v47 int32 = 307
                                                    _ = v47
                                                case "foo bar":
                                                    _ = len(v5)
                                                default:
                                                    var v48 uint64 = 0
                                                    _ = v48
                                                }
                                                v49 := GenS4{F0: 255, F1: 2147483647, F2: fi14}
                                                v41 = "hello"
                                                v1[("" + "foo bar")] = ((1 != 0) || (!(1 <= 2)))
                                                _ = v49
                                            }
                                            v40 = 2
                                        case "":
                                            fmt.Println(v12)
                                        default:
                                            v50 := GenS5{F0: ("\n\t" + v41)}
                                            var v51 float32 = 0.0
                                            var v52 float32 = (v51 - ((0.0 * v51) / v51))
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    if false {
                                        fmt.Println(v41, v43)
                                        var v53 int32 = 2147483647
                                        v2[9223372036854775807] = fi14
                                        v4 = append(v4, (v41 + v41))
                                        var v54 uint8 = 255
                                        _ = v53
                                        _ = v54
                                    } else {
                                        var v55 int = (fi38 % ((fi38) | 1))
                                        _ = v55
                                    }
                                    var v56 int = fi38
                                    _ = v56
                                }
                                var v57 float32 = float32(float64(1e10))
                                _ = v57
                            }
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                    case "":
                        var v58 uint = 319
                        _ = v58
                    }
                } else {
                    v59 := GenS2{F0: false, F1: 42, F2: 77}
                    _ = v59
                }
            } else {
                v60 := GenS2{F0: ((((float64(3.14159) + -2.25) / (1e10 * float64(((((1.5 / float32(0.1)) * ((((float32((3.14159 - ((1e10 * 0.0) * -2.25))) + (0.1 * ((-2.25 + -2.25) - 3.14159))) - 0.1) / (-2.25 + 0.0)) + (0.0 / ((1e10 / 1.0) + 1e10)))) / 1e10) + 1.5)))) * ((3.14159 - (-2.25 * float64(((1.5 / (1.5 + 0.0)) * (0.1 / 1e10))))) * (0.0 - float64(1e10)))) < (0.0 / 1.5)), F1: 946, F2: (v12 << ((334) & 63))}
                _ = v60
            }
            _ = fi14
        }
    }
    delete(v2, 32767)
    delete(v1, "a")
    for fi61 := 0; fi61 < 17; fi61++ {
        v62 := GenS3{F0: 0}
        v11 += v11
        v63 := GenS2{F0: (fi61 >= 128), F1: 65535, F2: v12}
        _ = fi61
        _ = v62
        _ = v63
    }
    var v64 float64 = (1e10 + 3.14159)
    _ = len(v2)
    var v65 uint8 = 42
    v12 = (v12 + -9223372036854775808)
    var v66 uint32 = 1
    v66 = v66
    var v67 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
