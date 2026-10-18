package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint32; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint, p2 int64) int32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch (fi4 >> ((p1) & 63)) {
            case 42:
                switch fi5 {
                case 1:
                    if (false || true) {
                        var v6 int8 = 100
                        _ = v6
                    }
                case 2147483647:
                    var v7 float32 = (3.14159 + 1e10)
                    _ = v7
                case 42:
                    v8 := GenS1{F0: -32768, F1: 42}
                    _ = v8
                default:
                    var v9 int = fi4
                    _ = v9
                }
            case -9223372036854775808:
                v10 := GenS2{F0: 0, F1: 127, F2: ("foo bar" + "hello")}
                _ = v10
            case 32768:
                v11 := []int{(^(fi4)), fi5, fi5, fi5}
                _ = v11
            }
            if true {
                v12 := GenS1{F0: 285, F1: 42}
                _ = v12
            }
            switch "\n\t" {
            case "":
                if (((fi4 ^ fi5) == -1) || (967 == 18446744073709551615)) {
                    if (!true) {
                        v13 := GenS2{F0: -128, F1: 1, F2: "hello"}
                        _ = v13
                    } else {
                        var v14 uint8 = 255
                        p2 -= (-(p0))
                        _ = v14
                    }
                    var v15 string = ("foo bar" + ("\n\t" + "\n\t"))
                    var v16 uint = 349
                    v17 := GenS3{F0: p1, F1: 2147483647, F2: 1}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                p2 = p2
                p1 = 18446744073709551615
            case "hello":
                var v18 int64 = p2
                _ = v18
            }
            if ((((p0 & (p0 ^ p2)) == p2) || (0.0 == 1e10)) || true) {
                if true {
                    fmt.Println(p2, p2)
                    var v19 int16 = -102
                    v20 := []int{(fi5 * (fi4 >> ((65535) & 63))), 9223372036854775807}
                    _ = v19
                    _ = v20
                }
                fmt.Println(p1, fi4, p0)
                if ((false && true) || false) {
                    fmt.Println(p1)
                    var v21 float64 = ((((0.1 - 1.5) * (((0.0 / 1.0) + -2.25) - float64((3.14159 / 0.1)))) * float64(((0.0 + float32(1e10)) - (0.1 * 0.0)))) + 0.0)
                    _ = v21
                } else {
                    p2 -= 42
                }
                v22 := GenS1{F0: -32768, F1: -377}
                var v23 int32 = -1
                _ = v22
                _ = v23
            } else {
                var v24 int8 = 42
                p1 = 18446744073709551615
                var v25 uint = p1
                _ = v24
                _ = v25
            }
            var v26 int = (^(fi4))
            _ = fi5
            _ = v26
        }
        _ = fi4
    }
    p0 += p0
    switch "a" {
    case "a":
        p1 -= p1
    }
    v27 := []int{42, 127, 42, int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 0
}

func main() {
    var v1 bool = (!((-2147483648 > -1) || false))
    var v2 bool = v1
    var v3 uint64 = 1
    v1 = ("\n\t" < ("\n\t" + "\n\t"))
    var v28 bool = false
    v29 := GenS1{F0: 42, F1: -32768}
    switch -217 {
    case 294:
        v30 := []int{127}
        _ = v30
    }
    v2 = (v28 && (0 <= 1))
    v31 := []int{32767}
    for fi32 := 0; fi32 < 18; fi32++ {
        if len(v31) > 0 {
            v31[0] = fi32
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 float32 = 1e10
            if v2 {
                if len(v31) > 0 {
                    v31[0] = (fi33 ^ (-(fi33)))
                }
            }
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    if v1 {
        for fi35 := 0; fi35 < 6; fi35++ {
            v1 = ((3.14159 == 1e10) || v1)
            _ = len(v31)
            v2 = ((!(!v2)) && ((888 < 0) && (false || ((!((("a" + "") + "\n\t") >= "a")) && v2))))
            _ = fi35
        }
        _ = len(v31)
        v36 := GenS4{F0: 4, F1: 564, F2: -2147483648}
        _ = v36
    } else {
        fmt.Println(v28, v2, v3)
    }
    fmt.Println(v3, v2, v3)
    var v37 uint32 = 0
    v1 = v1
    _ = len(v31)
    if ("\n\t" != "a") {
        fmt.Println(v2, v1)
        var v38 uint8 = 255
        v1 = (v3 == (v3 % ((v3) | 1)))
        _ = v38
    } else {
        v37 = 234
        for fi39 := 0; fi39 < 13; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                v41 := GenS1{F0: 1, F1: -32768}
                v3 = (v3 % (((v3 % ((0) | 1))) | 1))
                v1 = v28
                _ = fi40
                _ = v41
            }
            if v1 {
                for fi42 := 0; fi42 < 12; fi42++ {
                    switch ("" + ("a" + ("\n\t" + "hello"))) {
                    case "\n\t":
                        if v2 {
                            fmt.Println(v28, fi39, fi42)
                        } else {
                            switch int(v37) {
                            case 4294967295:
                                switch "\n\t" {
                                case "":
                                    var v43 uint64 = (v3 % (((^(v3))) | 1))
                                    _ = v43
                                case "foo bar":
                                    switch fi42 {
                                    case 128:
                                        var v44 float32 = 1e10
                                        _ = v44
                                    }
                                }
                            default:
                                var v45 string = "a"
                                _ = v45
                            }
                            v46 := GenS5{F0: 2}
                            var v47 int16 = 815
                            _ = v46
                            _ = v47
                        }
                    }
                    v48 := GenS5{F0: 4294967295}
                    var v49 int16 = 859
                    _ = fi42
                    _ = v48
                    _ = v49
                }
            }
            var v50 uint8 = 1
            _ = fi39
            _ = v50
        }
        v1 = v2
    }
    v51 := GenS6{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v37
    _ = v51
}
