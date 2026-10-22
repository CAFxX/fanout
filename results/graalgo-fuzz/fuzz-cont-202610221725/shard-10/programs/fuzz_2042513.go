package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 int, p1 int) string {
    if (false || true) {
        v5 := GenS1{F0: 1, F1: 18446744073709551615, F2: "a"}
        _ = v5
    } else {
        var v6 uint16 = 32767
        if false {
            var v7 int8 = 0
            switch "hello" {
            case "foo bar":
                v8 := []int{(p0 * p0), (p0 - (p0 % ((p1) | 1)))}
                _ = v8
            case "a":
                var v9 float64 = 1e10
                _ = v9
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 uint = 2
                    _ = fi10
                    _ = v11
                }
            }
            var v12 int64 = 9223372036854775807
            _ = v12
            _ = v7
        }
        v13 := GenS1{F0: 42, F1: 794, F2: "\n\t"}
        _ = v13
        _ = v6
    }
    var v14 bool = false
    _ = p0
    _ = p1
    _ = v14
    return "a"
}
func genFn1(p0 string, p1 uint8, p2 int16) uint64 {
    switch 128 {
    case 888:
        p2 = (p2 / ((p2) | 1))
    case 42:
        var v15 uint8 = p1
        _ = v15
    default:
        if true {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 float32 = 3.14159
                _ = fi16
                _ = v17
            }
        } else {
            var v18 uint16 = 696
            switch -22 {
            case 408:
                for fi19 := 0; fi19 < 15; fi19++ {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(fi20)
                        var v21 float64 = (1.5 * 0.1)
                        _ = fi20
                        _ = v21
                    }
                    if (false || true) {
                        for fi22 := 0; fi22 < 4; fi22++ {
                            var v23 bool = (p0 < ((("hello" + ("a" + (p0 + (p0 + "foo bar")))) + "") + "a"))
                            p1 = 74
                            _ = fi22
                            _ = v23
                        }
                        var v24 int8 = 62
                        _ = v24
                    } else {
                        var v25 string = ""
                        v26 := []int{}
                        _ = v25
                        _ = v26
                    }
                    _ = fi19
                }
            case 32767:
                for fi27 := 0; fi27 < 13; fi27++ {
                    switch (fi27 >> ((uint(p1)) & 63)) {
                    case -494:
                        var v28 int = (^(fi27))
                        _ = v28
                    case 209:
                        var v29 int64 = 825
                        _ = v29
                    default:
                        var v30 uint8 = p1
                        _ = v30
                    }
                    p2 = p2
                    p0 = ("\n\t" + p0)
                    _ = fi27
                }
            case -37:
                var v31 uint = 2
                _ = v31
            }
            var v32 uint32 = 42
            _ = v18
            _ = v32
        }
    }
    var v33 int = -783
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return 4294967295
}

func main() {
    v1 := []string{(((("\n\t" + "\n\t") + "a") + "") + "hello")}
    var v2 int = 32767
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    if false {
        fmt.Println(v2)
        for fi34 := 0; fi34 < 2; fi34++ {
            switch "a" {
            case "foo bar":
                for fi35 := 0; fi35 < 6; fi35++ {
                    v2 = fi35
                    _ = fi35
                }
            case "a":
                var v36 uint16 = 1
                _ = v36
            case "hello":
                if (((("hello" + (("foo bar" + ("hello" + (("" + (("a" + "") + ("\n\t" + "foo bar"))) + "foo bar"))) + "foo bar")) + ("\n\t" + ("hello" + ("hello" + "foo bar")))) + "hello") != "hello") {
                    if true {
                        v3[("" + ("\n\t" + "foo bar"))] = ("foo bar" >= "hello")
                        v37 := GenS1{F0: 0, F1: 438, F2: (("a" + ("" + (("" + ("a" + "foo bar")) + "hello"))) + "hello")}
                        v2 = -901
                        _ = v37
                    }
                    fmt.Println(v2, fi34)
                } else {
                    var v38 int16 = -32768
                    switch -9223372036854775808 {
                    case 127:
                        for fi39 := 0; fi39 < 20; fi39++ {
                            var v40 uint16 = 2
                            _ = fi39
                            _ = v40
                        }
                    case 9223372036854775807:
                        v41 := v3[("" + "")]
                        _ = v41
                    case -179:
                        if (0.1 != 1.5) {
                            switch v2 {
                            case 0:
                                fmt.Println(v38, v38)
                            case -1:
                                if true {
                                    if (0 != -2147483648) {
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                        var v42 int = 141
                                        var v43 int = 4294967295
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v2 = fi34
                                        switch "hello" {
                                        case "\n\t":
                                            var v44 uint8 = 169
                                            _ = v44
                                        case "":
                                            switch "foo bar" {
                                            case "foo bar":
                                                v2 += (fi34 << ((18446744073709551615) & 63))
                                            case "":
                                                _ = len(v1)
                                            case "hello":
                                                switch (v2 << ((2) & 63)) {
                                                case -655:
                                                    v45 := v4["hello"]
                                                    _ = v45
                                                case 4294967295:
                                                    v2 = v2
                                                case 42:
                                                    v2 -= v2
                                                default:
                                                    var v46 int32 = 849
                                                    _ = v46
                                                }
                                            default:
                                                var v47 int = fi34
                                                _ = v47
                                            }
                                        }
                                    }
                                }
                            default:
                                v1 = append(v1, "hello")
                            }
                        }
                    }
                    _ = v38
                }
            default:
                if (!(3.14159 == 3.14159)) {
                    var v48 int32 = -2147483648
                    _ = v48
                } else {
                    if true {
                        _ = len(v1)
                        v2 += fi34
                    }
                    var v49 uint32 = 570
                    _ = v49
                }
            }
            v2 -= (-(v2))
            if (true || true) {
                v2 -= v2
                for fi50 := 0; fi50 < 8; fi50++ {
                    if (-11 > 100) {
                        fmt.Println(fi34)
                        _ = len(v1)
                    }
                    if (!false) {
                        if true {
                            var v51 uint32 = 268
                            switch 128 {
                            case 128:
                                switch (fi50 % (((fi34 << ((18446744073709551615) & 63))) | 1)) {
                                case 42:
                                    var v52 bool = ((513 < 1) || true)
                                    _ = v52
                                case 130:
                                    v51 -= 2147483647
                                default:
                                    v51 = v51
                                }
                            case 554:
                                if (736 < 997) {
                                    v53 := GenS1{F0: 255, F1: genFn1("", 42, 699), F2: ("a" + (("foo bar" + "\n\t") + "hello"))}
                                    _ = v53
                                }
                            }
                            if false {
                                var v54 int8 = 35
                                var v55 bool = true
                                _ = v54
                                _ = v55
                            }
                            _ = v51
                        } else {
                            var v56 int8 = -11
                            v57 := GenS1{F0: 45, F1: 29, F2: ""}
                            _ = v56
                            _ = v57
                        }
                    }
                    var v58 float64 = ((-2.25 * 0.0) / 1e10)
                    _ = fi50
                    _ = v58
                }
                v59 := GenS2{F0: -758, F1: 127, F2: 1}
                _ = v59
            }
            _ = fi34
        }
        v60 := GenS2{F0: 0, F1: -1, F2: 42}
        _ = v60
    }
    var v61 uint = 578
    _ = len(v3)
    fmt.Println(v2, v61)
    v1 = append(v1, "foo bar")
    if len(v1) > 0 {
        v1[0] = ((("\n\t" + "foo bar") + "foo bar") + "\n\t")
    }
    v1 = append(v1, (((("hello" + (("hello" + "a") + ("hello" + "foo bar"))) + "hello") + ("" + "")) + ""))
    v2 = -1
    v62 := GenS3{F0: 0, F1: -9223372036854775808}
    v61 += v61
    if (22 > 892) {
        switch (v2 / ((v2) | 1)) {
        case 0:
            fmt.Println(v2, v2)
        case -2147483648:
            switch ("" + "\n\t") {
            case "a":
                v63 := GenS3{F0: 0, F1: 873}
                _ = v63
            }
        default:
            var v64 uint8 = 128
            _ = v64
        }
        v65 := GenS4{F0: false, F1: 2, F2: -1}
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
}
