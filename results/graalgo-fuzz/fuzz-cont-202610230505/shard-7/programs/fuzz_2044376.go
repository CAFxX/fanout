package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
type GenS5 struct { F0 string; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint64, p1 string, p2 uint64) uint {
    var v6 uint = 4294967295
    var v7 float64 = (1e10 + (3.14159 + -2.25))
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, (fi8 << ((v6) & 63)), -338}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            p0 += (p2 | p2)
            var v12 int16 = 32767
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int, p1 bool, p2 int16) string {
    for fi13 := 0; fi13 < 11; fi13++ {
        if p1 {
            var v14 float64 = 3.14159
            v15 := []int{p0, fi13, (p0 & (fi13 * -2147483648))}
            _ = v14
            _ = v15
        }
        p0 = fi13
        var v16 uint8 = 0
        _ = fi13
        _ = v16
    }
    v17 := GenS1{F0: 843, F1: ("foo bar" + ("" + ((("\n\t" + "hello") + (("" + "") + ("\n\t" + "foo bar"))) + "\n\t")))}
    v18 := []int{}
    v19 := GenS1{F0: p0, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return (("hello" + "") + "foo bar")
}

func main() {
    var v1 uint64 = 1
    v2 := []int{}
    var v3 uint8 = 127
    v4 := make(map[string]bool)
    var v5 uint64 = (^(v1))
    _ = len(v2)
    v3 += (v3 << ((785) & 63))
    var v20 uint = 65535
    var v21 uint8 = v3
    v20 += 4294967295
    _ = len(v2)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int32 = -649
        _ = fi22
        _ = v23
    }
    if false {
        switch 4294967295 {
        case 363:
            v24 := GenS1{F0: -1, F1: ("foo bar" + "hello")}
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            v21 = 76
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 15; fi27++ {
            v21 = v21
            if true {
                v2 = append(v2, fi26)
                switch fi26 {
                case -206:
                    if ((^(v20)) <= v20) {
                        var v28 int64 = 1
                        _ = v28
                    } else {
                        v4["a"] = false
                    }
                case 809:
                    var v29 int32 = -2147483648
                    _ = v29
                }
                fmt.Println(v1)
            }
            v30 := v4["a"]
            _ = fi27
            _ = v30
        }
        for fi31 := 0; fi31 < 18; fi31++ {
            v1 = v1
            _ = len(v2)
            switch (("a" + "hello") + "\n\t") {
            case "a":
                switch fi26 {
                case -9223372036854775808:
                    if false {
                        var v32 bool = (42 <= -152)
                        _ = len(v4)
                        switch (fi31 / ((fi26) | 1)) {
                        case 9223372036854775807:
                            var v33 uint = (v20 ^ 1)
                            _ = v33
                        case 105:
                            fmt.Println(v20, v32)
                        }
                        var v34 float32 = (0.1 + 3.14159)
                        _ = v32
                        _ = v34
                    }
                    v35 := GenS1{F0: (fi31 | -776), F1: "foo bar"}
                    if (3.14159 != 1e10) {
                        v5 -= (v1 ^ v5)
                        var v36 int = 1
                        var v37 float64 = (1e10 - 3.14159)
                        _ = v36
                        _ = v37
                    }
                    _ = v35
                }
            case "\n\t":
                v38 := v4[(("a" + ("a" + "foo bar")) + "")]
                _ = v38
            case "hello":
                fmt.Println(v21, v20, fi31)
                var v39 int64 = 1
                var v40 int = fi31
                if len(v2) > 0 {
                    v2[0] = fi31
                }
                _ = v39
                _ = v40
            default:
                if true {
                    if false {
                        switch "foo bar" {
                        case "a":
                            delete(v4, ("hello" + "a"))
                        default:
                            v41 := GenS1{F0: fi26, F1: "foo bar"}
                            _ = v41
                        }
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                fmt.Println(v3)
                            case "foo bar":
                                if (v21 <= v3) {
                                    fmt.Println(v21, v5)
                                    if true {
                                        var v42 bool = true
                                        if false {
                                            if ((("\n\t" + "\n\t") + ("" + "hello")) == "hello") {
                                                v43 := GenS1{F0: (fi26 >> ((v20) & 63)), F1: "\n\t"}
                                                var v44 int16 = 42
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 float64 = 1.5
                                            v46 := GenS1{F0: -2147483648, F1: ("" + "")}
                                            _ = len(v4)
                                            v5 -= v1
                                            _ = v45
                                            _ = v46
                                        } else {
                                            _ = len(v2)
                                            if v42 {
                                                v5 = v5
                                                fmt.Println(v5, v5, fi26)
                                                _ = len(v2)
                                                var v47 string = "a"
                                                _ = v47
                                            }
                                        }
                                        fmt.Println(v1, v5, fi31)
                                        switch "a" {
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v48 float32 = 1e10
                                            var v49 uint = (v20 % ((v20) | 1))
                                            _ = v48
                                            _ = v49
                                        case "hello":
                                            if (!v42) {
                                                var v50 uint32 = 456
                                                if (128 > (v3 << (((v20 << ((genFn0((v5 & (v5 >> ((genFn0((^(v1)), "", (v5 % ((v1) | 1)))) & 63))), ("foo bar" + ""), 0)) & 63))) & 63))) {
                                                    var v51 uint8 = v3
                                                    v52 := GenS1{F0: -347, F1: ("" + ("\n\t" + "a"))}
                                                    var v53 float64 = 1e10
                                                    v53 -= (v53 * 1.5)
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                }
                                                _ = v50
                                            }
                                        }
                                        var v54 uint64 = 0
                                        _ = v42
                                        _ = v54
                                    }
                                }
                            default:
                                var v55 uint32 = 868
                                _ = v55
                            }
                        case "hello":
                            if (2 <= 2) {
                                var v56 int = fi31
                                if false {
                                    var v57 uint16 = 0
                                    _ = v57
                                } else {
                                    v20 -= v20
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v21 += v3
                                _ = v56
                            }
                            delete(v4, (("foo bar" + (("\n\t" + ("" + "foo bar")) + "a")) + (("hello" + "a") + (("\n\t" + "") + "foo bar"))))
                            switch (("\n\t" + ((((("" + (("" + "a") + ("a" + "hello"))) + "\n\t") + ("\n\t" + "foo bar")) + ("a" + "\n\t")) + ("" + "a"))) + ("hello" + ("" + "a"))) {
                            case "a":
                                var v58 int = -2147483648
                                _ = v58
                            case "":
                                _ = len(v2)
                            }
                        case "foo bar":
                            v5 = 947
                        }
                    } else {
                        if false {
                            v5 -= (v5 | (v1 - v1))
                            v1 += 222
                        }
                    }
                    v20 -= v20
                    v21 += v3
                } else {
                    if (!true) {
                        if (-2.25 != 1.5) {
                            fmt.Println(fi26, v20, v21)
                            v59 := GenS1{F0: 127, F1: "a"}
                            v60 := GenS2{F0: 821}
                            delete(v4, ("\n\t" + "hello"))
                            v61 := GenS2{F0: 42}
                            var v62 int64 = 42
                            v63 := GenS2{F0: 65535}
                            _ = v59
                            _ = v60
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                    } else {
                        _ = len(v4)
                    }
                    if true {
                        var v64 float32 = (1.5 * 1.5)
                        if (127 >= 1) {
                            var v65 uint64 = v1
                            _ = v65
                        }
                        _ = v64
                    } else {
                        fmt.Println(fi31)
                    }
                    v3 += v3
                    fmt.Println(fi31, v21)
                    v20 = (v20 << ((65535) & 63))
                    var v66 uint32 = 0
                    switch fi31 {
                    case -1:
                        v67 := GenS3{F0: 2147483647, F1: fi26, F2: (!(true || false))}
                        _ = v67
                    case -651:
                        var v68 uint8 = 64
                        _ = v68
                    case 9223372036854775807:
                        v69 := GenS4{F0: 32767, F1: 9223372036854775807, F2: int8(v66)}
                        _ = v69
                    default:
                        if false {
                            if false {
                                v70 := GenS5{F0: "foo bar", F1: v3, F2: (^(v20))}
                                _ = v70
                            } else {
                                fmt.Println(v20)
                                switch "\n\t" {
                                case "a":
                                    var v71 float64 = ((1.5 / 1.5) / 1.5)
                                    _ = v71
                                }
                                if (-2.25 > 1e10) {
                                    switch "\n\t" {
                                    case "":
                                        var v72 float32 = (3.14159 / 1.0)
                                        _ = v72
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = (fi31 << ((uint(v5)) & 63))
                                        }
                                    case "foo bar":
                                        delete(v4, (("\n\t" + ("hello" + (("\n\t" + "") + (("foo bar" + ("" + ("\n\t" + "hello"))) + "")))) + "hello"))
                                    }
                                }
                            }
                            v20 = (v20 << ((v20) & 63))
                            v73 := v4["\n\t"]
                            var v74 string = "foo bar"
                            _ = v73
                            _ = v74
                        }
                    }
                    _ = len(v2)
                    v75 := GenS6{F0: uint8(v20), F1: ("hello" + (("\n\t" + "\n\t") + (("hello" + ("foo bar" + "foo bar")) + "")))}
                    _ = v66
                    _ = v75
                }
            }
            _ = fi31
        }
        _ = fi26
    }
    v3 = v21
    fmt.Println(v21, v3)
    if false {
        v76 := GenS6{F0: 159, F1: "\n\t"}
        v77 := GenS7{F0: -1}
        for fi78 := 0; fi78 < 8; fi78++ {
            delete(v4, (("\n\t" + ((("\n\t" + ("a" + "\n\t")) + "a") + "\n\t")) + (("foo bar" + "a") + ("a" + "foo bar"))))
            _ = fi78
        }
        _ = v76
        _ = v77
    }
    v20 += genFn0(4294967295, "foo bar", 678)
    v4["foo bar"] = (1.5 < -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
