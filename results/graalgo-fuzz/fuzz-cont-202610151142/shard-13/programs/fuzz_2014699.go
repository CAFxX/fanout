package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 bool, p1 int64, p2 bool) int16 {
    p2 = false
    var v7 int64 = p1
    v8 := GenS1{F0: 32767}
    if (904 > 2147483647) {
        if (!(p0 && false)) {
            v9 := GenS2{F0: (p1 + -1), F1: -32768}
            var v10 int16 = 32767
            if p0 {
                for fi11 := 0; fi11 < 13; fi11++ {
                    fmt.Println(v10, p1)
                    _ = fi11
                }
            } else {
                var v12 float64 = 0.0
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            var v13 int32 = 2147483647
            v13 -= v13
            _ = v13
        }
        var v14 uint = 255
        var v15 uint64 = 306
        _ = v14
        _ = v15
    } else {
        fmt.Println(p2, p2, p1)
        if p0 {
            if p0 {
                switch "a" {
                case "foo bar":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v7 = p1
                            var v18 uint32 = 42
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(fi16)
                        var v19 float32 = (float32((-2.25 + float64((0.1 + 3.14159)))) - 1e10)
                        _ = fi16
                        _ = v19
                    }
                case "a":
                    var v20 uint8 = 128
                    _ = v20
                default:
                    var v21 bool = (v7 < 69)
                    _ = v21
                }
                fmt.Println(p2, v7, v7)
                if ("foo bar" < "hello") {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 uint16 = 932
                        v23 = (v23 | v23)
                        v24 := GenS2{F0: -83, F1: 32767}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    var v25 string = ("" + ("" + ("" + "foo bar")))
                    v7 = v7
                    _ = v25
                }
            }
            if p0 {
                switch "foo bar" {
                case "a":
                    p0 = ("foo bar" != "hello")
                case "":
                    fmt.Println(v7)
                case "\n\t":
                    v26 := GenS1{F0: 1}
                    _ = v26
                default:
                    p2 = (58 <= 100)
                }
                switch 0 {
                case 4294967295:
                    switch 289 {
                    case 42:
                        for fi27 := 0; fi27 < 16; fi27++ {
                            if (!false) {
                                var v28 int16 = 0
                                _ = v28
                            }
                            var v29 int8 = 127
                            switch "a" {
                            case "a":
                                if (0 >= -1) {
                                    if p0 {
                                        for fi30 := 0; fi30 < 9; fi30++ {
                                            fmt.Println(p0)
                                            _ = fi30
                                        }
                                        p0 = (!p0)
                                    } else {
                                        v29 -= v29
                                        v29 = (v29 / ((100) | 1))
                                        v29 -= -1
                                    }
                                } else {
                                    for fi31 := 0; fi31 < 10; fi31++ {
                                        var v32 int32 = -269
                                        _ = fi31
                                        _ = v32
                                    }
                                    if false {
                                        p2 = p0
                                        if true {
                                            fmt.Println(p0, v29, v29)
                                        }
                                    }
                                }
                            default:
                                v33 := GenS1{F0: -1}
                                _ = v33
                            }
                            _ = fi27
                            _ = v29
                        }
                    case 682:
                        for fi34 := 0; fi34 < 18; fi34++ {
                            fmt.Println(p2)
                            var v35 int8 = 42
                            var v36 int16 = 42
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                case 2147483647:
                    v37 := []int{-9223372036854775808}
                    _ = v37
                }
            }
            fmt.Println(v7)
        }
        var v38 int8 = -128
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 uint, p1 int64) string {
    v39 := []int{127, 853, int(p0)}
    var v40 int8 = 0
    v41 := GenS2{F0: p1, F1: -1}
    for fi42 := 0; fi42 < 10; fi42++ {
        _ = len(v39)
        _ = fi42
    }
    _ = p0
    _ = p1
    _ = v39
    _ = v40
    _ = v41
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 int16 = 941
    v5 := []string{(((("foo bar" + "a") + "") + (("hello" + ("foo bar" + "foo bar")) + (("foo bar" + ("a" + ("" + "\n\t"))) + (("a" + "foo bar") + ("" + "hello"))))) + "\n\t"), "\n\t", "hello"}
    v6 := []string{"foo bar"}
    if (((!false) || (!false)) && (870 > -279)) {
        v43 := v3["hello"]
        _ = v43
    }
    if ((1 < 2) || false) {
        var v44 float64 = 0.1
        v44 += v44
        v44 += (3.14159 - ((((0.0 - v44) + v44) * v44) - 1.5))
        _ = v44
    }
    v4 -= (^(v4))
    v45 := GenS2{F0: -1, F1: v4}
    for fi46 := 0; fi46 < 15; fi46++ {
        v4 = v4
        _ = fi46
    }
    v4 -= 32767
    var v47 uint8 = 127
    if (-2147483648 > -342) {
        v6 = append(v6, (("a" + "\n\t") + ("" + "hello")))
        var v48 uint64 = 42
        switch 601 {
        case 2147483647:
            if (v4 <= int16(v47)) {
                v47 = v47
            }
        case 128:
            var v49 int32 = 42
            _ = v49
        }
        _ = v48
    }
    v4 += genFn0((v47 >= (v47 ^ 2)), 0, (v4 >= (v4 - (v4 + genFn0(false, 9223372036854775807, true)))))
    _ = len(v5)
    var v50 int32 = 896
    switch 1 {
    case 42:
        var v51 float64 = (((0.1 - ((0.0 - 3.14159) / (1e10 + -2.25))) - 0.1) * 1e10)
        _ = v51
    case 1:
        fmt.Println(v4, v47)
    default:
        fmt.Println(v47, v47)
    }
    switch "foo bar" {
    case "foo bar":
        var v52 int16 = v4
        _ = v52
    case "\n\t":
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
