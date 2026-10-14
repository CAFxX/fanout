package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 string }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 uint; F2 bool }
type GenS8 struct { F0 string; F1 int; F2 uint16 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint8 }
func genFn0(p0 bool, p1 string, p2 string) int {
    for fi6 := 0; fi6 < 4; fi6++ {
        if (false || false) {
            v7 := []int{826}
            _ = v7
        }
        fmt.Println(p0, p2, p1)
        p1 = (p1 + (("" + "foo bar") + p2))
        _ = fi6
    }
    var v8 int32 = 166
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(p0, p1)
        v10 := []int{}
        if (!(42 >= 2)) {
            for fi11 := 0; fi11 < 8; fi11++ {
                if ((128 != 2) && true) {
                    fmt.Println(p2)
                    var v12 string = p1
                    _ = v12
                } else {
                    p1 = ("a" + p2)
                    if (-771 != -1) {
                        var v13 float64 = (0.1 / float64(0.1))
                        _ = v13
                    }
                    if len(v10) > 0 {
                        v10[0] = 127
                    }
                }
                if (65535 >= 4294967295) {
                    if true {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v14 int64 = 1
                        _ = v14
                    } else {
                        v8 -= v8
                        if (1 >= -238) {
                            fmt.Println(v8, fi9, p2)
                            p0 = (1 != 644)
                            var v15 uint = 2
                            _ = v15
                        } else {
                            v16 := GenS1{F0: fi11, F1: "foo bar"}
                            _ = v16
                        }
                    }
                } else {
                    if ((0 >= -954) || true) {
                        var v17 bool = (p0 && (0.1 < (0.0 / 0.1)))
                        v18 := GenS2{F0: 123, F1: 42}
                        _ = v17
                        _ = v18
                    }
                    var v19 uint8 = 255
                    _ = v19
                }
                var v20 uint64 = 2
                _ = fi11
                _ = v20
            }
            var v21 int8 = 42
            _ = v21
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}
func genFn1(p0 bool) uint64 {
    switch ("" + "") {
    case "foo bar":
        if ((0 == 86) && p0) {
            var v22 int64 = 517
            var v23 uint32 = 96
            for fi24 := 0; fi24 < 16; fi24++ {
                for fi25 := 0; fi25 < 7; fi25++ {
                    v26 := GenS2{F0: 42, F1: 2}
                    switch fi25 {
                    case 2147483647:
                        v27 := GenS3{F0: 285}
                        _ = v27
                    case 0:
                        v23 = v23
                    default:
                        v28 := GenS4{F0: 700, F1: 1000, F2: "\n\t"}
                        _ = v28
                    }
                    v29 := GenS5{F0: 912, F1: ("hello" + "hello"), F2: (1 == 0)}
                    _ = fi25
                    _ = v26
                    _ = v29
                }
                var v30 int64 = (v22 ^ v22)
                _ = fi24
                _ = v30
            }
            _ = v22
            _ = v23
        } else {
            p0 = p0
        }
    default:
        if (42 == 2) {
            p0 = true
        }
    }
    var v31 uint16 = 1
    if (p0 && (!false)) {
        v32 := []int{9223372036854775807, -2147483648, 128}
        var v33 int64 = 0
        if (2 > 128) {
            v34 := GenS6{F0: 4294967295}
            _ = v34
        }
        _ = v32
        _ = v33
    } else {
        v35 := GenS2{F0: 106, F1: (v31 & v31)}
        v36 := []int{42, 1}
        var v37 uint16 = 32767
        _ = v35
        _ = v36
        _ = v37
    }
    var v38 bool = p0
    _ = p0
    _ = v31
    _ = v38
    return 0
}

func main() {
    var v1 int8 = -1
    v2 := []string{"", "", "a", "\n\t"}
    var v3 float64 = (0.0 + 0.0)
    v4 := make(map[string]bool)
    v5 := []uint32{42, 4294967295, 65535, 2}
    var v39 uint64 = 43
    v39 = v39
    switch 128 {
    case 443:
        _ = len(v4)
    case 233:
        v1 += (v1 - 34)
    }
    fmt.Println(v1, v39, v3)
    v40 := GenS4{F0: (v39 ^ (v39 & v39)), F1: 440, F2: "\n\t"}
    v1 = v1
    var v41 float64 = v3
    var v42 int16 = -273
    var v43 int64 = -888
    if (42 > (v42 - -32768)) {
        v4["hello"] = false
        delete(v4, "foo bar")
    }
    _ = len(v2)
    var v44 float64 = v3
    var v45 int16 = int16(v43)
    if false {
        var v46 uint = 18446744073709551615
        v3 -= 3.14159
        delete(v4, "")
        _ = v46
    } else {
        for fi47 := 0; fi47 < 14; fi47++ {
            switch fi47 {
            case 4294967295:
                v1 -= 0
            default:
                switch (("a" + "hello") + (((("a" + "foo bar") + ((("hello" + (("" + ("a" + "\n\t")) + "")) + "") + "foo bar")) + "\n\t") + ("a" + "\n\t"))) {
                case "a":
                    var v48 int32 = 2147483647
                    _ = v48
                case "\n\t":
                    if false {
                        fmt.Println(v42, v39, v41)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "":
                    v45 = (v45 | (-(v45)))
                }
            }
            _ = fi47
        }
        if (!true) {
            if (!(true || ((-(v43)) < int64(v42)))) {
                for fi49 := 0; fi49 < 3; fi49++ {
                    fmt.Println(v41)
                    v4["a"] = true
                    _ = fi49
                }
                var v50 int8 = -1
                fmt.Println(v39, v1, v44)
                _ = v50
            }
            var v51 string = "foo bar"
            _ = v51
        }
        if true {
            for fi52 := 0; fi52 < 4; fi52++ {
                if true {
                    v53 := GenS6{F0: 65535}
                    v54 := v4[("a" + "foo bar")]
                    _ = v53
                    _ = v54
                } else {
                    var v55 float32 = (float32((0.1 * (1.5 * (v41 * 3.14159)))) / -2.25)
                    v44 += -2.25
                    v41 += v41
                    _ = v55
                }
                var v56 int64 = 1
                _ = fi52
                _ = v56
            }
            var v57 int64 = -1
            var v58 float32 = 1e10
            _ = v57
            _ = v58
        } else {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                if (128 < -105) {
                    var v59 float64 = v3
                    if (((-(v1)) <= v1) || true) {
                        v60 := GenS2{F0: 0, F1: 42}
                        if false {
                            var v61 uint = 18446744073709551615
                            delete(v4, "a")
                            _ = v61
                        } else {
                            v42 = v42
                            for fi62 := 0; fi62 < 7; fi62++ {
                                v41 += ((v44 - 1e10) / v3)
                                _ = fi62
                            }
                        }
                        _ = v60
                    } else {
                        var v63 int16 = v42
                        switch "foo bar" {
                        case "a":
                            delete(v4, ("a" + "hello"))
                        case "":
                            v42 = -32768
                        default:
                            _ = len(v4)
                        }
                        _ = v63
                    }
                    _ = v59
                } else {
                    fmt.Println(v39)
                    var v64 bool = (!(0 <= uint32(v39)))
                    _ = v64
                }
            case "foo bar":
                v65 := GenS7{F0: true, F1: 4294967295, F2: false}
                _ = v65
            }
        }
    }
    switch "" {
    case "a":
        if (!false) {
            var v66 uint8 = 0
            _ = v66
        }
    }
    if false {
        for fi67 := 0; fi67 < 11; fi67++ {
            v68 := GenS3{F0: 42}
            _ = fi67
            _ = v68
        }
    }
    if ((v43 * v43) == 0) {
        var v69 uint32 = 812
        v70 := GenS8{F0: ("a" + ("foo bar" + "\n\t")), F1: 42, F2: 999}
        if true {
            fmt.Println(v3, v3, v41)
            v71 := GenS9{F0: genFn1(false), F1: v43, F2: 118}
            var v72 int8 = v1
            _ = v71
            _ = v72
        } else {
            fmt.Println(v69)
            for fi73 := 0; fi73 < 11; fi73++ {
                if false {
                    for fi74 := 0; fi74 < 18; fi74++ {
                        switch (("a" + "") + "\n\t") {
                        case "":
                            fmt.Println(fi74, v3)
                        default:
                            v1 += v1
                        }
                        fmt.Println(fi74, fi73, v42)
                        _ = fi74
                    }
                    v42 = -184
                    for fi75 := 0; fi75 < 9; fi75++ {
                        v3 = v44
                        _ = fi75
                    }
                }
                _ = fi73
            }
        }
        _ = v69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
