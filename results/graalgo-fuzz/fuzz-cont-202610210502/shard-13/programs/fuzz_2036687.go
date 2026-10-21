package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint; F2 int }
type GenS7 struct { F0 string }
func genFn0(p0 uint8, p1 int64) uint64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: -128, F1: 721}
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint32 = 644
            v8 := GenS2{F0: v7, F1: 139}
            var v9 uint8 = (p0 / ((p0) | 1))
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 float32 = -2.25
        _ = fi4
        _ = v10
        _ = v5
    }
    p1 += p1
    switch 32768 {
    case 32767:
        p0 += 159
    case -9223372036854775808:
        var v11 uint32 = 42
        _ = v11
    case 4294967295:
        v12 := []int{127, -50, 127}
        _ = v12
    default:
        v13 := []int{0, 1, -9223372036854775808}
        _ = v13
    }
    var v14 string = ""
    _ = p0
    _ = p1
    _ = v14
    return 1
}
func genFn1() int8 {
    if false {
        switch ((("" + "\n\t") + (("\n\t" + ("\n\t" + "foo bar")) + "foo bar")) + ("\n\t" + (("hello" + "hello") + ("foo bar" + "a")))) {
        case "":
            var v15 int = 977
            _ = v15
        case "\n\t":
            var v16 uint16 = 0
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint64 = 2
            switch ("a" + "a") {
            case "":
                if (!(v18 <= (v18 & genFn0(128, -866)))) {
                    var v19 float64 = ((0.0 - 0.0) * (1.5 - float64(0.0)))
                    _ = v19
                }
            case "a":
                v18 = (^(v18))
            case "\n\t":
                v18 = 42
            }
            v20 := []int{}
            _ = fi17
            _ = v18
            _ = v20
        }
    }
    var v21 string = "\n\t"
    fmt.Println(v21)
    var v22 uint = 255
    _ = v21
    _ = v22
    return -128
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS3{F0: (fi23 + fi23), F1: 42}
        if true {
            var v25 int32 = -1
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 uint64 = 1
            if false {
                switch "a" {
                case "hello":
                    var v28 float32 = 1.5
                    _ = v28
                case "":
                    v27 += v27
                case "a":
                    var v29 uint64 = v27
                    _ = v29
                }
                switch "\n\t" {
                case "hello":
                    v27 = genFn0(34, -9223372036854775808)
                case "":
                    var v30 int = fi23
                    _ = v30
                case "\n\t":
                    fmt.Println(fi23, v27)
                default:
                    v27 = 0
                }
                v27 -= (v27 - v27)
            } else {
                var v31 uint8 = 0
                v32 := []int{(fi26 << ((255) & 63)), 127, 342, (^(fi23))}
                _ = v31
                _ = v32
            }
            v33 := GenS1{F0: genFn1(), F1: 65535}
            _ = fi26
            _ = v27
            _ = v33
        }
        _ = fi23
        _ = v24
    }
    v34 := []int{128, -9223372036854775808, 128, -9223372036854775808}
    v35 := GenS2{F0: 180, F1: 129}
    switch ("a" + "a") {
    case "a":
        var v36 int = 484
        _ = v36
    default:
        if false {
            var v37 bool = (-32768 <= -1)
            v1[-992] = (58 > 42)
            fmt.Println(v37)
            _ = v37
        }
    }
    var v38 uint64 = 0
    v2[""] = 65535
    switch "\n\t" {
    case "hello":
        v38 = 2
    case "\n\t":
        var v39 int16 = 407
        _ = v39
    case "a":
        if true {
            var v40 uint64 = genFn0(255, 565)
            var v41 int64 = 9223372036854775807
            if true {
                var v42 float32 = 0.0
                if false {
                    if false {
                        switch 42 {
                        case 875:
                            v43 := GenS2{F0: 65535, F1: 1}
                            _ = v43
                        }
                        v44 := GenS3{F0: -1, F1: v41}
                        v40 -= v40
                        _ = v44
                    }
                    v41 -= v41
                    for fi45 := 0; fi45 < 7; fi45++ {
                        v41 = (v41 % (((v41 << ((0) & 63))) | 1))
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                        v46 := v2[""]
                        _ = fi45
                        _ = v46
                    }
                } else {
                    v41 -= v41
                    if (!((v42 + 0.1) > v42)) {
                        v42 = (float32(3.14159) * v42)
                        v47 := GenS3{F0: 970, F1: 42}
                        v40 = v40
                        _ = v47
                    } else {
                        var v48 uint64 = (v40 & (v38 >> ((1) & 63)))
                        var v49 int64 = v41
                        _ = v48
                        _ = v49
                    }
                }
                _ = v42
            } else {
                var v50 int = -9223372036854775808
                v50 += (-(v50))
                _ = v50
            }
            _ = v40
            _ = v41
        }
    }
    if false {
        v34 = append(v34, 0)
    } else {
        v51 := GenS1{F0: genFn1(), F1: 2}
        _ = v51
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        fmt.Println(v38, fi52)
        for fi53 := 0; fi53 < 11; fi53++ {
            if (!true) {
                switch fi52 {
                case -989:
                    v38 -= (v38 & v38)
                case -9223372036854775808:
                    switch "\n\t" {
                    case "a":
                        switch fi53 {
                        case 0:
                            v54 := GenS4{F0: 2}
                            _ = v54
                        default:
                            v38 = v38
                            if len(v34) > 0 {
                                v34[0] = 42
                            }
                            fmt.Println(fi52)
                        }
                    case "":
                        v55 := GenS2{F0: 911, F1: 1}
                        _ = v55
                    }
                }
            } else {
                switch fi53 {
                case 284:
                    v34 = append(v34, (fi52 - (-(fi53))))
                case -2147483648:
                    switch 4294967295 {
                    case 42:
                        v56 := GenS5{F0: 2147483647, F1: 1}
                        _ = v56
                    default:
                        if len(v34) > 0 {
                            v34[0] = (fi52 | (fi52 << ((42) & 63)))
                        }
                    }
                default:
                    v57 := GenS6{F0: true, F1: 255, F2: fi52}
                    _ = v57
                }
                v38 += (v38 << ((255) & 63))
                if (false || (32767 >= -118)) {
                    var v58 int8 = 42
                    var v59 bool = (genFn0(42, 0) >= 2)
                    fmt.Println(v59, v58, v58)
                    v60 := GenS7{F0: ("" + "a")}
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    var v61 int32 = 1
                    _ = v61
                }
                var v62 uint16 = 32767
                var v63 uint64 = v38
                var v64 uint32 = 771
                v64 = (v64 - v64)
                _ = v62
                _ = v63
                _ = v64
            }
            fmt.Println(fi53, fi53)
            var v65 uint = 1
            _ = fi53
            _ = v65
        }
        _ = fi52
    }
    v66 := GenS6{F0: true, F1: 128, F2: 2147483647}
    v38 = (v38 + (^(v38)))
    var v67 float32 = (1e10 / 1.5)
    switch "foo bar" {
    case "hello":
        v34 = append(v34, 127)
    case "a":
        _ = len(v34)
    case "foo bar":
        var v68 float32 = v67
        _ = v68
    default:
        if (!(!false)) {
            if false {
                switch 42 {
                case 1:
                    if true {
                        var v69 uint32 = 486
                        v67 = float32(0.0)
                        v38 = v38
                        _ = v69
                    }
                case 128:
                    if true {
                        v38 += (v38 << ((1) & 63))
                    }
                }
            } else {
                var v70 string = "a"
                var v71 uint32 = 2
                _ = v70
                _ = v71
            }
            switch ("foo bar" + "hello") {
            case "":
                for fi72 := 0; fi72 < 5; fi72++ {
                    fmt.Println(v67, v67, v67)
                    for fi73 := 0; fi73 < 6; fi73++ {
                        v67 = 1e10
                        v38 += 571
                        var v74 int64 = -137
                        _ = fi73
                        _ = v74
                    }
                    fmt.Println(fi72)
                    _ = fi72
                }
            }
        } else {
            fmt.Println(v67, v38)
            var v75 uint32 = 42
            v34 = append(v34, -1)
            _ = v75
        }
    }
    if (!true) {
        _ = len(v34)
        var v76 int32 = 2147483647
        _ = v76
    } else {
        v38 = v38
    }
    _ = len(v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v38:", v38)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v66
    _ = v67
}
