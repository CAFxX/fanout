package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0(p0 int8) int16 {
    v5 := GenS1{F0: -9223372036854775808, F1: -22, F2: 2147483648}
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(p0)
        p0 -= p0
        _ = fi6
    }
    if (((1e10 / 0.1) == 1e10) || true) {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int = fi7
            p0 = 127
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint64 = 42
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    } else {
        v11 := GenS2{F0: 42, F1: -541, F2: 0}
        v12 := []int{0, 0, 42}
        _ = v11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 12; fi13++ {
            p0 -= 12
            _ = fi13
        }
        if true {
            fmt.Println(p0)
        }
    } else {
        v14 := []int{1, int(p0)}
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 int8, p1 int64) int64 {
    if true {
        v15 := []int{4294967295, 32768, 2147483647, 32768}
        for fi16 := 0; fi16 < 18; fi16++ {
            v15 = append(v15, (fi16 - (fi16 + -1)))
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint8 = 107
                if false {
                    v19 := GenS3{F0: v18, F1: 2}
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        fmt.Println(p0, p1)
        _ = v15
    } else {
        v20 := GenS4{F0: 1000, F1: 255}
        _ = v20
    }
    switch ("foo bar" + ((("hello" + (("foo bar" + ("foo bar" + "hello")) + ((("hello" + "") + ("" + ("" + ""))) + ("\n\t" + ("hello" + "\n\t"))))) + "\n\t") + ("hello" + "hello"))) {
    case "foo bar":
        p1 = -9223372036854775808
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(p0)
        v22 := []int{128, (fi21 >> ((65535) & 63)), fi21, (fi21 / ((fi21) | 1))}
        _ = fi21
        _ = v22
    }
    if true {
        if (779 <= 429) {
            p1 -= (^(p1))
            fmt.Println(p0, p0)
        } else {
            v23 := GenS4{F0: 1, F1: 18446744073709551615}
            v24 := GenS2{F0: 65535, F1: genFn0((p0 / ((p0) | 1))), F2: 0}
            if true {
                switch "foo bar" {
                case "hello":
                    switch "\n\t" {
                    case "":
                        v25 := []int{42, 127, 32767}
                        _ = v25
                    default:
                        var v26 uint8 = 217
                        _ = v26
                    }
                case "foo bar":
                    p1 += (p1 & (p1 % ((0) | 1)))
                case "a":
                    var v27 int64 = p1
                    _ = v27
                default:
                    var v28 uint32 = 2147483648
                    _ = v28
                }
                p0 = (p0 & 0)
            }
            _ = v23
            _ = v24
        }
        if ("\n\t" != "") {
            v29 := []int{486, int(p0)}
            if (-2147483648 < 2147483647) {
                var v30 int = 2147483647
                _ = v30
            }
            for fi31 := 0; fi31 < 10; fi31++ {
                p0 -= 1
                _ = fi31
            }
            _ = v29
        }
    } else {
        p0 = (p0 % ((p0) | 1))
        var v32 int8 = 127
        switch 4294967295 {
        case 42:
            v32 = (p0 - v32)
        case 127:
            v32 -= (v32 + p0)
        default:
            if (false || false) {
                fmt.Println(v32, p1)
                v33 := []int{42, -2147483648, 128, 524}
                var v34 uint8 = 255
                _ = v33
                _ = v34
            } else {
                if true {
                    if (p1 >= 0) {
                        v35 := []int{}
                        if len(v35) > 0 {
                            v35[0] = 128
                        }
                        if (2 != 42) {
                            switch "foo bar" {
                            case "foo bar":
                                p0 += p0
                            case "hello":
                                fmt.Println(p1)
                            case "\n\t":
                                for fi36 := 0; fi36 < 4; fi36++ {
                                    v32 += p0
                                    if (193 != 904) {
                                        for fi37 := 0; fi37 < 3; fi37++ {
                                            v38 := GenS2{F0: 234, F1: 42, F2: 127}
                                            _ = fi37
                                            _ = v38
                                        }
                                    } else {
                                        if ("hello" != ("foo bar" + "foo bar")) {
                                            v39 := GenS5{F0: (fi36 % (((fi36 % (((fi36 << ((65535) & 63))) | 1))) | 1)), F1: 142}
                                            _ = v39
                                        } else {
                                            if (1 <= -1) {
                                                var v40 uint = 264
                                                for fi41 := 0; fi41 < 18; fi41++ {
                                                    fmt.Println(p0, v32)
                                                    v40 += 2
                                                    _ = fi41
                                                }
                                                _ = v40
                                            } else {
                                                p1 += 1
                                            }
                                        }
                                    }
                                    _ = fi36
                                }
                            }
                            var v42 float32 = float32(1.5)
                            fmt.Println(v32)
                            _ = v42
                        }
                        _ = v35
                    }
                    switch "hello" {
                    case "\n\t":
                        p0 -= (v32 << ((979) & 63))
                    case "":
                        var v43 bool = ((!((!((v32 << ((2) & 63)) <= v32)) || (143 > 1))) || (3.14159 == -2.25))
                        _ = v43
                    default:
                        v44 := GenS1{F0: (^(p1)), F1: (p1 << ((665) & 63)), F2: 42}
                        _ = v44
                    }
                    var v45 int = 1
                    _ = v45
                }
                switch ("a" + ("" + "")) {
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        var v46 uint32 = 0
                        _ = v46
                    case "":
                        p1 += (p1 ^ 369)
                    case "hello":
                        if (818 == 164) {
                            v32 += p0
                        }
                    default:
                        v47 := GenS3{F0: 127, F1: 4294967295}
                        _ = v47
                    }
                case "hello":
                    var v48 int64 = p1
                    _ = v48
                default:
                    var v49 int16 = -1
                    _ = v49
                }
            }
        }
        _ = v32
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 128
    var v3 uint8 = v2
    v4 := []string{}
    v2 -= (v2 * 128)
    v3 = v3
    switch (("hello" + (("\n\t" + "\n\t") + "")) + (("\n\t" + ("" + ("foo bar" + ""))) + "")) {
    case "foo bar":
        v4 = append(v4, ((("foo bar" + ("" + "hello")) + "") + "hello"))
    default:
        for fi50 := 0; fi50 < 4; fi50++ {
            var v51 int32 = 1
            delete(v1, (("" + "") + (((("a" + "\n\t") + "a") + "a") + "a")))
            _ = fi50
            _ = v51
        }
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        v2 += v3
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        for fi53 := 0; fi53 < 10; fi53++ {
            v1[("\n\t" + ("\n\t" + "\n\t"))] = 4294967295
            _ = fi53
        }
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v54 := v1[""]
            _ = v54
        }
        _ = fi52
    }
    fmt.Println(v3, v2)
    for fi55 := 0; fi55 < 3; fi55++ {
        v2 -= v2
        v3 += (^(v2))
        _ = fi55
    }
    v2 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
