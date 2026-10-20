package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 string }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 int16, p1 uint64) int32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 100}
        fmt.Println(p1, p1, p1)
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: -32768, F1: "foo bar"}
    v8 := []int{}
    for fi9 := 0; fi9 < 17; fi9++ {
        if len(v8) > 0 {
            v8[0] = fi9
        }
        _ = len(v8)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1() bool {
    var v10 uint32 = 815
    v11 := GenS2{F0: 486, F1: "\n\t"}
    switch "" {
    case "\n\t":
        var v12 string = (("a" + "a") + (("foo bar" + "") + "a"))
        _ = v12
    case "a":
        v13 := GenS2{F0: 42, F1: ("foo bar" + ("a" + ("" + "a")))}
        _ = v13
    case "foo bar":
        v10 -= 2147483647
    }
    v10 -= v10
    v14 := GenS2{F0: -975, F1: (("a" + "a") + "hello")}
    _ = v10
    _ = v11
    _ = v14
    return ((0 >= 18446744073709551615) && ((false && (false || (42 >= 374))) || (1 > 1)))
}

func main() {
    var v1 float64 = (0.1 + 3.14159)
    var v2 uint16 = 32767
    v3 := []int{-9223372036854775808, 0, 42, 2147483647}
    v4 := []string{((("a" + "") + "hello") + "a"), "foo bar", "hello", ""}
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(fi15, fi15, v1)
        fmt.Println(v1)
        _ = fi15
    }
    v4 = append(v4, "a")
    var v16 int16 = 1000
    fmt.Println(v16)
    if (0 != 0) {
        if (!(!(42 < 100))) {
            fmt.Println(v16)
            _ = len(v4)
            fmt.Println(v1, v16, v1)
        }
        _ = len(v4)
    } else {
        if false {
            var v17 int32 = -1
            fmt.Println(v1, v16)
            _ = v17
        } else {
            v18 := GenS1{F0: -128}
            var v19 float64 = 0.0
            if true {
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 float32 = 0.0
                    v16 = v16
                    _ = fi20
                    _ = v21
                }
                if ((v16 << ((837) & 63)) == 42) {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v1 += v1
                        _ = fi22
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v3 = append(v3, (fi23 << ((464) & 63)))
                        var v24 int = 2147483647
                        var v25 int8 = 127
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    switch -662 {
                    case 32767:
                        var v26 uint32 = 225
                        _ = v26
                    case 12:
                        v27 := GenS2{F0: 836, F1: ("\n\t" + ("foo bar" + "hello"))}
                        _ = v27
                    default:
                        v2 += v2
                    }
                    if true {
                        if true {
                            fmt.Println(v2, v19)
                            if (-2147483648 != genFn0((v16 - (v16 % (((v16 * (v16 - (v16 | v16)))) | 1))), 2)) {
                                if ((214 > 65535) || (true && (233 > 0))) {
                                    var v28 int8 = 100
                                    var v29 int = int(v28)
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = v29
                                        }
                                        v30 := GenS1{F0: v28}
                                        var v31 int8 = v28
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 float64 = v19
                                        _ = v32
                                    }
                                    _ = v28
                                    _ = v29
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                    fmt.Println(v19, v19)
                                    v33 := GenS1{F0: 127}
                                    _ = v33
                                }
                            } else {
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    switch fi34 {
                                    case 406:
                                        switch ("\n\t" + ("foo bar" + ("hello" + ("\n\t" + "hello")))) {
                                        case "a":
                                            for fi35 := 0; fi35 < 9; fi35++ {
                                                var v36 int8 = 127
                                                fmt.Println(v19, v19, v16)
                                                _ = fi35
                                                _ = v36
                                            }
                                        case "":
                                            v4 = append(v4, "foo bar")
                                        default:
                                            switch ("a" + "\n\t") {
                                            case "\n\t":
                                                v16 -= (-(v16))
                                            case "":
                                                fmt.Println(v16, v16, v16)
                                            default:
                                                v37 := GenS2{F0: 0, F1: "\n\t"}
                                                _ = v37
                                            }
                                        }
                                    case 1:
                                        switch (fi34 << ((1) & 63)) {
                                        case 671:
                                            for fi38 := 0; fi38 < 3; fi38++ {
                                                v39 := GenS3{F0: 0}
                                                fmt.Println(v1, v19)
                                                var v40 bool = (!true)
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                            }
                                        default:
                                            switch ("hello" + ((("" + "\n\t") + "hello") + "hello")) {
                                            case "":
                                                v41 := GenS4{F0: -1, F1: 0, F2: ""}
                                                _ = v41
                                            case "\n\t":
                                                if false {
                                                    if (true || false) {
                                                        for fi42 := 0; fi42 < 3; fi42++ {
                                                            var v43 int64 = 9223372036854775807
                                                            var v44 uint32 = 2
                                                            _ = fi42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        var v45 int = (fi34 ^ (fi34 >> ((2) & 63)))
                                                        _ = v45
                                                    }
                                                    fmt.Println(v19)
                                                } else {
                                                    switch ("foo bar" + "\n\t") {
                                                    case "a":
                                                        var v46 int32 = -2147483648
                                                        _ = v46
                                                    }
                                                }
                                            case "foo bar":
                                                var v47 string = "foo bar"
                                                _ = v47
                                            }
                                        }
                                    }
                                    _ = fi34
                                }
                            }
                            var v48 string = "hello"
                            _ = v48
                        }
                        fmt.Println(v16)
                        fmt.Println(v1, v1)
                    }
                    v19 += float64(3.14159)
                }
            }
            _ = v18
            _ = v19
        }
        if false {
            var v49 uint16 = 2
            v50 := GenS1{F0: 0}
            var v51 uint64 = 0
            _ = v49
            _ = v50
            _ = v51
        } else {
            var v52 uint16 = v2
            if true {
                var v53 uint = 1
                fmt.Println(v16)
                if len(v3) > 0 {
                    v3[0] = -521
                }
                _ = v53
            } else {
                var v54 int16 = (v16 - (v16 - v16))
                v54 = (v54 ^ 337)
                v55 := GenS5{F0: 196, F1: 42}
                _ = v54
                _ = v55
            }
            var v56 uint16 = 752
            _ = v52
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        v1 -= 1e10
        _ = fi57
    }
    if true {
        var v58 bool = true
        var v59 string = ("" + "")
        _ = v58
        _ = v59
    }
    v60 := GenS5{F0: -2147483648, F1: 524}
    var v61 int32 = 2147483647
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v62 uint32 = 786
    _ = len(v3)
    var v63 int32 = -1
    if false {
        var v64 int32 = v61
        if true {
            v4 = append(v4, ((("hello" + "foo bar") + "") + "foo bar"))
            v65 := GenS6{F0: 4294967295, F1: -443, F2: v16}
            _ = v65
        }
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
