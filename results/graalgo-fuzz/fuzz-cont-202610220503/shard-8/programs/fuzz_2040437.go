package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint64 }
func genFn0() int32 {
    v6 := []int{128, 0, -2147483648}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 bool = true
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: fi7}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return -1
}
func genFn1(p0 uint16) bool {
    v11 := []int{1}
    if len(v11) > 0 {
        v11[0] = 127
    }
    if (9223372036854775807 != 128) {
        for fi12 := 0; fi12 < 2; fi12++ {
            if (!(1 == 0)) {
                if (!("\n\t" < (("a" + "a") + ""))) {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        p0 = (p0 - p0)
                        _ = fi13
                    }
                    switch fi12 {
                    case -162:
                        if (("foo bar" + ("" + "foo bar")) >= "\n\t") {
                            var v14 int = fi12
                            if true {
                                p0 -= 42
                            } else {
                                p0 += p0
                            }
                            _ = v14
                        } else {
                            var v15 int16 = 897
                            if (9223372036854775807 >= 290) {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                if false {
                                    if true {
                                        v16 := GenS2{F0: fi12, F1: 73}
                                        switch fi12 {
                                        case -2147483648:
                                            var v17 string = (("a" + ("" + (("" + "") + (("\n\t" + ("foo bar" + "\n\t")) + "foo bar")))) + "\n\t")
                                            _ = v17
                                        }
                                        var v18 float32 = (-2.25 * 1.5)
                                        _ = v16
                                        _ = v18
                                    } else {
                                        var v19 string = ("hello" + "\n\t")
                                        _ = v19
                                    }
                                }
                                var v20 uint64 = 1
                                _ = v20
                            } else {
                                var v21 bool = ((fi12 / (((-(fi12))) | 1)) != (fi12 % ((fi12) | 1)))
                                var v22 float64 = (0.1 * 0.0)
                                _ = v21
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 20; fi23++ {
                                var v24 float64 = 0.0
                                v24 -= float64((-2.25 - 0.0))
                                _ = fi23
                                _ = v24
                            }
                            _ = v15
                        }
                    case 32767:
                        _ = len(v11)
                    case 563:
                        switch "" {
                        case "a":
                            fmt.Println(p0, fi12)
                        case "\n\t":
                            fmt.Println(fi12, p0)
                        case "hello":
                            p0 = (p0 % ((p0) | 1))
                        default:
                            var v25 int64 = -544
                            _ = v25
                        }
                    }
                    p0 += 255
                }
                var v26 string = "foo bar"
                if len(v11) > 0 {
                    v11[0] = -9223372036854775808
                }
                _ = v26
            }
            _ = fi12
        }
        v27 := GenS3{F0: 42, F1: (^(p0)), F2: 127}
        var v28 uint = 42
        _ = v27
        _ = v28
    } else {
        if true {
            var v29 uint8 = 228
            _ = v29
        }
        p0 += p0
    }
    switch -59 {
    case -578:
        v30 := GenS4{F0: -1}
        _ = v30
    case 4294967295:
        if false {
            if false {
                switch 32768 {
                case -965:
                    v11 = append(v11, 9223372036854775807)
                case 128:
                    switch "foo bar" {
                    case "a":
                        var v31 int16 = 42
                        _ = v31
                    default:
                        if (!true) {
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                p0 += p0
                            default:
                                for fi32 := 0; fi32 < 7; fi32++ {
                                    v33 := GenS5{F0: 4294967295, F1: "foo bar"}
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            p0 = 65535
                        }
                    }
                case 2147483647:
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v35 := GenS6{F0: true}
                        p0 = p0
                        var v36 int = (fi34 / ((fi34) | 1))
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                default:
                    if (0 < 789) {
                        v37 := GenS1{F0: 4294967295}
                        _ = v37
                    } else {
                        if len(v11) > 0 {
                            v11[0] = 2147483647
                        }
                        switch "" {
                        case "hello":
                            if (!true) {
                                if (!true) {
                                    p0 -= (p0 / ((p0) | 1))
                                }
                                if (42 < 909) {
                                    switch "foo bar" {
                                    case "":
                                        var v38 int = 128
                                        _ = v38
                                    default:
                                        v39 := GenS7{F0: 42, F1: 2147483647}
                                        _ = v39
                                    }
                                    fmt.Println(p0)
                                }
                            } else {
                                v40 := GenS3{F0: 297, F1: p0, F2: 55}
                                switch 32767 {
                                case 2147483647:
                                    p0 = (p0 ^ p0)
                                case 32768:
                                    for fi41 := 0; fi41 < 4; fi41++ {
                                        switch ("a" + ("foo bar" + ((("foo bar" + "foo bar") + ("\n\t" + ((("hello" + (("foo bar" + "") + "a")) + ("\n\t" + "a")) + ""))) + ("\n\t" + ("a" + "\n\t"))))) {
                                        case "\n\t":
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        case "":
                                            v42 := GenS8{F0: 742}
                                            _ = v42
                                        case "foo bar":
                                            p0 -= (p0 - 1)
                                        default:
                                            var v43 int8 = 127
                                            _ = v43
                                        }
                                        var v44 bool = true
                                        _ = fi41
                                        _ = v44
                                    }
                                }
                                _ = v40
                            }
                        case "\n\t":
                            fmt.Println(p0)
                        case "":
                            fmt.Println(p0)
                        default:
                            switch "\n\t" {
                            case "":
                                var v45 bool = (!("" <= ""))
                                _ = v45
                            case "\n\t":
                                if len(v11) > 0 {
                                    v11[0] = 994
                                }
                            case "foo bar":
                                p0 += p0
                            default:
                                v46 := GenS2{F0: -1, F1: 112}
                                _ = v46
                            }
                        }
                    }
                }
                v47 := GenS1{F0: 888}
                _ = v47
            }
        } else {
            var v48 int64 = 1
            var v49 float32 = 0.1
            switch "hello" {
            case "foo bar":
                var v50 int8 = -4
                _ = v50
            case "hello":
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            _ = v48
            _ = v49
        }
    default:
        switch -317 {
        case 4294967295:
            var v51 int16 = 1000
            _ = v51
        default:
            p0 = p0
        }
    }
    var v52 bool = true
    _ = p0
    _ = v11
    _ = v52
    return (!(-1 != 42))
}

func main() {
    v1 := []int{128}
    var v2 float64 = (((1e10 / 1e10) - (0.1 / 1e10)) * (1e10 * float64(0.1)))
    var v3 float64 = 1e10
    v4 := make(map[int]int)
    v5 := []uint32{2147483647, 4294967295, 404, 65535}
    if false {
        if false {
            fmt.Println(v2, v2)
            var v53 uint = 42
            if false {
                if false {
                    var v54 int = 32768
                    _ = v54
                }
                fmt.Println(v53, v53, v3)
                v55 := GenS8{F0: 488}
                _ = v55
            }
            _ = v53
        }
        v56 := GenS2{F0: 127, F1: 2}
        var v57 uint32 = 2
        _ = v56
        _ = v57
    }
    var v58 float32 = 0.1
    v59 := GenS3{F0: 42, F1: 42, F2: 102}
    for fi60 := 0; fi60 < 7; fi60++ {
        var v61 int64 = 0
        _ = len(v4)
        _ = len(v4)
        _ = fi60
        _ = v61
    }
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 float32 = (float32((float64(v58) + (v2 * 1.5))) * 1.5)
        var v64 int16 = 42
        _ = fi62
        _ = v63
        _ = v64
    }
    var v65 float64 = 0.0
    v58 = v58
    v66 := GenS6{F0: false}
    v2 = 1e10
    var v67 uint16 = 942
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi68 := 0; fi68 < 10; fi68++ {
        fmt.Println(v58, v3)
        _ = fi68
    }
    v58 = (0.1 * 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v65
    _ = v66
    _ = v67
}
