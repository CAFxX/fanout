package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16; F1 bool; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int64; F1 int32; F2 int8 }
type GenS9 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 uint32) int {
    fmt.Println(p0)
    switch (("foo bar" + ("hello" + "foo bar")) + "hello") {
    case "a":
        v5 := GenS1{F0: (p0 | p0), F1: (p0 >> ((0) & 63)), F2: 763}
        _ = v5
    default:
        var v6 uint = 4294967295
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            p0 = p0
            p0 -= p0
            fmt.Println(p0, p0)
        }
        _ = fi7
    }
    _ = p0
    return -1
}
func genFn1() uint16 {
    var v8 uint16 = 2
    v9 := []int{genFn0(707), genFn0(2147483647), -1}
    if false {
        v10 := GenS2{F0: -1, F1: 0}
        v8 -= v8
        _ = v10
    } else {
        var v11 uint64 = 0
        _ = v11
    }
    _ = v8
    _ = v9
    return (v8 / (((v8 & v8)) | 1))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18446744073709551615
    v3 := []string{}
    v4 := make(map[int]uint32)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS3{F0: 1000, F1: (("" + "hello") + ((("\n\t" + "\n\t") + "hello") + ""))}
        switch (fi12 + fi12) {
        case 32767:
            for fi14 := 0; fi14 < 3; fi14++ {
                if (255 <= 778) {
                    v2 += v2
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "foo bar" {
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        v2 -= v2
                    case "a":
                        v2 -= v2
                    case "hello":
                        v15 := GenS4{F0: genFn1(), F1: (true || true), F2: -1}
                        _ = v15
                    default:
                        v2 = (v2 % ((v2) | 1))
                    }
                case "a":
                    if (127 != 2) {
                        var v16 float32 = 3.14159
                        _ = v16
                    }
                case "\n\t":
                    v3 = append(v3, "foo bar")
                }
                _ = fi14
            }
        case -1:
            var v17 uint8 = 2
            _ = v17
        case 9223372036854775807:
            v2 = 2
        }
        _ = fi12
        _ = v13
    }
    v3 = append(v3, "")
    fmt.Println(v2)
    var v18 uint16 = 690
    v18 = v18
    if (9223372036854775807 <= 42) {
        if true {
            delete(v1, 1)
        } else {
            v1[2147483647] = -2147483648
            v2 = 4294967295
        }
        var v19 float32 = -2.25
        _ = v19
    } else {
        var v20 float32 = 0.1
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 8; fi22++ {
            if (!(128 >= 0)) {
                v23 := GenS4{F0: (^(v18)), F1: false, F2: -2147483648}
                var v24 int16 = 1
                _ = v23
                _ = v24
            }
            var v25 string = "foo bar"
            _ = fi22
            _ = v25
        }
        _ = fi21
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int64 = 517
        _ = fi26
        _ = v27
    }
    if false {
        if (false || true) {
            v18 += (v18 + 32767)
            _ = len(v3)
        } else {
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 float32 = 1.5
                v30 := GenS1{F0: 2147483648, F1: 65535, F2: 42}
                for fi31 := 0; fi31 < 17; fi31++ {
                    if (-2.25 <= 1.5) {
                        fmt.Println(fi31, v18)
                        var v32 uint = 0
                        _ = v32
                    } else {
                        v18 = v18
                    }
                    fmt.Println(fi31, fi31)
                    _ = len(v3)
                    _ = fi31
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            v33 := GenS4{F0: v18, F1: true, F2: -942}
            if true {
                switch -797 {
                case 0:
                    v1[42] = 1
                }
                v34 := GenS4{F0: (v18 % (((v18 << ((750) & 63))) | 1)), F1: true, F2: -215}
                _ = v34
            }
            _ = v33
        }
        v35 := GenS4{F0: v18, F1: (374 < 18446744073709551615), F2: -451}
        _ = v35
    } else {
        v36 := GenS5{F0: 65535}
        var v37 uint32 = 682
        if (!("foo bar" <= "a")) {
            for fi38 := 0; fi38 < 5; fi38++ {
                if (1 == 127) {
                    var v39 int64 = 9223372036854775807
                    v18 = (v18 & (v18 | v18))
                    var v40 bool = ((v18 != v18) && false)
                    _ = v39
                    _ = v40
                }
                var v41 int = 1
                v42 := GenS3{F0: 42, F1: (((("a" + "a") + "hello") + "a") + ("a" + ("foo bar" + "")))}
                _ = fi38
                _ = v41
                _ = v42
            }
            var v43 int = 127
            if (!(-128 <= 32)) {
                v18 = uint16(v2)
            } else {
                for fi44 := 0; fi44 < 13; fi44++ {
                    var v45 float32 = 1.5
                    for fi46 := 0; fi46 < 19; fi46++ {
                        v37 = (v37 & v37)
                        v2 -= v2
                        v47 := v4[-2147483648]
                        if (127 == -56) {
                            if (!(((v47 <= (^(v47))) || true) || true)) {
                                var v48 bool = true
                                v3 = append(v3, (("a" + "hello") + "a"))
                                _ = v48
                            } else {
                                switch ("" + ("foo bar" + "")) {
                                case "\n\t":
                                    v49 := GenS6{F0: "a"}
                                    _ = v49
                                default:
                                    switch (fi44 - 2147483647) {
                                    case 128:
                                        if true {
                                            var v50 uint32 = v37
                                            v51 := GenS7{F0: -397, F1: 255}
                                            v52 := v1[uint32(v18)]
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        } else {
                                            if (true || false) {
                                                if (fi44 > (fi46 + 2147483647)) {
                                                    switch genFn0(433) {
                                                    case -167:
                                                        v53 := v4[genFn0((v47 & 42))]
                                                        _ = v53
                                                    default:
                                                        fmt.Println(fi44, v43, v2)
                                                    }
                                                }
                                                var v54 uint16 = uint16(v37)
                                                _ = v54
                                            } else {
                                                var v55 int32 = -2147483648
                                                switch "a" {
                                                case "\n\t":
                                                    v43 -= 32768
                                                case "":
                                                    v55 -= (v55 + (v55 % ((-392) | 1)))
                                                case "foo bar":
                                                    v56 := GenS8{F0: -919, F1: (v55 ^ (-(v55))), F2: 127}
                                                    v57 := GenS5{F0: 999}
                                                    _ = v56
                                                    _ = v57
                                                }
                                                _ = v55
                                            }
                                            var v58 float32 = v45
                                            _ = v58
                                        }
                                        switch "\n\t" {
                                        case "":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            v59 := GenS9{F0: 127, F1: 226, F2: (^(v18))}
                                            _ = v59
                                        }
                                    case 713:
                                        if ("foo bar" < "a") {
                                            if (-2147483648 <= -2147483648) {
                                                v45 = ((v45 / 3.14159) + v45)
                                                v47 -= (v47 & (v37 ^ v37))
                                                v1[(v37 / ((v47) | 1))] = (v43 >> ((428) & 63))
                                            }
                                            fmt.Println(v18, v2, fi44)
                                            if false {
                                                v47 = (^(v47))
                                            } else {
                                                v18 -= v18
                                                v43 += genFn0((v37 ^ uint32(v18)))
                                                v18 += v18
                                                v60 := GenS6{F0: "a"}
                                                _ = v60
                                            }
                                        }
                                    case -1:
                                        var v61 uint64 = 747
                                        _ = v61
                                    default:
                                        v62 := GenS3{F0: -32768, F1: "hello"}
                                        _ = v62
                                    }
                                }
                                v45 = v45
                                if true {
                                    v63 := GenS1{F0: (v37 | 390), F1: (v37 * uint32(v2)), F2: 42}
                                    v2 -= (v2 >> ((2) & 63))
                                    _ = v63
                                }
                            }
                            fmt.Println(v47, v37, v47)
                        }
                        var v64 uint32 = (v37 & (^(v37)))
                        _ = fi46
                        _ = v47
                        _ = v64
                    }
                    _ = fi44
                    _ = v45
                }
                v18 = genFn1()
            }
            _ = v43
        } else {
            v65 := GenS3{F0: -32768, F1: ""}
            fmt.Println(v2)
            if true {
                v37 -= (v37 & (v37 + 65535))
                v18 = v18
            }
            _ = v65
        }
        _ = v36
        _ = v37
    }
    var v66 uint16 = (v18 & v18)
    if (128 < 1) {
        v67 := GenS1{F0: 1, F1: 180, F2: 0}
        _ = v67
    }
    switch 128 {
    case 42:
        if true {
            var v68 bool = false
            v68 = (389 == 9223372036854775807)
            _ = v68
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v69 int8 = 1
    v18 += (^(v18))
    fmt.Println(v66, v18, v66)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v69
}
