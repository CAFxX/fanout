package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 int; F2 uint }
func genFn0(p0 bool) string {
    v7 := GenS1{F0: 2, F1: 0, F2: "foo bar"}
    if false {
        switch -2147483648 {
        case -9223372036854775808:
            switch "\n\t" {
            case "a":
                var v8 uint16 = 65535
                _ = v8
            case "hello":
                fmt.Println(p0)
            }
        case 32767:
            v9 := GenS2{F0: 2, F1: 0}
            _ = v9
        default:
            switch ("hello" + ("a" + "foo bar")) {
            case "hello":
                var v10 bool = ((465 < 42) || (0 != 42))
                _ = v10
            case "\n\t":
                switch 32768 {
                case 1:
                    p0 = (p0 || true)
                case 2147483647:
                    v11 := []int{-532, 1, 4294967295, 1}
                    _ = v11
                default:
                    p0 = p0
                }
            case "a":
                for fi12 := 0; fi12 < 13; fi12++ {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 int64 = 1
                        if true {
                            var v15 float32 = (((1.5 + ((0.0 - ((0.0 - 1.5) + (0.0 / 1e10))) / (-2.25 * (3.14159 - 1e10)))) - (0.1 / 3.14159)) - 3.14159)
                            v14 = v14
                            _ = v15
                        } else {
                            var v16 int64 = v14
                            if true {
                                v16 -= v16
                                fmt.Println(v14, v14)
                            }
                            _ = v16
                        }
                        v17 := GenS3{F0: -128}
                        _ = fi13
                        _ = v14
                        _ = v17
                    }
                    _ = fi12
                }
            default:
                v18 := GenS2{F0: 0, F1: 255}
                _ = v18
            }
        }
        switch "a" {
        case "foo bar":
            var v19 float64 = 1e10
            _ = v19
        case "":
            p0 = p0
        case "a":
            p0 = (1 != 210)
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS3{F0: -1}
                p0 = (p0 || (true && true))
                var v22 uint32 = 42
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        v23 := GenS1{F0: 576, F1: -128, F2: ("foo bar" + "hello")}
        _ = v23
    } else {
        switch ("foo bar" + "\n\t") {
        case "":
            if (-1 == -1) {
                v24 := GenS4{F0: 0, F1: 8}
                _ = v24
            }
        case "foo bar":
            var v25 uint8 = 93
            _ = v25
        case "a":
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := []int{}
                if (1e10 == 1e10) {
                    switch "" {
                    case "hello":
                        p0 = (0 > 808)
                    case "\n\t":
                        if len(v27) > 0 {
                            v27[0] = fi26
                        }
                    case "":
                        var v28 uint64 = 4294967295
                        _ = v28
                    default:
                        var v29 int8 = -128
                        _ = v29
                    }
                    if p0 {
                        v30 := GenS5{F0: 1, F1: 34}
                        var v31 uint8 = 128
                        _ = v30
                        _ = v31
                    }
                    if p0 {
                        var v32 int16 = 1000
                        for fi33 := 0; fi33 < 8; fi33++ {
                            var v34 float32 = -2.25
                            switch (fi33 % ((fi26) | 1)) {
                            case 2147483647:
                                fmt.Println(v34)
                            default:
                                v32 += -1
                            }
                            var v35 int64 = -426
                            v27 = append(v27, fi26)
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS4{F0: -245, F1: 42}
                        _ = v32
                        _ = v36
                    }
                } else {
                    p0 = p0
                    var v37 string = "\n\t"
                    _ = v37
                }
                if false {
                    var v38 int64 = -865
                    for fi39 := 0; fi39 < 8; fi39++ {
                        var v40 uint8 = 201
                        if false {
                            if true {
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                                var v41 uint32 = 65535
                                _ = v41
                            }
                            if len(v27) > 0 {
                                v27[0] = (fi26 * fi26)
                            }
                            p0 = (v40 < v40)
                        }
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                } else {
                    p0 = false
                    fmt.Println(fi26)
                    p0 = false
                }
                _ = fi26
                _ = v27
            }
        }
        if false {
            var v42 int = -282
            _ = v42
        } else {
            v43 := []int{}
            if p0 {
                v43 = append(v43, -78)
                if false {
                    if ("" == "hello") {
                        for fi44 := 0; fi44 < 18; fi44++ {
                            var v45 uint8 = 42
                            var v46 uint = 2
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                    v47 := GenS2{F0: 0, F1: 65535}
                    _ = v47
                }
                var v48 int32 = -2147483648
                _ = v48
            }
            fmt.Println(p0)
            _ = v43
        }
        v49 := GenS5{F0: 1, F1: 96}
        _ = v49
    }
    _ = p0
    _ = v7
    return ("hello" + "foo bar")
}

func main() {
    var v1 int64 = 632
    v2 := make(map[int]string)
    var v3 uint = 995
    var v4 uint = v3
    var v5 float32 = 1.5
    var v6 uint32 = 885
    v50 := GenS5{F0: 510, F1: (^(v6))}
    switch 42 {
    case 9223372036854775807:
        v51 := GenS2{F0: 99, F1: (v4 / ((v3) | 1))}
        _ = v51
    case -436:
        switch "" {
        case "foo bar":
            fmt.Println(v3, v4)
        case "a":
            switch "" {
            case "hello":
                v52 := GenS4{F0: v1, F1: 1}
                _ = v52
            default:
                var v53 bool = true
                _ = v53
            }
        default:
            v54 := GenS1{F0: 18446744073709551615, F1: 71, F2: (("\n\t" + "\n\t") + ("hello" + ""))}
            _ = v54
        }
    default:
        var v55 string = "hello"
        _ = v55
    }
    if true {
        v3 = (^(v3))
    } else {
        v6 += v6
    }
    v1 = (v1 - v1)
    v56 := GenS6{F0: 42}
    var v57 float64 = 1.5
    v2[813] = ("a" + ("" + "a"))
    for fi58 := 0; fi58 < 7; fi58++ {
        v59 := GenS7{F0: (fi58 + fi58), F1: (fi58 & (^(fi58))), F2: (v4 % ((v4) | 1))}
        var v60 int64 = (^(v1))
        v6 -= (v6 - (v6 + (^(v6))))
        _ = fi58
        _ = v59
        _ = v60
    }
    var v61 int = 9223372036854775807
    if (false && false) {
        var v62 bool = (((float64(((v5 / v5) - 1.5)) * v57) / 1.5) > 0.1)
        _ = v62
    } else {
        for fi63 := 0; fi63 < 3; fi63++ {
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 string = ("" + "\n\t")
                v65 = (v65 + (("\n\t" + v65) + v65))
                var v66 int8 = 42
                _ = fi64
                _ = v65
                _ = v66
            }
            var v67 int64 = (-(v1))
            _ = fi63
            _ = v67
        }
    }
    v57 += 0.1
    switch 59 {
    case -332:
        v68 := []int{v61}
        _ = v68
    }
    v69 := GenS3{F0: 127}
    v57 = v57
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v56
    _ = v57
    _ = v6
    _ = v61
    _ = v69
}
