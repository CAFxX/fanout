package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
type GenS6 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint) int64 {
    p0 -= (^(p0))
    var v7 string = ""
    v7 = "a"
    _ = p0
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint32) int64 {
    switch "foo bar" {
    case "":
        p0 = (p0 & (p0 + (p0 * p0)))
    case "hello":
        if (("hello" <= "a") && (p1 <= 65535)) {
            p0 += p0
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(p1, fi8, p1)
                for fi9 := 0; fi9 < 3; fi9++ {
                    p0 = p0
                    var v10 uint8 = 95
                    v10 = v10
                    _ = fi9
                    _ = v10
                }
                p1 -= 604
                _ = fi8
            }
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS1{F0: genFn0(4294967295), F1: fi11}
                _ = fi11
                _ = v12
            }
        }
    case "foo bar":
        v13 := GenS1{F0: -470, F1: 145}
        _ = v13
    default:
        if (int64(p1) >= 0) {
            p1 -= p1
            if false {
                p0 = p0
            } else {
                var v14 float32 = 0.1
                v15 := GenS2{F0: 203, F1: 4294967295, F2: (p1 << ((0) & 63))}
                _ = v14
                _ = v15
            }
        } else {
            var v16 uint16 = 202
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        switch -829 {
        case 32767:
            for fi18 := 0; fi18 < 7; fi18++ {
                switch fi17 {
                case -2147483648:
                    v19 := GenS1{F0: -539, F1: fi17}
                    _ = v19
                }
                _ = fi18
            }
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                p0 = p0
                v21 := []int{(fi20 + (^(fi17))), (^(fi20)), (fi17 << ((0) & 63))}
                p1 = (p1 - (p1 ^ p1))
                _ = fi20
                _ = v21
            }
        }
        fmt.Println(p1, fi17)
        _ = fi17
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS3{F0: 255}
            _ = fi23
            _ = v24
        }
        v25 := GenS4{F0: (("\n\t" + ("hello" + ("\n\t" + "hello"))) + "a")}
        _ = fi22
        _ = v25
    }
    if true {
        if false {
            if ((668 < 557) && ((p1 ^ p1) >= 4294967295)) {
                fmt.Println(p1, p1)
                var v26 float64 = float64((1e10 + float32(3.14159)))
                v26 -= (v26 * v26)
                _ = v26
            } else {
                var v27 bool = (false || (420 > 2147483647))
                v28 := GenS4{F0: "hello"}
                p0 += (^(p0))
                _ = v27
                _ = v28
            }
            if (true && (false && (127 == 42))) {
                for fi29 := 0; fi29 < 19; fi29++ {
                    v30 := GenS4{F0: "a"}
                    if false {
                        fmt.Println(p1, p0)
                        for fi31 := 0; fi31 < 12; fi31++ {
                            var v32 float64 = ((0.1 / 1e10) - 1.5)
                            _ = fi31
                            _ = v32
                        }
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        var v34 bool = false
                        _ = fi33
                        _ = v34
                    }
                    _ = fi29
                    _ = v30
                }
                for fi35 := 0; fi35 < 8; fi35++ {
                    p0 = 42
                    if false {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            var v37 uint32 = p1
                            var v38 bool = (false && (p1 >= p1))
                            p0 = p0
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        fmt.Println(p1)
                        switch 32768 {
                        case 0:
                            var v39 int32 = -1
                            _ = v39
                        case 32767:
                            switch (fi35 * fi35) {
                            case -69:
                                v40 := GenS4{F0: "\n\t"}
                                _ = v40
                            case -1:
                                var v41 uint64 = 4294967295
                                _ = v41
                            case 42:
                                if (0 < 2) {
                                    var v42 int = fi35
                                    v42 = (fi35 | 42)
                                    switch (v42 / ((v42) | 1)) {
                                    case 883:
                                        var v43 float32 = 0.1
                                        _ = v43
                                    }
                                    _ = v42
                                }
                            default:
                                var v44 int16 = 42
                                _ = v44
                            }
                        }
                    }
                    _ = fi35
                }
            }
            v45 := GenS5{F0: 18446744073709551615, F1: 1, F2: 0}
            _ = v45
        } else {
            var v46 uint8 = 127
            v46 += 42
            _ = v46
        }
        var v47 uint = 42
        var v48 int = -9223372036854775808
        _ = v47
        _ = v48
    } else {
        var v49 int8 = 0
        _ = v49
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return genFn0(42)
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = (0.1 / 1.0)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := []string{}
    var v6 uint = 356
    v1 = v1
    v2 = ((v2 + 1.5) * ((v2 - v2) / v2))
    switch "foo bar" {
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        v50 := GenS6{F0: 127, F1: -9223372036854775808, F2: 1}
        _ = v50
    case "\n\t":
        v51 := v4["\n\t"]
        _ = v51
    }
    v52 := GenS7{F0: uint8(v6), F1: 941}
    if (376 < -993) {
        v6 = v6
        fmt.Println(v1, v2, v2)
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v53 uint64 = 0
        var v54 int16 = 20
        _ = v53
        _ = v54
    }
    var v55 float64 = ((v2 - v2) / 3.14159)
    v56 := v4[(("\n\t" + ("" + "foo bar")) + ((("\n\t" + ("a" + "hello")) + "hello") + "\n\t"))]
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v6
}
