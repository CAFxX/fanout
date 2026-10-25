package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int64; F1 string; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 string) int64 {
    v6 := []int{4294967295}
    p0 = ("hello" + ((("a" + p0) + p0) + p0))
    v7 := GenS1{F0: 0, F1: 2}
    var v8 int8 = 42
    for fi9 := 0; fi9 < 5; fi9++ {
        if (-611 == 0) {
            _ = len(v6)
            if true {
                switch "\n\t" {
                case "\n\t":
                    v10 := GenS2{F0: 1, F1: 1}
                    _ = v10
                case "a":
                    fmt.Println(v8, p0)
                case "":
                    fmt.Println(v8)
                default:
                    v11 := GenS3{F0: 420, F1: 100, F2: 0}
                    _ = v11
                }
                v8 += v8
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v8 += v8
                    var v13 float32 = 3.14159
                    _ = fi12
                    _ = v13
                }
                var v14 float32 = (float32(((0.1 + -2.25) * 1e10)) * float32(0.0))
                _ = v14
            }
        }
        _ = fi9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 686
}
func genFn1(p0 bool, p1 uint, p2 string) int16 {
    v15 := GenS4{F0: 4294967295, F1: 889}
    if (!p0) {
        var v16 uint64 = 949
        _ = v16
    }
    p0 = p0
    switch ((((("foo bar" + p2) + p2) + "\n\t") + "\n\t") + "foo bar") {
    case "foo bar":
        fmt.Println(p1, p0, p0)
    case "\n\t":
        var v17 bool = (!(-1 <= 42))
        _ = v17
    case "":
        if false {
            switch "hello" {
            case "":
                p1 = p1
            case "hello":
                switch ("" + p2) {
                case "\n\t":
                    v18 := []int{2147483647}
                    _ = v18
                default:
                    var v19 int16 = -515
                    _ = v19
                }
            case "foo bar":
                if (4294967295 == 2) {
                    v20 := []int{42}
                    _ = v20
                }
            default:
                var v21 float32 = (-2.25 * 0.0)
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := []int{(fi22 ^ 127)}
        for fi24 := 0; fi24 < 13; fi24++ {
            v23 = append(v23, fi22)
            fmt.Println(p1, p2)
            if p0 {
                fmt.Println(fi24, fi22)
                _ = len(v23)
            }
            if (170 >= 1) {
                p1 += (p1 ^ 960)
            } else {
                if ((fi24 << ((29) & 63)) == fi22) {
                    if (!(p0 && p0)) {
                        switch fi24 {
                        case 705:
                            var v25 uint32 = 1
                            _ = v25
                        }
                        fmt.Println(fi22, fi22)
                    }
                    if p0 {
                        var v26 int64 = -1
                        _ = v26
                    }
                    var v27 uint16 = 32767
                    _ = v27
                }
                switch -609 {
                case -647:
                    var v28 uint = p1
                    _ = v28
                case 0:
                    fmt.Println(p2, fi22)
                case 1:
                    fmt.Println(fi22, p2, fi24)
                default:
                    v29 := GenS3{F0: 1, F1: 127, F2: 42}
                    _ = v29
                }
                var v30 int8 = -128
                _ = v30
            }
            _ = fi24
        }
        p2 = (p2 + (p2 + "foo bar"))
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{803, 2147483648, 65535, 1}
    var v2 int = 127
    v3 := []string{("\n\t" + "a"), (("" + (("a" + "\n\t") + "a")) + "hello"), "a"}
    v4 := make(map[string]bool)
    var v5 bool = false
    for fi31 := 0; fi31 < 8; fi31++ {
        if v5 {
            fmt.Println(fi31, fi31, v5)
        }
        fmt.Println(v2)
        _ = fi31
    }
    var v32 uint8 = 255
    switch ("a" + ("hello" + "a")) {
    case "a":
        v33 := v4[("" + ("hello" + ("\n\t" + "hello")))]
        _ = v33
    }
    var v34 int64 = -406
    if true {
        var v35 uint64 = 18446744073709551615
        v32 += (v32 / ((241) | 1))
        _ = v35
    } else {
        var v36 int32 = 1
        var v37 int16 = 499
        v38 := GenS5{F0: (v34 ^ (v34 | 42)), F1: "hello", F2: 154}
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        fmt.Println(fi39, fi39, v5)
        switch v2 {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 128:
            if v5 {
                fmt.Println(v5, v34, v32)
                v40 := v4["hello"]
                var v41 string = ""
                _ = v40
                _ = v41
            } else {
                for fi42 := 0; fi42 < 8; fi42++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch 128 {
                    case -1:
                        fmt.Println(v2, v32, fi39)
                    case -2147483648:
                        if len(v3) > 0 {
                            v3[0] = ("a" + ("" + "hello"))
                        }
                        var v43 bool = true
                        _ = v43
                    default:
                        v32 += v32
                    }
                    v2 = (fi39 << ((4294967295) & 63))
                    switch "a" {
                    case "hello":
                        switch "a" {
                        case "":
                            v1 = append(v1, 62)
                        case "foo bar":
                            fmt.Println(v2)
                        case "a":
                            v5 = (!v5)
                        default:
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        }
                    }
                    var v44 bool = v5
                    fmt.Println(v5, fi42)
                    if (2147483647 == 2147483647) {
                        var v45 uint8 = (v32 % ((255) | 1))
                        if v44 {
                            var v46 int8 = 42
                            v34 -= (v34 / ((v34) | 1))
                            var v47 int8 = 1
                            _ = v46
                            _ = v47
                        } else {
                            switch "a" {
                            case "\n\t":
                                if ((114 >= 1) && (!true)) {
                                    if (!("hello" <= "")) {
                                        fmt.Println(v44)
                                        v2 = 598
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        v32 += (v45 << ((2) & 63))
                                        var v48 uint8 = v45
                                        v4[("hello" + ((("\n\t" + "foo bar") + "hello") + "hello"))] = (v5 || v5)
                                        _ = v48
                                    }
                                    fmt.Println(fi42, v5)
                                }
                            case "foo bar":
                                v45 -= 69
                            case "":
                                v32 -= (^(v32))
                            }
                            v5 = v5
                            v49 := GenS2{F0: int16(v45), F1: v45}
                            _ = v49
                        }
                        v34 = (-(v34))
                        _ = v45
                    } else {
                        v50 := GenS6{F0: (v34 + v34)}
                        v34 = (v34 ^ (v34 % ((v34) | 1)))
                        var v51 string = ("foo bar" + ("hello" + "foo bar"))
                        _ = v50
                        _ = v51
                    }
                    v2 = 141
                    v44 = (false || v44)
                    _ = fi42
                    _ = v44
                }
            }
        case 32768:
            if (true && ((v5 || v5) && v5)) {
                var v52 float32 = (1.5 + 3.14159)
                if (v5 || (4294967295 > 42)) {
                    v53 := GenS6{F0: (v34 + 42)}
                    _ = v53
                } else {
                    v54 := GenS7{F0: 100, F1: 762, F2: 102}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v5 = true
                    _ = v54
                }
                _ = v52
            } else {
                var v55 uint16 = 42
                _ = v55
            }
        default:
            var v56 int16 = genFn1((!(!(!v5))), 1, "")
            _ = v56
        }
        if v5 {
            v4[("\n\t" + "foo bar")] = v5
            fmt.Println(v32, v2)
        } else {
            var v57 int16 = -1
            _ = v57
        }
        _ = fi39
    }
    v2 = 629
    if ((v2 % ((v2) | 1)) == v2) {
        delete(v4, ((("foo bar" + ("foo bar" + (("hello" + "\n\t") + (("\n\t" + "a") + ("\n\t" + "hello"))))) + ("\n\t" + "hello")) + "\n\t"))
        if v5 {
            v58 := GenS7{F0: 24, F1: -32768, F2: 93}
            var v59 int64 = genFn0("a")
            _ = v58
            _ = v59
        } else {
            if (v32 > v32) {
                v34 += genFn0(("" + "foo bar"))
            }
            v32 = 255
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
