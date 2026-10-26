package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint; F1 string; F2 string }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 int16) uint64 {
    p0 -= (p0 / ((p0) | 1))
    v5 := GenS1{F0: 944, F1: false, F2: 0}
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v6, fi7, p0)
        fmt.Println(v6, fi7)
        var v8 string = ("hello" + ("a" + ("a" + "\n\t")))
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    var v4 float32 = float32(0.1)
    v3 += -41
    var v9 uint32 = 2
    for fi10 := 0; fi10 < 10; fi10++ {
        switch "hello" {
        case "hello":
            var v11 uint16 = 205
            _ = v11
        case "\n\t":
            v12 := GenS1{F0: 1, F1: false, F2: -9223372036854775808}
            _ = v12
        case "foo bar":
            v9 -= 2147483647
        default:
            if (!(-9223372036854775808 <= 353)) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v4 = v4
                    _ = len(v2)
                    if len(v1) > 0 {
                        v1[0] = (fi13 ^ fi13)
                    }
                    _ = fi13
                }
            }
        }
        if true {
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS2{F0: "a", F1: "a"}
                var v16 int16 = 272
                _ = fi14
                _ = v15
                _ = v16
            }
            v4 = 3.14159
            v17 := GenS3{F0: 2, F1: "hello", F2: "a"}
            _ = v17
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v18 float32 = (float32(0.0) - v4)
            _ = v18
        }
        _ = fi10
    }
    switch 225 {
    case -2147483648:
        var v19 int16 = 32767
        _ = v19
    case 96:
        var v20 bool = (18446744073709551615 <= genFn0(492))
        _ = v20
    case 1:
        if true {
            if (true && true) {
                fmt.Println(v9, v3, v9)
                v9 = 615
            }
        } else {
            for fi21 := 0; fi21 < 7; fi21++ {
                if len(v1) > 0 {
                    v1[0] = 127
                }
                v9 -= v9
                _ = fi21
            }
        }
    default:
        switch -171 {
        case 9223372036854775807:
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(v4, v3, v4)
                _ = fi22
            }
        case 32768:
            v9 = v9
        default:
            switch ("hello" + (((("hello" + "hello") + ("hello" + "hello")) + ("a" + "")) + "foo bar")) {
            case "a":
                v3 = (v3 | 42)
            case "":
                var v23 int16 = -444
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 bool = ((-(fi24)) < fi24)
            v27 := GenS3{F0: 65535, F1: ("foo bar" + "foo bar"), F2: ((("foo bar" + (("foo bar" + "\n\t") + ("a" + "foo bar"))) + "foo bar") + "foo bar")}
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int = -1
        switch "\n\t" {
        case "a":
            fmt.Println(v28, v28)
        case "hello":
            delete(v2, "")
        case "":
            fmt.Println(v9)
        }
        _ = fi24
        _ = v28
    }
    v9 = 65535
    for fi29 := 0; fi29 < 10; fi29++ {
        v3 -= 42
        if (v4 == -2.25) {
            v9 = v9
            var v30 float32 = (v4 * 0.1)
            _ = v30
        } else {
            var v31 int8 = v3
            for fi32 := 0; fi32 < 7; fi32++ {
                v2[("\n\t" + "\n\t")] = (v9 & (^(v9)))
                switch "hello" {
                case "":
                    switch "hello" {
                    case "":
                        v33 := GenS4{F0: 18446744073709551615, F1: 193}
                        _ = v33
                    }
                default:
                    if (32767 < -336) {
                        fmt.Println(fi32, v31, v3)
                        switch fi32 {
                        case 972:
                            switch (("\n\t" + "foo bar") + "hello") {
                            case "\n\t":
                                var v34 bool = false
                                _ = v34
                            case "a":
                                var v35 float32 = v4
                                _ = v35
                            default:
                                v4 += 0.1
                            }
                        case 128:
                            var v36 int8 = v31
                            _ = v36
                        case -178:
                            var v37 int16 = 1000
                            _ = v37
                        }
                        var v38 uint64 = 1
                        _ = v38
                    } else {
                        fmt.Println(v4)
                        switch (-(fi32)) {
                        case 127:
                            v39 := GenS4{F0: genFn0(-283), F1: genFn0(-1)}
                            _ = v39
                        case 295:
                            v40 := v2["\n\t"]
                            _ = v40
                        case 32767:
                            delete(v2, "a")
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v41 int32 = -1
                        _ = v41
                    }
                }
                v4 -= v4
                _ = fi32
            }
            _ = v31
        }
        _ = fi29
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        if (!((false || (!((("foo bar" + (("\n\t" + "a") + ("" + "a"))) + "hello") >= "a"))) || false)) {
            var v43 int8 = v3
            v3 += v3
            v4 -= v4
            _ = v43
        }
        if (722 == 0) {
            var v44 float64 = 1.5
            _ = v44
        }
        _ = fi42
    }
    v45 := v2[("foo bar" + ("foo bar" + ("\n\t" + "a")))]
    v46 := v2["\n\t"]
    v45 = v46
    if true {
        if (true && (!false)) {
            for fi47 := 0; fi47 < 19; fi47++ {
                v48 := v2[""]
                for fi49 := 0; fi49 < 17; fi49++ {
                    var v50 bool = false
                    v51 := GenS5{F0: 0, F1: 255}
                    if (((v45 % ((v48) | 1)) <= v9) && (!(65535 < v48))) {
                        v3 += (v3 + (v3 & v3))
                    }
                    v46 += (v9 * v9)
                    switch (fi47 + fi49) {
                    case 42:
                        var v52 uint8 = 42
                        _ = v52
                    case 0:
                        switch fi49 {
                        case 128:
                            var v53 int32 = 1
                            _ = v53
                        case 2147483647:
                            v50 = (v50 || v50)
                        case 9223372036854775807:
                            fmt.Println(fi47, v45)
                        default:
                            if v50 {
                                if v50 {
                                    if false {
                                        if (255 >= 1) {
                                            if v50 {
                                                var v54 int8 = (v3 + -128)
                                                if ((2147483647 <= (fi49 * fi47)) || (39 != -80)) {
                                                    var v55 int16 = 32767
                                                    v45 += 42
                                                    _ = v55
                                                }
                                                _ = v54
                                            } else {
                                                if v50 {
                                                    v2["foo bar"] = (v45 % (((^(v48))) | 1))
                                                    var v56 float32 = v4
                                                    if len(v1) > 0 {
                                                        v1[0] = fi49
                                                    }
                                                    _ = v56
                                                } else {
                                                    v50 = (!(false && false))
                                                    switch (fi47 / (((fi47 & fi49)) | 1)) {
                                                    case 42:
                                                        switch fi47 {
                                                        case 2147483647:
                                                            v57 := GenS5{F0: 9223372036854775807, F1: 255}
                                                            _ = v57
                                                        default:
                                                            delete(v2, ((("a" + "a") + "\n\t") + ("foo bar" + "a")))
                                                        }
                                                    case 158:
                                                        _ = len(v2)
                                                    case 32767:
                                                        var v58 bool = v50
                                                        _ = v58
                                                    }
                                                    switch (fi49 - (^(fi47))) {
                                                    case 32767:
                                                        v46 -= v48
                                                    }
                                                }
                                                var v59 float32 = ((v4 / v4) * (1e10 * (v4 - ((1.5 + v4) + (v4 + v4)))))
                                                var v60 float64 = 1e10
                                                _ = v59
                                                _ = v60
                                            }
                                            if v50 {
                                                v3 += v3
                                                var v61 int32 = 42
                                                _ = v61
                                            }
                                            v50 = (-2.25 >= ((0.1 - float64(v4)) / (0.1 / 0.1)))
                                        } else {
                                            v62 := GenS6{F0: 128}
                                            _ = v62
                                        }
                                        fmt.Println(v50, fi49)
                                    }
                                    var v63 float64 = 0.0
                                    v45 = v48
                                    _ = v63
                                }
                                if (!(!((33 < (v3 * (v3 * 1))) || (false && (298 == 1))))) {
                                    if (-31 >= 304) {
                                        var v64 int32 = 42
                                        _ = v64
                                    }
                                    switch fi47 {
                                    case 2147483647:
                                        if (!v50) {
                                            var v65 int = 128
                                            v65 = (fi47 * fi47)
                                            _ = v65
                                        } else {
                                            v9 -= (v45 << ((255) & 63))
                                            v66 := GenS5{F0: -9223372036854775808, F1: 255}
                                            _ = v66
                                        }
                                    case 32768:
                                        switch "a" {
                                        case "":
                                            v50 = (42 == 42)
                                        case "a":
                                            v67 := GenS3{F0: uint(v9), F1: (((("a" + "foo bar") + "\n\t") + "foo bar") + ("foo bar" + "\n\t")), F2: (("" + ("\n\t" + "a")) + "\n\t")}
                                            var v68 bool = (524 > -1)
                                            v69 := GenS7{F0: 597}
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        default:
                                            v4 -= v4
                                        }
                                    default:
                                        var v70 bool = true
                                        _ = v70
                                    }
                                } else {
                                    v71 := GenS8{F0: 255, F1: 958, F2: "\n\t"}
                                    v45 = (^(v9))
                                    _ = len(v1)
                                    delete(v2, ("a" + ("hello" + "foo bar")))
                                    v72 := GenS8{F0: 32767, F1: 442, F2: (("hello" + ((("foo bar" + (("" + "") + ("hello" + ("foo bar" + "a")))) + "hello") + "a")) + "a")}
                                    v73 := v2["\n\t"]
                                    if (v50 || ((-2.25 * 0.0) > float64(float32(0.1)))) {
                                        if len(v1) > 0 {
                                            v1[0] = 42
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = fi47
                                        }
                                        var v74 string = ""
                                        v75 := GenS6{F0: 42}
                                        fmt.Println(v45)
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                            }
                        }
                    case 128:
                        v2["\n\t"] = v9
                    }
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                _ = fi47
                _ = v48
            }
            var v76 bool = ("a" != ("" + ("hello" + "hello")))
            fmt.Println(v76)
            _ = v76
        }
        v2[("a" + ("a" + "a"))] = 127
        v46 = v9
    } else {
        var v77 uint16 = 255
        _ = v77
    }
    v4 -= v4
    switch "\n\t" {
    case "foo bar":
        switch "hello" {
        case "hello":
            v2["\n\t"] = v46
        case "\n\t":
            v2["\n\t"] = 1
        }
    case "hello":
        switch 263 {
        case 4294967295:
            var v78 int16 = 42
            _ = v78
        case -2147483648:
            var v79 int32 = -1
            _ = v79
        }
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v9
}
