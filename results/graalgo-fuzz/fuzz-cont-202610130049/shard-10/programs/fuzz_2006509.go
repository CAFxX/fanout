package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
type GenS3 struct { F0 uint32; F1 int8; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int32; F1 int64 }
func genFn0(p0 int8) int {
    var v4 int64 = -1
    v5 := []int{180, 32768}
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "a" {
        case "":
            if true {
                var v7 int8 = p0
                v4 = -488
                v7 = 1
                _ = v7
            }
        case "foo bar":
            v4 -= -9223372036854775808
        default:
            v5 = append(v5, (fi6 << ((440) & 63)))
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            switch (fi6 + fi6) {
            case -1:
                if len(v5) > 0 {
                    v5[0] = fi8
                }
            case 4294967295:
                if true {
                    v9 := GenS1{F0: 127, F1: fi6}
                    _ = v9
                }
            default:
                _ = len(v5)
            }
            var v10 int32 = -2147483648
            if ((1e10 + 0.1) <= 0.1) {
                if true {
                    v11 := GenS2{F0: 396, F1: 1, F2: 2}
                    if false {
                        if ((p0 + p0) <= p0) {
                            if true {
                                v4 += (v4 ^ -9223372036854775808)
                            } else {
                                switch "a" {
                                case "foo bar":
                                    v12 := GenS2{F0: 990, F1: 32767, F2: 4294967295}
                                    _ = v12
                                }
                                v10 = (v10 >> ((65535) & 63))
                            }
                            v10 = (^(v10))
                        } else {
                            v13 := GenS1{F0: 152, F1: (fi8 % ((fi6) | 1))}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v13
                        }
                        switch ("a" + "\n\t") {
                        case "a":
                            fmt.Println(p0)
                        default:
                            v14 := GenS3{F0: 2147483647, F1: 0, F2: 1}
                            _ = v14
                        }
                    } else {
                        var v15 int16 = -32768
                        v10 -= v10
                        v4 -= (v4 * (v4 & -9223372036854775808))
                        _ = v15
                    }
                    if false {
                        fmt.Println(v4, fi6)
                    }
                    _ = v11
                }
            } else {
                fmt.Println(fi6, fi8)
            }
            _ = fi8
            _ = v10
        }
        _ = fi6
    }
    if (-2147483648 != -2147483648) {
        var v16 int64 = (v4 + 242)
        fmt.Println(v4)
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int64 = 0
            if false {
                var v19 int = fi17
                _ = v19
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint = 65535
                    if len(v5) > 0 {
                        v5[0] = fi20
                    }
                    _ = fi20
                    _ = v21
                }
                var v22 uint16 = 775
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        var v23 uint64 = 18446744073709551615
                        _ = v23
                    default:
                        if true {
                            v16 = 265
                            var v24 int64 = v4
                            _ = v24
                        }
                    }
                    _ = len(v5)
                }
                _ = v22
            }
            v25 := GenS4{F0: 4294967295}
            _ = fi17
            _ = v18
            _ = v25
        }
        _ = v16
    } else {
        v26 := GenS4{F0: 42}
        _ = v26
    }
    _ = p0
    _ = v4
    _ = v5
    return 32768
}
func genFn1(p0 uint16) int16 {
    v27 := GenS5{F0: 1, F1: 32767}
    p0 = p0
    var v28 uint16 = (p0 >> ((65535) & 63))
    _ = p0
    _ = v27
    _ = v28
    return 42
}

func main() {
    v1 := []string{(("\n\t" + ((("foo bar" + (((("a" + "") + (("\n\t" + "foo bar") + "a")) + "a") + "foo bar")) + "foo bar") + "hello")) + "\n\t")}
    var v2 bool = false
    var v3 int32 = -98
    v3 = v3
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 uint32 = 2
        _ = fi29
        _ = v30
    }
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "hello")
    }
    var v31 uint32 = 0
    if false {
        var v32 float64 = ((3.14159 / float64(3.14159)) * (-2.25 * -2.25))
        var v33 float64 = 0.0
        _ = v32
        _ = v33
    }
    v34 := GenS3{F0: (v31 ^ 1), F1: 105, F2: 42}
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 uint64 = 1
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            fmt.Println(v3)
            _ = fi36
            _ = v37
        }
        _ = len(v1)
        _ = fi35
    }
    switch 4294967295 {
    case 42:
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 uint = 18446744073709551615
            v39 -= 2
            v40 := GenS4{F0: 42}
            _ = fi38
            _ = v39
            _ = v40
        }
    default:
        var v41 bool = false
        _ = v41
    }
    fmt.Println(v2, v31)
    for fi42 := 0; fi42 < 13; fi42++ {
        fmt.Println(v3, v31, v31)
        if (((1 > 4294967295) || ("a" != "a")) || (!(!((65535 >= 42) || (!("a" != (("" + ("" + "")) + "a"))))))) {
            _ = len(v1)
            v1 = append(v1, "a")
            if v2 {
                if (42 <= 42) {
                    v43 := GenS6{F0: v3, F1: 0}
                    _ = v43
                }
            }
        }
        for fi44 := 0; fi44 < 6; fi44++ {
            var v45 string = ("" + "a")
            switch fi42 {
            case 127:
                v46 := GenS2{F0: 392, F1: 42, F2: 42}
                _ = v46
            }
            v45 = v45
            _ = fi44
            _ = v45
        }
        _ = fi42
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v1 = append(v1, "\n\t")
    v2 = true
    for fi47 := 0; fi47 < 3; fi47++ {
        v48 := GenS2{F0: 0, F1: 2, F2: 1}
        v1 = append(v1, "hello")
        _ = fi47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
}
