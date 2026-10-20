package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 uint; F2 int8 }
func genFn0(p0 int32) int32 {
    v7 := []int{316}
    if len(v7) > 0 {
        v7[0] = -721
    }
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]uint32)
    v3 := []string{(("a" + ("a" + "hello")) + "")}
    var v4 int64 = -9223372036854775808
    v5 := []string{"", "\n\t", ("foo bar" + "a"), "\n\t"}
    var v6 float32 = ((0.0 + 1.5) / 1.0)
    if true {
        v4 = (v4 & (v4 - v4))
        var v8 int32 = 0
        _ = v8
    }
    var v9 int = v1
    _ = len(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v6 = v6
            _ = fi11
        }
        _ = len(v2)
        v12 := GenS1{F0: (!(!false)), F1: ("" + ("hello" + (("foo bar" + ("a" + "\n\t")) + "\n\t")))}
        _ = fi10
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if false {
            if (1 >= 42) {
                v14 := GenS1{F0: (!true), F1: ("\n\t" + "hello")}
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v9, fi13, fi15)
                    _ = fi15
                }
                v4 -= 533
                _ = v14
            } else {
                var v16 uint = 4294967295
                fmt.Println(v4)
                var v17 float32 = (v6 * v6)
                _ = v16
                _ = v17
            }
        } else {
            var v18 int64 = 42
            v6 += 1e10
            _ = v18
        }
        _ = fi13
    }
    if (!true) {
        var v19 uint64 = 647
        fmt.Println(v19, v4)
        _ = v19
    } else {
        if (true || true) {
            switch "" {
            case "":
                for fi20 := 0; fi20 < 7; fi20++ {
                    v6 += v6
                    var v21 uint8 = 0
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 uint = 42
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            case "foo bar":
                v1 -= 127
            case "\n\t":
                v24 := GenS2{F0: 0}
                _ = v24
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                v26 := GenS2{F0: 1000}
                _ = fi25
                _ = v26
            }
        } else {
            switch v1 {
            case -2147483648:
                var v27 int32 = 0
                _ = v27
            case -370:
                v4 += 9223372036854775807
            case -913:
                switch v9 {
                case -231:
                    switch "foo bar" {
                    case "hello":
                        var v28 int64 = (^(v4))
                        _ = v28
                    case "":
                        var v29 uint = 1
                        _ = v29
                    case "\n\t":
                        for fi30 := 0; fi30 < 16; fi30++ {
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 uint64 = 18446744073709551615
                                v33 := GenS3{F0: "hello"}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = fi30
                        }
                    default:
                        var v34 uint32 = 36
                        _ = v34
                    }
                }
            default:
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS1{F0: ((42 >= 0) && (1.5 > (0.0 - -2.25))), F1: "foo bar"}
                    _ = fi35
                    _ = v36
                }
            }
            v4 -= v4
            var v37 int64 = v4
            _ = v37
        }
    }
    switch (((("foo bar" + (("\n\t" + "") + ("a" + ""))) + ("a" + ("hello" + ""))) + "a") + "hello") {
    case "foo bar":
        if ((v6 * v6) != 0.1) {
            v2[int(v4)] = 65535
            fmt.Println(v4)
            v2[(v9 & v9)] = 1
        }
    case "hello":
        v9 = v9
    case "\n\t":
        switch "hello" {
        case "\n\t":
            if false {
                v38 := GenS4{F0: 255, F1: 42, F2: 239}
                v39 := GenS5{F0: ((127 > 128) || (!(!((false || true) || (3.14159 != v6))))), F1: 2, F2: 85}
                if true {
                    for fi40 := 0; fi40 < 1; fi40++ {
                        if true {
                            v6 -= (((1e10 - 0.1) * 0.0) / 1e10)
                            fmt.Println(v9, v1, v6)
                        }
                        v2[v1] = 914
                        _ = fi40
                    }
                    if false {
                        fmt.Println(v9)
                    }
                    var v41 int16 = 32767
                    _ = v41
                }
                _ = v38
                _ = v39
            } else {
                switch (^(v9)) {
                case 957:
                    if true {
                        for fi42 := 0; fi42 < 10; fi42++ {
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            v1 -= (^(fi42))
                            v6 += v6
                            _ = fi42
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        var v43 uint64 = 18446744073709551615
                        _ = v43
                    }
                case -959:
                    var v44 int32 = 2147483647
                    _ = v44
                case 42:
                    for fi45 := 0; fi45 < 10; fi45++ {
                        fmt.Println(v1, v4)
                        fmt.Println(v1)
                        var v46 int64 = (-(v4))
                        _ = fi45
                        _ = v46
                    }
                }
                for fi47 := 0; fi47 < 11; fi47++ {
                    v3 = append(v3, "\n\t")
                    _ = fi47
                }
                for fi48 := 0; fi48 < 6; fi48++ {
                    if false {
                        var v49 float64 = ((0.0 / 1.5) + -2.25)
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 9; fi50++ {
                            v51 := GenS6{F0: false}
                            _ = fi50
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 13; fi52++ {
                            var v53 float32 = (v6 + -2.25)
                            _ = fi52
                            _ = v53
                        }
                        switch "foo bar" {
                        case "hello":
                            var v54 int16 = 1000
                            _ = v54
                        }
                    }
                    _ = fi48
                }
            }
        case "foo bar":
            fmt.Println(v1, v1)
        default:
            var v55 uint64 = 1
            _ = v55
        }
    default:
        v4 = v4
    }
    var v56 uint = 4294967295
    if (18446744073709551615 < 1) {
        if ("\n\t" >= ("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + "hello")))) {
            v4 -= 42
        }
        v6 = float32(0.0)
        v6 -= float32(3.14159)
    }
    switch "a" {
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if false {
            var v57 int8 = 100
            fmt.Println(v1, v6, v4)
            _ = v57
        } else {
            switch 391 {
            case 2147483647:
                if (0 != 1) {
                    fmt.Println(v1, v6, v1)
                    var v58 int = 742
                    _ = v58
                }
            case 42:
                if true {
                    if (!true) {
                        if true {
                            v59 := GenS2{F0: 0}
                            _ = v59
                        }
                    }
                    v4 += v4
                } else {
                    switch "a" {
                    case "":
                        if (-2.25 != (0.0 + 0.1)) {
                            v56 = (v56 >> (((v56 ^ (^(v56)))) & 63))
                            if len(v5) > 0 {
                                v5[0] = ("foo bar" + "foo bar")
                            }
                        }
                    default:
                        fmt.Println(v6, v9)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            var v60 uint = v56
            v61 := GenS1{F0: (!false), F1: "foo bar"}
            _ = v60
            _ = v61
        }
    }
    var v62 uint32 = 2147483648
    v6 = ((v6 + v6) / ((v6 + 1e10) + (0.0 - 1.5)))
    for fi63 := 0; fi63 < 6; fi63++ {
        switch (("a" + "a") + ("\n\t" + (("\n\t" + "foo bar") + ("a" + "")))) {
        case "\n\t":
            fmt.Println(v56, v9)
        case "foo bar":
            v64 := GenS1{F0: false, F1: "foo bar"}
            _ = v64
        case "":
            var v65 bool = ((fi63 + v1) <= (fi63 / (((v1 | v9)) | 1)))
            _ = v65
        }
        if len(v3) > 0 {
            v3[0] = ("" + ("\n\t" + ("a" + "a")))
        }
        switch "\n\t" {
        case "":
            v56 += v56
        default:
            v2[(fi63 * (v9 ^ fi63))] = (^(v62))
        }
        _ = fi63
    }
    if false {
        if true {
            v2[662] = 0
            v66 := GenS2{F0: -428}
            _ = v66
        } else {
            v67 := GenS7{F0: 1, F1: v56, F2: 1}
            _ = v67
        }
    } else {
        var v68 int16 = 66
        _ = v68
    }
    var v69 int8 = -1
    var v70 bool = false
    var v71 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v62
    _ = v69
    _ = v70
    _ = v71
    _ = v9
}
