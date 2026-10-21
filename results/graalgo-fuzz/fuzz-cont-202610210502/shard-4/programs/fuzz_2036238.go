package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int16 }
func genFn0(p0 int64) int32 {
    switch 1 {
    case -2147483648:
        var v6 uint32 = 150
        _ = v6
    case -1:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{fi7}
            fmt.Println(p0)
            var v9 int64 = (p0 * (p0 / ((p0) | 1)))
            _ = fi7
            _ = v8
            _ = v9
        }
    case -577:
        var v10 int8 = 100
        _ = v10
    default:
        p0 += p0
    }
    var v11 uint16 = 263
    v12 := GenS1{F0: 1}
    v13 := []int{32768, -9223372036854775808}
    if (true || (true && (int16(v11) > int16(p0)))) {
        v13 = append(v13, -1)
        if false {
            fmt.Println(p0)
            v13 = append(v13, 42)
        } else {
            v14 := GenS2{F0: p0}
            _ = v14
        }
    } else {
        var v15 uint16 = (v11 % ((v11) | 1))
        if true {
            v16 := GenS3{F0: 2147483647}
            _ = v16
        } else {
            p0 -= p0
            v11 = v15
        }
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return -2147483648
}
func genFn1() string {
    v17 := []int{42, 9223372036854775807, 340, 32767}
    if (true && true) {
        var v18 int32 = 0
        _ = v18
    }
    _ = v17
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 2147483647
    v3 := make(map[string]string)
    v4 := []string{("hello" + "foo bar"), "a"}
    v5 := []string{"\n\t"}
    for fi19 := 0; fi19 < 20; fi19++ {
        if len(v4) > 0 {
            v4[0] = ((("foo bar" + "") + "hello") + ("hello" + "a"))
        }
        var v20 string = "foo bar"
        fmt.Println(fi19)
        _ = fi19
        _ = v20
    }
    v3[("foo bar" + ("a" + ((((((("\n\t" + "\n\t") + ("\n\t" + (("\n\t" + "") + "a"))) + ("foo bar" + "")) + (("" + ("\n\t" + "")) + "\n\t")) + "hello") + "foo bar") + (("hello" + ("" + "hello")) + ("\n\t" + "hello")))))] = "\n\t"
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 int8 = -119
            fmt.Println(v2)
            var v24 int32 = 426
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    switch -9223372036854775808 {
    case -489:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -2147483648:
        var v25 string = ("foo bar" + (("a" + "hello") + "\n\t"))
        _ = v25
    case 409:
        delete(v1, "\n\t")
    default:
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := GenS2{F0: 9223372036854775807}
            v28 := GenS4{F0: 110, F1: 57}
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch 127 {
                case -1:
                    fmt.Println(fi26)
                }
                if (0.0 <= 0.0) {
                    v2 += (v2 + v2)
                    for fi29 := 0; fi29 < 10; fi29++ {
                        if true {
                            v30 := GenS5{F0: -1, F1: true}
                            _ = v30
                        } else {
                            fmt.Println(fi29, fi26)
                            v31 := GenS5{F0: -53, F1: ((false || (fi29 != fi29)) && ("\n\t" >= ("foo bar" + "foo bar")))}
                            var v32 uint8 = 255
                            _ = v31
                            _ = v32
                        }
                        if (340 < 232) {
                            var v33 int64 = 9223372036854775807
                            var v34 int8 = 91
                            if true {
                                v34 = 0
                                v33 = v33
                                v35 := GenS4{F0: 127, F1: 424}
                                _ = v35
                            }
                            _ = v33
                            _ = v34
                        } else {
                            v2 += 237
                            v2 -= (fi29 ^ (v2 + -171))
                        }
                        var v36 float64 = ((-2.25 * 1e10) - -2.25)
                        v36 = (v36 * v36)
                        v2 = (fi29 / (((fi29 / ((42) | 1))) | 1))
                        _ = fi29
                        _ = v36
                    }
                    v37 := GenS6{F0: 1}
                    _ = v37
                } else {
                    if (38 >= 241) {
                        v2 = (v2 * (-(v2)))
                        v2 += (fi26 + (fi26 >> ((42) & 63)))
                        v2 += fi26
                    }
                    v2 = fi26
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    _ = len(v3)
    if (-28 == 0) {
        fmt.Println(v2)
        delete(v3, "hello")
        _ = len(v3)
    } else {
        for fi38 := 0; fi38 < 5; fi38++ {
            v2 = 128
            v2 += 703
            var v39 int8 = 82
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        v2 = fi40
        v41 := GenS7{F0: -128, F1: 932}
        _ = fi40
        _ = v41
    }
    v2 = (v2 % ((v2) | 1))
    var v42 uint64 = 899
    v43 := GenS2{F0: 0}
    v2 -= 247
    var v44 int8 = -29
    var v45 bool = true
    v45 = v45
    for fi46 := 0; fi46 < 2; fi46++ {
        if false {
            if (65535 < 1) {
                if true {
                    var v47 int32 = -1
                    _ = len(v4)
                    v5 = append(v5, ("\n\t" + "a"))
                    _ = v47
                } else {
                    v2 += 128
                    v48 := GenS8{F0: 1}
                    _ = v48
                }
            } else {
                v2 += fi46
                for fi49 := 0; fi49 < 9; fi49++ {
                    if len(v5) > 0 {
                        v5[0] = ("\n\t" + ("a" + (((("a" + "hello") + "a") + ((("a" + "foo bar") + "\n\t") + ("\n\t" + ("foo bar" + ("foo bar" + ((("foo bar" + "") + (("foo bar" + "hello") + ("" + "hello"))) + "hello")))))) + "")))
                    }
                    _ = fi49
                }
            }
        } else {
            v50 := GenS6{F0: -9223372036854775808}
            var v51 bool = v45
            _ = v50
            _ = v51
        }
        _ = fi46
    }
    v52 := GenS8{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v52
}
