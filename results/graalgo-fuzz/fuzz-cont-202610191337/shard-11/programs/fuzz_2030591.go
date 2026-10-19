package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 uint64, p1 uint, p2 int) int {
    if (p1 <= (p1 / ((p1) | 1))) {
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 float64 = 1.5
            _ = fi4
            _ = v5
        }
    } else {
        v6 := GenS1{F0: -1, F1: -24, F2: 1}
        _ = v6
    }
    var v7 int8 = -128
    p0 += 559
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 uint, p1 bool, p2 int) int8 {
    if p1 {
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{(p2 % ((32767) | 1)), (p2 + (p2 ^ (-(p2)))), fi8}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi8
                _ = v9
            }
        default:
            var v10 bool = (590 < 995)
            _ = v10
        }
    } else {
        var v11 int64 = -1
        var v12 int32 = 515
        _ = v11
        _ = v12
    }
    p0 = 893
    switch ("" + (("hello" + "hello") + (("foo bar" + "\n\t") + ("foo bar" + "hello")))) {
    case "\n\t":
        var v13 uint32 = 46
        _ = v13
    case "a":
        v14 := GenS1{F0: (p2 % (((p2 << ((p0) & 63))) | 1)), F1: 100, F2: 32767}
        _ = v14
    case "foo bar":
        if (1e10 > 3.14159) {
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 bool = (!(71 == -1))
                    _ = fi16
                    _ = v17
                }
                p2 = (p2 >> ((p0) & 63))
                _ = fi15
            }
            p2 -= (p2 / ((genFn0(18446744073709551615, p0, (-(p2)))) | 1))
            var v18 int = (p2 % (((-(p2))) | 1))
            _ = v18
        } else {
            v19 := []int{p2, p2, (-(p2)), p2}
            if len(v19) > 0 {
                v19[0] = (^(p2))
            }
            switch "foo bar" {
            case "hello":
                if ((p1 || (((("" + "a") + "") + "foo bar") < "")) || p1) {
                    fmt.Println(p2)
                }
            case "":
                if (p1 || (255 > (p0 ^ 665))) {
                    switch "" {
                    case "hello":
                        var v20 int8 = 100
                        _ = v20
                    case "\n\t":
                        for fi21 := 0; fi21 < 3; fi21++ {
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                if len(v19) > 0 {
                                    v19[0] = p2
                                }
                            default:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    fmt.Println(fi21)
                                    _ = len(v19)
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = fi22
                                }
                            }
                            p0 = 42
                            _ = fi21
                        }
                    case "":
                        if (302 < 32767) {
                            if ("foo bar" <= "hello") {
                                var v23 uint16 = 42
                                _ = v23
                            } else {
                                if true {
                                    fmt.Println(p2, p0, p0)
                                    p1 = p1
                                } else {
                                    v24 := GenS1{F0: genFn0(4294967295, (p0 | p0), 9223372036854775807), F1: 18, F2: -379}
                                    v19 = append(v19, p2)
                                    p2 = p2
                                    _ = v24
                                }
                            }
                            fmt.Println(p1, p1)
                        } else {
                            p1 = (0 > -2147483648)
                        }
                    default:
                        v19 = append(v19, p2)
                    }
                }
            }
            _ = v19
        }
    }
    var v25 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 1
}

func main() {
    var v1 int64 = 148
    var v2 float64 = (3.14159 / (-2.25 * -2.25))
    v3 := make(map[uint32]bool)
    for fi26 := 0; fi26 < 20; fi26++ {
        if (0 > 414) {
            fmt.Println(v1, v1)
            if true {
                for fi27 := 0; fi27 < 2; fi27++ {
                    if true {
                        v3[65535] = false
                    } else {
                        v28 := GenS1{F0: fi27, F1: 1, F2: -1}
                        delete(v3, 1)
                        fmt.Println(fi26, fi26)
                        _ = v28
                    }
                    v2 = ((float64((((0.0 - 3.14159) - 1e10) + 0.1)) * v2) * (-2.25 / (v2 * float64(-2.25))))
                    _ = fi27
                }
            } else {
                fmt.Println(v2, v2)
                for fi29 := 0; fi29 < 17; fi29++ {
                    var v30 uint32 = 678
                    if true {
                        fmt.Println(v30, v30)
                    } else {
                        fmt.Println(v2, v2, v2)
                    }
                    v2 = v2
                    _ = fi29
                    _ = v30
                }
            }
            var v31 uint = 65535
            _ = v31
        }
        var v32 uint16 = 32767
        _ = fi26
        _ = v32
    }
    v33 := []int{-2147483648, -2147483648, 128, 2147483647}
    v1 -= v1
    if (1.5 >= ((1e10 * 0.0) * 1e10)) {
        v1 -= 615
        fmt.Println(v1, v1)
        switch "\n\t" {
        case "a":
            var v34 bool = true
            _ = v34
        default:
            v35 := GenS2{F0: 42}
            _ = v35
        }
    } else {
        var v36 string = ""
        v37 := GenS3{F0: 42, F1: 2, F2: 1}
        _ = len(v33)
        _ = v36
        _ = v37
    }
    fmt.Println(v2, v1)
    for fi38 := 0; fi38 < 2; fi38++ {
        fmt.Println(v2)
        var v39 bool = true
        _ = fi38
        _ = v39
    }
    delete(v3, 387)
    v40 := GenS2{F0: 2147483647}
    delete(v3, 4294967295)
    v1 = v1
    for fi41 := 0; fi41 < 19; fi41++ {
        if (false || (!(((("foo bar" + ("\n\t" + "")) + ("foo bar" + "foo bar")) + "foo bar") != ("a" + "")))) {
            switch "" {
            case "hello":
                if (false || true) {
                    v42 := GenS3{F0: 430, F1: 585, F2: -32768}
                    _ = v42
                }
            case "\n\t":
                if false {
                    var v43 int32 = 1
                    _ = v43
                }
            case "":
                v44 := GenS3{F0: 0, F1: 4294967295, F2: 1000}
                _ = v44
            default:
                v3[1] = true
            }
        } else {
            v1 = v1
            _ = len(v3)
            if (1e10 >= (-2.25 / 1e10)) {
                for fi45 := 0; fi45 < 3; fi45++ {
                    if (32767 == 42) {
                        v46 := GenS4{F0: v1}
                        _ = v46
                    }
                    _ = fi45
                }
                var v47 int64 = (v1 * (v1 & int64(fi41)))
                v47 = v1
                _ = v47
            }
        }
        switch ("\n\t" + ((("foo bar" + "a") + ("hello" + "")) + "")) {
        case "a":
            v48 := GenS2{F0: 42}
            _ = v48
        case "foo bar":
            v49 := GenS5{F0: 127, F1: -1}
            _ = v49
        default:
            for fi50 := 0; fi50 < 15; fi50++ {
                var v51 bool = (!true)
                _ = fi50
                _ = v51
            }
        }
        var v52 string = (("\n\t" + "") + "a")
        _ = fi41
        _ = v52
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        var v54 uint32 = 4294967295
        for fi55 := 0; fi55 < 16; fi55++ {
            if true {
                switch (("" + "hello") + "\n\t") {
                case "foo bar":
                    if ((0.0 / (((3.14159 - 1.5) + 0.1) / -2.25)) <= 0.0) {
                        var v56 int8 = 127
                        _ = v56
                    }
                case "\n\t":
                    if (65 > -2147483648) {
                        var v57 uint32 = v54
                        _ = v57
                    } else {
                        v58 := GenS3{F0: 0, F1: v54, F2: int16(fi55)}
                        v2 -= (float64((-2.25 * -2.25)) / (float64(1.5) / v2))
                        fmt.Println(v2, v1)
                        _ = v58
                    }
                case "":
                    v2 = float64(0.1)
                }
            }
            v1 += (v1 & (v1 / ((v1) | 1)))
            fmt.Println(v54)
            _ = fi55
        }
        for fi59 := 0; fi59 < 14; fi59++ {
            var v60 int8 = -1
            var v61 uint64 = 841
            _ = fi59
            _ = v60
            _ = v61
        }
        _ = fi53
        _ = v54
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 bool = true
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v40
}
