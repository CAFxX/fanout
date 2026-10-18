package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int64, p1 uint16, p2 int16) uint16 {
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, p0)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 * 713)
}
func genFn1(p0 uint32, p1 uint64, p2 uint8) int64 {
    v8 := GenS1{F0: 132, F1: true, F2: p0}
    if (!(42 > p0)) {
        p1 = (p1 << ((837) & 63))
    } else {
        if (!false) {
            p2 -= (p2 | (p2 << ((428) & 63)))
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint64 = (^(p1))
                if false {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        p1 += p1
                        var v12 string = ("foo bar" + "foo bar")
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi9
                _ = v10
            }
        } else {
            if true {
                v13 := []int{744, 42, -2147483648, -9223372036854775808}
                fmt.Println(p0, p1)
                p2 += (p2 ^ p2)
                _ = v13
            }
            v14 := []int{0, 32768, 286}
            if ((^(p2)) == (p2 << ((1) & 63))) {
                if (("a" + "\n\t") <= (("" + "hello") + (("hello" + (("foo bar" + "foo bar") + "hello")) + "foo bar"))) {
                    p2 -= p2
                } else {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(p1)
                        if (65535 < 335) {
                            for fi16 := 0; fi16 < 1; fi16++ {
                                v14 = append(v14, fi16)
                                v14 = append(v14, (fi15 | fi15))
                                _ = fi16
                            }
                            v14 = append(v14, (fi15 / ((fi15) | 1)))
                        }
                        v14 = append(v14, fi15)
                        _ = fi15
                    }
                }
            }
            _ = v14
        }
        v17 := GenS1{F0: 726, F1: (9223372036854775807 >= -9223372036854775808), F2: 4294967295}
        var v18 uint64 = (p1 ^ (p1 >> ((42) & 63)))
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 float64 = 0.0
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := []string{("" + (("a" + "foo bar") + "a")), ("\n\t" + "a"), "foo bar"}
    switch 42 {
    case -208:
        v19 := v4[("\n\t" + "hello")]
        _ = v19
    case 32768:
        v2 = (v2 << ((65535) & 63))
    case 127:
        v1 = v1
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        switch (("\n\t" + (("a" + ("" + "hello")) + "\n\t")) + "hello") {
        case "\n\t":
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(v2, v2, fi20)
                if true {
                    v3[-67] = fi20
                    v2 = genFn1(4294967295, 42, 181)
                    v22 := GenS2{F0: 2, F1: 299}
                    _ = v22
                } else {
                    var v23 int32 = 222
                    v23 = 1
                    switch fi20 {
                    case -1:
                        v2 = 9223372036854775807
                    case 4294967295:
                        v24 := GenS3{F0: (v23 >> ((1) & 63)), F1: -824, F2: 65535}
                        _ = v24
                    }
                    _ = v23
                }
                switch 32767 {
                case 0:
                    switch ("foo bar" + ("foo bar" + ("foo bar" + "hello"))) {
                    case "hello":
                        if (v2 < genFn1(4294967295, 4294967295, 1)) {
                            fmt.Println(v2, v2, v1)
                            v2 = 0
                            v2 = (v2 >> ((255) & 63))
                        } else {
                            fmt.Println(v2, fi20, fi20)
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                            v2 = (^(v2))
                            fmt.Println(fi21, fi21)
                        }
                    case "":
                        fmt.Println(fi20, v1, v1)
                    }
                case 2147483647:
                    var v25 uint8 = 0
                    _ = v25
                default:
                    v1 = 0.0
                }
                _ = fi21
            }
        default:
            switch "\n\t" {
            case "foo bar":
                var v26 uint16 = 1
                _ = v26
            default:
                v4["foo bar"] = 65535
            }
        }
        switch "hello" {
        case "a":
            var v27 int16 = 0
            _ = v27
        case "foo bar":
            v2 = v2
        case "":
            var v28 float64 = (float64(-2.25) / v1)
            _ = v28
        }
        _ = fi20
    }
    var v29 int8 = 49
    switch ("" + ("foo bar" + "foo bar")) {
    case "":
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    default:
        var v30 int8 = v29
        _ = v30
    }
    if true {
        fmt.Println(v2, v1)
        v1 = -2.25
        delete(v3, -576)
    } else {
        delete(v4, "foo bar")
        for fi31 := 0; fi31 < 12; fi31++ {
            if true {
                v32 := v3[(fi31 * (^(fi31)))]
                v33 := GenS4{F0: 581, F1: 825}
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 4; fi34++ {
                v35 := GenS3{F0: 1, F1: 0, F2: 2147483647}
                fmt.Println(v2, v29, v29)
                switch 941 {
                case 9223372036854775807:
                    if (!(42 > 255)) {
                        fmt.Println(fi34, v29)
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                var v36 int16 = -1
                                var v37 int8 = (^(v29))
                                _ = v36
                                _ = v37
                            }
                        }
                        var v38 uint = 1
                        _ = v38
                    }
                case 4294967295:
                    v39 := GenS1{F0: 255, F1: false, F2: 2147483648}
                    _ = v39
                }
                _ = fi34
                _ = v35
            }
            _ = fi31
        }
        v1 = (((float64(((3.14159 + 1.5) * 1e10)) + (1e10 - 0.0)) - v1) + ((3.14159 / 3.14159) / v1))
    }
    fmt.Println(v1)
    v1 -= v1
    var v40 float64 = -2.25
    v2 = -465
    fmt.Println(v29, v29, v1)
    v41 := GenS4{F0: 65535, F1: 2147483647}
    v40 += 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
