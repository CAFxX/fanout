package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 int32; F2 uint }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0() int16 {
    v5 := []int{}
    var v6 int32 = 1
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 += 42
        _ = fi7
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8) string {
    switch "hello" {
    case "":
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(fi8)
            if false {
                switch fi8 {
                case -752:
                    v9 := GenS1{F0: -615, F1: -2147483648, F2: (1 >= 1)}
                    _ = v9
                case -9223372036854775808:
                    for fi10 := 0; fi10 < 14; fi10++ {
                        fmt.Println(p0, fi10, fi8)
                        p0 += p0
                        _ = fi10
                    }
                default:
                    v11 := GenS2{F0: "hello"}
                    _ = v11
                }
            } else {
                var v12 bool = (1e10 >= (float32(1.5) * (0.0 / 0.1)))
                _ = v12
            }
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint16 = 42
                switch ((((("" + "\n\t") + ("a" + "\n\t")) + ("" + "a")) + (("a" + "hello") + "hello")) + ("a" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t"))) {
                case "a":
                    v14 = v14
                }
                switch (fi13 - (fi8 / (((fi8 / ((fi8) | 1))) | 1))) {
                case -346:
                    p0 -= p0
                case 4294967295:
                    v15 := GenS2{F0: "foo bar"}
                    _ = v15
                case 465:
                    v16 := GenS2{F0: "\n\t"}
                    _ = v16
                default:
                    var v17 uint32 = 2
                    _ = v17
                }
                v18 := []int{fi13, (fi8 << ((42) & 63))}
                _ = fi13
                _ = v14
                _ = v18
            }
            _ = fi8
        }
    case "hello":
        var v19 int = 128
        _ = v19
    }
    switch "a" {
    case "hello":
        if true {
            for fi20 := 0; fi20 < 13; fi20++ {
                p0 += 1
                var v21 string = "\n\t"
                v21 = ("" + v21)
                _ = fi20
                _ = v21
            }
            if (true || false) {
                v22 := GenS3{F0: 88, F1: -1, F2: 18446744073709551615}
                _ = v22
            } else {
                switch "a" {
                case "a":
                    switch 32767 {
                    case 9223372036854775807:
                        v23 := []int{42, 42, 632, 0}
                        _ = v23
                    case -9223372036854775808:
                        var v24 float64 = 0.0
                        _ = v24
                    default:
                        v25 := []int{934}
                        _ = v25
                    }
                case "":
                    p0 -= p0
                default:
                    p0 = 1
                }
                switch 42 {
                case 42:
                    if (!true) {
                        if true {
                            if false {
                                p0 = p0
                                p0 -= p0
                            }
                            var v26 uint64 = 591
                            fmt.Println(p0, p0)
                            _ = v26
                        } else {
                            p0 -= p0
                            if false {
                                for fi27 := 0; fi27 < 3; fi27++ {
                                    var v28 int64 = 0
                                    if false {
                                        v29 := GenS2{F0: ""}
                                        var v30 uint = 18446744073709551615
                                        _ = v29
                                        _ = v30
                                    }
                                    v31 := GenS1{F0: (v28 / ((v28) | 1)), F1: 2147483647, F2: true}
                                    _ = fi27
                                    _ = v28
                                    _ = v31
                                }
                            }
                        }
                        switch 4294967295 {
                        case 42:
                            fmt.Println(p0)
                        default:
                            p0 = 0
                        }
                        v32 := GenS4{F0: 2, F1: 2147483647, F2: 0}
                        _ = v32
                    }
                case -1:
                    v33 := GenS1{F0: -1, F1: 2147483647, F2: (4294967295 > 2147483647)}
                    _ = v33
                default:
                    p0 -= p0
                }
            }
        }
    }
    _ = p0
    return ""
}

func main() {
    var v1 uint32 = 65535
    var v2 float64 = 0.0
    var v3 bool = true
    v4 := make(map[string]string)
    var v34 int8 = -128
    if (!false) {
        switch "a" {
        case "hello":
            if v3 {
                v2 = 0.1
                switch "\n\t" {
                case "\n\t":
                    switch -1 {
                    case 1:
                        v2 += (1.5 - v2)
                    default:
                        for fi35 := 0; fi35 < 15; fi35++ {
                            if true {
                                _ = len(v4)
                                switch 9223372036854775807 {
                                case 42:
                                    var v36 string = ("\n\t" + ("" + ("foo bar" + "")))
                                    _ = v36
                                case 95:
                                    for fi37 := 0; fi37 < 19; fi37++ {
                                        if false {
                                            v34 -= (v34 - v34)
                                            switch fi35 {
                                            case 127:
                                                v38 := GenS2{F0: ""}
                                                _ = v38
                                            }
                                        } else {
                                            v39 := GenS2{F0: "a"}
                                            var v40 uint64 = 18446744073709551615
                                            v41 := []int{-9223372036854775808, int(v1), (fi37 & fi37)}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        delete(v4, ("a" + "a"))
                                        _ = fi37
                                    }
                                case -1:
                                    switch "" {
                                    case "":
                                        var v42 string = ("a" + "\n\t")
                                        _ = v42
                                    case "foo bar":
                                        v3 = (v2 <= 3.14159)
                                    case "hello":
                                        fmt.Println(v1, v2, v2)
                                    }
                                }
                                v43 := []int{1}
                                _ = v43
                            } else {
                                if v3 {
                                    v34 = v34
                                }
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    v3 = true
                                    _ = fi44
                                }
                            }
                            for fi45 := 0; fi45 < 8; fi45++ {
                                _ = len(v4)
                                _ = fi45
                            }
                            _ = fi35
                        }
                    }
                default:
                    switch 0 {
                    case 127:
                        var v46 int16 = genFn0()
                        _ = v46
                    }
                }
            } else {
                var v47 uint8 = 0
                _ = v47
            }
        case "\n\t":
            var v48 float64 = v2
            _ = v48
        }
        if false {
            switch 4294967295 {
            case 1:
                var v49 string = "foo bar"
                _ = v49
            case 42:
                var v50 int = 0
                _ = v50
            }
        }
        v34 = -128
    } else {
        var v51 float64 = (v2 / (0.0 - (v2 - 0.0)))
        v34 = (^(v34))
        var v52 int16 = genFn0()
        _ = v51
        _ = v52
    }
    v53 := GenS3{F0: 191, F1: 0, F2: 18446744073709551615}
    v54 := []int{}
    switch "" {
    case "a":
        for fi55 := 0; fi55 < 11; fi55++ {
            v34 -= v34
            v56 := GenS3{F0: 42, F1: 0, F2: 2}
            if (!(907 < 0)) {
                if v3 {
                    v3 = v3
                    v34 = v34
                } else {
                    v57 := GenS3{F0: 128, F1: 1, F2: 275}
                    fmt.Println(v34, v34, v2)
                    _ = v57
                }
            }
            _ = fi55
            _ = v56
        }
    }
    fmt.Println(v3, v2, v2)
    if len(v54) > 0 {
        v54[0] = -200
    }
    for fi58 := 0; fi58 < 6; fi58++ {
        _ = len(v4)
        _ = fi58
    }
    var v59 int8 = (v34 % ((1) | 1))
    v60 := GenS5{F0: (^(v34)), F1: -32768}
    if true {
        var v61 uint = 965
        _ = v61
    } else {
        fmt.Println(v59)
        if (!v3) {
            var v62 uint64 = 1
            _ = v62
        } else {
            if false {
                v3 = (v2 == (-2.25 / v2))
                v3 = true
            }
            v63 := GenS2{F0: (("\n\t" + "a") + "")}
            switch 4294967295 {
            case 32767:
                var v64 uint64 = 4294967295
                _ = v64
            case 42:
                v3 = (71 > 1)
            }
            _ = v63
        }
    }
    v34 = v59
    if ((v2 < 3.14159) && false) {
        if true {
            v2 = ((v2 / (3.14159 - float64(0.1))) / 1.0)
        }
        fmt.Println(v34)
    }
    v3 = (-156 > 408)
    for fi65 := 0; fi65 < 10; fi65++ {
        var v66 uint64 = uint64(v1)
        v4[(("hello" + "\n\t") + "\n\t")] = "a"
        _ = fi65
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v53
    _ = v54
    _ = v59
    _ = v60
}
