package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int8; F2 string }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64, p1 bool, p2 int32) int64 {
    v7 := GenS1{F0: 2147483647, F1: -83}
    var v8 int32 = (p2 - (p2 ^ (-(p2))))
    v9 := GenS1{F0: 2147483648, F1: -1}
    v10 := GenS1{F0: 4294967295, F1: -68}
    var v11 float32 = (0.0 + 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return (-(p0))
}
func genFn1(p0 uint, p1 uint8, p2 string) uint {
    var v12 string = "foo bar"
    fmt.Println(v12, p1, p2)
    if (p1 <= p1) {
        v13 := GenS1{F0: 0, F1: 0}
        _ = v13
    } else {
        v12 = ("hello" + (p2 + (p2 + v12)))
        v14 := GenS2{F0: -2147483648}
        v12 = "a"
        _ = v14
    }
    if (1e10 > (1.5 - -2.25)) {
        v12 = "\n\t"
        v15 := GenS3{F0: 0, F1: 52}
        var v16 uint64 = 42
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p0 << ((p0) & 63))
}

func main() {
    var v1 int8 = 1
    var v2 int8 = (v1 / ((v1) | 1))
    var v3 int8 = v1
    v4 := make(map[string]int)
    var v5 bool = false
    v6 := []uint32{0, 4294967295, 4294967295}
    var v17 uint8 = 1
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            if len(v6) > 0 {
                v6[0] = 42
            }
            _ = fi19
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            v3 = (^(v3))
            v5 = v5
            switch (fi18 << ((281) & 63)) {
            case -159:
                v3 = v2
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            if (0.1 >= (0.1 * 0.0)) {
                if len(v6) > 0 {
                    v6[0] = 7
                }
                fmt.Println(v5)
            }
            var v21 string = "foo bar"
            _ = fi20
            _ = v21
        }
        _ = fi18
    }
    var v22 uint32 = 651
    v3 -= v2
    fmt.Println(v2, v17, v22)
    if true {
        switch -2147483648 {
        case 1:
            var v23 int = 0
            _ = v23
        default:
            v22 = v22
        }
        var v24 int8 = 100
        switch "" {
        case "hello":
            v25 := GenS1{F0: v22, F1: v2}
            _ = v25
        default:
            v17 -= v17
        }
        _ = v24
    } else {
        v5 = true
        if (v5 || ("\n\t" >= "a")) {
            if v5 {
                switch -184 {
                case 4294967295:
                    v26 := GenS3{F0: 122, F1: (v1 * (^(v1)))}
                    _ = v26
                case -1:
                    v17 = v17
                case 128:
                    switch 542 {
                    case 42:
                        v27 := GenS4{F0: -405}
                        _ = v27
                    default:
                        var v28 int64 = 0
                        _ = v28
                    }
                }
                switch ("foo bar" + ("" + ("a" + "a"))) {
                case "a":
                    v22 -= (^(v22))
                default:
                    if (v5 && ("foo bar" >= ("a" + "foo bar"))) {
                        var v29 float32 = 1e10
                        _ = v29
                    }
                }
                if v5 {
                    var v30 int8 = (-(v1))
                    v2 = 127
                    var v31 int16 = 0
                    _ = v30
                    _ = v31
                }
            } else {
                if (42 <= 42) {
                    var v32 uint32 = (v22 + 2147483647)
                    fmt.Println(v32, v2)
                    _ = v32
                } else {
                    _ = len(v4)
                    v33 := GenS3{F0: -1, F1: (v1 & 42)}
                    var v34 float64 = 1e10
                    _ = v33
                    _ = v34
                }
            }
            switch "\n\t" {
            case "foo bar":
                switch 32768 {
                case 127:
                    var v35 int = 4294967295
                    _ = v35
                case -1:
                    v4[("a" + "")] = -533
                case 128:
                    v36 := GenS5{F0: 18446744073709551615, F1: (v2 - v3), F2: "\n\t"}
                    _ = v36
                }
            case "hello":
                _ = len(v4)
            case "":
                var v37 float32 = float32((0.1 * 3.14159))
                _ = v37
            default:
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 string = ""
                    v40 := GenS6{F0: 869}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
            var v41 uint = 4294967295
            _ = v41
        }
        v5 = ((false || (true && v5)) && v5)
    }
    if v5 {
        for fi42 := 0; fi42 < 13; fi42++ {
            for fi43 := 0; fi43 < 20; fi43++ {
                v44 := GenS4{F0: -32768}
                _ = fi43
                _ = v44
            }
            if (42 != 1) {
                v4["\n\t"] = (fi42 ^ -485)
                v1 -= v3
                v5 = true
            } else {
                if false {
                    var v45 int16 = 296
                    switch (^(fi42)) {
                    case -98:
                        switch (-(fi42)) {
                        case 1:
                            v46 := GenS3{F0: (v45 - (v45 | (v45 + (-(v45))))), F1: v1}
                            _ = v46
                        }
                    }
                    _ = v45
                }
                for fi47 := 0; fi47 < 20; fi47++ {
                    var v48 int64 = 778
                    var v49 float32 = -2.25
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                switch (fi42 >> ((284) & 63)) {
                case 174:
                    fmt.Println(fi42)
                case 657:
                    for fi50 := 0; fi50 < 10; fi50++ {
                        var v51 bool = (1e10 == (-2.25 + (0.0 - 0.1)))
                        fmt.Println(v51, v51)
                        _ = fi50
                        _ = v51
                    }
                default:
                    for fi52 := 0; fi52 < 10; fi52++ {
                        switch fi42 {
                        case -331:
                            v5 = (v1 != v1)
                        }
                        _ = fi52
                    }
                }
            }
            _ = fi42
        }
        var v53 int = 9223372036854775807
        switch (("foo bar" + (("\n\t" + "hello") + ("a" + (((("foo bar" + "hello") + ("foo bar" + "")) + "") + ("" + (("foo bar" + (("\n\t" + (("foo bar" + (("\n\t" + "") + "hello")) + "hello")) + ("hello" + ""))) + ((("foo bar" + ("\n\t" + "foo bar")) + ("\n\t" + "a")) + "foo bar"))))))) + "\n\t") {
        case "":
            var v54 int32 = -393
            _ = v54
        }
        _ = v53
    } else {
        if (-32768 > 42) {
            switch -1 {
            case -9223372036854775808:
                var v55 float64 = float64(((1e10 / ((3.14159 * 1e10) + 0.1)) / 1.5))
                _ = v55
            case 1:
                for fi56 := 0; fi56 < 4; fi56++ {
                    var v57 int = fi56
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v6)
                    default:
                        var v58 int16 = -1
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 8; fi59++ {
                        switch "\n\t" {
                        case "foo bar":
                            v17 = v17
                        case "\n\t":
                            switch "hello" {
                            case "hello":
                                v3 = 0
                            default:
                                var v60 int = (fi56 * (fi56 * 2147483647))
                                _ = v60
                            }
                        case "":
                            var v61 uint64 = 1
                            _ = v61
                        }
                        _ = fi59
                    }
                    _ = fi56
                    _ = v57
                }
            case -576:
                var v62 bool = false
                _ = v62
            default:
                if ((true || false) || (!v5)) {
                    var v63 int32 = -140
                    _ = v63
                }
            }
            var v64 int8 = (v3 | v3)
            _ = v64
        }
        fmt.Println(v1, v5, v2)
    }
    v4[("\n\t" + "foo bar")] = 402
    var v65 uint8 = v17
    v5 = (v5 || (v5 && false))
    v65 = v17
    switch "" {
    case "\n\t":
        switch ("hello" + (("foo bar" + "\n\t") + ("" + "foo bar"))) {
        case "foo bar":
            _ = len(v6)
        case "\n\t":
            fmt.Println(v3, v65)
        case "":
            _ = len(v6)
        }
    default:
        v66 := GenS2{F0: -1}
        _ = v66
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        switch "a" {
        case "hello":
            v1 -= -128
        default:
            v22 = v22
        }
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
}
