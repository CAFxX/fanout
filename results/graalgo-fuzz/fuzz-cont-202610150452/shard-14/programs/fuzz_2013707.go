package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string; F2 int }
type GenS4 struct { F0 uint }
func genFn0(p0 int, p1 int8) uint32 {
    if (!(433 < 1)) {
        p1 = (-(p1))
    } else {
        if false {
            v6 := GenS1{F0: 255, F1: p0, F2: -2147483648}
            fmt.Println(p1)
            _ = v6
        }
        var v7 int32 = 42
        fmt.Println(p0)
        _ = v7
    }
    switch "hello" {
    case "a":
        var v8 uint64 = 829
        _ = v8
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            p1 += 0
            p1 = p1
            _ = fi9
        }
    }
    var v10 uint = 663
    p0 += p0
    _ = p0
    _ = p1
    _ = v10
    return 963
}
func genFn1(p0 bool) bool {
    if (true && (248 != 42)) {
        if ((1 < -128) || (!true)) {
            v11 := []int{239}
            v12 := GenS2{F0: 155}
            _ = v11
            _ = v12
        }
        p0 = (42 <= 127)
    } else {
        switch "foo bar" {
        case "":
            p0 = p0
        default:
            switch "hello" {
            case "foo bar":
                if false {
                    if (!((1 == 255) || p0)) {
                        p0 = (4294967295 != 262)
                    } else {
                        p0 = (1 < 0)
                        if p0 {
                            fmt.Println(p0)
                            for fi13 := 0; fi13 < 5; fi13++ {
                                p0 = (0.1 <= 1.5)
                                fmt.Println(fi13, fi13)
                                _ = fi13
                            }
                            if (1e10 < 1e10) {
                                switch "" {
                                case "":
                                    for fi14 := 0; fi14 < 15; fi14++ {
                                        switch -727 {
                                        case -513:
                                            var v15 uint64 = 609
                                            _ = v15
                                        case -1:
                                            v16 := []int{32768, (fi14 ^ (-(fi14))), fi14}
                                            _ = v16
                                        case 713:
                                            fmt.Println(fi14)
                                        }
                                        _ = fi14
                                    }
                                case "hello":
                                    for fi17 := 0; fi17 < 1; fi17++ {
                                        var v18 int = (-(fi17))
                                        p0 = true
                                        fmt.Println(v18, fi17, p0)
                                        _ = fi17
                                        _ = v18
                                    }
                                case "foo bar":
                                    fmt.Println(p0)
                                default:
                                    v19 := GenS3{F0: -560, F1: ("foo bar" + ("" + "hello")), F2: 2147483647}
                                    _ = v19
                                }
                            }
                        } else {
                            p0 = (false && p0)
                            var v20 float64 = 1.5
                            _ = v20
                        }
                    }
                } else {
                    p0 = (32767 > 42)
                }
            default:
                var v21 bool = (2 != 65535)
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                p0 = false
                v24 := []int{fi22, fi23, 2147483647}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        fmt.Println(p0)
    }
    var v25 uint32 = 42
    p0 = false
    fmt.Println(v25)
    _ = p0
    _ = v25
    return (p0 && ((!true) || (true || (2 < 1))))
}

func main() {
    var v1 uint64 = 147
    var v2 uint32 = 2147483648
    v3 := []uint32{42, 2147483647}
    var v4 uint = 255
    var v5 string = (("a" + "\n\t") + (("hello" + ("" + "a")) + "\n\t"))
    if (true && ((v5 <= "a") || true)) {
        v1 -= (v1 + 435)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v26 := GenS3{F0: 1, F1: v5, F2: 0}
        v4 = v4
        _ = v26
    }
    var v27 float64 = -2.25
    for fi28 := 0; fi28 < 9; fi28++ {
        switch v5 {
        case "\n\t":
            var v29 int = 32767
            _ = v29
        default:
            switch "hello" {
            case "foo bar":
                v2 -= (v2 % ((2147483647) | 1))
            case "a":
                switch fi28 {
                case 32767:
                    var v30 int64 = 1
                    _ = v30
                case -2147483648:
                    v3 = append(v3, (v2 >> (((v4 >> ((v4) & 63))) & 63)))
                }
            case "":
                fmt.Println(v5)
            }
        }
        v31 := GenS1{F0: 128, F1: fi28, F2: fi28}
        _ = fi28
        _ = v31
    }
    var v32 uint8 = uint8(v2)
    v3 = append(v3, 424)
    v4 = (^(v4))
    var v33 int32 = -1
    switch -524 {
    case -2147483648:
        fmt.Println(v4, v33)
    case 42:
        var v34 int = 32767
        _ = v34
    default:
        v27 = (float64(1e10) / (((v27 / float64(((((0.1 + (3.14159 - 1e10)) - 1.5) - 0.1) + 1e10))) + v27) - 3.14159))
    }
    v33 = v33
    var v35 uint = v4
    v5 = ""
    fmt.Println(v27)
    var v36 int16 = -497
    fmt.Println(v36, v32, v2)
    v5 = "hello"
    if len(v3) > 0 {
        v3[0] = v2
    }
    if false {
        var v37 int8 = 22
        _ = v37
    } else {
        v4 = v35
        if true {
            for fi38 := 0; fi38 < 19; fi38++ {
                _ = len(v3)
                v3 = append(v3, (v2 & genFn0(fi38, 0)))
                if (v36 <= v36) {
                    var v39 uint16 = 0
                    v27 -= v27
                    if (uint8(v1) >= v32) {
                        v4 += (^(v35))
                    } else {
                        v40 := GenS4{F0: 18446744073709551615}
                        fmt.Println(v36, v35)
                        v3 = append(v3, (v2 / ((0) | 1)))
                        _ = v40
                    }
                    _ = v39
                }
                _ = fi38
            }
        } else {
            var v41 bool = true
            var v42 float64 = -2.25
            var v43 int = -445
            _ = v41
            _ = v42
            _ = v43
        }
    }
    fmt.Println(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
