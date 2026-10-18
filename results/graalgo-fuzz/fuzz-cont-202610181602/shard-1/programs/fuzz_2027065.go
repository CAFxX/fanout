package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 bool; F1 int; F2 uint16 }
func genFn0(p0 bool, p1 uint8, p2 uint8) string {
    p1 = (p2 ^ (p1 / ((p2) | 1)))
    v7 := GenS1{F0: (!p0), F1: p2}
    var v8 uint16 = 0
    p0 = (-809 < 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint8) int16 {
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(p0, fi9)
        var v10 int8 = 1
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 1
    v12 := GenS1{F0: (468 < 65535), F1: (p0 ^ p0)}
    _ = p0
    _ = v11
    _ = v12
    return 32767
}

func main() {
    v1 := []string{("foo bar" + "a"), "foo bar", ""}
    var v2 uint64 = 1
    v3 := make(map[string]int)
    var v4 int64 = 42
    var v5 int16 = -32768
    var v6 int16 = 42
    if (2 != 1) {
        v4 = v4
        if true {
            v13 := GenS2{F0: (v4 | -1)}
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 string = ("a" + "foo bar")
                _ = fi14
                _ = v15
            }
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            _ = v13
        }
    } else {
        if (-2.25 <= (3.14159 * 0.1)) {
            switch (("\n\t" + "foo bar") + "a") {
            case "hello":
                v4 = v4
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                v1 = append(v1, "hello")
                v5 = (v5 ^ genFn1(203))
                _ = fi16
            }
            var v17 int8 = -1
            _ = v17
        } else {
            v5 -= v5
            v2 = v2
            _ = len(v3)
        }
        var v18 uint16 = 14
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint8 = 127
            switch "a" {
            case "a":
                var v21 uint16 = 0
                _ = v21
            case "foo bar":
                var v22 uint = 0
                _ = v22
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v23 int16 = -32768
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                if false {
                    v6 = (^(v5))
                    var v25 int16 = (v6 / ((v5) | 1))
                    var v26 uint16 = 927
                    switch 1 {
                    case 321:
                        v6 = genFn1(147)
                        v27 := GenS2{F0: v4}
                        var v28 float64 = float64((1e10 + -2.25))
                        fmt.Println(v26, v2, v5)
                        var v29 int16 = (^(v5))
                        _ = v27
                        _ = v28
                        _ = v29
                    case -847:
                        var v30 int32 = -814
                        _ = v30
                    default:
                        v1 = append(v1, "\n\t")
                    }
                    fmt.Println(v4, v6, v6)
                    _ = v25
                    _ = v26
                } else {
                    v31 := GenS3{F0: -113, F1: -93}
                    if ((v2 % ((v2) | 1)) == v2) {
                        if (true && false) {
                            v4 -= (v4 / ((-252) | 1))
                            v4 += (v4 / ((-9223372036854775808) | 1))
                        }
                        _ = len(v1)
                    }
                    _ = v31
                }
                switch ((("\n\t" + ("hello" + ("" + "hello"))) + "a") + "\n\t") {
                case "":
                    var v32 int64 = 0
                    _ = v32
                case "hello":
                    var v33 string = ("foo bar" + "foo bar")
                    _ = v33
                case "foo bar":
                    if true {
                        var v34 uint32 = 0
                        v35 := GenS2{F0: (v4 * (v4 % (((v4 << ((0) & 63))) | 1)))}
                        v36 := GenS3{F0: 100, F1: 74}
                        if (1 == -1) {
                            switch (("\n\t" + "foo bar") + "") {
                            case "":
                                if false {
                                    var v37 string = "a"
                                    var v38 int32 = 2147483647
                                    delete(v3, v37)
                                    _ = v37
                                    _ = v38
                                }
                            case "hello":
                                v3[""] = 128
                            case "\n\t":
                                var v39 uint8 = 1
                                _ = v39
                            }
                        }
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        var v40 bool = (((("foo bar" + ("a" + "a")) + ("" + "")) > ("" + "\n\t")) && ((1.5 + (float64(0.0) / 3.14159)) > (((((1.5 - 1e10) * (float64(0.1) * 3.14159)) - (((1.5 * 1e10) + 1.5) + float64(3.14159))) / (0.0 + ((0.0 * (((1e10 / (((-2.25 * (1.5 * 3.14159)) + (float64((0.0 / (-2.25 / (((0.0 + 0.1) / 1.5) / 3.14159)))) * 0.1)) + 1.5)) / 0.1) * 3.14159)) - 3.14159))) - ((-2.25 * 1e10) * (0.1 / 1.0)))))
                        _ = v40
                    }
                default:
                    var v41 string = ("\n\t" + ("\n\t" + "a"))
                    _ = v41
                }
                if true {
                    v2 = (v2 ^ v2)
                    var v42 uint64 = 360
                    _ = v42
                }
                _ = fi24
            }
            _ = fi19
            _ = v20
        }
        var v43 uint8 = 151
        _ = v43
    } else {
        for fi44 := 0; fi44 < 13; fi44++ {
            for fi45 := 0; fi45 < 3; fi45++ {
                var v46 uint16 = 2
                fmt.Println(v2, fi44)
                if len(v1) > 0 {
                    v1[0] = ("" + "foo bar")
                }
                v2 = v2
                var v47 uint = 18446744073709551615
                _ = fi45
                _ = v46
                _ = v47
            }
            var v48 int8 = 0
            _ = fi44
            _ = v48
        }
    }
    switch "hello" {
    case "\n\t":
        v6 = (-(v6))
    case "hello":
        v5 = v5
    case "a":
        fmt.Println(v5, v6)
    default:
        v49 := GenS4{F0: false}
        _ = v49
    }
    if true {
        if true {
            for fi50 := 0; fi50 < 11; fi50++ {
                v1 = append(v1, "foo bar")
                _ = fi50
            }
        } else {
            v5 = v6
            v3[("foo bar" + "foo bar")] = 1
            switch 127 {
            case 4294967295:
                fmt.Println(v4, v6, v5)
            case 128:
                v51 := GenS1{F0: false, F1: 128}
                _ = v51
            case 32768:
                v6 += (v5 >> ((397) & 63))
            default:
                v1 = append(v1, ("\n\t" + ""))
            }
        }
        for fi52 := 0; fi52 < 15; fi52++ {
            var v53 string = "foo bar"
            for fi54 := 0; fi54 < 2; fi54++ {
                switch -746 {
                case 4294967295:
                    var v55 int32 = -904
                    _ = v55
                }
                switch -650 {
                case 32767:
                    if (-2.25 != 0.1) {
                        v56 := GenS5{F0: v4, F1: fi52, F2: (fi52 >> ((65535) & 63))}
                        _ = v56
                    }
                case -2147483648:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v57 uint32 = 42
                    _ = v57
                }
                if false {
                    fmt.Println(v5)
                    switch v53 {
                    case "foo bar":
                        switch 42 {
                        case 4294967295:
                            v5 = (v6 + v5)
                        case 32768:
                            v58 := GenS6{F0: v53}
                            v6 = (-(v6))
                            var v59 int8 = -44
                            _ = v58
                            _ = v59
                        default:
                            v60 := v3[v53]
                            _ = v60
                        }
                    case "":
                        var v61 string = v53
                        _ = v61
                    default:
                        v62 := GenS7{F0: -70, F1: 882}
                        if true {
                            switch fi54 {
                            case 439:
                                v4 += 9223372036854775807
                            }
                        } else {
                            fmt.Println(fi52, fi54, v6)
                            if (!false) {
                                v63 := GenS2{F0: v4}
                                var v64 int = fi54
                                _ = v63
                                _ = v64
                            }
                            var v65 uint32 = 1
                            _ = v65
                        }
                        switch (-(fi54)) {
                        case 810:
                            v66 := GenS8{F0: (-73 < int8(v2)), F1: fi52, F2: 32767}
                            _ = v66
                        case -2147483648:
                            fmt.Println(fi54, fi54, v5)
                        default:
                            var v67 uint = 1
                            _ = v67
                        }
                        if true {
                            if len(v1) > 0 {
                                v1[0] = v53
                            }
                        }
                        v4 += v4
                        _ = v62
                    }
                    var v68 int = (fi54 | 9223372036854775807)
                    switch (v53 + (("" + (v53 + "")) + v53)) {
                    case "a":
                        var v69 uint = 420
                        _ = v69
                    case "hello":
                        if true {
                            delete(v3, "foo bar")
                            v70 := GenS3{F0: -128, F1: 42}
                            _ = v70
                        }
                    }
                    var v71 int8 = 1
                    _ = v68
                    _ = v71
                }
                _ = fi54
            }
            _ = fi52
            _ = v53
        }
    }
    fmt.Println(v5)
    var v72 int16 = (-(v5))
    for fi73 := 0; fi73 < 20; fi73++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi73
    }
    var v74 uint = 42
    _ = len(v1)
    var v75 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v74
    _ = v75
}
