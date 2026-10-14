package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 string; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
func genFn0() uint32 {
    v5 := GenS1{F0: 336}
    v6 := GenS1{F0: 2147483647}
    var v7 uint64 = 0
    _ = v5
    _ = v6
    _ = v7
    return 980
}
func genFn1(p0 int8, p1 int) int8 {
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{17}
        v10 := GenS1{F0: (fi8 | (p1 & 32768))}
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v11 := GenS1{F0: (p1 | 1)}
        _ = v11
    }
    v12 := []int{(p1 >> ((2) & 63)), (p1 | p1)}
    p0 = (p0 | (p0 * p0))
    _ = p0
    _ = p1
    _ = v12
    return 75
}

func main() {
    v1 := []uint32{968}
    v2 := []uint32{1, 353}
    v3 := make(map[string]bool)
    var v4 int = 1
    if true {
        if (2147483647 < 598) {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int8 = 42
                _ = fi13
                _ = v14
            }
        } else {
            v4 -= v4
            if (289 > 122) {
                var v15 float64 = 1e10
                for fi16 := 0; fi16 < 1; fi16++ {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 int16 = -376
                        var v19 int16 = -1
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v4 -= (v4 | (fi16 | 128))
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if true {
                            var v21 uint8 = 226
                            v21 = v21
                            _ = v21
                        }
                        if false {
                            if true {
                                var v22 int32 = 1
                                _ = v22
                            }
                            if false {
                                if ((false && true) && true) {
                                    fmt.Println(v4, v15)
                                }
                                if false {
                                    var v23 int32 = -2147483648
                                    var v24 uint16 = 1
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 uint32 = 65535
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v4)
                                    case "":
                                        v15 = ((1.5 / float64(0.1)) - (0.0 / 1.0))
                                    }
                                    if (!(0.1 >= (float32(1e10) + (0.1 * (0.1 + -2.25))))) {
                                        if (-992 > -2147483648) {
                                            v4 += (v4 - (^(fi16)))
                                            v26 := GenS1{F0: 9223372036854775807}
                                            _ = v26
                                        } else {
                                            v4 -= fi16
                                            var v27 uint8 = 86
                                            if false {
                                                var v28 string = "foo bar"
                                                if false {
                                                    v4 += -296
                                                    var v29 int64 = -9223372036854775808
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v30 int32 = 31
                                                        v31 := GenS1{F0: (fi16 % ((v4) | 1))}
                                                        v32 := GenS2{F0: v30, F1: 42, F2: (fi16 * (-(fi16)))}
                                                        v33 := GenS3{F0: 44, F1: "foo bar", F2: 32767}
                                                        var v34 bool = (true || (!true))
                                                        _ = len(v1)
                                                        var v35 int64 = (v29 << ((87) & 63))
                                                        v36 := GenS2{F0: int32(v27), F1: 42, F2: fi20}
                                                        var v37 int64 = v29
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    _ = v29
                                                }
                                                var v38 uint32 = v25
                                                _ = v28
                                                _ = v38
                                            }
                                            _ = v27
                                        }
                                        var v39 string = "a"
                                        v4 += (v4 % (((fi20 / ((fi20) | 1))) | 1))
                                        v25 = v25
                                        v25 = v25
                                        _ = v39
                                    }
                                    _ = v25
                                }
                                v4 = -624
                            } else {
                                v4 -= 0
                                fmt.Println(v4)
                            }
                        } else {
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v40 float32 = 1.5
                                    _ = v40
                                } else {
                                    switch (("a" + "\n\t") + "\n\t") {
                                    case "\n\t":
                                        switch -9223372036854775808 {
                                        case -545:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 2147483647:
                                            if (true || (938 <= -9223372036854775808)) {
                                                var v41 int8 = 42
                                                fmt.Println(fi20)
                                                _ = v41
                                            } else {
                                                v4 -= fi20
                                                if (820 > 2147483647) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v15 = v15
                                                    fmt.Println(v4, fi20, fi20)
                                                }
                                                if ((!(65535 == 909)) && (931 >= 255)) {
                                                    _ = len(v1)
                                                    var v42 int8 = genFn1(-118, fi20)
                                                    v3["\n\t"] = ((1e10 / (v15 - v15)) < 0.0)
                                                    _ = v42
                                                } else {
                                                    v43 := GenS3{F0: 64, F1: "foo bar", F2: (fi16 & (fi20 | (^(fi16))))}
                                                    v15 += 1.5
                                                    _ = v43
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        var v44 uint64 = 171
                                        _ = v44
                                    case "hello":
                                        v45 := GenS3{F0: 100, F1: "hello", F2: 618}
                                        _ = v45
                                    default:
                                        if false {
                                            switch (fi16 << ((322) & 63)) {
                                            case 309:
                                                switch "hello" {
                                                case "\n\t":
                                                    _ = len(v2)
                                                }
                                            case 23:
                                                var v46 string = ("foo bar" + ("" + "\n\t"))
                                                _ = v46
                                            }
                                            var v47 int64 = 0
                                            switch (("hello" + ("foo bar" + ((("a" + ((((("foo bar" + "") + "foo bar") + ("" + "")) + (("a" + "hello") + "a")) + "\n\t")) + (("hello" + ("\n\t" + ("a" + "foo bar"))) + "a")) + (((("a" + ("hello" + "a")) + "\n\t") + "") + "hello")))) + "hello") {
                                            case "hello":
                                                v4 += v4
                                            case "":
                                                if false {
                                                    v48 := GenS2{F0: -1, F1: 42, F2: (fi16 % ((fi20) | 1))}
                                                    switch "a" {
                                                    case "hello":
                                                        var v49 uint32 = 2147483648
                                                        _ = v49
                                                    case "a":
                                                        switch fi16 {
                                                        case 1:
                                                            v3[(("foo bar" + "\n\t") + "hello")] = ((0 == 4294967295) && false)
                                                        case 127:
                                                            switch (fi16 | 0) {
                                                            case -331:
                                                                v15 = (v15 + v15)
                                                            case 32768:
                                                                fmt.Println(v47)
                                                            }
                                                        }
                                                    }
                                                    _ = v48
                                                }
                                            }
                                            _ = v47
                                        }
                                    }
                                    var v50 uint64 = 4294967295
                                    var v51 int64 = 0
                                    _ = v50
                                    _ = v51
                                }
                            default:
                                if true {
                                    fmt.Println(fi20, fi20)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi20, fi16, v15)
                                } else {
                                    if (-1 > -686) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if false {
                                            fmt.Println(fi20, v4)
                                            var v52 float64 = float64(3.14159)
                                            _ = v52
                                        }
                                        v4 += -2147483648
                                    } else {
                                        switch 984 {
                                        case 128:
                                            if len(v2) > 0 {
                                                v2[0] = 947
                                            }
                                        case -2147483648:
                                            if len(v1) > 0 {
                                                v1[0] = genFn0()
                                            }
                                        case 2147483647:
                                            var v53 int8 = genFn1(-1, 32767)
                                            _ = v53
                                        }
                                        v4 -= fi20
                                        if false {
                                            var v54 uint = 42
                                            _ = v54
                                        }
                                    }
                                    v15 += float64(1.5)
                                }
                            }
                            v55 := GenS3{F0: 1, F1: "a", F2: v4}
                            v15 += v15
                            _ = v55
                        }
                        _ = fi20
                    }
                    _ = fi16
                }
                var v56 int = (v4 * v4)
                _ = v15
                _ = v56
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        fmt.Println(v4)
        for fi57 := 0; fi57 < 18; fi57++ {
            if (4294967295 < 1) {
                var v58 int32 = 2147483647
                v4 = (fi57 | v4)
                _ = v58
            }
            fmt.Println(v4, fi57)
            var v59 uint8 = 150
            _ = fi57
            _ = v59
        }
    }
    fmt.Println(v4)
    v60 := GenS4{F0: 0}
    v4 -= v4
    v4 -= v4
    for fi61 := 0; fi61 < 19; fi61++ {
        if (!true) {
            v4 += -9223372036854775808
        }
        v62 := GenS5{F0: 42}
        _ = fi61
        _ = v62
    }
    var v63 uint32 = 565
    v4 = v4
    for fi64 := 0; fi64 < 12; fi64++ {
        v65 := GenS6{F0: 4294967295}
        v63 = v63
        switch fi64 {
        case 127:
            if true {
                v63 = 2
                v4 = fi64
            }
        case 229:
            switch 128 {
            case 130:
                var v66 int16 = -179
                _ = v66
            case 1:
                fmt.Println(v4, v63)
            }
        }
        _ = fi64
        _ = v65
    }
    var v67 float32 = (3.14159 * ((-2.25 - 0.1) + (1.5 + float32(1.5))))
    v68 := GenS5{F0: 9223372036854775807}
    for fi69 := 0; fi69 < 6; fi69++ {
        v4 += 82
        v63 = v63
        switch ((("a" + "\n\t") + (((("" + "") + "foo bar") + ("a" + (((("a" + "hello") + "foo bar") + ("" + (("foo bar" + ((("" + "") + "hello") + ("a" + ""))) + (("a" + ("a" + "")) + "a")))) + "hello"))) + ("\n\t" + "foo bar"))) + "a") {
        case "foo bar":
            var v70 int = -9223372036854775808
            _ = v70
        case "a":
            if ((v4 >= v4) || false) {
                v71 := GenS3{F0: 0, F1: "", F2: fi69}
                _ = v71
            }
        case "hello":
            v63 += v63
        }
        _ = fi69
    }
    for fi72 := 0; fi72 < 17; fi72++ {
        v67 -= -2.25
        for fi73 := 0; fi73 < 5; fi73++ {
            fmt.Println(v63, v63, v67)
            v74 := GenS3{F0: 127, F1: "", F2: (v4 - (fi73 | (fi73 >> ((255) & 63))))}
            var v75 uint = 255
            _ = fi73
            _ = v74
            _ = v75
        }
        _ = fi72
    }
    if (56 != 255) {
        if len(v1) > 0 {
            v1[0] = (v63 + (v63 / ((768) | 1)))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v63
    _ = v67
    _ = v68
}
