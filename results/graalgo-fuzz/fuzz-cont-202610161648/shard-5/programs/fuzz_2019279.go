package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 bool; F2 uint }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint16) int8 {
    if (0 != -9223372036854775808) {
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 20; fi8++ {
            if (65535 < 346) {
                var v9 float32 = 1.5
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                if (-1 > -918) {
                    switch "foo bar" {
                    case "hello":
                        p0 = p0
                    case "":
                        if (false || false) {
                            switch ((("foo bar" + "foo bar") + ("" + ("hello" + ""))) + (("\n\t" + (("a" + "foo bar") + "")) + "hello")) {
                            case "":
                                v7 = 884
                            case "hello":
                                if false {
                                    var v11 int32 = 600
                                    switch (fi8 << ((18446744073709551615) & 63)) {
                                    case -1:
                                        fmt.Println(v7)
                                    case -902:
                                        var v12 float64 = ((-2.25 * 0.1) + -2.25)
                                        _ = v12
                                    case -464:
                                        var v13 float32 = 1e10
                                        _ = v13
                                    default:
                                        v11 -= v11
                                    }
                                    _ = v11
                                }
                            case "foo bar":
                                v14 := GenS1{F0: p0, F1: 2}
                                _ = v14
                            default:
                                var v15 int8 = -45
                                _ = v15
                            }
                            v7 -= v7
                            v7 = v7
                            v16 := []int{63, (fi10 | fi10), -903, 128}
                            var v17 int8 = 0
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        if ((139 == 127) || (!true)) {
                            fmt.Println(fi10, v7)
                            var v18 uint = 0
                            _ = v18
                        }
                    default:
                        fmt.Println(fi8)
                    }
                    var v19 uint32 = (v7 % (((v7 >> ((365) & 63))) | 1))
                    fmt.Println(v7, p0)
                    _ = v19
                }
                p0 += p0
                _ = fi10
            }
            _ = fi8
        }
        _ = v7
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{fi20, (fi20 & (fi20 << ((255) & 63))), (fi20 % ((fi20) | 1))}
            p0 -= 255
            fmt.Println(fi20, fi20)
            _ = fi20
            _ = v21
        }
        fmt.Println(p0)
    }
    v22 := GenS2{F0: 0, F1: ("foo bar" + (("hello" + "foo bar") + (("foo bar" + ("" + "")) + "a")))}
    _ = p0
    _ = v22
    return -128
}
func genFn1(p0 uint16) string {
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            p0 -= (p0 >> ((117) & 63))
            _ = fi24
        }
        var v25 string = ("" + ("\n\t" + "foo bar"))
        _ = fi23
        _ = v25
    }
    fmt.Println(p0)
    _ = p0
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -1
    var v3 uint64 = 2
    v4 := make(map[int]bool)
    var v5 float32 = float32(0.1)
    v6 := make(map[string]uint32)
    if (0.1 != 0.1) {
        var v26 float32 = 3.14159
        _ = v26
    } else {
        switch "a" {
        case "a":
            var v27 int = -651
            _ = v27
        case "\n\t":
            var v28 float64 = 1.5
            _ = v28
        case "hello":
            fmt.Println(v5, v5, v5)
        }
        switch 4294967295 {
        case -9223372036854775808:
            if false {
                v29 := GenS3{F0: -136, F1: -420}
                fmt.Println(v3)
                switch "foo bar" {
                case "a":
                    fmt.Println(v5)
                case "hello":
                    v30 := GenS1{F0: 65535, F1: uint32(v3)}
                    _ = v30
                }
                _ = v29
            } else {
                for fi31 := 0; fi31 < 20; fi31++ {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v33 := GenS4{F0: v3}
                        var v34 int16 = 731
                        v3 += v3
                        if false {
                            v35 := GenS4{F0: 456}
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 int8 = 1
                        if true {
                            delete(v1, fi36)
                        }
                        v37 -= (v37 | genFn0(407))
                        switch 688 {
                        case 32768:
                            v2 -= (v2 & -1)
                            delete(v6, (((("a" + "foo bar") + "") + ("\n\t" + "")) + "\n\t"))
                        default:
                            if ((!true) || false) {
                                v1[(fi36 ^ (^(fi31)))] = (!(0.1 <= v5))
                                v38 := GenS4{F0: v3}
                                _ = v38
                            } else {
                                v39 := GenS5{F0: -906, F1: true, F2: 4294967295}
                                v4[(fi31 / ((fi31) | 1))] = (65535 > 65535)
                                _ = v39
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                    _ = fi31
                }
            }
        }
        var v40 int64 = 42
        _ = v40
    }
    fmt.Println(v2)
    var v41 int8 = 64
    var v42 int32 = 42
    if (true && (114 < 128)) {
        fmt.Println(v41, v2)
        v43 := []int{32768, 42, int(v41), 635}
        _ = v43
    } else {
        v5 -= float32((0.0 / 0.1))
    }
    var v44 uint = 1
    var v45 float64 = 1.5
    for fi46 := 0; fi46 < 13; fi46++ {
        v41 += (^(v41))
        _ = fi46
    }
    fmt.Println(v5)
    var v47 string = "a"
    v48 := GenS6{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
