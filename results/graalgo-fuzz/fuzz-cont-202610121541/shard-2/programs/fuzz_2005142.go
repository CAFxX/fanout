package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 string; F1 bool; F2 string }
func genFn0(p0 uint8) string {
    v7 := GenS1{F0: ""}
    if false {
        p0 -= p0
        if ((!(("\n\t" < ("foo bar" + "hello")) || true)) || true) {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint8 = (p0 ^ 0)
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 uint16 = 0
                    var v12 uint16 = 1
                    v13 := GenS1{F0: ("a" + "hello")}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 bool = true
                _ = fi8
                _ = v14
                _ = v9
            }
        } else {
            var v15 int8 = -91
            v16 := GenS2{F0: 32767}
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint64 = 42
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    switch int(p0) {
                    case 128:
                        var v19 float64 = float64(float32(((1e10 - 0.0) - 0.0)))
                        _ = v19
                    case -1:
                        fmt.Println(p0, v15)
                    case -2147483648:
                        if false {
                            v15 -= v15
                        }
                    default:
                        v20 := []int{fi17, -748, (fi17 + (fi17 | (-(fi17)))), fi17}
                        _ = v20
                    }
                case "a":
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v18 = (v18 & 0)
                        v18 = 2
                        _ = fi21
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
    } else {
        switch "hello" {
        case "hello":
            var v22 uint32 = 0
            _ = v22
        case "\n\t":
            if true {
                if false {
                    switch "foo bar" {
                    case "":
                        p0 = 2
                    }
                } else {
                    p0 = (p0 << ((uint(p0)) & 63))
                    v23 := []int{}
                    v24 := GenS1{F0: ("" + "a")}
                    _ = v23
                    _ = v24
                }
                switch ("" + ("" + (((((("a" + "\n\t") + "\n\t") + "foo bar") + "") + (("\n\t" + ("\n\t" + "foo bar")) + "hello")) + ""))) {
                case "\n\t":
                    v25 := []int{9223372036854775807, -232, -824}
                    _ = v25
                case "a":
                    switch 532 {
                    case 2147483647:
                        var v26 int = -950
                        _ = v26
                    case -409:
                        switch (("\n\t" + (("hello" + "\n\t") + "\n\t")) + ("\n\t" + ("hello" + ((("" + ("" + "hello")) + "hello") + "")))) {
                        case "a":
                            p0 += p0
                        default:
                            v27 := GenS3{F0: 1}
                            _ = v27
                        }
                    case 1:
                        var v28 int32 = 385
                        _ = v28
                    default:
                        p0 += 127
                    }
                }
                switch "hello" {
                case "":
                    v29 := GenS1{F0: "\n\t"}
                    _ = v29
                }
            }
        case "":
            for fi30 := 0; fi30 < 14; fi30++ {
                v31 := GenS4{F0: 18446744073709551615, F1: "\n\t"}
                v32 := GenS2{F0: (fi30 | (fi30 >> ((4294967295) & 63)))}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    }
    if (1 <= 175) {
        switch "\n\t" {
        case "":
            for fi33 := 0; fi33 < 9; fi33++ {
                p0 += (^(p0))
                var v34 float64 = ((-2.25 * (3.14159 * 1e10)) * -2.25)
                fmt.Println(v34)
                _ = fi33
                _ = v34
            }
        case "\n\t":
            for fi35 := 0; fi35 < 16; fi35++ {
                switch ("foo bar" + "") {
                case "foo bar":
                    switch ("foo bar" + "hello") {
                    case "a":
                        var v36 float64 = float64(3.14159)
                        _ = v36
                    case "hello":
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    case "\n\t":
                        fmt.Println(fi35, fi35)
                    default:
                        switch -265 {
                        case 32768:
                            fmt.Println(p0, p0)
                        case -462:
                            if false {
                                fmt.Println(fi35)
                                var v38 bool = ((3.14159 + ((-2.25 - float32(0.0)) + ((3.14159 * -2.25) + 1.5))) < 1e10)
                                _ = v38
                            }
                        }
                    }
                case "\n\t":
                    p0 = p0
                default:
                    var v39 uint = 457
                    _ = v39
                }
                for fi40 := 0; fi40 < 14; fi40++ {
                    switch (-(fi40)) {
                    case 288:
                        v41 := GenS5{F0: 242}
                        _ = v41
                    }
                    if false {
                        fmt.Println(fi35)
                    } else {
                        v42 := []int{(^(fi35)), -1}
                        _ = v42
                    }
                    _ = fi40
                }
                _ = fi35
            }
        }
    } else {
        p0 = (p0 % (((^(p0))) | 1))
        p0 -= 111
        v43 := []int{32767}
        _ = v43
    }
    _ = p0
    _ = v7
    return ("hello" + "\n\t")
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    var v3 int16 = 1000
    var v4 uint32 = 700
    var v5 float32 = (1e10 / 1.0)
    var v6 bool = (2 != 42)
    var v44 int = int(v1)
    var v45 int64 = (^(v1))
    if true {
        v46 := v2[(((("foo bar" + "a") + ("hello" + "\n\t")) + (("" + "\n\t") + "\n\t")) + "\n\t")]
        v3 -= 42
        _ = v46
    } else {
        var v47 int8 = 1
        fmt.Println(v44, v3, v3)
        for fi48 := 0; fi48 < 17; fi48++ {
            if v6 {
                v49 := GenS5{F0: 65535}
                v1 = (v45 * (v45 + v1))
                _ = v49
            }
            _ = fi48
        }
        _ = v47
    }
    v5 -= (v5 + 1e10)
    fmt.Println(v4)
    if (1 != 255) {
        var v50 uint32 = 2147483647
        _ = v50
    } else {
        v51 := GenS5{F0: 579}
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        _ = len(v2)
        v53 := GenS3{F0: v1}
        v6 = v6
        _ = fi52
        _ = v53
    }
    var v54 string = "\n\t"
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 uint32 = v4
        _ = len(v2)
        _ = fi55
        _ = v56
    }
    v6 = v6
    for fi57 := 0; fi57 < 13; fi57++ {
        if (((1.5 + -2.25) == 1.5) && (v54 == ("foo bar" + v54))) {
            if v6 {
                var v58 uint32 = (v4 & (v4 >> ((42) & 63)))
                _ = v58
            } else {
                for fi59 := 0; fi59 < 20; fi59++ {
                    var v60 uint64 = 769
                    switch (v44 + fi57) {
                    case -747:
                        v61 := GenS6{F0: (v54 + v54)}
                        _ = v61
                    }
                    if (!(!(!v6))) {
                        if (-9223372036854775808 >= (v44 - v44)) {
                            v2[v54] = fi59
                            v62 := GenS3{F0: v1}
                            _ = v62
                        } else {
                            v63 := []int{}
                            _ = v63
                        }
                        v6 = v6
                    }
                    _ = fi59
                    _ = v60
                }
            }
            for fi64 := 0; fi64 < 18; fi64++ {
                var v65 int8 = 1
                v66 := GenS7{F0: (v3 / ((42) | 1)), F1: (v65 ^ -115)}
                v67 := v2[((v54 + ((v54 + "hello") + "\n\t")) + ("" + "\n\t"))]
                var v68 float64 = 3.14159
                _ = fi64
                _ = v65
                _ = v66
                _ = v67
                _ = v68
            }
        } else {
            v69 := GenS5{F0: 42}
            switch ((v54 + (v54 + "hello")) + v54) {
            case "hello":
                for fi70 := 0; fi70 < 9; fi70++ {
                    fmt.Println(v45)
                    delete(v2, v54)
                    var v71 uint = 995
                    _ = fi70
                    _ = v71
                }
            case "":
                var v72 uint16 = 367
                _ = v72
            case "foo bar":
                for fi73 := 0; fi73 < 3; fi73++ {
                    var v74 uint64 = 459
                    v75 := []int{(v44 | 32767), (v44 << ((185) & 63)), fi73}
                    v45 = (v1 * (-(v1)))
                    if v6 {
                        v76 := GenS8{F0: v54, F1: (!(v6 || true)), F2: "\n\t"}
                        var v77 bool = v6
                        if true {
                            v78 := GenS2{F0: 128}
                            fmt.Println(v4, v74)
                            switch v54 {
                            case "foo bar":
                                fmt.Println(v5)
                            default:
                                var v79 int8 = 1
                                _ = v79
                            }
                            _ = v78
                        } else {
                            v45 -= v1
                            if len(v75) >= 2 {
                                v75 = v75[1:len(v75)]
                            }
                        }
                        _ = v76
                        _ = v77
                    }
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            default:
                if false {
                    var v80 int64 = v45
                    _ = v80
                } else {
                    v54 = "a"
                    v5 = v5
                    if v6 {
                        if (4294967295 < 1) {
                            fmt.Println(v4, v3)
                            fmt.Println(v3, v54, v5)
                            if (v6 || v6) {
                                for fi81 := 0; fi81 < 4; fi81++ {
                                    if (v6 && (v4 < v4)) {
                                        fmt.Println(v44, v5)
                                    } else {
                                        fmt.Println(fi57, v45, v3)
                                    }
                                    fmt.Println(v4, v3, v3)
                                    _ = fi81
                                }
                                v54 = v54
                            }
                        }
                        if false {
                            fmt.Println(v54, v4, v54)
                            var v82 uint = 487
                            _ = v82
                        }
                    }
                }
            }
            _ = v69
        }
        fmt.Println(fi57)
        var v83 bool = false
        _ = fi57
        _ = v83
    }
    v44 -= 9223372036854775807
    var v84 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v54
    _ = v6
    _ = v84
}
