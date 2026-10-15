package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint16; F1 string; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 int16 = -32768
                v7 = v7
                p0 = (0 == 2)
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: 714, F1: (42 <= 128), F2: 32768}
            var v9 string = (("foo bar" + "\n\t") + "hello")
            _ = fi5
            _ = v8
            _ = v9
        }
        if (2 != 768) {
            v10 := GenS1{F0: 255, F1: (931 > 796), F2: 1}
            v11 := GenS1{F0: 31, F1: false, F2: 42}
            _ = v10
            _ = v11
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 uint16, p1 int8) bool {
    fmt.Println(p1, p1)
    fmt.Println(p1, p1)
    v12 := GenS2{F0: 255}
    v13 := GenS3{F0: 32767, F1: p0}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return ((18446744073709551615 == 2) || (p1 <= (p1 - (p1 / ((p1) | 1)))))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("a" + ("foo bar" + (("a" + "foo bar") + (("a" + ("hello" + (("\n\t" + "hello") + ("" + ("a" + ("a" + "\n\t")))))) + ("foo bar" + "foo bar")))))
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v14 uint64 = 0
    if false {
        var v15 uint = 4294967295
        switch 2147483647 {
        case -289:
            v16 := []int{796, 291, 9223372036854775807}
            _ = v16
        }
        var v17 float32 = float32(0.1)
        _ = v15
        _ = v17
    } else {
        var v18 string = "foo bar"
        var v19 uint = 2
        _ = v18
        _ = v19
    }
    v2 = ("foo bar" + v2)
    v14 = (v14 ^ (v14 % ((v14) | 1)))
    switch 868 {
    case 127:
        if true {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS4{F0: 65535, F1: v2, F2: 0}
                if true {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        _ = len(v4)
                        var v23 int = (fi20 + fi20)
                        switch fi22 {
                        case -672:
                            var v24 float64 = (1.5 - 0.1)
                            _ = v24
                        }
                        _ = fi22
                        _ = v23
                    }
                    var v25 float64 = ((float64(-2.25) - (float64(((0.1 / (0.1 - 1.5)) * -2.25)) * 1.5)) + 3.14159)
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := []int{(fi20 / (((fi20 & 1)) | 1)), fi26, fi26, -2147483648}
                        var v28 bool = false
                        _ = len(v4)
                        if v28 {
                            if ((!v28) || (-1 != fi26)) {
                                v14 = (v14 << ((127) & 63))
                                fmt.Println(v14, v14, v2)
                                switch (fi20 + (fi20 | -2147483648)) {
                                case 2147483647:
                                    if v28 {
                                        v29 := GenS5{F0: -47}
                                        v2 = ("foo bar" + "\n\t")
                                        _ = v29
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v2, fi26)
                                    var v30 string = "a"
                                    _ = len(v4)
                                    _ = v30
                                default:
                                    v31 := GenS6{F0: 42, F1: (v2 + v2)}
                                    _ = v31
                                }
                            } else {
                                var v32 float32 = 0.0
                                var v33 int8 = 97
                                v34 := GenS3{F0: -624, F1: 42}
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            v2 = "foo bar"
                        }
                        v2 = v2
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
            v2 = ((((v2 + ("a" + ((v2 + "") + v2))) + v2) + "") + v2)
        }
    case 32767:
        v35 := []int{430, 2147483647}
        _ = v35
    case 336:
        v3["hello"] = true
    }
    v14 = 0
    if (4294967295 >= 2) {
        for fi36 := 0; fi36 < 2; fi36++ {
            if false {
                var v37 string = v2
                v37 = ""
                _ = v37
            }
            v2 = ""
            fmt.Println(v2, v14)
            _ = fi36
        }
        v38 := []int{-550, 32767, -2147483648, -9223372036854775808}
        var v39 int16 = 0
        _ = v38
        _ = v39
    }
    var v40 bool = false
    v41 := []int{-842, 4294967295, 1, -93}
    for fi42 := 0; fi42 < 14; fi42++ {
        if (!((v2 + "") == ((v2 + "a") + "a"))) {
            fmt.Println(v40, v2)
            v43 := GenS7{F0: "\n\t"}
            _ = v43
        } else {
            var v44 float64 = 0.0
            v3[(v2 + v2)] = true
            switch "hello" {
            case "":
                var v45 uint = 0
                _ = v45
            case "hello":
                v44 -= ((v44 + v44) * v44)
            }
            _ = v44
        }
        v2 = v2
        for fi46 := 0; fi46 < 1; fi46++ {
            var v47 uint32 = 2147483647
            _ = fi46
            _ = v47
        }
        _ = fi42
    }
    v40 = (!((847 <= genFn0(((0 > 127) || v40))) || (-1 != 32767)))
    switch 9223372036854775807 {
    case 9223372036854775807:
        v14 -= v14
    case -1:
        _ = len(v41)
    default:
        v40 = (255 >= 255)
    }
    v14 -= 808
    v2 = ((v2 + "\n\t") + (v2 + v2))
    if (("" + (v2 + ("hello" + v2))) != (("foo bar" + v2) + "a")) {
        v40 = (483 > 42)
        if false {
            v14 = (v14 * 18446744073709551615)
        }
        v2 = ("" + "hello")
    } else {
        var v48 bool = false
        v40 = (-1 < 855)
        switch -1 {
        case -997:
            v48 = false
        default:
            var v49 int8 = -23
            _ = v49
        }
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
